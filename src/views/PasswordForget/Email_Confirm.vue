<template>
  <div class="container" v-if="state.current === 1">
    <div class="h1">EMAIL CONFIRMED</div>
    <div class="h2">メールアドレスが正常にが確認されました</div>
    <div class="boxtext">
      <div class="line">メールアドレスが確認されました。</div>
      <div class="line">以下のボタンよりパスワードの再設定に進んでください</div>
    </div>
    <Button class="btn" @click="changeEmail">パスワードを再設定する</Button>
  </div>
  <div class = "passwordnew"><password-new v-if="state.current === 2" @success-save = "onSuccessSave"/></div>
  <div class = "successsave"><success-save v-if="state.current === 3" @back-to-login="onBackToLogin"/></div>
  
</template>

<script>
import Button from "@/components/Button.vue";
import PasswordNew from "./PasswordNew.vue";

import { reactive } from "vue";
import SuccessSave from './SuccessSave.vue';
export default {
  name: "Email_Confirm",
  components: { Button, PasswordNew ,SuccessSave},
  setup() {
    const state = reactive({ email: "", pass: "", current: 1 });
    return { state };
  },
  methods: {
    changeEmail() {
      this.state.current = 2;
    },
    onSuccessSave(pass){
        //API
        this.state.pass = pass
        this.state.current = 3;

    },
    onBackToLogin(){
        this.$router.push({name:"Login"})
    }
  },
};
</script>

<style lang = "scss" scoped>
.container {
  position: absolute;
  top: 210px;
  left: 450px;
  height: 353px;
  width: 600px;
  background-color: white;
  font-family: "Noto Sans JP";
  border-radius: 5px;
  display: flex;
  flex-direction: column;
  align-items: center;
  .h1 {
    color: #564c7e;
    font-size: 12px;
    font-weight: 500;
    margin-top: 50px;
  }
  .h2 {
    margin-top: 5px;
    color: #564c7e;
    font-size: 20px;
    font-weight: 500;
  }
  .boxtext {
    margin-top: 50px;
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: space-between;
    font-size: 14px;
    font-weight: 400;
    color: #666666;
  }
  .btn {
    height: 50px;
    width: 500px;
    display: flex;
    justify-content: center;
    align-items: center;
    color: white;
    font-size: 14px;
    margin-top: 50px;
  }
}
.passwordnew {
    position: absolute;
  top: 210px;
  left: 450px;
}
</style>