package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int32; F2 int16 }
func genFn0() bool {
    var v4 int = 1
    var v5 uint16 = 32767
    var v6 int = 2147483647
    _ = v4
    _ = v5
    _ = v6
    return true
}

func main() {
    var v1 int64 = 844
    v2 := []uint32{382, 477, 613}
    var v3 float32 = 0.1
    v1 = (v1 & v1)
    v7 := GenS1{F0: -869, F1: -95, F2: 32767}
    var v8 uint64 = 18446744073709551615
    if (2 < 255) {
        for fi9 := 0; fi9 < 15; fi9++ {
            v1 = v1
            var v10 int32 = 864
            var v11 int64 = (v1 >> ((717) & 63))
            _ = fi9
            _ = v10
            _ = v11
        }
        for fi12 := 0; fi12 < 19; fi12++ {
            v13 := GenS1{F0: -1, F1: -2147483648, F2: 42}
            v2 = append(v2, 4294967295)
            v14 := GenS1{F0: -32768, F1: 0, F2: -32768}
            _ = fi12
            _ = v13
            _ = v14
        }
        v3 = 1e10
    }
    var v15 uint64 = (v8 - v8)
    fmt.Println(v3, v3, v8)
    var v16 float64 = (3.14159 / 1e10)
    var v17 int8 = 93
    v8 -= 2
    var v18 float64 = 0.0
    v8 += (v8 * (v8 & (v8 << ((0) & 63))))
    var v19 int32 = -2147483648
    var v20 float32 = v3
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v7
    _ = v8
}
