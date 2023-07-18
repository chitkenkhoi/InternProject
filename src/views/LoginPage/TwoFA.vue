<template>
    <div class = "container">
        <div class = h1>認証</div>
        <div class = h2>二段階認証</div>
        <div class = "text">認証システムアプリを起動し、6桁の認証コードをご確認の上、画面に表示されている入力項目にご入力ください。</div>
        <input type = text placeholder="認証コード" v-model="code"/>
        <div class = 'btnlist'>
            <div class = "isactive" @click = "backtoLogin">キャンセル</div>
            <div class = "btn2" :class = "{'isactive' : code.match(rex) && code.length === 6} " @click = "next">認証する</div>
        </div>
    </div>
  
</template>

<script>
import {ref} from 'vue'
export default {
    name:"TwoFA",
    setup(){
        const code = ref('')
        const rex = /[0-9][0-9][0-9][0-9][0-9][0-9]/
        return {code,rex}
    },
    methods:{
        backtoLogin(){
            this.$router.push({ name: "Login" });
        },
        next(){
            //API
            if(this.code.match(this.rex) && this.code.length === 6)
            {this.$router.push({ name: "LayoutPage" });}
            else 
            return
        }
    }


}
</script>

<style lang = "scss" scoped>
.container{
    width:600px;
    height:423px;
    background-color: white;
    border-radius:5px;
    position:absolute;
    top: 200px;
    left: 450px;
    display:flex;
    flex-direction: column;
    align-items: center;
    font-family:'Noto Sans JP';
    font-weight: 400;
    .h1{
        font-size:12px;
        color:#564C7E;
        margin-top:50px;

    }
    .h2{
        margin-top:5px;
        color:#564C7E;
        font-size:20px;
    }
    .text{
        width:500px;
        height:48px;
        display:flex;
        justify-content: center;
        align-items: center;
        font-size:14px;
        color:#666666;
        margin-top:50px;
    }
    input{
        margin-top:50px;
        width:470px;
        height:20px;
        padding:15px;
        border:solid lightgrey 1px;
        border-radius:4px;

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
        
        .btn2{
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

</style>