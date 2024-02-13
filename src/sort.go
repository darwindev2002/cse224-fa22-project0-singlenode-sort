package main

import (
	"bufio"
	"bytes"
	"io"
	"log"
	"os"
	"slices"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	if len(os.Args) != 3 {
		log.Fatalf("Usage: %v inputfile outputfile\n", os.Args[0])
	}

	log.Printf("Sorting %s to %s\n", os.Args[1], os.Args[2])

	// Open input file
	inFile, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer inFile.Close()

	// Create output file
	outFile, err := os.OpenFile(os.Args[2], os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer outFile.Close()

	buf := make([]byte, 100)
	reader := bufio.NewReader(inFile)
	writer := bufio.NewWriter(outFile)
	// dict := make(map[string][]byte)
	lst := [][]byte{}
	// count := 0
	for {
		// n, err := reader.Read(buf) // n = # bytes just read
		// When using reader.Read, sometimes it doesn't read the full buf bytes
		// Solution: use io.ReadFull to ensure a full 100-byte read
		n, err := io.ReadFull(reader, buf)
		if err != nil && err != io.EOF {
			log.Fatal(err)
		}
		if err == io.EOF {
			break
		}

		// Old approach - using a map[string]bytes
		// key, val := string(buf[:10]), buf[10:]
		// fmt.Println(key, val)
		// dict[key] = make([]byte, 90)
		// copy(dict[key], val)
		// fmt.Println("Created:", dict[key])
		// fmt.Println("We have buf[:n] =", buf[:n])
		// tmp := make([]byte, 0)
		// tmp = append(tmp, buf[:n]...)

		// Better approach - simply a list of bytes, and customized sort function
		var tmp []byte
		tmp = append(tmp, buf[:n]...)
		lst = append(lst, tmp)

		// fmt.Println("Entry", count, "done.")
		// count++
	}

	// Old approach
	// keys := make([]string, 0, len(dict))
	// for k := range dict {
	// 	keys = append(keys, k)
	// }
	// slices.Sort(keys)
	// for _, k := range keys {
	// 	_, err := writer.WriteString(k)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	// fmt.Println(string(k), dict[k])
	// 	_, err = writer.Write(dict[k])
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// }

	// New approach
	slices.SortFunc(lst, func(i, j []byte) int {
		// Debug use - check indicies and lenghts
		// fmt.Printf("Lengths: i at lst[%v] with len = %v, j at lst[%v] with len = %v \n",
		// 	slices.IndexFunc(lst, func(b []byte) bool { return bytes.Equal(b, i) }),
		// 	len(i),
		// 	slices.IndexFunc(lst, func(b []byte) bool { return bytes.Equal(b, j) }),
		// 	len(j))
		return bytes.Compare(i[:10], j[:10])
	})
	for _, j := range lst {
		_, err := writer.Write(j)
		if err != nil {
			log.Fatal(err)
		}
	}

	// Flush the sorted list from buffer
	writer.Flush()
}
