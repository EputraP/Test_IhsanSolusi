package service

import "github.com/EputraP/Test_IhsanSolusi/internal/repository"

type TestService interface {
	TestService() (string, error)
}

type testService struct {
	testRepo repository.TestRepository
}

type TestServiceConfig struct {
	TestRepo repository.TestRepository
}

func NewTestService(config TestServiceConfig) TestService {
	return &testService{
		testRepo: config.TestRepo,
	}
}

func (s testService) TestService() (string, error) {

	query, _ := s.testRepo.TestRepo()

	return query, nil

}
