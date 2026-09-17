package holiday

import (
	"sync"
	"sync/atomic"
	"testing"
)

func TestGetYearDataInitializesOnceConcurrently(t *testing.T) {
	const year = -1
	var calls int32
	want := Day{Name: "Test holiday", Date: "2007-01-01", IsOffDay: true}
	entry := &lazyYearData{init: func() map[string]Day {
		atomic.AddInt32(&calls, 1)
		return map[string]Day{want.Date: want}
	}}
	yearData[year] = entry
	t.Cleanup(func() { delete(yearData, year) })
	if calls != 0 || entry.data != nil {
		t.Fatal("year data initialized before its first request")
	}

	const callers = 64
	start := make(chan struct{})
	results := make(chan map[string]Day, callers)
	var wg sync.WaitGroup
	for i := 0; i < callers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			<-start
			data := GetYearData(year)
			if got := data[want.Date]; got != want {
				t.Errorf("incomplete data: got %+v, want %+v", got, want)
			}
			results <- data
			// Each caller can modify its own map while others read the cache.
			delete(data, want.Date)
			data["caller-owned"] = want
		}()
	}
	close(start)
	wg.Wait()
	close(results)
	for data := range results {
		if got := data["caller-owned"]; got != want {
			t.Errorf("incomplete data: got %+v, want %+v", got, want)
		}
	}
	if got := GetYearData(year); len(got) != 1 || got[want.Date] != want {
		t.Fatalf("caller mutation changed cached data: %v", got)
	}
	if got := atomic.LoadInt32(&calls); got != 1 {
		t.Fatalf("initializer called %d times, want 1", got)
	}
	if got := GetYearData(year - 1); got != nil {
		t.Errorf("unsupported year returned %v, want nil", got)
	}
}
