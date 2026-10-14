package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint32; F2 int64 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 string; F1 int32 }
func genFn0() uint64 {
    v6 := GenS1{F0: -32768, F1: 1, F2: -947}
    var v7 uint16 = 1
    if true {
        var v8 uint32 = 138
        v7 -= (^(v7))
        _ = v8
    } else {
        fmt.Println(v7)
        v7 -= 255
    }
    _ = v6
    _ = v7
    return 1
}
func genFn1(p0 uint64, p1 uint32) int16 {
    if ((72 <= 127) || (true && ((1.5 + (3.14159 + ((0.0 * (((float64(0.0) - (3.14159 - 1e10)) / (3.14159 - 1.5)) / 1e10)) + (1.5 * 1e10)))) <= ((1e10 / 1.0) / 1.0)))) {
        p1 -= (p1 + (p1 & 452))
        v9 := []int{2147483647, 494, -723}
        _ = v9
    } else {
        v10 := []int{32767, 1, 9223372036854775807, 509}
        _ = v10
    }
    v11 := GenS1{F0: 973, F1: (^(p1)), F2: 1}
    for fi12 := 0; fi12 < 17; fi12++ {
        var v13 int64 = -9223372036854775808
        _ = fi12
        _ = v13
    }
    _ = p0
    _ = p1
    _ = v11
    return 0
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 int16 = 0
    var v3 int32 = -56
    var v4 float64 = (1e10 / (((((1.5 * 3.14159) * 3.14159) / 0.1) * 0.1) / 0.1))
    var v5 bool = true
    if (false || (!v5)) {
        if (false && (-1 < 100)) {
            v14 := GenS1{F0: genFn1(18446744073709551615, 387), F1: 42, F2: 0}
            v3 = (v3 + (^(v3)))
            _ = v14
        }
    }
    if (1 != 2147483648) {
        switch 858 {
        case 986:
            v4 += 1.5
        case -9223372036854775808:
            v4 = 1e10
        default:
            v15 := []int{32768, -417, -2147483648, 9223372036854775807}
            _ = v15
        }
        v16 := []int{2147483647, -907, -98, -9223372036854775808}
        _ = v16
    } else {
        var v17 int = 9223372036854775807
        _ = v17
    }
    delete(v1, 2147483648)
    var v18 string = "hello"
    v5 = ((!true) && false)
    var v19 int = 655
    var v20 bool = (genFn1(genFn0(), 881) < (v2 - v2))
    var v21 uint32 = 14
    for fi22 := 0; fi22 < 18; fi22++ {
        v23 := v1[619]
        v24 := GenS1{F0: (v2 >> ((255) & 63)), F1: v21, F2: -9223372036854775808}
        fmt.Println(v5, v19)
        _ = fi22
        _ = v23
        _ = v24
    }
    if ((v2 >> ((4294967295) & 63)) <= genFn1(4294967295, (v21 >> ((2) & 63)))) {
        fmt.Println(v2, v2)
        var v25 uint16 = 65535
        for fi26 := 0; fi26 < 17; fi26++ {
            v19 = (fi26 ^ fi26)
            _ = fi26
        }
        _ = v25
    }
    switch v19 {
    case 9223372036854775807:
        switch "a" {
        case "foo bar":
            fmt.Println(v20, v18)
        default:
            v21 -= v21
        }
    case -996:
        for fi27 := 0; fi27 < 9; fi27++ {
            v28 := []int{}
            fmt.Println(v18)
            _ = fi27
            _ = v28
        }
    default:
        v29 := GenS2{F0: v21}
        _ = v29
    }
    for fi30 := 0; fi30 < 13; fi30++ {
        if v20 {
            var v31 int16 = (v2 | 0)
            var v32 uint16 = 255
            _ = v31
            _ = v32
        }
        v33 := GenS3{F0: 2}
        v4 += v4
        _ = fi30
        _ = v33
    }
    v20 = v5
    if (0 != 375) {
        v2 = (-(v2))
        for fi34 := 0; fi34 < 11; fi34++ {
            if true {
                v35 := GenS4{F0: v18, F1: (v3 | (v3 | 2147483647))}
                var v36 uint64 = 1
                _ = v35
                _ = v36
            } else {
                v20 = (1 <= -9223372036854775808)
            }
            v37 := GenS4{F0: v18, F1: (-(v3))}
            var v38 int8 = 42
            _ = fi34
            _ = v37
            _ = v38
        }
    }
    for fi39 := 0; fi39 < 9; fi39++ {
        for fi40 := 0; fi40 < 9; fi40++ {
            var v41 int64 = 0
            v42 := GenS1{F0: v2, F1: 2147483647, F2: -207}
            _ = fi40
            _ = v41
            _ = v42
        }
        var v43 int64 = 0
        _ = fi39
        _ = v43
    }
    for fi44 := 0; fi44 < 2; fi44++ {
        for fi45 := 0; fi45 < 1; fi45++ {
            v46 := GenS3{F0: 41}
            fmt.Println(v20, v18, v18)
            _ = fi45
            _ = v46
        }
        _ = fi44
    }
    var v47 int16 = v2
    for fi48 := 0; fi48 < 5; fi48++ {
        var v49 float64 = (0.0 * v4)
        _ = fi48
        _ = v49
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v47:", v47)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v47
    _ = v5
}
