/**
 * Created by nwu on 9/16/17.
 */
import React, {Component} from 'react';
import backgroundPicture from './roommate.jpg'
import './home.css'

export default class HomeComponent extends Component {
  render() {
    return (
      <div>
      <div className="home-background" style={{backgroundImage:"url("+backgroundPicture+")", backgroundSize:"cover"}}/>
      <div className="home-content">
      	<div>
      	  <h1>Find Your Perfect Roommate</h1>
      	  <h2>Start Your Matching Process</h2>
      	</div>
      </div>
      </div>
    )
  }
}