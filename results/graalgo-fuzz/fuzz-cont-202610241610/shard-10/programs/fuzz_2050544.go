package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int }

func main() {
    v1 := make(map[uint32]string)
    var v2 uint16 = 32767
    var v3 float32 = 1.5
    var v4 int16 = 1
    v5 := make(map[int]bool)
    v6 := []int{9223372036854775807, -360}
    v2 += (v2 & 255)
    delete(v1, 176)
    var v7 int32 = 2147483647
    v5[int(v7)] = false
    fmt.Println(v4)
    v5[-2147483648] = (1 < 72)
    v8 := GenS1{F0: 0, F1: 32767}
    var v9 uint32 = 1
    var v10 uint32 = (v9 * (v9 % ((759) | 1)))
    for fi11 := 0; fi11 < 5; fi11++ {
        var v12 int16 = 1000
        v10 += v9
        for fi13 := 0; fi13 < 11; fi13++ {
            if (!true) {
                v14 := v5[(-(fi13))]
                _ = v14
            }
            switch -9223372036854775808 {
            case 9223372036854775807:
                v15 := GenS1{F0: uint32(v2), F1: (^(fi11))}
                v1[(^(v10))] = "foo bar"
                _ = v15
            case -535:
                v2 = 255
            case 42:
                if (float32(0.0) <= v3) {
                    v7 -= v7
                    v16 := GenS1{F0: (^(v9)), F1: 127}
                    _ = v16
                } else {
                    v17 := GenS1{F0: 263, F1: (^(fi13))}
                    v5[4294967295] = ("\n\t" <= "a")
                    switch 1 {
                    case 0:
                        var v18 int16 = (v12 / ((1) | 1))
                        _ = v18
                    case 522:
                        var v19 uint16 = 2
                        _ = v19
                    case 127:
                        _ = len(v5)
                    }
                    _ = v17
                }
            default:
                switch "hello" {
                case "a":
                    if (!false) {
                        if len(v6) > 0 {
                            v6[0] = (fi13 * (fi11 * -104))
                        }
                        var v20 uint = 457
                        var v21 int8 = 0
                        _ = v20
                        _ = v21
                    } else {
                        v22 := GenS1{F0: (v9 % ((v9) | 1)), F1: (fi13 - (fi11 ^ fi11))}
                        _ = v22
                    }
                default:
                    v7 += -766
                }
            }
            v6 = append(v6, fi13)
            _ = fi13
        }
        _ = fi11
        _ = v12
    }
    v5[32767] = true
    var v23 int32 = (v7 ^ v7)
    for fi24 := 0; fi24 < 16; fi24++ {
        _ = len(v5)
        _ = fi24
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
