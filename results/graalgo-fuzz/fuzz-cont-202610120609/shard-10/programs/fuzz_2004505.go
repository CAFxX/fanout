package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint }
type GenS2 struct { F0 int16; F1 bool }
type GenS3 struct { F0 string; F1 int64; F2 int64 }
type GenS4 struct { F0 uint8; F1 bool }
type GenS5 struct { F0 uint8; F1 uint32 }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 uint64; F1 int32; F2 bool }
type GenS8 struct { F0 uint; F1 int64; F2 int64 }
type GenS9 struct { F0 string }
type GenS10 struct { F0 int32; F1 uint32 }
type GenS11 struct { F0 uint32; F1 uint16; F2 uint32 }
type GenS12 struct { F0 int16; F1 bool }
func genFn0() bool {
    for fi5 := 0; fi5 < 9; fi5++ {
        switch fi5 {
        case 1:
            v6 := GenS1{F0: 0, F1: 4294967295}
            _ = v6
        case 2147483647:
            var v7 int32 = 0
            _ = v7
        case 0:
            var v8 string = "foo bar"
            _ = v8
        default:
            var v9 float64 = 3.14159
            _ = v9
        }
        _ = fi5
    }
    var v10 uint16 = 614
    if ((true && false) || true) {
        v11 := []int{-9223372036854775808, 10}
        var v12 uint16 = (v10 / ((68) | 1))
        _ = v11
        _ = v12
    }
    switch "" {
    case "foo bar":
        v13 := []int{32768, -43, 0}
        _ = v13
    }
    switch 2147483647 {
    case 128:
        fmt.Println(v10)
    case 1:
        v14 := []int{0, -356, 348, 2147483647}
        _ = v14
    }
    _ = v10
    return true
}
func genFn1(p0 uint64, p1 int) uint {
    var v15 int32 = 2147483647
    v16 := GenS2{F0: 32767, F1: (!(2147483647 < 0))}
    if ((((((((0.0 + 1.5) / (0.0 + (3.14159 / (((1e10 / 1.0) - 1.5) - 0.1)))) + 0.1) * (((1e10 + 3.14159) + (-2.25 * (1e10 / 1e10))) - float32((1e10 * (-2.25 + (0.1 / (0.1 + 3.14159))))))) / (1.5 + 0.1)) + 0.0) / -2.25) != 0.1) {
        for fi17 := 0; fi17 < 6; fi17++ {
            p1 -= p1
            var v18 int64 = 9223372036854775807
            _ = fi17
            _ = v18
        }
    } else {
        var v19 int64 = 0
        v20 := []int{(^(p1)), p1}
        v21 := GenS2{F0: -1, F1: false}
        _ = v19
        _ = v20
        _ = v21
    }
    v15 -= 10
    v22 := []int{(p1 + (p1 >> ((2) & 63)))}
    _ = p0
    _ = p1
    _ = v15
    _ = v16
    _ = v22
    return 0
}

func main() {
    v1 := make(map[uint32]int)
    var v2 int8 = -1
    var v3 float64 = 1e10
    var v4 uint8 = 68
    v23 := GenS1{F0: -1, F1: 672}
    v24 := []int{-9223372036854775808}
    v2 = v2
    var v25 string = "hello"
    for fi26 := 0; fi26 < 10; fi26++ {
        var v27 uint8 = 0
        if (!true) {
            if (1 <= 2147483648) {
                var v28 string = v25
                if (v2 >= (v2 * -42)) {
                    var v29 uint64 = 263
                    v30 := GenS2{F0: 400, F1: (true || (!false))}
                    _ = v29
                    _ = v30
                } else {
                    v2 = v2
                    var v31 uint32 = 2147483647
                    _ = v31
                }
                _ = v28
            }
            if len(v24) >= 2 {
                v24 = v24[1:len(v24)]
            }
            if true {
                v32 := GenS3{F0: "\n\t", F1: 9223372036854775807, F2: 42}
                v2 += v2
                var v33 int8 = (^(v2))
                _ = v32
                _ = v33
            } else {
                for fi34 := 0; fi34 < 3; fi34++ {
                    v4 -= (v27 + 127)
                    var v35 uint16 = 32767
                    _ = fi34
                    _ = v35
                }
                for fi36 := 0; fi36 < 2; fi36++ {
                    v37 := GenS1{F0: -101, F1: 4294967295}
                    _ = fi36
                    _ = v37
                }
            }
        } else {
            switch 127 {
            case 9223372036854775807:
                fmt.Println(v3, fi26)
            case -601:
                v38 := v1[630]
                _ = v38
            default:
                v39 := GenS4{F0: v4, F1: true}
                _ = v39
            }
        }
        if (v27 <= v4) {
            v40 := GenS2{F0: 32767, F1: (!true)}
            if false {
                v4 -= (v4 / (((v27 >> ((255) & 63))) | 1))
                fmt.Println(v4, v3)
                delete(v1, 65535)
            } else {
                var v41 bool = (!((false || false) && false))
                for fi42 := 0; fi42 < 12; fi42++ {
                    var v43 uint64 = 1
                    var v44 string = v25
                    _ = fi42
                    _ = v43
                    _ = v44
                }
                if v41 {
                    _ = len(v1)
                    for fi45 := 0; fi45 < 4; fi45++ {
                        v41 = (1 < -2147483648)
                        v25 = v25
                        _ = fi45
                    }
                    v4 = v4
                } else {
                    fmt.Println(v4)
                    if (18446744073709551615 <= 526) {
                        v27 += (v4 / (((v27 | v4)) | 1))
                        fmt.Println(v2, v25, v27)
                        v46 := GenS5{F0: (v4 / ((v4) | 1)), F1: 65535}
                        _ = v46
                    }
                    fmt.Println(v4)
                }
                _ = v41
            }
            _ = v40
        } else {
            for fi47 := 0; fi47 < 9; fi47++ {
                v27 = 127
                _ = fi47
            }
            fmt.Println(v3, v3, v4)
            v48 := GenS6{F0: 2147483647}
            _ = v48
        }
        _ = fi26
        _ = v27
    }
    switch 4294967295 {
    case -751:
        fmt.Println(v4, v25)
    default:
        _ = len(v24)
    }
    fmt.Println(v4, v25)
    for fi49 := 0; fi49 < 7; fi49++ {
        if (255 <= 16) {
            v50 := GenS5{F0: v4, F1: 587}
            if len(v24) >= 2 {
                v24 = v24[1:len(v24)]
            }
            if false {
                if (false && (v25 >= (v25 + v25))) {
                    v51 := GenS4{F0: v4, F1: (false || false)}
                    _ = v51
                } else {
                    var v52 int32 = 0
                    if true {
                        v1[2] = fi49
                        if false {
                            v3 -= v3
                        }
                    }
                    _ = v52
                }
                for fi53 := 0; fi53 < 8; fi53++ {
                    var v54 int = fi53
                    switch 128 {
                    case 789:
                        v55 := GenS1{F0: 1, F1: 255}
                        _ = v55
                    case -407:
                        var v56 uint8 = v4
                        if false {
                            v57 := GenS4{F0: v4, F1: (true && (true && true))}
                            _ = v57
                        } else {
                            switch v25 {
                            case "":
                                if ((!(!(39 >= 65535))) && false) {
                                    if ((-32768 == -34) || false) {
                                        if true {
                                            fmt.Println(fi53)
                                            v54 += (fi49 | fi49)
                                            v25 = "foo bar"
                                            var v58 string = "a"
                                            v2 += v2
                                            _ = v58
                                        } else {
                                            switch (fi49 * (fi53 >> ((42) & 63))) {
                                            case -1:
                                                var v59 uint8 = (v4 - (v56 << ((4294967295) & 63)))
                                                _ = v59
                                            case 127:
                                                v2 += v2
                                            }
                                        }
                                        var v60 uint32 = 260
                                        _ = v60
                                    } else {
                                        v61 := GenS7{F0: 159, F1: 147, F2: true}
                                        _ = len(v1)
                                        v4 = (v4 << ((18446744073709551615) & 63))
                                        _ = v61
                                    }
                                    var v62 uint8 = 42
                                    _ = v62
                                } else {
                                    switch 934 {
                                    case 42:
                                        if false {
                                            fmt.Println(v54, v2)
                                            if len(v24) > 0 {
                                                v24[0] = 1
                                            }
                                            v2 += v2
                                        }
                                    case 536:
                                        if (42 < -1) {
                                            _ = len(v24)
                                        } else {
                                            v25 = ((("foo bar" + "a") + v25) + v25)
                                            fmt.Println(v2, v56)
                                            v63 := GenS8{F0: 2, F1: 42, F2: -859}
                                            _ = v63
                                        }
                                        var v64 string = (v25 + (v25 + v25))
                                        _ = v64
                                    case 1:
                                        v65 := GenS3{F0: v25, F1: -1, F2: int64(v56)}
                                        _ = v65
                                    }
                                    if (!(!true)) {
                                        v66 := GenS5{F0: (v4 ^ v56), F1: 2147483648}
                                        _ = len(v1)
                                        if false {
                                            var v67 uint64 = 621
                                            if (((v67 / ((v67) | 1)) < 0) && (v25 == v25)) {
                                                v68 := v1[998]
                                                _ = v68
                                            }
                                            var v69 uint16 = 32767
                                            if len(v24) > 0 {
                                                v24[0] = (-(fi49))
                                            }
                                            var v70 uint64 = 589
                                            _ = v67
                                            _ = v69
                                            _ = v70
                                        } else {
                                            v71 := GenS9{F0: "foo bar"}
                                            var v72 bool = true
                                            switch "a" {
                                            case "\n\t":
                                                v25 = (v25 + (v25 + (v25 + "foo bar")))
                                            case "":
                                                v56 = v4
                                            case "foo bar":
                                                var v73 int8 = v2
                                                v1[200] = v54
                                                _ = v73
                                            default:
                                                if v72 {
                                                    if (869 <= 1) {
                                                        var v74 float32 = (float32(1e10) * -2.25)
                                                        switch fi53 {
                                                        case 2147483647:
                                                            fmt.Println(v72)
                                                        case 9223372036854775807:
                                                            var v75 int16 = 55
                                                            _ = v75
                                                        }
                                                        fmt.Println(v25, v74)
                                                        v76 := GenS10{F0: 2147483647, F1: 798}
                                                        _ = v74
                                                        _ = v76
                                                    } else {
                                                        v24 = append(v24, (fi53 | v54))
                                                        var v77 uint64 = 1
                                                        _ = v77
                                                    }
                                                    v2 = (v2 & (v2 << ((65535) & 63)))
                                                    if (false || v72) {
                                                        if true {
                                                            v25 = v25
                                                        }
                                                        var v78 int32 = -658
                                                        if (!(42 > 1)) {
                                                            v54 = (fi49 << ((0) & 63))
                                                            v79 := GenS7{F0: 596, F1: v78, F2: true}
                                                            _ = v79
                                                        } else {
                                                            var v80 int16 = 1
                                                            v81 := GenS6{F0: 1}
                                                            var v82 uint32 = 0
                                                            v83 := GenS11{F0: v82, F1: 32767, F2: 0}
                                                            _ = v80
                                                            _ = v81
                                                            _ = v82
                                                            _ = v83
                                                        }
                                                        var v84 int8 = (v2 * (v2 / (((v2 * (^(v2)))) | 1)))
                                                        v1[2] = -9223372036854775808
                                                        v1[2147483647] = v54
                                                        _ = v78
                                                        _ = v84
                                                    }
                                                } else {
                                                    v72 = v72
                                                    v72 = true
                                                }
                                            }
                                            switch (v25 + "") {
                                            case "foo bar":
                                                fmt.Println(v4, v3, v25)
                                            default:
                                                v25 = "\n\t"
                                            }
                                            switch v54 {
                                            case -2147483648:
                                                var v85 uint = 581
                                                v1[915] = (fi49 ^ (fi49 & (fi49 << (((v85 * v85)) & 63))))
                                                v86 := GenS6{F0: 42}
                                                _ = v85
                                                _ = v86
                                            }
                                            fmt.Println(fi49, v25)
                                            _ = v71
                                            _ = v72
                                        }
                                        _ = v66
                                    }
                                    v4 = (v56 >> ((4294967295) & 63))
                                }
                            case "hello":
                                _ = len(v1)
                            case "foo bar":
                                var v87 bool = false
                                _ = v87
                            }
                            switch (fi53 << ((4294967295) & 63)) {
                            case 4294967295:
                                switch (v54 | (-(fi53))) {
                                case 42:
                                    switch "hello" {
                                    case "hello":
                                        v88 := v1[788]
                                        _ = v88
                                    }
                                }
                            case -2147483648:
                                v25 = "a"
                            case 1:
                                fmt.Println(v25)
                            default:
                                v2 += (v2 ^ 42)
                            }
                            v56 = 71
                            v89 := GenS12{F0: 1000, F1: (((^(v4)) > v56) || true)}
                            _ = v89
                        }
                        _ = v56
                    case 42:
                        v25 = ((v25 + "foo bar") + v25)
                    default:
                        v54 = 572
                    }
                    _ = fi53
                    _ = v54
                }
            } else {
                fmt.Println(v25, v2)
                v90 := GenS1{F0: 286, F1: 18446744073709551615}
                _ = v90
            }
            _ = v50
        }
        if len(v24) >= 2 {
            v24 = v24[1:len(v24)]
        }
        v2 = (v2 >> ((0) & 63))
        _ = fi49
    }
    _ = len(v24)
    v4 -= (v4 & (v4 | 128))
    switch v25 {
    case "hello":
        v3 += ((float64(0.1) / (v3 - 1e10)) / v3)
    }
    v4 = v4
    for fi91 := 0; fi91 < 13; fi91++ {
        v1[272] = int(v4)
        v25 = (v25 + "")
        v25 = ((("" + v25) + ((v25 + "a") + v25)) + "")
        _ = fi91
    }
    if true {
        fmt.Println(v25, v3, v3)
        _ = len(v1)
    }
    var v92 bool = (true || (!(!false)))
    var v93 bool = v92
    var v94 int64 = 1
    var v95 string = v25
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v92:", v92)
    fmt.Println("v93:", v93)
    fmt.Println("v94:", v94)
    fmt.Println("v95:", v95)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v92
    _ = v93
    _ = v94
    _ = v95
}
