<template>
  <div class="page">
    <input
      type="text"
      class="searchbar"
      placeholder="メンバーを検索する"
      v-model="search"
      @keyup.enter="Search"
    /><span class="material-symbols-outlined"> search </span>
    <div class="list">
      <div class="header">
        <div id="h1">メンバー氏名</div>
        <div id="h2">アカウントの種類</div>
      </div>
      <Item
        v-for="mem in state.view"
        v-bind:key="mem.id"
        :viewProp="mem"
        :isInternal="true"
        :show="state.show"
        v-on:turn-off="turnoff"
        v-on:turn-on="turnon"
        v-on:popup-2fa="popup2fa"
        v-on:popup-block="popupblock"
        v-on:popup-delete="popupdelete"
        @view-profile-internal="viewProfile"
      />
      <re-auth
        v-if="state.popup == 1"
        v-on:no-pop-up="nopop"
        v-on:to-reset="reset2fa"
      />
      <mem-block
        v-if="state.popup == 2"
        v-on:no-pop-up="nopop"
        v-on:to-block="block"
        :memberProp="
          state.show
            ? state.view.filter((ele) => ele.id === state.show)[0]
            : none
        "
      />
      <mem-delete
        v-if="state.popup == 3"
        v-on:no-pop-up="nopop"
        v-on:to-delete="del"
        :memberProp="
          state.show
            ? state.view.filter((ele) => ele.id === state.show)[0]
            : none
        "
      />
      <div class = "nodata" v-if="state.view.length === 0">
        <div class = "line" id = "line1">現在、このフォルダにはメンバーがいません。</div>
        <div class = "line" id = "line2">ボタン「メンバーを招待する」を選択して、新しいメンバーを追加してください。</div>
      </div>
    </div>
  </div>
</template>

<script>
import Item from "./Item.vue";
import { reactive } from "vue";
import ReAuth from "@/popup/Memlist/ReAuth.vue";
import MemBlock from "@/popup/Memlist/MemberBlock.vue";
import MemDelete from "@/popup/Memlist/MemberDelete.vue";
export default {
  name: "MemOne",
  components: { Item, ReAuth, MemBlock, MemDelete },
  props: ["dataPropInternal"],
  setup(props) {
    const state = reactive({
      search: "",
      view: props.dataPropInternal,
      show: 0,
      popup: 0,
    });
    return { state };
  },
  methods: {
    Search() {
      //API
    },
    turnoff() {
      this.state.show = 0;
    },
    turnon(id) {
      this.state.show = id;
    },
    popup2fa() {
      this.state.popup = 1;
    },
    reset2fa() {
      //API

      this.state.popup = 0;
      this.state.show = 0;
    },
    popupblock() {
      this.state.popup = 2;
    },
    block() {
      //API
      this.state.view = this.state.view.filter(
        (ele) => ele.id !== this.state.show
      );
      this.$emit('onDeleteInternal',this.state.show)
      this.state.popup = 0;
      this.state.show = 0;
    },
    popupdelete() {
      this.state.popup = 3;
    },
    del() {
      //API

      this.state.popup = 0;
      this.state.show = 0;
    },
    nopop() {
      this.state.popup = 0;
    },
    viewProfile(){
      var path = '/profile.admin/Internal/' + JSON.stringify(this.state.show)
      this.$router.push({path: path})
    }
  },
};
</script>

<style lang = "scss" scoped>
.page {
  .material-symbols-outlined {
    position: fixed;
    height: 20px;
    width: 20px;
    top: 173px;
    left: 1160px;
  }
  .searchbar {
    position: absolute;
    top: 30px;
    left: 0px;
    height: 50px;
    width: 893px;
    font-size: 14px;
    border-radius: 5px;
    border: 1px lightgrey solid;
  }
  .list {
    display: flex;
    flex-direction: column;
    margin-top: 100px;
    .header {
      display: flex;
      flex-direction: row;
      font-size: 12px;
      color: #564c7e;
      width: 900px;
      border-bottom: 3px solid black;
      #h1 {
        margin-left: 20px;
      }
      #h2 {
        margin-left: 215px;
      }
    }
    .nodata{
      display:flex;
      flex-direction: column;
      align-items: center;
      margin-top: 50px;
      width: 900px;
      height:47px;
      font-family: 'Noto Sans JP';
      font-weight: 400;
      font-size: 14px;
      color: #666666;
    }
  }
}
</style>