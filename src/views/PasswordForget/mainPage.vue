<template>
  <div class = "container" v-if="state.current === 1">
    <div class = "h1">FORGET PASSWORD</div>
    <dv class = "h2">パスワードのリセット</dv>
    <div class = "line">登録したメールアドレスを入力してください。</div>
    <input type = text placeholder = "メールアドレス" v-model = "state.email"/>
    <div class ="btnlist">
        <div class = "isactive" @click = "backtoLogin">キャンセル</div>
        <div class = "inactive" :class = "{'isactive' : state.email.match(rex) } " @click = "next">リセットする</div>
    </div>
  </div>
  <Confirming class = "confirm" :emailProp="state.email" v-if = "state.current === 2"/>
  
</template>

<script>
import Confirming from './Confirming.vue'
import {reactive} from 'vue'
export default {
name: 'mainPage',
components:{Confirming} ,
setup(){
    const state = reactive({email: '',password: '',current:1})
    const rex = /[0-9A-Za-z]+@gmail/
    return {state,rex}
},
methods:{
    backtoLogin(){
        this.$router.push({ name: "Login" });
    },
    next(){
        if(this.state.email.match(this.rex)){
            this.state.current = 2
        }
    }
}

}
</script>

<style lang = "scss" scoped>
.container{
    width:600px;
    height:395px;
    position: absolute;
    top:200px;
    left:450px;
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
    .line{
        margin-top:50px;
        color:#666666;
        font-size:14px;
        font-weight: 400;

    }
    input{
        height:20px;
        width:470px;
        padding:15px;
        border-radius:5px;
        border: solid lightgrey 1px;
        margin-top:50px;
    }   
    input:focus{
        outline:none;
    }
    .btnlist{
        margin-top: 20px;
        display:flex;
        flex-direction: row;
        justify-content: space-between;
        align-items: center;
        height:50px;
        width:500px;
        font-size:14px;
        
        .inactive{
            height:50px;
            width:245px;
            background-color: #F0F0F0;
            color:#AAAAAA;
            border:none;
            border-radius:5px;
            display:flex;
            justify-content: center;
            align-items: center;
        }
        .isactive{
            height:50px;
            width:245px;
            border:solid 2px #D64D10;
            background-color: white;
            color:#D64D10;
            display:flex;
            justify-content: center;
            align-items: center;
            border-radius:5px;
            cursor:pointer;
        }
    }
}
.confirm{
    position:absolute;
    top: 210px;
    left:450px;
}

</style>