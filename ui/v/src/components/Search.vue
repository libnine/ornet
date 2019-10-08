<template>
  <div class="md-auto">
    <div id="auto-c">
      <md-autocomplete
        v-model="selected"
        :md-options="users"
        :md-fuzzy-search="false"
        >
        <label id="placehold" v-if="selected == null || selected == ''">Start typing...</label>

        <template slot="md-autocomplete-item" slot-scope="{ item, term }">
          <md-highlight-text :md-term="term">{{ item }}</md-highlight-text>
        </template>

        <template slot="md-autocomplete-empty" slot-scope="{ term }" v-if="term != null">
          "{{ term }}" is not currently on file. <a @click="$router.push({ name: 'Add', params: { user: term }})">You can add them here</a>.
        </template>
      </md-autocomplete>
    </div>

    <div class="md-layout md-gutter">
      <transition name="fade">
        <a class="selectlink" v-if="selected != null && selected !='' && view_link" @click="show_data()">Select</a>
        <a class="selectlink" v-if="selected != null && selected !='' && view_next_link" @click="next()">Next</a>
      </transition>
    </div>

    <div class="md-layout md-gutter">
      <transition name="fade">
        <span v-if="data_on && view_data">
          <div class="md-auto">
            <md-card>
              <md-list>
                <md-list-item>
                  <md-icon style="color:#91bafb">face</md-icon>
                  <span class="md-list-item-text">{{ data_display.first_name }} {{ data_display.last_name }}</span>
                </md-list-item>
                <md-list-item>
                  <md-icon style="color:#91bafb">headset</md-icon>
                  <span class="md-list-item-text">{{ data_display.role }}</span>
                </md-list-item>
                <md-list-item>
                  <md-icon style="color:#91bafb">people_outline</md-icon>
                  <span class="md-list-item-text">{{ data_display.company }}</span>
                </md-list-item>
                <md-list-item>
                  <md-icon style="color:#91bafb">email</md-icon>
                  <span class="md-list-item-text">{{ data_display.email }}</span>
                </md-list-item>
                <md-list-item>
                  <md-icon style="color:#91bafb">phone</md-icon>
                  <span class="md-list-item-text">{{ data_display.phone }}</span>
                </md-list-item>
              </md-list>
            </md-card>
          </div>
        </span>
      </transition>
    </div>
  </div>
</template>

<script>
import axios from 'axios'

export default {
  name: 'Search',
  data () {
    return {
      case_byusers: 'http://localhost:81/api/cu',
      data_display: null,
      dump: null,
      selected: null,
      users: [],
      users_path: 'http://localhost:81/api/u',
      view_data: false,
      view_next_link: false,
      view_link: false
    }
  },

  methods: {
    show_data () {
      var i = this.users.indexOf(this.selected)
      this.view_data = true
      this.view_next_link = true
      this.view_link = false
      this.data_display = this.dump[i]
    },

    next () {
      axios.post(this.case_byusers, 
      {
        first_name: this.data_display.first_name,
        last_name: this.data_display.last_name,
        role: this.data_display.role,
        company: this.data_display.company,
        phone: this.data_display.phone,
        email: this.data_display.email
      },
      {
        timeout: 6000,
        headers: {}
      }).then((response) => {
        this.$router.push({ name: "User", params: { dump: JSON.parse(JSON.stringify(response.data)), name: `${this.data_display.first_name} ${this.data_display.last_name}`, user: `${this.data_display.first_name.toLowerCase()}_${this.data_display.last_name.toLowerCase()}` }})
      })
    }
  },

  computed: {
    data_on () {
      if (this.selected != null && this.selected != '' && this.users.includes(this.selected)) {
        this.view_link = true
        return true
      } else {
        this.data_display = null
        this.view_data = false
        return false
      }
    }
  },

  mounted () {
    axios.get(this.users_path, {
      timeout: 6000,
      headers: {}
    }).then((response) => {
      this.dump = JSON.parse(JSON.stringify(response.data))

      let i = 0
      for (i = 0; i < this.dump.length; i++) {
        this.users.push(`${this.dump[i].first_name} ${this.dump[i].last_name}`)
      }
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

.card-icons {
  color: #91bafb;
}

.card-text {
  color: #5e9afa;
  font-weight: 200;
}

#placehold {
  color:#5e9afa;
  font-weight: 100;
}

.md-auto {
  display: block;
  margin: 0 auto;
  padding: 15px;
  width: 600px;
}

.md-autocomplete + strong {
  display: block;
  margin: 0 auto;
  width: 600px;
}

.md-highlight-text {
  font-weight: 400;
  font-size: 15px;
  color:#5e9afa;
}

.fade-enter-active, .fade-leave-active {
  transition: opacity .55s;
}

.fade-enter, .fade-leave-to {
  opacity: 0;
}

</style>
