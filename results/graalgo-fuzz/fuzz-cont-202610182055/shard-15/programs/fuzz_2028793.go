package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16 }

func main() {
    var v1 float64 = (1e10 * 1.5)
    var v2 int = 127
    v3 := make(map[string]int)
    var v4 uint8 = 2
    v5 := []uint32{529, 774, 735}
    if ((v2 | (-(v2))) <= v2) {
        fmt.Println(v2, v1)
        v2 = v2
        v1 -= (v1 - (0.1 + 0.1))
    } else {
        if false {
            fmt.Println(v2, v2)
            fmt.Println(v2, v1, v4)
        } else {
            fmt.Println(v2, v1, v1)
            v4 = (v4 | 128)
            for fi6 := 0; fi6 < 15; fi6++ {
                if (-1 <= -782) {
                    var v7 bool = (!false)
                    _ = v7
                }
                _ = fi6
            }
        }
        for fi8 := 0; fi8 < 19; fi8++ {
            v2 -= (fi8 << ((327) & 63))
            _ = fi8
        }
        v1 = v1
    }
    for fi9 := 0; fi9 < 19; fi9++ {
        var v10 uint32 = 549
        var v11 uint = 1
        for fi12 := 0; fi12 < 20; fi12++ {
            var v13 uint8 = (v4 | 0)
            _ = fi12
            _ = v13
        }
        _ = fi9
        _ = v10
        _ = v11
    }
    v14 := GenS1{F0: true, F1: int16(v2)}
    v15 := GenS1{F0: false, F1: -1}
    var v16 int8 = 42
    v16 = (v16 % (((-(v16))) | 1))
    var v17 uint16 = 414
    delete(v3, "hello")
    var v18 int8 = (^(v16))
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
