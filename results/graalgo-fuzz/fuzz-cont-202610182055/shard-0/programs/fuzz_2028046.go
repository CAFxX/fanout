package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int32; F1 string }
type GenS3 struct { F0 uint64 }
func genFn0(p0 uint16, p1 string) uint {
    if (0 < -708) {
        v7 := GenS1{F0: 2}
        v8 := []int{127}
        var v9 uint = 1
        _ = v7
        _ = v8
        _ = v9
    } else {
        var v10 uint8 = 86
        _ = v10
    }
    v11 := GenS2{F0: 0, F1: ((("hello" + (p1 + "\n\t")) + (p1 + p1)) + "")}
    if (!(true && true)) {
        v12 := []int{}
        for fi13 := 0; fi13 < 3; fi13++ {
            v12 = append(v12, (fi13 << ((603) & 63)))
            fmt.Println(fi13, fi13, p1)
            _ = fi13
        }
        _ = v12
    }
    p0 += p0
    if (true || (-2147483648 <= 32768)) {
        for fi14 := 0; fi14 < 20; fi14++ {
            var v15 uint16 = (^(p0))
            for fi16 := 0; fi16 < 1; fi16++ {
                var v17 string = "\n\t"
                v17 = p1
                var v18 int64 = 0
                v15 = p0
                _ = fi16
                _ = v17
                _ = v18
            }
            p1 = (((((p1 + "foo bar") + ((((("\n\t" + p1) + p1) + "hello") + "a") + (((p1 + ((p1 + (((p1 + ((((p1 + "hello") + p1) + p1) + p1)) + (p1 + p1)) + (p1 + p1))) + (p1 + p1))) + p1) + (p1 + ((p1 + "\n\t") + p1))))) + p1) + (p1 + p1)) + p1)
            _ = fi14
            _ = v15
        }
        var v19 int64 = 513
        _ = v19
    } else {
        v20 := []int{9223372036854775807, 237}
        switch -2147483648 {
        case -1:
            fmt.Println(p0, p0)
        case 128:
            if len(v20) >= 2 {
                v20 = v20[1:len(v20)]
            }
        default:
            v21 := GenS2{F0: -605, F1: p1}
            _ = v21
        }
        p0 -= (p0 >> ((65535) & 63))
        _ = v20
    }
    _ = p0
    _ = p1
    _ = v11
    return 91
}
func genFn1(p0 uint8) int8 {
    switch "\n\t" {
    case "hello":
        if (!false) {
            for fi22 := 0; fi22 < 17; fi22++ {
                v23 := []int{-568, (fi22 / (((fi22 + (fi22 << ((2) & 63)))) | 1))}
                _ = fi22
                _ = v23
            }
        } else {
            var v24 float32 = 1e10
            var v25 float64 = (((1e10 / 1.5) + -2.25) / (3.14159 + -2.25))
            _ = v24
            _ = v25
        }
    case "":
        v26 := GenS3{F0: 1}
        _ = v26
    }
    fmt.Println(p0)
    fmt.Println(p0)
    _ = p0
    return 85
}

func main() {
    var v1 float32 = 1e10
    var v2 int16 = -659
    var v3 int = -320
    v4 := make(map[string]string)
    var v5 int32 = -2147483648
    var v6 uint16 = 42
    v3 = (v3 * (v3 * 128))
    v5 = (v5 * v5)
    var v27 float64 = ((-2.25 + -2.25) * float64(1.5))
    v28 := []int{v3, -2147483648, (v3 | (v3 ^ (v3 - v3)))}
    switch "" {
    case "foo bar":
        v3 = (v3 & -14)
    default:
        var v29 string = "a"
        _ = v29
    }
    for fi30 := 0; fi30 < 7; fi30++ {
        switch "\n\t" {
        case "":
            v5 = v5
        default:
            for fi31 := 0; fi31 < 12; fi31++ {
                if len(v28) >= 2 {
                    v28 = v28[1:len(v28)]
                }
                _ = fi31
            }
        }
        v32 := GenS2{F0: -2147483648, F1: ("foo bar" + "\n\t")}
        _ = fi30
        _ = v32
    }
    v1 = ((v1 - v1) - v1)
    v3 = 127
    v33 := GenS3{F0: 18446744073709551615}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v28:", len(v28), v28)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}
