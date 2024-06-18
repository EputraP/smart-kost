package service

type TestService interface {
	Test() string
}

type testService struct {
}

func NewTestService() TestService {
	return &testService{}
}

func (ts *testService) Test() string {
	return "test"

}
