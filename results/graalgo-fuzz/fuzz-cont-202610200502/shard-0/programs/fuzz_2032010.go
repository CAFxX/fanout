package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int64 }
type GenS2 struct { F0 uint64; F1 bool; F2 uint16 }
type GenS3 struct { F0 uint16; F1 uint16 }
func genFn0(p0 int, p1 int16) uint32 {
    var v5 int32 = 2147483647
    var v6 int = (^(p0))
    if (!true) {
        v7 := []int{}
        var v8 int32 = (v5 / ((v5) | 1))
        _ = len(v7)
        _ = v7
        _ = v8
    }
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return 1
}

func main() {
    var v1 uint64 = 42
    var v2 uint64 = (v1 >> ((65535) & 63))
    v3 := []uint32{4294967295, 65535, 497}
    v4 := make(map[int]int)
    v9 := GenS1{F0: false, F1: 677}
    var v10 bool = (false || false)
    switch 32768 {
    case 128:
        v2 -= 2
    case 761:
        v11 := v4[-900]
        _ = v11
    case 4294967295:
        fmt.Println(v1)
    default:
        fmt.Println(v10, v2)
    }
    if v10 {
        for fi12 := 0; fi12 < 1; fi12++ {
            for fi13 := 0; fi13 < 18; fi13++ {
                var v14 int = (-(fi13))
                v2 += v2
                _ = fi13
                _ = v14
            }
            var v15 int = fi12
            _ = fi12
            _ = v15
        }
    }
    if true {
        switch "a" {
        case "hello":
            switch "foo bar" {
            case "hello":
                fmt.Println(v1, v1, v10)
            case "\n\t":
                v2 -= (v2 * (v2 % ((174) | 1)))
            case "foo bar":
                v16 := v4[32768]
                _ = v16
            default:
                v17 := GenS1{F0: (!v10), F1: 632}
                _ = v17
            }
        case "foo bar":
            v1 -= v1
        case "a":
            for fi18 := 0; fi18 < 2; fi18++ {
                v10 = v10
                _ = fi18
            }
        }
        _ = len(v3)
        v2 = v1
    } else {
        if v10 {
            var v19 uint64 = v2
            for fi20 := 0; fi20 < 6; fi20++ {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                _ = fi20
            }
            if (-32768 != 42) {
                fmt.Println(v10, v1, v10)
                var v21 uint16 = 255
                _ = v21
            }
            _ = v19
        } else {
            var v22 bool = (!(-9223372036854775808 >= 0))
            v23 := GenS1{F0: v10, F1: -9223372036854775808}
            v1 = v2
            _ = v22
            _ = v23
        }
        var v24 float64 = 1.5
        _ = v24
    }
    v1 = (^(v1))
    fmt.Println(v2, v10, v2)
    if (!(-1 <= 1)) {
        var v25 float64 = (3.14159 * 3.14159)
        _ = v25
    } else {
        fmt.Println(v10, v10, v2)
        v26 := GenS1{F0: (18446744073709551615 >= (v1 & (^(v2)))), F1: 248}
        _ = v26
    }
    var v27 bool = v10
    for fi28 := 0; fi28 < 16; fi28++ {
        var v29 uint32 = 181
        _ = fi28
        _ = v29
    }
    if ((!(437 != 1)) || true) {
        v30 := GenS2{F0: 1, F1: v27, F2: 113}
        _ = v30
    } else {
        switch -356 {
        case 2147483647:
            var v31 uint = 368
            _ = v31
        case 42:
            v2 += v2
        }
        var v32 int8 = 127
        v3 = append(v3, 84)
        _ = v32
    }
    var v33 int32 = 0
    fmt.Println(v2)
    switch ("hello" + "a") {
    case "\n\t":
        v1 -= v1
    }
    for fi34 := 0; fi34 < 1; fi34++ {
        v2 -= v2
        v35 := GenS3{F0: 32767, F1: 42}
        v36 := v4[fi34]
        _ = fi34
        _ = v35
        _ = v36
    }
    var v37 uint = 18446744073709551615
    var v38 int = 278
    if false {
        var v39 uint16 = 2
        _ = v39
    } else {
        v38 = v38
        if (uint8(v37) >= 127) {
            v4[v38] = 446
            if (0 < -128) {
                _ = len(v4)
            } else {
                v4[v38] = (v38 >> (((v37 << ((v37) & 63))) & 63))
                v40 := v4[(v38 * -2147483648)]
                _ = v40
            }
            fmt.Println(v37, v37, v37)
        } else {
            v4[1] = (v38 / ((1) | 1))
        }
        var v41 int16 = 42
        _ = v41
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v33:", v33)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v10
    _ = v2
    _ = v27
    _ = v3
    _ = v33
    _ = v37
    _ = v38
    _ = v4
    _ = v9
}
