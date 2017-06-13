package db

type step0 struct{}

func (s step0) StepNumber() int { return 0 }
func (s step0) Check() bool     { return true }
func (s step0) Apply() {

}
