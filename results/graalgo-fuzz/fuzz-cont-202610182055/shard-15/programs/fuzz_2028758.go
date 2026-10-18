package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int; F2 int64 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint16; F1 int8; F2 bool }
type GenS4 struct { F0 bool; F1 uint32; F2 int32 }
type GenS5 struct { F0 uint64; F1 int16 }
type GenS6 struct { F0 uint; F1 uint32 }
type GenS7 struct { F0 int; F1 int16; F2 uint8 }
type GenS8 struct { F0 uint64 }
type GenS9 struct { F0 uint64; F1 int64; F2 int64 }
type GenS10 struct { F0 bool }
type GenS11 struct { F0 uint64 }

func main() {
    var v1 uint64 = 11
    v2 := make(map[uint32]uint32)
    var v3 float32 = 1.5
    var v4 uint64 = v1
    v5 := []string{"hello"}
    var v6 int8 = -1
    if (-9223372036854775808 > 9223372036854775807) {
        if (702 < 325) {
            v7 := GenS1{F0: 42, F1: 2147483647, F2: -660}
            switch 2147483647 {
            case 2147483647:
                v4 = v4
            case 812:
                var v8 string = "a"
                _ = v8
            case 42:
                _ = len(v5)
            }
            for fi9 := 0; fi9 < 11; fi9++ {
                var v10 float64 = (-2.25 * (0.0 - (((1e10 + -2.25) / 3.14159) * 0.1)))
                _ = fi9
                _ = v10
            }
            _ = v7
        }
        if true {
            v11 := GenS2{F0: 0}
            if false {
                switch "hello" {
                case "\n\t":
                    var v12 uint32 = 2147483648
                    _ = v12
                case "foo bar":
                    var v13 int8 = v6
                    _ = v13
                }
                v1 = v4
            }
            fmt.Println(v1, v4, v1)
            _ = v11
        } else {
            if (130 == 376) {
                for fi14 := 0; fi14 < 17; fi14++ {
                    v15 := v2[2]
                    for fi16 := 0; fi16 < 13; fi16++ {
                        v15 -= 4294967295
                        _ = fi16
                    }
                    _ = fi14
                    _ = v15
                }
                if false {
                    v3 = ((((v3 - (1e10 * (v3 * v3))) - (v3 / float32(1.5))) * 0.0) - (v3 + v3))
                    v17 := GenS3{F0: 42, F1: -128, F2: (1 != 65535)}
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    _ = v17
                }
                if true {
                    switch 32767 {
                    case -9223372036854775808:
                        fmt.Println(v4)
                    case 345:
                        v1 = 2
                    case -2147483648:
                        var v18 string = "\n\t"
                        _ = v18
                    }
                } else {
                    var v19 bool = (255 != 42)
                    v20 := GenS2{F0: -641}
                    fmt.Println(v19, v6)
                    _ = v19
                    _ = v20
                }
            }
            var v21 int16 = -231
            if true {
                var v22 int32 = -401
                fmt.Println(v6, v22, v1)
                if (!true) {
                    if ((166 == 225) && true) {
                        fmt.Println(v6, v3, v4)
                    } else {
                        v21 += (v21 | (v21 / (((v21 >> ((460) & 63))) | 1)))
                        switch 2147483647 {
                        case 42:
                            if true {
                                v3 += v3
                                var v23 int16 = (v21 / ((v21) | 1))
                                _ = v23
                            }
                        case 9223372036854775807:
                            var v24 uint32 = uint32(v4)
                            _ = v24
                        case 32768:
                            switch 0 {
                            case -480:
                                v3 = v3
                            case 1:
                                fmt.Println(v3)
                            case -815:
                                if false {
                                    var v25 float64 = float64(v3)
                                    for fi26 := 0; fi26 < 10; fi26++ {
                                        if len(v5) > 0 {
                                            v5[0] = "a"
                                        }
                                        _ = len(v2)
                                        v27 := GenS4{F0: false, F1: 4294967295, F2: (v22 ^ v22)}
                                        _ = fi26
                                        _ = v27
                                    }
                                    v1 += (v4 & v4)
                                    _ = v25
                                }
                            default:
                                for fi28 := 0; fi28 < 15; fi28++ {
                                    if len(v5) >= 2 {
                                        v5 = v5[1:len(v5)]
                                    }
                                    fmt.Println(v22, v4)
                                    _ = fi28
                                }
                            }
                        default:
                            v5 = append(v5, "foo bar")
                        }
                    }
                    v5 = append(v5, "hello")
                    for fi29 := 0; fi29 < 9; fi29++ {
                        switch (fi29 & 2147483647) {
                        case -1:
                            switch (fi29 ^ fi29) {
                            case 402:
                                _ = len(v2)
                            case -804:
                                switch ("" + ("hello" + "")) {
                                case "\n\t":
                                    v30 := GenS4{F0: ((^(v22)) > v22), F1: 65535, F2: (^(v22))}
                                    _ = v30
                                case "":
                                    for fi31 := 0; fi31 < 4; fi31++ {
                                        switch "a" {
                                        case "a":
                                            var v32 uint8 = 2
                                            _ = v32
                                        case "\n\t":
                                            switch 9223372036854775807 {
                                            case 1:
                                                v33 := GenS2{F0: fi31}
                                                _ = v33
                                            case 32768:
                                                v34 := GenS2{F0: (-(fi29))}
                                                _ = v34
                                            }
                                        case "":
                                            if false {
                                                var v35 int32 = (^(v22))
                                                v36 := GenS5{F0: 1, F1: (v21 - v21)}
                                                var v37 bool = true
                                                _ = v35
                                                _ = v36
                                                _ = v37
                                            }
                                        default:
                                            if (!false) {
                                                var v38 float32 = v3
                                                var v39 string = ("a" + "\n\t")
                                                v22 -= 68
                                                _ = v38
                                                _ = v39
                                            }
                                        }
                                        _ = fi31
                                    }
                                case "foo bar":
                                    v22 -= 42
                                }
                            case -790:
                                v22 -= -394
                            }
                        default:
                            fmt.Println(v22, v1)
                        }
                        _ = fi29
                    }
                } else {
                    var v40 bool = (!false)
                    switch "foo bar" {
                    case "a":
                        if (127 <= 1) {
                            var v41 float32 = float32(1e10)
                            if v40 {
                                v40 = v40
                            } else {
                                for fi42 := 0; fi42 < 16; fi42++ {
                                    for fi43 := 0; fi43 < 1; fi43++ {
                                        if true {
                                            if v40 {
                                                v44 := GenS6{F0: uint(v4), F1: 1}
                                                if v40 {
                                                    var v45 float32 = 1e10
                                                    v40 = (4294967295 > 39)
                                                    _ = v45
                                                } else {
                                                    var v46 int = (fi42 % ((int(v6)) | 1))
                                                    if (v40 || ((v3 < (v41 - (v41 * (v3 * v3)))) || (1e10 == ((1.5 / (-2.25 * (1.5 / 1.0))) / ((-2.25 / 1e10) / (0.1 - (float64(((-2.25 - (0.0 / v41)) * 0.0)) + 0.0))))))) {
                                                        var v47 bool = v40
                                                        _ = v47
                                                    } else {
                                                        v48 := GenS5{F0: (v1 & v1), F1: 32767}
                                                        _ = v48
                                                    }
                                                    v3 += 1e10
                                                    if true {
                                                        if v40 {
                                                            switch fi42 {
                                                            case 989:
                                                                v3 = ((float32(3.14159) - v3) + v3)
                                                            case 0:
                                                                if (!false) {
                                                                    switch (-(v46)) {
                                                                    case 4294967295:
                                                                        v5 = append(v5, "a")
                                                                    default:
                                                                        var v49 bool = (false && true)
                                                                        _ = v49
                                                                    }
                                                                }
                                                                delete(v2, 203)
                                                                var v50 string = "a"
                                                                _ = v50
                                                            case 32767:
                                                                if (42 <= 65535) {
                                                                    var v51 int8 = (-(v6))
                                                                    _ = v51
                                                                }
                                                            }
                                                        } else {
                                                            if ((("foo bar" + "foo bar") + "\n\t") != "\n\t") {
                                                                v21 = (^(v21))
                                                                v52 := GenS2{F0: fi43}
                                                                _ = v52
                                                            }
                                                            v41 = v41
                                                            fmt.Println(v22, v41)
                                                        }
                                                        if ((v6 & v6) < 42) {
                                                            v53 := GenS7{F0: fi42, F1: v21, F2: 42}
                                                            _ = v53
                                                        } else {
                                                            var v54 int8 = 127
                                                            var v55 bool = (!(false && ((-2.25 != 3.14159) || v40)))
                                                            var v56 uint8 = 127
                                                            var v57 int16 = v21
                                                            _ = v54
                                                            _ = v55
                                                            _ = v56
                                                            _ = v57
                                                        }
                                                    }
                                                    _ = v46
                                                }
                                                _ = v44
                                            }
                                            switch ("\n\t" + "a") {
                                            case "":
                                                var v58 float32 = v41
                                                _ = v58
                                            }
                                        }
                                        v21 -= (-(v21))
                                        _ = fi43
                                    }
                                    _ = fi42
                                }
                                var v59 uint32 = 0
                                _ = len(v2)
                                _ = v59
                            }
                            var v60 int = 1
                            _ = v41
                            _ = v60
                        } else {
                            _ = len(v5)
                            var v61 int64 = -9223372036854775808
                            _ = v61
                        }
                    case "foo bar":
                        switch ("hello" + ("\n\t" + ("a" + ""))) {
                        case "foo bar":
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                        case "a":
                            v21 += v21
                        }
                    case "hello":
                        if v40 {
                            v62 := GenS8{F0: (v1 * v4)}
                            v63 := GenS6{F0: 1, F1: 65535}
                            _ = v62
                            _ = v63
                        }
                    }
                    _ = v40
                }
                _ = v22
            } else {
                if (((-2.25 / v3) * v3) < v3) {
                    for fi64 := 0; fi64 < 17; fi64++ {
                        if (483 <= 431) {
                            for fi65 := 0; fi65 < 2; fi65++ {
                                var v66 int8 = v6
                                if ((^(v21)) > v21) {
                                    var v67 uint16 = 0
                                    switch "a" {
                                    case "\n\t":
                                        fmt.Println(v4, v6)
                                    case "":
                                        if ((^(v1)) <= v4) {
                                            var v68 int = fi65
                                            v21 += v21
                                            fmt.Println(v67, v4, v3)
                                            _ = v68
                                        } else {
                                            v69 := GenS9{F0: (v4 & v1), F1: -1, F2: -9223372036854775808}
                                            var v70 uint8 = 42
                                            v67 = v67
                                            _ = v69
                                            _ = v70
                                        }
                                    case "foo bar":
                                        var v71 uint32 = 2
                                        _ = v71
                                    default:
                                        if true {
                                            if ((v67 + (v67 - 0)) <= v67) {
                                                v21 = 903
                                            }
                                        }
                                    }
                                    v72 := GenS7{F0: 9223372036854775807, F1: (v21 / ((42) | 1)), F2: 128}
                                    _ = v67
                                    _ = v72
                                }
                                _ = fi65
                                _ = v66
                            }
                            fmt.Println(v4)
                            var v73 int16 = (v21 - (v21 & (v21 + 1000)))
                            _ = v73
                        }
                        _ = len(v5)
                        _ = fi64
                    }
                    delete(v2, 767)
                    v74 := GenS10{F0: (127 > -994)}
                    _ = v74
                }
            }
            _ = v21
        }
        if len(v5) > 0 {
            v5[0] = "foo bar"
        }
    } else {
        var v75 uint = 0
        _ = v75
    }
    v6 -= v6
    fmt.Println(v4, v1)
    var v76 int16 = -32768
    var v77 float32 = (v3 - v3)
    var v78 int16 = 1000
    var v79 int32 = 1
    v4 -= (v1 ^ 0)
    for fi80 := 0; fi80 < 15; fi80++ {
        var v81 uint8 = 1
        v82 := GenS6{F0: 459, F1: 2}
        v81 += v81
        _ = fi80
        _ = v81
        _ = v82
    }
    _ = len(v2)
    if false {
        v83 := GenS3{F0: 55, F1: 127, F2: (-550 != 128)}
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        v84 := GenS6{F0: 0, F1: 832}
        _ = v83
        _ = v84
    } else {
        _ = len(v2)
    }
    var v85 uint = 18446744073709551615
    if false {
        if len(v5) > 0 {
            v5[0] = "a"
        }
        if false {
            var v86 int = 1
            if false {
                var v87 int32 = (v79 - -434)
                v88 := v2[1]
                v89 := GenS6{F0: (^(v85)), F1: (v88 << ((v85) & 63))}
                _ = v87
                _ = v88
                _ = v89
            } else {
                v5 = append(v5, "foo bar")
            }
            _ = v86
        }
    } else {
        if true {
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            fmt.Println(v85, v79, v76)
        } else {
            switch "foo bar" {
            case "foo bar":
                var v90 uint64 = (v4 - v4)
                _ = v90
            case "":
                v91 := GenS8{F0: (^(v4))}
                _ = v91
            }
        }
        v92 := GenS7{F0: 42, F1: v78, F2: 128}
        v93 := GenS11{F0: (v1 | v1)}
        _ = v92
        _ = v93
    }
    v94 := v2[2]
    var v95 uint32 = 484
    fmt.Println(v77)
    v96 := GenS9{F0: v4, F1: -9223372036854775808, F2: 42}
    var v97 int = 2147483647
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v76:", v76)
    fmt.Println("v77:", v77)
    fmt.Println("v78:", v78)
    fmt.Println("v79:", v79)
    fmt.Println("v85:", v85)
    fmt.Println("v94:", v94)
    fmt.Println("v95:", v95)
    fmt.Println("v97:", v97)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v76
    _ = v77
    _ = v78
    _ = v79
    _ = v85
    _ = v94
    _ = v95
    _ = v96
    _ = v97
}
