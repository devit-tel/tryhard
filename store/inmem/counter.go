package inmem

func (s *Store) AddCounterWithTTL(id string, second int) {
	s.addCounter(id, timeAfterSecond(second))
}

func (s *Store) AddCounter(id string) {
	s.addCounter(id, 0)
}

func (s *Store) GetCounter(id string) int {
	if count, exist := s.ListData[id]; exist {
		return count
	}

	return 0
}

func (s *Store) ClearCounter(id string) {
	delete(s.ListData, id)
	delete(s.ListTTL, id)
}
