import Vue from 'vue'

import App from './App'

window.onload = function () {
  var topic = document.createElement("div");
  topic.id = 'vue-app';
  document.body.insertBefore(topic, document.body.childNodes[0])  

  new Vue({
    el: '#vue-app',
    components: { App },
    template: '<App/>'
  })
}

//Reload on keypress 'r'
document.addEventListener('keyup', function(e){
if(e.keyCode == 82)
window.location.reload();
})

