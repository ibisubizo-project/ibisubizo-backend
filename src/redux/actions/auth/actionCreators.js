import { AUTHENTICATE_USER, UNAUTHENTICATE_USER } from "../actions";


export const authenticateUser = (user) => ({
    type: AUTHENTICATE_USER,
    user
});

export const unauthenticateuser = (user) => ({
    type: UNAUTHENTICATE_USER,
    user
});

