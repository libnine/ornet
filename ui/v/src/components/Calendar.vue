<template>
  <div>
    <div class="md-layout md-gutter md-alignment-center-center">
      <div class="md-layout-item md-size-25" @click="$router.push({ name: 'Case', params: { dump: dump_filter_bydate(5), date_formatted: dates_list[5], case_date: `${dates_mmddyyyy[5]}`}})">
        <md-card md-with-hover>
          <md-card-area>
            <md-card-header>
              <span class="md-title" style="font-weight:100">{{ dates_list[5][0] }}</span><br>
              <span class="md-title" style="font-weight:400">{{ dates_list[5][1] }} {{ dates_list[5][2] }} {{ dates_list[5][3] }}</span><br>
              <span class="md-subhead-2"><b>{{ cases_count[5] }}</b> cases.</span><br><br><br>
              <span class="md-subhead">Starting at <b>6:30 AM</b> EST.</span><br><br><br>
            </md-card-header>
          </md-card-area>
        </md-card>
      </div>

      <div class="md-layout-item md-size-25" @click="$router.push({ name: 'Case', params: { dump: dump_filter_bydate(4), date_formatted: dates_list[4], case_date: `${dates_mmddyyyy[4]}`}})">
        <md-card md-with-hover>
          <md-card-area>
            <md-card-header>
              <span class="md-title" style="font-weight:100">{{ dates_list[4][0] }}</span><br>
              <span class="md-title" style="font-weight:400">{{ dates_list[4][1] }} {{ dates_list[4][2] }} {{ dates_list[4][3] }}</span><br>
              <span class="md-subhead-2"><b>{{ cases_count[4] }}</b> cases.</span><br><br><br>
              <span class="md-subhead">Starting at <b>7:30 AM</b> EST.</span><br><br><br>
            </md-card-header>
          </md-card-area>
        </md-card>
      </div>

      <div class="md-layout-item md-size-25" @click="$router.push({ name: 'Case', params: { dump: dump_filter_bydate(3), date_formatted: dates_list[3], case_date: `${dates_mmddyyyy[3]}`}})">
        <md-card md-with-hover>
          <md-card-area>
            <md-card-header>
              <span class="md-title" style="font-weight:100">{{ dates_list[3][0] }}</span><br>
              <span class="md-title" style="font-weight:400">{{ dates_list[3][1] }} {{ dates_list[3][2] }} {{ dates_list[3][3] }}</span><br>
              <span class="md-subhead-2"><b>{{ cases_count[3] }}</b> cases.</span><br><br><br>
              <span class="md-subhead">Starting at <b>8:15 AM</b> EST.</span><br><br><br>
            </md-card-header>
          </md-card-area>
        </md-card>
      </div>
    </div>

    <div class="md-layout md-gutter md-alignment-center-center">
      <div class="md-layout-item md-size-25" @click="$router.push({ name: 'Case', params: { dump: dump_filter_bydate(2), date_formatted: dates_list[2], case_date: `${dates_mmddyyyy[2]}`}})">
        <md-card md-with-hover>
          <md-card-area>
            <md-card-header>
              <span class="md-title" style="font-weight:100">{{ dates_list[2][0] }}</span><br>
              <span class="md-title" style="font-weight:400">{{ dates_list[2][1] }} {{ dates_list[2][2] }} {{ dates_list[2][3] }}</span><br>
              <span class="md-subhead-2"><b>{{ cases_count[2] }}</b> cases.</span><br><br><br>
              <span class="md-subhead">Starting at <b>6:30 AM</b> EST.</span><br><br><br>
            </md-card-header>
          </md-card-area>
        </md-card>
      </div>

      <div class="md-layout-item md-size-25" @click="$router.push({ name: 'Case', params: { dump: dump_filter_bydate(1), date_formatted: dates_list[1], case_date: `${dates_mmddyyyy[1]}`}})">
        <md-card md-with-hover>
          <md-card-area>
            <md-card-header>
              <span class="md-title" style="font-weight:100">{{ dates_list[1][0] }}</span><br>
              <span class="md-title" style="font-weight:400">{{ dates_list[1][1] }} {{ dates_list[1][2] }} {{ dates_list[1][3] }}</span><br>
              <span class="md-subhead-2"><b>{{ cases_count[1] }}</b> cases.</span><br><br><br>
              <span class="md-subhead">Starting at <b>7:30 AM</b> EST.</span><br><br><br>
            </md-card-header>
          </md-card-area>
        </md-card>
      </div>

      <div class="md-layout-item md-size-25"  @click="$router.push({ name: 'Case', params: { dump: dump_filter_bydate(0), date_formatted: dates_list[0], case_date: `${dates_mmddyyyy[0]}`}})">
        <md-card md-with-hover>
          <md-card-area>
            <md-card-header>
              <span class="md-title" style="font-weight:100">{{ dates_list[0][0] }}</span><br>
              <span class="md-title" style="font-weight:400">{{ dates_list[0][1] }} {{ dates_list[0][2] }} {{ dates_list[0][3] }}</span><br>
              <span class="md-subhead-2"><b>{{ cases_count[0] }}</b> cases.</span><br><br><br>
              <span class="md-subhead">Starting at <b>8:15 AM</b> EST.</span><br><br><br>
            </md-card-header>
          </md-card-area>
        </md-card>
      </div>
    </div>
  </div>
</template>

<script>
import axios from 'axios'
import moment from 'moment'

export default {
  name: 'Calendar',
  data () {
    return {
      cases_count : [],
      dates_mongo: [],
      dates_mmddyyyy: [],
      dates_list: [],
      dump: null,
      cases_path: 'http://localhost:81/api/c'
    }
  },

  methods: {
    dump_filter_bydate: function (d) {
      let date_filtered = this.dates_mongo[d]
      return this.dump.filter(function (x) {return x.case_date.substring(0, 10) == date_filtered.toString()})
    }
  },

  mounted () {
    axios.get(this.cases_path, {
      timeout: 6000,
      headers: {}
    }).then((response) => {
      this.dump = JSON.parse(JSON.stringify(response.data))

      for (var i = -3; i < 3; i++) {
        this.dates_list.push(moment().add(i, 'days').format('LLLL').split(' ').slice(0, 4).join(' ').replace(',', '').split(' '))
        this.dates_mmddyyyy.push(moment().add(i, 'days').format('MMDDYYYY'))
        this.dates_mongo.push(moment().add(i, 'days').format('YYYY-MM-DD'))
      }

      try {
        for (var n = 0; n < 6; n++) {
          let df = this.dump_filter_bydate(n).length
          if (isNaN(df) || typeof(df) == "undefined") {
            this.cases_count.push('0')
          } else {
            this.cases_count.push(df)          
          }
        }
      } catch (_) {}
    })
  }
}
</script>

<style lang="scss" scoped>
.md-gutter {
  width: 100%;
}

.md-subhead-2 {
  color: #91bafb;
  font-size: 16px;
  font-weight: 300;
}

.md-card {
  color:#5e9afa;
  display: block;
  height: 200px;
  margin: 0 auto;
  padding: 20px;
  position: relative;
  margin-top: 40px;
  width: 325px;
}

</style>
