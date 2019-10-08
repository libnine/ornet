<template>
  <div id="app">
    <t-header id="t-head"/>
    <transition
      name="fade"
      mode="out-in">
      <router-view />
    </transition>
  </div>
</template>

<script>
import Header from './components/Header'

export default {
  name: 'App',

  components: {
    't-header': Header
  },

  created () {
    document.title = 'Ornet Medical Network'
  }
}
</script>

<style>
#app {
  justify-items: center;
  align-items: center;
}

#t-head {
  align-content: center;
  padding: 15px;
}

.fade-enter-active,
.fade-leave-active {
  transition-duration: 0.3s;
  transition-property: opacity;
  transition-timing-function: ease;
}

.fade-enter,
.fade-leave-active {
  opacity: 0
}
</style>
