package structs

import "testing"

func TestPerimeter(t *testing.T) {
	values := Rectangle{10.2, 20.5} //perimeter=2(l+w)
	got := Perimeter(values)
	want := 61.4
	if got != want {
		t.Errorf("got %.2f, want %.2f ", got, want) //prints 2 values after .
	}
}

/*func TestArea(t *testing.T) {
	checkArea := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("got %v, want %v ", got, want)
		}

	}
	t.Run("rectangles", func(t *testing.T) {
		//rectangle := Rectangle{12.0, 6.0}
		rectangle := Rectangle{12.0, 6.0} //perimeter=2(l+w)
		checkArea(t, rectangle, 72.0)

	})
	t.Run("circles", func(t *testing.T) {
		circle := Circle{5.0}
		checkArea(t, circle, 78.53981633974483)

	})
}*/
//Table driven tests
func TestArea(t *testing.T) {
	checkAreas := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{name: "Rectangle", shape: Rectangle{h: 12, w: 6}, want: 72.0},
		{name: "circle", shape: Circle{r: 5.0}, want: 78.53981633974483},
		{name: "Triangle", shape: Triangle{b: 2, h: 5}, want: 5},
	}
	for _, tt := range checkAreas {
		t.Run(tt.name, func(t *testing.T) { //it will give each shape in TestArea
			got := tt.shape.Area()
			if got != tt.want {
				t.Errorf("%#v got %v, want %v ", tt.shape, got, tt.want) //%#v will print the shape, for identifying which test has failed
			}
		})
	}

}
