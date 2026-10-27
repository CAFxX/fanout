package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32 }
func genFn0(p0 string, p1 int16, p2 int16) string {
    var v6 int = 473
    var v7 int64 = -1
    switch (^(v6)) {
    case 32767:
        p2 = (p1 | p2)
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return "hello"
}

func main() {
    var v1 uint8 = 2
    v2 := []uint32{4294967295, 514, 4294967295, 1}
    v3 := make(map[string]bool)
    var v4 int16 = -32768
    v5 := []string{"foo bar", ("" + "\n\t")}
    for fi8 := 0; fi8 < 12; fi8++ {
        var v9 float32 = ((0.0 - 3.14159) - 1.5)
        var v10 int32 = 16
        _ = fi8
        _ = v10
        _ = v9
    }
    v1 += v1
    v1 = 128
    switch -684 {
    case 32768:
        v2 = append(v2, 0)
    case -922:
        switch "a" {
        case "a":
            if false {
                fmt.Println(v1)
                if false {
                    v4 -= (^(v4))
                    v4 -= (v4 & (v4 * v4))
                }
            }
        default:
            switch 32768 {
            case -457:
                switch 32768 {
                case -9223372036854775808:
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                case 0:
                    fmt.Println(v1)
                case 32768:
                    v5 = append(v5, "hello")
                default:
                    var v11 uint8 = 1
                    _ = v11
                }
            case -2147483648:
                var v12 int = 1
                _ = v12
            case 247:
                var v13 bool = (false || false)
                _ = v13
            default:
                if false {
                    v4 += 32767
                    v4 -= (v4 ^ (v4 - v4))
                }
            }
        }
    default:
        v14 := GenS1{F0: 0, F1: 2}
        _ = v14
    }
    var v15 uint16 = 2
    var v16 int = -1
    if (!(((v1 << ((4294967295) & 63)) < v1) || (-128 > 127))) {
        v1 += v1
    } else {
        fmt.Println(v16, v16)
        switch -340 {
        case -171:
            var v17 int16 = (v4 ^ v4)
            _ = v17
        case -153:
            var v18 string = "hello"
            _ = v18
        }
    }
    v2 = append(v2, 0)
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
