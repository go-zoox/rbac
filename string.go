package rbac

// StringMatch returns a Matcher that returns true
// if the target string matches s.
func StringMatch(s string) Matcher {
	return func(target string) (bool, error) {
		return target == s, nil
	}
}

// NewStringPermission returns a Permission that use StringMatchers for the specified action and target strings.
func NewStringPermission(action, target string) Permission {
	return NewPermission(StringMatch(action), StringMatch(target))
}
