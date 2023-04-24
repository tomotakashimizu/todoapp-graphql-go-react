import React, { useState } from 'react';
import { useMutation } from '@apollo/client';
import { CreateTodoDocument } from '../types/gen/api';

function CreateTodoForm() {
  const [text, setText] = useState('');
  const [createTodo] = useMutation(CreateTodoDocument);

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
