package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64 }
type GenS2 struct { F0 string; F1 uint; F2 int }
func genFn0(p0 int32, p1 uint) int64 {
    fmt.Println(p1, p1)
    var v7 bool = true
    for fi8 := 0; fi8 < 6; fi8++ {
        switch (fi8 | (fi8 / ((-2147483648) | 1))) {
        case 32767:
            for fi9 := 0; fi9 < 12; fi9++ {
                var v10 float64 = (float64(1e10) + (1.5 * 0.1))
                v7 = v7
                var v11 int64 = 42
                _ = fi9
                _ = v10
                _ = v11
            }
        case 32768:
            fmt.Println(p0, fi8, fi8)
        }
        _ = fi8
    }
    for fi12 := 0; fi12 < 4; fi12++ {
        v13 := []int{128, (fi12 | (-(fi12)))}
        _ = fi12
        _ = v13
    }
    if v7 {
        v14 := GenS1{F0: 175, F1: 4294967295}
        switch "foo bar" {
        case "a":
            p0 += (p0 - p0)
        case "":
            v15 := GenS1{F0: 2, F1: 0}
            _ = v15
        case "foo bar":
            var v16 float64 = 1.5
            _ = v16
        default:
            if (v7 || ((^(p1)) != 4294967295)) {
                var v17 uint = (p1 + (p1 & (p1 ^ 4294967295)))
                var v18 int64 = 0
                switch 1 {
                case -400:
                    p1 += (p1 | (v17 % (((p1 * 2)) | 1)))
                default:
                    var v19 int64 = v18
                    _ = v19
                }
                _ = v17
                _ = v18
            } else {
                fmt.Println(v7, p1)
            }
        }
        _ = v14
    } else {
        v20 := []int{-290, -448, 4294967295}
        var v21 int8 = -128
        _ = v20
        _ = v21
    }
    _ = p0
    _ = p1
    _ = v7
    return -847
}
func genFn1() uint16 {
    if false {
        var v22 int8 = -128
        _ = v22
    }
    v23 := []int{9223372036854775807, -244, 9223372036854775807, 127}
    v24 := GenS2{F0: "hello", F1: 443, F2: 32768}
    _ = v23
    _ = v24
    return 2
}

func main() {
    var v1 int = 32768
    var v2 uint32 = 0
    v3 := make(map[int]int)
    var v4 int = (v1 | (v1 >> ((297) & 63)))
    var v5 int32 = 42
    var v6 int16 = 143
    for fi25 := 0; fi25 < 14; fi25++ {
        v6 += (v6 * 32767)
        _ = fi25
    }
    fmt.Println(v6, v4, v2)
    v5 += (-(v5))
    var v26 uint = 18446744073709551615
    var v27 int8 = -128
    for fi28 := 0; fi28 < 10; fi28++ {
        fmt.Println(v26, v1, v2)
        _ = fi28
    }
    var v29 float32 = 0.0
    switch (("a" + "") + "a") {
    case "a":
        var v30 uint64 = 1
        _ = v30
    case "\n\t":
        fmt.Println(v5, v6)
    default:
        if (true || (!(!(false || (42 < v27))))) {
            fmt.Println(v2)
        } else {
            var v31 float32 = (1e10 + 0.1)
            v32 := GenS2{F0: ("hello" + "\n\t"), F1: 985, F2: (v4 ^ (v4 - (v1 | 9223372036854775807)))}
            if (!(false || true)) {
                v5 -= v5
            } else {
                if false {
                    var v33 int32 = (v5 + (-(v5)))
                    v27 = -1
                    _ = v33
                }
                v27 = v27
            }
            _ = v31
            _ = v32
        }
    }
    switch "" {
    case "\n\t":
        v4 = (-(v1))
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
