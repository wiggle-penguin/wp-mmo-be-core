package model

type Locatable struct {
	X, Y, Z  int64
	Rotation float
}

type Movable struct {
	Locatable
	Speed int64
}
