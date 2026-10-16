package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 string; F2 bool }
type GenS2 struct { F0 int; F1 uint; F2 uint }
type GenS3 struct { F0 int64; F1 int8 }
type GenS4 struct { F0 uint8; F1 uint16; F2 string }
type GenS5 struct { F0 uint16 }
func genFn0(p0 uint, p1 uint64) int32 {
    fmt.Println(p0)
    v5 := GenS1{F0: 0, F1: "foo bar", F2: (1 > 348)}
    v6 := GenS2{F0: 84, F1: (^(p0)), F2: p0}
    var v7 uint32 = 618
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    _ = v7
    return -2147483648
}
func genFn1(p0 int64) int8 {
    fmt.Println(p0)
    if true {
        if (!true) {
            fmt.Println(p0)
            switch 525 {
            case -660:
                v8 := GenS1{F0: 32767, F1: ("foo bar" + "hello"), F2: true}
                _ = v8
            case -614:
                p0 -= p0
            case 127:
                switch "" {
                case "":
                    if (false || true) {
                        p0 = p0
                        var v9 bool = true
                        _ = v9
                    }
                case "foo bar":
                    v10 := []int{-9223372036854775808, 0}
                    _ = v10
                default:
                    v11 := GenS3{F0: p0, F1: -128}
                    _ = v11
                }
            default:
                v12 := GenS1{F0: -32768, F1: ((("" + "\n\t") + "hello") + "a"), F2: true}
                _ = v12
            }
        }
        p0 = p0
        p0 += p0
    } else {
        for fi13 := 0; fi13 < 11; fi13++ {
            var v14 int = -9223372036854775808
            _ = fi13
            _ = v14
        }
    }
    _ = p0
    return -1
}

func main() {
    var v1 bool = (255 < 2)
    v2 := []int{0}
    v3 := make(map[string]string)
    var v4 int8 = -97
    for fi15 := 0; fi15 < 13; fi15++ {
        v16 := GenS2{F0: fi15, F1: 65535, F2: 1}
        v17 := GenS3{F0: 42, F1: v4}
        _ = fi15
        _ = v16
        _ = v17
    }
    switch "\n\t" {
    case "a":
        v18 := GenS1{F0: 887, F1: "a", F2: (-571 < 605)}
        _ = v18
    case "\n\t":
        v4 -= 100
    }
    v19 := GenS4{F0: 2, F1: 1, F2: ("foo bar" + "a")}
    fmt.Println(v1)
    v20 := GenS5{F0: 2}
    if len(v2) > 0 {
        v2[0] = 128
    }
    var v21 int8 = genFn1(-888)
    var v22 float32 = (1e10 + (float32(0.0) / ((0.1 / 3.14159) * 1.5)))
    fmt.Println(v21, v21, v1)
    var v23 float32 = (((v22 / (0.0 - ((v22 + v22) - v22))) / ((-2.25 + 1e10) - -2.25)) * (0.0 - (float32(1.5) / (-2.25 + v22))))
    var v24 float64 = float64(v23)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
}
