package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8; F2 uint32 }
func genFn0() int {
    var v4 bool = false
    v4 = (4294967295 > 2)
    for fi5 := 0; fi5 < 19; fi5++ {
        var v6 bool = v4
        var v7 int8 = -1
        _ = fi5
        _ = v6
        _ = v7
    }
    var v8 int16 = 14
    v9 := []int{32768}
    _ = v4
    _ = v8
    _ = v9
    return -2147483648
}

func main() {
    var v1 int16 = 360
    var v2 uint8 = 42
    var v3 int8 = 127
    v2 += (v2 + 231)
    var v10 uint = 4294967295
    var v11 float64 = 0.1
    v12 := []int{-352}
    var v13 int32 = 2147483647
    v2 += 0
    v2 = (v2 << ((v10) & 63))
    v14 := GenS1{F0: 2, F1: v2, F2: 42}
    var v15 int64 = 499
    var v16 int64 = (v15 / ((173) | 1))
    fmt.Println(v10, v16, v2)
    v13 = (v13 & -2147483648)
    var v17 uint32 = 658
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
}
