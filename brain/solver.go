package brain

import (
	"bufio"
	"fmt"
	"os"
)

func SolveProblem() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("What is your problem?: ")

	problem, _ := reader.ReadString('\n')

	fmt.Println("Your problem is: ", problem)

	solutions := GenerateIdeas(problem, 5)

	fmt.Println("Here are some ideas: ")
	for _, solution := range solutions {
		fmt.Println(solution)
	}
}
