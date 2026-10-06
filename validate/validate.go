// Package validate provides simple syntactic validators for common
// user-input formats.
//
// Homework — Task 2 (Lesson 6: File I/O, JSON and Testing):
// Implement ValidateEmail and/or ValidatePhone below (your mentor may
// ask for just one) and extend the test tables in validate_test.go to
// at least 8 cases each, including edge cases.
package validate

import (
	"strings"
	"unicode"
)

// ValidateEmail reports whether s is a syntactically valid email address.
//
// We accept a normal ASCII-style address of the form local@domain.tld,
// with exactly one @, a non-empty local part, and a domain containing at
// least one dot. Trailing dots, consecutive dots, and Unicode characters
// are rejected to keep the validation simple and predictable.
func ValidateEmail(s string) bool {
	if s == "" || strings.ContainsAny(s, " \t\r\n") {
		return false
	}
	if strings.Count(s, "@") != 1 {
		return false
	}

	for _, r := range s {
		if r > unicode.MaxASCII {
			return false
		}
	}

	parts := strings.SplitN(s, "@", 2)
	local := parts[0]
	domain := parts[1]
	if local == "" || domain == "" {
		return false
	}
	if strings.HasPrefix(local, ".") || strings.HasSuffix(local, ".") || strings.Contains(local, "..") {
		return false
	}
	if strings.HasPrefix(domain, ".") || strings.HasSuffix(domain, ".") || strings.Contains(domain, "..") {
		return false
	}
	if !strings.Contains(domain, ".") {
		return false
	}

	for _, ch := range local {
		if (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') || (ch >= '0' && ch <= '9') {
			continue
		}
		switch ch {
		case '.', '-', '_', '+':
			continue
		case '!', '#', '$', '%', '&', '\'', '*', '/', '=', '?', '^', '`', '{', '|', '}', '~':
			continue
		default:
			return false
		}
	}

	labels := strings.Split(domain, ".")
	for _, label := range labels {
		if label == "" {
			return false
		}
		if strings.HasPrefix(label, "-") || strings.HasSuffix(label, "-") {
			return false
		}
		for _, ch := range label {
			if (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') || (ch >= '0' && ch <= '9') || ch == '-' {
				continue
			}
			return false
		}
	}

	return true
}

// ValidatePhone reports whether s is a syntactically valid phone number.
//
// We accept Ukrainian mobile numbers in the common forms 0501234567,
// 050-123-4567, 380501234567, and +380501234567. Letters, whitespace,
// dots and parentheses are rejected, and only the + prefix or a leading 0/380
// country format is accepted.
func ValidatePhone(s string) bool {
	if s == "" || strings.ContainsAny(s, " \t\r\n") {
		return false
	}
	if strings.ContainsAny(s, "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ") {
		return false
	}
	if strings.ContainsAny(s, "().") {
		return false
	}

	cleaned := strings.ReplaceAll(s, "-", "")

	if cleaned == "" {
		return false
	}
	if strings.HasPrefix(cleaned, "+") {
		cleaned = strings.TrimPrefix(cleaned, "+")
		if len(cleaned) != 12 {
			return false
		}
		return strings.HasPrefix(cleaned, "380") && isDigits(cleaned)
	}
	if isDigits(cleaned) {
		if len(cleaned) == 10 && strings.HasPrefix(cleaned, "0") {
			return true
		}
		if len(cleaned) == 12 && strings.HasPrefix(cleaned, "380") {
			return true
		}
	}
	return false
}

func isDigits(s string) bool {
	if s == "" {
		return false
	}
	for _, ch := range s {
		if ch < '0' || ch > '9' {
			return false
		}
	}
	return true
}
