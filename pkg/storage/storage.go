package storage

import (
	"errors"
	"fmt"
	"log"
	"students/pkg/user"
)

type Storage struct {
	store map[int]*user.User
}

func NewStorage() Storage {
	return Storage{make(map[int]*user.User, 0)}
}

func (s Storage) Get(id int) (*user.User, error) {
	if exist := s.contains(id); !exist {
		return nil, errors.New(fmt.Sprint("User not found"))
	}

	return s.store[id], nil
}

func (s Storage) GetFriends(id int) ([]*user.User, error) {
	u, err := s.Get(id)
	if err != nil {
		return nil, err
	}

	friends := make([]*user.User, 0, len(u.Friends))
	for _, friendId := range u.Friends {
		if val, ok := s.store[friendId]; ok {
			friends = append(friends, val)
		}
	}

	return friends, nil
}

func (s Storage) Add(u *user.User) (bool, error) {
	u.Id = len(s.store) + 1
	s.store[u.Id] = u

	log.Println(s.store)

	return true, nil
}

func (s Storage) Put(id int, u *user.User) (bool, error) {
	eu, err := s.Get(id)
	if err != nil {
		return false, err
	}

	if u.Name != "" {
		eu.Name = u.Name
	}

	if u.Age != 0 {
		eu.Age = u.Age
	}

	s.store[eu.Id] = eu

	return true, nil
}

func (s Storage) Attach(suId int, tuId int) (bool, error) {
	source, err := s.Get(suId)
	if err != nil {
		return false, err
	}

	target, err := s.Get(tuId)
	if err != nil {
		return false, err
	}

	founded := false
	for _, fid := range source.Friends {
		if fid == target.Id {
			founded = true
			break
		}
	}

	if !founded {
		source.Friends = append(source.Friends, target.Id)
	}

	s.store[source.Id] = source

	return true, nil
}

func (s Storage) Delete(id int) (bool, error) {
	u, err := s.Get(id)
	if err != nil {
		return false, err
	}

	for id, us := range s.store {
		for i, friendId := range us.Friends {
			if u.Id == friendId {
				s.store[id].Friends = append(us.Friends[:i], us.Friends[i+1:]...)
				break
			}
		}
	}

	delete(s.store, u.Id)

	return true, nil
}

func (s Storage) contains(id int) bool {
	if _, ok := s.store[id]; ok {
		return true
	}

	return false
}
