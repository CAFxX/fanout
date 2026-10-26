package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int64; F2 string }
type GenS2 struct { F0 uint; F1 uint64; F2 uint }
func genFn0(p0 int64, p1 uint32, p2 int) uint {
    fmt.Println(p1, p2, p2)
    v7 := GenS1{F0: 255, F1: 9223372036854775807, F2: "\n\t"}
    p1 += 65535
    for fi8 := 0; fi8 < 20; fi8++ {
        v9 := []int{p2, p2}
        p1 = (^(p1))
        _ = fi8
        _ = v9
    }
    var v10 uint32 = 42
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v7
    return 4294967295
}

func main() {
    v1 := make(map[int]int)
    var v2 int8 = -128
    var v3 int = 366
    var v4 bool = false
    v5 := []uint32{2147483648, 2, 2}
    var v6 int64 = 1
    v11 := GenS1{F0: 42, F1: v6, F2: "foo bar"}
    var v12 uint16 = 2
    fmt.Println(v12)
    v2 -= (v2 / ((v2) | 1))
    var v13 uint = 995
    var v14 bool = v4
    fmt.Println(v4)
    if len(v5) > 0 {
        v5[0] = 942
    }
    fmt.Println(v14, v13, v6)
    var v15 uint32 = 65535
    var v16 int64 = v6
    switch "hello" {
    case "a":
        var v17 bool = v4
        _ = v17
    case "":
        var v18 uint16 = (v12 + (v12 + v12))
        _ = v18
    case "hello":
        for fi19 := 0; fi19 < 9; fi19++ {
            v20 := GenS2{F0: 4294967295, F1: 4294967295, F2: v13}
            fmt.Println(v15, v13, v3)
            _ = fi19
            _ = v20
        }
    }
    v12 = v12
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
