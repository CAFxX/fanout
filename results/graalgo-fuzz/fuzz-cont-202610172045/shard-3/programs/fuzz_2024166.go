package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32 }
type GenS2 struct { F0 uint32; F1 uint16; F2 int64 }
type GenS3 struct { F0 uint8; F1 uint16 }
type GenS4 struct { F0 int32; F1 int16; F2 string }
type GenS5 struct { F0 string; F1 uint16 }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 int8 }
type GenS8 struct { F0 int32; F1 int16; F2 uint32 }
type GenS9 struct { F0 uint16; F1 int32 }
type GenS10 struct { F0 int64; F1 int8 }
type GenS11 struct { F0 int8 }
type GenS12 struct { F0 int8 }
type GenS13 struct { F0 uint32 }
type GenS14 struct { F0 uint32; F1 int }
type GenS15 struct { F0 int64 }
type GenS16 struct { F0 int64; F1 int32 }
type GenS17 struct { F0 int64; F1 string }
type GenS18 struct { F0 string }
type GenS19 struct { F0 uint32; F1 int64; F2 int16 }
type GenS20 struct { F0 uint; F1 int32; F2 string }
type GenS21 struct { F0 int }
type GenS22 struct { F0 int64 }
type GenS23 struct { F0 int32; F1 int32 }
type GenS24 struct { F0 int16; F1 int8 }
type GenS25 struct { F0 int; F1 int8; F2 uint16 }
type GenS26 struct { F0 int8; F1 int32; F2 string }
type GenS27 struct { F0 int64; F1 uint16 }
type GenS28 struct { F0 uint8; F1 uint64 }
type GenS29 struct { F0 uint16; F1 uint16; F2 uint }
type GenS30 struct { F0 uint32 }
type GenS31 struct { F0 uint64; F1 int64 }
type GenS32 struct { F0 uint }
type GenS33 struct { F0 uint8; F1 string; F2 uint8 }
type GenS34 struct { F0 int8; F1 uint32 }
type GenS35 struct { F0 int; F1 uint64 }
type GenS36 struct { F0 int16; F1 int64 }
type GenS37 struct { F0 uint16 }
type GenS38 struct { F0 int32 }
type GenS39 struct { F0 uint16; F1 int; F2 uint8 }
type GenS40 struct { F0 int64 }
type GenS41 struct { F0 uint; F1 int }
type GenS42 struct { F0 int32 }
type GenS43 struct { F0 int64; F1 int8 }
type GenS44 struct { F0 uint }
type GenS45 struct { F0 int8; F1 bool }
type GenS46 struct { F0 int32 }
type GenS47 struct { F0 uint16 }
func genFn0(p0 int16, p1 int) uint16 {
    v6 := []int{-1, (p1 + p1), (p1 >> ((4294967295) & 63))}
    var v7 int32 = 0
    p1 = (-(p1))
    if true {
        v7 -= v7
    } else {
        v8 := GenS1{F0: "a", F1: 710}
        _ = v8
    }
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 920
}
func genFn1(p0 int8) uint16 {
    v9 := GenS2{F0: 0, F1: genFn0(-1, -301), F2: 535}
    if (0 <= -32768) {
        if (808 <= 255) {
            fmt.Println(p0)
            p0 = (^(p0))
        }
        if (!true) {
            var v10 int = 32767
            _ = v10
        } else {
            p0 = p0
            for fi11 := 0; fi11 < 2; fi11++ {
                p0 = (p0 / (((p0 % ((-86) | 1))) | 1))
                p0 -= (p0 % (((-(p0))) | 1))
                _ = fi11
            }
        }
    } else {
        v12 := []int{-9223372036854775808, -751, 4294967295, -1}
        _ = v12
    }
    v13 := []int{}
    v14 := GenS1{F0: "foo bar", F1: 65535}
    var v15 int8 = (p0 / ((1) | 1))
    _ = p0
    _ = v13
    _ = v14
    _ = v15
    _ = v9
    return 611
}

func main() {
    var v1 uint64 = 0
    var v2 string = ("hello" + "foo bar")
    var v3 uint16 = 32767
    v4 := make(map[string]string)
    v5 := make(map[string]string)
    v16 := v5["foo bar"]
    if true {
        v17 := []int{}
        for fi18 := 0; fi18 < 2; fi18++ {
            for fi19 := 0; fi19 < 8; fi19++ {
                switch (v2 + v16) {
                case "foo bar":
                    v20 := GenS1{F0: v16, F1: 0}
                    _ = v20
                }
                _ = fi19
            }
            _ = fi18
        }
        _ = v17
    }
    v1 -= v1
    v21 := v5[(("a" + (v2 + ("hello" + v16))) + "a")]
    switch v2 {
    case "hello":
        switch -461 {
        case 128:
            v1 = (v1 | v1)
        case -693:
            switch 4294967295 {
            case 42:
                for fi22 := 0; fi22 < 14; fi22++ {
                    v23 := []int{(fi22 * (fi22 / ((fi22) | 1))), 127, fi22, fi22}
                    for fi24 := 0; fi24 < 15; fi24++ {
                        v21 = v21
                        fmt.Println(fi22, fi22)
                        _ = fi24
                    }
                    v25 := GenS2{F0: 2147483648, F1: 1, F2: -9223372036854775808}
                    _ = fi22
                    _ = v23
                    _ = v25
                }
            case 850:
                if true {
                    for fi26 := 0; fi26 < 5; fi26++ {
                        v16 = v2
                        switch ("foo bar" + "") {
                        case "hello":
                            if false {
                                switch v2 {
                                case "hello":
                                    for fi27 := 0; fi27 < 15; fi27++ {
                                        switch fi27 {
                                        case 2147483647:
                                            fmt.Println(v1)
                                        case 576:
                                            if ((!(false && false)) || (-1 >= -1)) {
                                                switch fi27 {
                                                case 0:
                                                    v28 := GenS2{F0: 218, F1: v3, F2: -1}
                                                    _ = v28
                                                case 4294967295:
                                                    fmt.Println(fi26, v2)
                                                case 42:
                                                    fmt.Println(fi27, v1, fi27)
                                                }
                                                var v29 int16 = -197
                                                _ = v29
                                            } else {
                                                var v30 uint32 = 42
                                                v31 := []int{fi27, 4294967295, 32768}
                                                _ = v30
                                                _ = v31
                                            }
                                        case 9223372036854775807:
                                            switch (fi27 + fi27) {
                                            case 1:
                                                var v32 int8 = -97
                                                _ = v32
                                            }
                                        }
                                        _ = fi27
                                    }
                                case "":
                                    switch 2147483647 {
                                    case 2147483647:
                                        for fi33 := 0; fi33 < 15; fi33++ {
                                            if (((float64(0.0) + -2.25) - 0.0) == ((1e10 - 1.5) - float64(1.5))) {
                                                if (103 >= 995) {
                                                    delete(v5, "")
                                                }
                                            } else {
                                                v16 = v21
                                            }
                                            v16 = (v21 + ("hello" + ("\n\t" + "a")))
                                            v34 := []int{fi26, fi26, fi33}
                                            _ = fi33
                                            _ = v34
                                        }
                                    case 0:
                                        var v35 int = 586
                                        _ = v35
                                    default:
                                        for fi36 := 0; fi36 < 2; fi36++ {
                                            v37 := GenS3{F0: 127, F1: (v3 % (((v3 & (v3 - v3))) | 1))}
                                            var v38 float64 = 1e10
                                            v3 += genFn1(int8(fi26))
                                            _ = fi36
                                            _ = v37
                                            _ = v38
                                        }
                                    }
                                case "\n\t":
                                    var v39 int16 = -32768
                                    _ = v39
                                }
                                v21 = v21
                            } else {
                                if true {
                                    var v40 int32 = 2147483647
                                    var v41 uint8 = 2
                                    _ = v40
                                    _ = v41
                                }
                            }
                        case "foo bar":
                            var v42 int64 = 0
                            _ = v42
                        default:
                            fmt.Println(v21)
                        }
                        _ = fi26
                    }
                } else {
                    v3 += 665
                    switch 2147483647 {
                    case 32767:
                        for fi43 := 0; fi43 < 20; fi43++ {
                            if true {
                                var v44 int64 = -9223372036854775808
                                delete(v5, v2)
                                _ = v44
                            } else {
                                v45 := v5[(v21 + v2)]
                                _ = v45
                            }
                            for fi46 := 0; fi46 < 7; fi46++ {
                                var v47 int = -1
                                if false {
                                    var v48 int16 = 1000
                                    if false {
                                        var v49 string = v2
                                        if ((v1 | 0) <= v1) {
                                            fmt.Println(v48)
                                        } else {
                                            v50 := GenS4{F0: 1, F1: v48, F2: v21}
                                            _ = v50
                                        }
                                        var v51 uint64 = 1
                                        _ = v49
                                        _ = v51
                                    }
                                    if (!(100 < 42)) {
                                        v52 := GenS1{F0: v2, F1: 42}
                                        if false {
                                            fmt.Println(v2)
                                            if true {
                                                v53 := []int{fi43, v47}
                                                _ = len(v53)
                                                v54 := GenS2{F0: 0, F1: genFn0(32767, fi43), F2: int64(v3)}
                                                _ = v53
                                                _ = v54
                                            } else {
                                                v16 = (("foo bar" + v16) + "foo bar")
                                                delete(v4, (("\n\t" + "\n\t") + v21))
                                                v55 := GenS1{F0: v21, F1: 227}
                                                _ = v55
                                            }
                                            fmt.Println(v16, v16)
                                        } else {
                                            fmt.Println(fi43, v1, v2)
                                            var v56 uint32 = 65535
                                            switch (fi43 & (fi46 ^ fi43)) {
                                            case -9223372036854775808:
                                                v3 -= (^(v3))
                                            case -2147483648:
                                                v1 -= (^(v1))
                                            default:
                                                switch v2 {
                                                case "hello":
                                                    var v57 uint16 = v3
                                                    v57 = (^(v3))
                                                    v58 := []int{(v47 / ((32768) | 1)), (fi46 + 1), -748, 719}
                                                    _ = v57
                                                    _ = v58
                                                case "":
                                                    switch v21 {
                                                    case "\n\t":
                                                        fmt.Println(v48, v1, v2)
                                                    case "a":
                                                        var v59 uint64 = 4294967295
                                                        _ = v59
                                                    }
                                                    fmt.Println(v56)
                                                    if false {
                                                        fmt.Println(fi43, v56)
                                                        if true {
                                                            v60 := GenS3{F0: 1, F1: genFn0((v48 / ((1) | 1)), fi43)}
                                                            var v61 uint = 876
                                                            _ = v60
                                                            _ = v61
                                                        }
                                                        if true {
                                                            var v62 float32 = ((1e10 - float32((1.5 - float64((1.5 - -2.25))))) + ((0.0 + float32(-2.25)) / (-2.25 / (0.0 + -2.25))))
                                                            v63 := GenS2{F0: v56, F1: v3, F2: -9223372036854775808}
                                                            var v64 int16 = v48
                                                            v65 := GenS5{F0: v21, F1: (v3 & (v3 - (v3 ^ (v3 ^ genFn0(-993, v47)))))}
                                                            _ = v62
                                                            _ = v63
                                                            _ = v64
                                                            _ = v65
                                                        }
                                                    } else {
                                                        var v66 uint64 = 42
                                                        if false {
                                                            v67 := GenS6{F0: 0}
                                                            var v68 string = ("" + v2)
                                                            v1 -= (v66 / ((v1) | 1))
                                                            if (!(0 > 42)) {
                                                                var v69 int32 = 572
                                                                v70 := GenS7{F0: 59}
                                                                if ((1e10 - -2.25) == 0.0) {
                                                                    if true {
                                                                        var v71 uint64 = (v1 % (((v66 - (v66 ^ v1))) | 1))
                                                                        if (false || false) {
                                                                            v72 := []int{v47, fi46}
                                                                            _ = v72
                                                                        }
                                                                        v73 := []int{v47, v47, fi43, int(v3)}
                                                                        _ = v71
                                                                        _ = v73
                                                                    } else {
                                                                        if (86 >= 255) {
                                                                            if false {
                                                                                v69 = (-(v69))
                                                                                v16 = v16
                                                                                if ((^(v69)) < v69) {
                                                                                    var v74 float32 = 1.5
                                                                                    v75 := GenS8{F0: v69, F1: (-(v48)), F2: (v56 >> ((2) & 63))}
                                                                                    v69 = v69
                                                                                    _ = v74
                                                                                    _ = v75
                                                                                } else {
                                                                                    if ((v69 * v69) <= -2147483648) {
                                                                                        switch fi46 {
                                                                                        case 9223372036854775807:
                                                                                            v76 := []int{fi46, 32768, -9223372036854775808, fi43}
                                                                                            _ = v76
                                                                                        default:
                                                                                            var v77 uint8 = 1
                                                                                            _ = v77
                                                                                        }
                                                                                        v78 := []int{int(v69), v47, (^(fi43)), 2147483647}
                                                                                        _ = v78
                                                                                    } else {
                                                                                        v79 := GenS4{F0: (v69 ^ v69), F1: v48, F2: ((((v21 + "a") + "a") + v68) + v16)}
                                                                                        v68 = v21
                                                                                        _ = len(v4)
                                                                                        if false {
                                                                                            v66 = (^(v66))
                                                                                            if true {
                                                                                                v2 = "\n\t"
                                                                                            }
                                                                                            v80 := GenS8{F0: v69, F1: (v48 * v48), F2: 0}
                                                                                            v81 := []int{}
                                                                                            _ = v80
                                                                                            _ = v81
                                                                                        } else {
                                                                                            fmt.Println(v66, fi46, v47)
                                                                                            if false {
                                                                                                v82 := GenS9{F0: (^(v3)), F1: (v69 ^ (v69 * v69))}
                                                                                                _ = v82
                                                                                            }
                                                                                            v83 := []int{(fi43 % (((fi46 >> ((255) & 63))) | 1)), 42, 32768}
                                                                                            v48 -= 889
                                                                                            v84 := GenS7{F0: -1}
                                                                                            _ = v83
                                                                                            _ = v84
                                                                                        }
                                                                                        _ = v79
                                                                                    }
                                                                                    var v85 int8 = -64
                                                                                    _ = v85
                                                                                }
                                                                            }
                                                                            if true {
                                                                                switch (v68 + (("\n\t" + ("\n\t" + ((v21 + (("\n\t" + v68) + v16)) + "a"))) + v2)) {
                                                                                case "foo bar":
                                                                                    switch 128 {
                                                                                    case -9223372036854775808:
                                                                                        v21 = v2
                                                                                    default:
                                                                                        v47 += fi46
                                                                                    }
                                                                                case "\n\t":
                                                                                    v86 := GenS10{F0: -821, F1: 100}
                                                                                    _ = v86
                                                                                case "hello":
                                                                                    var v87 float32 = 1.5
                                                                                    _ = v87
                                                                                default:
                                                                                    v21 = v68
                                                                                }
                                                                                fmt.Println(v69, v66)
                                                                                var v88 uint64 = v1
                                                                                _ = v88
                                                                            }
                                                                            v89 := GenS11{F0: -113}
                                                                            fmt.Println(v56, v47)
                                                                            var v90 int32 = int32(v56)
                                                                            _ = v89
                                                                            _ = v90
                                                                        }
                                                                        if ((v66 ^ 42) <= 4294967295) {
                                                                            v21 = "foo bar"
                                                                            switch (fi46 << ((234) & 63)) {
                                                                            case 4294967295:
                                                                                var v91 int32 = (v69 + 0)
                                                                                _ = v91
                                                                            case 128:
                                                                                fmt.Println(v21, v68, v56)
                                                                            default:
                                                                                if true {
                                                                                    var v92 uint16 = v3
                                                                                    switch fi46 {
                                                                                    case 32768:
                                                                                        if (false && true) {
                                                                                            var v93 bool = (v66 > (v1 - v1))
                                                                                            _ = v93
                                                                                        } else {
                                                                                            v68 = (v16 + v21)
                                                                                            v21 = ""
                                                                                        }
                                                                                        fmt.Println(fi43)
                                                                                        v94 := []int{(-(fi46)), -9223372036854775808}
                                                                                        v95 := GenS12{F0: 100}
                                                                                        _ = v94
                                                                                        _ = v95
                                                                                    case 32767:
                                                                                        v96 := GenS8{F0: (v69 | v69), F1: v48, F2: 773}
                                                                                        _ = v96
                                                                                    case 128:
                                                                                        var v97 int32 = v69
                                                                                        _ = v97
                                                                                    default:
                                                                                        var v98 float64 = 0.1
                                                                                        _ = v98
                                                                                    }
                                                                                    _ = v92
                                                                                } else {
                                                                                    _ = len(v4)
                                                                                    var v99 int = fi43
                                                                                    _ = v99
                                                                                }
                                                                            }
                                                                            v100 := GenS13{F0: (v56 >> ((0) & 63))}
                                                                            v68 = v21
                                                                            if (128 <= 0) {
                                                                                var v101 int32 = (-(v69))
                                                                                v102 := []int{fi46, v47, 127, (v47 + fi46)}
                                                                                _ = v101
                                                                                _ = v102
                                                                            } else {
                                                                                var v103 int32 = int32(fi43)
                                                                                if ((false && false) && true) {
                                                                                    var v104 uint8 = 42
                                                                                    if (!true) {
                                                                                        var v105 float32 = ((3.14159 * -2.25) * float32((3.14159 + (0.1 + (1.5 - 1.5)))))
                                                                                        _ = v105
                                                                                    }
                                                                                    v106 := GenS14{F0: v56, F1: (fi43 + 0)}
                                                                                    _ = v104
                                                                                    _ = v106
                                                                                } else {
                                                                                    var v107 float64 = -2.25
                                                                                    var v108 bool = (false || (!false))
                                                                                    if (false || true) {
                                                                                        var v109 uint8 = 70
                                                                                        v103 = 0
                                                                                        _ = v109
                                                                                    }
                                                                                    _ = v107
                                                                                    _ = v108
                                                                                }
                                                                                _ = v103
                                                                            }
                                                                            _ = v100
                                                                        } else {
                                                                            v110 := []int{}
                                                                            delete(v5, v16)
                                                                            _ = v110
                                                                        }
                                                                    }
                                                                    var v111 uint = 42
                                                                    _ = v111
                                                                } else {
                                                                    if (9223372036854775807 > 666) {
                                                                        v112 := GenS14{F0: (v56 | (v56 / ((373) | 1))), F1: -504}
                                                                        if false {
                                                                            var v113 uint16 = v3
                                                                            var v114 string = v2
                                                                            v2 = ((v16 + ("a" + ((((v21 + (v2 + ("foo bar" + "a"))) + v2) + "") + "\n\t"))) + (v68 + ("a" + (v16 + v21))))
                                                                            _ = v113
                                                                            _ = v114
                                                                        } else {
                                                                            var v115 uint32 = 65535
                                                                            _ = v115
                                                                        }
                                                                        v116 := GenS15{F0: 1}
                                                                        _ = v112
                                                                        _ = v116
                                                                    }
                                                                }
                                                                _ = v69
                                                                _ = v70
                                                            } else {
                                                                v117 := GenS12{F0: int8(v56)}
                                                                switch v2 {
                                                                case "hello":
                                                                    var v118 uint = 55
                                                                    _ = v118
                                                                case "foo bar":
                                                                    v2 = ("" + v16)
                                                                case "\n\t":
                                                                    if (!false) {
                                                                        fmt.Println(v47, v66)
                                                                        switch fi46 {
                                                                        case 2147483647:
                                                                            v119 := GenS3{F0: 1, F1: (v3 ^ (v3 * (v3 / (((^(v3))) | 1))))}
                                                                            _ = v119
                                                                        }
                                                                        if (-589 > 9223372036854775807) {
                                                                            fmt.Println(v48)
                                                                            var v120 int16 = v48
                                                                            if false {
                                                                                v121 := GenS16{F0: -1, F1: 42}
                                                                                if true {
                                                                                    if false {
                                                                                        v122 := []int{(fi46 + 0), fi43, 32767}
                                                                                        var v123 uint32 = (v56 * 65535)
                                                                                        v48 -= (v120 - v120)
                                                                                        _ = v122
                                                                                        _ = v123
                                                                                    } else {
                                                                                        switch (v2 + "a") {
                                                                                        case "":
                                                                                            v5[v68] = v16
                                                                                        case "\n\t":
                                                                                            v1 -= 2
                                                                                        case "foo bar":
                                                                                            v16 = (v16 + v2)
                                                                                        }
                                                                                        if true {
                                                                                            fmt.Println(fi46, v48, fi46)
                                                                                        } else {
                                                                                            var v124 int64 = 0
                                                                                            _ = v124
                                                                                        }
                                                                                        v125 := []int{fi46, (v47 + fi46), (fi43 / (((v47 << ((328) & 63))) | 1))}
                                                                                        _ = v125
                                                                                    }
                                                                                    v120 = (v48 / ((-1) | 1))
                                                                                    var v126 int8 = -128
                                                                                    _ = v126
                                                                                }
                                                                                fmt.Println(v47, v2)
                                                                                _ = v121
                                                                            }
                                                                            _ = v120
                                                                        }
                                                                    } else {
                                                                        if true {
                                                                            switch v2 {
                                                                            case "":
                                                                                if (v16 == (("foo bar" + ((v68 + ("\n\t" + ("a" + v68))) + "foo bar")) + v21)) {
                                                                                    if (v68 == ("" + v2)) {
                                                                                        v127 := GenS17{F0: -1, F1: (v16 + "")}
                                                                                        if (1 != 42) {
                                                                                            if ((fi46 >= (v47 ^ fi43)) && false) {
                                                                                                v68 = v16
                                                                                            }
                                                                                            v128 := GenS3{F0: 42, F1: 440}
                                                                                            if false {
                                                                                                v3 += 2
                                                                                                if false {
                                                                                                    v129 := GenS14{F0: (v56 >> ((18446744073709551615) & 63)), F1: v47}
                                                                                                    fmt.Println(v47)
                                                                                                    var v130 int32 = -2147483648
                                                                                                    _ = v129
                                                                                                    _ = v130
                                                                                                } else {
                                                                                                    switch 355 {
                                                                                                    case 72:
                                                                                                        if (!(!true)) {
                                                                                                            v131 := GenS15{F0: -955}
                                                                                                            if (v3 >= (v3 % ((65535) | 1))) {
                                                                                                                v5[v68] = "hello"
                                                                                                            } else {
                                                                                                                var v132 int64 = int64(v47)
                                                                                                                _ = v132
                                                                                                            }
                                                                                                            _ = v131
                                                                                                        }
                                                                                                    }
                                                                                                    var v133 uint64 = 42
                                                                                                    fmt.Println(v1, v47, fi43)
                                                                                                    _ = v133
                                                                                                }
                                                                                                fmt.Println(v3, v1, v68)
                                                                                            } else {
                                                                                                var v134 uint32 = v56
                                                                                                v135 := GenS18{F0: ((v16 + "") + "")}
                                                                                                if true {
                                                                                                    switch fi46 {
                                                                                                    case 32768:
                                                                                                        v68 = "a"
                                                                                                    case 128:
                                                                                                        fmt.Println(v47, v21)
                                                                                                    case -9223372036854775808:
                                                                                                        switch (v21 + (v16 + (v68 + v68))) {
                                                                                                        case "foo bar":
                                                                                                            if (!(-6 <= 1)) {
                                                                                                                fmt.Println(fi46, v48, fi46)
                                                                                                                switch v47 {
                                                                                                                case -1:
                                                                                                                    v136 := GenS19{F0: v56, F1: 123, F2: 173}
                                                                                                                    _ = v136
                                                                                                                case 2147483647:
                                                                                                                    var v137 uint32 = v56
                                                                                                                    _ = v137
                                                                                                                case -2147483648:
                                                                                                                    switch 9223372036854775807 {
                                                                                                                    case 128:
                                                                                                                        v2 = v2
                                                                                                                    case -9223372036854775808:
                                                                                                                        v138 := GenS17{F0: 9223372036854775807, F1: (v68 + (v21 + "\n\t"))}
                                                                                                                        _ = v138
                                                                                                                    default:
                                                                                                                        if ("foo bar" != v16) {
                                                                                                                            if (1 < 255) {
                                                                                                                                v2 = v16
                                                                                                                            }
                                                                                                                            var v139 uint32 = (v134 / ((318) | 1))
                                                                                                                            _ = v139
                                                                                                                        }
                                                                                                                    }
                                                                                                                }
                                                                                                                switch (fi43 - fi46) {
                                                                                                                case -683:
                                                                                                                    var v140 uint16 = v3
                                                                                                                    _ = v140
                                                                                                                case 42:
                                                                                                                    v21 = "a"
                                                                                                                case -618:
                                                                                                                    switch (-(v47)) {
                                                                                                                    case 0:
                                                                                                                        v134 += v56
                                                                                                                    case -1:
                                                                                                                        v16 = ""
                                                                                                                    }
                                                                                                                default:
                                                                                                                    if false {
                                                                                                                        v16 = (v2 + v68)
                                                                                                                        v141 := GenS20{F0: 4294967295, F1: 42, F2: ""}
                                                                                                                        _ = v141
                                                                                                                    } else {
                                                                                                                        if false {
                                                                                                                            v142 := []int{}
                                                                                                                            var v143 float64 = 3.14159
                                                                                                                            _ = v142
                                                                                                                            _ = v143
                                                                                                                        } else {
                                                                                                                            var v144 int8 = -128
                                                                                                                            switch v2 {
                                                                                                                            case "hello":
                                                                                                                                v145 := GenS21{F0: fi46}
                                                                                                                                _ = v145
                                                                                                                            default:
                                                                                                                                fmt.Println(v16, v68)
                                                                                                                            }
                                                                                                                            var v146 int = v47
                                                                                                                            _ = v144
                                                                                                                            _ = v146
                                                                                                                        }
                                                                                                                        v147 := GenS14{F0: 578, F1: v47}
                                                                                                                        _ = v147
                                                                                                                    }
                                                                                                                }
                                                                                                            } else {
                                                                                                                v148 := GenS22{F0: 9223372036854775807}
                                                                                                                v149 := GenS23{F0: 536, F1: 2147483647}
                                                                                                                var v150 int16 = v48
                                                                                                                _ = v148
                                                                                                                _ = v149
                                                                                                                _ = v150
                                                                                                            }
                                                                                                        default:
                                                                                                            fmt.Println(fi46, v3, v1)
                                                                                                        }
                                                                                                    }
                                                                                                    _ = len(v5)
                                                                                                    v5[v21] = (v21 + (v68 + v68))
                                                                                                    if true {
                                                                                                        v151 := []int{}
                                                                                                        v3 = v3
                                                                                                        _ = v151
                                                                                                    }
                                                                                                } else {
                                                                                                    var v152 int64 = 0
                                                                                                    if (v152 <= (v152 / ((v152) | 1))) {
                                                                                                        var v153 uint16 = v3
                                                                                                        var v154 uint = 504
                                                                                                        v155 := GenS24{F0: v48, F1: 100}
                                                                                                        _ = v153
                                                                                                        _ = v154
                                                                                                        _ = v155
                                                                                                    } else {
                                                                                                        fmt.Println(v134)
                                                                                                        v156 := GenS25{F0: int(v152), F1: 25, F2: v3}
                                                                                                        v134 = (v56 / (((v134 << ((18446744073709551615) & 63))) | 1))
                                                                                                        _ = v156
                                                                                                    }
                                                                                                    switch v47 {
                                                                                                    case 844:
                                                                                                        fmt.Println(v2)
                                                                                                    case 730:
                                                                                                        fmt.Println(v47, v134, fi46)
                                                                                                    default:
                                                                                                        switch 2147483647 {
                                                                                                        case -1:
                                                                                                            switch 128 {
                                                                                                            case 932:
                                                                                                                switch (v68 + v68) {
                                                                                                                case "hello":
                                                                                                                    v134 -= v56
                                                                                                                    v2 = v16
                                                                                                                    v21 = v2
                                                                                                                case "a":
                                                                                                                    if ((v66 - (v66 | (v1 % ((v66) | 1)))) < uint64(v134)) {
                                                                                                                        v56 = v134
                                                                                                                        switch (^(v47)) {
                                                                                                                        case 193:
                                                                                                                            v47 -= (v47 * (fi43 << ((42) & 63)))
                                                                                                                        case 365:
                                                                                                                            v157 := GenS26{F0: 42, F1: 42, F2: v21}
                                                                                                                            _ = v157
                                                                                                                        }
                                                                                                                    }
                                                                                                                }
                                                                                                            case -990:
                                                                                                                var v158 float32 = -2.25
                                                                                                                _ = v158
                                                                                                            default:
                                                                                                                var v159 string = ("hello" + "")
                                                                                                                var v160 uint32 = 65535
                                                                                                                _ = v159
                                                                                                                _ = v160
                                                                                                            }
                                                                                                        case 9223372036854775807:
                                                                                                            v161 := GenS3{F0: 186, F1: v3}
                                                                                                            v4["hello"] = "\n\t"
                                                                                                            switch (fi46 << ((1) & 63)) {
                                                                                                            case -1:
                                                                                                                v5["foo bar"] = "foo bar"
                                                                                                            case 127:
                                                                                                                v5["a"] = (v68 + v68)
                                                                                                            case 32768:
                                                                                                                if (1e10 < -2.25) {
                                                                                                                    switch v16 {
                                                                                                                    case "\n\t":
                                                                                                                        v1 = v66
                                                                                                                    }
                                                                                                                    v68 = "hello"
                                                                                                                } else {
                                                                                                                    v3 = 2
                                                                                                                    v2 = (v2 + (v16 + v16))
                                                                                                                }
                                                                                                            default:
                                                                                                                if (13 == 94) {
                                                                                                                    if ((v152 | (v152 - v152)) >= (v152 | -1)) {
                                                                                                                        v1 = v66
                                                                                                                        var v162 int8 = 0
                                                                                                                        _ = v162
                                                                                                                    } else {
                                                                                                                        v66 += (v66 + 4294967295)
                                                                                                                    }
                                                                                                                } else {
                                                                                                                    v66 = 18446744073709551615
                                                                                                                    var v163 int16 = (-(v48))
                                                                                                                    v164 := []int{fi46, (v47 ^ (v47 % (((fi46 | 32767)) | 1))), (fi46 >> ((42) & 63))}
                                                                                                                    _ = v163
                                                                                                                    _ = v164
                                                                                                                }
                                                                                                                switch fi43 {
                                                                                                                case 127:
                                                                                                                    v21 = v21
                                                                                                                default:
                                                                                                                    v2 = "foo bar"
                                                                                                                }
                                                                                                                if (!false) {
                                                                                                                    fmt.Println(v1, v2, v134)
                                                                                                                    var v165 uint16 = 32767
                                                                                                                    _ = v165
                                                                                                                }
                                                                                                                var v166 int32 = -863
                                                                                                                if false {
                                                                                                                    v4["hello"] = v21
                                                                                                                    _ = len(v5)
                                                                                                                }
                                                                                                                _ = v166
                                                                                                            }
                                                                                                            v66 -= v66
                                                                                                            _ = v161
                                                                                                        case 478:
                                                                                                            v167 := GenS27{F0: (v152 + -1), F1: v3}
                                                                                                            _ = v167
                                                                                                        }
                                                                                                        delete(v4, (v21 + (v68 + v16)))
                                                                                                        switch fi43 {
                                                                                                        case 42:
                                                                                                            v168 := GenS15{F0: v152}
                                                                                                            v1 = 42
                                                                                                            v134 = (^(v134))
                                                                                                            _ = v168
                                                                                                        case -9223372036854775808:
                                                                                                            var v169 uint16 = v3
                                                                                                            _ = v169
                                                                                                        default:
                                                                                                            switch v16 {
                                                                                                            case "a":
                                                                                                                v68 = v68
                                                                                                            case "":
                                                                                                                var v170 uint64 = v1
                                                                                                                _ = v170
                                                                                                            default:
                                                                                                                var v171 int8 = -1
                                                                                                                _ = v171
                                                                                                            }
                                                                                                        }
                                                                                                    }
                                                                                                    _ = v152
                                                                                                }
                                                                                                _ = v134
                                                                                                _ = v135
                                                                                            }
                                                                                            _ = v128
                                                                                        } else {
                                                                                            v1 -= 42
                                                                                        }
                                                                                        _ = v127
                                                                                    }
                                                                                }
                                                                            case "a":
                                                                                var v172 float64 = 1.5
                                                                                _ = v172
                                                                            case "foo bar":
                                                                                fmt.Println(v16, v66, v66)
                                                                            default:
                                                                                v173 := GenS8{F0: -2147483648, F1: 42, F2: (v56 + (v56 * (v56 * (v56 & v56))))}
                                                                                v174 := []int{(^(fi43)), v47, fi46, 353}
                                                                                _ = v173
                                                                                _ = v174
                                                                            }
                                                                            v68 = ("a" + "hello")
                                                                            v175 := []int{fi46, 180, 32767}
                                                                            v176 := GenS7{F0: 0}
                                                                            _ = v175
                                                                            _ = v176
                                                                        }
                                                                        var v177 uint64 = v1
                                                                        if true {
                                                                            v68 = "hello"
                                                                            if true {
                                                                                fmt.Println(fi43, v48)
                                                                                v178 := []int{}
                                                                                _ = v178
                                                                            } else {
                                                                                var v179 uint16 = (v3 ^ v3)
                                                                                var v180 uint16 = genFn0((v48 >> ((42) & 63)), fi46)
                                                                                v181 := []int{(fi43 * fi43), (fi46 / ((fi43) | 1))}
                                                                                _ = v179
                                                                                _ = v180
                                                                                _ = v181
                                                                            }
                                                                            var v182 string = (v21 + ((v21 + v21) + "a"))
                                                                            _ = v182
                                                                        } else {
                                                                            v183 := v5[v16]
                                                                            _ = v183
                                                                        }
                                                                        _ = v177
                                                                    }
                                                                default:
                                                                    v184 := GenS28{F0: 42, F1: (v1 >> ((0) & 63))}
                                                                    v185 := GenS29{F0: genFn0((v48 << ((315) & 63)), (fi43 % ((v47) | 1))), F1: (^(v3)), F2: 1}
                                                                    fmt.Println(v16, v16)
                                                                    _ = v184
                                                                    _ = v185
                                                                }
                                                                switch v21 {
                                                                case "hello":
                                                                    v186 := GenS30{F0: v56}
                                                                    _ = v186
                                                                default:
                                                                    var v187 uint16 = v3
                                                                    _ = v187
                                                                }
                                                                _ = v117
                                                            }
                                                            v188 := []int{}
                                                            switch "\n\t" {
                                                            case "hello":
                                                                if false {
                                                                    var v189 uint64 = (v66 % ((0) | 1))
                                                                    _ = v189
                                                                } else {
                                                                    if true {
                                                                        if true {
                                                                            if false {
                                                                                switch "\n\t" {
                                                                                case "a":
                                                                                    v190 := GenS11{F0: 0}
                                                                                    _ = v190
                                                                                case "":
                                                                                    switch (fi46 + fi46) {
                                                                                    case 4294967295:
                                                                                        if true {
                                                                                            var v191 uint32 = (v56 * v56)
                                                                                            _ = v191
                                                                                        } else {
                                                                                            fmt.Println(v21, v21)
                                                                                        }
                                                                                    case 0:
                                                                                        if (((v56 + (v56 & v56)) != 967) || false) {
                                                                                            var v192 int16 = (^(v48))
                                                                                            _ = v192
                                                                                        }
                                                                                        v193 := GenS24{F0: (^(v48)), F1: -1}
                                                                                        fmt.Println(v16, v3)
                                                                                        _ = v193
                                                                                    case 621:
                                                                                        fmt.Println(v2, v68, v66)
                                                                                    }
                                                                                case "\n\t":
                                                                                    v21 = "hello"
                                                                                }
                                                                                switch v16 {
                                                                                case "hello":
                                                                                    v56 += v56
                                                                                case "foo bar":
                                                                                    v21 = v2
                                                                                case "a":
                                                                                    switch 2147483647 {
                                                                                    case 42:
                                                                                        if true {
                                                                                            v47 -= (fi43 + (fi46 * (-(fi46))))
                                                                                            v194 := GenS19{F0: (^(v56)), F1: 42, F2: v48}
                                                                                            v66 = (v66 - 780)
                                                                                            _ = v194
                                                                                        }
                                                                                    case 954:
                                                                                        v66 = (v1 ^ uint64(v3))
                                                                                    default:
                                                                                        var v195 float64 = (0.1 + 1e10)
                                                                                        if true {
                                                                                            v48 += (v48 << ((42) & 63))
                                                                                            if false {
                                                                                                v196 := GenS31{F0: v66, F1: -134}
                                                                                                _ = v196
                                                                                            }
                                                                                            if ((uint16(v1) == v3) && false) {
                                                                                                v56 = v56
                                                                                                _ = len(v188)
                                                                                            } else {
                                                                                                v197 := GenS8{F0: -2147483648, F1: (^(v48)), F2: v56}
                                                                                                if true {
                                                                                                    var v198 int16 = v48
                                                                                                    switch (((((v2 + v2) + v68) + v21) + v68) + "") {
                                                                                                    case "hello":
                                                                                                        v199 := GenS32{F0: 0}
                                                                                                        _ = v199
                                                                                                    default:
                                                                                                        switch ("foo bar" + "hello") {
                                                                                                        case "hello":
                                                                                                            v47 = v47
                                                                                                            switch "hello" {
                                                                                                            case "hello":
                                                                                                                v200 := GenS1{F0: v16, F1: 65535}
                                                                                                                _ = v200
                                                                                                            case "foo bar":
                                                                                                                v66 -= (v1 | v66)
                                                                                                            case "a":
                                                                                                                var v201 uint32 = (^(v56))
                                                                                                                _ = v201
                                                                                                            }
                                                                                                        case "\n\t":
                                                                                                            var v202 int = (v47 / ((v47) | 1))
                                                                                                            _ = v202
                                                                                                        default:
                                                                                                            if (false || false) {
                                                                                                                switch "a" {
                                                                                                                case "hello":
                                                                                                                    if (!(3.14159 <= 0.0)) {
                                                                                                                        v3 += (v3 ^ 749)
                                                                                                                        fmt.Println(v3, v1, v66)
                                                                                                                    } else {
                                                                                                                        fmt.Println(v2, v195)
                                                                                                                        var v203 uint64 = v66
                                                                                                                        switch v47 {
                                                                                                                        case 2147483647:
                                                                                                                            switch fi43 {
                                                                                                                            case -258:
                                                                                                                                _ = len(v188)
                                                                                                                            case 127:
                                                                                                                                v47 = (fi43 << ((1) & 63))
                                                                                                                            case 4294967295:
                                                                                                                                fmt.Println(v21, v195)
                                                                                                                            default:
                                                                                                                                v56 -= v56
                                                                                                                            }
                                                                                                                        case 128:
                                                                                                                            v204 := GenS7{F0: 71}
                                                                                                                            _ = v204
                                                                                                                        case 42:
                                                                                                                            fmt.Println(fi43)
                                                                                                                        default:
                                                                                                                            var v205 bool = false
                                                                                                                            _ = v205
                                                                                                                        }
                                                                                                                        _ = v203
                                                                                                                    }
                                                                                                                case "\n\t":
                                                                                                                    v195 = (v195 / v195)
                                                                                                                case "":
                                                                                                                    if len(v188) >= 2 {
                                                                                                                        v188 = v188[1:len(v188)]
                                                                                                                    }
                                                                                                                    v66 = v1
                                                                                                                    if (18446744073709551615 >= 42) {
                                                                                                                        fmt.Println(v195, v66, v3)
                                                                                                                    }
                                                                                                                }
                                                                                                                if (42 <= uint(v3)) {
                                                                                                                    var v206 uint16 = 255
                                                                                                                    _ = v206
                                                                                                                }
                                                                                                            }
                                                                                                        }
                                                                                                    }
                                                                                                    var v207 int64 = 9223372036854775807
                                                                                                    _ = v198
                                                                                                    _ = v207
                                                                                                }
                                                                                                _ = v197
                                                                                            }
                                                                                            if (34 > 100) {
                                                                                                if (false || (((((-2.25 <= 0.1) || (-54 <= 1)) && true) && false) || (746 > (v56 & (v56 % ((65535) | 1)))))) {
                                                                                                    fmt.Println(v195)
                                                                                                    v21 = ("foo bar" + "\n\t")
                                                                                                    _ = len(v4)
                                                                                                }
                                                                                                switch fi43 {
                                                                                                case 128:
                                                                                                    switch fi43 {
                                                                                                    case -587:
                                                                                                        if false {
                                                                                                            var v208 uint32 = v56
                                                                                                            var v209 uint8 = 127
                                                                                                            _ = v208
                                                                                                            _ = v209
                                                                                                        } else {
                                                                                                            var v210 int32 = -1
                                                                                                            v1 = v66
                                                                                                            if len(v188) > 0 {
                                                                                                                v188[0] = -661
                                                                                                            }
                                                                                                            _ = v210
                                                                                                        }
                                                                                                        v68 = ("foo bar" + "foo bar")
                                                                                                    default:
                                                                                                        switch v2 {
                                                                                                        case "foo bar":
                                                                                                            fmt.Println(v2)
                                                                                                        default:
                                                                                                            if true {
                                                                                                                if false {
                                                                                                                    if true {
                                                                                                                        var v211 uint32 = (v56 ^ (v56 * v56))
                                                                                                                        v212 := GenS33{F0: 127, F1: "a", F2: 0}
                                                                                                                        _ = v211
                                                                                                                        _ = v212
                                                                                                                    }
                                                                                                                    fmt.Println(v16, v56, v48)
                                                                                                                }
                                                                                                            }
                                                                                                        }
                                                                                                    }
                                                                                                }
                                                                                                var v213 string = (v2 + "")
                                                                                                delete(v5, (v68 + v21))
                                                                                                v48 += 32767
                                                                                                _ = v213
                                                                                            }
                                                                                            var v214 uint16 = (v3 << ((255) & 63))
                                                                                            v215 := v5["\n\t"]
                                                                                            delete(v4, v68)
                                                                                            _ = v214
                                                                                            _ = v215
                                                                                        }
                                                                                        _ = v195
                                                                                    }
                                                                                }
                                                                                var v216 int64 = 0
                                                                                _ = v216
                                                                            } else {
                                                                                v47 = v47
                                                                                var v217 float64 = 0.1
                                                                                if true {
                                                                                    v218 := GenS4{F0: 1, F1: (^(v48)), F2: (("\n\t" + (v16 + v2)) + "\n\t")}
                                                                                    var v219 int16 = (v48 / ((int16(v3)) | 1))
                                                                                    v219 = (-(v48))
                                                                                    if true {
                                                                                        var v220 int16 = (v48 / ((v219) | 1))
                                                                                        _ = v220
                                                                                    }
                                                                                    v221 := GenS33{F0: 15, F1: v68, F2: 128}
                                                                                    var v222 int8 = 1
                                                                                    switch v16 {
                                                                                    case "\n\t":
                                                                                        if (!(255 == 42)) {
                                                                                            var v223 uint = 65535
                                                                                            switch v68 {
                                                                                            case "":
                                                                                                fmt.Println(v2, v48, fi43)
                                                                                            }
                                                                                            _ = v223
                                                                                        }
                                                                                    case "hello":
                                                                                        var v224 int16 = (v48 % (((v219 >> ((0) & 63))) | 1))
                                                                                        _ = v224
                                                                                    case "a":
                                                                                        fmt.Println(v2, v68)
                                                                                    }
                                                                                    v68 = v16
                                                                                    v1 = 1
                                                                                    var v225 uint = 33
                                                                                    if true {
                                                                                        if true {
                                                                                            fmt.Println(v2)
                                                                                        }
                                                                                        v4[v16] = v21
                                                                                        v226 := GenS32{F0: (v225 - (v225 + (^(v225))))}
                                                                                        _ = v226
                                                                                    }
                                                                                    if false {
                                                                                        if false {
                                                                                            var v227 string = ((("\n\t" + v21) + "foo bar") + "\n\t")
                                                                                            switch v47 {
                                                                                            case 829:
                                                                                                v222 = (v222 << ((v225) & 63))
                                                                                            }
                                                                                            _ = v227
                                                                                        } else {
                                                                                            v68 = (v16 + v16)
                                                                                            switch (v16 + ((v21 + "foo bar") + v21)) {
                                                                                            case "foo bar":
                                                                                                delete(v4, v16)
                                                                                            default:
                                                                                                v217 += 3.14159
                                                                                            }
                                                                                        }
                                                                                        v188 = append(v188, (fi43 | (v47 + fi43)))
                                                                                    } else {
                                                                                        v217 -= v217
                                                                                    }
                                                                                    _ = v218
                                                                                    _ = v219
                                                                                    _ = v221
                                                                                    _ = v222
                                                                                    _ = v225
                                                                                } else {
                                                                                    var v228 float64 = 1.5
                                                                                    _ = v228
                                                                                }
                                                                                var v229 int32 = 659
                                                                                v47 = fi43
                                                                                switch ("\n\t" + "foo bar") {
                                                                                case "\n\t":
                                                                                    fmt.Println(v229, v16)
                                                                                    var v230 int8 = -128
                                                                                    _ = v230
                                                                                case "foo bar":
                                                                                    v231 := GenS34{F0: 1, F1: uint32(v66)}
                                                                                    var v232 uint64 = (^(v1))
                                                                                    v68 = v21
                                                                                    _ = v231
                                                                                    _ = v232
                                                                                case "":
                                                                                    switch 41 {
                                                                                    case 32768:
                                                                                        v2 = ("foo bar" + v21)
                                                                                    case 829:
                                                                                        if false {
                                                                                            if true {
                                                                                                if (1.5 <= 3.14159) {
                                                                                                    v233 := GenS35{F0: (fi43 - fi46), F1: (v1 ^ (v66 % ((v1) | 1)))}
                                                                                                    if true {
                                                                                                        switch v2 {
                                                                                                        case "":
                                                                                                            switch fi46 {
                                                                                                            case 1:
                                                                                                                var v234 uint = 255
                                                                                                                _ = v234
                                                                                                            }
                                                                                                        case "hello":
                                                                                                            var v235 uint8 = 157
                                                                                                            _ = v235
                                                                                                        case "\n\t":
                                                                                                            var v236 int16 = (v48 % (((v48 << ((0) & 63))) | 1))
                                                                                                            _ = v236
                                                                                                        default:
                                                                                                            v237 := GenS36{F0: (-(v48)), F1: -1}
                                                                                                            _ = v237
                                                                                                        }
                                                                                                        _ = len(v188)
                                                                                                    }
                                                                                                    switch "hello" {
                                                                                                    case "a":
                                                                                                        var v238 uint64 = v66
                                                                                                        _ = v238
                                                                                                    case "hello":
                                                                                                        if true {
                                                                                                            switch v68 {
                                                                                                            case "a":
                                                                                                                switch (v21 + "foo bar") {
                                                                                                                case "\n\t":
                                                                                                                    var v239 string = "hello"
                                                                                                                    _ = v239
                                                                                                                }
                                                                                                            case "hello":
                                                                                                                fmt.Println(v68, v68, v1)
                                                                                                            }
                                                                                                        } else {
                                                                                                            switch v16 {
                                                                                                            case "hello":
                                                                                                                var v240 string = "hello"
                                                                                                                _ = v240
                                                                                                            default:
                                                                                                                v21 = v16
                                                                                                            }
                                                                                                            fmt.Println(fi43)
                                                                                                            _ = len(v188)
                                                                                                        }
                                                                                                    }
                                                                                                    _ = v233
                                                                                                } else {
                                                                                                    switch v16 {
                                                                                                    case "\n\t":
                                                                                                        v66 = (v1 - (^(v66)))
                                                                                                    case "hello":
                                                                                                        var v241 int16 = v48
                                                                                                        _ = v241
                                                                                                    case "foo bar":
                                                                                                        fmt.Println(v2)
                                                                                                    }
                                                                                                    if true {
                                                                                                        if true {
                                                                                                            v66 -= 0
                                                                                                            v229 -= 0
                                                                                                            v5[v16] = v16
                                                                                                        } else {
                                                                                                            v16 = v16
                                                                                                            switch "a" {
                                                                                                            case "hello":
                                                                                                                if true {
                                                                                                                    v66 = (v1 | 18446744073709551615)
                                                                                                                } else {
                                                                                                                    fmt.Println(v1, v1, fi43)
                                                                                                                    var v242 float32 = float32(v217)
                                                                                                                    _ = v242
                                                                                                                }
                                                                                                            case "a":
                                                                                                                v243 := GenS12{F0: 0}
                                                                                                                _ = v243
                                                                                                            case "foo bar":
                                                                                                                v188 = append(v188, (fi46 * (v47 | 42)))
                                                                                                            }
                                                                                                        }
                                                                                                        var v244 uint64 = v1
                                                                                                        if false {
                                                                                                            fmt.Println(v229, v68, v66)
                                                                                                        } else {
                                                                                                            if true {
                                                                                                                switch 4294967295 {
                                                                                                                case 631:
                                                                                                                    v2 = ((v16 + v2) + v21)
                                                                                                                case 289:
                                                                                                                    if (18446744073709551615 == 4294967295) {
                                                                                                                        var v245 int8 = 100
                                                                                                                        v246 := GenS37{F0: 1}
                                                                                                                        _ = v245
                                                                                                                        _ = v246
                                                                                                                    }
                                                                                                                default:
                                                                                                                    fmt.Println(v244, v2, v48)
                                                                                                                }
                                                                                                            } else {
                                                                                                                var v247 int = (v47 % ((v47) | 1))
                                                                                                                v248 := GenS25{F0: int(v48), F1: -128, F2: (v3 + (v3 * v3))}
                                                                                                                if true {
                                                                                                                    var v249 string = ("\n\t" + v21)
                                                                                                                    var v250 float64 = 1.5
                                                                                                                    var v251 int32 = 364
                                                                                                                    _ = v249
                                                                                                                    _ = v250
                                                                                                                    _ = v251
                                                                                                                } else {
                                                                                                                    var v252 uint64 = (v244 >> ((914) & 63))
                                                                                                                    _ = v252
                                                                                                                }
                                                                                                                _ = v247
                                                                                                                _ = v248
                                                                                                            }
                                                                                                        }
                                                                                                        _ = v244
                                                                                                    }
                                                                                                    var v253 int64 = 0
                                                                                                    v3 = uint16(v66)
                                                                                                    v1 += (v1 | (v66 / ((949) | 1)))
                                                                                                    v2 = "hello"
                                                                                                    var v254 float64 = (v217 * float64(1.5))
                                                                                                    _ = v253
                                                                                                    _ = v254
                                                                                                }
                                                                                            }
                                                                                        }
                                                                                    case 0:
                                                                                        var v255 int16 = v48
                                                                                        _ = v255
                                                                                    }
                                                                                default:
                                                                                    v21 = v16
                                                                                }
                                                                                var v256 float32 = 1.5
                                                                                _ = v217
                                                                                _ = v229
                                                                                _ = v256
                                                                            }
                                                                            var v257 int16 = (v48 * (v48 << ((uint(v56)) & 63)))
                                                                            v56 = uint32(v1)
                                                                            _ = v257
                                                                        }
                                                                        switch "foo bar" {
                                                                        case "":
                                                                            v258 := GenS38{F0: -2147483648}
                                                                            _ = v258
                                                                        case "a":
                                                                            var v259 uint = 2
                                                                            _ = v259
                                                                        case "foo bar":
                                                                            v260 := GenS39{F0: (^(v3)), F1: (^(fi46)), F2: 42}
                                                                            _ = v260
                                                                        }
                                                                    }
                                                                    var v261 uint32 = v56
                                                                    switch ("hello" + (v68 + v2)) {
                                                                    case "\n\t":
                                                                        var v262 float64 = 0.0
                                                                        _ = v262
                                                                    case "foo bar":
                                                                        v2 = "\n\t"
                                                                    case "hello":
                                                                        fmt.Println(v2, v68, v3)
                                                                    default:
                                                                        fmt.Println(v56)
                                                                    }
                                                                    _ = v261
                                                                }
                                                                v21 = "hello"
                                                                switch 377 {
                                                                case 488:
                                                                    if true {
                                                                        var v263 uint8 = 1
                                                                        v264 := GenS13{F0: v56}
                                                                        _ = v263
                                                                        _ = v264
                                                                    }
                                                                case -821:
                                                                    v48 -= v48
                                                                }
                                                            }
                                                            v265 := GenS25{F0: v47, F1: -41, F2: 32767}
                                                            _ = v188
                                                            _ = v265
                                                            _ = v67
                                                            _ = v68
                                                        }
                                                        delete(v4, v16)
                                                        _ = v66
                                                    }
                                                case "foo bar":
                                                    v266 := []int{4294967295, (^(fi46))}
                                                    _ = v266
                                                }
                                                switch (((v21 + ("hello" + (((v2 + ("hello" + (v16 + v2))) + v2) + "hello"))) + v2) + v2) {
                                                case "\n\t":
                                                    v267 := v5[v16]
                                                    _ = v267
                                                }
                                                if (-128 > 42) {
                                                    v1 = uint64(v56)
                                                    var v268 uint8 = 255
                                                    _ = v268
                                                } else {
                                                    v21 = v16
                                                }
                                                switch ((v2 + "a") + v16) {
                                                case "\n\t":
                                                    v269 := GenS40{F0: 0}
                                                    _ = v269
                                                case "foo bar":
                                                    var v270 float64 = float64(1e10)
                                                    _ = v270
                                                }
                                            }
                                            _ = v56
                                        }
                                        _ = v52
                                    }
                                    v16 = ""
                                    v3 += (v3 - genFn1(int8(v3)))
                                    v271 := GenS41{F0: 100, F1: v47}
                                    _ = v271
                                    _ = v48
                                }
                                _ = fi46
                                _ = v47
                            }
                            for fi272 := 0; fi272 < 6; fi272++ {
                                fmt.Println(fi43, v2, v16)
                                var v273 uint16 = (v3 + 0)
                                v274 := []int{32767}
                                _ = fi272
                                _ = v273
                                _ = v274
                            }
                            _ = fi43
                        }
                    case 0:
                        if false {
                            v275 := []int{}
                            _ = v275
                        }
                    case -61:
                        if false {
                            switch v21 {
                            case "hello":
                                v16 = "a"
                            case "foo bar":
                                for fi276 := 0; fi276 < 19; fi276++ {
                                    v16 = v16
                                    var v277 uint = 1
                                    _ = fi276
                                    _ = v277
                                }
                            default:
                                delete(v4, (v2 + v21))
                            }
                        } else {
                            v2 = (v16 + v16)
                            if false {
                                var v278 uint16 = v3
                                fmt.Println(v3, v21, v3)
                                v279 := GenS4{F0: -413, F1: 0, F2: v2}
                                _ = v278
                                _ = v279
                            }
                        }
                    default:
                        var v280 uint8 = 127
                        _ = v280
                    }
                }
            case 2147483647:
                for fi281 := 0; fi281 < 3; fi281++ {
                    v2 = ((("foo bar" + "a") + ((v2 + (("" + "\n\t") + v2)) + "hello")) + "\n\t")
                    var v282 string = (v2 + "hello")
                    v282 = v282
                    _ = fi281
                    _ = v282
                }
            default:
                v21 = ""
            }
        }
    case "\n\t":
        var v283 int = 0
        _ = v283
    case "":
        var v284 float32 = (1.5 - 0.1)
        _ = v284
    default:
        var v285 uint = 65535
        _ = v285
    }
    if false {
        if true {
            v286 := []int{152, -570}
            v287 := GenS1{F0: "\n\t", F1: uint32(v3)}
            _ = v286
            _ = v287
        } else {
            for fi288 := 0; fi288 < 16; fi288++ {
                if (32767 == (v3 + genFn0(32767, fi288))) {
                    if (!(42 >= 0)) {
                        for fi289 := 0; fi289 < 9; fi289++ {
                            if (-9223372036854775808 < 1) {
                                if (false && true) {
                                    v290 := GenS24{F0: -952, F1: 42}
                                    if false {
                                        _ = len(v5)
                                    } else {
                                        delete(v5, (((("a" + v21) + ((v16 + ("hello" + "hello")) + ("" + ("a" + v2)))) + (("hello" + "") + "foo bar")) + (v21 + ("" + v21))))
                                    }
                                    var v291 int16 = 25
                                    _ = v290
                                    _ = v291
                                }
                                var v292 float32 = -2.25
                                fmt.Println(v16)
                                _ = v292
                            }
                            _ = fi289
                        }
                        v293 := v4[(("a" + v2) + v16)]
                        v294 := GenS42{F0: -123}
                        _ = v293
                        _ = v294
                    }
                    for fi295 := 0; fi295 < 14; fi295++ {
                        fmt.Println(v3)
                        v296 := GenS43{F0: 0, F1: -123}
                        if false {
                            switch v16 {
                            case "\n\t":
                                v297 := GenS11{F0: 1}
                                _ = v297
                            case "hello":
                                v3 -= 65535
                            case "":
                                if (1000 <= 153) {
                                    var v298 int16 = 32767
                                    _ = v298
                                } else {
                                    var v299 int = fi295
                                    v299 += (^(fi295))
                                    v300 := []int{127, v299}
                                    _ = v299
                                    _ = v300
                                }
                            }
                            var v301 uint64 = v1
                            _ = v301
                        } else {
                            _ = len(v4)
                            v2 = v21
                        }
                        _ = fi295
                        _ = v296
                    }
                    delete(v4, "\n\t")
                } else {
                    v3 -= genFn1(0)
                    if true {
                        var v302 int8 = 0
                        _ = v302
                    } else {
                        _ = len(v4)
                    }
                }
                var v303 string = ""
                _ = fi288
                _ = v303
            }
        }
        v304 := GenS37{F0: 2}
        fmt.Println(v1)
        _ = v304
    }
    var v305 uint = 255
    if ((-2.25 < 1e10) || false) {
        var v306 float32 = (float32(float64(0.0)) * 1.5)
        if false {
            var v307 int16 = 32767
            _ = v307
        }
        for fi308 := 0; fi308 < 9; fi308++ {
            v309 := GenS44{F0: v305}
            v310 := GenS45{F0: 1, F1: true}
            _ = fi308
            _ = v309
            _ = v310
        }
        _ = v306
    } else {
        switch 128 {
        case -9223372036854775808:
            if ((^(v1)) < v1) {
                var v311 int = -17
                for fi312 := 0; fi312 < 8; fi312++ {
                    for fi313 := 0; fi313 < 10; fi313++ {
                        v314 := GenS25{F0: -1, F1: -1, F2: 42}
                        _ = fi313
                        _ = v314
                    }
                    for fi315 := 0; fi315 < 13; fi315++ {
                        var v316 bool = false
                        _ = fi315
                        _ = v316
                    }
                    var v317 uint64 = v1
                    _ = fi312
                    _ = v317
                }
                if true {
                    switch v311 {
                    case -71:
                        v305 = (v305 + v305)
                    case 127:
                        v318 := GenS6{F0: 42}
                        _ = v318
                    case 236:
                        fmt.Println(v2, v305, v16)
                    }
                    fmt.Println(v1, v21, v2)
                } else {
                    if (-171 <= int32(v1)) {
                        v305 += v305
                    } else {
                        var v319 uint8 = 178
                        var v320 bool = true
                        _ = v319
                        _ = v320
                    }
                    v321 := GenS46{F0: 0}
                    v322 := GenS47{F0: 2}
                    _ = v321
                    _ = v322
                }
                _ = v311
            }
        case 630:
            if ((0.1 - float64(1.5)) != -2.25) {
                fmt.Println(v305, v305)
                v3 = v3
            } else {
                var v323 uint64 = (v1 + 42)
                var v324 float32 = 0.1
                _ = v323
                _ = v324
            }
        default:
            var v325 int16 = 421
            _ = v325
        }
        v1 = v1
        var v326 int = 128
        _ = v326
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v305:", v305)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v16
    _ = v2
    _ = v21
    _ = v3
    _ = v305
    _ = v4
    _ = v5
}
