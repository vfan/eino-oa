import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'

// 后端默认端口。若后端改端口，设置 VITE_BACKEND_URL 即可，例如：
// VITE_BACKEND_URL=http://localhost:8081 npm run dev
const backend = process.env.VITE_BACKEND_URL || 'http://localhost:8080'
const toBackend = () => ({ target: backend, changeOrigin: true })

// https://vite.dev/config/
export default defineConfig({
  plugins: [react()],
  server: {
    port: 3000,
    proxy: {
      '/api': toBackend(),
      // OAuth 授权页与 token 端点需要与前端同源，Cookie 才能正常携带
      '/oauth': toBackend(),
      '/.well-known': toBackend(),
      // Agent 通过网页地址下载技能文档
      '/skills': toBackend(),
    },
  },
})
