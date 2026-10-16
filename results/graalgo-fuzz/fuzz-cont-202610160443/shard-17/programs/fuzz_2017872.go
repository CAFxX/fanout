package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int64; F2 uint8 }
type GenS2 struct { F0 uint32; F1 int8 }
func genFn0(p0 uint64, p1 string) uint16 {
    for fi6 := 0; fi6 < 6; fi6++ {
        p0 += p0
        var v7 uint8 = 255
        fmt.Println(p0, v7, p1)
        _ = fi6
        _ = v7
    }
    switch -2147483648 {
    case 1:
        var v8 uint = 642
        _ = v8
    case -9223372036854775808:
        p1 = ""
    case 9223372036854775807:
        var v9 float64 = float64(1e10)
        _ = v9
    }
    p1 = p1
    switch -482 {
    case -1:
        p1 = ((p1 + p1) + p1)
    default:
        fmt.Println(p1, p0)
    }
    p1 = "foo bar"
    _ = p0
    _ = p1
    return 0
}

func main() {
    var v1 uint32 = 697
    var v2 uint = 66
    var v3 uint64 = 18446744073709551615
    v4 := make(map[uint32]bool)
    var v5 int = 189
    v10 := []int{127, v5}
    v4[v1] = false
    var v11 uint32 = v1
    var v12 uint16 = 2
    switch "foo bar" {
    case "a":
        v3 -= (v3 * v3)
    case "":
        if false {
            _ = len(v4)
            _ = len(v4)
            fmt.Println(v3, v2)
        }
    }
    v4[(v1 % ((1) | 1))] = true
    var v13 uint8 = 1
    var v14 uint32 = 42
    for fi15 := 0; fi15 < 10; fi15++ {
        var v16 int8 = 65
        v12 += genFn0(v3, "a")
        v17 := GenS1{F0: 0, F1: -501, F2: v13}
        _ = fi15
        _ = v16
        _ = v17
    }
    v11 = (v11 >> (((v2 >> ((v2) & 63))) & 63))
    v18 := GenS2{F0: v1, F1: -30}
    fmt.Println(v5, v2)
    v14 = (v11 % (((v1 << ((65535) & 63))) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
