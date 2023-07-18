<template>
  <div class="mid">
    <div class="h1" @click="backToList">メンバー一覧</div>
    <span class="material-symbols-outlined" id="arow"> chevron_right </span>
    <div class="h2">メンバープロフィール</div>
  </div>
  <div class="profile">
    <div class="ProfileSetting">
      <div class="h">プロフィール設定</div>
      <ItemProfile
        v-for="item in fakedata"
        v-bind:key="item.id"
        :itemProp="item"
        @pop-up="popUp"
      />
      <UpdateAvatar v-if="state.pop === 1" @no-pop-up="turnOff" />
      <UpdateName
        v-if="state.pop === 2"
        @no-pop-up="turnOff"
        :popprop="fakedata[1]"
        @data-change="setData"
      />
      <UpdateRole
        v-if="state.pop === 3"
        @no-pop-up="turnOff"
        :popprop="fakedata[2]"
        @data-change="setData"
      />
      <UpdatePhoneNumber
        v-if="state.pop === 4"
        @no-pop-up="turnOff"
        :popprop="fakedata[3]"
        @data-change="setData"
      />
      <UpdateEmail
        v-if="state.pop === 5"
        @no-pop-up="turnOff"
        :popprop="fakedata[4]"
        @data-change="setData"
      />
      <AddSigner v-if="state.pop === 6" @no-pop-up="turnOff" />
      <UpdatePassword v-if="state.pop === 7" @no-pop-up="turnOff" />
    </div>
    <div class="SignerSetting">
      <div class="h">
        <div class="text">署名者の設定</div>
        <div class="AddSigner" @click="SignerPopup">
          <span class="material-symbols-outlined"> add </span>署名者を追加
        </div>
      </div>
      <div class="box">
        <div class="l1">
          アラート通知先のメールアドレスをデフォルト設定できます。
        </div>
        <div class="l2">
          また、他のメニューから契約書ごとに変更することもできます。
        </div>
      </div>
    </div>
    <div class="PasswordSetting">
      <div class="h">
        <div class="text">セキュリティ設定</div>
      </div>
      <div class="box2">
        <div class="tittle">パスワード</div>
        <img :src="HiddenPassword" />
        <button-edit class="Edit" @click="popupPassword" />
      </div>
    </div>
    <div class = admin><slot/> </div>
  </div>
</template>

<script>
import HiddenPassword from "../../../../public/icon/pwd_dot.png";
import UpdateAvatar from "@/popup/Profile/UpdateAvatar.vue";
import UpdateName from "@/popup/Profile/UpdateName.vue";
import UpdateRole from "@/popup/Profile/UpdateRole.vue";
import UpdatePhoneNumber from "@/popup/Profile/UpdatePhoneNumber.vue";
import UpdateEmail from "@/popup/Profile/UpdateEmail.vue";
import avatar from "@/layout/LayoutPage/img/avatar.jpg";
import { reactive, ref, onMounted } from "vue";
import ItemProfile from "./ItemProfile.vue";
import { useRoute } from "vue-router";
import AddSigner from "@/popup/Profile/AddSigner.vue";
import ButtonEdit from "../Company/ButtonEdit.vue";
import UpdatePassword from "@/popup/Profile/UpdatePassword.vue";
export default {
  name: "ProFile",
  components: {
    ItemProfile,
    UpdateAvatar,
    UpdateName,
    UpdateRole,
    UpdateEmail,
    UpdatePhoneNumber,
    AddSigner,
    ButtonEdit,
    UpdatePassword,
  },
  setup() {
    const state = reactive({
      ID: false,
      type: false,
      pop: 0,
    });
    const route = useRoute();
    onMounted(() => {
      state.ID = route.params.id;
      state.type = route.params.type;
    });
    const fakedata = ref([
      { id: 1, tittle: "プロフィール画像", data: avatar },
      { id: 2, tittle: "氏名", data: "水原 コーコー" },
      { id: 3, tittle: "役職", data: "デザイナー" },
      { id: 4, tittle: "電話番号", data: 8132159865 },
      { id: 5, tittle: "メールアドレス", data: "coco_mizuhara@gmail.com" },
      { id: 6, tittle: "アカウントの種類", data: "メンバー" },
    ]);

    return { state, fakedata, HiddenPassword };
  },
  methods: {
    popUp(id) {
      this.state.pop = id;
    },
    turnOff() {
      this.state.pop = 0;
    },
    setData(data) {
      var id = this.state.pop - 1;
      this.fakedata[id].data = data;
      this.state.pop = 0;
    },
    backToList() {
      this.$router.push({ path: "/management/member" });
    },
    SignerPopup() {
      this.state.pop = 6;
    },
    popupPassword() {
      this.state.pop = 7;
    },
  },
};
</script>

<style lang = "scss" scoped>
.profile {
  position: fixed;
  top: 127px;
  left: 300px;
  text-decoration: none;
  width: 1235px;
  height: 680px;
  overflow-y: auto;
  display: flex;
  flex-direction: column;
  .SignerSetting {
    margin-top: 50px;
    .h {
      margin-bottom: 10px;
      font-size: 16px;
      color: #271b5a;
      font-family: "Noto Sans JP";
      font-weight: 400;
      display: flex;

      .AddSigner {
        margin-left: 690px;
        color: rgb(247, 84, 25);
        display: flex;
        cursor: pointer;
        &:hover {
          color: rgb(255, 165, 0);
        }
        font-size: 14px;
      }
    }
    .box {
      width: 900px;
      height: 152px;
      border-top: 3px solid black;
      border-bottom: 2px solid lightgrey;
      display: flex;
      flex-direction: column;
      justify-content: center;
      align-items: center;
      .l1 .l2 {
        color: #666666;
      }
    }
  }
  .PasswordSetting {
    margin-top: 50px;
    .box2 {
      width: 900px;
      height: 100px;
      border-top: 3px solid black;
      border-bottom: 2px solid lightgrey;
      display: flex;
      flex-direction: row;

      align-items: center;
      .tittle {
        font-size: 12px;
        color: #564c7e;
        font-family: "Noto Sans JP";
        font-weight: 400;
        margin-left: 20px;
      }
      img {
        position: absolute;
        left: 233px;
      }
      .Edit {
        border: none;
        background: none;
        cursor: pointer;
        position: absolute;
        left: 831px;
      }
    }
  }
  .h {
    margin-bottom: 10px;
    font-size: 16px;
    color: #271b5a;
    font-family: "Noto Sans JP";
    font-weight: 400;
    display: flex;

    .AddSigner {
      margin-left: 700px;
      color: rgb(247, 84, 25);
      display: flex;
      margin-right: 10px;
      cursor: pointer;
      &:hover {
        color: rgb(255, 165, 0);
      }
      font-size: 14px;
    }
  }
  .ProfileSetting {
    display: flex;
    flex-direction: column;
    margin-top: 20px;
  }
}
.mid {
  position: fixed;
  display: flex;
  // align-items: center;
  left: 300px;
  top: 91px;
  .h1 {
    font-size: 14px;
    font-family: "Noto Sans JP";
    font-weight: 400;
    color: #aaaaaa;
    cursor: pointer;
  }
  .h2 {
    font-size: 14px;
    font-family: "Noto Sans JP";
    font-weight: 400;
    color: #271b5a;
    margin-left: 5px;
  }
  #arow {
    margin-left: 21px;
  }
}
</style>