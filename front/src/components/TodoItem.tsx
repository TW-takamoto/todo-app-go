import React from 'react';
import { CheckCircle2, XCircle, Trash2 } from 'lucide-react';
import { Todo } from '../types';

interface TodoItemProps {
  todo: Todo;
  onToggle: (id: string) => void;
  onDelete: (id: string) => void;
}

export const TodoItem: React.FC<TodoItemProps> = ({ todo, onToggle, onDelete }) => {
  return (
    <div
      className={`flex items-center justify-between p-4 rounded-lg transition-all ${
        todo.completed ? 'bg-green-50' : 'bg-gray-50'
      }`}
    >
      <div className="flex items-center gap-3 flex-1">
        <button
          onClick={() => onToggle(todo.id)}
          className={`focus:outline-none ${
            todo.completed ? 'text-green-500' : 'text-gray-400'
          }`}
        >
          {todo.completed ? <CheckCircle2 /> : <XCircle />}
        </button>
        <span
          className={`flex-1 text-gray-700 ${
            todo.completed ? 'line-through text-gray-400' : ''
          }`}
        >
          {todo.text}
        </span>
      </div>
      <button
        onClick={() => onDelete(todo.id)}
        className="text-red-500 hover:text-red-600 focus:outline-none"
      >
        <Trash2 className="w-5 h-5" />
      </button>
    </div>
  );
};