package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8 }
func genFn0(p0 int, p1 int32, p2 uint64) uint64 {
    switch "foo bar" {
    case "":
        p0 = 1
    default:
        switch p0 {
        case 127:
            p0 = p0
        case -338:
            var v5 float64 = (3.14159 - float64(3.14159))
            _ = v5
        case 32768:
            p0 = p0
        default:
            p1 = (-(p1))
        }
    }
    for fi6 := 0; fi6 < 2; fi6++ {
        v7 := []int{1}
        for fi8 := 0; fi8 < 6; fi8++ {
            if len(v7) >= 2 {
                v7 = v7[1:len(v7)]
            }
            p2 -= p2
            fmt.Println(p2)
            _ = fi8
        }
        p1 += (p1 % (((-(p1))) | 1))
        _ = fi6
        _ = v7
    }
    for fi9 := 0; fi9 < 10; fi9++ {
        v10 := []int{fi9, (p0 / ((-301) | 1))}
        v10 = append(v10, (fi9 / (((p0 ^ 127)) | 1)))
        _ = fi9
        _ = v10
    }
    fmt.Println(p1)
    var v11 int32 = p1
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    return (p2 >> ((65535) & 63))
}
func genFn1(p0 int32) uint16 {
    var v12 string = "foo bar"
    if (((float64((0.1 + (float32((0.1 / 3.14159)) - 1e10))) / 1e10) <= (0.0 + (3.14159 / ((float64((1.5 - ((3.14159 / 0.1) * (0.0 / -2.25)))) + ((((-2.25 - float64(float32(((1e10 - ((1.5 / 0.1) + 0.0)) + (1e10 * 3.14159))))) - (0.1 * ((3.14159 / float64((1e10 / ((-2.25 + 1.5) / 3.14159)))) - 0.1))) * ((0.1 / 0.1) / 0.1)) * -2.25)) - 0.1)))) || true) {
        p0 += p0
    } else {
        p0 = 834
        v12 = ((((((v12 + (v12 + "foo bar")) + "\n\t") + "\n\t") + "hello") + v12) + ("" + (v12 + v12)))
        p0 -= (p0 % ((p0) | 1))
    }
    _ = p0
    _ = v12
    return 65535
}

func main() {
    var v1 uint = 4294967295
    v2 := []uint32{1, 1}
    v3 := make(map[string]uint32)
    v4 := make(map[uint32]string)
    if (-32768 == -1) {
        var v13 string = "hello"
        var v14 float64 = ((0.1 * ((3.14159 / 1.0) * (0.1 * (float64(3.14159) / (1.5 / 0.1))))) / -2.25)
        _ = len(v2)
        _ = v13
        _ = v14
    } else {
        v1 += v1
        if (!(!(!("foo bar" <= "hello")))) {
            v1 = v1
        }
    }
    var v15 uint16 = 1
    var v16 int64 = -9223372036854775808
    var v17 float32 = ((3.14159 / (1.5 - 1e10)) - ((float32(((((float64((0.0 + -2.25)) + ((0.1 * ((float64(-2.25) / -2.25) * 1e10)) - 1.5)) - 3.14159) - 3.14159) * (3.14159 + (1e10 + 0.1)))) / (0.1 * -2.25)) * 0.0))
    if len(v2) > 0 {
        v2[0] = 7
    }
    var v18 int = 9223372036854775807
    var v19 uint32 = 0
    v4[v19] = (("a" + ("hello" + "hello")) + "foo bar")
    if false {
        for fi20 := 0; fi20 < 10; fi20++ {
            v19 = v19
            _ = fi20
        }
        for fi21 := 0; fi21 < 10; fi21++ {
            for fi22 := 0; fi22 < 9; fi22++ {
                v23 := GenS1{F0: (v15 >> (((^(v1))) & 63)), F1: 42}
                if ((v16 - v16) > (^(v16))) {
                    var v24 int16 = 608
                    v1 = (v1 + v1)
                    var v25 string = ("hello" + ("" + "\n\t"))
                    _ = v24
                    _ = v25
                } else {
                    fmt.Println(v1, v17)
                    v2 = append(v2, 738)
                    if false {
                        v17 += float32((float64(v17) * float64(v17)))
                        var v26 int64 = (v16 % ((1) | 1))
                        _ = v26
                    }
                }
                _ = fi22
                _ = v23
            }
            for fi27 := 0; fi27 < 7; fi27++ {
                fmt.Println(fi21, v18)
                _ = fi27
            }
            _ = fi21
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
}
