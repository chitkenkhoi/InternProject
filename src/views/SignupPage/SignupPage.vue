<template>
  <div class = page>
  <div class="container">
    <div class="h1">CREATE ACCOUNT</div>
    <div class="h2">アカウントを作成する</div>
    <keep-alive><div class="box" v-if="state.current === 1">
      <input type="text" placeholder="メールアドレス" class="inputbox" v-model = "state.email"/>
      <div class = "element">
        <input
          type="password"
          placeholder="パスワード"
          class="inputbox"
          v-model="state.password"
        />
      </div>
      <input
        type="password"
        placeholder="パスワードを再入力する"
        class="inputbox"
        v-model = "state.recheck"
      />
    
      <div v-if="state.recheck !== state.password" class = "alert">パスワードが違います</div>
      <Button class="signupbutton" @click="signup">作成する</Button>
    </div>
  </keep-alive>
    <div v-if="state.current === 2" class = "box2">
      <img :src = "mailSuccess"/>
      <div class = "success">
        確認メ <div class = "mail">{{state.email}}</div> <div class = "margin"> に送信されました。受信トレイを確認し,</div>
        
      </div>
    </div>

    <div class="or">または</div>
    <div class="XID">
      <img :src="XID" />
      <div class="buttonText">xIDでログインする</div>
    </div>
    <div class="boxtext">
      xIDでログインに問題がありますか？よくある質問をご覧ください。 App
      StoreまたはGoogle Playを使用してxIDをインストールできます。
    </div>
    <div class="logoList">
      <img :src="Appstore" class="logo1" />
      <img :src="GGplay" class="logo2" />
    </div>
  </div>
  <div class = "check_password"><check-password v-if="state.password && !(check_character()&&check_length()&&check_number())" :check_character="check_character()" :check_number="check_number()" :check_length="check_length()"/></div>
  <div class="secondcontainer">
    <div>既にアカウントをお持ちの方は</div>
    <router-link to="/auth/login" class = "routerlink">こちら</router-link>
  </div>
  
  
</div>

</template>

<script>
import Button from "@/components/Button.vue";
import mailSuccess from "../../../public/logo/Group 7785.png"
import XID from "../../../public/logo/XID.png";
import Appstore from "../../../public/logo/Appstore.png";
import GGplay from "../../../public/logo/ggplay.png";
import checkPassword from "./checkPassword.vue"; 
import { reactive } from "vue";
export default {
  name: "SignupPage",
  components: { Button,checkPassword },
  setup() {
    const state = reactive({password: '',recheck:'',current:1,email:''})
   
    return { XID, Appstore, GGplay ,state,mailSuccess};
  },
  methods: {
    signup() {
      //API
      if( (this.state.password && this.state.email &&this.state.password === this.state.recheck)) this.state.current = 2
    },
    check_character() {
      const rex = /[A-Za-z]/;
      if (this.state.password.match(rex)) return true;
      else return false;
    },
    check_number() {
      const rex = /[0-9]/;
      if (this.state.password.match(rex)) return true;
      else return false;
    },
    check_length() {
      if (this.state.password.length >= 8) return true;
      else return false;
    },
  },
};
</script>

<style lang = "scss" scoped>
.page{
  display:flex;
  flex-direction: column;
  .container {
  position: absolute;
  top: 60px;
  left: 450px;
  width: 600px;
  height: 750px;
  background-color: white;
  border-radius: 5px;
  font-family: "Noto Sans JP";
  font-weight: 400;
  display: flex;
  flex-direction: column;
  align-items: center;
  .h1 {
    margin-top: 50px;
    color: #564c7e;
    font-size: 12px;
  }
  .h2 {
    font-size: 20px;
    color: #564c7e;
    margin-top: 5px;
  }
  .box {
    margin-top: 50px;
    width: 500px;
    height: 260px;
    display: flex;
    flex-direction: column;
    justify-content: space-between;
    input {
      padding: 15px;
      height: 20px;
      width: 470px;
      border-radius: 5px;
      border: solid lightgrey 1px;
    }
    input:focus {
      outline: none;
    }
    .alert{
      position:absolute;
      top:341px;
      font-size:12px;
      color:#aaaaaa
    }
  
      
    
    .signupbutton {
      display: flex;
      justify-content: center;
      align-items: center;
      color: white;
      height: 50px;
      width: 500px;
      font-size: 14px;
    }
  }
  .box2{
    display:flex;
    flex-direction: column;
    width:500px;
    margin-top:67px;
    align-items:center;
    img{
      margin-top:25px;
      width:180px;
      height:100px;
    }
    .success{
      margin-top:97px;
    
      display:flex;
      justify-content: center;
      
      width:500px;
      font-size:14px;
      .mail{
        color:#17B182;
        margin-left:5px;
      }
      .margin{
        margin-left:5px;
      }
    }
  }
  .or {
    margin-top: 50px;
    color: #aaaaaa;
    font-size: 14px;
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
  .boxtext{
    width:500px;
    height: 50px;
    text-align:center;
    font-size:14px;
    color:#666666;
    margin-top:10px;
  }
  .logoList{
    margin-top: 30px;
    .logo2{
      margin-left: 30px;
    }
    
  }
}
.secondcontainer{
  position:absolute;
  left: 450px;
  top:820px;
  width:600px;
  height:60px;
  text-decoration: none;
  background-color:white;
  border-radius:5px;
  display:flex;
  align-items: center;
  justify-content: center;
  font-family:'Noto Sans JP';
  font-weight:400;
  .routerlink{
    text-decoration: none;
    color:#D64D10;
  }
  
}
.check_password{
        position: absolute;
        top:310px;
        left:500px;
        z-index:100;
      }

}

</style>