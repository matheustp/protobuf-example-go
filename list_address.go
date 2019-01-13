package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"

	"github.com/golang/protobuf/proto"

	pb "github.com/matheustp/protobuf-example-go/src/tutorial"
)

func writePerson(w io.Writer, p *pb.Person) {
	fmt.Fprintln(w, "Person ID:", p.Id)
	fmt.Fprintln(w, "\tName:", p.Name)
	if p.Email != "" {
		fmt.Fprintln(w, "\tE-mail address:", p.Email)
	}

	for _, pn := range p.Phones {
		switch pn.Type {
		case pb.Person_MOBILE:
			fmt.Fprint(w, "\tMobile phone #: ")
		case pb.Person_HOME:
			fmt.Fprint(w, "\tHome phone #: ")
		case pb.Person_WORK:
			fmt.Fprint(w, "\tWork phone #: ")
		}
		fmt.Fprintln(w, pn.Number)
	}
}

func listPeople(w io.Writer, ab *pb.AddressBook) {
	for _, p := range ab.People {
		writePerson(w, p)
	}
}

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("Usage: %s ADDRESS_BOOK_FILE\n", os.Args[0])
	}
	fname := os.Args[1]
	in, err := ioutil.ReadFile(fname)
	if err != nil {
		log.Fatalf("Error reading the file ", err)
	}
	ab := &pb.AddressBook{}
	if err := proto.Unmarshal(in, ab); err != nil {
		log.Fatalf("Error parsing the file ", err)
	}

	listPeople(os.Stdout, ab)
}
