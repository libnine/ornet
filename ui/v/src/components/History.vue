<template>
  <div class="md-auto">
    <div class="md-layout md-gutter md-alignment-center-center">
      <md-datepicker v-model="date_from" :md-open-on-focus="true" style="color:#91bafb" md-immediately>
        <label style="color:#91bafb">From</label>
      </md-datepicker>
      <md-datepicker v-model="date_to" :md-open-on-focus="true" style="color:#91bafb; margin-bottom: 25px;" md-immediately>'
        <label style="color:#91bafb">To</label>
      </md-datepicker>
    </div>
    <div class="md-layout md-gutter">
      <transition name="fade">
        <a class="selectlink" v-if="date_from != null && date_from !='' && date_to != null && date_to !=''">Search</a>
      </transition>
    </div>
  </div>
</template>

<script>
import axios from 'axios'

export default {
  name: 'History',
  data () {
    return {
      date_from: null,
      date_to: null,
      dump: null,
      historical_path: 'http://localhost:81/api/h'
    }
  },

  mounted () {
    axios.get(this.historical_path, {
      timeout: 6000,
      headers: {}
    }).then((response) => {
      this.dump = JSON.parse(JSON.stringify(response.data))
    })
  }
}
</script>

<style lang="scss" scoped>
.selectlink {
  border: 2px solid transparent;
  font-weight: 200;
  display: block;
  font-size: 24px;
  margin: 0 auto;
  padding: 5px;
  text-decoration: none;
  transition: all 200ms ease;

  &:hover {
    text-decoration: none;
    border-color:#91bafb;
    transform: scale(1.075);
  }

  &:active {
    background-color: #91bafb;
    font-size: 28px;
    color:white;
  }
}

.md-auto {
  display: block;
  margin: 0 auto;
  padding: 15px;
  width: 300px;
}

.md-select {
  margin: 15px;
}

.fade-enter-active, .fade-leave-active {
  transition: opacity .55s;
}

.fade-enter, .fade-leave-to {
  opacity: 0;
}

</style>
