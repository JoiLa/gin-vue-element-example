<!--
* Created by lfs
* Email: 1144828910@qq.com
* Date: 2021/3/8
* Time: 20:47
-->
<template>
  <div>

    <el-dialog
      :visible.sync="dialogVisible"
      :close-on-click-modal="false"
      :close-on-press-escape="false"
      :center="true"
      width="370px"
    >
      <div class="box" v-loading="loadingStatus" v-if="dialogVisible">
        <div class="bgBox">
          <el-image class="verImg" :lazy="true" :src="verImg" fit="fill" draggable="false"></el-image>
          <el-image class="verify" :lazy="true" :src="verIfy" fit="fill" :style="{top:verIfyVertical+'px',left:slideChunkLeft+'px'}" draggable="false"/>
          <i class="el-icon-refresh-right refresh" @click="refreshSlideChunk"/>
        </div>
        <div class="handle" ref="refHandleWrapper">
          <div :class="['slideChunk',(showErrorsStatus===true?'slideChunkErrors':'')]"
               :style="{left:slideChunkLeft+'px'}"
               ref="refSlideChunk">
            <i class="el-icon-right" v-if="showErrorsStatus===false"/>
            <i class="el-icon-close" v-else/>
          </div>
          <div :class="['handleMoveBorder',(showErrorsStatus===true?'handleMoveBorderErrors':'')] "
               :style="{width:(slideChunkLeft+40)+'px'}" ref="refHandleMoveBorder"/>
          <span class="message">请滑动图片</span>
        </div>
      </div>
    </el-dialog>

  </div>
</template>

<script>


  import slide from "@/api/slide/slide";

  export default {
    // 父传子数据传递
    props     : {},
    // 注册组件
    components: {},
    // 定义属性
    data() {
      return {
        slide            : {
          verImg        : null,//底图图片
          verIfy        : null,//滑动图片
          verIfyVertical: 0,//滑动图片垂直位置
          seed          : ''//滑动图片的种子Id
        },
        loadingStatus    : false,//加载状态
        dialogVisible    : false,//对话框显示状态
        slideChunkLeft   : 0,//滑动块的左边位置
        slideChunkMaxLeft: 0,//滑动块滑动到左边最大位置
        startMouseDownX  : 0,//起始鼠标横向坐标
        showErrorsStatus : false,//显示错误状态
        parentCallback   : null//父级回调函数
      }
    },
    // 创建完毕状态
    created() {
    },
    // 挂载完毕状态
    mounted() {

    },
    watch     : {
      dialogVisible(e) {
        if (e === false) this.removeMouseEventListener()
      }
    },
    // 事件方法处理集合
    methods   : {
      /**
       * @description 进入滑动图片
       * @param {function}ParentCallback 父级回调函数
       */
      enterSlide(ParentCallback) {
        let _this            = this
        _this.dialogVisible  = true
        _this.parentCallback = ParentCallback
        _this.refreshSlideChunk()
      },
      /**
       * @description 刷新滑块
       */
      refreshSlideChunk() {
        let _this           = this
        _this.loadingStatus = true
        if (_this.$refs.refSlideChunk !== undefined) {
          //移除监听器
          _this.$refs.refSlideChunk.removeEventListener('mousedown', _this.SlideMousedown)
        }
        //获取滑动验证码
        slide.getGenerateImage().then(res => {
          _this.slide.verImg         = res.data.verImg
          _this.slide.verIfy         = res.data.verIfy
          _this.slide.verIfyVertical = res.data.y
          _this.slide.seed           = res.data.seed
          _this.loadingStatus        = false
          _this.slideChunkLeft       = 0
          //添加监听器
          _this.$refs.refSlideChunk.addEventListener('mousedown', _this.SlideMousedown)
        })
      },

      /**
       * @description 检查滑动块位置是否正确
       * @param {Number}slideChunkLeft
       * @param {string}Seed
       * @constructor
       */
      checkSlideChunkPosition(slideChunkLeft, Seed) {
        //console.log('用户给的横向答案位置->', slideChunkLeft, '种子->', Seed)
        let _this           = this
        _this.loadingStatus = true//设置加载状态
        _this.$refs.refSlideChunk.removeEventListener('mousedown', _this.SlideMousedown)//移除监听器
        /**
         * @description 执行错误
         */
        let runErrors = () => {
          _this.loadingStatus    = false
          _this.showErrorsStatus = true//显示错误提示
          let duration           = 1500//提示框关闭毫秒
          setTimeout(() => {
            _this.showErrorsStatus = false//关闭错误提示
            _this.slideChunkLeft   = 0
            _this.$refs.refSlideChunk.addEventListener('mousedown', _this.SlideMousedown)//添加监听器，让用户可用重新滑动
          }, duration)
        }
        slide.checkSlideChunkPosition({slideChunkLeft: slideChunkLeft, seed: Seed}).then(res => {
          if (typeof _this.parentCallback === 'function') {
            _this.dialogVisible = false
            _this.loadingStatus = false
            _this.parentCallback()//执行回调函数
          }
        }).catch((e) => {
          if ([4001, 4002].includes(e.code)) {
            _this.refreshSlideChunk()//重新获取滑动验证码
          } else {
            runErrors()
          }

        })
      },
      /**
       * @description 移除监听器
       */
      removeMouseEventListener() {
        //移除监听器
        this.bodyContainer.removeEventListener('mousemove', this.SlideMousemove)
        this.bodyContainer.removeEventListener('mouseup', this.SlideMouseup)
      },

      /**
       * @description 鼠标按下
       * @param Event
       * @constructor
       */
      SlideMousedown(Event) {
        Event.stopPropagation()//停止广播
        //console.log('鼠标按下')
        this.startMouseDownX   = Event.clientX - this.$refs.refSlideChunk.getBoundingClientRect().left//计算鼠标按下滑动块的起始位置
        this.slideChunkMaxLeft = this.$refs.refHandleWrapper.offsetWidth - this.$refs.refSlideChunk.getBoundingClientRect().width;//计算滑动块滑动最大位置
        this.removeMouseEventListener()
        this.$refs.refHandleMoveBorder.style.borderColor = '#1991fa'//设置拖动边框颜色
        //添加监听器
        this.bodyContainer.addEventListener('mousemove', this.SlideMousemove)
        this.bodyContainer.addEventListener('mouseup', this.SlideMouseup)
      },
      /**
       * @description 鼠标抬起
       * @param Event
       * @constructor
       */
      SlideMouseup(Event) {
        Event.stopPropagation()//停止广播
        //console.log('鼠标抬起')
        this.$refs.refHandleMoveBorder.style.borderColor = '#e0303000'//设置拖动边框颜色
        this.removeMouseEventListener()//移除监听器
        this.checkSlideChunkPosition(this.slideChunkLeft, this.slide.seed)

      },
      SlideMousemove(Event) {
        Event.stopPropagation()//停止广播
        let position = Event.clientX - this.startMouseDownX - this.$refs.refHandleWrapper.getBoundingClientRect().left//计算滑动块的位置
        position < 0 ? position = 0 : position > this.slideChunkMaxLeft && (position = this.slideChunkMaxLeft);
        this.slideChunkLeft = position
      }
    },

    computed: {
      verImg() {
        return this.slide.verImg
      },
      verIfy() {
        return this.slide.verIfy
      },
      verIfyVertical() {
        return this.slide.verIfyVertical
      },
      bodyContainer() {
        return document.getElementsByTagName('body')[0]//获取body容器
      }
    }
  }
</script>

<style lang="scss" scoped>
  .box {
    position: relative;
    width: 320px;

    .bgBox {
      overflow: hidden;
      position: relative;
      border-style: none;

      .verImg {
        width: 100%;
        -webkit-user-drag: none;
        height: 180px;
        border-radius: 5px;
        user-select: none;
      }

      .verify {
        position: absolute;
        -webkit-user-drag: none;
        left: 0;
        top: 0;
        z-index: 10;
      }

      .refresh {
        color: #ececec;
        cursor: pointer;
        position: absolute;
        right: 10px;
        top: 10px;
        font-size: 25px;
        z-index: 35;
      }

      .refresh:hover {
        color: #fff;
        transform: scale(1.1, 1.1);
        transition: all 200ms;
      }
    }

    .handle {
      position: relative;
      display: flex;
      margin-top: 10px;
      width: auto;
      text-align: center;
      height: 40px;
      line-height: 40px;
      background-color: #fff;
      border-radius: 5px;
      border: 1px solid #e4e7eb;
      -webkit-box-align: center;
      align-items: center;
      user-select: none;

      .handleMoveBorder {
        position: absolute;
        left: 0;
        top: 0;
        width: 0;
        height: 100%;
        border: 1px solid #e0303000;
        background: #D1E9FE;
        border-radius: 2px
      }

      .slideChunk {
        position: absolute;
        top: 0;
        left: 0;
        text-align: center;
        cursor: pointer;
        background-color: #fff;
        border: 1px solid #1991fa;
        border-radius: 2px;
        box-shadow: 0 0 3px rgb(0 0 0 / 30%);
        transition: background .2s linear;
        width: 40px;
        height: 100%;
        z-index: 30;

        i {
          font-size: 16px;
          color: #2b2f3a;
        }
      }

      .slideChunk:hover {
        background-color: #1991FA;

        i {
          color: #fff;
        }
      }

      .slideChunkErrors {
        background-color: rgb(224, 126, 122);
        border: 1px solid rgb(215, 96, 91);

        i {
          color: #fff;
        }
      }

      .slideChunkErrors:hover {
        background-color: rgb(224, 126, 122);

        i {
          color: #fff;
        }
      }

      .handleMoveBorderErrors {
        background-color: rgb(248, 160, 155);
        border: 1px solid rgb(248, 160, 155);

      }


      .message {
        margin: auto;
        text-align: center;
        line-height: 40px;
        color: #727272;
      }

    }
  }
</style>
