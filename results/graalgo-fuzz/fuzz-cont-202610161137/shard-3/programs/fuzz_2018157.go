package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint8 }
func genFn0(p0 uint64) int16 {
    if (-22 <= -973) {
        var v4 uint32 = 2
        v5 := GenS1{F0: "a"}
        v6 := GenS2{F0: 127}
        _ = v4
        _ = v5
        _ = v6
    }
    p0 += (p0 | 748)
    p0 = (^(p0))
    p0 = p0
    _ = p0
    return 0
}

func main() {
    v1 := []int{}
    var v2 int32 = 1
    v3 := []uint32{42, 1, 65535, 324}
    v1 = append(v1, -2147483648)
    var v7 int8 = 127
    v7 += (-(v7))
    fmt.Println(v7, v2)
    if (true && ((!true) || (42 == 65535))) {
        if (9223372036854775807 > 42) {
            fmt.Println(v2, v2)
            v2 -= v2
        }
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    }
    v2 += (^(v2))
    fmt.Println(v7, v7)
    switch "\n\t" {
    case "":
        v2 = v2
    case "\n\t":
        var v8 uint32 = 42
        _ = v8
    case "foo bar":
        v1 = append(v1, -591)
    default:
        v2 = -1
    }
    fmt.Println(v2)
    var v9 bool = false
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v7
    _ = v9
}
