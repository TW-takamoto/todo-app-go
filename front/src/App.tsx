import React from 'react';
import { useQuery, useMutation } from '@apollo/client';
import { GET_TODOS, ADD_TODO, TOGGLE_TODO, DELETE_TODO } from './graphql/queries';
import { TodoItem } from './components/TodoItem';
import { TodoForm } from './components/TodoForm';
import { Todo } from './types';

function App() {
  const { loading, error, data } = useQuery(GET_TODOS);
  const [addTodo] = useMutation(ADD_TODO, {
    refetchQueries: [{ query: GET_TODOS }],
  });
  const [toggleTodo] = useMutation(TOGGLE_TODO);
  const [deleteTodo] = useMutation(DELETE_TODO, {
    refetchQueries: [{ query: GET_TODOS }],
  });

  const handleAddTodo = async (text: string) => {
    try {
      await addTodo({ variables: { text } });
    } catch (err) {
      console.error('Error adding todo:', err);
    }
  };

  const handleToggleTodo = async (id: string) => {
    try {
      await toggleTodo({ variables: { id } });
    } catch (err) {
      console.error('Error toggling todo:', err);
    }
  };

  const handleDeleteTodo = async (id: string) => {
    try {
      await deleteTodo({ variables: { id } });
    } catch (err) {
      console.error('Error deleting todo:', err);
    }
  };

  if (loading) {
    return (
      <div className="min-h-screen bg-gradient-to-br from-purple-50 to-blue-50 flex items-center justify-center">
        <div className="text-xl text-gray-600">読み込み中...</div>
      </div>
    );
  }

  if (error) {
    return (
      <div className="min-h-screen bg-gradient-to-br from-purple-50 to-blue-50 flex items-center justify-center">
        <div className="text-xl text-red-600">エラーが発生しました: {error.message}</div>
      </div>
    );
  }

  return (
    <div className="min-h-screen bg-gradient-to-br from-purple-50 to-blue-50 py-12 px-4 sm:px-6 lg:px-8">
      <div className="max-w-md mx-auto">
        <div className="bg-white rounded-xl shadow-lg overflow-hidden">
          <div className="px-6 py-8">
            <h1 className="text-3xl font-bold text-gray-900 text-center mb-8">
              TODOリスト
            </h1>

            <TodoForm onSubmit={handleAddTodo} />

            <div className="space-y-3">
              {data.todos.map((todo: Todo) => (
                <TodoItem
                  key={todo.id}
                  todo={todo}
                  onToggle={handleToggleTodo}
                  onDelete={handleDeleteTodo}
                />
              ))}
              {data.todos.length === 0 && (
                <p className="text-center text-gray-500 py-4">
                  タスクがありません。新しいタスクを追加してください。
                </p>
              )}
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}

export default App;