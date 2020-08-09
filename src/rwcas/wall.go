package rwcas

import (
	"github.com/atedja/go-vector"
)

// Wall :
type Wall struct {
	ID       int
	IsConvex bool
	NextWall *Wall
	PrevWall *Wall
	Point    vector.Vector
	UnitDir  vector.Vector
}

// NewWall : To create new Wall object
func NewWall(id int, isConvex bool, nextWall *Wall, prevWall *Wall, point vector.Vector, unitDir vector.Vector) *Wall {
	o := &Wall{
		ID:       id,
		IsConvex: isConvex,
		NextWall: nextWall,
		PrevWall: prevWall,
		Point:    point,
		UnitDir:  unitDir,
	}
	return o
}

// NewEmptyWall : To create new Wall object
func NewEmptyWall(id int) *Wall {
	return NewWall(id, false, nil, nil, vector.NewWithValues([]float64{0.0, 0.0, 0.0}), vector.NewWithValues([]float64{0.0, 0.0, 0.0}))
}
