package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 bool }
func genFn0(p0 int64, p1 bool, p2 uint32) int64 {
    p1 = (-9223372036854775808 <= (^(p0)))
    switch -401 {
    case -943:
        var v5 bool = (((42 < 0) && true) || (float64(-2.25) != (0.0 / -2.25)))
        _ = v5
    case 128:
        if false {
            var v6 uint = 65535
            var v7 int64 = (p0 / (((p0 - -9223372036854775808)) | 1))
            _ = v6
            _ = v7
        }
    case 9223372036854775807:
        fmt.Println(p0, p0, p0)
    }
    var v8 int = -9223372036854775808
    v9 := GenS1{F0: 626, F1: p1}
    var v10 uint16 = 65535
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v8
    _ = v9
    return 235
}
func genFn1(p0 string, p1 uint8) int16 {
    v11 := []int{655}
    fmt.Println(p1, p1)
    p1 -= 255
    _ = p0
    _ = p1
    _ = v11
    return -1
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 float32 = 0.0
    v3 := make(map[int]int)
    var v4 uint32 = 2
    v3[-9223372036854775808] = 689
    var v12 int8 = 100
    for fi13 := 0; fi13 < 3; fi13++ {
        var v14 uint64 = 816
        v12 = v12
        _ = fi13
        _ = v14
    }
    var v15 int64 = 9223372036854775807
    var v16 int16 = int16(v4)
    fmt.Println(v15, v16, v15)
    v4 += (v4 - 339)
    var v17 int = 32768
    v18 := GenS1{F0: v15, F1: (2147483647 <= 2147483647)}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
}
