package raspiConfig

import (
	"os/exec"
	"bytes"
	"strings"
)

func execute_command(prog string, arguments []string) (stdoutString string, stderrString string) {
	cmd := exec.Command(prog, arguments...);
	var outBuffer, errBuffer bytes.Buffer
	cmd.Stdout = &outBuffer;
	cmd.Stderr = &errBuffer;

	cmd.Wait()
	return outBuffer.String(), errBuffer.String();
}

func has_root()(hasRoot bool) {

	cmd := exec.Command("sudo", "true");
	err := cmd.Start();

	if err != nil { 
	    hasRoot = false;
	} else {
		hasRoot = true;
	}
	return;
}


func Expand_file_system()(success bool) {
	if !has_root() {
		return false;
	} 

	stdin, _ := execute_command("readlink", []string{"/dev/root"});
	if strings.Index(stdin, "mmcblk0p") == -1 {

	}
	return true;
}
