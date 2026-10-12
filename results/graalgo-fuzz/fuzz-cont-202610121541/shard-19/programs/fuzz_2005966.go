package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 string }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int64; F1 uint64; F2 string }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 int16; F1 int }
type GenS7 struct { F0 string; F1 int32; F2 int8 }
type GenS8 struct { F0 string; F1 uint16; F2 uint32 }
type GenS9 struct { F0 uint64; F1 int8; F2 bool }
type GenS10 struct { F0 int16; F1 uint16; F2 bool }
type GenS11 struct { F0 uint32; F1 uint16 }
type GenS12 struct { F0 int16; F1 int32; F2 int64 }
type GenS13 struct { F0 uint }
type GenS14 struct { F0 int32; F1 int; F2 int }
type GenS15 struct { F0 uint; F1 int32; F2 int64 }
type GenS16 struct { F0 int8; F1 uint }
type GenS17 struct { F0 uint; F1 int }
type GenS18 struct { F0 int; F1 int8 }

func main() {
    var v1 uint8 = 140
    var v2 bool = (true || true)
    v3 := []string{"a", "a", "a", ""}
    v4 := make(map[string]int)
    var v5 uint = 326
    v6 := make(map[uint32]bool)
    var v7 string = ("\n\t" + ("" + "a"))
    fmt.Println(v7)
    if (!(507 < -258)) {
        if v2 {
            var v8 uint8 = (v1 ^ v1)
            if true {
                switch v7 {
                case "foo bar":
                    var v9 int16 = 1
                    _ = v9
                }
                v1 = v8
                if (!(!v2)) {
                    switch 2147483647 {
                    case -9223372036854775808:
                        v10 := GenS1{F0: 1, F1: "a"}
                        _ = v10
                    default:
                        for fi11 := 0; fi11 < 4; fi11++ {
                            v12 := GenS1{F0: 9223372036854775807, F1: v7}
                            _ = fi11
                            _ = v12
                        }
                    }
                } else {
                    var v13 int8 = -128
                    _ = v13
                }
            }
            fmt.Println(v1, v1)
            _ = v8
        } else {
            v2 = v2
            fmt.Println(v5, v2)
        }
        v2 = (1e10 < (-2.25 - -2.25))
    }
    if v2 {
        var v14 int8 = 1
        v3 = append(v3, "a")
        _ = v14
    }
    v15 := GenS2{F0: -1}
    if (!v2) {
        v2 = v2
    } else {
        var v16 int64 = -665
        switch "hello" {
        case "":
            switch 42 {
            case 9223372036854775807:
                switch v7 {
                case "a":
                    if (!true) {
                        for fi17 := 0; fi17 < 6; fi17++ {
                            for fi18 := 0; fi18 < 1; fi18++ {
                                v19 := GenS1{F0: 0, F1: (v7 + v7)}
                                v2 = true
                                v20 := GenS2{F0: (v16 / (((-(v16))) | 1))}
                                fmt.Println(fi17, v2)
                                v5 += v5
                                if v2 {
                                    v21 := GenS2{F0: (v16 >> (((v5 >> ((255) & 63))) & 63))}
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                    _ = v21
                                }
                                fmt.Println(v5, v1)
                                switch (^(fi18)) {
                                case -1:
                                    v22 := v6[553]
                                    _ = v22
                                }
                                v23 := GenS2{F0: v16}
                                _ = fi18
                                _ = v19
                                _ = v20
                                _ = v23
                            }
                            v1 = v1
                            _ = fi17
                        }
                        var v24 int = 128
                        _ = v24
                    } else {
                        if false {
                            v25 := GenS1{F0: (v16 % ((v16) | 1)), F1: (("a" + v7) + v7)}
                            if v2 {
                                v26 := GenS3{F0: (v16 / ((v16) | 1)), F1: uint64(v5), F2: v7}
                                _ = v26
                            }
                            var v27 uint8 = v1
                            _ = v25
                            _ = v27
                        } else {
                            v28 := GenS4{F0: -754}
                            var v29 uint8 = v1
                            _ = len(v3)
                            _ = v28
                            _ = v29
                        }
                        v7 = v7
                    }
                default:
                    var v30 int8 = -98
                    _ = v30
                }
            case 0:
                v31 := GenS5{F0: false}
                _ = v31
            case 4294967295:
                if true {
                    v1 = v1
                    v2 = (1e10 > -2.25)
                    if ((!(v2 || true)) || (((v1 | (v1 % (((v1 & (^(v1)))) | 1))) > 71) && true)) {
                        if v2 {
                            v2 = (!true)
                            switch "a" {
                            case "hello":
                                v32 := GenS6{F0: 1000, F1: 9223372036854775807}
                                _ = v32
                            default:
                                var v33 float32 = 1.5
                                _ = v33
                            }
                        } else {
                            v34 := GenS7{F0: ("foo bar" + v7), F1: 2147483647, F2: -128}
                            _ = v34
                        }
                        if len(v3) > 0 {
                            v3[0] = "\n\t"
                        }
                    } else {
                        v1 -= (v1 + v1)
                        if true {
                            if (!v2) {
                                for fi35 := 0; fi35 < 2; fi35++ {
                                    if true {
                                        var v36 float32 = (0.0 + -2.25)
                                        for fi37 := 0; fi37 < 20; fi37++ {
                                            v16 = (v16 - v16)
                                            if v2 {
                                                if v2 {
                                                    v38 := GenS8{F0: v7, F1: 130, F2: 58}
                                                    if len(v3) > 0 {
                                                        v3[0] = (v7 + (v7 + "\n\t"))
                                                    }
                                                    _ = v38
                                                } else {
                                                    switch v7 {
                                                    case "foo bar":
                                                        var v39 uint8 = 1
                                                        _ = v39
                                                    case "a":
                                                        if v2 {
                                                            _ = len(v4)
                                                            if ((v1 % (((v1 >> ((18446744073709551615) & 63))) | 1)) > v1) {
                                                                var v40 uint8 = v1
                                                                fmt.Println(v1)
                                                                if ((fi35 <= int(v16)) || false) {
                                                                    delete(v4, "hello")
                                                                    v2 = true
                                                                } else {
                                                                    fmt.Println(v2)
                                                                }
                                                                v5 += v5
                                                                _ = v40
                                                            } else {
                                                                v7 = (((v7 + "hello") + ((((v7 + v7) + "hello") + ((v7 + v7) + (v7 + (v7 + v7)))) + ("hello" + v7))) + v7)
                                                            }
                                                        }
                                                    case "\n\t":
                                                        var v41 int64 = -9223372036854775808
                                                        _ = v41
                                                    default:
                                                        var v42 int32 = -2147483648
                                                        if v2 {
                                                            v36 += (0.0 * 1e10)
                                                            v43 := GenS8{F0: "\n\t", F1: 1, F2: 2147483648}
                                                            _ = v43
                                                        } else {
                                                            v36 -= (v36 / (1.5 + float32((3.14159 - 3.14159))))
                                                            if (v2 || (v2 || (!false))) {
                                                                v3 = append(v3, v7)
                                                                v7 = v7
                                                                switch ("" + "hello") {
                                                                case "foo bar":
                                                                    v44 := GenS4{F0: (-(v16))}
                                                                    _ = v44
                                                                case "\n\t":
                                                                    v7 = v7
                                                                case "hello":
                                                                    switch 128 {
                                                                    case 42:
                                                                        switch (fi35 >> (((v5 ^ (^(v5)))) & 63)) {
                                                                        case 42:
                                                                            var v45 uint = (v5 + 255)
                                                                            _ = v45
                                                                        default:
                                                                            v3 = append(v3, "\n\t")
                                                                        }
                                                                    }
                                                                default:
                                                                    if ((v1 >> ((uint(v1)) & 63)) > v1) {
                                                                        _ = len(v6)
                                                                    }
                                                                }
                                                            } else {
                                                                v36 = v36
                                                            }
                                                            if true {
                                                                v7 = "a"
                                                                switch 4294967295 {
                                                                case 0:
                                                                    v1 = v1
                                                                case -1:
                                                                    v46 := GenS8{F0: ("\n\t" + "\n\t"), F1: 1, F2: 65535}
                                                                    _ = v46
                                                                default:
                                                                    var v47 string = ((((v7 + v7) + ("\n\t" + "hello")) + (("a" + v7) + "hello")) + "")
                                                                    _ = v47
                                                                }
                                                            }
                                                        }
                                                        delete(v6, 2147483648)
                                                        _ = v42
                                                    }
                                                    var v48 int32 = 0
                                                    _ = v48
                                                }
                                                fmt.Println(fi37, v16, v36)
                                            } else {
                                                v49 := GenS9{F0: 140, F1: -125, F2: (2147483648 > 4294967295)}
                                                switch fi37 {
                                                case 766:
                                                    _ = len(v3)
                                                }
                                                v50 := GenS10{F0: 1, F1: 32767, F2: (uint(v1) > (v5 | v5))}
                                                v1 -= (v1 & v1)
                                                fmt.Println(v16, fi35)
                                                _ = len(v3)
                                                v51 := GenS11{F0: 2147483647, F1: 255}
                                                var v52 uint32 = 42
                                                _ = v49
                                                _ = v50
                                                _ = v51
                                                _ = v52
                                            }
                                            v53 := GenS10{F0: -1, F1: 0, F2: v2}
                                            _ = fi37
                                            _ = v53
                                        }
                                        _ = v36
                                    } else {
                                        v54 := v4["\n\t"]
                                        if v2 {
                                            v54 -= 588
                                            for fi55 := 0; fi55 < 19; fi55++ {
                                                fmt.Println(v5)
                                                v56 := GenS9{F0: 0, F1: 72, F2: false}
                                                v57 := GenS11{F0: 42, F1: 42}
                                                _ = fi55
                                                _ = v56
                                                _ = v57
                                            }
                                            v2 = false
                                        } else {
                                            v58 := GenS12{F0: 905, F1: int32(v5), F2: v16}
                                            if ((v2 || (!v2)) && false) {
                                                var v59 int = 929
                                                _ = v59
                                            } else {
                                                v16 += v16
                                                var v60 uint16 = 886
                                                _ = v60
                                            }
                                            _ = v58
                                        }
                                        _ = v54
                                    }
                                    if ((false && (false || ((("\n\t" + v7) + "foo bar") <= v7))) && (v16 != v16)) {
                                        v61 := GenS13{F0: 18446744073709551615}
                                        v1 = 231
                                        for fi62 := 0; fi62 < 11; fi62++ {
                                            v1 = (v1 % ((193) | 1))
                                            _ = fi62
                                        }
                                        _ = v61
                                    } else {
                                        switch 144 {
                                        case -360:
                                            v2 = v2
                                        case 128:
                                            fmt.Println(v16)
                                        case -2147483648:
                                            v16 = v16
                                        default:
                                            for fi63 := 0; fi63 < 12; fi63++ {
                                                var v64 uint32 = 377
                                                _ = fi63
                                                _ = v64
                                            }
                                        }
                                    }
                                    for fi65 := 0; fi65 < 14; fi65++ {
                                        v66 := GenS6{F0: 1, F1: (fi35 % ((int(v5)) | 1))}
                                        var v67 uint32 = 65535
                                        _ = fi65
                                        _ = v66
                                        _ = v67
                                    }
                                    _ = fi35
                                }
                                if false {
                                    v1 -= v1
                                    if false {
                                        v2 = ((v7 + (v7 + (v7 + v7))) < (v7 + "hello"))
                                        if v2 {
                                            v2 = v2
                                        }
                                    } else {
                                        v3 = append(v3, v7)
                                        var v68 uint16 = 132
                                        _ = v68
                                    }
                                    var v69 uint64 = 42
                                    _ = v69
                                } else {
                                    switch -1 {
                                    case 1:
                                        if (((v16 * -544) != v16) || ((v1 << ((v5) & 63)) < (v1 / (((^(v1))) | 1)))) {
                                            if (!("" > v7)) {
                                                var v70 int32 = -2147483648
                                                switch 9223372036854775807 {
                                                case 4294967295:
                                                    switch 127 {
                                                    case 128:
                                                        for fi71 := 0; fi71 < 12; fi71++ {
                                                            var v72 float64 = 3.14159
                                                            for fi73 := 0; fi73 < 7; fi73++ {
                                                                if len(v3) > 0 {
                                                                    v3[0] = (((v7 + v7) + ((v7 + v7) + (v7 + (v7 + v7)))) + v7)
                                                                }
                                                                if (!((!v2) && (-88 == 472))) {
                                                                    var v74 float64 = float64((0.1 - 1e10))
                                                                    v2 = (v1 != (v1 - v1))
                                                                    if (v2 || v2) {
                                                                        fmt.Println(fi71, v70, v16)
                                                                        _ = len(v3)
                                                                        fmt.Println(fi71)
                                                                    }
                                                                    _ = v74
                                                                } else {
                                                                    var v75 bool = v2
                                                                    fmt.Println(v72, fi71, v16)
                                                                    _ = len(v3)
                                                                    _ = v75
                                                                }
                                                                var v76 uint = v5
                                                                _ = fi73
                                                                _ = v76
                                                            }
                                                            _ = fi71
                                                            _ = v72
                                                        }
                                                    }
                                                default:
                                                    if (v2 && true) {
                                                        for fi77 := 0; fi77 < 1; fi77++ {
                                                            var v78 float32 = -2.25
                                                            fmt.Println(v5)
                                                            switch ((((((v7 + (("hello" + (("\n\t" + "") + v7)) + ((v7 + "a") + "a"))) + (v7 + v7)) + (v7 + v7)) + "") + v7) + (v7 + "foo bar")) {
                                                            case "foo bar":
                                                                if (-128 < 127) {
                                                                    var v79 int64 = (v16 % (((^(v16))) | 1))
                                                                    _ = v79
                                                                } else {
                                                                    for fi80 := 0; fi80 < 4; fi80++ {
                                                                        if v2 {
                                                                            v81 := GenS14{F0: (v70 << (((v5 / ((42) | 1))) & 63)), F1: (^(fi77)), F2: fi80}
                                                                            var v82 uint64 = 109
                                                                            _ = v81
                                                                            _ = v82
                                                                        } else {
                                                                            switch v7 {
                                                                            case "a":
                                                                                if v2 {
                                                                                    v5 -= v5
                                                                                    var v83 uint64 = 459
                                                                                    _ = v83
                                                                                }
                                                                            }
                                                                        }
                                                                        v78 -= (v78 / v78)
                                                                        _ = fi80
                                                                    }
                                                                }
                                                            case "":
                                                                fmt.Println(v5, v1, v1)
                                                            case "hello":
                                                                var v84 float64 = 1e10
                                                                _ = v84
                                                            }
                                                            _ = fi77
                                                            _ = v78
                                                        }
                                                    }
                                                }
                                                _ = v70
                                            } else {
                                                _ = len(v4)
                                                v85 := GenS11{F0: 42, F1: 0}
                                                _ = v85
                                            }
                                            for fi86 := 0; fi86 < 2; fi86++ {
                                                for fi87 := 0; fi87 < 14; fi87++ {
                                                    v7 = "a"
                                                    var v88 int8 = 1
                                                    _ = fi87
                                                    _ = v88
                                                }
                                                var v89 float64 = 1.5
                                                _ = fi86
                                                _ = v89
                                            }
                                        }
                                    }
                                    if (!false) {
                                        if v2 {
                                            var v90 float64 = -2.25
                                            if len(v3) > 0 {
                                                v3[0] = ("foo bar" + v7)
                                            }
                                            _ = v90
                                        } else {
                                            var v91 float64 = 0.0
                                            _ = v91
                                        }
                                        var v92 int64 = -1
                                        v3 = append(v3, (v7 + ((v7 + ((("foo bar" + ("\n\t" + v7)) + "") + ("" + "foo bar"))) + v7)))
                                        _ = v92
                                    }
                                }
                                _ = len(v4)
                            }
                            v7 = (v7 + v7)
                            for fi93 := 0; fi93 < 14; fi93++ {
                                if (((((v7 + "hello") + (v7 + "")) + "hello") + v7) <= ("" + v7)) {
                                    switch v7 {
                                    case "foo bar":
                                        delete(v4, v7)
                                    }
                                } else {
                                    v94 := GenS3{F0: (v16 - (v16 % ((-9223372036854775808) | 1))), F1: 1, F2: (v7 + "a")}
                                    var v95 uint64 = 151
                                    _ = v94
                                    _ = v95
                                }
                                var v96 uint16 = 669
                                v7 = ("" + "hello")
                                _ = fi93
                                _ = v96
                            }
                        } else {
                            var v97 string = "\n\t"
                            var v98 int32 = -987
                            if (((824 <= v5) && (2 > 2)) && v2) {
                                v5 -= (v5 * 0)
                                v7 = v97
                            } else {
                                for fi99 := 0; fi99 < 3; fi99++ {
                                    for fi100 := 0; fi100 < 18; fi100++ {
                                        if false {
                                            if v2 {
                                                if (v16 < v16) {
                                                    v1 -= v1
                                                }
                                            } else {
                                                v5 = (v5 % (((v5 * 2)) | 1))
                                                if ((!(v2 && true)) && (v97 == "foo bar")) {
                                                    var v101 int8 = -128
                                                    v102 := GenS15{F0: (v5 << ((v5) & 63)), F1: v98, F2: v16}
                                                    var v103 uint64 = 2
                                                    fmt.Println(v103, v97)
                                                    _ = len(v3)
                                                    var v104 uint16 = 65535
                                                    _ = v101
                                                    _ = v102
                                                    _ = v103
                                                    _ = v104
                                                } else {
                                                    if true {
                                                        _ = len(v3)
                                                    } else {
                                                        switch v7 {
                                                        case "hello":
                                                            var v105 int16 = int16(v1)
                                                            _ = v105
                                                        case "a":
                                                            v106 := GenS6{F0: -32768, F1: fi99}
                                                            _ = v106
                                                        default:
                                                            v107 := GenS16{F0: 127, F1: uint(v1)}
                                                            _ = v107
                                                        }
                                                        v108 := GenS17{F0: v5, F1: 9223372036854775807}
                                                        _ = v108
                                                    }
                                                }
                                            }
                                            var v109 float32 = 1e10
                                            var v110 int32 = v98
                                            v4[""] = fi99
                                            _ = v109
                                            _ = v110
                                        } else {
                                            v111 := GenS1{F0: (v16 ^ (v16 / ((v16) | 1))), F1: ((v97 + v97) + v7)}
                                            delete(v4, v7)
                                            _ = v111
                                        }
                                        v98 = (v98 ^ 169)
                                        v5 += 2
                                        _ = fi100
                                    }
                                    _ = fi99
                                }
                            }
                            _ = v97
                            _ = v98
                        }
                    }
                }
            }
        case "hello":
            var v112 string = v7
            _ = v112
        default:
            v113 := GenS1{F0: (v16 & v16), F1: ((v7 + v7) + "hello")}
            _ = v113
        }
        var v114 int32 = 42
        _ = v114
        _ = v16
    }
    v5 = v5
    v2 = (1 < 475)
    if ("foo bar" <= v7) {
        var v115 string = "foo bar"
        _ = v115
    } else {
        if true {
            var v116 uint64 = 4294967295
            v116 = v116
            for fi117 := 0; fi117 < 19; fi117++ {
                if false {
                    switch v7 {
                    case "a":
                        if ((0.1 * (0.0 * 1e10)) <= 1.5) {
                            switch fi117 {
                            case -2147483648:
                                v6[65535] = (0 != 255)
                            case -825:
                                v7 = (v7 + v7)
                            default:
                                var v118 int = (^(fi117))
                                _ = v118
                            }
                            switch fi117 {
                            case 224:
                                var v119 uint8 = (v1 >> (((^(v5))) & 63))
                                _ = v119
                            case 9223372036854775807:
                                if ((^(v1)) > v1) {
                                    v1 = (v1 / (((v1 - (v1 + 2))) | 1))
                                    v2 = false
                                }
                            }
                        } else {
                            var v120 bool = (false && ((^(v116)) <= (^(v116))))
                            for fi121 := 0; fi121 < 5; fi121++ {
                                var v122 int = fi121
                                _ = fi121
                                _ = v122
                            }
                            v120 = (true && v2)
                            _ = v120
                        }
                    case "\n\t":
                        var v123 int16 = 32767
                        _ = v123
                    case "foo bar":
                        var v124 int = -9223372036854775808
                        _ = v124
                    default:
                        v3 = append(v3, "a")
                    }
                    var v125 int32 = 2147483647
                    _ = v125
                } else {
                    for fi126 := 0; fi126 < 13; fi126++ {
                        v127 := GenS18{F0: 32767, F1: -124}
                        v5 -= (v5 << ((uint(v116)) & 63))
                        _ = fi126
                        _ = v127
                    }
                    v7 = (v7 + v7)
                }
                _ = fi117
            }
            _ = v116
        }
        v3 = append(v3, v7)
    }
    var v128 int8 = 103
    var v129 string = v7
    v5 -= (v5 / ((v5) | 1))
    var v130 bool = true
    v131 := v4[v129]
    v129 = ("" + "hello")
    v128 -= 34
    var v132 bool = (((1e10 + 0.1) < 0.1) && ((v131 >> ((1) & 63)) != v131))
    fmt.Println("v1:", v1)
    fmt.Println("v128:", v128)
    fmt.Println("v129:", v129)
    fmt.Println("v130:", v130)
    fmt.Println("v131:", v131)
    fmt.Println("v132:", v132)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v128
    _ = v129
    _ = v130
    _ = v131
    _ = v132
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
