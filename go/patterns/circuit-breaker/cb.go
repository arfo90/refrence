package circuitbreaker

import (
	"context"
	"errors"
	"sync"
	"time"
)

type Circuit func(context.Context) (string, error)

func Breaker(circuit Circuit, failureThreshold uint) Circuit {
	var consecutiveFailuer int = 0
	var lastAttemp = time.Now()
	var m sync.RWMutex

	return func(c context.Context) (string, error) {
		m.RLock()

		d := consecutiveFailuer - int(failureThreshold)

		if d >= 0 {
			retryAt := lastAttemp.Add(time.Second * 2 << d)
			if !time.Now().After(retryAt) {
				m.RUnlock()
				return "", errors.New("service unreachable")
			}
		}

		m.RUnlock()

		reponse, err := circuit(c)

		m.Lock()
		defer m.Unlock()

		if err != nil {
			consecutiveFailuer++
			return reponse, err
		}

		consecutiveFailuer = 0

		return reponse, nil
	}
}
