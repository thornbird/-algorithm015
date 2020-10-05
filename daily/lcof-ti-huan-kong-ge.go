import "fmt"

func replaceSpace(s string) string {
    res := make([]byte, 0)
    replace := []byte{'%', '2', '0'}

    for i:=0; i<len(s); i++ {
        if s[i] == ' '{
            res = append(res, replace...)
        } else {
            res = append(res, s[i])
        }
    }
    return fmt.Sprintf("%s", res)
}