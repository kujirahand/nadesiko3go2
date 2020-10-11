package token

// TokenID is ID of Token
type TokenID int

// List of tokens
const (
	Unknown TokenID = iota
	EOF
	Comment
	Plus
	Minus
	Mul
	Div
	Mod
	PlusStr
)

var tokenNames = map[TokenID]string{
	Unknown: "Unknown",
	EOF:     "EOF",
	Comment: "Comment",
	Plus:    "Plus",
	Minus:   "Minus",
	Mul:     "Mul",
	Div:     "Div",
	Mod:     "Mod",
	PlusStr: "PlusStr",
}

// GetTokenName returns Token name
func GetTokenName(id TokenID) string {
	v, ok := tokenNames[id]
	if ok {
		return v
	}
	return "InvalidTokenID"
}
