package agent

type Agent interface {
	DoJob() error
	Stop() error
}
