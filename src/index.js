import React from 'react';
import ReactDOM from 'react-dom';

export default class Hello extends React.Component {
    render() {
        return (<p>Hello From the other side</p>)
    }
}


ReactDOM.render(
    <div><Hello /></div>, document.getElementById("app")
)
