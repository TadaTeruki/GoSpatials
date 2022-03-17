
package rtree

import (
	"testing"
	"math/rand"
)

func FuzzSearch(f *testing.F){

	f.Add(int64(0), 0)

	f.Fuzz(func(t *testing.T, seed int64, data_num int) {

		if data_num >= 1000000 || data_num < 0{
			t.Skip()
		}

		rand.Seed(seed)

		search_start, search_end := []float64{ rand.Float64(), rand.Float64() }, []float64{ rand.Float64(), rand.Float64() }

		for i := 0; i < 2; i++ {
			search_start[i], search_end[i]  = min(search_start[i], search_end[i]), max(search_start[i], search_end[i])
		}
	
		rtree := New[float64, struct{}](2, 10)

		points_num_inside_of_search_range := 0

		for i := 0; i < data_num; i++ {
		
			point := []float64{ rand.Float64(), rand.Float64() }
	
			point_is_inside := 
				search_start[0] < point[0] && point[0] < search_end[0] &&
				search_start[1] < point[1] && point[1] < search_end[1]

			if point_is_inside {
				points_num_inside_of_search_range++
			}

			rtree.StorePoint(point, struct{}{})
		}

		want := points_num_inside_of_search_range 


		get := 0
		rtree.Search(search_start, search_end, func(start, end []float64, obj struct{}){
			get++
		})

		if get != want {
			t.Fatalf("want=%v, got=%v", want, get)
		}


	})
	
}


func TestSearch(t *testing.T){
	
	t.Run("success", func(t *testing.T){

		data_num := 100000

		search_start, search_end := []float64{ 0.2, 0.5 }, []float64{ 0.3, 0.7 }
	
		rtree := New[float64, struct{}](2, 2)

		points_num_inside_of_search_range := 0

		for i := 0; i < data_num; i++ {
		
			point := []float64{ rand.Float64(), rand.Float64() }
	
			point_is_inside := 
				search_start[0] < point[0] && point[0] < search_end[0] &&
				search_start[1] < point[1] && point[1] < search_end[1]

			if point_is_inside {
				rtree.StorePoint(point, struct{}{})
				points_num_inside_of_search_range++
			} else {
				rtree.StorePoint(point, struct{}{})
			}
			
		}

		want := points_num_inside_of_search_range 


		get := 0
		rtree.Search(search_start, search_end, func(start, end []float64, obj struct{}){
			get++
		})

		if get != want {
			t.Fatalf("want=%v, got=%v", want, get)
		}

	})

}

