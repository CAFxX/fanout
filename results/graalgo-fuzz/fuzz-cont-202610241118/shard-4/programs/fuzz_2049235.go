package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 bool; F2 int8 }
type GenS2 struct { F0 int8; F1 uint8; F2 bool }
type GenS3 struct { F0 int32; F1 int8 }
type GenS4 struct { F0 uint8; F1 int16; F2 int32 }
type GenS5 struct { F0 uint16; F1 int8; F2 uint }
func genFn0(p0 uint32, p1 int64, p2 string) uint64 {
    var v5 int64 = p1
    var v6 uint8 = 25
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return uint64(v6)
}

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[uint32]uint32)
    var v3 int64 = 9223372036854775807
    var v4 uint = 42
    v7 := []int{9223372036854775807, 32768, 0, -2147483648}
    var v8 int32 = 1
    for fi9 := 0; fi9 < 18; fi9++ {
        if len(v7) > 0 {
            v7[0] = fi9
        }
        var v10 float64 = ((1.5 / 3.14159) * 0.0)
        switch (fi9 % ((fi9) | 1)) {
        case 4294967295:
            switch "a" {
            case "foo bar":
                if (float64((0.1 + 3.14159)) <= v10) {
                    v11 := GenS1{F0: 0, F1: false, F2: 100}
                    var v12 int64 = (v3 | v3)
                    _ = v11
                    _ = v12
                } else {
                    if false {
                        var v13 uint64 = 42
                        _ = v13
                    } else {
                        _ = len(v1)
                        v8 = v8
                    }
                    for fi14 := 0; fi14 < 18; fi14++ {
                        v15 := GenS1{F0: 0, F1: ((!((!false) && ((!(-128 > -43)) && (v10 >= 0.1)))) && ((v10 - (float64(-2.25) + float64(0.1))) >= 1.5)), F2: 1}
                        v3 = (v3 & (v3 * (-(v3))))
                        v16 := GenS2{F0: 1, F1: 1, F2: (v3 >= (v3 - 1))}
                        _ = fi14
                        _ = v15
                        _ = v16
                    }
                }
            default:
                fmt.Println(v4, v10)
            }
        }
        _ = fi9
        _ = v10
    }
    fmt.Println(v4, v8)
    if false {
        v8 -= (v8 << ((v4) & 63))
    } else {
        fmt.Println(v4)
        v8 = 0
        fmt.Println(v8)
    }
    var v17 int = 195
    switch ("hello" + (("foo bar" + ("foo bar" + "")) + "hello")) {
    case "foo bar":
        var v18 int8 = 1
        _ = v18
    case "a":
        for fi19 := 0; fi19 < 8; fi19++ {
            switch (fi19 + v17) {
            case -484:
                for fi20 := 0; fi20 < 20; fi20++ {
                    v21 := GenS1{F0: 0, F1: false, F2: 100}
                    v3 += v3
                    _ = fi20
                    _ = v21
                }
            case 124:
                var v22 uint32 = 1
                _ = v22
            }
            if (42 == 24) {
                var v23 uint = (v4 | 65535)
                _ = v23
            }
            v17 -= 455
            _ = fi19
        }
    default:
        switch v17 {
        case -9223372036854775808:
            for fi24 := 0; fi24 < 8; fi24++ {
                var v25 float64 = 1e10
                delete(v1, (v17 & 2147483647))
                var v26 int64 = v3
                _ = fi24
                _ = v25
                _ = v26
            }
        case -1:
            fmt.Println(v17)
        default:
            if len(v7) > 0 {
                v7[0] = 1
            }
        }
    }
    v17 = (-(v17))
    var v27 uint64 = 1
    switch (^(v17)) {
    case 778:
        for fi28 := 0; fi28 < 2; fi28++ {
            if len(v7) >= 2 {
                v7 = v7[1:len(v7)]
            }
            var v29 int32 = 42
            v3 -= (v3 >> ((v4) & 63))
            _ = fi28
            _ = v29
        }
    case -652:
        var v30 uint32 = 0
        _ = v30
    default:
        var v31 uint = 783
        _ = v31
    }
    var v32 uint32 = 0
    fmt.Println(v4, v3)
    switch "" {
    case "\n\t":
        if true {
            v33 := GenS3{F0: v8, F1: 127}
            switch "" {
            case "a":
                if ((true && false) && false) {
                    var v34 uint = v4
                    switch v17 {
                    case -9223372036854775808:
                        switch ("foo bar" + ((((("hello" + "foo bar") + "a") + ("\n\t" + "")) + ("foo bar" + "hello")) + ("" + ""))) {
                        case "a":
                            var v35 bool = (((3.14159 + 1e10) < 0.1) || false)
                            _ = v35
                        }
                    case -2147483648:
                        v36 := GenS4{F0: 33, F1: 42, F2: v8}
                        _ = v36
                    case 32768:
                        var v37 int16 = 0
                        _ = v37
                    default:
                        for fi38 := 0; fi38 < 14; fi38++ {
                            v27 = v27
                            _ = fi38
                        }
                    }
                    var v39 uint16 = 793
                    _ = v34
                    _ = v39
                } else {
                    if len(v7) >= 2 {
                        v7 = v7[1:len(v7)]
                    }
                    if ((v27 * 1) <= (v27 & (v27 * genFn0(v32, (v3 / ((v3) | 1)), ("\n\t" + "hello"))))) {
                        v40 := GenS4{F0: 127, F1: -367, F2: v8}
                        var v41 uint64 = uint64(v4)
                        v8 = (v8 >> (((v4 * (v4 ^ v4))) & 63))
                        _ = v40
                        _ = v41
                    }
                }
            }
            for fi42 := 0; fi42 < 3; fi42++ {
                delete(v2, 42)
                v43 := GenS5{F0: 158, F1: -1, F2: (v4 - v4)}
                switch (v17 / (((-(fi42))) | 1)) {
                case -892:
                    var v44 uint16 = 32767
                    _ = v44
                }
                _ = fi42
                _ = v43
            }
            _ = v33
        }
    case "a":
        var v45 int32 = (v8 / ((v8) | 1))
        _ = v45
    case "hello":
        v8 = 1
    }
    for fi46 := 0; fi46 < 4; fi46++ {
        v8 += (v8 / (((v8 - int32(v3))) | 1))
        v27 += (v27 * (v27 << ((v4) & 63)))
        fmt.Println(v4, v3)
        _ = fi46
    }
    var v47 float32 = (0.0 / 1.0)
    v48 := GenS3{F0: (^(v8)), F1: 100}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v47:", v47)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v17
    _ = v2
    _ = v27
    _ = v3
    _ = v32
    _ = v4
    _ = v47
    _ = v48
    _ = v7
    _ = v8
}
