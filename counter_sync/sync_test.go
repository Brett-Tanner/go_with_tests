package counter_sync

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	t.Run("incrementing counter 3 times leaves it at 3", func(t *testing.T) {
		counter := NewCounter()
		incrementNTimes(counter, 3)

		assertCount(t, counter, 3)
	})

	t.Run("runs safely concurrently", func(t *testing.T) {
		wantedCount := 1000
		counter := NewCounter()

		var wg sync.WaitGroup
		wg.Add(wantedCount)

		for i := 0; i < wantedCount; i++ {
			go func() {
				counter.Inc()
				wg.Done()
			}()
		}
		wg.Wait()

		assertCount(t, counter, wantedCount)
	})
}

func NewCounter() *Counter {
	return &Counter{}
}

func incrementNTimes(counter *Counter, n int) {
	for i := 0; i < n; i++ {
		counter.Inc()
	}
}

func assertCount(t testing.TB, counter *Counter, want int) {
	t.Helper()

	if counter.Value() != want {
		t.Errorf("counter should be %d but is %d", want, counter.Value())
	}
}
