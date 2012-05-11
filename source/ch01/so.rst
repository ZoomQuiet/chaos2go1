.. include:: ../LINKS.rst


那么
==============

接下来出于 程序员的直觉 应该折腾什么?!

- 提高响应速度!



.. warning:: 

    - 走着...



ps:

`app.yaml` 修订版本号::

.. code-block:: yaml

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

   插图.1-? 在 `GAE`_ 后台指定默认应用版本

