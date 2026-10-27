package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint8; F2 uint }
type GenS2 struct { F0 uint16 }
func genFn0() int8 {
    switch "" {
    case "hello":
        if true {
            var v7 float32 = (3.14159 / 3.14159)
            var v8 uint32 = 1
            _ = v7
            _ = v8
        } else {
            v9 := []int{537}
            var v10 int = -2147483648
            var v11 string = ("\n\t" + ("\n\t" + "hello"))
            _ = v10
            _ = v11
            _ = v9
        }
    }
    var v12 uint = 65535
    for fi13 := 0; fi13 < 13; fi13++ {
        fmt.Println(fi13, fi13)
        _ = fi13
    }
    _ = v12
    return 34
}

func main() {
    v1 := make(map[int]bool)
    v2 := []string{"", "a", "\n\t", ((("\n\t" + "\n\t") + "\n\t") + "hello")}
    v3 := make(map[int]int)
    v4 := []string{"foo bar"}
    var v5 uint64 = 18446744073709551615
    var v6 float32 = 1e10
    switch "\n\t" {
    case "hello":
        v6 = v6
    case "\n\t":
        v4 = append(v4, "")
    case "a":
        v4 = append(v4, "hello")
    }
    if (65535 < 0) {
        for fi14 := 0; fi14 < 14; fi14++ {
            var v15 int32 = -1
            v6 = v6
            var v16 int64 = -9223372036854775808
            _ = fi14
            _ = v15
            _ = v16
        }
        var v17 int = 9223372036854775807
        _ = v17
    }
    v6 += float32(-2.25)
    var v18 int32 = 42
    v19 := GenS1{F0: 2147483648, F1: 127, F2: 182}
    switch -897 {
    case 32767:
        v5 -= v5
    case 42:
        var v20 bool = (32767 <= 42)
        _ = v20
    }
    var v21 int64 = 415
    for fi22 := 0; fi22 < 7; fi22++ {
        var v23 float64 = (-2.25 - (-2.25 * 3.14159))
        switch 546 {
        case 4294967295:
            v18 += (v18 * (-(v18)))
        case -9223372036854775808:
            fmt.Println(fi22, v18)
        default:
            v24 := GenS1{F0: 0, F1: 128, F2: 255}
            _ = v24
        }
        v23 += ((0.1 + (v23 * 1.5)) * 1.5)
        _ = fi22
        _ = v23
    }
    if (!false) {
        var v25 string = "a"
        _ = v25
    } else {
        v5 = v5
    }
    for fi26 := 0; fi26 < 13; fi26++ {
        var v27 int8 = -104
        if len(v2) > 0 {
            v2[0] = ""
        }
        _ = fi26
        _ = v27
    }
    if ((float64(v6) + ((0.0 * 1e10) / 1.0)) < 3.14159) {
        var v28 uint16 = 326
        for fi29 := 0; fi29 < 12; fi29++ {
            v30 := GenS1{F0: 2147483647, F1: 194, F2: 42}
            _ = fi29
            _ = v30
        }
        switch "" {
        case "hello":
            switch 560 {
            case 2147483647:
                var v31 float64 = -2.25
                _ = v31
            case 0:
                switch "hello" {
                case "hello":
                    for fi32 := 0; fi32 < 17; fi32++ {
                        var v33 int64 = v21
                        _ = fi32
                        _ = v33
                    }
                default:
                    v34 := GenS2{F0: v28}
                    _ = v34
                }
            case 32767:
                var v35 int32 = v18
                _ = v35
            }
        case "\n\t":
            fmt.Println(v21)
        case "a":
            fmt.Println(v28, v6)
        }
        _ = v28
    }
    v5 -= v5
    v6 = (0.0 + 1.5)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
