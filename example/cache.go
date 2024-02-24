package example

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

var (
	errUserNotInCache = errors.New("the user isn't in cache")
)

type user struct {
	Id    int64  `json:"id"`
	Email string `json:"email"`
}

type cachedUser struct {
	user
	expireAtTimestamp int64
}

type localCache struct {
	stop  chan struct{}
	mu    sync.RWMutex
	users map[int64]cachedUser
}

func newLocalCache(interval time.Duration) *localCache {
	lc := localCache{
		stop:  make(chan struct{}),
		mu:    sync.RWMutex{},
		users: make(map[int64]cachedUser),
	}

	go func(cleanupInterval time.Duration) {
		lc.cleanupLoop(cleanupInterval)

	}(interval)

	return &lc
}

func (l *localCache) cleanupLoop(interval time.Duration) {
	t := time.NewTicker(interval)
	defer t.Stop()

	for {
		select {
		case <-l.stop:
			return
		case <-t.C:
			l.mu.Lock()

			for key, value := range l.users {
				if value.expireAtTimestamp <= time.Now().Unix() {
					delete(l.users, key)
				}
			}
			l.mu.Unlock()
		}
	}

}

func (l *localCache) stopCleanup() {
	l.stop <- struct{}{}
}

func (l *localCache) update(u user, expireAtTimestamp int64) {
	l.mu.Lock()
	defer l.mu.Unlock()

	l.users[u.Id] = cachedUser{
		user:              u,
		expireAtTimestamp: expireAtTimestamp,
	}
}

func (l *localCache) read(id int64) (user, error) {
	l.mu.RLock()
	defer l.mu.RUnlock()

	cu, ok := l.users[id]
	if !ok {
		return user{}, errUserNotInCache
	}

	return cu.user, nil
}

func TestCache() {
	newUser := user{
		Id:    1,
		Email: "felipe.meriga@gmail.com",
	}
	lc := newLocalCache(time.Second * 3)

	lc.update(newUser, time.Now().Add(time.Second*6).Unix())

	lc.stopCleanup()
	waitingTicker := time.NewTicker(time.Second * 15)

	example, err := lc.read(1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(example)

	select {
	case <-waitingTicker.C:
		example, err := lc.read(1)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(example)
	}
}
