package main

import (
	"fmt"
	"path/filepath"
)

var modFolder = "c:\\Users\\admin\\Documents\\Paradox Interactive\\Hearts of Iron IV\\mod"
var filePath = "oldworldblues\\map\\cities.bmp"

func main() {
	img, err := Identify(filepath.Join(modFolder, filePath))
	if err != nil {
		panic(err)
	}
	fmt.Println(img)
}
