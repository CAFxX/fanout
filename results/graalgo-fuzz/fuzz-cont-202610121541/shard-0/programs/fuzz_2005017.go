package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int }
type GenS2 struct { F0 uint; F1 int32; F2 uint64 }
type GenS3 struct { F0 uint32; F1 int64 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 int; F1 int32; F2 int32 }
type GenS6 struct { F0 string; F1 uint; F2 uint }
type GenS7 struct { F0 int8 }
type GenS8 struct { F0 uint64 }
type GenS9 struct { F0 uint32; F1 uint32 }
type GenS10 struct { F0 string; F1 uint; F2 bool }
func genFn0() int64 {
    var v5 float64 = 3.14159
    if (!(!true)) {
        var v6 uint64 = 1
        _ = v6
    } else {
        fmt.Println(v5)
        var v7 string = "foo bar"
        var v8 float64 = ((1.5 * v5) / float64(((0.0 / 1e10) + (float32(v5) + 3.14159))))
        _ = v7
        _ = v8
    }
    v9 := GenS1{F0: 161, F1: 2147483647}
    for fi10 := 0; fi10 < 5; fi10++ {
        for fi11 := 0; fi11 < 7; fi11++ {
            fmt.Println(fi10, fi11)
            _ = fi11
        }
        _ = fi10
    }
    for fi12 := 0; fi12 < 5; fi12++ {
        v5 += 0.1
        for fi13 := 0; fi13 < 1; fi13++ {
            v5 = float64(0.0)
            _ = fi13
        }
        _ = fi12
    }
    _ = v5
    _ = v9
    return 9223372036854775807
}
func genFn1(p0 uint64, p1 uint16) int8 {
    v14 := GenS1{F0: 127, F1: 127}
    p0 = (p0 & p0)
    var v15 float64 = ((-2.25 + float64(1.5)) * (((1e10 + (3.14159 - 0.0)) - float64((-2.25 + (0.1 / (1e10 * 1.5))))) * 0.0))
    for fi16 := 0; fi16 < 3; fi16++ {
        var v17 int8 = 100
        v18 := []int{-1, fi16, fi16, (fi16 - fi16)}
        _ = fi16
        _ = v17
        _ = v18
    }
    v19 := GenS1{F0: 2, F1: 42}
    _ = p0
    _ = p1
    _ = v14
    _ = v15
    _ = v19
    return 0
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 int32 = 0
    v3 := make(map[string]uint32)
    v4 := []string{("hello" + ("hello" + "\n\t")), ""}
    v20 := GenS2{F0: 0, F1: v2, F2: 18446744073709551615}
    v2 += v2
    if ((!false) && false) {
        fmt.Println(v2)
        var v21 int32 = 376
        v22 := GenS1{F0: 207, F1: -2147483648}
        _ = v21
        _ = v22
    }
    if (2147483648 != 2147483648) {
        v1[574] = 2147483647
        switch 42 {
        case 4294967295:
            v23 := GenS3{F0: 997, F1: 1}
            _ = v23
        case 864:
            if false {
                var v24 int8 = -35
                fmt.Println(v24, v24)
                _ = v24
            } else {
                var v25 int8 = 127
                _ = v25
            }
        case 128:
            v2 = v2
        default:
            var v26 int8 = -1
            _ = v26
        }
        if false {
            for fi27 := 0; fi27 < 17; fi27++ {
                v28 := v3["a"]
                if ((false || true) && ((fi27 - (fi27 << ((277) & 63))) >= (fi27 + (fi27 & fi27)))) {
                    var v29 int = 9223372036854775807
                    _ = v29
                } else {
                    v30 := GenS4{F0: -1}
                    switch (fi27 ^ (fi27 + fi27)) {
                    case 32768:
                        for fi31 := 0; fi31 < 18; fi31++ {
                            delete(v1, v28)
                            _ = fi31
                        }
                    case 2147483647:
                        var v32 int32 = (v2 + (v2 & v2))
                        _ = v32
                    default:
                        v33 := GenS5{F0: (fi27 % (((fi27 << ((255) & 63))) | 1)), F1: (v2 ^ v2), F2: v2}
                        _ = v33
                    }
                    v34 := GenS3{F0: (v28 * (v28 >> ((42) & 63))), F1: 1}
                    _ = v30
                    _ = v34
                }
                if len(v4) > 0 {
                    v4[0] = "hello"
                }
                _ = fi27
                _ = v28
            }
            var v35 uint64 = 978
            _ = v35
        } else {
            var v36 uint8 = 2
            if (false && true) {
                if false {
                    fmt.Println(v2, v2)
                    var v37 uint32 = 222
                    _ = v37
                }
                var v38 float32 = 0.1
                for fi39 := 0; fi39 < 19; fi39++ {
                    fmt.Println(v38, v2, v36)
                    _ = len(v4)
                    _ = fi39
                }
                _ = v38
            } else {
                switch "foo bar" {
                case "\n\t":
                    if true {
                        for fi40 := 0; fi40 < 3; fi40++ {
                            v2 -= (v2 + 2147483647)
                            _ = fi40
                        }
                        var v41 uint8 = (v36 ^ 4)
                        _ = v41
                    }
                case "a":
                    if false {
                        var v42 float64 = 0.0
                        if (false && (false && false)) {
                            for fi43 := 0; fi43 < 20; fi43++ {
                                for fi44 := 0; fi44 < 1; fi44++ {
                                    if false {
                                        v2 = 42
                                        _ = len(v4)
                                        v42 = (float64((0.0 - 1e10)) / -2.25)
                                    } else {
                                        if len(v4) >= 2 {
                                            v4 = v4[1:len(v4)]
                                        }
                                        var v45 uint16 = 791
                                        var v46 int64 = 635
                                        _ = len(v4)
                                        switch (fi43 % (((fi43 / ((-617) | 1))) | 1)) {
                                        case 32768:
                                            var v47 float64 = -2.25
                                            switch (fi43 + 32767) {
                                            case -735:
                                                var v48 float64 = v42
                                                var v49 uint8 = 127
                                                _ = v48
                                                _ = v49
                                            case 9223372036854775807:
                                                v2 = 607
                                            case 0:
                                                v50 := v1[uint32(v36)]
                                                _ = v50
                                            }
                                            fmt.Println(v45, v45)
                                            _ = v47
                                        default:
                                            var v51 uint64 = 0
                                            _ = v51
                                        }
                                        _ = v45
                                        _ = v46
                                    }
                                    var v52 bool = true
                                    _ = fi44
                                    _ = v52
                                }
                                _ = fi43
                            }
                            for fi53 := 0; fi53 < 16; fi53++ {
                                for fi54 := 0; fi54 < 16; fi54++ {
                                    delete(v1, 17)
                                    _ = len(v4)
                                    _ = fi54
                                }
                                for fi55 := 0; fi55 < 3; fi55++ {
                                    if false {
                                        v2 -= v2
                                        v56 := GenS5{F0: (-(fi55)), F1: v2, F2: v2}
                                        var v57 uint = 42
                                        _ = v56
                                        _ = v57
                                    } else {
                                        var v58 uint32 = 245
                                        _ = v58
                                    }
                                    if true {
                                        if (9223372036854775807 == -458) {
                                            fmt.Println(fi53, v2, fi53)
                                            _ = len(v4)
                                            if len(v4) >= 2 {
                                                v4 = v4[1:len(v4)]
                                            }
                                        } else {
                                            switch (fi53 % ((fi55) | 1)) {
                                            case -9223372036854775808:
                                                v59 := GenS1{F0: (v36 - (v36 + v36)), F1: -9223372036854775808}
                                                _ = v59
                                            case 485:
                                                v60 := GenS1{F0: (v36 & (v36 | v36)), F1: (fi55 - 0)}
                                                _ = v60
                                            case 66:
                                                if false {
                                                    v2 = (v2 | 1)
                                                }
                                            }
                                            v36 += (v36 * v36)
                                            var v61 int32 = (v2 << ((65535) & 63))
                                            if true {
                                                switch "foo bar" {
                                                case "foo bar":
                                                    var v62 uint8 = (v36 * v36)
                                                    _ = v62
                                                default:
                                                    switch (fi53 - 42) {
                                                    case 128:
                                                        var v63 int8 = 104
                                                        _ = v63
                                                    case -27:
                                                        var v64 uint32 = 0
                                                        _ = v64
                                                    }
                                                }
                                                switch "\n\t" {
                                                case "a":
                                                    v65 := v3[("hello" + "\n\t")]
                                                    _ = v65
                                                case "":
                                                    v36 = v36
                                                case "hello":
                                                    v66 := GenS1{F0: v36, F1: fi55}
                                                    _ = v66
                                                }
                                                if true {
                                                    switch fi53 {
                                                    case 32768:
                                                        switch ("" + (("foo bar" + "") + (((((("" + "foo bar") + "\n\t") + "") + "\n\t") + ((((("a" + "\n\t") + "") + "\n\t") + "\n\t") + (("hello" + "foo bar") + "\n\t"))) + ""))) {
                                                        case "\n\t":
                                                            v61 = (v2 >> ((42) & 63))
                                                        case "a":
                                                            switch "a" {
                                                            case "\n\t":
                                                                var v67 int64 = -1
                                                                _ = v67
                                                            }
                                                        default:
                                                            if false {
                                                                switch "\n\t" {
                                                                case "":
                                                                    v68 := GenS6{F0: "\n\t", F1: 994, F2: 18446744073709551615}
                                                                    delete(v3, (("\n\t" + "hello") + "hello"))
                                                                    _ = v68
                                                                case "hello":
                                                                    if true {
                                                                        switch "\n\t" {
                                                                        case "a":
                                                                            if ((("hello" + "hello") + "\n\t") >= "a") {
                                                                                v69 := GenS7{F0: 127}
                                                                                switch ("a" + (("\n\t" + "hello") + "a")) {
                                                                                case "":
                                                                                    var v70 int16 = 0
                                                                                    _ = v70
                                                                                case "foo bar":
                                                                                    v61 = 1
                                                                                }
                                                                                fmt.Println(v36, v42, fi55)
                                                                                _ = v69
                                                                            }
                                                                        case "hello":
                                                                            var v71 int16 = 1
                                                                            _ = v71
                                                                        case "foo bar":
                                                                            v72 := GenS3{F0: 78, F1: 42}
                                                                            _ = v72
                                                                        default:
                                                                            var v73 uint64 = 18446744073709551615
                                                                            _ = v73
                                                                        }
                                                                        if (1 <= (^(v36))) {
                                                                            var v74 int = -223
                                                                            var v75 bool = (true && false)
                                                                            fmt.Println(v42, fi55)
                                                                            _ = v74
                                                                            _ = v75
                                                                        }
                                                                    } else {
                                                                        var v76 bool = true
                                                                        var v77 bool = (-740 == 42)
                                                                        v1[2147483648] = 2
                                                                        var v78 uint64 = 18446744073709551615
                                                                        _ = len(v3)
                                                                        v78 -= v78
                                                                        v42 = v42
                                                                        if v77 {
                                                                            var v79 int8 = int8(fi53)
                                                                            var v80 uint16 = 675
                                                                            _ = len(v3)
                                                                            _ = v79
                                                                            _ = v80
                                                                        }
                                                                        v81 := GenS8{F0: (v78 | 1)}
                                                                        fmt.Println(v2)
                                                                        var v82 int8 = -1
                                                                        _ = v76
                                                                        _ = v77
                                                                        _ = v78
                                                                        _ = v81
                                                                        _ = v82
                                                                    }
                                                                case "a":
                                                                    v83 := GenS1{F0: (v36 ^ (v36 + v36)), F1: 388}
                                                                    _ = v83
                                                                }
                                                                v42 = (0.0 - v42)
                                                            } else {
                                                                v84 := v1[2]
                                                                if (!(true || (fi55 <= (^(fi53))))) {
                                                                    var v85 uint64 = 0
                                                                    _ = v85
                                                                }
                                                                switch "a" {
                                                                case "":
                                                                    var v86 bool = (!(!(!(("\n\t" + ("a" + "foo bar")) >= ""))))
                                                                    _ = v86
                                                                default:
                                                                    v2 = (v2 - v2)
                                                                }
                                                                _ = v84
                                                            }
                                                        }
                                                        fmt.Println(fi55, fi55)
                                                        _ = len(v4)
                                                    case 643:
                                                        if true {
                                                            var v87 int64 = genFn0()
                                                            fmt.Println(v87)
                                                            var v88 int64 = -61
                                                            var v89 uint8 = (^(v36))
                                                            _ = v87
                                                            _ = v88
                                                            _ = v89
                                                        }
                                                    default:
                                                        v61 = v2
                                                    }
                                                }
                                            }
                                            _ = v61
                                        }
                                    }
                                    var v90 int16 = -32768
                                    fmt.Println(v36, v2, v42)
                                    _ = len(v4)
                                    var v91 int8 = -1
                                    _ = fi55
                                    _ = v90
                                    _ = v91
                                }
                                _ = fi53
                            }
                            var v92 int32 = v2
                            _ = v92
                        }
                        switch 0 {
                        case 42:
                            var v93 int32 = (v2 >> ((29) & 63))
                            _ = v93
                        case 0:
                            var v94 string = "a"
                            _ = v94
                        case 108:
                            switch 128 {
                            case -9223372036854775808:
                                v95 := GenS9{F0: 1, F1: 1}
                                _ = v95
                            case 293:
                                if false {
                                    v96 := GenS10{F0: ("foo bar" + "\n\t"), F1: 4294967295, F2: true}
                                    _ = v96
                                } else {
                                    var v97 int16 = 1
                                    _ = v97
                                }
                            case -362:
                                fmt.Println(v2, v36, v36)
                            default:
                                v36 = (v36 / ((6) | 1))
                            }
                        default:
                            v36 = v36
                        }
                        _ = v42
                    } else {
                        v2 = (v2 >> ((2) & 63))
                        for fi98 := 0; fi98 < 7; fi98++ {
                            var v99 bool = (1e10 == 1.5)
                            for fi100 := 0; fi100 < 18; fi100++ {
                                v36 += (^(v36))
                                if false {
                                    fmt.Println(v99)
                                    v101 := GenS8{F0: 771}
                                    _ = v101
                                } else {
                                    if v99 {
                                        var v102 bool = (true || (v2 > (v2 ^ v2)))
                                        delete(v3, "")
                                        fmt.Println(v102, v36, v2)
                                        _ = v102
                                    }
                                }
                                _ = fi100
                            }
                            for fi103 := 0; fi103 < 6; fi103++ {
                                v104 := GenS8{F0: 18446744073709551615}
                                var v105 int = (^(fi103))
                                var v106 uint64 = 4294967295
                                var v107 string = "hello"
                                _ = fi103
                                _ = v104
                                _ = v105
                                _ = v106
                                _ = v107
                            }
                            _ = fi98
                            _ = v99
                        }
                        fmt.Println(v36, v2)
                    }
                }
                v36 = (^(v36))
                var v108 int16 = -986
                _ = v108
            }
            _ = v36
        }
    } else {
        fmt.Println(v2)
        var v109 bool = true
        v110 := GenS7{F0: -1}
        _ = v109
        _ = v110
    }
    v2 += v2
    _ = len(v4)
    v2 = -2147483648
    v111 := GenS10{F0: ("" + "a"), F1: 4294967295, F2: false}
    var v112 float32 = 3.14159
    if (false && true) {
        var v113 int = 1
        fmt.Println(v113, v2)
        _ = v113
    } else {
        switch "" {
        case "a":
            v114 := v1[65535]
            _ = v114
        }
    }
    v112 += v112
    v2 += v2
    if (true && (("hello" + ("\n\t" + "")) > ("a" + ""))) {
        fmt.Println(v2)
        for fi115 := 0; fi115 < 20; fi115++ {
            v3["foo bar"] = 770
            _ = fi115
        }
    }
    for fi116 := 0; fi116 < 6; fi116++ {
        switch fi116 {
        case 2147483647:
            if true {
                v2 = (^(v2))
            }
        }
        _ = fi116
    }
    var v117 uint32 = 42
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v112:", v112)
    fmt.Println("v117:", v117)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v111
    _ = v112
    _ = v117
    _ = v2
    _ = v20
    _ = v3
    _ = v4
}
