package dupclean

import "bufio"
import "fmt"
import "os"
import "strings"

// ConfirmDelete asks user for confirmation
func ConfirmDelete(file string) bool {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Delete %s? [y/N]: ", file)
	resp, _ := reader.ReadString('\n')
	resp = strings.TrimSpace(resp)
	return resp == "y" || resp == "Y"
}
