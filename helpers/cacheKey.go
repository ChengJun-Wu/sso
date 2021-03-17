package helpers

func cacheKey(key string) string {
	return "sso_" + key
}

func FailedFrequency() string {
	return cacheKey("failed_frequency")
}