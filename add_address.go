package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/golang/protobuf/proto"

	pb "github.com/matheustp/protobuf-example-go/src/tutorial"
)

func promptForAddress(r io.Reader) (*pb.Person, error) {
	p := &pb.Person{}
	rd := bufio.NewReader(r)
	fmt.Print("Enter person ID: ")
	if _, err := fmt.Fscanf(rd, "%d\n", &p.Id); err != nil {
		return p, err
	}

	fmt.Print("Enter name: ")
	name, err := rd.ReadString('\n')
	if err != nil {
		return p, err
	}
	p.Name = strings.TrimSpace(name)

	fmt.Print("Enter email (blank for none): ")
	email, err := rd.ReadString('\n')
	if err != nil {
		return p, err
	}
	p.Email = strings.TrimSpace(email)

	for {
		fmt.Print("Enter phone number (blank to finish): ")
		phone, err := rd.ReadString('\n')
		if err != nil {
			return p, err
		}
		phone = strings.TrimSpace(phone)
		if phone == "" {
			break
		}

		pn := &pb.Person_PhoneNumber{
			Number: phone,
		}

		fmt.Print("Is this a mobile, home or work phone? ")
		ptype, err := rd.ReadString('\n')
		if err != nil {
			return p, err
		}
		ptype = strings.TrimSpace(ptype)
		switch ptype {
		case "mobile":
			pn.Type = pb.Person_MOBILE
		case "home":
			pn.Type = pb.Person_HOME
		case "work":
			pn.Type = pb.Person_WORK
		default:
			fmt.Printf("Unknown phone type %q. Using default. \n", ptype)
		}
		p.Phones = append(p.Phones, pn)
	}
	return p, nil
}

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("Usage: %s ADDRESS_BOOK_FILE\n", os.Args[0])
	}
	fname := os.Args[1]

	in, err := ioutil.ReadFile(fname)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Fprintf(os.Stdout, "The file %s does not exist, it will be created now\n", fname)
		} else {
			log.Fatal("An error happened while reading the file")
		}
	}
	ab := &pb.AddressBook{}

	if err := proto.Unmarshal(in, ab); err != nil {
		log.Fatalln("Failed to parse address book", err)
	}

	addr, err := promptForAddress(os.Stdin)
	if err != nil {
		log.Fatalln("Error prompting the address", err)
	}

	ab.People = append(ab.People, addr)
	out, err := proto.Marshal(ab)
	if err != nil {
		log.Fatalln("Failed to encode address book", err)
	}
	if err := ioutil.WriteFile(fname, out, 0644); err != nil {
		log.Fatalln("Failed to write to the file", err)
	}
}
