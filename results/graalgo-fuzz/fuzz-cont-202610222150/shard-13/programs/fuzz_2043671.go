package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 bool; F2 int }

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[int]int)
    v3 := []string{"\n\t"}
    var v4 int64 = -9223372036854775808
    var v5 int8 = 1
    fmt.Println(v5, v5)
    switch ("" + "a") {
    case "":
        if (false && (false || true)) {
            v5 += v5
            var v6 float32 = 1e10
            fmt.Println(v6, v4)
            _ = v6
        }
    default:
        if false {
            var v7 int = 2147483647
            _ = v7
        }
    }
    var v8 int64 = 472
    v5 -= 127
    for fi9 := 0; fi9 < 3; fi9++ {
        if true {
            v5 += (v5 - (v5 - (-(v5))))
            v10 := v2[(-(fi9))]
            _ = v10
        } else {
            fmt.Println(fi9, v5)
            v8 += (v8 & (v4 + 9223372036854775807))
        }
        if false {
            var v11 bool = false
            _ = len(v3)
            _ = v11
        } else {
            _ = len(v3)
            v5 = (v5 ^ 42)
            v5 += 64
        }
        if (0 >= -2147483648) {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            v4 = -831
        } else {
            var v12 uint64 = 660
            var v13 uint = 65535
            _ = len(v2)
            _ = v12
            _ = v13
        }
        _ = fi9
    }
    v14 := GenS1{F0: 2, F1: (42 == -788), F2: -169}
    if true {
        var v15 float64 = 1e10
        _ = v15
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
