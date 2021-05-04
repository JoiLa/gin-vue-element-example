/**
 * Created by lfs
 * Email: 1144828910@qq.com
 * Date: 2021/3/8
 * Time: 21:48
 */
import {apiPost} from "@/utils/request";

export default {
  /**
   * 获取生成的滑动图片
   * @return {AxiosPromise}
   */
  getGenerateImage(){
    return apiPost('/side/generate-image')
  },
  /**
   * 检查滑块位置
   * @return {AxiosPromise}
   */
  checkSlideChunkPosition(Params){
    return apiPost('/side/verify-slide',Params)
  }
}
