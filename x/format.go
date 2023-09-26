package x

import (
	"fmt"
)

func MessageFromMsgAndArgs(msgAndArgs ...any) string {
	if len(msgAndArgs) == 0 || msgAndArgs == nil {
		return ""
	}
	if s, ok := msgAndArgs[0].(string); ok {
		if len(msgAndArgs) == 1 {
			return s
		}
		return fmt.Sprintf(s, msgAndArgs[1:]...)
	}
	return fmt.Sprintf("%#v", msgAndArgs...)
}
