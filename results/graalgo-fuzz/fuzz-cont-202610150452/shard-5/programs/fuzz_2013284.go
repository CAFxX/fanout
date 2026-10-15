package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int16; F1 int32 }
type GenS3 struct { F0 int; F1 string }
type GenS4 struct { F0 uint16; F1 bool; F2 bool }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 uint32; F1 uint8; F2 int16 }
type GenS7 struct { F0 uint16; F1 uint8; F2 uint32 }
type GenS8 struct { F0 uint16; F1 uint16 }
type GenS9 struct { F0 uint8 }
type GenS10 struct { F0 uint8; F1 uint; F2 uint16 }
type GenS11 struct { F0 uint64; F1 int8; F2 int16 }
type GenS12 struct { F0 uint32 }
type GenS13 struct { F0 string; F1 string }
func genFn0(p0 string, p1 string, p2 int32) uint32 {
    fmt.Println(p2)
    v4 := []int{0, -285, 2147483647}
    var v5 uint8 = 127
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return 366
}
func genFn1() string {
    for fi6 := 0; fi6 < 7; fi6++ {
        for fi7 := 0; fi7 < 17; fi7++ {
            var v8 uint = 255
            switch ("" + ("\n\t" + "\n\t")) {
            case "":
                v9 := []int{}
                _ = v9
            case "a":
                v10 := GenS1{F0: -218}
                _ = v10
            case "\n\t":
                if ((false || true) && ("foo bar" != "")) {
                    v11 := GenS1{F0: fi6}
                    if true {
                        var v12 int32 = 0
                        fmt.Println(v12, fi7, fi7)
                        var v13 int32 = 0
                        if (fi7 <= (fi6 >> (((v8 << ((18446744073709551615) & 63))) & 63))) {
                            if true {
                                v12 += -168
                                v14 := GenS2{F0: 1, F1: (v12 | v13)}
                                switch fi6 {
                                case -9223372036854775808:
                                    v15 := []int{fi7, 42}
                                    _ = v15
                                case -154:
                                    if (v8 != v8) {
                                        v8 = 255
                                    }
                                case 9223372036854775807:
                                    if true {
                                        var v16 int64 = 42
                                        v17 := []int{fi6}
                                        fmt.Println(v16)
                                        _ = v16
                                        _ = v17
                                    } else {
                                        v18 := GenS3{F0: fi6, F1: "a"}
                                        switch ("hello" + "") {
                                        case "hello":
                                            var v19 uint64 = 42
                                            _ = v19
                                        case "":
                                            v8 += (v8 + (v8 - 2))
                                        }
                                        _ = v18
                                    }
                                default:
                                    fmt.Println(fi7, v8)
                                }
                                _ = v14
                            } else {
                                v8 -= v8
                            }
                        }
                        _ = v12
                        _ = v13
                    } else {
                        var v20 float64 = 0.0
                        _ = v20
                    }
                    _ = v11
                } else {
                    v8 = (v8 | 407)
                    if true {
                        v21 := []int{(-(fi6)), -981}
                        _ = v21
                    }
                }
            default:
                v8 -= v8
            }
            v22 := []int{}
            _ = fi7
            _ = v22
            _ = v8
        }
        v23 := []int{}
        var v24 int32 = -1
        _ = fi6
        _ = v23
        _ = v24
    }
    var v25 uint32 = 2147483647
    for fi26 := 0; fi26 < 8; fi26++ {
        v27 := GenS1{F0: -2147483648}
        _ = fi26
        _ = v27
    }
    if true {
        var v28 uint32 = 2
        for fi29 := 0; fi29 < 20; fi29++ {
            v25 -= (v28 >> ((1) & 63))
            v30 := GenS4{F0: 65535, F1: true, F2: (-1 != -71)}
            var v31 int32 = 845
            _ = fi29
            _ = v30
            _ = v31
        }
        _ = v28
    } else {
        var v32 int32 = -435
        _ = v32
    }
    v33 := []int{}
    _ = v25
    _ = v33
    return ""
}

func main() {
    var v1 int8 = 100
    var v2 float64 = 3.14159
    v3 := make(map[string]uint32)
    switch "" {
    case "a":
        if true {
            var v34 int16 = 1
            var v35 int32 = 42
            for fi36 := 0; fi36 < 16; fi36++ {
                _ = len(v3)
                switch "a" {
                case "\n\t":
                    v1 = int8(v35)
                }
                _ = fi36
            }
            _ = v34
            _ = v35
        } else {
            v37 := GenS1{F0: 32767}
            v2 += (v2 + 1e10)
            _ = v37
        }
    case "foo bar":
        if false {
            delete(v3, "\n\t")
        } else {
            v1 += 0
            v38 := []int{-9223372036854775808}
            _ = v38
        }
    default:
        v1 += v1
    }
    if (true || (821 >= 141)) {
        switch 32767 {
        case 1:
            fmt.Println(v2, v1)
        case -1:
            for fi39 := 0; fi39 < 17; fi39++ {
                var v40 string = "\n\t"
                _ = fi39
                _ = v40
            }
        default:
            for fi41 := 0; fi41 < 3; fi41++ {
                switch ((("hello" + "\n\t") + "") + "\n\t") {
                case "foo bar":
                    v2 = float64(0.0)
                case "hello":
                    v42 := GenS3{F0: (fi41 >> ((439) & 63)), F1: "a"}
                    _ = v42
                case "":
                    v2 = (v2 / 1e10)
                default:
                    for fi43 := 0; fi43 < 7; fi43++ {
                        v44 := []int{fi43, (fi41 >> ((42) & 63)), fi41}
                        var v45 int = -115
                        _ = fi43
                        _ = v44
                        _ = v45
                    }
                }
                _ = fi41
            }
        }
        v46 := GenS1{F0: -9223372036854775808}
        v47 := GenS5{F0: 4294967295}
        _ = v46
        _ = v47
    } else {
        v1 = v1
        fmt.Println(v1, v2)
        v48 := []int{0, -2147483648}
        _ = v48
    }
    for fi49 := 0; fi49 < 13; fi49++ {
        fmt.Println(fi49, v1, v1)
        _ = fi49
    }
    switch "\n\t" {
    case "foo bar":
        v2 = (v2 + v2)
    case "hello":
        fmt.Println(v1, v2)
    case "":
        v1 = v1
    }
    for fi50 := 0; fi50 < 19; fi50++ {
        v2 = 1.5
        v51 := GenS6{F0: 2147483648, F1: 23, F2: -32768}
        _ = fi50
        _ = v51
    }
    for fi52 := 0; fi52 < 7; fi52++ {
        if true {
            if ((false && false) || (-2.25 != ((0.0 * 0.0) * (-2.25 / 1.5)))) {
                if (755 < 0) {
                    v1 = (v1 % ((v1) | 1))
                } else {
                    v53 := GenS6{F0: 2147483648, F1: 42, F2: 1000}
                    var v54 float32 = float32(v2)
                    v55 := GenS7{F0: 0, F1: 243, F2: genFn0("", "hello", -2147483648)}
                    _ = v53
                    _ = v54
                    _ = v55
                }
                var v56 float64 = (v2 * v2)
                for fi57 := 0; fi57 < 14; fi57++ {
                    fmt.Println(v1, v56, v56)
                    _ = fi57
                }
                _ = v56
            } else {
                for fi58 := 0; fi58 < 11; fi58++ {
                    var v59 int32 = -180
                    _ = fi58
                    _ = v59
                }
            }
        } else {
            fmt.Println(fi52)
            v1 = (-(v1))
        }
        for fi60 := 0; fi60 < 12; fi60++ {
            v61 := GenS4{F0: 32767, F1: false, F2: false}
            v62 := GenS2{F0: -104, F1: 0}
            _ = fi60
            _ = v61
            _ = v62
        }
        _ = fi52
    }
    switch "a" {
    case "\n\t":
        if (42 <= 234) {
            v1 = v1
        } else {
            v1 = (-(v1))
            for fi63 := 0; fi63 < 15; fi63++ {
                v64 := []int{}
                if false {
                    v1 -= 1
                }
                _ = fi63
                _ = v64
            }
        }
    case "hello":
        var v65 string = ((("" + ("a" + "a")) + "\n\t") + "foo bar")
        _ = v65
    }
    var v66 int32 = -1
    v67 := GenS7{F0: 65535, F1: 200, F2: 2147483648}
    v68 := GenS7{F0: 2, F1: 2, F2: 2}
    fmt.Println(v66, v1)
    var v69 int = -839
    for fi70 := 0; fi70 < 10; fi70++ {
        var v71 uint64 = 1
        var v72 int32 = v66
        for fi73 := 0; fi73 < 20; fi73++ {
            if false {
                if (!(4294967295 >= (v71 + v71))) {
                    var v74 string = "hello"
                    _ = v74
                }
            }
            if false {
                v71 = 2
            } else {
                fmt.Println(v71, v1)
                v3["hello"] = 1
                if false {
                    v1 -= (v1 << ((1) & 63))
                }
            }
            _ = fi73
        }
        _ = fi70
        _ = v71
        _ = v72
    }
    v75 := GenS8{F0: 32767, F1: 2}
    var v76 uint16 = 1
    if true {
        v77 := GenS9{F0: 0}
        v78 := GenS10{F0: 90, F1: 255, F2: v76}
        v79 := []int{(v69 - v69), (v69 & v69), (-(v69))}
        _ = v77
        _ = v78
        _ = v79
    } else {
        if false {
            for fi80 := 0; fi80 < 3; fi80++ {
                var v81 int16 = -1
                _ = fi80
                _ = v81
            }
            fmt.Println(v76, v76)
            switch ((("\n\t" + "") + (("hello" + "foo bar") + "a")) + "hello") {
            case "\n\t":
                var v82 int32 = (-(v66))
                _ = v82
            case "":
                if true {
                    v1 = v1
                    if (!((42 > -1) && false)) {
                        if true {
                            for fi83 := 0; fi83 < 11; fi83++ {
                                v84 := GenS10{F0: 242, F1: 4294967295, F2: 32767}
                                v85 := GenS11{F0: 85, F1: 77, F2: 42}
                                var v86 bool = (42 == 18446744073709551615)
                                _ = fi83
                                _ = v84
                                _ = v85
                                _ = v86
                            }
                            for fi87 := 0; fi87 < 6; fi87++ {
                                for fi88 := 0; fi88 < 2; fi88++ {
                                    var v89 int8 = v1
                                    _ = fi88
                                    _ = v89
                                }
                                _ = fi87
                            }
                        }
                        switch v69 {
                        case -9223372036854775808:
                            if false {
                                if (true && false) {
                                    _ = len(v3)
                                }
                                if true {
                                    fmt.Println(v66, v76, v76)
                                }
                                v66 -= -1
                            }
                        case 1:
                            v69 += v69
                        }
                        for fi90 := 0; fi90 < 2; fi90++ {
                            v1 = (v1 | (v1 | v1))
                            v91 := GenS8{F0: 65535, F1: (v76 | v76)}
                            _ = fi90
                            _ = v91
                        }
                    } else {
                        v76 += (v76 + v76)
                        v92 := GenS12{F0: 1}
                        switch "\n\t" {
                        case "\n\t":
                            var v93 float64 = float64(0.0)
                            _ = v93
                        default:
                            v69 -= (v69 | v69)
                        }
                        _ = v92
                    }
                    switch "foo bar" {
                    case "foo bar":
                        v94 := GenS8{F0: v76, F1: (v76 / (((v76 + 65535)) | 1))}
                        _ = v94
                    case "a":
                        if (255 >= 1) {
                            v95 := []int{4294967295, (v69 & v69), (v69 % (((-(v69))) | 1))}
                            var v96 uint16 = 65535
                            _ = v95
                            _ = v96
                        }
                    default:
                        var v97 float64 = 1.5
                        _ = v97
                    }
                } else {
                    v98 := []int{(v69 ^ v69), v69}
                    if true {
                        var v99 int16 = -32768
                        _ = v99
                    }
                    if len(v98) >= 2 {
                        v98 = v98[1:len(v98)]
                    }
                    _ = v98
                }
            case "foo bar":
                _ = len(v3)
            }
        }
        if ((false || true) || ((!(false && false)) || (0.0 >= 0.0))) {
            var v100 uint32 = 837
            v101 := GenS13{F0: "\n\t", F1: ""}
            _ = v100
            _ = v101
        } else {
            v102 := GenS13{F0: "hello", F1: "a"}
            v2 += v2
            _ = v102
        }
    }
    if (uint64(v76) < 18446744073709551615) {
        if (!true) {
            for fi103 := 0; fi103 < 4; fi103++ {
                var v104 int32 = (v66 & v66)
                for fi105 := 0; fi105 < 16; fi105++ {
                    v69 += v69
                    v106 := GenS8{F0: v76, F1: 65535}
                    v107 := GenS4{F0: (v76 * v76), F1: (false && false), F2: (1 <= 2)}
                    _ = fi105
                    _ = v106
                    _ = v107
                }
                _ = fi103
                _ = v104
            }
            if true {
                for fi108 := 0; fi108 < 17; fi108++ {
                    v66 = v66
                    if false {
                        if true {
                            for fi109 := 0; fi109 < 14; fi109++ {
                                var v110 uint8 = 2
                                var v111 uint64 = 2
                                var v112 uint = 42
                                v113 := []int{(fi109 + 42), fi108}
                                _ = fi109
                                _ = v110
                                _ = v111
                                _ = v112
                                _ = v113
                            }
                            var v114 uint64 = 18446744073709551615
                            _ = v114
                        }
                        if (42 == 408) {
                            var v115 uint8 = 242
                            _ = v115
                        } else {
                            for fi116 := 0; fi116 < 11; fi116++ {
                                v117 := GenS13{F0: "hello", F1: ""}
                                switch v69 {
                                case 9223372036854775807:
                                    var v118 uint32 = genFn0("\n\t", ("\n\t" + "\n\t"), (v66 & v66))
                                    _ = v118
                                case 127:
                                    var v119 bool = false
                                    _ = v119
                                default:
                                    v120 := GenS6{F0: 4294967295, F1: 127, F2: -367}
                                    _ = v120
                                }
                                _ = fi116
                                _ = v117
                            }
                        }
                    } else {
                        if true {
                            _ = len(v3)
                            v121 := v3[""]
                            v122 := v3["hello"]
                            _ = v121
                            _ = v122
                        } else {
                            var v123 float32 = float32((float64(3.14159) + v2))
                            if true {
                                _ = len(v3)
                                for fi124 := 0; fi124 < 14; fi124++ {
                                    var v125 int64 = -682
                                    switch "a" {
                                    case "a":
                                        v123 = (0.0 / -2.25)
                                    default:
                                        fmt.Println(v69, fi108)
                                    }
                                    var v126 int = fi124
                                    _ = fi124
                                    _ = v125
                                    _ = v126
                                }
                            } else {
                                fmt.Println(v2, v2, v2)
                                for fi127 := 0; fi127 < 2; fi127++ {
                                    v2 = 0.1
                                    fmt.Println(fi127, v1)
                                    _ = fi127
                                }
                            }
                            _ = v123
                        }
                    }
                    switch int(v1) {
                    case 42:
                        var v128 float32 = 3.14159
                        _ = v128
                    }
                    _ = fi108
                }
                if (v2 > v2) {
                    v76 += v76
                    v1 = v1
                } else {
                    if false {
                        fmt.Println(v76, v1)
                        switch "a" {
                        case "foo bar":
                            var v129 int32 = (v66 - v66)
                            _ = v129
                        case "":
                            fmt.Println(v76)
                        case "\n\t":
                            fmt.Println(v2, v1, v76)
                        default:
                            v130 := GenS11{F0: 935, F1: int8(v69), F2: 32767}
                            _ = v130
                        }
                        if false {
                            var v131 uint32 = 496
                            _ = v131
                        }
                    }
                }
            }
            var v132 int16 = 1000
            _ = v132
        } else {
            v133 := []int{(^(v69))}
            _ = v133
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v66:", v66)
    fmt.Println("v69:", v69)
    fmt.Println("v76:", v76)
    _ = v1
    _ = v2
    _ = v3
    _ = v66
    _ = v67
    _ = v68
    _ = v69
    _ = v75
    _ = v76
}
