package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32 }
func genFn0(p0 int, p1 string) int8 {
    p1 = (p1 + p1)
    v6 := GenS1{F0: 63, F1: 42}
    for fi7 := 0; fi7 < 9; fi7++ {
        var v8 uint8 = 128
        for fi9 := 0; fi9 < 11; fi9++ {
            fmt.Println(p0)
            switch (-(p0)) {
            case 1:
                var v10 float32 = 3.14159
                _ = v10
            default:
                if (1 != 4294967295) {
                    fmt.Println(p1, p0, fi7)
                    var v11 int16 = -32768
                    v12 := []int{}
                    _ = v11
                    _ = v12
                }
            }
            p0 += fi9
            _ = fi9
        }
        _ = fi7
        _ = v8
    }
    v13 := GenS1{F0: 217, F1: -618}
    _ = p0
    _ = p1
    _ = v13
    _ = v6
    return 1
}

func main() {
    v1 := make(map[int]int)
    var v2 uint32 = 2147483647
    v3 := []uint32{42, (v2 + v2), 42}
    var v4 int32 = 730
    var v5 int16 = 1000
    v4 = (v4 ^ (v4 >> ((18446744073709551615) & 63)))
    v2 += v2
    for fi14 := 0; fi14 < 7; fi14++ {
        var v15 uint64 = 42
        _ = fi14
        _ = v15
    }
    delete(v1, 0)
    switch ("foo bar" + "\n\t") {
    case "a":
        var v16 float32 = 0.1
        _ = v16
    case "foo bar":
        v2 -= (^(v2))
    default:
        var v17 float64 = 1e10
        _ = v17
    }
    fmt.Println(v4, v2, v4)
    for fi18 := 0; fi18 < 3; fi18++ {
        if len(v3) > 0 {
            v3[0] = v2
        }
        _ = fi18
    }
    if true {
        if true {
            v19 := GenS1{F0: 91, F1: (v4 * v4)}
            _ = v19
        }
        var v20 uint8 = 2
        v21 := GenS1{F0: v20, F1: (-(v4))}
        _ = v20
        _ = v21
    } else {
        fmt.Println(v4)
    }
    for fi22 := 0; fi22 < 19; fi22++ {
        var v23 float64 = 1e10
        v4 = (-(v4))
        _ = fi22
        _ = v23
    }
    for fi24 := 0; fi24 < 15; fi24++ {
        var v25 int16 = (v5 % ((-32768) | 1))
        _ = fi24
        _ = v25
    }
    if ((false || (false || (!false))) && ((0.1 - 1.5) != 1e10)) {
        for fi26 := 0; fi26 < 8; fi26++ {
            _ = len(v1)
            _ = fi26
        }
        var v27 uint16 = 1
        var v28 uint8 = 69
        _ = v27
        _ = v28
    } else {
        switch "a" {
        case "foo bar":
            var v29 int32 = v4
            _ = v29
        case "a":
            var v30 int16 = (v5 >> ((4294967295) & 63))
            _ = v30
        case "\n\t":
            switch 32768 {
            case -1:
                _ = len(v1)
            case 1:
                if true {
                    switch 32767 {
                    case 9223372036854775807:
                        v2 -= v2
                    }
                }
            }
        }
        var v31 int16 = (v5 ^ 32767)
        v1[0] = -1
        _ = v31
    }
    if (-536 != 310) {
        switch -440 {
        case -500:
            v2 -= (v2 & v2)
        }
        v32 := GenS1{F0: 127, F1: (v4 / ((v4) | 1))}
        v4 -= (-(v4))
        _ = v32
    } else {
        v4 += (v4 << ((1) & 63))
        fmt.Println(v4, v5, v4)
    }
    var v33 float32 = float32((0.0 / 1.0))
    v34 := GenS1{F0: 255, F1: v4}
    var v35 int = 128
    var v36 int16 = (-(v5))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v33:", v33)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v4
    _ = v5
}
