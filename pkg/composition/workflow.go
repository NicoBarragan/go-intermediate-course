package composition

type WorkFlow interface {
	StartWorking(bool) string
	EndWorking(bool) string
}
