package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int32 }
func genFn0(p0 string, p1 uint8, p2 bool) uint8 {
    if (42 < 32768) {
        switch 708 {
        case -84:
            fmt.Println(p1)
        case -736:
            var v6 string = p0
            _ = v6
        case 2147483647:
            p0 = p0
        default:
            p1 = p1
        }
        var v7 uint32 = 492
        _ = v7
    } else {
        p2 = p2
        v8 := GenS1{F0: p2, F1: 865}
        v9 := GenS2{F0: 127}
        _ = v8
        _ = v9
    }
    var v10 int64 = 1
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    return (p1 - (p1 >> ((0) & 63)))
}
func genFn1(p0 int32, p1 uint) int16 {
    var v11 uint16 = 1
    var v12 int16 = -32768
    var v13 int8 = 100
    var v14 float64 = (((1.5 * (((3.14159 * float64(0.0)) / 0.1) / 1e10)) * ((-2.25 / 1.5) + ((-2.25 + ((((0.1 / 3.14159) / float64(1.5)) / 1e10) + (1.5 / 0.1))) / (1.5 / 1.5)))) + (3.14159 / 1.0))
    switch 128 {
    case 42:
        v15 := []int{128, 32767}
        _ = v15
    default:
        var v16 uint8 = 127
        _ = v16
    }
    _ = p0
    _ = p1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    return (-(v12))
}

func main() {
    v1 := make(map[string]bool)
    v2 := make(map[uint32]int)
    v3 := make(map[uint32]string)
    var v4 int = 567
    var v5 uint16 = 143
    if (4294967295 > 2) {
        v5 = v5
        v4 += v4
        var v17 int64 = -494
        _ = v17
    } else {
        fmt.Println(v5)
    }
    switch v4 {
    case -1:
        var v18 int32 = 42
        _ = v18
    case -9223372036854775808:
        var v19 uint8 = 127
        _ = v19
    default:
        fmt.Println(v5, v5)
    }
    v5 = (v5 | (v5 | v5))
    v5 = v5
    fmt.Println(v5)
    v5 -= v5
    v20 := []int{}
    _ = len(v20)
    if true {
        if len(v20) > 0 {
            v20[0] = (v4 - (v4 ^ v4))
        }
        _ = len(v20)
    } else {
        _ = len(v20)
        v21 := GenS1{F0: true, F1: 362}
        var v22 int16 = 0
        _ = v21
        _ = v22
    }
    var v23 bool = false
    for fi24 := 0; fi24 < 14; fi24++ {
        for fi25 := 0; fi25 < 6; fi25++ {
            v26 := v2[65535]
            v27 := GenS1{F0: v23, F1: 42}
            v28 := GenS3{F0: 151}
            switch v4 {
            case 0:
                var v29 uint8 = genFn0(("" + (("" + "a") + "")), 255, (1e10 < 1.5))
                _ = v29
            case 9223372036854775807:
                v20 = append(v20, (v26 - 32768))
            default:
                var v30 uint32 = 65535
                _ = v30
            }
            _ = fi25
            _ = v26
            _ = v27
            _ = v28
        }
        if (1e10 <= 3.14159) {
            v31 := GenS2{F0: 90}
            _ = v31
        } else {
            delete(v1, "hello")
            fmt.Println(v4)
        }
        _ = fi24
    }
    var v32 int16 = 182
    switch ("foo bar" + (("hello" + "foo bar") + "")) {
    case "\n\t":
        fmt.Println(v32)
    case "foo bar":
        v33 := GenS3{F0: -2147483648}
        _ = v33
    default:
        var v34 uint32 = 0
        _ = v34
    }
    for fi35 := 0; fi35 < 20; fi35++ {
        _ = len(v3)
        var v36 uint = 42
        _ = fi35
        _ = v36
    }
    v5 += v5
    switch v4 {
    case -2147483648:
        for fi37 := 0; fi37 < 16; fi37++ {
            v4 = 1
            for fi38 := 0; fi38 < 17; fi38++ {
                v39 := v2[42]
                fmt.Println(v23, v4, v32)
                _ = fi38
                _ = v39
            }
            _ = fi37
        }
    case 2147483647:
        for fi40 := 0; fi40 < 1; fi40++ {
            v41 := GenS1{F0: v23, F1: 47}
            v5 -= v5
            _ = fi40
            _ = v41
        }
    case 0:
        _ = len(v20)
    }
    v32 = (v32 * v32)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v20
    _ = v23
    _ = v3
    _ = v32
    _ = v4
    _ = v5
}
