package repositories

type TaskRepository interface {
	Add()
	Update()
	Delete()
	List()
}
