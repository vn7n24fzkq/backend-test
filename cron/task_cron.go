package cron

import (
	"fmt"
	"time"
	"vn7n24fzkq/backend-test/common"
	"vn7n24fzkq/backend-test/service"

	"github.com/go-co-op/gocron"
)

func NewTaskCorn(userService *service.UserService, taskService *service.TaskService) *TaskCorn {
	return &TaskCorn{userService, taskService}
}

type TaskCorn struct {
	userService *service.UserService
	taskService *service.TaskService
}

func (p *TaskCorn) InitTaskExpirationNotify() {
	s := gocron.NewScheduler(time.UTC)
	s.Every(10).Seconds().Do(func() {
		tasks, err := p.taskService.GetAllNeedNotifyTask(time.Now().Unix())
		if err != nil {
			fmt.Println(err)
			return
		}
		for _, task := range tasks {
			author, err := p.userService.GetUserByID(task.UserID)
			if err != nil {
				fmt.Println(err)
				continue
			}
			sendMessage := common.DiscordNotificationObject{
				Embeds: []common.DiscordNotificationEmbed{
					{
						Title:       "Task expired!",
						Description: "Expired at ",
						Color:       14177041,
						Fields: []common.DiscordNotificationField{
							{
								Name:   "Title",
								Value:  task.Title,
								Inline: true,
							},
							{
								Name:   "Content",
								Value:  task.Content,
								Inline: true,
							},
							{
								Name:   "Author",
								Value:  author.Username,
								Inline: true,
							},
						},
					},
				},
			}
			// if send notification without any error
			if err := common.SendDiscordWebHookMsg(sendMessage); err == nil {
				task.SendNotify = true
				err := task.Update(p.taskService.TaskDAO, task)
				if err != nil {
					fmt.Println(err)
				}
			} else {
				fmt.Println(err)
			}
		}
	})
	s.StartAsync()
}
