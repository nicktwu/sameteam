import React, { Component } from 'react';
import {HeaderBar} from './components/HeaderBar/index';
import './App.css';
import {BrowserRouter, Route} from 'react-router-dom';
import pageList from './pages';

class App extends Component {
  render() {
    return (
      <BrowserRouter>
        <div>
          <HeaderBar pages={pageList}/>
          { pageList.map((page, index) => {
            if (page.exact) {
              return <Route exact path={page.url} component={page.comp}/>
            } else {
              return <Route path={page.url} component={page.comp}/>
            }
          })}

        </div>
      </BrowserRouter>
    );
  }
}

export default App;
