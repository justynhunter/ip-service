package main

type ResponseStatus string

const (
	Ok ResponseStatus = "OK"
	Error ResponseStatus = "ERROR"
)