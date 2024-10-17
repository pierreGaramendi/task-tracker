package tasks

// Definir el tipo TaskStatus
type TaskStatus string

// Constantes que representan los diferentes estados
const (
	Todo       TaskStatus = "todo"
	InProgress TaskStatus = "in-progress"
	Done       TaskStatus = "done"
)

func (ts TaskStatus) IsValid() bool {
	switch ts {
	case Todo, InProgress, Done:
		return true
	}
	return false
}
