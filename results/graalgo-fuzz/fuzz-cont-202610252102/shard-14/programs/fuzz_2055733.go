package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint }
type GenS2 struct { F0 uint; F1 int32; F2 int64 }
func genFn0(p0 uint, p1 uint16) uint16 {
    for fi6 := 0; fi6 < 9; fi6++ {
        var v7 int = (fi6 % ((fi6) | 1))
        _ = fi6
        _ = v7
    }
    if false {
        p1 = p1
        var v8 int32 = -1
        for fi9 := 0; fi9 < 8; fi9++ {
            v10 := []int{fi9}
            v11 := GenS1{F0: 2147483648, F1: (p0 / ((p0) | 1))}
            _ = fi9
            _ = v10
            _ = v11
        }
        _ = v8
    } else {
        var v12 int8 = -128
        var v13 uint8 = uint8(p0)
        if true {
            fmt.Println(v13)
            var v14 string = ""
            fmt.Println(p1, v12, v13)
            _ = v14
        }
        _ = v12
        _ = v13
    }
    var v15 uint16 = 2
    _ = p0
    _ = p1
    _ = v15
    return (^(v15))
}

func main() {
    var v1 float32 = 1e10
    v2 := []string{"foo bar", "\n\t"}
    v3 := []int{-610, 1, 128, 32767}
    var v4 float32 = -2.25
    v5 := make(map[uint32]int)
    switch "a" {
    case "foo bar":
        var v16 string = "\n\t"
        _ = v16
    }
    if (float32(-2.25) <= 3.14159) {
        for fi17 := 0; fi17 < 8; fi17++ {
            if (65535 == 255) {
                for fi18 := 0; fi18 < 2; fi18++ {
                    v1 += 1.5
                    _ = fi18
                }
                fmt.Println(v4)
                fmt.Println(v1, fi17, v4)
            }
            _ = fi17
        }
        var v19 uint16 = 1
        var v20 float64 = (1.5 + 3.14159)
        _ = v19
        _ = v20
    } else {
        v2 = append(v2, ("\n\t" + "\n\t"))
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        v21 := GenS1{F0: 65535, F1: 2}
        _ = v21
    }
    switch -9223372036854775808 {
    case 466:
        v4 -= ((v4 / (v4 + ((v1 / v1) / (-2.25 + (v1 - 0.0))))) * (1.5 * v1))
    }
    var v22 int64 = 465
    if ((3.14159 / (((1.5 * ((v4 - (v4 + v1)) * (((float32(0.1) / 0.1) + 0.1) * v1))) - (3.14159 * v1)) - (-2.25 - (((0.0 / float32(-2.25)) - v4) * 0.1)))) > float32((0.1 * 1e10))) {
        _ = len(v5)
        v5[4294967295] = 32768
    }
    delete(v5, 0)
    v23 := GenS2{F0: 4294967295, F1: 42, F2: (v22 >> ((255) & 63))}
    switch -2147483648 {
    case -213:
        var v24 int8 = 127
        _ = v24
    default:
        var v25 uint = 42
        _ = v25
    }
    for fi26 := 0; fi26 < 6; fi26++ {
        var v27 uint64 = 2
        _ = fi26
        _ = v27
    }
    v22 -= -9223372036854775808
    switch "" {
    case "foo bar":
        var v28 int = 42
        _ = v28
    }
    v4 += ((float32(1e10) / v1) + float32((3.14159 - (1.5 + (-2.25 + ((-2.25 * 3.14159) * 1e10))))))
    v1 = v1
    if false {
        v4 += ((float32((-2.25 / 3.14159)) - v1) - ((0.0 * ((v4 / v4) + (((0.1 / v4) / (-2.25 / v4)) * 0.0))) / v4))
        if (!true) {
            fmt.Println(v4)
            fmt.Println(v4, v1)
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}
