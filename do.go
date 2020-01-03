package tryhard

func (t *TryHard) Do(doJob Handler) {
	retryCount := 1

	for true {
		err := doJob()
		if err != nil {
			retryCount++
		}

		if t.isExceedMaximumRetry(retryCount) {
			// TODO: add handler to manage error that exceed maximum retry
			break
		}

		if err != nil {
			t.delay(t.backOff)
			continue
		}

		break
	}
}

func (t *TryHard) isExceedMaximumRetry(retryCount int) bool {
	return retryCount > t.maximumRetry
}

func (t *TryHard) DoWithKey(key string, doJob Handler) {
	t.doWithKey(key, 0, doJob)
}

func (t *TryHard) DoWithKeyAndTTL(key string, ttl int, doJob Handler) {
	t.doWithKey(key, ttl, doJob)
}

func (t *TryHard) doWithKey(key string, ttl int, doJob Handler) {
	for true {
		counter := t.store.GetCounter(key)
		if t.isExceedMaximumRetry(counter) {
			// TODO: add handler to manage error that exceed maximum retry
			break
		}

		err := doJob()
		if err != nil {
			t.store.AddCounterWithTTL(key, ttl)

			if !t.isExceedMaximumRetry(counter + 1) {
				t.delay(t.backOff)
			}
			continue
		}

		t.store.ClearCounter(key)
		break
	}
}
