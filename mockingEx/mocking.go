package mockingEx

type Cube interface {
	GetVolume() int
	GetArea() int
}
type Measurements struct {
	Height int
	Length int
	Width  int
}

func GetVolumeandSurfaceCube(val Cube) (int, int) {
	volume := val.GetVolume()
	area := val.GetArea()
	return volume, area
}
func (m Measurements) GetVolume() int {
	result := m.Height * m.Length * m.Width
	return result
}
func (m Measurements) GetArea() int {
	result := (m.Height * m.Height * 2) + (m.Length * m.Length * 2) + (m.Width * m.Width * 2)
	return result
}
