/**
 * Created by nwu on 9/17/17.
 */
import React, {Component} from 'react';
import {Panel} from 'react-bootstrap';
import {connect} from 'react-redux';

class ProfileComponent extends Component {
  render() {
    return (
      <div className="panel-container">
        <Panel>
          <h1>Hi there{this.props.user ? ", "+this.props.user.username : null}!</h1>
          <p>This is an internal profile page</p>
        </Panel>
      </div>
    )
  }
}

function mapStateToProps(state) {
  return {
    user: state.session.user,
  }
}

export default connect(mapStateToProps, null)(ProfileComponent);
