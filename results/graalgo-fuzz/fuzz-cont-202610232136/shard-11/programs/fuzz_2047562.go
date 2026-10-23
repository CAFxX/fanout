package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int16; F1 uint64; F2 uint }
type GenS3 struct { F0 int; F1 uint32; F2 int }
func genFn0(p0 bool, p1 int16) uint32 {
    var v5 int16 = p1
    switch 1 {
    case -730:
        var v6 uint32 = 2147483647
        _ = v6
    case -615:
        if true {
            var v7 int64 = 9223372036854775807
            v8 := []int{-2147483648, 1, 266}
            _ = v7
            _ = v8
        }
    }
    p0 = (42 < 513)
    v9 := GenS1{F0: 1}
    v10 := []int{4294967295, 225, 1}
    _ = p0
    _ = p1
    _ = v10
    _ = v5
    _ = v9
    return 546
}
func genFn1(p0 int16) bool {
    p0 = p0
    p0 += 0
    var v11 uint8 = 42
    _ = p0
    _ = v11
    return true
}

func main() {
    var v1 uint64 = 477
    v2 := make(map[int]int)
    var v3 int16 = -430
    var v4 uint16 = 42
    v1 -= 1
    v12 := GenS1{F0: -1}
    for fi13 := 0; fi13 < 5; fi13++ {
        v14 := GenS2{F0: int16(v1), F1: v1, F2: 377}
        _ = fi13
        _ = v14
    }
    var v15 string = "foo bar"
    v16 := GenS1{F0: 42}
    var v17 uint32 = 2147483648
    v18 := GenS2{F0: (v3 << ((2) & 63)), F1: (^(v1)), F2: uint(v4)}
    fmt.Println(v3, v1, v1)
    if true {
        v19 := []int{42, 42}
        var v20 int = 1
        var v21 bool = false
        _ = v19
        _ = v20
        _ = v21
    } else {
        if false {
            if true {
                var v22 uint = 1
                _ = v22
            }
        }
    }
    var v23 uint16 = uint16(v17)
    var v24 uint32 = v17
    v25 := GenS3{F0: 0, F1: (^(v17)), F2: 42}
    var v26 bool = (1e10 <= float32(0.0))
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
}
