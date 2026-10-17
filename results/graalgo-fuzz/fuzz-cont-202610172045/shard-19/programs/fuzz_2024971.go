package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint16; F2 int }
type GenS2 struct { F0 bool; F1 bool }
func genFn0(p0 uint16, p1 int32) string {
    p0 = p0
    v4 := []int{128, -1, 823, -2147483648}
    var v5 int64 = 562
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return "\n\t"
}
func genFn1() uint16 {
    v6 := GenS1{F0: 0, F1: 0, F2: -881}
    switch ((("\n\t" + ("foo bar" + "foo bar")) + "\n\t") + "\n\t") {
    case "foo bar":
        var v7 int64 = 9223372036854775807
        _ = v7
    case "hello":
        var v8 float32 = 1.5
        _ = v8
    case "\n\t":
        v9 := GenS2{F0: (68 != -50), F1: true}
        _ = v9
    default:
        switch -25 {
        case 42:
            v10 := GenS2{F0: (32 == -128), F1: true}
            _ = v10
        case -998:
            var v11 float64 = -2.25
            _ = v11
        default:
            var v12 int16 = 1000
            _ = v12
        }
    }
    _ = v6
    return 255
}

func main() {
    var v1 float64 = 0.1
    var v2 uint = 42
    var v3 int32 = 647
    switch 1 {
    case 1:
        v3 += v3
    }
    var v13 int8 = 127
    fmt.Println(v13, v3)
    v2 = (v2 << (((v2 ^ v2)) & 63))
    var v14 float64 = (-2.25 + (3.14159 - v1))
    v15 := GenS2{F0: false, F1: (-9223372036854775808 == -1)}
    var v16 float32 = 1e10
    fmt.Println(v3, v13)
    v1 -= v1
    v17 := GenS2{F0: (!false), F1: false}
    switch 2147483647 {
    case -960:
        fmt.Println(v1, v16, v3)
    case -577:
        switch ("\n\t" + "") {
        case "":
            var v18 uint64 = 370
            _ = v18
        default:
            for fi19 := 0; fi19 < 15; fi19++ {
                fmt.Println(v14)
                fmt.Println(fi19, v14)
                _ = fi19
            }
        }
    case 42:
        v16 -= ((v16 - 1e10) / 0.1)
    }
    for fi20 := 0; fi20 < 12; fi20++ {
        v16 = (0.0 / v16)
        _ = fi20
    }
    for fi21 := 0; fi21 < 18; fi21++ {
        v2 = v2
        _ = fi21
    }
    switch 32767 {
    case 128:
        var v22 float64 = v1
        _ = v22
    case 32768:
        v1 = 3.14159
    default:
        fmt.Println(v13, v16, v1)
    }
    v16 += v16
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
}
