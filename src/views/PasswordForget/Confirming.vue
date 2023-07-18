<template>
  <div class = container>
    <div class = h1>FORGET PASSWORD</div>
    <div class = h2>パスワードのリセット</div>
    <div class = notification>
        <div class = line id = line><div class = "email">{{ emailProp }}</div>にメールを送信しました。</div>
        <div class = line>メールに記載の手順に従って再設定を行ってください。</div>
    </div>
    <Button class = "btn" @click = "backtoLogin">ログインに戻る</Button>

  </div>
</template>

<script>
import Button from '@/components/Button.vue';
export default {
name:"ConfirmPass",
props: ['emailProp'],
components: {Button},
methods:{
    backtoLogin(){
        this.$router.push({name: "Login"})
    }
}
}
</script>

<style lang = "scss" scoped>
.container{
    width:600px;
    height:353px;
    background-color:white;
    border-radius:5px;
    font-family:'Noto Sans JP';
    display:flex;
    flex-direction: column;
    align-items: center;
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
    .notification{
        margin-top:50px;
        width:500px;
        height:48px;
        color:#666666;
        font-size:14px;
        display:flex;
        flex-direction: column;
        align-items: center;
        #line{
            display:flex;

        }
        .email{
            color:#17B182;
        }
    }
    .btn{
        height:50px;
        width:500px;
        color:white;
        display:flex;
        justify-content: center;
        align-items: center;
        margin-top:50px;
    }

}

</style>