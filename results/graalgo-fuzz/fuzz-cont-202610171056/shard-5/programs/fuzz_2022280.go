package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8; F2 string }
func genFn0(p0 int64, p1 uint32, p2 uint16) uint16 {
    if true {
        var v7 int64 = (p0 + (p0 % (((p0 & -9223372036854775808)) | 1)))
        _ = v7
    }
    var v8 uint32 = 2147483647
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return 447
}

func main() {
    var v1 int32 = -712
    var v2 string = "hello"
    v3 := make(map[string]int)
    var v4 uint32 = 8
    var v5 uint32 = (^(v4))
    var v6 float32 = -2.25
    fmt.Println(v5, v4, v2)
    v9 := GenS1{F0: 42, F1: 127, F2: v2}
    fmt.Println(v6, v2, v1)
    v4 = 2147483647
    v1 += (v1 + -1)
    var v10 bool = (100 <= 127)
    v10 = (65535 < 42)
    var v11 int32 = -1
    switch "\n\t" {
    case "a":
        var v12 float64 = 1e10
        _ = v12
    case "foo bar":
        v13 := v3[("foo bar" + (("\n\t" + v2) + v2))]
        _ = v13
    case "":
        var v14 uint32 = 2147483648
        _ = v14
    }
    var v15 int = -1
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
