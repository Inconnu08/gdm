package main

import "gdm/gdm"

func main() {
	defer gdm.Elapsed()()
	gdm.Download("downloads/")
}
