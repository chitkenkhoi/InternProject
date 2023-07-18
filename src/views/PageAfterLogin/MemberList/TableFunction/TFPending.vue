<template>
  <span class="material-icons"> arrow_drop_up </span>
  <div class="dropdown">
    <div class="item" id="item1" @click = "viewProfile">
      <img :src="iconview" />
      <div class="text">プロフィールを見る</div>
    </div>
    <div class="item" id="item2" @click = "popupresend">
      <img :src="iconresend" />
      <div class="text">招待メールを再送する</div>
    </div>
    <div class="item" id="item3" @click = "popupcancel">
      <img :src="iconremove" />
      <div class="text">招待を取り消す</div>
    </div>
  </div>
</template>

<script>
import iconview from "../../../../../public/icon/icon_view.png";
import iconresend from "../../../../../public/icon/icon_resend.png";
import iconremove from "../../../../../public/icon/icon_remove.png";
export default {
  name: "TF2",
  setup(_,context) {
    const popupresend = ()=>{
        context.emit("popup-resend")
    }
    const popupcancel = ()=>{
      context.emit("popup-cancel")
    }
    const viewProfile = ()=>{
      context.emit('view-profile')
    }
    return { iconremove, iconresend, iconview, popupresend, popupcancel,viewProfile };
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
  height: 150px;
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
    margin-top: 15px;
  }
  #item3 {
    color: #d00d41;
  }
}
</style>