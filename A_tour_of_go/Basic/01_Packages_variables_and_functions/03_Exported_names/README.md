# Exported names
In Go, a name is exported if it begins with a capital letter. For example, Car is an exported name, as is Pi, which is exported from the math package.

car and pi do not start with a capital letter, so they are not exported.

When importing a package, you can refer only to its exported names. Any "unexported" names are not accessible from outside the package.
