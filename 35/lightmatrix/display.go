package display

type display struct {
	values [][]int
}
type Pixel struct {
	x, y int
}

func (d display) Neighbors(pixel Pixel) int {
	var testPixel Pixel
	neighborsOn := 0
	for i := -1; i < 2; i++ {
		for j := -1; j < 2; j++ {
			testPixel.x = pixel.x + j
			testPixel.y = pixel.y + i
			switch {
			case testPixel.x < 0:
				continue
			case testPixel.y < 0:
				continue
			case testPixel.x >= len(d.values):
				continue
			case testPixel.y >= len(d.values[0]):
				continue
			case testPixel == pixel:
				continue
			case d.values[testPixel.y][testPixel.x] == 1:
				neighborsOn++
			case d.values[testPixel.y][testPixel.x] == 0:
				continue
			}
		}
	}
	return neighborsOn
}
func NewDisplay(pixel Pixel) display {
	var disp display
	values := make([][]int, pixel.x)
	for i := range values {
		values[i] = make([]int, pixel.y)
	}
	return disp
}
func (d *display) ReplaceDisplay(newDisplay [][]int) {
	d.values = newDisplay
}
func (d display) GetDisplay() [][]int {
	return d.values
}
func (d *display) Animate() {
	nextDisplay := make([][]int, len(d.values))
	outDisplay := make([][]int, len(d.values))
	for i := range nextDisplay {
		nextDisplay[i] = make([]int, len(d.values[0]))
		outDisplay[i] = make([]int, len(d.values[0]))

	}
	for i := 0; i < len(d.values); i++ {
		for j := 0; j < len(d.values[0]); j++ {
			nextDisplay[i][j] = d.Neighbors(NewPixel(i, j))
		}
	}

	for i := 0; i < len(d.values); i++ {
		for j := 0; j < len(d.values[0]); j++ {
			outDisplay[i][j] = test(d.values[i][j], nextDisplay[i][j])
		}
	}
	d.values = outDisplay
}
func test(previous, neighbors int) int {
	switch {
	case previous == 1 && (2 == neighbors || 3 == neighbors):
		return 1
	case previous == 0 && neighbors == 3:
		return 1
	default:
		return 0
	}
}
func (d display) SumDisplay() int {
	total := 0
	for _, v := range d.values {
		for _, x := range v {
			total += x
		}
	}
	return total
}
func NewPixel(x, y int) Pixel {
	var p Pixel
	p.x = x
	p.y = y
	return p
}
