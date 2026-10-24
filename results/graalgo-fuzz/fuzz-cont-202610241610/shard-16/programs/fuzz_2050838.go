package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint16; F1 int16; F2 uint16 }
type GenS3 struct { F0 int8; F1 int8; F2 uint }
func genFn0(p0 uint8, p1 uint8) int8 {
    fmt.Println(p0)
    v6 := GenS1{F0: p0}
    v7 := GenS2{F0: 2, F1: 1000, F2: 2}
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return -28
}
func genFn1(p0 uint, p1 string) uint8 {
    for fi8 := 0; fi8 < 3; fi8++ {
        var v9 float64 = 1e10
        _ = fi8
        _ = v9
    }
    v10 := []int{-2147483648}
    for fi11 := 0; fi11 < 4; fi11++ {
        for fi12 := 0; fi12 < 17; fi12++ {
            switch "" {
            case "foo bar":
                var v13 uint64 = 77
                _ = v13
            case "hello":
                p0 = 2
            case "":
                p1 = p1
            default:
                if len(v10) >= 2 {
                    v10 = v10[1:len(v10)]
                }
            }
            fmt.Println(p0, p1, p1)
            _ = fi12
        }
        _ = fi11
    }
    var v14 uint = (p0 * 42)
    _ = p0
    _ = p1
    _ = v10
    _ = v14
    return 1
}

func main() {
    v1 := make(map[string]uint32)
    var v2 int32 = 42
    v3 := make(map[int]int)
    var v4 uint8 = 201
    v5 := make(map[string]bool)
    v2 = v2
    v15 := []int{1}
    v2 -= v2
    v4 -= (v4 | 255)
    switch 127 {
    case 0:
        v2 += -1
    default:
        var v16 int32 = -1
        _ = v16
    }
    v4 += (v4 * 255)
    v17 := GenS1{F0: 196}
    if true {
        switch 32767 {
        case 0:
            fmt.Println(v2, v2)
        default:
            v4 = v4
        }
        v2 = -1
    } else {
        var v18 uint32 = 1
        _ = len(v15)
        _ = len(v5)
        _ = v18
    }
    var v19 bool = (!((3.14159 / (1.5 - 0.1)) >= (float64(-2.25) / 1.0)))
    for fi20 := 0; fi20 < 1; fi20++ {
        var v21 uint8 = v4
        fmt.Println(fi20, v4)
        _ = fi20
        _ = v21
    }
    switch ("hello" + (("hello" + "") + "\n\t")) {
    case "\n\t":
        if v19 {
            if v19 {
                v22 := v1["a"]
                v2 = v2
                v23 := GenS1{F0: (v4 >> ((255) & 63))}
                _ = v22
                _ = v23
            }
        }
    case "a":
        switch (("foo bar" + "hello") + (("\n\t" + "\n\t") + "foo bar")) {
        case "a":
            v15 = append(v15, 836)
        }
    default:
        for fi24 := 0; fi24 < 14; fi24++ {
            v25 := GenS3{F0: 1, F1: 1, F2: 489}
            _ = fi24
            _ = v25
        }
    }
    v26 := GenS2{F0: 512, F1: 0, F2: 32767}
    switch "hello" {
    case "foo bar":
        var v27 int64 = 0
        _ = v27
    case "\n\t":
        v19 = v19
    case "":
        v4 = 1
    default:
        for fi28 := 0; fi28 < 4; fi28++ {
            v2 = -2147483648
            if v19 {
                var v29 int32 = (v2 % ((v2) | 1))
                v4 = 255
                var v30 int = (fi28 / ((fi28) | 1))
                _ = v29
                _ = v30
            }
            _ = fi28
        }
    }
    v4 += (^(v4))
    var v31 uint8 = 1
    var v32 uint = 0
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v15
    _ = v17
    _ = v19
    _ = v2
    _ = v26
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v5
}
