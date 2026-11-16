package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint16; F2 int16 }

func main() {
    v1 := make(map[int]uint32)
    var v2 float32 = 3.14159
    var v3 uint16 = 255
    var v4 bool = true
    v5 := make(map[string]bool)
    v6 := make(map[int]uint32)
    for fi7 := 0; fi7 < 14; fi7++ {
        v4 = false
        v8 := []int{(fi7 / ((fi7) | 1)), (fi7 << ((18446744073709551615) & 63))}
        if (!(255 >= 42)) {
            for fi9 := 0; fi9 < 6; fi9++ {
                if len(v8) > 0 {
                    v8[0] = (fi7 - (fi7 | fi7))
                }
                _ = fi9
            }
            switch "hello" {
            case "hello":
                if (v4 || false) {
                    fmt.Println(fi7, v3)
                    v10 := GenS1{F0: (("" + "") + ""), F1: 42, F2: int16(v3)}
                    for fi11 := 0; fi11 < 3; fi11++ {
                        v4 = v4
                        if len(v8) > 0 {
                            v8[0] = (fi11 / ((fi11) | 1))
                        }
                        _ = fi11
                    }
                    _ = v10
                }
            default:
                switch "hello" {
                case "":
                    var v12 int32 = 739
                    _ = v12
                }
            }
            var v13 uint8 = 42
            _ = v13
        } else {
            var v14 uint16 = (v3 % (((v3 | (v3 << ((65535) & 63)))) | 1))
            _ = v14
        }
        _ = fi7
        _ = v8
    }
    for fi15 := 0; fi15 < 14; fi15++ {
        v16 := []int{}
        v17 := GenS1{F0: ("hello" + ("" + "\n\t")), F1: v3, F2: 0}
        _ = fi15
        _ = v16
        _ = v17
    }
    v2 = 0.0
    v2 -= float32(0.0)
    switch "foo bar" {
    case "hello":
        v18 := GenS1{F0: "", F1: v3, F2: 0}
        _ = v18
    default:
        fmt.Println(v2)
    }
    for fi19 := 0; fi19 < 15; fi19++ {
        if (v4 && v4) {
            var v20 int32 = -2147483648
            _ = v20
        }
        _ = fi19
    }
    switch 184 {
    case 833:
        var v21 int64 = 1
        _ = v21
    case 338:
        var v22 int = 42
        _ = v22
    default:
        var v23 uint64 = 42
        _ = v23
    }
    v2 += (((((v2 * (((((v2 * v2) + v2) * ((float32((0.1 - ((3.14159 / 0.1) * (-2.25 / 3.14159)))) * -2.25) + 0.0)) * 1e10) - ((1.5 * (v2 - float32((float64((v2 - (v2 / ((((1e10 - (v2 * v2)) * v2) * v2) / (v2 * (0.1 + 0.0)))))) / 1e10)))) + ((v2 / v2) * (v2 * -2.25))))) + v2) * ((v2 * 0.1) - ((1.5 - (v2 * ((float32(1e10) - v2) / v2))) * (v2 + -2.25)))) - (v2 + float32(float64(v2)))) - (0.1 + v2))
    v24 := GenS1{F0: "", F1: (v3 | v3), F2: -1}
    if v4 {
        v3 = v3
        fmt.Println(v4, v2, v3)
        var v25 int = 128
        _ = v25
    }
    v26 := GenS1{F0: "", F1: v3, F2: -1}
    var v27 uint = 229
    for fi28 := 0; fi28 < 15; fi28++ {
        v3 -= (v3 & v3)
        v29 := []int{2147483647, fi28}
        v30 := GenS1{F0: "hello", F1: v3, F2: 0}
        _ = fi28
        _ = v29
        _ = v30
    }
    v3 += v3
    var v31 string = ("" + ("" + "a"))
    v32 := []int{32767, 128, 362, 127}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v32:", len(v32), v32)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v24
    _ = v26
    _ = v27
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
