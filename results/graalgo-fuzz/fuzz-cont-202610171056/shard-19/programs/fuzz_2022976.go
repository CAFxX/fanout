package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
func genFn0() uint32 {
    var v4 int64 = 1
    switch "" {
    case "\n\t":
        v4 += -711
    case "foo bar":
        v4 -= 547
    case "a":
        v5 := GenS1{F0: 4294967295}
        _ = v5
    default:
        if (false || false) {
            v4 += v4
            v6 := []int{802, -1, 371, 1}
            _ = v6
        }
    }
    v7 := []int{}
    _ = v4
    _ = v7
    return 2
}

func main() {
    v1 := []string{"hello", "", "\n\t", "\n\t"}
    v2 := make(map[uint32]bool)
    var v3 float64 = 1e10
    v3 = v3
    if false {
        var v8 uint32 = 435
        var v9 int8 = -1
        _ = v8
        _ = v9
    }
    var v10 uint16 = 0
    v3 = v3
    for fi11 := 0; fi11 < 7; fi11++ {
        fmt.Println(fi11, fi11)
        _ = fi11
    }
    var v12 uint = 18446744073709551615
    if (((v12 & v12) == (v12 >> (((v12 / (((v12 ^ (v12 >> ((917) & 63)))) | 1))) & 63))) && true) {
        var v13 int = 533
        _ = v13
    } else {
        v14 := GenS1{F0: 0}
        v12 = v12
        for fi15 := 0; fi15 < 6; fi15++ {
            v1 = append(v1, "\n\t")
            var v16 uint32 = 310
            _ = fi15
            _ = v16
        }
        _ = v14
    }
    v12 = (v12 + v12)
    _ = len(v1)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v12
    _ = v2
    _ = v3
}
