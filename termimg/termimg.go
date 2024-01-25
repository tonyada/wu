package termimg

import (
	"encoding/base64"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

var (
	width  = flag.String("width", "", "width (e.g. 100px, 10%, or auto)")
	height = flag.String("height", "", "height (e.g. 100px, 10%, or auto)")
	size   = flag.String("size", "", "width,height in pixels (e.g. 1024px,768px or 3,3)")
)

// Display image to iterm2
func Display(filename string) error {
	width, height := widthAndHeight()
	return DisplayPro(filename, width, height)
}

// Display image with width and height to iterm2
func DisplayPro(filename, width, height string) error {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	// b64filename := base64.StdEncoding.EncodeToString([]byte(filename))
	// width, height = widthAndHeight()
	if width == "" && height == "" {
		width = "auto"
		height = "auto"
	}

	fmt.Print("\033]1337;")
	// fmt.Printf("File=inline=1;preserveAspectRatio=1;name='%s'", b64filename)
	fmt.Printf("File=inline=1")
	if width != "" || height != "" {
		if width != "" {
			fmt.Printf(";width=%s", width)
		}
		if height != "" {
			fmt.Printf(";height=%s", height)
		}
	}
	// fmt.Print("preserveAspectRatio=1")
	fmt.Print(":")
	fmt.Printf("%s", base64.StdEncoding.EncodeToString(data))
	fmt.Print("\a\n")
	// fmt.Printf("\033[A%s\n", filename)

	return nil
}

func widthAndHeight() (w, h string) {
	if *width != "" {
		w = *width
	}
	if *height != "" {
		h = *height
	}
	if *size != "" {
		sp := strings.SplitN(*size, ",", -1)
		if len(sp) == 2 {
			w = sp[0]
			h = sp[1]
		}
	}
	return
}
