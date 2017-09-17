/**
 * Created by nwu on 9/17/17.
 */
import React, {Component} from 'react';
import {Redirect} from 'react-router-dom';
import * as sessionActions from '../../session/actions';
import {connect} from 'react-redux';

class LogoutComponent extends Component {
  render() {
    this.props.logout();
    return <Redirect to="/" />
  }
}

function mapDispatchToProps(dispatch) {
  return {
    logout: sessionActions.logoutUser(dispatch)
  }
}

export default connect(null, mapDispatchToProps)(LogoutComponent);