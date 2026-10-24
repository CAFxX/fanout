package main

import (
    "fmt"
)

func genFn0() uint16 {
    var v4 float32 = (3.14159 - -2.25)
    var v5 float64 = -2.25
    v5 = (v5 / ((v5 * v5) / v5))
    _ = v4
    _ = v5
    return 1
}

func main() {
    v1 := []string{(("a" + "\n\t") + ("hello" + (("hello" + ("\n\t" + "")) + ""))), "", "a"}
    var v2 bool = false
    var v3 float32 = ((1e10 + 1.5) * (1.5 - 1e10))
    var v6 bool = (true || (2 > 127))
    fmt.Println(v3)
    var v7 uint16 = 1
    v3 -= v3
    var v8 int8 = 42
    switch ("a" + "\n\t") {
    case "hello":
        v1 = append(v1, "hello")
    }
    var v9 int32 = -870
    var v10 uint8 = 1
    v2 = (!(531 <= 4294967295))
    var v11 uint = 65535
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
