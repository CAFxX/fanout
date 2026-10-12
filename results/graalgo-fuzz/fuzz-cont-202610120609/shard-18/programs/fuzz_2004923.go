package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint8; F1 uint }
type GenS4 struct { F0 uint16; F1 int32 }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 int64 }
func genFn0(p0 int16, p1 uint32, p2 int64) string {
    fmt.Println(p2, p1, p0)
    p2 -= (p2 * 924)
    _ = p0
    _ = p1
    _ = p2
    return ("" + "hello")
}
func genFn1(p0 int16) uint8 {
    p0 = p0
    switch "" {
    case "foo bar":
        var v6 bool = false
        _ = v6
    case "hello":
        var v7 int8 = -128
        _ = v7
    default:
        v8 := GenS1{F0: true}
        _ = v8
    }
    v9 := GenS2{F0: 2}
    _ = p0
    _ = v9
    return 42
}

func main() {
    var v1 uint32 = 147
    var v2 int = 588
    var v3 float32 = ((1e10 / 1.0) + (3.14159 - (3.14159 / 1.0)))
    var v4 float32 = 0.1
    var v5 int32 = 279
    for fi10 := 0; fi10 < 10; fi10++ {
        var v11 float64 = ((0.0 + (1.5 - 0.0)) - (0.1 * float64((float32(((3.14159 + 0.1) * -2.25)) + 0.0))))
        _ = fi10
        _ = v11
    }
    fmt.Println(v1, v3)
    var v12 int64 = 628
    if (true || (((("foo bar" + "hello") + "foo bar") + "\n\t") >= "a")) {
        switch "foo bar" {
        case "hello":
            var v13 bool = true
            _ = v13
        default:
            v14 := []int{v2, (v2 | -2147483648), 718, (-(v2))}
            _ = v14
        }
    } else {
        var v15 uint = 18446744073709551615
        for fi16 := 0; fi16 < 10; fi16++ {
            v17 := []int{fi16, fi16, (v2 - fi16), (fi16 & (fi16 % (((^(v2))) | 1)))}
            _ = fi16
            _ = v17
        }
        _ = v15
    }
    v18 := GenS3{F0: 9, F1: 65535}
    v19 := []int{42, v2, v2, (^(v2))}
    v19 = append(v19, v2)
    for fi20 := 0; fi20 < 17; fi20++ {
        v3 = float32((1.5 * float64(v4)))
        switch "" {
        case "hello":
            var v21 int = fi20
            _ = v21
        case "\n\t":
            v22 := GenS1{F0: true}
            _ = v22
        case "foo bar":
            v4 = (-2.25 + -2.25)
        default:
            var v23 int32 = -804
            _ = v23
        }
        _ = fi20
    }
    if (!false) {
        if true {
            switch "" {
            case "\n\t":
                var v24 uint16 = uint16(v1)
                _ = v24
            default:
                for fi25 := 0; fi25 < 1; fi25++ {
                    switch ("a" + "foo bar") {
                    case "hello":
                        v19 = append(v19, 128)
                    case "\n\t":
                        v26 := GenS4{F0: 255, F1: v5}
                        _ = v26
                    default:
                        v27 := GenS5{F0: 1}
                        _ = v27
                    }
                    v28 := GenS5{F0: uint16(v1)}
                    _ = fi25
                    _ = v28
                }
            }
        } else {
            if (!(((v5 * v5) <= (v5 + 1)) && (2 < 823))) {
                if false {
                    v29 := GenS6{F0: -1}
                    _ = v29
                } else {
                    if len(v19) >= 2 {
                        v19 = v19[1:len(v19)]
                    }
                }
            }
            if len(v19) >= 2 {
                v19 = v19[1:len(v19)]
            }
        }
        var v30 int32 = v5
        switch v2 {
        case -2147483648:
            v19 = append(v19, v2)
        default:
            for fi31 := 0; fi31 < 16; fi31++ {
                var v32 int64 = (-(v12))
                fmt.Println(v3)
                _ = fi31
                _ = v32
            }
        }
        _ = v30
    } else {
        v4 = ((-2.25 - (-2.25 + (v3 * v4))) + v3)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
