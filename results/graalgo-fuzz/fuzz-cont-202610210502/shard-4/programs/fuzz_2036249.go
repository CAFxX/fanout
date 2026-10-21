package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int64 }
type GenS2 struct { F0 int16; F1 int16 }

func main() {
    v1 := []int{32768, -9223372036854775808, 345, 4294967295}
    var v2 bool = (-78 > -9223372036854775808)
    var v3 uint16 = 516
    var v4 string = "\n\t"
    switch 747 {
    case 758:
        v5 := GenS1{F0: 100, F1: int64(v3)}
        _ = v5
    case -844:
        if ((v4 + v4) <= "") {
            v6 := GenS1{F0: 1, F1: 441}
            _ = v6
        } else {
            v3 = v3
        }
    default:
        for fi7 := 0; fi7 < 7; fi7++ {
            var v8 uint = 346
            var v9 int32 = 42
            v10 := GenS1{F0: 0, F1: 524}
            _ = fi7
            _ = v10
            _ = v8
            _ = v9
        }
    }
    var v11 float32 = 3.14159
    var v12 int64 = 892
    v12 += v12
    var v13 int16 = 0
    switch (v4 + "\n\t") {
    case "hello":
        var v14 int64 = 42
        _ = v14
    case "":
        fmt.Println(v12, v3, v11)
    default:
        v15 := GenS1{F0: 100, F1: v12}
        _ = v15
    }
    var v16 int16 = v13
    if false {
        if false {
            v17 := GenS2{F0: 1000, F1: v13}
            if ((v12 * (v12 - (v12 - v12))) >= 145) {
                v16 -= v13
                v2 = (v11 < (0.0 / v11))
                if len(v1) > 0 {
                    v1[0] = 0
                }
            }
            _ = v17
        }
    } else {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    }
    if v2 {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        var v18 int16 = (-(v13))
        _ = v18
    } else {
        switch ((v4 + v4) + v4) {
        case "foo bar":
            v11 -= v11
        case "hello":
            var v19 int16 = v13
            _ = v19
        case "":
            v2 = (0.1 > (((-2.25 * 0.1) * 3.14159) / 1.0))
        default:
            fmt.Println(v16, v3)
        }
        v13 -= v16
    }
    switch -896 {
    case 4294967295:
        v16 = (v13 + 1000)
    case 884:
        var v20 uint32 = 0
        _ = v20
    case 32767:
        var v21 int = 128
        _ = v21
    }
    for fi22 := 0; fi22 < 18; fi22++ {
        fmt.Println(v4, v12, v3)
        _ = fi22
    }
    v3 -= v3
    v12 += v12
    v23 := GenS2{F0: (-(v13)), F1: (-(v16))}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v16
    _ = v2
    _ = v23
    _ = v3
    _ = v4
}
