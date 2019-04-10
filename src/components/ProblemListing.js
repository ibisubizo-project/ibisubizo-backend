import React from 'react';
import { connect } from 'react-redux';
import { bindActionCreators } from "redux";
import fetchProblemListings, { fetchingProblems } from '../redux/actions/posts/problems'; 
import SinglePost from './SinglePost';




// class ProblemListing extends React.Component {
//     componentDidMount() {
//         //this.props.dispatch.getVisibleListings()
//         //console.log(this.props.fetchProblemListings());
//     }

//     render() {
//         // const { visibleListings } = this.props;

//         // if (error) {
//         //     return <div>Error! {error.message}</div>;
//         // }

//         // if (loading) {
//         //     return <div>Loading...</div>;
//         // }


//         return (
//             <div className="container mx-auto px-6 py-24">
//                 <div className="flex items-baseline justify-between border-b-2 border-grey-light">
//                     <span className="font-display font-bold tracking-wide uppercase py-4 border-b-2 border-indigo" style={{margin: '-2px'}}>Trending Problems</span>
//                     <div className="problems__categories flex">
//                         <span className="mr-4">Latest Problems</span>
//                         <span>My Problems</span>
//                     </div>
//                 </div>


//                 <div className="post_listings">
//                 <SinglePost />
//                     {/* {problems.map(problem => {
//                         <SinglePost />
//                     })} */}
//                 </div>

//              </div>
//         )
//     }
// }


const ProblemListing = ({problems, getVisibleListings}) => {
    //console.dir(problems());
    return (
        <div>
            <SinglePost />>
        </div>
    )
}

const mapStateToProps = state => ({
  //problems: state
});

const mapDispatchToProps = dispatch => {
    //return bindActionCreators({ fetchProblemListings }, dispatch)
    return {
        //getVisibleListings: () => dispatch({ type: 'GET_VISIBLE_LISTINGS'}),
        //problems : () => dispatch(fetchingProblems())
    }
};

export default connect(mapStateToProps, mapDispatchToProps)(ProblemListing);