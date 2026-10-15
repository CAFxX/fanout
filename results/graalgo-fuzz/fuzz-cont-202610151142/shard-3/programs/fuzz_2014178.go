package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int }
func genFn0(p0 uint64) uint8 {
    switch "" {
    case "a":
        switch ("foo bar" + "hello") {
        case "a":
            v5 := []int{46}
            _ = v5
        case "hello":
            p0 = (^(p0))
        case "foo bar":
            v6 := []int{}
            _ = v6
        default:
            p0 += 42
        }
    default:
        var v7 int8 = -124
        _ = v7
    }
    p0 = (p0 >> ((753) & 63))
    _ = p0
    return 213
}
func genFn1(p0 int64, p1 uint64, p2 uint) uint64 {
    v8 := []int{4294967295, -1}
    if (2 <= 470) {
        if len(v8) > 0 {
            v8[0] = 1
        }
        if (292 < 82) {
            for fi9 := 0; fi9 < 4; fi9++ {
                v10 := GenS1{F0: (p2 ^ p2), F1: fi9}
                _ = fi9
                _ = v10
            }
        } else {
            var v11 uint = 18446744073709551615
            _ = v11
        }
    } else {
        var v12 uint8 = 2
        _ = v12
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return 0
}

func main() {
    v1 := []int{127, 4294967295, -266, -2147483648}
    v2 := []string{}
    var v3 uint = 18446744073709551615
    v4 := make(map[string]string)
    var v13 int32 = -160
    if true {
        v1 = append(v1, -2147483648)
    }
    var v14 uint8 = genFn0(genFn1(1, 4294967295, v3))
    var v15 uint64 = 4294967295
    if len(v2) > 0 {
        v2[0] = "\n\t"
    }
    var v16 int8 = 0
    var v17 int = -1
    var v18 uint = 65535
    v19 := GenS1{F0: v3, F1: (v17 + v17)}
    v14 += (^(v14))
    for fi20 := 0; fi20 < 12; fi20++ {
        v1 = append(v1, (v17 - (v17 / (((fi20 + (v17 & 430))) | 1))))
        _ = fi20
    }
    var v21 int64 = -1
    switch "\n\t" {
    case "hello":
        var v22 int16 = -332
        _ = v22
    case "":
        if false {
            for fi23 := 0; fi23 < 12; fi23++ {
                var v24 uint16 = 255
                _ = len(v4)
                v25 := GenS1{F0: (v18 * v3), F1: (v17 + v17)}
                _ = fi23
                _ = v24
                _ = v25
            }
        }
    case "a":
        v26 := GenS1{F0: v3, F1: (v17 + v17)}
        _ = v26
    }
    _ = len(v4)
    var v27 float32 = float32(0.0)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v21
    _ = v27
    _ = v3
    _ = v4
}
