package manpage

// SubCommand defines a sub command
type SubCommand struct {
	Name        string
	Description string
	Flags       []*Flag // NOTE: this doesn't work yet
}
