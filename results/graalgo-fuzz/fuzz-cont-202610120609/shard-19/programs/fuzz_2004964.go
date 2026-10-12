package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint32; F1 uint32; F2 uint8 }
type GenS4 struct { F0 int; F1 uint64; F2 int }
type GenS5 struct { F0 int32; F1 string; F2 uint }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 int64 }
func genFn0(p0 uint, p1 uint16, p2 int32) int16 {
    if ((42 < 0) || true) {
        if false {
            var v7 int8 = 1
            var v8 uint16 = p1
            _ = v7
            _ = v8
        } else {
            if (((!(!false)) || true) && (2147483647 == 2)) {
                var v9 uint32 = 648
                if true {
                    v10 := []int{32768, int(p1), 32767, -1}
                    _ = v10
                }
                _ = v9
            }
        }
        p1 += p1
    }
    var v11 uint64 = uint64(p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    return 32767
}
func genFn1(p0 uint64, p1 uint) string {
    fmt.Println(p0)
    v12 := GenS1{F0: 431, F1: (p1 << (((^(p1))) & 63))}
    fmt.Println(p0, p1)
    _ = p0
    _ = p1
    _ = v12
    return "hello"
}

func main() {
    var v1 string = "foo bar"
    v2 := make(map[uint32]uint32)
    v3 := []uint32{0}
    v4 := make(map[uint32]int)
    var v5 uint = 244
    var v6 uint8 = 39
    v13 := GenS2{F0: genFn0((v5 ^ (^(v5))), 0, 375)}
    for fi14 := 0; fi14 < 1; fi14++ {
        var v15 int16 = 42
        v5 = 0
        _ = fi14
        _ = v15
    }
    fmt.Println(v5, v1, v5)
    for fi16 := 0; fi16 < 4; fi16++ {
        v5 = v5
        v5 = 42
        fmt.Println(v6, v1, v1)
        _ = fi16
    }
    var v17 uint = (^(v5))
    v5 -= 486
    v5 = 657
    if (2 != 908) {
        if false {
            if true {
                var v18 float64 = (0.0 * float64(0.0))
                _ = v18
            } else {
                for fi19 := 0; fi19 < 15; fi19++ {
                    switch (v1 + v1) {
                    case "foo bar":
                        for fi20 := 0; fi20 < 20; fi20++ {
                            fmt.Println(v1, v17, fi20)
                            switch v1 {
                            case "\n\t":
                                var v21 int64 = -1
                                _ = v21
                            case "foo bar":
                                switch -1 {
                                case 9223372036854775807:
                                    v22 := GenS1{F0: 42, F1: v17}
                                    v1 = "foo bar"
                                    var v23 int8 = int8(v5)
                                    _ = v22
                                    _ = v23
                                case -883:
                                    if false {
                                        v5 = v17
                                        v24 := GenS3{F0: 65535, F1: 348, F2: (v6 / (((v6 + v6)) | 1))}
                                        fmt.Println(v6)
                                        _ = v24
                                    }
                                default:
                                    fmt.Println(fi20, v5, v5)
                                }
                            }
                            if false {
                                switch 42 {
                                case 53:
                                    v5 -= 4294967295
                                }
                            } else {
                                v5 = 42
                            }
                            _ = fi20
                        }
                    case "hello":
                        for fi25 := 0; fi25 < 5; fi25++ {
                            fmt.Println(fi19, v6, v1)
                            _ = fi25
                        }
                    case "\n\t":
                        v26 := GenS4{F0: int(v6), F1: 934, F2: (fi19 - fi19)}
                        _ = v26
                    default:
                        v6 = (v6 - (v6 / ((v6) | 1)))
                    }
                    var v27 int8 = 50
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    _ = fi19
                    _ = v27
                }
                for fi28 := 0; fi28 < 18; fi28++ {
                    v4[2147483648] = -2147483648
                    var v29 string = v1
                    _ = fi28
                    _ = v29
                }
                if len(v3) > 0 {
                    v3[0] = 725
                }
            }
            fmt.Println(v5)
        }
        fmt.Println(v1)
        switch 128 {
        case -935:
            var v30 uint32 = 65535
            _ = v30
        case -462:
            var v31 uint32 = 340
            _ = v31
        case -1:
            fmt.Println(v6, v17, v17)
        default:
            switch -23 {
            case -71:
                for fi32 := 0; fi32 < 19; fi32++ {
                    v3 = append(v3, 0)
                    v1 = v1
                    if true {
                        switch int(v6) {
                        case 32768:
                            fmt.Println(v17)
                        case 0:
                            var v33 int = 9223372036854775807
                            _ = v33
                        case 127:
                            _ = len(v3)
                        default:
                            var v34 float64 = (float64(1.5) - 0.1)
                            _ = v34
                        }
                        for fi35 := 0; fi35 < 10; fi35++ {
                            var v36 int16 = 1
                            fmt.Println(v17, v17)
                            var v37 int8 = -1
                            _ = fi35
                            _ = v36
                            _ = v37
                        }
                        switch "hello" {
                        case "a":
                            var v38 float32 = 1.5
                            _ = v38
                        }
                    }
                    _ = fi32
                }
            case 32767:
                var v39 uint64 = 0
                _ = v39
            default:
                var v40 uint64 = 533
                _ = v40
            }
        }
    }
    if ((v1 + (v1 + v1)) == "hello") {
        v41 := GenS4{F0: 2147483647, F1: 42, F2: 128}
        _ = v41
    }
    v6 = 1
    for fi42 := 0; fi42 < 3; fi42++ {
        v6 = v6
        v17 = 42
        _ = fi42
    }
    v43 := GenS4{F0: 32768, F1: 1, F2: 32767}
    for fi44 := 0; fi44 < 4; fi44++ {
        for fi45 := 0; fi45 < 1; fi45++ {
            var v46 uint64 = 768
            if (65535 <= 331) {
                v6 = v6
            }
            _ = fi45
            _ = v46
        }
        if true {
            _ = len(v3)
            _ = len(v4)
            if false {
                if (!true) {
                    for fi47 := 0; fi47 < 1; fi47++ {
                        v5 = (v5 >> (((v5 - v17)) & 63))
                        _ = fi47
                    }
                    if true {
                        for fi48 := 0; fi48 < 19; fi48++ {
                            var v49 int16 = genFn0(2, 2, 393)
                            var v50 int64 = -1
                            var v51 uint = v17
                            _ = fi48
                            _ = v49
                            _ = v50
                            _ = v51
                        }
                        var v52 int32 = -820
                        if ((!true) || (!((false || ((v6 - 2) >= v6)) && ((42 > genFn0(v17, 0, -1)) || (!(!(591 != 1))))))) {
                            if (-107 != -1) {
                                if (!false) {
                                    for fi53 := 0; fi53 < 12; fi53++ {
                                        if (false && false) {
                                            fmt.Println(v5, v17, v1)
                                            v5 -= (^(v5))
                                            var v54 int8 = 0
                                            v1 = (v1 + "hello")
                                            _ = v54
                                        } else {
                                            v52 = v52
                                        }
                                        switch ((v1 + v1) + v1) {
                                        case "hello":
                                            var v55 uint32 = 2
                                            _ = v55
                                        }
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                        _ = fi53
                                    }
                                    v5 = (v17 - (v17 + (v17 % (((^(v17))) | 1))))
                                    v5 = v5
                                } else {
                                    for fi56 := 0; fi56 < 17; fi56++ {
                                        v6 = (^(v6))
                                        v6 += v6
                                        v57 := v2[880]
                                        var v58 uint = v17
                                        v17 = 42
                                        switch fi56 {
                                        case 0:
                                            if (fi44 < 2147483647) {
                                                var v59 uint64 = 4294967295
                                                fmt.Println(v57)
                                                v60 := GenS5{F0: (v52 * (v52 / (((v52 | v52)) | 1))), F1: "a", F2: 1}
                                                _ = v59
                                                _ = v60
                                            }
                                        case -1:
                                            var v61 bool = (v17 > (v58 % ((725) | 1)))
                                            _ = v61
                                        case -2147483648:
                                            var v62 int = 9223372036854775807
                                            _ = v62
                                        default:
                                            v17 -= uint(v57)
                                        }
                                        _ = fi56
                                        _ = v57
                                        _ = v58
                                    }
                                    var v63 uint = (^(v17))
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                    _ = v63
                                }
                            } else {
                                switch (fi44 + fi44) {
                                case 206:
                                    v64 := GenS6{F0: v52}
                                    _ = v64
                                default:
                                    v17 = v17
                                }
                                v65 := GenS2{F0: 1}
                                var v66 uint = (^(v17))
                                _ = v65
                                _ = v66
                            }
                            v17 = v17
                            v1 = "hello"
                        } else {
                            fmt.Println(v6, fi44)
                            v52 -= v52
                            if len(v3) > 0 {
                                v3[0] = 42
                            }
                        }
                        _ = v52
                    }
                    v1 = v1
                }
                for fi67 := 0; fi67 < 17; fi67++ {
                    if (!false) {
                        var v68 int8 = -105
                        _ = v68
                    }
                    _ = fi67
                }
                var v69 int = (fi44 % ((fi44) | 1))
                _ = v69
            }
        } else {
            v70 := GenS3{F0: 383, F1: uint32(v5), F2: v6}
            v6 = (v6 / ((42) | 1))
            var v71 int32 = -1
            _ = v70
            _ = v71
        }
        _ = fi44
    }
    fmt.Println(v1, v17)
    v72 := GenS7{F0: 242}
    switch 128 {
    case -2147483648:
        v73 := GenS6{F0: 0}
        _ = v73
    case -177:
        if (4294967295 <= 0) {
            for fi74 := 0; fi74 < 17; fi74++ {
                switch "hello" {
                case "\n\t":
                    for fi75 := 0; fi75 < 11; fi75++ {
                        if true {
                            if true {
                                v5 = 0
                            }
                            switch (v1 + (v1 + v1)) {
                            case "\n\t":
                                var v76 int8 = 127
                                _ = v76
                            case "foo bar":
                                var v77 float64 = 1.5
                                _ = v77
                            }
                            var v78 int8 = 127
                            _ = v78
                        } else {
                            var v79 bool = (false || (((false || (!(1 != 520))) || true) || (true || (-7 >= -1))))
                            v80 := GenS5{F0: 0, F1: v1, F2: 4294967295}
                            var v81 float64 = 1e10
                            fmt.Println(v17, v17)
                            switch 32768 {
                            case 32768:
                                v81 -= (v81 / float64(1.5))
                            case -1:
                                v82 := GenS1{F0: 5, F1: (^(v5))}
                                _ = v82
                            }
                            _ = v79
                            _ = v80
                            _ = v81
                        }
                        v83 := GenS7{F0: 1}
                        _ = fi75
                        _ = v83
                    }
                case "hello":
                    for fi84 := 0; fi84 < 3; fi84++ {
                        var v85 uint64 = 95
                        v5 = (v17 | 255)
                        _ = fi84
                        _ = v85
                    }
                default:
                    var v86 int8 = 0
                    _ = v86
                }
                v1 = v1
                _ = fi74
            }
            v87 := GenS6{F0: -832}
            var v88 int = -892
            _ = v87
            _ = v88
        } else {
            fmt.Println(v5, v17, v5)
        }
    case 42:
        var v89 float32 = -2.25
        _ = v89
    default:
        v4[611] = -233
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v5
    _ = v6
    _ = v72
}
