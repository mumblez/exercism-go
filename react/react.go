package react

const testVersion = 4

type reactor struct {
	cells []*computeCell
}

type computeCell struct {
	compute func() int
	cbacks  map[CallbackHandle]func(int)
}

type inputCell struct {
	value int
	*reactor
}

// New returns a reactor
func New() *reactor {
	return &reactor{cells: make([]*computeCell, 0)}
}

// CreateInput ...
func (r *reactor) CreateInput(v int) InputCell {
	return &inputCell{v, r}
}

// CreateCompute1 ...
func (r *reactor) CreateCompute1(c Cell, f func(int) int) ComputeCell {
	v := f(c.Value())
	cc := &computeCell{cbacks: make(map[CallbackHandle]func(int), 0)}
	cc.compute = func() int {
		newv := f(c.Value())
		if newv != v {
			for _, cb := range cc.cbacks {
				cb(newv)
			}
			v = newv
		}
		return newv
	}
	r.cells = append(r.cells, cc)
	return cc
}

// CreateCompute2 ...
func (r *reactor) CreateCompute2(c1, c2 Cell, f func(int, int) int) ComputeCell {
	v := f(c1.Value(), c2.Value())
	cc := &computeCell{cbacks: make(map[CallbackHandle]func(int), 0)}
	cc.compute = func() int {
		newv := f(c1.Value(), c2.Value())
		if newv != v {
			for _, cb := range cc.cbacks {
				cb(newv)
			}
			v = newv
		}
		return newv
	}
	r.cells = append(r.cells, cc)
	return cc
}

// SetValue ...
func (ic *inputCell) SetValue(v int) {
	ic.value = v
	for _, cc := range ic.cells {
		cc.compute()
	}
}

// AddCallback ...
func (cc *computeCell) AddCallback(f func(int)) CallbackHandle {
	cc.cbacks[&f] = f
	return &f
}

// RemoveCallback ...
func (cc *computeCell) RemoveCallback(cbh CallbackHandle) {
	delete(cc.cbacks, cbh)
}

// Value ...
func (cc *computeCell) Value() int {
	return cc.compute()
}

// Value ...
func (ic *inputCell) Value() int {
	return ic.value
}
