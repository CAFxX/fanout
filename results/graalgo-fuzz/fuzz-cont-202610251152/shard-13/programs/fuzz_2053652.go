package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint64; F1 int64; F2 int16 }
func genFn0() uint64 {
    for fi7 := 0; fi7 < 15; fi7++ {
        var v8 int = (^(fi7))
        _ = fi7
        _ = v8
    }
    var v9 int = 666
    _ = v9
    return 0
}
func genFn1() int16 {
    var v10 uint32 = 1
    var v11 uint16 = 216
    _ = v10
    _ = v11
    return 1
}

func main() {
    var v1 int8 = 100
    var v2 int8 = (v1 + (-(v1)))
    var v3 int32 = 42
    v4 := make(map[uint32]uint32)
    var v5 float32 = 1e10
    var v6 uint = 255
    fmt.Println(v1, v3, v5)
    v2 += (v2 << ((152) & 63))
    v5 = v5
    var v12 uint32 = 2
    var v13 string = "a"
    for fi14 := 0; fi14 < 14; fi14++ {
        v13 = v13
        if true {
            fmt.Println(v5)
            v5 -= 1e10
        }
        v13 = v13
        _ = fi14
    }
    v15 := GenS1{F0: 42}
    if false {
        v16 := GenS2{F0: 0, F1: 9223372036854775807, F2: 1}
        _ = v16
    }
    v17 := []int{858}
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v15
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
