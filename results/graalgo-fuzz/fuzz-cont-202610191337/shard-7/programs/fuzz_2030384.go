package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int64; F1 uint32 }
func genFn0(p0 int8, p1 uint8, p2 uint) int32 {
    v5 := GenS1{F0: 255, F1: 220}
    for fi6 := 0; fi6 < 10; fi6++ {
        p0 -= p0
        _ = fi6
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 157
}
func genFn1(p0 int8) int {
    v7 := GenS1{F0: 1, F1: 4294967295}
    if true {
        fmt.Println(p0)
        var v8 float32 = float32(float64(float32(1.5)))
        v9 := GenS1{F0: 255, F1: 0}
        _ = v8
        _ = v9
    }
    _ = p0
    _ = v7
    return 128
}

func main() {
    v1 := make(map[string]string)
    var v2 bool = false
    v3 := make(map[int]uint32)
    v4 := []string{"foo bar"}
    v2 = v2
    switch "\n\t" {
    case "hello":
        switch 32768 {
        case -121:
            var v10 uint64 = 42
            _ = v10
        case 1:
            var v11 int8 = -93
            _ = v11
        default:
            if (v2 && v2) {
                if ("" >= "\n\t") {
                    var v12 float32 = (0.0 * 1.5)
                    for fi13 := 0; fi13 < 17; fi13++ {
                        var v14 int = (fi13 >> ((255) & 63))
                        _ = fi13
                        _ = v14
                    }
                    _ = v12
                } else {
                    _ = len(v3)
                }
                if (false || v2) {
                    var v15 int64 = 42
                    _ = v15
                } else {
                    v2 = ((true && true) && v2)
                    v2 = false
                }
                v2 = v2
            }
        }
    case "\n\t":
        fmt.Println(v2)
    }
    v2 = true
    for fi16 := 0; fi16 < 4; fi16++ {
        v2 = ((fi16 / (((fi16 ^ genFn1(0))) | 1)) == (fi16 / ((genFn1(1)) | 1)))
        _ = fi16
    }
    if len(v4) > 0 {
        v4[0] = (("a" + "foo bar") + "")
    }
    switch 9223372036854775807 {
    case -838:
        v2 = v2
    case 2147483647:
        if v2 {
            v17 := GenS2{F0: 2}
            fmt.Println(v2)
            _ = v17
        }
    default:
        if (false && (42 != 4294967295)) {
            delete(v1, ("a" + "\n\t"))
            _ = len(v4)
            v2 = (!true)
        } else {
            v18 := GenS3{F0: 9223372036854775807, F1: 4294967295}
            if (((0.0 <= 3.14159) && v2) && true) {
                var v19 uint = 2
                _ = v19
            }
            _ = v18
        }
    }
    v2 = (genFn1(0) >= -745)
    v2 = ((3.14159 - (1e10 / -2.25)) >= 0.1)
    var v20 uint16 = 255
    var v21 string = ("a" + ("foo bar" + "foo bar"))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
}
