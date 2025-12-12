package main

import "fmt"

// func readInput() []string {
// 	_, filename, _, ok := runtime.Caller(0)
// 	if !ok {
// 		panic("cannot get current file info")
// 	}

// 	dir := filepath.Dir(filename)

// 	// path := filepath.Join(dir, "../sampleinput.txt")
// 	path := filepath.Join(dir, "../input.txt")

// 	file, err := os.Open(path)
// 	if err != nil {
// 		panic(err)
// 	}

// 	defer file.Close()

// 	scanner := bufio.NewScanner(file)
// 	text := make([]string, 0)
// 	for scanner.Scan() {
// 		line := scanner.Text()
// 		text = append(text, line)

// 	}
// 	return text
// }

func main() {
	matrix := [][]int{
		// A B C D E F G
		{1, 0, 0, 1, 0, 0, 1}, // Row A
		{1, 0, 0, 1, 0, 0, 0}, // Row B   <- in solution
		{0, 0, 0, 1, 1, 0, 1}, // Row C
		{0, 0, 1, 0, 1, 1, 0}, // Row D   <- in solution
		{0, 1, 1, 0, 0, 1, 1}, // Row E
		{0, 1, 0, 0, 0, 0, 1}, // Row F   <- in solution
	}
	dlx := NewDLX(matrix)
	dlx.Search()

	if !dlx.HasSolution {
		fmt.Printf("No Solution Found\n")
	} else {
		fmt.Printf("Solution rows:\n")
		for i := range dlx.Solution {
			fmt.Printf("%c ", 'A'+dlx.Solution[i])
		}
	}
}
