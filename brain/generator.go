package brain

type Solution struct {
	problem string
	idea    string
}

func GenerateIdeas(problem string, num int) []Solution {
	solutions := make([]Solution, num)

	// generate num ideas using chat completion
	for i := 0; i < num; i++ {
		// generate a random idea
		idea := "random idea"

		solutions[i] = Solution{problem, idea}
	}

	return solutions
}
