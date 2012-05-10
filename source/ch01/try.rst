.. include:: ../LINKS.rst


7:42" 突入 
==============

嗯嗯嗯,现在可以开始计时了...


本地运行
---------------------------




URIsAok
-----------------------

对于,我们的目标任务: 包装 `金山网址云安全开放API <http://code.ijinshan.com/api/devmore4.html#md1>`_ 为 `REST`_ 在线服务; 

不用研究透 `Bottle`_ ,仅仅需要作到以下几点,就可以完成核心功能了:

- 对指定 URL 接收 `POST` 方式提交来的待查网址
- 根据文档, 对查询参数项进行合理的 `base64` / `md5` 编码
- 对金山网址云,发起合理请求,并收集结果



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



27:00" 小结
---------------------------

以上这一小堆代码,二十分钟,整出来不难吧? 因为,基本上没有涉及太多 `Bottle`_ 的特殊能力,
几乎全部是标准的本地脚本写法儿,想来:

- 其实,关键功能性行为代码,就8行

    - 仅仅有一行,是需要钻研文档的,,,
    - 即: `eval(urilib.urlopen(api_url).read())`

.. _fig_1_4:
.. figure:: ../_static/figs/chaos1-4-urllib.png

    插图 1-4 访问外网的涉及文档

    - 之前版本文档中,吼关闭了多数对外访问的模块,只能使用 `urllib2`
    - 后来 `SAE`_ 的快速进化,又重新开放了主要的常见几个对外网络访问的库模块 
    - 但是,没有例子,没有推荐链接,真心一句话,是很需要心灵感应才知道怎么作的..

- 其余,都是力气活儿

    - 只要别抄錯
    - 都是赋值,赋值,赋值,赋值,,,,

- 只要注意每一步,随时都可以使用 `print` 吼回来,测试确认无误,就可以继续前进了,,,

`这就是脚本语言的直觉式开发调试体验!`

