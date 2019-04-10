import { initialState } from "..";


const problemReducer  = (state = initialState, action) => {
    switch(action.type) {
        case "GET_VISIBLE_PROBLEMS":
        case "GET_VISIBLE_PROBLEMS_SUCCESS":
            return {
                ...state,
                isFetchingData: true
            }
        case "GET_VISIBLE_PROBLEMS_FAILURE": 
            return {
                ...state,
                isFetchingData: true,
                //error: payload.error
            }
        default:
            return state
    }
}


export default problemsReducer;