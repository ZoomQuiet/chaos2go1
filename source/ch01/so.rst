.. include:: ../LINKS.rst


那么
==============

接下来出于 程序员的直觉 应该折腾什么?!

- 提高响应速度!



.. warning:: 

    - 走着...


`$ /opt/sbin/google_appengine_go/appcfg.py update urisaok/`

::
   
    Application: urisago1; version: 2
    Host: appengine.google.com

    Starting update of app: urisago1, version: 2
    Getting current resource limits.
    2012-05-03 11:10:37,215 WARNING appengine_rpc.py:436 ssl module not found.
    Without the ssl module, the identity of the remote host cannot be verified, and
    connections may NOT be secure. To fix this, please install the ssl module from
    http://pypi.python.org/pypi/ssl .
    To learn more, see http://code.google.com/appengine/kb/general.html#rpcssl . 
    Email: Zoom.Quiet@gmail.com
    Password for Zoom.Quiet@gmail.com: 
    Scanning files on local disk.
    Cloning 5 application files.
    Uploading 4 files and blobs.
    Uploaded 4 files and blobs
    Compilation starting.
    Compilation: 1 files left.
    Compilation completed.
    Starting deployment.
    Checking if deployment succeeded.
    Will check again in 1 seconds.
    Checking if deployment succeeded.
    Will check again in 2 seconds.
    Checking if deployment succeeded.
    Will check again in 4 seconds.
    Checking if deployment succeeded.
    Will check again in 8 seconds.
    Checking if deployment succeeded.
    Will check again in 16 seconds.
    Checking if deployment succeeded.
    Deployment successful.
    Checking if updated app version is serving.
    Completed update of app: urisago1, version: 2



ps:
------------

表忘记及时使用 `saecloud`_  将全新代码部署到 `SAE`_ 中吼! 这样大家才用得上你的新应用呢,,

pps:
--------------

`config.yaml` 修订版本号::

    ---
    name: urisaok
    version: 2
    ...

- 再进行 `saecloud deploy` 时,会自动部署为新版本应用
- 当然, 认为一切 OK 时,也记得要到 `SAE`_ 后台,指定默认版本到新版本吼...

