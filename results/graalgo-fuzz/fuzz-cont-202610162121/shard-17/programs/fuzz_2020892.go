package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int8; F1 int16; F2 string }
type GenS3 struct { F0 int8; F1 int8 }
type GenS4 struct { F0 uint32; F1 int8 }
type GenS5 struct { F0 uint16; F1 uint }
type GenS6 struct { F0 string; F1 bool }
type GenS7 struct { F0 bool; F1 uint32 }
func genFn0(p0 int8) uint32 {
    v5 := []int{0, 4294967295, 2147483647}
    if false {
        if len(v5) > 0 {
            v5[0] = 32768
        }
    }
    for fi6 := 0; fi6 < 15; fi6++ {
        var v7 int8 = 0
        p0 += p0
        _ = fi6
        _ = v7
    }
    _ = p0
    _ = v5
    return 65535
}
func genFn1(p0 string) string {
    switch p0 {
    case "":
        p0 = p0
    case "\n\t":
        var v8 int8 = 100
        _ = v8
    default:
        var v9 string = "foo bar"
        _ = v9
    }
    p0 = ("" + "")
    p0 = p0
    for fi10 := 0; fi10 < 17; fi10++ {
        p0 = p0
        if false {
            switch ((p0 + (p0 + p0)) + "") {
            case "foo bar":
                if false {
                    fmt.Println(p0, p0)
                }
            }
        } else {
            if (!false) {
                v11 := GenS1{F0: 0}
                _ = v11
            }
            for fi12 := 0; fi12 < 8; fi12++ {
                p0 = ((p0 + p0) + p0)
                var v13 int32 = 0
                v13 = v13
                _ = fi12
                _ = v13
            }
        }
        _ = fi10
    }
    p0 = p0
    _ = p0
    return p0
}

func main() {
    var v1 float32 = 1e10
    v2 := make(map[string]string)
    var v3 uint64 = 1
    var v4 int16 = -1
    var v14 uint32 = 42
    switch 128 {
    case 127:
        v1 += 1e10
    case -9223372036854775808:
        switch "hello" {
        case "hello":
            v15 := []int{}
            _ = v15
        case "a":
            v16 := []int{-858}
            _ = v16
        }
    default:
        if (uint(v14) == 429) {
            for fi17 := 0; fi17 < 11; fi17++ {
                v1 = (v1 * v1)
                fmt.Println(v14, v3)
                v2[""] = "foo bar"
                _ = fi17
            }
        } else {
            var v18 uint64 = v3
            var v19 int8 = 100
            v19 = (v19 % (((^(v19))) | 1))
            _ = v18
            _ = v19
        }
    }
    v20 := GenS1{F0: 4294967295}
    v3 += v3
    if false {
        if ("\n\t" == "a") {
            if false {
                var v21 uint16 = 2
                _ = v21
            } else {
                if false {
                    var v22 string = ("foo bar" + ("foo bar" + ""))
                    _ = v22
                }
                v2[("" + (((("" + ("hello" + "")) + "a") + "foo bar") + "a"))] = (("hello" + "hello") + ("\n\t" + "hello"))
                switch ("" + "") {
                case "foo bar":
                    for fi23 := 0; fi23 < 6; fi23++ {
                        switch 128 {
                        case 585:
                            var v24 uint64 = v3
                            _ = v24
                        case -171:
                            if ((!(884 >= 972)) && ((true && false) || false)) {
                                var v25 bool = ((fi23 & (fi23 * fi23)) == 128)
                                _ = v25
                            }
                        }
                        v3 = v3
                        v26 := GenS2{F0: 100, F1: -1, F2: (("foo bar" + (((("foo bar" + "\n\t") + "") + "") + "a")) + "hello")}
                        _ = fi23
                        _ = v26
                    }
                case "\n\t":
                    var v27 float64 = 0.0
                    _ = v27
                default:
                    for fi28 := 0; fi28 < 5; fi28++ {
                        var v29 uint = 65535
                        _ = fi28
                        _ = v29
                    }
                }
            }
            v14 -= (v14 & v14)
            delete(v2, "a")
        }
        var v30 uint64 = (^(v3))
        _ = v30
    } else {
        var v31 string = ("foo bar" + "a")
        _ = v31
    }
    if (true && (true && ((v3 + 18446744073709551615) <= (v3 - v3)))) {
        v32 := []int{-9223372036854775808, 4294967295, 32768}
        switch "" {
        case "hello":
            for fi33 := 0; fi33 < 16; fi33++ {
                if (uint16(v14) < 577) {
                    var v34 uint = 1
                    switch fi33 {
                    case 32768:
                        v34 += (v34 | v34)
                    case -2147483648:
                        v34 = v34
                    }
                    if true {
                        if (false || false) {
                            if (((true || (701 >= 1)) || (!((v3 | v3) != 18446744073709551615))) || true) {
                                v4 = 274
                            } else {
                                v3 = v3
                                var v35 int8 = -70
                                _ = v35
                            }
                            var v36 int32 = 29
                            _ = v36
                        } else {
                            var v37 int8 = 127
                            _ = v37
                        }
                        v38 := GenS3{F0: 42, F1: int8(v14)}
                        _ = v38
                    }
                    _ = v34
                } else {
                    var v39 bool = (fi33 == (fi33 / (((^(fi33))) | 1)))
                    _ = v39
                }
                fmt.Println(fi33)
                _ = fi33
            }
        default:
            switch 2147483647 {
            case -641:
                v2["foo bar"] = ""
            case 922:
                switch "foo bar" {
                case "foo bar":
                    if false {
                        v3 = (v3 << ((2) & 63))
                    }
                case "hello":
                    v4 += v4
                case "a":
                    v3 = v3
                }
            default:
                var v40 uint32 = 572
                _ = v40
            }
        }
        _ = v32
    } else {
        switch -53 {
        case 32768:
            v41 := GenS2{F0: 110, F1: -81, F2: ""}
            _ = v41
        case -9223372036854775808:
            fmt.Println(v4, v3, v14)
        case 9223372036854775807:
            if true {
                v14 = uint32(v3)
                var v42 float32 = v1
                v43 := []int{9223372036854775807, 4294967295}
                _ = v42
                _ = v43
            } else {
                v44 := v2[("a" + "foo bar")]
                switch 468 {
                case 0:
                    v14 -= (v14 / (((v14 * 4294967295)) | 1))
                default:
                    switch 314 {
                    case 4294967295:
                        for fi45 := 0; fi45 < 15; fi45++ {
                            v2[v44] = "foo bar"
                            _ = fi45
                        }
                    case 0:
                        var v46 uint32 = (v14 - (v14 << ((0) & 63)))
                        _ = v46
                    }
                }
                _ = v44
            }
        default:
            var v47 uint16 = 65535
            _ = v47
        }
        switch "hello" {
        case "foo bar":
            var v48 uint32 = v14
            _ = v48
        default:
            var v49 bool = true
            _ = v49
        }
        for fi50 := 0; fi50 < 5; fi50++ {
            for fi51 := 0; fi51 < 9; fi51++ {
                v14 = (v14 % ((v14) | 1))
                v52 := GenS1{F0: v3}
                if false {
                    if false {
                        switch "\n\t" {
                        case "":
                            v3 = v3
                        case "foo bar":
                            var v53 int8 = 42
                            _ = v53
                        case "hello":
                            var v54 uint8 = 0
                            fmt.Println(fi51, v4)
                            v55 := GenS1{F0: v3}
                            v56 := GenS2{F0: -4, F1: (^(v4)), F2: ("foo bar" + "")}
                            var v57 uint16 = 255
                            v14 -= 2147483647
                            _ = v54
                            _ = v55
                            _ = v56
                            _ = v57
                        default:
                            if false {
                                var v58 uint16 = 32767
                                var v59 uint = 255
                                if (false && (false && (((((fi50 & fi50) == -2147483648) && (54 <= 0)) && (!(!true))) || (false || (((v59 % (((v59 << ((v59) & 63))) | 1)) <= (v59 - v59)) || false))))) {
                                    v60 := []int{-9223372036854775808, fi51}
                                    _ = v60
                                } else {
                                    if (1 != -2147483648) {
                                        var v61 uint = 1
                                        fmt.Println(v61)
                                        _ = v61
                                    } else {
                                        v59 += 0
                                        fmt.Println(fi50)
                                        v62 := GenS3{F0: 127, F1: -128}
                                        _ = v62
                                    }
                                    fmt.Println(v4)
                                    if true {
                                        fmt.Println(fi51)
                                    } else {
                                        if (v58 >= v58) {
                                            switch fi51 {
                                            case -929:
                                                if ((((true && (v59 < (v59 - v59))) && ((347 >= v58) && false)) && (false || false)) && (true || false)) {
                                                    fmt.Println(fi50)
                                                    var v63 int16 = v4
                                                    _ = v63
                                                }
                                            case 42:
                                                var v64 int64 = -950
                                                _ = v64
                                            default:
                                                switch int(v59) {
                                                case 0:
                                                    delete(v2, "hello")
                                                case -2147483648:
                                                    v58 += v58
                                                    if false {
                                                        if true {
                                                            if (true || false) {
                                                                fmt.Println(fi51, v4, v1)
                                                                fmt.Println(v1, v1, v58)
                                                                if false {
                                                                    delete(v2, ("hello" + ("\n\t" + "foo bar")))
                                                                    switch "foo bar" {
                                                                    case "hello":
                                                                        var v65 int = 127
                                                                        _ = v65
                                                                    case "a":
                                                                        fmt.Println(fi51, v14)
                                                                    case "foo bar":
                                                                        v14 += 2147483647
                                                                    default:
                                                                        if (!((!(v14 != genFn0(23))) && (false || (true && true)))) {
                                                                            switch fi51 {
                                                                            case 0:
                                                                                if (643 <= (v59 << (((v59 / (((^(v59))) | 1))) & 63))) {
                                                                                    var v66 float32 = 1.5
                                                                                    v67 := GenS1{F0: v3}
                                                                                    _ = v66
                                                                                    _ = v67
                                                                                } else {
                                                                                    v68 := GenS1{F0: v3}
                                                                                    v69 := GenS4{F0: 4294967295, F1: 127}
                                                                                    v4 = 554
                                                                                    v59 = uint(v14)
                                                                                    v70 := v2["a"]
                                                                                    _ = v68
                                                                                    _ = v69
                                                                                    _ = v70
                                                                                }
                                                                                v14 += 65535
                                                                            case -9223372036854775808:
                                                                                var v71 string = ("a" + "")
                                                                                _ = v71
                                                                            case -329:
                                                                                var v72 int = 32768
                                                                                _ = v72
                                                                            }
                                                                            fmt.Println(v1, v59)
                                                                            var v73 int8 = -27
                                                                            _ = v73
                                                                        } else {
                                                                            if true {
                                                                                v14 = (v14 ^ genFn0(-1))
                                                                                v74 := GenS1{F0: v3}
                                                                                _ = v74
                                                                            } else {
                                                                                switch "\n\t" {
                                                                                case "":
                                                                                    if true {
                                                                                        var v75 uint64 = 897
                                                                                        _ = v75
                                                                                    } else {
                                                                                        v76 := GenS1{F0: v3}
                                                                                        if false {
                                                                                            var v77 int32 = 0
                                                                                            _ = v77
                                                                                        } else {
                                                                                            v1 = (v1 - (v1 * v1))
                                                                                            _ = len(v2)
                                                                                            v58 -= 255
                                                                                        }
                                                                                        if false {
                                                                                            v78 := GenS5{F0: v58, F1: 154}
                                                                                            if true {
                                                                                                var v79 int32 = 1
                                                                                                _ = v79
                                                                                            } else {
                                                                                                if (!true) {
                                                                                                    _ = len(v2)
                                                                                                    v80 := GenS6{F0: "", F1: (true && false)}
                                                                                                    switch (fi50 * fi50) {
                                                                                                    case 2147483647:
                                                                                                        var v81 int32 = 42
                                                                                                        _ = v81
                                                                                                    default:
                                                                                                        v82 := GenS2{F0: -128, F1: v4, F2: ""}
                                                                                                        _ = v82
                                                                                                    }
                                                                                                    _ = v80
                                                                                                }
                                                                                                var v83 int64 = -500
                                                                                                v84 := []int{fi50, fi51, (fi50 ^ 9223372036854775807)}
                                                                                                _ = v83
                                                                                                _ = v84
                                                                                            }
                                                                                            _ = v78
                                                                                        } else {
                                                                                            v85 := []int{9223372036854775807}
                                                                                            _ = len(v85)
                                                                                            switch ("foo bar" + ("foo bar" + ("foo bar" + "foo bar"))) {
                                                                                            case "":
                                                                                                v3 += (v3 << ((uint(v58)) & 63))
                                                                                            case "a":
                                                                                                switch fi51 {
                                                                                                case 32768:
                                                                                                    v85 = append(v85, -915)
                                                                                                }
                                                                                            case "\n\t":
                                                                                                var v86 int = fi50
                                                                                                _ = v86
                                                                                            }
                                                                                            _ = v85
                                                                                        }
                                                                                        var v87 int64 = 42
                                                                                        v4 += (-(v4))
                                                                                        _ = v76
                                                                                        _ = v87
                                                                                    }
                                                                                case "\n\t":
                                                                                    v4 = v4
                                                                                default:
                                                                                    v14 += (v14 + (v14 % (((v14 ^ (v14 ^ genFn0(42)))) | 1)))
                                                                                    v88 := GenS1{F0: 0}
                                                                                    var v89 float64 = 0.0
                                                                                    v90 := GenS7{F0: (true && true), F1: v14}
                                                                                    _ = v88
                                                                                    _ = v89
                                                                                    _ = v90
                                                                                }
                                                                                var v91 int = 9223372036854775807
                                                                                v1 = (v1 + (v1 * ((1.5 - v1) + v1)))
                                                                                _ = v91
                                                                            }
                                                                            if true {
                                                                                var v92 int64 = 0
                                                                                _ = v92
                                                                            } else {
                                                                                v58 -= v58
                                                                                v2["\n\t"] = ((("\n\t" + "\n\t") + ("\n\t" + ("hello" + "hello"))) + ("\n\t" + (((("foo bar" + "foo bar") + "foo bar") + ((("a" + "\n\t") + "\n\t") + "a")) + ("hello" + ("foo bar" + ("" + "\n\t"))))))
                                                                                var v93 float32 = 0.1
                                                                                _ = v93
                                                                            }
                                                                            fmt.Println(v1)
                                                                        }
                                                                    }
                                                                    if true {
                                                                        v94 := []int{(fi51 * (fi51 & (fi51 >> ((v59) & 63))))}
                                                                        var v95 int32 = 2147483647
                                                                        var v96 int64 = -654
                                                                        _ = v94
                                                                        _ = v95
                                                                        _ = v96
                                                                    }
                                                                }
                                                                if ((false && ((-33 < 0) && false)) && false) {
                                                                    var v97 uint = v59
                                                                    _ = v97
                                                                }
                                                            } else {
                                                                v59 = (v59 + v59)
                                                            }
                                                            v3 = 1
                                                        }
                                                        if (1e10 == 3.14159) {
                                                            var v98 uint32 = v14
                                                            var v99 int32 = 1
                                                            if (true && (0.1 != 1e10)) {
                                                                v98 = v98
                                                            } else {
                                                                fmt.Println(v3)
                                                                if ((v98 ^ genFn0(100)) > v98) {
                                                                    fmt.Println(fi51, fi50)
                                                                }
                                                                v100 := GenS6{F0: "\n\t", F1: (false || ("hello" != "\n\t"))}
                                                                _ = v100
                                                            }
                                                            _ = v98
                                                            _ = v99
                                                        } else {
                                                            v101 := GenS1{F0: v3}
                                                            v102 := []int{(fi50 / (((fi51 >> (((^(v59))) & 63))) | 1))}
                                                            _ = v101
                                                            _ = v102
                                                        }
                                                    } else {
                                                        v103 := v2[("foo bar" + "a")]
                                                        _ = v103
                                                    }
                                                default:
                                                    v104 := []int{-48, fi50, (fi50 ^ (fi50 >> ((v59) & 63)))}
                                                    _ = v104
                                                }
                                            }
                                            if false {
                                                v105 := v2[("" + ("a" + ("a" + "hello")))]
                                                _ = v105
                                            } else {
                                                var v106 uint64 = (v3 / (((v3 | 1)) | 1))
                                                _ = v106
                                            }
                                        }
                                        if (!false) {
                                            v1 = v1
                                        }
                                        v1 = v1
                                    }
                                }
                                v58 -= v58
                                _ = v58
                                _ = v59
                            }
                        }
                        v107 := []int{-387}
                        if (uint16(v3) <= 1) {
                            switch "\n\t" {
                            case "a":
                                fmt.Println(v1)
                            default:
                                v108 := v2["foo bar"]
                                _ = v108
                            }
                            v109 := v2[("\n\t" + "\n\t")]
                            _ = v109
                        }
                        _ = v107
                    } else {
                        v4 = (v4 & v4)
                        var v110 bool = false
                        var v111 string = "a"
                        var v112 float64 = -2.25
                        var v113 uint8 = 42
                        _ = v110
                        _ = v111
                        _ = v112
                        _ = v113
                    }
                    v114 := v2["hello"]
                    v14 -= v14
                    _ = v114
                }
                _ = fi51
                _ = v52
            }
            _ = fi50
        }
    }
    for fi115 := 0; fi115 < 5; fi115++ {
        switch "foo bar" {
        case "a":
            if true {
                v4 = 32767
            }
        case "\n\t":
            for fi116 := 0; fi116 < 5; fi116++ {
                v117 := GenS4{F0: v14, F1: 100}
                var v118 int8 = 42
                v3 = 2
                _ = fi116
                _ = v117
                _ = v118
            }
        case "foo bar":
            v14 = 334
        default:
            v14 -= uint32(v3)
        }
        if (-2147483648 > 568) {
            switch "a" {
            case "foo bar":
                var v119 uint16 = 65535
                _ = v119
            default:
                if false {
                    var v120 int32 = 0
                    if ((v120 == 0) || (((v3 % ((881) | 1)) >= 4294967295) || true)) {
                        if false {
                            v121 := []int{fi115, (fi115 * (fi115 + (-(fi115))))}
                            switch -1 {
                            case 32767:
                                v122 := GenS5{F0: 65535, F1: uint(v14)}
                                _ = v122
                            }
                            var v123 uint = 65535
                            _ = v121
                            _ = v123
                        } else {
                            var v124 uint = 42
                            var v125 int8 = -1
                            _ = v124
                            _ = v125
                        }
                        var v126 int = int(v14)
                        v126 = (fi115 - (fi115 ^ int(v4)))
                        _ = v126
                    }
                    _ = v120
                } else {
                    v1 -= v1
                }
            }
            v1 += v1
        }
        fmt.Println(fi115, v1)
        _ = fi115
    }
    for fi127 := 0; fi127 < 7; fi127++ {
        var v128 int = (fi127 & fi127)
        _ = fi127
        _ = v128
    }
    for fi129 := 0; fi129 < 8; fi129++ {
        var v130 int = 1
        _ = fi129
        _ = v130
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v2
    _ = v20
    _ = v3
    _ = v4
}
