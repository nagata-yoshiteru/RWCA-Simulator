package rwcas

import (
	"sort"
	"time"

	"github.com/nagata-yoshiteru/go-vector"
)

var (
	Rwcas *RWCASimulator
)

func init() {

}

// RWCASimulator : Type Definition of RWCA-Simulator Object
type RWCASimulator struct {
	CurrentTime       time.Time
	Agents            map[int]*Agent
	Walls             map[int]*Wall        // ObstacleVertices
	WallVertices      map[int]*WallVertice // Obstacles
	Obstacles         map[int]*Obstacle
	TimeStep          float64
	NextAgentID       int
	NextWallID        int
	NextWallVerticeID int
	NextObstacleID    int
}

// NewRWCASimulator : RWCASimulator with options
func NewRWCASimulator(currentTime time.Time, timeStep float64) *RWCASimulator {

	sim := &RWCASimulator{
		CurrentTime:       currentTime,
		Agents:            make(map[int]*Agent),
		Walls:             make(map[int]*Wall),
		WallVertices:      make(map[int]*WallVertice),
		Obstacles:         make(map[int]*Obstacle),
		TimeStep:          timeStep, // 0.1 for 10 frames per sec
		NextAgentID:       0,
		NextWallID:        0,
		NextWallVerticeID: 0,
		NextObstacleID:    0,
	}
	Rwcas = sim

	return sim
}

// GetCurrentTime : Get current time
func (rwcas *RWCASimulator) GetCurrentTime() time.Time {
	return rwcas.CurrentTime
}

// SetCurrentTime : Set current time
func (rwcas *RWCASimulator) SetCurrentTime(currentTime time.Time) {
	rwcas.CurrentTime = currentTime
}

// GetCurrentUnixTimestamp : Get current UNIX timestamp
func (rwcas *RWCASimulator) GetCurrentUnixTimestamp() int64 {
	return rwcas.CurrentTime.Unix()
}

// SetCurrentTimeFromUnixTimestamp : Set current time from UNIX timestamp
func (rwcas *RWCASimulator) SetCurrentTimeFromUnixTimestamp(unixTimestamp int64) {
	rwcas.CurrentTime = time.Unix(unixTimestamp, 0)
}

// GetCurrentUnixNanoTimestamp : Get current UNIX nano timestamp
func (rwcas *RWCASimulator) GetCurrentUnixNanoTimestamp() int64 {
	return rwcas.CurrentTime.UnixNano()
}

// SetCurrentTimeFromUnixNanoTimestamp : Set current time from UNIX nano timestamp
func (rwcas *RWCASimulator) SetCurrentTimeFromUnixNanoTimestamp(unixNanoTimestamp int64) {
	nanoSec := unixNanoTimestamp % 1000000000
	rwcas.CurrentTime = time.Unix((unixNanoTimestamp-nanoSec)/1000000000, nanoSec)
}

// GetAgentsMap : Get agents map
func (rwcas *RWCASimulator) GetAgentsMap() map[int]*Agent {
	return rwcas.Agents
}

// GetWallsMap : Get walls map
func (rwcas *RWCASimulator) GetWallsMap() map[int]*Wall {
	return rwcas.Walls
}

// GetObstaclesMap : Get obstacles map
func (rwcas *RWCASimulator) GetObstaclesMap() map[int]*Obstacle {
	return rwcas.Obstacles
}

// GetTimeStep : Get time step
func (rwcas *RWCASimulator) GetTimeStep() float64 {
	return rwcas.TimeStep
}

// SetTimeStep : Set time step
func (rwcas *RWCASimulator) SetTimeStep(timeStep float64) {
	rwcas.TimeStep = timeStep
}

// AddAgent : Add agent with options
func (rwcas *RWCASimulator) AddAgent(agentType int, position vector.Vector, prevVelocity vector.Vector, prefVelocity vector.Vector, goal vector.Vector) (int, bool) {
	agent := NewAgent(rwcas.NextAgentID, agentType, position, prevVelocity, prefVelocity, goal)
	rwcas.Agents[rwcas.NextAgentID] = agent
	rwcas.NextAgentID++
	return agent.ID, true
}

// GetAgent : Get agent by agentID
func (rwcas *RWCASimulator) GetAgent(agentID int) (*Agent, bool) {
	if agent, exists := rwcas.Agents[agentID]; exists {
		return agent, true
	}
	return nil, false
}

// GetAgentIDs : Get agent IDs
func (rwcas *RWCASimulator) GetAgentIDs() []int {
	ids := make([]int, 0, len(rwcas.Agents))
	for k := range rwcas.Agents {
		ids = append(ids, k)
	}
	sort.Ints(ids)
	return ids
}

// SetAgent : Set agent by agentID
func (rwcas *RWCASimulator) SetAgent(agentID int, agent *Agent) bool {
	if _, exists := rwcas.Agents[agentID]; exists {
		rwcas.Agents[agentID] = agent
		return true
	}
	return false
}

// RemoveAgent : Remove agent by agentID
func (rwcas *RWCASimulator) RemoveAgent(agentID int) bool {
	if agent, exists := rwcas.Agents[agentID]; exists {
		agent.Status = 3
		return true
	}
	return false
}

// AddWall : Add wall with options
func (rwcas *RWCASimulator) AddWall(wallType int, vertices []*vector.Vector) (int, bool) {

	if len(vertices) < 2 {
		return -1, false
	}

	wall := NewWall(rwcas.NextWallID, vertices, rwcas.NextWallVerticeID)
	rwcas.Walls[rwcas.NextWallID] = wall
	rwcas.NextWallID++

	for i := 0; i < len(vertices); i++ {
		wallVertice := NewWallVertice(rwcas.NextWallVerticeID, false, nil, nil, *(vertices[i]), nil)
		rwcas.NextWallVerticeID++

		if i != 0 {
			wallVertice.PrevWallVertice = rwcas.WallVertices[wallVertice.ID-1]
			wallVertice.PrevWallVertice.NextWallVertice = wallVertice
		}

		if i == len(vertices)-1 {
			wallVertice.NextWallVertice = rwcas.WallVertices[wall.Head]
			wallVertice.NextWallVertice.PrevWallVertice = wallVertice
		}

		ti := i + 1
		if i == len(vertices)-1 {
			ti = 0
		}

		wallVertice.UnitDir = vector.Unit(vector.Subtract(*(vertices[ti]), *(vertices[i])))

		ki := i - 1
		if i == 0 {
			ki = len(vertices) - 1
		}

		if len(vertices) == 2 {
			wallVertice.IsConvex = true
		} else {
			wallVertice.IsConvex = (vector.LeftOf(*(vertices[ki]), *(vertices[i]), *(vertices[ti])) >= 0.0)
		}
	}

	return wall.ID, true
}

// GetWall : Get wall by wallID
func (rwcas *RWCASimulator) GetWall(wallID int) (*Wall, bool) {
	if wall, exists := rwcas.Walls[wallID]; exists {
		return wall, true
	}
	return nil, false
}

// GetWallIDs : Get wall IDs
func (rwcas *RWCASimulator) GetWallIDs() []int {
	ids := make([]int, 0, len(rwcas.Walls))
	for k := range rwcas.Walls {
		ids = append(ids, k)
	}
	sort.Ints(ids)
	return ids
}

// RemoveWall : Remove wall by wallID
func (rwcas *RWCASimulator) RemoveWall(wallID int) bool {
	if wall, exists := rwcas.Walls[wallID]; exists {
		for i := wall.Head; i < wall.Head+len(wall.Vertices); i++ {
			delete(rwcas.WallVertices, i)
		}
		delete(rwcas.Walls, wallID)
		return true
	}
	return false
}
