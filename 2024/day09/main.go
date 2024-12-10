package main

import (
	"fmt"
	"os"
	"slices"
)

type File struct {
	fileId  int
	start   int
	length  int
	end     int
	checked bool
}

func main() {
	data, err := os.ReadFile("input")
	if err != nil {
		panic(fmt.Errorf("could not read input: %w", err))
	}

	files := []File{}

	index := 0
	fileId := 0

	addFile := true
	for _, d := range data {
		if addFile {
			files = append(files, File{
				fileId: fileId,
				start:  index,
				end:    index + int(d-'0'),
				length: int(d - '0'),
			})
			fileId++
		}

		index += int(d - '0')
		addFile = !addFile
	}

	// print(files)

	for i := len(files) - 1; i >= 0; i-- {
		file := files[i]

		if file.checked {
			continue
		}

		files[i].checked = true
		file.checked = true

		for j := 1; j < i; j++ {
			space := files[j].start - files[j-1].end

			if file.length <= space {
				files = slices.Delete(files, i, i+1)
				file.start = files[j-1].end
				file.end = file.start + file.length

				files = slices.Insert(files, j, file)

				i = len(files)

				// print(files)

				break
			}
		}
	}

	print(files)

	var sum int = 0
	for _, fs := range files {
		for i := fs.start; i < fs.end; i++ {
			sum += i * fs.fileId
		}
	}

	fmt.Printf("Checksum is: %d", sum)
}

func print(files []File) {
	for i := 0; i < len(files); i++ {
		for l := 0; l < files[i].length; l++ {
			fmt.Print(string([]byte{byte(files[i].fileId + '0')}))
		}

		if i < len(files)-1 {
			for l := files[i].end; l < files[i+1].start; l++ {
				fmt.Print(".")
			}
		}
	}

	fmt.Println()
}
