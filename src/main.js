import Vue from 'vue'
import App from './App.vue'
import router from './router'
import firebase from 'firebase'

Vue.config.productionTip = false

const firebaseConfig = {
  apiKey: "AIzaSyCZKqKEVc8wSMrzgrfk0SuU7EA4qGQY-xg",
  authDomain: "vue-go-firebase-950a3.firebaseapp.com",
  databaseURL: "https://vue-go-firebase-950a3.firebaseio.com",
  projectId: "vue-go-firebase-950a3",
  storageBucket: "vue-go-firebase-950a3.appspot.com",
  messagingSenderId: "659365831211",
  appId: "1:659365831211:web:0c1a24f70a66c51f91006e"
};
firebase.initializeApp(firebaseConfig)

let app

firebase.auth().onAuthStateChanged(() => {
  /* eslint-disable no-new */
  if (!app) {
    new Vue({
      el: '#app',
      router,
      components: { App },
      template: '<App/>'
    })
  }
})