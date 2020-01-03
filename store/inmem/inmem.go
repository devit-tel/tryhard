package inmem

import (
	"time"

	"github.com/devit-tel/gotime"
)

type Store struct {
	ListData map[string]int
	ListTTL  map[string]int64
}

func New() *Store {
	s := &Store{
		ListData: map[string]int{},
		ListTTL:  map[string]int64{},
	}

	go s.watchExpireKeyAndDelete()

	return s
}

func (s *Store) watchExpireKeyAndDelete() {
	//TODO: Refactor for test
	for true {
		deleteList := map[string]struct{}{}
		for key, timestamp := range s.ListTTL {
			if timestamp > 0 && timestamp <= gotime.Now().Unix() {
				deleteList[key] = struct{}{}
			}
		}

		for key, _ := range deleteList {
			delete(s.ListTTL, key)
			delete(s.ListData, key)
		}

		time.Sleep(time.Second)
	}
}

func (s *Store) addCounter(id string, expireTimestamp int64) {
	if count, exist := s.ListData[id]; exist {
		s.ListData[id] = count + 1
	} else {
		s.ListData[id] = 1
		s.ListTTL[id] = expireTimestamp
	}
}

func timeAfterSecond(second int) int64 {
	return time.Now().Add(time.Second * time.Duration(second)).Unix()
}
