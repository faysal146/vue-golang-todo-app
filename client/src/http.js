import axios from 'axios';

const axiosIns = axios.create({ baseURL: 'http://127.0.0.1:8000/api/v1' });

export const getTodos = () => axiosIns.get('/todos');
export const createNewTodo = val =>
    axiosIns.post('/todos', JSON.stringify(val));

export const updateTodo = (id, val) =>
    axiosIns.patch('/todos/' + id, JSON.stringify(val));

export const removeTodo = id => axiosIns.delete(`/todos/${id}`);
