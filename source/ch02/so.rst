.. include:: ../LINKS.rst


整个儿的
==============

最终,完成所有功能的配置和代码:

- 版本仓库: https://github.com/ZoomQuiet/urisaok/tree/GAE
- 应用发布: urisago1.appsp0t.com/



app.yaml
---------------

::

    application: urisago1
    version: 3
    runtime: go
    api_version: go1

    handlers:
    - url: /.*
      script: _go_app



urisa.go
---------------

.. literalinclude:: urisa.go
    :language: go
    :linenos:



以上!
---------------

从0基础开始, 以确切的任务为目标, 专注尝试, `42`_ 分钟,在 `Go`_ 简洁而强大的表述帮助,
以及 `GAE`_ 完备的本地调试环境为基础,真心可以完成可用的服务,并实时发布到互联网中吼!!!

接下来?! 当然的继续折腾,享受 `Go`_ ~ 互联网时代的 `C` 语言哪! 
