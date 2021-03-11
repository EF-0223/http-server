package poker

func NewInMemoryPlayerStore() *InMemoryPlayerStore {
	m := make(map[string]int)
	return &InMemoryPlayerStore{m}
}

type InMemoryPlayerStore struct {
	store map[string]int
}

func (i *InMemoryPlayerStore) RecordWin(name string) {
	i.store[name]++
}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return i.store[name]
}

func (i *InMemoryPlayerStore) GetLeague() (league League) {
	for name, score := range i.store {
		league = append(league, Player{
			name,
			score,
		})
	}
	return
}
