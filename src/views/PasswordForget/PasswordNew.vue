<template>
  <div class = "container">
    <div class = h1>NEW PASSWORD</div>
    <div class = h2>パスワードの再設定</div>
    <input type = password id = "margin" v-model = "state.pass" placeholder="新しいパスワード"/>
    <input type = password v-model = "state.newpass" placeholder = "新しいパスワードの再入力"/>
    <Button class = "btn" :class = "{inactive:state.pass === '' || state.newpass==='' || state.newpass !== state.pass || !(check_character()&&check_length()&&check_number())}" @click = "savePassword">再設定</Button>
  </div>
  <div class = "check"><check-password v-if="state.pass && !(check_character()&&check_length()&&check_number())" :check_character="check_character()" :check_number="check_number()" :check_length="check_length()"/></div>
  
 
</template>

<script>
import checkPassword from "../SignupPage/checkPassword.vue";
import Button from "@/components/Button.vue";
import { reactive } from "vue";
export default {
  name: "PasswordNew",
  components: { checkPassword,Button },
  setup() {
    const state = reactive({ pass: "", newpass: "" });
    return { state };
  },
  methods:{
    check_character() {
      const rex = /[A-Za-z]/;
      if (this.state.pass.match(rex)) return true;
      else return false;
    },
    check_number() {
      const rex = /[0-9]/;
      if (this.state.pass.match(rex)) return true;
      else return false;
    },
    check_length() {
      if (this.state.pass.length >= 8) return true;
      else return false;
    },
    savePassword(){
        if(this.state.pass === '' || this.state.newpass==='' || this.state.newpass !== this.state.pass || !(this.check_character()&&this.check_length()&&this.check_number())) return
        else {
            this.$emit('success-save',this.state.pass)
        }
    }
  }
};
</script>

<style lang="scss" scoped>
.container{
    width:600px;
    height:395px;
    background-color: white;
    border-radius:5px;
    display:flex;
    flex-direction: column;
    align-items: center;
    font-family:'Noto Sans JP';
    .h1{
        color: #564C7E;
        font-size:12px;
        font-weight: 500;
        margin-top: 50px;
    }
    .h2{
        margin-top:5px;
        color:#564C7E;
        font-size:20px;
        font-weight: 500;
    }
    input{
        margin-top:20px;
        padding:15px;
        height:20px;
        width:470px;
        border-radius:5px;
        border:solid lightgrey 1px;

    }
    #margin{
        margin-top:50px;
    }
    input:focus{
        outline:none;
    }
    .btn{
        margin-top:20px;
        display:flex;
        height:50px;
        width:500px;
        justify-content: center;
        align-items: center;
        color:white;
        border-radius:5px;
        font-size: 14px;
        font-weight: 500;
    }
    .inactive{
        background-color:#F0F0F0;
        color: #AAAAAA;
        cursor:auto;

    }

}
.check{
position:absolute;
top:180px;
left:50px;
}
</style>