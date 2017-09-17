/**
 * Created by nwu on 9/17/17.
 */
import React, {Component} from 'react';
import {Button, Panel, Form, FormGroup, FormControl, Col, ControlLabel} from 'react-bootstrap';
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
    console.log(this.state);
    this.props.login(this.state);
  }

  render() {
    if (sessionActions.isAuth()) {
      return <Redirect to="/profile"/>
    }
    return (
      <div className="panel-container">
        <Panel>
          <h1>Login</h1>
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

function mapDispatchToProps(dispatch) {
  return {
    login: sessionActions.loginUser(dispatch)
  }
}

export default connect(null, mapDispatchToProps)(LoginComponent);