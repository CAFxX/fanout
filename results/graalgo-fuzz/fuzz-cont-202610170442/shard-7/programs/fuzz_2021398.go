package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
func genFn0(p0 string, p1 uint16) int {
    v6 := GenS1{F0: p1}
    switch p0 {
    case "":
        v7 := []int{128, 128, 1, 1}
        _ = v7
    default:
        var v8 uint = 4294967295
        _ = v8
    }
    _ = p0
    _ = p1
    _ = v6
    return 0
}

func main() {
    v1 := make(map[int]int)
    v2 := []uint32{731, 729, 210}
    var v3 int64 = 0
    var v4 uint16 = 1
    var v5 bool = ((2147483648 != 294) || false)
    for fi9 := 0; fi9 < 7; fi9++ {
        v4 -= (^(v4))
        var v10 uint64 = 2
        _ = fi9
        _ = v10
    }
    if v5 {
        v11 := GenS1{F0: (v4 >> ((255) & 63))}
        _ = v11
    }
    fmt.Println(v4, v3, v3)
    if ("" >= "") {
        switch 32768 {
        case 32768:
            var v12 uint64 = 18446744073709551615
            _ = v12
        }
    }
    v3 = (v3 / ((v3) | 1))
    fmt.Println(v3, v3)
    v13 := GenS1{F0: (^(v4))}
    switch "foo bar" {
    case "foo bar":
        fmt.Println(v5, v3)
    default:
        var v14 int64 = (v3 / ((-916) | 1))
        _ = v14
    }
    var v15 int16 = 31
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
