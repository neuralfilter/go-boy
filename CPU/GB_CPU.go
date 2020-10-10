package cpu

type register struct {
	a  uint8
	b  uint8
	c  uint8
	d  uint8
	e  uint8
	f  uint8
	h  uint8
	l  uint8
	sp uint16
	pc uint16
}

type flag struct {
	Z bool
	N bool
	H bool
	C bool
}
