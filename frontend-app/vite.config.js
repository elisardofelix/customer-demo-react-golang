import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'
import reactRouterDom from 'react-router-dom'


export default defineConfig({
  test: {
    environment: 'jsdom',
  },
  plugins: [react({
    dependencies: [
      reactRouterDom,
    ],
  })],
  server: {
    host: true, // needed for the Docker Container port
    strictPort: true,
    port: 3000
  }
})

