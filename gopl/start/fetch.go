// Fetch prints the content found at a URL.
package main

//func main() {
//	urls := [100]string{"http://baidu.com"}
//	for _, url := range urls {
//		resp, err := http.Get(url)
//		fmt.Println(resp)
//		if err != nil {
//			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
//			os.Exit(1)
//		}
//		b, err := ioutil.ReadAll(resp.Body)
//		resp.Body.Close()
//		if err != nil {
//			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
//			os.Exit(1)
//		}
//		fmt.Printf("%s", b)
//	}
//}
