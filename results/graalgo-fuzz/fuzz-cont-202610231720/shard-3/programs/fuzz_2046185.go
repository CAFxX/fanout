package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 string; F1 uint8; F2 bool }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint16; F1 int8 }
type GenS5 struct { F0 uint16; F1 uint16; F2 uint64 }
type GenS6 struct { F0 uint64; F1 int16 }
type GenS7 struct { F0 int64; F1 uint16; F2 uint64 }
func genFn0(p0 string) string {
    fmt.Println(p0)
    var v6 int32 = 42
    _ = p0
    _ = v6
    return (p0 + ((p0 + p0) + p0))
}
func genFn1(p0 uint, p1 int64) int16 {
    fmt.Println(p1, p0)
    for fi7 := 0; fi7 < 2; fi7++ {
        var v8 uint8 = 1
        v9 := []int{32767, 550, fi7, (fi7 % ((-2147483648) | 1))}
        _ = fi7
        _ = v8
        _ = v9
    }
    var v10 uint32 = 42
    var v11 uint32 = 0
    v12 := []int{2147483647, 313, int(p1)}
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    _ = v12
    return -1
}

func main() {
    var v1 float32 = 0.0
    var v2 string = "a"
    var v3 float64 = (((0.0 + ((1e10 / 1.0) - 1e10)) + 1.5) / (1.5 - -2.25))
    var v4 int32 = 0
    v5 := make(map[int]string)
    var v13 uint16 = 270
    for fi14 := 0; fi14 < 16; fi14++ {
        v15 := GenS1{F0: 42}
        _ = fi14
        _ = v15
    }
    switch 4294967295 {
    case 127:
        var v16 int16 = 42
        _ = v16
    case 2147483647:
        for fi17 := 0; fi17 < 19; fi17++ {
            var v18 float64 = v3
            for fi19 := 0; fi19 < 12; fi19++ {
                var v20 int = 656
                v13 -= (v13 ^ (v13 - 214))
                v13 += (v13 >> ((0) & 63))
                _ = fi19
                _ = v20
            }
            _ = fi17
            _ = v18
        }
    case 128:
        switch "\n\t" {
        case "foo bar":
            if true {
                switch (v2 + v2) {
                case "\n\t":
                    var v21 bool = (false && (v13 != (v13 ^ v13)))
                    _ = v21
                case "":
                    fmt.Println(v4)
                case "a":
                    if false {
                        if ((674 < 1000) || true) {
                            if ((751 != -1) || (!(v4 > (v4 | (v4 ^ 1))))) {
                                fmt.Println(v13)
                                switch v2 {
                                case "":
                                    fmt.Println(v1, v3, v4)
                                case "foo bar":
                                    for fi22 := 0; fi22 < 5; fi22++ {
                                        for fi23 := 0; fi23 < 19; fi23++ {
                                            var v24 int64 = 42
                                            v25 := GenS2{F0: "", F1: 65, F2: (!false)}
                                            var v26 int16 = genFn1(65535, (^(v24)))
                                            var v27 float32 = v1
                                            _ = len(v5)
                                            v2 = "foo bar"
                                            v28 := GenS3{F0: v13}
                                            if true {
                                                switch "hello" {
                                                case "a":
                                                    var v29 uint8 = 31
                                                    _ = v29
                                                }
                                                if true {
                                                    if (false && (!false)) {
                                                        v30 := []int{}
                                                        _ = len(v5)
                                                        var v31 float64 = (1e10 + float64(v27))
                                                        _ = v30
                                                        _ = v31
                                                    } else {
                                                        var v32 int64 = v24
                                                        _ = v32
                                                    }
                                                    v2 = ""
                                                    if false {
                                                        v4 -= (v4 & v4)
                                                    }
                                                    if false {
                                                        delete(v5, (fi23 / (((fi23 - (fi23 % (((-(fi23))) | 1)))) | 1)))
                                                        v3 -= (v3 * v3)
                                                    } else {
                                                        v33 := v5[2147483647]
                                                        v24 = (v24 | (v24 ^ v24))
                                                        v24 += 1
                                                        _ = v33
                                                    }
                                                } else {
                                                    switch v2 {
                                                    case "\n\t":
                                                        var v34 uint = 2
                                                        _ = v34
                                                    default:
                                                        v4 -= (v4 + v4)
                                                    }
                                                }
                                            }
                                            v35 := GenS3{F0: 260}
                                            if true {
                                                if ((127 <= 42) && (!(963 >= 2))) {
                                                    v1 += v27
                                                    var v36 int32 = 971
                                                    _ = v36
                                                }
                                            } else {
                                                switch v2 {
                                                case "\n\t":
                                                    v37 := GenS1{F0: int16(v13)}
                                                    _ = v37
                                                default:
                                                    fmt.Println(fi23, v2)
                                                }
                                                _ = len(v5)
                                                v38 := GenS2{F0: v2, F1: 212, F2: (-1 == (^(v24)))}
                                                _ = v38
                                            }
                                            v39 := GenS1{F0: 32767}
                                            _ = fi23
                                            _ = v24
                                            _ = v25
                                            _ = v26
                                            _ = v27
                                            _ = v28
                                            _ = v35
                                            _ = v39
                                        }
                                        for fi40 := 0; fi40 < 5; fi40++ {
                                            v41 := GenS4{F0: 543, F1: -1}
                                            if (!((739 != 18446744073709551615) && ((-2.25 / 1e10) < v3))) {
                                                var v42 int8 = 100
                                                _ = v42
                                            }
                                            if (true && true) {
                                                if true {
                                                    var v43 uint16 = (v13 << ((4294967295) & 63))
                                                    delete(v5, (^(fi40)))
                                                    _ = len(v5)
                                                    _ = v43
                                                }
                                                v3 -= ((v3 / (v3 - v3)) / (1e10 + (0.0 + float64(v1))))
                                            } else {
                                                var v44 bool = (false || true)
                                                _ = v44
                                            }
                                            _ = fi40
                                            _ = v41
                                        }
                                        fmt.Println(v4, v13, fi22)
                                        _ = fi22
                                    }
                                }
                                for fi45 := 0; fi45 < 3; fi45++ {
                                    v3 -= (v3 - v3)
                                    _ = fi45
                                }
                            } else {
                                var v46 int32 = (v4 ^ v4)
                                _ = len(v5)
                                _ = v46
                            }
                            var v47 float64 = v3
                            _ = v47
                        } else {
                            v3 += v3
                            fmt.Println(v13, v1, v1)
                            var v48 uint16 = (v13 - (v13 & v13))
                            _ = v48
                        }
                    }
                default:
                    var v49 uint16 = v13
                    _ = v49
                }
                switch v2 {
                case "a":
                    v50 := v5[128]
                    _ = v50
                default:
                    v4 += v4
                }
            } else {
                v51 := GenS5{F0: v13, F1: 1, F2: 2}
                _ = v51
            }
        default:
            v52 := []int{2147483647, 9223372036854775807}
            _ = v52
        }
    default:
        var v53 int8 = -19
        _ = v53
    }
    var v54 uint = 423
    for fi55 := 0; fi55 < 11; fi55++ {
        v56 := []int{(fi55 >> ((v54) & 63)), 0, 501}
        var v57 int64 = 42
        _ = fi55
        _ = v56
        _ = v57
    }
    v54 = (v54 % ((32) | 1))
    _ = len(v5)
    var v58 int = 4294967295
    _ = len(v5)
    v1 = (v1 / v1)
    if (!false) {
        var v59 int64 = 1
        _ = v59
    } else {
        var v60 uint8 = 127
        _ = v60
    }
    if false {
        for fi61 := 0; fi61 < 13; fi61++ {
            var v62 uint32 = 2147483648
            fmt.Println(v58)
            _ = fi61
            _ = v62
        }
    }
    if false {
        v63 := GenS5{F0: 255, F1: (v13 % (((v13 - 32767)) | 1)), F2: 4294967295}
        _ = v63
    } else {
        switch v58 {
        case -699:
            var v64 uint32 = 0
            _ = v64
        case -9223372036854775808:
            for fi65 := 0; fi65 < 18; fi65++ {
                if false {
                    var v66 uint64 = 18446744073709551615
                    fmt.Println(v1, v3)
                    _ = v66
                }
                switch v2 {
                case "":
                    v67 := []int{}
                    _ = v67
                case "\n\t":
                    fmt.Println(v54, fi65)
                case "foo bar":
                    switch fi65 {
                    case 981:
                        var v68 float32 = v1
                        _ = v68
                    case 127:
                        fmt.Println(v1, v1, v4)
                    case 42:
                        v13 = v13
                    }
                }
                _ = fi65
            }
        case -1:
            v69 := GenS4{F0: 65535, F1: 127}
            _ = v69
        default:
            switch v2 {
            case "foo bar":
                delete(v5, v58)
            case "hello":
                var v70 int8 = 1
                _ = v70
            case "a":
                v13 = v13
            default:
                fmt.Println(v58)
            }
        }
        if (0.1 >= v3) {
            switch v2 {
            case "foo bar":
                if (!false) {
                    for fi71 := 0; fi71 < 8; fi71++ {
                        for fi72 := 0; fi72 < 8; fi72++ {
                            v73 := []int{(-(fi72)), (fi71 >> ((v54) & 63)), 941, 0}
                            _ = fi72
                            _ = v73
                        }
                        _ = fi71
                    }
                }
            case "\n\t":
                delete(v5, v58)
            default:
                switch ("hello" + "") {
                case "a":
                    switch "hello" {
                    case "a":
                        if ((true || false) || true) {
                            v74 := GenS2{F0: v2, F1: 42, F2: true}
                            _ = v74
                        }
                    default:
                        v2 = ("hello" + "")
                    }
                case "":
                    v75 := v5[-356]
                    _ = v75
                }
            }
        } else {
            for fi76 := 0; fi76 < 12; fi76++ {
                v13 -= (v13 >> ((v54) & 63))
                _ = fi76
            }
            switch (v58 / (((v58 | 4294967295)) | 1)) {
            case -2147483648:
                switch -2147483648 {
                case 686:
                    switch "a" {
                    case "\n\t":
                        var v77 float64 = v3
                        _ = v77
                    case "a":
                        v78 := []int{v58, v58}
                        _ = v78
                    default:
                        for fi79 := 0; fi79 < 12; fi79++ {
                            for fi80 := 0; fi80 < 2; fi80++ {
                                var v81 uint32 = 1
                                _ = fi80
                                _ = v81
                            }
                            v58 = (v58 | 127)
                            _ = fi79
                        }
                    }
                case -690:
                    if true {
                        v82 := GenS6{F0: 42, F1: 75}
                        _ = v82
                    }
                case 42:
                    if (377 == v58) {
                        fmt.Println(v2, v54, v3)
                        delete(v5, v58)
                        v83 := []int{v58, 507, (v58 + v58)}
                        _ = v83
                    }
                }
            case -1:
                if ((427 <= 0) && false) {
                    switch v2 {
                    case "\n\t":
                        switch "hello" {
                        case "a":
                            if false {
                                fmt.Println(v54)
                                var v84 int64 = -9223372036854775808
                                v85 := GenS7{F0: (v84 | -1), F1: v13, F2: 4294967295}
                                _ = v84
                                _ = v85
                            } else {
                                v86 := []int{}
                                _ = v86
                            }
                        case "":
                            var v87 uint8 = 0
                            _ = v87
                        case "foo bar":
                            switch 9223372036854775807 {
                            case 1:
                                var v88 int16 = 1000
                                _ = v88
                            case -2147483648:
                                var v89 uint64 = 2
                                _ = v89
                            case 2147483647:
                                v90 := []int{4294967295, -9223372036854775808}
                                _ = v90
                            default:
                                var v91 uint32 = 2
                                _ = v91
                            }
                        }
                    }
                } else {
                    fmt.Println(v3, v2, v1)
                    v5[v58] = v2
                }
            default:
                if true {
                    v4 += 2147483647
                }
            }
            v58 = v58
        }
        v58 = (v58 * (-(v58)))
    }
    v3 = v3
    fmt.Println(v3)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v54:", v54)
    fmt.Println("v58:", v58)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v54
    _ = v58
}
