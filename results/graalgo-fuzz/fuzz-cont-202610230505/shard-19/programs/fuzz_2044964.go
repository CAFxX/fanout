package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint32; F1 uint8; F2 int64 }
type GenS4 struct { F0 uint16; F1 string }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 uint16; F1 bool; F2 uint64 }
type GenS7 struct { F0 uint64; F1 int }
type GenS8 struct { F0 int32 }
type GenS9 struct { F0 string; F1 int8 }
type GenS10 struct { F0 uint32; F1 string; F2 uint }
type GenS11 struct { F0 uint32 }
type GenS12 struct { F0 int64; F1 int }
func genFn0() uint {
    switch 127 {
    case 32767:
        var v7 uint = 762
        _ = v7
    case -958:
        var v8 uint64 = 42
        _ = v8
    case -486:
        v9 := []int{127}
        _ = v9
    }
    if false {
        for fi10 := 0; fi10 < 14; fi10++ {
            for fi11 := 0; fi11 < 3; fi11++ {
                if (9223372036854775807 != -221) {
                    fmt.Println(fi11, fi11)
                    var v12 string = "foo bar"
                    v13 := []int{}
                    if ((0 < 852) || (1 != 2)) {
                        v12 = ""
                        if true {
                            var v14 int8 = -79
                            var v15 uint32 = 1
                            _ = v14
                            _ = v15
                        } else {
                            if len(v13) > 0 {
                                v13[0] = -382
                            }
                            if (fi10 < fi10) {
                                fmt.Println(fi10, fi10, fi11)
                                v16 := GenS1{F0: 17}
                                _ = v16
                            }
                        }
                        v12 = (v12 + v12)
                    } else {
                        var v17 int16 = 1000
                        v17 = (v17 % ((v17) | 1))
                        fmt.Println(fi11, v12, v17)
                        _ = v17
                    }
                    _ = v12
                    _ = v13
                } else {
                    v18 := []int{fi10, 1, (fi11 - 2147483647), fi10}
                    v19 := GenS2{F0: 981}
                    if (("foo bar" + "\n\t") > "") {
                        var v20 uint64 = 0
                        if len(v18) >= 2 {
                            v18 = v18[1:len(v18)]
                        }
                        v21 := GenS3{F0: 1, F1: 2, F2: -9223372036854775808}
                        _ = v20
                        _ = v21
                    }
                    var v22 uint64 = 66
                    var v23 bool = (123 <= 0)
                    var v24 int64 = -876
                    if true {
                        _ = len(v18)
                        v25 := GenS4{F0: 32767, F1: ("a" + "hello")}
                        _ = v25
                    } else {
                        if len(v18) > 0 {
                            v18[0] = fi11
                        }
                        v24 -= -9223372036854775808
                    }
                    _ = v18
                    _ = v19
                    _ = v22
                    _ = v23
                    _ = v24
                }
                var v26 int = 32767
                var v27 int32 = 0
                _ = fi11
                _ = v26
                _ = v27
            }
            var v28 uint64 = 2
            _ = fi10
            _ = v28
        }
    } else {
        var v29 bool = true
        v30 := GenS4{F0: 255, F1: "a"}
        if v29 {
            v31 := []int{4294967295}
            var v32 int64 = 9223372036854775807
            _ = v31
            _ = v32
        } else {
            v29 = (1.5 > 0.0)
        }
        _ = v29
        _ = v30
    }
    var v33 uint = 553
    v34 := []int{458, 374}
    switch "foo bar" {
    case "\n\t":
        fmt.Println(v33)
    case "a":
        var v35 uint32 = 2
        _ = v35
    case "hello":
        var v36 uint = 255
        _ = v36
    default:
        if (1.5 > (0.1 + -2.25)) {
            fmt.Println(v33)
            if false {
                if true {
                    var v37 float64 = 1e10
                    var v38 int = -180
                    _ = v37
                    _ = v38
                } else {
                    v39 := GenS5{F0: 0}
                    _ = v39
                }
                for fi40 := 0; fi40 < 12; fi40++ {
                    var v41 string = ("a" + "a")
                    _ = fi40
                    _ = v41
                }
            }
            for fi42 := 0; fi42 < 15; fi42++ {
                v43 := GenS4{F0: 65535, F1: ("foo bar" + ("\n\t" + "\n\t"))}
                _ = fi42
                _ = v43
            }
        }
    }
    _ = v33
    _ = v34
    return (^(v33))
}
func genFn1() uint8 {
    v44 := GenS2{F0: 1}
    switch 0 {
    case -1:
        if ((!true) || (2147483647 >= -2147483648)) {
            var v45 bool = false
            _ = v45
        } else {
            if (0 >= 1) {
                var v46 bool = false
                _ = v46
            } else {
                var v47 int = 127
                var v48 int64 = -9223372036854775808
                v48 = 9223372036854775807
                _ = v47
                _ = v48
            }
            var v49 int = 237
            _ = v49
        }
    case 118:
        var v50 int8 = -128
        _ = v50
    case 32768:
        for fi51 := 0; fi51 < 2; fi51++ {
            for fi52 := 0; fi52 < 1; fi52++ {
                var v53 string = "a"
                if (127 < 245) {
                    var v54 uint = 255
                    v55 := []int{(fi51 << ((v54) & 63))}
                    fmt.Println(v53, v53)
                    _ = v54
                    _ = v55
                } else {
                    v56 := []int{fi52, (fi52 - fi51)}
                    _ = v56
                }
                v57 := []int{}
                _ = fi52
                _ = v53
                _ = v57
            }
            var v58 uint32 = 2147483648
            _ = fi51
            _ = v58
        }
    default:
        if ((!false) && true) {
            var v59 int64 = 9223372036854775807
            _ = v59
        }
    }
    var v60 uint = 343
    _ = v44
    _ = v60
    return 231
}

func main() {
    var v1 float64 = -2.25
    v2 := make(map[int]bool)
    v3 := make(map[uint32]int)
    var v4 uint16 = 527
    var v5 int32 = 42
    var v6 int8 = 1
    v61 := GenS6{F0: (v4 - (v4 - 0)), F1: (!(!true)), F2: 350}
    _ = len(v2)
    switch 128 {
    case 478:
        for fi62 := 0; fi62 < 7; fi62++ {
            switch (fi62 + (fi62 << ((65535) & 63))) {
            case 1:
                v6 = (v6 & (^(v6)))
            default:
                if (0.1 != 0.1) {
                    if true {
                        switch (-(fi62)) {
                        case 32768:
                            if true {
                                v5 = (-(v5))
                                var v63 int32 = (^(v5))
                                _ = v63
                            } else {
                                switch ("\n\t" + "a") {
                                case "\n\t":
                                    for fi64 := 0; fi64 < 18; fi64++ {
                                        fmt.Println(fi62, v4)
                                        v65 := GenS7{F0: 18446744073709551615, F1: (fi62 | 313)}
                                        _ = fi64
                                        _ = v65
                                    }
                                case "a":
                                    if true {
                                        var v66 int = 32768
                                        for fi67 := 0; fi67 < 13; fi67++ {
                                            v68 := v3[876]
                                            _ = fi67
                                            _ = v68
                                        }
                                        var v69 int64 = 9223372036854775807
                                        _ = v66
                                        _ = v69
                                    }
                                default:
                                    if ((v4 % (((v4 + v4)) | 1)) <= 42) {
                                        if ((("foo bar" + ((("a" + "") + "a") + ("" + (("foo bar" + "a") + "foo bar")))) + "foo bar") <= "hello") {
                                            v4 -= 723
                                            v4 -= v4
                                            if (-2.25 != 0.1) {
                                                v70 := []int{fi62}
                                                for fi71 := 0; fi71 < 15; fi71++ {
                                                    v72 := GenS1{F0: 42}
                                                    v4 -= 164
                                                    if true {
                                                        v73 := GenS2{F0: 2}
                                                        _ = v73
                                                    } else {
                                                        fmt.Println(fi71, v5, fi62)
                                                        if true {
                                                            switch "foo bar" {
                                                            case "a":
                                                                switch (fi62 + (^(fi62))) {
                                                                case 319:
                                                                    var v74 float32 = 1e10
                                                                    _ = v74
                                                                }
                                                                v75 := GenS3{F0: 256, F1: 8, F2: 0}
                                                                var v76 int8 = v6
                                                                _ = v75
                                                                _ = v76
                                                            }
                                                            v5 = (-(v5))
                                                            var v77 uint = 113
                                                            v1 = float64(((((1.5 + (float32(float64(1e10)) * -2.25)) + 0.1) - ((-2.25 + 0.1) + 1e10)) - 0.1))
                                                            switch ("hello" + "a") {
                                                            case "foo bar":
                                                                if false {
                                                                    if ((97 < 0) && false) {
                                                                        v78 := GenS8{F0: (^(v5))}
                                                                        var v79 string = ("hello" + ("foo bar" + "hello"))
                                                                        _ = v78
                                                                        _ = v79
                                                                    }
                                                                    v80 := GenS9{F0: (("foo bar" + "hello") + ("\n\t" + "hello")), F1: 100}
                                                                    _ = v80
                                                                } else {
                                                                    v4 -= v4
                                                                    v6 -= (v6 % (((v6 ^ (v6 * v6))) | 1))
                                                                }
                                                            default:
                                                                v4 += 42
                                                            }
                                                            if true {
                                                                var v81 float32 = 0.0
                                                                if (true && ((v6 >> ((genFn0()) & 63)) != (v6 | v6))) {
                                                                    v82 := GenS10{F0: 2147483647, F1: ("\n\t" + ""), F2: v77}
                                                                    var v83 int32 = 0
                                                                    _ = v82
                                                                    _ = v83
                                                                } else {
                                                                    var v84 int32 = v5
                                                                    _ = v84
                                                                }
                                                                _ = v81
                                                            } else {
                                                                var v85 int16 = 1000
                                                                _ = v85
                                                            }
                                                            _ = len(v3)
                                                            _ = v77
                                                        } else {
                                                            var v86 int16 = 637
                                                            switch 438 {
                                                            case 42:
                                                                switch (fi71 ^ 127) {
                                                                case -34:
                                                                    fmt.Println(v86)
                                                                case 128:
                                                                    var v87 bool = ((4294967295 <= 163) && ((32767 <= v4) && false))
                                                                    _ = v87
                                                                }
                                                                var v88 int16 = (v86 * v86)
                                                                _ = v88
                                                            case -1:
                                                                if (!(!((!(("foo bar" + "\n\t") <= ("a" + ""))) && ("" == "\n\t")))) {
                                                                    fmt.Println(v86, fi62)
                                                                    fmt.Println(v1, v1, v6)
                                                                    var v89 uint16 = (v4 >> ((42) & 63))
                                                                    _ = v89
                                                                }
                                                            case 4294967295:
                                                                if len(v70) > 0 {
                                                                    v70[0] = fi71
                                                                }
                                                            }
                                                            v86 = v86
                                                            v6 = (v6 - 0)
                                                            var v90 bool = true
                                                            _ = v86
                                                            _ = v90
                                                        }
                                                        v1 = (1.5 - 3.14159)
                                                    }
                                                    if false {
                                                        v6 += 14
                                                        switch fi71 {
                                                        case -683:
                                                            fmt.Println(v5, fi62, fi71)
                                                        case 2147483647:
                                                            switch "a" {
                                                            case "":
                                                                fmt.Println(v6)
                                                            default:
                                                                switch fi71 {
                                                                case 472:
                                                                    if (42 > 68) {
                                                                        var v91 float64 = v1
                                                                        _ = v91
                                                                    }
                                                                case -382:
                                                                    fmt.Println(v5)
                                                                }
                                                            }
                                                        }
                                                    } else {
                                                        v1 += 0.0
                                                        v92 := GenS2{F0: 18446744073709551615}
                                                        _ = v92
                                                    }
                                                    var v93 uint = 1
                                                    _ = fi71
                                                    _ = v72
                                                    _ = v93
                                                }
                                                for fi94 := 0; fi94 < 20; fi94++ {
                                                    var v95 int8 = (v6 ^ v6)
                                                    v2[-598] = false
                                                    var v96 int64 = -459
                                                    _ = fi94
                                                    _ = v95
                                                    _ = v96
                                                }
                                                _ = v70
                                            }
                                        } else {
                                            for fi97 := 0; fi97 < 1; fi97++ {
                                                fmt.Println(v1, fi62)
                                                var v98 int32 = v5
                                                _ = fi97
                                                _ = v98
                                            }
                                            fmt.Println(v6)
                                            v99 := GenS8{F0: v5}
                                            _ = v99
                                        }
                                        var v100 float64 = v1
                                        switch ("hello" + ("foo bar" + "\n\t")) {
                                        case "":
                                            switch ("a" + "\n\t") {
                                            case "a":
                                                var v101 int32 = v5
                                                _ = v101
                                            case "foo bar":
                                                v102 := GenS8{F0: (-(v5))}
                                                _ = v102
                                            }
                                        }
                                        _ = v100
                                    } else {
                                        switch "hello" {
                                        case "\n\t":
                                            v2[fi62] = (!false)
                                        case "foo bar":
                                            var v103 uint32 = 2
                                            _ = v103
                                        case "a":
                                            _ = len(v2)
                                        default:
                                            fmt.Println(fi62, fi62, fi62)
                                        }
                                        if (!(false || (!(float32(v1) != 1.5)))) {
                                            v104 := GenS11{F0: 2}
                                            _ = v104
                                        } else {
                                            v105 := GenS11{F0: 1}
                                            fmt.Println(v6, v5)
                                            _ = v105
                                        }
                                    }
                                }
                                fmt.Println(fi62)
                            }
                        default:
                            v1 += 1.5
                        }
                        v106 := GenS2{F0: 2}
                        _ = v106
                    }
                } else {
                    var v107 uint32 = 750
                    var v108 float64 = v1
                    _ = v107
                    _ = v108
                }
            }
            var v109 uint64 = 42
            _ = fi62
            _ = v109
        }
    case 0:
        switch "" {
        case "\n\t":
            if false {
                v4 += 538
                v5 -= v5
                fmt.Println(v1, v1, v4)
            }
        case "":
            v110 := GenS8{F0: (v5 % ((v5) | 1))}
            _ = v110
        default:
            v4 += v4
        }
    case 1:
        if (18446744073709551615 != 18446744073709551615) {
            var v111 int = 32768
            for fi112 := 0; fi112 < 5; fi112++ {
                v1 += v1
                v113 := GenS11{F0: 0}
                var v114 bool = (false || true)
                _ = fi112
                _ = v113
                _ = v114
            }
            var v115 bool = (!true)
            _ = v111
            _ = v115
        }
    default:
        for fi116 := 0; fi116 < 13; fi116++ {
            v6 = (v6 | 68)
            var v117 uint64 = 42
            v5 = (v5 ^ (v5 & v5))
            _ = fi116
            _ = v117
        }
    }
    switch "a" {
    case "\n\t":
        var v118 int16 = -32768
        _ = v118
    case "":
        if (209 == 203) {
            v6 = -100
            fmt.Println(v1, v1)
        } else {
            v119 := []int{2147483647, 42, -9223372036854775808}
            var v120 int8 = (v6 % ((v6) | 1))
            _ = v119
            _ = v120
        }
    default:
        if ((-(v5)) != (v5 / ((0) | 1))) {
            v121 := GenS7{F0: 290, F1: 128}
            var v122 uint64 = 1
            var v123 int16 = -882
            _ = v121
            _ = v122
            _ = v123
        }
    }
    v124 := []int{683}
    var v125 uint64 = 478
    v126 := GenS12{F0: -1, F1: 32767}
    switch -1 {
    case -2147483648:
        for fi127 := 0; fi127 < 17; fi127++ {
            for fi128 := 0; fi128 < 19; fi128++ {
                var v129 int = (fi127 / (((fi128 << ((255) & 63))) | 1))
                delete(v2, fi127)
                _ = fi128
                _ = v129
            }
            _ = fi127
        }
    default:
        v4 = v4
    }
    v4 += (v4 ^ 32767)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v124:", len(v124), v124)
    fmt.Println("v125:", v125)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v124
    _ = v125
    _ = v126
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v61
}
