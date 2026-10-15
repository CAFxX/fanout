package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string }
type GenS2 struct { F0 int32; F1 uint16; F2 bool }
type GenS3 struct { F0 uint64; F1 bool; F2 bool }
type GenS4 struct { F0 int; F1 uint8; F2 int32 }
type GenS5 struct { F0 uint16; F1 bool; F2 uint64 }
type GenS6 struct { F0 string }
type GenS7 struct { F0 int64; F1 int16 }
type GenS8 struct { F0 int64 }
type GenS9 struct { F0 uint16; F1 uint64 }
type GenS10 struct { F0 uint64; F1 bool }
type GenS11 struct { F0 int; F1 uint8 }
type GenS12 struct { F0 bool }
func genFn0(p0 int, p1 int16) int64 {
    v4 := GenS1{F0: p0, F1: "hello"}
    v5 := GenS1{F0: p0, F1: ((("a" + "foo bar") + "foo bar") + ("a" + ("" + "")))}
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return -766
}

func main() {
    var v1 int32 = 42
    var v2 uint64 = 18446744073709551615
    v3 := []uint32{2147483647, 2147483648}
    for fi6 := 0; fi6 < 13; fi6++ {
        var v7 int32 = (v1 | v1)
        var v8 uint32 = 43
        if false {
            if (255 < 42) {
                for fi9 := 0; fi9 < 2; fi9++ {
                    switch "\n\t" {
                    case "\n\t":
                        v3 = append(v3, 793)
                    case "foo bar":
                        var v10 int32 = (v7 ^ (v1 + 42))
                        _ = v10
                    case "a":
                        _ = len(v3)
                    default:
                        v8 += v8
                    }
                    v7 -= v1
                    v11 := GenS1{F0: (fi6 - -105), F1: "foo bar"}
                    _ = fi9
                    _ = v11
                }
            }
            for fi12 := 0; fi12 < 19; fi12++ {
                var v13 string = "a"
                switch v13 {
                case "a":
                    v14 := GenS2{F0: v1, F1: 32767, F2: false}
                    _ = v14
                case "\n\t":
                    var v15 int = -1
                    _ = v15
                case "":
                    fmt.Println(v13, v2, v8)
                }
                _ = fi12
                _ = v13
            }
            var v16 uint = 900
            _ = v16
        } else {
            v2 += v2
        }
        _ = fi6
        _ = v7
        _ = v8
    }
    var v17 uint32 = 42
    for fi18 := 0; fi18 < 6; fi18++ {
        var v19 int32 = v1
        if (-96 >= 0) {
            v19 = int32(v17)
            fmt.Println(v17, fi18, v19)
        } else {
            switch (^(fi18)) {
            case 32768:
                for fi20 := 0; fi20 < 19; fi20++ {
                    fmt.Println(v17)
                    v17 = (^(v17))
                    v21 := GenS1{F0: (fi18 | (fi20 ^ (fi18 - -503))), F1: ""}
                    v17 += v17
                    var v22 uint16 = 32767
                    fmt.Println(v1)
                    _ = fi20
                    _ = v21
                    _ = v22
                }
            }
            v3 = append(v3, v17)
        }
        _ = fi18
        _ = v19
    }
    var v23 uint = 18446744073709551615
    if len(v3) > 0 {
        v3[0] = (^(v17))
    }
    fmt.Println(v17)
    if false {
        v3 = append(v3, (v17 % (((v17 % (((^(v17))) | 1))) | 1)))
    } else {
        v23 += v23
        v23 = 2
    }
    for fi24 := 0; fi24 < 7; fi24++ {
        if ((357 > 1) && true) {
            if true {
                fmt.Println(v23, v1, v2)
                v1 += v1
                var v25 uint = v23
                _ = v25
            } else {
                if true {
                    if true {
                        fmt.Println(v1, fi24)
                        v26 := GenS3{F0: v2, F1: (false && ((!true) && true)), F2: ((^(v2)) <= 0)}
                        v27 := GenS1{F0: (fi24 << (((v23 | 65535)) & 63)), F1: "\n\t"}
                        _ = v26
                        _ = v27
                    }
                } else {
                    v28 := GenS3{F0: v2, F1: (false || (fi24 >= fi24)), F2: (!((("\n\t" + "a") + (("hello" + "hello") + "foo bar")) > "\n\t"))}
                    v3 = append(v3, v17)
                    for fi29 := 0; fi29 < 1; fi29++ {
                        var v30 uint32 = 0
                        var v31 int8 = 14
                        switch (-(fi29)) {
                        case 892:
                            var v32 float64 = -2.25
                            _ = v32
                        case 4294967295:
                            if len(v3) > 0 {
                                v3[0] = v17
                            }
                        case 2147483647:
                            if (fi24 != fi29) {
                                v33 := GenS4{F0: 128, F1: 128, F2: 1}
                                v34 := GenS5{F0: 65535, F1: false, F2: uint64(v23)}
                                _ = v33
                                _ = v34
                            } else {
                                v23 -= v23
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                var v35 int16 = 32767
                                if true {
                                    var v36 int32 = v1
                                    switch "foo bar" {
                                    case "\n\t":
                                        if (((v17 - (v17 ^ 4294967295)) >= (v30 & v30)) && ((!(!(0.0 != (1.5 + 3.14159)))) || ((v30 % (((v17 - v30)) | 1)) < (^(v17))))) {
                                            var v37 float64 = 0.1
                                            _ = v37
                                        } else {
                                            v36 -= 404
                                        }
                                    case "":
                                        switch ("hello" + "hello") {
                                        case "a":
                                            if (int(v23) == fi24) {
                                                var v38 uint = v23
                                                fmt.Println(v38)
                                                if len(v3) >= 2 {
                                                    v3 = v3[1:len(v3)]
                                                }
                                                v39 := GenS2{F0: (v1 % ((-28) | 1)), F1: 649, F2: true}
                                                _ = v38
                                                _ = v39
                                            }
                                        case "":
                                            v17 = (v30 / (((v30 & v17)) | 1))
                                        case "foo bar":
                                            if true {
                                                v30 = v30
                                                if false {
                                                    switch (fi24 % ((0) | 1)) {
                                                    case 1:
                                                        fmt.Println(fi24, fi29, v2)
                                                    case 32767:
                                                        v40 := GenS3{F0: v2, F1: ((false && (true && true)) && false), F2: (!false)}
                                                        _ = v40
                                                    }
                                                    v1 = 982
                                                } else {
                                                    switch fi24 {
                                                    case 4294967295:
                                                        if false {
                                                            v41 := GenS1{F0: -222, F1: ("hello" + ("hello" + "\n\t"))}
                                                            v42 := GenS4{F0: fi24, F1: uint8(v30), F2: v1}
                                                            v36 -= (^(v1))
                                                            v1 = v1
                                                            v1 += int32(v30)
                                                            fmt.Println(fi24, v23, v17)
                                                            _ = v41
                                                            _ = v42
                                                        } else {
                                                            switch 42 {
                                                            case 13:
                                                                var v43 int8 = (v31 | v31)
                                                                _ = v43
                                                            }
                                                        }
                                                    default:
                                                        switch "" {
                                                        case "a":
                                                            if ((!true) && false) {
                                                                v44 := GenS1{F0: (fi24 % ((int(v31)) | 1)), F1: "a"}
                                                                if (174 <= 116) {
                                                                    if ((1.5 / (0.1 + 0.1)) > 0.0) {
                                                                        switch ((("hello" + "") + "") + ((("foo bar" + ("\n\t" + ("a" + "a"))) + "foo bar") + ("foo bar" + ("\n\t" + "a")))) {
                                                                        case "hello":
                                                                            v45 := GenS6{F0: "\n\t"}
                                                                            _ = v45
                                                                        case "a":
                                                                            v17 -= 953
                                                                        }
                                                                        var v46 bool = ("a" == ("\n\t" + ("foo bar" + "foo bar")))
                                                                        _ = v46
                                                                    } else {
                                                                        if (1000 == (v35 / (((v35 >> ((v23) & 63))) | 1))) {
                                                                            v47 := GenS7{F0: 9223372036854775807, F1: (v35 ^ (v35 ^ v35))}
                                                                            fmt.Println(v17, fi24, v23)
                                                                            _ = v47
                                                                        } else {
                                                                            v35 += -41
                                                                            v1 = v36
                                                                        }
                                                                        if (!(644 >= 0)) {
                                                                            var v48 uint32 = v17
                                                                            _ = v48
                                                                        } else {
                                                                            fmt.Println(v30, v17)
                                                                            var v49 bool = (255 <= 42)
                                                                            _ = v49
                                                                        }
                                                                    }
                                                                    v50 := GenS8{F0: 0}
                                                                    _ = v50
                                                                } else {
                                                                    if false {
                                                                        var v51 int32 = 1
                                                                        fmt.Println(fi29, v51, v2)
                                                                        var v52 uint = 2
                                                                        var v53 float32 = (1e10 + 0.0)
                                                                        fmt.Println(v53)
                                                                        _ = v51
                                                                        _ = v52
                                                                        _ = v53
                                                                    }
                                                                    switch "foo bar" {
                                                                    case "":
                                                                        if len(v3) >= 2 {
                                                                            v3 = v3[1:len(v3)]
                                                                        }
                                                                    default:
                                                                        if false {
                                                                            var v54 string = (("" + ("" + ("foo bar" + "foo bar"))) + "\n\t")
                                                                            var v55 int32 = (^(v36))
                                                                            var v56 uint8 = 112
                                                                            _ = v54
                                                                            _ = v55
                                                                            _ = v56
                                                                        }
                                                                    }
                                                                    v57 := GenS9{F0: 0, F1: (v2 + v2)}
                                                                    _ = v57
                                                                }
                                                                if false {
                                                                    if len(v3) > 0 {
                                                                        v3[0] = 4294967295
                                                                    }
                                                                } else {
                                                                    if true {
                                                                        if false {
                                                                            var v58 bool = false
                                                                            _ = v58
                                                                        }
                                                                        var v59 string = "hello"
                                                                        _ = v59
                                                                    }
                                                                    v3 = append(v3, v30)
                                                                    v23 += (v23 << ((v23) & 63))
                                                                }
                                                                _ = v44
                                                            }
                                                        case "\n\t":
                                                            if ((0 <= v1) || (!false)) {
                                                                switch "" {
                                                                case "hello":
                                                                    var v60 uint64 = v2
                                                                    _ = v60
                                                                case "a":
                                                                    var v61 bool = true
                                                                    _ = v61
                                                                default:
                                                                    v36 = v36
                                                                }
                                                                var v62 int64 = 0
                                                                _ = v62
                                                            } else {
                                                                switch "\n\t" {
                                                                case "hello":
                                                                    var v63 int32 = -1
                                                                    v3 = append(v3, 694)
                                                                    _ = v63
                                                                case "\n\t":
                                                                    switch (fi24 ^ (^(fi24))) {
                                                                    case -9223372036854775808:
                                                                        if ((3.14159 + (0.0 - (-2.25 * (1e10 / (1.5 * 3.14159))))) <= ((1.5 * -2.25) * 0.1)) {
                                                                            var v64 uint = 1
                                                                            v30 -= 65535
                                                                            _ = v64
                                                                        } else {
                                                                            var v65 float64 = float64(3.14159)
                                                                            v66 := GenS10{F0: v2, F1: false}
                                                                            _ = v65
                                                                            _ = v66
                                                                        }
                                                                    case 42:
                                                                        var v67 int = (fi29 | (fi24 & 32768))
                                                                        _ = v67
                                                                    }
                                                                case "a":
                                                                    var v68 uint64 = 716
                                                                    _ = v68
                                                                default:
                                                                    if ((v23 & 18446744073709551615) == v23) {
                                                                        var v69 int64 = 969
                                                                        _ = len(v3)
                                                                        var v70 uint16 = 42
                                                                        _ = v69
                                                                        _ = v70
                                                                    }
                                                                }
                                                                fmt.Println(v30, fi29, fi24)
                                                                var v71 int16 = v35
                                                                v35 = (v35 * (v71 ^ -886))
                                                                v23 = 42
                                                                _ = v71
                                                            }
                                                            if len(v3) > 0 {
                                                                v3[0] = 937
                                                            }
                                                            if false {
                                                                fmt.Println(v2)
                                                                v72 := GenS8{F0: 357}
                                                                v1 = (v36 - (v36 << (((^(v23))) & 63)))
                                                                _ = v72
                                                            } else {
                                                                v31 -= (-(v31))
                                                                fmt.Println(v23, v31)
                                                                if false {
                                                                    v3 = append(v3, v30)
                                                                    if len(v3) >= 2 {
                                                                        v3 = v3[1:len(v3)]
                                                                    }
                                                                } else {
                                                                    var v73 uint = (v23 | v23)
                                                                    switch ("" + "foo bar") {
                                                                    case "hello":
                                                                        var v74 uint64 = 42
                                                                        var v75 uint32 = v30
                                                                        _ = v74
                                                                        _ = v75
                                                                    case "a":
                                                                        var v76 int8 = (v31 + v31)
                                                                        _ = v76
                                                                    }
                                                                    _ = v73
                                                                }
                                                            }
                                                        }
                                                        switch fi29 {
                                                        case -797:
                                                            var v77 bool = true
                                                            _ = v77
                                                        default:
                                                            var v78 int8 = v31
                                                            _ = v78
                                                        }
                                                        var v79 float64 = 1e10
                                                        _ = v79
                                                    }
                                                    var v80 int64 = 0
                                                    var v81 uint32 = (v30 & 4294967295)
                                                    _ = v80
                                                    _ = v81
                                                }
                                                v1 = -291
                                                switch -2147483648 {
                                                case 0:
                                                    v3 = append(v3, (^(v17)))
                                                case 128:
                                                    if false {
                                                        if true {
                                                            var v82 int32 = v1
                                                            var v83 int8 = (v31 << ((v23) & 63))
                                                            _ = v82
                                                            _ = v83
                                                        } else {
                                                            if len(v3) > 0 {
                                                                v3[0] = (v30 / (((^(v30))) | 1))
                                                            }
                                                            switch (fi24 / ((333) | 1)) {
                                                            case 32768:
                                                                var v84 int = (fi24 >> ((v23) & 63))
                                                                _ = v84
                                                            case -1:
                                                                v35 = (v35 / ((v35) | 1))
                                                            case 627:
                                                                v2 = 323
                                                                var v85 bool = false
                                                                if len(v3) >= 2 {
                                                                    v3 = v3[1:len(v3)]
                                                                }
                                                                _ = v85
                                                            }
                                                            if (158 >= 0) {
                                                                v86 := GenS11{F0: int(v31), F1: 33}
                                                                var v87 int = 42
                                                                _ = v86
                                                                _ = v87
                                                            } else {
                                                                var v88 uint = (v23 | v23)
                                                                var v89 uint32 = 65535
                                                                switch "a" {
                                                                case "foo bar":
                                                                    switch (fi29 | (fi29 + (fi29 + (fi29 + 127)))) {
                                                                    case 32767:
                                                                        var v90 float64 = 3.14159
                                                                        _ = v90
                                                                    case 0:
                                                                        var v91 string = "hello"
                                                                        _ = v91
                                                                    case 9223372036854775807:
                                                                        v3 = append(v3, v17)
                                                                    default:
                                                                        v2 += (^(v2))
                                                                    }
                                                                case "a":
                                                                    v92 := GenS12{F0: ("" <= ((("hello" + "hello") + "foo bar") + (("a" + "a") + "")))}
                                                                    _ = v92
                                                                case "\n\t":
                                                                    switch fi24 {
                                                                    case 0:
                                                                        fmt.Println(v89, v89, v1)
                                                                    case -772:
                                                                        _ = len(v3)
                                                                    default:
                                                                        switch ((("\n\t" + "\n\t") + (("" + "foo bar") + ((("foo bar" + "\n\t") + "foo bar") + "hello"))) + "foo bar") {
                                                                        case "\n\t":
                                                                            var v93 string = ("a" + "foo bar")
                                                                            _ = v93
                                                                        case "foo bar":
                                                                            v94 := GenS12{F0: false}
                                                                            _ = v94
                                                                        case "":
                                                                            if len(v3) > 0 {
                                                                                v3[0] = (v89 % (((v17 + v17)) | 1))
                                                                            }
                                                                        default:
                                                                            var v95 uint8 = 1
                                                                            _ = v95
                                                                        }
                                                                    }
                                                                default:
                                                                    v89 -= (v30 ^ 2147483648)
                                                                }
                                                                _ = v88
                                                                _ = v89
                                                            }
                                                        }
                                                    }
                                                case -245:
                                                    var v96 uint32 = v30
                                                    _ = v96
                                                }
                                                v1 += -136
                                                v3 = append(v3, v30)
                                                v97 := GenS7{F0: 42, F1: v35}
                                                _ = v97
                                            }
                                        }
                                    }
                                    _ = v36
                                }
                                switch fi24 {
                                case 4294967295:
                                    fmt.Println(v35, fi29, fi24)
                                case 32767:
                                    v30 = 65535
                                default:
                                    v98 := GenS11{F0: (fi24 | int(v35)), F1: 118}
                                    _ = v98
                                }
                                _ = v35
                            }
                        default:
                            v99 := GenS5{F0: 919, F1: (!(-424 < -1)), F2: v2}
                            _ = v99
                        }
                        v31 = -1
                        if (v23 <= (v23 * (v23 % ((v23) | 1)))) {
                            var v100 int16 = -509
                            if (v2 != 4294967295) {
                                if true {
                                    fmt.Println(v100)
                                    fmt.Println(v30, fi24, v1)
                                }
                            } else {
                                if len(v3) > 0 {
                                    v3[0] = v17
                                }
                                var v101 uint64 = v2
                                _ = v101
                            }
                            _ = v100
                        } else {
                            v3 = append(v3, v30)
                            v2 = (v2 ^ 602)
                            v102 := GenS5{F0: 42, F1: true, F2: 558}
                            _ = v102
                        }
                        v2 -= v2
                        _ = fi29
                        _ = v30
                        _ = v31
                    }
                    _ = v28
                }
                v3 = append(v3, v17)
                v3 = append(v3, v17)
            }
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        } else {
            fmt.Println(v17)
        }
        _ = fi24
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v17
    _ = v2
    _ = v23
    _ = v3
}
