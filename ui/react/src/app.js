var React = require('react');
var ReactDOM = require('react-dom');

require("./style.css");

class Root extends React.Component {
  render() {
    return(
      <h1 className="topic">
        Hello Gotron / React
      </h1>
    )
  }
}

window.onload = function () {
    var div = document.createElement("div");
    ReactDOM.render(<Root />, div);
    document.body.appendChild(div);
}

//Reload on keypress 'r'
document.addEventListener('keyup', function(e){
if(e.keyCode == 82)
  window.location.reload();
})
