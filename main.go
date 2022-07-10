package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"io/ioutil"
	"log"
	"os"
	"protobufer/pb/gen/pb"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("Run app with arg: path to pb files")
	}
	dir := os.Args[1]
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fmt.Printf("File %s/%s\n", dir, file.Name())
		log.Println(tryDecode(dir + "/" + file.Name()))
	}
}

func tryDecode(fileName string) (string, error) {
	in, err := ioutil.ReadFile(fileName)
	if err != nil {
		return fmt.Sprintf("File %s not readed", fileName), nil
	}

	city := &pb.Cities{}
	if err := proto.Unmarshal(in, city); err == nil {
		log.Printf("%v", city)
		return "This is cities!", nil
	}
	name := &pb.Names{}
	if err := proto.Unmarshal(in, name); err == nil {
		log.Printf("%v", name)
		return "This is names!", nil
	}
	person := &pb.Person{}
	if err := proto.Unmarshal(in, person); err == nil {
		log.Printf("%v", person)
		return "This is person!", nil
	}
	points := &pb.Points{}
	if err := proto.Unmarshal(in, points); err == nil {
		log.Printf("%v", points)
		return "This is points!", nil
	}
	teams := &pb.Teams{}
	if err := proto.Unmarshal(in, teams); err == nil {
		log.Printf("%v", teams)
		return "This is teams!", nil
	}
	return "Unknown proto!", err
}
