package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint8; F1 int }
type GenS3 struct { F0 bool; F1 uint64 }
type GenS4 struct { F0 string; F1 uint8; F2 uint16 }
type GenS5 struct { F0 int8; F1 int16; F2 int8 }
type GenS6 struct { F0 int16; F1 int }
func genFn0(p0 uint) int32 {
    fmt.Println(p0)
    var v7 uint64 = 4294967295
    v7 = v7
    v8 := GenS1{F0: (p0 >> ((747) & 63))}
    if (false || (-1 >= -1)) {
        p0 = 255
        for fi9 := 0; fi9 < 2; fi9++ {
            p0 += 18446744073709551615
            _ = fi9
        }
        switch "\n\t" {
        case "hello":
            var v10 int8 = -128
            _ = v10
        default:
            for fi11 := 0; fi11 < 6; fi11++ {
                var v12 uint = 1
                _ = fi11
                _ = v12
            }
        }
    } else {
        v13 := GenS1{F0: 803}
        _ = v13
    }
    _ = p0
    _ = v7
    _ = v8
    return 42
}

func main() {
    var v1 int16 = 168
    var v2 bool = (true && ((-2.25 - 1.5) != 1.5))
    var v3 int = -75
    v4 := make(map[string]int)
    v5 := make(map[uint32]uint32)
    var v6 float64 = 0.0
    v14 := GenS1{F0: 301}
    switch v3 {
    case 1:
        v15 := GenS1{F0: 255}
        _ = v15
    case -556:
        switch "a" {
        case "foo bar":
            v16 := GenS1{F0: 1}
            _ = v16
        case "\n\t":
            fmt.Println(v2, v1)
        }
    }
    var v17 float64 = 0.1
    v18 := []int{}
    var v19 float64 = (v6 / 1.0)
    for fi20 := 0; fi20 < 17; fi20++ {
        var v21 int16 = v1
        _ = fi20
        _ = v21
    }
    switch (-(v3)) {
    case -522:
        if v2 {
            for fi22 := 0; fi22 < 10; fi22++ {
                v23 := GenS1{F0: 306}
                v24 := GenS1{F0: 4294967295}
                if len(v18) > 0 {
                    v18[0] = 32767
                }
                _ = fi22
                _ = v23
                _ = v24
            }
            fmt.Println(v3, v3)
            var v25 uint8 = 114
            _ = v25
        }
    case -1:
        if true {
            fmt.Println(v1, v1)
            var v26 bool = v2
            if ((v26 && (!(v26 || v2))) || v26) {
                _ = len(v18)
                var v27 uint32 = 2147483647
                _ = v27
            }
            _ = v26
        }
    }
    var v28 int = (-(v3))
    switch "" {
    case "hello":
        for fi29 := 0; fi29 < 8; fi29++ {
            v19 -= 0.0
            v17 = ((v17 / float64(((1.5 - 3.14159) + 0.1))) / v19)
            _ = fi29
        }
    }
    var v30 int8 = 1
    _ = len(v18)
    v31 := GenS2{F0: 1, F1: v28}
    var v32 uint32 = 2
    v33 := GenS3{F0: v2, F1: 1}
    if ((v2 || true) || (v2 || v2)) {
        for fi34 := 0; fi34 < 18; fi34++ {
            for fi35 := 0; fi35 < 13; fi35++ {
                var v36 float32 = 0.1
                _ = fi35
                _ = v36
            }
            _ = fi34
        }
        v6 -= 1.5
        switch v28 {
        case -1:
            switch -1 {
            case 30:
                fmt.Println(v19, v32, v28)
            case -9223372036854775808:
                var v37 int = (v3 - 42)
                _ = v37
            case -439:
                var v38 uint64 = 4294967295
                _ = v38
            default:
                fmt.Println(v3, v6, v6)
            }
        case -650:
            if (!(-2147483648 == -2147483648)) {
                if ((v1 * v1) > (v1 * v1)) {
                    for fi39 := 0; fi39 < 8; fi39++ {
                        v3 = v28
                        _ = fi39
                    }
                } else {
                    v40 := GenS4{F0: "hello", F1: 141, F2: 2}
                    v6 = v17
                    _ = v40
                }
                v41 := GenS5{F0: -128, F1: v1, F2: (^(v30))}
                _ = v41
            } else {
                if (127 != (v30 * -128)) {
                    v18 = append(v18, (v3 / ((v28) | 1)))
                }
            }
        case 128:
            v42 := GenS6{F0: v1, F1: -2147483648}
            _ = v42
        default:
            var v43 float32 = 1e10
            _ = v43
        }
    } else {
        v30 = 121
    }
    for fi44 := 0; fi44 < 12; fi44++ {
        fmt.Println(fi44, v2)
        var v45 uint = 803
        _ = fi44
        _ = v45
    }
    if (4294967295 >= v28) {
        for fi46 := 0; fi46 < 10; fi46++ {
            if v2 {
                v47 := GenS6{F0: v1, F1: (v28 / ((fi46) | 1))}
                v19 -= 1e10
                if len(v18) > 0 {
                    v18[0] = (^(v28))
                }
                _ = v47
            }
            _ = fi46
        }
    }
    if v2 {
        v48 := GenS5{F0: v30, F1: v1, F2: v30}
        var v49 string = "\n\t"
        _ = v48
        _ = v49
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v28
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}
