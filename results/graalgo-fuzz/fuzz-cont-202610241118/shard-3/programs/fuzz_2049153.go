package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int; F2 string }
type GenS2 struct { F0 int16; F1 int8; F2 string }

func main() {
    v1 := make(map[string]int)
    var v2 bool = (true && ((255 < 65535) && false))
    v3 := []uint32{99, 220}
    var v4 int = 128
    var v5 float64 = (-2.25 * (1.5 / 3.14159))
    v6 := make(map[int]int)
    for fi7 := 0; fi7 < 10; fi7++ {
        for fi8 := 0; fi8 < 7; fi8++ {
            v5 += 0.1
            _ = fi8
        }
        v4 += 32768
        for fi9 := 0; fi9 < 11; fi9++ {
            var v10 int8 = 38
            var v11 uint = 909
            v10 = (v10 & v10)
            _ = fi9
            _ = v10
            _ = v11
        }
        _ = fi7
    }
    var v12 string = ("a" + "foo bar")
    v4 = (v4 + 9223372036854775807)
    v5 += ((0.1 - float64(0.0)) - 0.1)
    fmt.Println(v12, v12)
    v1[v12] = v4
    var v13 float64 = float64(3.14159)
    v4 -= (v4 >> ((4294967295) & 63))
    for fi14 := 0; fi14 < 8; fi14++ {
        fmt.Println(fi14, v2, v2)
        _ = fi14
    }
    var v15 int16 = 1000
    if false {
        if v2 {
            var v16 int32 = 15
            var v17 int32 = v16
            if ((v16 * (^(v16))) >= -391) {
                var v18 int64 = 9223372036854775807
                _ = v18
            } else {
                v5 += 3.14159
            }
            _ = v16
            _ = v17
        }
        fmt.Println(v15, v4, v5)
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    }
    var v19 uint8 = 88
    for fi20 := 0; fi20 < 13; fi20++ {
        if (v12 <= "foo bar") {
            v13 = (0.0 * v5)
            var v21 float32 = ((((1e10 - 0.1) * (float32(0.0) / (0.1 - ((float32((v5 / (float64(0.0) - (v5 * (v13 - (float64(0.0) / v13)))))) - (float32(0.1) * (((-2.25 / (1e10 + -2.25)) + 1.5) / -2.25))) + 1.5)))) + (0.0 + (-2.25 - (1e10 + 3.14159)))) + 1.5)
            _ = v21
        } else {
            switch "" {
            case "foo bar":
                delete(v6, (fi20 / (((v4 >> ((2) & 63))) | 1)))
            case "":
                for fi22 := 0; fi22 < 2; fi22++ {
                    fmt.Println(fi20)
                    switch -897 {
                    case -948:
                        fmt.Println(v15, v2, v4)
                    case 127:
                        v23 := GenS1{F0: fi20, F1: (fi22 / ((fi20) | 1)), F2: v12}
                        _ = v23
                    case 32768:
                        v24 := GenS2{F0: (v15 | (v15 << ((255) & 63))), F1: 1, F2: v12}
                        _ = v24
                    }
                    if len(v3) > 0 {
                        v3[0] = 907
                    }
                    switch "" {
                    case "foo bar":
                        fmt.Println(fi22, v13)
                    case "a":
                        fmt.Println(v2)
                    default:
                        v13 = 1.5
                    }
                    _ = fi22
                }
            }
        }
        if false {
            v2 = false
        } else {
            v25 := v1["hello"]
            v26 := GenS2{F0: v15, F1: -1, F2: v12}
            _ = v25
            _ = v26
        }
        _ = len(v3)
        _ = fi20
    }
    var v27 int8 = 0
    var v28 int16 = 211
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v12
    _ = v13
    _ = v15
    _ = v19
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
