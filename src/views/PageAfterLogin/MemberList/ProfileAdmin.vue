<template>
  <Profile class="profile">
    <div class="box">
      <div class="function" @click="popUp(1)">
        <img :src="iconOwner" />
        <div class="tittle">オーナー変更</div>
      </div>
      <div class="function" @click="popUp(2)">
        <img :src="icon2fa" />
        <div class="tittle">二段階認証をリセット</div>
      </div>
      <div class="function" @click="popUp(3)">
        <img :src="iconBlock" />
        <div class="tittle">メンバーのブロック</div>
      </div>
      <div class="function" id="last" @click="popUp(4)">
        <img :src="iconDelete" />
        <div class="tittle">メンバーを削除</div>
      </div>
    </div>
    <change-owner v-if="popup == 1" @no-pop-up="turnOff" :OwnerProp="OwnerList"/>
    <re-auth v-if="popup == 2" @no-pop-up="turnOff" @to-reset="turnOff" />
    <member-block
      v-if="popup == 3"
      @no-pop-up="turnOff"
      @to-block="turnOff"
      :memberProp="member" />
    <member-delete
      v-if="popup == 4"
      @no-pop-up="turnOff"
      @to-delete="turnOff"
      :memberProp="member"
  /></Profile>
</template>

<script>
import iconOwner from "../../../../public/icon/iconOwner.png";
import icon2fa from "../../../../public/icon/icon_2fa.png";
import iconBlock from "../../../../public/icon/icon_block.png";
import iconDelete from "../../../../public/icon/icon_delete.png";
import Profile from "./Profile.vue";
import ReAuth from "@/popup/Memlist/ReAuth.vue";
import MemberBlock from "@/popup/Memlist/MemberBlock.vue";
import MemberDelete from "@/popup/Memlist/MemberDelete.vue";
import { ref } from "vue";
import ChangeOwner from "@/popup/Profile/ChangeOwner.vue";
export default {
  name: "ProfileAdmin",
  components: { Profile, ReAuth, MemberBlock, MemberDelete, ChangeOwner },

  setup() {
    const popup = ref(0);
    const member = { name: "Le Quang Khoi", email: "abc@gmail.com" };
    const OwnerList = ref([
      {
        id: 1,
        name: "公的　太郎",
        email: "taro_yamada@gmail.com",
        company: "Company",
      },
      {
        id: 2,
        name: "公的　太郎",
        email: "taro_yamada@gmail.com",
        company: "Company",
      },
      {
        id: 3,
        name: "公的　太郎",
        email: "taro_yamada@gmail.com",
        company: "Company",
      },
      {
        id: 4,
        name: "公的　太郎",
        email: "taro_yamada@gmail.com",
        company: "Company",
      },
      {
        id: 5,
        name: "公的　太郎",
        email: "taro_yamada@gmail.com",
        company: "Company",
      },
      {
        id: 6,
        name: "公的　太郎",
        email: "taro_yamada@gmail.com",
        company: "Company",
      },
      {
        id: 7,
        name: "公的　太郎",
        email: "taro_yamada@gmail.com",
        company: "Company",
      },
    ]);
    return { iconOwner, icon2fa, iconBlock, iconDelete, popup, member,OwnerList };
  },
  methods: {
    turnOff() {
      this.popup = 0;
    },
    popUp(id) {
      this.popup = id;
      console.log(id);
    },
  },
};
</script>

<style lang = "scss" scoped>
.box {
  position: fixed;
  top: 150px;
  left: 1230px;
  display: flex;
  flex-direction: column;
  justify-content: space-between;
  width: 240px;
  height: 140px;
  z-index: 20;

  .function {
    width: 240px;
    height: 20px;
    display: flex;
    align-items: center;
    flex-direction: row;
    color: #d64d10;
    cursor: pointer;
    img {
      height: 20px;
      width: 20px;
    }
    .tittle {
      margin-left: 10px;
      font-family: "Noto Sans JP";
      font-size: 14px;
      font-weight: 400;
    }
  }
  #last {
    color: #d00d41;
  }
}
</style>