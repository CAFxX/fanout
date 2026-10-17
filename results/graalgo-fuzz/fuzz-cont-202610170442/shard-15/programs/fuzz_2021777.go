package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 bool; F2 int32 }
type GenS2 struct { F0 int; F1 int16; F2 int32 }
type GenS3 struct { F0 int16; F1 int }
type GenS4 struct { F0 uint; F1 uint64; F2 uint }
type GenS5 struct { F0 uint32; F1 int; F2 bool }
type GenS6 struct { F0 uint32; F1 int8; F2 string }
type GenS7 struct { F0 int16 }
type GenS8 struct { F0 int16; F1 uint32 }
func genFn0() uint {
    var v5 bool = false
    v6 := GenS1{F0: 1, F1: (42 > 255), F2: -2147483648}
    var v7 uint64 = 1
    for fi8 := 0; fi8 < 8; fi8++ {
        if v5 {
            if v5 {
                v5 = (false && false)
                var v9 uint = 4294967295
                var v10 uint64 = (v7 | v7)
                _ = v10
                _ = v9
            } else {
                v7 = v7
                v7 -= v7
            }
            if false {
                if v5 {
                    for fi11 := 0; fi11 < 18; fi11++ {
                        v5 = true
                        v12 := []int{42}
                        var v13 int16 = 1
                        var v14 uint64 = (v7 / (((^(v7))) | 1))
                        v14 -= (v14 >> ((2) & 63))
                        if len(v12) > 0 {
                            v12[0] = fi8
                        }
                        _ = fi11
                        _ = v12
                        _ = v13
                        _ = v14
                    }
                    v15 := GenS2{F0: fi8, F1: 32767, F2: -2147483648}
                    if (1e10 >= 3.14159) {
                        switch ("\n\t" + "") {
                        case "foo bar":
                            v16 := []int{(fi8 - fi8)}
                            _ = v16
                        default:
                            var v17 uint = 221
                            _ = v17
                        }
                        if false {
                            v18 := []int{}
                            fmt.Println(fi8)
                            v7 -= (v7 << ((0) & 63))
                            _ = v18
                        } else {
                            v19 := []int{fi8, fi8, (fi8 * 1), fi8}
                            if false {
                                for fi20 := 0; fi20 < 9; fi20++ {
                                    if (v5 || false) {
                                        v7 += v7
                                        if v5 {
                                            var v21 int8 = -1
                                            v22 := GenS1{F0: -248, F1: (542 == 4294967295), F2: 603}
                                            _ = v21
                                            _ = v22
                                        } else {
                                            var v23 uint8 = 255
                                            v24 := GenS2{F0: 9223372036854775807, F1: 545, F2: -2147483648}
                                            var v25 bool = v5
                                            _ = v23
                                            _ = v24
                                            _ = v25
                                        }
                                        if len(v19) >= 2 {
                                            v19 = v19[1:len(v19)]
                                        }
                                    }
                                    var v26 bool = v5
                                    fmt.Println(v5, v26)
                                    _ = fi20
                                    _ = v26
                                }
                            }
                            var v27 uint = 255
                            _ = v19
                            _ = v27
                        }
                    } else {
                        v28 := []int{}
                        if v5 {
                            _ = len(v28)
                            var v29 int32 = 877
                            _ = v29
                        }
                        if v5 {
                            for fi30 := 0; fi30 < 19; fi30++ {
                                v7 = 661
                                _ = fi30
                            }
                            fmt.Println(v5, v5, v5)
                        }
                        _ = v28
                    }
                    _ = v15
                } else {
                    v5 = (float64(1.5) > 1e10)
                }
                for fi31 := 0; fi31 < 8; fi31++ {
                    var v32 float64 = (1e10 / 1.0)
                    _ = fi31
                    _ = v32
                }
                v33 := []int{(fi8 + (fi8 / (((fi8 ^ fi8)) | 1)))}
                _ = v33
            } else {
                for fi34 := 0; fi34 < 2; fi34++ {
                    v7 -= (v7 / ((v7) | 1))
                    _ = fi34
                }
                v7 += (v7 % (((^(v7))) | 1))
            }
        }
        _ = fi8
    }
    _ = v5
    _ = v6
    _ = v7
    return 263
}
func genFn1(p0 uint, p1 uint, p2 uint16) uint {
    for fi35 := 0; fi35 < 15; fi35++ {
        v36 := GenS1{F0: 271, F1: false, F2: 0}
        for fi37 := 0; fi37 < 16; fi37++ {
            v38 := GenS3{F0: 0, F1: fi35}
            _ = fi37
            _ = v38
        }
        _ = fi35
        _ = v36
    }
    if (1.5 < (3.14159 - float32(1.5))) {
        for fi39 := 0; fi39 < 10; fi39++ {
            if (true && (2147483647 < 2147483648)) {
                p0 += (p0 & 0)
            }
            var v40 string = "foo bar"
            var v41 int16 = 32767
            _ = fi39
            _ = v40
            _ = v41
        }
        switch ("hello" + (("\n\t" + (("foo bar" + ("" + ("a" + ("" + ("foo bar" + "\n\t"))))) + "")) + "a")) {
        case "a":
            for fi42 := 0; fi42 < 2; fi42++ {
                p0 = p1
                var v43 int64 = 9223372036854775807
                _ = fi42
                _ = v43
            }
        }
    }
    _ = p0
    _ = p1
    _ = p2
    return 1
}

func main() {
    v1 := []string{"hello"}
    v2 := []int{32767}
    var v3 string = "\n\t"
    var v4 float64 = 1.5
    for fi44 := 0; fi44 < 18; fi44++ {
        if ((!(("\n\t" + v3) != ("a" + "hello"))) || true) {
            for fi45 := 0; fi45 < 18; fi45++ {
                var v46 float32 = 3.14159
                v47 := GenS2{F0: fi45, F1: 336, F2: 701}
                _ = fi45
                _ = v46
                _ = v47
            }
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        }
        if (!false) {
            if len(v2) > 0 {
                v2[0] = 128
            }
            var v48 float64 = float64(3.14159)
            _ = v48
        } else {
            for fi49 := 0; fi49 < 7; fi49++ {
                v4 = v4
                _ = fi49
            }
            fmt.Println(fi44)
            v4 -= float64(1e10)
        }
        _ = fi44
    }
    var v50 int32 = 42
    switch 4294967295 {
    case 127:
        if false {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            switch "hello" {
            case "foo bar":
                var v51 uint32 = 0
                _ = v51
            case "\n\t":
                var v52 int32 = (v50 | (v50 & (v50 % (((^(v50))) | 1))))
                _ = v52
            default:
                switch (v3 + (("\n\t" + v3) + ("a" + "foo bar"))) {
                case "a":
                    v53 := GenS1{F0: -1, F1: (false && (!(!true))), F2: (v50 / ((v50) | 1))}
                    _ = v53
                case "":
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                case "foo bar":
                    var v54 float64 = float64(float32(0.0))
                    _ = v54
                default:
                    for fi55 := 0; fi55 < 12; fi55++ {
                        var v56 int64 = 0
                        if true {
                            switch (("hello" + ((v3 + (v3 + v3)) + "a")) + "\n\t") {
                            case "":
                                switch "foo bar" {
                                case "a":
                                    if (v3 < v3) {
                                        v56 -= (v56 / ((v56) | 1))
                                        v50 = (v50 >> ((4294967295) & 63))
                                    }
                                case "foo bar":
                                    v3 = v3
                                case "":
                                    fmt.Println(v4)
                                }
                            case "a":
                                if true {
                                    var v57 int8 = 1
                                    _ = v57
                                }
                            }
                        }
                        _ = fi55
                        _ = v56
                    }
                }
            }
            v58 := GenS4{F0: 65535, F1: 18446744073709551615, F2: genFn1(0, 18446744073709551615, 42)}
            _ = v58
        } else {
            switch 569 {
            case 585:
                if true {
                    for fi59 := 0; fi59 < 13; fi59++ {
                        if (true || (1 == 32767)) {
                            for fi60 := 0; fi60 < 1; fi60++ {
                                var v61 float32 = 1.5
                                _ = fi60
                                _ = v61
                            }
                            var v62 int64 = -9223372036854775808
                            v4 -= v4
                            _ = v62
                        } else {
                            for fi63 := 0; fi63 < 8; fi63++ {
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                                var v64 int64 = -845
                                var v65 int = fi63
                                _ = fi63
                                _ = v64
                                _ = v65
                            }
                            fmt.Println(v3)
                        }
                        v66 := GenS2{F0: fi59, F1: 1000, F2: v50}
                        if (!(!(2147483647 != 42))) {
                            var v67 uint = 305
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            _ = v67
                        }
                        _ = fi59
                        _ = v66
                    }
                    v68 := GenS5{F0: 65535, F1: 2147483647, F2: (255 <= 161)}
                    _ = v68
                }
            case 127:
                v50 = v50
            }
            v50 = v50
        }
    case 366:
        if (33 < 0) {
            v2 = append(v2, 1)
            var v69 uint = 4294967295
            if false {
                if true {
                    v50 = (-(v50))
                } else {
                    if (!(255 >= 0)) {
                        _ = len(v2)
                        if len(v1) > 0 {
                            v1[0] = (v3 + v3)
                        }
                    }
                }
                _ = len(v2)
                v4 = (1e10 * 0.1)
            }
            _ = v69
        }
    case -764:
        v70 := GenS2{F0: 811, F1: -999, F2: v50}
        _ = v70
    default:
        var v71 bool = false
        _ = v71
    }
    var v72 int = -403
    if (1 >= 124) {
        fmt.Println(v3, v50, v50)
        fmt.Println(v3, v50, v72)
        v4 = (((v4 + v4) + float64(3.14159)) + (v4 - 0.1))
    } else {
        for fi73 := 0; fi73 < 14; fi73++ {
            var v74 int32 = v50
            var v75 uint16 = 688
            _ = fi73
            _ = v74
            _ = v75
        }
        var v76 uint = 65535
        v1 = append(v1, (((("" + v3) + (v3 + v3)) + ("foo bar" + v3)) + "\n\t"))
        _ = v76
    }
    v77 := GenS6{F0: 2147483648, F1: 42, F2: v3}
    _ = len(v2)
    switch v3 {
    case "foo bar":
        var v78 uint32 = 65535
        _ = v78
    case "\n\t":
        if (!false) {
            if (false || (false && true)) {
                v4 -= v4
            }
        }
    case "":
        for fi79 := 0; fi79 < 13; fi79++ {
            var v80 string = v3
            _ = fi79
            _ = v80
        }
    default:
        v3 = (v3 + "foo bar")
    }
    switch v72 {
    case -2147483648:
        switch (v72 | v72) {
        case -9223372036854775808:
            var v81 uint = 2
            _ = v81
        case 639:
            v3 = (v3 + (((v3 + v3) + (v3 + v3)) + (v3 + "foo bar")))
        default:
            v82 := GenS7{F0: -32768}
            _ = v82
        }
    case 32767:
        v83 := GenS8{F0: -391, F1: 2}
        _ = v83
    default:
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    }
    for fi84 := 0; fi84 < 19; fi84++ {
        v50 -= (v50 + 1)
        _ = fi84
    }
    if len(v1) > 0 {
        v1[0] = v3
    }
    var v85 int = (v72 | v72)
    if len(v2) > 0 {
        v2[0] = (v72 * (v85 - -2147483648))
    }
    v3 = v3
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v50:", v50)
    fmt.Println("v72:", v72)
    fmt.Println("v85:", v85)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v50
    _ = v72
    _ = v77
    _ = v85
}
