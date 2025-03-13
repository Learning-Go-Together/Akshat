package main

func main() {

}

// In this interface the method is taking two string input but we cannot be sure what are they
// also it returns int value and we don't know what is this
type Copier interface {
	Copy(string, string) int
}

// So we can use named parameter to provide more detailed view of interfaces
type Copier2 interface {
	Copy(sourceFile string, destinationFile string) (bytesCopied int)
}
