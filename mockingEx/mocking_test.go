package mockingEx

import "testing"

type MockStruct struct {
	Elem int
}

func (s MockStruct) GetVolume() int {
	return s.Elem * s.Elem * s.Elem
}
func (s MockStruct) GetArea() int {
	return s.Elem * s.Elem * 6
}

func TestCube(t *testing.T) {
	for key, value := range map[MockStruct][]int{
		MockStruct{Elem: 5}: []int{125, 150},
		MockStruct{Elem: 2}: []int{8, 24},
		MockStruct{Elem: 1}: []int{1, 6},
	} {
		t.Run("mock values", func(t *testing.T) {
			volume, area := GetVolumeandSurfaceCube(&key)
			if volume != value[0] {
				t.Errorf("expected %d, got %d", value[0], volume)
			}
			if area != value[1] {
				t.Errorf("expected %d, got %d", value[1], area)
			}
		})

	}
}
func TestGetVolume(t *testing.T) {
	value := &Measurements{5, 5, 5}
	volume := value.GetVolume()
	if volume != 125 {
		t.Errorf("expected 125, got %d", volume)
	}
}
func TestGetArea(t *testing.T) {
	value := &Measurements{5, 5, 5}
	volume := value.GetArea()
	if volume != 150 {
		t.Errorf("expected 125, got %d", volume)
	}
}
