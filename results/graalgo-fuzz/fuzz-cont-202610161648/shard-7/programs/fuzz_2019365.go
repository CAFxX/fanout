package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32 }
type GenS2 struct { F0 int32; F1 int32; F2 uint8 }
type GenS3 struct { F0 int32; F1 int64 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 int8; F1 uint; F2 int }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 uint; F1 uint8; F2 uint32 }
type GenS8 struct { F0 uint32; F1 uint64 }
func genFn0() uint8 {
    v5 := GenS1{F0: 0, F1: 0}
    v6 := []int{42, 2147483647, 4294967295, 32767}
    v7 := GenS1{F0: 42, F1: 2147483647}
    _ = v5
    _ = v6
    _ = v7
    return 127
}
func genFn1(p0 uint16, p1 uint64, p2 int) uint {
    v8 := []int{(p2 << ((641) & 63)), p2, p2}
    var v9 uint32 = 55
    p0 = p0
    p1 -= p1
    p1 = (p1 - (p1 * (p1 & p1)))
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    _ = v9
    return 255
}

func main() {
    var v1 bool = false
    var v2 uint64 = 18446744073709551615
    v3 := make(map[string]int)
    var v4 uint = 4294967295
    switch ((("hello" + "foo bar") + (("foo bar" + "") + ("foo bar" + (("" + "foo bar") + "\n\t")))) + "hello") {
    case "":
        var v10 int = 0
        _ = v10
    case "\n\t":
        fmt.Println(v1, v1, v2)
    case "foo bar":
        var v11 uint16 = 1
        _ = v11
    }
    switch -659 {
    case 42:
        for fi12 := 0; fi12 < 3; fi12++ {
            var v13 int32 = 871
            var v14 int64 = -384
            var v15 int64 = (v14 / ((-9223372036854775808) | 1))
            _ = fi12
            _ = v13
            _ = v14
            _ = v15
        }
    case -1:
        if ((false || v1) || v1) {
            if (!((((-2.25 - -2.25) - 0.1) + 0.1) < float32(1.5))) {
                switch -736 {
                case 127:
                    if (42 == -535) {
                        for fi16 := 0; fi16 < 15; fi16++ {
                            fmt.Println(fi16, v4, v2)
                            _ = fi16
                        }
                    } else {
                        if ((v4 >> (((v4 & 253)) & 63)) <= genFn1(332, v2, -2147483648)) {
                            if v1 {
                                for fi17 := 0; fi17 < 11; fi17++ {
                                    for fi18 := 0; fi18 < 18; fi18++ {
                                        v19 := []int{fi18, -89, 933, fi17}
                                        _ = len(v19)
                                        v20 := GenS1{F0: 42, F1: 714}
                                        if (!(-235 != 253)) {
                                            var v21 bool = false
                                            v1 = v1
                                            v4 -= v4
                                            _ = v21
                                        }
                                        fmt.Println(v4, v1)
                                        _ = fi18
                                        _ = v19
                                        _ = v20
                                    }
                                    if v1 {
                                        v4 -= 807
                                        fmt.Println(v2, v2)
                                    }
                                    _ = fi17
                                }
                                switch "foo bar" {
                                case "foo bar":
                                    var v22 uint64 = (v2 / (((v2 - (^(v2)))) | 1))
                                    _ = v22
                                case "":
                                    switch "foo bar" {
                                    case "foo bar":
                                        for fi23 := 0; fi23 < 16; fi23++ {
                                            if (true || ((^(v4)) != v4)) {
                                                v24 := []int{-45, fi23}
                                                switch ("" + ("a" + ((("\n\t" + "\n\t") + ("hello" + "\n\t")) + (("\n\t" + ("" + ("hello" + "\n\t"))) + (("hello" + ("a" + ("a" + "a"))) + (((("foo bar" + "a") + "") + ("\n\t" + "")) + ("a" + "hello"))))))) {
                                                case "":
                                                    for fi25 := 0; fi25 < 19; fi25++ {
                                                        if (32767 <= 327) {
                                                            var v26 int16 = 1000
                                                            v27 := GenS1{F0: 19, F1: 921}
                                                            _ = v26
                                                            _ = v27
                                                        }
                                                        var v28 float64 = -2.25
                                                        _ = fi25
                                                        _ = v28
                                                    }
                                                case "foo bar":
                                                    if v1 {
                                                        var v29 bool = false
                                                        fmt.Println(v29, v1, v29)
                                                        _ = v29
                                                    } else {
                                                        if (102 != 42) {
                                                            if v1 {
                                                                switch "hello" {
                                                                case "hello":
                                                                    v3["hello"] = fi23
                                                                case "":
                                                                    for fi30 := 0; fi30 < 10; fi30++ {
                                                                        v2 -= 18446744073709551615
                                                                        _ = fi30
                                                                    }
                                                                case "foo bar":
                                                                    var v31 int8 = 127
                                                                    _ = v31
                                                                }
                                                                v32 := GenS1{F0: 2, F1: 0}
                                                                fmt.Println(v1, v1, v4)
                                                                _ = v32
                                                            }
                                                            v4 -= v4
                                                            v33 := GenS1{F0: 65535, F1: 65535}
                                                            _ = v33
                                                        } else {
                                                            var v34 float64 = float64((-2.25 + 1.5))
                                                            var v35 uint32 = 65535
                                                            _ = len(v24)
                                                            _ = v34
                                                            _ = v35
                                                        }
                                                        v1 = false
                                                        var v36 int16 = 1000
                                                        _ = v36
                                                    }
                                                default:
                                                    var v37 uint8 = genFn0()
                                                    _ = v37
                                                }
                                                _ = v24
                                            }
                                            var v38 uint64 = v2
                                            var v39 uint64 = (^(v2))
                                            _ = fi23
                                            _ = v38
                                            _ = v39
                                        }
                                    case "a":
                                        fmt.Println(v4, v1, v2)
                                    case "\n\t":
                                        fmt.Println(v1, v2, v2)
                                    }
                                }
                                delete(v3, ("hello" + ((("foo bar" + "a") + "a") + "foo bar")))
                            } else {
                                var v40 int32 = 305
                                for fi41 := 0; fi41 < 11; fi41++ {
                                    for fi42 := 0; fi42 < 19; fi42++ {
                                        if v1 {
                                            v4 += genFn1(186, 1, fi41)
                                            switch (fi42 - fi42) {
                                            case -9223372036854775808:
                                                v1 = (v4 == 65535)
                                            case 128:
                                                v3[((("hello" + "") + "a") + ("a" + "a"))] = 0
                                            case 32768:
                                                var v43 uint = (^(v4))
                                                _ = v43
                                            }
                                            v44 := GenS2{F0: 2147483647, F1: -48, F2: uint8(v2)}
                                            _ = v44
                                        } else {
                                            v2 -= v2
                                        }
                                        _ = fi42
                                    }
                                    var v45 float64 = 1e10
                                    _ = fi41
                                    _ = v45
                                }
                                _ = v40
                            }
                            if ((-149 > 210) || v1) {
                                switch -766 {
                                case 42:
                                    v46 := GenS2{F0: 2147483647, F1: 1, F2: 0}
                                    _ = v46
                                }
                                switch 32767 {
                                case 32767:
                                    var v47 int32 = 0
                                    _ = v47
                                }
                            }
                        }
                    }
                }
            }
            _ = len(v3)
        }
    }
    if (1 == -1) {
        v48 := GenS2{F0: 0, F1: 562, F2: 127}
        _ = v48
    } else {
        v49 := []int{-851, 128}
        v1 = (false && v1)
        _ = v49
    }
    v50 := GenS1{F0: 0, F1: 953}
    var v51 float64 = 3.14159
    switch 128 {
    case 4294967295:
        v1 = ((42 >= 255) && (2147483648 == 2))
    case 32767:
        v52 := []int{32768, 0, 864}
        _ = v52
    case 32768:
        fmt.Println(v4, v2, v4)
    default:
        if true {
            v53 := []int{32767}
            _ = v53
        } else {
            var v54 uint64 = (v2 << ((v4) & 63))
            _ = v54
        }
    }
    if (v1 && v1) {
        v55 := GenS1{F0: 255, F1: 2147483648}
        if (v1 || (!(4294967295 <= (v4 << ((v4) & 63))))) {
            v2 -= v2
            v56 := []int{}
            switch 0 {
            case 4294967295:
                var v57 uint = (^(v4))
                _ = v57
            }
            _ = v56
        } else {
            switch 1 {
            case 4294967295:
                v58 := GenS1{F0: 0, F1: 726}
                _ = v58
            case 439:
                var v59 string = ("hello" + "")
                _ = v59
            case -2147483648:
                for fi60 := 0; fi60 < 11; fi60++ {
                    v51 -= ((float64(1.5) + v51) / (((float64(0.0) / (((((0.1 - (v51 - 0.1)) * (-2.25 - v51)) + (v51 * 1.5)) * ((v51 * v51) * v51)) - v51)) + v51) - v51))
                    var v61 uint = v4
                    _ = fi60
                    _ = v61
                }
            default:
                var v62 int8 = 123
                _ = v62
            }
        }
        var v63 uint16 = 632
        _ = v55
        _ = v63
    } else {
        var v64 uint16 = 2
        for fi65 := 0; fi65 < 1; fi65++ {
            var v66 uint8 = 1
            if v1 {
                v67 := GenS3{F0: 2147483647, F1: 765}
                _ = v67
            }
            _ = fi65
            _ = v66
        }
        _ = v64
    }
    v68 := GenS3{F0: -1, F1: -1}
    var v69 uint32 = 4294967295
    if true {
        for fi70 := 0; fi70 < 4; fi70++ {
            if (v2 != (v2 / ((v2) | 1))) {
                for fi71 := 0; fi71 < 13; fi71++ {
                    if (!(v1 || (635 <= int64(v69)))) {
                        switch ("a" + "") {
                        case "":
                            if v1 {
                                v72 := []int{}
                                fmt.Println(fi70, v1)
                                v1 = (!v1)
                                _ = v72
                            }
                        case "\n\t":
                            if (-735 <= 1) {
                                var v73 float64 = 3.14159
                                v3[("\n\t" + ("" + "hello"))] = fi71
                                if v1 {
                                    switch -1 {
                                    case 9223372036854775807:
                                        fmt.Println(v69, v73)
                                    case 344:
                                        switch "" {
                                        case "":
                                            v74 := GenS4{F0: 1}
                                            _ = v74
                                        case "a":
                                            switch -9223372036854775808 {
                                            case 264:
                                                var v75 bool = ((v2 % (((v2 % ((642) | 1))) | 1)) != (v2 & (v2 | uint64(v69))))
                                                _ = v75
                                            case 32768:
                                                var v76 int64 = 42
                                                _ = v76
                                            default:
                                                v77 := GenS5{F0: -1, F1: 4294967295, F2: 2147483647}
                                                _ = v77
                                            }
                                        case "foo bar":
                                            v4 = (v4 ^ genFn1(0, (v2 | uint64(v4)), fi70))
                                        default:
                                            v78 := v3[(("\n\t" + "\n\t") + "\n\t")]
                                            _ = v78
                                        }
                                    }
                                    v79 := v3["a"]
                                    var v80 uint = 4294967295
                                    v81 := GenS4{F0: 555}
                                    v82 := GenS4{F0: 2}
                                    _ = v79
                                    _ = v80
                                    _ = v81
                                    _ = v82
                                }
                                _ = v73
                            } else {
                                v83 := []int{fi70, (fi71 % (((fi71 - fi70)) | 1))}
                                v84 := GenS2{F0: -2147483648, F1: 446, F2: 42}
                                _ = v83
                                _ = v84
                            }
                        }
                        var v85 uint64 = 18446744073709551615
                        v1 = v1
                        v4 -= v4
                        v86 := GenS5{F0: 1, F1: 42, F2: (fi70 - -3)}
                        var v87 float32 = (((((3.14159 / 3.14159) + float32(v51)) / 1.5) - (1.5 / 1.5)) / -2.25)
                        _ = v85
                        _ = v86
                        _ = v87
                    } else {
                        if ((v69 ^ (^(v69))) < (^(v69))) {
                            v3["foo bar"] = (fi70 & fi71)
                            if true {
                                delete(v3, "\n\t")
                                v88 := GenS4{F0: 255}
                                _ = v88
                            }
                            switch fi70 {
                            case -9223372036854775808:
                                v2 -= v2
                            default:
                                v2 = 1
                            }
                        }
                        v69 = (v69 * (v69 - v69))
                        var v89 int64 = -541
                        _ = v89
                    }
                    v2 = v2
                    v90 := GenS1{F0: 1, F1: (v69 >> (((v4 % ((v4) | 1))) & 63))}
                    _ = fi71
                    _ = v90
                }
                fmt.Println(v2, fi70)
            } else {
                v69 = v69
                _ = len(v3)
            }
            v2 -= (v2 << (((v4 & (v4 + (v4 ^ (v4 - v4))))) & 63))
            switch fi70 {
            case -9223372036854775808:
                fmt.Println(v69, fi70)
            case 42:
                switch "\n\t" {
                case "foo bar":
                    for fi91 := 0; fi91 < 1; fi91++ {
                        v51 = v51
                        v92 := GenS6{F0: (v4 / (((v4 << ((42) & 63))) | 1))}
                        if (!true) {
                            v1 = ("hello" <= "a")
                            v4 = (v4 * (v4 | genFn1(65535, v2, (fi70 - fi91))))
                            v69 -= (v69 << ((genFn1(2, (v2 & (v2 + (v2 * 1))), fi91)) & 63))
                            switch ("" + "hello") {
                            case "\n\t":
                                var v93 int = (fi70 - fi70)
                                _ = v93
                            case "":
                                delete(v3, "hello")
                            default:
                                switch ((((("a" + ((("a" + ("\n\t" + "foo bar")) + "foo bar") + ("\n\t" + ("\n\t" + "foo bar")))) + ("" + "\n\t")) + "a") + "a") + "") {
                                case "a":
                                    v94 := GenS2{F0: 0, F1: -363, F2: genFn0()}
                                    _ = v94
                                case "foo bar":
                                    var v95 float32 = -2.25
                                    _ = v95
                                }
                            }
                        } else {
                            fmt.Println(v2)
                            var v96 bool = v1
                            switch ("foo bar" + "\n\t") {
                            case "\n\t":
                                var v97 uint16 = 18
                                switch "hello" {
                                case "\n\t":
                                    fmt.Println(fi91, v51, v2)
                                case "a":
                                    v1 = false
                                case "":
                                    v1 = (0 > fi91)
                                default:
                                    v98 := GenS3{F0: -881, F1: -9223372036854775808}
                                    _ = v98
                                }
                                v99 := []int{-2147483648, 2147483647}
                                v2 = (v2 - uint64(v97))
                                _ = v97
                                _ = v99
                            case "a":
                                if v96 {
                                    _ = len(v3)
                                    v2 = v2
                                    switch (^(fi70)) {
                                    case 42:
                                        var v100 bool = (3.14159 != 1.5)
                                        _ = v100
                                    case 128:
                                        v101 := GenS5{F0: 1, F1: v4, F2: (fi70 & (-(fi70)))}
                                        _ = v101
                                    }
                                    v1 = (32767 != 0)
                                    var v102 bool = (!v96)
                                    v103 := []int{}
                                    if false {
                                        var v104 bool = v102
                                        _ = v104
                                    }
                                    v103 = append(v103, (fi70 + (fi70 | fi70)))
                                    _ = v102
                                    _ = v103
                                } else {
                                    if v96 {
                                        v96 = true
                                        fmt.Println(fi91, v4)
                                    } else {
                                        v105 := []int{fi70, (fi91 - -1), fi70, fi91}
                                        _ = v105
                                    }
                                    v51 = 0.1
                                }
                            }
                            _ = v96
                        }
                        _ = fi91
                        _ = v92
                    }
                }
            case 1:
                var v106 bool = (!v1)
                _ = v106
            }
            _ = fi70
        }
        var v107 int16 = 32767
        v51 = v51
        _ = v107
    } else {
        switch 9223372036854775807 {
        case -68:
            v1 = (42 != 65535)
        case -2147483648:
            v108 := []int{4294967295, 51, -1, -9223372036854775808}
            _ = v108
        case 267:
            fmt.Println(v2)
        default:
            v2 -= v2
        }
        var v109 int8 = -128
        v110 := v3["a"]
        _ = v109
        _ = v110
    }
    if (-32768 > 1) {
        v1 = (1 <= 72)
        fmt.Println(v1, v51, v2)
        v111 := []int{}
        _ = v111
    } else {
        for fi112 := 0; fi112 < 20; fi112++ {
            var v113 uint32 = v69
            _ = fi112
            _ = v113
        }
        v114 := GenS6{F0: (v4 << ((14) & 63))}
        v115 := []int{32767, -2147483648}
        _ = v114
        _ = v115
    }
    for fi116 := 0; fi116 < 15; fi116++ {
        var v117 int = fi116
        v118 := GenS7{F0: (v4 + v4), F1: 0, F2: v69}
        _ = fi116
        _ = v117
        _ = v118
    }
    var v119 uint = genFn1(1, v2, -1)
    var v120 float64 = v51
    var v121 int16 = -663
    if false {
        for fi122 := 0; fi122 < 19; fi122++ {
            v119 = (v4 / (((v4 >> ((65535) & 63))) | 1))
            if v1 {
                v123 := GenS8{F0: v69, F1: 690}
                _ = v123
            }
            _ = fi122
        }
        v121 -= (v121 ^ 1)
        switch -2147483648 {
        case 2147483647:
            if v1 {
                var v124 uint8 = uint8(v69)
                _ = v124
            } else {
                v125 := []int{}
                _ = v125
            }
        default:
            switch "a" {
            case "a":
                var v126 uint16 = 32767
                _ = v126
            case "\n\t":
                for fi127 := 0; fi127 < 20; fi127++ {
                    v128 := []int{(fi127 >> ((65535) & 63))}
                    if (!v1) {
                        v51 = (float64(-2.25) - 0.1)
                    } else {
                        for fi129 := 0; fi129 < 13; fi129++ {
                            var v130 uint = (v119 & v4)
                            if (!(((false || (!(596 == -2147483648))) && v1) && v1)) {
                                v120 += v51
                                v69 -= 1
                            }
                            if (42 < -2147483648) {
                                var v131 int16 = v121
                                v4 = v4
                                _ = v131
                            }
                            _ = fi129
                            _ = v130
                        }
                        v51 = v51
                        if v1 {
                            v132 := GenS7{F0: (v119 * (v119 & genFn1(478, v2, (fi127 + (fi127 - fi127))))), F1: uint8(v69), F2: v69}
                            v133 := v3["a"]
                            v51 += v120
                            _ = v132
                            _ = v133
                        }
                    }
                    _ = fi127
                    _ = v128
                }
            case "hello":
                switch 2147483647 {
                case 4294967295:
                    v119 += uint(v69)
                case -2147483648:
                    delete(v3, "a")
                default:
                    v134 := []int{2147483647, 0, 2147483647}
                    _ = v134
                }
            }
        }
    } else {
        _ = len(v3)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v119:", v119)
    fmt.Println("v120:", v120)
    fmt.Println("v121:", v121)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v51:", v51)
    fmt.Println("v69:", v69)
    _ = v1
    _ = v119
    _ = v120
    _ = v121
    _ = v2
    _ = v3
    _ = v4
    _ = v50
    _ = v51
    _ = v68
    _ = v69
}
