# AWS Lambda Strategy Pattern with Go

## Overview
This project demonstrates the implementation of the Strategy design pattern using AWS Lambda in Go. The project structure incorporates GORM and other best practices to achieve modular, maintainable, and scalable code.

### Key Features
- **AWS Lambda Integration**: A serverless architecture leveraging AWS Lambda to handle tasks dynamically based on actions.
- **Strategy Pattern**: Implementation of the Strategy design pattern to separate and encapsulate algorithms, allowing flexible task execution.
- **Event-driven Processing**: Event-specific task handling using JSON payloads.
- **SAM CLI Integration**: Simplified build and local testing using AWS Serverless Application Model (SAM).

## Project Structure
```plaintext
.
├── app
│   ├── bootstrap
│   ├── controller
│   │   ├── route
│   │   │   └── route.go
│   │   └── task
│   │       ├── model
│   │       │   ├── task_request.go
│   │       │   └── task_response.go
│   │       ├── task_controller.go
│   │       └── task_interface.go
│   ├── domain
│   │   ├── task_domain.go
│   │   └── task_domain_interface.go
│   ├── go.mod
│   ├── go.sum
│   ├── main.go
│   ├── usecase
│   │   ├── run_task_service.go
│   │   ├── task
│   │   │   ├── delete_logs.go
│   │   │   ├── task_strategy.go
│   │   │   └── update_database.go
│   │   └── task_usecase.go
│   └── view
│       └── convert_domain_to_response.go
├── events
│   ├── delete_logs.json
│   └── update_database.json
├── README.md
├── run-local.sh
└── template.yaml
```

## Implementation Details

### Task Execution with Strategy Pattern
The **Strategy Pattern** is implemented using a registry of task strategies, allowing dynamic registration and execution of tasks based on the provided action.

#### `run_task_service.go`
```go
package usecase

import (
    "bootstrap/domain"
    "bootstrap/usecase/task"
    "log"
)

const DELETE string = "delete_logs"
const UPDATE string = "update_database"

func (t taskUseCase) Task(taskInterface domain.TaskDomainInterface) string {

    taskStrategy := task.NewStrategy()
    taskStrategy.Register(DELETE, task.DeleteLogs{})
    taskStrategy.Register(UPDATE, task.UpdateDataBase{})

    run, err := taskStrategy.Get(taskInterface.GetAction())
    if err != nil {
        log.Panic("Strategy not found")
    }
    return run.Run(taskInterface.GetParameters())
}
```

#### `task_strategy.go`
```go
package task

import "fmt"

type TaskStrategy interface {
    Run(data string) string
}

type StrategyRegistry struct {
    strategies map[string]TaskStrategy
}

func NewStrategy() *StrategyRegistry {
    return &StrategyRegistry{strategies: make(map[string]TaskStrategy)}
}

func (s *StrategyRegistry) Register(action string, strategy TaskStrategy) {
    s.strategies[action] = strategy
}

func (s *StrategyRegistry) Get(action string) (TaskStrategy, error) {
    strategy, exists := s.strategies[action]
    if !exists {
        return nil, fmt.Errorf("Strategy '%s' not found", action)
    }
    return strategy, nil
}
```

### Local Development Script
The `run-local.sh` script simplifies local development by automating the build process, setting up environment variables, and invoking Lambda functions with sample events.

#### `run-local.sh`
```bash
#!/bin/bash

rm -rf .aws-sam/

export AWS_ACCESS_KEY_ID="your_access_key_id"
export AWS_SECRET_ACCESS_KEY="your_secret_access_key"
export AWS_REGION="your_region"

export GOOS=linux
export GOARCH=amd64
export CGO_ENABLED=0

# Compile Go binary
cd app
go mod tidy
go build
if [[ $? -ne 0 ]]; then
    echo "Error compiling the binary. Exiting."
    exit 1
fi
cd ..

# Build SAM
sam build --template-file template.yaml --build-dir .aws-sam/build
if [[ $? -ne 0 ]]; then
    echo "Error executing SAM build. Exiting."
    exit 1
fi

# Select event
EVENT_DIR="./events"

echo "Select an event to run:"
EVENT_FILES=()
i=1
for file in "$EVENT_DIR"/*; do
    if [[ -f "$file" ]]; then
        EVENT_FILES+=("$file")
        echo "[$i] $(basename "$file")"
        ((i++))
    fi
done

read -p "Enter the event number: " EVENT_CHOICE

if [[ "$EVENT_CHOICE" -ge 1 && "$EVENT_CHOICE" -le "${#EVENT_FILES[@]}" ]]; then
    SELECTED_EVENT=${EVENT_FILES[$EVENT_CHOICE-1]}
    echo "Running with event: $(basename "$SELECTED_EVENT")"

    sam local invoke -e "$SELECTED_EVENT" StrategyFunction
else
    echo "Invalid option. Exiting."
    exit 1
fi
```

### How to Run
1. **Build the Project:**
   ```bash
   ./run-local.sh
   ```
2. **Deploy with SAM:**
   ```bash
   sam build
   sam deploy --guided
   ```

## Supported Actions
- **delete_logs**: Deletes logs based on the provided parameters.
- **update_database**: Updates the database using specified inputs.

## Sample Events
- `delete_logs.json`
- `update_database.json`

These sample JSON files simulate input events for the Lambda function.

## License
This project is licensed under the Apache 2.0 License.