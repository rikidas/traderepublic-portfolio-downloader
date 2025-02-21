package transaction

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var ErrNoMatch = errors.New("value did not match the pattern")

func ParseFloatWithPeriod(src string) (float64, error) {
	pattern := regexp.MustCompile(`^[^\d]*(\d+)\.?(\d*)$`)
	matches := pattern.FindStringSubmatch(src)

	if len(matches) == 0 {
		return 0, ErrNoMatch
	}

	value := matches[1] + "." + matches[2]

	valueFloat, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return 0, fmt.Errorf("could not parse float from '%s': %w", value, err)
	}

	return valueFloat, nil
}

// ParseFloatWithComma converts a number with a decimal comma and removes extra symbols
func ParseFloatWithComma(src string, isNegative bool) (float64, error) {
	// Step 1: Remove currency symbols (€,$,MXN,COP) and surrounding spaces
	src = strings.ReplaceAll(src, "€", "")
	src = strings.ReplaceAll(src, "$", "")
	src = strings.ReplaceAll(src, "COP", "")
	src = strings.ReplaceAll(src, "MXN", "")
	src = strings.ReplaceAll(src, "USD", "")
	src = strings.ReplaceAll(src, "+", "")
	src = strings.TrimSpace(src) // Trim spaces before and after

	// Step 2: Remove the '%' symbol if present
	if strings.Contains(src, "%") {
		src = strings.ReplaceAll(src, "%", "")
	}

	// Step 4: Regular expression to capture numbers and decimal commas while ignoring any letters
	pattern := regexp.MustCompile(`[^0-9,.-]`) // Remove any character that is not a number, comma, dot, or sign
	src = pattern.ReplaceAllString(src, "")    // Remove letters and other non-numeric characters

	// Step 3: If no numbers are present (empty string), assign "0"
	if src == "" {
		src = "0"
	}

	// Step 5: Simpler regular expression to handle numbers with commas and dots
	pattern = regexp.MustCompile(`^\s*([+-]?\d+(?:\.\d{3})*)(?:,(\d+))?\s*$`)
	matches := pattern.FindStringSubmatch(src)

	if len(matches) == 0 {
		return 0, fmt.Errorf("value did not match the pattern: '%s'", src)
	}

	// Remove thousand separators (dots) and convert the decimal comma to a dot
	value := strings.ReplaceAll(matches[1], ".", "") // Remove dots

	// If there's a decimal part, append it after replacing the comma with a dot
	if matches[2] != "" {
		value += "." + matches[2]
	}

	// Convert the string to a float
	valueFloat, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return 0, fmt.Errorf("could not parse float from '%s': %w", value, err)
	}

	// If the original text contained "%", divide by 100
	if strings.Contains(src, "%") {
		valueFloat /= 100
	}

	// Apply a negative sign if necessary
	if isNegative {
		valueFloat = -valueFloat
	}

	return valueFloat, nil
}

func ParseNumericValueFromString(src string) (string, error) {
	pattern := regexp.MustCompile(`(\d+\.?\d*,?\d+)`)
	matches := pattern.FindStringSubmatch(src)

	if len(matches) == 0 {
		return "", ErrNoMatch
	}

	return matches[1], nil
}
