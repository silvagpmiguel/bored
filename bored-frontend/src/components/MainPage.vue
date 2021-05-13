<template>
  <div>
    <div class="panel panel-dark">
      <h3>Are you not Bored?</h3>
      <p>Don't worry, now there is a simple way to it</p>
      <div class="panel-component">
        <button @click="getBored()">Get Bored</button>
      </div>
    </div>
    <div v-show="json.activity" class="panel panel-dark">
        <p>{{ activity }}{{ url }}</p>
    </div>
  </div>
</template>

<script>
import { fetchBored } from '../services/service.js'
export default {
  name: 'MainPage',
  data() {
    return {
      json: {},
    }
  },
  methods: {
    async getBored() {
      try {
        const res = await fetchBored()
        this.json = res
      } catch (e) {
        console.error(e)
      }
    },
  },
  computed: {
    activity(){
      return this.json?.activity
    },
    url(){
      return this.json?.url
    }
  }
}
</script>

<style scoped>
.panel{
  margin-top: 7.5vh;
  padding-top: 1.5vh;
  padding-bottom: 1.5vh;
  width: 20vw;
  height: 100%;
  margin-left: auto;
  margin-right: auto;
}

.panel .panel-component{
  padding-top: 2.5vh;
  padding-bottom: 2.5vh;
}

.panel-dark{
  color: whitesmoke;
  background-color:slategray;
}

button{
  border: none;
  padding: 0.75rem 1rem;
  border-radius: 0.25rem;
  background-color: silver;
  color: whitesmoke;
  cursor: pointer;
}

button:hover{
  color: slategray;
  background-color: whitesmoke;
}
</style>
