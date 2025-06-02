package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/mjschwenne/pollux-report/proto-test/gopb"
	"google.golang.org/protobuf/proto"
)

func main() {
	var i32o = flag.String("i32o", "", "Output file name for i32 struct")
	var i32i = flag.String("i32i", "", "File name to read a protobuf blob from")
	var i64o = flag.String("i64o", "", "Output file name for i64 struct")
	var u32i = flag.String("u32i", "", "File name to read a protobuf blob from")
	var u32o = flag.String("u32o", "", "Output file name for u32 struct")
	var boolo = flag.String("boolo", "", "Output file name for bool struct")
	var booli = flag.String("booli", "", "File name to read a protobuf blob from")
	var s32o = flag.String("s32o", "", "Output file name for s32 struct")
	var s32i = flag.String("s32i", "", "File name to read a protobuf blob from")
	var s64o = flag.String("s64o", "", "Output file name for s32 struct")
	var s64i = flag.String("s64i", "", "File name to read a protobuf blob from")
	flag.Parse()

	if *i32o != "" {
		test := &gopb.I32{
			Field: -9,
		}

		out, err := proto.Marshal(test)
		if err != nil {
			log.Fatalln("Failed to encode i32 struct:", err)
		}
		if err := os.WriteFile(*i32o, out, 0644); err != nil {
			log.Fatalln("Failed to write i32 struct to disk:", err)
		}

		fmt.Printf("i32 encoded value: %d (dec)\n", test.Field)
		fmt.Printf("i32 encoded value: %064b (bin)\n", uint32(test.Field))
	}

	if *i64o != "" {
		test := &gopb.I64{
			Field: (int64(1<<34) - 1),
		}

		out, err := proto.Marshal(test)
		if err != nil {
			log.Fatalln("Failed to encode i64 struct:", err)
		}
		if err := os.WriteFile(*i64o, out, 0644); err != nil {
			log.Fatalln("Failed to write i64 struct to disk:", err)
		}

		fmt.Printf("i64 encoded value: %d (dec)\n", test.Field)
		fmt.Printf("i64 encoded value: %064b (bin)\n", uint64(test.Field))
	}

	if *boolo != "" {
		test := &gopb.Bool{
			Field: true,
		}

		out, err := proto.Marshal(test)
		if err != nil {
			log.Fatalln("Failed to encode bool struct:", err)
		}
		if err := os.WriteFile(*boolo, out, 0644); err != nil {
			log.Fatalln("Failed to write bool struct to disk:", err)
		}
	}

	if *s32o != "" {
		test := &gopb.S32{
			Field: 2,
		}

		out, err := proto.Marshal(test)
		if err != nil {
			log.Fatalln("Failed to encode s32 struct:", err)
		}
		if err := os.WriteFile(*s32o, out, 0644); err != nil {
			log.Fatalln("Failed to write s32 struct to disk:", err)
		}

		fmt.Printf("s32 encoded value: %v (dec)\n", test.Field)
		fmt.Printf("s32 encoded value: %064b (bin)\n", (test.Field<<1)^(test.Field>>31))
	}

	if *s64o != "" {
		test := &gopb.S64{
			Field: -(int64(1<<23) - 3),
		}

		out, err := proto.Marshal(test)
		if err != nil {
			log.Fatalln("Failed to encode s64 struct:", err)
		}
		if err := os.WriteFile(*s64o, out, 0644); err != nil {
			log.Fatalln("Failed to write s64 struct to disk:", err)
		}

		fmt.Printf("s64 encoded value: %v (dec)\n", test.Field)
		fmt.Printf("s64 encoded value: %064b (bin)\n", (test.Field<<1)^(test.Field>>63))
	}

	if *u32o != "" {
		test := &gopb.U32{
			Field: 23,
		}

		out, err := proto.Marshal(test)
		if err != nil {
			log.Fatalln("Failed to encode u32 struct:", err)
		}
		if err := os.WriteFile(*u32o, out, 0644); err != nil {
			log.Fatalln("Failed to write u32 struct to disk:", err)
		}

		fmt.Printf("u32 encoded value: %v (dec)\n", test.Field)
		fmt.Printf("u32 encoded value: %064b (bin)\n", test.Field)
	}

	if *i32i != "" {
		in, err := os.ReadFile(*i32i)
		if err != nil {
			log.Fatalln("Error reading i32 input file:", err)
		}

		test := &gopb.I32{}
		if err := proto.Unmarshal(in, test); err != nil {
			log.Fatalln("Failed to parse i32 test input:", err)
		}

		fmt.Printf("i32 encoded value: %064b (bin)\n", uint32(test.Field))
		fmt.Printf("i32 encoded value: %d (dec)\n", test.Field)
	}

	if *u32i != "" {
		in, err := os.ReadFile(*u32i)
		if err != nil {
			log.Fatalln("Error reading u32 input file:", err)
		}

		test := &gopb.U32{}
		if err := proto.Unmarshal(in, test); err != nil {
			log.Fatalln("Failed to parse u32 test input:", err)
		}

		fmt.Printf("u32 decoded value: %v\n", test.Field)
	}

	if *booli != "" {
		in, err := os.ReadFile(*booli)
		if err != nil {
			log.Fatalln("Error reading bool input file:", err)
		}

		test := &gopb.Bool{}
		if err := proto.Unmarshal(in, test); err != nil {
			log.Fatalln("Failed to parse bool test input:", err)
		}

		fmt.Printf("bool decoded value: %v\n", test.Field)
	}

	if *s32i != "" {
		in, err := os.ReadFile(*s32i)
		if err != nil {
			log.Fatalln("Error reading s32 input file:", err)
		}

		test := &gopb.S32{}
		if err := proto.Unmarshal(in, test); err != nil {
			log.Fatalln("Failed to parse s32 test input:", err)
		}

		fmt.Printf("s32 decoded value: %064b (bin)\n",
			uint32((test.Field<<1)^(test.Field>>31)))
		fmt.Printf("s32 decoded value: %v (dec)\n", test.Field)
	}

	if *s64i != "" {
		in, err := os.ReadFile(*s64i)
		if err != nil {
			log.Fatalln("Error reading s64 input file:", err)
		}

		test := &gopb.S64{}
		if err := proto.Unmarshal(in, test); err != nil {
			log.Fatalln("Failed to parse s64 test input:", err)
		}

		fmt.Printf("s64 decoded value: %064b (bin)\n",
			uint64((test.Field<<1)^(test.Field>>63)))
		fmt.Printf("s64 decoded value: %v (dec)\n", test.Field)
	}
}
