package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int8; F1 int32; F2 uint64 }
func genFn0(p0 int8, p1 uint32) uint16 {
    var v6 bool = (0 != 42)
    v7 := []int{-1, 9223372036854775807}
    p0 = p0
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 255
}

func main() {
    var v1 uint8 = 37
    v2 := []int{}
    var v3 float64 = float64(1.5)
    v4 := []int{9223372036854775807, 128}
    v5 := []int{159, 9223372036854775807, 128, 127}
    if len(v4) > 0 {
        v4[0] = -2147483648
    }
    if (15 > 334) {
        for fi8 := 0; fi8 < 2; fi8++ {
            for fi9 := 0; fi9 < 3; fi9++ {
                v3 -= v3
                v3 -= v3
                switch fi8 {
                case 4294967295:
                    var v10 uint64 = 524
                    _ = v10
                case 9223372036854775807:
                    v1 -= 127
                case -632:
                    var v11 float32 = (0.0 / (float32((v3 - v3)) - 0.1))
                    _ = v11
                default:
                    fmt.Println(fi9, fi9)
                }
                _ = fi9
            }
            v5 = append(v5, fi8)
            _ = fi8
        }
        _ = len(v2)
    }
    var v12 uint64 = 42
    switch 42 {
    case 303:
        var v13 int64 = -1
        _ = v13
    case -9223372036854775808:
        switch 42 {
        case 128:
            fmt.Println(v3, v1, v12)
        }
    case -961:
        for fi14 := 0; fi14 < 12; fi14++ {
            for fi15 := 0; fi15 < 5; fi15++ {
                if ("hello" == "\n\t") {
                    var v16 uint32 = 2147483648
                    var v17 int = (fi14 ^ (fi14 ^ (^(fi15))))
                    switch "\n\t" {
                    case "a":
                        var v18 uint32 = (v16 | (v16 - (v16 << ((580) & 63))))
                        _ = v18
                    case "hello":
                        v19 := GenS1{F0: ((("hello" + ("a" + "\n\t")) + "\n\t") + "foo bar")}
                        _ = v19
                    case "":
                        v20 := GenS2{F0: -108, F1: 42, F2: (v12 & (^(v12)))}
                        _ = v20
                    default:
                        v17 -= (^(v17))
                    }
                    _ = v16
                    _ = v17
                }
                fmt.Println(v3)
                v21 := GenS2{F0: -1, F1: -906, F2: v12}
                _ = fi15
                _ = v21
            }
            _ = fi14
        }
    default:
        v22 := GenS1{F0: "\n\t"}
        _ = v22
    }
    v1 = v1
    var v23 uint = 371
    _ = len(v5)
    if len(v4) > 0 {
        v4[0] = 0
    }
    fmt.Println(v3, v1)
    var v24 uint32 = 2147483648
    if ((v12 / ((v12) | 1)) != (v12 + v12)) {
        fmt.Println(v1, v23)
    }
    var v25 string = "foo bar"
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v12
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
