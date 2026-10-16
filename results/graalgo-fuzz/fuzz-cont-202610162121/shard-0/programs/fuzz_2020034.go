package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint }
type GenS2 struct { F0 uint32; F1 int8; F2 uint8 }
type GenS3 struct { F0 uint64; F1 uint; F2 bool }
type GenS4 struct { F0 uint8; F1 int8; F2 uint64 }
type GenS5 struct { F0 uint32; F1 string; F2 int16 }
type GenS6 struct { F0 uint; F1 uint32 }
type GenS7 struct { F0 string; F1 int64 }
type GenS8 struct { F0 int8 }
func genFn0(p0 int32) uint8 {
    v4 := GenS1{F0: false, F1: 2}
    for fi5 := 0; fi5 < 13; fi5++ {
        var v6 float32 = (3.14159 - float32(1.5))
        if (!((false || (2147483647 != 585)) && true)) {
            for fi7 := 0; fi7 < 4; fi7++ {
                switch fi5 {
                case -278:
                    fmt.Println(v6, fi7)
                case 378:
                    var v8 string = "a"
                    _ = v8
                default:
                    var v9 uint8 = 1
                    _ = v9
                }
                _ = fi7
            }
            v6 = 0.1
            v6 += v6
        } else {
            for fi10 := 0; fi10 < 1; fi10++ {
                var v11 uint32 = 126
                var v12 float32 = ((float32(1.5) + (v6 / 1e10)) / -2.25)
                v13 := GenS1{F0: (1 < 100), F1: 42}
                _ = fi10
                _ = v11
                _ = v12
                _ = v13
            }
            switch "hello" {
            case "":
                if true {
                    v6 = v6
                    p0 = (p0 + p0)
                }
            case "a":
                var v14 uint32 = 1
                _ = v14
            case "foo bar":
                var v15 int32 = -451
                _ = v15
            default:
                switch "foo bar" {
                case "hello":
                    p0 += (p0 << ((65535) & 63))
                case "a":
                    var v16 string = (("\n\t" + (("" + "") + ("foo bar" + ("a" + ("\n\t" + "a"))))) + "foo bar")
                    _ = v16
                default:
                    v6 -= (v6 * v6)
                }
            }
            for fi17 := 0; fi17 < 9; fi17++ {
                v18 := GenS2{F0: 65535, F1: -1, F2: 255}
                var v19 uint16 = 65535
                _ = fi17
                _ = v18
                _ = v19
            }
        }
        _ = fi5
        _ = v6
    }
    v20 := []int{127}
    var v21 uint8 = 128
    _ = p0
    _ = v20
    _ = v21
    _ = v4
    return (v21 % (((v21 >> ((914) & 63))) | 1))
}
func genFn1(p0 int64, p1 uint64) int8 {
    fmt.Println(p1)
    fmt.Println(p0, p1)
    var v22 int8 = -1
    for fi23 := 0; fi23 < 15; fi23++ {
        for fi24 := 0; fi24 < 10; fi24++ {
            v25 := GenS1{F0: (true && false), F1: 18446744073709551615}
            _ = fi24
            _ = v25
        }
        _ = fi23
    }
    _ = p0
    _ = p1
    _ = v22
    return v22
}

func main() {
    v1 := make(map[string]bool)
    var v2 int = 32768
    var v3 uint16 = 2
    var v26 uint8 = 128
    var v27 int16 = -446
    v28 := GenS3{F0: 0, F1: 255, F2: (-128 != -1)}
    for fi29 := 0; fi29 < 1; fi29++ {
        if (1.5 != (1.5 - 0.1)) {
            v27 = 33
            var v30 int64 = 59
            for fi31 := 0; fi31 < 8; fi31++ {
                v2 += fi31
                var v32 int8 = 70
                _ = fi31
                _ = v32
            }
            _ = v30
        } else {
            var v33 uint32 = uint32(v3)
            var v34 uint8 = v26
            if true {
                fmt.Println(fi29)
            } else {
                var v35 int = (fi29 ^ fi29)
                _ = v35
            }
            _ = v33
            _ = v34
        }
        var v36 int8 = 0
        _ = fi29
        _ = v36
    }
    switch int(v26) {
    case 4294967295:
        v37 := []int{v2, (v2 + 42), -72}
        _ = v37
    }
    v38 := GenS2{F0: 0, F1: 27, F2: genFn0(42)}
    if true {
        if (42 <= 890) {
            v3 -= v3
            if (v2 != (v2 >> ((0) & 63))) {
                for fi39 := 0; fi39 < 6; fi39++ {
                    v40 := GenS1{F0: false, F1: 999}
                    _ = fi39
                    _ = v40
                }
                var v41 int64 = 9223372036854775807
                if true {
                    for fi42 := 0; fi42 < 15; fi42++ {
                        for fi43 := 0; fi43 < 6; fi43++ {
                            v26 += 255
                            var v44 uint = 984
                            _ = fi43
                            _ = v44
                        }
                        v45 := []int{(fi42 * v2)}
                        v46 := GenS1{F0: (!(((0.1 * 0.0) + (1e10 / 1.5)) == 0.0)), F1: 1}
                        _ = fi42
                        _ = v45
                        _ = v46
                    }
                    var v47 uint64 = 18446744073709551615
                    _ = v47
                } else {
                    v27 += (v27 - (v27 % ((42) | 1)))
                }
                _ = v41
            } else {
                if false {
                    var v48 int64 = 42
                    v49 := []int{32768, -1, 2147483647}
                    fmt.Println(v48, v48, v26)
                    _ = v48
                    _ = v49
                }
                v26 = 75
                if false {
                    if false {
                        v50 := []int{(v2 << ((1) & 63)), (v2 ^ v2), (v2 % ((492) | 1))}
                        fmt.Println(v2, v3, v3)
                        _ = v50
                    } else {
                        var v51 int32 = 42
                        switch v2 {
                        case 549:
                            v51 = (v51 % ((v51) | 1))
                        case -9223372036854775808:
                            v27 = v27
                        case -764:
                            v52 := GenS1{F0: false, F1: 42}
                            _ = v52
                        default:
                            v53 := []int{(v2 / ((-2147483648) | 1)), v2, (v2 | (v2 << ((4294967295) & 63))), -1000}
                            _ = v53
                        }
                        v54 := GenS2{F0: 65535, F1: -128, F2: v26}
                        _ = v51
                        _ = v54
                    }
                    if (!(-1 >= -2147483648)) {
                        v26 -= 218
                        v27 = v27
                    }
                    if true {
                        var v55 int16 = v27
                        for fi56 := 0; fi56 < 16; fi56++ {
                            for fi57 := 0; fi57 < 5; fi57++ {
                                fmt.Println(v3)
                                if true {
                                    var v58 uint64 = 393
                                    var v59 int64 = 1
                                    var v60 uint64 = 0
                                    _ = v58
                                    _ = v59
                                    _ = v60
                                } else {
                                    switch 1 {
                                    case 0:
                                        v61 := GenS3{F0: 0, F1: 63, F2: false}
                                        _ = v61
                                    case -2147483648:
                                        if false {
                                            var v62 uint = 371
                                            if false {
                                                v55 += v55
                                            } else {
                                                var v63 int16 = (v27 & v27)
                                                var v64 uint32 = 104
                                                if (true || (false || (fi56 > int(v55)))) {
                                                    var v65 int16 = int16(v2)
                                                    _ = v65
                                                } else {
                                                    v1[("" + ("\n\t" + ""))] = false
                                                }
                                                _ = v63
                                                _ = v64
                                            }
                                            v2 -= (fi56 & -2147483648)
                                            _ = v62
                                        } else {
                                            v55 = -1
                                        }
                                    default:
                                        fmt.Println(v3)
                                        if (v3 == v3) {
                                            if true {
                                                var v66 float32 = (float32(1e10) / (-2.25 / 1.0))
                                                v67 := GenS4{F0: (v26 - (^(v26))), F1: -128, F2: 4294967295}
                                                var v68 float64 = 0.1
                                                _ = v66
                                                _ = v67
                                                _ = v68
                                            } else {
                                                switch fi56 {
                                                case 32768:
                                                    v26 = (v26 & (v26 ^ (v26 % ((v26) | 1))))
                                                case 2147483647:
                                                    v55 -= v55
                                                default:
                                                    switch fi57 {
                                                    case -9223372036854775808:
                                                        var v69 int32 = 1
                                                        _ = v69
                                                    case 32767:
                                                        v26 = (v26 / ((v26) | 1))
                                                    }
                                                }
                                                var v70 bool = ((v3 << ((0) & 63)) != (^(v3)))
                                                v71 := []int{fi57, (v2 >> ((255) & 63)), 127}
                                                if ((!(!v70)) || v70) {
                                                    v2 = (^(fi56))
                                                    v72 := GenS4{F0: (v26 * (^(v26))), F1: 100, F2: 273}
                                                    _ = v72
                                                } else {
                                                    v26 += 255
                                                    v26 -= v26
                                                }
                                                if (!(0.1 < 3.14159)) {
                                                    v70 = (("\n\t" + ("a" + ("a" + "foo bar"))) >= ("" + ("" + "\n\t")))
                                                } else {
                                                    if (-9223372036854775808 == 9223372036854775807) {
                                                        v27 = (v27 >> ((4294967295) & 63))
                                                        switch (fi57 << ((151) & 63)) {
                                                        case 32768:
                                                            delete(v1, "foo bar")
                                                        case 42:
                                                            v2 += fi57
                                                        case 820:
                                                            v55 -= (^(v55))
                                                        }
                                                        var v73 int8 = 23
                                                        v74 := GenS5{F0: 0, F1: ((("hello" + ("foo bar" + "a")) + "\n\t") + ("a" + "\n\t")), F2: (v55 ^ -756)}
                                                        _ = v73
                                                        _ = v74
                                                    } else {
                                                        var v75 int32 = 0
                                                        delete(v1, ((("a" + ("" + "hello")) + "\n\t") + ("a" + ("\n\t" + ("\n\t" + "")))))
                                                        var v76 int64 = 0
                                                        _ = v75
                                                        _ = v76
                                                    }
                                                }
                                                _ = v70
                                                _ = v71
                                            }
                                            v77 := GenS4{F0: genFn0(203), F1: -37, F2: 0}
                                            _ = v77
                                        } else {
                                            if ((v26 % (((v26 % ((genFn0(-2147483648)) | 1))) | 1)) >= (v26 - 128)) {
                                                fmt.Println(v3)
                                                var v78 int16 = v27
                                                if false {
                                                    fmt.Println(v27, fi56, v27)
                                                }
                                                _ = v78
                                            }
                                            v79 := GenS4{F0: v26, F1: 53, F2: 513}
                                            fmt.Println(fi57, fi56)
                                            _ = v79
                                        }
                                    }
                                    v80 := []int{fi56, 32768, (v2 >> ((4294967295) & 63))}
                                    v81 := GenS5{F0: 2147483648, F1: (("\n\t" + "") + ("hello" + ("a" + "\n\t"))), F2: v27}
                                    _ = v80
                                    _ = v81
                                }
                                v27 = v55
                                _ = fi57
                            }
                            v82 := GenS2{F0: 65535, F1: 0, F2: (v26 - v26)}
                            fmt.Println(v3, v27, v3)
                            _ = fi56
                            _ = v82
                        }
                        _ = v55
                    }
                }
            }
        } else {
            switch v2 {
            case 1:
                v83 := []int{(v2 & v2), 0}
                _ = v83
            case 127:
                v2 = v2
            }
            v3 += 398
        }
    } else {
        for fi84 := 0; fi84 < 6; fi84++ {
            var v85 string = ""
            for fi86 := 0; fi86 < 12; fi86++ {
                var v87 bool = (float32(3.14159) <= 0.0)
                fmt.Println(v85, v87)
                fmt.Println(fi84, fi84, fi86)
                _ = fi86
                _ = v87
            }
            _ = fi84
            _ = v85
        }
        fmt.Println(v2, v26, v3)
        v88 := []int{(v2 ^ (v2 + (v2 << ((0) & 63)))), v2, 128}
        _ = v88
    }
    var v89 int64 = -9223372036854775808
    fmt.Println(v27)
    var v90 uint16 = (v3 >> ((uint(v3)) & 63))
    if false {
        var v91 int = -941
        fmt.Println(v3, v91, v26)
        _ = v91
    } else {
        for fi92 := 0; fi92 < 11; fi92++ {
            fmt.Println(v2)
            _ = fi92
        }
        switch -9223372036854775808 {
        case 32767:
            if false {
                if true {
                    v93 := GenS6{F0: 2, F1: 42}
                    if true {
                        var v94 float64 = 3.14159
                        if (("" + "foo bar") <= ("\n\t" + (("hello" + "\n\t") + ("" + "hello")))) {
                            for fi95 := 0; fi95 < 15; fi95++ {
                                v90 += (v3 + v3)
                                v96 := v1[""]
                                v2 = (fi95 * (^(fi95)))
                                _ = fi95
                                _ = v96
                            }
                        } else {
                            var v97 bool = true
                            v2 = v2
                            fmt.Println(v26, v3)
                            _ = v97
                        }
                        switch "\n\t" {
                        case "":
                            v3 = (v90 | (^(v3)))
                        case "\n\t":
                            var v98 int = v2
                            _ = v98
                        default:
                            v89 = v89
                        }
                        _ = v94
                    }
                    if false {
                        v99 := GenS7{F0: "hello", F1: v89}
                        _ = v99
                    } else {
                        fmt.Println(v3, v90, v26)
                        var v100 int16 = (v27 / ((v27) | 1))
                        v101 := GenS8{F0: 0}
                        _ = v100
                        _ = v101
                    }
                    _ = v93
                } else {
                    v26 = (v26 << ((255) & 63))
                    v26 = 0
                }
                var v102 int8 = int8(v3)
                _ = v102
            }
        case 42:
            var v103 uint16 = 1
            _ = v103
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v89:", v89)
    fmt.Println("v90:", v90)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v38
    _ = v89
    _ = v90
}
