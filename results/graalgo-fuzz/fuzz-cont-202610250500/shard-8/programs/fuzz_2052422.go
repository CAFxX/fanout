package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8 }
type GenS2 struct { F0 uint32; F1 int; F2 uint16 }

func main() {
    var v1 int = 127
    var v2 int16 = -1
    var v3 int8 = 0
    var v4 bool = false
    var v5 int64 = -1
    v6 := []int{v1, v1, (-(v1)), v1}
    v3 -= 100
    var v7 string = "hello"
    for fi8 := 0; fi8 < 5; fi8++ {
        var v9 uint64 = 42
        _ = fi8
        _ = v9
    }
    v3 -= 0
    fmt.Println(v4)
    if (114 > 162) {
        v3 += 100
        if v4 {
            v10 := GenS1{F0: (v5 * (v5 - int64(v3))), F1: v3}
            v11 := GenS1{F0: (v5 * v5), F1: (-(v3))}
            v12 := GenS1{F0: v5, F1: v3}
            _ = v10
            _ = v11
            _ = v12
        }
    } else {
        for fi13 := 0; fi13 < 19; fi13++ {
            if len(v6) > 0 {
                v6[0] = (-(v1))
            }
            var v14 int8 = (v3 + v3)
            _ = fi13
            _ = v14
        }
    }
    _ = len(v6)
    v15 := GenS2{F0: 782, F1: (v1 * v1), F2: 0}
    fmt.Println(v5)
    var v16 float64 = (0.1 - (1e10 - 1e10))
    for fi17 := 0; fi17 < 8; fi17++ {
        var v18 string = "hello"
        fmt.Println(v7, v3)
        var v19 uint8 = 127
        _ = fi17
        _ = v18
        _ = v19
    }
    var v20 bool = v4
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
