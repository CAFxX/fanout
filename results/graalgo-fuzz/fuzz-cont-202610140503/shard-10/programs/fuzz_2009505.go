package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32 }
func genFn0() uint {
    switch 2147483647 {
    case 9223372036854775807:
        var v5 uint32 = 1
        _ = v5
    default:
        switch 1 {
        case 128:
            var v6 float32 = 1.5
            _ = v6
        default:
            var v7 uint32 = 0
            _ = v7
        }
    }
    var v8 float32 = 1.5
    fmt.Println(v8)
    if true {
        for fi9 := 0; fi9 < 12; fi9++ {
            v10 := GenS1{F0: 13, F1: 1}
            v11 := []int{128, fi9, fi9}
            fmt.Println(fi9, v8)
            _ = fi9
            _ = v10
            _ = v11
        }
    }
    _ = v8
    return 4294967295
}
func genFn1(p0 int64, p1 uint8) int64 {
    fmt.Println(p0, p1)
    var v12 int64 = (p0 | (p0 - -712))
    v13 := []int{-252, 9223372036854775807, 32767, -9223372036854775808}
    if len(v13) > 0 {
        v13[0] = -76
    }
    p1 += 209
    _ = p0
    _ = p1
    _ = v12
    _ = v13
    return -367
}

func main() {
    var v1 int8 = -128
    v2 := []string{"hello"}
    var v3 float64 = (((3.14159 * 0.0) - 1e10) + ((0.0 - (0.0 - 3.14159)) / 1.5))
    var v4 int16 = -1
    var v14 int8 = (v1 * v1)
    var v15 int64 = 1
    switch ("" + ("a" + "\n\t")) {
    case "hello":
        if true {
            v2 = append(v2, "\n\t")
            var v16 string = "foo bar"
            var v17 uint8 = 128
            _ = v16
            _ = v17
        }
    case "":
        v4 -= 1000
    default:
        switch int(v4) {
        case 598:
            for fi18 := 0; fi18 < 14; fi18++ {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                if false {
                    v3 += v3
                    var v19 int = fi18
                    _ = v19
                }
                if (((float64((3.14159 + -2.25)) - v3) + -2.25) < v3) {
                    v15 -= int64(v1)
                }
                _ = fi18
            }
        }
    }
    switch 4294967295 {
    case 0:
        var v20 float32 = ((-2.25 * (((1e10 / (float32((v3 / 1.5)) * 3.14159)) + (float32(v3) - 0.0)) * (((((1.5 / 1.0) * -2.25) * 1.5) + 0.0) * (float32(((v3 * float64(0.1)) / 1e10)) / ((1.5 * (0.1 - (0.0 + (1e10 - 1e10)))) / 3.14159))))) + (1.5 * 1.5))
        _ = v20
    default:
        if len(v2) > 0 {
            v2[0] = ""
        }
    }
    switch 42 {
    case 2147483647:
        var v21 int16 = (v4 ^ v4)
        _ = v21
    case 204:
        var v22 uint = 42
        _ = v22
    }
    v23 := GenS1{F0: 65535, F1: 42}
    var v24 int8 = -128
    fmt.Println(v14, v14)
    if false {
        fmt.Println(v14, v15, v15)
        v15 = v15
        v25 := GenS1{F0: 65535, F1: 0}
        _ = v25
    }
    v1 += (v1 / (((^(v1))) | 1))
    v26 := GenS1{F0: 65535, F1: 1}
    fmt.Println(v3, v14, v3)
    for fi27 := 0; fi27 < 15; fi27++ {
        for fi28 := 0; fi28 < 2; fi28++ {
            v4 -= (v4 | (v4 / (((v4 - 378)) | 1)))
            v15 += (v15 | 9223372036854775807)
            var v29 uint16 = 389
            _ = fi28
            _ = v29
        }
        _ = fi27
    }
    var v30 uint16 = 154
    fmt.Println(v14, v4, v4)
    if len(v2) > 0 {
        v2[0] = "\n\t"
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v23
    _ = v24
    _ = v26
    _ = v3
    _ = v30
    _ = v4
}
