require("./style.css");

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

  document.body.appendChild(container);
}

//Reload on keypress 'r'
document.addEventListener('keyup', function (e) {
  if (e.keyCode == 82)
    window.location.reload();
})