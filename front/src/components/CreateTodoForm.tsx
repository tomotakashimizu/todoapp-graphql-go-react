import React, { useState } from 'react';
import { useMutation } from '@apollo/client';
import { CreateTodoDocument, GetAllTodosDocument, GetAllTodosQuery } from '../types/gen/api';

function CreateTodoForm() {
  const [text, setText] = useState('');
  const [createTodo] = useMutation(CreateTodoDocument, {
    // キャッシュを更新し、そのキャッシュを参照しているすべてのコンポーネントに通知し、新しいデータをもとに再レンダリング=>UIが自動的に更新
    update(cache, { data: { createTodo } }) {
      const existingTodos = cache.readQuery<GetAllTodosQuery>({ query: GetAllTodosDocument });

      if (existingTodos && createTodo) {
        cache.writeQuery({
          query: GetAllTodosDocument,
          data: { todos: [...existingTodos.todos, createTodo] },
        });
      }
    },
  });

  const handleSubmit = async (event: { preventDefault: () => void }) => {
    event.preventDefault();

    if (!text.trim()) return;

    try {
      const input = { text };
      await createTodo({ variables: { input } });

      setText('');
    } catch (error) {
      console.error('Error adding todo: ', error);
    }
  };

  return (
    <form onSubmit={handleSubmit}>
      <input type='text' value={text} onChange={(event) => setText(event.target.value)} placeholder='Add new todo' />
      <button type='submit'>Add</button>
    </form>
  );
}

export default CreateTodoForm;
