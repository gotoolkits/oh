# oh

![gif](images/oh.gif)

Oh is a Unix shell. If you've used other Unix shells, oh should feel
familiar. The following commands behave as expected:

    echo "Hello, World!"
    cal 01 2030
    date >greeting
    echo "Hello, World!" >>greeting
    wc <greeting
    cat greeting | wc	# Useless use of cat.
    tail -n1 greeting; cal 01 2030
    grep impossible *[a-z]ing &
    wait
    mkdir junk && cd junk
    cd ..
    rm -r greeting junk || echo "rm failed!"

Where oh diverges from traditional Unix shells is in its programming
language features.

At its core, oh is a heavily modified dialect of the Scheme programming
language complete with first-class continuations and proper tail
recursion. Like early Scheme implementations, oh exposes environments
as first-class values. Oh extends environments to allow both public and
private members and uses these extended first-class environments as the
basis for its prototype-based object system.

Written in Go, oh is also a concurrent programming language. It exposes
channels, in addition to pipes, as first-class values. As oh uses the
same syntax for code and data, channels and pipes can, in many cases, be
used interchangeably. This homoiconic nature also allows oh to support
fexprs which, in turn, allow oh to be easily extended. In fact, much of
oh is written in oh.

For more detail see: [Using oh](MANUAL.md)

## Installing

Until Go 1.5 is released, building oh requires building Go from source,
after switching to the master (development) branch. For instructions on
how to do this see:
[Installing Go from source](http://golang.org/doc/install/source)

After building Go, you can install oh by typing:

    go get github.com/michaelmacinnis/oh

Oh compiles and runs, but should be considered experimental, on Windows.

## License

Oh is released under an MIT-style license.

