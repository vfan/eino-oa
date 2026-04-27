import axios from 'axios';

const api = axios.create({
  baseURL: '/api',
  timeout: 10000,
});

export interface PaginatedResult<T> {
  items: T[];
  total: number;
  page: number;
  page_size: number;
}

export interface Student {
  id?: number;
  student_no: string;
  name: string;
  gender: string;
  age: number;
  phone: string;
  email: string;
  department: string;
  major: string;
  class_name: string;
  enrollment_year: number;
  status: string;
  created_at?: string;
  updated_at?: string;
}

export interface Course {
  id?: number;
  course_code: string;
  name: string;
  credits: number;
  department: string;
  teacher: string;
  semester: string;
  description: string;
  created_at?: string;
  updated_at?: string;
}

export interface Grade {
  id?: number;
  student_id: number;
  course_id: number;
  score: number;
  grade_point?: number;
  semester: string;
  status: string;
  student?: Student;
  course?: Course;
  created_at?: string;
  updated_at?: string;
}

export interface GradeStats {
  course_id: number;
  course_name: string;
  avg_score: number;
  max_score: number;
  min_score: number;
  pass_rate: number;
  count: number;
}

export interface StudentGPA {
  student_id: number;
  student_no: string;
  student_name: string;
  gpa: number;
  total_courses: number;
}

export const studentApi = {
  list: (params?: Record<string, unknown>) =>
    api.get<PaginatedResult<Student>>('/students', { params }),
  get: (id: number) => api.get<Student>(`/students/${id}`),
  create: (data: Student) => api.post<Student>('/students', data),
  update: (id: number, data: Student) => api.put<Student>(`/students/${id}`, data),
  delete: (id: number) => api.delete(`/students/${id}`),
};

export const courseApi = {
  list: (params?: Record<string, unknown>) =>
    api.get<PaginatedResult<Course>>('/courses', { params }),
  get: (id: number) => api.get<Course>(`/courses/${id}`),
  create: (data: Course) => api.post<Course>('/courses', data),
  update: (id: number, data: Course) => api.put<Course>(`/courses/${id}`, data),
  delete: (id: number) => api.delete(`/courses/${id}`),
};

export const gradeApi = {
  list: (params?: Record<string, unknown>) =>
    api.get<PaginatedResult<Grade>>('/grades', { params }),
  get: (id: number) => api.get<Grade>(`/grades/${id}`),
  create: (data: Grade) => api.post<Grade>('/grades', data),
  update: (id: number, data: Grade) => api.put<Grade>(`/grades/${id}`, data),
  delete: (id: number) => api.delete(`/grades/${id}`),
  statsByCourse: (courseId: number) => api.get<GradeStats>(`/grades/stats/course/${courseId}`),
  calculateGPA: (studentId: number) => api.get<StudentGPA>(`/grades/gpa/student/${studentId}`),
  academicWarning: (threshold?: number) =>
    api.get<{ threshold: number; students: StudentGPA[] }>('/grades/warning', {
      params: { threshold },
    }),
};

export interface ChatSession {
  id: string;
  title: string;
  created_at: string;
  updated_at: string;
}

export interface ChatMessageItem {
  id: number;
  session_id: string;
  role: 'user' | 'assistant';
  content: string;
  created_at: string;
}

export interface ChatSessionDetail extends ChatSession {
  messages: ChatMessageItem[];
}

export const sessionApi = {
  list: () => api.get<ChatSession[]>('/sessions'),
  get: (id: string) => api.get<ChatSessionDetail>(`/sessions/${id}`),
  create: (title?: string) => api.post<ChatSession>('/sessions', { title }),
  delete: (id: string) => api.delete(`/sessions/${id}`),
  chat: (id: string, message: string) =>
    api.post<{ reply: string }>(`/sessions/${id}/chat`, { message }, { timeout: 120000 }),
};

export default api;
