.. include:: ../LINKS.rst


那么
==============

接下来出于 程序员的直觉 应该折腾什么?!

- 提高响应速度!



.. warning:: 

    - 走着...



PS:
----------------
`app.yaml` 修订版本号::

    application: urisago1
    version: 2
    runtime: go
    api_version: go1

    handlers:
    - url: /.*
      script: _go_app



- 再进行 `appcfg.py update ` 时,会自动部署为对应版本
- 当然如 :ref:`fig_1_v` 所示需要在后台配置,才可能将当前版本发布为正式服务
- 否则,可以通过带版本号的应用网址进行访问 ::

    http://2.urisago1.appsp0t.com




.. _fig_1_v:
.. figure:: ../_static/figs/ch1-deploy-versions.png

   插图.1-4 在 `GAE`_ 后台指定默认应用版本



JSON 解析外一篇
--------------------------------

其实呢对于 `JSON`_ 这种2.0时代的明星数据格式解析, `Go`_ 是非常重视的:

- 官方文档中专门章节: `JSON and Go - The Go Programming Language <http://golang.org/doc/articles/json_and_go.html>`_
    
    - 当然,已经有了中译: `JSON和Go - Go语言Wiki <http://golangwiki.org/wiki/index.php?title=JSON%E5%92%8CGo>`_

- 从形式上看,其实 `使用interface{}的通用JSON` 更加简单
- 对于不知道结构,或是不关心结构的情景可以,可以用Unmarshal将其解码成interface{}： 

.. code-block:: go

    b == []byte(`{"Name": "Wednesday", "Age": 6, "Parents": ["Gomez", "Morticia"]}`)
    var f interface{}
    err := json.Unmarshal(b, &f)
    // 此处f中值的Go形式将是一个映射，该映射的键为字符串，值为以空接口值保存的自身
    /*
    map[string]interface{}{
        "Name": "Wednesday",
        "Age": 6,
        "Parents": []interface{}{
            "Gomez",
            "Morticia",
        },
    }
    */
    // 要访问这些数据，可以使用一个类型断言来访问f底层的map[string]interface{}：
    m := f.(map[string]interface{})
    // 然后使用一个range语句来对该映射进行迭代，并使用一个类型switch来以他们的实际类型来访问其值
    for k, v := range m {
        switch vv := v.(type) {
        case string:
            fmt.Println(k, "is string", vv)
        case int:
            fmt.Println(k, "is int", vv)
        case []interface{}:
            fmt.Println(k, "is an array:")
            for i, u := range vv {
                fmt.Println(i, u)
            }
        default:
            fmt.Println(k, "is of a type I don't know how to handle")
        }
    }


或是更加简洁的::

    d := json.NewDecoder (Response.Body)
    data := new(map[string]interface{})
    d.Decode (data)


`Golang-China`_ 列表中 `minux minux.ma # gmail.com` 及时吼回:

- 这样unmarshal确实简单，但是你想过用data的时候么？ `data["success"].(int)` 这么用么？你如何保证success就是int呢？
-  况且，实际上由于javascript的数值类型就只是浮点，所以json根本不区分整数和浮点数，因此那样unmarshal后 `data["success"]` 的类型是float64。

    - http://play.golang.org/p/5aafUAkx0f

- 我一直是不建议unmarshal到interface{}的，同理，也绝对不建议用map[string]interface{}的，这样等于把Go的type-safe完全给丢弃了。对比这个：

    - http://play.golang.org/p/kbh8rYAQI0

- `encoding/json` 包还替你检查了success是不是能放到你所指定的int8类型里去，等于又多了一层保障。


嗯嗯嗯,这样想来,使用 `interface` 解析形式简单,但是使用时,有太多隐患,得不偿失,,,


PPS:
^^^^^^^^^^^^^^^^^^^^^^

顺便说一句，处理用字符串方式返回数字的，只要在struct tag里面这样写就行了：`json:"success,string"`
