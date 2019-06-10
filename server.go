// Stephan ACQUATELLA 11/2017
package main

import (
	"github.com/rs/zerolog"
	"html/template"
	"os"
	"path/filepath"
	//	_ "slides-remark-server/statik"
	_ "github.com/sacquatella/slides-remark-server/statik"

	"flag"
	"github.com/gin-contrib/logger"
	"github.com/gin-contrib/static"
	"strings"
	//	"github.com/gin-gonic/contrib/ginrus"
	"github.com/gin-gonic/gin"
	"github.com/rakyll/statik/fs"
	"github.com/rs/zerolog/log"
	"net/http"
)

var prefix = ""
var slidepath = "."
var themepath = "BINDATA"
var theme = "remark" // remark theme also available
var port = "3000"    // default http port
var title = "SlideAsCode Server (SACS)"
var useGB = true
var ratio = "16:9"
var ratiolist = []string{"4:3", "16:9"}

var showhelp *bool
var help string = `Environment variables :
	- SLIDE_PREFIX : markdown file prefix. Only markdown files that match defined prefix will be displayed the "menu slide".
	- PORT: http port, default port is 3000.
	- SLIDE_TITLE: redefine your title in slide menu.
	- SLIDES_PATH : markdown file location.
	- THEME : slides theme to use. Default theme is remark.
    - RATIO : slide ratio, available value are "4:3" and "16:9". Default value is "16:9"
	`

// check if a string exist in a slice.
func stringInSlice(str string, list []string) bool {
	for _, v := range list {
		if v == str {
			return true
		}
	}
	return false
}

func init() {
	showhelp = flag.Bool("help", false, help)
}

// get env variables and initialise logs
func initialise() {

	log.Logger = log.Output(
		zerolog.ConsoleWriter{
			Out:     os.Stderr,
			NoColor: false,
		},
	)
	if gin.IsDebugging() {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}
	if os.Getenv("SLIDES_PATH") != "" {
		slidepath = os.Getenv("SLIDES_PATH")
	}

	if os.Getenv("SLIDE_PREFIX") != "" {
		prefix = os.Getenv("SLIDE_PREFIX")
	}
	if os.Getenv("THEME") != "" {
		theme = os.Getenv("THEME")
	}
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}

	if os.Getenv("THEME_PATH") != "" {
		themepath = os.Getenv("THEME_PATH")
	}

	if os.Getenv("USEGB") == "FALSE" {
		useGB = false
	}

	if os.Getenv("SLIDE_TITLE") != "" {
		title = os.Getenv("SLIDE_TITLE")
	}

	if os.Getenv("RATIO") != "" {
		ratio = os.Getenv("RATIO")
		if !stringInSlice(ratio, ratiolist) {
			ratio = "4:3"
			log.Info().Msgf("Provided ratio %s not supported .. ", ratio)
		}
	}
}

// build home.md slide
func buildHome() {
	f, err := os.Create(slidepath + "/home.md")
	if err != nil {
		log.Fatal().Msg(err.Error())
	}
	defer f.Close()
	_, err = f.WriteString("# " + title + "\n \n")

	files, err := filepath.Glob(slidepath + "/" + prefix + "*.md")
	if err != nil {
		log.Fatal().Msg(err.Error())
	}

	for _, slide := range files {

		slidepath = strings.TrimPrefix(slidepath, "./")

		replaceStr := slidepath + string(os.PathSeparator)
		cslide := strings.Replace(slide, replaceStr, "", -1)
		csslide := strings.Replace(cslide, ".md", "", -1)
		mdline := "- [" + csslide + "](?slides=" + cslide + ") \n"
		_, err := f.WriteString(mdline)
		if err != nil {
			log.Fatal().Msg(err.Error())
		}
	}
	f.Sync()
	return
}

// redirect / to /index
func IndexMiddleware(c *gin.Context) {

	//log.Debugf("Path is :%s", c.Request.URL.Path)

	if c.Request.URL.Path == "/" {
		c.Redirect(301, "/index")
	}

	c.Next()
}

// main
func main() {

	initialise()
	flag.Parse()
	buildHome()
	// use statik to serve js and themes css
	statikFS, err := fs.New()
	if err != nil {
		log.Fatal().Msg(err.Error())
	}

	router := gin.New()

	if gin.Mode() == "release" {
		router.Use(logger.SetLogger())
	} else {
		router.Use(gin.Logger())
	}

	// define template
	// Use binddata for html template
	var t *template.Template
	if useGB {
		b, _ := Asset("templates/index.tmpl")
		t, _ = template.New("index.tmpl").Parse(string(b))
		router.SetHTMLTemplate(t)
		log.Debug().Msg("Use embeded index.tmpl template, go-bindata")
	} else {
		router.LoadHTMLFiles("templates/index.tmpl")
		log.Debug().Msg("Use local index.tmpl template localized in templates folder")
	}

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// Use middleware to redirect / to /index
	router.Use(IndexMiddleware)

	router.StaticFS("/public/", statikFS)
	router.Use(static.Serve("/", static.LocalFile(slidepath+"/", false)))

	router.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"Theme": theme,
			"Ratio": ratio,
		})
	})

	router.Run(":" + port) // listen and serve on PORT
}
