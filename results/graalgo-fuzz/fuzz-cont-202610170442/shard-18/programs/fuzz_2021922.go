package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int64; F1 uint8 }
type GenS5 struct { F0 int64; F1 uint16; F2 int }
type GenS6 struct { F0 string }
type GenS7 struct { F0 int8; F1 int64 }
type GenS8 struct { F0 int32; F1 int }
type GenS9 struct { F0 uint16 }
type GenS10 struct { F0 int }
type GenS11 struct { F0 uint32; F1 string; F2 int32 }
func genFn0() uint32 {
    var v7 int = 32767
    v8 := []int{(v7 + v7), -742, (v7 * v7)}
    _ = v7
    _ = v8
    return 42
}
func genFn1(p0 uint64) string {
    v9 := GenS1{F0: 2147483647, F1: 1}
    for fi10 := 0; fi10 < 11; fi10++ {
        if (((^(p0)) <= (p0 + 18446744073709551615)) || (!false)) {
            var v11 int64 = -1
            p0 = p0
            _ = v11
        }
        v12 := GenS1{F0: (fi10 / (((-(fi10))) | 1)), F1: int32(p0)}
        if ((fi10 % ((fi10) | 1)) > -2147483648) {
            if true {
                if true {
                    fmt.Println(fi10, p0)
                    if ((true || (32767 <= 65535)) || (678 != 65535)) {
                        for fi13 := 0; fi13 < 4; fi13++ {
                            var v14 int32 = 0
                            v14 = (v14 & v14)
                            if false {
                                v15 := GenS2{F0: 255}
                                var v16 int64 = 0
                                _ = v15
                                _ = v16
                            } else {
                                var v17 bool = (!false)
                                switch (fi13 | fi13) {
                                case 128:
                                    var v18 uint = 1
                                    _ = v18
                                case 4294967295:
                                    var v19 int64 = -1
                                    _ = v19
                                default:
                                    fmt.Println(v17)
                                }
                                v20 := []int{(fi13 & fi10)}
                                if false {
                                    if len(v20) >= 2 {
                                        v20 = v20[1:len(v20)]
                                    }
                                    fmt.Println(v17)
                                    var v21 bool = false
                                    _ = v21
                                }
                                _ = v17
                                _ = v20
                            }
                            v22 := []int{fi10, (fi10 >> ((0) & 63))}
                            _ = fi13
                            _ = v14
                            _ = v22
                        }
                    }
                    var v23 uint32 = 2147483648
                    _ = v23
                } else {
                    for fi24 := 0; fi24 < 1; fi24++ {
                        v25 := []int{fi10, 0}
                        _ = fi24
                        _ = v25
                    }
                }
                var v26 bool = ((fi10 << ((1) & 63)) == (fi10 / ((fi10) | 1)))
                switch fi10 {
                case -9223372036854775808:
                    v27 := GenS2{F0: 0}
                    _ = v27
                case -1:
                    switch "" {
                    case "foo bar":
                        p0 = 42
                    }
                }
                _ = v26
            } else {
                for fi28 := 0; fi28 < 13; fi28++ {
                    var v29 uint8 = 128
                    if ((v29 << ((1) & 63)) != (v29 - (v29 * 97))) {
                        p0 += uint64(v29)
                        v30 := GenS1{F0: (fi10 ^ (fi28 * (fi10 + (fi10 & fi10)))), F1: -2147483648}
                        _ = v30
                    } else {
                        var v31 uint16 = 42
                        var v32 string = "hello"
                        fmt.Println(p0, v32)
                        v33 := []int{127}
                        v34 := GenS3{F0: (^(v31))}
                        v31 += 975
                        _ = v31
                        _ = v32
                        _ = v33
                        _ = v34
                    }
                    _ = fi28
                    _ = v29
                }
                p0 = (p0 ^ p0)
            }
        } else {
            for fi35 := 0; fi35 < 14; fi35++ {
                var v36 uint64 = 856
                switch fi10 {
                case 49:
                    p0 += 998
                case -296:
                    switch (-(fi35)) {
                    case 1:
                        v36 += 18446744073709551615
                    case 303:
                        v37 := GenS2{F0: 315}
                        _ = v37
                    default:
                        v38 := []int{(fi35 << ((1) & 63)), fi35}
                        _ = v38
                    }
                case 947:
                    p0 = (p0 & (v36 ^ 0))
                default:
                    if (2147483648 <= 468) {
                        p0 -= (p0 / ((42) | 1))
                        v39 := []int{(fi35 + 128), fi10, 32768}
                        var v40 uint32 = 108
                        _ = v39
                        _ = v40
                    } else {
                        fmt.Println(p0, p0)
                    }
                }
                fmt.Println(p0)
                _ = fi35
                _ = v36
            }
        }
        _ = fi10
        _ = v12
    }
    for fi41 := 0; fi41 < 12; fi41++ {
        p0 -= p0
        p0 = (p0 / ((2) | 1))
        _ = fi41
    }
    var v42 int64 = -9223372036854775808
    _ = p0
    _ = v42
    _ = v9
    return ("hello" + (("hello" + "foo bar") + ("hello" + "hello")))
}

func main() {
    var v1 uint64 = 1
    var v2 bool = false
    var v3 uint32 = 42
    var v4 uint32 = 431
    v5 := make(map[int]string)
    var v6 bool = v2
    switch -2147483648 {
    case 4294967295:
        v6 = true
    }
    v43 := []int{657}
    v1 += v1
    v5[9223372036854775807] = "a"
    v43 = append(v43, -1)
    if (42 != genFn0()) {
        if v6 {
            v44 := GenS4{F0: 125, F1: 1}
            for fi45 := 0; fi45 < 4; fi45++ {
                fmt.Println(v4, v4, v4)
                _ = fi45
            }
            switch 9223372036854775807 {
            case 128:
                v46 := v5[-130]
                _ = v46
            case 127:
                v1 = v1
            case 32767:
                fmt.Println(v2, v2)
            default:
                for fi47 := 0; fi47 < 9; fi47++ {
                    if (true || (42 == 65535)) {
                        if (!v2) {
                            if (v2 && (0.1 <= 3.14159)) {
                                v4 = (v4 + 42)
                            } else {
                                v5[914] = "a"
                                fmt.Println(fi47, v1)
                                switch ((((((("\n\t" + ((("a" + "") + "a") + "")) + ("foo bar" + "")) + "foo bar") + ("foo bar" + "hello")) + "") + "") + ("\n\t" + "")) {
                                case "\n\t":
                                    v4 = 0
                                case "foo bar":
                                    if ((v4 / ((v4) | 1)) < v4) {
                                        switch "" {
                                        case "foo bar":
                                            var v48 uint32 = 65535
                                            _ = v48
                                        default:
                                            var v49 int32 = 2147483647
                                            _ = v49
                                        }
                                    } else {
                                        var v50 uint64 = v1
                                        v3 = (^(v3))
                                        v2 = true
                                        _ = v50
                                    }
                                }
                            }
                            var v51 string = "hello"
                            _ = v51
                        } else {
                            v52 := GenS1{F0: (fi47 << ((42) & 63)), F1: -811}
                            _ = v52
                        }
                    }
                    switch ("hello" + "a") {
                    case "\n\t":
                        for fi53 := 0; fi53 < 3; fi53++ {
                            if ("hello" < "foo bar") {
                                v2 = false
                            }
                            v54 := GenS5{F0: 0, F1: 930, F2: (fi47 << ((1) & 63))}
                            if false {
                                switch 367 {
                                case -842:
                                    if v6 {
                                        var v55 int = 42
                                        _ = v55
                                    }
                                case 9223372036854775807:
                                    var v56 float64 = 0.1
                                    _ = v56
                                case 1:
                                    fmt.Println(fi47)
                                default:
                                    if false {
                                        if (0 < (-(fi53))) {
                                            v3 -= v3
                                        } else {
                                            if false {
                                                v57 := GenS6{F0: ("a" + "foo bar")}
                                                v2 = (!(285 <= 42))
                                                _ = v57
                                            } else {
                                                var v58 int16 = -1
                                                _ = v58
                                            }
                                            v2 = true
                                        }
                                        v59 := GenS7{F0: -128, F1: 9223372036854775807}
                                        _ = v59
                                    } else {
                                        v60 := GenS8{F0: -592, F1: fi53}
                                        v3 -= (^(v4))
                                        var v61 uint = 0
                                        _ = v60
                                        _ = v61
                                    }
                                }
                                fmt.Println(v1)
                                fmt.Println(v1, fi47)
                            }
                            _ = fi53
                            _ = v54
                        }
                    case "foo bar":
                        v4 = v4
                    default:
                        v43 = append(v43, (fi47 % ((fi47) | 1)))
                    }
                    if v2 {
                        for fi62 := 0; fi62 < 13; fi62++ {
                            v2 = v6
                            v63 := GenS4{F0: 1, F1: 0}
                            _ = fi62
                            _ = v63
                        }
                        v2 = (fi47 != fi47)
                        var v64 uint16 = 165
                        _ = v64
                    }
                    _ = fi47
                }
            }
            _ = v44
        } else {
            if v6 {
                if (v2 && v6) {
                    for fi65 := 0; fi65 < 11; fi65++ {
                        var v66 uint = 255
                        switch "foo bar" {
                        case "\n\t":
                            v67 := GenS2{F0: v66}
                            _ = v67
                        case "foo bar":
                            for fi68 := 0; fi68 < 9; fi68++ {
                                var v69 string = ("foo bar" + (("hello" + "a") + ("hello" + (("a" + ("a" + "")) + "hello"))))
                                _ = fi68
                                _ = v69
                            }
                        case "":
                            v70 := GenS4{F0: 0, F1: 42}
                            _ = v70
                        default:
                            if false {
                                v4 = genFn0()
                                if v2 {
                                    var v71 uint64 = (v1 & v1)
                                    v3 = (v3 * (^(v3)))
                                    _ = v71
                                } else {
                                    var v72 uint16 = 573
                                    var v73 bool = v6
                                    _ = v72
                                    _ = v73
                                }
                                v3 -= (v3 / ((v3) | 1))
                            }
                        }
                        _ = fi65
                        _ = v66
                    }
                } else {
                    var v74 uint = 0
                    for fi75 := 0; fi75 < 8; fi75++ {
                        v76 := GenS6{F0: ("a" + "\n\t")}
                        if (!false) {
                            var v77 int8 = 1
                            if len(v43) >= 2 {
                                v43 = v43[1:len(v43)]
                            }
                            fmt.Println(v77, v3)
                            _ = v77
                        } else {
                            for fi78 := 0; fi78 < 9; fi78++ {
                                v79 := GenS6{F0: ("foo bar" + ("foo bar" + ""))}
                                _ = fi78
                                _ = v79
                            }
                        }
                        _ = fi75
                        _ = v76
                    }
                    v1 += (^(v1))
                    _ = v74
                }
                v6 = v6
            } else {
                v80 := GenS3{F0: 0}
                _ = v80
            }
        }
        for fi81 := 0; fi81 < 10; fi81++ {
            if (678 == 1) {
                for fi82 := 0; fi82 < 13; fi82++ {
                    if len(v43) > 0 {
                        v43[0] = fi81
                    }
                    _ = fi82
                }
                if (1.5 > -2.25) {
                    v2 = ((v4 + uint32(v1)) < 65535)
                } else {
                    fmt.Println(v4, v1, v2)
                }
            }
            v3 += v4
            _ = fi81
        }
    } else {
        if ((v3 | (^(v3))) != 2) {
            v83 := GenS5{F0: -9223372036854775808, F1: 142, F2: 32768}
            v84 := GenS6{F0: "hello"}
            var v85 int16 = 42
            _ = v83
            _ = v84
            _ = v85
        } else {
            v2 = (!v6)
            var v86 string = ("\n\t" + "")
            v1 = v1
            _ = v86
        }
        var v87 int16 = 1000
        for fi88 := 0; fi88 < 17; fi88++ {
            for fi89 := 0; fi89 < 2; fi89++ {
                switch ("\n\t" + "a") {
                case "foo bar":
                    if len(v43) >= 2 {
                        v43 = v43[1:len(v43)]
                    }
                }
                if (!false) {
                    if (779 > -1) {
                        v90 := GenS4{F0: 0, F1: 222}
                        _ = v90
                    } else {
                        v4 -= uint32(v1)
                    }
                    fmt.Println(fi89, v4)
                    v91 := GenS8{F0: 1, F1: (^(fi88))}
                    var v92 int16 = (v87 / ((v87) | 1))
                    var v93 float64 = ((float64(float32(((1.5 * 0.1) - float64(-2.25)))) / 1.5) / 3.14159)
                    if ((!true) && (127 <= 127)) {
                        delete(v5, 32768)
                        if (v1 == 2) {
                            switch "a" {
                            case "":
                                v87 = v92
                            case "hello":
                                var v94 bool = (true || true)
                                if true {
                                    if false {
                                        var v95 int16 = 1
                                        _ = v95
                                    }
                                } else {
                                    switch fi89 {
                                    case 1:
                                        var v96 uint64 = v1
                                        var v97 int64 = 42
                                        v92 += 1000
                                        _ = v96
                                        _ = v97
                                    case 0:
                                        v98 := GenS9{F0: 2}
                                        _ = v98
                                    }
                                    v87 += v92
                                }
                                if v94 {
                                    var v99 bool = ((1e10 == (1.5 - 0.1)) || v6)
                                    _ = v99
                                }
                                v94 = (v94 || (((((("a" + "\n\t") + "foo bar") + "hello") + ("a" + "")) + "hello") != ""))
                                _ = v94
                            case "a":
                                v100 := GenS10{F0: fi89}
                                _ = v100
                            default:
                                var v101 float64 = v93
                                _ = v101
                            }
                            var v102 bool = v2
                            v103 := GenS11{F0: v4, F1: "", F2: 1}
                            _ = v102
                            _ = v103
                        }
                        if ((v2 && v6) || (0 != 31)) {
                            v4 = (v3 & (v4 | v3))
                        } else {
                            if v6 {
                                switch "a" {
                                case "":
                                    var v104 float64 = 1e10
                                    _ = v104
                                case "hello":
                                    var v105 int32 = 2147483647
                                    var v106 int16 = 32767
                                    _ = v105
                                    _ = v106
                                case "a":
                                    fmt.Println(v92)
                                    var v107 uint16 = 32767
                                    v108 := GenS4{F0: -9223372036854775808, F1: 42}
                                    _ = v107
                                    _ = v108
                                }
                            } else {
                                v92 = v87
                                v2 = (2147483647 < v4)
                                var v109 float32 = (0.0 - 3.14159)
                                _ = v109
                            }
                            v110 := GenS1{F0: 9223372036854775807, F1: 2147483647}
                            _ = v110
                        }
                    }
                    _ = v91
                    _ = v92
                    _ = v93
                }
                var v111 int64 = -10
                v43 = append(v43, -991)
                _ = fi89
                _ = v111
            }
            var v112 int16 = (v87 >> ((4294967295) & 63))
            v3 -= (v4 / (((v4 & 908)) | 1))
            _ = fi88
            _ = v112
        }
        _ = v87
    }
    var v113 uint = 672
    var v114 bool = v6
    var v115 int64 = -1
    v116 := GenS6{F0: "a"}
    _ = len(v5)
    var v117 uint32 = genFn0()
    var v118 int = 128
    switch v118 {
    case -2147483648:
        v119 := GenS9{F0: 802}
        _ = v119
    default:
        var v120 int = (v118 << ((v113) & 63))
        _ = v120
    }
    fmt.Println("v1:", v1)
    fmt.Println("v113:", v113)
    fmt.Println("v114:", v114)
    fmt.Println("v115:", v115)
    fmt.Println("v117:", v117)
    fmt.Println("v118:", v118)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v43:", len(v43), v43)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v113
    _ = v114
    _ = v115
    _ = v116
    _ = v117
    _ = v118
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v5
    _ = v6
}
