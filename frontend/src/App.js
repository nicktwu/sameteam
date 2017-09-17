import React, { Component } from 'react';
import {HeaderBar} from './components/HeaderBar/index';
import './App.css';
import {BrowserRouter, Route} from 'react-router-dom';
import pageList from './pages';

class App extends Component {
  render() {
    return (
      <BrowserRouter>
        <div className="App">
          <HeaderBar pages={pageList}/>
          <div className="main-content">
          { pageList.map((page, index) => {
            if (page.exact) {
              return <Route exact key={index} path={page.url} component={page.comp}/>
            } else {
              return <Route key={index} path={page.url} component={page.comp}/>
            }
          })}
          </div>

        </div>
      </BrowserRouter>
    );
  }
}

export default App;
