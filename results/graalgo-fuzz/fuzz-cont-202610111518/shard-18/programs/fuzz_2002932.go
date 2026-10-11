package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint32 }
type GenS2 struct { F0 string; F1 uint }

func main() {
    var v1 int16 = 0
    v2 := make(map[int]int)
    var v3 float32 = ((0.1 - 1.5) - 1e10)
    v4 := make(map[int]string)
    var v5 int16 = (^(v1))
    v6 := []int{32768, 523, 127, 47}
    switch -9223372036854775808 {
    case 4294967295:
        if len(v6) > 0 {
            v6[0] = 818
        }
    case -612:
        var v7 uint64 = 2
        _ = v7
    }
    v3 = v3
    if (1 >= 1) {
        if true {
            v3 -= 1.5
        } else {
            v3 = v3
        }
        fmt.Println(v3, v5, v1)
        switch 4294967295 {
        case -753:
            v1 -= 42
        default:
            fmt.Println(v1, v1, v3)
        }
    } else {
        _ = len(v6)
        for fi8 := 0; fi8 < 3; fi8++ {
            for fi9 := 0; fi9 < 19; fi9++ {
                _ = len(v6)
                _ = fi9
            }
            var v10 int = -674
            _ = fi8
            _ = v10
        }
    }
    var v11 bool = true
    switch "a" {
    case "a":
        if true {
            fmt.Println(v5)
            fmt.Println(v1, v3, v1)
            v1 += v5
        }
    case "":
        var v12 int32 = 0
        _ = v12
    case "foo bar":
        var v13 float64 = 0.0
        _ = v13
    default:
        v3 = (v3 - v3)
    }
    v5 += (v1 / (((v1 % ((v1) | 1))) | 1))
    if ((("a" + "\n\t") + "a") < ("a" + "a")) {
        v14 := GenS1{F0: 42, F1: 755}
        v11 = true
        _ = v14
    } else {
        v11 = v11
        if (!v11) {
            if v11 {
                v5 -= (v5 ^ 0)
                v11 = false
            }
            v15 := GenS2{F0: "\n\t", F1: 1}
            var v16 int = 2147483647
            _ = v15
            _ = v16
        } else {
            v3 = (v3 * (float32(3.14159) - (v3 / v3)))
            v5 = v1
        }
        v6 = append(v6, 142)
    }
    switch 4294967295 {
    case 447:
        if v11 {
            fmt.Println(v11)
            _ = len(v6)
            var v17 uint16 = 65535
            _ = v17
        }
    case 0:
        var v18 int = 2147483647
        _ = v18
    case -2147483648:
        for fi19 := 0; fi19 < 10; fi19++ {
            fmt.Println(v5, v1)
            _ = fi19
        }
    default:
        v3 = 3.14159
    }
    switch -419 {
    case 42:
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
