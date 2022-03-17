
package rtree

import "testing"
import "math/rand"


func TestPoint(t *testing.T) {

	var dim uint = 2

	rtree := New[float64, string](dim, 5)
	/*
	rtree.Store([]float64{ 3.0, 3.0 }, []float64{ 4.0, 4.0 }, "hello")
	rtree.Store([]float64{ 2.0, 1.0 }, []float64{ 3.0, 2.0 }, "bye-bye")
	rtree.Store([]float64{ 8.0, 3.0 }, []float64{ 9.0, 5.0 }, "yes")
	rtree.Store([]float64{ 3.0, 4.5 }, []float64{ 6.0, 5.0 }, "no")
	*/
	
	r := 110.0
	
	for i := 0; i < 200; i++ {
		start := make([]float64, dim)
		end   := make([]float64, dim)
		for d := 0; d < int(dim); d++{
			
			start[d] = r*rand.Float64()	
			end  [d] = start[d]+rand.Float64()
		}
		rtree.Store(start, end, "")
	}
	
	

	rtree.WriteToPNG(500)
	
}

