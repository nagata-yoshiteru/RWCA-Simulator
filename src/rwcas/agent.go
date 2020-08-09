package rwcas

import (
	"github.com/nagata-yoshiteru/go-vector"
)

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

// IsReachedGoal :
func (agent *Agent) IsReachedGoal() bool {
	/* Check if agent have reached their goals. */
	if agent.Status > 1 {
		return true
	}
	if vector.Subtract(agent.Goal, agent.Position).Magnitude() > AgentTypes[agent.AgentType].Radius {
		return false
	}
	agent.Status = 2
	return true
}

// GetGoal :
func (agent *Agent) GetGoal() vector.Vector {
	return agent.Goal
}

// SetGoal :
func (agent *Agent) SetGoal(goal vector.Vector) {
	agent.Goal = goal
}

// GetAgentType :
func (agent *Agent) GetAgentType() int {
	return agent.AgentType
}

// SetAgentType :
func (agent *Agent) SetAgentType(agentType int) {
	agent.AgentType = agentType
}

// GetPosition :
func (agent *Agent) GetPosition() vector.Vector {
	return agent.Position
}

// SetPosition :
func (agent *Agent) SetPosition(position vector.Vector) {
	agent.Position = position
}

// GetPrevVelocity :
func (agent *Agent) GetPrevVelocity() vector.Vector {
	return agent.PrevVelocity
}

// SetPrevVelocity :
func (agent *Agent) SetPrevVelocity(prevVelocity vector.Vector) {
	agent.PrevVelocity = prevVelocity
}

// GetPrefVelocity :
func (agent *Agent) GetPrefVelocity() vector.Vector {
	return agent.PrefVelocity
}

// SetPrefVelocity :
func (agent *Agent) SetPrefVelocity(prefVelocity vector.Vector) {
	agent.PrefVelocity = prefVelocity
}

// GetNextVelocity :
func (agent *Agent) GetNextVelocity() vector.Vector {
	return agent.NextVelocity
}

// SetNextVelocity :
func (agent *Agent) SetNextVelocity(nextVelocity vector.Vector) {
	agent.NextVelocity = nextVelocity
}

// GetWallNeighbors :
func (agent *Agent) GetWallNeighbors() []*WallNeighbor {
	return agent.WallNeighbors
}

// SetWallNeighbors :
func (agent *Agent) SetWallNeighbors(wallNeighbors []*WallNeighbor) {
	agent.WallNeighbors = wallNeighbors
}

// GetObstacleNeighbors :
func (agent *Agent) GetObstacleNeighbors() []*ObstacleNeighbor {
	return agent.ObstacleNeighbors
}

// SetObstacleNeighbors :
func (agent *Agent) SetObstacleNeighbors(obstacleNeighbors []*ObstacleNeighbor) {
	agent.ObstacleNeighbors = obstacleNeighbors
}

// GetAgentNeighbors :
func (agent *Agent) GetAgentNeighbors() []*AgentNeighbor {
	return agent.AgentNeighbors
}

// SetAgentNeighbors :
func (agent *Agent) SetAgentNeighbors(agentNeighbors []*AgentNeighbor) {
	agent.AgentNeighbors = agentNeighbors
}

// GetStatus :
func (agent *Agent) GetStatus() int {
	return agent.Status
}

// SetStatus :
func (agent *Agent) SetStatus(status int) {
	agent.Status = status
}

// GetID :
func (agent *Agent) GetID() int {
	return agent.ID
}

// GetOrcaLines :
func (agent *Agent) GetOrcaLines() []*Line {
	return agent.OrcaLines
}

// SetOrcaLines :
func (agent *Agent) SetOrcaLines(orcaLines []*Line) {
	agent.OrcaLines = orcaLines
}
