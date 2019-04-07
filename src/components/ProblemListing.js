import React from 'react';
import SinglePost from './SinglePost';


const ProblemListing = () => {
    return (
        <div className="container mx-auto px-6 py-24">
            <div className="flex items-baseline justify-between border-b-2 border-grey-light">
                <span className="font-display font-bold tracking-wide uppercase py-4 border-b-2 border-indigo" style={{margin: '-2px'}}>Trending Problems</span>
                <div className="problems__categories flex">
                    <span className="mr-4">Latest Problems</span>
                    <span>My Problems</span>
                </div>
            </div>


            <div className="post_listings">
                <SinglePost />
                <SinglePost />
                <SinglePost />
                <SinglePost />
                <SinglePost />
            </div>

        </div>
    )
}

export default ProblemListing;