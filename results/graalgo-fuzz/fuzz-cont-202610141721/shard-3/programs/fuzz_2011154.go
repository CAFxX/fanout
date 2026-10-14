package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string }
type GenS2 struct { F0 bool }
func genFn0(p0 bool, p1 int32) int64 {
    var v5 int8 = 100
    for fi6 := 0; fi6 < 19; fi6++ {
        var v7 uint16 = 677
        v7 -= v7
        v8 := GenS1{F0: v5, F1: ("hello" + "")}
        _ = fi6
        _ = v7
        _ = v8
    }
    p1 -= (p1 << ((1) & 63))
    if p0 {
        switch 2147483647 {
        case 1:
            for fi9 := 0; fi9 < 17; fi9++ {
                var v10 int64 = 394
                v11 := []int{fi9}
                _ = fi9
                _ = v10
                _ = v11
            }
        }
        fmt.Println(p1, p1)
        v5 = v5
    }
    _ = p0
    _ = p1
    _ = v5
    return 118
}

func main() {
    var v1 bool = (4294967295 > 227)
    var v2 uint16 = 42
    v3 := []uint32{}
    var v4 uint = 18446744073709551615
    v2 -= 1
    var v12 uint16 = v2
    if v1 {
        for fi13 := 0; fi13 < 4; fi13++ {
            switch ("hello" + ("hello" + ("foo bar" + "hello"))) {
            case "\n\t":
                v4 = uint(v12)
            case "a":
                for fi14 := 0; fi14 < 13; fi14++ {
                    if len(v3) > 0 {
                        v3[0] = 1
                    }
                    var v15 uint16 = v12
                    fmt.Println(fi13, v4, v4)
                    _ = fi14
                    _ = v15
                }
            default:
                if true {
                    var v16 uint32 = 399
                    v17 := GenS1{F0: -128, F1: (("hello" + "") + "hello")}
                    var v18 int32 = 1
                    _ = v16
                    _ = v17
                    _ = v18
                }
            }
            var v19 int64 = 9223372036854775807
            _ = fi13
            _ = v19
        }
    }
    fmt.Println(v4)
    switch "a" {
    case "":
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    case "a":
        for fi20 := 0; fi20 < 2; fi20++ {
            v2 = (^(v2))
            v1 = (-1 == -1)
            _ = fi20
        }
    }
    switch ("" + "\n\t") {
    case "a":
        var v21 uint16 = 734
        _ = v21
    case "\n\t":
        fmt.Println(v1, v4, v4)
    }
    for fi22 := 0; fi22 < 17; fi22++ {
        var v23 float32 = (-2.25 * 1e10)
        v24 := GenS2{F0: (v1 && (false && v1))}
        _ = fi22
        _ = v23
        _ = v24
    }
    if (!false) {
        fmt.Println(v1, v4)
    } else {
        fmt.Println(v2, v12)
        var v25 int16 = 1000
        _ = v25
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v4 += (v4 / ((255) | 1))
    var v26 int32 = -1
    v1 = v1
    var v27 int = 439
    v3 = append(v3, 4294967295)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v4
}
