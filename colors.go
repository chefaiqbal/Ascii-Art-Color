package ascii

func GetColor(color string) string {
	codes := map[string]string{
		"black":   "\u001b[30m",
		"red":     "\u001b[31m",
		"green":   "\u001b[32m",
		"yellow":  "\u001b[33m",
		"blue":    "\u001b[34m",
		"magenta": "\u001b[35m",
		"cyan":    "\u001b[36m",
		"white":   "\u001b[37m",
		"orange":  "\u001b[38;5;208m",
		"reset":   "\u001b[0m",
	}
	return codes[color]
}

func IsColorSupported(color string) bool {
	supportedColors := map[string]bool{
		"black":   true,
		"red":     true,
		"green":   true,
		"yellow":  true,
		"blue":    true,
		"magenta": true,
		"cyan":    true,
		"white":   true,
		"orange":   true,
		"reset":   true,
	}

	_, ok := supportedColors[color]
	return ok
}
