require("./style.css");

window.onload = function () {
  var topic = document.createElement("h1");
  topic.className = 'topic';
  topic.innerHTML = 'Hello Gotron / JS';

  document.body.appendChild(topic);
}

//Reload on keypress 'r'
document.addEventListener('keyup', function (e) {
  if (e.keyCode == 82)
    window.location.reload();
})