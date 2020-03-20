//存取队列
//提供内存存取队列的功能，也可以作为栈使用。支持存取列表和存取键值表所支持的所有类型（易语言所有基础数据类型 和 存取列表、存取键值表）。支持阻塞和非阻塞式队列操作，阻塞队列时支持超时设置。未初始化时默认为非阻塞队列，且不限制队列最大数量。您可以从搜索引擎查询队列和栈的结构、用法等。本对象线程安全。
package os

import (
	E "github.com/duolabmeng6/goefun/core"
	"github.com/gogf/gf/container/gqueue"
)

type E存取队列 struct {
	queue *gqueue.Queue
}

func New存取队列() *E存取队列 {
	this := new(E存取队列)
	this.E初始化()
	return this
}

//调用格式： 〈逻辑型〉 对象．初始化 （长整数型 最大数量，整数型 超时时间） - E2EE互联网服务器套件2.2.3->存取队列
//英文名称：Init
//初始化当前队列。本命令为初级对象成员命令。
//参数<1>的名称为“最大数量”，类型为“长整数型（int64）”，初始值为“-1”。队列存储数据的最大数量。[-1] 不限制数量(无限制)。如果等于0 或 小于“-1” 则会返回失败。
//参数<2>的名称为“超时时间”，类型为“整数型（int）”，初始值为“0”。弹出队列时的超时时间(单位：毫秒)。在弹出队列超过当前时间设置都没得到数据则直接返回(假) [0] 不设置超时：弹出队列没有新数据时直接返回(假) [-1] 阻塞直到得到有效数据。
//
//操作系统需求： Windows
func (this *E存取队列) E初始化() {
	this.queue = gqueue.New()
}

//调用格式： 〈无返回值〉 对象．清空 （） - E2EE互联网服务器套件2.2.3->存取队列
//英文名称：Clear
//清空当前队列。本命令为初级对象成员命令。
//
//操作系统需求： Windows
func (this *E存取队列) E清空() {
	this.queue.Close()
}

//调用格式： 〈长整数型〉 对象．取数量 （） - E2EE互联网服务器套件2.2.3->存取队列
//英文名称：Size
//获取当前队列的数量。本命令为初级对象成员命令。
//
//操作系统需求： Windows
func (this *E存取队列) E取数量() int64 {
	return E.E到整数(this.queue.Len())
}

//调用格式： 〈逻辑型〉 对象．压入队列 （通用型 要压入的数据） - E2EE互联网服务器套件2.2.3->存取队列
//英文名称：Push
//将数据压入队列。被加入的数据会添加到队列尾部。如果队列已满则无论是否阻塞都将返回假。本命令为初级对象成员命令。
//参数<1>的名称为“要压入的数据”，类型为“通用型（all）”。支持所有易语言基础数据类型和存取列表、存取键值表。如果是不支持的类型则直接返回假。
//
//操作系统需求： Windows
func (this *E存取队列) E压入队列(v interface{}) {
	this.queue.Push(v)
}

//调用格式： 〈逻辑型〉 对象．压入顶部 （通用型 要压入的数据） - E2EE互联网服务器套件2.2.3->存取队列
//英文名称：PushFront
//将数据压入队列顶部。如果队列已满则无论是否阻塞都将返回假。本命令为初级对象成员命令。
//参数<1>的名称为“要压入的数据”，类型为“通用型（all）”。支持所有易语言基础数据类型和存取列表、存取键值表。如果是不支持的类型则直接返回假。
//
//操作系统需求： Windows
func (this *E存取队列) E压入顶部(v interface{}) {
	this.queue.Push(v)
}

//调用格式： 〈逻辑型〉 对象．弹出队列 （通用型变量 要弹出的数据，［整数型变量 处理结果］） - E2EE互联网服务器套件2.2.3->存取队列
//英文名称：Pop
//将数据从队列弹出。弹出的内容为队列中的第一个值。阻塞时超时未得到数据返回假，非阻塞时未得到数据返回假，类型转换失败返回假。数值、文本、时间 和 字节集之间的数据可通过类型自动转换不需要判断类型，如果队列中的数据为存取列表和存取键值表类型的数据，而接收数据的类型不一致时直接返回假。本命令为初级对象成员命令。
//参数<1>的名称为“要弹出的数据”，类型为“通用型（all）”，提供参数数据时只能提供变量。如果队列中数据类型为存取列表、存取键值表时，此处的变量类型不匹配则直接返回假。如果此处变量为易语言基础数据类型，则队列中的数据将自动转换为此处的变量类型的数据。
//参数<2>的名称为“处理结果”，类型为“整数型（int）”，可以被省略，提供参数数据时只能提供变量。指定存储处理结果的整数变量。其内容的含义 [0] 获取数据成功 [-1] 获取数据失败 [-2] 获取数据超时 [-3] 类型转换失败。
//
//操作系统需求： Windows
func (this *E存取队列) E弹出队列() interface{} {
	return this.queue.Pop()
}

//调用格式： 〈逻辑型〉 对象．弹出尾部 （通用型变量 要弹出的数据，［整数型变量 处理结果］） - E2EE互联网服务器套件2.2.3->存取队列
//英文名称：PopLast
//将队列最后的数据弹出。弹出的内容为队列中的最后一个值。阻塞时超时未得到数据返回假，非阻塞时未得到数据返回假，类型转换失败返回假。数值、文本、时间 和 字节集之间的数据可通过类型自动转换不需要判断类型，如果队列中的数据为存取列表和存取键值表类型的数据，而接收数据的类型不一致时直接返回假。本命令为初级对象成员命令。
//参数<1>的名称为“要弹出的数据”，类型为“通用型（all）”，提供参数数据时只能提供变量。如果队列中数据类型为存取列表、存取键值表时，此处的变量类型不匹配则直接返回假。如果此处变量为易语言基础数据类型，则队列中的数据将自动转换为此处的变量类型的数据。
//参数<2>的名称为“处理结果”，类型为“整数型（int）”，可以被省略，提供参数数据时只能提供变量。指定存储处理结果的整数变量。其内容的含义 [0] 获取数据成功 [-1] 获取数据失败 [-2] 获取数据超时 [-3] 类型转换失败。
//
//操作系统需求： Windows
func (this *E存取队列) E弹出尾部() interface{} {
	return this.queue.Pop()
}
