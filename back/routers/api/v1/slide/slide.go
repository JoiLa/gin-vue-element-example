/**
 * description: 滑动验证码
 * author: lfs <1144828910@qq.com>
 * Date: 2021/3/4
 */
package slide

import (
    `bytes`
    `crypto/md5`
    `encoding/base64`
    `encoding/json`
    `fmt`
    `image`
    `image/png`
    `log`
    `math`
    `math/rand`
    `strconv`
    `strings`
    `time`

    `Api/pkg/e`
    `Api/pkg/helper`
    `Api/service/cache`
    "github.com/fogleman/gg"
    `github.com/gin-gonic/gin`
    `github.com/manucorporat/try`
)

/**
 * @Description: 滑动方块大小
 */
type SlideCubeSize struct {
    W float64
    H float64
}

/**
 * @Description: 滑动方块位置
 */
type SlideCubePosition struct {
    X float64
    Y float64
}

/**
 * @Description: 滑动图片源
 */
var slideImagesSource []image.Image

/**
 * @Description: 滑动块大小
 */
var slideCubeSize = SlideCubeSize{
    W: 40,
    H: 40,
}

func init() {
    // 获取当前目录下的文件或目录名(包含路径)
    // temp, err := filepath.Glob(helper.StoragePath("side/*"))
    temp := helper.AppBox.StorageSide.List()
    for _, filePath := range temp {
        r, err := helper.AppBox.StorageSide.Find(filePath)
        if err != nil {
            log.Fatal(err)
        } else {
            img, err := png.Decode(bytes.NewBuffer(r))
            if err != nil {
                log.Fatal(err)
            }
            slideImagesSource = append(slideImagesSource, img)
        }
    }
}

/**
 * @description: 生成滑动图片
 * @param c 客户端会话
 */
func GenerateImage(c *gin.Context) {
    // 随机获取一张图片
    img := RandomGetOneImage()
    // 创建画布，把滑动的底图画上去
    verImg := gg.NewContext(320, 180)
    // 创建画布，把滑块的底图画上去
    verIfy := gg.NewContext(int(slideCubeSize.W), int(slideCubeSize.H))
    // 画上底图
    verImg.DrawImage(img, 0, 0)
    verImgW := verImg.Image().Bounds().Dx()           // 底图宽度
    verImgH := verImg.Image().Bounds().Dy()           // 底图高度
    generateRandLineQuantity := helper.RandInt(5, 12) // 生成随机干扰线数量
    // 画上干扰线
    for i := 0; i < generateRandLineQuantity; i++ {
        x1 := rand.Float64() * float64(verImgW)
        y1 := rand.Float64() * float64(verImgH)
        x2 := rand.Float64() * float64(verImgW)
        y2 := rand.Float64() * float64(verImgH)
        r := rand.Float64()
        g := rand.Float64()
        b := rand.Float64()
        a := rand.Float64()*0.7 + 0.3         // 设置颜色透明度
        lineWidth := rand.Float64()*1.5 + 0.5 // 设置线的宽度
        verImg.SetRGBA(r, g, b, a)            // 设置颜色
        verImg.SetLineWidth(lineWidth)        // 设置线宽度
        verImg.DrawLine(x1, y1, x2, y2)       // 画线
        verImg.Stroke()                       // 描边
    }
    // 随机生成滑块位置
    slideCubePosition := randomGetSlideCubePosition(verImg.Width(), verImg.Height())
    // fmt.Println("滑动图片缺块位置：", slideCubePosition)

    borderAlphaValue := helper.RandInt(200, 255)        // 定义边框透明色值
    borderWidth := helper.RandFloat64(0.3, 1)           // 缺口和滑块的边框宽度
    missingBlockBorderColor := helper.RandInt(180, 255) // 缺块边框颜色
    missingBlockFillColor := helper.RandInt(220, 255)   // 缺块填充颜色

    // =============================================================================================================================
    // 随机生成滑块 形状
    switch helper.RandInt(1, 4) {
    case 1:
        // =============================================================================================================================
        // 绘制矩形验证块
        // 绘制图形
        verIfy.DrawImageAnchored(verImg.Image(), int(slideCubePosition.X*-1), int(slideCubePosition.Y*-1), 0, 0)
        // 绘制矩形
        verIfy.DrawRectangle(0, 0, slideCubeSize.W, slideCubeSize.H)
        // =============================================================================================================================
        // 绘制底图缺块
        // 画一个 方块，做为 拼接滑块的位置
        verImg.DrawRectangle(slideCubePosition.X, slideCubePosition.Y, slideCubeSize.W, slideCubeSize.H)
        // =============================================================================================================================

        break
    case 2:
        // =============================================================================================================================
        // 绘制圆形验证块
        verIfy.SetRGBA(0, 0, 0, 0)
        verIfy.DrawCircle(slideCubeSize.W/2, slideCubeSize.H/2, slideCubeSize.H/2)
        verIfy.StrokePreserve()
        verIfy.Clip()
        verIfy.DrawImageAnchored(verImg.Image(), int(slideCubePosition.X*-1), int(slideCubePosition.Y*-1), 0.0, 0.0)
        verIfy.DrawCircle(slideCubeSize.W/2, slideCubeSize.H/2, slideCubeSize.H/2)
        // =============================================================================================================================
        // 绘制底图缺块
        // 画一个 圆形，做为 拼接滑块的位置
        verImg.DrawCircle(slideCubePosition.X+slideCubeSize.W/2, slideCubePosition.Y+slideCubeSize.H/2, slideCubeSize.H/2)
        // =============================================================================================================================
        break
    case 3:
        // =============================================================================================================================
        // 绘制菱形验证块
        verIfy.SetRGBA(0, 0, 0, 0)
        drawDiamond := func() { // 画菱形
            verIfy.NewSubPath()
            verIfy.MoveTo(20, 0)
            verIfy.LineTo(40, 20)
            verIfy.LineTo(20, 40)
            verIfy.LineTo(0, 20)
            verIfy.ClosePath()
        }
        drawDiamond()
        verIfy.StrokePreserve()
        verIfy.Clip()
        verIfy.DrawImageAnchored(verImg.Image(), int(slideCubePosition.X*-1), int(slideCubePosition.Y*-1), 0.0, 0.0)
        drawDiamond()
        // =============================================================================================================================
        verImg.NewSubPath()
        verImg.MoveTo(slideCubePosition.X+20, slideCubePosition.Y+0)
        verImg.LineTo(slideCubePosition.X+40, slideCubePosition.Y+20)
        verImg.LineTo(slideCubePosition.X+20, slideCubePosition.Y+40)
        verImg.LineTo(slideCubePosition.X+0, slideCubePosition.Y+20)
        verImg.LineTo(slideCubePosition.X+20, slideCubePosition.Y+0)
        verImg.ClosePath()
        // =============================================================================================================================
        break
    case 4:
        // =============================================================================================================================
        // 绘制五角星形验证块
        verIfy.SetRGBA(0, 0, 0, 0)
        r := slideCubeSize.H / 2 / helper.RandFloat64(1.7, 2)
        rot := helper.RandFloat64(1, 360)
        drawStar(verIfy, slideCubeSize.H/2, r, slideCubeSize.W/2, slideCubeSize.H/2, rot)
        verIfy.StrokePreserve()
        verIfy.Clip()
        verIfy.DrawImageAnchored(verImg.Image(), int(slideCubePosition.X*-1), int(slideCubePosition.Y*-1), 0.0, 0.0)
        drawStar(verIfy, slideCubeSize.H/2, r, slideCubeSize.W/2, slideCubeSize.H/2, rot)
        // =============================================================================================================================
        drawStar(verImg, slideCubeSize.H/2, r, slideCubePosition.X+slideCubeSize.W/2, slideCubePosition.Y+slideCubeSize.H/2, rot)
        // =============================================================================================================================
        break

    }
    // =============================================================================================================================
    // 滑动块边框上色
    verIfy.SetRGBA255(missingBlockBorderColor, missingBlockBorderColor, missingBlockBorderColor, borderAlphaValue) // 画颜色为白色
    verIfy.SetLineWidth(borderWidth + helper.RandFloat64(0.1, 0.3))                                                // 设置线条宽度
    verIfy.Stroke()                                                                                                // 描边
    // =============================================================================================================================
    verImg.SetRGBA255(missingBlockFillColor, missingBlockFillColor, missingBlockFillColor, helper.RandInt(100, 170)) // 画底图缺块颜色
    verImg.FillPreserve()                                                                                            // 填充
    verImg.SetRGBA255(missingBlockBorderColor, missingBlockBorderColor, missingBlockBorderColor, borderAlphaValue)   // 画底图缺块边框颜色
    verImg.SetLineWidth(borderWidth)                                                                                 // 设置线条宽度
    verImg.Stroke()                                                                                                  // 描边
    // =============================================================================================================================

    // 种子标记
    seed := generateSeed(slideCubePosition)
    // 保存滑动答案
    cache.SaveSlideAnswer(seed, slideCubePosition)
    // 反馈数据
    helper.NewResult(c).Success(e.SUCCESS, e.GetMsg(e.SUCCESS), gin.H{
        "seed":   seed,
        "y":      slideCubePosition.Y,
        "verImg": `data:image/png;base64,` + imageToBase64(verImg.Image()),
        "verIfy": `data:image/png;base64,` + imageToBase64(verIfy.Image()),
    })
}

/**
 * @Description: 效验滑动是否正确
 * @param c 客户端会话
 */
func VerifySlide(c *gin.Context) {

    try.This(func() {
        slideChunkLeft, err := strconv.ParseFloat(c.Request.PostForm.Get("slideChunkLeft"), 10) // 滑块左边位置
        if err != nil {
            panic(e.InvalidParams)
        }
        seed := c.Request.PostForm.Get("seed") // 种子
        if seed == "" {
            panic(e.InvalidParams)
        }
        result, err := cache.GetSlideValue(seed)
        if err != nil {
            // 缓存数据不存在
            panic(e.VerificationHasFailedPleaseReAcquire)
        }
        var slideCubePosition SlideCubePosition
        cacheResult := result
        err = json.Unmarshal(cacheResult, &slideCubePosition)
        if err != nil {
            panic(e.InvalidParams)
        }
        var diff = math.Abs(slideChunkLeft - slideCubePosition.X) // 阈值范围
        var threshold float64 = 5
        if diff == 0 {
            // 滑动到准确的位置了
            helper.NewResult(c).Success(e.SUCCESS, e.GetMsg(e.SUCCESS), nil)
        } else if threshold >= diff && diff <= threshold {
            // 滑动到阈值范围
            helper.NewResult(c).Success(e.SUCCESS, e.GetMsg(e.SUCCESS), nil)
        } else {
            helper.NewResult(c).Success(e.VerificationFailedPleaseTryAgain, e.GetMsg(e.VerificationFailedPleaseTryAgain), nil)
        }
    }).Catch(func(err try.E) {
        e1 := err.(int)
        // 不存在缓存数据
        helper.NewResult(c).Error(e1, e.GetMsg(e1))
    })

}

/**
 * @Description: 随机获取一张图片
 * @return string
 */
func RandomGetOneImage() image.Image {
    rand.Seed(time.Now().UnixNano()) // 将时间戳设置成种子数
    return slideImagesSource[rand.Intn(len(slideImagesSource))]
}

/**
 * @Description: 随机获取滑动位置
 */
func randomGetSlideCubePosition(W, H int) SlideCubePosition {
    rand.Seed(time.Now().UnixNano()) // 将时间戳设置成种子数
    return SlideCubePosition{
        X: float64(helper.RandInt(int(slideCubeSize.W)+helper.RandInt(30, 60), W-int(slideCubeSize.W)-helper.RandInt(10, 40))),
        Y: float64(helper.RandInt(helper.RandInt(10, 50), H-int(slideCubeSize.H)-helper.RandInt(1, 20))),
    }
}

/**
 * @Description: 图片转base64
 * @param im
 * @return string
 */
func imageToBase64(im image.Image) string {
    buf := new(bytes.Buffer)
    _ = png.Encode(buf, im)
    return base64.StdEncoding.EncodeToString(buf.Bytes())
}

/**
 * @Description: 生成种子
 * @param Position
 * @return string
 */
func generateSeed(Position SlideCubePosition) string {
    data := []byte(strings.Join([]string{
        strconv.FormatInt(int64(Position.X), 10),
        strconv.FormatInt(int64(Position.Y), 10),
        strconv.FormatInt(time.Now().UnixNano(), 10),
    }, ","))
    return fmt.Sprintf("%x", md5.Sum(data))
}

/**
 * @Description:  画五角星
 * @param img
 * @param R 大圆半径
 * @param r 小圆半径
 * @param x 方向圆心位置
 * @param y 方向圆心位置
 * @param rot 旋转角度
 */
func drawStar(img *gg.Context, R, r, x, y, rot float64) {
    img.NewSubPath()
    for i := 0; i < 5; i++ {
        // 因为角度是逆时针计算的，而旋转是顺时针旋转，所以是度数是负值。
        img.LineTo(x+math.Cos((18+72*float64(i)-rot)/180*math.Pi)*R, y-math.Sin((18+72*float64(i)-rot)/180*math.Pi)*R)
        img.LineTo(x+math.Cos((54+72*float64(i)-rot)/180*math.Pi)*r, y-math.Sin((54+72*float64(i)-rot)/180*math.Pi)*r)
    }
    img.ClosePath()
}
