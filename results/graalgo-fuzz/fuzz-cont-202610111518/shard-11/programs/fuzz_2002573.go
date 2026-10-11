package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32; F2 int16 }
type GenS2 struct { F0 int32; F1 uint64; F2 int32 }
func genFn0(p0 uint16, p1 int32) string {
    fmt.Println(p1, p0)
    v7 := []int{128, -2147483648}
    fmt.Println(p1, p1)
    _ = p0
    _ = p1
    _ = v7
    return "a"
}

func main() {
    var v1 int64 = 9223372036854775807
    var v2 uint = 1
    v3 := make(map[int]uint32)
    v4 := make(map[int]string)
    var v5 float32 = (0.0 + -2.25)
    v6 := make(map[string]uint32)
    v1 = v1
    if (!false) {
        v8 := GenS1{F0: -9223372036854775808, F1: 2147483647, F2: -777}
        _ = v8
    }
    var v9 uint64 = 18446744073709551615
    if (492 <= -9223372036854775808) {
        for fi10 := 0; fi10 < 17; fi10++ {
            v11 := GenS1{F0: (fi10 ^ 127), F1: 273, F2: -1}
            v12 := GenS1{F0: (fi10 % ((fi10) | 1)), F1: 791, F2: 42}
            for fi13 := 0; fi13 < 8; fi13++ {
                v2 = v2
                _ = fi13
            }
            _ = fi10
            _ = v11
            _ = v12
        }
        v5 -= float32(3.14159)
    }
    for fi14 := 0; fi14 < 9; fi14++ {
        var v15 uint8 = 255
        switch "" {
        case "":
            fmt.Println(fi14, v1, v9)
        }
        _ = fi14
        _ = v15
    }
    v16 := []int{}
    if false {
        var v17 float64 = float64((((v5 / 1.0) - v5) - (v5 * (v5 - v5))))
        if false {
            if (v1 != (v1 & (v1 + (v1 / ((9223372036854775807) | 1))))) {
                v17 = ((v17 - float64(v5)) - (3.14159 / v17))
                v2 = 255
                for fi18 := 0; fi18 < 8; fi18++ {
                    v1 += v1
                    _ = fi18
                }
            } else {
                var v19 int32 = -2147483648
                fmt.Println(v2)
                for fi20 := 0; fi20 < 1; fi20++ {
                    fmt.Println(v2)
                    fmt.Println(v1)
                    v5 = (1.5 + 0.0)
                    _ = fi20
                }
                _ = v19
            }
            v21 := GenS2{F0: -2147483648, F1: 615, F2: -79}
            _ = v21
        } else {
            v2 = (v2 & v2)
        }
        for fi22 := 0; fi22 < 16; fi22++ {
            if (2 > 42) {
                var v23 uint8 = 1
                fmt.Println(v9)
                var v24 uint64 = (^(v9))
                _ = v23
                _ = v24
            }
            _ = fi22
        }
        _ = v17
    }
    fmt.Println(v9)
    v16 = append(v16, 251)
    switch -9223372036854775808 {
    case 417:
        var v25 float64 = (((((((((((float64((((v5 - 1.5) + v5) * v5)) / ((((1e10 - 1e10) - 1e10) - 1e10) / 0.1)) * 1e10) * 0.0) + float64(v5)) / (1.5 / (-2.25 * 1e10))) - -2.25) + (float64(v5) + float64(v5))) + (((float64(v5) * 1.5) - (0.1 - (float64(((v5 / v5) + (1.5 - (3.14159 * v5)))) * 0.0))) + 1.5)) * ((((3.14159 / (1e10 - 0.1)) + 0.0) - ((0.1 + ((0.0 / (-2.25 - 1.5)) / (1.5 + float64(1e10)))) / (float64((v5 - ((((((v5 * (0.0 * v5)) * v5) / -2.25) + v5) * 0.0) * v5))) + 1e10))) * float64(-2.25))) / 1.0) * 1e10)
        _ = v25
    default:
        switch "hello" {
        case "foo bar":
            if (("a" + "hello") > "a") {
                var v26 uint16 = 42
                if len(v16) >= 2 {
                    v16 = v16[1:len(v16)]
                }
                _ = v26
            }
        case "\n\t":
            v2 += 2
        default:
            var v27 int64 = v1
            _ = v27
        }
    }
    var v28 float64 = (((-2.25 * (3.14159 / 1.5)) * (0.1 - 1.5)) - (float64(1e10) * 3.14159))
    var v29 bool = (float32(v28) <= ((v5 * (((v5 - 1.5) * ((float32(float64(0.0)) + ((v5 * 3.14159) + v5)) / v5)) + v5)) * (1.5 + v5)))
    v30 := GenS1{F0: -904, F1: 1, F2: 0}
    fmt.Println(v5)
    switch "\n\t" {
    case "hello":
        v31 := GenS2{F0: 0, F1: 2, F2: 24}
        _ = v31
    case "":
        v1 += -263
    }
    fmt.Println(v9, v29)
    var v32 int32 = 42
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v16
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v32
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
