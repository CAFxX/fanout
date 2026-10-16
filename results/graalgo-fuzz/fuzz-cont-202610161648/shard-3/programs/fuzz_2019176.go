package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint; F1 uint8 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 int64; F1 uint8 }
type GenS6 struct { F0 uint32; F1 uint64 }
type GenS7 struct { F0 int32; F1 bool; F2 int8 }
type GenS8 struct { F0 uint8; F1 int8; F2 uint16 }
type GenS9 struct { F0 int32; F1 int16 }
type GenS10 struct { F0 uint }
func genFn0(p0 uint64, p1 int) int8 {
    if true {
        fmt.Println(p0, p1)
    }
    for fi7 := 0; fi7 < 12; fi7++ {
        p0 -= (p0 * p0)
        v8 := GenS1{F0: ""}
        switch 571 {
        case -9223372036854775808:
            if (255 >= 1) {
                p0 = 18446744073709551615
                for fi9 := 0; fi9 < 6; fi9++ {
                    fmt.Println(p0, fi7)
                    p0 -= (^(p0))
                    _ = fi9
                }
                p0 = (^(p0))
            }
        case 0:
            if false {
                var v10 int8 = 11
                _ = v10
            }
        }
        _ = fi7
        _ = v8
    }
    _ = p0
    _ = p1
    return 1
}
func genFn1() int {
    for fi11 := 0; fi11 < 2; fi11++ {
        v12 := GenS1{F0: "a"}
        var v13 int16 = -32768
        _ = fi11
        _ = v12
        _ = v13
    }
    var v14 float64 = 0.1
    v14 = v14
    v15 := GenS1{F0: ((("a" + ("a" + ("hello" + "\n\t"))) + "\n\t") + "\n\t")}
    if (true && true) {
        var v16 bool = true
        var v17 int16 = 42
        if (false && true) {
            switch 463 {
            case 127:
                switch 128 {
                case -1:
                    v17 = 0
                }
            default:
                v18 := GenS2{F0: 808, F1: 234}
                _ = v18
            }
        } else {
            v19 := []int{167, 795, 32767, 42}
            v16 = true
            _ = v19
        }
        _ = v16
        _ = v17
    } else {
        var v20 uint32 = 4294967295
        _ = v20
    }
    _ = v14
    _ = v15
    return 42
}

func main() {
    var v1 uint8 = 0
    v2 := []uint32{0}
    var v3 int16 = 1
    var v4 int64 = 0
    v5 := make(map[int]string)
    var v6 float32 = 0.1
    v1 -= v1
    var v21 uint = 1
    if false {
        if ((v21 % (((v21 << ((333) & 63))) | 1)) > 255) {
            if false {
                if len(v2) > 0 {
                    v2[0] = 4294967295
                }
                var v22 float64 = 1e10
                fmt.Println(v3, v1)
                _ = v22
            } else {
                if false {
                    v3 = v3
                    switch (("foo bar" + ("hello" + (("a" + ("a" + "a")) + "a"))) + ("hello" + "\n\t")) {
                    case "a":
                        switch "hello" {
                        case "\n\t":
                            var v23 int64 = (v4 | (v4 & v4))
                            _ = v23
                        }
                    default:
                        for fi24 := 0; fi24 < 10; fi24++ {
                            if (3.14159 > 0.1) {
                                if (fi24 < (fi24 - fi24)) {
                                    fmt.Println(v3, v4, v1)
                                    v25 := v5[(fi24 | fi24)]
                                    _ = v25
                                }
                                v4 = int64(v1)
                            }
                            for fi26 := 0; fi26 < 2; fi26++ {
                                v27 := GenS1{F0: "\n\t"}
                                v28 := v5[(-(fi26))]
                                v29 := GenS2{F0: (v21 + (^(v21))), F1: (^(v1))}
                                _ = fi26
                                _ = v27
                                _ = v28
                                _ = v29
                            }
                            var v30 int8 = -36
                            _ = fi24
                            _ = v30
                        }
                    }
                    for fi31 := 0; fi31 < 14; fi31++ {
                        for fi32 := 0; fi32 < 10; fi32++ {
                            fmt.Println(v21, v3)
                            _ = fi32
                        }
                        switch ("foo bar" + ((("foo bar" + "") + ((("a" + "hello") + "") + "foo bar")) + ((("foo bar" + "foo bar") + "hello") + ("\n\t" + "hello")))) {
                        case "foo bar":
                            v1 -= v1
                        case "":
                            v21 = 65535
                        case "a":
                            v33 := GenS3{F0: uint8(v21)}
                            _ = v33
                        default:
                            if false {
                                for fi34 := 0; fi34 < 9; fi34++ {
                                    switch "foo bar" {
                                    case "":
                                        v1 = 128
                                    default:
                                        fmt.Println(v4)
                                    }
                                    _ = fi34
                                }
                                v6 = v6
                                var v35 uint32 = 2147483647
                                _ = v35
                            } else {
                                var v36 uint = v21
                                _ = v36
                            }
                        }
                        _ = fi31
                    }
                } else {
                    fmt.Println(v4)
                    fmt.Println(v3, v4, v1)
                    v37 := GenS1{F0: "a"}
                    _ = v37
                }
                switch -682 {
                case -1:
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                case 32768:
                    var v38 uint32 = 17
                    _ = v38
                case -2147483648:
                    v39 := GenS1{F0: "a"}
                    _ = v39
                }
            }
        }
        v4 = 890
    } else {
        for fi40 := 0; fi40 < 17; fi40++ {
            v6 = (0.1 / v6)
            var v41 int64 = v4
            _ = fi40
            _ = v41
        }
        v2 = append(v2, 2147483647)
        fmt.Println(v3, v1, v3)
    }
    if false {
        var v42 int16 = 0
        var v43 int16 = (-(v3))
        v44 := GenS2{F0: (v21 >> ((v21) & 63)), F1: v1}
        _ = v42
        _ = v43
        _ = v44
    }
    fmt.Println(v6, v21, v3)
    switch 2147483647 {
    case -1:
        v4 = 657
    case 5:
        v21 += (v21 | v21)
    case 4294967295:
        delete(v5, -229)
    }
    v6 = ((((((1e10 - 3.14159) - 1e10) * -2.25) / 3.14159) + 0.0) - 1e10)
    var v45 int32 = -1
    for fi46 := 0; fi46 < 1; fi46++ {
        if (!((v4 | v4) != v4)) {
            if true {
                fmt.Println(v21, fi46, v1)
                var v47 uint8 = v1
                v48 := v5[(^(fi46))]
                _ = v47
                _ = v48
            } else {
                var v49 bool = false
                var v50 float64 = (-2.25 * ((1e10 / 1.5) / 3.14159))
                var v51 uint8 = (v1 ^ 1)
                _ = v49
                _ = v50
                _ = v51
            }
            for fi52 := 0; fi52 < 7; fi52++ {
                var v53 uint = (^(v21))
                v4 = -336
                _ = fi52
                _ = v53
            }
            if true {
                var v54 float32 = (float32(-2.25) / 3.14159)
                var v55 int = fi46
                if (-9223372036854775808 > v4) {
                    switch "foo bar" {
                    case "a":
                        switch "" {
                        case "\n\t":
                            _ = len(v2)
                        default:
                            v56 := GenS3{F0: (v1 * v1)}
                            _ = v56
                        }
                    case "hello":
                        v57 := v5[0]
                        _ = v57
                    case "\n\t":
                        var v58 uint8 = 34
                        _ = v58
                    default:
                        if (v21 > (v21 - (v21 >> (((v21 / (((v21 % ((v21) | 1))) | 1))) & 63)))) {
                            v6 -= v6
                            if false {
                                var v59 uint8 = (^(v1))
                                var v60 float64 = 0.1
                                var v61 int32 = (v45 & (-(v45)))
                                _ = v59
                                _ = v60
                                _ = v61
                            } else {
                                v3 = (v3 | v3)
                                v55 -= genFn1()
                                var v62 int8 = 1
                                _ = v62
                            }
                        }
                    }
                }
                _ = v54
                _ = v55
            } else {
                if (100 <= -128) {
                    fmt.Println(v21, v4)
                    v21 -= uint(v1)
                } else {
                    v63 := GenS2{F0: (v21 | 18446744073709551615), F1: (v1 >> ((65535) & 63))}
                    if true {
                        v64 := GenS4{F0: 128}
                        v2 = append(v2, 172)
                        _ = v64
                    } else {
                        v65 := GenS5{F0: (v4 * v4), F1: v1}
                        v6 += ((float32((float64(((((((0.0 * ((v6 - (float32(3.14159) / (v6 / v6))) * v6)) * v6) - (v6 * float32((0.1 + float64((1.5 - ((v6 - 1e10) + float32(0.1)))))))) / v6) - -2.25) - ((v6 + 0.0) / float32(1e10)))) * 0.1)) * (v6 - 0.1)) - (1.5 * (((v6 - 0.0) / v6) / 1.5)))
                        v66 := GenS6{F0: 2147483647, F1: 500}
                        _ = v65
                        _ = v66
                    }
                    _ = v63
                }
                var v67 bool = (!true)
                var v68 uint16 = 493
                _ = v67
                _ = v68
            }
        } else {
            if (!(!(1e10 >= float64((v6 / v6))))) {
                if true {
                    if true {
                        v69 := GenS1{F0: "\n\t"}
                        _ = v69
                    } else {
                        v70 := GenS4{F0: v1}
                        _ = v70
                    }
                } else {
                    v71 := GenS2{F0: uint(v1), F1: (v1 + v1)}
                    for fi72 := 0; fi72 < 10; fi72++ {
                        if len(v2) > 0 {
                            v2[0] = 2147483647
                        }
                        if false {
                            if true {
                                var v73 int32 = (^(v45))
                                var v74 uint = v21
                                v75 := GenS4{F0: v1}
                                _ = v73
                                _ = v74
                                _ = v75
                            } else {
                                if true {
                                    if ("" <= "foo bar") {
                                        var v76 int8 = 0
                                        _ = v76
                                    } else {
                                        var v77 float32 = ((v6 - (((float32(1e10) / ((1.5 - -2.25) + v6)) * 1e10) / 0.1)) / (v6 + 1.5))
                                        if true {
                                            if true {
                                                v5[(fi72 ^ (-(fi46)))] = "\n\t"
                                            }
                                            var v78 float32 = 3.14159
                                            v4 = 9223372036854775807
                                            _ = v78
                                        } else {
                                            var v79 int = (-(fi46))
                                            _ = v79
                                        }
                                        _ = v77
                                    }
                                    if (42 != 1) {
                                        v80 := GenS4{F0: (v1 | (^(v1)))}
                                        var v81 int16 = (v3 >> (((^(v21))) & 63))
                                        v81 = 1
                                        _ = v80
                                        _ = v81
                                    }
                                }
                                v1 = (v1 / (((v1 + v1)) | 1))
                                var v82 int64 = -9223372036854775808
                                if (false && ((-(v45)) < (v45 * v45))) {
                                    var v83 bool = true
                                    _ = v83
                                } else {
                                    if false {
                                        _ = len(v2)
                                    } else {
                                        v82 += v4
                                        fmt.Println(fi46, v4, v82)
                                        v21 += (v21 << ((596) & 63))
                                    }
                                    var v84 float32 = v6
                                    v85 := GenS5{F0: (v82 % ((1) | 1)), F1: v1}
                                    fmt.Println(v84)
                                    v86 := GenS7{F0: (v45 ^ (v45 | v45)), F1: false, F2: -128}
                                    var v87 uint32 = uint32(v1)
                                    if false {
                                        v88 := GenS8{F0: v1, F1: 1, F2: 42}
                                        if (!(!true)) {
                                            _ = len(v2)
                                            v2 = append(v2, v87)
                                            if true {
                                                if true {
                                                    v89 := GenS4{F0: 2}
                                                    _ = v89
                                                }
                                            } else {
                                                var v90 bool = (-99 <= -63)
                                                v91 := GenS3{F0: v1}
                                                _ = v90
                                                _ = v91
                                            }
                                        }
                                        _ = v88
                                    }
                                    _ = v84
                                    _ = v85
                                    _ = v86
                                    _ = v87
                                }
                                _ = v82
                            }
                            if (false && false) {
                                v6 -= v6
                                v92 := GenS5{F0: v4, F1: (^(v1))}
                                _ = v92
                            }
                        } else {
                            _ = len(v5)
                        }
                        switch (fi46 % ((genFn1()) | 1)) {
                        case 128:
                            v6 = v6
                        case 32767:
                            var v93 uint64 = 587
                            _ = v93
                        }
                        _ = fi72
                    }
                    _ = v71
                }
                for fi94 := 0; fi94 < 18; fi94++ {
                    fmt.Println(v4)
                    var v95 int8 = 100
                    _ = fi94
                    _ = v95
                }
                for fi96 := 0; fi96 < 20; fi96++ {
                    _ = len(v5)
                    if (!false) {
                        v3 -= int16(v1)
                        var v97 float32 = 1e10
                        _ = v97
                    } else {
                        var v98 string = ""
                        _ = v98
                    }
                    _ = fi96
                }
            } else {
                var v99 float64 = (0.1 - (1e10 / 3.14159))
                _ = v99
            }
            for fi100 := 0; fi100 < 18; fi100++ {
                _ = len(v5)
                v101 := GenS9{F0: v45, F1: v3}
                v102 := GenS10{F0: (v21 * (v21 >> ((179) & 63)))}
                _ = fi100
                _ = v101
                _ = v102
            }
        }
        _ = fi46
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v45
    _ = v5
    _ = v6
}
