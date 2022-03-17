
package rtree

import(
	"os"
	"image"
	"image/png"
	"image/color"
	"fmt"
)

func rectangle(img *image.RGBA, width, height int, color color.Color, sx, sy, ex, ey float64, fill bool) {

	dsx := int(float64(width)*sx)
	dsy := int(float64(height)*sy)
	dex := int(float64(width)*ex)
	dey := int(float64(height)*ey)

    for y := dsy; y <= dey; y++ {
        for x := dsx; x <= dex; x++ {
			if !fill && (y != dsy && y != dey && x != dsx && x != dex){ continue }
            img.Set(x, y, color)
        }
    }
	
}

func (node *rnode[T, OBJ]) drawRnode(img *image.RGBA, width, height int, root_rect *rect[T], depth, max_depth float64){
	
	sx := float64(node.rect.start[0]-root_rect.start[0])/float64(root_rect.end[0]-root_rect.start[0])
	sy := float64(node.rect.start[1]-root_rect.start[1])/float64(root_rect.end[1]-root_rect.start[1])
	ex := float64(node.rect.end[0]-root_rect.start[0])/float64(root_rect.end[0]-root_rect.start[0])
	ey := float64(node.rect.end[1]-root_rect.start[1])/float64(root_rect.end[1]-root_rect.start[1])

	
	if node.isObject(){
		col := uint8(0)
		rectangle(img, width, height, color.RGBA{col, col, col, 255}, sx, sy, ex, ey, true)
	} else {
		col := 255-uint8(depth/max_depth*255.0)
		rectangle(img, width, height, color.RGBA{col, col, col, 255}, sx, sy, ex, ey ,false)
	}
	
	for i := range node.children {
		node.children[i].drawRnode(img, width, height, root_rect, depth+1, max_depth)
	}
}

func (rtree *Rtree[T, OBJ]) WriteToPNG(width int){

	height := int(float64(width)/float64(rtree.root.rect.end[0]-rtree.root.rect.start[0])*float64(rtree.root.rect.end[1]-rtree.root.rect.start[1]))
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	
	file, _ := os.Create("image.png")
    defer file.Close()

	fmt.Println(rtree.root.getDepth())

	rtree.root.drawRnode(img, width, height, &rtree.root.rect, 0, float64(rtree.root.getDepth()))

    if err := png.Encode(file, img); err != nil {
        panic(err)
    }

}