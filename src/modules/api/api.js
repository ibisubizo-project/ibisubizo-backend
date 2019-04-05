import axios from "axios"

let API_URL = "http://localhost:8000/api"

let api = {
    Login : (form) => {
        return axios.post(`${API_URL}/auth/login`, form, { headers: {'Accept': 'application/json'} })
    },
    Register: (credentials) => {
        return axios.post(`${API_URL}/auth/register`, credentials, { headers: {'Accept': 'application/json'} })
    },
    GetUserById: (user_id) => {
        return axios.get(`${API_URL}/users/id/${user_id}`)
    }
}


export default api