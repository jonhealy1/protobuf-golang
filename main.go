package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	complexpb "github.com/simplesteph/protobuf-example-go/src/complex"
	enumpb "github.com/simplesteph/protobuf-example-go/src/enum_example"
	simplepb "github.com/simplesteph/protobuf-example-go/src/simple"
)

func main() {
	sm := doSimple()
	readAndWriteDemo(sm)
	jsonDemo(sm)

	doEnum()

	doComplex()
}

func doComplex() {
	cm := complexpb.ComplexMessage{
		OneDummy: &complexpb.DummyMessage{
			Id:   1,
			Name: "first message",
		},
		MultipleDummy: []*complexpb.DummyMessage{
			&complexpb.DummyMessage{
				Id:   2,
				Name: "second message",
			},
			&complexpb.DummyMessage{
				Id:   3,
				Name: "third message",
			},
		},
	}
	fmt.Println(cm)
}

func doEnum() {
	em := enumpb.EnumMessage{
		Id:           42,
		DayOfTheWeek: enumpb.DayOfTheWeek_THURSDAY,
	}
	em.DayOfTheWeek = enumpb.DayOfTheWeek_MONDAY
	fmt.Println(em)
}

func jsonDemo(sm proto.Message) {
	smAsString := toJSON(sm)
	fmt.Println(smAsString)

	sm2 := &simplepb.SimpleMessage{}
	fromJSON(smAsString, sm2)
	fmt.Println("Successfully create proto structs")
}

func toJSON(pb proto.Message) string {
	marshaler := jsonpb.Marshaler{}
	out, err := marshaler.MarshalToString(pb)
	if err != nil {
		log.Fatalln("Can't convert to JSON", err)
		return ""
	}
	return out
}

func fromJSON(in string, pb proto.Message) {
	err := jsonpb.UnmarshalString(in, pb)
	if err != nil {
		log.Fatalln("Can't convert to JSON", err)
	}
}

func readAndWriteDemo(sm proto.Message) {
	writeToFile("simple.bin", sm)
	sm2 := &simplepb.SimpleMessage{}
	readFromFile("simple.bin", sm2)
	fmt.Println("Read the content:", sm2)
}

func writeToFile(fname string, pb proto.Message) error {
	out, err := proto.Marshal(pb)
	if err != nil {
		log.Fatalln("Can't serialize to bytes", err)
		return err
	}
	if err := ioutil.WriteFile(fname, out, 0644); err != nil {
		log.Fatalln("Can't write to file", err)
		return err
	}
	fmt.Println("Data has been written.")
	return nil
}

func readFromFile(fname string, pb proto.Message) error {
	in, err := ioutil.ReadFile(fname)
	if err != nil {
		log.Fatalln("Error reading file", err)
		return err
	}
	err2 := proto.Unmarshal(in, pb)
	if err2 != nil {
		log.Fatalln("Couldn't put bytes into protobuf struct", err)
		return err2
	}
	return nil
}

func doSimple() *simplepb.SimpleMessage {
	sm := simplepb.SimpleMessage{
		Id:         12345,
		IsSimple:   true,
		Name:       "My simple message",
		SampleList: []int32{1, 4, 7, 8},
	}
	sm.Name = "I renamed you"
	//fmt.Println(sm)

	//fmt.Println("The ID is:", sm.GetId())
	return &sm
}
