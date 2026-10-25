package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32 }
type GenS2 struct { F0 uint16; F1 int64 }
func genFn0(p0 int16, p1 string) uint8 {
    fmt.Println(p1)
    var v6 uint32 = 2147483647
    var v7 int64 = -772
    if false {
        v7 += (v7 + v7)
        for fi8 := 0; fi8 < 10; fi8++ {
            fmt.Println(p0, fi8)
            _ = fi8
        }
        var v9 int8 = 1
        _ = v9
    } else {
        v7 -= -1
        v6 += (v6 + (v6 << ((0) & 63)))
        var v10 float64 = 1.5
        _ = v10
    }
    var v11 int32 = 42
    _ = p0
    _ = p1
    _ = v11
    _ = v6
    _ = v7
    return 0
}
func genFn1(p0 uint) uint64 {
    for fi12 := 0; fi12 < 1; fi12++ {
        var v13 int8 = 1
        p0 += (p0 & p0)
        for fi14 := 0; fi14 < 7; fi14++ {
            if (!true) {
                var v15 float64 = 3.14159
                p0 = 1
                v16 := []int{(fi14 * fi12)}
                _ = v15
                _ = v16
            }
            _ = fi14
        }
        _ = fi12
        _ = v13
    }
    p0 -= p0
    p0 = p0
    if false {
        var v17 int64 = -1
        v18 := []int{-9223372036854775808}
        switch "a" {
        case "hello":
            if ((true && false) && true) {
                var v19 int32 = 42
                _ = v19
            }
        case "\n\t":
            for fi20 := 0; fi20 < 17; fi20++ {
                if len(v18) > 0 {
                    v18[0] = fi20
                }
                if ((94 == 2) || true) {
                    var v21 int16 = -867
                    if (!true) {
                        if len(v18) > 0 {
                            v18[0] = (-(fi20))
                        }
                    }
                    v22 := GenS1{F0: (fi20 & -9223372036854775808), F1: 2}
                    _ = v21
                    _ = v22
                } else {
                    v17 -= (v17 >> ((42) & 63))
                    switch "\n\t" {
                    case "a":
                        for fi23 := 0; fi23 < 18; fi23++ {
                            if false {
                                switch "\n\t" {
                                case "foo bar":
                                    v17 -= (-(v17))
                                }
                            }
                            var v24 int = fi20
                            if ((false || true) && true) {
                                v17 += v17
                            }
                            p0 += 749
                            _ = fi23
                            _ = v24
                        }
                    case "hello":
                        p0 = p0
                    case "":
                        v17 = (v17 >> ((2) & 63))
                    }
                    _ = len(v18)
                }
                v25 := GenS2{F0: 390, F1: (v17 >> ((0) & 63))}
                _ = fi20
                _ = v25
            }
        case "a":
            fmt.Println(v17)
        default:
            fmt.Println(p0, p0)
        }
        _ = v17
        _ = v18
    }
    _ = p0
    return 590
}

func main() {
    var v1 uint16 = 2
    v2 := make(map[int]string)
    var v3 int64 = 1
    var v4 uint16 = 255
    var v5 int64 = v3
    _ = len(v2)
    fmt.Println(v3)
    fmt.Println(v5)
    var v26 int8 = 17
    var v27 int64 = (v3 | -536)
    switch "hello" {
    case "hello":
        fmt.Println(v3)
    default:
        v3 = v3
    }
    for fi28 := 0; fi28 < 8; fi28++ {
        fmt.Println(v1, fi28)
        var v29 int32 = 42
        _ = fi28
        _ = v29
    }
    var v30 string = "hello"
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
