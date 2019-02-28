package main


import (
//	"fmt"
//	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func init() {
	gin.SetMode(gin.TestMode)
}

type testStruct struct {
	T *testing.T
}

func TestStringInSlice(t *testing.T) {
	ratio := "4:3"
	ratiolist := []string{"4:3", "16:9"}
	assert.True(t, stringInSlice(ratio, ratiolist))
}