package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io/ioutil"
	"log"
	"net/http/httputil"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/unrolled/secure"
)

const EmulatedBand = 100000 // 100KBps

func DebugHeaders(ctx *gin.Context) {
	fmt.Println(ctx.Request.Host, ctx.Request.RemoteAddr, ctx.Request.RequestURI)

	// Save a copy of this request for debugging.
	requestDump, err := httputil.DumpRequest(ctx.Request, false)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(requestDump))
	clS := ctx.Request.Header.Get("Content-Length")

	contentLength, err := strconv.Atoi(clS)
	if err == nil {
		sleepMS := 200.0 + 100.0*contentLength/EmulatedBand
		log.Printf("Sleeping %d ms for %d bytes\n", sleepMS, contentLength)
		time.Sleep(time.Duration(sleepMS) * time.Millisecond)
	}

	ctx.Next()
}

func PrintContent(c *gin.Context) {
	jsonData, err := ioutil.ReadAll(c.Request.Body)
	encoding := "normal"
	if err != nil {
		log.Printf("ERROR %s", err)
		c.Next()
		return
	}
	if strings.Contains(c.Request.Header.Get("Content-Encoding"), "gzip") {
		gr, _ := gzip.NewReader(bytes.NewReader(jsonData))
		defer gr.Close()
		jsonData, err = ioutil.ReadAll(gr)
		if err != nil {
			log.Printf("ERROR %s", err)
		}
		encoding = "gzip"
	}
	log.Printf("Received [%s] %s\n", encoding, string(jsonData))
	c.Next()
}

func main() {
	r := gin.Default()

	r.Use(DebugHeaders)
	r.Use(PrintContent)

	//r.Use(TlsHandler())
	r.POST("/*request", func(c *gin.Context) {
		// jsonData, err := ioutil.ReadAll(c.Request.Body)
		// encoding := "normal"
		// if err == nil {

		// 	if strings.Contains(c.Request.Header.Get("Content-Encoding"), "gzip") {
		// 		gr, _ := gzip.NewReader(bytes.NewReader(jsonData))
		// 		defer gr.Close()
		// 		jsonData, err = ioutil.ReadAll(gr)
		// 		encoding = "gzip"
		// 	}
		// 	log.Printf("Received [%s] %s\n", encoding, string(jsonData))
		// }
		c.JSON(200, gin.H{"POST": "OK"})
	})
	r.PUT("/*request", func(c *gin.Context) {
		// jsonData, err := ioutil.ReadAll(c.Request.Body)
		// if err == nil {
		// 	log.Println("Received", string(jsonData))
		// }
		c.JSON(200, gin.H{"PUT": "OK"})
	})

	r.Run()
	// r.RunTLS(":8080", "cert.pem", "key.pem") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func TlsHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		secureMiddleware := secure.New(secure.Options{
			SSLRedirect: true,
			SSLHost:     "localhost:8080",
		})
		err := secureMiddleware.Process(c.Writer, c.Request)

		// If there was an error, do not continue.
		if err != nil {
			return
		}

		c.Next()
	}
}
