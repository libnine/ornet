<template>
  <div>
    <md-toolbar class="main-toolbar">
      <div class="md-toolbar-section-start">
        <md-button class="md-icon-button" @click="$router.go(-1)">
            <md-icon style="color:skyblue">first_page</md-icon>
        </md-button>
        <md-button class="md-icon-button" @click="$router.push({ name: 'Home' })">
            <md-icon style="color:skyblue">home</md-icon>
        </md-button>
        <md-menu>
          <md-button class="md-icon-button" md-menu-trigger>
              <md-icon style="color:skyblue">menu</md-icon>
          </md-button>
          <md-menu-content class="mdmenuStyle">
            <md-menu-item @click="$router.push({ name: 'Calendar' })">
              <span style="color:#91bafb">Cases</span>
              <md-icon style="color:#91bafb">assignment</md-icon>
            </md-menu-item>
            <md-menu-item @click="$router.push({ name: 'Create' })">
              <span style="color:#91bafb">Create</span>
              <md-icon style="color:#91bafb">create</md-icon>
            </md-menu-item>
            <md-menu-item @click="$router.push({ name: 'Search' })">
                <span style="color:#91bafb">Search</span>
                <md-icon style="color:#91bafb">search</md-icon>
            </md-menu-item>
            <md-menu-item @click="$router.push({ name: 'History' })">
              <span style="color:#91bafb">History</span>
              <md-icon style="color:#91bafb">history</md-icon>
            </md-menu-item>
            <md-menu-item>
              <span style="color:#91bafb">Profile</span>
              <md-icon style="color:#91bafb">account_circle</md-icon>
            </md-menu-item>
            <md-menu-item>
              <span style="color:#91bafb">Help</span>
              <md-icon style="color:#91bafb">help_outline</md-icon>
            </md-menu-item>
          </md-menu-content>
        </md-menu>
      </div>
      <md-button class="md-primary" id="login">Login</md-button>
      <md-button class="md-primary">Register</md-button>
    </md-toolbar>
  </div>
</template>

<script>
export default {
  name: 'Header',
  data () {
    return {
      msg: null
    }
  }
}
</script>

<style>
#login {
  font-weight: 100;
}

#main-toolbar {
  position: sticky;
}

.mdmenuStyle {
  max-height: none;
}

</style>
