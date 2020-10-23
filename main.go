package main

func main() {
	octocli()
}

func octocli() []string {
	var arg []string
	var optionalArgs []string

	arg = append(arg, "--project=${INPUT_PROJECT}")
	arg = append(arg, "--server=${INPUT_OCTOPUS_SERVER}")
	arg = append(arg, "--apikey=${API_KEY}")

	if set == true {
		optionalArgs = append(arg, "--space=${INPUT_SPACE}")
		return optionalArgs
	}

	return []string{""}
}
