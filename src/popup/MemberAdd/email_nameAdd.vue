<template>
  <Popup @turn-off="turnoff">
    <div class="container">
      <div class="h">
        <div>メンバーを招待する</div>
        <div class="style">メンバー</div>
        <span class="material-symbols-outlined" @click="turnoff">close</span>
      </div>
      <div class="boxtext">
        <div>メンバーリスト</div>
        <div class="btnAdd" @click="add">
          <span class="material-symbols-outlined"> add </span>
          <div>メンバーを招待する</div>
        </div>
      </div>
      <div class="AddingList"></div>
      <div class="AddingInput">
        <input type="text" v-model="state.email" placeholder="メールアドレス" />
        <input type="text" v-model="state.name" placeholder="名前" />
      </div>
    </div>
  </Popup>
</template>
    
    <script>
import Popup from "../Popup.vue";
import { reactive } from "vue";
export default {
  name: "EmailNameAdd",
  components: { Popup },
  setup() {
    const state = reactive({
      AddingList: [],
      email: "",
      name: "",
      alert: false,
    });
    return { state };
  },
  methods: {
    turnoff() {
      this.$emit("no-pop-up");
    },
    next() {
      this.$emit("next");
    },
    add() {
      if (this.state.email || this.state.name) {
        this.state.alert = true;
      } else {
        var id = this.state.AddingList.length + 1;
        const member = {
          id: id,
          email: this.state.email,
          name: this.state.name,
        };
        this.state.AddingList.push(member);
        this.state.email = "";
        this.state.name = "";
      }
    },
    delete(id){
        this.state.AddingList = this.state.AddingList.filter((ele)=>ele.id !== id)
        var x = id-1
        var size = this.state.AddingList.length - 1
        while(x <= size){
            this.state.AddingList[x].id -=1
        }

    }
  },
};
</script>
    
    <style lang="scss" scoped>
.container {
  position: fixed;
  background-color: white;
  z-index: 99;
  top: 40px;
  left: 500px;
  width: 600px;
  height: 621px;
  border-radius: 5px;
  display: flex;
  flex-direction: column;
  align-items: center;
  .h {
    display: flex;
    flex-direction: row;
    justify-content: space-between;
    width: 500px;
    height: 30px;
    margin-top: 50px;
    color: #271b5a;
    font-size: 16px;
    font-family: "Noto Sans JP";
    font-weight: 500;
    .material-symbols-outlined {
      color: black;
      cursor: pointer;
      &:hover {
        opacity: 0.5;
      }
    }
  }
  .boxtext {
    width: 500px;
    height: 48px;
    display: flex;
    flex-direction: column;
    align-items: center;
    margin-top: 50px;
    font-size: 14px;
    font-family: "Noto Sans JP";
    font-weight: 400;
    justify-content: space-between;
    color: #666666;
  }
  .boxrole {
    margin-top: 39px;
    width: 500px;
    height: 373px;
    display: flex;
    flex-direction: column;
    justify-content: space-evenly;
    .role {
      width: 500px;
      height: 107px;
      border: 2px lightgrey solid;
      border-radius: 5px;
      display: flex;
      justify-content: space-evenly;
      align-items: center;
      .material-symbols-outlined {
        color: black;
        cursor: pointer;
        &:hover {
          opacity: 0.5;
        }
      }
      .box {
        font-size: 12px;
        width: 420px;
        height: 60px;
        display: flex;
        flex-direction: column;
        font-family: "Noto Sans JP";
        font-weight: 400;
        color: #666666;
        .head {
          color: #d64d10;
          font-size: 14px;
          font-family: "Noto Sans JP";
          font-weight: 500;
        }
        .firstline {
          margin-top: 5px;
        }
      }
    }
  }
}
</style>