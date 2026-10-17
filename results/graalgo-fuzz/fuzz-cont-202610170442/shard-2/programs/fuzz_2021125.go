package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int64; F1 int16 }
type GenS3 struct { F0 int8; F1 int32 }
func genFn0(p0 int64) uint32 {
    v5 := GenS1{F0: 74}
    v6 := []int{-9223372036854775808, -869, 9223372036854775807}
    v7 := GenS2{F0: p0, F1: -32768}
    v8 := GenS2{F0: -124, F1: -311}
    _ = p0
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    return 2
}

func main() {
    var v1 int64 = 9223372036854775807
    var v2 uint = 657
    v3 := make(map[string]bool)
    var v4 string = (("hello" + "a") + "\n\t")
    v2 += (v2 / (((v2 * v2)) | 1))
    for fi9 := 0; fi9 < 5; fi9++ {
        v1 += v1
        _ = fi9
    }
    var v10 uint8 = 0
    var v11 int16 = -1
    var v12 int32 = 1
    v10 = v10
    v12 += 2147483647
    for fi13 := 0; fi13 < 7; fi13++ {
        v11 = v11
        v14 := GenS3{F0: int8(v2), F1: v12}
        var v15 int16 = 0
        _ = fi13
        _ = v14
        _ = v15
    }
    var v16 float64 = 1e10
    var v17 bool = false
    var v18 int32 = v12
    var v19 uint8 = (v10 / ((v10) | 1))
    v12 = (-(v12))
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
}
