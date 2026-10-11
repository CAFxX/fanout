package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint16; F2 int }

func main() {
    var v1 int64 = 133
    v2 := []uint32{65535, 2, 529, 2147483647}
    var v3 uint32 = 4294967295
    var v4 uint8 = uint8(v3)
    var v5 int32 = -2147483648
    for fi6 := 0; fi6 < 19; fi6++ {
        fmt.Println(fi6, v4, v1)
        fmt.Println(fi6)
        _ = fi6
    }
    v7 := GenS1{F0: v1, F1: 777, F2: -850}
    v2 = append(v2, v3)
    switch -174 {
    case 9223372036854775807:
        v3 = v3
    case 32768:
        var v8 int32 = (v5 + (-(v5)))
        _ = v8
    case -9223372036854775808:
        var v9 string = ""
        _ = v9
    }
    _ = len(v2)
    v2 = append(v2, (v3 >> ((340) & 63)))
    for fi10 := 0; fi10 < 1; fi10++ {
        if false {
            var v11 int32 = (^(v5))
            var v12 int32 = v11
            var v13 int8 = 0
            _ = v11
            _ = v12
            _ = v13
        }
        _ = len(v2)
        _ = fi10
    }
    v14 := GenS1{F0: (v1 - v1), F1: 1, F2: 32767}
    fmt.Println(v5, v5)
    var v15 float32 = float32((0.1 / 1e10))
    switch -531 {
    case 42:
        _ = len(v2)
    case 231:
        switch 1 {
        case -1:
            v16 := GenS1{F0: (v1 & (v1 << ((0) & 63))), F1: 1, F2: 827}
            _ = v16
        case -608:
            v15 = 0.1
        default:
            if (v1 < (-(v1))) {
                if len(v2) > 0 {
                    v2[0] = uint32(v4)
                }
                var v17 string = ""
                var v18 int64 = (v1 << ((882) & 63))
                _ = v17
                _ = v18
            } else {
                var v19 int16 = 32767
                fmt.Println(v4)
                if true {
                    for fi20 := 0; fi20 < 6; fi20++ {
                        var v21 float32 = float32((((((1.5 / 1e10) - float64(float32(0.1))) / 1e10) + (3.14159 / 1e10)) + (0.1 / ((1e10 / float64(v15)) - ((0.1 * 1.5) / (-2.25 / -2.25))))))
                        _ = fi20
                        _ = v21
                    }
                    var v22 int16 = (v19 / ((1000) | 1))
                    _ = v22
                } else {
                    for fi23 := 0; fi23 < 10; fi23++ {
                        for fi24 := 0; fi24 < 3; fi24++ {
                            fmt.Println(fi24, v19, v5)
                            var v25 int32 = (-(v5))
                            _ = fi24
                            _ = v25
                        }
                        _ = fi23
                    }
                }
                _ = v19
            }
        }
    }
    var v26 uint64 = 18446744073709551615
    var v27 float32 = ((1.5 + ((v15 - ((float32(1.5) / v15) - float32(float64(3.14159)))) / ((v15 + (((v15 + 1.5) * (1.5 - v15)) * v15)) * v15))) - v15)
    v5 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v7
}
