package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32; F2 bool }
func genFn0() int8 {
    var v4 int8 = -1
    v5 := []int{0, 832, -205, 32768}
    _ = v4
    _ = v5
    return -56
}
func genFn1(p0 uint8) int32 {
    var v6 int = 210
    var v7 uint16 = 32767
    v8 := []int{(v6 << ((2) & 63)), (^(v6))}
    if true {
        switch "a" {
        case "a":
            var v9 int64 = 9223372036854775807
            _ = v9
        default:
            var v10 bool = false
            _ = v10
        }
        p0 = 109
    }
    var v11 int8 = -128
    _ = p0
    _ = v11
    _ = v6
    _ = v7
    _ = v8
    return 67
}

func main() {
    v1 := make(map[int]string)
    var v2 int8 = 100
    var v3 uint16 = 2
    v1[721] = ("foo bar" + ("a" + "\n\t"))
    v3 = (v3 / (((v3 << ((255) & 63))) | 1))
    v12 := v1[9223372036854775807]
    v13 := GenS1{F0: 1, F1: 0, F2: (!false)}
    v2 += (v2 + v2)
    for fi14 := 0; fi14 < 19; fi14++ {
        for fi15 := 0; fi15 < 18; fi15++ {
            v16 := GenS1{F0: (fi15 * fi15), F1: 65535, F2: true}
            _ = fi15
            _ = v16
        }
        _ = fi14
    }
    if (1 > -32768) {
        var v17 uint32 = 2147483647
        _ = v17
    } else {
        v18 := []int{4294967295}
        var v19 int32 = 878
        _ = v18
        _ = v19
    }
    v2 = v2
    var v20 int8 = -128
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v20
    _ = v3
}
