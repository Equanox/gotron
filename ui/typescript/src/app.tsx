import * as React from 'react';
import * as ReactDOM from "react-dom";

// https://github.com/zilverline/react-tap-event-plugin
import * as injectTapEventPlugin from 'react-tap-event-plugin';
injectTapEventPlugin();

import { Root } from 'root';

ReactDOM.render(
  <Root />,
  document.getElementsByTagName("body")[0]
);

//Reload on keypress 'r'
document.addEventListener('keyup', function(e){
if(e.keyCode == 82)
  window.location.reload();
})
