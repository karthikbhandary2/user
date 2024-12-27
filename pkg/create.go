package create

import (
    "fmt"
    "os/exec"
    // "bytes"
)

const ShellToUse = "bash"

func CreateUser(username, name string) error {
    // var stdout bytes.Buffer
    // var stderr bytes.Buffer
    cmd := exec.Command(ShellToUse, "-c", fmt.Sprintf("/home/u2006092/user/user/pkg/adduser.sh %s %s", username, name))
    // cmd.Stdout = &stdout
    // cmd.Stderr = &stderr
    err := cmd.Run()
    if err != nil {
        fmt.Println("Error:", err)
		return err
    }
    // fmt.Println("Stdout:", stdout.String())
    // fmt.Println("Stderr:", stderr.String())
	return nil
}