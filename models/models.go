package models

type Picture struct {
	Image    []byte
	Title    string
	Date     string
	Category []string
	Tags     []string
}
