# GoSpatials
Spatial index collection written in Go 1.18


### Rtree


![image0](https://user-images.githubusercontent.com/69315285/158785147-b04547d4-63e7-4661-973c-02df481ae6ef.png)

![image1](https://user-images.githubusercontent.com/69315285/158785148-09387ba9-7c98-4903-8662-e13d28cd80ac.png)

##### example

```go
package main

import "github.com/TadaTeruki/GoSpatials/Rtree"

func main(){

	var dim uint = 2
  // Make new structure (dimentions, maximum branching factors)
	rtree := New[float64, string](dim, 5)

  // Store objects
	rtree.Store([]float64{ 3.0, 3.0 }, []float64{ 4.0, 4.0 }, "hello")
	rtree.Store([]float64{ 2.0, 1.0 }, []float64{ 3.0, 2.0 }, "bye-bye")
	rtree.Store([]float64{ 8.0, 3.0 }, []float64{ 9.0, 5.0 }, "yes")
	rtree.Store([]float64{ 3.0, 4.5 }, []float64{ 6.0, 5.0 }, "no")

  // Search objects
	rtree.Search([]float64{ 1.0, 2.0 }, []float64{ 6.0, 7.0 }, func(start, end []float64, obj string){
		fmt.Println(obj)
	})

}

```
