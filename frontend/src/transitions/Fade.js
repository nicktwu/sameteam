/**
 * Created by nwu on 9/17/17.
 */
import React from 'react';
import ReactCSSTransitionGroup from 'react-addons-css-transition-group';

export const Fade = Content => (
  props => (
    <ReactCSSTransitionGroup transitionAppear={true}
                      transitionAppearTimeout={1000}
                      transitionEnterTimeout={1000}
                      transitionLeaveTimeout={1000}
                      transitionName="Fade">
      <Content {...props}/>
    </ReactCSSTransitionGroup>
  )
);