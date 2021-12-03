package problemsolvers

type ProblemSolver interface {
	Initialize(initData string) error
	Solve() (string, error)
}
