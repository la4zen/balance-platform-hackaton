package store

import "github.com/la4zen/balance-platform-hackaton/pkg/models"

func (s *Store) AddTask(task *models.Task) error {
	r := s.DB.Create(task)
	return r.Error
}

func (s *Store) DeleteTask(task *models.Task) error {
	r := s.DB.Delete(task, "id=?", task.ID)
	return r.Error
}

func (s *Store) GetTasks() ([]models.Task, error) {
	tasks := []models.Task{}
	r := s.DB.Find(&tasks)
	return tasks, r.Error
}
