package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint16; F2 int8 }
type GenS2 struct { F0 int8; F1 uint32; F2 uint32 }
type GenS3 struct { F0 int16; F1 int8 }
type GenS4 struct { F0 bool; F1 uint32; F2 uint }
type GenS5 struct { F0 int64 }
func genFn0(p0 int) uint32 {
    p0 += 2147483647
    if false {
        fmt.Println(p0)
    }
    var v4 string = "a"
    _ = p0
    _ = v4
    return 2147483648
}

func main() {
    v1 := []string{""}
    var v2 float32 = (0.1 + 0.1)
    var v3 int16 = 0
    _ = len(v1)
    fmt.Println(v2, v2)
    v5 := GenS1{F0: 2147483647, F1: 65535, F2: 87}
    _ = len(v1)
    v6 := GenS1{F0: 736, F1: 1, F2: -1}
    v7 := GenS1{F0: 170, F1: 1, F2: -75}
    v3 += -1
    fmt.Println(v2)
    switch ("a" + "a") {
    case "":
        v2 = v2
    case "\n\t":
        for fi8 := 0; fi8 < 18; fi8++ {
            for fi9 := 0; fi9 < 9; fi9++ {
                var v10 uint16 = 42
                _ = fi9
                _ = v10
            }
            _ = fi8
        }
    }
    v2 -= (((v2 - ((0.1 + ((float32(1.5) - v2) * v2)) + ((float32(1e10) - (v2 - v2)) + v2))) * (v2 - ((v2 / ((3.14159 + -2.25) * 1.5)) / (3.14159 / 1.5)))) / (((((v2 * v2) * (v2 - ((1e10 + float32(float64((float32(3.14159) / 3.14159)))) + v2))) / v2) * v2) * v2))
    var v11 float64 = float64((-2.25 / v2))
    v11 = (float64((3.14159 + v2)) - v11)
    v12 := GenS2{F0: -105, F1: genFn0(199), F2: 924}
    for fi13 := 0; fi13 < 13; fi13++ {
        var v14 uint32 = 4294967295
        var v15 int64 = -1
        if false {
            fmt.Println(v2, v15)
            v16 := GenS3{F0: (v3 - v3), F1: 42}
            v17 := GenS4{F0: false, F1: v14, F2: 730}
            _ = v16
            _ = v17
        }
        _ = fi13
        _ = v14
        _ = v15
    }
    for fi18 := 0; fi18 < 20; fi18++ {
        v2 += 1e10
        for fi19 := 0; fi19 < 1; fi19++ {
            v3 = (v3 - (v3 >> ((33) & 63)))
            fmt.Println(v11)
            _ = fi19
        }
        _ = fi18
    }
    var v20 int16 = (-(v3))
    v21 := GenS2{F0: -128, F1: genFn0(95), F2: 0}
    v22 := GenS5{F0: -1}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v5
    _ = v6
    _ = v7
}
