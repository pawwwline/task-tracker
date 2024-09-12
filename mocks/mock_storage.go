package mocks

import "task-tracker/models"

type MockStorage struct {
	SaveInfoCalled bool
	SaveInfoErr    error
	SaveInfoParams models.Task

	LoadInfoCalled bool
	LoadInfoErr    error
	LoadInfoResult []models.Task

	UpdateInfoCalled bool
	UpdateInfoParams []models.Task
	UpdateInfoErr    error
}

func (m *MockStorage) SaveInfo(task models.Task) error {
	m.SaveInfoCalled = true
	m.SaveInfoParams = task
	if m.SaveInfoErr != nil {
		return m.SaveInfoErr
	}
	return nil
}

func (m *MockStorage) LoadInfo() ([]models.Task, error) {
	m.LoadInfoCalled = true
	if m.LoadInfoErr != nil {
		return nil, m.LoadInfoErr
	}
	return m.LoadInfoResult, nil
}

func (m *MockStorage) UpdateInfo(tasks []models.Task) error {
	m.UpdateInfoCalled = true
	m.UpdateInfoParams = tasks
	if m.UpdateInfoErr != nil {
		return m.UpdateInfoErr
	}
	return nil
}

//type Storage interface {
//	SaveInfo(task models.Task)
//	LoadInfo() []models.Task
//	UpdateInfo([]models.Task)
//}
