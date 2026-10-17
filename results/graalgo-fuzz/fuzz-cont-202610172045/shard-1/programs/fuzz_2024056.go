package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int16; F1 uint8; F2 uint }
type GenS3 struct { F0 int32; F1 uint8 }
func genFn0(p0 int8, p1 uint) uint32 {
    p0 = p0
    var v6 float64 = ((1.5 + -2.25) / (3.14159 * (float64(float32(1e10)) * ((float64(float32((1.5 / 0.1))) + -2.25) - float64(-2.25)))))
    switch ("foo bar" + (("" + "foo bar") + ("" + "a"))) {
    case "foo bar":
        fmt.Println(p0)
    case "hello":
        p0 += (-(p0))
    case "":
        var v7 string = (("" + "") + "hello")
        _ = v7
    default:
        p1 = (p1 & 65535)
    }
    _ = p0
    _ = p1
    _ = v6
    return 4294967295
}
func genFn1(p0 uint) uint {
    v8 := GenS1{F0: 98}
    v9 := []int{128, -9223372036854775808, 128, 128}
    p0 -= (p0 * 221)
    p0 = 18446744073709551615
    _ = p0
    _ = v8
    _ = v9
    return (p0 + (p0 ^ 4294967295))
}

func main() {
    var v1 int = 42
    var v2 float64 = float64(1.5)
    var v3 uint16 = 255
    var v4 int = 32767
    var v5 uint32 = 0
    v10 := []int{}
    if (-104 <= 1) {
        v5 = (v5 - 84)
        switch 4294967295 {
        case 48:
            if len(v10) > 0 {
                v10[0] = int(v5)
            }
        case -855:
            var v11 uint32 = (v5 - 2147483648)
            _ = v11
        case -857:
            var v12 uint8 = 1
            _ = v12
        default:
            fmt.Println(v3)
        }
        v3 -= v3
    } else {
        if true {
            switch "foo bar" {
            case "\n\t":
                switch v4 {
                case 0:
                    v2 = (3.14159 / 1e10)
                case 42:
                    v4 = 526
                case 76:
                    if (float32(1e10) < ((0.0 + (float32(v2) - 0.1)) - (-2.25 * -2.25))) {
                        if len(v10) >= 2 {
                            v10 = v10[1:len(v10)]
                        }
                    } else {
                        v3 = 0
                        for fi13 := 0; fi13 < 10; fi13++ {
                            v2 = (((-2.25 - v2) * 1e10) - v2)
                            _ = fi13
                        }
                    }
                }
            case "foo bar":
                v4 = (-(v1))
            case "hello":
                if (2147483648 == 4294967295) {
                    fmt.Println(v2, v3)
                    v4 = -2147483648
                    for fi14 := 0; fi14 < 10; fi14++ {
                        v4 += fi14
                        v15 := GenS1{F0: v3}
                        fmt.Println(v5, v2)
                        _ = fi14
                        _ = v15
                    }
                } else {
                    v1 = (v4 * (v1 >> ((1) & 63)))
                }
            default:
                v16 := GenS2{F0: 1000, F1: 42, F2: 2}
                _ = v16
            }
        } else {
            v17 := GenS1{F0: v3}
            var v18 int32 = -1
            _ = v17
            _ = v18
        }
    }
    for fi19 := 0; fi19 < 19; fi19++ {
        switch "hello" {
        case "":
            if (0 != 42) {
                var v20 uint16 = v3
                var v21 string = ""
                v21 = v21
                _ = v20
                _ = v21
            }
        case "foo bar":
            v4 = v1
        }
        v22 := GenS3{F0: 42, F1: 253}
        fmt.Println(v3)
        _ = fi19
        _ = v22
    }
    var v23 uint64 = 42
    fmt.Println(v5, v4, v23)
    var v24 string = (("hello" + "") + "hello")
    if ((0 < 32767) || false) {
        var v25 uint64 = (^(v23))
        if (false && false) {
            var v26 string = v24
            _ = v26
        } else {
            var v27 int32 = 42
            var v28 float64 = (0.0 - 1.5)
            _ = v27
            _ = v28
        }
        var v29 uint32 = v5
        _ = v25
        _ = v29
    }
    if (!false) {
        var v30 int16 = 42
        _ = v30
    } else {
        v1 -= v4
    }
    var v31 int8 = int8(v3)
    _ = len(v10)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v31
    _ = v4
    _ = v5
}
