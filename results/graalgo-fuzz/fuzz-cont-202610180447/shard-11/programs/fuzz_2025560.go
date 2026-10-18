package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint; F1 int8; F2 uint }

func main() {
    var v1 int16 = -350
    var v2 bool = (719 < 199)
    var v3 int8 = -59
    for fi4 := 0; fi4 < 10; fi4++ {
        for fi5 := 0; fi5 < 7; fi5++ {
            if (-2147483648 != -243) {
                var v6 float64 = ((0.0 + 1.5) - 3.14159)
                v7 := GenS1{F0: (v3 << ((4294967295) & 63))}
                _ = v6
                _ = v7
            }
            fmt.Println(v2, v1)
            fmt.Println(v2)
            _ = fi5
        }
        for fi8 := 0; fi8 < 2; fi8++ {
            v9 := []int{fi4, fi4, 127, (fi4 * (fi8 | (fi8 / ((-2147483648) | 1))))}
            _ = fi8
            _ = v9
        }
        v10 := GenS1{F0: 127}
        _ = fi4
        _ = v10
    }
    fmt.Println(v2, v1)
    v11 := GenS2{F0: (-301 == 9223372036854775807)}
    var v12 bool = (!(!v2))
    var v13 int32 = 0
    v14 := []int{-696, int(v1)}
    fmt.Println(v2, v1, v3)
    var v15 uint16 = 675
    for fi16 := 0; fi16 < 4; fi16++ {
        v15 = v15
        v17 := GenS3{F0: uint(v15), F1: v3, F2: 255}
        fmt.Println(v1, v13)
        _ = fi16
        _ = v17
    }
    fmt.Println(v2, v15, v2)
    for fi18 := 0; fi18 < 5; fi18++ {
        v12 = (!false)
        _ = fi18
    }
    var v19 float32 = 1e10
    v1 = (v1 / (((^(v1))) | 1))
    fmt.Println(v13, v19, v3)
    var v20 int16 = v1
    switch "\n\t" {
    case "hello":
        v12 = ((^(v3)) != 1)
    case "a":
        v14 = append(v14, 32767)
    }
    if len(v14) > 0 {
        v14[0] = 4294967295
    }
    switch -198 {
    case -9223372036854775808:
        v20 = (v1 * (v1 >> ((155) & 63)))
    default:
        fmt.Println(v2, v13)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v15:", v15)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v19
    _ = v2
    _ = v20
    _ = v3
}
