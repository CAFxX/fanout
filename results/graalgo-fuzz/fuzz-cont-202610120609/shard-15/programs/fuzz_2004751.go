package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32 }
type GenS2 struct { F0 int8; F1 uint16 }

func main() {
    v1 := make(map[int]bool)
    v2 := []uint32{1, 1}
    v3 := make(map[int]int)
    var v4 uint64 = 0
    v2 = append(v2, 7)
    var v5 float64 = 1.5
    var v6 uint32 = 170
    var v7 bool = (-764 <= 0)
    v7 = (-136 > 719)
    for fi8 := 0; fi8 < 7; fi8++ {
        v9 := GenS1{F0: 232, F1: 2147483647}
        switch 2147483647 {
        case -2147483648:
            v4 = (v4 / ((0) | 1))
        case -9223372036854775808:
            v6 += uint32(v4)
        case -933:
            fmt.Println(fi8)
        }
        _ = fi8
        _ = v9
    }
    var v10 uint = 255
    var v11 uint64 = v4
    var v12 int32 = 42
    switch ("" + "") {
    case "a":
        var v13 int32 = v12
        _ = v13
    default:
        if ((v6 <= v6) || v7) {
            for fi14 := 0; fi14 < 14; fi14++ {
                v15 := GenS2{F0: -93, F1: 32767}
                _ = fi14
                _ = v15
            }
        }
    }
    if len(v2) > 0 {
        v2[0] = v6
    }
    switch -177 {
    case -884:
        var v16 int16 = 614
        _ = v16
    case -9223372036854775808:
        v7 = v7
    case 2147483647:
        v10 = 1
    }
    var v17 int8 = 88
    var v18 float64 = v5
    if false {
        for fi19 := 0; fi19 < 14; fi19++ {
            v20 := GenS1{F0: 1, F1: 2147483647}
            _ = fi19
            _ = v20
        }
        switch "" {
        case "hello":
            var v21 uint8 = 7
            _ = v21
        case "a":
            for fi22 := 0; fi22 < 2; fi22++ {
                var v23 float32 = 1.5
                _ = fi22
                _ = v23
            }
        case "\n\t":
            v17 = v17
        default:
            if v7 {
                v10 -= (^(v10))
                v7 = (-109 != 1000)
                var v24 uint = v10
                _ = v24
            }
        }
        var v25 int = -628
        _ = v25
    }
    var v26 string = ("\n\t" + "")
    fmt.Println(v26)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v17
    _ = v18
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
