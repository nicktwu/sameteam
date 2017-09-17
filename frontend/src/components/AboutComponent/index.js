/**
 * Created by nwu on 9/16/17.
 */
import React, {Component} from 'react';
import './about.css';
import {Panel} from 'react-bootstrap'

export default class AboutComponent extends Component {
  render() {
    return (
      <div className="panel-container">
        <Panel>
          <h1>About Us</h1>
          <p>We're just a bunch of idiots clowning around at a hackathon. Specifically, we came up with this idea for HackMIT 2017, which we weren't really taking seriously.</p>
        </Panel>
      </div>
    )
  }
}