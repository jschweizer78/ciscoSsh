package main

import "os"

func main() {
	perm := os.FileMode(0777)
	os.Chmod("../test", perm)

}
