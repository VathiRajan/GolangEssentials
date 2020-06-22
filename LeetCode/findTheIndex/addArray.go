    package main
     
    import (
    	"bufio"
    	"bytes"
    	"fmt"
    	"os"
    	"strconv"
    )
     
    func main() {
    	var buffer bytes.Buffer
     
    	scanner := bufio.NewScanner(os.Stdin)
    	scanner.Split(bufio.ScanWords)
     
    	scanner.Scan()
    	N, _ := strconv.Atoi(scanner.Text())
     
    	arr := make([]int, N)
    	for i := 0; i < N; i++ {
    		scanner.Scan()
    		arr[i], _ = strconv.Atoi(scanner.Text())
    	}
     
    	for i := 0; i < N-1; i++ {
    		scanner.Scan()
    		x, _ := strconv.Atoi(scanner.Text())
    		buffer.WriteString(strconv.Itoa(arr[i]+x) + " ")
    	}
    	scanner.Scan()
    	x, _ := strconv.Atoi(scanner.Text())
    	buffer.WriteString(strconv.Itoa(arr[N-1]+x) + "\n")
     
    	fmt.Print(buffer.String())
    }