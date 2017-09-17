/**
 * Created by nwu on 9/16/17.
 */
import React, {Component} from 'react';
import './header.css'

export class HeaderBar extends Component {

  render() {
    return (
      <div className="header-bar">
        <ul>
          {this.props.pages.map((page, index) => {
            return <li key={index}>{page.name}</li>
          })}
        </ul>
      </div>
    )
  }
}