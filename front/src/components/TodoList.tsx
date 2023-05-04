import React from 'react';
import { DeleteTodoDocument, useGetAllTodosQuery } from '../types/gen/api';
import { useMutation } from '@apollo/client';

export const TodoList: React.FC = () => {
  const { data, loading, error, refetch } = useGetAllTodosQuery();
  const [deleteTodo] = useMutation(DeleteTodoDocument);

  const handleDeleteTodo = async (todoId: string) => {
    try {
      await deleteTodo({ variables: { todoId } });
      refetch();
    } catch (error) {
      console.error('Error deleting todo: ', error);
    }
  };

  if (loading) return <p>Loading...</p>;
  if (error) return <p>Error: {error.message}</p>;

  return (
    <div>
      <h2>Todoリスト</h2>
      <ul>
        {data?.todos.map((todo) => (
          <li key={todo.id}>
            {todo.text}
            <button onClick={() => handleDeleteTodo(todo.id)}>削除</button>
          </li>
        ))}
      </ul>
    </div>
  );
};
