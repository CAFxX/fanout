package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint; F1 uint16; F2 uint8 }
func genFn0(p0 int32) int16 {
    var v4 float64 = 3.14159
    var v5 int32 = -1
    if true {
        var v6 float32 = 1.5
        _ = v6
    } else {
        for fi7 := 0; fi7 < 7; fi7++ {
            v8 := GenS1{F0: 9223372036854775807}
            v5 = (p0 / ((v5) | 1))
            _ = fi7
            _ = v8
        }
        for fi9 := 0; fi9 < 2; fi9++ {
            v4 -= (3.14159 - v4)
            for fi10 := 0; fi10 < 2; fi10++ {
                if true {
                    fmt.Println(p0)
                    var v11 int = (fi9 | fi9)
                    var v12 uint16 = 199
                    _ = v11
                    _ = v12
                } else {
                    v13 := []int{fi9, -369, (-(fi9))}
                    _ = v13
                }
                _ = fi10
            }
            v14 := GenS1{F0: -9223372036854775808}
            _ = fi9
            _ = v14
        }
        v15 := []int{0, 128, -2147483648}
        _ = v15
    }
    _ = p0
    _ = v4
    _ = v5
    return 0
}

func main() {
    var v1 string = "\n\t"
    var v2 uint32 = 65535
    var v3 int16 = 608
    v1 = v1
    v16 := []int{-2147483648, -9223372036854775808, 9223372036854775807, 42}
    if false {
        v17 := GenS1{F0: -680}
        var v18 uint16 = 34
        _ = v17
        _ = v18
    }
    v3 += genFn0(0)
    var v19 int32 = -2147483648
    for fi20 := 0; fi20 < 16; fi20++ {
        v16 = append(v16, fi20)
        var v21 int64 = -1
        var v22 uint16 = 65535
        _ = fi20
        _ = v21
        _ = v22
    }
    switch v1 {
    case "hello":
        v19 += -829
    }
    if true {
        v23 := GenS2{F0: 4294967295, F1: 32767, F2: uint8(v2)}
        fmt.Println(v2, v3, v3)
        _ = v23
    }
    switch 0 {
    case 1:
        var v24 int16 = (v3 / ((v3) | 1))
        _ = v24
    case 0:
        if (false && (1 >= 4)) {
            v3 = 1
            fmt.Println(v3, v2)
        } else {
            var v25 uint = 18446744073709551615
            var v26 float32 = (float32(((1.5 / (1.5 + 1.5)) * 3.14159)) * (1e10 * 3.14159))
            var v27 int32 = (v19 % (((v19 & v19)) | 1))
            _ = v25
            _ = v26
            _ = v27
        }
    case 4294967295:
        v16 = append(v16, 127)
    default:
        var v28 bool = true
        _ = v28
    }
    var v29 int = 660
    v29 -= 32768
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v16
    _ = v19
    _ = v2
    _ = v29
    _ = v3
}
