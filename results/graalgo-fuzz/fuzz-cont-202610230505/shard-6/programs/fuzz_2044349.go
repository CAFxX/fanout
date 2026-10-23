package main

import (
    "fmt"
)


func main() {
    v1 := []int{0}
    var v2 uint = 4294967295
    v3 := make(map[int]bool)
    var v4 uint = (^(v2))
    var v5 uint8 = 42
    var v6 int64 = 9223372036854775807
    var v7 uint = (v2 - 373)
    v2 = v7
    v6 -= -1
    var v8 int32 = -637
    var v9 uint32 = 2147483647
    v7 = v2
    var v10 bool = (v2 == 2)
    switch ((("\n\t" + "hello") + "\n\t") + ((("\n\t" + "") + "foo bar") + ("hello" + ("foo bar" + "")))) {
    case "\n\t":
        v7 = 18446744073709551615
    case "a":
        v2 += 78
    }
    v7 = (^(v4))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
