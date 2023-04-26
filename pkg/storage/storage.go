package storage

import (
	"errors"
	"fmt"
	"log"
	pUser "students/pkg/user"
)

type Storage struct {
	store map[int]*pUser.User
}

func NewStorage() Storage {
	return Storage{make(map[int]*pUser.User, 0)}
}

func (s Storage) Get(id int) (*pUser.User, error) {
	if exist := s.contains(id); !exist {
		return nil, errors.New(fmt.Sprint("User not found"))
	}

	return s.store[id], nil
}

func (s Storage) GetFriends(id int) ([]*pUser.User, error) {
	user, err := s.Get(id)
	if err != nil {
		return nil, err
	}

	friends := make([]*pUser.User, 0, len(user.Friends))
	for _, friendId := range user.Friends {
		if val, ok := s.store[friendId]; ok {
			friends = append(friends, val)
		}
	}

	return friends, nil
}

func (s Storage) Add(u *pUser.User) (bool, error) {
	u.Id = len(s.store) + 1
	s.store[u.Id] = u

	log.Println(s.store)

	return true, nil
}

func (s Storage) Put(id int, user *pUser.User) (bool, error) {
	existUser, err := s.Get(id)
	if err != nil {
		return false, err
	}

	if user.Name != "" {
		existUser.Name = user.Name
	}

	if user.Age != 0 {
		existUser.Age = user.Age
	}

	s.store[existUser.Id] = existUser

	return true, nil
}

func (s Storage) Attach(suId int, tuId int) (bool, error) {
	sourceUser, err := s.Get(suId)
	if err != nil {
		return false, err
	}

	targetUser, err := s.Get(tuId)
	if err != nil {
		return false, err
	}

	founded := false
	for _, fid := range sourceUser.Friends {
		if fid == targetUser.Id {
			founded = true
			break
		}
	}

	if !founded {
		sourceUser.Friends = append(sourceUser.Friends, targetUser.Id)
	}

	s.store[sourceUser.Id] = sourceUser

	return true, nil
}

func (s Storage) Delete(id int) (bool, error) {
	user, err := s.Get(id)
	if err != nil {
		return false, err
	}

	for id, userFroStore := range s.store {
		for i, friendId := range userFroStore.Friends {
			if user.Id == friendId {
				s.store[id].Friends = append(userFroStore.Friends[:i], userFroStore.Friends[i+1:]...)
				break
			}
		}
	}

	delete(s.store, user.Id)

	return true, nil
}

func (s Storage) contains(id int) bool {
	if _, ok := s.store[id]; ok {
		return true
	}

	return false
}
