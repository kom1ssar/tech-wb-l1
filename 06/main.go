package main

func main() {

	// завершаемся по каналу
	exitGoroutineByChExit()

	// завершаемся по статусу
	exitGoroutineByChStatus()

	// завершаемся с помощью контекста cancel()
	exitGoroutineByCtx()

	// завершаемся с помощью таймайута котекста
	exitGoroutineByCtxTimer()

	// завершаемся с помощью Goexit
	exitGoroutineByGoExit()
}
