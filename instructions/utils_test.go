package instructions_test

import (
	"encoding/binary"
	"os"
	"os/exec"
	"testing"

	"alon.kr/x/aarch64codegen/instructions"
	"github.com/stretchr/testify/assert"
)

func writeAssemblySource(t *testing.T, tempDir, asm string) (srcFile *os.File) {
	srcFile, err := os.CreateTemp(tempDir, "*.s")
	assert.NoError(t, err)
	defer srcFile.Close()

	_, err = srcFile.WriteString(asm)
	assert.NoError(t, err)

	return srcFile
}

func assembleSourceFile(t *testing.T, tempDir string, srcFile *os.File) (objFile *os.File) {
	objFile, err := os.CreateTemp(tempDir, "*.o")
	assert.NoError(t, err)
	defer objFile.Close()

	command := exec.Command("aarch64-linux-gnu-as", "-o", objFile.Name(), srcFile.Name())
	if err := command.Run(); err != nil {
		t.Fatalf("Failed to assemble: %s", err)
	}

	return objFile
}

func objectFileToRawBinary(t *testing.T, tempDir string, objFile *os.File) (binFile *os.File) {
	binFile, err := os.CreateTemp(tempDir, "*.bin")
	assert.NoError(t, err)
	defer binFile.Close()

	command := exec.Command("aarch64-linux-gnu-objcopy", "-O", "binary", objFile.Name(), binFile.Name())
	if err := command.Run(); err != nil {
		t.Fatalf("Failed to extract raw binary: %s", err)
	}

	return binFile
}

func rawBinaryFileToUint32(t *testing.T, binFile *os.File) uint32 {
	data, err := os.ReadFile(binFile.Name())
	assert.NoError(t, err)
	assert.Len(t, data, 4)
	return binary.LittleEndian.Uint32(data)
}

func assembleInstruction(t *testing.T, asm string) uint32 {
	tempDir, err := os.MkdirTemp("", "aarch64codegen_test_*")
	assert.NoError(t, err)
	defer os.RemoveAll(tempDir)

	srcFile := writeAssemblySource(t, tempDir, asm)
	objFile := assembleSourceFile(t, tempDir, srcFile)
	binFile := objectFileToRawBinary(t, tempDir, objFile)

	return rawBinaryFileToUint32(t, binFile)
}

func AssertExpectedInstruction(
	t *testing.T,
	expected string,
	actual instructions.Instruction,
) {
	expectedBinary := assembleInstruction(t, expected)
	assert.Equal(t, expectedBinary, actual.Binary())
	assert.Equal(t, expected, actual.String())
}
