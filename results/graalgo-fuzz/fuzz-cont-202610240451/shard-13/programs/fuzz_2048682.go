package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int32; F1 int64; F2 int8 }
func genFn0(p0 uint32, p1 uint) int16 {
    p1 += (p1 / (((p1 + 0)) | 1))
    p0 += p0
    v7 := []int{-100, -53}
    _ = p0
    _ = p1
    _ = v7
    return 32767
}

func main() {
    var v1 int16 = 0
    v2 := []uint32{2}
    v3 := make(map[int]uint32)
    v4 := make(map[uint32]string)
    var v5 float32 = (-2.25 - 1e10)
    v6 := make(map[uint32]int)
    if false {
        v5 = (v5 + 0.0)
        v8 := GenS1{F0: 65535}
        _ = v8
    }
    switch "\n\t" {
    case "hello":
        fmt.Println(v1)
    case "a":
        v5 = ((float32((1.5 + 1.5)) * float32(float64(-2.25))) / 1.0)
    case "\n\t":
        v5 = v5
    }
    v9 := v3[1]
    v1 -= (^(v1))
    var v10 string = "\n\t"
    v11 := GenS1{F0: 42}
    fmt.Println(v1, v1)
    var v12 int = -861
    v9 = (v9 * v9)
    for fi13 := 0; fi13 < 17; fi13++ {
        switch (v12 - 347) {
        case 2147483647:
            _ = len(v3)
        default:
            v14 := GenS2{F0: -1, F1: 1, F2: 100}
            _ = v14
        }
        var v15 uint16 = 516
        _ = fi13
        _ = v15
    }
    for fi16 := 0; fi16 < 1; fi16++ {
        var v17 uint32 = 65535
        v17 -= (v17 + v17)
        fmt.Println(v1, v10, v10)
        _ = fi16
        _ = v17
    }
    var v18 int8 = 127
    if (v9 >= v9) {
        v19 := GenS2{F0: -1, F1: 1, F2: -128}
        var v20 int16 = -247
        _ = v19
        _ = v20
    } else {
        v21 := GenS2{F0: -1, F1: 66, F2: (v18 & (v18 ^ (v18 & v18)))}
        v2 = append(v2, 65535)
        var v22 int8 = v18
        _ = v21
        _ = v22
    }
    v9 = 65535
    if true {
        v23 := GenS2{F0: -199, F1: 42, F2: (v18 % (((v18 * -95)) | 1))}
        _ = v23
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
