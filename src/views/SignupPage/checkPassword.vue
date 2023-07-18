<template>
  <span class="material-symbols-outlined" id = "arrow">
arrow_drop_up
</span>
  <div class="box">
    <div class="item">
      <div class="circle" :class = "{okay:check_character}">
        <span class="material-symbols-outlined" id="check"> done </span>
      </div>
      <div class = "tittle">1文字以上の英字</div>
    </div>
    <div class="item">
      <div class="circle" :class = "{okay:check_number}">
        <span class="material-symbols-outlined" id="check"> done </span>
      </div>
      <div class = "tittle">１文字以上の数字</div>
    </div>
    <div class="item">
      <div class="circle" :class = "{okay:check_length}">
        <span class="material-symbols-outlined" id="check"> done </span>
      </div>
      <div class = "tittle">8文字以上</div>
    </div>
  </div>
</template>

<script>
export default {
    name:"checkPassword",
    props: ['check_character','check_number','check_length'],
    setup(){

    }

}
</script>

<style lang = "scss" scoped>
#arrow{
    width:10px;
    height:5px;
    color: white;
    margin-top:13px;
}
.box{
    margin-top:-14px;
    width:470px;
    background-color:white;
    height:80px;
    display:flex;
    flex-direction: column;
    padding:15px;   
    box-shadow: rgba(99, 99, 99, 0.2) 0px 0px 8px 0px;
    justify-content: space-between;
    .item{
        width:470px;
        height:20px;
        display:flex;
        align-items: center;
        .circle{
            width:20px;
            height:20px;
            border-radius:50%;
            background-color: #DDDDDD;
            display:flex;
            align-items: center;
            justify-content: center;
            #check{
                font-size:15px;
                margin-top:1px;
                color: white;
            }
        }
        .okay{
            background-color: #17B182;
        }
        .tittle{
            margin-left:10px;
            color:#666666;
            font-size:12px;
            font-family:'Noto Sans JP';
            font-weight:400;
        }
    }

}
</style>    