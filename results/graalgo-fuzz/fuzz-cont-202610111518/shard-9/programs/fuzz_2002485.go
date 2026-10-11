package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int32; F2 int8 }
type GenS2 struct { F0 uint32; F1 uint }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 string }
func genFn0(p0 int) bool {
    p0 = (p0 + (p0 - p0))
    v7 := []int{}
    _ = p0
    _ = v7
    return (2 == 4294967295)
}
func genFn1(p0 string, p1 int32, p2 string) uint32 {
    fmt.Println(p0, p0, p2)
    v8 := []int{42, -728, -742, -952}
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return 0
}

func main() {
    var v1 int8 = -22
    var v2 uint8 = 218
    var v3 string = "a"
    var v4 uint = 438
    var v5 int8 = v1
    var v6 int8 = (^(v1))
    for fi9 := 0; fi9 < 4; fi9++ {
        if ((true && ((v2 << (((v4 / ((v4) | 1))) & 63)) == 255)) || false) {
            var v10 uint8 = (^(v2))
            switch ("" + "\n\t") {
            case "a":
                if ((0.0 + ((1.5 * float64((-2.25 / (1e10 * 1.5)))) / -2.25)) < (0.0 / ((((0.1 - 0.0) - float64(1e10)) * 1e10) - 0.1))) {
                    v2 += (v2 + (v10 + 2))
                } else {
                    var v11 int = fi9
                    v6 -= (v5 ^ v5)
                    for fi12 := 0; fi12 < 6; fi12++ {
                        fmt.Println(v4, v2)
                        v3 = v3
                        var v13 int8 = (v6 % ((v5) | 1))
                        _ = fi12
                        _ = v13
                    }
                    _ = v11
                }
            default:
                fmt.Println(v5)
            }
            _ = v10
        }
        _ = fi9
    }
    v14 := GenS1{F0: (-(v6)), F1: 1, F2: 1}
    v15 := GenS2{F0: 775, F1: (v4 * v4)}
    fmt.Println(v6, v4, v1)
    v1 -= -45
    switch v3 {
    case "a":
        for fi16 := 0; fi16 < 7; fi16++ {
            fmt.Println(v2, v1, v2)
            var v17 string = (("" + v3) + v3)
            _ = fi16
            _ = v17
        }
    default:
        v18 := []int{}
        _ = v18
    }
    v19 := GenS3{F0: 32767}
    v4 -= (v4 / ((v4) | 1))
    var v20 float32 = (-2.25 + float32(0.0))
    if true {
        v21 := []int{128}
        _ = v21
    } else {
        v22 := GenS4{F0: (v3 + v3)}
        var v23 int = -1
        v24 := GenS4{F0: ""}
        _ = v22
        _ = v23
        _ = v24
    }
    var v25 int16 = int16(v6)
    v26 := GenS2{F0: 4294967295, F1: v4}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v19
    _ = v2
    _ = v20
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
