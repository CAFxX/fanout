package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8 }
func genFn0(p0 string, p1 int16, p2 int) int32 {
    var v7 int = -1
    var v8 int32 = 751
    var v9 uint32 = 2
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    _ = v9
    return v8
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 uint32 = 317
    var v3 uint16 = 32767
    var v4 uint64 = 42
    v5 := []uint32{v2, 0, v2}
    v6 := make(map[uint32]int)
    var v10 uint32 = v2
    for fi11 := 0; fi11 < 12; fi11++ {
        var v12 float32 = 1.5
        _ = fi11
        _ = v12
    }
    if false {
        for fi13 := 0; fi13 < 13; fi13++ {
            var v14 uint = 186
            v15 := GenS1{F0: -543, F1: 127}
            _ = fi13
            _ = v14
            _ = v15
        }
        v6[(v2 & uint32(v4))] = 4294967295
        v10 -= 317
    } else {
        delete(v6, 1)
        v2 = 65535
        v16 := GenS1{F0: 42, F1: 1}
        _ = v16
    }
    var v17 uint8 = 128
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    var v18 uint = 958
    for fi19 := 0; fi19 < 12; fi19++ {
        if true {
            v2 -= v2
            v17 = (v17 | v17)
        } else {
            v4 = v4
            v20 := GenS1{F0: 1, F1: 0}
            var v21 uint = v18
            _ = v20
            _ = v21
        }
        _ = fi19
    }
    v22 := GenS1{F0: -9223372036854775808, F1: 0}
    switch int(v10) {
    case 128:
        v3 -= v3
    case 32768:
        fmt.Println(v18, v17)
    }
    if len(v5) > 0 {
        v5[0] = v2
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v10
    _ = v17
    _ = v18
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
