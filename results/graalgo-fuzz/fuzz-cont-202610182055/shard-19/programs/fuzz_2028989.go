package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8; F2 string }
type GenS2 struct { F0 int32; F1 uint8; F2 uint64 }
type GenS3 struct { F0 string; F1 uint16 }
func genFn0(p0 int16, p1 uint) int16 {
    v5 := GenS1{F0: 237, F1: 0, F2: ("hello" + "a")}
    var v6 uint16 = 453
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return p0
}
func genFn1() uint16 {
    switch "foo bar" {
    case "a":
        var v7 string = ("" + "hello")
        _ = v7
    case "hello":
        var v8 int32 = 973
        _ = v8
    case "\n\t":
        v9 := []int{301}
        _ = v9
    }
    var v10 float32 = float32(1.5)
    v10 = ((v10 / 1.5) - (v10 * float32(1e10)))
    v11 := GenS2{F0: 256, F1: 1, F2: 18446744073709551615}
    var v12 float32 = (v10 + (1.5 / float32(-2.25)))
    _ = v10
    _ = v11
    _ = v12
    return 2
}

func main() {
    var v1 uint64 = 0
    var v2 float64 = 0.0
    var v3 int64 = 334
    var v4 float32 = 1e10
    v13 := []int{-186, -906}
    var v14 uint16 = 2
    v1 = (^(v1))
    if (2 <= 903) {
        var v15 int16 = 949
        v13 = append(v13, int(v14))
        _ = v15
    } else {
        var v16 int16 = 42
        _ = v16
    }
    for fi17 := 0; fi17 < 5; fi17++ {
        fmt.Println(v14, v14)
        _ = fi17
    }
    switch 225 {
    case 9223372036854775807:
        v2 -= ((float64(0.1) * ((v2 * 1.5) * v2)) + 1e10)
    case -210:
        var v18 uint16 = v14
        _ = v18
    case 0:
        for fi19 := 0; fi19 < 10; fi19++ {
            _ = len(v13)
            v14 += v14
            _ = fi19
        }
    default:
        _ = len(v13)
    }
    var v20 uint8 = 0
    v21 := GenS3{F0: ("\n\t" + "foo bar"), F1: v14}
    v14 += (^(v14))
    for fi22 := 0; fi22 < 17; fi22++ {
        v23 := GenS1{F0: (v14 / ((2) | 1)), F1: 7, F2: ""}
        _ = fi22
        _ = v23
    }
    var v24 int = 32767
    for fi25 := 0; fi25 < 3; fi25++ {
        var v26 uint8 = (v20 ^ v20)
        _ = fi25
        _ = v26
    }
    v3 = v3
    var v27 float32 = ((v4 / ((v4 + 0.0) * ((((float32(-2.25) - v4) / (v4 - (((v4 + -2.25) * 1e10) * (v4 - (v4 * v4))))) - (float32(((float64(float32(float64(v4))) - (3.14159 - v2)) + v2)) + (v4 / v4))) - ((((v4 + v4) - 0.1) + ((v4 / 1.5) / 1.0)) / 3.14159)))) / v4)
    v13 = append(v13, v24)
    switch "foo bar" {
    case "":
        var v28 int64 = (v3 * v3)
        _ = v28
    case "a":
        fmt.Println(v14)
    }
    if (!("\n\t" <= "foo bar")) {
        fmt.Println(v24)
    }
    var v29 uint32 = uint32(v14)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v24:", v24)
    fmt.Println("v27:", v27)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v20
    _ = v21
    _ = v24
    _ = v27
    _ = v29
    _ = v3
    _ = v4
}
