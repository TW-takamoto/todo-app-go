import { gql } from '@apollo/client';

export const GET_TODOS = gql`
  query GetTodos {
    todos {
      id
      text
      completed
    }
  }
`;

export const ADD_TODO = gql`
  mutation AddTodo($text: String!) {
    createTodo(input: { text: $text }) {
      id
      text
      completed
    }
  }
`;

export const TOGGLE_TODO = gql`
  mutation ToggleTodo($id: ID!) {
    toggleTodo(id: $id) {
      id
      completed
    }
  }
`;

export const DELETE_TODO = gql`
  mutation DeleteTodo($id: ID!) {
    deleteTodo(id: $id)
  }
`;