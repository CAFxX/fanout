package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint64; F1 bool }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 string; F1 bool }
type GenS8 struct { F0 int16; F1 uint8; F2 uint8 }
type GenS9 struct { F0 uint8; F1 int8 }
type GenS10 struct { F0 int8; F1 uint16 }
type GenS11 struct { F0 int; F1 int8 }
type GenS12 struct { F0 uint16; F1 int; F2 uint }
type GenS13 struct { F0 int; F1 int }
type GenS14 struct { F0 uint16; F1 string; F2 int }
func genFn0() uint64 {
    switch -9223372036854775808 {
    case 32768:
        var v7 uint = 0
        _ = v7
    case 128:
        var v8 int8 = 100
        _ = v8
    case 127:
        switch "hello" {
        case "foo bar":
            v9 := []int{}
            _ = v9
        case "\n\t":
            for fi10 := 0; fi10 < 12; fi10++ {
                var v11 uint = 1
                var v12 float32 = (0.0 / (0.1 * 0.1))
                var v13 uint = v11
                _ = fi10
                _ = v11
                _ = v12
                _ = v13
            }
        case "hello":
            v14 := []int{287, 0, 1}
            _ = v14
        }
    default:
        var v15 bool = (467 < 1)
        _ = v15
    }
    var v16 uint = 610
    _ = v16
    return 446
}

func main() {
    var v1 int = 127
    v2 := make(map[string]bool)
    var v3 float64 = (-2.25 / ((1e10 + 0.1) / -2.25))
    var v4 float64 = (float64((float32(v3) - float32(-2.25))) - ((float64(1.5) * float64(1.5)) + (v3 / 1.5)))
    var v5 float64 = v4
    v6 := make(map[string]uint32)
    var v17 int64 = -588
    fmt.Println(v4, v1)
    var v18 int8 = -73
    var v19 uint32 = 1
    v6["hello"] = 2147483648
    v20 := []int{-1}
    var v21 uint64 = genFn0()
    v17 = 514
    if (true && true) {
        v20 = append(v20, 127)
        v1 += (v1 << ((1) & 63))
    }
    for fi22 := 0; fi22 < 5; fi22++ {
        fmt.Println(v19)
        _ = fi22
    }
    v23 := GenS1{F0: 2}
    v24 := GenS2{F0: uint16(v19)}
    v5 += (v4 - ((((float64((-2.25 / float32(v5))) * 1.5) + v3) * ((v4 * v5) - ((v4 * v3) / (0.0 - v3)))) - (float64((3.14159 / (0.0 - 0.1))) / v5)))
    if false {
        for fi25 := 0; fi25 < 9; fi25++ {
            if (!false) {
                var v26 uint32 = (v19 - uint32(v21))
                _ = v26
            }
            _ = fi25
        }
        v5 = 1.5
        var v27 uint32 = 190
        _ = v27
    } else {
        var v28 int16 = 0
        switch 9223372036854775807 {
        case -9223372036854775808:
            if (v28 == -841) {
                for fi29 := 0; fi29 < 15; fi29++ {
                    fmt.Println(v1, v18, v18)
                    if ((v19 % ((v19) | 1)) >= v19) {
                        var v30 bool = false
                        if (v3 == (v3 / 1.0)) {
                            v19 = (v19 * 2147483648)
                            var v31 float64 = (0.0 + float64(1.5))
                            v32 := GenS2{F0: 255}
                            _ = v31
                            _ = v32
                        } else {
                            var v33 int8 = v18
                            var v34 float32 = 3.14159
                            v18 = v18
                            _ = v33
                            _ = v34
                        }
                        _ = v30
                    } else {
                        var v35 float32 = (3.14159 - (-2.25 - 0.0))
                        for fi36 := 0; fi36 < 4; fi36++ {
                            if (((!true) || true) && (!true)) {
                                switch "a" {
                                case "":
                                    v35 = float32(v5)
                                }
                                v17 = (v17 | (v17 | 0))
                                fmt.Println(fi36)
                                switch (fi29 - fi29) {
                                case 9223372036854775807:
                                    v28 -= -541
                                    v37 := GenS3{F0: v28}
                                    v38 := GenS1{F0: 180}
                                    _ = v37
                                    _ = v38
                                case 127:
                                    if len(v20) >= 2 {
                                        v20 = v20[1:len(v20)]
                                    }
                                default:
                                    if false {
                                        if false {
                                            switch (fi36 - v1) {
                                            case -2147483648:
                                                var v39 int64 = (v17 ^ (v17 / ((9223372036854775807) | 1)))
                                                var v40 float32 = (1.5 / (v35 + v35))
                                                _ = v39
                                                _ = v40
                                            case 0:
                                                var v41 int = -186
                                                _ = v41
                                            case -9223372036854775808:
                                                var v42 uint64 = (^(v21))
                                                _ = v42
                                            default:
                                                fmt.Println(v17, v17, fi36)
                                            }
                                        } else {
                                            v43 := GenS4{F0: (v21 & 4294967295), F1: false}
                                            v44 := GenS5{F0: v19}
                                            _ = v43
                                            _ = v44
                                        }
                                        var v45 bool = false
                                        _ = v45
                                    }
                                }
                                if (v19 >= (v19 - v19)) {
                                    var v46 int16 = (v28 / ((v28) | 1))
                                    v5 = float64((v35 - v35))
                                    v47 := GenS5{F0: (^(v19))}
                                    _ = v46
                                    _ = v47
                                }
                            } else {
                                if (true || (111 == 681)) {
                                    fmt.Println(v1)
                                    switch fi29 {
                                    case 359:
                                        var v48 int = (^(v1))
                                        _ = v48
                                    case 2147483647:
                                        v20 = append(v20, (v1 % (((fi29 * fi29)) | 1)))
                                        delete(v2, "a")
                                        v49 := GenS6{F0: (-(v18))}
                                        _ = v49
                                    case -9223372036854775808:
                                        if true {
                                            var v50 uint32 = v19
                                            _ = v50
                                        } else {
                                            var v51 uint64 = 668
                                            v52 := GenS4{F0: v21, F1: false}
                                            _ = v51
                                            _ = v52
                                        }
                                    }
                                    var v53 uint16 = 137
                                    _ = v53
                                } else {
                                    switch fi29 {
                                    case 128:
                                        v54 := GenS3{F0: 42}
                                        _ = v54
                                    case 1:
                                        switch "a" {
                                        case "hello":
                                            var v55 int32 = 42
                                            _ = v55
                                        case "\n\t":
                                            if (!true) {
                                                if false {
                                                    fmt.Println(v5, v1, fi36)
                                                    var v56 int8 = v18
                                                    _ = v56
                                                }
                                                switch fi36 {
                                                case 127:
                                                    fmt.Println(v4)
                                                default:
                                                    v19 += (v19 | 2)
                                                }
                                                switch ("a" + "\n\t") {
                                                case "":
                                                    var v57 float64 = v4
                                                    _ = v57
                                                case "\n\t":
                                                    if (!((!(!(!true))) && (v35 > 0.0))) {
                                                        delete(v2, "a")
                                                        v58 := GenS5{F0: (^(v19))}
                                                        v1 -= fi36
                                                        _ = v58
                                                    }
                                                case "foo bar":
                                                    fmt.Println(v28, v4, fi36)
                                                default:
                                                    v59 := GenS4{F0: uint64(v19), F1: false}
                                                    _ = v59
                                                }
                                            } else {
                                                v60 := GenS7{F0: "\n\t", F1: false}
                                                fmt.Println(v35, fi36, v1)
                                                if false {
                                                    var v61 int64 = (v17 >> ((460) & 63))
                                                    _ = v61
                                                } else {
                                                    var v62 float32 = v35
                                                    if (1 == 0) {
                                                        v63 := GenS1{F0: 128}
                                                        if ("hello" <= "a") {
                                                            if (genFn0() > 100) {
                                                                fmt.Println(v3)
                                                                var v64 int16 = v28
                                                                _ = v64
                                                            }
                                                            var v65 string = "hello"
                                                            v65 = v65
                                                            _ = v65
                                                        } else {
                                                            switch (fi36 >> ((0) & 63)) {
                                                            case -2147483648:
                                                                if (false || (9223372036854775807 == (^(v17)))) {
                                                                    var v66 int32 = 0
                                                                    _ = v66
                                                                } else {
                                                                    v35 += float32(float64(1e10))
                                                                    var v67 int32 = 1
                                                                    fmt.Println(v35, v19)
                                                                    _ = v67
                                                                }
                                                            case 525:
                                                                v21 = uint64(v19)
                                                            case 0:
                                                                var v68 int8 = (v18 * (^(v18)))
                                                                _ = v68
                                                            default:
                                                                switch "hello" {
                                                                case "hello":
                                                                    if len(v20) >= 2 {
                                                                        v20 = v20[1:len(v20)]
                                                                    }
                                                                case "\n\t":
                                                                    var v69 uint16 = 255
                                                                    _ = v69
                                                                }
                                                            }
                                                            switch "" {
                                                            case "a":
                                                                delete(v2, "foo bar")
                                                            case "hello":
                                                                switch fi29 {
                                                                case 128:
                                                                    if (v28 >= 42) {
                                                                        v18 = v18
                                                                        if (!(!true)) {
                                                                            var v70 int64 = 1
                                                                            _ = v70
                                                                        } else {
                                                                            v28 -= (v28 - v28)
                                                                            switch v1 {
                                                                            case 563:
                                                                                v71 := v6["hello"]
                                                                                _ = v71
                                                                            case 724:
                                                                                v5 = (float64(-2.25) + v4)
                                                                            }
                                                                            v20 = append(v20, 799)
                                                                        }
                                                                    } else {
                                                                        v72 := GenS6{F0: -5}
                                                                        v73 := GenS1{F0: 127}
                                                                        v62 += float32(v3)
                                                                        _ = v72
                                                                        _ = v73
                                                                    }
                                                                case 4294967295:
                                                                    var v74 int8 = int8(v28)
                                                                    _ = v74
                                                                default:
                                                                    fmt.Println(v21, v35)
                                                                }
                                                            case "":
                                                                if false {
                                                                    if (float32((float64(v35) * v4)) <= v35) {
                                                                        v75 := GenS7{F0: ("a" + "foo bar"), F1: (!(!(((false || (!(true || true))) && (false || true)) && (-1 >= -2147483648))))}
                                                                        var v76 int16 = (v28 >> ((18446744073709551615) & 63))
                                                                        _ = v75
                                                                        _ = v76
                                                                    } else {
                                                                        fmt.Println(v1, fi29, v17)
                                                                        fmt.Println(v1, v21, v1)
                                                                        switch (("a" + "\n\t") + "") {
                                                                        case "foo bar":
                                                                            switch fi36 {
                                                                            case 32767:
                                                                                if false {
                                                                                    switch (fi36 * fi29) {
                                                                                    case 32767:
                                                                                        var v77 int64 = (v17 * (v17 << ((582) & 63)))
                                                                                        _ = v77
                                                                                    default:
                                                                                        var v78 int8 = (-(v18))
                                                                                        _ = v78
                                                                                    }
                                                                                    var v79 float32 = v35
                                                                                    if len(v20) > 0 {
                                                                                        v20[0] = -2147483648
                                                                                    }
                                                                                    _ = v79
                                                                                }
                                                                            }
                                                                        case "\n\t":
                                                                            var v80 int32 = -377
                                                                            _ = v80
                                                                        }
                                                                    }
                                                                    var v81 int8 = (^(v18))
                                                                    _ = v81
                                                                } else {
                                                                    if len(v20) > 0 {
                                                                        v20[0] = 9223372036854775807
                                                                    }
                                                                    var v82 int16 = v28
                                                                    var v83 uint64 = (v21 << ((18446744073709551615) & 63))
                                                                    v20 = append(v20, (^(fi36)))
                                                                    if false {
                                                                        switch (v1 ^ (-(v1))) {
                                                                        case 32768:
                                                                            v35 -= v62
                                                                        case 567:
                                                                            fmt.Println(fi29, v5)
                                                                        }
                                                                        v84 := GenS8{F0: -29, F1: 42, F2: 255}
                                                                        v85 := GenS9{F0: uint8(v19), F1: v18}
                                                                        var v86 float64 = v4
                                                                        _ = v84
                                                                        _ = v85
                                                                        _ = v86
                                                                    }
                                                                    _ = v82
                                                                    _ = v83
                                                                }
                                                            }
                                                            if (v28 > 409) {
                                                                fmt.Println(v21)
                                                                switch "\n\t" {
                                                                case "a":
                                                                    if ((9223372036854775807 <= 851) || false) {
                                                                        switch "\n\t" {
                                                                        case "":
                                                                            fmt.Println(fi36, fi29, v35)
                                                                            switch "" {
                                                                            case "foo bar":
                                                                                v5 = v4
                                                                            }
                                                                        case "a":
                                                                            _ = len(v2)
                                                                        default:
                                                                            fmt.Println(v4)
                                                                        }
                                                                    }
                                                                case "\n\t":
                                                                    var v87 int = (^(fi29))
                                                                    _ = v87
                                                                }
                                                                v4 -= v3
                                                            } else {
                                                                var v88 int16 = (v28 & (-(v28)))
                                                                if true {
                                                                    v88 = (v88 - (v88 / (((v28 * (v28 % ((int16(fi29)) | 1)))) | 1)))
                                                                    if true {
                                                                        v89 := v6[("foo bar" + "\n\t")]
                                                                        v21 -= (v21 / (((v21 | v21)) | 1))
                                                                        _ = v89
                                                                    }
                                                                }
                                                                v90 := v2["\n\t"]
                                                                _ = v88
                                                                _ = v90
                                                            }
                                                            switch (fi29 & (fi29 & (fi36 << ((317) & 63)))) {
                                                            case 9223372036854775807:
                                                                var v91 int64 = v17
                                                                _ = v91
                                                            case 42:
                                                                var v92 uint64 = genFn0()
                                                                _ = v92
                                                            case 166:
                                                                v93 := GenS7{F0: (("foo bar" + "a") + "a"), F1: (345 < 786)}
                                                                _ = v93
                                                            }
                                                            v28 = (v28 * (-(v28)))
                                                            if ((((!false) && (v19 < v19)) && ((!(!true)) || true)) && true) {
                                                                var v94 uint8 = 127
                                                                if (true || true) {
                                                                    var v95 float64 = v3
                                                                    _ = v95
                                                                } else {
                                                                    v96 := GenS5{F0: v19}
                                                                    fmt.Println(v35)
                                                                    v35 = 3.14159
                                                                    _ = v96
                                                                }
                                                                switch ("a" + "a") {
                                                                case "a":
                                                                    v97 := GenS10{F0: 100, F1: 255}
                                                                    v98 := GenS10{F0: 1, F1: uint16(v21)}
                                                                    _ = v97
                                                                    _ = v98
                                                                case "\n\t":
                                                                    if true {
                                                                        v21 = uint64(v94)
                                                                    } else {
                                                                        var v99 int16 = (-(v28))
                                                                        if len(v20) > 0 {
                                                                            v20[0] = 4294967295
                                                                        }
                                                                        _ = v99
                                                                    }
                                                                case "":
                                                                    var v100 int32 = 42
                                                                    _ = v100
                                                                default:
                                                                    v21 = v21
                                                                }
                                                                v101 := GenS11{F0: (-(fi29)), F1: v18}
                                                                switch fi29 {
                                                                case -9223372036854775808:
                                                                    var v102 uint32 = v19
                                                                    _ = v102
                                                                case -2147483648:
                                                                    var v103 uint32 = (v19 << ((18446744073709551615) & 63))
                                                                    _ = v103
                                                                case 703:
                                                                    var v104 float64 = (v5 * v5)
                                                                    _ = v104
                                                                default:
                                                                    if len(v20) > 0 {
                                                                        v20[0] = (fi29 * fi29)
                                                                    }
                                                                }
                                                                _ = v101
                                                                _ = v94
                                                            } else {
                                                                v3 -= ((v4 / v3) * (v3 / (1e10 * v3)))
                                                                var v105 float64 = 1e10
                                                                _ = v105
                                                            }
                                                        }
                                                        v35 = v35
                                                        v106 := GenS12{F0: 1, F1: -977, F2: 18446744073709551615}
                                                        if (v35 < v62) {
                                                            v107 := GenS13{F0: 0, F1: (v1 % ((fi36) | 1))}
                                                            _ = v107
                                                        }
                                                        _ = v106
                                                        _ = v63
                                                    }
                                                    _ = v62
                                                }
                                                _ = v60
                                            }
                                        case "a":
                                            fmt.Println(fi29, fi36)
                                        default:
                                            var v108 int8 = v18
                                            _ = v108
                                        }
                                    }
                                }
                                v28 -= v28
                                if true {
                                    fmt.Println(v4)
                                    v109 := GenS13{F0: v1, F1: (^(fi29))}
                                    v110 := GenS10{F0: v18, F1: 65535}
                                    if true {
                                        v111 := GenS14{F0: 0, F1: "hello", F2: (fi36 | fi29)}
                                        var v112 float64 = (float64(v35) - 3.14159)
                                        v18 -= 108
                                        fmt.Println(v1, v35)
                                        _ = v111
                                        _ = v112
                                    } else {
                                        if true {
                                            var v113 float32 = v35
                                            var v114 int32 = -2147483648
                                            _ = v113
                                            _ = v114
                                        } else {
                                            fmt.Println(v17, fi29, v3)
                                        }
                                        fmt.Println(v19)
                                    }
                                    _ = v109
                                    _ = v110
                                } else {
                                    if len(v20) >= 2 {
                                        v20 = v20[1:len(v20)]
                                    }
                                    _ = len(v20)
                                    var v115 uint16 = 255
                                    _ = v115
                                }
                                _ = len(v20)
                            }
                            var v116 int = (fi29 / ((int(v19)) | 1))
                            fmt.Println(fi29, v3, v1)
                            var v117 uint64 = v21
                            _ = fi36
                            _ = v116
                            _ = v117
                        }
                        v28 = v28
                        _ = v35
                    }
                    if len(v20) > 0 {
                        v20[0] = v1
                    }
                    _ = fi29
                }
            } else {
                if (!true) {
                    for fi118 := 0; fi118 < 19; fi118++ {
                        var v119 uint8 = 127
                        _ = fi118
                        _ = v119
                    }
                }
                fmt.Println(v5)
                if false {
                    v5 += v3
                    var v120 int32 = 1
                    var v121 int8 = v18
                    _ = v120
                    _ = v121
                }
            }
        }
        _ = v28
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
