package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int32 }
type GenS2 struct { F0 int64; F1 uint8; F2 uint32 }
type GenS3 struct { F0 uint8 }
func genFn0(p0 uint, p1 int) uint8 {
    var v7 uint32 = 42
    for fi8 := 0; fi8 < 6; fi8++ {
        switch fi8 {
        case 0:
            v9 := []int{-1, fi8, p1}
            _ = v9
        default:
            var v10 int32 = 42
            _ = v10
        }
        var v11 float32 = (1e10 / -2.25)
        _ = fi8
        _ = v11
    }
    p0 -= p0
    v12 := GenS1{F0: "\n\t", F1: 1}
    _ = p0
    _ = p1
    _ = v12
    _ = v7
    return 46
}

func main() {
    var v1 bool = true
    var v2 int8 = 42
    v3 := make(map[uint32]int)
    var v4 int64 = 1
    var v5 uint16 = 65535
    v6 := []int{-843, -2147483648, 127, 32767}
    var v13 string = "hello"
    v14 := GenS2{F0: v4, F1: 127, F2: 539}
    var v15 uint16 = (v5 << ((0) & 63))
    if (-471 >= 1) {
        v1 = (!true)
    }
    var v16 int16 = -32768
    delete(v3, 721)
    switch -9223372036854775808 {
    case 611:
        switch 554 {
        case 835:
            v17 := GenS2{F0: (-(v4)), F1: 69, F2: 934}
            _ = v17
        }
    default:
        v18 := GenS3{F0: 2}
        _ = v18
    }
    for fi19 := 0; fi19 < 8; fi19++ {
        v2 += (v2 & (-(v2)))
        _ = len(v3)
        if false {
            _ = len(v3)
        } else {
            v1 = (0 >= 2147483648)
        }
        _ = fi19
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
