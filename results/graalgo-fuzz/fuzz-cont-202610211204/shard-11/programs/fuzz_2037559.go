package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32; F2 uint64 }
type GenS2 struct { F0 int64 }
func genFn0(p0 uint, p1 int32) uint8 {
    p0 += (^(p0))
    if ((1 == -844) && (-749 < 9223372036854775807)) {
        v6 := GenS1{F0: 9223372036854775807, F1: p1, F2: 944}
        p1 = p1
        _ = v6
    } else {
        var v7 int64 = -1
        _ = v7
    }
    v8 := GenS2{F0: 1}
    var v9 string = "hello"
    _ = p0
    _ = p1
    _ = v8
    _ = v9
    return 152
}

func main() {
    var v1 uint32 = 2147483648
    var v2 uint8 = 221
    v3 := []int{-223}
    var v4 int32 = -2147483648
    v5 := make(map[string]uint32)
    v1 = v1
    fmt.Println(v2, v2)
    var v10 uint = 255
    var v11 uint8 = 2
    if (!false) {
        v4 -= (^(v4))
        switch "hello" {
        case "foo bar":
            v12 := GenS1{F0: 222, F1: (-(v4)), F2: 800}
            _ = v12
        case "\n\t":
            fmt.Println(v10)
        }
        _ = len(v5)
    }
    v13 := GenS2{F0: -1}
    _ = len(v5)
    v11 += (^(v2))
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v10
    _ = v11
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
