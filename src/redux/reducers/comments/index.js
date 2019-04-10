import { comment } from "postcss-selector-parser";


const commentsReducer = (state, action) => {
    switch(action.type) {
        case 'ADD_COMMENT':
            return state
        case 'DELETE_COMMENT':
            return state
        case 'EDIT_COMMENT':
            return state
        default:
            return state
    }
}

export default commentsReducer;