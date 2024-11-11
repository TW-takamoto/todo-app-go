import React, { useState } from 'react';
import { PlusCircle } from 'lucide-react';

interface TodoFormProps {
  onSubmit: (text: string) => void;
}

export const TodoForm: React.FC<TodoFormProps> = ({ onSubmit }) => {
  const [input, setInput] = useState('');

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    if (input.trim()) {
      onSubmit(input.trim());
      setInput('');
    }
  };

  return (
    <form onSubmit={handleSubmit} className="mb-6">
      <div className="flex gap-2">
        <input
          type="text"
          value={input}
          onChange={(e) => setInput(e.target.value)}
          placeholder="新しいタスクを入力..."
          className="flex-1 rounded-lg border-gray-300 shadow-sm focus:border-purple-500 focus:ring-purple-500 px-4 py-2"
        />
        <button
          type="submit"
          className="inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-lg shadow-sm text-white bg-purple-600 hover:bg-purple-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-purple-500"
        >
          <PlusCircle className="w-5 h-5" />
        </button>
      </div>
    </form>
  );
};