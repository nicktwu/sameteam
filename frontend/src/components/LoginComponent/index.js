/**
 * Created by nwu on 9/17/17.
 */
import React, {Component} from 'react';
import {Alert, Button, Panel, Form, FormGroup, FormControl, Col, ControlLabel} from 'react-bootstrap';
import {connect} from 'react-redux';
import {Redirect} from 'react-router-dom';
import * as sessionActions from '../../session/actions';

class LoginComponent extends Component {
  constructor(props) {
    super(props);
    this.state = {
      username: "",
      password: "",
    };
    this.handleSubmit = this.handleSubmit.bind(this);
  }

  handleSubmit(event) {
    event.preventDefault();
    this.props.login(this.state);
  }

  render() {
    if (this.props.token && this.props.token.length > 0) {
      return <Redirect to="/profile"/>
    }
    if (sessionActions.isAuth()) {
      return <Redirect to="/profile"/>
    }
    return (
      <div className="panel-container">
        <Panel>
          <h1>Login</h1>
          { this.props.error ? <Alert bsStyle="danger">{this.props.error}</Alert> : null }
          <br/>
          <Form horizontal onSubmit={this.handleSubmit}>
            <FormGroup controlId="username">
              <Col componentClass={ControlLabel} sm={3}>
                Username
              </Col>
              <Col sm={9}>
                <FormControl type="text" onChange={event=>this.setState({"username":event.target.value})} value={this.state.username} placeholder="Username"/>
              </Col>
            </FormGroup>
            <FormGroup controlId="password">
              <Col componentClass={ControlLabel} sm={3}>
                Password
              </Col>
              <Col sm={9}>
                <FormControl type="password" onChange={event=>this.setState({"password":event.target.value})} value={this.state.password} placeholder="Password"/>
              </Col>
            </FormGroup>
            <FormGroup>
              <Col sm={12}>
                <Button type="submit">
                  Login
                </Button>
              </Col>
            </FormGroup>
          </Form>
        </Panel>
      </div>
    )
  }
}

function mapStateToProps(state) {
  return {
    token: state.session.token,
    error: state.session.error,
  }
}

function mapDispatchToProps(dispatch) {
  return {
    login: sessionActions.loginUser(dispatch)
  }
}

export default connect(mapStateToProps, mapDispatchToProps)(LoginComponent);