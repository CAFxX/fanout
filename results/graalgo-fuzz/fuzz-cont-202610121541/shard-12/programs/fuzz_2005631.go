package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int8; F2 uint16 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 string; F1 int64 }
func genFn0(p0 uint8, p1 int64, p2 uint) int16 {
    switch "hello" {
    case "hello":
        p1 -= 42
    case "foo bar":
        v5 := GenS1{F0: 4294967295, F1: -46, F2: 42}
        _ = v5
    case "":
        for fi6 := 0; fi6 < 18; fi6++ {
            v7 := GenS1{F0: 947, F1: 46, F2: 255}
            p1 -= p1
            _ = fi6
            _ = v7
        }
    default:
        switch 4294967295 {
        case 128:
            var v8 float32 = (((-2.25 - (0.0 / (float32(1.5) * 3.14159))) - 0.1) / 3.14159)
            _ = v8
        case 384:
            p2 -= (p2 | (^(p2)))
        default:
            v9 := []int{32768, -15, 9223372036854775807, -724}
            _ = v9
        }
    }
    var v10 float32 = (1.5 / 1e10)
    for fi11 := 0; fi11 < 13; fi11++ {
        var v12 uint8 = p0
        v12 -= 1
        _ = fi11
        _ = v12
    }
    if false {
        var v13 uint = (^(p2))
        _ = v13
    } else {
        var v14 float64 = 1.5
        for fi15 := 0; fi15 < 6; fi15++ {
            for fi16 := 0; fi16 < 20; fi16++ {
                v17 := []int{(fi16 | 4294967295), 32768}
                var v18 uint32 = 2147483647
                switch fi16 {
                case 32768:
                    v19 := GenS1{F0: 716, F1: 127, F2: 58}
                    _ = v19
                case 68:
                    v20 := GenS1{F0: 4294967295, F1: 127, F2: 2}
                    _ = v20
                case -2147483648:
                    v10 -= v10
                default:
                    var v21 uint8 = p0
                    v22 := GenS1{F0: 352, F1: 120, F2: 65535}
                    _ = len(v17)
                    _ = v21
                    _ = v22
                }
                _ = fi16
                _ = v17
                _ = v18
            }
            v23 := []int{127, fi15, -544}
            if len(v23) >= 2 {
                v23 = v23[1:len(v23)]
            }
            _ = fi15
            _ = v23
        }
        _ = v14
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    return -429
}
func genFn1(p0 int32, p1 uint64, p2 uint16) int8 {
    p0 = p0
    v24 := GenS1{F0: (p1 % ((p1) | 1)), F1: 43, F2: (p2 + p2)}
    _ = p0
    _ = p1
    _ = p2
    _ = v24
    return 14
}

func main() {
    v1 := []int{-78, 9223372036854775807, 9223372036854775807, -445}
    v2 := make(map[string]uint32)
    v3 := []int{-9223372036854775808, 9223372036854775807, 128, -2147483648}
    var v4 int = 127
    var v25 int32 = -507
    v25 = v25
    var v26 uint64 = 0
    v26 = v26
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v27 := GenS2{F0: 268}
    v28 := v2[((("" + ("" + ("" + (("\n\t" + ("a" + "hello")) + ("\n\t" + ("" + "hello")))))) + ("" + "a")) + ("" + ("a" + ("hello" + "hello"))))]
    var v29 uint16 = 65535
    for fi30 := 0; fi30 < 8; fi30++ {
        if true {
            v29 = v29
            v29 = (v29 ^ v29)
        } else {
            var v31 float64 = 0.0
            v32 := GenS3{F0: "\n\t", F1: 0}
            if len(v1) > 0 {
                v1[0] = (^(fi30))
            }
            _ = v31
            _ = v32
        }
        var v33 uint = 18446744073709551615
        v34 := GenS3{F0: "foo bar", F1: int64(v29)}
        _ = fi30
        _ = v33
        _ = v34
    }
    v25 -= v25
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
}
