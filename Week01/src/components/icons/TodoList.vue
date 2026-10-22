<script setup>
import {ref} from 'vue';

let id = 0;
const todos = ref([
    {id: id++, work: "read book"},
    {id: id++, work: "code"},
    {id: id++, work: "sleep"},
]);

const newTodo = ref("");

const addTodo = () => {
    let todoOb = {
        id: id++,
        work: newTodo.value,

    };
    todos.value.push(todoOb);
};

const delTodo = (todo) => {
    console.log(todo);
    todos.value = todos.value.filter((td)=>td != todo)

}
</script>

<template>
    <h1> TODO List</h1>
    <input v-model = "newTodo" />
    <button @click = "addTodo()"> Add </button> 
    <ul>
        <li v-for = "todo in todos" :key = "todo.id">
        <h2> {{  todo.work }}</h2>
        <button @click="delTodo(todo)" class = "close">X</button>
        </li>
    </ul>
</template> 