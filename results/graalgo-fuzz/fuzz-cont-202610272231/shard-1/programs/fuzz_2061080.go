package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 bool; F2 int }
type GenS2 struct { F0 string; F1 uint32 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 uint32; F1 uint; F2 uint8 }
type GenS5 struct { F0 uint8 }
func genFn0(p0 int32, p1 int16, p2 int32) string {
    v6 := []int{-9223372036854775808, 42}
    p1 -= (p1 & p1)
    fmt.Println(p2, p2)
    fmt.Println(p2, p0)
    fmt.Println(p0, p2, p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return "hello"
}

func main() {
    var v1 bool = ("hello" != "")
    v2 := []uint32{}
    var v3 int = -764
    var v4 uint = 718
    var v5 int16 = 1000
    var v7 float32 = (1.5 / -2.25)
    switch "" {
    case "\n\t":
        for fi8 := 0; fi8 < 10; fi8++ {
            for fi9 := 0; fi9 < 9; fi9++ {
                fmt.Println(v5, v4, v3)
                _ = fi9
            }
            _ = fi8
        }
    default:
        if v1 {
            v4 = (v4 + v4)
        }
    }
    v10 := GenS1{F0: 2, F1: ((65535 >= 0) || true), F2: v3}
    v3 += int(v5)
    if v1 {
        var v11 uint8 = 12
        switch (-(v3)) {
        case 128:
            var v12 int16 = 812
            _ = v12
        case 32768:
            fmt.Println(v7, v11)
        }
        _ = v11
    } else {
        v13 := GenS2{F0: ("" + ((("\n\t" + ((("" + "a") + "foo bar") + "hello")) + "a") + "a")), F1: 494}
        _ = v13
    }
    for fi14 := 0; fi14 < 14; fi14++ {
        v7 -= (0.0 + 0.0)
        v1 = (65535 > v4)
        fmt.Println(v4, v5)
        _ = fi14
    }
    for fi15 := 0; fi15 < 4; fi15++ {
        if false {
            var v16 bool = (v1 || false)
            var v17 uint = (v4 | v4)
            _ = v16
            _ = v17
        }
        _ = fi15
    }
    var v18 uint32 = 2
    fmt.Println(v7, v4, v1)
    v19 := GenS3{F0: v3}
    _ = len(v2)
    var v20 uint64 = 2
    for fi21 := 0; fi21 < 11; fi21++ {
        for fi22 := 0; fi22 < 2; fi22++ {
            v23 := GenS4{F0: 2147483647, F1: (^(v4)), F2: 128}
            fmt.Println(v20, v3)
            if v1 {
                if len(v2) > 0 {
                    v2[0] = 4294967295
                }
            } else {
                var v24 uint16 = 728
                var v25 string = "\n\t"
                _ = v24
                _ = v25
            }
            _ = fi22
            _ = v23
        }
        v26 := GenS2{F0: "a", F1: 2147483647}
        v27 := GenS5{F0: 127}
        _ = fi21
        _ = v26
        _ = v27
    }
    v4 += v4
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v7
}
