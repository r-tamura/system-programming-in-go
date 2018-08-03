package main

func getArgs(args []string) []string {

	if len(args) == 1 {
		return args
	}

	if args[1] == "--" {
		return append(args[:1], args[2:]...)
	}

	return args
}
