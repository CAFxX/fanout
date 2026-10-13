package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 string; F2 uint8 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 string; F1 int32 }
type GenS5 struct { F0 int8; F1 uint8 }
type GenS6 struct { F0 int; F1 uint16; F2 bool }
type GenS7 struct { F0 uint8 }
type GenS8 struct { F0 int8 }
type GenS9 struct { F0 int }
type GenS10 struct { F0 uint32 }
type GenS11 struct { F0 int8; F1 uint8 }
type GenS12 struct { F0 uint16 }
type GenS13 struct { F0 uint16 }
type GenS14 struct { F0 uint8; F1 int8 }
func genFn0(p0 uint) uint8 {
    v7 := []int{391, 128, -9223372036854775808, -2147483648}
    for fi8 := 0; fi8 < 5; fi8++ {
        for fi9 := 0; fi9 < 16; fi9++ {
            switch (-(fi9)) {
            case 42:
                p0 -= 4294967295
            case 396:
                p0 -= p0
            case -787:
                p0 -= (^(p0))
            default:
                fmt.Println(p0)
            }
            p0 -= p0
            switch ("foo bar" + "hello") {
            case "":
                p0 -= 752
            case "foo bar":
                if false {
                    p0 -= p0
                    p0 += p0
                    var v10 uint16 = 2
                    switch fi9 {
                    case 42:
                        v11 := GenS1{F0: 42, F1: "foo bar", F2: 2}
                        _ = v11
                    case -814:
                        if (4294967295 <= 424) {
                            if len(v7) > 0 {
                                v7[0] = (fi9 * (fi8 - 42))
                            }
                        }
                    default:
                        v10 = (v10 % ((v10) | 1))
                    }
                    v12 := GenS1{F0: 340, F1: ("" + "hello"), F2: 255}
                    _ = v10
                    _ = v12
                }
            default:
                var v13 int = 9223372036854775807
                _ = v13
            }
            fmt.Println(p0, fi8)
            _ = fi9
        }
        if (false || false) {
            if (2147483647 != -1) {
                v7 = append(v7, fi8)
                var v14 uint16 = 2
                _ = v14
            }
            if false {
                var v15 string = ""
                _ = v15
            } else {
                switch -9223372036854775808 {
                case -564:
                    var v16 bool = true
                    _ = v16
                case 1:
                    switch "" {
                    case "hello":
                        for fi17 := 0; fi17 < 7; fi17++ {
                            p0 -= 94
                            var v18 bool = ((false || false) && true)
                            _ = fi17
                            _ = v18
                        }
                    case "a":
                        for fi19 := 0; fi19 < 13; fi19++ {
                            v20 := GenS1{F0: -1, F1: ("" + (("\n\t" + "\n\t") + ("hello" + "a"))), F2: 1}
                            if false {
                                if len(v7) > 0 {
                                    v7[0] = -628
                                }
                            } else {
                                v21 := GenS2{F0: 42}
                                var v22 float64 = (0.0 - 0.0)
                                if false {
                                    v22 += v22
                                    fmt.Println(fi19, p0)
                                } else {
                                    var v23 int32 = 1
                                    v24 := GenS1{F0: (v23 * -2147483648), F1: "\n\t", F2: 255}
                                    _ = v23
                                    _ = v24
                                }
                                _ = v21
                                _ = v22
                            }
                            var v25 string = "hello"
                            v26 := GenS3{F0: fi8}
                            _ = fi19
                            _ = v20
                            _ = v25
                            _ = v26
                        }
                    case "foo bar":
                        p0 -= (^(p0))
                    }
                }
                var v27 int8 = -77
                p0 = p0
                _ = v27
            }
        }
        _ = fi8
    }
    fmt.Println(p0)
    var v28 uint = p0
    _ = p0
    _ = v28
    _ = v7
    return 2
}
func genFn1(p0 int8, p1 uint8) int32 {
    for fi29 := 0; fi29 < 9; fi29++ {
        var v30 int32 = -1
        var v31 int = (fi29 - 4294967295)
        if true {
            fmt.Println(p1, p0, p0)
            switch ((("a" + "") + ("" + "\n\t")) + ("foo bar" + ("\n\t" + "foo bar"))) {
            case "hello":
                var v32 int = v31
                _ = v32
            case "a":
                var v33 int = v31
                _ = v33
            case "\n\t":
                v34 := GenS3{F0: (v31 >> ((702) & 63))}
                _ = v34
            default:
                v35 := []int{}
                _ = v35
            }
            v36 := []int{v31, 4294967295}
            _ = v36
        } else {
            var v37 int64 = 626
            _ = v37
        }
        _ = fi29
        _ = v30
        _ = v31
    }
    fmt.Println(p1, p0)
    v38 := []int{}
    var v39 int = 42
    var v40 uint32 = uint32(p1)
    _ = p0
    _ = p1
    _ = v38
    _ = v39
    _ = v40
    return -2147483648
}

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[uint32]uint32)
    v3 := []string{("foo bar" + "")}
    var v4 bool = false
    var v5 bool = v4
    v6 := make(map[string]bool)
    for fi41 := 0; fi41 < 11; fi41++ {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = fi41
    }
    v42 := GenS2{F0: 81}
    var v43 uint16 = 291
    if v5 {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    } else {
        switch 9223372036854775807 {
        case 32767:
            var v44 uint = 18446744073709551615
            _ = v44
        case 0:
            var v45 int16 = -32768
            _ = v45
        case 306:
            v4 = ("" == "foo bar")
        default:
            fmt.Println(v43, v43)
        }
        for fi46 := 0; fi46 < 11; fi46++ {
            if v4 {
                v47 := GenS4{F0: "hello", F1: 401}
                if (int32(fi46) >= 2147483647) {
                    v43 += v43
                    v3 = append(v3, "")
                }
                v48 := GenS1{F0: 2147483647, F1: "", F2: 96}
                _ = v47
                _ = v48
            } else {
                v43 += (v43 & v43)
                for fi49 := 0; fi49 < 19; fi49++ {
                    switch ("foo bar" + ("hello" + "foo bar")) {
                    case "a":
                        var v50 string = ("" + "\n\t")
                        _ = v50
                    default:
                        var v51 float32 = 3.14159
                        _ = v51
                    }
                    fmt.Println(v43, v5, v4)
                    v4 = false
                    _ = fi49
                }
            }
            for fi52 := 0; fi52 < 16; fi52++ {
                _ = len(v6)
                _ = fi52
            }
            v4 = v5
            _ = fi46
        }
        v53 := GenS2{F0: 1}
        _ = v53
    }
    for fi54 := 0; fi54 < 16; fi54++ {
        if (uint(v43) == 299) {
            var v55 int = (fi54 / ((fi54) | 1))
            _ = v55
        }
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        delete(v2, 2)
        _ = fi54
    }
    for fi56 := 0; fi56 < 4; fi56++ {
        if (fi56 != 0) {
            switch (fi56 & (fi56 >> ((4294967295) & 63))) {
            case -2147483648:
                v57 := GenS1{F0: genFn1(-128, 42), F1: (("\n\t" + "foo bar") + "foo bar"), F2: 45}
                _ = v57
            case 42:
                for fi58 := 0; fi58 < 3; fi58++ {
                    if ((v4 && (902 >= -667)) || (42 <= 255)) {
                        v59 := GenS5{F0: -120, F1: 174}
                        v60 := GenS1{F0: 42, F1: "\n\t", F2: 107}
                        v61 := GenS6{F0: fi58, F1: (v43 << ((483) & 63)), F2: v4}
                        _ = v59
                        _ = v60
                        _ = v61
                    }
                    var v62 int64 = 0
                    _ = fi58
                    _ = v62
                }
            }
            if (32767 == -183) {
                v5 = (2147483647 != 2147483647)
            } else {
                switch fi56 {
                case -515:
                    for fi63 := 0; fi63 < 6; fi63++ {
                        if ((0 == 0) || (150 <= 1)) {
                            v43 += (v43 - v43)
                            var v64 bool = true
                            var v65 uint = 2
                            var v66 int32 = 1
                            _ = v64
                            _ = v65
                            _ = v66
                        }
                        _ = fi63
                    }
                case 127:
                    v5 = v5
                default:
                    v43 += v43
                }
                v4 = false
                var v67 float64 = (3.14159 * (1e10 + float64((-2.25 - 1.5))))
                _ = v67
            }
        } else {
            v5 = v4
            var v68 string = ("hello" + "hello")
            var v69 int8 = 0
            _ = v68
            _ = v69
        }
        _ = fi56
    }
    switch "\n\t" {
    case "\n\t":
        if (1 != 0) {
            switch -495 {
            case 9223372036854775807:
                v70 := v2[0]
                _ = v70
            case -844:
                v71 := GenS7{F0: 128}
                _ = v71
            case 4294967295:
                if len(v3) > 0 {
                    v3[0] = "hello"
                }
            }
            if (32767 < -385) {
                switch 128 {
                case 2147483647:
                    v72 := v6["hello"]
                    _ = v72
                case -9223372036854775808:
                    var v73 bool = v4
                    _ = v73
                case 127:
                    v74 := GenS8{F0: 100}
                    _ = v74
                }
                if ((1.5 < (float64(-2.25) / 3.14159)) || v4) {
                    switch (("foo bar" + "\n\t") + "\n\t") {
                    case "a":
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                    case "":
                        if v4 {
                            for fi75 := 0; fi75 < 15; fi75++ {
                                v43 = (^(v43))
                                _ = fi75
                            }
                            v3 = append(v3, (("" + "a") + (("\n\t" + ("" + ("\n\t" + "a"))) + "foo bar")))
                        }
                    default:
                        var v76 bool = true
                        _ = v76
                    }
                    v77 := GenS9{F0: 9223372036854775807}
                    _ = v77
                } else {
                    if len(v3) > 0 {
                        v3[0] = (("\n\t" + (("hello" + "a") + ("" + "a"))) + "\n\t")
                    }
                }
            } else {
                v5 = (1e10 <= 0.1)
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            }
        } else {
            if v5 {
                for fi78 := 0; fi78 < 13; fi78++ {
                    v79 := v2[2]
                    switch fi78 {
                    case -2147483648:
                        v80 := v1[fi78]
                        _ = v80
                    case 42:
                        for fi81 := 0; fi81 < 19; fi81++ {
                            v82 := GenS1{F0: -797, F1: "hello", F2: uint8(v79)}
                            var v83 int16 = -242
                            _ = fi81
                            _ = v82
                            _ = v83
                        }
                    default:
                        if (4294967295 <= 1) {
                            for fi84 := 0; fi84 < 16; fi84++ {
                                v4 = (false || ((((-2.25 - (1e10 + (1.5 + ((1.5 / 1.0) + (3.14159 * 3.14159))))) / 0.1) / 1.0) != (1e10 * 0.1)))
                                delete(v1, fi78)
                                _ = fi84
                            }
                            if v4 {
                                if (1 == v43) {
                                    v79 = 65535
                                } else {
                                    v6[("a" + ("hello" + "foo bar"))] = (-1 == 29)
                                    v85 := GenS8{F0: -1}
                                    _ = v85
                                }
                                for fi86 := 0; fi86 < 7; fi86++ {
                                    var v87 uint16 = (v43 / ((2) | 1))
                                    if len(v3) > 0 {
                                        v3[0] = ""
                                    }
                                    fmt.Println(v4, v87, v87)
                                    _ = fi86
                                    _ = v87
                                }
                            }
                        } else {
                            v88 := GenS10{F0: (^(v79))}
                            _ = v88
                        }
                    }
                    _ = fi78
                    _ = v79
                }
            }
        }
    }
    var v89 float64 = -2.25
    for fi90 := 0; fi90 < 16; fi90++ {
        v91 := GenS5{F0: -19, F1: 0}
        var v92 uint16 = v43
        var v93 int32 = -872
        _ = fi90
        _ = v91
        _ = v92
        _ = v93
    }
    v43 = (v43 & v43)
    v43 = 42
    if v5 {
        var v94 int64 = -260
        if (v94 >= v94) {
            var v95 int16 = 437
            _ = v95
        } else {
            var v96 int16 = 42
            _ = v96
        }
        for fi97 := 0; fi97 < 9; fi97++ {
            var v98 int64 = -1
            v99 := GenS11{F0: 42, F1: 1}
            _ = fi97
            _ = v98
            _ = v99
        }
        _ = v94
    }
    if len(v3) > 0 {
        v3[0] = "hello"
    }
    if v5 {
        v43 -= (v43 & (^(v43)))
        v5 = v5
    }
    if ((^(v43)) <= (v43 - v43)) {
        _ = len(v6)
    } else {
        if (v4 && false) {
            if (2 >= 2147483647) {
                v100 := GenS12{F0: v43}
                var v101 int64 = 359
                v4 = true
                _ = v100
                _ = v101
            }
        } else {
            for fi102 := 0; fi102 < 8; fi102++ {
                for fi103 := 0; fi103 < 2; fi103++ {
                    v5 = v5
                    _ = fi103
                }
                _ = fi102
            }
        }
        if len(v3) > 0 {
            v3[0] = ""
        }
    }
    switch -9223372036854775808 {
    case -646:
        var v104 int16 = 1000
        _ = v104
    case -9223372036854775808:
        v105 := GenS13{F0: (v43 / ((2) | 1))}
        _ = v105
    case -2147483648:
        _ = len(v1)
    }
    for fi106 := 0; fi106 < 2; fi106++ {
        v107 := GenS14{F0: 2, F1: 64}
        var v108 float32 = (((3.14159 * -2.25) - 0.1) * 3.14159)
        _ = fi106
        _ = v107
        _ = v108
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v89:", v89)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v43
    _ = v5
    _ = v6
    _ = v89
}
