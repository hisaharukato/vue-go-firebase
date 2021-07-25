import Vue from 'vue'
import App from './App.vue'
import router from './router'
import firebase from 'firebase'

Vue.config.productionTip = false

const firebaseConfig = {
  apiKey: "AIzaSyDWYThry1dULbYDTHseaGnV3LPuN1zOQcI",
  authDomain: "vue-golang-a3fb7.firebaseapp.com",
  databaseURL: "https://vue-golang-a3fb7.firebaseio.com",
  projectId: "vue-golang-a3fb7",
  storageBucket: "vue-golang-a3fb7.appspot.com",
  messagingSenderId: "26946518503",
  appId: "1:26946518503:web:d36ab36cfb3bc390c09142"
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