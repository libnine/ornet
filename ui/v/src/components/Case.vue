<template>
  <div style="overflow-x: hidden">
    <div class="md-layout md-gutter md-alignment-center-center">
      <h1 style="text-align:center"><span style="font-weight:100">{{ $route.params.date_formatted[0] }}</span> {{ $route.params.date_formatted[1] }} {{ $route.params.date_formatted[2] }} {{ $route.params.date_formatted[3] }}</h1>
      <md-button v-if="dump_cases.length >= 2" class="md-icon-button" style="position:absolute; left: 500px; margin-top: 75px" @click="dump_reverse()">
        <md-icon style="color:#91bafb">swap_vert</md-icon>
      </md-button>
      <md-dialog :md-active.sync="show_dialog" style="padding:5px">
        <md-dialog-title style="font-weight: 200; color: #5e9afa">{{ dialog_data.case_id }}</md-dialog-title>
          <h4><strong>Patient</strong>: {{ dialog_data.patient }}</h4>
          <h4><strong>Body part</strong>: {{ dialog_data.body_part }}</h4>
          <h4><strong>Left/Right</strong>: {{ dialog_data.left_right }}</h4>
          <h4><strong>Description</strong>: {{ dialog_data.description }}</h4>
          <h4><strong>Hospital</strong>: {{ dialog_data.hospital }}</h4>
          <h4><strong>Doctor</strong>: {{ dialog_data.doctor }}</h4>
          <h4><strong>Rep</strong>: {{ dialog_data.rep }}</h4>
          <h4><strong>Date & Time</strong>: {{ dialog_data.case_date }}</h4>
        <md-dialog-actions>
          <md-button class="md-primary" @click="noop()">Details</md-button>
          <md-button class="md-primary" @click="show_dialog = false">Close</md-button>
        </md-dialog-actions>
      </md-dialog>
    </div>

    <div class="md-layout md-gutter md-alignment-center-center">
      <md-list class="md-double-line" style="width: 275px; background-color:rgb(250, 250, 250)">
        <md-list-item style="padding: 2px; margin: 3px; height: 75px" v-for="d in dump_cases" v-bind:key="d.case_id" @click="dialog_func(d)">
          <md-icon class="md-primary">local_hospital</md-icon>
          <div class="md-list-item-text">
            <span style="color: #5e9afa">{{ d.case_id }}</span>
            <span style="color: #91bafb">{{ d.case_date.substring(11, 16) }}</span>
          </div>
        </md-list-item>
      </md-list>
    </div>
  </div>
</template>

<script>
export default {
  name: 'Case',
  props: ['dump', 'date_formatted'],

  data () {
    return {
      dump_cases: null,
      dump_cases_rev: null,
      dialog_data: {
        '_id': null,
        'body_part': null,
        'case_date': null,
        'case_id': null,
        'date_created': null,
        'description': null,
        'doctor': null,
        'hospital': null,
        'left_right': null,
        'patient': null,
        'rep': null
      },
      show_dialog: false
    }
  },

  methods: {
    dialog_func: function (d) {
      this.dialog_data = d
      this.show_dialog = true
    },

    dump_reverse () {
      this.dump_cases = this.dump_cases.slice().reverse()
    }
  },

  mounted () {
    this.dump_cases = this.$route.params.dump
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

h4 {
  margin: 0;
  margin-left: 15px;
  font-weight: 400;
}

</style>
