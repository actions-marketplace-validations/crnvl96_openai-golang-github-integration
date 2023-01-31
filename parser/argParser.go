package parser

func ArgParser(args []string) (result []string) {
	argsSlice := args[1:]
	result = append(result, argsSlice...)
	return
}