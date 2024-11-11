import { CodegenConfig } from '@graphql-codegen/cli';

const config: CodegenConfig = {
  schema: 'http://localhost:8080/query',  // GoのGraphQLサーバーのスキーマURL
  documents: ['src/**/*.tsx', 'src/**/*.ts'],
  generates: {
    './src/generated/': {
      preset: 'client',
      plugins: [
        'typescript',
        'typescript-operations',
        'typescript-react-apollo'
      ],
    },
  },
};

export default config;