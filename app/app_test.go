package app

import (
	"errors"
	"task-tracker/mocks"
	"task-tracker/models"
	"testing"
)

func compareIgnoringTime(t1, t2 models.Task) bool {
	return t1.Id == t2.Id &&
		t1.Description == t2.Description &&
		t1.Status == t2.Status
}

func compareTaskSlices(expected, actual []models.Task) bool {
	if len(expected) != len(actual) {
		return false
	}

	for i := range expected {
		if !compareIgnoringTime(expected[i], actual[i]) {
			return false
		}
	}

	return true
}

func TestAddTask(t *testing.T) {
	tests := []struct {
		name          string
		input         string
		mockError     error
		expectError   bool
		expectedId    int
		expectCalled  bool
		expectedParam models.Task
	}{
		{name: "success case",
			input:         "Buy groceries",
			mockError:     nil,
			expectError:   false,
			expectCalled:  true,
			expectedId:    1,
			expectedParam: models.Task{Id: 1, Description: "Buy groceries", Status: "todo"},
		},
		{name: "error case",
			input:         "Buy groceries",
			mockError:     errors.New("mock error"),
			expectError:   true,
			expectCalled:  true,
			expectedId:    -1,
			expectedParam: models.Task{Id: 1, Description: "Buy groceries", Status: "todo"},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mockStorage := &mocks.MockStorage{SaveInfoErr: test.mockError}
			testApp := NewApp(mockStorage)
			id, err := testApp.AddTask(test.input)
			if test.expectError && err == nil {
				t.Errorf("Wanted error but got nil")
			}
			if !test.expectError && err != nil {
				t.Errorf("Wanted no error but got %s", err)
			}
			if id != test.expectedId && !test.expectError {
				t.Errorf("Wanted id %d but got %d", test.expectedId, id)
			}
			if mockStorage.SaveInfoCalled != test.expectCalled {
				t.Errorf("expected SaveInfoCalled to be %v but got %v", test.expectCalled, mockStorage.SaveInfoCalled)
			}
			if test.expectCalled && !compareIgnoringTime(test.expectedParam, mockStorage.SaveInfoParams) {
				t.Errorf("expected %v but got %v", test.expectedParam, mockStorage.SaveInfoParams)
			}
		})
	}
}

func TestUpdateTask(t *testing.T) {
	tests := []struct {
		name            string
		inputId         int
		inputString     string
		mockError       error
		mockLoadError   error
		mockLoadResult  []models.Task
		mockUpdateError error
		expectError     bool
		expectCalled    bool
		expectedParam   []models.Task
	}{
		{name: "success case",
			inputId:         1,
			inputString:     "Updated",
			mockError:       nil,
			mockLoadResult:  []models.Task{{Id: 1, Description: "Old task", Status: "todo"}},
			mockLoadError:   nil,
			mockUpdateError: nil,
			expectError:     false,
			expectCalled:    true,
			expectedParam:   []models.Task{{Id: 1, Description: "Updated", Status: "todo"}},
		},
		{name: "Error loading info",
			inputId:         1,
			inputString:     "Updated task",
			mockLoadResult:  nil,
			mockLoadError:   errors.New("load error"),
			mockUpdateError: nil,
			expectError:     true,
			expectCalled:    false,
			expectedParam:   nil,
		},
		{name: "Error updating info",
			inputId:         1,
			inputString:     "Updated task",
			mockLoadResult:  nil,
			mockLoadError:   nil,
			mockUpdateError: errors.New("update error"),
			expectError:     true,
			expectCalled:    true,
			expectedParam:   nil,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mockStorage := &mocks.MockStorage{LoadInfoResult: test.mockLoadResult,
				LoadInfoErr:   test.mockLoadError,
				UpdateInfoErr: test.mockUpdateError}
			testApp := NewApp(mockStorage)
			err := testApp.UpdateTask(test.inputId, test.inputString)
			if test.expectError && err == nil {
				t.Errorf("Wanted error but got nil")
			}
			if !test.expectError && err != nil {
				t.Errorf("Wanted no error but got %s", err)
			}
			if mockStorage.UpdateInfoCalled != test.expectCalled {
				t.Errorf("expected called %v got %v", mockStorage.UpdateInfoCalled, test.expectCalled)
			}
			if !errors.Is(test.mockLoadError, test.mockLoadError) {
				t.Errorf("Wanted error %v got %v", test.mockLoadError, test.mockLoadError)
			}
			if !errors.Is(test.mockUpdateError, test.mockUpdateError) {
				t.Errorf("Wanted error %v got %v", test.mockUpdateError, test.mockUpdateError)
			}
			if test.expectCalled && !compareTaskSlices(test.expectedParam, mockStorage.UpdateInfoParams) {
				t.Errorf("expected %v but got %v", test.expectedParam, mockStorage.UpdateInfoParams)
			}
		})
	}
}

func TestDeleteTask(t *testing.T) {
	tests := []struct {
		name            string
		inputId         int
		mockError       error
		mockLoadError   error
		mockUpdateError error
		mockLoadResult  []models.Task
		expectError     bool
		expectCalled    bool
		expectedParam   []models.Task
	}{
		{name: "success case",
			inputId:   1,
			mockError: nil,
			mockLoadResult: []models.Task{
				{Id: 1,
					Description: "Old task",
					Status:      "todo",
				},
				{Id: 2,
					Description: "Delete",
					Status:      "todo",
				}},
			mockLoadError:   nil,
			mockUpdateError: nil,
			expectError:     false,
			expectCalled:    true,
			expectedParam: []models.Task{
				{Id: 2,
					Description: "Delete",
					Status:      "todo",
				}},
		},
		{name: "Error loading info",
			inputId:         2,
			mockLoadResult:  nil,
			mockLoadError:   errors.New("load error"),
			mockUpdateError: nil,
			expectError:     true,
			expectCalled:    false,
			expectedParam:   nil,
		},
		{name: "Error updating info",
			inputId: 2,
			mockLoadResult: []models.Task{
				{Id: 1,
					Description: "Old task",
					Status:      "todo",
				},
				{Id: 2,
					Description: "Delete",
					Status:      "todo",
				}},
			mockLoadError:   nil,
			mockUpdateError: errors.New("update error"),
			expectError:     true,
			expectCalled:    false,
			expectedParam:   nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mockStorage := &mocks.MockStorage{
				LoadInfoResult: test.mockLoadResult,
				LoadInfoErr:    test.mockLoadError,
				UpdateInfoErr:  test.mockUpdateError}
			testApp := NewApp(mockStorage)
			err := testApp.DeleteTask(test.inputId)
			if test.expectError && err == nil {
				t.Errorf("Wanted error but got nil")
			}
			if !test.expectError && err != nil {
				t.Errorf("Wanted no error but got %s", err)
			}
			if mockStorage.UpdateInfoCalled != test.expectCalled {
				t.Errorf("Expected UpdateInfoCalled to be %v but got %v", test.expectCalled, mockStorage.UpdateInfoCalled)
			}
			if test.expectCalled && !compareTaskSlices(test.expectedParam, mockStorage.UpdateInfoParams) {
				t.Errorf("Expected %v but got %v", test.expectedParam, mockStorage.UpdateInfoParams)
			}

		})
	}
}

//type TaskService interface {
//	AddTask(task string) (int, error)
//	UpdateTask(id int, task string) error
//	DeleteTask(id int) error
//	ListAllTasks() ([]models.Task, error)
//	ListDoneTasks() ([]models.Task, error)
//	ListProgressTasks() ([]models.Task, error)
//	ListToDoTasks() ([]models.Task, error)
//	MarkInProgress(id int) error
//	MarkDone(id int) error
//	MarkToDo(id int) error
//}
