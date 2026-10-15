package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8; F2 int16 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint16; F1 uint }
type GenS4 struct { F0 uint8; F1 int; F2 int64 }
type GenS5 struct { F0 int16; F1 int16; F2 uint32 }
type GenS6 struct { F0 int32; F1 int; F2 string }
type GenS7 struct { F0 int8; F1 uint16 }
type GenS8 struct { F0 int16; F1 uint32 }
type GenS9 struct { F0 int32; F1 int32 }
type GenS10 struct { F0 string; F1 uint64 }
type GenS11 struct { F0 uint64; F1 uint64 }
type GenS12 struct { F0 string; F1 uint; F2 bool }
func genFn0() int {
    for fi6 := 0; fi6 < 12; fi6++ {
        v7 := []int{(fi6 * fi6), 4294967295, -788, (fi6 / ((fi6) | 1))}
        var v8 int = (fi6 ^ fi6)
        var v9 int32 = 0
        _ = fi6
        _ = v7
        _ = v8
        _ = v9
    }
    v10 := GenS1{F0: 32768, F1: 0, F2: -1}
    for fi11 := 0; fi11 < 16; fi11++ {
        var v12 int64 = 1
        var v13 uint = 4294967295
        _ = fi11
        _ = v12
        _ = v13
    }
    _ = v10
    return 1
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 int64 = -1
    var v3 uint16 = 2
    var v4 uint32 = 2147483647
    var v5 uint = 1
    var v14 uint8 = 0
    switch 32767 {
    case -9223372036854775808:
        v15 := GenS2{F0: 32767}
        _ = v15
    case 1:
        v5 = (v5 - (v5 ^ (v5 % ((1) | 1))))
    case -1:
        v16 := GenS2{F0: -815}
        _ = v16
    }
    v17 := []int{-2147483648}
    if false {
        switch "hello" {
        case "foo bar":
            if (-2.25 <= (0.0 + 1.5)) {
                v18 := GenS3{F0: (^(v3)), F1: 2}
                _ = v18
            }
        case "":
            var v19 float32 = (1e10 + 0.1)
            _ = v19
        default:
            fmt.Println(v3, v3, v4)
        }
        for fi20 := 0; fi20 < 19; fi20++ {
            if true {
                for fi21 := 0; fi21 < 19; fi21++ {
                    v14 = 0
                    v22 := GenS3{F0: 770, F1: 562}
                    v3 = v3
                    if (false && (false && (127 > -128))) {
                        var v23 float64 = -2.25
                        _ = len(v17)
                        var v24 string = "\n\t"
                        _ = v23
                        _ = v24
                    } else {
                        if (!true) {
                            var v25 uint16 = 65535
                            _ = v25
                        } else {
                            if false {
                                v17 = append(v17, fi21)
                                v4 += (v4 ^ (v4 << (((^(v5))) & 63)))
                                v2 = (v2 | 9223372036854775807)
                            }
                            fmt.Println(fi21, fi20, fi21)
                            v17 = append(v17, int(v5))
                            var v26 uint8 = 78
                            _ = v26
                        }
                        v27 := v1[(v4 & v4)]
                        var v28 int32 = 0
                        var v29 string = ("" + "\n\t")
                        v30 := GenS3{F0: (v3 << ((uint(v3)) & 63)), F1: v5}
                        _ = v27
                        _ = v28
                        _ = v29
                        _ = v30
                    }
                    _ = fi21
                    _ = v22
                }
                v31 := v1[(^(v4))]
                v1[v4] = ((v5 - (^(v5))) == (^(v5)))
                _ = v31
            }
            for fi32 := 0; fi32 < 3; fi32++ {
                v33 := GenS4{F0: 42, F1: fi20, F2: (v2 ^ (v2 / ((v2) | 1)))}
                v1[v4] = (-102 <= -197)
                switch fi32 {
                case 0:
                    var v34 int8 = -82
                    _ = v34
                case 2147483647:
                    switch "\n\t" {
                    case "\n\t":
                        v35 := GenS2{F0: (-(fi32))}
                        _ = v35
                    case "a":
                        if (-2147483648 == 1) {
                            fmt.Println(v4, v5)
                            if true {
                                if (uint8(v4) >= v14) {
                                    if true {
                                        v36 := GenS5{F0: -32768, F1: 42, F2: uint32(v3)}
                                        fmt.Println(v3, v2, fi20)
                                        _ = v36
                                    } else {
                                        var v37 float32 = (0.1 - 1e10)
                                        var v38 int8 = 127
                                        v4 -= (^(v4))
                                        fmt.Println(v5, v5, v5)
                                        if (!false) {
                                            var v39 int8 = (-(v38))
                                            _ = v39
                                        }
                                        var v40 int64 = (-(v2))
                                        switch "a" {
                                        case "\n\t":
                                            v41 := GenS2{F0: (fi20 | genFn0())}
                                            _ = v41
                                        default:
                                            v40 = v40
                                            v4 = (v4 >> (((v5 ^ (^(v5)))) & 63))
                                        }
                                        _ = v37
                                        _ = v38
                                        _ = v40
                                    }
                                    _ = len(v1)
                                    v42 := GenS6{F0: 0, F1: fi20, F2: "\n\t"}
                                    switch "hello" {
                                    case "foo bar":
                                        var v43 int = genFn0()
                                        _ = v43
                                    case "\n\t":
                                        fmt.Println(fi32)
                                    }
                                    _ = v42
                                }
                                fmt.Println(fi32, v5, v4)
                            } else {
                                var v44 int = (fi20 * fi32)
                                if false {
                                    if false {
                                        if len(v17) >= 2 {
                                            v17 = v17[1:len(v17)]
                                        }
                                        if false {
                                            fmt.Println(v44, fi32)
                                        } else {
                                            if false {
                                                v14 += 128
                                            } else {
                                                var v45 uint16 = v3
                                                var v46 bool = (42 >= 955)
                                                v47 := GenS6{F0: 42, F1: v44, F2: "a"}
                                                _ = v45
                                                _ = v46
                                                _ = v47
                                            }
                                            v48 := GenS7{F0: 33, F1: (^(v3))}
                                            v49 := GenS4{F0: v14, F1: 2147483647, F2: (v2 | -1)}
                                            _ = v48
                                            _ = v49
                                        }
                                        v5 = v5
                                    }
                                } else {
                                    if true {
                                        v2 = v2
                                        var v50 uint = v5
                                        var v51 int64 = v2
                                        v52 := GenS7{F0: 75, F1: uint16(v14)}
                                        if ((!((!false) || (false || (1.5 < -2.25)))) && false) {
                                            v2 = v2
                                            if true {
                                                v17 = append(v17, (fi32 * v44))
                                            } else {
                                                v5 -= (v5 % ((uint(v4)) | 1))
                                                var v53 int = v44
                                                if true {
                                                    v14 += (v14 / ((v14) | 1))
                                                    var v54 bool = false
                                                    var v55 int8 = 127
                                                    switch v53 {
                                                    case -1:
                                                        v56 := GenS2{F0: int(v55)}
                                                        _ = v56
                                                    default:
                                                        v2 = v2
                                                    }
                                                    v17 = append(v17, (v44 << ((v50) & 63)))
                                                    _ = v54
                                                    _ = v55
                                                } else {
                                                    fmt.Println(fi20)
                                                    switch (v53 % (((fi32 ^ (fi20 + v44))) | 1)) {
                                                    case 127:
                                                        var v57 float32 = float32(0.1)
                                                        _ = v57
                                                    case 1:
                                                        v14 = (v14 & (^(v14)))
                                                    default:
                                                        var v58 uint16 = (v3 ^ (v3 + uint16(v5)))
                                                        _ = v58
                                                    }
                                                    v50 = (^(v5))
                                                }
                                                _ = v53
                                            }
                                            if ((v5 % (((v50 / (((^(v50))) | 1))) | 1)) == v5) {
                                                if len(v17) > 0 {
                                                    v17[0] = -948
                                                }
                                            }
                                        }
                                        _ = v50
                                        _ = v51
                                        _ = v52
                                    } else {
                                        v2 = v2
                                    }
                                    v5 -= v5
                                    fmt.Println(v4, fi32, fi32)
                                }
                                _ = v44
                            }
                            if ("\n\t" != "a") {
                                v14 -= uint8(v4)
                                v59 := GenS8{F0: 1, F1: v4}
                                var v60 string = "\n\t"
                                _ = v59
                                _ = v60
                            } else {
                                fmt.Println(v4, v14, v5)
                                v61 := GenS2{F0: 0}
                                var v62 float32 = 1.5
                                var v63 int32 = 2147483647
                                if true {
                                    v63 += v63
                                    v62 += v62
                                    v62 = 0.0
                                    v5 = v5
                                } else {
                                    v64 := GenS7{F0: -1, F1: 65535}
                                    var v65 uint64 = 4294967295
                                    _ = v64
                                    _ = v65
                                }
                                _ = v61
                                _ = v62
                                _ = v63
                            }
                        }
                    case "foo bar":
                        fmt.Println(v4)
                    }
                case 32767:
                    fmt.Println(v14, v2, fi20)
                    v66 := GenS4{F0: v14, F1: fi32, F2: (v2 | -1)}
                    _ = v66
                }
                _ = fi32
                _ = v33
            }
            _ = fi20
        }
        v67 := GenS9{F0: -1, F1: 511}
        _ = v67
    } else {
        var v68 int64 = 1
        _ = v68
    }
    var v69 uint64 = 2
    var v70 uint64 = v69
    for fi71 := 0; fi71 < 8; fi71++ {
        for fi72 := 0; fi72 < 2; fi72++ {
            v70 = (v70 << (((v5 + (v5 >> ((65535) & 63)))) & 63))
            _ = fi72
        }
        switch (fi71 + 32768) {
        case 127:
            v73 := GenS2{F0: 262}
            _ = v73
        case 32767:
            if len(v17) >= 2 {
                v17 = v17[1:len(v17)]
            }
        case -889:
            var v74 uint64 = (v69 ^ (^(v70)))
            _ = v74
        }
        if ((true || true) || false) {
            switch "a" {
            case "":
                var v75 float32 = (float32(-2.25) + -2.25)
                _ = v75
            case "hello":
                var v76 bool = (42 > 42)
                _ = v76
            case "\n\t":
                if (false || false) {
                    switch genFn0() {
                    case -1:
                        v4 -= v4
                    default:
                        fmt.Println(v70, v3, v70)
                    }
                    v4 = (v4 - v4)
                    var v77 uint16 = (v3 ^ (^(v3)))
                    _ = v77
                }
            }
        }
        _ = fi71
    }
    var v78 uint32 = (^(v4))
    if (!(-1 >= -165)) {
        v79 := GenS10{F0: "", F1: (^(v70))}
        fmt.Println(v78, v3)
        _ = v79
    } else {
        for fi80 := 0; fi80 < 13; fi80++ {
            for fi81 := 0; fi81 < 20; fi81++ {
                var v82 bool = (!((^(v3)) <= (v3 % ((v3) | 1))))
                var v83 bool = v82
                v1[2147483647] = ((-(fi81)) >= int(v78))
                _ = fi81
                _ = v82
                _ = v83
            }
            v84 := GenS8{F0: 933, F1: (v78 | (v4 % ((42) | 1)))}
            _ = fi80
            _ = v84
        }
        v5 = (v5 | uint(v14))
    }
    if false {
        v3 += uint16(v78)
        switch "a" {
        case "":
            if false {
                v85 := GenS11{F0: v69, F1: (v70 << ((42) & 63))}
                _ = v85
            } else {
                v86 := GenS10{F0: (("a" + "hello") + "hello"), F1: 945}
                _ = v86
            }
        }
    }
    v2 += v2
    for fi87 := 0; fi87 < 4; fi87++ {
        var v88 float64 = -2.25
        for fi89 := 0; fi89 < 4; fi89++ {
            v17 = append(v17, fi89)
            if false {
                fmt.Println(v14, v14, fi87)
                var v90 string = "hello"
                v91 := GenS12{F0: (v90 + v90), F1: (v5 & (v5 >> (((v5 & (v5 ^ (v5 & uint(v3))))) & 63))), F2: ((v70 | v70) != 710)}
                _ = v90
                _ = v91
            }
            _ = fi89
        }
        _ = fi87
        _ = v88
    }
    if (!("hello" >= "\n\t")) {
        for fi92 := 0; fi92 < 19; fi92++ {
            var v93 int8 = -13
            v2 = v2
            _ = fi92
            _ = v93
        }
        if (false || (v78 > 964)) {
            var v94 int16 = -627
            v3 += v3
            _ = v94
        }
    } else {
        v95 := GenS4{F0: v14, F1: 32767, F2: 42}
        _ = v95
    }
    var v96 bool = (!false)
    var v97 uint8 = v14
    var v98 int32 = 42
    v3 += (v3 + v3)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v69:", v69)
    fmt.Println("v70:", v70)
    fmt.Println("v78:", v78)
    fmt.Println("v96:", v96)
    fmt.Println("v97:", v97)
    fmt.Println("v98:", v98)
    _ = v1
    _ = v14
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v69
    _ = v70
    _ = v78
    _ = v96
    _ = v97
    _ = v98
}
