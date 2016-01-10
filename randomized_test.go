package randomized

import (
	"math/rand"
	"testing"
	"time"
)

type weight struct {
	weight    float64
	threshold float64
	count     int
}

func BenchmarkRandomized(b *testing.B) {
	b.StopTimer()

	lists := []*weight{&weight{weight: 0.2}, &weight{weight: 0.3}, &weight{weight: 0.4}, &weight{weight: 0.1}}

	totalW := 0.0
	for _, v := range lists {
		(*v).threshold = totalW
		totalW += v.weight
	}

	rand.Seed(time.Now().UnixNano())

	b.StartTimer()

	for i := 0; i < b.N; i++ {
		random := rand.Float64()
		for i := len(lists) - 1; i >= 0; i-- {
			if lists[i].threshold <= random {
				lists[i].count++
				break
			}
		}
	}
}

type weightInt struct {
	value int
	count int
}

func BenchmarkRandomizedInt(b *testing.B) {
	b.StopTimer()
	lists := []*weightInt{&weightInt{2, 0}, &weightInt{2, 0}, &weightInt{3, 0}, &weightInt{3, 0}, &weightInt{3, 0}, &weightInt{4, 0}, &weightInt{4, 0}, &weightInt{4, 0}, &weightInt{4, 0}, &weightInt{1, 0}}
	rand.Seed(time.Now().UnixNano())

	b.StartTimer()

	for i := 0; i < b.N; i++ {
		lists[rand.Intn(len(lists))].count++
	}
}
