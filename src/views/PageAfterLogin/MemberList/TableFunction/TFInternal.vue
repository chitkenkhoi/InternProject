<template>
  <span class="material-icons" > arrow_drop_up </span>
  <div class="dropdown">
    <div class="item" id="item1" @click = "viewProfile">
      <img :src="iconview" />
      <div class="text">プロフィールを見る</div>
    </div>
    <div class="item" id="item2" @click = "popup2fa">
      <img :src="icon2fa" />
      <div class="text">二段階認証をリセット</div>
    </div>
    <div class="item" id="item3" @click = "popupblock">
      <img :src="iconblock" />
      <div class="text">メンバーのブロック</div>
    </div>
    <div class="item" id="item4" @click = "popupdelete">
      <img :src="icondelete" />
      <div class="text">メンバーを削除</div>
    </div>
  </div>
</template>

<script>
import iconview from "../../../../../public/icon/icon_view.png";
import icon2fa from "../../../../../public/icon/icon_2fa.png";
import iconblock from "../../../../../public/icon/icon_block.png";
import icondelete from "../../../../../public/icon/icon_delete.png";
export default {
  name: "TF1",
  setup(_,context) {
    const popup2fa = ()=>{
      context.emit('popup-2fa')
    }
    const popupblock = ()=>{
      context.emit('popup-block')
    }
    const popupdelete = ()=>{
      context.emit('popup-delete')
    }
    const viewProfile = ()=>{
      context.emit('view-profile')
      console.log("hello1")
    }
    return { icon2fa, iconblock, icondelete, iconview, popup2fa, popupblock, popupdelete,viewProfile};
  },
};
</script>

<style lang = "scss" scoped>
.material-icons {
  position: relative;
  top: -13px;
  color: white;
  z-index: 91;
}
.dropdown {
  width: 250px;
  height: 200px;
  display: flex;
  flex-direction: column;
  border-radius: 5px;
  z-index: 90;
  background-color: white;
  position: relative;
  left: -227px;
  top: -28px;
  box-shadow: rgba(99, 99, 99, 0.2) 0px 0px 8px 0px;
  .item {
    height: 40px;
    width: 250px;
    display: flex;
    align-items: center;
    font-size: 14px;
    font-family: "Noto Sans JP";
    font-weight: 500;
    color: #d64d10;
    cursor: pointer;
    img {
      margin-left: 20px;
    }
    .text {
      margin-left: 10px;
    }
    &:hover {
      background-color: #f0f0f0;
    }
  }
  #item1 {
    margin-top: 20px;
  }
  #item4 {
    color: #d00d41;
  }
}
</style>