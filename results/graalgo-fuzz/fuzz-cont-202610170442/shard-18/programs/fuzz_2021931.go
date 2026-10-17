package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32 }
type GenS2 struct { F0 uint32; F1 int }

func main() {
    var v1 int16 = 32767
    v2 := make(map[uint32]uint32)
    v3 := []uint32{42, 1, 90}
    v4 := []uint32{4294967295, 42, 454, 2147483648}
    v5 := []int{32767, -403, -9223372036854775808}
    var v6 int8 = -128
    _ = len(v5)
    var v7 float64 = (3.14159 - float64(0.0))
    v7 -= v7
    switch 4294967295 {
    case -1:
        for fi8 := 0; fi8 < 5; fi8++ {
            fmt.Println(fi8)
            if true {
                if true {
                    v6 = v6
                }
            }
            _ = fi8
        }
    default:
        v4 = append(v4, 601)
    }
    if (!true) {
        var v9 int8 = v6
        if true {
            if false {
                switch 764 {
                case 4294967295:
                    for fi10 := 0; fi10 < 3; fi10++ {
                        var v11 int = fi10
                        v7 += (v7 * (v7 + 1e10))
                        _ = fi10
                        _ = v11
                    }
                case 0:
                    if len(v5) > 0 {
                        v5[0] = int(v1)
                    }
                }
                fmt.Println(v7, v6)
                v12 := GenS1{F0: 0, F1: 2147483647}
                _ = v12
            } else {
                switch "\n\t" {
                case "foo bar":
                    fmt.Println(v9)
                default:
                    fmt.Println(v1)
                }
            }
        }
        _ = v9
    } else {
        if false {
            var v13 int8 = 127
            var v14 float32 = -2.25
            fmt.Println(v7, v7)
            _ = v13
            _ = v14
        }
    }
    if (false && false) {
        var v15 int16 = v1
        _ = v15
    }
    v4 = append(v4, 935)
    v6 = v6
    v16 := GenS2{F0: 4294967295, F1: -1}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
