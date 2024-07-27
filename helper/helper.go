package helper

var version = "1.0.0" // Tidak bisa diakses dari luar package
var Application = "GoLang"

func sayGoodBye(name string) string { // Tidak bisa diakses dari luar package
	return "Goodbye " + name
}

func SayHello(name string) string {
	return "Hello " + name
}
