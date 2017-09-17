/**
 * Created by nwu on 9/16/17.
 */
import React, {Component} from 'react';
import './home.css'
import {Link} from 'react-router-dom';
import {Button} from 'react-bootstrap';

export default class HomeComponent extends Component {
  render() {
    return (
      <div>
        <div className="home-content">
      	  <div>
      	    <h1>Find Your Perfect Roommate</h1>
              <h2>Moving into a new apartment or dorm? Want a dank ass roommate? We got you covered!</h2>
      	    <Link to="/signup"><Button>Start Your Matching Process</Button></Link>
      	  </div>
        </div>
      </div>
    )
  }
}