import React from 'react';
import './App.css';
import { TodoList } from './components/TodoList';
import CreateTodoForm from './components/CreateTodoForm';

function App() {
  return (
    <div className='App'>
      <h1>Todo App</h1>
      <CreateTodoForm />
      <TodoList />
    </div>
  );
}

export default App;
