package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int8; F1 int; F2 uint64 }
func genFn0() int64 {
    v4 := []int{-599, 32768, 55, 178}
    var v5 uint32 = 933
    _ = v4
    _ = v5
    return 83
}

func main() {
    var v1 uint16 = 528
    v2 := make(map[string]int)
    var v3 int32 = 42
    v3 = (v3 & (v3 << ((4294967295) & 63)))
    v3 = 2147483647
    v6 := GenS1{F0: -87}
    v2[("foo bar" + ("hello" + "hello"))] = 729
    v7 := GenS1{F0: -128}
    for fi8 := 0; fi8 < 6; fi8++ {
        fmt.Println(v3, v1)
        for fi9 := 0; fi9 < 4; fi9++ {
            v1 += (v1 << ((4294967295) & 63))
            var v10 int32 = -1
            _ = fi9
            _ = v10
        }
        _ = fi8
    }
    fmt.Println(v1)
    var v11 int32 = (v3 ^ -1)
    v1 = 255
    if false {
        var v12 int8 = -128
        var v13 uint16 = (v1 & 42)
        _ = v12
        _ = v13
    } else {
        if (uint64(v1) >= 0) {
            var v14 float64 = 0.0
            switch -2147483648 {
            case 42:
                for fi15 := 0; fi15 < 11; fi15++ {
                    fmt.Println(fi15)
                    var v16 int8 = 0
                    _ = fi15
                    _ = v16
                }
            case 127:
                var v17 int = 186
                _ = v17
            case 0:
                switch "a" {
                case "\n\t":
                    fmt.Println(v3)
                case "":
                    var v18 float64 = (((v14 - v14) * ((v14 * 0.1) - v14)) - (v14 - v14))
                    _ = v18
                case "hello":
                    v2[("a" + "")] = -2147483648
                default:
                    var v19 int8 = 100
                    _ = v19
                }
            default:
                v20 := GenS2{F0: 42, F1: -771, F2: 18446744073709551615}
                _ = v20
            }
            fmt.Println(v11)
            _ = v14
        } else {
            v21 := GenS2{F0: -128, F1: -9223372036854775808, F2: 0}
            _ = v21
        }
        v11 += -972
    }
    var v22 int8 = 42
    v11 += v3
    fmt.Println(v22, v22)
    var v23 int = 1
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v11
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v6
    _ = v7
}
