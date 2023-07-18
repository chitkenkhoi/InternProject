<template>
  <div class="ele">
    <div class="tittle" :class = "{'margin':turnProp.id === 10}">{{ turnProp.tittle }}</div>
    <div class="data">{{ turnProp.data }}</div>
    <div class="btn">
      <div class="bg" :class="{ after1: turnProp.db }" @click = "toSwitch">
        <div
          class="circle"
          :class="{ after2: turnProp.db }"
          
        ></div>
      </div>
    </div>
  </div>
</template>

<script>
import { ref } from "vue";
export default {
  name: "EleMent",
  props: ["turnProp"],
  setup(props) {
    const a = props.turnProp.db === true ? ref(true) : ref(false);
    return { a };
  },
  methods: {
    async toSwitch() {
      if (this.turnProp.db === true) {
        this.$emit("switch", this.turnProp.id);
      } else {
        await this.$emit("pop-up", this.turnProp.id);
      }
    },
  },
};
</script>

<style lang = "scss" scoped>
.ele {
  display: flex;
  align-items: center;
  height: 50px;
  width: 900px;
  border-top: 5px solid black;
  margin-bottom:50px;
  .tittle {
    color: rgb(72, 61, 139);
    margin-left: 15px;
    font-size: 12px;
    width:150px;
    
  }
  .margin{
      margin-top:28px;
    }
  .data {
    position: absolute;
    left: 250px;
    font-size:14px;
  }
  .btn{
    position: absolute;
    left:870px;
    .bg {
    height: 17px;
    width: 30px;
    background-color: lightgray;
    border-radius: 17px;
    display:flex;
    align-items: center;
    cursor: pointer;
    }
  .circle {
    position: relative;
    cursor: pointer;
    background-color: #ccc;
    -webkit-transition: 0.4s;
    transition: 0.4s;
    position: relative;
    content: "";
    height: 12px;
    width: 12px;
    top: 0px;
    left: 3px;
    border-radius: 50%;
    background-color: white;
  }
  .after1 {
    background-color: rgb(247, 84, 25);
  }
  .after2 {
    left: 15px;
  }
  }
}
</style>
