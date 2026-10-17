package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int32 }
func genFn0(p0 int16, p1 int64) uint32 {
    var v6 uint16 = 2
    v6 = (v6 | v6)
    if ((v6 + v6) > (v6 - v6)) {
        var v7 int8 = 127
        v6 = 32767
        _ = v7
    }
    if (!false) {
        v8 := []int{-486, -2147483648, 154, -426}
        _ = v8
    }
    _ = p0
    _ = p1
    _ = v6
    return 404
}

func main() {
    var v1 int32 = 0
    var v2 uint32 = 2147483648
    var v3 uint16 = 51
    var v4 int16 = -32768
    v5 := make(map[string]int)
    var v9 uint16 = (v3 * v3)
    for fi10 := 0; fi10 < 3; fi10++ {
        switch (-(fi10)) {
        case 2147483647:
            fmt.Println(v9)
        default:
            v4 += 1000
        }
        _ = fi10
    }
    v11 := GenS1{F0: -128, F1: -1}
    v1 += v1
    var v12 bool = true
    fmt.Println(v4)
    var v13 int64 = -121
    switch 127 {
    case 42:
        for fi14 := 0; fi14 < 7; fi14++ {
            for fi15 := 0; fi15 < 20; fi15++ {
                v1 += v1
                _ = fi15
            }
            _ = fi14
        }
    default:
        v16 := []int{32768, -853, 32767}
        _ = v16
    }
    v17 := GenS1{F0: 0, F1: v1}
    fmt.Println(v1, v1)
    switch "" {
    case "\n\t":
        var v18 int8 = 100
        _ = v18
    case "":
        fmt.Println(v13)
    case "a":
        if (!((-79 != v4) && false)) {
            var v19 int = -1
            for fi20 := 0; fi20 < 17; fi20++ {
                switch -9223372036854775808 {
                case -334:
                    v3 -= 255
                case 128:
                    var v21 uint8 = 127
                    _ = v21
                case 4294967295:
                    v12 = (v12 && false)
                }
                v1 -= (v1 ^ v1)
                switch "\n\t" {
                case "foo bar":
                    var v22 int64 = (v13 | int64(fi20))
                    _ = v22
                case "":
                    var v23 uint16 = v3
                    _ = v23
                default:
                    v24 := v5["hello"]
                    _ = v24
                }
                _ = fi20
            }
            _ = v19
        }
    }
    v13 -= (v13 & (v13 / (((v13 % (((^(v13))) | 1))) | 1)))
    v4 = (-(v4))
    v25 := GenS1{F0: -128, F1: int32(v4)}
    for fi26 := 0; fi26 < 15; fi26++ {
        if (2 >= 97) {
            fmt.Println(v9, fi26, fi26)
            var v27 uint8 = 42
            _ = v27
        } else {
            var v28 int32 = (v1 & (-(v1)))
            _ = v28
        }
        _ = fi26
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v17
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
