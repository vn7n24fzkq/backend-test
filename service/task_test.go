package service

import (
	"testing"
	"time"
	"vn7n24fzkq/backend-test/dao"

	"github.com/google/go-cmp/cmp"
)

func TestCreateTask(t *testing.T) {
	taskService, preCreatedUserID := getTaskService(t)
	var testTask = dao.Task{
		Title:   "test",
		Content: "content",
		UserID:  preCreatedUserID,
	}
	_, err := taskService.CreateTask(testTask)
	if err != nil {
		t.Fatalf("Should not get any error when creating task. %s", err)
	}
}

func TestCreateTaskFail(t *testing.T) {
	taskService, preCreatedUserID := getTaskService(t)
	var testTask = dao.Task{
		Title:   "test",
		Content: "content",
		UserID:  preCreatedUserID + 1,
	}
	taskService.CreateTask(testTask)
	_, err := taskService.CreateTask(testTask)
	if err == nil {
		t.Fatalf("Should get an error when creating task belong a not exists user. %s", err)
	}
}

func TestGetTaskByID(t *testing.T) {
	taskService, preCreatedUserID := getTaskService(t)
	var testTask = dao.Task{
		Title:   "test",
		Content: "content",
		UserID:  preCreatedUserID,
	}

	var task, err = taskService.CreateTask(testTask)
	var createdTaskID = task.ID
	task, err = taskService.GetTaskByID(createdTaskID)

	if err != nil {
		t.Fatalf("Should not get any error when getting an exists task. %s", err)
	} else if createdTaskID != task.ID {
		t.Fatalf("The task ID is %d, and it should be %d", task.ID, createdTaskID)
	}
}

func TestGetTaskByIDFail(t *testing.T) {
	taskService, _ := getTaskService(t)
	_, err := taskService.GetTaskByID(1)
	if err == nil {
		t.Fatal("Should get an error when try to get a task which is not exists.")
	}
}

func TestUpdateTaskByID(t *testing.T) {
	taskService, preCreatedUserID := getTaskService(t)
	var testTask = dao.Task{
		Title:   "test",
		Content: "content",
		UserID:  preCreatedUserID,
	}
	var task, err = taskService.CreateTask(testTask)
	testTask.Title = "updateTitle"
	testTask.Content = "updatedContent"
	testTask.ExpiredAt = time.Now().Unix()
	err = taskService.UpdateTaskByID(task.ID, testTask)
	if err != nil {
		t.Fatalf("Should not get any error when updating task. %s", err)
	}

	updatedTask, _ := taskService.GetTaskByID(task.ID)
	if updatedTask.Title != testTask.Title || updatedTask.Content != testTask.Content || updatedTask.ExpiredAt != testTask.ExpiredAt {
		t.Fatalf("Except the ID, task should be same with \n%+v\n,but it is \n%+v", testTask, updatedTask)
	}
}

func TestUpdateTaskByIDFail(t *testing.T) {
	taskService, preCreatedUserID := getTaskService(t)
	var testTask = dao.Task{
		Title:   "test",
		Content: "content",
		UserID:  preCreatedUserID,
	}
	err := taskService.UpdateTaskByID(2, testTask)
	if err == nil {
		t.Fatalf("Should get an error when updating a not exists task")
	}
}

func TestDeleteTaskByID(t *testing.T) {
	taskService, preCreatedUserID := getTaskService(t)
	var testTask = dao.Task{
		Title:   "test",
		Content: "content",
		UserID:  preCreatedUserID,
	}
	var task, err = taskService.CreateTask(testTask)
	err = taskService.DeleteTaskByID(task.ID)
	if err != nil {
		t.Fatalf("Should not get any error when deleting task. %s", err)
	}
}

func TestDeleteTaskByIDFail(t *testing.T) {
	taskService, _ := getTaskService(t)
	err := taskService.DeleteTaskByID(1)
	if err == nil {
		t.Fatalf("Should get an error when deleting a not exists task. %s", err)
	}
}

func TestGetAllTaskByUserID(t *testing.T) {
	taskService, preCreatedUserID := getTaskService(t)
	var exceptTaskArray = []dao.Task{}
	var testTask = dao.Task{
		Title:     "test",
		Content:   "content",
		UserID:    preCreatedUserID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	var tasks, err = taskService.GetAllTaskByUserID(preCreatedUserID)
	if !cmp.Equal(exceptTaskArray, tasks) {
		t.Fatalf("Should equal with an empty task array. %s", err)
	}

	var task1, _ = taskService.CreateTask(testTask)
	exceptTaskArray = append(exceptTaskArray, task1)
	var task2, _ = taskService.CreateTask(testTask)
	exceptTaskArray = append(exceptTaskArray, task2)
	var task3, _ = taskService.CreateTask(testTask)
	exceptTaskArray = append(exceptTaskArray, task3)

	tasks, err = taskService.GetAllTaskByUserID(preCreatedUserID)

	for i := 0; i < 3; i++ {
		if !cmp.Equal(exceptTaskArray[i], tasks[i]) {
			t.Fatalf("Should equal with task array. %s", err)
		}
	}
	taskService.DeleteTaskByID(task1.ID)
	tasks, err = taskService.GetAllTaskByUserID(preCreatedUserID)
	if len(tasks) == len(exceptTaskArray) {
		t.Fatalf("Should not has same length. %s", err)
	}
}

// Return TaskService and pre-created user ID
func getTaskService(t *testing.T) (*TaskService, int) {
	db := GetTestDB(t)
	// Initialize DAO
	taskDAO := dao.NewTaskDAO(db)
	userDAO := dao.NewUserDAO(db)
	user := dao.User{
		Username:       "testUser",
		PasswordDigest: "password",
	}

	user, err := userDAO.CreateUser(user)
	if err != nil {
		t.Fatalf("Get and error when pre-creating user")
	}

	return NewTaskService(taskDAO, NewUserService(userDAO)), user.ID
}
