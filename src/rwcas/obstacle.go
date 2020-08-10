package rwcas

import "github.com/nagata-yoshiteru/go-vector"

// Obstacle :
type Obstacle struct {
	ID       int
	Position vector.Vector
	Shape    []vector.Vector
	Velocity vector.Vector
}

// NewObstacle : To create new obstacle object
func NewObstacle(id int, position vector.Vector, shape []vector.Vector, velocity vector.Vector) *Obstacle {
	o := &Obstacle{
		ID:       id,
		Position: position,
		Shape:    shape,
		Velocity: velocity,
	}
	return o
}

// NewEmptyObstacle : To create new empty obstacle object
func NewEmptyObstacle(id int) *Obstacle {
	return NewObstacle(id, vector.NewWithValues([]float64{0.0, 0.0, 0.0}), nil, vector.NewWithValues([]float64{0.0, 0.0, 0.0}))
}

// NewUnitObstacle : To create new unit obstacle object
func NewUnitObstacle(id int, position vector.Vector, velocity vector.Vector) *Obstacle {
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
	return NewObstacle(id, position, shape, velocity)
}
