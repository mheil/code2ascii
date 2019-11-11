Description
===========
This repository contains some simple command line utilities for conversion of ascii <-> ascii-code.
Reason for this was that I had some text from another application that I can't read, but assumed
that this might be ascii-encoded text, so I needed something to convert the ascii-codes to text.
The first version was written in Java, but I though It would be a nice excercise in learning Go.

code2ascii
----------
Command line utility that converts a list of ascii-codes (in decimal form) to a readable string.

#### Example
``` code2ascii 71,111,32,76,97,110,103 ``` results in ```Go Lang```

ascii2code
----------
Command line utility that converts a string to a list of ascii-codes (in decimal form).

#### Example
``` ascii2code "Go Lang"``` results in ```71,111,32,76,97,110,103```
