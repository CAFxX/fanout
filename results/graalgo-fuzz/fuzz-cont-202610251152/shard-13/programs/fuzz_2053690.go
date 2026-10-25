package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8 }
type GenS2 struct { F0 uint; F1 bool; F2 bool }

func main() {
    var v1 uint16 = 255
    var v2 uint64 = 2
    var v3 uint = 4294967295
    var v4 int64 = -1
    v5 := []int{32768}
    for fi6 := 0; fi6 < 6; fi6++ {
        for fi7 := 0; fi7 < 3; fi7++ {
            v8 := GenS1{F0: (v1 * (v1 ^ v1)), F1: 127}
            _ = fi7
            _ = v8
        }
        var v9 uint32 = uint32(v2)
        _ = fi6
        _ = v9
    }
    fmt.Println(v4, v4)
    v10 := GenS1{F0: v1, F1: 113}
    fmt.Println(v3, v4, v2)
    var v11 int16 = -1
    v11 -= v11
    var v12 uint16 = (v1 % ((v1) | 1))
    for fi13 := 0; fi13 < 19; fi13++ {
        v14 := GenS2{F0: v3, F1: (!false), F2: (!(((!true) && false) || (1 < -408)))}
        _ = fi13
        _ = v14
    }
    fmt.Println(v12, v3)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
