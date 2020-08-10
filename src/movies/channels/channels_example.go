package channels

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/indikamaligaspe/go-concurrnecy/src/movies"
)

var chcache = map[int]movies.Movie{}
var rnd = rand.New(rand.NewSource(time.Now().UnixNano()))

// StartChannels : StartChannels
func StartChannels() {
	wg := &sync.WaitGroup{}
	m := &sync.RWMutex{}
	cacheCh := make(chan movies.Movie)
	dbCh := make(chan movies.Movie)

	for i := 1; i < 10; i++ {
		id := rnd.Intn(10) + 1
		wg.Add(2)
		go func(id int, wg *sync.WaitGroup, m *sync.RWMutex, ch chan<- movies.Movie) {
			if movie, ok := queryCahce(id, m); ok {
				ch <- movie
			}
			wg.Done()
		}(id, wg, m, cacheCh)
		go func(id int, wg *sync.WaitGroup, m *sync.RWMutex, ch chan<- movies.Movie) {
			if movie, ok := queryDatabase(id, m); ok {
				m.Lock()
				chcache[id] = movie
				m.Unlock()
				ch <- movie
			}
			wg.Done()
		}(id, wg, m, dbCh)

		go func(cacheCh, dbCh <-chan movies.Movie) {
			select {
			case movie := <-cacheCh:
				fmt.Println("From Cache ->")
				fmt.Println(movie)
				<-dbCh
			case movie := <-dbCh:
				fmt.Println("From Database ->")
				fmt.Println(movie)
			}
		}(cacheCh, dbCh)
		time.Sleep(150 * time.Millisecond)
	}
	wg.Wait()
}

func queryCahce(id int, m *sync.RWMutex) (movies.Movie, bool) {
	m.RLock()
	movie, ok := chcache[id]
	m.RUnlock()
	return movie, ok
}

func queryDatabase(id int, m *sync.RWMutex) (movies.Movie, bool) {
	time.Sleep(300 * time.Millisecond)
	for _, movie := range movies.Movies {
		if movie.ID == id {
			m.Lock()
			chcache[id] = movie
			m.Unlock()
			return movie, true
		}
	}
	return movies.Movie{}, false
}
