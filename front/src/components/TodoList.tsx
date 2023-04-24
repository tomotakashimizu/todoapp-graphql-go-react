import React from 'react';
import { useGetAllTodosQuery } from '../types/gen/api';

export const TodoList: React.FC = () => {
  const { data, loading, error } = useGetAllTodosQuery();

  if (loading) return <p>Loading...</p>;
  if (error) return <p>Error: {error.message}</p>;

  return (
    <div>
      <h2>Todoリスト</h2>
      <ul>
        {data?.todos.map((todo) => (
          <li key={todo.id}>{todo.text}</li>
        ))}
      </ul>
    </div>
  );
};
