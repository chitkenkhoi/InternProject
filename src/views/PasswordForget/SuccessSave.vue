<template>
    <div class = "container">
      <div class = h1>NEW PASSWORD</div>
      <div class = h2>パスワードの再設定</div>
      <img :src = "success"/>
      <div class = line>新しいパスワードを使用してログインしてください</div>
      <Button class = "btn" @click = "backtoLogin">再設定</Button>
    </div>
    
    
   
  </template>
  
  <script>
  import success from '../../../public/icon/success.png'
  import Button from "@/components/Button.vue";

  export default {
    name: "PasswordNew",
    components: {Button },
    setup() {
      
      return {success };
    },
    methods:{
        backtoLogin(){
            this.$emit('back-to-login')
        }
    }
  };
  </script>
  
  <style lang="scss" scoped>
  .container{
      width:600px;
      height:425px;
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
      img{
        width:60px;
        height:80px;
        margin-top:50px;
      }
      .line{
        margin-top:20px;
        color:#666666;
        font-size:14px;
        font-weight: 400;
      }
      .btn{
          margin-top:50px;
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
  
  }
  .check{
  position:absolute;
  top:180px;
  left:50px;
  }
  </style>