# Basic types
Go's basic types are

_**bool**_

_**string**_

_**int  int8  int16  int32  int64**_

_**uint uint8 uint16 uint32 uint64 uintptr**_

_**byte // alias for uint8**_

_**rune // alias for int32**_
     // represents a Unicode code point

_**float32 float64**_

_**complex64 complex128**_

The example shows variables of several types, and also that variable declarations may be "factored" into blocks, as with import statements.

The int, uint, and uintptr types are usually 32 bits wide on 32-bit systems and 64 bits wide on 64-bit systems. When you need an integer value you should use int unless you have a specific reason to use a sized or unsigned integer type.
