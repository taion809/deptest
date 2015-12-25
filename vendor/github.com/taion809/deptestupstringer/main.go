package deptestupstringer

import (
	stringer "github.com/taion809/depteststringer"
	"strings"
)

func Upper(f stringer.Stringer) string {
	return strings.ToUpper(f.Get())
}
