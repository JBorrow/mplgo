mplgo
=====

A small package for using `matplotlib` colour maps in `golang`.

Works by shelling out to `python` and extracting a color map directly from matplotlib,
to a `golang` struct.

Example Usage
-------------

Download the package:

```
go get github.com/jborrow/mplgo
```

```go
package main

import (
	"image/png"
	"log"
	"os"
	"github.com/jborrow/mplgo"
)

func MapArrayToPNG(m mplgo.ColorMap, in [][]float64, file_name string) error {
	f, err := os.Create(file_name)
	if err != nil {
		return err
	}
	defer f.Close()

	// Encode the image to PNG and save it to the file
	if err := png.Encode(f, m.MapArrayToImage(in)); err != nil {
		panic(err)
	}

	return nil
}

func main() {
	colorMap, err := mplgo.GetCmap("viridis", 128)

	if err != nil {
		log.Fatal(err)
	}

	// Example data
	data := make([][]float64, 128)

	for i := range data {
		line := make([]float64, 128)
		for j := range line {
			line[j] = (float64(i) / 128.0 * float64(j) / 128.0)
		}
		data[i] = line
	}

	// Write to the world
	err = MapArrayToPNG(colorMap, data, "hello_world.png")

	if err != nil {
		log.Fatal(err)
	}

	return
}
```