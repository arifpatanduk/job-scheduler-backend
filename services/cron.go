// services/cron.go
package services

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"job-scheduler-backend/models"
	"log"
	"net/http"

	"github.com/robfig/cron/v3"
	"gorm.io/gorm"
)

var db *gorm.DB
var broadcast = make(chan Message)

type Message struct {
    JobID    uint   `json:"job_id"`
    Status   string `json:"status"`
    Response string `json:"response"`
}

func StartCronScheduler(database *gorm.DB) {
    db = database
    c := cron.New()

    schedulers, err := loadActiveSchedulers()
    if err != nil {
        log.Fatal(err)
    }

    for _, scheduler := range schedulers {
        scheduleScheduler(c, scheduler)
    }

    c.Start()
}

func loadActiveSchedulers() ([]models.Scheduler, error) {
    var schedulers []models.Scheduler
    result := db.Preload("Jobs").Where("is_active = ?", true).Find(&schedulers)
    return schedulers, result.Error
}

func scheduleScheduler(c *cron.Cron, scheduler models.Scheduler) {
    c.AddFunc(scheduler.Interval, func() {
        go executeScheduler(scheduler)
    })
}

func executeScheduler(scheduler models.Scheduler) {
    for _, job := range scheduler.Jobs {
        executeJob(job)
    }
}

func executeJob(job models.Jobs) {
    var response string
    switch job.Activity.Name {
    case "hit_api":
        response = sendTaskToAPIService(job)
    default:
        log.Printf("Unknown activity type: %s\n", job.Activity.Name)
        response = "Unknown activity type"
    }
    broadcast <- Message{JobID: job.ID, Status: "completed", Response: response}
    logExecution(job.ID, "Job executed successfully")
}

func sendTaskToAPIService(job models.Jobs) string {
    payload := map[string]interface{}{
        "url":    job.Url,
        // "method": job.Method,
        // "body":   job.Body,
    }
    return sendTask("http://api-service:8081/execute", payload)
}

func sendTask(url string, payload map[string]interface{}) string {
    jsonPayload, _ := json.Marshal(payload)
    resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonPayload))
    if err != nil {
        log.Printf("Failed to send task: %v\n", err)
        return "Failed to send task"
    }
    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Printf("Failed to read response: %v\n", err)
        return "Failed to read response"
    }

    if resp.StatusCode != http.StatusOK {
        log.Printf("Failed to execute task, received status code: %d\n", resp.StatusCode)
        return "Failed to execute task"
    }

    return string(body)
}

func logExecution(jobID uint, message string) {
    db.Create(&models.Logs{JobID: jobID, Message: message})
}
