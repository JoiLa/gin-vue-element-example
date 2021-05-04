/**
 * description:
 * author: lfs <1144828910@qq.com>
 * Date: 2021/3/8
 */
package main

import (
    `math`

    _ `github.com/manucorporat/try`
)

type Point struct {
    X, Y float64
}

func Polygon(n int) []Point {
    result := make([]Point, n)
    for i := 0; i < n; i++ {
        a := float64(i)*2*math.Pi/float64(n) - math.Pi/2
        result[i] = Point{math.Cos(a), math.Sin(a)}
    }
    return result
}

func main() {
    // img := slide.RandomGetOneImage()
    // dc := gg.NewContext(img.Bounds().Dx(), img.Bounds().Dy())
    // hk := gg.NewContext(40, 40)
    // hk := image.NewRGBA(image.Rect(0, 0, img.Bounds().Dx(), img.Bounds().Dy())).SubImage(image.Rect(0, 0, img.Bounds().Dx(), img.Bounds().Dy()))
    // gg.SavePNG("out.png", hk)
    /*hk.Push()
      hk.NewSubPath()
      hk.MoveTo(20, 0)
      hk.LineTo(40, 20)
      hk.LineTo(20, 40)
      hk.LineTo(0, 20)
      hk.SetRGBA255(255, 255, 255, 160)
      // hk.DrawImage(img, 0, 0)
      hk.Fill()
      hk.Pop()*/
    /*dc.DrawImage(img, 0, 0)
      dc.DrawImage(hk.Image(), 30, 30)
      _ = dc.SavePNG("out.png")*/

}
