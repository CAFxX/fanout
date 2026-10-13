package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint32 }
type GenS2 struct { F0 uint8; F1 uint8 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint32; F1 uint16; F2 uint8 }
func genFn0(p0 uint32, p1 uint64) int8 {
    if (!((!(1000 == 727)) || (0.1 < 0.0))) {
        fmt.Println(p0, p0)
        p1 += 101
        p0 += 2147483647
    }
    v6 := []int{-39, 128}
    for fi7 := 0; fi7 < 20; fi7++ {
        var v8 int = fi7
        v9 := GenS1{F0: 18446744073709551615, F1: (p0 << ((883) & 63))}
        _ = fi7
        _ = v8
        _ = v9
    }
    _ = p0
    _ = p1
    _ = v6
    return 0
}
func genFn1(p0 uint, p1 uint, p2 uint64) bool {
    var v10 float64 = 1.5
    v11 := []int{}
    var v12 int8 = -128
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    _ = v12
    return (42 == 1)
}

func main() {
    var v1 uint16 = 42
    v2 := []string{}
    v3 := make(map[uint32]string)
    var v4 float64 = 1.5
    var v5 uint64 = 368
    var v13 int8 = 100
    v5 = 622
    fmt.Println(v13, v1)
    v14 := GenS2{F0: 255, F1: 128}
    v3[657] = "\n\t"
    var v15 float32 = 1e10
    v16 := GenS3{F0: false}
    var v17 bool = (!true)
    v17 = false
    fmt.Println(v13)
    v13 -= int8(v1)
    if (0.1 == (float64(((v15 - -2.25) - (((((1e10 / (0.0 - v15)) - v15) - ((v15 / -2.25) * float32(v4))) * (((float32((float64(((1.5 / float32(v4)) * (v15 / 1.5))) - float64(v15))) / -2.25) * (v15 - v15)) - 1e10)) - (((v15 + ((v15 + v15) - 0.1)) / v15) - v15)))) / 0.1)) {
        v4 += 0.0
    }
    v18 := GenS4{F0: 2, F1: uint16(v5), F2: uint8(v5)}
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
