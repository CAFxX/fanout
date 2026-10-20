package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int16; F2 int64 }
type GenS2 struct { F0 string; F1 uint64; F2 int16 }
type GenS3 struct { F0 uint64; F1 int8; F2 uint }
func genFn0(p0 int8, p1 int, p2 uint8) bool {
    fmt.Println(p2, p0, p0)
    var v4 int32 = -2147483648
    fmt.Println(p0, v4, p1)
    p0 = -26
    v5 := []int{p1, -2147483648, -1}
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return (false || false)
}

func main() {
    var v1 uint8 = 2
    v2 := []string{"a", "a", "\n\t", ""}
    var v3 uint16 = 32767
    v6 := GenS1{F0: 291, F1: -32768, F2: 42}
    _ = len(v2)
    v1 = v1
    var v7 uint = 65535
    if len(v2) > 0 {
        v2[0] = "a"
    }
    switch "a" {
    case "foo bar":
        v3 += v3
    case "":
        if len(v2) > 0 {
            v2[0] = (("hello" + "") + "")
        }
    case "a":
        for fi8 := 0; fi8 < 10; fi8++ {
            for fi9 := 0; fi9 < 1; fi9++ {
                if (v1 != 42) {
                    var v10 float32 = 0.0
                    var v11 uint = v7
                    switch fi9 {
                    case -607:
                        if true {
                            v10 = (0.1 / v10)
                        } else {
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            fmt.Println(v3, fi9)
                        }
                    case -748:
                        var v12 float64 = 1.5
                        _ = v12
                    }
                    _ = v10
                    _ = v11
                }
                _ = fi9
            }
            v1 = 2
            _ = fi8
        }
    }
    if (0.1 > (3.14159 * 1.5)) {
        var v13 int = -978
        v7 -= (v7 + 1)
        _ = v13
    }
    v2 = append(v2, ("\n\t" + "\n\t"))
    v7 += 295
    v14 := GenS1{F0: 0, F1: 1, F2: 0}
    var v15 int = -1
    switch "a" {
    case "\n\t":
        v16 := GenS1{F0: 2, F1: -336, F2: -691}
        _ = v16
    case "a":
        v15 = (v15 << ((0) & 63))
    case "":
        var v17 uint16 = 575
        _ = v17
    }
    if (float64(((3.14159 - 0.0) + float32(1e10))) >= (3.14159 + ((0.0 / 1.0) / ((-2.25 / -2.25) + ((0.0 - 0.1) + 3.14159))))) {
        switch v15 {
        case -467:
            v18 := GenS2{F0: ("" + "\n\t"), F1: 2, F2: -32768}
            _ = v18
        case -670:
            v19 := GenS1{F0: 1, F1: -504, F2: -969}
            _ = v19
        default:
            v20 := GenS3{F0: 2, F1: -40, F2: (v7 % ((18446744073709551615) | 1))}
            _ = v20
        }
    }
    if (v1 < 1) {
        var v21 int16 = 42
        v1 -= 42
        var v22 uint8 = (v1 % ((v1) | 1))
        _ = v21
        _ = v22
    } else {
        if len(v2) > 0 {
            v2[0] = ""
        }
        fmt.Println(v3, v1, v1)
        var v23 int64 = 0
        _ = v23
    }
    v7 = v7
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v6
    _ = v7
}
