import axios from 'axios';

const apiBaseUrl = 'http://localhost:3000/api/v1';

export const getAll = async () => {
    const { data } = await axios.get(`${apiBaseUrl}/products`); 
    return data;
}

export const get = async (name) => {
    const { data } = await axios.get(`${apiBaseUrl}/products/${name}`);   
    return data;
}

export const create = async (product) => {
    const { data } = await axios.post(`${apiBaseUrl}/products`, product);   
    return data;
}

export const update = async (name, product) => {
    const { data } = await axios.put(`${apiBaseUrl}/products/${name}`, product);
    return data;
}

export const remove = async (name) => {
    const { data } = await axios.delete(`${apiBaseUrl}/products/${name}`);
    return data;
}

