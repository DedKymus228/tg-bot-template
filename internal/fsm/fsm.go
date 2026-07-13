package fsm

import "sync"

type StateManager interface {
	SetState(userID int64, state string)
	GetState(userID int64) string
	ClearState(userID int64)
}

type MemoryFSM struct {
	mu     sync.Mutex
	states map[int64]string
}

func NewMemoryFSM() *MemoryFSM {
	return &MemoryFSM{
		states: make(map[int64]string),
	}
}

func (m *MemoryFSM) SetState(userID int64, state string) {
	m.mu.Lock()
	m.states[userID] = state
	m.mu.Unlock()
}

func (m *MemoryFSM) GetState(userID int64) string {
	m.mu.Lock()
	state := m.states[userID]
	m.mu.Unlock()
	return state
}

func (m *MemoryFSM) ClearState(userID int64) {
	m.mu.Lock()
	m.states[userID] = ""
	m.mu.Unlock()
}
