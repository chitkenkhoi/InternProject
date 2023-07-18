<template>
  <div class = "button">
    <img class = "btn" :src = "edit">
    <div class = "text">編集</div>
  </div>
</template>

<script>
import edit from '../../../../public/icon/edit.png'
export default {
name: "ButtonEdit",
setup(){
  return {edit}
}

}
</script>

<style lang = "scss" scoped>
.button{
    display: flex;
    color: orangered;
    // &:hover{
    //         color: orange;
    //     }
    .btn{
      margin-right:5px;
    }
    .text{
      margin-left:5px;
      font-size:14px;
    }
}
</style>