package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int64 }
func genFn0() uint32 {
    v5 := GenS1{F0: 1}
    for fi6 := 0; fi6 < 19; fi6++ {
        for fi7 := 0; fi7 < 19; fi7++ {
            var v8 bool = false
            _ = fi7
            _ = v8
        }
        for fi9 := 0; fi9 < 20; fi9++ {
            var v10 uint64 = 2
            var v11 uint8 = 190
            _ = fi9
            _ = v10
            _ = v11
        }
        var v12 uint16 = 0
        _ = fi6
        _ = v12
    }
    if true {
        v13 := GenS2{F0: 257}
        var v14 uint32 = 0
        v15 := []int{0, 128, -2147483648}
        _ = v13
        _ = v14
        _ = v15
    }
    var v16 int = -1
    _ = v16
    _ = v5
    return 678
}
func genFn1(p0 uint32) uint16 {
    p0 -= 1
    v17 := GenS1{F0: 1}
    _ = p0
    _ = v17
    return 2
}

func main() {
    var v1 uint16 = 0
    v2 := []int{178, -9223372036854775808, -1, -1}
    var v3 string = "\n\t"
    v4 := make(map[uint32]uint32)
    for fi18 := 0; fi18 < 9; fi18++ {
        var v19 float64 = 0.1
        v3 = "\n\t"
        if true {
            for fi20 := 0; fi20 < 3; fi20++ {
                v19 -= v19
                var v21 int8 = 42
                v22 := GenS3{F0: (fi20 % (((fi20 * (fi20 | (fi20 & fi20)))) | 1))}
                _ = fi20
                _ = v21
                _ = v22
            }
            fmt.Println(fi18)
            v23 := GenS4{F0: 1}
            _ = v23
        }
        _ = fi18
        _ = v19
    }
    _ = len(v2)
    v4[886] = 2147483647
    var v24 int32 = 322
    var v25 string = v3
    var v26 int16 = int16(v1)
    v24 += (v24 >> ((771) & 63))
    var v27 int8 = 100
    var v28 uint8 = 236
    var v29 int = -247
    v3 = ("hello" + v25)
    fmt.Println(v3, v25, v29)
    switch "a" {
    case "\n\t":
        v26 = v26
    }
    var v30 float32 = 1.5
    v4[1] = 4294967295
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
}
