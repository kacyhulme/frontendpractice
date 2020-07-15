require("expose-loader?$!expose-loader?jQuery!jquery");
require("bootstrap/dist/js/bootstrap.bundle.js");
require('jquery-ui-bundle');
import React from "react";
import ReactDOM from "react-dom";

$(() => {

  $(function () {
    $("#post-Date").datepicker();
  });

  const e = React.createElement;

  class LikeButton extends React.Component {
    constructor(props) {
      super(props);
      this.state = { liked: false };
    }

    render() {
      if (this.state.liked) {
        return 'You liked this.';
      }

      return e(
        'button',
        { onClick: () => this.setState({ liked: true }) },
        'Like'
      );
    }
  }
  const domContainer = document.querySelector('#practice-container');
  ReactDOM.render(e(LikeButton), domContainer);

});
