package service

import (
	"errors"

	"github.com/wuhaitao/eino-oa/backend/internal/model"
	"github.com/wuhaitao/eino-oa/backend/internal/repository"
)

type StudentService struct {
	repo *repository.StudentRepo
}

func NewStudentService(repo *repository.StudentRepo) *StudentService {
	return &StudentService{repo: repo}
}

func (s *StudentService) List(q repository.StudentQuery) (*repository.PaginatedResult[model.Student], error) {
	return s.repo.List(q)
}

func (s *StudentService) GetByID(id uint) (*model.Student, error) {
	return s.repo.GetByID(id)
}

func (s *StudentService) GetByStudentNo(no string) (*model.Student, error) {
	return s.repo.GetByStudentNo(no)
}

func (s *StudentService) Create(student *model.Student) error {
	if student.StudentNo == "" || student.Name == "" {
		return errors.New("学号和姓名不能为空")
	}
	if student.Status == "" {
		student.Status = model.StudentStatusActive
	}
	return s.repo.Create(student)
}

func (s *StudentService) Update(student *model.Student) error {
	if student.ID == 0 {
		return errors.New("学生ID不能为空")
	}
	return s.repo.Update(student)
}

func (s *StudentService) Delete(id uint) error {
	return s.repo.Delete(id)
}
