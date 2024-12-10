package main

import (
	"fmt"
	"os"
)

type File struct {
	fileId int
	start  int
	length int
	end    int
	isPart bool
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

	newFs := []File{
		files[0],
	}

	fileIndex := 1
	for i := len(files) - 1; i >= 0; i-- {
		file := files[i]

		if fileIndex == i {
			if file.length > 0 {
				newFs = append(newFs, File{
					start:  newFs[len(newFs)-1].end,
					end:    newFs[len(newFs)-1].end + file.length,
					length: file.length,
					fileId: file.fileId,
				})
			}
			break
		}

		space := files[fileIndex].start - newFs[len(newFs)-1].end

		if space <= file.length {
			newFs = append(newFs, File{
				start:  newFs[len(newFs)-1].end,
				end:    newFs[len(newFs)-1].end + space,
				length: space,
				fileId: file.fileId,
			})

			files[i].length -= space

			if files[i].length > 0 {
				i++
			}

			newFs = append(newFs, files[fileIndex])
			fileIndex++

		} else {
			newFs = append(newFs, File{
				start:  newFs[len(newFs)-1].end,
				end:    newFs[len(newFs)-1].end + file.length,
				length: file.length,
				fileId: file.fileId,
			})
		}
	}

	// for _, fs := range newFs {
	// 	for i := 0; i < fs.length; i++ {
	// 		fmt.Print(string([]byte{byte(fs.fileId + '0')}))
	// 	}
	// }

	var sum int = 0
	for _, fs := range newFs {
		for i := fs.start; i < fs.end; i++ {
			sum += i * fs.fileId
		}
	}

	fmt.Printf("Checksum is: %d", sum)
}
