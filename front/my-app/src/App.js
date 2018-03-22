import React  from 'react';
import './App.css';

import { BrowserRouter , Route} from 'react-router-dom'

import Weather from './Weather.js';

class App extends  React.Component  {
  render() {
    return (
      <BrowserRouter>
        <div>
          <Route exact path="/" component={Info}/>
          <Route path="/:city" component={Weather} />
        </div>
      </BrowserRouter>
    );
  }
}

const Info = ({ match }) => {
  return <h2>Your should type a city<br/> Exemple : <a href="http://localhost:3000/Paris">Paris</a></h2>
}

/*
const testweather = {{ match }} => {
  <div>
    {match.params.city}
  </div>
}
*/
export default App;
