package write


type sender interface {
	Send()
}

type senderStatusSetter interface {
	Send()
	Status(statusCode int) sender
}