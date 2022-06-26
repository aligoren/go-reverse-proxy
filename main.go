package main

import "reverse_proxy/cmd"

func main() {
	cmd.ServeApp(":80")
}
