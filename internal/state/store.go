/*store.go has exactly four responsibilities.

Store

├── Store runtime state
├── Publish runtime state
├── Serve runtime state
└── Protect concurrent access  */
package state 
import(
   "context"
    "sync"
)
func (s *Store) PublishObservation(
	ctx context.Context,
	state *ObservationState,
) error {

	s.mu.Lock()
	defer s.mu.Unlock()

	s.observation = state

	return nil
}

func (s *Store) CurrentObservation(
	ctx context.Context,
) (*ObservationState, error) {

	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.observation, nil
}