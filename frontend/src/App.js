import React, { Component } from "react";
import {HeaderBar} from "./components/HeaderBar/index";
import './App.css';
import BrowserRouter from "react-router-dom/BrowserRouter";
import Route from "react-router-dom/Route";
import pageList from './pages';
import {Fade} from './transitions/Fade';
import backgroundPicture from './roommate.jpg'
import {Switch} from "react-router-dom";

class App extends Component {
  render() {
    return (
      <BrowserRouter>
        <div className="App">
          <HeaderBar pages={pageList}/>
          <div className="main-content">
            <div className="background" style={{background:"url("+backgroundPicture+") no-repeat center center fixed", backgroundSize:"cover"}}/>
            <Switch>
            { pageList.map((page, index) => {
              if (page.exact) {
                return (
                  <Route exact key={index} path={page.url} component={Fade(page.comp)}/>
                )
              } else {
                return (
                  <Route key={index} path={page.url} component={Fade(page.comp)}/>
                )
              }
            })}
            </Switch>
          </div>
        </div>
      </BrowserRouter>
    );
  }
}

export default App;
