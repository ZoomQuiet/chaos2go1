.. include:: ../LINKS.rst


整备
==============

如果真心没有学习过 `Go`_ ,抄过了 `Hollo World` 也许没法儿继续了,,,

- 不过,其实, `Go`_ 真心好学习,最多两小时,就可以完成各种常见任务了,,,
- 但是...


前题!
-----------

有前题的吼!

- 至少有过 C/C++/C# 什么的任何一种编译型语言的体验
- 又或是,至少有任何一种开发语言的经验,明白计算机程序的基本元素
- 这样才好快速复用起以往的经验


.. seealso:: 教程推荐
    
    - 三天包会 `Go`_ 由创始人 Rob Pike 主持的课程:

        - `GoCourseDay1.pdf <http://go.googlecode.com/hg-history/release-branch.r60/doc/GoCourseDay1.pdf>`_
        - `GoCourseDay2.pdf <http://go.googlecode.com/hg-history/release-branch.r60/doc/GoCourseDay2.pdf>`_
        - `GoCourseDay3.pdf <http://go.googlecode.com/hg-history/release-branch.r60/doc/GoCourseDay3.pdf>`_

    - `Go`_ 中国达人 `Fango 的精彩翻译 <http://code.google.com/p/ac-me/downloads/list>`_

        - `胡文 Go.ogle <http://ac-me.googlecode.com/files/go.ogle.pdf>`_ ~ 最快乐的 `Go`_ 体验小说
        - `Go导读/效率手册/规范 三合一 <http://ac-me.googlecode.com/files/fango.pdf>`_


参考: `Go语言教程 - Go语言Wiki <http://golangwiki.org/wiki/index.php?title=Go%E8%AF%AD%E8%A8%80%E6%95%99%E7%A8%8B>`_



Go 精粹
--------------------

.. code-block:: cpp

    /*
        多行
        注释
    */

    package main    // 每个文件,必须声明为包

    import (        // 统一在头部邮件 各种包的加载
        "os"
        "flag"      // 单行注释
    )

    var omitNewline = flag.Bool("n", false, "don't print final newline")

    /*  变量的声明,以下都是合法的 ;-)
        var s string = ""
        var s = ""
        s := ""

        常量可以是:
        var a uint64 = 0  // a的类型为uint64，值为0
        a := uint64(0)    // 与以上相同；使用了一次“转换”
        i := 0x1234       // i获得了默认的类型：int
        var j int = 1e6   // 合法的 - 1000000是一个int
        x := 1.5          // float64类型，这是浮点常量的默认类型
        i3div2 := 3/2     // 整数除法 - 结果是1
        f3div2 := 3./2.   // 浮点数除法 - 结果是1.5
    */

    const (         // 容器 ;-)
        Space   = " "
        Newline = "\n"
    )

    func main() {   // { 必须跟在行尾,单起一行,将编译不过!-)
        flag.Parse() // Scans the arg list and sets up flags
        var s string = ""
        for i := 0; i < flag.NArg(); i++ {  
        // for 是唯一的循环形式,也是唯一可能出现 ; 的语句
            if i > 0 {
                s += Space
            }
            s += flag.Arg(i)
        }
        if !*omitNewline {
            s += Newline
        }
        os.Stdout.WriteString(s)
    }


.. sidebar:: 搜索
    :subtitle: go 是常用词,引擎基本不索引的

    其实,有专用的定制搜索,
    用来搜索 `Go`_ 相关的各种文件/代码/邮件
    
    `A Search Engine for Go Information <http://go-lang.cat-v.org/go-search>`_




另外, 参考 QCon2012北京的講演: `go，互联网时代的c语言 许式伟 <http://www.slideshare.net/Zoom.Quiet/01s0401-goc>`_ 可以了解各种 `Go`_ 的核心特性;

就笔者的体验,最爽直的有一点就是 `Go`_ 的形式非常人性!

- 参考: `螺旋形（C/C++）和顺序（Go）的声明语法 « Yi Wang's Tech Notes <http://floss.zoomquiet.org/data/20110622145115/index.html>`_
- C/C++ 的代码形式,进行各种声明时,词的顺序和意义是大幅度扭曲的:

::

                     +--------------------+
                     | +---+              |
                     | |+-+|              |
                     | |^ ||              |
                char *(*fp)( int, float *);
                 ^   ^ ^  ||              |
                 |   | +--+|              |
                 |   +-----+              |
                 +------------------------+

读作:
    - fp是一个指针，
    - 指向一个函数（螺旋路径被fp右边的括号封死，绕到左边的`*`）
    - 有一个整形和一个浮点指针参数（两个参数一起读因为他们被一个括号括起来了）
    - 并返回一个指针，指向一个字符

`何其蛋痛!!!` ~ 这才回想起当年,为毛使用 `Turbo C++` 时,永远的挫败感了...

对等的 `Go`_ 声明就有爱的多...

::

    f func(func(int,int) int, int) func(int, int) int


读作:
    - f是一个函数
    - 他的参数包括一个函数(有两个整数参数并返回一个整数)，和一个整数，
    - 并且返回一个函数，他有两个整数参数并返回一个整数  

这样,思想和书写统一,少了很多转化,舒服很多 ;-)
