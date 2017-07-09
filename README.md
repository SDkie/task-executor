# Task-executor
Task Executor in Golang

### Installation:
 - Set `$GOPATH`
 - git clone https://github.com/SDkie/task-executor.git
 
### Setup:
 - Install Mysql
 - Create database for this project
 - updated Env variables in run.sh
 
### Build:
- go build
 
### Run:
 - execute run.sh

### Post task:
 - Method : POST
 - URL: localhost:7000/task
 - Sample Body - {
	"url": "https://www.facebook.com", 
	"method": "POST",
	"data": {
	"name":"Kumar"
	},
	"max_retry": 20, 
	"retry_until":"2017-07-09T20:50:36Z"
}

### View Task Status:
 - Goto http://localhost:7000/tasks/status URL
