package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8; F2 uint }
type GenS2 struct { F0 uint16; F1 int16 }
type GenS3 struct { F0 uint8; F1 int64 }
type GenS4 struct { F0 uint64; F1 uint; F2 int16 }
type GenS5 struct { F0 uint16 }
func genFn0(p0 int8) uint {
    v6 := []int{}
    if len(v6) > 0 {
        v6[0] = 46
    }
    _ = p0
    _ = v6
    return 18446744073709551615
}

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[uint32]uint32)
    var v3 int8 = 127
    v4 := []uint32{2, 2147483647, 812}
    var v5 uint16 = 312
    if (18446744073709551615 <= 65535) {
        var v7 float64 = ((0.0 / 1.0) + 3.14159)
        for fi8 := 0; fi8 < 12; fi8++ {
            v5 = (^(v5))
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            v9 := GenS1{F0: 255, F1: 0, F2: 42}
            _ = fi8
            _ = v9
        }
        _ = v7
    } else {
        switch "" {
        case "hello":
            var v10 uint = 407
            _ = v10
        }
        var v11 uint = 1
        _ = v11
    }
    v12 := GenS1{F0: 2, F1: 128, F2: 255}
    if false {
        v5 = v5
        var v13 float32 = (-2.25 - -2.25)
        _ = v13
    }
    if true {
        var v14 bool = ((float64(1e10) - (0.1 * 1.5)) <= 1.5)
        for fi15 := 0; fi15 < 3; fi15++ {
            v16 := GenS2{F0: (v5 & v5), F1: -347}
            v17 := GenS1{F0: 255, F1: 2, F2: 2}
            var v18 int64 = 0
            _ = fi15
            _ = v16
            _ = v17
            _ = v18
        }
        for fi19 := 0; fi19 < 13; fi19++ {
            fmt.Println(fi19)
            _ = fi19
        }
        _ = v14
    } else {
        for fi20 := 0; fi20 < 7; fi20++ {
            v21 := v1[0]
            _ = fi20
            _ = v21
        }
        if true {
            fmt.Println(v3, v5)
            var v22 uint64 = 1
            _ = v22
        } else {
            delete(v2, 763)
            for fi23 := 0; fi23 < 14; fi23++ {
                var v24 float64 = (1e10 * -2.25)
                v5 = (v5 | v5)
                if (1 == 1) {
                    v4 = append(v4, 2)
                }
                _ = fi23
                _ = v24
            }
        }
        for fi25 := 0; fi25 < 12; fi25++ {
            v3 += (v3 ^ (v3 ^ (v3 % ((-120) | 1))))
            _ = fi25
        }
    }
    switch "\n\t" {
    case "":
        v3 -= 100
    case "foo bar":
        v3 = v3
    default:
        if ((v5 >> ((782) & 63)) != v5) {
            v26 := GenS1{F0: 2, F1: 128, F2: 1}
            var v27 uint16 = 216
            _ = v26
            _ = v27
        }
    }
    switch "a" {
    case "\n\t":
        for fi28 := 0; fi28 < 16; fi28++ {
            var v29 int16 = 1
            var v30 int32 = 2147483647
            _ = fi28
            _ = v29
            _ = v30
        }
    default:
        if ((!(!false)) && true) {
            if false {
                var v31 int64 = 42
                _ = v31
            } else {
                switch 32768 {
                case 42:
                    var v32 uint16 = v5
                    _ = v32
                case 32768:
                    fmt.Println(v5, v5)
                }
                for fi33 := 0; fi33 < 11; fi33++ {
                    if len(v4) > 0 {
                        v4[0] = 42
                    }
                    if (!true) {
                        v5 = 255
                    }
                    _ = len(v1)
                    _ = fi33
                }
                v34 := v2[4294967295]
                _ = v34
            }
            for fi35 := 0; fi35 < 12; fi35++ {
                var v36 uint32 = 4294967295
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                _ = fi35
                _ = v36
            }
            switch ("hello" + "\n\t") {
            case "foo bar":
                var v37 uint64 = 2
                _ = v37
            case "a":
                if true {
                    switch 751 {
                    case 1:
                        if true {
                            var v38 uint16 = v5
                            for fi39 := 0; fi39 < 13; fi39++ {
                                var v40 uint16 = (v5 ^ (v38 >> ((2) & 63)))
                                _ = fi39
                                _ = v40
                            }
                            var v41 bool = ((!true) || false)
                            _ = v38
                            _ = v41
                        }
                    default:
                        for fi42 := 0; fi42 < 16; fi42++ {
                            v5 -= 1
                            switch "" {
                            case "hello":
                                for fi43 := 0; fi43 < 11; fi43++ {
                                    v3 -= v3
                                    _ = fi43
                                }
                            case "\n\t":
                                delete(v1, (fi42 % ((4294967295) | 1)))
                            }
                            fmt.Println(v3, fi42, v5)
                            _ = fi42
                        }
                    }
                    for fi44 := 0; fi44 < 4; fi44++ {
                        var v45 uint = 935
                        v46 := GenS1{F0: genFn0((v3 + (v3 << ((0) & 63)))), F1: 42, F2: (v45 * (v45 / (((^(v45))) | 1)))}
                        _ = fi44
                        _ = v45
                        _ = v46
                    }
                    for fi47 := 0; fi47 < 18; fi47++ {
                        switch (("foo bar" + "foo bar") + "") {
                        case "hello":
                            var v48 bool = (false || (false || false))
                            _ = v48
                        case "a":
                            switch "" {
                            case "hello":
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                            case "a":
                                for fi49 := 0; fi49 < 18; fi49++ {
                                    var v50 uint64 = 202
                                    fmt.Println(fi47, fi47, v3)
                                    _ = fi49
                                    _ = v50
                                }
                            }
                        }
                        var v51 uint = 1
                        _ = fi47
                        _ = v51
                    }
                }
            }
        } else {
            _ = len(v2)
            if false {
                v52 := GenS3{F0: 255, F1: -1}
                switch -813 {
                case 42:
                    var v53 int64 = 511
                    _ = v53
                case -9223372036854775808:
                    fmt.Println(v5, v3)
                default:
                    var v54 int16 = 29
                    _ = v54
                }
                var v55 int64 = -9223372036854775808
                _ = v52
                _ = v55
            } else {
                v5 = 1
                var v56 uint8 = 2
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                _ = v56
            }
            fmt.Println(v5, v5)
        }
    }
    switch "a" {
    case "hello":
        if ((0.1 + 0.1) != 0.1) {
            if (712 != 0) {
                if false {
                    var v57 bool = (0.1 != (float32(3.14159) / (0.1 + (0.0 / -2.25))))
                    var v58 string = "foo bar"
                    _ = v57
                    _ = v58
                } else {
                    v59 := v1[0]
                    if (2 == 14) {
                        switch 170 {
                        case 2147483647:
                            var v60 bool = v59
                            _ = v60
                        case 766:
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                        case 4294967295:
                            for fi61 := 0; fi61 < 10; fi61++ {
                                if (!(1.5 >= (1e10 / (1e10 + -2.25)))) {
                                    v62 := GenS3{F0: 41, F1: -432}
                                    v4 = append(v4, 65535)
                                    v3 += v3
                                    _ = v62
                                } else {
                                    v3 += 14
                                }
                                v3 = (v3 * -128)
                                delete(v1, 1)
                                _ = fi61
                            }
                        }
                    }
                    _ = v59
                }
                v3 = v3
                var v63 uint = 419
                _ = v63
            }
        } else {
            v5 = 42
            for fi64 := 0; fi64 < 3; fi64++ {
                v65 := GenS4{F0: 0, F1: 236, F2: 1}
                v66 := GenS5{F0: (v5 * 0)}
                _ = fi64
                _ = v65
                _ = v66
            }
            v3 += (v3 >> ((0) & 63))
        }
    case "":
        for fi67 := 0; fi67 < 5; fi67++ {
            v3 -= v3
            _ = fi67
        }
    case "a":
        for fi68 := 0; fi68 < 7; fi68++ {
            v69 := GenS4{F0: 746, F1: genFn0((v3 | int8(v5))), F2: 32767}
            for fi70 := 0; fi70 < 14; fi70++ {
                switch 9223372036854775807 {
                case 356:
                    if false {
                        v3 -= (v3 & (-(v3)))
                        v71 := GenS1{F0: 2, F1: 128, F2: 316}
                        if false {
                            delete(v2, 1)
                            var v72 uint32 = 298
                            _ = v72
                        } else {
                            delete(v2, 2)
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                            switch (^(fi68)) {
                            case -524:
                                var v73 string = (("\n\t" + "a") + "a")
                                _ = v73
                            case -987:
                                v3 = 0
                            default:
                                var v74 string = "a"
                                _ = v74
                            }
                        }
                        _ = v71
                    }
                case -284:
                    fmt.Println(v5)
                }
                _ = fi70
            }
            _ = fi68
            _ = v69
        }
    }
    fmt.Println(v3, v5)
    for fi75 := 0; fi75 < 1; fi75++ {
        v5 += v5
        _ = fi75
    }
    switch "a" {
    case "a":
        v76 := v1[42]
        _ = v76
    default:
        fmt.Println(v3)
    }
    _ = len(v4)
    var v77 int8 = v3
    var v78 int32 = 1
    fmt.Println(v78, v77)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v77:", v77)
    fmt.Println("v78:", v78)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v77
    _ = v78
}
