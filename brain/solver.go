package brain

import (
	"bufio"
	"fmt"
)

func SolveProblem() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("What is your problem?: ")

	problem, _ := reader.ReadString('\n')

	fmt.Println("Your problem is: ", problem)
}
