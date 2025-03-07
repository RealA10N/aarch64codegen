// filepath: /Users/alonkr/Developer/aarch64codegen/instructions/svc_test.go
package instructions_test

import (
	"testing"

	"alon.kr/x/aarch64codegen/instructions"
)

func TestSvc(t *testing.T) {
	// SVC #0 (typically used for Linux system calls)
	AssertExpectedInstruction(t, "SVC #0", instructions.SVC(0))

	// SVC #1 (often used for system calls in other OS'es)
	AssertExpectedInstruction(t, "SVC #1", instructions.SVC(1))

	// SVC with larger immediate value
	AssertExpectedInstruction(t, "SVC #65535", instructions.SVC(65535))
}
