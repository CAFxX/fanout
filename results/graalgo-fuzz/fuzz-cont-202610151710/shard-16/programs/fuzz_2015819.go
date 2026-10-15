package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int16 }
type GenS2 struct { F0 uint16; F1 bool }
type GenS3 struct { F0 uint; F1 int64 }
type GenS4 struct { F0 bool; F1 uint16 }
type GenS5 struct { F0 uint64; F1 string; F2 int }
type GenS6 struct { F0 int16; F1 uint; F2 bool }
type GenS7 struct { F0 string }
func genFn0() uint64 {
    v7 := []int{-2147483648, 32768, 455, 126}
    _ = len(v7)
    switch "a" {
    case "foo bar":
        var v8 uint = 0
        _ = v8
    case "\n\t":
        var v9 uint8 = 2
        _ = v9
    }
    switch 9223372036854775807 {
    case -2147483648:
        if false {
            if (-128 <= 42) {
                v7 = append(v7, -9223372036854775808)
                switch "\n\t" {
                case "hello":
                    var v10 bool = ((true || ((!false) && (668 >= -9223372036854775808))) && true)
                    _ = v10
                case "":
                    var v11 uint64 = 42
                    _ = v11
                }
                var v12 int8 = 0
                _ = v12
            }
            var v13 float64 = 3.14159
            if true {
                for fi14 := 0; fi14 < 5; fi14++ {
                    if (-1 <= 127) {
                        _ = len(v7)
                    }
                    if false {
                        switch "foo bar" {
                        case "a":
                            for fi15 := 0; fi15 < 9; fi15++ {
                                var v16 string = ("a" + "hello")
                                _ = fi15
                                _ = v16
                            }
                        }
                        for fi17 := 0; fi17 < 1; fi17++ {
                            v13 -= v13
                            var v18 int8 = 0
                            if false {
                                _ = len(v7)
                                v13 = 0.0
                            }
                            _ = fi17
                            _ = v18
                        }
                    }
                    _ = fi14
                }
                var v19 int16 = 59
                _ = v19
            } else {
                v20 := GenS1{F0: 65535, F1: 485}
                if true {
                    fmt.Println(v13)
                    v21 := GenS2{F0: 1, F1: (!false)}
                    _ = v21
                } else {
                    var v22 int64 = 9223372036854775807
                    _ = v22
                }
                fmt.Println(v13)
                _ = v20
            }
            _ = v13
        } else {
            switch "foo bar" {
            case "a":
                if true {
                    v23 := GenS3{F0: 356, F1: 42}
                    var v24 bool = true
                    if v24 {
                        v25 := GenS2{F0: 382, F1: (false && (v24 || (float32((-2.25 - (3.14159 / (0.1 / 0.1)))) >= (float32(float64((1.5 + 1e10))) * 3.14159))))}
                        v26 := GenS1{F0: 2147483647, F1: 333}
                        v27 := GenS1{F0: 4294967295, F1: -1}
                        _ = v25
                        _ = v26
                        _ = v27
                    }
                    _ = v23
                    _ = v24
                }
            }
            if true {
                var v28 uint32 = 2147483647
                var v29 string = "\n\t"
                _ = v28
                _ = v29
            } else {
                var v30 int32 = -1
                var v31 uint = 447
                v32 := GenS1{F0: 2, F1: 1000}
                _ = v30
                _ = v31
                _ = v32
            }
            if true {
                for fi33 := 0; fi33 < 20; fi33++ {
                    fmt.Println(fi33)
                    var v34 int8 = -45
                    v7 = append(v7, fi33)
                    _ = fi33
                    _ = v34
                }
                var v35 uint32 = 4294967295
                _ = v35
            }
        }
    case -398:
        var v36 bool = (float32((-2.25 - (0.1 - (3.14159 + (1.5 / 3.14159))))) == (0.1 * ((1e10 + (float32((float64(3.14159) * 3.14159)) - 1e10)) * (float32(0.1) + (0.1 / 0.1)))))
        _ = v36
    default:
        var v37 int32 = 1
        _ = v37
    }
    var v38 int16 = 800
    _ = v38
    _ = v7
    return 4294967295
}
func genFn1(p0 string, p1 int64) int32 {
    var v39 float64 = float64(3.14159)
    var v40 uint8 = 255
    for fi41 := 0; fi41 < 17; fi41++ {
        for fi42 := 0; fi42 < 18; fi42++ {
            v43 := GenS4{F0: true, F1: 223}
            _ = fi42
            _ = v43
        }
        if true {
            v44 := []int{9223372036854775807}
            _ = v44
        } else {
            v40 += (v40 + (v40 & 131))
            for fi45 := 0; fi45 < 6; fi45++ {
                var v46 int64 = (p1 & (p1 << ((42) & 63)))
                _ = fi45
                _ = v46
            }
            switch p0 {
            case "foo bar":
                v39 = (v39 / (v39 - (0.0 + v39)))
            }
        }
        switch fi41 {
        case 375:
            v47 := GenS1{F0: 2, F1: 1000}
            _ = v47
        case 32767:
            v48 := []int{fi41, fi41, 9223372036854775807}
            _ = v48
        case 9223372036854775807:
            v49 := GenS1{F0: 940, F1: -1}
            _ = v49
        }
        _ = fi41
    }
    fmt.Println(p1)
    _ = p0
    _ = p1
    _ = v39
    _ = v40
    return 754
}

func main() {
    var v1 bool = false
    var v2 int16 = -990
    var v3 float64 = (3.14159 * (((3.14159 * (1e10 + (3.14159 - float64((-2.25 * 3.14159))))) / 1.0) * (float64(-2.25) - 0.0)))
    v4 := make(map[uint32]uint32)
    var v5 uint16 = 65535
    v6 := []uint32{345, 166, 65535, 4294967295}
    delete(v4, 2147483648)
    v50 := v4[4294967295]
    for fi51 := 0; fi51 < 15; fi51++ {
        if (42 < 29) {
            if (1 <= 42) {
                v6 = append(v6, v50)
                v52 := GenS5{F0: genFn0(), F1: ("hello" + "a"), F2: 2147483647}
                for fi53 := 0; fi53 < 1; fi53++ {
                    var v54 uint = 759
                    _ = len(v4)
                    if ((((v3 * v3) == ((3.14159 * (v3 - -2.25)) * v3)) || v1) && (int64(v54) <= 0)) {
                        v5 += v5
                    } else {
                        switch "" {
                        case "a":
                            var v55 int64 = 1
                            _ = v55
                        }
                        fmt.Println(v54, v54, v3)
                        var v56 int64 = 0
                        switch "\n\t" {
                        case "hello":
                            fmt.Println(v5)
                        case "a":
                            if len(v6) > 0 {
                                v6[0] = (v50 >> ((4294967295) & 63))
                            }
                        case "":
                            if v1 {
                                switch fi51 {
                                case 906:
                                    fmt.Println(v3, v54, v5)
                                }
                                _ = len(v6)
                            } else {
                                v50 -= 559
                                var v57 uint32 = v50
                                delete(v4, v57)
                                _ = v57
                            }
                        default:
                            var v58 uint32 = 2147483648
                            _ = v58
                        }
                        if (!(3.14159 == -2.25)) {
                            if v1 {
                                var v59 uint32 = (v50 % ((42) | 1))
                                v59 -= 42
                                _ = v59
                            }
                        } else {
                            fmt.Println(v5, v2)
                            v60 := GenS5{F0: 0, F1: "a", F2: fi51}
                            _ = v60
                        }
                        _ = v56
                    }
                    _ = fi53
                    _ = v54
                }
                _ = v52
            }
            switch 32768 {
            case 550:
                v61 := GenS6{F0: 32767, F1: 1, F2: true}
                _ = v61
            case 32767:
                _ = len(v6)
            }
        } else {
            switch fi51 {
            case 633:
                for fi62 := 0; fi62 < 2; fi62++ {
                    if len(v6) > 0 {
                        v6[0] = (v50 % ((v50) | 1))
                    }
                    _ = fi62
                }
            case 225:
                if v1 {
                    v63 := GenS1{F0: v50, F1: v2}
                    _ = v63
                }
            default:
                var v64 uint64 = 18446744073709551615
                _ = v64
            }
            var v65 int64 = 9223372036854775807
            _ = v65
        }
        fmt.Println(fi51)
        _ = fi51
    }
    v66 := GenS5{F0: 558, F1: "hello", F2: -837}
    for fi67 := 0; fi67 < 10; fi67++ {
        switch 2147483647 {
        case -2147483648:
            fmt.Println(fi67, fi67)
        case 1:
            switch ("" + "\n\t") {
            case "":
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
            default:
                v5 = (v5 & (v5 << ((2) & 63)))
            }
        case 4294967295:
            v2 = 0
        default:
            if (-31 >= -128) {
                v68 := GenS7{F0: ("a" + (("\n\t" + (((("hello" + "a") + "a") + "\n\t") + "a")) + "hello"))}
                fmt.Println(v2, v1, v2)
                _ = v68
            } else {
                if (!(128 > 36)) {
                    var v69 int32 = genFn1("foo bar", -9223372036854775808)
                    if len(v6) > 0 {
                        v6[0] = (v50 & 2147483647)
                    }
                    v70 := GenS5{F0: 679, F1: ("a" + "a"), F2: 128}
                    _ = v69
                    _ = v70
                }
                _ = len(v4)
                for fi71 := 0; fi71 < 18; fi71++ {
                    v72 := GenS6{F0: 1000, F1: 93, F2: false}
                    _ = fi71
                    _ = v72
                }
            }
        }
        if (!(v1 && (1 >= v50))) {
            var v73 int8 = -6
            v74 := GenS6{F0: v2, F1: uint(v5), F2: (true && false)}
            _ = len(v4)
            _ = v73
            _ = v74
        } else {
            var v75 uint = 2
            _ = v75
        }
        _ = fi67
    }
    var v76 uint16 = (v5 & 893)
    v1 = v1
    v76 += (v76 % (((v5 / ((744) | 1))) | 1))
    var v77 string = "foo bar"
    var v78 int64 = 0
    if v1 {
        var v79 int32 = 42
        var v80 int32 = -1
        v6 = append(v6, (v50 ^ v50))
        _ = v79
        _ = v80
    } else {
        switch ("hello" + v77) {
        case "a":
            v1 = true
        case "\n\t":
            v78 = v78
        case "hello":
            if v1 {
                v81 := GenS4{F0: v1, F1: v76}
                _ = v81
            } else {
                if true {
                    if len(v6) > 0 {
                        v6[0] = v50
                    }
                    v50 -= (v50 | (v50 * (v50 % (((v50 ^ v50)) | 1))))
                    if false {
                        for fi82 := 0; fi82 < 4; fi82++ {
                            v77 = ((("a" + ("foo bar" + v77)) + "\n\t") + v77)
                            v1 = v1
                            v83 := GenS7{F0: (v77 + ("hello" + (v77 + (v77 + v77))))}
                            _ = fi82
                            _ = v83
                        }
                        v3 = 0.0
                    }
                } else {
                    for fi84 := 0; fi84 < 20; fi84++ {
                        v78 = v78
                        if (1 >= 0) {
                            var v85 int = (fi84 / ((fi84) | 1))
                            for fi86 := 0; fi86 < 16; fi86++ {
                                v85 = fi86
                                var v87 uint8 = 128
                                _ = fi86
                                _ = v87
                            }
                            _ = v85
                        } else {
                            v88 := GenS5{F0: uint64(v76), F1: "a", F2: fi84}
                            switch (fi84 & (fi84 ^ (fi84 ^ fi84))) {
                            case 142:
                                v3 = v3
                            case 127:
                                if len(v6) >= 2 {
                                    v6 = v6[1:len(v6)]
                                }
                            }
                            var v89 int8 = 0
                            _ = v88
                            _ = v89
                        }
                        if false {
                            v90 := v4[76]
                            _ = v90
                        } else {
                            fmt.Println(v5)
                            switch int(v50) {
                            case 645:
                                fmt.Println(v78, v76)
                            }
                            var v91 int16 = v2
                            _ = v91
                        }
                        _ = fi84
                    }
                    v77 = ((v77 + "a") + (v77 + ""))
                    for fi92 := 0; fi92 < 10; fi92++ {
                        var v93 int16 = (-(v2))
                        _ = fi92
                        _ = v93
                    }
                }
            }
        }
    }
    v94 := GenS7{F0: (v77 + v77)}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v76:", v76)
    fmt.Println("v77:", v77)
    fmt.Println("v78:", v78)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v50
    _ = v6
    _ = v66
    _ = v76
    _ = v77
    _ = v78
    _ = v94
}
