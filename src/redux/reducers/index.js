


export const initialState = {
    isFetchingData: false,
    isAuthenticated: false,
    authenticatedUser: {},
    error: '',
    problems: [],
    comments: []
};

function rootReducer(state = initialState, action) {
    return state
}

export default rootReducer