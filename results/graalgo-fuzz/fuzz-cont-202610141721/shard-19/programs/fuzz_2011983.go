package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint16 }
type GenS2 struct { F0 uint32; F1 int16 }

func main() {
    var v1 int8 = -1
    var v2 int8 = v1
    v3 := []uint32{491, 0, 2}
    var v4 int16 = 1
    var v5 int8 = (v1 << ((1) & 63))
    var v6 int64 = -283
    var v7 string = "\n\t"
    v6 = v6
    if (2147483648 < 567) {
        v8 := GenS1{F0: v6, F1: 255}
        v7 = (v7 + v7)
        v9 := GenS2{F0: 4294967295, F1: v4}
        _ = v8
        _ = v9
    }
    _ = len(v3)
    v5 = v1
    v10 := GenS2{F0: 2147483648, F1: (v4 ^ (v4 - 0))}
    if false {
        for fi11 := 0; fi11 < 15; fi11++ {
            for fi12 := 0; fi12 < 4; fi12++ {
                _ = len(v3)
                _ = fi12
            }
            for fi13 := 0; fi13 < 17; fi13++ {
                var v14 int32 = 571
                _ = fi13
                _ = v14
            }
            fmt.Println(v7, v2)
            _ = fi11
        }
    }
    var v15 uint64 = 13
    fmt.Println(v2)
    var v16 uint16 = 0
    v7 = (v7 + "hello")
    var v17 bool = true
    if len(v3) > 0 {
        v3[0] = uint32(v16)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
