<script setup>
import { ref, onMounted } from 'vue';
import { createNewTodo, updateTodo, getTodos, removeTodo } from './http';

const todoInput = ref('');
const todos = ref([]);
const updatTodoData = ref(null);

onMounted(() => {
    getTodos()
        .then(res => {
            todos.value = res.data.sort((a, b) => {
                return new Date(b.updated_at) - new Date(a.updated_at);
            });
        })
        .catch(console.log);
});

function handleCreateNewTodo() {
    createNewTodo({ text: todoInput.value })
        .then(res => {
            todos.value.push(res.data);
            todoInput.value = '';
        })
        .catch(console.log);
}

function handleUpdateTodo() {
    if (updatTodoData.value) {
        updateTodo(updatTodoData.value.id, { text: todoInput.value })
            .then(res => {
                const todoIndex = todos.value.findIndex(
                    td => td.id === res.data.id
                );
                todos.value[todoIndex] = res.data;
                todoInput.value = '';
            })
            .catch(console.log)
            .finally(() => {
                updatTodoData.value = null;
            });
    }
}

function handleRemoveTodo(id) {
    removeTodo(id)
        .then(res => {
            todos.value = todos.value.filter(it => it.id !== res.data.id);
        })
        .catch(console.log);
}

function editTodo(todo) {
    updatTodoData.value = todo;
    todoInput.value = todo.text;
}
</script>
<template>
    <main>
        <h1 class="title">Todo App build With VueJS And GoLang</h1>
        <form v-if="updatTodoData" @submit.prevent="handleUpdateTodo">
            <input v-model="todoInput" type="text" name="todo" />
            <button>Update</button>
        </form>
        <form v-else @submit.prevent="handleCreateNewTodo">
            <input v-model="todoInput" type="text" name="todo" />
            <button>Add</button>
        </form>

        <section>
            <ul>
                <li v-for="todo in todos" :key="todo.id">
                    <div class="todo-text">
                        <span class="text">
                            {{ todo.text }}
                        </span>
                        <div>
                            <button @click="editTodo(todo)">
                                <svg
                                    xmlns="http://www.w3.org/2000/svg"
                                    viewBox="-2.5 -2.5 24 24"
                                    width="18"
                                    fill="currentColor"
                                >
                                    <path
                                        d="M16.318 6.11l-3.536-3.535 1.415-1.414c.63-.63 2.073-.755 2.828 0l.707.707c.755.755.631 2.198 0 2.829L16.318 6.11zm-1.414 1.415l-9.9 9.9-4.596 1.06 1.06-4.596 9.9-9.9 3.536 3.536z"
                                    ></path>
                                </svg>
                            </button>
                            <button @click="handleRemoveTodo(todo.id)">
                                <svg
                                    xmlns="http://www.w3.org/2000/svg"
                                    viewBox="-3 -2 24 24"
                                    width="18"
                                    fill="currentColor"
                                >
                                    <path
                                        d="M6 2V1a1 1 0 0 1 1-1h4a1 1 0 0 1 1 1v1h4a2 2 0 0 1 2 2v1a2 2 0 0 1-2 2h-.133l-.68 10.2a3 3 0 0 1-2.993 2.8H5.826a3 3 0 0 1-2.993-2.796L2.137 7H2a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h4zm10 2H2v1h14V4zM4.141 7l.687 10.068a1 1 0 0 0 .998.932h6.368a1 1 0 0 0 .998-.934L13.862 7h-9.72zM7 8a1 1 0 0 1 1 1v7a1 1 0 0 1-2 0V9a1 1 0 0 1 1-1zm4 0a1 1 0 0 1 1 1v7a1 1 0 0 1-2 0V9a1 1 0 0 1 1-1z"
                                    ></path>
                                </svg>
                            </button>
                        </div>
                    </div>
                    <div class="todo-info">
                        <span>
                            Created At:
                            {{ new Date(todo.created_at).toDateString() }}</span
                        >
                        <span>
                            Updated At:
                            {{ new Date(todo.updated_at).toDateString() }}</span
                        >
                    </div>
                </li>
            </ul>
        </section>
    </main>
</template>

<style>
@import url('https://fonts.googleapis.com/css2?family=Poppins:wght@400;600&display=swap');
body {
    margin: 0;
    background: rgb(6, 28, 68);
    color: rgba(255, 255, 255, 0.741);
    font-family: 'Poppins', sans-serif;
}
button,
input,
a {
    color: initial;
    font-family: inherit;
    font-size: inherit;
}
.title {
    font-size: 40px;
    text-align: center;
    margin: 50px 0;
}
form {
    padding: 10px;
    border-radius: 3px;
    width: 500px;
    margin: auto;
    display: flex;
}
form input {
    display: block;
    padding: 12px;
    border: none;
    background: rgba(255, 255, 255, 0.892);
    border-radius: 3px;
    flex-grow: 1;
    margin-right: 10px;
}
form button {
    width: 100px;
    display: block;
    cursor: pointer;
    background-color: rgb(7, 211, 133);
    color: white;
    border: none;
    border-radius: 3px;
}
ul {
    list-style: none;
    padding: 20px;
    width: 500px;
    margin: auto;
}
li {
    padding: 10px;
    background-color: rgb(6, 169, 160);
    border-radius: 4px;
    margin: 10px;
    /* display: flex; */
}
ul li button {
    font-size: 10px;
    display: flex;
    justify-content: center;
    align-items: center;
    cursor: pointer;
    border: none;
    border-radius: 4px;
    color: #fff;
}

ul li button svg {
    display: inline-block;
}

.todo-text {
    display: flex;
    justify-content: space-between;
}

.todo-text .text {
    font-weight: bold;
    font-size: 18px;
}

.todo-text > div {
    display: flex;
}
.todo-text button:first-child {
    background-color: rebeccapurple;
}
.todo-text button:last-child {
    background-color: rgb(199, 0, 0);
}

.todo-text button:not(:last-child) {
    margin-right: 5px;
}

.todo-info {
    font-size: 12px;
}
.todo-info span {
    margin-left: 20px;
}
</style>
