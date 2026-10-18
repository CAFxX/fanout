package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }

func main() {
    var v1 int64 = -1
    v2 := make(map[string]bool)
    var v3 int64 = (v1 % ((v1) | 1))
    v4 := []uint32{2147483648, 565, 42, 65535}
    var v5 float32 = 0.1
    v5 -= v5
    switch 679 {
    case 127:
        v6 := GenS1{F0: 53}
        _ = v6
    default:
        v5 = v5
    }
    _ = len(v2)
    if (-313 != 2147483647) {
        fmt.Println(v1, v1)
        v4 = append(v4, 1)
    }
    var v7 uint64 = 4294967295
    var v8 int = -9223372036854775808
    delete(v2, ("foo bar" + ("" + "")))
    var v9 uint = 760
    var v10 uint8 = 128
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
    _ = v9
}
