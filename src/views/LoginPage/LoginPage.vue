<template>
  <div class="container">
    <div class="head1">現場の声から生まれた電子契約サービス</div>
    <div class="head2">Digital Signにログイン</div>
    <input
      type="text"
      placeholder="メールアドレス"
      v-model="email"
      class="inputbox1"
    />
    <input
      type="text"
      placeholder="パスワード"
      v-model="password"
      class="inputbox2"
    />
    <Button class="signinbutton" @click="confirm2FA">ログイン</Button>
    <div class="forgetPassword">
      <div class = "text">パスワードを忘れた場合は</div>
      <router-link to = "/auth/password_forget" class = "here">こちら</router-link>
    </div>
    <div class="or">または</div>
    <div class="XID">
      <img :src="XID" />
      <div class="buttonText">xIDでログインする</div>
    </div>
  </div>
  <div class="secondContainer">
    <div class = "text">初めてDigital Signをご利用の方は</div>
    <router-link to="/auth/signup" class = "SignUp">こちら</router-link>
  </div>
</template>
  
  <script>
import { ref } from "vue";
import Button from "@/components/Button.vue";
import XID from "../../../public/logo/XID.png";
export default {
  name: "SignupPage",
  components: { Button },
  setup() {
    const email = ref("");
    const password = ref("");
    return { email, password, XID };
  },
  methods: {
    confirm2FA() {
      //API
      
      this.$router.push({ name: "TwoFA" });
    },
  },
};
</script>
  
<style lang="scss" scoped>
.container{
  position:absolute;
  top:140px;
  left:450px;
  background-color:white;
  width:600px;
  height:575px;
  display:flex;
  flex-direction: column;
  align-items: center;
  border-radius:5px;
  .head1{
    margin-top:50px;
    color:#564C7E;
    font-size:12px;
    font-family: 'Noto Sans JP';
    font-weight: 500;
  }
  .head2{
    color:#564C7E;
    font-size:20px;
    font-family: 'Noto Sans JP';
    font-weight: 500;
    margin-top:5px;
  }
  input{
    width:470px;
    height:20px;
    border-radius:3px;
    padding:15px;
    border:solid lightgrey 1px;
  }
  input:focus{
    outline:none;
  }
  .inputbox1{
    margin-top:50px;
    
  }
  .inputbox2{
    margin-top:20px;
    
  }
  .signinbutton{
    color:white;
    border-radius:5px;
    height:50px;
    width:500px;
    margin-top:20px;
    display:flex;
    justify-content: center;
    align-items: center;
  }
  .forgetPassword{
    display:flex;
    margin-top:10px;
    color:#666666;
    font-size:14px;
    font-family:'Noto Sans JP';
    font-weight: 400;
    .here{
      color:#D64D10;
      text-decoration: none;
    }
  }
  .or{
    margin-top:31px;
    color:#AAAAAA;
    font-size:14px;
  }
  .XID{
    height:50px;
    width: 500px;
    background-color: #33CEAA;
    cursor:pointer;
    color:white;
    display:flex;
    align-items: center;
    border-radius:5px;
    margin-top:50px;
    .buttonText{
      margin-left: 140px;
    }
  }
  
}
.secondContainer{
  width:600px;
  height:60px;
  display:flex;
  justify-content: center;
  align-items: center;
  border-radius:5px;
  position:absolute;
  background-color: white;;
  top:750px;
  left:450px;
  color:#666666;
  font-size:14px;
  font-weight: 400;
  font-family:'Noto Sans JP';

  .SignUp{
    text-decoration: none;
    color:#D64D10;
  }

}
</style>