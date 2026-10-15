package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool }

func main() {
    var v1 string = "a"
    v2 := make(map[int]bool)
    v3 := []uint32{}
    var v4 string = v1
    var v5 uint32 = 461
    v6 := make(map[int]int)
    switch -2147483648 {
    case 2147483647:
        v5 -= v5
    default:
        var v7 int32 = -283
        _ = v7
    }
    for fi8 := 0; fi8 < 2; fi8++ {
        v9 := GenS1{F0: 255, F1: false}
        v6[-783] = (fi8 - 1)
        v3 = append(v3, (^(v5)))
        _ = fi8
        _ = v9
    }
    v5 += (v5 / (((v5 | 349)) | 1))
    if false {
        for fi10 := 0; fi10 < 6; fi10++ {
            for fi11 := 0; fi11 < 13; fi11++ {
                fmt.Println(v1, fi10, fi11)
                _ = fi11
            }
            _ = fi10
        }
        v4 = v1
        v5 = (v5 ^ (^(v5)))
    } else {
        v1 = (("a" + ("a" + v4)) + v1)
        v12 := v2[32768]
        v13 := v6[-9223372036854775808]
        _ = v12
        _ = v13
    }
    var v14 int32 = 1
    v3 = append(v3, v5)
    if ((v14 * (-(v14))) < 0) {
        var v15 float32 = -2.25
        _ = v15
    }
    v16 := v2[127]
    v17 := v2[1]
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v14
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
