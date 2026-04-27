package service

import (
	"errors"

	"github.com/wuhaitao/eino-oa/backend/internal/model"
	"github.com/wuhaitao/eino-oa/backend/internal/repository"
)

type CourseService struct {
	repo *repository.CourseRepo
}

func NewCourseService(repo *repository.CourseRepo) *CourseService {
	return &CourseService{repo: repo}
}

func (s *CourseService) List(q repository.CourseQuery) (*repository.PaginatedResult[model.Course], error) {
	return s.repo.List(q)
}

func (s *CourseService) GetByID(id uint) (*model.Course, error) {
	return s.repo.GetByID(id)
}

func (s *CourseService) Create(course *model.Course) error {
	if course.CourseCode == "" || course.Name == "" {
		return errors.New("课程编号和课程名不能为空")
	}
	return s.repo.Create(course)
}

func (s *CourseService) Update(course *model.Course) error {
	if course.ID == 0 {
		return errors.New("课程ID不能为空")
	}
	return s.repo.Update(course)
}

func (s *CourseService) Delete(id uint) error {
	return s.repo.Delete(id)
}
