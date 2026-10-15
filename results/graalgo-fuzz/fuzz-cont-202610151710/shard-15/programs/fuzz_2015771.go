package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int16 }

func main() {
    var v1 float64 = 1e10
    var v2 bool = (18446744073709551615 > 2)
    var v3 int32 = 565
    if false {
        var v4 int16 = 1
        switch 127 {
        case -722:
            v5 := []int{254, 32768, -709, 1}
            _ = v5
        case -2147483648:
            fmt.Println(v2, v1, v2)
        }
        _ = v4
    } else {
        v6 := []int{-482, -474, 2147483647}
        _ = v6
    }
    v1 = ((v1 * (float64((1e10 + 0.0)) - v1)) * float64((3.14159 * 1.5)))
    v7 := GenS1{F0: 32767}
    for fi8 := 0; fi8 < 12; fi8++ {
        for fi9 := 0; fi9 < 14; fi9++ {
            v10 := []int{}
            _ = fi9
            _ = v10
        }
        if false {
            v11 := GenS1{F0: 1000}
            var v12 float32 = float32(((3.14159 * -2.25) * 0.0))
            v2 = ((v12 > v12) || (-531 <= 32767))
            _ = v11
            _ = v12
        } else {
            for fi13 := 0; fi13 < 8; fi13++ {
                v14 := []int{657}
                _ = fi13
                _ = v14
            }
        }
        v15 := GenS1{F0: 32767}
        _ = fi8
        _ = v15
    }
    var v16 uint64 = 977
    v17 := GenS2{F0: 1}
    for fi18 := 0; fi18 < 8; fi18++ {
        var v19 uint = 309
        v2 = (142 != 2)
        _ = fi18
        _ = v19
    }
    v20 := GenS2{F0: 1}
    v21 := []int{-2147483648, -295, 0, 142}
    var v22 int32 = (^(v3))
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v7
}
