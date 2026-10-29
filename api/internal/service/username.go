package service

import (
	"context"
	"errors"
	"regexp"
	"strconv"
	"strings"

	"github.com/google/uuid"
	"github.com/marcoshack/taskwondo/internal/model"
)

// usernamePattern is the accepted username shape after normalization:
// starts with a letter or digit, then lowercase letters, digits, '-' or '_'.
var usernamePattern = regexp.MustCompile(`^[a-z0-9][a-z0-9_-]{1,29}$`)

var usernameInvalidChars = regexp.MustCompile(`[^a-z0-9_-]`)

// normalizeUsername lowercases and sanitizes a raw identifier into a valid
// username, returning false when nothing usable remains.
func normalizeUsername(raw string) (string, bool) {
	u := usernameInvalidChars.ReplaceAllString(strings.ToLower(strings.TrimSpace(raw)), "")
	u = strings.TrimLeft(u, "-_")
	if len(u) > 30 {
		u = u[:30]
	}
	return u, usernamePattern.MatchString(u)
}

// usernameFromEmail derives a username candidate from an email's local part.
func usernameFromEmail(email string) string {
	if u, ok := normalizeUsername(strings.SplitN(email, "@", 2)[0]); ok {
		return u
	}
	return "user"
}

type usernameChecker interface {
	GetByLogin(ctx context.Context, identifier string) (*model.User, error)
}

// uniqueUsername returns base or a -N suffixed variant not yet taken.
func uniqueUsername(ctx context.Context, users usernameChecker, base string) string {
	candidate := base
	for i := 2; i <= 100; i++ {
		_, err := users.GetByLogin(ctx, candidate)
		if errors.Is(err, model.ErrNotFound) {
			return candidate
		}
		if err != nil {
			return candidate
		}
		candidate = base + "-" + strconv.Itoa(i)
	}
	return base + "-" + uuid.NewString()[:8]
}
