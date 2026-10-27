package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16 }
type GenS2 struct { F0 int16; F1 uint8; F2 string }
type GenS3 struct { F0 uint64; F1 bool }
type GenS4 struct { F0 string; F1 uint64 }
type GenS5 struct { F0 int64; F1 int16; F2 int16 }
type GenS6 struct { F0 int32; F1 bool }
func genFn0(p0 uint64, p1 uint64) uint32 {
    var v6 uint = 1
    switch 2147483647 {
    case -424:
        var v7 uint16 = 1
        _ = v7
    case 1:
        var v8 int64 = 0
        _ = v8
    case -791:
        var v9 int32 = int32(v6)
        _ = v9
    }
    v10 := GenS1{F0: 422, F1: -1}
    p0 += (p1 | p0)
    if true {
        for fi11 := 0; fi11 < 14; fi11++ {
            for fi12 := 0; fi12 < 15; fi12++ {
                v13 := []int{-1, (fi11 & (fi12 >> (((v6 + v6)) & 63))), -1}
                switch 2147483647 {
                case -727:
                    if false {
                        fmt.Println(fi11, fi11, p0)
                        fmt.Println(v6, p0)
                        if len(v13) >= 2 {
                            v13 = v13[1:len(v13)]
                        }
                    }
                }
                if len(v13) > 0 {
                    v13[0] = fi11
                }
                _ = fi12
                _ = v13
            }
            _ = fi11
        }
    }
    _ = p0
    _ = p1
    _ = v10
    _ = v6
    return 4294967295
}
func genFn1(p0 uint, p1 int, p2 int8) int {
    for fi14 := 0; fi14 < 10; fi14++ {
        for fi15 := 0; fi15 < 14; fi15++ {
            var v16 uint = p0
            v17 := GenS2{F0: -719, F1: 136, F2: ("foo bar" + ((("\n\t" + "a") + ("hello" + ("foo bar" + (("hello" + ("\n\t" + "foo bar")) + "hello")))) + ""))}
            _ = fi15
            _ = v16
            _ = v17
        }
        fmt.Println(p2)
        var v18 int8 = p2
        _ = fi14
        _ = v18
    }
    switch (("" + (("a" + "\n\t") + "\n\t")) + ("" + "\n\t")) {
    case "a":
        fmt.Println(p1, p0, p1)
    case "foo bar":
        p0 += (p0 >> (((p0 >> (((^(p0))) & 63))) & 63))
    }
    switch "hello" {
    case "a":
        for fi19 := 0; fi19 < 11; fi19++ {
            fmt.Println(p2)
            _ = fi19
        }
    }
    for fi20 := 0; fi20 < 14; fi20++ {
        v21 := GenS3{F0: 18446744073709551615, F1: (!(42 < 0))}
        _ = fi20
        _ = v21
    }
    _ = p0
    _ = p1
    _ = p2
    return (p1 + (p1 & (p1 >> ((p0) & 63))))
}

func main() {
    v1 := make(map[int]uint32)
    var v2 int16 = 42
    var v3 int = 764
    var v4 string = "foo bar"
    var v5 float32 = ((0.1 - 0.0) * -2.25)
    var v22 uint64 = 18446744073709551615
    for fi23 := 0; fi23 < 19; fi23++ {
        var v24 uint64 = 1
        _ = fi23
        _ = v24
    }
    var v25 int64 = -9223372036854775808
    fmt.Println(v22, v5)
    switch (-(v3)) {
    case 9223372036854775807:
        switch v3 {
        case 32768:
            var v26 int64 = (v25 ^ v25)
            _ = v26
        case 1:
            v1[genFn1(0, (v3 & 9223372036854775807), 1)] = 2147483648
        case 701:
            if (425 == 42) {
                switch "hello" {
                case "hello":
                    v27 := GenS3{F0: v22, F1: (!true)}
                    _ = v27
                case "\n\t":
                    switch (v3 / ((128) | 1)) {
                    case 12:
                        v28 := GenS4{F0: (v4 + "hello"), F1: v22}
                        _ = v28
                    default:
                        if ((v3 & int(v22)) < v3) {
                            var v29 int32 = 42
                            for fi30 := 0; fi30 < 12; fi30++ {
                                for fi31 := 0; fi31 < 17; fi31++ {
                                    var v32 int64 = (v25 * (v25 >> ((346) & 63)))
                                    fmt.Println(fi30, v4, v29)
                                    _ = fi31
                                    _ = v32
                                }
                                _ = fi30
                            }
                            switch v3 {
                            case 351:
                                var v33 bool = (true && ((v2 >> ((1) & 63)) == v2))
                                _ = v33
                            case 321:
                                v2 = (-(v2))
                            case -2147483648:
                                v34 := []int{v3, (v3 + genFn1(35, -2147483648, 127)), (-(v3))}
                                _ = v34
                            }
                            _ = v29
                        }
                    }
                case "a":
                    for fi35 := 0; fi35 < 7; fi35++ {
                        for fi36 := 0; fi36 < 7; fi36++ {
                            switch (fi35 + fi35) {
                            case -517:
                                switch (fi36 | fi36) {
                                case -528:
                                    fmt.Println(v5, v22, fi36)
                                default:
                                    fmt.Println(v3, fi35)
                                }
                            }
                            v37 := GenS2{F0: v2, F1: 42, F2: v4}
                            _ = fi36
                            _ = v37
                        }
                        _ = fi35
                    }
                default:
                    if false {
                        var v38 float32 = v5
                        v4 = "\n\t"
                        _ = v38
                    }
                }
                v4 = v4
                v4 = "a"
            } else {
                fmt.Println(v5, v25)
                fmt.Println(v2, v22, v25)
            }
        default:
            v39 := GenS5{F0: (-(v25)), F1: v2, F2: (-(v2))}
            _ = v39
        }
    }
    var v40 string = "a"
    var v41 uint8 = 17
    delete(v1, (v3 << ((255) & 63)))
    var v42 string = v40
    v2 -= v2
    v43 := GenS3{F0: 896, F1: (!(true && (-128 <= 1)))}
    var v44 int = genFn1(65535, v3, 100)
    if false {
        switch 2147483647 {
        case 93:
            for fi45 := 0; fi45 < 18; fi45++ {
                fmt.Println(v2)
                _ = fi45
            }
        case -1:
            fmt.Println(v4)
        case 4294967295:
            fmt.Println(v5)
        default:
            var v46 uint64 = uint64(v41)
            _ = v46
        }
        fmt.Println(v44, v2)
    } else {
        switch "" {
        case "a":
            v44 += 9223372036854775807
        case "hello":
            if (1.5 != 0.1) {
                switch "\n\t" {
                case "foo bar":
                    fmt.Println(v22)
                default:
                    v42 = "a"
                }
                switch v42 {
                case "foo bar":
                    for fi47 := 0; fi47 < 6; fi47++ {
                        var v48 int32 = int32(v44)
                        for fi49 := 0; fi49 < 20; fi49++ {
                            v40 = "foo bar"
                            _ = fi49
                        }
                        _ = fi47
                        _ = v48
                    }
                case "\n\t":
                    fmt.Println(v40, v44, v2)
                }
            }
        case "foo bar":
            if (!(0 >= 0)) {
                var v50 uint = 42
                for fi51 := 0; fi51 < 15; fi51++ {
                    v44 -= (v3 & fi51)
                    var v52 uint = (v50 ^ (^(v50)))
                    var v53 int64 = v25
                    _ = fi51
                    _ = v52
                    _ = v53
                }
                var v54 uint8 = (v41 & (^(v41)))
                _ = v50
                _ = v54
            } else {
                v55 := GenS6{F0: -199, F1: false}
                v40 = v40
                _ = v55
            }
        }
        v40 = ((v42 + (((v40 + (v42 + "")) + ("a" + (v40 + v42))) + v42)) + ("\n\t" + v42))
    }
    for fi56 := 0; fi56 < 10; fi56++ {
        if false {
            v57 := GenS2{F0: v2, F1: v41, F2: (v42 + ((v42 + "a") + v42))}
            _ = v57
        }
        if (true && (70 != v25)) {
            v41 = 0
            var v58 uint8 = 0
            _ = v58
        } else {
            if (3.14159 <= 0.0) {
                for fi59 := 0; fi59 < 4; fi59++ {
                    fmt.Println(v40, fi56)
                    var v60 bool = (!false)
                    _ = fi59
                    _ = v60
                }
                for fi61 := 0; fi61 < 17; fi61++ {
                    if false {
                        v5 = ((v5 / ((1e10 / v5) / 1.5)) + v5)
                        var v62 uint = 598
                        _ = v62
                    }
                    if (float64(((1.5 - v5) / (float32(float64(-2.25)) / v5))) != (3.14159 / float64(1e10))) {
                        if (v22 != 4294967295) {
                            v63 := []int{32768, -2147483648}
                            v2 = (v2 & v2)
                            v25 += (v25 / ((-1) | 1))
                            _ = v63
                        } else {
                            v64 := v1[(fi56 * fi61)]
                            _ = v64
                        }
                        switch 4294967295 {
                        case 127:
                            v41 = uint8(v22)
                        default:
                            fmt.Println(v41, v40, v42)
                        }
                        fmt.Println(v4, v4)
                        var v65 string = ("a" + v42)
                        var v66 float64 = (1e10 * float64((v5 * 0.1)))
                        _ = v65
                        _ = v66
                    } else {
                        v67 := []int{(fi56 % (((fi56 ^ (v3 | (v44 & v3)))) | 1)), 4294967295, 440}
                        fmt.Println(v25, v41, v22)
                        v25 = v25
                        switch 826 {
                        case -469:
                            fmt.Println(v5)
                        case 484:
                            var v68 float64 = 0.0
                            _ = v68
                        }
                        v44 += 672
                        if false {
                            v41 -= 165
                        }
                        switch v40 {
                        case "a":
                            v67 = append(v67, 231)
                            switch v3 {
                            case 0:
                                if len(v67) >= 2 {
                                    v67 = v67[1:len(v67)]
                                }
                            case 263:
                                var v69 uint8 = v41
                                _ = v69
                            }
                            switch (fi56 >> ((42) & 63)) {
                            case 9223372036854775807:
                                if true {
                                    v5 -= v5
                                    v5 = 0.0
                                    var v70 uint = 65535
                                    _ = v70
                                }
                            }
                            v25 -= (v25 % ((-869) | 1))
                            v5 -= ((1.5 + (v5 * ((v5 * ((float32(0.0) - v5) / (-2.25 * -2.25))) * v5))) * 3.14159)
                            v2 += 1000
                        case "hello":
                            v4 = v40
                        default:
                            v5 = (float32(((0.1 * float64((1.5 * (v5 / v5)))) - float64((0.0 + v5)))) - v5)
                        }
                        var v71 uint8 = v41
                        _ = v67
                        _ = v71
                    }
                    _ = fi61
                }
                for fi72 := 0; fi72 < 17; fi72++ {
                    if (v5 > (0.1 / 3.14159)) {
                        var v73 uint = 0
                        _ = v73
                    } else {
                        fmt.Println(v2, v2, fi72)
                        switch (fi56 - (-(fi72))) {
                        case 32768:
                            switch ((v42 + v4) + (("hello" + "\n\t") + "hello")) {
                            case "hello":
                                if true {
                                    _ = len(v1)
                                }
                            }
                        }
                    }
                    var v74 uint32 = 65535
                    _ = fi72
                    _ = v74
                }
            } else {
                var v75 string = (v42 + v42)
                var v76 string = v75
                _ = v75
                _ = v76
            }
        }
        _ = fi56
    }
    var v77 uint64 = 0
    v4 = "\n\t"
    v78 := []int{v44, 42}
    for fi79 := 0; fi79 < 2; fi79++ {
        for fi80 := 0; fi80 < 8; fi80++ {
            v42 = ((v40 + v42) + "foo bar")
            _ = fi80
        }
        v44 = -9223372036854775808
        _ = fi79
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v44:", v44)
    fmt.Println("v5:", v5)
    fmt.Println("v77:", v77)
    fmt.Println("slicelen v78:", len(v78), v78)
    _ = v1
    _ = v2
    _ = v22
    _ = v25
    _ = v3
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v43
    _ = v44
    _ = v5
    _ = v77
    _ = v78
}
