package main

import (
    "strings"
    "image"
    "image/jpeg"
    "image/png"
    "log"
    "net/http"
    "path"
    "os"

    "github.com/codegangsta/martini"
    "github.com/nfnt/resize"
)

func preThumb() image.Image {
	imageName := "200.jpg"
	imageFile, err := os.Open(imageName)
	if err != nil {
		log.Fatal(err)
	}
	img, err := jpeg.Decode(imageFile)
	if err != nil {
		log.Fatal(err)
	}
	return img
}

func thumb() image.Image {
    imageName := "test.jpg"
    imageFile, err := os.Open(imageName)
    if err != nil {
        log.Fatal(err)
    }
    
    var myImage *image.Image

    switch strings.ToLower(path.Ext(imageName)) {
		case ".jpg", ".jpeg":
		img, err := jpeg.Decode(imageFile)
		if err != nil {
			log.Fatal(err)
		}
		myImage = &img
		case ".png":
		img, err := png.Decode(imageFile)
		if err != nil {
			log.Fatal(err)
		}
		myImage = &img
    }
	imageFile.Close()

    m := resize.Resize(0, 200, *myImage, resize.MitchellNetravali)

    return m
}

func main() {
    m := martini.Classic()

    m.Get("/", func(res http.ResponseWriter, req *http.Request) {
        res.Header().Set("Content-Type", "image/jpeg")
        err := jpeg.Encode(res, thumb(), &jpeg.Options{75})
        if err != nil {
            res.WriteHeader(500)
        } else {
            res.WriteHeader(200)
        }
    })

	m.Get("/cached", func(response http.ResponseWriter, req *http.Request) {
			response.Header().Set("Content-Type", "image/jpeg")
			err := jpeg.Encode(response, preThumb(), &jpeg.Options{75})
			if err != nil {
				response.WriteHeader(500)
			} else {
				response.WriteHeader(200)
			}
		})


	log.Fatal(http.ListenAndServe(":10010", m))
    m.Run()
}
