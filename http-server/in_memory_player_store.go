package main

import "sync"

func NewInMemoryPlayerStore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{
		store:   map[string]int{},
		RWMutex: sync.RWMutex{},
	}
}

type InMemoryPlayerStore struct {
	store map[string]int
	sync.RWMutex
}

func (i *InMemoryPlayerStore) RecordWin(name string) {
	i.Lock()
	defer i.Unlock()
	i.store[name]++
}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	i.RLock()
	defer i.RUnlock()
	return i.store[name]
}

func (i *InMemoryPlayerStore) GetLeague() []Player {
	var league []Player
	for name, wins := range i.store {
		league = append(league, Player{name, wins})
	}
	return league
}
