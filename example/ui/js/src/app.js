require("./style.css");
var W3CWebSocket = require('websocket').w3cwebsocket;
const backendPath = '/web/app/events'

window.onload = function () {
  var container = document.createElement("div");

  var topic = document.createElement("h1");
  topic.className = 'topic';
  topic.innerHTML = 'Hello Gotron / JS';

  container.appendChild(topic);

  //Show backend port
  var port = document.createElement("p");
  port.className = 'topic';
  port.innerHTML = `Backend Port: ${global.backendPort}`;
  container.appendChild(port);

  //websocket
  client = new W3CWebSocket('ws://127.0.0.1:' + global.backendPort + backendPath, [])
  

  client.onmessage = function(message){
    
    console.log(message);
    console.log(JSON.parse(message));
  };

  //button for websocket test
  var button = document.createElement("button");
  var t = document.createTextNode("click me");
  button.appendChild(t);
  button.onclick = function(){client.send(JSON.stringify({event: "hello", data: "Hello from frontend"}))};
  container.appendChild(button);

  document.body.appendChild(container);
}

//Reload on keypress 'r'
document.addEventListener('keyup', function (e) {
  if (e.keyCode == 82)
    window.location.reload();
})