package display

type display struct {
	values [][]int
}
type Pixel struct {
	y, x int
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

func (d *display) Animate(corners bool) {
	nextDisplay := make([][]int, len(d.values))
	outDisplay := make([][]int, len(d.values))
	for i := range nextDisplay {
		nextDisplay[i] = make([]int, len(d.values[0]))
		outDisplay[i] = make([]int, len(d.values[0]))
	}
	for i, v := range d.values {
		for j := range v {
			nextDisplay[j][i] = d.Neighbors(NewPixel(i, j))
		}
	}
	for i, v := range d.values {
		for j := range v {
			outDisplay[j][i] = test(d.values[j][i], nextDisplay[j][i])
		}
	}
	d.values = outDisplay
	if corners {
		d.stuckCorners()
	}
}
func (d *display) stuckCorners() {
	lenX := len(d.values) - 1
	lenY := len(d.values) - 1
	d.values[0][0] = 1
	d.values[0][lenY] = 1
	d.values[lenX][0] = 1
	d.values[lenX][lenY] = 1
}

func test(previous, neighbors int) int {
	switch {
	case previous == 1 && (neighbors == 2 || neighbors == 3):
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
