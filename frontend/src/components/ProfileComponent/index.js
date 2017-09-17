/**
 * Created by nwu on 9/17/17.
 */
import React, {Component} from 'react';
import {Panel} from 'react-bootstrap';

export default class LoginComponent extends Component {
  render() {
    return (
      <div className="panel-container">
        <Panel>
          <h1>Hi there!</h1>
          <p>This is an internal profile page</p>
        </Panel>
      </div>
    )
  }
}