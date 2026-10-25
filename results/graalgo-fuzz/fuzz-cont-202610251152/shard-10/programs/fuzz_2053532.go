package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint32; F2 uint8 }
func genFn0(p0 int16) int {
    v6 := GenS1{F0: 42, F1: 0, F2: 26}
    var v7 uint32 = 65535
    _ = p0
    _ = v6
    _ = v7
    return -9223372036854775808
}

func main() {
    v1 := make(map[string]int)
    var v2 uint = 65535
    var v3 int8 = 100
    var v4 int16 = 32767
    var v5 string = ""
    for fi8 := 0; fi8 < 15; fi8++ {
        switch fi8 {
        case 9223372036854775807:
            v3 += (v3 >> ((18446744073709551615) & 63))
        case 74:
            var v9 uint64 = 964
            _ = v9
        case 127:
            for fi10 := 0; fi10 < 19; fi10++ {
                v11 := []int{-1, (fi10 ^ (fi8 ^ (fi10 * fi10))), 127}
                var v12 float64 = float64(((3.14159 + 1e10) / 1e10))
                _ = fi10
                _ = v11
                _ = v12
            }
        }
        _ = fi8
    }
    var v13 int8 = (^(v3))
    switch v5 {
    case "":
        if false {
            v1[v5] = 779
            for fi14 := 0; fi14 < 20; fi14++ {
                v15 := GenS1{F0: 65535, F1: 2147483647, F2: 128}
                if ((953 < v2) && (false || false)) {
                    v16 := GenS1{F0: 65535, F1: 49, F2: 69}
                    v5 = v5
                    _ = v16
                }
                var v17 int8 = v13
                _ = fi14
                _ = v15
                _ = v17
            }
        }
    case "hello":
        v18 := GenS1{F0: (v2 + 809), F1: 1, F2: 1}
        _ = v18
    default:
        v13 = (v13 << ((v2) & 63))
    }
    v3 = v3
    switch "" {
    case "foo bar":
        fmt.Println(v13)
    case "\n\t":
        var v19 uint16 = 32767
        _ = v19
    }
    for fi20 := 0; fi20 < 16; fi20++ {
        var v21 uint16 = 42
        v22 := []int{}
        _ = fi20
        _ = v21
        _ = v22
    }
    v2 += 2
    v5 = v5
    v23 := []int{-886, -9223372036854775808, -9223372036854775808, 42}
    var v24 string = ("" + v5)
    fmt.Println(v4, v13)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
}
