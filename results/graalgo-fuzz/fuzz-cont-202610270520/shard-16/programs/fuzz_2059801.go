package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint; F2 string }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int8; F1 string }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 int }
type GenS8 struct { F0 int64; F1 int8 }
type GenS9 struct { F0 uint16; F1 uint8 }
type GenS10 struct { F0 int8; F1 uint }
type GenS11 struct { F0 int8; F1 uint64 }
type GenS12 struct { F0 int64; F1 int; F2 int }
type GenS13 struct { F0 uint8; F1 int32; F2 int8 }
func genFn0(p0 int8, p1 uint32, p2 int64) uint32 {
    v4 := []int{-2147483648, 124}
    if true {
        for fi5 := 0; fi5 < 15; fi5++ {
            fmt.Println(p0)
            v6 := GenS1{F0: p2, F1: 65535, F2: "\n\t"}
            p2 -= (p2 | (p2 ^ 9223372036854775807))
            _ = fi5
            _ = v6
        }
        var v7 bool = false
        var v8 string = ""
        _ = v7
        _ = v8
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return (p1 << ((2) & 63))
}
func genFn1() int16 {
    for fi9 := 0; fi9 < 6; fi9++ {
        var v10 int32 = -495
        _ = fi9
        _ = v10
    }
    for fi11 := 0; fi11 < 12; fi11++ {
        switch fi11 {
        case -9223372036854775808:
            for fi12 := 0; fi12 < 11; fi12++ {
                var v13 uint64 = 0
                var v14 int64 = 9223372036854775807
                var v15 bool = true
                var v16 int64 = v14
                v16 -= v16
                _ = fi12
                _ = v13
                _ = v14
                _ = v15
                _ = v16
            }
        case 42:
            var v17 float32 = ((3.14159 + (1.5 / 1.0)) + 1e10)
            _ = v17
        }
        var v18 uint32 = genFn0(-28, 43, -405)
        if true {
            v19 := GenS1{F0: 1, F1: 2, F2: ("\n\t" + "foo bar")}
            for fi20 := 0; fi20 < 7; fi20++ {
                fmt.Println(v18, fi20)
                switch "a" {
                case "":
                    var v21 int8 = -1
                    _ = v21
                default:
                    if (1 < 1) {
                        if false {
                            var v22 uint32 = 42
                            v22 = genFn0(100, 425, -9223372036854775808)
                            switch 690 {
                            case -1:
                                var v23 float64 = float64(3.14159)
                                var v24 uint = 668
                                _ = v23
                                _ = v24
                            default:
                                v25 := []int{fi11, (fi20 & -1), fi11}
                                _ = v25
                            }
                            var v26 int32 = -622
                            _ = v22
                            _ = v26
                        }
                        var v27 int16 = 32767
                        var v28 float32 = -2.25
                        switch 32768 {
                        case -2147483648:
                            var v29 float32 = v28
                            _ = v29
                        case -9223372036854775808:
                            v30 := []int{1}
                            _ = v30
                        case -275:
                            if (true && true) {
                                var v31 bool = (v27 != (v27 + 136))
                                var v32 int16 = v27
                                v33 := GenS2{F0: "foo bar"}
                                _ = v31
                                _ = v32
                                _ = v33
                            } else {
                                v18 = (v18 >> ((18446744073709551615) & 63))
                                var v34 int16 = v27
                                _ = v34
                            }
                        }
                        _ = v27
                        _ = v28
                    }
                }
                v35 := []int{fi11, (fi11 & -9223372036854775808)}
                var v36 float32 = (((3.14159 - 0.1) + ((1.5 * -2.25) / (float32(1.5) + ((((-2.25 * (((-2.25 * 3.14159) * (1e10 - (0.0 - 1.5))) - ((0.1 + 1e10) / 1.0))) + -2.25) / 1.0) + 0.0)))) * -2.25)
                if (!((0.0 - -2.25) >= (0.0 * (0.0 * 1e10)))) {
                    var v37 uint = 65535
                    v36 = float32(0.0)
                    _ = v37
                }
                _ = fi20
                _ = v35
                _ = v36
            }
            var v38 uint8 = 2
            _ = v19
            _ = v38
        }
        _ = fi11
        _ = v18
    }
    return 0
}

func main() {
    var v1 int64 = -764
    var v2 int16 = 1
    var v3 float32 = 0.0
    v1 = v1
    fmt.Println(v2)
    v39 := GenS3{F0: 65535}
    v1 = v1
    for fi40 := 0; fi40 < 7; fi40++ {
        for fi41 := 0; fi41 < 20; fi41++ {
            v42 := []int{(fi40 << ((18446744073709551615) & 63)), 2147483647, (-(fi40)), fi40}
            v43 := GenS4{F0: 8, F1: ("\n\t" + "foo bar")}
            if len(v42) >= 2 {
                v42 = v42[1:len(v42)]
            }
            _ = fi41
            _ = v42
            _ = v43
        }
        v2 = (v2 | (v2 % (((v2 | (v2 + (-(v2))))) | 1)))
        _ = fi40
    }
    if ((391 >= 2147483647) || true) {
        var v44 float32 = 0.1
        _ = v44
    }
    v45 := []int{}
    v1 = v1
    fmt.Println(v2)
    fmt.Println(v3, v1)
    for fi46 := 0; fi46 < 3; fi46++ {
        var v47 uint32 = 42
        _ = fi46
        _ = v47
    }
    fmt.Println(v2)
    v48 := GenS5{F0: false}
    v3 += v3
    for fi49 := 0; fi49 < 2; fi49++ {
        for fi50 := 0; fi50 < 20; fi50++ {
            switch "a" {
            case "":
                var v51 int64 = v1
                _ = v51
            case "foo bar":
                v1 -= v1
            case "hello":
                var v52 uint = 18446744073709551615
                _ = v52
            }
            var v53 float64 = 1e10
            _ = fi50
            _ = v53
        }
        var v54 uint = 42
        var v55 int16 = (^(v2))
        _ = fi49
        _ = v54
        _ = v55
    }
    for fi56 := 0; fi56 < 18; fi56++ {
        for fi57 := 0; fi57 < 5; fi57++ {
            var v58 uint = 18446744073709551615
            _ = fi57
            _ = v58
        }
        switch ("" + ("hello" + ("\n\t" + "\n\t"))) {
        case "hello":
            v1 -= v1
        case "\n\t":
            switch ("a" + ("\n\t" + "a")) {
            case "hello":
                v59 := GenS1{F0: (v1 << ((1) & 63)), F1: 18446744073709551615, F2: ("foo bar" + ("a" + "foo bar"))}
                _ = v59
            case "":
                var v60 uint = 222
                _ = v60
            case "a":
                var v61 int64 = (v1 & (v1 + 1))
                _ = v61
            }
        case "":
            v2 += (v2 % ((42) | 1))
        default:
            if false {
                if (!(42 != 127)) {
                    for fi62 := 0; fi62 < 8; fi62++ {
                        var v63 int8 = 100
                        var v64 int32 = -1
                        _ = fi62
                        _ = v63
                        _ = v64
                    }
                }
                for fi65 := 0; fi65 < 15; fi65++ {
                    v1 = 0
                    _ = fi65
                }
            } else {
                for fi66 := 0; fi66 < 4; fi66++ {
                    if len(v45) > 0 {
                        v45[0] = fi66
                    }
                    v45 = append(v45, 0)
                    if (!(32767 >= fi66)) {
                        if (65535 > 42) {
                            v67 := GenS6{F0: 42}
                            v68 := GenS2{F0: (((("\n\t" + "a") + "hello") + "hello") + "a")}
                            if true {
                                var v69 int32 = -350
                                switch "\n\t" {
                                case "\n\t":
                                    if len(v45) >= 2 {
                                        v45 = v45[1:len(v45)]
                                    }
                                case "foo bar":
                                    v70 := GenS7{F0: fi66}
                                    _ = v70
                                case "hello":
                                    var v71 int8 = int8(v2)
                                    _ = v71
                                default:
                                    if false {
                                        if true {
                                            switch "a" {
                                            case "":
                                                if false {
                                                    v2 += (v2 & v2)
                                                    var v72 int = 4294967295
                                                    _ = v72
                                                }
                                            }
                                            v69 -= (v69 % (((v69 ^ (v69 % ((v69) | 1)))) | 1))
                                            v73 := GenS2{F0: ("foo bar" + ("hello" + ""))}
                                            _ = v73
                                        } else {
                                            fmt.Println(v69, v1)
                                        }
                                        var v74 int16 = v2
                                        _ = v74
                                    } else {
                                        v69 -= (v69 >> ((255) & 63))
                                        if true {
                                            var v75 uint = 42
                                            _ = v75
                                        }
                                    }
                                }
                                switch "a" {
                                case "hello":
                                    var v76 int = int(v1)
                                    if (3.14159 > 0.1) {
                                        v1 = 9223372036854775807
                                    } else {
                                        v1 += (-(v1))
                                        switch (v76 | 32767) {
                                        case 32767:
                                            v1 = v1
                                        case 32768:
                                            var v77 uint16 = 2
                                            _ = v77
                                        }
                                    }
                                    if len(v45) > 0 {
                                        v45[0] = (-(v76))
                                    }
                                    if len(v45) >= 2 {
                                        v45 = v45[1:len(v45)]
                                    }
                                    _ = v76
                                default:
                                    switch "\n\t" {
                                    case "\n\t":
                                        v78 := GenS5{F0: true}
                                        _ = v78
                                    case "a":
                                        if true {
                                            var v79 float64 = -2.25
                                            switch fi66 {
                                            case 32768:
                                                if len(v45) > 0 {
                                                    v45[0] = fi56
                                                }
                                                var v80 uint32 = 593
                                                _ = v80
                                            case 541:
                                                v81 := GenS8{F0: v1, F1: 0}
                                                _ = v81
                                            case 128:
                                                var v82 float64 = (1e10 / ((v79 + 1.5) * v79))
                                                _ = v82
                                            }
                                            _ = v79
                                        }
                                    }
                                }
                                v83 := GenS1{F0: (^(v1)), F1: 4294967295, F2: "hello"}
                                _ = v69
                                _ = v83
                            }
                            var v84 int32 = 219
                            _ = v67
                            _ = v68
                            _ = v84
                        } else {
                            v85 := GenS6{F0: 120}
                            if true {
                                v2 -= (v2 + 242)
                            } else {
                                if true {
                                    v1 += (v1 >> ((18446744073709551615) & 63))
                                } else {
                                    var v86 uint = 163
                                    var v87 string = ""
                                    fmt.Println(fi66, v87, v86)
                                    _ = v86
                                    _ = v87
                                }
                                if len(v45) > 0 {
                                    v45[0] = fi66
                                }
                                fmt.Println(fi56, v2, v1)
                            }
                            _ = v85
                        }
                        switch fi66 {
                        case 0:
                            v2 += (^(v2))
                        case 127:
                            switch ("a" + ("" + "a")) {
                            case "foo bar":
                                switch ("hello" + "foo bar") {
                                case "a":
                                    v45 = append(v45, fi56)
                                case "\n\t":
                                    var v88 uint16 = 2
                                    _ = v88
                                case "hello":
                                    var v89 uint8 = 42
                                    _ = v89
                                default:
                                    v90 := GenS6{F0: 42}
                                    _ = v90
                                }
                            case "":
                                v1 += (v1 & 1)
                            case "\n\t":
                                if false {
                                    var v91 int = 128
                                    fmt.Println(fi56)
                                    _ = v91
                                }
                            default:
                                switch "" {
                                case "\n\t":
                                    fmt.Println(fi56, v1)
                                case "a":
                                    v3 = 3.14159
                                default:
                                    var v92 uint = 0
                                    _ = v92
                                }
                                v1 -= (v1 | (v1 << ((18446744073709551615) & 63)))
                                switch (("a" + "") + (("" + ("a" + "hello")) + (("" + ("a" + "hello")) + "\n\t"))) {
                                case "\n\t":
                                    if len(v45) > 0 {
                                        v45[0] = -693
                                    }
                                case "a":
                                    v93 := GenS9{F0: 42, F1: 109}
                                    _ = v93
                                case "hello":
                                    v3 = 0.1
                                    var v94 string = ""
                                    _ = v94
                                default:
                                    if (!(!true)) {
                                        v95 := GenS9{F0: 2, F1: 1}
                                        var v96 uint64 = 30
                                        var v97 int16 = (v2 % ((v2) | 1))
                                        _ = v95
                                        _ = v96
                                        _ = v97
                                    }
                                }
                            }
                        default:
                            v2 -= -737
                        }
                        if (42 >= 127) {
                            var v98 uint32 = 4294967295
                            var v99 uint32 = 671
                            var v100 uint32 = (v99 * v99)
                            var v101 int32 = 42
                            var v102 float32 = ((v3 - float32(((1.5 / 3.14159) - 0.1))) / v3)
                            _ = v100
                            _ = v101
                            _ = v102
                            _ = v98
                            _ = v99
                        } else {
                            var v103 uint64 = 42
                            _ = v103
                        }
                    }
                    _ = fi66
                }
                for fi104 := 0; fi104 < 2; fi104++ {
                    switch 9223372036854775807 {
                    case 42:
                        v105 := GenS10{F0: 42, F1: 828}
                        if len(v45) > 0 {
                            v45[0] = (fi56 >> ((962) & 63))
                        }
                        if true {
                            fmt.Println(v3)
                            v45 = append(v45, (fi104 << ((42) & 63)))
                            v3 -= v3
                            v1 += (v1 + (-(v1)))
                        } else {
                            v1 = v1
                            if false {
                                if ((!("hello" != "foo bar")) || ((v3 / (v3 / -2.25)) != float32(3.14159))) {
                                    switch fi56 {
                                    case 32768:
                                        var v106 string = "\n\t"
                                        _ = v106
                                    }
                                    if len(v45) >= 2 {
                                        v45 = v45[1:len(v45)]
                                    }
                                    fmt.Println(fi104, fi56, fi56)
                                    var v107 uint = 630
                                    _ = v107
                                }
                                if len(v45) > 0 {
                                    v45[0] = fi56
                                }
                            }
                        }
                        _ = v105
                    }
                    if true {
                        v108 := GenS11{F0: -1, F1: 542}
                        if true {
                            v2 = -1
                            fmt.Println(fi56, fi56)
                            fmt.Println(fi104, v2)
                        } else {
                            if true {
                                v109 := GenS12{F0: (v1 | (v1 >> ((65535) & 63))), F1: (fi104 % ((fi56) | 1)), F2: int(v1)}
                                var v110 int32 = 476
                                if ("\n\t" == "hello") {
                                    v3 += v3
                                    if false {
                                        switch ("hello" + (("a" + (("a" + "a") + (("a" + "hello") + ""))) + "foo bar")) {
                                        case "hello":
                                            v110 = v110
                                        case "":
                                            v111 := GenS13{F0: 255, F1: v110, F2: 127}
                                            fmt.Println(v3)
                                            v3 = -2.25
                                            v110 -= v110
                                            var v112 float64 = 1e10
                                            _ = v111
                                            _ = v112
                                        default:
                                            if false {
                                                v3 = v3
                                                var v113 int16 = 424
                                                _ = v113
                                            }
                                        }
                                        var v114 uint = 477
                                        v110 -= (^(v110))
                                        _ = v114
                                    }
                                    if len(v45) > 0 {
                                        v45[0] = fi56
                                    }
                                } else {
                                    v2 = (v2 % (((v2 % ((v2) | 1))) | 1))
                                }
                                v1 = v1
                                _ = v109
                                _ = v110
                            }
                        }
                        if len(v45) >= 2 {
                            v45 = v45[1:len(v45)]
                        }
                        _ = v108
                    }
                    var v115 uint8 = 127
                    _ = fi104
                    _ = v115
                }
            }
        }
        if len(v45) >= 2 {
            v45 = v45[1:len(v45)]
        }
        _ = fi56
    }
    v1 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v45:", len(v45), v45)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v45
    _ = v48
}
