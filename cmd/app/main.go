package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	arg := os.Args[1]
	pwd, _ := os.Getwd()
	err := filepath.Walk(pwd+arg,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			name := info.Name()
			if name[len(name)-1] == 'o' && name[len(name)-2] == 'g' && name[len(name)-3] == '.' {
				readFile, err := os.Open(path)
				fileScanner := bufio.NewScanner(readFile)
				fileScanner.Split(bufio.ScanLines)
				var fileTextLines []string

				for fileScanner.Scan() {
					fileTextLines = append(fileTextLines, fileScanner.Text())
				}

				defer readFile.Close()

				for index, eachline := range fileTextLines {
					if strings.Contains(eachline, ").Get(") {
						i := strings.Index(eachline, ").Get(")
						s := eachline[i+5:]
						end := strings.Index(s, ")")
						params := s[:end+1]

						split := strings.Split(params, ",")

						if len(split) > 1 {
							// check for context.TODO()
							//first add context.TODO()
							if strings.Contains(params, ",") && params[1:strings.Index(params, ",")] != "context.TODO()" {
								params = params[0:1] + "context.TODO()," + params[1:]
							}
						}

						if len(split) == 1 {
							if strings.Contains(params, ")") && params[1:strings.Index(params, ")")] != "context.TODO()" {
								params = params[0:1] + "context.TODO()," + params[1:]
							}
						}

						split = strings.Split(params, ",")

						if len(split) == 2 {
							params = params[0:len(params)-1] + ", metav1.GetOptions{}" + params[len(params)-1:]
						}

						fileTextLines[index] = fileTextLines[index][:i+5] + params + s[end+1:]

					}
					if strings.Contains(eachline, ").Create(") {
						i := strings.Index(eachline, ").Create(")
						s := eachline[i+8:]
						end := strings.Index(s, ")")
						params := s[:end+1]

						split := strings.Split(params, ",")

						if len(split) > 1 {
							// check for context.TODO()
							//first add context.TODO()
							if strings.Contains(params, ",") && params[1:strings.Index(params, ",")] != "context.TODO()" {
								params = params[0:1] + "context.TODO()," + params[1:]
							}
						}

						if len(split) == 1 {
							if strings.Contains(params, ")") && params[1:strings.Index(params, ")")] != "context.TODO()" {
								params = params[0:1] + "context.TODO()," + params[1:]
							}
						}

						split = strings.Split(params, ",")

						if len(split) == 2 {
							params = params[0:len(params)-1] + ", metav1.CreateOptions{}" + params[len(params)-1:]
						}

						fileTextLines[index] = fileTextLines[index][:i+8] + params + s[end+1:]

					}
					if strings.Contains(eachline, ").List(") {
						i := strings.Index(eachline, ").List(")
						s := eachline[i+6:]
						end := strings.Index(s, ")")
						params := s[:end+1]
						split := strings.Split(params, ",")

						if len(split) > 1 {
							// check for context.TODO()
							//first add context.TODO()
							if strings.Contains(params, ",") && params[1:strings.Index(params, ",")] != "context.TODO()" {
								params = params[0:1] + "context.TODO()," + params[1:]
							}
						}

						if len(split) == 1 {
							if strings.Contains(params, ")") && params[1:strings.Index(params, ")")] != "context.TODO()" {
								params = params[0:1] + "context.TODO()," + params[1:]
							}
						}

						split = strings.Split(params, ",")
						// if strings.Contains(split[1], ".") {
						// 	optionIndex := strings.Index(split[1], ".")
						// 	if split[1][optionIndex+1 : optionIndex+12] != "ListOptions"
						// }
						if len(split) == 2 {
							if strings.Contains(split[1], ".") {
								optionIndex := strings.Index(split[1], ".")
								if split[1][optionIndex+1:optionIndex+12] != "ListOptions" {
									params = params[0:len(params)-1] + ", metav1.ListOptions{}" + params[len(params)-1:]
								}
							}
						}

						fileTextLines[index] = fileTextLines[index][:i+6] + params + s[end+1:]

					}
					if strings.Contains(eachline, ").Delete(") {
						i := strings.Index(eachline, ").Delete(")
						s := eachline[i+8:]
						end := strings.Index(s, ")")
						params := s[:end+1]

						split := strings.Split(params, ",")

						if len(split) > 1 {
							// check for context.TODO()
							//first add context.TODO()
							if strings.Contains(params, ",") && params[1:strings.Index(params, ",")] != "context.TODO()" {
								params = params[0:1] + "context.TODO()," + params[1:]
							}
						}

						if len(split) == 1 {
							if strings.Contains(params, ")") && params[1:strings.Index(params, ")")] != "context.TODO()" {
								params = params[0:1] + "context.TODO()," + params[1:]
							}
						}

						split = strings.Split(params, ",")

						if len(split) == 2 {
							params = params[0:len(params)-1] + ", metav1.DeleteOptions{}" + params[len(params)-1:]
						}

						fileTextLines[index] = fileTextLines[index][:i+8] + params + s[end+1:]

					}
					if strings.Contains(eachline, ").Update(") {
						i := strings.Index(eachline, ").Update(")
						s := eachline[i+8:]
						end := strings.Index(s, ")")
						params := s[:end+1]

						split := strings.Split(params, ",")

						if len(split) > 1 {
							// check for context.TODO()
							//first add context.TODO()
							if strings.Contains(params, ",") && params[1:strings.Index(params, ",")] != "context.TODO()" {
								params = params[0:1] + "context.TODO()," + params[1:]
							}
						}

						if len(split) == 1 {
							if strings.Contains(params, ")") && params[1:strings.Index(params, ")")] != "context.TODO()" {
								params = params[0:1] + "context.TODO()," + params[1:]
							}
						}

						split = strings.Split(params, ",")

						if len(split) == 2 {
							params = params[0:len(params)-1] + ", metav1.UpdateOptions{}" + params[len(params)-1:]
						}

						fileTextLines[index] = fileTextLines[index][:i+8] + params + s[end+1:]

					}

					output := strings.Join(fileTextLines, "\n")
					err = ioutil.WriteFile(path, []byte(output), 0644)
					if err != nil {
						log.Fatalln(err)
					}
				}
			}
			return nil
		})

	if err != nil {
		fmt.Println(err)
	}

}
