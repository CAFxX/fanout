package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int; F1 int16; F2 uint8 }
type GenS4 struct { F0 string }
func genFn0(p0 int64, p1 int, p2 int) int8 {
    for fi4 := 0; fi4 < 18; fi4++ {
        var v5 int = -973
        p0 += p0
        var v6 uint16 = 2
        _ = fi4
        _ = v5
        _ = v6
    }
    switch ("" + ("\n\t" + "\n\t")) {
    case "":
        p2 -= p2
    case "\n\t":
        v7 := []int{p1}
        _ = v7
    }
    _ = p0
    _ = p1
    _ = p2
    return -54
}
func genFn1(p0 int) uint8 {
    fmt.Println(p0)
    p0 += p0
    v8 := []int{(p0 / ((p0) | 1)), p0}
    var v9 int32 = -768
    p0 = p0
    _ = p0
    _ = v8
    _ = v9
    return 1
}

func main() {
    var v1 uint64 = 42
    var v2 uint32 = 2147483648
    var v3 int8 = 27
    var v10 float64 = ((((1e10 * 1.5) / float64(1.5)) + 1.5) / (((-2.25 - float64((0.0 + float32(3.14159)))) - (((3.14159 * (float64((1.5 - 1.5)) * 1e10)) + 0.0) * 1e10)) / (1.5 - 1e10)))
    for fi11 := 0; fi11 < 5; fi11++ {
        switch (fi11 ^ (fi11 + (fi11 >> ((944) & 63)))) {
        case 1:
            fmt.Println(fi11, v3, v3)
        case 9223372036854775807:
            v3 -= genFn0(42, (^(fi11)), fi11)
        default:
            if (v1 > v1) {
                if (genFn0(-420, fi11, -543) <= (v3 << ((1) & 63))) {
                    v3 += v3
                }
            } else {
                v12 := GenS1{F0: 255}
                _ = v12
            }
        }
        switch int(v2) {
        case 42:
            var v13 float32 = 3.14159
            _ = v13
        }
        _ = fi11
    }
    var v14 uint = 82
    if true {
        var v15 uint32 = (v2 >> (((^(v14))) & 63))
        v2 = v15
        var v16 string = "foo bar"
        _ = v15
        _ = v16
    }
    v17 := []int{0, 9223372036854775807, 35, 442}
    v10 = (0.0 * (-2.25 + (((((3.14159 - v10) / v10) / ((0.1 + 0.1) + (v10 / -2.25))) * v10) * v10)))
    var v18 uint32 = 1
    if false {
        fmt.Println(v1, v3, v14)
        v19 := GenS2{F0: 4294967295}
        _ = v19
    } else {
        if true {
            v1 -= v1
            var v20 int16 = -1
            v1 = (v1 << ((v14) & 63))
            _ = v20
        }
        v21 := GenS3{F0: int(v1), F1: 783, F2: 0}
        v22 := GenS4{F0: "foo bar"}
        _ = v21
        _ = v22
    }
    v14 += v14
    var v23 float32 = 0.1
    var v24 uint32 = (v18 + (v2 ^ 2))
    for fi25 := 0; fi25 < 9; fi25++ {
        fmt.Println(v10)
        _ = fi25
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v14
    _ = v17
    _ = v18
    _ = v2
    _ = v23
    _ = v24
    _ = v3
}
