package rwcas

import "github.com/atedja/go-vector"

var (
	AgentTypes []*AgentType
)

func init() {
	AgentTypes = append(AgentTypes, &AgentType{ // append sample agent (walker)
		ID:               0,
		Name:             "Default Agent",
		Radius:           1.0, // radius size of agent
		TimeHorizonAgent: 5.0,
		TimeHorizonObst:  5.0,
		TimeHorizonWall:  10.0,
		MaxNeighbors:     100,
		NeighborDist:     10,
		MaxSpeed:         5,
	})
}

// Agent :
type Agent struct {
	ID                int
	AgentType         int
	Position          vector.Vector
	PrevVelocity      vector.Vector
	PrefVelocity      vector.Vector
	NextVelocity      vector.Vector
	WallNeighbors     []*WallNeighbor
	ObstacleNeighbors []*ObstacleNeighbor
	AgentNeighbors    []*AgentNeighbor
	Goal              vector.Vector
	OrcaLines         []*Line
	Status            int // 0: created, 1: moving, 2: goal, 3: to be deleted
}

// AgentType :
type AgentType struct {
	ID               int
	Name             string
	Radius           float64 // radius size of agent
	TimeHorizonAgent float64
	TimeHorizonObst  float64
	TimeHorizonWall  float64
	MaxNeighbors     int
	NeighborDist     float64
	MaxSpeed         float64
}

// WallNeighbor :
type WallNeighbor struct {
	DistSq float64
	Wall   *Wall
}

// ObstacleNeighbor :
type ObstacleNeighbor struct {
	DistSq   float64
	Obstacle *Obstacle
}

// AgentNeighbor :
type AgentNeighbor struct {
	DistSq float64
	Agent  *Agent
}

// Line :
type Line struct {
	Point     vector.Vector
	Direction vector.Vector
}

// NewAgent :
func NewAgent(id int, agentType int, position vector.Vector, prevVelocity vector.Vector, prefVelocity vector.Vector, goal vector.Vector) *Agent {
	a := &Agent{
		ID:                id,
		AgentType:         agentType,
		Position:          position,
		PrevVelocity:      prevVelocity,
		PrefVelocity:      prefVelocity,
		NextVelocity:      vector.NewWithValues([]float64{0.0, 0.0, 0.0}),
		WallNeighbors:     make([]*WallNeighbor, 0),
		ObstacleNeighbors: make([]*ObstacleNeighbor, 0),
		AgentNeighbors:    make([]*AgentNeighbor, 0),
		Goal:              goal,
		OrcaLines:         make([]*Line, 0),
		Status:            0, // 0: created, 1: moving, 2: goal, 3: to be deleted
	}
	return a
}

// NewEmptyAgent :
func NewEmptyAgent(id int) *Agent {
	return NewAgent(
		id,
		0,
		vector.NewWithValues([]float64{0.0, 0.0, 0.0}), // position = (0, 0, 0)
		vector.NewWithValues([]float64{0.0, 0.0, 0.0}), // prevVelocity = (0, 0, 0)
		vector.NewWithValues([]float64{0.0, 0.0, 0.0}), // prefVelocity = (0, 0, 0)
		vector.NewWithValues([]float64{0.0, 0.0, 0.0}), // goal = (0, 0, 0)
	)
}
