package structsmethodsinterfaces

import "testing"

func TestArea(t *testing.T) {
	t.Run("area of rectangle", func(t *testing.T) {
		rectangle := Rectangle{
			Length:  50.0,
			Breadth: 30.0,
		}
		got := rectangle.Area()
		want := 1500.0

		if got != want {
			t.Errorf("expected %f , got %f", want, got)
		}
	})
	t.Run("area of circle", func(t *testing.T) {
		circle := Circle{
			Radius: 10,
		}
		got := circle.Area()
		want := 314.1592653589793
		if got != want {
			t.Errorf("expected %f , got %f", want, got)
		}
	})

}
