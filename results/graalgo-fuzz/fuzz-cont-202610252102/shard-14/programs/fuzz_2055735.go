package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint; F1 uint16 }
func genFn0(p0 uint) uint8 {
    p0 -= p0
    v6 := []int{}
    v7 := GenS1{F0: 697}
    _ = p0
    _ = v6
    _ = v7
    return 128
}

func main() {
    var v1 int = 0
    var v2 int32 = 756
    v3 := make(map[uint32]uint32)
    var v4 uint64 = 2
    var v5 int64 = 1
    fmt.Println(v4, v4)
    v8 := GenS1{F0: 255}
    v9 := []int{(v1 % ((-208) | 1)), -274, v1, (v1 * (-(v1)))}
    if ((1.5 / 1.0) < float64(1.5)) {
        if (2 > 2) {
            var v10 string = "hello"
            _ = v10
        }
        v2 += (^(v2))
    } else {
        fmt.Println(v5, v2)
        var v11 uint64 = (v4 % (((^(v4))) | 1))
        _ = v11
    }
    var v12 uint16 = 331
    fmt.Println(v1, v5)
    _ = len(v9)
    if false {
        var v13 float64 = 0.1
        var v14 string = "foo bar"
        _ = v13
        _ = v14
    } else {
        var v15 bool = false
        if (((((!((v1 << ((4294967295) & 63)) <= (v1 ^ 1))) || v15) || v15) || (true || v15)) && v15) {
            if len(v9) >= 2 {
                v9 = v9[1:len(v9)]
            }
        } else {
            var v16 uint16 = (^(v12))
            _ = v16
        }
        switch "foo bar" {
        case "a":
            _ = len(v9)
        case "foo bar":
            v17 := GenS2{F0: 221, F1: 615}
            _ = v17
        case "":
            fmt.Println(v2, v12)
        }
        _ = v15
    }
    _ = len(v3)
    v18 := GenS1{F0: (v12 | v12)}
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v12
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
