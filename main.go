package main

import (
	"fmt"
	"github/goWithTesting/mockingEx"
)

func main() {
	data := mockingEx.Measurements{
		Height: 6,
		Length: 3,
		Width:  2,
	}
	volume, area := mockingEx.GetVolumeandSurfaceCube(data)
	fmt.Printf("volume %d, area %d", volume, area)
}
