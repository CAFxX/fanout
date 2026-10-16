package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint }
type GenS2 struct { F0 uint16; F1 int; F2 uint }
type GenS3 struct { F0 uint; F1 int; F2 int32 }
type GenS4 struct { F0 int8; F1 int16 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 uint; F1 uint64; F2 bool }
type GenS7 struct { F0 uint16 }
type GenS8 struct { F0 uint }
type GenS9 struct { F0 int64; F1 bool }
type GenS10 struct { F0 uint32 }
func genFn0(p0 bool) uint64 {
    v7 := GenS1{F0: 655, F1: 42}
    for fi8 := 0; fi8 < 18; fi8++ {
        v9 := []int{0, (fi8 * (fi8 | (^(fi8))))}
        var v10 uint8 = 255
        _ = fi8
        _ = v10
        _ = v9
    }
    v11 := GenS1{F0: 2147483647, F1: 100}
    var v12 int16 = 32767
    for fi13 := 0; fi13 < 16; fi13++ {
        for fi14 := 0; fi14 < 13; fi14++ {
            var v15 int64 = 900
            v16 := []int{fi13, fi14, fi13}
            _ = fi14
            _ = v15
            _ = v16
        }
        if p0 {
            var v17 int = 629
            _ = v17
        } else {
            switch (^(fi13)) {
            case 4294967295:
                v18 := []int{0}
                _ = v18
            default:
                v19 := []int{fi13, (fi13 * fi13), (fi13 >> ((42) & 63))}
                _ = v19
            }
            switch "foo bar" {
            case "foo bar":
                switch "foo bar" {
                case "a":
                    var v20 float64 = 0.0
                    _ = v20
                case "foo bar":
                    v21 := []int{(fi13 + (-(fi13))), (fi13 / ((fi13) | 1)), (fi13 / ((fi13) | 1)), fi13}
                    _ = v21
                default:
                    var v22 bool = (v12 < (^(v12)))
                    _ = v22
                }
            case "hello":
                for fi23 := 0; fi23 < 18; fi23++ {
                    var v24 float64 = 0.0
                    _ = fi23
                    _ = v24
                }
            case "\n\t":
                v12 -= (v12 | 42)
            }
            switch ("" + "a") {
            case "foo bar":
                for fi25 := 0; fi25 < 4; fi25++ {
                    fmt.Println(p0, p0, p0)
                    if (65535 <= 807) {
                        switch (fi13 << ((233) & 63)) {
                        case 1:
                            p0 = false
                        case 247:
                            switch 4294967295 {
                            case 4294967295:
                                v26 := GenS2{F0: 0, F1: 42, F2: 671}
                                _ = v26
                            case 9223372036854775807:
                                v27 := GenS1{F0: 0, F1: 1}
                                _ = v27
                            case -16:
                                var v28 int8 = 100
                                _ = v28
                            }
                        }
                        switch fi25 {
                        case 9223372036854775807:
                            v29 := []int{fi25, 127}
                            _ = v29
                        case 32768:
                            v30 := GenS3{F0: 2, F1: fi25, F2: 0}
                            _ = v30
                        }
                    }
                    v12 = 32767
                    _ = fi25
                }
            default:
                v31 := GenS4{F0: 0, F1: v12}
                _ = v31
            }
        }
        var v32 int16 = (v12 >> ((598) & 63))
        _ = fi13
        _ = v32
    }
    _ = p0
    _ = v11
    _ = v12
    _ = v7
    return 0
}
func genFn1(p0 int64, p1 string) uint8 {
    var v33 float64 = float64(1.5)
    p0 += (-(p0))
    v34 := []int{-2147483648, -73}
    _ = p0
    _ = p1
    _ = v33
    _ = v34
    return 127
}

func main() {
    v1 := make(map[int]bool)
    var v2 float32 = 1e10
    var v3 int32 = -1
    var v4 float64 = 0.1
    v5 := make(map[uint32]string)
    var v6 int = 9223372036854775807
    if (v2 <= v2) {
        var v35 uint64 = 832
        _ = v35
    } else {
        for fi36 := 0; fi36 < 8; fi36++ {
            if (genFn0((!(234 < 42))) != 42) {
                v1[(^(fi36))] = false
                v37 := v5[2]
                _ = v37
            } else {
                v6 += 203
            }
            var v38 string = "a"
            v39 := GenS4{F0: 1, F1: 0}
            _ = fi36
            _ = v38
            _ = v39
        }
        v6 += v6
    }
    if true {
        switch ("\n\t" + "") {
        case "\n\t":
            v4 = (v4 - -2.25)
        case "hello":
            var v40 string = ""
            _ = v40
        default:
            var v41 uint8 = 0
            _ = v41
        }
        v2 += v2
        var v42 uint32 = 65535
        _ = v42
    } else {
        for fi43 := 0; fi43 < 9; fi43++ {
            var v44 int32 = (v3 % (((v3 / (((v3 >> ((326) & 63))) | 1))) | 1))
            v45 := []int{}
            _ = fi43
            _ = v44
            _ = v45
        }
        v46 := []int{}
        v47 := GenS2{F0: 567, F1: (v6 >> ((90) & 63)), F2: 391}
        _ = v46
        _ = v47
    }
    v48 := GenS5{F0: (("" + "hello") + ("a" + (("hello" + "\n\t") + "hello")))}
    fmt.Println(v2, v6)
    if true {
        v2 -= v2
    }
    for fi49 := 0; fi49 < 11; fi49++ {
        fmt.Println(v6, fi49)
        _ = fi49
    }
    v2 = ((0.1 / float32(-2.25)) * float32(v4))
    v4 = 3.14159
    v50 := GenS1{F0: (v3 ^ (v3 + 565)), F1: 18446744073709551615}
    var v51 uint16 = 586
    v6 = v6
    fmt.Println(v2)
    switch v6 {
    case 128:
        v52 := GenS6{F0: 0, F1: genFn0(false), F2: (1e10 == float64((1e10 / (v2 * v2))))}
        _ = v52
    default:
        var v53 uint64 = 1
        _ = v53
    }
    v54 := []int{v6, (v6 / ((v6) | 1))}
    for fi55 := 0; fi55 < 7; fi55++ {
        if true {
            fmt.Println(v2, fi55, v4)
            for fi56 := 0; fi56 < 18; fi56++ {
                switch "a" {
                case "":
                    v4 = v4
                case "\n\t":
                    v51 = v51
                case "a":
                    v2 = ((v2 * v2) * (1.5 / v2))
                default:
                    v57 := GenS4{F0: 42, F1: -32768}
                    _ = v57
                }
                var v58 int16 = 307
                if true {
                    v59 := GenS7{F0: 42}
                    _ = v59
                }
                _ = fi56
                _ = v58
            }
        } else {
            v60 := GenS4{F0: 127, F1: 979}
            fmt.Println(fi55)
            fmt.Println(v2)
            _ = v60
        }
        fmt.Println(v3, fi55)
        switch "foo bar" {
        case "\n\t":
            var v61 string = "foo bar"
            _ = v61
        case "a":
            if ((^(v51)) <= v51) {
                var v62 int32 = (^(v3))
                if len(v54) >= 2 {
                    v54 = v54[1:len(v54)]
                }
                _ = v62
            } else {
                if (((((0.0 * (float64(v2) / (((v4 + v4) - ((float64(v2) * ((((-2.25 * (v4 + v4)) / v4) - (((v4 / 1e10) - v4) / (float64(v2) * v4))) / v4)) / v4)) + 3.14159))) * v4) * (v4 - (float64(v2) + ((float64(v2) * -2.25) + v4)))) + ((float64((v2 * ((0.1 + v2) * 1.5))) * v4) + v4)) != (v4 / ((v4 / 1e10) / v4))) {
                    var v63 int16 = 314
                    var v64 float64 = (v4 / 1e10)
                    _ = v63
                    _ = v64
                }
                for fi65 := 0; fi65 < 1; fi65++ {
                    v2 -= 3.14159
                    _ = fi65
                }
                var v66 uint = 122
                _ = v66
            }
        case "foo bar":
            switch 342 {
            case -649:
                var v67 uint16 = 32767
                _ = v67
            case -750:
                fmt.Println(v51)
            default:
                fmt.Println(v4)
            }
        }
        _ = fi55
    }
    v3 = v3
    for fi68 := 0; fi68 < 13; fi68++ {
        switch fi68 {
        case 4294967295:
            if false {
                if false {
                    for fi69 := 0; fi69 < 6; fi69++ {
                        if false {
                            v70 := GenS8{F0: 4294967295}
                            v6 = fi69
                            fmt.Println(v51)
                            _ = v70
                        } else {
                            fmt.Println(v2, v6, v51)
                            delete(v5, 2147483647)
                            v6 = v6
                        }
                        var v71 uint32 = 239
                        _ = fi69
                        _ = v71
                    }
                    for fi72 := 0; fi72 < 11; fi72++ {
                        v1[9223372036854775807] = (959 == 42)
                        _ = fi72
                    }
                    var v73 float32 = v2
                    _ = v73
                }
            } else {
                v51 = (v51 % ((v51) | 1))
                var v74 int64 = 42
                _ = v74
            }
        case 0:
            var v75 float32 = (v2 * v2)
            _ = v75
        default:
            switch "" {
            case "a":
                _ = len(v54)
            case "foo bar":
                v76 := GenS9{F0: 0, F1: true}
                _ = v76
            default:
                if (false && true) {
                    var v77 int32 = (-(v3))
                    _ = v77
                } else {
                    v3 = v3
                    if (2 > 4294967295) {
                        for fi78 := 0; fi78 < 17; fi78++ {
                            if false {
                                if (127 <= -9223372036854775808) {
                                    fmt.Println(v2)
                                    v79 := GenS10{F0: 4294967295}
                                    _ = v79
                                }
                                switch "foo bar" {
                                case "\n\t":
                                    fmt.Println(fi78)
                                }
                            }
                            var v80 uint16 = v51
                            _ = fi78
                            _ = v80
                        }
                        switch ("a" + "\n\t") {
                        case "":
                            v81 := GenS4{F0: -116, F1: 1}
                            _ = v81
                        default:
                            v3 += -674
                        }
                        var v82 uint8 = 65
                        _ = v82
                    }
                }
            }
        }
        var v83 int64 = -1
        _ = fi68
        _ = v83
    }
    if true {
        for fi84 := 0; fi84 < 6; fi84++ {
            for fi85 := 0; fi85 < 1; fi85++ {
                switch ("" + "foo bar") {
                case "\n\t":
                    if (65535 >= 968) {
                        var v86 bool = (false && true)
                        var v87 int32 = (^(v3))
                        v4 = v4
                        v2 -= float32(v4)
                        _ = v86
                        _ = v87
                    }
                case "foo bar":
                    if len(v54) >= 2 {
                        v54 = v54[1:len(v54)]
                    }
                }
                _ = fi85
            }
            _ = len(v5)
            _ = fi84
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v51:", v51)
    fmt.Println("slicelen v54:", len(v54), v54)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v48
    _ = v5
    _ = v50
    _ = v51
    _ = v54
    _ = v6
}
