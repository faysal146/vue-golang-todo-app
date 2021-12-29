<script setup>
import { ref } from 'vue';

const todo = ref('');
function handleSubmit() {
    console.log(todo.value);
}
</script>
<template>
    <main>
        <h1 class="title">Todo App build With VueJS And GoLang</h1>
        <form @submit.prevent="handleSubmit">
            <input v-model="todo" type="text" name="todo" />
            <button>Add Todo</button>
        </form>
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
</style>
