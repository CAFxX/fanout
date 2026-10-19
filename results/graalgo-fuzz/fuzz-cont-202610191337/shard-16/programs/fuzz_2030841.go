package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int32; F2 int32 }
type GenS2 struct { F0 uint16; F1 int8 }
type GenS3 struct { F0 string; F1 int64 }
type GenS4 struct { F0 int16; F1 uint }
type GenS5 struct { F0 uint8; F1 uint8 }
type GenS6 struct { F0 bool; F1 uint32 }
type GenS7 struct { F0 uint64 }
type GenS8 struct { F0 int16; F1 int64; F2 uint64 }
func genFn0(p0 int64, p1 int) uint16 {
    p1 = (^(p1))
    v7 := []int{int(p0), (p1 >> ((4294967295) & 63)), (^(p1)), 4294967295}
    var v8 uint8 = 1
    fmt.Println(p0)
    var v9 bool = (!(528 >= 4294967295))
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    _ = v9
    return 42
}
func genFn1(p0 uint16, p1 uint8) int32 {
    var v10 int16 = -32768
    p0 -= (^(p0))
    switch ("foo bar" + "foo bar") {
    case "a":
        var v11 int8 = -128
        _ = v11
    case "hello":
        switch -2147483648 {
        case 9223372036854775807:
            p0 = 65535
        case 127:
            p0 += genFn0(-459, int(p0))
        default:
            v12 := GenS1{F0: "a", F1: 1, F2: 42}
            _ = v12
        }
    default:
        switch "" {
        case "\n\t":
            switch "hello" {
            case "":
                var v13 float32 = 0.1
                _ = v13
            case "foo bar":
                for fi14 := 0; fi14 < 17; fi14++ {
                    switch 4294967295 {
                    case -9223372036854775808:
                        for fi15 := 0; fi15 < 6; fi15++ {
                            p0 = (p0 / ((255) | 1))
                            v16 := []int{1, -1, fi14, (fi14 ^ 42)}
                            var v17 uint16 = 2
                            _ = fi15
                            _ = v16
                            _ = v17
                        }
                    }
                    fmt.Println(p1)
                    _ = fi14
                }
            }
        case "foo bar":
            v18 := []int{-427, 446, -645}
            _ = v18
        case "":
            v19 := GenS1{F0: "", F1: -651, F2: 42}
            _ = v19
        default:
            switch 9223372036854775807 {
            case -2147483648:
                v20 := GenS1{F0: "", F1: -849, F2: -2147483648}
                _ = v20
            case -551:
                if true {
                    if true {
                        var v21 uint32 = 2147483647
                        for fi22 := 0; fi22 < 7; fi22++ {
                            switch "hello" {
                            case "":
                                var v23 string = ("\n\t" + "hello")
                                _ = v23
                            case "hello":
                                switch ("foo bar" + "a") {
                                case "a":
                                    for fi24 := 0; fi24 < 17; fi24++ {
                                        p0 = p0
                                        var v25 int64 = -556
                                        var v26 int16 = 32767
                                        if true {
                                            fmt.Println(fi24, v26, fi22)
                                            fmt.Println(v21, p0, v25)
                                            p1 -= 0
                                            var v27 bool = (!(!false))
                                            _ = v27
                                        } else {
                                            v26 -= (v10 + v26)
                                        }
                                        _ = fi24
                                        _ = v25
                                        _ = v26
                                    }
                                case "":
                                    p1 = p1
                                default:
                                    fmt.Println(p1, p0)
                                }
                            case "\n\t":
                                v21 = v21
                            }
                            for fi28 := 0; fi28 < 20; fi28++ {
                                p1 = 70
                                var v29 int32 = 42
                                v30 := GenS1{F0: ("\n\t" + ("" + "foo bar")), F1: (v29 ^ v29), F2: (v29 % ((701) | 1))}
                                _ = fi28
                                _ = v29
                                _ = v30
                            }
                            _ = fi22
                        }
                        _ = v21
                    } else {
                        if (2147483647 == 0) {
                            var v31 bool = true
                            for fi32 := 0; fi32 < 19; fi32++ {
                                v33 := []int{(fi32 ^ (fi32 << ((18446744073709551615) & 63))), fi32, (fi32 >> ((1) & 63))}
                                _ = fi32
                                _ = v33
                            }
                            p1 -= (p1 % ((p1) | 1))
                            _ = v31
                        } else {
                            v34 := []int{-306, 759, -9223372036854775808}
                            _ = v34
                        }
                        switch 739 {
                        case 32768:
                            p1 = (p1 % (((p1 - 1)) | 1))
                        default:
                            if (((("\n\t" + "foo bar") + ("hello" + "hello")) + (("\n\t" + "\n\t") + ("foo bar" + "a"))) != (("a" + ("hello" + "foo bar")) + "")) {
                                v10 = (v10 ^ v10)
                            } else {
                                v35 := GenS1{F0: "hello", F1: -2147483648, F2: -1}
                                var v36 float64 = 3.14159
                                fmt.Println(p1, v10, p1)
                                _ = v35
                                _ = v36
                            }
                        }
                    }
                }
            case 1:
                switch ("a" + "foo bar") {
                case "\n\t":
                    v10 += 387
                case "foo bar":
                    var v37 int16 = (v10 >> ((18446744073709551615) & 63))
                    _ = v37
                case "a":
                    fmt.Println(p0, p1)
                default:
                    v38 := GenS1{F0: "foo bar", F1: 1, F2: 0}
                    _ = v38
                }
            default:
                switch -587 {
                case -9223372036854775808:
                    if false {
                        v39 := []int{9223372036854775807, 1, 679, 342}
                        _ = v39
                    } else {
                        p0 = (p0 >> ((42) & 63))
                        var v40 int32 = -1
                        _ = v40
                    }
                case 32767:
                    var v41 int = -819
                    _ = v41
                }
            }
        }
    }
    for fi42 := 0; fi42 < 10; fi42++ {
        for fi43 := 0; fi43 < 20; fi43++ {
            if (!(p1 >= uint8(p0))) {
                var v44 int8 = 1
                switch fi42 {
                case 1:
                    var v45 uint32 = 2147483647
                    _ = v45
                case 32768:
                    fmt.Println(fi42)
                }
                p1 -= (p1 >> ((18446744073709551615) & 63))
                _ = v44
            } else {
                p1 = (p1 / (((p1 / ((p1) | 1))) | 1))
                switch "hello" {
                case "a":
                    v46 := GenS1{F0: "a", F1: 0, F2: 1}
                    _ = v46
                case "\n\t":
                    if (980 > 18446744073709551615) {
                        var v47 int32 = 0
                        var v48 bool = ((false || true) || (42 == 401))
                        if (1e10 < 3.14159) {
                            var v49 int8 = 1
                            var v50 int = 127
                            v47 += 962
                            v51 := GenS2{F0: (p0 * (p0 - (p0 / ((p0) | 1)))), F1: v49}
                            var v52 bool = ((^(p1)) != p1)
                            _ = v49
                            _ = v50
                            _ = v51
                            _ = v52
                        }
                        _ = v47
                        _ = v48
                    } else {
                        if true {
                            fmt.Println(fi43, p1, fi42)
                            v53 := GenS2{F0: p0, F1: 42}
                            p0 = (^(p0))
                            var v54 uint8 = 221
                            _ = v53
                            _ = v54
                        } else {
                            v55 := GenS1{F0: "", F1: 2147483647, F2: 42}
                            v56 := []int{4294967295, -888, fi43}
                            v57 := GenS3{F0: ("a" + ("foo bar" + "")), F1: 42}
                            _ = v55
                            _ = v56
                            _ = v57
                        }
                        var v58 float32 = 1e10
                        v10 = (v10 * v10)
                        var v59 uint = 4294967295
                        var v60 uint8 = (p1 + p1)
                        v61 := GenS4{F0: -956, F1: v59}
                        _ = v58
                        _ = v59
                        _ = v60
                        _ = v61
                    }
                default:
                    v10 += 609
                }
            }
            v62 := []int{0}
            if false {
                v63 := GenS1{F0: "a", F1: 2147483647, F2: 1}
                v10 = (v10 >> ((4294967295) & 63))
                switch "" {
                case "hello":
                    _ = len(v62)
                case "a":
                    v10 += (v10 / ((v10) | 1))
                }
                _ = v63
            } else {
                v10 += (v10 >> ((157) & 63))
                if ((-(v10)) >= (v10 ^ v10)) {
                    switch 0 {
                    case 9223372036854775807:
                        p1 = (p1 / (((p1 | (p1 + 249))) | 1))
                    case 512:
                        if true {
                            v64 := GenS1{F0: "foo bar", F1: 1, F2: 2147483647}
                            _ = v64
                        } else {
                            _ = len(v62)
                            fmt.Println(p1)
                            if len(v62) >= 2 {
                                v62 = v62[1:len(v62)]
                            }
                            var v65 bool = false
                            _ = v65
                        }
                    case -63:
                        var v66 float64 = 1.5
                        _ = v66
                    default:
                        fmt.Println(v10, p0, fi42)
                        switch (-(fi43)) {
                        case -413:
                            if len(v62) >= 2 {
                                v62 = v62[1:len(v62)]
                            }
                        case 4294967295:
                            p1 += 53
                        case 809:
                            switch "foo bar" {
                            case "\n\t":
                                v67 := GenS5{F0: (p1 & 26), F1: p1}
                                _ = v67
                            case "a":
                                if true {
                                    v68 := GenS4{F0: -1, F1: 18446744073709551615}
                                    _ = v68
                                }
                            case "":
                                if true {
                                    if (20 >= 4294967295) {
                                        if len(v62) >= 2 {
                                            v62 = v62[1:len(v62)]
                                        }
                                        switch 0 {
                                        case 0:
                                            switch (-(fi43)) {
                                            case 939:
                                                var v69 uint8 = 0
                                                _ = v69
                                            case 4294967295:
                                                v70 := GenS6{F0: true, F1: 65535}
                                                _ = v70
                                            case -1:
                                                p1 = 42
                                                v10 += v10
                                            default:
                                                v10 += (^(v10))
                                            }
                                        case -9223372036854775808:
                                            v10 = -32768
                                            var v71 uint = 42
                                            _ = v71
                                        }
                                        v62 = append(v62, fi43)
                                    } else {
                                        fmt.Println(p1, v10)
                                        if false {
                                            p0 += (^(p0))
                                            if len(v62) > 0 {
                                                v62[0] = fi42
                                            }
                                        } else {
                                            v72 := GenS1{F0: (("a" + "foo bar") + "hello"), F1: 1, F2: 1}
                                            _ = v72
                                        }
                                        if false {
                                            if len(v62) >= 2 {
                                                v62 = v62[1:len(v62)]
                                            }
                                            var v73 float32 = (float32(0.0) + 0.0)
                                            var v74 uint16 = p0
                                            _ = v73
                                            _ = v74
                                        }
                                    }
                                    p1 += p1
                                    if (-347 <= -872) {
                                        switch "a" {
                                        case "hello":
                                            var v75 float64 = 0.1
                                            _ = v75
                                        }
                                        v76 := GenS7{F0: 2}
                                        _ = v76
                                    } else {
                                        v62 = append(v62, -9223372036854775808)
                                        v10 = -649
                                        p1 = 2
                                    }
                                }
                            default:
                                fmt.Println(fi43, p0, v10)
                            }
                        }
                    }
                } else {
                    switch "" {
                    case "foo bar":
                        fmt.Println(p0, p0, fi43)
                    case "":
                        v77 := GenS4{F0: v10, F1: 18446744073709551615}
                        _ = v77
                    default:
                        if false {
                            var v78 bool = (p0 >= (p0 + p0))
                            var v79 bool = (171 >= (p0 % ((genFn0(-506, -2147483648)) | 1)))
                            _ = v78
                            _ = v79
                        }
                    }
                    if (0 > 42) {
                        var v80 float64 = 0.1
                        _ = v80
                    } else {
                        var v81 uint8 = p1
                        p0 = p0
                        var v82 int8 = 127
                        v83 := GenS7{F0: 18446744073709551615}
                        _ = v81
                        _ = v82
                        _ = v83
                    }
                    var v84 int16 = (v10 - (v10 << ((255) & 63)))
                    _ = v84
                }
                _ = len(v62)
                v85 := GenS8{F0: v10, F1: 104, F2: 438}
                _ = v85
            }
            switch ("" + ("" + "a")) {
            case "\n\t":
                if len(v62) >= 2 {
                    v62 = v62[1:len(v62)]
                }
            default:
                fmt.Println(fi42, fi43)
            }
            _ = fi43
            _ = v62
        }
        var v86 int32 = -1
        _ = fi42
        _ = v86
    }
    _ = p0
    _ = p1
    _ = v10
    return -374
}

func main() {
    var v1 bool = false
    var v2 int = -536
    var v3 uint64 = 2
    var v4 int64 = 0
    v5 := []uint32{1}
    var v6 int8 = 34
    if v1 {
        v3 = (^(v3))
    }
    for fi87 := 0; fi87 < 20; fi87++ {
        switch ("foo bar" + "") {
        case "a":
            var v88 uint8 = 22
            _ = v88
        case "hello":
            v3 = v3
        case "\n\t":
            if true {
                fmt.Println(fi87, v6, v1)
            }
        }
        _ = fi87
    }
    fmt.Println(v4, v1, v3)
    var v89 string = "foo bar"
    for fi90 := 0; fi90 < 4; fi90++ {
        for fi91 := 0; fi91 < 17; fi91++ {
            v92 := GenS8{F0: 42, F1: v4, F2: 18446744073709551615}
            var v93 int8 = v6
            _ = fi91
            _ = v92
            _ = v93
        }
        _ = fi90
    }
    for fi94 := 0; fi94 < 17; fi94++ {
        var v95 int64 = -988
        _ = fi94
        _ = v95
    }
    var v96 int = (^(v2))
    for fi97 := 0; fi97 < 16; fi97++ {
        var v98 int = -9223372036854775808
        _ = fi97
        _ = v98
    }
    v3 = (^(v3))
    var v99 string = (v89 + v89)
    v4 = 9223372036854775807
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    var v100 int = (v96 - 9223372036854775807)
    fmt.Println("v1:", v1)
    fmt.Println("v100:", v100)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v89:", v89)
    fmt.Println("v96:", v96)
    fmt.Println("v99:", v99)
    _ = v1
    _ = v100
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v89
    _ = v96
    _ = v99
}
