.. include:: ../LINKS.rst


7:42" 突入 
==============

嗯嗯嗯,现在可以开始计时了...



.. note:: (~_~)

    - 新手就是新手,没有帮助是会崩溃的,所以,,,
    - 一定要订阅相关的开发列表( `maillist`_ )
    - 官方的忒活跃了,而且都是E文, 如 :ref:`fig_1_0`
    - 所以,推荐中国弟兄们簇集的 `Golang-China <http://groups.google.com/group/golang-china>`_


.. _fig_1_0:
.. figure:: ../_static/figs/120427-golang-nuts.png

   插图.1-0 `golang-nuts <http://groups.google.com/group/golang-nuts>`_ 官方主力列表






URIsAok
-----------------------

对于,我们的目标任务: 包装 `金山网址云安全开放API <http://code.ijinshan.com/api/devmore4.html#md1>`_ 为 `REST`_ 在线服务; 

不用研究透 `Go`_ ,仅仅需要作到以下几点,就可以完成核心功能了:

- 对指定 URL 接收 `POST` 方式提交来的待查网址
- 根据文档, 对查询参数项进行合理的 `base64` / `md5` 编码
- 向 `金山网址云API`_ 发起合理请求,并解析返回的 `JSON`_ 格式数据



获得POST数据
^^^^^^^^^^^^^^^^^^^^^^^
有的抄 ;-)

- 参考: `Handling Forms - Google App Engine — Google Developers <https://developers.google.com/appengine/docs/go/gettingstarted/handlingforms>`_


::

    ...
    func init() {
        http.HandleFunc("/", root)
        http.HandleFunc("/sign", sign)
    }

    ...
    func sign(w http.ResponseWriter, r *http.Request) {
        err := signTemplate.Execute(w, r.FormValue("content"))


就可以得出:


.. code-block:: go

    package urisa

    import (
        "fmt"
        "net/http"
        
        "appengine"
    )

    func init() {
        http.HandleFunc("/", help)
        http.HandleFunc("/chk", chk)
    }

    func help(w http.ResponseWriter, r *http.Request) {
        fmt.Fprint(w, usageHelp)
    }
    const usageHelp = `
    URIsA ~ KSC 4 GAE powdered by go1
    {v12.05.3}
    usage:
        $ curl -d "uri=http://sina.com" urisago1.appsp0t.com/chk
    `
    func chk(w http.ResponseWriter, r *http.Request) {
        c := appengine.NewContext(r)
        url := r.FormValue("uri")
        c.Infof("url~\t %v\n", url)
    }


如 :ref:`fig_1_1` 所示:

.. _fig_1_1:
.. figure:: ../_static/figs/ch1-1-post.png

   插图.1-1 使用 `cURL`_ 测试POST数据捕获

完全吻合预期, 增补了两招儿:

- 使用 `http.HandleFunc("/chk", chk)` 声明了守候的路由
- 使用 `r.FormValue("uri")` 直接抓到 POST 来的数据



时间戳
^^^^^^^^^^^^^^^^^^^^^^^

根据 `金山网址云API`_ 文档, 需要形如 `1336730157.204` 的 `Unix时间戳(Unix timestamp) <http://tool.chinaz.com/Tools/unixtime.aspx>`_

- 官方包,当然有的了: `func Unix(sec int64, nsec int64) Time <http://golang.org/pkg/time/#Unix>`_
- 问题是返回的是整数吼,没有小数点后3位!
- 再细看的,有 `func (t Time) UnixNano() int64 <http://golang.org/pkg/time/#Time.UnixNano>`_

可是怎么使用?

- 返回的 `int64` 长整数,如何可以格式化为有3个小数的浮点数? 或是字串?
- 吼了列表没有获得及时指点,就使用以往经验:

    - 硬来!
    - 总是可以将整数变成字串
    - 然后使用字串格式化的方式输出想要的形式

果然!

- 有 `strconv <http://golang.org/pkg/strconv/>`_ 字串整形包!
- 进一步的,定位到 `func FormatInt(i int64, base int) string <http://golang.org/pkg/strconv/#FormatInt>`_

    - 囧的是: `FormatUint returns the string representation of i in the given base`
    - 这个 `base` 神马意思?!

::

    now := time.Now()
    c.Infof("%v , %v", now.Unix(), now.UnixNano())
    nano := strconv.FormatInt(now.UnixNano(),2)
    c.Infof("timestamp ~ %v", nano)


先如此尝试一下,,,

- 可以运行,输出: `timestamp ~ 1001010001101000001011110011100000100110011111110100011000000`
- `FormatInt(now.UnixNano(),3)` ,输出: `timestamp ~ 22220110102022002202120220120110211221`
- `FormatInt(now.UnixNano(),4)` ,输出: `timestamp ~ 112150136724241452010`
- 嗯嗯嗯?! 有点儿感觉了,使用 `FormatInt(now.UnixNano(),16)` ?!

    - 果然输出: `timestamp ~ 128d05f359d03738`
    - 原来! 这 `base` 就是制式的意思
    - `FormatInt()` 能将整数,处理成指定制式的数字! 从二进制到16进制!

比照 `Python`_ 中的时间戮::

    $ ipython
    Python 2.7.1 (r271:86832, Jul 31 2011, 19:30:53) 
    Type "copyright", "credits" or "license" for more information.

    In [1]: import time

    In [2]: "%.3f"% time.time()
    Out[2]: '1336731591.484'


完成 `Go`_ 的::

    now := time.Now()
    nano := strconv.FormatInt(now.UnixNano(),10)
    c.Infof("timestamp ~ %v %v.%v", nano, nano[0:10],nano[10:13])




MD5
^^^^^^^^^^^^^^^^^^^^^^^


base64
^^^^^^^^^^^^^^^^^^^^^^^



外网请求(urlfetch)
^^^^^^^^^^^^^^^^^^^^^^^

.. _URL Fetch Go API Overview - Google App Engine — Google Developers: https://developers.google.com/appengine/docs/go/urlfetch/overview

::

    import (
        "fmt"
        "net/http"

        "appengine"
        "appengine/urlfetch"
    )

    func handler(w http.ResponseWriter, r *http.Request) {
        c := appengine.NewContext(r)
        client := urlfetch.Client(c)
        resp, err := client.Get("http://www.google.com/")
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        fmt.Fprintf(w, "HTTP GET returned status %v", resp.Status)
    }


JSON 解析
^^^^^^^^^^^^^^^^^^^^^^^


type KSC struct {
        success int  
        msg     string 
}

type KSC struct {
        Success int    `json:"success"`
        Msg     string `json:"msg"`
}


友情提醒，不论是 encoding 还是 decoding struct 的 field
都必须首字母大写：

否则就会无论如何取不到值……

json - The Go Programming Language
http://golang.org/pkg/encoding/json/
The empty values are false, 0, any nil pointer or interface value, and any array, slice, map, or string of length zero. The object's default key string is the struct field name but can be specified in the struct field's tag value. The "json" key in struct field's tag value is the key name, followed by an optional comma and options. Examples:

// Field is ignored by this package.
Field int `json:"-"`

// Field appears in JSON as key "myName".
Field int `json:"myName"`

// Field appears in JSON as key "myName" and
// the field is omitted from the object if its value is empty,
// as defined above.
Field int `json:"myName,omitempty"`

// Field appears in JSON as key "Field" (the default), but
// the field is skipped if empty.
// Note the leading comma.
Field int `json:",omitempty"`

The "string" option signals that a field is stored as JSON inside a JSON-encoded string. This extra level of encoding is sometimes used when communicating with JavaScript programs:

Int64String int64 `json:",string"`

The key name will be used if it's a non-empty string consisting of only Unicode letters, digits, dollar signs, percent signs, hyphens, underscores and slashes.

Map values encode as JSON objects. The map's key type must be string; the object keys are used directly as map keys.

Pointer values encode as the value pointed to. A nil pointer encodes as the null JSON object.

Interface values encode as the value contained in the interface. A nil interface value encodes as the null JSON object.

Channel, complex, and function values cannot be encoded in JSON. Attempting to encode such a value causes Marshal to return an InvalidTypeError.

JSON cannot represent cyclic data structures and Marshal does not handle them. Passing cyclic structures to Marshal will result in an infinite recursion. 


200->400 ?!@#@$!%$#%^5467
------------------------------------------------

::

    2012/05/11 03:16:00 INFO: HTTP GET returned status 200 OK
    2012/05/11 03:16:00 INFO: resp.ContentLength 173
    2012/05/11 03:16:00 INFO: resp.Body <html>
    <head><title>400 Bad Request</title></head>
    <body bgcolor="white">
    <center><h1>400 Bad Request</h1></center>
    <hr><center>nginx/1.0.11</center>
    </body>
    </html>



tcpdump
-----------------
::

    $ tcpdump -ien0 -n -w fixed-get.log  host open.pc120.com
    tcpdump: listening on en0, link-type EN10MB (Ethernet), capture size 65535 bytes
    ^C
    10 packets captured
    100 packets received by filter
    0 packets dropped by kernel


::

    0000  00 0f e2 d3 c2 78 3c 07  54 2c 04 85 08 00 45 00   .....x<. T,....E.
    0010  02 01 16 11 40 00 40 06  00 00 0a 14 e5 17 77 93   ....@.@. ......w.
    0020  92 4d d6 c0 00 50 7e 8a  8a f7 31 fd 5d dc 50 18   .M...P~. ..1.].P.
    0030  ff ff fa ff 00 00 47 45  54 20 2f 70 68 69 73 68   ......GE T /phish
    0040  2f 3f 61 70 70 6b 65 79  3d 6b 2d 36 30 36 36 36   /?appkey =k-60666
    0050  26 71 3d 61 48 52 30 63  44 6f 76 4c 32 52 6b 4d   &q=aHR0c DovL2RkM
    0060  79 35 7a 65 6e 52 73 5a  69 35 75 5a 58 51 36 4f   y5zenRsZ i5uZXQ6O
    0070  44 41 77 4c 32 59 78 4d  44 41 77 4c 77 3d 3d 00   DAwL2YxM DAwLw==.
    0080  00 00 00 00 00 00 00 00  00 00 00 00 00 00 00 00   ........ ........
    0090  00 00 00 00 00 00 00 00  00 00 00 00 00 00 00 00   ........ ........
    00a0  00 00 00 00 00 00 00 00  00 00 00 00 00 00 00 00   ........ ........
    00b0  00 00 00 00 00 00 00 00  00 00 00 00 00 00 00 00   ........ ........
    00c0  00 00 00 00 00 00 00 00  00 00 00 00 00 00 00 00   ........ ........
    00d0  00 00 00 00 00 00 00 00  00 00 00 00 00 00 00 00   ........ ........
    00e0  00 00 00 00 00 00 00 00  00 00 00 00 00 00 00 00   ........ ........
    00f0  00 00 00 00 00 00 00 00  00 00 00 00 00 00 00 00   ........ ........
    0100  00 00 00 00 00 00 00 00  00 00 00 00 00 00 00 00   ........ ........
    0110  00 00 00 00 00 00 00 00  00 00 00 00 00 00 00 00   ........ ........
    0120  00 00 00 00 00 00 00 00  00 00 00 00 00 00 00 00   ........ ........
    0130  00 00 00 00 00 00 00 00  00 00 00 00 00 00 00 00   ........ ........
    0140  00 00 00 00 00 00 00 00  00 00 00 00 00 00 00 00   ........ ........
    0150  00 00 00 26 74 69 6d 65  73 74 61 6d 70 3d 31 33   ...&time stamp=13
    0160  33 36 37 30 35 34 39 38  2e 32 30 31 26 73 69 67   36705498 .201&sig
    0170  6e 3d 34 65 65 32 36 66  66 30 39 66 32 33 62 39   n=4ee26f f09f23b9
    0180  33 62 32 39 64 65 66 65  39 64 36 37 31 66 35 37   3b29defe 9d671f57
    0190  61 30 20 48 54 54 50 2f  31 2e 31 0d 0a 48 6f 73   a0 HTTP/ 1.1..Hos
    01a0  74 3a 20 6f 70 65 6e 2e  70 63 31 32 30 2e 63 6f   t: open. pc120.co
    01b0  6d 0d 0a 41 63 63 65 70  74 2d 45 6e 63 6f 64 69   m..Accep t-Encodi
    01c0  6e 67 3a 20 67 7a 69 70  0d 0a 55 73 65 72 2d 41   ng: gzip ..User-A
    01d0  67 65 6e 74 3a 20 41 70  70 45 6e 67 69 6e 65 2d   gent: Ap pEngine-
    01e0  47 6f 6f 67 6c 65 3b 20  28 2b 68 74 74 70 3a 2f   Google;  (+http:/
    01f0  2f 63 6f 64 65 2e 67 6f  6f 67 6c 65 2e 63 6f 6d   /code.go ogle.com
    0200  2f 61 70 70 65 6e 67 69  6e 65 29 0d 0a 0d 0a      /appengi ne).... 



27:00" 小结
---------------------------

以上这一小堆代码,二十分钟,整出来不难吧? 因为,基本上没有涉及太多 `Go`_ 的特殊能力,
几乎全部是标准的本地脚本写法儿,想来:

- 其实,关键功能性行为代码,就8行

    - 仅仅有一行,是需要钻研文档的,,,

    - 之前版本文档中,吼关闭了多数对外访问的模块,只能使用 `urllib2`
    - 但是,没有例子,没有推荐链接,真心一句话,是很需要心灵感应才知道怎么作的..

- 其余,都是力气活儿,而且基本是相同的流程:

    1. 建立容器,分配合理内存
    1. 处理数据
    1. 输出到对应容器

- 其实也都是赋值,赋值,赋值,赋值,,,,
- 只要注意每一步,随时都可以使用 `print` 吼回来,测试确认无误,就可以继续前进了,,,

`这就是脚本语言的直觉式开发调试体验!`

