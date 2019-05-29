package aunt

type Aunt struct {
	attributes map[string]int
	id         int
}

// New creates a new aunt
func New(id int) Aunt {
	var a Aunt
	a.id = id
	a.attributes = make(map[string]int, 0)
	return a
}

// AddValue adds a value and key to the attributes map in Aunt
func (a *Aunt) AddValue(key string, value int) {
	a.attributes[key] = value
}

// CheckKey checks for a key in the Aunt map. Returns the value if true, else 0 and false
func (a *Aunt) CheckKey(key string) (int, bool) {
	elem, check := a.attributes[key]
	if check == true {
		return elem, check
	}
	return 0, check
}

// GetID returns the ID of the specific aunt
func (a Aunt) GetID() int {
	return a.id
}
func (a Aunt) Compare(b Aunt) bool {
	for k, v := range a.attributes {
		val1, _ := b.CheckKey(k)
		switch {
		case k == "cats":
			if v <= val1 {
				return false
			}
		case k == "trees":
			if v <= val1 {
				return false
			}
		case k == "pomeranians":
			if v >= val1 {
				return false
			}
		case k == "goldfish":
			if v >= val1 {
				return false
			}
		default:
			if val1 != v {
				return false
			}
		}
	}
	return true
}
