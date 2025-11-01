package main

import "BlogSystem/cmd"

func main() {

	defer cmd.Clean()

	cmd.Start()
}
