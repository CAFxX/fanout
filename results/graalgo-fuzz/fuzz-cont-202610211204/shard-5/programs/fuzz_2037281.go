package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16 }

func main() {
    var v1 int = 9223372036854775807
    var v2 int = 191
    var v3 uint = 742
    fmt.Println(v2)
    v4 := GenS1{F0: 808, F1: 1000}
    v5 := GenS1{F0: 1, F1: -1}
    v3 += (v3 << (((^(v3))) & 63))
    v2 = v2
    var v6 float64 = (3.14159 - 1.5)
    v7 := GenS1{F0: 0, F1: 397}
    v8 := GenS1{F0: 32767, F1: 1000}
    if true {
        for fi9 := 0; fi9 < 1; fi9++ {
            if (!(v3 <= 933)) {
                v1 -= (fi9 | (fi9 / ((v2) | 1)))
            }
            v10 := GenS1{F0: 42, F1: 1}
            v3 -= v3
            _ = fi9
            _ = v10
        }
    }
    var v11 int32 = -1
    var v12 uint32 = uint32(v3)
    for fi13 := 0; fi13 < 14; fi13++ {
        v14 := []int{}
        _ = fi13
        _ = v14
    }
    var v15 string = ("" + "hello")
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
