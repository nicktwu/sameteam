/**
 * Created by nwu on 9/16/17.
 */
import React, {Component} from 'react';
import './header.css'
import {LinkContainer, IndexLinkContainer} from 'react-router-bootstrap';
import {Navbar, Nav, NavItem} from 'react-bootstrap';
import {isAuth} from "../../session/actions"

export class HeaderBar extends Component {

  render() {
    let auth = isAuth();
    return (
      <Navbar inverse>
        <Navbar.Header>
          <Navbar.Brand>
            SameTeam
          </Navbar.Brand>
          <Navbar.Toggle/>
        </Navbar.Header>
        <Navbar.Collapse>
          <Nav pullRight>
            {this.props.pages.map((page, index) => {
              if (page.exact) {
                return <IndexLinkContainer to={page.url} key={index}><NavItem>{page.name}</NavItem></IndexLinkContainer>
              }
              if (page.show) {
                return <LinkContainer to={page.url} key={index}><NavItem>{page.name}</NavItem></LinkContainer>
              }
              if ((page.auth && !auth) || (auth && !page.auth)) {
                return null;
              }
              return <LinkContainer to={page.url} key={index}><NavItem>{page.name}</NavItem></LinkContainer>
            })}
          </Nav>
        </Navbar.Collapse>
      </Navbar>
    )
  }
}