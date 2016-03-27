=====
goose
=====

Synposis
========
goose [options] `[webroot]`_

.. _[webroot]: webroot_

Description
===========
goose is a cli-webserver intended to be a drop-in replacement to :code:`python2 -m SimpleHTTPServer`.
It comes with some nice things like being able to easily specify webroot.
To prevent acceidental exposure it also binds to `127.0.0.1` be default.

webroot
-------
:code:`webroot` defaults to current working directory.

Options
=======
--export, -e    Bind server to 0.0.0.0
--port int, -p  Port to bind (default 8080)
--quiet, -q     Run in quiet mode, ie. no logs
