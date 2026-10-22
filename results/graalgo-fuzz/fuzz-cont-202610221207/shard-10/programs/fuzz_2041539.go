package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8 }
type GenS2 struct { F0 uint; F1 bool; F2 bool }
type GenS3 struct { F0 int64; F1 int64 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 uint64; F1 uint16 }
type GenS6 struct { F0 uint; F1 string }
func genFn0() int32 {
    if true {
        var v7 int = -1
        for fi8 := 0; fi8 < 3; fi8++ {
            var v9 string = "\n\t"
            var v10 int16 = -1
            for fi11 := 0; fi11 < 13; fi11++ {
                var v12 bool = (1.5 >= 3.14159)
                _ = fi11
                _ = v12
            }
            _ = fi8
            _ = v10
            _ = v9
        }
        _ = v7
    } else {
        var v13 uint32 = 748
        v14 := GenS1{F0: 32767, F1: -121}
        for fi15 := 0; fi15 < 4; fi15++ {
            v16 := []int{32768, fi15, fi15, (^(fi15))}
            fmt.Println(fi15)
            v13 = 2
            _ = fi15
            _ = v16
        }
        _ = v13
        _ = v14
    }
    if false {
        if false {
            v17 := GenS2{F0: 18446744073709551615, F1: (32767 == 42), F2: false}
            for fi18 := 0; fi18 < 6; fi18++ {
                for fi19 := 0; fi19 < 12; fi19++ {
                    fmt.Println(fi19, fi18)
                    _ = fi19
                }
                if (false || ((!true) && true)) {
                    if true {
                        var v20 uint32 = 2147483648
                        v21 := []int{0, fi18}
                        fmt.Println(v20, fi18)
                        _ = v20
                        _ = v21
                    }
                    var v22 int8 = 127
                    _ = v22
                }
                for fi23 := 0; fi23 < 8; fi23++ {
                    if true {
                        switch (fi23 & fi23) {
                        case 2147483647:
                            var v24 uint32 = 4294967295
                            _ = v24
                        case 32768:
                            var v25 int64 = 1
                            _ = v25
                        case 127:
                            var v26 string = ""
                            fmt.Println(v26, fi18, fi23)
                            v26 = v26
                            _ = v26
                        }
                        v27 := GenS1{F0: 255, F1: 59}
                        _ = v27
                    }
                    _ = fi23
                }
                _ = fi18
            }
            for fi28 := 0; fi28 < 20; fi28++ {
                var v29 int16 = 1000
                fmt.Println(v29, v29)
                _ = fi28
                _ = v29
            }
            _ = v17
        } else {
            if (!(32767 > 407)) {
                if ((752 == 42) && (2 < 1)) {
                    var v30 bool = (2 == 18446744073709551615)
                    v31 := GenS1{F0: 42, F1: 52}
                    _ = v30
                    _ = v31
                }
                var v32 float64 = (float64(float32((3.14159 / 3.14159))) / (1e10 - 0.1))
                _ = v32
            } else {
                var v33 uint16 = 868
                v34 := GenS1{F0: (^(v33)), F1: 1}
                v35 := []int{895, -2147483648, 9223372036854775807}
                _ = v33
                _ = v34
                _ = v35
            }
            if true {
                for fi36 := 0; fi36 < 8; fi36++ {
                    for fi37 := 0; fi37 < 1; fi37++ {
                        if (!(483 == 478)) {
                            fmt.Println(fi37, fi37)
                            var v38 float32 = 0.1
                            _ = v38
                        }
                        v39 := GenS2{F0: 4294967295, F1: true, F2: true}
                        var v40 int64 = 9223372036854775807
                        _ = fi37
                        _ = v39
                        _ = v40
                    }
                    fmt.Println(fi36)
                    if true {
                        if ((!(0 >= 214)) || (!((^(fi36)) <= fi36))) {
                            for fi41 := 0; fi41 < 11; fi41++ {
                                var v42 uint64 = 18446744073709551615
                                v42 -= v42
                                _ = fi41
                                _ = v42
                            }
                        } else {
                            var v43 int16 = 698
                            _ = v43
                        }
                        var v44 int16 = -1
                        if false {
                            v44 = (v44 + v44)
                            if true {
                                fmt.Println(v44, v44)
                                fmt.Println(v44)
                                for fi45 := 0; fi45 < 9; fi45++ {
                                    v46 := []int{fi36, fi45, 7}
                                    v46 = append(v46, (fi36 * (fi36 % ((fi36) | 1))))
                                    v44 -= v44
                                    _ = fi45
                                    _ = v46
                                }
                            }
                        } else {
                            for fi47 := 0; fi47 < 8; fi47++ {
                                v48 := GenS1{F0: 0, F1: -1}
                                v49 := GenS3{F0: 42, F1: -863}
                                v44 += -1
                                _ = fi47
                                _ = v48
                                _ = v49
                            }
                        }
                        _ = v44
                    }
                    _ = fi36
                }
            } else {
                var v50 string = "foo bar"
                v51 := []int{127}
                _ = v50
                _ = v51
            }
        }
        if (-32768 > 94) {
            var v52 uint = 0
            _ = v52
        }
    }
    var v53 uint64 = 976
    v53 = (v53 | (v53 / (((^(v53))) | 1)))
    _ = v53
    return 1
}

func main() {
    var v1 float64 = 1.5
    var v2 bool = true
    v3 := make(map[uint32]string)
    v4 := make(map[int]string)
    v5 := make(map[int]string)
    v6 := make(map[string]uint32)
    if (!v2) {
        v54 := GenS1{F0: 1, F1: 127}
        _ = v54
    }
    var v55 int16 = 0
    var v56 float64 = v1
    switch "" {
    case "\n\t":
        var v57 string = ("foo bar" + "")
        _ = v57
    case "":
        if (58 >= 18446744073709551615) {
            var v58 bool = v2
            var v59 float64 = v56
            for fi60 := 0; fi60 < 12; fi60++ {
                var v61 int64 = -641
                if (!(209 == 4294967295)) {
                    if ((764 <= 1) || (true || ((44 > 1) || (1 == 18446744073709551615)))) {
                        v2 = (1 >= 1)
                    }
                    var v62 int = -2147483648
                    _ = v62
                }
                var v63 float32 = 1.5
                _ = fi60
                _ = v61
                _ = v63
            }
            _ = v58
            _ = v59
        } else {
            fmt.Println(v2)
        }
    default:
        if (v2 && (4294967295 <= 42)) {
            v64 := GenS4{F0: 0}
            _ = v64
        } else {
            if (v2 && v2) {
                for fi65 := 0; fi65 < 4; fi65++ {
                    v66 := GenS4{F0: 179}
                    _ = fi65
                    _ = v66
                }
            } else {
                v56 = v56
            }
            v3[4294967295] = "foo bar"
        }
    }
    switch 42 {
    case -9223372036854775808:
        v56 -= (0.0 * v56)
    case -2147483648:
        v67 := GenS1{F0: 255, F1: -112}
        _ = v67
    }
    for fi68 := 0; fi68 < 8; fi68++ {
        if (!(833 >= 853)) {
            var v69 int64 = 462
            delete(v3, 2)
            _ = v69
        } else {
            var v70 uint8 = 1
            var v71 int8 = 0
            _ = v70
            _ = v71
        }
        var v72 uint8 = 0
        _ = fi68
        _ = v72
    }
    if ((^(v55)) >= -32768) {
        if ("" >= "a") {
            fmt.Println(v1)
            for fi73 := 0; fi73 < 5; fi73++ {
                var v74 uint64 = 2
                fmt.Println(fi73)
                _ = fi73
                _ = v74
            }
        } else {
            switch "hello" {
            case "\n\t":
                v55 += v55
            case "a":
                switch -195 {
                case -100:
                    var v75 uint32 = 1
                    _ = v75
                case -542:
                    v56 = v56
                case 127:
                    v1 = v56
                }
            case "":
                v2 = v2
            default:
                var v76 float64 = v1
                _ = v76
            }
        }
        var v77 int8 = 1
        for fi78 := 0; fi78 < 1; fi78++ {
            v2 = (v2 && (566 <= 9223372036854775807))
            _ = fi78
        }
        _ = v77
    }
    if v2 {
        var v79 int32 = 243
        switch 836 {
        case -1:
            v79 -= 2147483647
        case 9223372036854775807:
            for fi80 := 0; fi80 < 4; fi80++ {
                for fi81 := 0; fi81 < 10; fi81++ {
                    v82 := GenS1{F0: 32767, F1: -1}
                    var v83 bool = (v2 || (1000 <= 4294967295))
                    _ = fi81
                    _ = v82
                    _ = v83
                }
                for fi84 := 0; fi84 < 7; fi84++ {
                    var v85 uint32 = 108
                    v85 -= (v85 - (v85 ^ 4294967295))
                    switch (fi84 / (((fi84 - fi80)) | 1)) {
                    case 2147483647:
                        v1 += v1
                    case 4294967295:
                        if ((-2.25 + (v1 - (v56 / 0.1))) != (v1 + v1)) {
                            var v86 int = fi84
                            _ = v86
                        }
                    case 32768:
                        var v87 int16 = -1
                        _ = v87
                    default:
                        if (399 > (v85 >> ((416) & 63))) {
                            v88 := GenS3{F0: -1, F1: 0}
                            v85 -= v85
                            if (430 < 0) {
                                if (!(!v2)) {
                                    var v89 uint64 = 11
                                    if (255 == 316) {
                                        v56 -= v1
                                        var v90 uint64 = v89
                                        v2 = (v2 || v2)
                                        _ = len(v6)
                                        _ = v90
                                    } else {
                                        var v91 int64 = -1
                                        v89 -= (v89 << ((42) & 63))
                                        switch -9223372036854775808 {
                                        case 32768:
                                            var v92 int = -2147483648
                                            _ = v92
                                        case 42:
                                            switch 1 {
                                            case 42:
                                                v91 = (v91 / ((v91) | 1))
                                            default:
                                                fmt.Println(v79)
                                            }
                                            fmt.Println(fi84, v79)
                                            v2 = v2
                                        case -907:
                                            fmt.Println(v79)
                                        default:
                                            if (!v2) {
                                                fmt.Println(v56, v55)
                                                v79 = (-(v79))
                                                var v93 uint = 625
                                                _ = v93
                                            }
                                        }
                                        var v94 int8 = 127
                                        _ = v91
                                        _ = v94
                                    }
                                    v95 := GenS2{F0: 2, F1: false, F2: v2}
                                    if ((-(fi84)) == 2147483647) {
                                        v2 = v2
                                        delete(v5, -284)
                                        v89 = 286
                                    }
                                    _ = v89
                                    _ = v95
                                }
                                if (!(556 == 2)) {
                                    var v96 float64 = 1.5
                                    var v97 uint = 1
                                    _ = v96
                                    _ = v97
                                }
                            }
                            v98 := v3[2147483648]
                            _ = v88
                            _ = v98
                        }
                        fmt.Println(v2, v85)
                    }
                    _ = fi84
                    _ = v85
                }
                _ = fi80
            }
        case 0:
            for fi99 := 0; fi99 < 2; fi99++ {
                v79 = -1
                v100 := []int{fi99}
                v55 -= 629
                _ = fi99
                _ = v100
            }
        default:
            var v101 int16 = v55
            _ = v101
        }
        _ = v79
    }
    for fi102 := 0; fi102 < 10; fi102++ {
        var v103 int64 = -9223372036854775808
        _ = fi102
        _ = v103
    }
    for fi104 := 0; fi104 < 5; fi104++ {
        v105 := v5[(-(fi104))]
        switch fi104 {
        case 32767:
            v106 := GenS1{F0: 122, F1: 0}
            _ = v106
        case -556:
            v1 = ((1e10 * v1) / float64(float32(v1)))
        default:
            fmt.Println(v55, v105)
        }
        _ = fi104
        _ = v105
    }
    if v2 {
        var v107 float64 = v1
        var v108 int16 = (v55 << ((405) & 63))
        v109 := GenS5{F0: 42, F1: 1}
        _ = v107
        _ = v108
        _ = v109
    } else {
        for fi110 := 0; fi110 < 13; fi110++ {
            for fi111 := 0; fi111 < 14; fi111++ {
                v112 := GenS2{F0: 4294967295, F1: (848 > 732), F2: ("foo bar" > "a")}
                if v2 {
                    fmt.Println(fi110, v55, v56)
                } else {
                    var v113 uint32 = 218
                    var v114 uint32 = (v113 - v113)
                    v115 := GenS6{F0: 4294967295, F1: "foo bar"}
                    _ = v113
                    _ = v114
                    _ = v115
                }
                _ = fi111
                _ = v112
            }
            switch 1 {
            case -2147483648:
                fmt.Println(fi110, v1)
            case 113:
                var v116 uint = 1
                _ = v116
            }
            _ = fi110
        }
        for fi117 := 0; fi117 < 20; fi117++ {
            var v118 string = ""
            _ = fi117
            _ = v118
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v55:", v55)
    fmt.Println("v56:", v56)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v55
    _ = v56
    _ = v6
}
