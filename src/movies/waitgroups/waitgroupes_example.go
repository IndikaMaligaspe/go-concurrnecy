package waitgroups

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/indikamaligaspe/go-concurrnecy/src/movies"
)

var wgcache = map[int]movies.Movie{}
var rnd = rand.New(rand.NewSource(time.Now().UnixNano()))

// StartWaitGroup : StartWaitGroup
func StartWaitGroup() {
	wg := &sync.WaitGroup{}
	m := &sync.RWMutex{}
	for i := 0; i < 10; i++ {
		fmt.Printf("Run %v : ", (i + 1))
		id := rnd.Intn(10) + 1
		wg.Add(2)
		go func(id int, wg *sync.WaitGroup, m *sync.RWMutex) {
			movie, ok := queryCahce(id, m)
			if ok {
				fmt.Println("From Cache")
				fmt.Println(movie)
			}
			wg.Done()
		}(id, wg, m)
		go func(id int, wg *sync.WaitGroup, m *sync.RWMutex) {
			movie, ok := queryDatabase(id, m)
			if ok {
				fmt.Println("From Database")
				fmt.Println(movie)
			}
			wg.Done()
		}(id, wg, m)
		time.Sleep(150 * time.Millisecond)
	}
	wg.Wait()
}

func queryCahce(id int, m *sync.RWMutex) (movies.Movie, bool) {
	m.RLock()
	movie, ok := wgcache[id]
	m.RUnlock()
	return movie, ok
}

func queryDatabase(id int, m *sync.RWMutex) (movies.Movie, bool) {
	time.Sleep(300 * time.Millisecond)
	for _, movie := range movies.Movies {
		if movie.ID == id {
			m.Lock()
			wgcache[id] = movie
			m.Unlock()
			return movie, true
		}
	}
	return movies.Movie{}, false
}
