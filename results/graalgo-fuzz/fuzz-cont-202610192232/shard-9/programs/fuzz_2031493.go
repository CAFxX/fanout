package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8 }
func genFn0(p0 int32, p1 uint16, p2 int32) int64 {
    p2 += (p0 & p0)
    if (("foo bar" + "a") == "a") {
        var v4 uint16 = (^(p1))
        for fi5 := 0; fi5 < 12; fi5++ {
            p2 = p0
            var v6 float32 = float32(1e10)
            _ = fi5
            _ = v6
        }
        _ = v4
    }
    _ = p0
    _ = p1
    _ = p2
    return 680
}

func main() {
    var v1 uint8 = 42
    var v2 int64 = -1
    var v3 uint64 = 18446744073709551615
    v7 := GenS1{F0: 599, F1: 255}
    for fi8 := 0; fi8 < 13; fi8++ {
        v1 -= (v1 | v1)
        v1 += (v1 - (^(v1)))
        _ = fi8
    }
    for fi9 := 0; fi9 < 8; fi9++ {
        var v10 float64 = -2.25
        _ = fi9
        _ = v10
    }
    switch "" {
    case "hello":
        fmt.Println(v3)
    }
    var v11 int32 = 851
    if false {
        v12 := []int{128, 32768}
        var v13 uint8 = (v1 * (v1 | v1))
        _ = v12
        _ = v13
    } else {
        for fi14 := 0; fi14 < 17; fi14++ {
            v15 := GenS1{F0: 18446744073709551615, F1: 2}
            _ = fi14
            _ = v15
        }
        if true {
            fmt.Println(v11)
            if false {
                switch "hello" {
                case "hello":
                    v16 := []int{32767, -784, 255}
                    _ = v16
                case "a":
                    v3 -= 0
                case "":
                    v3 -= (v3 >> ((1) & 63))
                }
            }
        } else {
            var v17 float32 = (3.14159 - 3.14159)
            v18 := []int{127, -382, -1, 42}
            var v19 float32 = (((v17 * (v17 * 0.1)) * ((float32((((1e10 * -2.25) - -2.25) + 1.5)) + (v17 / 1.5)) + v17)) - (3.14159 + v17))
            _ = v17
            _ = v18
            _ = v19
        }
    }
    if (127 != -926) {
        v20 := []int{946, -2147483648}
        var v21 float32 = float32(3.14159)
        _ = v20
        _ = v21
    }
    v22 := GenS1{F0: 255, F1: 2}
    v1 += (^(v1))
    if ((1e10 / 0.1) == (float64(0.0) + -2.25)) {
        var v23 uint32 = 2
        v24 := []int{42}
        v24 = append(v24, 17)
        _ = v23
        _ = v24
    }
    v25 := GenS1{F0: 898, F1: v1}
    fmt.Println(v11, v1, v11)
    v26 := []int{-282, 1}
    if true {
        switch -2147483648 {
        case 4294967295:
            var v27 string = ("foo bar" + "\n\t")
            _ = v27
        case 32768:
            for fi28 := 0; fi28 < 7; fi28++ {
                v2 -= v2
                var v29 int32 = v11
                var v30 uint32 = 717
                _ = fi28
                _ = v29
                _ = v30
            }
        default:
            v31 := GenS1{F0: 18446744073709551615, F1: (v1 + 42)}
            _ = v31
        }
        v11 += v11
        v26 = append(v26, -876)
    } else {
        v11 = (v11 << ((18446744073709551615) & 63))
    }
    var v32 int8 = 17
    _ = len(v26)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    _ = v1
    _ = v11
    _ = v2
    _ = v22
    _ = v25
    _ = v26
    _ = v3
    _ = v32
    _ = v7
}
