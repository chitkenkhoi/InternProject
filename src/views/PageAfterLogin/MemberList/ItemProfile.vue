<template>
    <div
      class="item"
      :class="{
        first: itemProp.id === 1,
        middle: itemProp.id !== 1 && itemProp.id!== 6,
        last: itemProp.id ===6,
      }"
    >
      <div class="tittle">{{ itemProp.tittle }}</div>

      <div class="data" v-if="itemProp.id !== 1 && itemProp.id!==6">{{ itemProp.data }}</div>
      <img class="img" v-else-if="itemProp.id === 1" :src ="`${itemProp.data}`" />
      <div class = "role" v-else-if ="itemProp.id ===6"> {{itemProp.data}}</div>

      <button v-if="itemProp.id !== 6" class="btn" @click="show">
        <button-edit />
      </button>
    </div>
  </template>

<script>
import ButtonEdit from '@/views/PageAfterLogin/Company/ButtonEdit.vue';
export default {
name: "ItemProfile",
props: ['itemProp'],
components :{ButtonEdit},
setup(props,context){
    const show = ()=>{
        context.emit('pop-up',props.itemProp.id)
    }
return {show}

}
}
</script>

<style lang = "scss" scoped>
.item{
    display: flex;
    flex-direction: row;
    align-items: center;
    .tittle{
        font-size:12px;
        color:#564C7E;
        font-family:'Noto Sans JP';
        font-weight: 400;
        margin-left:20px;
    }
    img{
        height:40px;
        width: 40px;
        border-radius:50%;
        position:absolute;
        left: 233px;
    }
    .data{
        position:absolute;
        left: 233px;
        font-size: 14px;
        font-family: 'Noto Sans JP';
        font-weight:400px;

    }
    .role {
    width: 150px;
    height: 30px;
    position:absolute;
    left: 233px;
    text-align: center;
    border-radius: 30px;
    border: 2px solid lightgrey;
    font-size: 12px;
    display: flex;
    align-items: center;
    justify-content: center;
    font-family: 'Noto Sans JP';
    font-weight:400;
  }
    .btn {
    border: none;
    background: none;
    cursor: pointer;
    position:absolute;
    left: 821px;
  }

}
.first{
    width: 900px;
    height: 80px;
    border-top:3px solid black;
}
.middle{
    width: 900px;
    height: 60px;
    border-top:1px solid lightgrey;
    
}
.last{
    width: 900px;
    height: 80px;
    border-top:1px solid lightgrey;
}

</style>