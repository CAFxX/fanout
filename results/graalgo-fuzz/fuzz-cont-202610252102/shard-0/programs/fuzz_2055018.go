package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 string; F1 bool }
type GenS4 struct { F0 bool; F1 string }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 uint32; F1 string; F2 string }
type GenS7 struct { F0 uint16; F1 uint8 }
type GenS8 struct { F0 uint64 }
type GenS9 struct { F0 uint64; F1 int16 }
func genFn0(p0 int64, p1 int64, p2 bool) bool {
    for fi5 := 0; fi5 < 10; fi5++ {
        for fi6 := 0; fi6 < 18; fi6++ {
            v7 := GenS1{F0: 210}
            var v8 int = fi6
            p2 = true
            _ = fi6
            _ = v7
            _ = v8
        }
        _ = fi5
    }
    switch "" {
    case "hello":
        var v9 int8 = 69
        _ = v9
    case "":
        var v10 float32 = float32(0.0)
        _ = v10
    case "a":
        p2 = (0 <= 2)
    }
    v11 := []int{-561, 1}
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    return p2
}

func main() {
    v1 := []string{}
    v2 := make(map[int]string)
    v3 := make(map[string]string)
    v4 := make(map[int]string)
    if ((2 >= 30) || (!false)) {
        var v12 uint8 = 42
        _ = v12
    }
    for fi13 := 0; fi13 < 9; fi13++ {
        var v14 float64 = ((float64((1.5 - 3.14159)) - float64((float32(1e10) * (3.14159 * 0.1)))) * 3.14159)
        switch fi13 {
        case -2147483648:
            fmt.Println(fi13)
        case -9223372036854775808:
            v14 -= 0.1
        case 9223372036854775807:
            delete(v2, fi13)
        default:
            v15 := GenS2{F0: fi13}
            _ = v15
        }
        var v16 uint32 = 65535
        _ = fi13
        _ = v14
        _ = v16
    }
    var v17 uint64 = 2
    if (!(!true)) {
        v18 := v3["\n\t"]
        _ = v18
    } else {
        if (!((42 < 2) || (true || false))) {
            switch 128 {
            case -161:
                _ = len(v1)
            case 340:
                if (!true) {
                    if ((true && true) || (!true)) {
                        v19 := GenS3{F0: ("" + ""), F1: true}
                        fmt.Println(v17)
                        _ = v19
                    }
                    v17 += v17
                } else {
                    switch 325 {
                    case -2147483648:
                        var v20 int64 = -9223372036854775808
                        _ = v20
                    case 1:
                        if false {
                            delete(v4, 0)
                            switch int(v17) {
                            case 0:
                                v17 = (^(v17))
                            case 1:
                                var v21 int8 = -8
                                _ = v21
                            }
                        } else {
                            fmt.Println(v17)
                        }
                    case 42:
                        v22 := GenS4{F0: (65535 <= 125), F1: "\n\t"}
                        _ = v22
                    }
                    if (18446744073709551615 > 2) {
                        v17 = (v17 | v17)
                        v2[127] = "\n\t"
                    }
                }
            case 32768:
                if (false || (v17 >= (v17 * (v17 / (((^(v17))) | 1))))) {
                    if true {
                        var v23 int64 = 42
                        v17 += 756
                        _ = v23
                    }
                } else {
                    if true {
                        for fi24 := 0; fi24 < 1; fi24++ {
                            fmt.Println(v17)
                            v17 = (v17 ^ (v17 & 1))
                            if (4294967295 <= 65535) {
                                if (v17 < 1) {
                                    switch fi24 {
                                    case 127:
                                        v25 := GenS1{F0: 60}
                                        _ = v25
                                    case -2147483648:
                                        v26 := GenS5{F0: -128}
                                        _ = v26
                                    case -189:
                                        v17 = v17
                                    default:
                                        v17 = (v17 / ((v17) | 1))
                                    }
                                    v17 += (v17 | 264)
                                } else {
                                    for fi27 := 0; fi27 < 19; fi27++ {
                                        v28 := GenS2{F0: fi24}
                                        v1 = append(v1, "a")
                                        _ = fi27
                                        _ = v28
                                    }
                                    v17 = v17
                                    v17 -= (v17 << ((1) & 63))
                                }
                            }
                            _ = fi24
                        }
                        for fi29 := 0; fi29 < 13; fi29++ {
                            fmt.Println(v17, v17)
                            for fi30 := 0; fi30 < 10; fi30++ {
                                var v31 int16 = -145
                                var v32 int16 = v31
                                var v33 string = "hello"
                                _ = fi30
                                _ = v31
                                _ = v32
                                _ = v33
                            }
                            _ = fi29
                        }
                    } else {
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        _ = len(v3)
                    }
                    var v34 uint32 = 2147483647
                    _ = v34
                }
            default:
                for fi35 := 0; fi35 < 17; fi35++ {
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    if ((false && ((false || (((!true) && true) || false)) || true)) && true) {
                        var v36 uint32 = 2
                        var v37 int64 = -243
                        _ = v36
                        _ = v37
                    } else {
                        var v38 int64 = 42
                        fmt.Println(v38, v17)
                        _ = len(v2)
                        _ = v38
                    }
                    var v39 float32 = 1.5
                    _ = fi35
                    _ = v39
                }
            }
            if (1 < -1) {
                _ = len(v1)
                var v40 int32 = -1
                for fi41 := 0; fi41 < 17; fi41++ {
                    if (-32768 >= int16(v17)) {
                        v42 := GenS4{F0: true, F1: ("a" + "")}
                        if (!(1 == 4294967295)) {
                            fmt.Println(v17)
                            if false {
                                v40 = v40
                            } else {
                                for fi43 := 0; fi43 < 4; fi43++ {
                                    var v44 float32 = (3.14159 + (1.5 * float32(1e10)))
                                    _ = fi43
                                    _ = v44
                                }
                                switch fi41 {
                                case 128:
                                    if (-1 <= 0) {
                                        var v45 uint = 18446744073709551615
                                        _ = v45
                                    } else {
                                        if true {
                                            v46 := GenS2{F0: fi41}
                                            _ = v46
                                        }
                                        for fi47 := 0; fi47 < 15; fi47++ {
                                            if true {
                                                v48 := GenS3{F0: ("foo bar" + ("a" + "foo bar")), F1: (false && false)}
                                                var v49 float64 = 3.14159
                                                v17 = 539
                                                _ = v48
                                                _ = v49
                                            } else {
                                                switch ("hello" + "\n\t") {
                                                case "a":
                                                    v40 += (v40 % ((42) | 1))
                                                case "foo bar":
                                                    v17 = 1
                                                    v40 = 140
                                                    v50 := GenS1{F0: 214}
                                                    _ = v50
                                                case "hello":
                                                    var v51 string = ""
                                                    _ = v51
                                                }
                                            }
                                            _ = fi47
                                        }
                                    }
                                case 4294967295:
                                    var v52 int16 = int16(v40)
                                    _ = v52
                                case 127:
                                    for fi53 := 0; fi53 < 3; fi53++ {
                                        fmt.Println(v40, v40, fi41)
                                        v40 = -2147483648
                                        if len(v1) > 0 {
                                            v1[0] = "hello"
                                        }
                                        var v54 int8 = -61
                                        switch "hello" {
                                        case "foo bar":
                                            if false {
                                                var v55 float32 = (3.14159 * 1e10)
                                                if false {
                                                    v56 := GenS2{F0: fi41}
                                                    switch fi41 {
                                                    case -2147483648:
                                                        v54 = int8(v40)
                                                        var v57 int32 = (-(v40))
                                                        _ = v57
                                                    }
                                                    if len(v1) > 0 {
                                                        v1[0] = "\n\t"
                                                    }
                                                    _ = v56
                                                } else {
                                                    if len(v1) >= 2 {
                                                        v1 = v1[1:len(v1)]
                                                    }
                                                    if (!false) {
                                                        var v58 int8 = (v54 - -1)
                                                        if true {
                                                            var v59 int = fi41
                                                            delete(v2, fi53)
                                                            _ = v59
                                                        } else {
                                                            var v60 uint32 = 0
                                                            v4[(^(fi53))] = "hello"
                                                            _ = v60
                                                        }
                                                        var v61 bool = true
                                                        _ = v58
                                                        _ = v61
                                                    } else {
                                                        var v62 float64 = (0.1 + (1e10 + 0.0))
                                                        var v63 uint64 = (v17 * v17)
                                                        v64 := GenS6{F0: 284, F1: ("\n\t" + ("" + "hello")), F2: ("a" + "\n\t")}
                                                        _ = v62
                                                        _ = v63
                                                        _ = v64
                                                    }
                                                }
                                                if false {
                                                    var v65 float64 = 1e10
                                                    _ = v65
                                                } else {
                                                    var v66 uint64 = (v17 % (((v17 % (((v17 & (v17 ^ v17))) | 1))) | 1))
                                                    _ = v66
                                                }
                                                _ = v55
                                            } else {
                                                var v67 int32 = 0
                                                v40 += (-(v67))
                                                _ = v67
                                            }
                                        case "\n\t":
                                            var v68 string = "foo bar"
                                            _ = v68
                                        default:
                                            var v69 int16 = 0
                                            _ = v69
                                        }
                                        _ = fi53
                                        _ = v54
                                    }
                                default:
                                    if (!(2 >= 734)) {
                                        var v70 int16 = 32767
                                        v71 := GenS4{F0: (!true), F1: ""}
                                        _ = v70
                                        _ = v71
                                    }
                                }
                            }
                            v17 -= v17
                        } else {
                            _ = len(v1)
                            if len(v1) > 0 {
                                v1[0] = "hello"
                            }
                        }
                        fmt.Println(fi41, v17)
                        _ = v42
                    } else {
                        if (v40 <= 515) {
                            v72 := GenS2{F0: (fi41 * fi41)}
                            var v73 int8 = 100
                            _ = v72
                            _ = v73
                        }
                        _ = len(v4)
                        v17 = 137
                    }
                    _ = fi41
                }
                _ = v40
            }
        }
        var v74 float64 = 0.1
        _ = v74
    }
    for fi75 := 0; fi75 < 10; fi75++ {
        var v76 int64 = 42
        _ = fi75
        _ = v76
    }
    var v77 uint = 531
    v78 := GenS6{F0: 2147483648, F1: "", F2: "a"}
    var v79 uint = v77
    v79 = v79
    var v80 int64 = -9223372036854775808
    var v81 int32 = 103
    for fi82 := 0; fi82 < 2; fi82++ {
        v1 = append(v1, ("\n\t" + ""))
        _ = fi82
    }
    _ = len(v1)
    var v83 int = 32768
    if (v77 >= (v77 | (^(v79)))) {
        fmt.Println(v77, v79)
    } else {
        switch 127 {
        case 837:
            if (127 >= -1) {
                switch (v83 / ((9223372036854775807) | 1)) {
                case 32767:
                    fmt.Println(v79, v79)
                case 128:
                    switch -1 {
                    case 407:
                        v81 += v81
                    }
                default:
                    if (!false) {
                        var v84 uint16 = 42
                        v81 -= (v81 | (v81 / ((v81) | 1)))
                        _ = v84
                    } else {
                        var v85 int = (v83 >> ((v77) & 63))
                        v86 := GenS4{F0: false, F1: (("" + "\n\t") + ("foo bar" + ("" + ((("" + (("\n\t" + "hello") + "\n\t")) + ("hello" + ("a" + "hello"))) + "hello"))))}
                        fmt.Println(v83)
                        _ = v85
                        _ = v86
                    }
                }
            } else {
                v87 := GenS7{F0: 255, F1: 2}
                v88 := GenS8{F0: 4294967295}
                switch 171 {
                case -9223372036854775808:
                    v89 := GenS8{F0: v17}
                    _ = v89
                default:
                    if true {
                        for fi90 := 0; fi90 < 13; fi90++ {
                            switch ("\n\t" + ("a" + "a")) {
                            case "foo bar":
                                v91 := GenS7{F0: uint16(v17), F1: 2}
                                _ = v91
                            case "":
                                v77 = 293
                            case "\n\t":
                                v17 -= v17
                            default:
                                for fi92 := 0; fi92 < 3; fi92++ {
                                    var v93 int = fi90
                                    _ = fi92
                                    _ = v93
                                }
                            }
                            v94 := GenS9{F0: (v17 | (^(v17))), F1: 0}
                            _ = fi90
                            _ = v94
                        }
                    }
                }
                _ = v87
                _ = v88
            }
        case -576:
            v77 -= (^(v77))
        case 32767:
            var v95 uint = 428
            _ = v95
        default:
            var v96 uint32 = 2147483647
            _ = v96
        }
    }
    if (uint(v17) < 255) {
        switch "a" {
        case "":
            v17 = 1
        case "\n\t":
            v97 := v2[v83]
            _ = v97
        case "hello":
            v79 = 42
        }
        if (2 >= 855) {
            if (1 <= 0) {
                v83 += v83
                v98 := GenS5{F0: 0}
                fmt.Println(v83)
                _ = v98
            } else {
                var v99 float64 = float64((0.0 * 0.0))
                _ = v99
            }
            delete(v4, v83)
        }
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v77:", v77)
    fmt.Println("v79:", v79)
    fmt.Println("v80:", v80)
    fmt.Println("v81:", v81)
    fmt.Println("v83:", v83)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v77
    _ = v78
    _ = v79
    _ = v80
    _ = v81
    _ = v83
}
