package rwcas

import "github.com/nagata-yoshiteru/go-vector"

// Obstacle :
type Obstacle struct {
	ID       int
	IsConvex bool
	Point    vector.Vector
	Shape    []vector.Vector
	UnitDir  vector.Vector
}

// NewObstacle : To create new obstacle object
func NewObstacle(id int, isConvex bool, point vector.Vector, shape []vector.Vector, unitDir vector.Vector) *Obstacle {
	o := &Obstacle{
		ID:       id,
		IsConvex: isConvex,
		Point:    point,
		Shape:    shape,
		UnitDir:  unitDir,
	}
	return o
}

// NewEmptyObstacle : To create new empty obstacle object
func NewEmptyObstacle(id int) *Obstacle {
	return NewObstacle(id, false, vector.NewWithValues([]float64{0.0, 0.0, 0.0}), nil, vector.NewWithValues([]float64{0.0, 0.0, 0.0}))
}

// NewUnitObstacle : To create new unit obstacle object
func NewUnitObstacle(id int, point vector.Vector, unitDir vector.Vector) *Obstacle {
	shape := []vector.Vector{
		vector.NewWithValues([]float64{-0.5, -0.5, -0.5}),
		vector.NewWithValues([]float64{-0.5, -0.5, 0.5}),
		vector.NewWithValues([]float64{-0.5, 0.5, -0.5}),
		vector.NewWithValues([]float64{-0.5, 0.5, 0.5}),
		vector.NewWithValues([]float64{0.5, -0.5, -0.5}),
		vector.NewWithValues([]float64{0.5, -0.5, 0.5}),
		vector.NewWithValues([]float64{0.5, 0.5, -0.5}),
		vector.NewWithValues([]float64{0.5, 0.5, 0.5}),
	}
	return NewObstacle(id, false, point, shape, unitDir)
}
