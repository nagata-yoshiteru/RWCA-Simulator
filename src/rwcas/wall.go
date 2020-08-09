package rwcas

import (
	"github.com/nagata-yoshiteru/go-vector"
)

// Wall :
type Wall struct {
	ID       int
	Vertices []*vector.Vector
	Head     int
}

// NewWall : To create new Wall object
func NewWall(id int, vertices []*vector.Vector, head int) *Wall {
	o := &Wall{
		ID:       id,
		Vertices: vertices,
		Head:     head,
	}
	return o
}

// NewEmptyWall : To create new Wall object
func NewEmptyWall(id int) *Wall {
	return NewWall(id, nil, -1)
}

// WallVertice :
type WallVertice struct {
	ID              int
	IsConvex        bool
	NextWallVertice *WallVertice
	PrevWallVertice *WallVertice
	Vertice         vector.Vector
	UnitDir         vector.Vector
}

// NewWallVertice : To create new WallVertice object
func NewWallVertice(id int, isConvex bool, nextWallVertice *WallVertice, prevWallVertice *WallVertice, vertice vector.Vector, unitDir vector.Vector) *WallVertice {
	o := &WallVertice{
		ID:              id,
		IsConvex:        isConvex,
		NextWallVertice: nextWallVertice,
		PrevWallVertice: prevWallVertice,
		Vertice:         vertice,
		UnitDir:         unitDir,
	}
	return o
}

// NewEmptyWallVertice : To create new WallVertice object
func NewEmptyWallVertice(id int) *WallVertice {
	return NewWallVertice(id, false, nil, nil, vector.NewWithValues([]float64{0.0, 0.0, 0.0}), vector.NewWithValues([]float64{0.0, 0.0, 0.0}))
}
