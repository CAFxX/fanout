package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint; F2 int }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int; F1 uint64 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 uint64; F1 int8 }
type GenS6 struct { F0 uint64; F1 int8; F2 uint }
func genFn0(p0 int16, p1 string, p2 int8) string {
    p2 += (p2 - 1)
    v7 := []int{32767, 494}
    p1 = ((p1 + "foo bar") + "\n\t")
    if len(v7) > 0 {
        v7[0] = 377
    }
    switch p1 {
    case "":
        if ((p2 & -1) == p2) {
            if true {
                for fi8 := 0; fi8 < 10; fi8++ {
                    for fi9 := 0; fi9 < 10; fi9++ {
                        p2 = (p2 | (p2 >> ((65535) & 63)))
                        p2 += (-(p2))
                        _ = fi9
                    }
                    _ = fi8
                }
                v10 := GenS1{F0: 42, F1: 255, F2: 2147483647}
                v11 := GenS1{F0: 255, F1: 18446744073709551615, F2: 32767}
                _ = v10
                _ = v11
            } else {
                v12 := GenS1{F0: 127, F1: 0, F2: 32768}
                _ = v12
            }
        }
    case "a":
        if true {
            fmt.Println(p2, p1, p1)
        } else {
            if true {
                var v13 uint = 18446744073709551615
                var v14 uint = v13
                p1 = "foo bar"
                _ = v13
                _ = v14
            }
            fmt.Println(p0)
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return ("" + "a")
}

func main() {
    v1 := make(map[string]int)
    var v2 float32 = 0.1
    var v3 int64 = 42
    var v4 int64 = 9223372036854775807
    v5 := make(map[string]int)
    var v6 uint64 = 2
    v15 := GenS1{F0: 255, F1: 4294967295, F2: 42}
    var v16 uint64 = (^(v6))
    if false {
        v2 -= 1e10
    }
    var v17 int = 4294967295
    switch (v17 / ((v17) | 1)) {
    case 0:
        if false {
            var v18 uint64 = v6
            v3 -= v3
            var v19 int = (v17 / ((v17) | 1))
            _ = v18
            _ = v19
        }
    case 444:
        for fi20 := 0; fi20 < 11; fi20++ {
            v4 = int64(v17)
            _ = fi20
        }
    default:
        if ((!((float32(1e10) <= v2) && (-32768 > -1))) || false) {
            v16 -= v16
        }
    }
    var v21 uint8 = 0
    fmt.Println(v4, v17)
    var v22 float32 = v2
    if (v16 > (v16 % ((v6) | 1))) {
        var v23 int16 = -32768
        var v24 uint64 = v16
        if true {
            if true {
                switch "" {
                case "foo bar":
                    fmt.Println(v21, v6)
                case "hello":
                    v25 := GenS1{F0: (v21 | v21), F1: 255, F2: v17}
                    _ = v25
                }
            } else {
                if true {
                    if true {
                        v21 += v21
                        switch "a" {
                        case "":
                            var v26 int8 = 127
                            _ = v26
                        case "hello":
                            fmt.Println(v16, v6, v23)
                        case "a":
                            var v27 float64 = 1e10
                            _ = v27
                        }
                        v28 := GenS1{F0: (v21 * 2), F1: 255, F2: v17}
                        _ = v28
                    }
                    for fi29 := 0; fi29 < 2; fi29++ {
                        var v30 string = "hello"
                        _ = fi29
                        _ = v30
                    }
                    v31 := GenS2{F0: v4}
                    _ = v31
                }
                v32 := []int{4294967295, (v17 ^ 991)}
                var v33 int8 = -128
                _ = v32
                _ = v33
            }
            if (false || (v17 == 32768)) {
                switch (v17 ^ v17) {
                case 9223372036854775807:
                    _ = len(v5)
                case 32767:
                    var v34 uint32 = 1
                    _ = v34
                case 153:
                    switch "\n\t" {
                    case "\n\t":
                        if false {
                            v35 := GenS3{F0: int(v3), F1: uint64(v21)}
                            _ = v35
                        }
                    case "":
                        if (!(2 < 249)) {
                            v36 := GenS4{F0: 4294967295}
                            v6 = (^(v24))
                            _ = v36
                        }
                    }
                }
                var v37 int8 = 127
                v38 := GenS4{F0: v6}
                _ = v37
                _ = v38
            }
        }
        _ = v23
        _ = v24
    } else {
        if true {
            fmt.Println(v4, v17, v22)
            var v39 int8 = 1
            for fi40 := 0; fi40 < 12; fi40++ {
                for fi41 := 0; fi41 < 5; fi41++ {
                    v5["\n\t"] = v17
                    _ = fi41
                }
                delete(v5, ("hello" + "\n\t"))
                _ = fi40
            }
            _ = v39
        }
        if false {
            if true {
                fmt.Println(v4, v16, v6)
            } else {
                var v42 uint64 = (v16 * v16)
                _ = v42
            }
            if (2147483647 >= 0) {
                switch "\n\t" {
                case "hello":
                    _ = len(v1)
                }
            } else {
                for fi43 := 0; fi43 < 6; fi43++ {
                    v6 = (v16 % ((v16) | 1))
                    if true {
                        v3 = v4
                        var v44 int = v17
                        _ = v44
                    }
                    if true {
                        v3 -= v3
                    }
                    _ = fi43
                }
                v45 := GenS4{F0: v16}
                if false {
                    if false {
                        for fi46 := 0; fi46 < 8; fi46++ {
                            v16 += uint64(v21)
                            if false {
                                fmt.Println(v21, v4, v17)
                            } else {
                                for fi47 := 0; fi47 < 7; fi47++ {
                                    var v48 int8 = 42
                                    v49 := GenS4{F0: v16}
                                    _ = fi47
                                    _ = v48
                                    _ = v49
                                }
                            }
                            switch ("hello" + "a") {
                            case "foo bar":
                                switch ("hello" + "") {
                                case "":
                                    var v50 uint16 = 32767
                                    _ = v50
                                case "foo bar":
                                    switch "a" {
                                    case "hello":
                                        var v51 int8 = 99
                                        _ = v51
                                    }
                                case "hello":
                                    var v52 uint64 = (^(v16))
                                    _ = v52
                                }
                            case "hello":
                                var v53 int = fi46
                                _ = v53
                            case "\n\t":
                                if true {
                                    for fi54 := 0; fi54 < 11; fi54++ {
                                        if (fi46 <= fi46) {
                                            if (42 < (^(v6))) {
                                                v55 := GenS3{F0: (fi54 / ((v17) | 1)), F1: (v16 % ((v6) | 1))}
                                                _ = len(v5)
                                                fmt.Println(v4, v4)
                                                if ("a" >= ("\n\t" + "foo bar")) {
                                                    switch v17 {
                                                    case -759:
                                                        v56 := []int{601, fi46, fi46}
                                                        _ = v56
                                                    case 127:
                                                        if true {
                                                            fmt.Println(v21, fi46)
                                                            var v57 int32 = 42
                                                            v58 := GenS3{F0: v17, F1: 18446744073709551615}
                                                            _ = v57
                                                            _ = v58
                                                        }
                                                    case 9223372036854775807:
                                                        fmt.Println(v21, v2, v21)
                                                    }
                                                    var v59 bool = (0 <= 793)
                                                    _ = v59
                                                }
                                                v60 := []int{}
                                                _ = v55
                                                _ = v60
                                            }
                                            v17 -= v17
                                            var v61 int = v17
                                            v3 = v4
                                            v62 := GenS5{F0: v6, F1: int8(v21)}
                                            _ = v61
                                            _ = v62
                                        } else {
                                            if (67 <= 127) {
                                                var v63 uint = 476
                                                _ = v63
                                            } else {
                                                fmt.Println(v3, v21, fi54)
                                                var v64 uint16 = 1
                                                _ = v64
                                            }
                                            v16 = v16
                                        }
                                        _ = fi54
                                    }
                                    for fi65 := 0; fi65 < 1; fi65++ {
                                        v66 := GenS3{F0: 1, F1: (^(v16))}
                                        _ = fi65
                                        _ = v66
                                    }
                                }
                            default:
                                v6 = (^(v16))
                            }
                            _ = fi46
                        }
                    }
                    switch ("" + "\n\t") {
                    case "\n\t":
                        var v67 float64 = ((1e10 + float64(v2)) + (3.14159 * ((0.0 + 0.0) + 0.0)))
                        _ = v67
                    case "":
                        switch v17 {
                        case 967:
                            if (v17 != -1) {
                                for fi68 := 0; fi68 < 12; fi68++ {
                                    v69 := GenS6{F0: (v16 % ((531) | 1)), F1: -128, F2: 103}
                                    fmt.Println(fi68)
                                    for fi70 := 0; fi70 < 14; fi70++ {
                                        v4 = -987
                                        _ = fi70
                                    }
                                    _ = fi68
                                    _ = v69
                                }
                            } else {
                                v2 += -2.25
                                v71 := GenS5{F0: 442, F1: 0}
                                _ = v71
                            }
                        }
                    default:
                        v4 = -784
                    }
                    fmt.Println(v17, v17)
                }
                _ = v45
            }
            for fi72 := 0; fi72 < 1; fi72++ {
                if (!true) {
                    if (!(!(3.14159 >= 1.5))) {
                        var v73 uint16 = 255
                        _ = v73
                    } else {
                        fmt.Println(fi72, v6, v4)
                        var v74 uint8 = v21
                        _ = v74
                    }
                    for fi75 := 0; fi75 < 9; fi75++ {
                        switch "foo bar" {
                        case "foo bar":
                            var v76 uint32 = 2
                            _ = v76
                        default:
                            v77 := []int{(fi72 ^ fi75)}
                            _ = v77
                        }
                        if (false && false) {
                            fmt.Println(v22, v22)
                        } else {
                            var v78 int32 = 2147483647
                            fmt.Println(v78)
                            if false {
                                v16 += v6
                                v78 -= 415
                            } else {
                                if (422 <= 152) {
                                    var v79 int8 = -1
                                    var v80 float32 = -2.25
                                    _ = v79
                                    _ = v80
                                } else {
                                    v2 -= 1e10
                                }
                                if true {
                                    var v81 uint64 = (v6 % ((v16) | 1))
                                    var v82 float32 = ((((float32((-2.25 + float64(v2))) + v2) + v22) / v2) / 0.1)
                                    switch "a" {
                                    case "a":
                                        v83 := []int{-597, fi72, v17, fi72}
                                        _ = v83
                                    case "\n\t":
                                        fmt.Println(v16)
                                    }
                                    _ = v81
                                    _ = v82
                                } else {
                                    delete(v1, "")
                                }
                                var v84 uint8 = v21
                                _ = v84
                            }
                            _ = v78
                        }
                        var v85 uint32 = 118
                        _ = fi75
                        _ = v85
                    }
                    var v86 bool = ("foo bar" < "a")
                    _ = v86
                } else {
                    switch "foo bar" {
                    case "\n\t":
                        _ = len(v1)
                    case "hello":
                        v4 += (v3 ^ (v4 << ((4294967295) & 63)))
                    }
                    v22 = 0.1
                }
                _ = fi72
            }
        } else {
            v87 := GenS4{F0: (v16 / ((v16) | 1))}
            var v88 int32 = -2147483648
            _ = v87
            _ = v88
        }
    }
    var v89 int32 = 42
    v90 := GenS6{F0: 42, F1: 127, F2: uint(v6)}
    var v91 int32 = v89
    v91 = v91
    if ((((v3 >> ((18446744073709551615) & 63)) <= (^(v3))) && false) && true) {
        v22 += 0.0
        var v92 float64 = 1.5
        _ = v92
    } else {
        if (false && (!true)) {
            delete(v5, "foo bar")
        }
    }
    fmt.Println(v22, v3, v16)
    v3 = int64(v16)
    switch v17 {
    case 32768:
        fmt.Println(v91, v6)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v89:", v89)
    fmt.Println("v91:", v91)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v89
    _ = v90
    _ = v91
}
