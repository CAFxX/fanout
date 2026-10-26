package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int16; F1 string }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 uint; F1 int64 }

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint32 = 42
    v3 := []string{"\n\t", "foo bar", "a"}
    v4 := make(map[uint32]uint32)
    var v5 float32 = ((0.1 * 1.5) * (0.0 - 0.1))
    if false {
        for fi6 := 0; fi6 < 17; fi6++ {
            fmt.Println(v5)
            if true {
                v4[2] = (v2 * (v2 << ((300) & 63)))
                v7 := v1[v2]
                _ = v7
            } else {
                if true {
                    var v8 uint8 = 255
                    _ = v8
                }
            }
            _ = fi6
        }
        v9 := GenS1{F0: v2, F1: 42}
        v3 = append(v3, "hello")
        _ = v9
    }
    var v10 uint32 = (^(v2))
    v11 := GenS2{F0: -812}
    var v12 float64 = float64(v5)
    if len(v3) > 0 {
        v3[0] = ("foo bar" + ("" + "foo bar"))
    }
    _ = len(v4)
    switch "a" {
    case "foo bar":
        v10 = v10
    case "hello":
        v13 := GenS1{F0: v2, F1: 2}
        _ = v13
    default:
        var v14 uint = 1
        _ = v14
    }
    v2 += v10
    _ = len(v3)
    var v15 int = -2147483648
    v15 = 128
    v4[710] = (v2 / ((v10) | 1))
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    fmt.Println(v2)
    if (9223372036854775807 > -9223372036854775808) {
        v16 := GenS3{F0: -32768, F1: "\n\t"}
        switch (v15 * v15) {
        case 32768:
            v12 = ((v12 / (((v12 / v12) / float64(v5)) - v12)) * v12)
        case 2147483647:
            v17 := GenS1{F0: 1, F1: 1}
            _ = v17
        case 512:
            if true {
                for fi18 := 0; fi18 < 1; fi18++ {
                    switch "hello" {
                    case "a":
                        v19 := v1[(v10 / ((v10) | 1))]
                        _ = v19
                    default:
                        var v20 float64 = 3.14159
                        _ = v20
                    }
                    _ = fi18
                }
            } else {
                switch "" {
                case "foo bar":
                    var v21 uint = 255
                    _ = v21
                case "a":
                    if false {
                        switch (v15 | v15) {
                        case -262:
                            v2 += v2
                        case 127:
                            v22 := GenS4{F0: 412}
                            _ = v22
                        default:
                            var v23 float64 = (v12 + v12)
                            _ = v23
                        }
                        v24 := GenS5{F0: 18446744073709551615, F1: 42}
                        v1[v10] = (627 < 807)
                        _ = v24
                    } else {
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        var v25 int8 = 100
                        _ = v25
                    }
                case "hello":
                    v10 -= 2
                default:
                    v15 -= (-(v15))
                }
                v5 -= v5
            }
        }
        _ = v16
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
