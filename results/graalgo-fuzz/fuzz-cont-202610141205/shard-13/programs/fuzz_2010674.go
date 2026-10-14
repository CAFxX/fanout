package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint16 }
type GenS2 struct { F0 int32; F1 uint64 }
func genFn0() string {
    switch 2147483647 {
    case 2147483647:
        var v6 int64 = -950
        _ = v6
    case 42:
        v7 := GenS1{F0: 249, F1: 2}
        _ = v7
    case 128:
        var v8 int64 = -498
        _ = v8
    default:
        var v9 int8 = 1
        _ = v9
    }
    var v10 int32 = 1
    _ = v10
    return ("foo bar" + "hello")
}

func main() {
    var v1 uint16 = 2
    var v2 uint = 65535
    var v3 string = "\n\t"
    v4 := []int{-9223372036854775808}
    v5 := make(map[string]string)
    v1 -= (v1 & (v1 * v1))
    var v11 int8 = -1
    switch int(v11) {
    case -1:
        v2 += (v2 * 1)
    default:
        v2 = v2
    }
    if ((false || ((42 != (v1 ^ (v1 + v1))) && true)) && (false || true)) {
        for fi12 := 0; fi12 < 7; fi12++ {
            v3 = v3
            for fi13 := 0; fi13 < 7; fi13++ {
                var v14 bool = (v11 > v11)
                var v15 int16 = 0
                v16 := GenS1{F0: 270, F1: (v1 ^ 42)}
                _ = fi13
                _ = v14
                _ = v15
                _ = v16
            }
            _ = fi12
        }
        v11 -= v11
        switch -1 {
        case 1:
            v17 := GenS2{F0: 1, F1: 18446744073709551615}
            _ = v17
        case 9223372036854775807:
            v18 := GenS1{F0: 2147483647, F1: (v1 - 255)}
            _ = v18
        default:
            var v19 uint32 = 2147483647
            _ = v19
        }
    }
    v20 := GenS2{F0: 2147483647, F1: 2}
    if (1e10 < 0.1) {
        var v21 uint16 = (v1 + (v1 + (v1 / (((v1 - (^(v1)))) | 1))))
        _ = v21
    } else {
        var v22 int8 = v11
        v4 = append(v4, 1)
        if (0.1 < float64(1e10)) {
            v3 = "a"
        }
        _ = v22
    }
    for fi23 := 0; fi23 < 16; fi23++ {
        v2 = v2
        _ = fi23
    }
    switch "" {
    case "hello":
        var v24 int8 = v11
        _ = v24
    case "foo bar":
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    default:
        fmt.Println(v3, v1, v3)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v11
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
