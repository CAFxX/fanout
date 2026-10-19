package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int64; F2 string }
type GenS2 struct { F0 uint64; F1 string; F2 uint64 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint; F1 int64; F2 bool }
type GenS5 struct { F0 int32; F1 uint; F2 uint8 }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 string; F1 string; F2 int32 }
type GenS8 struct { F0 int32; F1 int8 }
type GenS9 struct { F0 int64 }
type GenS10 struct { F0 int64 }
type GenS11 struct { F0 int8 }
type GenS12 struct { F0 string; F1 string }
type GenS13 struct { F0 uint }
type GenS14 struct { F0 string; F1 int64; F2 uint32 }
type GenS15 struct { F0 int; F1 uint16 }
type GenS16 struct { F0 uint64; F1 uint32 }
type GenS17 struct { F0 int; F1 int; F2 int16 }
type GenS18 struct { F0 uint16 }
type GenS19 struct { F0 uint32 }
type GenS20 struct { F0 uint64; F1 uint8 }
type GenS21 struct { F0 uint16; F1 int8; F2 int }
type GenS22 struct { F0 int32; F1 string }

func main() {
    var v1 int64 = 1
    var v2 uint8 = 127
    var v3 float32 = 3.14159
    var v4 int16 = 1
    var v5 string = (("\n\t" + "foo bar") + ("" + (("a" + "hello") + "foo bar")))
    var v6 int64 = (-(v1))
    switch "foo bar" {
    case "a":
        var v7 uint64 = 18446744073709551615
        _ = v7
    default:
        fmt.Println(v1, v5)
    }
    v8 := GenS1{F0: (v1 & v6), F1: (-(v6)), F2: ""}
    v5 = "hello"
    if false {
        fmt.Println(v3, v2)
        v6 = v6
    }
    v1 = v6
    if false {
        var v9 uint32 = 4294967295
        _ = v9
    } else {
        if (false && true) {
            v10 := GenS2{F0: 1, F1: ("hello" + (("" + v5) + ("\n\t" + v5))), F2: 1}
            _ = v10
        } else {
            switch (v5 + "\n\t") {
            case "\n\t":
                for fi11 := 0; fi11 < 11; fi11++ {
                    switch "foo bar" {
                    case "hello":
                        if false {
                            fmt.Println(v2)
                            for fi12 := 0; fi12 < 5; fi12++ {
                                if true {
                                    var v13 bool = true
                                    switch v5 {
                                    case "":
                                        v3 -= v3
                                    case "foo bar":
                                        v13 = v13
                                    case "\n\t":
                                        fmt.Println(fi12)
                                    }
                                    var v14 int8 = -1
                                    _ = v13
                                    _ = v14
                                }
                                v15 := GenS3{F0: false}
                                if (false && true) {
                                    v16 := GenS4{F0: 0, F1: v1, F2: (!false)}
                                    var v17 string = v5
                                    switch 128 {
                                    case 42:
                                        if (!(42 >= 2)) {
                                            var v18 string = ("\n\t" + "")
                                            _ = v18
                                        }
                                    case 32768:
                                        if (v17 >= "\n\t") {
                                            var v19 bool = (-31 > 100)
                                            var v20 bool = v19
                                            switch (fi12 / ((fi11) | 1)) {
                                            case 1:
                                                var v21 uint32 = 2
                                                _ = v21
                                            case 128:
                                                if (!(!((v1 ^ (v1 + (v6 * -247))) < v1))) {
                                                    if v20 {
                                                        v22 := GenS3{F0: (1 <= -1)}
                                                        switch fi12 {
                                                        case 1:
                                                            switch fi11 {
                                                            case 639:
                                                                switch v5 {
                                                                case "":
                                                                    switch ("hello" + v5) {
                                                                    case "a":
                                                                        switch (^(fi11)) {
                                                                        case 4294967295:
                                                                            if ("a" == "") {
                                                                                if (0 < (-(fi12))) {
                                                                                    var v23 uint32 = 2147483647
                                                                                    if true {
                                                                                        var v24 int64 = (v6 ^ (v6 - v6))
                                                                                        if (962 > 255) {
                                                                                            v19 = (v20 && true)
                                                                                            v5 = ("" + ((v17 + v17) + v5))
                                                                                            if (42 > 42) {
                                                                                                switch 127 {
                                                                                                case 4294967295:
                                                                                                    if v20 {
                                                                                                        switch v5 {
                                                                                                        case "\n\t":
                                                                                                            fmt.Println(v6, v23, v6)
                                                                                                        case "a":
                                                                                                            switch int(v2) {
                                                                                                            case 416:
                                                                                                                switch v17 {
                                                                                                                case "foo bar":
                                                                                                                    v5 = v5
                                                                                                                case "\n\t":
                                                                                                                    v1 -= 226
                                                                                                                    v25 := GenS5{F0: 1, F1: 0, F2: 128}
                                                                                                                    _ = v25
                                                                                                                case "":
                                                                                                                    switch int(v4) {
                                                                                                                    case 42:
                                                                                                                        v20 = ((!true) || (42 > 42))
                                                                                                                    default:
                                                                                                                        v26 := []int{(fi12 & (fi11 << ((65535) & 63))), (fi11 + (fi12 + fi11))}
                                                                                                                        _ = v26
                                                                                                                    }
                                                                                                                }
                                                                                                            case -2147483648:
                                                                                                                v5 = v17
                                                                                                            }
                                                                                                        case "":
                                                                                                            if false {
                                                                                                                var v27 int = fi12
                                                                                                                _ = v27
                                                                                                            }
                                                                                                        default:
                                                                                                            v28 := GenS1{F0: v6, F1: (v6 ^ (-(v1))), F2: (v5 + v5)}
                                                                                                            _ = v28
                                                                                                        }
                                                                                                    }
                                                                                                case 246:
                                                                                                    var v29 int8 = 0
                                                                                                    _ = v29
                                                                                                }
                                                                                                var v30 uint8 = 255
                                                                                                if (!v19) {
                                                                                                    v31 := GenS3{F0: (!((v19 || (1 >= -1)) && v20))}
                                                                                                    switch fi11 {
                                                                                                    case 2147483647:
                                                                                                        var v32 int8 = 86
                                                                                                        _ = v32
                                                                                                    case 32768:
                                                                                                        v33 := []int{(fi12 / ((-2147483648) | 1)), (fi12 >> ((2) & 63)), fi12}
                                                                                                        _ = v33
                                                                                                    default:
                                                                                                        var v34 int64 = (-(v6))
                                                                                                        fmt.Println(v5, fi11, v30)
                                                                                                        v35 := GenS5{F0: -2147483648, F1: 255, F2: v2}
                                                                                                        _ = v34
                                                                                                        _ = v35
                                                                                                    }
                                                                                                    _ = v31
                                                                                                } else {
                                                                                                    v36 := GenS4{F0: 257, F1: (v1 | v6), F2: (false || (((100 == 100) || v19) && (!v19)))}
                                                                                                    switch (fi12 + 42) {
                                                                                                    case 0:
                                                                                                        if (v3 > ((3.14159 - 1.5) * (3.14159 - 3.14159))) {
                                                                                                            if v20 {
                                                                                                                v20 = ((((((v5 + v17) + v5) + "hello") + v5) + "a") >= v5)
                                                                                                                v3 += (v3 / ((-2.25 + v3) * v3))
                                                                                                                v4 += (v4 / (((v4 % (((v4 << ((89) & 63))) | 1))) | 1))
                                                                                                                v37 := GenS3{F0: (2147483647 == 646)}
                                                                                                                _ = v37
                                                                                                            } else {
                                                                                                                v38 := []int{(fi11 & fi11)}
                                                                                                                v39 := GenS6{F0: 42}
                                                                                                                var v40 uint64 = 2
                                                                                                                _ = v38
                                                                                                                _ = v39
                                                                                                                _ = v40
                                                                                                            }
                                                                                                            var v41 int64 = v1
                                                                                                            if false {
                                                                                                                var v42 float32 = ((v3 + (1.5 * v3)) * v3)
                                                                                                                _ = v42
                                                                                                            }
                                                                                                            _ = v41
                                                                                                        }
                                                                                                    }
                                                                                                    _ = v36
                                                                                                }
                                                                                                _ = v30
                                                                                            } else {
                                                                                                var v43 bool = (!v20)
                                                                                                if ((v2 & v2) == (v2 + (v2 & 128))) {
                                                                                                    v44 := []int{(fi11 & (-(fi11))), (^(fi12))}
                                                                                                    var v45 string = "\n\t"
                                                                                                    _ = v44
                                                                                                    _ = v45
                                                                                                }
                                                                                                _ = v43
                                                                                            }
                                                                                        }
                                                                                        var v46 int64 = (v24 << ((2) & 63))
                                                                                        _ = v24
                                                                                        _ = v46
                                                                                    }
                                                                                    v6 += v1
                                                                                    _ = v23
                                                                                }
                                                                            }
                                                                        case 686:
                                                                            var v47 uint64 = 1
                                                                            _ = v47
                                                                        default:
                                                                            fmt.Println(v4, fi11)
                                                                        }
                                                                    default:
                                                                        var v48 uint16 = 2
                                                                        _ = v48
                                                                    }
                                                                default:
                                                                    fmt.Println(fi11, v5)
                                                                }
                                                            case 9223372036854775807:
                                                                var v49 uint16 = 0
                                                                _ = v49
                                                            }
                                                        case 127:
                                                            var v50 bool = (!(!v20))
                                                            var v51 int32 = 42
                                                            if ((fi11 / ((32768) | 1)) > 42) {
                                                                if true {
                                                                    v20 = (!((v5 == "\n\t") && false))
                                                                    v2 = v2
                                                                    v51 -= (v51 + (v51 ^ -389))
                                                                } else {
                                                                    if (2 <= 966) {
                                                                        var v52 int8 = -128
                                                                        _ = v52
                                                                    }
                                                                    if (v51 <= v51) {
                                                                        var v53 uint64 = 2
                                                                        v54 := []int{(fi11 + -490)}
                                                                        _ = v53
                                                                        _ = v54
                                                                    } else {
                                                                        v55 := GenS7{F0: v5, F1: v5, F2: v51}
                                                                        v56 := []int{fi12}
                                                                        _ = v55
                                                                        _ = v56
                                                                    }
                                                                }
                                                            } else {
                                                                fmt.Println(v17)
                                                                var v57 int16 = (v4 - 32767)
                                                                _ = v57
                                                            }
                                                            _ = v50
                                                            _ = v51
                                                        default:
                                                            v58 := GenS8{F0: 1, F1: -78}
                                                            _ = v58
                                                        }
                                                        var v59 int8 = 0
                                                        v60 := GenS9{F0: 1}
                                                        var v61 int32 = 1
                                                        _ = v22
                                                        _ = v59
                                                        _ = v60
                                                        _ = v61
                                                    }
                                                    switch fi11 {
                                                    case -9223372036854775808:
                                                        var v62 int = (fi12 - 32767)
                                                        _ = v62
                                                    case -937:
                                                        switch "foo bar" {
                                                        case "a":
                                                            v3 = (((((v3 / v3) + float32(1.5)) + ((float32(float64((0.1 * ((v3 * 0.0) / float32(1.5))))) * v3) + float32(0.0))) + v3) * v3)
                                                        case "\n\t":
                                                            v63 := []int{32767, (fi11 ^ fi12), fi12, (fi12 + (-(fi11)))}
                                                            _ = v63
                                                        case "foo bar":
                                                            if v20 {
                                                                v64 := []int{fi12, (fi12 / ((127) | 1)), (^(fi11))}
                                                                var v65 string = "foo bar"
                                                                v20 = v20
                                                                var v66 float32 = float32((1e10 - 3.14159))
                                                                _ = len(v64)
                                                                var v67 string = v17
                                                                v68 := GenS4{F0: 0, F1: v1, F2: v20}
                                                                _ = v64
                                                                _ = v65
                                                                _ = v66
                                                                _ = v67
                                                                _ = v68
                                                            } else {
                                                                var v69 int8 = 1
                                                                _ = v69
                                                            }
                                                        default:
                                                            fmt.Println(v19, v6, v3)
                                                        }
                                                        v70 := GenS9{F0: 608}
                                                        v5 = (v17 + ("\n\t" + ("\n\t" + "")))
                                                        _ = v70
                                                    case 44:
                                                        v71 := []int{(fi11 - (fi12 >> ((255) & 63))), 892, int(v6)}
                                                        _ = v71
                                                    }
                                                    v72 := []int{(-(fi12))}
                                                    switch (fi12 >> ((65535) & 63)) {
                                                    case -581:
                                                        switch (fi12 & fi11) {
                                                        case 321:
                                                            v3 += 1e10
                                                        case 127:
                                                            switch -1 {
                                                            case 0:
                                                                v73 := GenS10{F0: 1}
                                                                _ = v73
                                                            default:
                                                                _ = len(v72)
                                                            }
                                                        default:
                                                            switch v5 {
                                                            case "hello":
                                                                v20 = true
                                                            default:
                                                                if true {
                                                                    var v74 bool = v20
                                                                    v72 = append(v72, fi11)
                                                                    if (42 > 32767) {
                                                                        if v74 {
                                                                            var v75 uint32 = 788
                                                                            fmt.Println(v17, v19, v75)
                                                                            _ = v75
                                                                        }
                                                                        v72 = append(v72, fi11)
                                                                        switch (v5 + v5) {
                                                                        case "\n\t":
                                                                            var v76 int64 = -9223372036854775808
                                                                            _ = v76
                                                                        }
                                                                    } else {
                                                                        fmt.Println(v20, v2, fi11)
                                                                    }
                                                                    _ = v74
                                                                }
                                                                v77 := GenS11{F0: -128}
                                                                var v78 int8 = -94
                                                                v5 = "a"
                                                                _ = v77
                                                                _ = v78
                                                            }
                                                        }
                                                    case 32768:
                                                        var v79 int8 = -1
                                                        _ = v79
                                                    default:
                                                        if v19 {
                                                            switch v5 {
                                                            case "\n\t":
                                                                var v80 uint32 = 991
                                                                _ = v80
                                                            default:
                                                                var v81 int32 = 42
                                                                _ = v81
                                                            }
                                                            fmt.Println(fi11, v20, v1)
                                                            var v82 uint32 = 2
                                                            _ = v82
                                                        }
                                                    }
                                                    v72 = append(v72, fi11)
                                                    _ = v72
                                                } else {
                                                    v2 = 255
                                                    v83 := []int{}
                                                    switch v5 {
                                                    case "hello":
                                                        var v84 int = (fi12 % (((fi11 % ((128) | 1))) | 1))
                                                        _ = v84
                                                    case "foo bar":
                                                        var v85 int64 = (v1 % ((v1) | 1))
                                                        _ = v85
                                                    case "":
                                                        var v86 float64 = 1e10
                                                        _ = v86
                                                    }
                                                    if v20 {
                                                        var v87 int8 = -31
                                                        v88 := GenS8{F0: 1, F1: (v87 + int8(fi11))}
                                                        _ = v87
                                                        _ = v88
                                                    } else {
                                                        _ = len(v83)
                                                    }
                                                    v6 -= (v6 & v6)
                                                    var v89 uint = 426
                                                    _ = v83
                                                    _ = v89
                                                }
                                            }
                                            _ = v19
                                            _ = v20
                                        } else {
                                            v90 := []int{(^(fi12))}
                                            switch fi12 {
                                            case -2147483648:
                                                if len(v90) > 0 {
                                                    v90[0] = (fi11 + (fi11 & fi11))
                                                }
                                            case 9223372036854775807:
                                                switch v17 {
                                                case "a":
                                                    v5 = "hello"
                                                case "\n\t":
                                                    if (2 <= 255) {
                                                        if (42 <= 18446744073709551615) {
                                                            if true {
                                                                fmt.Println(v3)
                                                                var v91 int64 = v6
                                                                _ = v91
                                                            } else {
                                                                var v92 int8 = -1
                                                                if len(v90) > 0 {
                                                                    v90[0] = fi12
                                                                }
                                                                v93 := GenS12{F0: "\n\t", F1: "hello"}
                                                                _ = v92
                                                                _ = v93
                                                            }
                                                        } else {
                                                            if len(v90) >= 2 {
                                                                v90 = v90[1:len(v90)]
                                                            }
                                                            if true {
                                                                if (true || true) {
                                                                    v94 := GenS13{F0: 65535}
                                                                    if (!true) {
                                                                        if len(v90) >= 2 {
                                                                            v90 = v90[1:len(v90)]
                                                                        }
                                                                        var v95 bool = ((^(v4)) < (v4 >> ((927) & 63)))
                                                                        _ = v95
                                                                    } else {
                                                                        switch v5 {
                                                                        case "hello":
                                                                            _ = len(v90)
                                                                        case "":
                                                                            if (42 > (v4 % (((^(v4))) | 1))) {
                                                                                var v96 int32 = 42
                                                                                _ = v96
                                                                            } else {
                                                                                v97 := GenS14{F0: "hello", F1: v1, F2: 2}
                                                                                var v98 uint8 = 177
                                                                                if false {
                                                                                    fmt.Println(v98)
                                                                                } else {
                                                                                    v3 -= v3
                                                                                }
                                                                                _ = v97
                                                                                _ = v98
                                                                            }
                                                                        case "\n\t":
                                                                            v99 := GenS4{F0: 65535, F1: (v1 << ((861) & 63)), F2: false}
                                                                            _ = v99
                                                                        }
                                                                        v100 := GenS15{F0: fi12, F1: 772}
                                                                        if true {
                                                                            v5 = v5
                                                                            if true {
                                                                                var v101 int32 = -2147483648
                                                                                switch ((v5 + v5) + v5) {
                                                                                case "\n\t":
                                                                                    if len(v90) >= 2 {
                                                                                        v90 = v90[1:len(v90)]
                                                                                    }
                                                                                case "":
                                                                                    v102 := GenS7{F0: "a", F1: ((v5 + ("hello" + v17)) + (v17 + "hello")), F2: v101}
                                                                                    _ = v102
                                                                                case "a":
                                                                                    var v103 uint = 919
                                                                                    _ = v103
                                                                                default:
                                                                                    if (int32(v4) <= v101) {
                                                                                        v6 = v1
                                                                                    } else {
                                                                                        v3 -= v3
                                                                                        if (false || (1.5 != 0.1)) {
                                                                                            v104 := GenS16{F0: 2, F1: 2}
                                                                                            _ = v104
                                                                                        }
                                                                                    }
                                                                                }
                                                                                _ = v101
                                                                            }
                                                                            var v105 string = ""
                                                                            v6 += v6
                                                                            _ = v105
                                                                        }
                                                                        _ = v100
                                                                    }
                                                                    fmt.Println(v2, fi12)
                                                                    _ = v94
                                                                } else {
                                                                    v2 += (v2 | (v2 >> ((42) & 63)))
                                                                }
                                                                if (0 <= 0) {
                                                                    v1 = (v6 * 1)
                                                                    var v106 int16 = v4
                                                                    _ = v106
                                                                }
                                                                var v107 int8 = -55
                                                                _ = v107
                                                            } else {
                                                                v2 = v2
                                                            }
                                                            v108 := GenS17{F0: 4294967295, F1: -505, F2: (v4 / ((v4) | 1))}
                                                            _ = v108
                                                        }
                                                    } else {
                                                        fmt.Println(v5, v5)
                                                    }
                                                    v4 += v4
                                                    if len(v90) > 0 {
                                                        v90[0] = fi12
                                                    }
                                                }
                                            }
                                            if (false || false) {
                                                var v109 uint = 65535
                                                v90 = append(v90, (fi11 + (fi12 | fi12)))
                                                fmt.Println(v2)
                                                _ = v109
                                            } else {
                                                v6 = v6
                                                var v110 uint16 = 0
                                                _ = v110
                                            }
                                            v2 -= v2
                                            _ = v90
                                        }
                                    case 32767:
                                        var v111 uint16 = 994
                                        if ((-2.25 - float64(3.14159)) < 1.5) {
                                            var v112 int32 = -1
                                            _ = v112
                                        }
                                        v113 := []int{6}
                                        _ = v111
                                        _ = v113
                                    }
                                    _ = v16
                                    _ = v17
                                } else {
                                    switch (v5 + ("hello" + "a")) {
                                    case "hello":
                                        var v114 uint64 = 226
                                        _ = v114
                                    case "\n\t":
                                        fmt.Println(v5)
                                    case "a":
                                        var v115 int32 = 0
                                        var v116 uint16 = 2
                                        v117 := []int{fi11}
                                        _ = v115
                                        _ = v116
                                        _ = v117
                                    }
                                }
                                var v118 bool = ((-777 >= (v1 % ((-31) | 1))) || true)
                                if true {
                                    if v118 {
                                        v119 := GenS18{F0: 765}
                                        switch (v5 + ("foo bar" + v5)) {
                                        case "foo bar":
                                            v120 := []int{int(v2)}
                                            _ = v120
                                        case "a":
                                            var v121 bool = (uint32(v2) <= 2)
                                            _ = v121
                                        case "\n\t":
                                            fmt.Println(v4, v3, v4)
                                        }
                                        if v118 {
                                            if v118 {
                                                v122 := []int{}
                                                fmt.Println(v2, v3)
                                                _ = v122
                                            } else {
                                                var v123 uint = 2
                                                v124 := GenS19{F0: 1}
                                                switch fi12 {
                                                case 128:
                                                    v125 := GenS20{F0: 42, F1: 2}
                                                    _ = v125
                                                }
                                                _ = v123
                                                _ = v124
                                            }
                                            v126 := GenS21{F0: 412, F1: -1, F2: 32768}
                                            if (false || v118) {
                                                v127 := []int{}
                                                fmt.Println(v2, v4, v3)
                                                var v128 uint32 = 0
                                                _ = v127
                                                _ = v128
                                            } else {
                                                fmt.Println(v1)
                                                var v129 uint8 = 127
                                                _ = v129
                                            }
                                            _ = v126
                                        } else {
                                            v130 := GenS12{F0: "hello", F1: v5}
                                            var v131 string = v5
                                            var v132 int = int(v4)
                                            _ = v130
                                            _ = v131
                                            _ = v132
                                        }
                                        _ = v119
                                    } else {
                                        fmt.Println(v118, v118)
                                        switch fi11 {
                                        case 32767:
                                            v133 := GenS13{F0: 2}
                                            v118 = (-94 <= 21)
                                            _ = v133
                                        case 1:
                                            v5 = "foo bar"
                                        }
                                        var v134 uint = 4294967295
                                        _ = v134
                                    }
                                } else {
                                    var v135 int64 = (-(v1))
                                    v136 := []int{0}
                                    v137 := GenS2{F0: 1, F1: v5, F2: 0}
                                    _ = v135
                                    _ = v136
                                    _ = v137
                                }
                                var v138 int16 = v4
                                var v139 uint8 = 255
                                var v140 float32 = (float32(0.0) * 1e10)
                                _ = fi12
                                _ = v118
                                _ = v138
                                _ = v139
                                _ = v140
                                _ = v15
                            }
                        }
                    default:
                        if (-1 <= -682) {
                            for fi141 := 0; fi141 < 9; fi141++ {
                                if false {
                                    v142 := GenS22{F0: 1, F1: (v5 + ("" + "hello"))}
                                    _ = v142
                                } else {
                                    if true {
                                        var v143 int8 = -128
                                        var v144 uint = 2
                                        _ = v143
                                        _ = v144
                                    }
                                    var v145 uint64 = 0
                                    var v146 int16 = (v4 ^ v4)
                                    _ = v145
                                    _ = v146
                                }
                                fmt.Println(v5)
                                v147 := GenS15{F0: (fi141 * fi141), F1: 32767}
                                _ = fi141
                                _ = v147
                            }
                        } else {
                            v4 -= 1000
                            for fi148 := 0; fi148 < 1; fi148++ {
                                var v149 int8 = 127
                                v150 := GenS1{F0: (v1 * v1), F1: v6, F2: (v5 + ("" + (v5 + (v5 + v5))))}
                                switch (fi11 % ((fi148) | 1)) {
                                case -1:
                                    var v151 uint32 = 901
                                    _ = v151
                                case -813:
                                    var v152 int64 = (v6 >> ((4294967295) & 63))
                                    _ = v152
                                case 32767:
                                    var v153 int8 = (v149 & (v149 / (((v149 + (-(v149)))) | 1)))
                                    v154 := GenS17{F0: fi148, F1: -2147483648, F2: 931}
                                    fmt.Println(v149, v5, v3)
                                    _ = v153
                                    _ = v154
                                default:
                                    if false {
                                        fmt.Println(v4, v3)
                                        if false {
                                            v155 := GenS1{F0: 1, F1: (v1 * v1), F2: ""}
                                            _ = v155
                                        } else {
                                            if (42 > 1) {
                                                v5 = ""
                                                fmt.Println(fi11)
                                                var v156 int64 = v6
                                                _ = v156
                                            } else {
                                                v6 -= v1
                                            }
                                        }
                                    }
                                }
                                _ = fi148
                                _ = v149
                                _ = v150
                            }
                            v157 := GenS2{F0: 502, F1: ("" + v5), F2: 531}
                            _ = v157
                        }
                    }
                    _ = fi11
                }
            default:
                v158 := []int{-9223372036854775808, int(v6), 502, 127}
                _ = v158
            }
            var v159 uint64 = 574
            var v160 int64 = v6
            _ = v159
            _ = v160
        }
    }
    v161 := GenS16{F0: 819, F1: 2}
    v3 = (float32(0.1) + v3)
    var v162 uint8 = 1
    v163 := []int{}
    fmt.Println("v1:", v1)
    fmt.Println("v162:", v162)
    fmt.Println("slicelen v163:", len(v163), v163)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v161
    _ = v162
    _ = v163
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
