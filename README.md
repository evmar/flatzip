flatzip creates a mtime-equivalent mirror of a directory with all
files filled with zeros for easy compression.  It is useful for
distributing a directory tree as seen by a build system for
performance discussions.

Build it with [go](http://golang.org) 1:

    go build

Run it to mirror a directory:

    flatzip \path\to\src outputdir

Then zip the output and distribute.
