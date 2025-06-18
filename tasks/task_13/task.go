package task_13

type vehicle struct {
	Model string
	Color string
	Speed int
}

type crossover interface {
	DriveSlow()
	BigTale()
}
