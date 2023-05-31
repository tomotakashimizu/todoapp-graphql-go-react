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
      <h3>未完了のTODO</h3>
      <ul>
        {data?.todos.map(
          (todo) =>
            todo.done || (
              <li key={todo.id}>
                {todo.text}
                <button onClick={() => handleDeleteTodo(todo.id)}>削除</button>
              </li>
            )
        )}
      </ul>

      <h3>完了したTODO</h3>
      <ul>
        {data?.todos.map(
          (todo) =>
            todo.done && (
              <li key={todo.id}>
                {todo.text}
                <button onClick={() => handleDeleteTodo(todo.id)}>削除</button>
              </li>
            )
        )}
      </ul>
    </div>
  );
};
