package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint8 }
type GenS2 struct { F0 uint8; F1 int8 }
type GenS3 struct { F0 string; F1 int32; F2 int32 }
func genFn0(p0 uint32, p1 uint, p2 int) int16 {
    p0 = (p0 * (p0 / (((^(p0))) | 1)))
    p0 = p0
    var v4 bool = true
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 881
}
func genFn1(p0 string, p1 uint16) uint16 {
    if true {
        if (2147483647 < 42) {
            var v5 int32 = 1
            var v6 string = p0
            _ = v5
            _ = v6
        }
        v7 := []int{4294967295}
        var v8 int64 = -238
        _ = v7
        _ = v8
    } else {
        var v9 int32 = 2147483647
        v10 := []int{32768, 128, -1}
        v9 = 1
        _ = v10
        _ = v9
    }
    for fi11 := 0; fi11 < 4; fi11++ {
        v12 := GenS1{F0: fi11, F1: 1}
        var v13 int32 = -1
        _ = fi11
        _ = v12
        _ = v13
    }
    for fi14 := 0; fi14 < 15; fi14++ {
        for fi15 := 0; fi15 < 12; fi15++ {
            if false {
                p1 = p1
            }
            if (fi14 < fi14) {
                switch ((p0 + "hello") + "a") {
                case "foo bar":
                    if false {
                        p1 -= p1
                        fmt.Println(p1)
                        v16 := []int{-9223372036854775808, (fi14 * fi15), (fi14 | fi14), (fi15 % (((^(fi15))) | 1))}
                        v17 := GenS2{F0: 1, F1: 100}
                        _ = v16
                        _ = v17
                    }
                case "":
                    switch (fi14 << ((4294967295) & 63)) {
                    case -527:
                        v18 := []int{4294967295, (fi14 & (fi14 + (fi15 << ((2) & 63)))), fi14}
                        _ = v18
                    case -116:
                        fmt.Println(fi14, p1)
                        if false {
                            v19 := GenS1{F0: (^(fi15)), F1: 0}
                            _ = v19
                        } else {
                            p1 -= (p1 / (((p1 & (p1 & (p1 + (p1 % (((p1 / ((664) | 1))) | 1)))))) | 1))
                            var v20 uint16 = (p1 + p1)
                            _ = v20
                        }
                    }
                    var v21 float64 = (1.5 / ((float64(-2.25) / 1.0) / 0.1))
                    _ = v21
                case "hello":
                    var v22 int16 = 42
                    _ = v22
                }
                var v23 int64 = 0
                _ = v23
            } else {
                p1 -= (p1 % (((^(p1))) | 1))
            }
            p0 = p0
            fmt.Println(p1, p1, fi14)
            _ = fi15
        }
        v24 := []int{(fi14 & (fi14 ^ (fi14 + fi14)))}
        v24 = append(v24, fi14)
        _ = fi14
        _ = v24
    }
    if ("foo bar" <= ("a" + ("hello" + p0))) {
        var v25 uint32 = 65535
        fmt.Println(p1, p1, v25)
        var v26 int8 = -128
        _ = v25
        _ = v26
    }
    var v27 uint64 = 0
    _ = p0
    _ = p1
    _ = v27
    return (^(p1))
}

func main() {
    var v1 int = 1
    v2 := []uint32{42, 1, 2147483648, 300}
    v3 := []int{}
    if len(v3) > 0 {
        v3[0] = -361
    }
    switch v1 {
    case -1:
        var v28 float64 = (float64((((3.14159 + 3.14159) * -2.25) / 1.5)) + 1e10)
        _ = v28
    }
    fmt.Println(v1)
    fmt.Println(v1)
    if ((729 <= 9223372036854775807) && (!(true && (!(2147483647 <= 65535))))) {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    }
    switch v1 {
    case -1:
        v1 = -2147483648
    }
    for fi29 := 0; fi29 < 17; fi29++ {
        if (0 <= 268) {
            v2 = append(v2, 65535)
            v30 := GenS3{F0: "\n\t", F1: 1, F2: -1}
            var v31 int32 = -2147483648
            _ = v30
            _ = v31
        } else {
            v32 := GenS2{F0: 127, F1: 1}
            v1 = (fi29 / (((-(v1))) | 1))
            switch ("hello" + (("foo bar" + "\n\t") + "a")) {
            case "a":
                v1 += fi29
            case "\n\t":
                switch 1 {
                case -747:
                    fmt.Println(fi29, fi29)
                case -9223372036854775808:
                    for fi33 := 0; fi33 < 8; fi33++ {
                        if (false || (2147483647 <= -784)) {
                            v1 += fi29
                        }
                        var v34 uint = 220
                        _ = fi33
                        _ = v34
                    }
                default:
                    for fi35 := 0; fi35 < 3; fi35++ {
                        v1 = -2147483648
                        _ = fi35
                    }
                }
            default:
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            }
            _ = v32
        }
        _ = fi29
    }
    var v36 int16 = 0
    fmt.Println(v36, v1)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v36:", v36)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
}
