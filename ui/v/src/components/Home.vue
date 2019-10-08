<template>
  <div class="home">
    <h1><span style="font-weight: 100">{{ today[0] }}</span> <span style="font-weight: 400">{{ today[1] }}</span>, <span style="font-weight:400">{{ today[2].split(' ')[1] }}</span></h1>
    <img src="../../static/hospital-reception-svg.png">
  </div>
</template>

<script>
import moment from 'moment'

export default {
  name: 'Home',
  data () {
    return {
      msg: 'Welcome.',
      today: null
    }
  },

  methods: {

  },

  mounted () {
    this.today = moment().format('LLLL').split(',')
  }
}
</script>

<style scoped>
.home {
  margin-top: 40px;
}

body {
  align-content: center;
  color: #5e9afa;
}

h1, h2 {
  color:#91bafb;
  font-weight: 400;
  text-align: center;
}

img {
  display: block;
  height: 350px;
  margin: 0 auto;
  margin-top: 50px;
  width: 400px;
}

</style>
