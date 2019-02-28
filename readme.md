# Basic Slides Server


It's a simple, one executable server, for markdown-driven slideshow tools name [remark](https://github.com/gnab/remark).

You just have to create your slides as markdown files in your current folder , and start slides-server. 
You slides will be available through http://localhost:3000/.

slides-server create automatically a "home-page" slide, with the list of slides that he found in the current (or provided) folder.

## Build

```console
go get github.com/sacquatella/slides-remark-server
```

## Usage 

Download slides-server binary for your platform or build it.

Create your slides as a markdown files . You can use [remark wiki](https://github.com/gnab/remark/wiki) to help you.
Start slides-server binary from the folder that content your slides, 
```bash
$ slides-server
```

Your slides are available at http://localhost:3000.


## Customizing run

 A set of environment variables can be set to customize rendering.

- SLIDE_PREFIX : markdown file prefix. Only markdown files that match defined prefix will be displayed the "menu slide".
> linux/OSX
```bash
 SLIDE_PREFIX="Slide_" slides-server
```
> Windows, PowerShell
```bash
 $env:SLIDE_PREFIX="Slide_" slides-server
```
- PORT : http port. default port is 3000
- SLIDES_PATH : markdown files location (your slides).
- SLIDE_TITLE : You can redefine the title of the "menu slide".
- THEME : Theme to use. The default theme is `remark` theme. Embedded themes are located in public/js/themes folder. 
> Sample : 
```bash
THEME=remark slides-server
```
- RATIO : Ratio for your slides. Available values ares  "4:3" and "16:9". Default value is "16:9".
- THEME_PATH : If  THEME_PATH is defined slides-server will not use embedded theme, the theme  will be searched in provided folder.

## Run it as docker container

```bash
docker run -p 3000:3000 -v $(pwd):/slides -e SLIDES_PATH=/slides onlans/slides-remark-server:0.5
```


## Development

See [CONTRIBUTING.md](CONTRIBUTING.md)

