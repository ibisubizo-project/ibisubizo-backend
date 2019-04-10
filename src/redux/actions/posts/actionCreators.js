

export function getVisibleProblems() {
    return {
        type: 'GET_VISIBLE_PROBLEMS',
        isFetchingData: true
    }
}

export function getVisibleProblemsSuccess(response) {
    console.log(response)
    return {
        type: 'GET_VISIBLE_PROBLEMS_SUCCESS',
        isFetchingData: false,
        problems: response.data
    }
}

export function getVisibleProblemsFailure(error) {
    return {
        type: 'GET_VISIBLE_PROBLEMS_FAILURE',
        isFetchingData: false,
        error : error.message
    }
}