[freeCodeCamp.org tutorial](https://www.youtube.com/watch?v=YS4e4q9oBaU 'freeCodeCamp.org tutorial')

## Some key features

- Strong and statically typed
- Garbage collection
- Built-in concurrency
- Compiles to standalone binaries (Go runtime, libraries, etc)

## Basic Variables

- Can be reassigned, but not redeclared
- Shadowing is allowed
- Must be used (compile-time error if not)

Creating:

- Declared explicity and assigned later
- Declared and assigned explicity
- Declared and assigned implicitly

Naming:

- Pascal or camelCase
- Length of the name should reflect the lifespan of the variable (ex. "i" for loop since it's only used once in a loop)
- Acronyms should be all uppercase (ex. theURL, theHTTPRequest)

Scoping:

| Location                 | Scope                      |
| ------------------------ | -------------------------- |
| Package-level, uppercase | Exported and global        |
| Package-level, lowercase | Scoped to the package only |
| Block scope              | Block                      |

Uninitialized primitive variables are considered "zero values":

- booleans are false
- numeric types are their equivalent to 0

Numeric types:

- Signed integers: int8, int16, int32, int64 ("int" is based on the platform architecture)
- Unsigned integers: uint8 (alias "byte"), uint16, uint32
- Floating point: float32, float64 (also allow exponential and mixed notation)
- Complex numbers: complex64, complex128
- Types must match before attempting calculations on them

Numeric operators:

- Signed/unsigned int: + - \* / % (int division and remainder), & | ^ &^ << >>
- Float: + - \* /
- Complex: + - \* /

Text:

- String: string
  - UTF-8 character(s), using ""
  - Immutable
  - \+ (concatenation)
  - Array syntax gives the uint8 value for that character
- Rune: rune (alias for int32)
  - UTF-32 character(s), using ''

Constant variables:

- Initialized with the "const" keyword
- Must be assigned a value and cannot be reassigned
- Cannot be assigned to something that is determined at runtime (ex. a function that gives output)
- Use all the same operators as their normal variable counterparts

Enumerated constants and expressions using "iota":

- Allows related constants to be created easily
- Starts at 0 in each const block and increments by 1
- Operations that can be determined at compile time are allowed: arithmetic, bitwise operations, bitshifting
