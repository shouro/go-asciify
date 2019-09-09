package main

import (
	"log"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/disintegration/imaging"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/memstore"
	"github.com/gin-gonic/gin"
	"github.com/shouro/go-asciify/asciify"
)

func randomString() string {
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, 8)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

func makeASCII(imgPath string) string {
	srcImg, err := imaging.Open(imgPath, imaging.AutoOrientation(true))
	if err != nil {
		log.Fatalf("failed to open image: %v", err)
		return ""
	}
	//width := srcImg.Bounds().Max.X
	//height := srcImg.Bounds().Max.Y
	//width := 80
	//height := 0
	var dark uint8
	//resizedImg := imaging.Resize(srcImg, width, height, imaging.Lanczos)
	ascii := asciify.ToASCII(srcImg, dark)
	out := ""
	for _, row := range ascii {
		out = out + row + "\n"
	}
	return out
}

func main() {
	//defer profile.Start().Stop()
	r := gin.Default()
	// Set a lower memory limit for multipart forms (default is 32 MiB)
	r.MaxMultipartMemory = 8 << 20 // 8 MiB
	r.LoadHTMLGlob("static/templates/*")
	store := memstore.NewStore([]byte("supersecret")) //TODO: set proper secret
	r.Use(sessions.Sessions("asciifysession", store))

	r.GET("/", func(c *gin.Context) {
		session := sessions.Default(c)
		asciiout := session.Get("asciiout")
		session.Set("asciiout", nil)
		flash := session.Get("flash")
		session.Set("flash", nil)
		session.Save()
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"asciiout": asciiout,
			"flash":    flash,
		})
	})

	r.POST("/goascii", func(c *gin.Context) {
		session := sessions.Default(c)
		file, err := c.FormFile("imagefile")
		if err != nil {
			//c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
			session.Set("flash", "Please select a file.")
			session.Save()
			c.Redirect(http.StatusMovedPermanently, "/")
			return
		}
		file.Filename = randomString()
		filename := filepath.Base(file.Filename)
		imgPath := "/tmp/" + filename + ".image"
		if err := c.SaveUploadedFile(file, imgPath); err != nil {
			//c.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
			log.Printf("failed to save file in server: %s\n", err.Error())
			session.Set("flash", "Something went wrong!!, failed to save file in server.")
			session.Save()
			c.Redirect(http.StatusMovedPermanently, "/")
			return
		}
		asciiout := makeASCII(imgPath)
		session.Set("asciiout", asciiout)
		session.Save()
		err = os.Remove(imgPath)
		if err != nil {
			log.Printf("failed to remove file: %s\n", err.Error())
		}
		c.Redirect(http.StatusMovedPermanently, "/")
	})
	r.Run()
}
