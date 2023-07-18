<template>
  <div
    class="item"
    :class="{
      first: isFirst,
      mid: isMid,
      last: isLast,
      tachra: itemProp.popup,
    }"
  >
    <div class="tittle">{{ itemProp.tittle }}</div>
    <div class="data" v-if="itemProp.id !== 1 && itemProp.id!==9">{{ itemProp.data }}</div>
    <img class="img" v-else :src ="`${itemProp.data}`" />
    <button v-if="itemProp.popup" class="btn" @click="show">
      <button-edit />
    </button>
  </div>
</template>

<script>
import ButtonEdit from "./ButtonEdit.vue";
import { ref } from "vue";
export default {
  name: "CompanyItem",
  props: ["itemProp"],
  components: { ButtonEdit },
  setup(props, context) {
    const isFirst = ref(true);
    const isMid = ref(true);
    const isLast = ref(true);
    if (props.itemProp.id === 1) {
      isFirst.value = true;
      isMid.value = false;
      isLast.value = false;
    } else if (props.itemProp.id === 9) {
      isFirst.value = false;
      isMid.value = false;
      isLast.value = true;
    } else {
      isFirst.value = false;
      isMid.value = true;
      isLast.value = false;
    }
    const show = () => {
      context.emit("Pop-up", props.itemProp.id);
    };
    return { isFirst, isMid, isLast, show, };
  },
};
</script>

<style lang = "scss" scoped>
.item {
  display: flex;
  align-items: center;
  height: 50px;
  width: 900px;
  .tittle {
    color: rgb(72, 61, 139);
    margin-left: 15px;
    font-size: 12px;
  }
  .data {
    position: absolute;
    left: 250px;
    font-size:14px;
  }
  .img {
    position: absolute;
    left: 250px;
    border-radius: 50%;
    height: 40px;
    width: 40px;
  }
  .btn {
    border: none;
    background: none;
    cursor: pointer;
  }
}
.first {
  border-top: 5px solid black;
  border-bottom: 1px solid lightgrey;
  height: 70px;
  width: 900px;
}
.mid {
  border-bottom: 1px solid lightgrey;
}
.last {
  background-color: white;
  height: 70px;
  width: 900px;
}
.tachra {
  justify-content: space-between;
  
}
</style>