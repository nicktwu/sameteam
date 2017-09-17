/**
 * Created by nwu on 9/16/17.
 */
import React, {Component} from 'react';
import './header.css'
import {Link} from 'react-router-dom';

export class HeaderBar extends Component {

  render() {
    return (
      <div className="header-bar">
        <ul>
          {this.props.pages.map((page, index) => {
            return <Link to={page.url} key={index}><li>{page.name}</li></Link>
          })}
        </ul>
      </div>
    )
  }
}