package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint8; F1 uint64 }
type GenS3 struct { F0 int32; F1 uint8; F2 uint }
type GenS4 struct { F0 uint8; F1 uint }
type GenS5 struct { F0 uint8; F1 int8 }
type GenS6 struct { F0 uint16; F1 uint16 }
type GenS7 struct { F0 uint32; F1 int8; F2 uint32 }
type GenS8 struct { F0 int32; F1 uint; F2 int8 }
type GenS9 struct { F0 int32 }
type GenS10 struct { F0 uint; F1 int16; F2 uint }
type GenS11 struct { F0 string; F1 int32; F2 int32 }
type GenS12 struct { F0 int32; F1 int16; F2 uint32 }
func genFn0() uint32 {
    var v4 uint64 = 42
    v4 = (v4 | v4)
    var v5 uint8 = 42
    v6 := GenS1{F0: -661}
    v7 := GenS2{F0: v5, F1: uint64(v5)}
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    return 65535
}
func genFn1(p0 uint8, p1 int32, p2 int64) int32 {
    fmt.Println(p1)
    switch 0 {
    case 32767:
        for fi8 := 0; fi8 < 2; fi8++ {
            var v9 uint8 = p0
            if (-553 != 1000) {
                var v10 int32 = p1
                p1 = (v10 >> ((887) & 63))
                _ = v10
            } else {
                var v11 int = fi8
                if false {
                    var v12 uint64 = 0
                    _ = v12
                } else {
                    var v13 uint16 = 2
                    _ = v13
                }
                _ = v11
            }
            if (-1 > 128) {
                var v14 bool = (genFn0() < 2147483648)
                switch fi8 {
                case -1:
                    v15 := GenS3{F0: 1, F1: 128, F2: uint(v9)}
                    _ = v15
                case 42:
                    v16 := GenS4{F0: 42, F1: 42}
                    _ = v16
                case 128:
                    v9 -= v9
                default:
                    v17 := GenS2{F0: (v9 / ((p0) | 1)), F1: 0}
                    _ = v17
                }
                _ = v14
            }
            _ = fi8
            _ = v9
        }
    }
    v18 := GenS5{F0: (p0 % ((p0) | 1)), F1: 6}
    p0 += (p0 ^ p0)
    switch int(p2) {
    case 32768:
        var v19 bool = (0 <= 289)
        _ = v19
    default:
        var v20 int32 = -692
        _ = v20
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v18
    return p1
}

func main() {
    v1 := make(map[uint32]string)
    var v2 int16 = -464
    var v3 bool = true
    v21 := []int{-9223372036854775808, 32768}
    var v22 string = "\n\t"
    switch 128 {
    case -1:
        for fi23 := 0; fi23 < 18; fi23++ {
            if len(v21) >= 2 {
                v21 = v21[1:len(v21)]
            }
            for fi24 := 0; fi24 < 18; fi24++ {
                v25 := GenS6{F0: 1, F1: 32767}
                v22 = "\n\t"
                delete(v1, genFn0())
                _ = fi24
                _ = v25
            }
            v26 := v1[2147483648]
            _ = fi23
            _ = v26
        }
    default:
        delete(v1, 4294967295)
    }
    v22 = v22
    if v3 {
        v27 := GenS4{F0: 127, F1: 4294967295}
        _ = len(v21)
        v3 = false
        _ = v27
    } else {
        v3 = v3
        for fi28 := 0; fi28 < 17; fi28++ {
            for fi29 := 0; fi29 < 12; fi29++ {
                v21 = append(v21, -2147483648)
                _ = fi29
            }
            _ = fi28
        }
    }
    v30 := GenS6{F0: 255, F1: 1}
    if v3 {
        fmt.Println(v22, v22)
        for fi31 := 0; fi31 < 9; fi31++ {
            v32 := GenS5{F0: 0, F1: 127}
            for fi33 := 0; fi33 < 6; fi33++ {
                v34 := GenS7{F0: 472, F1: 1, F2: 147}
                _ = fi33
                _ = v34
            }
            _ = fi31
            _ = v32
        }
    } else {
        if v3 {
            var v35 uint64 = 947
            var v36 uint32 = 32
            _ = v35
            _ = v36
        } else {
            if (!(42 >= 950)) {
                var v37 uint32 = 766
                if false {
                    for fi38 := 0; fi38 < 3; fi38++ {
                        delete(v1, (v37 / (((v37 & (^(v37)))) | 1)))
                        v2 += -867
                        for fi39 := 0; fi39 < 20; fi39++ {
                            fmt.Println(v37)
                            var v40 int64 = -1
                            switch v22 {
                            case "hello":
                                var v41 float32 = (-2.25 - 0.1)
                                _ = v41
                            case "":
                                v40 += v40
                            }
                            _ = fi39
                            _ = v40
                        }
                        _ = fi38
                    }
                    for fi42 := 0; fi42 < 11; fi42++ {
                        for fi43 := 0; fi43 < 7; fi43++ {
                            v44 := GenS6{F0: 255, F1: 42}
                            _ = len(v1)
                            var v45 uint64 = 1
                            _ = fi43
                            _ = v44
                            _ = v45
                        }
                        switch (fi42 / ((128) | 1)) {
                        case -1:
                            switch fi42 {
                            case -807:
                                if true {
                                    var v46 int = fi42
                                    if v3 {
                                        fmt.Println(v46)
                                        v46 += v46
                                    }
                                    v21 = append(v21, fi42)
                                    _ = v46
                                } else {
                                    v2 -= v2
                                }
                            default:
                                if (true || ((fi42 & fi42) != (fi42 + fi42))) {
                                    v47 := GenS8{F0: 42, F1: 211, F2: 0}
                                    _ = v47
                                } else {
                                    v48 := v1[v37]
                                    v49 := GenS9{F0: 1}
                                    _ = len(v1)
                                    _ = v48
                                    _ = v49
                                }
                            }
                        case -48:
                            if len(v21) > 0 {
                                v21[0] = (fi42 / ((fi42) | 1))
                            }
                        }
                        for fi50 := 0; fi50 < 11; fi50++ {
                            if (2 <= 127) {
                                fmt.Println(v22)
                                v51 := GenS2{F0: 127, F1: 0}
                                v22 = (v22 + "\n\t")
                                _ = v51
                            } else {
                                v3 = (v3 || (false && true))
                                switch v22 {
                                case "foo bar":
                                    fmt.Println(v2)
                                case "hello":
                                    var v52 int8 = -106
                                    _ = v52
                                default:
                                    v37 = v37
                                }
                            }
                            if ((!(1 != 65535)) && true) {
                                v2 = v2
                                switch v22 {
                                case "\n\t":
                                    var v53 uint8 = 128
                                    fmt.Println(v3, v53)
                                    var v54 int8 = int8(v2)
                                    if false {
                                        fmt.Println(v37)
                                        if v3 {
                                            v55 := GenS6{F0: 255, F1: 32767}
                                            _ = v55
                                        }
                                        v56 := v1[v37]
                                        _ = v56
                                    }
                                    _ = v53
                                    _ = v54
                                case "":
                                    if true {
                                        switch fi42 {
                                        case 4294967295:
                                            v22 = ("a" + "\n\t")
                                        }
                                    } else {
                                        if (false || ((fi50 ^ (fi50 % ((fi42) | 1))) != fi42)) {
                                            var v57 float32 = 0.0
                                            fmt.Println(v22, v2)
                                            _ = v57
                                        }
                                        var v58 float32 = ((0.0 - 0.1) * float32((1e10 + (1.5 - 3.14159))))
                                        var v59 uint64 = 4294967295
                                        _ = v58
                                        _ = v59
                                    }
                                }
                            } else {
                                fmt.Println(v22)
                            }
                            var v60 string = v22
                            var v61 int = fi50
                            _ = fi50
                            _ = v60
                            _ = v61
                        }
                        _ = fi42
                    }
                } else {
                    v2 = v2
                    v62 := GenS7{F0: genFn0(), F1: 0, F2: v37}
                    if false {
                        for fi63 := 0; fi63 < 6; fi63++ {
                            var v64 uint16 = 65535
                            _ = fi63
                            _ = v64
                        }
                    } else {
                        _ = len(v1)
                    }
                    _ = v62
                }
                v2 += (v2 >> ((1) & 63))
                _ = v37
            }
            var v65 uint = 255
            fmt.Println(v2)
            _ = v65
        }
        v66 := GenS10{F0: 304, F1: (v2 & (v2 / (((v2 * v2)) | 1))), F2: 4294967295}
        v2 -= v2
        _ = v66
    }
    v21 = append(v21, -9223372036854775808)
    var v67 string = v22
    var v68 uint8 = 42
    v69 := GenS10{F0: 289, F1: (v2 - v2), F2: 255}
    for fi70 := 0; fi70 < 6; fi70++ {
        switch "foo bar" {
        case "a":
            for fi71 := 0; fi71 < 10; fi71++ {
                v72 := GenS4{F0: v68, F1: 515}
                var v73 int = 128
                v2 += 0
                _ = fi71
                _ = v72
                _ = v73
            }
        }
        v67 = v22
        var v74 float64 = 0.0
        _ = fi70
        _ = v74
    }
    v2 = (v2 % ((v2) | 1))
    for fi75 := 0; fi75 < 4; fi75++ {
        for fi76 := 0; fi76 < 13; fi76++ {
            v77 := GenS8{F0: 2147483647, F1: 1, F2: -19}
            switch v67 {
            case "hello":
                v68 = v68
                switch v67 {
                case "\n\t":
                    var v78 int8 = 42
                    _ = v78
                case "foo bar":
                    v21 = append(v21, 32767)
                case "hello":
                    v68 -= 18
                    v2 += v2
                    v67 = v67
                default:
                    switch "a" {
                    case "":
                        v21 = append(v21, (^(fi75)))
                    case "\n\t":
                        switch (^(fi75)) {
                        case -2147483648:
                            var v79 int32 = -1
                            _ = v79
                        case -9223372036854775808:
                            fmt.Println(v22, v22)
                        default:
                            var v80 int8 = 100
                            _ = v80
                        }
                    }
                }
            case "\n\t":
                if ((423 > -2147483648) && (673 == 490)) {
                    v81 := GenS2{F0: (v68 / (((v68 + (v68 ^ (v68 ^ (v68 | v68))))) | 1)), F1: 4294967295}
                    var v82 uint64 = 1
                    if v3 {
                        var v83 string = v22
                        var v84 float64 = 0.0
                        switch "a" {
                        case "":
                            fmt.Println(v84)
                        default:
                            var v85 int8 = 10
                            _ = v85
                        }
                        _ = v83
                        _ = v84
                    }
                    _ = v81
                    _ = v82
                } else {
                    if v3 {
                        if v3 {
                            if v3 {
                                var v86 uint16 = 42
                                _ = v86
                            }
                            v87 := GenS9{F0: 2147483647}
                            if (!false) {
                                fmt.Println(v22, v67, v68)
                                v1[2147483647] = v22
                                switch 142 {
                                case 2147483647:
                                    var v88 uint64 = uint64(v68)
                                    _ = len(v21)
                                    fmt.Println(fi76, fi76, v22)
                                    v89 := GenS11{F0: (v22 + ((v67 + (v67 + (v67 + (v22 + "hello")))) + v67)), F1: -1, F2: -1}
                                    v88 -= (v88 - v88)
                                    _ = v88
                                    _ = v89
                                }
                                var v90 float32 = 0.1
                                v67 = v67
                                _ = v90
                            } else {
                                fmt.Println(v67)
                            }
                            _ = v87
                        } else {
                            v21 = append(v21, (-(fi75)))
                        }
                    }
                    v91 := v1[887]
                    _ = v91
                }
            case "a":
                v3 = false
            default:
                if v3 {
                    v3 = (3.14159 >= 1e10)
                    fmt.Println(fi75, v2, v67)
                } else {
                    fmt.Println(v22, v3)
                    fmt.Println(v68, fi75, fi76)
                    if v3 {
                        var v92 int64 = int64(fi76)
                        fmt.Println(v2, v3, fi75)
                        switch (-(fi75)) {
                        case 1:
                            fmt.Println(v68)
                            v93 := GenS12{F0: genFn1((v68 & (v68 % ((0) | 1))), 0, (v92 - 1)), F1: v2, F2: 535}
                            if false {
                                v2 -= v2
                            }
                            if (v3 || ((!((((3.14159 - 1.5) * (1.5 / 1e10)) + -2.25) <= (1.5 * 0.0))) && v3)) {
                                fmt.Println(v22, v92)
                                switch (-(fi76)) {
                                case 9223372036854775807:
                                    var v94 uint8 = v68
                                    _ = v94
                                case 32767:
                                    if len(v21) >= 2 {
                                        v21 = v21[1:len(v21)]
                                    }
                                }
                                v92 = v92
                            } else {
                                if (false || (!false)) {
                                    var v95 uint16 = 266
                                    var v96 bool = v3
                                    _ = v95
                                    _ = v96
                                } else {
                                    v3 = (3.14159 != ((0.0 / 3.14159) / -2.25))
                                }
                            }
                            v2 -= v2
                            _ = v93
                        case 32768:
                            _ = len(v21)
                        }
                        _ = v92
                    } else {
                        v68 += 255
                        v67 = v67
                        v1[0] = ((v67 + ("a" + (v67 + ((v22 + (v22 + v22)) + v22)))) + v22)
                    }
                }
            }
            switch (fi75 / (((-(fi75))) | 1)) {
            case 32767:
                var v97 int8 = -28
                _ = v97
            default:
                var v98 int8 = 127
                _ = v98
            }
            _ = fi76
            _ = v77
        }
        fmt.Println(v22, v2, v2)
        _ = fi75
    }
    var v99 uint32 = 65535
    switch v67 {
    case "a":
        v67 = v67
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v67:", v67)
    fmt.Println("v68:", v68)
    fmt.Println("v99:", v99)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v30
    _ = v67
    _ = v68
    _ = v69
    _ = v99
}
