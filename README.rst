.. vi: set tw=100:

Halalruns
=========


About
-----

Halalruns is a library for interfacing with the speedrun.com_ API. For more information about the
API, check out the official (and outdated) `API documentation`_ on their GitHub page. The aim of
Halalruns is to leverage golangs easy to use concurrency to create a library that is both fast and
easy to use.

.. _speedrun.com: https://www.speedrun.com
.. _API documentation: https://www.github.com/speedruncomorg/api

Documentation
-------------

All functions, methods, structs, etc. provided by the library are documented in the form of UNIX
manual pages. These can be found in the `man/`_ directory in the repository. To install the manual
pages simply run the command ``make docs`` with root permissions. Once this is done you should be
able to view the documentation for any part of the library by running the command ``man <TERM>``
where ``TERM`` is a function, method, struct, etc. For example, to view the documentation for the
``FetchUser`` function you can run the command ``man FetchUser``. In the extremely rare
circumstance that a manual page name conflicts with an existing manual on your system, simply
provide the ``3go`` section to the ``man`` command by running ``man 3go <TERM>``.

If you are on Windows, well that's kind of unfortunate. Just get WSL lol.

.. _man/: man/


Example Usage
-------------

Here is an example program which takes a keyword as a command line argument and prints links to the
first 400 users that match the keyword. This could be useful for finding bots to report in the
`Bots Deletion Thread`_. The number 400 is used simply to show how the user does not need to worry
about manual pagination. If you want to use this as a real application you should probably add in
error checking.

.. code-block:: go

        package main

        import (
                "fmt"
                "os"

                "github.com/speedrunning/halalruns"
        )

        func main() {
                users, err := halalruns.FetchUsers(halalruns.UserFilter{Name: os.Args[1], Max: 400})
                if err == nil {
                        for _, u := range users {
                                fmt.Println(u.Weblink)
                        }
                }
        }

.. _Bots Deletion Thread: https://www.speedrun.com/the_site/thread/7p1bg

License
-------

Halalruns is licensed under the `BSD Zero Clause License`_. In simple terms this means you can do
whatever the hell you want with the code in this repo.

.. _BSD Zero Clause License: LICENSE
