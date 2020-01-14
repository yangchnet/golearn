package tempconv
/**
 * @Description: 研究接口的使用
 * @author lichang
 * @date 20-1-14
 * @time 下午3:29
*/
import (
	"flag"
	"fmt"
)

type Celsius float64  //摄氏度
type Fahrenheit float64	//华氏度

func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9.0/5.0 + 32.0) }
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32.0) * 5.0 / 9.0) }

/**
   @Description: 这里实现Stringer接口，String方法格式化标记的值，并用在命令行帮助消息中
 */
func (c Celsius) String() string { return fmt.Sprintf("%g°C", c) }
func (f Fahrenheit) String() string {return fmt.Sprintf("%g°F", f)}

/**
   @Description: 实现flag.Value接口类型
 */
type celsiusFlag struct{ Celsius }
type fahrenheitFlag struct {Fahrenheit}

/** useless
   @Description: 之所以一定要实现String和Set方法，是因为在接口内部已经做出了规定
	type Value interface {
		String() string
		Set(string) error
	}
 */
//func (f *fahrenheitFlag) Set(s string) error {
//	/**
//	   @Description: 接收命令行参数并做处理
//	   @Param:
//			s: 输入的字符串
//	   @Return:
//	        error ：错误信息
//	 */
//	var unit string
//	var value float64
//	// Sscanf将s中的信息提取出来，存储到value与unit中
//	_, _ = fmt.Sscanf(s, "%f%s", &value, &unit)
//	switch unit{  // 根据单位做出相应的操作
//	case "C", "°C":
//		f.Fahrenheit = CToF(Celsius(value))
//		return nil
//	case "F", "°F":
//		f.Fahrenheit = Fahrenheit(value)
//		return nil
//	}
//	return fmt.Errorf("invalid temperature %q", s)
//}


func (f *celsiusFlag)Set (s string) error {
	var unit string
	var value float64
	_, _ = fmt.Sscanf(s, "%f%s", &value, &unit)
	switch unit{
	case "C", "°C":
		f.Celsius = Celsius(value)
		return nil
	case "F", "°F":
		f.Celsius = FToC(Fahrenheit(value))
		return nil
	}
	return fmt.Errorf("invalid temperature %q", s)
}

func CelsiusFlag(name string, value Celsius, usage string) *Celsius {
	f := celsiusFlag{value}
	flag.CommandLine.Var(&f, name, usage) //Var方法将标记加入应用的命令行标记集合中
	return &f.Celsius
	/**
	   @Description: 返回的Celsius字段是一个会通过Set方法在标记处理的过程中更新的变量
	 */
}

//useless
//func FahrenheitFlag(name string, value Fahrenheit, usage string) *Fahrenheit{
//	f := fahrenheitFlag{value}
//	flag.CommandLine.Var(&f, name, usage)
//	return &f.Fahrenheit
//}