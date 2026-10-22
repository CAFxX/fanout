package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint32 }
type GenS2 struct { F0 uint32; F1 bool; F2 bool }
type GenS3 struct { F0 int8; F1 uint; F2 uint }
type GenS4 struct { F0 uint64; F1 uint16; F2 int32 }
type GenS5 struct { F0 int64; F1 string }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 string; F1 int32 }
type GenS8 struct { F0 string; F1 int }
type GenS9 struct { F0 uint32; F1 uint16; F2 uint32 }
type GenS10 struct { F0 uint }
type GenS11 struct { F0 uint8; F1 uint64; F2 int16 }
type GenS12 struct { F0 string; F1 int64 }
func genFn0() int {
    var v4 uint32 = 4294967295
    var v5 uint16 = 2
    var v6 uint16 = v5
    _ = v4
    _ = v5
    _ = v6
    return 128
}
func genFn1(p0 uint, p1 uint32) uint16 {
    v7 := GenS1{F0: 0, F1: p1}
    var v8 int = 32768
    fmt.Println(p0, p1)
    if false {
        v8 = (v8 ^ (v8 & -437))
        v9 := []int{genFn0(), (v8 ^ (-(v8))), genFn0(), (v8 / ((v8) | 1))}
        for fi10 := 0; fi10 < 13; fi10++ {
            v11 := GenS1{F0: -2147483648, F1: (p1 | (^(p1)))}
            _ = fi10
            _ = v11
        }
        _ = v9
    } else {
        if false {
            for fi12 := 0; fi12 < 18; fi12++ {
                switch (v8 >> ((p0) & 63)) {
                case 4294967295:
                    var v13 int64 = 0
                    _ = v13
                case 32767:
                    switch (fi12 ^ (v8 ^ 557)) {
                    case 4294967295:
                        var v14 uint64 = 18446744073709551615
                        _ = v14
                    case 930:
                        v15 := []int{}
                        _ = v15
                    default:
                        v16 := GenS2{F0: p1, F1: false, F2: false}
                        _ = v16
                    }
                case 42:
                    if true {
                        v17 := []int{42}
                        if (18446744073709551615 > 42) {
                            if false {
                                v8 -= 2147483647
                            }
                        }
                        _ = v17
                    } else {
                        p1 = (^(p1))
                    }
                default:
                    var v18 uint8 = 231
                    _ = v18
                }
                if (414 >= (p0 / (((p0 | (p0 + p0))) | 1))) {
                    v19 := []int{(fi12 / ((32767) | 1)), (v8 - (v8 % ((genFn0()) | 1))), (v8 - (v8 ^ fi12)), (v8 & fi12)}
                    _ = v19
                } else {
                    p1 = p1
                }
                _ = fi12
            }
            v20 := GenS3{F0: 0, F1: 42, F2: 42}
            _ = v20
        }
        p1 = uint32(p0)
    }
    p0 = 42
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 65535
}

func main() {
    var v1 uint32 = 1
    var v2 uint16 = 524
    v3 := make(map[uint32]uint32)
    var v21 int8 = 35
    v2 = v2
    v22 := GenS4{F0: 99, F1: 138, F2: 2147483647}
    for fi23 := 0; fi23 < 3; fi23++ {
        v2 -= 42
        var v24 int64 = -1
        _ = fi23
        _ = v24
    }
    v1 -= (v1 & v1)
    v21 -= v21
    v25 := GenS5{F0: 347, F1: "foo bar"}
    if (false && false) {
        v1 += 1
        var v26 int64 = -1
        _ = v26
    }
    v1 = (v1 / ((v1) | 1))
    v2 -= (v2 % ((32767) | 1))
    if true {
        for fi27 := 0; fi27 < 19; fi27++ {
            v28 := GenS2{F0: (v1 % (((v1 - v1)) | 1)), F1: ((^(fi27)) < fi27), F2: (true || true)}
            _ = fi27
            _ = v28
        }
        for fi29 := 0; fi29 < 7; fi29++ {
            v2 = (v2 ^ genFn1(65535, 65535))
            _ = fi29
        }
        fmt.Println(v1, v21, v2)
    } else {
        switch 0 {
        case 32768:
            v21 = 127
        case 866:
            _ = len(v3)
        }
        for fi30 := 0; fi30 < 17; fi30++ {
            if (!(!true)) {
                var v31 int64 = -986
                _ = v31
            }
            _ = fi30
        }
        if (-1 == -2147483648) {
            for fi32 := 0; fi32 < 1; fi32++ {
                switch fi32 {
                case 484:
                    v21 = (v21 << ((4294967295) & 63))
                case 2147483647:
                    fmt.Println(fi32)
                case 1:
                    v2 = v2
                }
                for fi33 := 0; fi33 < 18; fi33++ {
                    v1 += 307
                    fmt.Println(v2)
                    if (float32(float64(1.5)) < 0.1) {
                        switch "" {
                        case "\n\t":
                            v2 = v2
                        case "a":
                            var v34 uint = 18446744073709551615
                            _ = v34
                        case "foo bar":
                            v2 = (v2 / ((genFn1(555, (v1 % (((^(v1))) | 1)))) | 1))
                            v1 = (v1 & 300)
                        default:
                            v2 = 1
                        }
                        if (0.1 < -2.25) {
                            var v35 int8 = (v21 % ((97) | 1))
                            switch "hello" {
                            case "\n\t":
                                v36 := GenS6{F0: 42}
                                _ = v36
                            case "a":
                                var v37 uint16 = v2
                                _ = v37
                            case "":
                                if false {
                                    var v38 int = (fi33 * fi32)
                                    if (226 != -781) {
                                        if (fi32 > fi32) {
                                            v38 += v38
                                            var v39 int8 = v35
                                            var v40 int = genFn0()
                                            v2 = genFn1(915, v1)
                                            if (!true) {
                                                var v41 int32 = 42
                                                var v42 bool = ((fi33 ^ (^(v40))) != (v38 << ((255) & 63)))
                                                v43 := v3[(v1 * 65535)]
                                                _ = v41
                                                _ = v42
                                                _ = v43
                                            } else {
                                                if (true && (genFn1(42, v1) < genFn1(942, (v1 | v1)))) {
                                                    v44 := v3[(^(v1))]
                                                    v45 := []int{fi33, fi33, (-(fi32)), (fi32 >> ((255) & 63))}
                                                    _ = v44
                                                    _ = v45
                                                }
                                                switch 32767 {
                                                case 127:
                                                    fmt.Println(v21, fi32, v2)
                                                case 332:
                                                    if false {
                                                        if (!true) {
                                                            var v46 int32 = 1
                                                            var v47 string = (("hello" + ("a" + ("\n\t" + ""))) + "foo bar")
                                                            _ = v46
                                                            _ = v47
                                                        }
                                                    } else {
                                                        switch "a" {
                                                        case "\n\t":
                                                            var v48 int8 = (v21 >> ((623) & 63))
                                                            _ = v48
                                                        case "hello":
                                                            if (((v39 + v21) == (v39 << ((65535) & 63))) && (!true)) {
                                                                var v49 int8 = 42
                                                                _ = v49
                                                            } else {
                                                                var v50 int8 = 0
                                                                var v51 float32 = 0.1
                                                                _ = v50
                                                                _ = v51
                                                            }
                                                        case "":
                                                            delete(v3, (v1 << ((42) & 63)))
                                                        default:
                                                            var v52 int64 = 1
                                                            _ = v52
                                                        }
                                                    }
                                                case -935:
                                                    var v53 int32 = 0
                                                    _ = v53
                                                }
                                            }
                                            _ = v39
                                            _ = v40
                                        }
                                        v1 = 4294967295
                                    } else {
                                        v54 := GenS7{F0: "a", F1: 42}
                                        var v55 float32 = 1e10
                                        fmt.Println(v35, v2)
                                        v56 := GenS8{F0: "\n\t", F1: fi32}
                                        switch genFn0() {
                                        case 9223372036854775807:
                                            if (!(false && true)) {
                                                if (v1 > (v1 / ((v1) | 1))) {
                                                    v2 = 1
                                                }
                                                v57 := []int{(fi33 & 2147483647), (fi32 + (fi32 ^ genFn0()))}
                                                if false {
                                                    fmt.Println(v35, v38)
                                                    v1 += 1
                                                    v21 = 1
                                                } else {
                                                    fmt.Println(v55, fi32)
                                                }
                                                _ = v57
                                            }
                                        case 0:
                                            var v58 string = "hello"
                                            _ = v58
                                        case 32767:
                                            v59 := GenS4{F0: 42, F1: genFn1(1, 65535), F2: 1}
                                            _ = v59
                                        default:
                                            v60 := v3[(v1 | v1)]
                                            _ = v60
                                        }
                                        _ = v54
                                        _ = v55
                                        _ = v56
                                    }
                                    if (!false) {
                                        switch ("hello" + "foo bar") {
                                        case "hello":
                                            if (2 < 255) {
                                                switch 92 {
                                                case -9223372036854775808:
                                                    v61 := []int{0, genFn0(), 127, genFn0()}
                                                    _ = v61
                                                case 127:
                                                    fmt.Println(v38, v2)
                                                case 207:
                                                    v1 = v1
                                                }
                                                v38 = (fi33 + (-(fi32)))
                                                var v62 int = (fi32 >> ((255) & 63))
                                                if ("" <= "a") {
                                                    var v63 uint16 = v2
                                                    var v64 uint = 647
                                                    _ = v63
                                                    _ = v64
                                                } else {
                                                    _ = len(v3)
                                                    if false {
                                                        v2 = genFn1(733, v1)
                                                        v65 := []int{fi32, v38, v62}
                                                        var v66 bool = (!(!false))
                                                        _ = v65
                                                        _ = v66
                                                    }
                                                    var v67 uint16 = genFn1(432, (v1 - (^(v1))))
                                                    _ = v67
                                                }
                                                _ = v62
                                            }
                                        }
                                        v68 := v3[2147483648]
                                        _ = v68
                                    } else {
                                        var v69 string = ("hello" + "foo bar")
                                        _ = v69
                                    }
                                    var v70 uint32 = (v1 + 739)
                                    fmt.Println(v35)
                                    _ = v38
                                    _ = v70
                                } else {
                                    fmt.Println(v35, v1)
                                    v71 := []int{fi32, (-(fi33)), genFn0(), (fi32 | fi32)}
                                    _ = len(v71)
                                    _ = v71
                                }
                            default:
                                v72 := GenS9{F0: v1, F1: (v2 % ((391) | 1)), F2: v1}
                                _ = v72
                            }
                            v1 = v1
                            switch "\n\t" {
                            case "\n\t":
                                v73 := v3[v1]
                                _ = v73
                            case "foo bar":
                                v21 += v35
                            }
                            v21 = (v21 + (v35 >> ((4294967295) & 63)))
                            fmt.Println(v21)
                            if (false && true) {
                                v74 := GenS5{F0: 42, F1: "foo bar"}
                                fmt.Println(v21, fi33)
                                var v75 int32 = 0
                                _ = v74
                                _ = v75
                            } else {
                                var v76 int32 = -219
                                var v77 uint = 255
                                _ = v76
                                _ = v77
                            }
                            v1 = 251
                            v1 = (v1 + v1)
                            _ = v35
                        }
                    }
                    _ = fi33
                }
                _ = fi32
            }
            for fi78 := 0; fi78 < 19; fi78++ {
                v79 := []int{(fi78 * -408), fi78, fi78}
                fmt.Println(v21, v2)
                switch "foo bar" {
                case "":
                    if (1 != 127) {
                        for fi80 := 0; fi80 < 5; fi80++ {
                            var v81 uint = 255
                            switch fi78 {
                            case -1:
                                switch "hello" {
                                case "":
                                    fmt.Println(v81)
                                }
                                switch 127 {
                                case 128:
                                    v2 = 42
                                    fmt.Println(v2, v1)
                                    if ((127 != 42) && false) {
                                        v21 = 0
                                        v2 = 42
                                        if true {
                                            v81 = 255
                                            var v82 uint64 = 704
                                            _ = v82
                                        }
                                    }
                                case -1:
                                    v21 = v21
                                case 32767:
                                    _ = len(v79)
                                default:
                                    v83 := GenS9{F0: v1, F1: (^(v2)), F2: (v1 % ((v1) | 1))}
                                    _ = v83
                                }
                            }
                            fmt.Println(fi78)
                            _ = fi80
                            _ = v81
                        }
                        var v84 uint = 894
                        if (false && (v1 > (v1 * (v1 << ((v84) & 63))))) {
                            v85 := GenS10{F0: (v84 & v84)}
                            if false {
                                var v86 int16 = 1
                                var v87 string = (("hello" + (("foo bar" + "hello") + "hello")) + "\n\t")
                                _ = v86
                                _ = v87
                            } else {
                                v88 := GenS11{F0: uint8(v1), F1: 18446744073709551615, F2: 536}
                                if len(v79) >= 2 {
                                    v79 = v79[1:len(v79)]
                                }
                                _ = v88
                            }
                            _ = v85
                        } else {
                            for fi89 := 0; fi89 < 7; fi89++ {
                                if (42 >= 596) {
                                    if (-128 <= -1) {
                                        v84 = v84
                                        var v90 int = (fi78 ^ genFn0())
                                        switch "\n\t" {
                                        case "":
                                            var v91 string = "foo bar"
                                            _ = v91
                                        case "\n\t":
                                            var v92 uint8 = 128
                                            _ = v92
                                        case "hello":
                                            switch ("foo bar" + "") {
                                            case "a":
                                                var v93 string = "foo bar"
                                                var v94 uint64 = 2
                                                if len(v79) >= 2 {
                                                    v79 = v79[1:len(v79)]
                                                }
                                                _ = v93
                                                _ = v94
                                            default:
                                                if (!((!true) || (1 >= 1))) {
                                                    switch (fi89 ^ genFn0()) {
                                                    case 128:
                                                        var v95 uint32 = 2
                                                        _ = v95
                                                    case 387:
                                                        var v96 float64 = (((float64((0.1 - float32(((float64(((0.0 * 0.0) - (3.14159 - (((float32((3.14159 / 1.0)) - (0.0 * 3.14159)) + float32(0.0)) - 0.1)))) + float64(0.1)) * 0.0)))) + 0.1) * float64(((1.5 / float32(0.1)) + 1e10))) + float64(1e10))
                                                        _ = v96
                                                    case -2147483648:
                                                        _ = len(v3)
                                                    }
                                                }
                                            }
                                        default:
                                            fmt.Println(v1, v2)
                                        }
                                        _ = v90
                                    } else {
                                        v97 := GenS6{F0: 1}
                                        v21 -= (v21 ^ (v21 - (v21 & v21)))
                                        var v98 uint8 = 41
                                        _ = v97
                                        _ = v98
                                    }
                                } else {
                                    _ = len(v3)
                                }
                                v21 += v21
                                _ = fi89
                            }
                        }
                        _ = v84
                    } else {
                        if (true || (-1 > 0)) {
                            switch fi78 {
                            case 596:
                                var v99 bool = false
                                _ = v99
                            case -476:
                                for fi100 := 0; fi100 < 19; fi100++ {
                                    _ = len(v3)
                                    var v101 float64 = 0.1
                                    var v102 uint32 = v1
                                    _ = fi100
                                    _ = v101
                                    _ = v102
                                }
                            }
                            var v103 uint32 = (v1 * v1)
                            v2 = (v2 - genFn1(1, v103))
                            _ = v103
                        }
                        if (9223372036854775807 == -1) {
                            if true {
                                if true {
                                    for fi104 := 0; fi104 < 13; fi104++ {
                                        v105 := GenS12{F0: "hello", F1: -492}
                                        var v106 uint16 = uint16(v1)
                                        _ = fi104
                                        _ = v105
                                        _ = v106
                                    }
                                    v1 += 155
                                } else {
                                    v2 = 0
                                    for fi107 := 0; fi107 < 10; fi107++ {
                                        if (42 < 0) {
                                            var v108 int = -373
                                            fmt.Println(v21)
                                            v108 += genFn0()
                                            _ = v108
                                        }
                                        _ = fi107
                                    }
                                    v1 = 65535
                                }
                                v109 := v3[2147483648]
                                _ = v109
                            }
                        }
                    }
                case "hello":
                    var v110 uint = 1
                    _ = v110
                case "foo bar":
                    var v111 int = fi78
                    _ = v111
                default:
                    v112 := GenS12{F0: "foo bar", F1: -9223372036854775808}
                    _ = v112
                }
                _ = fi78
                _ = v79
            }
            v2 = (v2 ^ v2)
        } else {
            v113 := []int{32767, 0, 42}
            _ = len(v113)
            v114 := v3[0]
            _ = v113
            _ = v114
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v25
    _ = v3
}
