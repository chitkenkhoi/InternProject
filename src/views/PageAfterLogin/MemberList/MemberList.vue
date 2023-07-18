<template>
  <div class="mid">
    <div
      class="head"
      id="h1"
      @click="state.current = 'MemberInternal'"
      :class="{ show: state.current === 'MemberInternal' }"
    >
      <div class="text">メンバー一覧 ({{ state.data_internal.length }})</div>
    </div>
    <div
      class="head"
      id="h2"
      @click="state.current = 'MemberPending'"
      :class="{ show: state.current === 'MemberPending' }"
    >
      <div class="text">招待中 ({{ state.data_pending.length }})</div>
    </div>
  </div>
 <keep-alive><component
    :is="state.current"
    class="body"
    :dataPropInternal="state.data_internal"
    @onDeleteInternal="onDeleteInternal"
    :dataPropPending="state.data_pending"
    @onDeletePending="onDeletePending"
  ></component></keep-alive>

  <!-- <keep-alive
    >
    <div>
      <MemberInternal
        class="body"
        id="body1"
        v-if="state.current == '1'"
        :dataProp="state.data_internal"
        v-on:onDelete="onDeleteInternal"
      />
      <MemberPending
        class="body"
        id="body2"
        v-else
        :dataProp="state.data_pending"
        v-on:onDelete="onDeletePending"
      /></div>
  </keep-alive> -->
  <Button class="btnadd" @click = "state.pop = 1">
    <div class="inner">
      <span class="material-symbols-outlined"> add </span>
      <div class="textBtn">メンバーを招待する</div>
    </div>
  </Button>
  <RoleAdd v-if = "state.pop == 1" @no-pop-up="state.pop = 0" @next="state.pop = 2"/>
  
</template>

<script>
import { reactive } from "vue";

import Button from "@/components/Button.vue";
import MemberInternal from "./MemberInternal.vue";
import MemberPending from "./MemberPending.vue";
import RoleAdd from '@/popup/MemberAdd/RoleAdd.vue'
export default {
  name: "MemberList",
  components: { Button, MemberInternal, MemberPending,RoleAdd },
  setup() {
    const state = reactive({
      current: "MemberInternal",
      pop: 0,
      data_internal: [
        {
          id: 1,
          name: "公的　太郎　",
          email: "taro_yamada@gmail.com",
          status: "マネージャー",
          twoFactorAuthentication: false,
        },
        {
          id: 2,
          name: "abc　",
          email: "abc@gmail.com",
          status: "マネージャー",
          twoFactorAuthentication: false,
        },
        {
          id: 3,
          name: "公的　太郎　",
          email: "taro_yamada@gmail.com",
          status: "マネージャー",
          twoFactorAuthentication: false,
        },
        {
          id: 4,
          name: "公的　太郎　",
          email: "taro_yamada@gmail.com",
          status: "マネージャー",
          twoFactorAuthentication: false,
        },
        {
          id: 5,
          name: "公的　太郎　",
          email: "taro_yamada@gmail.com",
          status: "マネージャー",
          twoFactorAuthentication: false,
        },
        {
          id: 6,
          name: "公的　太郎　",
          email: "taro_yamada@gmail.com",
          status: "マネージャー",
          twoFactorAuthentication: false,
        },
        {
          id: 7,
          name: "公的　太郎　",
          email: "taro_yamada@gmail.com",
          status: "マネージャー",
          twoFactorAuthentication: false,
        },
        {
          id: 8,
          name: "公的　太郎　",
          email: "taro_yamada@gmail.com",
          status: "マネージャー",
          twoFactorAuthentication: false,
        },
        {
          id: 9,
          name: "公的　太郎　",
          email: "taro_yamada@gmail.com",
          status: "マネージャー",
          twoFactorAuthentication: false,
        },
        {
          id: 10,
          name: "公的　太郎　",
          email: "taro_yamada@gmail.com",
          status: "マネージャー",
          twoFactorAuthentication: false,
        },
        {
          id: 11,
          name: "Le Quang Khoi　",
          email: "chitkenkhoi@gmail.com",
          status: "マネージャー",
          twoFactorAuthentication: false,
        },
      ],
      data_pending: [
        {
          id: 1,
          name: "公的　太郎　",
          email: "taro_yamada@gmail.com",
          status: "マネージャー",
          twoFactorAuthentication: false,
        },
        {
          id: 2,
          name: "公的　太郎　",
          email: "taro_yamada@gmail.com",
          status: "マネージャー",
          twoFactorAuthentication: false,
        },
        {
          id: 3,
          name: "公的　太郎　",
          email: "taro_yamada@gmail.com",
          status: "マネージャー",
          twoFactorAuthentication: false,
        },
        {
          id: 4,
          name: "公的　太郎　",
          email: "taro_yamada@gmail.com",
          status: "マネージャー",
          twoFactorAuthentication: false,
        },
        {
          id: 5,
          name: "公的　太郎　",
          email: "taro_yamada@gmail.com",
          status: "マネージャー",
          twoFactorAuthentication: false,
        },
        {
          id: 6,
          name: "公的　太郎　",
          email: "taro_yamada@gmail.com",
          status: "マネージャー",
          twoFactorAuthentication: false,
        },
        {
          id: 7,
          name: "公的　太郎　",
          email: "taro_yamada@gmail.com",
          status: "マネージャー",
          twoFactorAuthentication: false,
        },
        {
          id: 8,
          name: "公的　太郎　",
          email: "taro_yamada@gmail.com",
          status: "マネージャー",
          twoFactorAuthentication: false,
        },
        {
          id: 9,
          name: "公的　太郎　",
          email: "taro_yamada@gmail.com",
          status: "マネージャー",
          twoFactorAuthentication: false,
        },
        {
          id: 10,
          name: "公的　太郎　",
          email: "taro_yamada@gmail.com",
          status: "マネージャー",
          twoFactorAuthentication: false,
        },
        {
          id: 11,
          name: "Le Quang Khoi　",
          email: "chitkenkhoi@gmail.com",
          status: "マネージャー",
          twoFactorAuthentication: false,
        },
      ],
    });

    const onDeleteInternal = (id) => {
      state.data_internal = state.data_internal.filter((ele) => ele.id !== id);
      console.log("hêlo");
    };
    const onDeletePending = (id) => {
      state.data_pending = state.data_pending.filter((ele) => ele.id !== id);
    };
    return { state, onDeleteInternal, onDeletePending };
  },
};
</script>

<style lang = "scss" scoped>
.mid {
  position: fixed;
  display: flex;
  top: 76px;
  left: 270px;
  height: 50px;
  .head {
    color: #aaaaaa;
    height: 50px;
    cursor: pointer;
    display: flex;
    align-items: center;
  }
  #h1 {
    margin-left: 30px;
  }
  #h2 {
    margin-left: 51px;
  }
  .show {
    border-bottom: 3px solid #d64d10;
    color: #271b5a;
  }
}
.btnadd {
  position: absolute;
  right: 300px;
  top: 30px;
  width: 240px;
  height: 50px;
  .inner {
    display: flex;
    height: 20px;
    width: 210px;
    color: white;
    .material-symbols-outlined {
      margin-left: 5px;
    }
    .textBtn {
      margin-left: 25px;
      font-size: 14px;
      font-weight: 500;
    }
  }
}
</style>