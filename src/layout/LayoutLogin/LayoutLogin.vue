<template>
  <div class = "layout">
<back-ground>
    <div class = "LayoutLogin_main">
        <router-view></router-view>
    </div>
  </back-ground>
</div>
</template>

<script>
import BackGround from './BackGround.vue'
export default {
  components: { BackGround },

name: "LayoutLogin",
conponent: {BackGround}
}
</script>

<style lang = "scss" scoped>
.layout{
  position:fixed;
  top:0px;
  left:0px;
  width:100vw;
  height:800px;
  overflow-y: auto;

}

</style>