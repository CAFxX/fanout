package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8 }
func genFn0(p0 uint, p1 int16) int32 {
    var v7 uint16 = 42
    var v8 uint8 = 117
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return -615
}

func main() {
    var v1 float32 = ((float32((0.0 - 0.0)) + -2.25) + 0.1)
    var v2 int32 = 265
    var v3 float32 = (((3.14159 + v1) / 0.1) - (-2.25 - (v1 / 3.14159)))
    var v4 uint = 42
    var v5 int = 2147483647
    var v6 int16 = 536
    if (((((!false) || true) || (255 == 1)) || (false || (128 >= 2))) || false) {
        v9 := GenS1{F0: int16(v5), F1: 0}
        var v10 uint64 = 591
        _ = v10
        _ = v9
    }
    v11 := GenS1{F0: v6, F1: 42}
    var v12 uint32 = 65535
    v13 := GenS1{F0: v6, F1: 179}
    fmt.Println(v12, v6)
    switch 0 {
    case -9223372036854775808:
        v14 := GenS1{F0: (v6 | (-(v6))), F1: 254}
        _ = v14
    case 163:
        fmt.Println(v1)
    case 0:
        v2 = (v2 / (((v2 * v2)) | 1))
    default:
        v1 += (v3 * (v3 * v3))
    }
    if ((genFn0(v4, 1) > v2) && true) {
        var v15 uint16 = 2
        _ = v15
    }
    v16 := []int{(v5 % ((v5) | 1)), 584}
    switch "\n\t" {
    case "foo bar":
        if true {
            var v17 uint = (v4 % ((65535) | 1))
            var v18 int8 = 1
            for fi19 := 0; fi19 < 18; fi19++ {
                var v20 bool = true
                switch "" {
                case "hello":
                    for fi21 := 0; fi21 < 2; fi21++ {
                        switch fi21 {
                        case 761:
                            fmt.Println(v20)
                        case 0:
                            var v22 int16 = (v6 >> ((42) & 63))
                            _ = v22
                        case 4294967295:
                            fmt.Println(v5, v1)
                        }
                        _ = fi21
                    }
                }
                _ = fi19
                _ = v20
            }
            _ = v17
            _ = v18
        } else {
            v16 = append(v16, (v5 + (v5 * v5)))
            var v23 uint64 = 2
            _ = v23
        }
    case "\n\t":
        var v24 int8 = -128
        _ = v24
    }
    var v25 bool = true
    if (v25 && v25) {
        v3 = 0.1
        v1 = 0.1
    } else {
        v1 -= (v3 + (v1 + v1))
        var v26 uint = 41
        _ = v26
    }
    var v27 string = "\n\t"
    switch (v27 + ("foo bar" + v27)) {
    case "foo bar":
        fmt.Println(v1, v4)
    default:
        v25 = (!(-810 > 1))
    }
    switch v5 {
    case 32768:
        var v28 float32 = v3
        _ = v28
    case 382:
        v29 := GenS1{F0: 366, F1: 108}
        _ = v29
    case -259:
        switch (v5 / ((v5) | 1)) {
        case 9223372036854775807:
            if true {
                v16 = append(v16, (v5 | (v5 << (((v4 | (v4 >> ((263) & 63)))) & 63))))
            }
        case 409:
            v30 := GenS1{F0: v6, F1: 128}
            _ = v30
        }
    default:
        for fi31 := 0; fi31 < 6; fi31++ {
            if len(v16) >= 2 {
                v16 = v16[1:len(v16)]
            }
            v5 -= 1
            _ = fi31
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v16
    _ = v2
    _ = v25
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
