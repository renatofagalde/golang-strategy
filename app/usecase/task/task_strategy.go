package task

import "fmt"

type TaskStrategy interface {
	Run(data string) string
}

type StrategyRegistry struct {
	strategies map[string]TaskStrategy
}

func NewStrategy() *StrategyRegistry {
	return &StrategyRegistry{strategies: make(map[string]TaskStrategy)}
}

func (s *StrategyRegistry) Register(action string, strategy TaskStrategy) {
	s.strategies[action] = strategy
}

func (s *StrategyRegistry) Get(action string) (TaskStrategy, error) {
	strategy, exists := s.strategies[action]
	if !exists {
		return nil, fmt.Errorf("Strategy '%s' not found", action)
	}
	return strategy, nil
}
