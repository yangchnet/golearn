package main
/**
 * @Description: 从所给出的HTML中提取出所有的链接
 * @author lichang
 * @date 20-1-6 
 * @time 下午4:35
*/
import (
	"fmt"
	"golang.org/x/net/html"
	"os"
)

func main(){
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprint(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	for _, link := range visit(nil, doc){
		fmt.Println(link)
	}
}


func visit(links []string, n *html.Node)[]string {
	if n.Type == html.ElementNode && n.Data =="a"{
		for _,a := range n.Attr{
			if a.Key == "href"{
				links = append(links, a.Val)
			}

		}
	}
	for c := n.FirstChild; c!= nil; c=c.NextSibling{
		links = visit(links, c)
	}
	return links
}

func visit1(links []string, n *html.Node)[]string {
	/**
	   @Description: 将上面的循环访问 n.FirstChild 链表改为递归访问，暂时不会改
	   @Param: 
	        
	   @Return: 
	        
	 */
	if n.Type == html.ElementNode && n.Data == "a"{
		for _, a := range n.Attr{
			if a.Key == "href"{
				links = append(links, a.Val)
			}
		}
	}

}