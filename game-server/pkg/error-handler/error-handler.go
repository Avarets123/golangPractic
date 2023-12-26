package error_handler

func HandleErrorByFn(errHandleFn func(error)) func(error) {
	return func(e error) {
		errHandleFn(e)
	}
}
