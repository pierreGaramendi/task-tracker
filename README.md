# Task Tracker

Solution for the [Task Tracker](https://roadmap.sh/projects/task-tracker) backend challenge from [roadmap.sh](https://roadmap.sh/).



### INSTALLATION
    git clone https://github.com/pierreGaramendi/task-tracker.git
    cd task-tracker
    go build -o task-tracker  

### COMMANDS AND OPTIONS

The task-tracker tool allows you to manage tasks via different commands. Below are the available commands and their options.

    ./task-tracker [COMMAND] [OPTIONS] <arguments>

#### List Usage:

    ./task-tracker list                         List all tasks

Lists all tasks. You can filter tasks by status:

    ./task-tracker list -status todo            Show all tasks that are not done
    ./task-tracker list -status in-progress     Show all tasks that are in progress
    ./task-tracker list -status done            Show all tasks that are done

#### Add Task Usage:

    ./task-tracker add -task "Task description"     Adds a new task with the given description.

#### Update Task Usage:

    ./task-tracker update -id <taskId> -desc "New task description"     Edit task description
    ./task-tracker mark-in-progress -id 1     Marking a task as in progress
    ./task-tracker mark-done -id 1     Marking a task as done

#### Delete Task Usage:

    ./task-tracker delete -id 1    Deleting tasks