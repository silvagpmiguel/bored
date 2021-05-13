import axios from 'axios'

const API_URL='/api'

export async function fetchBored(){
    const response = await axios.get(`${API_URL}/bored`)
    return response.data
}