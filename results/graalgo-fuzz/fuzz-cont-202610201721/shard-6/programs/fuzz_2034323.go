package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int }
type GenS2 struct { F0 string; F1 int32; F2 uint32 }
type GenS3 struct { F0 uint; F1 int8 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 int64; F1 uint32; F2 uint16 }
func genFn0(p0 bool, p1 uint64, p2 uint32) bool {
    var v6 int8 = 127
    var v7 int8 = v6
    v8 := []int{4294967295, -2147483648, -271, 0}
    for fi9 := 0; fi9 < 13; fi9++ {
        var v10 uint64 = (p1 >> ((65535) & 63))
        v11 := GenS1{F0: 9223372036854775807, F1: (-(fi9))}
        var v12 int64 = 0
        _ = fi9
        _ = v10
        _ = v11
        _ = v12
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    _ = v8
    return p0
}
func genFn1() int8 {
    var v13 int64 = -116
    v13 = v13
    v14 := []int{-81, 127, 2147483647}
    var v15 uint64 = 0
    v13 += (^(v13))
    _ = v13
    _ = v14
    _ = v15
    return 42
}

func main() {
    var v1 int8 = 19
    var v2 uint32 = 2
    var v3 int16 = 213
    var v4 uint64 = 42
    var v5 uint16 = 625
    for fi16 := 0; fi16 < 1; fi16++ {
        v5 = v5
        if true {
            var v17 uint = 18446744073709551615
            var v18 bool = (!(false || ((v17 >> ((18446744073709551615) & 63)) != 0)))
            _ = v17
            _ = v18
        } else {
            var v19 int16 = int16(v2)
            var v20 int64 = 9223372036854775807
            var v21 int32 = 2147483647
            _ = v19
            _ = v20
            _ = v21
        }
        _ = fi16
    }
    for fi22 := 0; fi22 < 18; fi22++ {
        for fi23 := 0; fi23 < 11; fi23++ {
            switch (fi22 % (((^(fi23))) | 1)) {
            case 4294967295:
                if true {
                    var v24 float64 = float64(0.1)
                    var v25 uint64 = v4
                    var v26 uint64 = uint64(v5)
                    v26 -= v25
                    _ = v24
                    _ = v25
                    _ = v26
                }
            default:
                fmt.Println(fi23, v1)
            }
            v4 = (v4 & (^(v4)))
            var v27 uint64 = uint64(v5)
            switch (^(fi23)) {
            case 0:
                var v28 uint16 = 653
                _ = v28
            case -9223372036854775808:
                if true {
                    fmt.Println(v2)
                    v29 := []int{(fi23 + 2147483647), fi22, fi23}
                    _ = v29
                } else {
                    if (-75 < -32768) {
                        var v30 float32 = 1.5
                        var v31 uint32 = 1
                        v2 += v2
                        _ = v30
                        _ = v31
                    }
                }
            case -2147483648:
                v32 := []int{32767}
                _ = v32
            }
            _ = fi23
            _ = v27
        }
        if false {
            v4 += uint64(v2)
            v2 += 2147483647
        } else {
            fmt.Println(v2, v1)
        }
        for fi33 := 0; fi33 < 11; fi33++ {
            var v34 uint64 = (^(v4))
            v35 := []int{-48}
            _ = fi33
            _ = v34
            _ = v35
        }
        _ = fi22
    }
    v3 += 32767
    v36 := []int{32767}
    for fi37 := 0; fi37 < 6; fi37++ {
        v38 := GenS2{F0: "foo bar", F1: 42, F2: uint32(v4)}
        _ = fi37
        _ = v38
    }
    for fi39 := 0; fi39 < 17; fi39++ {
        if true {
            if (42 == v4) {
                for fi40 := 0; fi40 < 9; fi40++ {
                    if (((((0.1 + (0.0 + float64((-2.25 - 1.5)))) / (0.0 + 0.1)) * float64(0.1)) / 1e10) < -2.25) {
                        var v41 uint = 42
                        _ = v41
                    }
                    fmt.Println(v3, fi39, v5)
                    _ = fi40
                }
            } else {
                v5 += (v5 | (v5 / ((v5) | 1)))
            }
            v4 = (v4 / ((0) | 1))
            switch "" {
            case "foo bar":
                for fi42 := 0; fi42 < 15; fi42++ {
                    v4 -= (v4 - (^(v4)))
                    _ = len(v36)
                    v1 += (v1 >> ((255) & 63))
                    var v43 uint64 = v4
                    _ = fi42
                    _ = v43
                }
            }
        } else {
            switch 9223372036854775807 {
            case 127:
                var v44 int64 = 0
                _ = v44
            default:
                fmt.Println(v3)
            }
        }
        if false {
            if true {
                _ = len(v36)
            }
        } else {
            fmt.Println(v5, v4, v5)
            switch (("hello" + "") + "a") {
            case "\n\t":
                if true {
                    v45 := GenS3{F0: 632, F1: genFn1()}
                    _ = v45
                } else {
                    if len(v36) >= 2 {
                        v36 = v36[1:len(v36)]
                    }
                    v46 := GenS1{F0: (fi39 * fi39), F1: (fi39 & (fi39 >> ((2) & 63)))}
                    if false {
                        if len(v36) > 0 {
                            v36[0] = int(v1)
                        }
                        var v47 uint16 = (v5 & (v5 & v5))
                        v47 -= v5
                        _ = v47
                    } else {
                        var v48 uint = uint(v2)
                        switch "\n\t" {
                        case "a":
                            v4 -= 4294967295
                        case "":
                            v49 := GenS1{F0: int(v2), F1: (fi39 ^ (fi39 - fi39))}
                            _ = v49
                        case "\n\t":
                            for fi50 := 0; fi50 < 14; fi50++ {
                                if true {
                                    v2 = v2
                                } else {
                                    switch "a" {
                                    case "":
                                        var v51 uint64 = v4
                                        _ = v51
                                    case "foo bar":
                                        if len(v36) > 0 {
                                            v36[0] = (-(fi39))
                                        }
                                    default:
                                        if false {
                                            _ = len(v36)
                                        }
                                    }
                                    var v52 float32 = 0.0
                                    _ = v52
                                }
                                var v53 int16 = v3
                                var v54 int = (-(fi39))
                                _ = fi50
                                _ = v53
                                _ = v54
                            }
                        }
                        _ = v48
                    }
                    _ = v46
                }
            case "":
                for fi55 := 0; fi55 < 3; fi55++ {
                    switch (fi39 / ((-1) | 1)) {
                    case 0:
                        fmt.Println(v1, v2)
                    case 4294967295:
                        if false {
                            if len(v36) > 0 {
                                v36[0] = int(v1)
                            }
                            var v56 int32 = -2147483648
                            v3 = -1
                            _ = v56
                        }
                    default:
                        switch ("\n\t" + ((("hello" + "a") + "hello") + "a")) {
                        case "foo bar":
                            v5 = (v5 ^ (^(v5)))
                        }
                    }
                    v3 = 0
                    _ = fi55
                }
            case "hello":
                if ((^(v5)) >= 1) {
                    var v57 uint = 63
                    fmt.Println(v57, v5)
                    v3 = (v3 << ((v57) & 63))
                    _ = v57
                }
            default:
                for fi58 := 0; fi58 < 7; fi58++ {
                    if false {
                        v1 -= genFn1()
                        var v59 int = fi58
                        _ = v59
                    } else {
                        v4 = v4
                    }
                    _ = fi58
                }
            }
        }
        var v60 uint8 = 153
        _ = fi39
        _ = v60
    }
    v36 = append(v36, 778)
    var v61 int = 32768
    var v62 string = "hello"
    var v63 float32 = (0.0 + 1e10)
    if len(v36) > 0 {
        v36[0] = (v61 << ((255) & 63))
    }
    var v64 uint16 = v5
    for fi65 := 0; fi65 < 16; fi65++ {
        for fi66 := 0; fi66 < 16; fi66++ {
            switch fi65 {
            case 128:
                v67 := GenS2{F0: "", F1: 849, F2: uint32(v4)}
                _ = v67
            }
            if true {
                v2 = v2
                if (-9223372036854775808 > 9223372036854775807) {
                    if true {
                        v61 -= v61
                        var v68 uint = 0
                        _ = v68
                    } else {
                        switch v62 {
                        case "a":
                            switch (v62 + ("hello" + v62)) {
                            case "\n\t":
                                var v69 uint16 = v64
                                _ = v69
                            default:
                                if len(v36) >= 2 {
                                    v36 = v36[1:len(v36)]
                                }
                            }
                        case "":
                            v61 -= v61
                        }
                        fmt.Println(v2, fi66, fi66)
                        fmt.Println(v64, fi66, v64)
                        var v70 int8 = (^(v1))
                        _ = v70
                    }
                }
            }
            _ = fi66
        }
        fmt.Println(v2, v4)
        v71 := GenS3{F0: 2, F1: 42}
        _ = fi65
        _ = v71
    }
    v61 -= -621
    if true {
        fmt.Println(v2)
        fmt.Println(v1, v63)
        v63 += v63
    } else {
        for fi72 := 0; fi72 < 10; fi72++ {
            if ((v62 + (v62 + (v62 + ("a" + v62)))) <= ("foo bar" + v62)) {
                switch v62 {
                case "hello":
                    for fi73 := 0; fi73 < 5; fi73++ {
                        v1 = (v1 | (-(v1)))
                        var v74 int64 = -9223372036854775808
                        _ = fi73
                        _ = v74
                    }
                case "foo bar":
                    switch (v62 + v62) {
                    case "a":
                        for fi75 := 0; fi75 < 6; fi75++ {
                            v63 += ((((1.5 + 1e10) / v63) * v63) / 1e10)
                            _ = fi75
                        }
                    case "hello":
                        if len(v36) >= 2 {
                            v36 = v36[1:len(v36)]
                        }
                    case "foo bar":
                        var v76 int16 = v3
                        _ = v76
                    default:
                        v77 := GenS4{F0: uint64(v64)}
                        _ = v77
                    }
                case "a":
                    var v78 int64 = 788
                    _ = v78
                }
                fmt.Println(fi72, v63)
                var v79 uint64 = (v4 * (v4 ^ v4))
                _ = v79
            } else {
                if ((v61 <= int(v64)) || (v62 > "\n\t")) {
                    fmt.Println(v61, v64)
                    v64 += (v64 * v5)
                    if false {
                        for fi80 := 0; fi80 < 1; fi80++ {
                            if false {
                                v2 = v2
                                var v81 uint32 = (^(v2))
                                var v82 int8 = 42
                                if ((202 > 96) || false) {
                                    if len(v36) >= 2 {
                                        v36 = v36[1:len(v36)]
                                    }
                                    v83 := GenS1{F0: v61, F1: 1}
                                    _ = v83
                                } else {
                                    v1 = v82
                                    var v84 int16 = (v3 - v3)
                                    _ = v84
                                }
                                _ = v81
                                _ = v82
                            } else {
                                v2 = (v2 + v2)
                                fmt.Println(v3, v2)
                            }
                            _ = fi80
                        }
                        for fi85 := 0; fi85 < 8; fi85++ {
                            var v86 int8 = 0
                            _ = fi85
                            _ = v86
                        }
                        for fi87 := 0; fi87 < 13; fi87++ {
                            fmt.Println(v2)
                            if len(v36) > 0 {
                                v36[0] = (fi72 % (((v61 | fi72)) | 1))
                            }
                            switch 90 {
                            case 4294967295:
                                v88 := GenS5{F0: -62, F1: (v2 >> ((255) & 63)), F2: (v5 * v64)}
                                _ = v88
                            case 9223372036854775807:
                                v63 += float32(0.0)
                            default:
                                v63 -= 3.14159
                            }
                            _ = fi87
                        }
                    } else {
                        var v89 float64 = 3.14159
                        if true {
                            v64 += uint16(v4)
                            switch v62 {
                            case "":
                                v61 += 9223372036854775807
                            case "foo bar":
                                v3 = -32768
                            }
                        }
                        _ = v89
                    }
                } else {
                    fmt.Println(v63, v3, fi72)
                    var v90 int16 = -32768
                    _ = v90
                }
                fmt.Println(v63, fi72, v61)
                var v91 uint32 = 0
                _ = v91
            }
            var v92 uint16 = v64
            _ = fi72
            _ = v92
        }
        v93 := GenS5{F0: -9223372036854775808, F1: 937, F2: v64}
        _ = v93
    }
    var v94 bool = (true || ((v2 * 2) > v2))
    var v95 uint = 552
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v36:", len(v36), v36)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v61:", v61)
    fmt.Println("v62:", v62)
    fmt.Println("v63:", v63)
    fmt.Println("v64:", v64)
    fmt.Println("v94:", v94)
    fmt.Println("v95:", v95)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v4
    _ = v5
    _ = v61
    _ = v62
    _ = v63
    _ = v64
    _ = v94
    _ = v95
}
