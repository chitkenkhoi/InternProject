<template>
  <Footer class="Footer" v-on:scroll="scrollto" />
  <div class="head">企業情報</div>

  <div class="Company">
    <div class="info1" ref="ref1">
      <div class="header">企業情報の設定</div>
      <div class="info">
        <item
          v-for="ele in state.accountarray"
          :key="ele.id"
          :itemProp="ele"
          v-on:Pop-up="set"
        />
      </div>
    </div>
    <div class="info2" ref="ref2">
      <div class="header">セキュリティ設定</div>
      <div class="element">
        <ele
          class="turn"
          :turnProp="state.turn[0]"
          v-on:pop-up="setpop"
          v-on:switch="switchdata"
        />
      </div>
    </div>
    <div class="info2" ref="ref3">
      <div class="header">メール通知設定</div>
      <div class="element">
        <ele
          class="turn"
          :turnProp="state.turn[1]"
          v-on:pop-up="setpop"
          v-on:switch="switchdata"
        />
      </div>
    </div>
    <div class="info2" ref="ref4">
      <div class="header">
        <div class="text">アラート通知先の設定</div>
        <div class="emailBtn" @click="emailPopup">
          <span class="material-symbols-outlined"> add </span
          >アラート通知先の設定をする
        </div>
      </div>
      <div class="last">
        <div class="line" id="line1">
          アラート通知先のメールアドレスをデフォルト設定できます。
        </div>
        <div class="line" id="line2">
          また、他のメニューから契約書ごとに変更することもできます。
        </div>
        <div class="space"></div>
        <div class="space" id="none"></div>
      </div>
    </div>
  </div>
  <EditPhoto
    v-if="state.showmodal === true && state.ID === 1"
    v-on:no-pop-up="turnoff"
  />
  <EditName
    v-if="state.showmodal === true && state.ID === 2"
    v-on:no-pop-up="turnoff"
    v-on:data-change="setdata"
    :popprop="state.accountarray[1]"
  />
  <EditAddress
    v-if="state.showmodal === true && state.ID === 3"
    v-on:no-pop-up="turnoff"
    v-on:data-change="setdata"
    :popprop="state.accountarray[2]"
  />
  <EditPhonenumber
    v-if="state.showmodal === true && state.ID === 4"
    v-on:no-pop-up="turnoff"
    v-on:data-change="setdata"
    :popprop="state.accountarray[3]"
  />
  <EditRename
    v-if="state.showmodal === true && state.ID === 5"
    v-on:no-pop-up="turnoff"
    v-on:data-change="setdata"
    :popprop="state.accountarray[4]"
  />
  <EditSign
    v-if="state.showmodal === true && state.ID === 9"
    v-on:no-pop-up="turnoff"
  />
  <Confirm
    v-if="state.showmodal === true && (state.ID === 10 || state.ID === 11)"
    v-on:no-pop-up="turnoff"
    v-on:to-switch="switchdata(state.ID)"
  />
  <AddEmail
    v-if="state.showmodal === true && state.ID === 12"
    v-on:no-pop-up="turnoff"
  />
</template>


<script>
import EditPhoto from "../../../popup/Company/UpdateLogo.vue";
import EditName from "../../../popup/Company/EditCompanyName.vue";
import EditAddress from "../../../popup/Company/EditAddress.vue";
import EditPhonenumber from "../../../popup/Company/EditPhonenumber.vue";
import EditRename from "../../../popup/Company/EditRepresentativeName.vue";
import item from "./CompanyItem.vue";
import Footer from "../../../layout/LayoutPage/Footer.vue";
import Confirm from "../../../popup/Company/ConfirmOFFON.vue";
import AddEmail from "@/popup/Company/AddEmail.vue";
import EditSign from "@/popup/Company/EditSignature.vue";
import { reactive } from "vue";
import ele from "./Element.vue";
import sign from "../../../../public/logo/stamp.png";
import HTTP from "@/http-common";
export default {
  name: " CompanyPage",
  components: {
    ele,
    item,
    EditPhoto,
    EditAddress,
    EditPhonenumber,
    EditRename,
    Footer,
    Confirm,
    AddEmail,
    EditSign,
    EditName,
  },
  setup() {
    const state = reactive({
      accountarray: [
        {
          id: 1,
          tittle: "企業のロゴ",
          data: "",
          type: "img",
          popup: "EditPhoto",
        },
        {
          id: 2,
          tittle: "企業名",
          data: "〇〇株式会社",
          type: "string",
          popup: "EditCompanyName",
        },
        {
          id: 3,
          tittle: "企業住所",
          data: "〒100-0014東京都千代田区永田町2-14-3",
          type: "string",
          popup: "EditAddress",
        },
        {
          id: 4,
          tittle: "電話番号",
          data: "",
          type: "string",
          popup: "EditPhonenumber",
        },
        {
          id: 5,
          tittle: "代表者氏名",
          data: "山田　太郎",
          type: "string",
          popup: "EditRename",
        },
        {
          id: 6,
          tittle: "パートナー会社",
          data: "株式会社Wiz",
          type: "string",
          popup: false,
        },
        {
          id: 7,
          tittle: "紹介営業担当者",
          data: "Garment",
          type: "",
          popup: false,
        },
        {
          id: 8,
          tittle: "管理番号",
          data: "0123456789",
          type: "string",
          popup: false,
        },
        {
          id: 9,
          tittle: "オリジナル印影",
          data: sign,
          type: "img",
          popup: "EditSign",
        },
      ],
      turn: [
        {
          id: 10,
          data: "オン",
          tittle: "すべてのメンバーに 二段階認証を要求する",
          db: true,
        },
        { id: 11, data: "オフ", tittle: "メール通知を許可する", db: false },
      ],
      ID: 0,
      showmodal: false,
    });

    return { state };
  },
  created() {
        console.log("hello")
        HTTP.get(`getCompanyInformation/0123456789`)
        .then((response) => {
          const obj = response.data.data
          this.state.accountarray[0].data = obj.logo
          this.state.accountarray[1].data = obj.c_name
          this.state.accountarray[2].data = obj.address
          this.state.accountarray[3].data = obj.phone_number
          this.state.accountarray[4].data = obj.representative_name
          this.state.accountarray[5].data = obj.partner_company
          this.state.accountarray[8].data = obj.signature

        })
        .catch((e) => {
          console.log(e);
        });
    },
  methods: {
   
    switchdata(id) {
      if (this.state.turn[id - 10].data === "オフ") {
        this.state.turn[id - 10].data = "オン";
        this.state.turn[id - 10].db = true;
      } else {
        this.state.turn[id - 10].data = "オフ";
        this.state.turn[id - 10].db = false;
      }
      this.state.showmodal = false;
    },
    set(id) {
      this.state.showmodal = true;
      this.state.ID = id;
    },
    setpop(id) {
      this.state.showmodal = true;
      this.state.ID = id;
    },
    turnoff() {
      this.state.showmodal = false;
    },
    createObject(){
      return {
        control_number: this.state.accountarray[7].data,
        c_name: this.state.accountarray[1].data,
        logo: this.state.accountarray[0].data,
        address: this.state.accountarray[2].data,
        phone_number: this.state.accountarray[3].data,
        representative_name: this.state.accountarray[4].data,
        partner_company: this.state.accountarray[5].data,
        signature: this.state.accountarray[8].data,
        created_at: null
      }
    },
    setdata(data) {
      this.state.showmodal = false;
      var id = this.state.ID;
      this.state.accountarray[id - 1].data = data;
      
      
      HTTP.put(`updateCompany/0123456789`,{
        body: this.createObject()
      })
      .then(response => {
        console.log(response)
      })
      .catch(e => {
      console.log("loi",e)
    })
    },
    emailPopup() {
      this.state.showmodal = true;
      this.state.ID = 12;
    },
    scrollto(id) {
      if (id === 1) this.$refs.ref1.scrollIntoView({ behavior: "smooth" });
      if (id === 2) this.$refs.ref2.scrollIntoView({ behavior: "smooth" });
      if (id === 3) this.$refs.ref3.scrollIntoView({ behavior: "smooth" });
      if (id === 4) this.$refs.ref4.scrollIntoView({ behavior: "smooth" });
    },
  },
};
</script>

<style lang = "scss" scoped>
.Footer {
  position: fixed;
  top: 140px;
  right: 20px;
  z-index: 20;
}
.head {
  position: fixed;
  top: 90px;
  left: 300px;
  color: rgb(72, 61, 139);
  font-size: 14px;
}
.Company {
  position: fixed;
  margin-top: 0px;
  text-decoration: none;
  width: 1237px;
  height: 680px;
  overflow-y: auto;
  .info1 {
    margin-top: 20px;
  }
  .info2 {
    margin-top: 30px;
    width: 900px;
    .last {
      border-top: 5px solid black;
      display: flex;
      flex-direction: column;
      align-items: center;
      .line {
        font-size: 14px;
      }
      #line1 {
        margin-top: 40px;
      }
      #line2 {
        margin-top: 5px;
      }
      .space {
        width: 900px;
        height: 40px;
        border-bottom: 1px solid lightgrey;
      }
      #none {
        border-bottom: none;
      }
    }
  }

  .header {
    margin-bottom: 10px;
    color: rgb(72, 61, 139);
    font-size: 16px;
    display: flex;
    justify-content: space-between;
    .emailBtn {
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
  .info {
    display: flex;
    flex-direction: column;
    color: black;
  }
}
</style>