import * as React from 'react';
import * as ReactDOM from "react-dom";


import { Root } from 'root';

const id = 'reactRoot'
let body = document.getElementsByTagName("body")[0];
let reactRoot = document.createElement("div");
reactRoot.id = id
body.appendChild(reactRoot);

ReactDOM.render(
  <Root />,
  document.getElementById(id)
);

//Reload on keypress 'r'
document.addEventListener('keyup', function(e){
if(e.keyCode == 82)
  window.location.reload();
})
