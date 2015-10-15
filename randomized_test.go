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

	lists := []*weight{&weight{weight: 1.0}, &weight{weight: 3.0}, &weight{weight: 5.0}, &weight{weight: 1.0}}

	totalW := 0.0
	for _, v := range lists {
		(*v).threshold = totalW
		totalW += v.weight
	}

	rand.Seed(time.Now().UnixNano())

	b.StartTimer()

	for i := 0; i < b.N; i++ {
		rTotalW := rand.Float64() * totalW
		for i := len(lists) - 1; i >= 0; i-- {
			if lists[i].threshold <= rTotalW {
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
	lists := []*weightInt{&weightInt{1, 0}, &weightInt{3, 0}, &weightInt{3, 0}, &weightInt{3, 0}, &weightInt{5, 0}, &weightInt{5, 0}, &weightInt{5, 0}, &weightInt{5, 0}, &weightInt{5, 0}, &weightInt{1, 0}}
	rand.Seed(time.Now().UnixNano())

	b.StartTimer()

	for i := 0; i < b.N; i++ {
		lists[rand.Intn(len(lists))].count++
	}
}
