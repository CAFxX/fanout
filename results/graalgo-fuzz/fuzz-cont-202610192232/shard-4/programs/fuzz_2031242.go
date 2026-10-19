package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint32 }
type GenS2 struct { F0 uint }

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[uint32]string)
    var v3 int32 = -2147483648
    var v4 float32 = 1.5
    var v5 float64 = 1.5
    var v6 int64 = -9223372036854775808
    v4 = v4
    v7 := []int{-24, 152, 4294967295}
    var v8 int = -2147483648
    if (!((!(2147483647 <= 42)) || ((0 > (v3 % ((1) | 1))) || true))) {
        v9 := GenS1{F0: (127 == 123), F1: 2147483648}
        _ = v9
    } else {
        fmt.Println(v6, v5)
        fmt.Println(v8, v5, v4)
        var v10 int32 = 2147483647
        _ = v10
    }
    if len(v7) >= 2 {
        v7 = v7[1:len(v7)]
    }
    switch (v8 >> ((473) & 63)) {
    case 9223372036854775807:
        v3 -= 1
    case -602:
        if false {
            var v11 uint64 = 2
            var v12 uint = 18446744073709551615
            _ = v11
            _ = v12
        }
    default:
        var v13 int32 = (v3 / ((0) | 1))
        _ = v13
    }
    switch "hello" {
    case "foo bar":
        v14 := GenS2{F0: 324}
        _ = v14
    default:
        var v15 uint32 = 1
        _ = v15
    }
    var v16 string = "hello"
    v17 := GenS1{F0: true, F1: 42}
    v6 += (v6 - (v6 - v6))
    fmt.Println(v6)
    v3 += 143
    if (2147483648 <= 544) {
        v18 := GenS1{F0: true, F1: 4294967295}
        v19 := GenS2{F0: 880}
        delete(v1, 1)
        _ = v18
        _ = v19
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
