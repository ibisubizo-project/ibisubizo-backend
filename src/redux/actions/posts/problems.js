import api from '../../../modules/api/api';
import { getVisibleProblems, getVisibleProblemsFailure } from './actionCreators';

const fetchProblemListings = () => {
    return dispatch => api.GetProblemListing()
        .then((response) => {
            console.dir(response);
        }).catch((error) => {
            console.error(error);
        });
}

export const fetchingProblems = () => {
    return dispatch => {
        dispatch(getVisibleProblems());
        api.GetProblemListing().then((response) => {
            dispatch(getVisibleProblems(response));
        })
        .catch((error) => {
            dispatch(getVisibleProblemsFailure(error));
        });
    }
}

export default fetchProblemListings