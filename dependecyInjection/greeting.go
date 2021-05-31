package dependencyInjection

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

// how to use to write data in other cases
// func main() {
// 	Greet(os.Stdout, "Elodie")
// }

// web
func MyGreeterHandler(w http.ResponseWriter, r *http.Request ) {
	Greet(w, "world")
}

func main() {
	log.Fatal(http.ListenAndServe(":5000", http.HandlerFunc(MyGreeterHandler)))
}

