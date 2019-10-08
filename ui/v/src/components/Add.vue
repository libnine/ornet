<template>
  <div style="overflow-x: hidden">
    <div class="md-layout md-gutter md-alignment-center-center">
      <md-dialog :md-active.sync="show_dialog" style="padding:5px">
        <md-dialog-title style="font-weight: 200">All fields must be completed.</md-dialog-title>
        <md-dialog-actions>
          <md-button class="md-primary" @click="show_dialog = false">OK</md-button>
        </md-dialog-actions>
      </md-dialog>
    </div>

    <div class="md-layout md-gutter md-alignment-center-center">
      <md-dialog :md-active.sync="success_dialog" style="padding:5px">
        <md-dialog-title style="font-weight: 200"><strong>{{ first_name }} {{ last_name }}</strong> successfully created.</md-dialog-title>
        <md-dialog-actions>
          <md-button class="md-primary" @click="success_close">OK</md-button>
        </md-dialog-actions>
      </md-dialog>
    </div>

    <div class="md-layout md-gutter md-alignment-center-center" style="display: block; margin: 0 auto; margin-top: 35px; width: 475px;">
      <transition name="fade">
        <md-card>
          <md-list>
            <md-list-item>
              <md-icon style="color:#91bafb">face</md-icon>
              <md-field>
                <label style="color:#5e9afa;">Name</label>
                <md-input v-model="name"></md-input>
              </md-field>
            </md-list-item>
            <md-list-item>
              <md-icon style="color:#91bafb">headset</md-icon>
              <md-field>
                <label style="color:#5e9afa;">Role</label>
                <md-input v-model="role"></md-input>
              </md-field>
            </md-list-item>
            <md-list-item>
              <md-icon style="color:#91bafb">people_outline</md-icon>
              <md-field>
                <label style="color:#5e9afa;">Company</label>
                <md-input v-model="company"></md-input>
              </md-field>
            </md-list-item>
            <md-list-item>
              <md-icon style="color:#91bafb">email</md-icon>
              <md-field>
                <label style="color:#5e9afa;">Email</label>
                <md-input v-model="email"></md-input>
              </md-field>
            </md-list-item>
            <md-list-item>
              <md-icon style="color:#91bafb">phone</md-icon>
              <md-field>
                <label style="color:#5e9afa;">Phone</label>
                <md-input v-model="phone"></md-input>
              </md-field>
            </md-list-item>
          </md-list>
        </md-card>
      </transition>
      <div class="md-layout md-gutter md-alignment-center-center">
        <md-button class="md-raised" style="color:#5e9afa; margin-top:25px;" @click="submit()">Create</md-button>
      </div>
    </div>
  </div>
</template>

<script>
import axios from 'axios'
import moment from 'moment'

export default {
  name: 'Add',
  data () {
    return {
      add_path: 'http://localhost:81/api/u',
      company: null,
      email: null,
      first_name: null,
      last_name: null,
      name: null,
      phone: null,
      role: null,
      show_dialog: false,
      success_dialog: false,
      today: null
    }
  },

  methods: {
    submit () {
      if
      (this.company == null ||
      this.email == null ||
      this.name == null ||
      this.phone == null ||
      this.role == null) {
        this.show_dialog = true
        return
      }

      let full_name = this.name.split(" ")
      if (full_name.length == 2) {
        this.first_name = full_name[0]
        this.last_name = full_name[1]
      } else if (full_name.length == 3) {
        this.first_name = `${full_name[0]} ${full_name[1]}`
        this.last_name = full_name[2]
      } else if (full_name.length >= 4) {
        let n = full_name.length
        this.first_name = `${full_name[0]} ${full_name[1]}`
        this.last_name = full_name.slice(2, full_name.length).join(" ")
      }

      axios.post(this.add_path,
      {
        first_name: this.first_name,
        last_name: this.last_name,
        role: this.role,
        company: this.company,
        phone: this.phone,
        email: this.email,
        username: 'n/a',
        date_created: moment(Date.now()).toISOString()
      },
      {
        timeout: 6000,
        headers: {}
      }).then((response) => {
        this.success_dialog = true
      })
    },

    success_close() {
      this.success_dialog = false
      this.$router.back()
    }
  },

  mounted () {
    this.name = this.$route.params.user
  }
}
</script>

<style scoped>
body {
  align-content: center;
  color: #5e9afa;
}

h1, h2 {
  color:#91bafb;
  font-weight: 400;
  text-align: center;
}

.md-auto {
  display: block;
  margin: 0 auto;
  padding: 5px;
  width: 600px;
}

</style>
