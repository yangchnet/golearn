package main

import (
	"fmt"
	"os"
	"sort"
	"text/tabwriter"
	"time"
	)

/**
    @Destription:学习对于sort.Interface接口的使用，其定义如下： 
		type Interface interface {
			Len() int
			Less(i, j int) bool
			Swap(i, j int)
		}
	要想使用这个接口，必须对某个类型定义这三个方法
	@Author:lichang
    @Date:2020/1/16
    @Time:下午8:54
*/

// 定义基础类型
type Track struct {
	Title string
	Artist string
	Album string
	Year int
	Length time.Duration
}

var tracks = []*Track{
	{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
	{"Go", "Moby", "Moby", 1992, length("3m37s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
}


// sort.Interface 的具体类型不一定是切片型
type customSort struct {
	t    []*Track
	less func(x, y *Track) bool
}
func (x customSort) Len() int           { return len(x.t) }
func (x customSort) Less(i, j int) bool { return x.less(x.t[i], x.t[j]) }
func (x customSort) Swap(i, j int)      { x.t[i], x.t[j] = x.t[j], x.t[i] }

// 排序参考为Artist
type byArtist []*Track
func (x byArtist) Len()int {return len(x)}
func (x byArtist) Less(i, j int) bool {return x[i].Artist < x[j].Artist}
func (x byArtist) Swap(i, j int) {x[i], x[j] = x[j], x[i]}

// 排序参考为Year
type byYear []*Track
func (x byYear) Len()int {return len(x)}
func (x byYear) Less(i, j int) bool {return x[i].Year < x[j].Year}
func (x byYear) Swap(i ,j int) {x[i], x[j] = x[j], x[i]}

func length(s string) time.Duration{
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

func printTracks(tracks []*Track){
	/**
	    @Description:格式化输出
	    @Parm:
	    	tracks: 输入的表格
	    @Return:
			
	 */
	const format = "%v\t%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Title", "Artist", "Alblum", "Year", "Length")
	fmt.Fprintf(tw, format, "-----", "------", "------", "------", "------")
	for _, t := range tracks {
		fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
	}
	tw.Flush()
}

func main() {
	fmt.Println("byArtist:")
	sort.Sort(byArtist(tracks))
	printTracks(tracks)

	fmt.Println("\nReverse(byArtist):")
	sort.Sort(sort.Reverse(byArtist(tracks)))
	printTracks(tracks)

	fmt.Println("\nbyYear:")
	sort.Sort(byYear(tracks))
	printTracks(tracks)

	fmt.Println("\nCustom:")
	//!+customcall
	sort.Sort(customSort{tracks, func(x, y *Track) bool {
		// 匿名排序函数，按标题、年、运行时间
		if x.Title != y.Title {
			return x.Title < y.Title
		}
		if x.Year != y.Year {
			return x.Year < y.Year
		}
		if x.Length != y.Length {
			return x.Length < y.Length
		}
		return false
	}})
	//!-customcall
	printTracks(tracks)
}






