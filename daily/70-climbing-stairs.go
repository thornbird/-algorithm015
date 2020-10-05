func climbStairs(n int) int {
    if n <= 2 {
        return n
    }
    
    oneStepsBefore := 2
    twoStepsBefore := 1
    allSteps := 0

    for i:=3; i<=n; i++ {
		allSteps = twoStepsBefore + oneStepsBefore
		twoStepsBefore = oneStepsBefore
        oneStepsBefore = allSteps
    }

    return allSteps
}