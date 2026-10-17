package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 string; F2 int16 }
type GenS2 struct { F0 int16; F1 uint }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 int64; F1 string; F2 int16 }
type GenS6 struct { F0 string; F1 int64 }
type GenS7 struct { F0 string; F1 int16; F2 uint16 }
func genFn0(p0 int8, p1 uint) uint64 {
    for fi5 := 0; fi5 < 2; fi5++ {
        switch fi5 {
        case 127:
            var v6 uint16 = 255
            _ = v6
        case 271:
            var v7 int32 = 1
            _ = v7
        default:
            var v8 string = ("foo bar" + "\n\t")
            _ = v8
        }
        _ = fi5
    }
    if false {
        p1 = 0
    }
    _ = p0
    _ = p1
    return 42
}
func genFn1() bool {
    var v9 int8 = -43
    if true {
        for fi10 := 0; fi10 < 17; fi10++ {
            var v11 int16 = 42
            _ = fi10
            _ = v11
        }
        if (!false) {
            for fi12 := 0; fi12 < 13; fi12++ {
                var v13 uint64 = 4294967295
                var v14 string = "foo bar"
                _ = fi12
                _ = v13
                _ = v14
            }
            if false {
                fmt.Println(v9)
                v15 := []int{-1, -2147483648, -979}
                _ = v15
            } else {
                var v16 int64 = 42
                _ = v16
            }
        } else {
            switch "foo bar" {
            case "hello":
                fmt.Println(v9)
            case "foo bar":
                if true {
                    if true {
                        switch "\n\t" {
                        case "hello":
                            switch "hello" {
                            case "foo bar":
                                fmt.Println(v9)
                            case "hello":
                                var v17 int32 = -2147483648
                                _ = v17
                            }
                        case "foo bar":
                            for fi18 := 0; fi18 < 18; fi18++ {
                                var v19 string = ("" + "")
                                v9 += (v9 - 100)
                                _ = fi18
                                _ = v19
                            }
                        case "\n\t":
                            for fi20 := 0; fi20 < 4; fi20++ {
                                fmt.Println(v9, v9)
                                _ = fi20
                            }
                        }
                        v9 += (v9 / ((v9) | 1))
                        v21 := GenS1{F0: 0, F1: "a", F2: 1000}
                        _ = v21
                    }
                    if true {
                        switch 610 {
                        case 2147483647:
                            v22 := []int{59, -2147483648, 128, 2147483647}
                            _ = v22
                        case -343:
                            for fi23 := 0; fi23 < 17; fi23++ {
                                fmt.Println(v9)
                                _ = fi23
                            }
                        }
                        v24 := []int{-9223372036854775808, 4294967295}
                        _ = v24
                    }
                }
            case "\n\t":
                if false {
                    switch "hello" {
                    case "\n\t":
                        fmt.Println(v9)
                    case "foo bar":
                        var v25 int64 = -1
                        _ = v25
                    case "":
                        for fi26 := 0; fi26 < 2; fi26++ {
                            v9 -= -105
                            v9 = v9
                            switch ("" + "foo bar") {
                            case "foo bar":
                                v27 := GenS2{F0: 0, F1: 386}
                                _ = v27
                            case "":
                                if (!true) {
                                    v28 := GenS2{F0: 53, F1: 1}
                                    switch "a" {
                                    case "\n\t":
                                        var v29 int8 = -1
                                        _ = v29
                                    case "a":
                                        v30 := GenS2{F0: 0, F1: 452}
                                        _ = v30
                                    case "":
                                        var v31 string = "\n\t"
                                        _ = v31
                                    }
                                    _ = v28
                                } else {
                                    v9 += v9
                                    v32 := []int{(fi26 | fi26), (fi26 ^ (fi26 / (((fi26 | -470)) | 1))), -1}
                                    _ = v32
                                }
                            case "a":
                                v9 = (v9 + -1)
                            }
                            _ = fi26
                        }
                    default:
                        v9 = 127
                    }
                    if (!false) {
                        for fi33 := 0; fi33 < 4; fi33++ {
                            for fi34 := 0; fi34 < 12; fi34++ {
                                if ((0.1 / (1.5 / 3.14159)) > 1e10) {
                                    v9 = 1
                                } else {
                                    v35 := []int{(^(fi34))}
                                    _ = v35
                                }
                                var v36 uint16 = 67
                                if (false || true) {
                                    fmt.Println(fi34, fi34)
                                    if (-1 <= -477) {
                                        v37 := []int{}
                                        v9 = (v9 & v9)
                                        _ = v37
                                    }
                                    v38 := []int{(fi33 >> ((65535) & 63)), 0}
                                    v9 = v9
                                    _ = v38
                                } else {
                                    if true {
                                        var v39 int16 = 42
                                        if (!false) {
                                            v40 := GenS3{F0: 1}
                                            _ = v40
                                        }
                                        _ = v39
                                    }
                                    if (-849 <= 1) {
                                        var v41 uint32 = 42
                                        if (!false) {
                                            var v42 float64 = 0.0
                                            switch "foo bar" {
                                            case "foo bar":
                                                v36 = (v36 | (v36 | 255))
                                            }
                                            _ = v42
                                        }
                                        _ = v41
                                    }
                                }
                                _ = fi34
                                _ = v36
                            }
                            _ = fi33
                        }
                    } else {
                        fmt.Println(v9)
                        v9 = 118
                        fmt.Println(v9)
                    }
                    fmt.Println(v9)
                }
            default:
                v43 := []int{}
                _ = v43
            }
            v9 -= 86
        }
    }
    v44 := []int{-2147483648, 391, 127}
    _ = v44
    _ = v9
    return true
}

func main() {
    v1 := make(map[string]string)
    var v2 string = ("hello" + "")
    var v3 float32 = 0.1
    var v4 int32 = 2147483647
    v45 := []int{0, -2147483648}
    switch (v2 + v2) {
    case "hello":
        for fi46 := 0; fi46 < 17; fi46++ {
            switch v2 {
            case "a":
                v47 := GenS4{F0: 255}
                _ = v47
            case "foo bar":
                if false {
                    var v48 float32 = v3
                    delete(v1, v2)
                    _ = len(v1)
                    _ = v48
                } else {
                    if true {
                        var v49 int64 = 42
                        _ = v49
                    } else {
                        var v50 int64 = 42
                        var v51 int16 = -434
                        _ = v50
                        _ = v51
                    }
                }
            case "":
                if (false || false) {
                    fmt.Println(v3, v3)
                    for fi52 := 0; fi52 < 10; fi52++ {
                        v3 -= v3
                        _ = fi52
                    }
                    for fi53 := 0; fi53 < 8; fi53++ {
                        v54 := GenS5{F0: 0, F1: v2, F2: 534}
                        fmt.Println(v2, v3)
                        var v55 int = 2147483647
                        _ = fi53
                        _ = v54
                        _ = v55
                    }
                } else {
                    if false {
                        _ = len(v45)
                        v4 = (v4 << ((4294967295) & 63))
                    } else {
                        v56 := v1[(((v2 + "") + ("a" + (v2 + "\n\t"))) + "\n\t")]
                        v57 := GenS1{F0: 42, F1: v2, F2: -1}
                        v56 = v56
                        _ = v56
                        _ = v57
                    }
                    var v58 int16 = 1000
                    _ = v58
                }
            default:
                switch (fi46 + (fi46 << ((18446744073709551615) & 63))) {
                case 32768:
                    var v59 uint32 = 115
                    _ = v59
                case 0:
                    var v60 int = (^(fi46))
                    _ = v60
                }
            }
            v61 := GenS4{F0: 299}
            _ = fi46
            _ = v61
        }
    case "a":
        for fi62 := 0; fi62 < 17; fi62++ {
            v4 = 0
            _ = fi62
        }
    default:
        switch "a" {
        case "":
            _ = len(v45)
        case "\n\t":
            var v63 uint = 587
            _ = v63
        }
    }
    switch -9223372036854775808 {
    case -272:
        var v64 uint8 = 227
        _ = v64
    case 9223372036854775807:
        for fi65 := 0; fi65 < 9; fi65++ {
            var v66 bool = (true && (3.14159 > 3.14159))
            if v66 {
                for fi67 := 0; fi67 < 15; fi67++ {
                    v66 = true
                    _ = len(v45)
                    switch v2 {
                    case "foo bar":
                        var v68 uint16 = 888
                        _ = v68
                    default:
                        var v69 uint = 1
                        _ = v69
                    }
                    if true {
                        fmt.Println(v2, v4, v4)
                        v66 = (((!v66) || ((^(v4)) >= v4)) && ((!true) && v66))
                        var v70 bool = (((v4 <= (v4 / ((-2147483648) | 1))) && (false && (127 <= (-(fi65))))) || (2 > 1))
                        _ = v70
                    }
                    _ = len(v45)
                    _ = fi67
                }
                for fi71 := 0; fi71 < 18; fi71++ {
                    v72 := GenS6{F0: v2, F1: -433}
                    _ = len(v45)
                    _ = fi71
                    _ = v72
                }
            } else {
                var v73 int = 0
                _ = v73
            }
            v4 = (v4 - 42)
            _ = fi65
            _ = v66
        }
    }
    for fi74 := 0; fi74 < 8; fi74++ {
        var v75 uint16 = 2
        var v76 uint16 = v75
        _ = fi74
        _ = v75
        _ = v76
    }
    v1[v2] = v2
    fmt.Println(v3, v2)
    switch v2 {
    case "a":
        v3 += ((v3 / 1.5) * (3.14159 + v3))
    }
    if true {
        if (false && ((!(genFn0(42, 4294967295) <= 42)) || true)) {
            fmt.Println(v3, v2, v2)
            var v77 int32 = (v4 | v4)
            _ = v77
        }
        v2 = ""
    }
    var v78 float32 = (v3 * v3)
    var v79 int32 = (v4 + (^(v4)))
    switch v2 {
    case "\n\t":
        v3 = 1e10
    default:
        var v80 float64 = (1e10 * 3.14159)
        _ = v80
    }
    fmt.Println(v3, v78, v2)
    var v81 bool = ((0.0 / -2.25) != (((1e10 + float64(v78)) * -2.25) - (1e10 + ((3.14159 * 1.5) + (0.0 - 1e10)))))
    v82 := GenS7{F0: (v2 + v2), F1: 1000, F2: 537}
    fmt.Println(v79, v4, v78)
    v4 -= (v4 * v4)
    v2 = (("foo bar" + v2) + v2)
    if (0.1 <= (1.5 * ((1e10 * 1.5) / 1e10))) {
        var v83 int16 = -1
        _ = v83
    } else {
        _ = len(v1)
        v84 := GenS5{F0: 622, F1: ((v2 + "") + (v2 + v2)), F2: 1000}
        _ = v84
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v45:", len(v45), v45)
    fmt.Println("v78:", v78)
    fmt.Println("v79:", v79)
    fmt.Println("v81:", v81)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v78
    _ = v79
    _ = v81
    _ = v82
}
