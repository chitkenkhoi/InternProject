<template>
  <div class="item" :class="{ isMarked: show === viewProp.id }">
    <img class="logo" :src="logo" />
    <div class="box">
      <div class="name">
        {{ viewProp.name }}
      </div>
      <div class="email">
        {{ viewProp.email }}
      </div>
    </div>
    <div class="status">
      {{ viewProp.status }}
    </div>
    <div class="function">
      <span
        class="material-symbols-outlined"
        :class="{ isShow: show === viewProp.id }"
        @click="handle"
      >
        more_horiz
      </span>
      <TF1 v-if="show === viewProp.id && isInternal" v-on:popup-2fa = "popup2fa" v-on:popup-block="popupblock" v-on:popup-delete="popupdelete" @view-profile = "viewProfileInternal"> </TF1>
      <TF2 v-if="show === viewProp.id && !isInternal" v-on:popup-resend = "popupresend" v-on:popup-cancel = "popupcancel" @view-profile = "viewProfilePending"> </TF2>
    </div>
  </div>
</template>

<script>
import TF1 from "./TableFunction/TFInternal.vue";
import TF2 from "./TableFunction/TFPending.vue";
// import {ref} from 'vue'
import logo from "@/assets/logo.png";
export default {
  name: "MemBer",
  props: ["viewProp", "isInternal", "show"],
  components: {TF1, TF2},
  setup(props, context) {
    const handle = () => {
      if (props.show) {
        if (props.show === props.viewProp.id) {
          context.emit("turn-off");
        } else {
          context.emit("turn-on", props.viewProp.id);
        }
      } else {
        context.emit("turn-on", props.viewProp.id);
      }
    };
    const popup2fa = ()=>{
      context.emit('popup-2fa')
    }
    const popupblock = ()=>{
      context.emit('popup-block')
    }
    const popupdelete = ()=>{
      context.emit('popup-delete')
    }
    const popupresend = ()=>{
      context.emit('popup-resend')
    }
    const popupcancel = ()=>{
      context.emit('popup-cancel')
    }
    const viewProfileInternal = ()=>{
      context.emit('view-profile-internal')
    }
    const viewProfilePending = ()=>{
      context.emit('view-profile-pending')
    }
    return {logo, handle, popup2fa,popupblock,popupdelete,popupresend,popupcancel,viewProfileInternal,viewProfilePending};
  },
  methods: {},
};
</script>

<style lang = "scss" scoped>
.item {
  display: flex;
  width: 900px;
  height: 60px;
  align-items: center;
  border-bottom: 1px solid lightgrey;
  .logo {
    height: 30px;
    width: 30px;
    border-radius: 50%;
    margin-left: 20px;
  }
  .box {
    margin-left: 20px;
    display: flex;
    flex-direction: column;
    width: 200px;
    .name {
      font-size: 14px;
      height: 20px;
    }
    .email {
      font-size: 12px;
      height: 16px;
    }
  }
  .status {
    width: 150px;
    height: 30px;
    margin-left: 10px;
    text-align: center;
    border-radius: 30px;
    border: 2px solid lightgrey;
    font-size: 12px;
    display: flex;
    align-items: center;
    justify-content: center;
  }
  .function {
    height: 30px;
    width: 30px;
    margin-left: 420px;

    .material-symbols-outlined {
      &:hover {
        background-color: #f0f0f0;
        border-radius: 5px;
      }
      color: #d64d10;
      cursor: pointer;
    }
    .isShow {
      background-color: #f0f0f0;
      border-radius: 5px;
    }
  }
}
.isMarked {
  background-color: #f8f8f8;
}
</style>