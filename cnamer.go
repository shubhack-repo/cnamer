package main
import (
	"bufio"
	"fmt"
	"os"
	"net"
	"strings"
	"sync"
	"flag"
)

func main(){

	concurrency := 20
	flag.IntVar(&concurrency,"c", 20, "Set the concurrency level")
	flag.Parse()

	var wg sync.WaitGroup
	jobs := make(chan string)
		
	for i := 0; i < concurrency ; i++ {
		wg.Add(1)

		go func(){
			for domain := range jobs {
			cname,err := net.LookupCNAME(domain)
			if err != nil {
				fmt.Println(domain,"No CNAME record")
				continue
			}
			fmt.Println(domain,strings.TrimSuffix(cname,"."))
			}
			wg.Done()
		}()
	}

	reader := bufio.NewScanner(os.Stdin)
	for reader.Scan(){
		jobs <- reader.Text()
	}

	close(jobs)

	wg.Wait()
}
