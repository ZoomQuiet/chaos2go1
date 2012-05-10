.. include:: ../LINKS.rst


27:42" 重构
============================


gae datastore get , return value is wrong.. - golang-nuts | Google 网上论坛
http://groups.google.com/group/golang-nuts/browse_thread/thread/996aa7e83aa29fde/f73e109d94735a00?show_docid=f73e109d94735a00



**BINGO!**



42:01" 小结
---------------------------------

~ 这一处增强,纯粹是根据文档配合后台日志,尝试几个回和而已,一刻钟,整出来不难吧?

- 但是,过程中的心理冲突,绝对不轻
- 比如,文档中未言明的各种细节, 是否重要? 怎么测试确认?
- 怎么设计 `print` 点输出的格式,以便从后台日志中明确的识别出?
- 等等,都需要补课,老实查阅文档,认真领悟,大胆尝试,建立靠谱的思路和反应,,,

不过,整体上,只要思路明确,方向正确,真心只是个轻松的过程而已,,,


.. note:: `KVDB`_ 的 `key`

    - 因为 `KVDB`_ 按文档吼,是对 `memcached`_ 接口的精简仿制;
    - 所以,根据 `Very long URL aliases not correctly cached in memcache <http://2bits.com/articles/very-long-url-aliases-not-correctly-cached-memcache.html>`_  等相关文章的分享,如果使用原样儿 URL 作 `key` 很可能出问题...
    - 笔者就曾经通过后台日志确认,只要使用正常的 URL 作 `key` 是保存不进 `KVDB`_ 的,
    - 于是使用 `urlsafe_b64encode(uri)` 进行处理就好...
    - 可是,没有想到 `SAE`_ 日新月异的发展中,现在再试,居然,平静的接受了! `叫声好!`


