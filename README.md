# Task Tracker

Solution for the [Task Tracker](https://roadmap.sh/projects/task-tracker) backend challenge from [roadmap.sh](https://roadmap.sh/).



### INSTALLATION
    git clone https://github.com/pierreGaramendi/task-tracker.git
    cd task-tracker
    go build -o task-tracker  

### USAGE AND OPTIONS

    ./task-tracker [OPTION] [--][FLAG] <argument>

#### List Usage:

    ./task-tracker list                         List all tasks
    ./task-tracker list -status todo            List all tasks that are not done
    ./task-tracker list -status in-progress     List all tasks that are in progress
    ./task-tracker list -status done            List all tasks that are done

#### Add Task Usage:

    ./task-tracker add -task "Task description"     Add new task to json file