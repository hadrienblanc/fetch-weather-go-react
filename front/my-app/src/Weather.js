import React from 'react';

class Wheather extends React.Component {
    constructor(props) {
        super(props);
        this.state = {color: "#d1d1d1", loading: "loading"};
    }

    componentDidMount() {
        fetch("http://localhost:8080/" + this.props.match.params.city)
        .then(response => response.json())
        .then(myJson => this.setState({ loading: "", color: myJson.color , temperature: myJson.temperature}))
        .catch(error => console.error(error));
    }

    componentWillUnmount() {

    }

    render() {
        return (
            <div id="weather" style={{backgroundColor: this.state.color}}>
                <h1>{this.props.match.params.city}</h1>
                <p>{this.state.temperature}° celsius</p>
            </div>
        );
    }
}

export default Wheather;
