package service

import (
	"testing"
	"time"
	"vn7n24fzkq/backend-test/dao"
)

func TestCreateTask(t *testing.T) {
	taskService, userID := getTaskService(t)
	var testTask = dao.Task{
		Title:   "test",
		Content: "content",
		UserID:  userID,
	}
	_, err := taskService.TaskDAO.CreateTask(testTask)
	if err != nil {
		t.Fatalf("Should not get any error when creating task. %s", err)
	}

	testTask.UserID = 0
	_, err2 := taskService.TaskDAO.CreateTask(testTask)
	if err2 == nil {
		t.Fatalf("Should get any error when creating task belong a not exists task. %s", err)
	}

}

func TestGetTaskByID(t *testing.T) {
	taskService, userID := getTaskService(t)
	var testTask = dao.Task{
		Title:   "test",
		Content: "content",
		UserID:  userID,
	}
	var task, err = taskService.GetTaskByID(1)
	if err == nil {
		t.Fatal("Should get an error when try to get a task which is not exists.")
	}

	task, err = taskService.CreateTask(testTask)
	var createdTaskID = task.ID
	task, err = taskService.GetTaskByID(createdTaskID)

	if err != nil {
		t.Fatalf("Should not get any error when getting an exists task. %s", err)
	} else if createdTaskID != task.ID {
		t.Fatalf("The task ID is %d, and it should be %d", task.ID, createdTaskID)
	}
}

func TestUpdateTaskByID(t *testing.T) {
	taskService, userID := getTaskService(t)
	var testTask = dao.Task{
		Title:   "test",
		Content: "content",
		UserID:  userID,
	}
	er := taskService.UpdateTaskByID(2, testTask)
	if er == nil {
		t.Fatalf("Should get an error when updating a not exists task")
	}
	var task, err = taskService.CreateTask(testTask)
	testTask.Title = "updateTitle"
	testTask.Content = "updatedContent"
	testTask.ExpiredAt = time.Now()
	err = taskService.UpdateTaskByID(testTask.ID, testTask)
	if err != nil {
		t.Fatalf("Should not get any error when updating task. %s", err)
	}

	updatedTask, _ := taskService.GetTaskByID(task.ID)
	if updatedTask.Title != testTask.Title || updatedTask.Content != testTask.Content || updatedTask.ExpiredAt != testTask.ExpiredAt {
		t.Fatalf("Except the ID, task should be same with \n%+v\n,but it is \n%+v", testTask, updatedTask)
	}
}

func TestDeleteTaskByID(t *testing.T) {
	taskService, userID := getTaskService(t)
	var testTask = dao.Task{
		Title:   "test",
		Content: "content",
		UserID:  userID,
	}
	var task, _ = taskService.CreateTask(testTask)
	var err = taskService.DeleteTaskByID(task.ID)
	if err != nil {
		t.Fatalf("Should not get any error when deleting a exists task. %s", err)
	}
	task, err = taskService.GetTaskByID(task.ID)
	if err == nil {
		t.Fatalf("Should any error when getting a deleted task. %s", err)
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
		Salt:           "salt",
	}

	user, err := userDAO.CreateUser(user)
	if err != nil {
		t.Fatalf("Get and error when pre-creating user")
	}

	return NewTaskService(taskDAO), user.ID
}
