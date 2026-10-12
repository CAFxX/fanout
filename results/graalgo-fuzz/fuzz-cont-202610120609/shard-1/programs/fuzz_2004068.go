package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16; F2 uint64 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint16; F1 int; F2 int }
type GenS5 struct { F0 int64; F1 string; F2 int16 }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 int; F1 int32 }
func genFn0() int64 {
    var v5 float64 = (float64(1e10) * 1e10)
    v6 := []int{127, 42}
    v5 += 3.14159
    v7 := GenS1{F0: 997, F1: 0, F2: 18446744073709551615}
    _ = v5
    _ = v6
    _ = v7
    return 9223372036854775807
}
func genFn1(p0 uint16, p1 string, p2 string) uint16 {
    v8 := GenS1{F0: 962, F1: 539, F2: 18446744073709551615}
    for fi9 := 0; fi9 < 13; fi9++ {
        fmt.Println(fi9)
        _ = fi9
    }
    if (18446744073709551615 >= 490) {
        var v10 int8 = -1
        _ = v10
    } else {
        switch p1 {
        case "hello":
            v11 := GenS2{F0: 42}
            _ = v11
        case "\n\t":
            for fi12 := 0; fi12 < 15; fi12++ {
                var v13 uint32 = 0
                for fi14 := 0; fi14 < 7; fi14++ {
                    var v15 uint32 = uint32(p0)
                    fmt.Println(v13, fi14)
                    if (false || false) {
                        switch "a" {
                        case "hello":
                            var v16 uint32 = v13
                            _ = v16
                        }
                        v13 = (v15 ^ v13)
                    }
                    var v17 int64 = 42
                    _ = fi14
                    _ = v15
                    _ = v17
                }
                fmt.Println(fi12, p1, p2)
                _ = fi12
                _ = v13
            }
        }
        for fi18 := 0; fi18 < 18; fi18++ {
            v19 := GenS2{F0: 2}
            v20 := GenS3{F0: 1}
            p1 = p2
            _ = fi18
            _ = v19
            _ = v20
        }
    }
    p1 = "\n\t"
    var v21 string = "foo bar"
    _ = p0
    _ = p1
    _ = p2
    _ = v21
    _ = v8
    return p0
}

func main() {
    v1 := []uint32{200}
    var v2 int8 = 127
    v3 := make(map[uint32]string)
    var v4 int64 = 42
    switch 32767 {
    case 0:
        v1 = append(v1, 2147483648)
    case 535:
        v22 := GenS4{F0: 459, F1: -839, F2: 494}
        _ = v22
    case 1:
        switch ("" + "hello") {
        case "foo bar":
            var v23 bool = ((42 < 42) || (-1 > 1))
            _ = v23
        case "hello":
            var v24 int16 = 0
            _ = v24
        }
    default:
        var v25 int32 = -742
        _ = v25
    }
    v26 := GenS3{F0: 255}
    delete(v3, 29)
    var v27 string = (("\n\t" + ("hello" + ("\n\t" + "foo bar"))) + "foo bar")
    if (0 != 24) {
        v2 = v2
        delete(v3, 42)
        for fi28 := 0; fi28 < 8; fi28++ {
            fmt.Println(v2)
            switch -929 {
            case 2147483647:
                switch (v27 + (v27 + v27)) {
                case "foo bar":
                    v1 = append(v1, 61)
                case "a":
                    if false {
                        v27 = v27
                    }
                case "hello":
                    v2 = (^(v2))
                default:
                    _ = len(v3)
                }
            case 0:
                fmt.Println(v4, v27, v2)
            case 1:
                for fi29 := 0; fi29 < 17; fi29++ {
                    switch v27 {
                    case "a":
                        v30 := GenS5{F0: v4, F1: v27, F2: 42}
                        _ = v30
                    default:
                        var v31 uint8 = 128
                        _ = v31
                    }
                    if false {
                        var v32 uint8 = 127
                        var v33 float64 = 1.5
                        _ = v32
                        _ = v33
                    } else {
                        var v34 uint16 = 0
                        if true {
                            delete(v3, 86)
                            var v35 int64 = (v4 & (v4 ^ genFn0()))
                            _ = v35
                        }
                        v4 -= v4
                        var v36 int = (^(fi28))
                        _ = len(v1)
                        switch "foo bar" {
                        case "foo bar":
                            v27 = "foo bar"
                        case "\n\t":
                            v4 = (v4 | 0)
                        case "a":
                            fmt.Println(v4, v4)
                        default:
                            fmt.Println(v36)
                        }
                        _ = v34
                        _ = v36
                    }
                    _ = fi29
                }
            default:
                v4 = -9223372036854775808
            }
            _ = fi28
        }
    } else {
        v2 = (v2 - (v2 / ((v2) | 1)))
    }
    v2 -= v2
    fmt.Println(v27)
    switch v27 {
    case "hello":
        if false {
            var v37 uint32 = 800
            _ = v37
        }
    case "":
        for fi38 := 0; fi38 < 6; fi38++ {
            var v39 uint16 = 111
            for fi40 := 0; fi40 < 12; fi40++ {
                switch int(v39) {
                case 2147483647:
                    var v41 bool = false
                    _ = v41
                case 4294967295:
                    v27 = v27
                case -174:
                    v2 += v2
                }
                fmt.Println(v4)
                fmt.Println(fi40, fi40, fi40)
                _ = fi40
            }
            _ = fi38
            _ = v39
        }
    case "foo bar":
        v3[632] = "foo bar"
    default:
        v42 := GenS6{F0: 131}
        _ = v42
    }
    v27 = v27
    if ((v4 ^ 960) <= genFn0()) {
        v43 := GenS7{F0: -9223372036854775808, F1: 42}
        v27 = v27
        v2 = (-(v2))
        _ = v43
    }
    delete(v3, 555)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v4
}
