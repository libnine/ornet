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
        <md-dialog-title style="font-weight: 200"><strong>{{ case_id }}</strong> successfully created.</md-dialog-title>
        <md-dialog-actions>
          <md-button class="md-primary" @click="success_close">OK</md-button>
        </md-dialog-actions>
      </md-dialog>
    </div>

    <div class="md-layout md-gutter md-alignment-center-center">
      <div class="md-layout-item md-size-25">
        <h1 style="font-weight: 400">{{ case_id }}</h1>
      </div>
    </div>

    <div class="md-layout md-gutter md-alignment-center-center">
      <div class="md-layout-item md-size-25">
        <md-field>
          <label for="times" style="color:#5e9afa;">Case Time</label>
          <md-select v-model="case_time" name="times">
            <md-option v-for="(time, key) in times" :key="key" :value="key">{{ time }}</md-option>
          </md-select>
        </md-field>
      </div>

      <div class="md-layout-item md-size-25">
        <md-datepicker v-model="case_date" :md-open-on-focus="true" md-immediately>
          <label style="color:#5e9afa;">Case Date</label>
        </md-datepicker>
      </div>
    </div>

    <div class="md-layout md-gutter md-alignment-center-center">
      <div class="md-layout-item md-size-25">
        <md-field>
          <label style="color:#5e9afa;">Patient</label>
          <md-input v-model="patient"></md-input>
        </md-field>
      </div>

      <div class="md-layout-item md-size-25">
        <md-datepicker v-model="patient_dob" :md-open-on-focus="true" md-immediately>
          <label style="color:#5e9afa;">Patient DOB</label>
        </md-datepicker>
      </div>
    </div>

    <div class="md-layout md-gutter md-alignment-center-center">
      <div class="md-layout-item md-size-25">
        <md-field>
          <label style="color:#5e9afa;">Body Part</label>
          <md-input v-model="body_part"></md-input>
        </md-field>
      </div>

      <div class="md-layout-item md-size-25">
        <md-autocomplete
        v-model="left_right"
        :md-options="left_right_options"
        md-dense>
          <label style="color:#5e9afa;">Left/Right</label>
        </md-autocomplete>
      </div>
    </div>

    <div class="md-layout md-gutter md-alignment-center-center">
      <div class="md-layout-item md-size-25">
        <md-field>
          <label style="color:#5e9afa;">Doctor</label>
          <md-input v-model="doctor"></md-input>
        </md-field>
      </div>

      <div class="md-layout-item md-size-25">
        <md-field>
          <label style="color:#5e9afa;">Hospital</label>
          <md-input v-model="hospital"></md-input>
        </md-field>
      </div>
    </div>

    <div class="md-layout md-gutter md-alignment-center-center">
      <div class="md-layout-item md-size-25">
        <md-field>
          <label style="color:#5e9afa;">Description</label>
          <md-input v-model="description"></md-input>
        </md-field>
      </div>

      <div class="md-layout-item md-size-25">
        <md-autocomplete
        v-model="subspecialty"
        :md-options="subspecialty_options"
        md-dense>
          <label style="color:#5e9afa;">Subspecialty</label>
        </md-autocomplete>
      </div>
    </div>

    <div class="md-layout md-gutter md-alignment-center-center">
      <md-autocomplete
        v-model="rep"
        :md-options="users"
        :md-fuzzy-search="false"
        id="users-search">
        <label style="color:#5e9afa;">Search for reps...</label>

        <template slot="md-autocomplete-item" slot-scope="{ item, term }">
          <md-highlight-text :md-term="term">{{ item }}</md-highlight-text>
        </template>

        <template slot="md-autocomplete-empty" slot-scope="{ term }" v-if="term != null">
          "{{ term }}" is not currently on file. <a @click="$router.push({ name: 'Add', params: { user: term }})">You can add them here</a>.
        </template>
      </md-autocomplete>
    </div>

    <div class="md-layout md-gutter md-alignment-center-center">
      <md-field style="width: 500px">
        <label style="color:#5e9afa">Notes...</label>
        <md-textarea v-model="notes" style="height: 150px;"></md-textarea>
      </md-field>
    </div>

    <div class="md-layout md-gutter md-alignment-center-center">
      <md-button class="md-raised" style="color:#5e9afa; margin-top:15px;" @click="submit()">Create</md-button>
    </div>
  </div>
</template>

<script>
import axios from 'axios'
import moment from 'moment'
import format from 'date-fns/format'

export default {
  name: 'Create',
  data () {
    let dateFormat = this.$material.locale.dateFormat || 'yyyy-MM-dd'
    let now = new Date()

    return {
      body_part: null,
      case_date: format(now, dateFormat),
      case_id: Math.random().toString(36).replace(/[^a-z0-9]+/g, '').substr(1, 7),
      case_time: null,
      create_path: 'http://localhost:81/api/c',
      description: null,
      doctor: null,
      hospital: null,
      left_right: null,
      left_right_options: ['Left', 'Right', 'N/A'],
      notes: null,
      notes_array: null,
      patient: null,
      patient_dob: null,
      rep: null,
      show_dialog: false,
      subspecialty: null,
      subspecialty_options: ['Orthopedics', 'Neurology', 'Cardiothoracic', 'Oral Maxillofacial',
        'Pediatric', 'Vascular', 'Urology', 'General Surgery', 'ENT', 'Plastic'].sort(),
      success_dialog: false,
      times: {
        '00:00:00': '12:00 AM',
        '01:00:00': '1:00 AM',
        '02:00:00': '2:00 AM',
        '03:00:00': '3:00 AM',
        '04:00:00': '4:00 AM',
        '05:00:00': '5:00 AM',
        '06:00:00': '6:00 AM',
        '07:00:00': '7:00 AM',
        '08:00:00': '8:00 AM',
        '09:00:00': '9:00 AM',
        '10:00:00': '10:00 AM',
        '11:00:00': '11:00 AM',
        '12:00:00': '12:00 AM',
        '13:00:00': '1:00 PM',
        '14:00:00': '2:00 PM',
        '15:00:00': '3:00 PM',
        '16:00:00': '4:00 PM',
        '17:00:00': '5:00 PM',
        '18:00:00': '6:00 PM',
        '19:00:00': '7:00 PM',
        '20:00:00': '8:00 PM',
        '21:00:00': '9:00 PM',
        '22:00:00': '10:00 PM',
        '23:00:00': '11:00 PM'
      },
      users_path: 'http://localhost:81/api/u',
      users: []
    }
  },

  methods: {
    submit () {
      if
      (this.body_part == null ||
      this.case_date == null ||
      this.case_time == null ||
      this.doctor == null ||
      this.hospital == null ||
      this.left_right == null ||
      this.patient == null ||
      this.patient_dob == null ||
      this.rep == null ||
      this.subspecialty == null) {
        this.show_dialog = true
        return
      }

      if (this.notes != null) {
        let msg = this.notes
        let now = moment(Date.now()).toISOString()
        let user = 'n/a'

        this.notes_array = [
          {
            "date_created": now,
            "msg": msg,
            "user": user
          }
        ]
      } else {
        this.notes_array = null
      }

      axios.post(this.create_path,
      {
        body_part: this.body_part,
        case_date: moment(`${this.case_date} ${this.case_time}`).toISOString(),
        case_id: this.case_id,
        description: this.description || '',
        doctor: this.doctor,
        hospital: this.hospital,
        left_right: this.left_right,
        patient: this.patient,
        patient_dob: this.patient_dob.toString(),
        rep: this.rep,
        subspecialty: this.subspecialty,
        date_created: moment(Date.now()).toISOString(),
        notes: this.notes_array
      },
      {
        timeout: 6000,
        headers: {}
      }).then((response) => {
        this.success_dialog = true
      })
    },

    success_close () {
      this.success_dialog = false
      window.location.reload()
    }
  },

  mounted () {
    axios.get(this.users_path,
    {
      timeout: 6000,
      headers: {}
    }).then((response) => {
      this.dump = JSON.parse(JSON.stringify(response.data))

      try {
        for (var i = 0; i < this.dump.length; i++) {
          this.users.push(`${this.dump[i].first_name} ${this.dump[i].last_name}`)
        }
      } catch (_) {}
    })
  }
}
</script>

<style>
h1, h2 {
  color:#5e9afa;
  font-weight: 400;
  text-align: center;
}

label {
  color: #91bafb;
}

.md-field {
  width: 300px;
  margin: 5px;
}

.md-highlight-text {
  font-weight: 400;
  font-size: 15px;
  color:#5e9afa;
}

#users-search {
  display: block;
  margin: 0 auto;
  margin-top: 5px;
  margin-bottom: 15px;
  width: 500px;
}

</style>
