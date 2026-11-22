package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16; F2 int16 }
type GenS2 struct { F0 uint8; F1 int8; F2 int8 }
type GenS3 struct { F0 int16; F1 int8 }
func genFn0(p0 int64, p1 uint, p2 int16) string {
    v5 := []int{}
    for fi6 := 0; fi6 < 4; fi6++ {
        var v7 float64 = -2.25
        _ = fi6
        _ = v7
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return ""
}
func genFn1(p0 int, p1 uint8, p2 int16) uint8 {
    for fi8 := 0; fi8 < 11; fi8++ {
        p2 -= (^(p2))
        _ = fi8
    }
    v9 := GenS1{F0: p1, F1: 245, F2: p2}
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return (p1 << ((487) & 63))
}

func main() {
    var v1 int8 = 127
    var v2 int8 = (v1 >> ((368) & 63))
    var v3 int16 = -32768
    v4 := make(map[uint32]bool)
    for fi10 := 0; fi10 < 9; fi10++ {
        fmt.Println(v1, v2)
        v1 = v2
        _ = fi10
    }
    v3 = v3
    if false {
        v1 += (^(v1))
    }
    if true {
        v1 = v1
    }
    var v11 int16 = 895
    for fi12 := 0; fi12 < 18; fi12++ {
        v11 = (v11 / ((int16(fi12)) | 1))
        var v13 uint8 = 210
        _ = fi12
        _ = v13
    }
    if false {
        var v14 float32 = (-2.25 + 1.5)
        _ = v14
    }
    v15 := GenS2{F0: 61, F1: (v2 % (((v2 << ((18446744073709551615) & 63))) | 1)), F2: v2}
    var v16 int64 = 0
    v4[2147483647] = true
    var v17 uint32 = 172
    v18 := GenS3{F0: v3, F1: (^(v1))}
    v19 := GenS2{F0: 0, F1: v2, F2: (v1 & (-(v1)))}
    if (v3 < v3) {
        v17 = (v17 - (v17 / ((v17) | 1)))
        v20 := GenS2{F0: 2, F1: v1, F2: (v1 + -89)}
        _ = v20
    } else {
        v1 = (v2 >> ((4294967295) & 63))
    }
    v1 = (v2 % (((v1 / ((v1) | 1))) | 1))
    fmt.Println(v17, v3, v3)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v11
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
}
