import React from 'react';
import NavBar from './NavBar';
import ProblemForm from '../components/ProblemForm';
import ProblemListing from '../components/ProblemListing';


const Main = () => {
    return (
        <div className="main__component">
            <NavBar />

            <div className="container mx-auto mt-10">
                <ProblemForm />
                <ProblemListing />
            </div>
        </div>
    )
}

export default Main;