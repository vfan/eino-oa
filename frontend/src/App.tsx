import { BrowserRouter, Routes, Route, Navigate } from 'react-router-dom';
import { ConfigProvider } from 'antd';
import zhCN from 'antd/locale/zh_CN';
import AppLayout from './components/Layout';
import StudentsPage from './pages/students/StudentsPage';
import CoursesPage from './pages/courses/CoursesPage';
import GradesPage from './pages/grades/GradesPage';
import ChatPage from './pages/chat/ChatPage';

function App() {
  return (
    <ConfigProvider locale={zhCN}>
      <BrowserRouter>
        <Routes>
          <Route path="/" element={<AppLayout />}>
            <Route index element={<Navigate to="/students" replace />} />
            <Route path="students" element={<StudentsPage />} />
            <Route path="courses" element={<CoursesPage />} />
            <Route path="grades" element={<GradesPage />} />
            <Route path="chat" element={<ChatPage />} />
          </Route>
        </Routes>
      </BrowserRouter>
    </ConfigProvider>
  );
}

export default App;
