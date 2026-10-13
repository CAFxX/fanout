package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int16; F1 string; F2 int64 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 uint8; F1 uint16 }
type GenS6 struct { F0 int64; F1 int32; F2 uint32 }
type GenS7 struct { F0 int16; F1 uint64 }
type GenS8 struct { F0 int32; F1 string }
type GenS9 struct { F0 uint16; F1 int64 }
type GenS10 struct { F0 int16 }
func genFn0(p0 int16) uint32 {
    var v7 bool = (2 == 87)
    for fi8 := 0; fi8 < 20; fi8++ {
        v7 = v7
        fmt.Println(v7, fi8, v7)
        _ = fi8
    }
    var v9 int16 = p0
    v10 := GenS1{F0: 255, F1: 64}
    v11 := GenS2{F0: v7}
    _ = p0
    _ = v10
    _ = v11
    _ = v7
    _ = v9
    return 2147483648
}
func genFn1(p0 int) bool {
    for fi12 := 0; fi12 < 2; fi12++ {
        if (!true) {
            for fi13 := 0; fi13 < 14; fi13++ {
                if true {
                    switch "" {
                    case "hello":
                        var v14 float32 = 0.0
                        _ = v14
                    case "foo bar":
                        var v15 int32 = 1
                        _ = v15
                    default:
                        var v16 bool = (1 != 60)
                        _ = v16
                    }
                    switch "a" {
                    case "hello":
                        var v17 uint = 255
                        _ = v17
                    }
                }
                v18 := []int{p0, fi12}
                v19 := GenS1{F0: 0, F1: 0}
                _ = fi13
                _ = v18
                _ = v19
            }
            var v20 uint = 0
            v20 += (^(v20))
            _ = v20
        } else {
            fmt.Println(fi12)
        }
        var v21 float32 = 0.0
        switch "a" {
        case "\n\t":
            var v22 bool = true
            _ = v22
        case "hello":
            switch "" {
            case "a":
                var v23 bool = (false && (42 >= -1))
                _ = v23
            case "\n\t":
                p0 -= 32767
            }
        default:
            var v24 uint = 18446744073709551615
            _ = v24
        }
        _ = fi12
        _ = v21
    }
    var v25 float32 = 1.5
    v26 := []int{(-(p0))}
    _ = p0
    _ = v25
    _ = v26
    return true
}

func main() {
    var v1 uint16 = 1
    v2 := []string{"foo bar", "", ("\n\t" + ("foo bar" + "a")), "hello"}
    var v3 int64 = -334
    v4 := []uint32{65535, 2}
    var v5 uint8 = 0
    v6 := make(map[uint32]string)
    v3 = v3
    v3 += (v3 | (v3 ^ int64(v1)))
    for fi27 := 0; fi27 < 3; fi27++ {
        v28 := GenS3{F0: -279, F1: "\n\t", F2: v3}
        for fi29 := 0; fi29 < 19; fi29++ {
            fmt.Println(v5, v3)
            if (194 >= 65535) {
                if (3.14159 >= float32(float64(0.0))) {
                    var v30 float32 = 1.5
                    _ = len(v4)
                    v1 = (v1 ^ (v1 / (((v1 * v1)) | 1)))
                    _ = v30
                } else {
                    if true {
                        var v31 uint = 662
                        if ((2147483647 < -1) && (false || false)) {
                            _ = len(v2)
                            switch (("" + "hello") + ("foo bar" + ("foo bar" + ("hello" + "a")))) {
                            case "a":
                                if true {
                                    v5 = (v5 / ((v5) | 1))
                                    if true {
                                        v32 := GenS1{F0: v1, F1: 95}
                                        var v33 float64 = 0.1
                                        _ = v32
                                        _ = v33
                                    } else {
                                        v34 := GenS4{F0: fi29}
                                        _ = v34
                                    }
                                    v6[0] = ""
                                    var v35 int = (fi29 + (fi29 / (((fi27 ^ (fi29 - (fi27 * (fi29 / ((2147483647) | 1)))))) | 1)))
                                    _ = v35
                                } else {
                                    v1 = v1
                                    delete(v6, 42)
                                    if false {
                                        fmt.Println(v31)
                                        if (2 >= v5) {
                                            switch "hello" {
                                            case "":
                                                if (0 <= -1) {
                                                    switch "\n\t" {
                                                    case "\n\t":
                                                        v3 = -959
                                                    case "hello":
                                                        v4 = append(v4, 65535)
                                                    default:
                                                        v4 = append(v4, 4294967295)
                                                    }
                                                    var v36 int8 = -74
                                                    v37 := GenS2{F0: (!false)}
                                                    _ = v36
                                                    _ = v37
                                                }
                                            case "\n\t":
                                                fmt.Println(fi29, v1)
                                            case "a":
                                                v38 := GenS5{F0: 1, F1: v1}
                                                _ = v38
                                            }
                                            v39 := GenS1{F0: v1, F1: -1}
                                            v3 = -1
                                            _ = v39
                                        }
                                        var v40 uint8 = 42
                                        _ = v40
                                    }
                                    var v41 int8 = 64
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                    _ = v41
                                }
                            case "foo bar":
                                switch "\n\t" {
                                case "hello":
                                    var v42 int16 = 42
                                    _ = v42
                                case "foo bar":
                                    var v43 float64 = 0.0
                                    _ = v43
                                default:
                                    fmt.Println(v1, fi27, v31)
                                }
                            case "\n\t":
                                v44 := GenS2{F0: false}
                                _ = v44
                            default:
                                switch (("hello" + "") + ("foo bar" + "a")) {
                                case "foo bar":
                                    if false {
                                        var v45 uint16 = (v1 % (((v1 * (v1 << ((18446744073709551615) & 63)))) | 1))
                                        v46 := v6[2147483647]
                                        v47 := GenS6{F0: v3, F1: int32(v3), F2: 4294967295}
                                        _ = v45
                                        _ = v46
                                        _ = v47
                                    } else {
                                        v5 += (v5 & 235)
                                        if (0 >= int8(v5)) {
                                            if ((!((1 == 1) && (v31 == 0))) && false) {
                                                v1 = 42
                                                var v48 uint = (v31 - 18446744073709551615)
                                                _ = v48
                                            }
                                            v3 -= 1
                                            if false {
                                                if len(v2) > 0 {
                                                    v2[0] = "hello"
                                                }
                                            }
                                            var v49 uint8 = (v5 + 127)
                                            _ = v49
                                        }
                                        _ = len(v6)
                                    }
                                case "a":
                                    var v50 int = (-(fi27))
                                    _ = v50
                                }
                            }
                        }
                        if true {
                            if true {
                                v51 := GenS7{F0: int16(v3), F1: 42}
                                var v52 uint = 1
                                switch -27 {
                                case 42:
                                    var v53 int64 = (^(v3))
                                    _ = v53
                                case -600:
                                    fmt.Println(fi27)
                                    var v54 bool = true
                                    switch ("foo bar" + "foo bar") {
                                    case "":
                                        fmt.Println(v5)
                                    default:
                                        var v55 int8 = 100
                                        _ = v55
                                    }
                                    var v56 float32 = ((0.0 / (3.14159 - (0.0 - (((1e10 / float32(float64(1e10))) - 1.5) / 1.0)))) + ((float32(1e10) / 0.1) / 1.5))
                                    _ = v54
                                    _ = v56
                                case 4294967295:
                                    if false {
                                        var v57 string = ("\n\t" + "\n\t")
                                        fmt.Println(v5, v5, v3)
                                        var v58 int16 = -687
                                        _ = v57
                                        _ = v58
                                    } else {
                                        v59 := v6[4294967295]
                                        var v60 int64 = (-(v3))
                                        _ = v59
                                        _ = v60
                                    }
                                default:
                                    if false {
                                        var v61 uint32 = 65535
                                        v3 = (v3 | v3)
                                        v62 := GenS1{F0: (v1 >> ((373) & 63)), F1: -1}
                                        _ = v61
                                        _ = v62
                                    } else {
                                        switch 4294967295 {
                                        case 127:
                                            if true {
                                                if true {
                                                    delete(v6, genFn0(32767))
                                                } else {
                                                    v63 := GenS4{F0: (fi29 - 1)}
                                                    v3 = 9223372036854775807
                                                    _ = v63
                                                }
                                                var v64 float32 = 0.0
                                                var v65 float32 = v64
                                                _ = v64
                                                _ = v65
                                            }
                                            if (true || true) {
                                                var v66 uint32 = 671
                                                fmt.Println(v1, v52)
                                                _ = v66
                                            }
                                            if true {
                                                fmt.Println(v31, v52, v5)
                                            } else {
                                                v67 := GenS4{F0: fi27}
                                                if false {
                                                    switch "\n\t" {
                                                    case "\n\t":
                                                        v68 := GenS1{F0: (v1 & v1), F1: 127}
                                                        _ = v68
                                                    case "":
                                                        _ = len(v4)
                                                    case "a":
                                                        _ = len(v4)
                                                    default:
                                                        var v69 int16 = -32768
                                                        _ = v69
                                                    }
                                                }
                                                v31 = (v31 >> (((^(v52))) & 63))
                                                v1 -= (^(v1))
                                                _ = v67
                                            }
                                        case 4294967295:
                                            var v70 string = ("hello" + "")
                                            _ = v70
                                        case -719:
                                            fmt.Println(fi27)
                                            fmt.Println(v3, v52, fi29)
                                            v31 -= 720
                                        default:
                                            fmt.Println(v5)
                                        }
                                    }
                                }
                                if (((-2.25 - (float32(((1.5 / 1.0) * 1.5)) + -2.25)) > 0.1) && false) {
                                    var v71 int8 = -128
                                    if true {
                                        var v72 uint32 = uint32(v31)
                                        _ = v72
                                    } else {
                                        if (1 < 515) {
                                            if true {
                                                if (float32((1e10 + float64(float32((0.1 / float64((float32(1e10) * (1e10 / 3.14159)))))))) == (((float32(((3.14159 + 0.0) / (0.1 - 0.0))) / 1.5) * ((float32(1e10) * 0.1) * ((0.0 + ((1.5 / (-2.25 - float32((0.0 * 3.14159)))) - 3.14159)) - 3.14159))) - -2.25)) {
                                                    switch (fi29 / ((-1) | 1)) {
                                                    case -1:
                                                        var v73 uint8 = (v5 << ((v31) & 63))
                                                        _ = v73
                                                    case 2147483647:
                                                        v31 = (v31 + 1)
                                                    default:
                                                        var v74 string = "a"
                                                        _ = v74
                                                    }
                                                }
                                            }
                                            switch "" {
                                            case "a":
                                                if len(v4) > 0 {
                                                    v4[0] = 2147483647
                                                }
                                            default:
                                                switch "a" {
                                                case "":
                                                    var v75 uint16 = (v1 / ((v1) | 1))
                                                    _ = v75
                                                case "\n\t":
                                                    switch (fi27 << ((v31) & 63)) {
                                                    case 128:
                                                        var v76 uint64 = 18446744073709551615
                                                        _ = v76
                                                    case 2147483647:
                                                        v77 := v6[434]
                                                        _ = v77
                                                    }
                                                default:
                                                    v71 += (v71 + (v71 / ((v71) | 1)))
                                                }
                                            }
                                        }
                                        v1 = 32767
                                        v52 = uint(v5)
                                    }
                                    fmt.Println(v5)
                                    _ = v71
                                } else {
                                    var v78 uint16 = 65535
                                    v5 = 0
                                    v79 := GenS7{F0: 705, F1: 0}
                                    v80 := GenS8{F0: 0, F1: "hello"}
                                    _ = v78
                                    _ = v79
                                    _ = v80
                                }
                                fmt.Println(v52, fi29)
                                fmt.Println(v52)
                                fmt.Println(v52, fi29, v52)
                                _ = v51
                                _ = v52
                            }
                            var v81 uint64 = 4294967295
                            var v82 int32 = 42
                            v31 = (v31 << (((v31 / (((v31 | (^(v31)))) | 1))) & 63))
                            v31 += v31
                            _ = v81
                            _ = v82
                        } else {
                            if (0 <= 390) {
                                v1 = (v1 & 378)
                            } else {
                                if (-2147483648 <= 2147483647) {
                                    delete(v6, 2147483648)
                                    var v83 int16 = 0
                                    _ = v83
                                } else {
                                    switch 792 {
                                    case 9223372036854775807:
                                        v1 = v1
                                    case 32768:
                                        v84 := GenS3{F0: 42, F1: "", F2: (v3 - v3)}
                                        _ = v84
                                    case 1:
                                        v1 = (v1 % ((1) | 1))
                                    }
                                }
                                var v85 string = "hello"
                                _ = v85
                            }
                            var v86 bool = (42 == 982)
                            _ = v86
                        }
                        _ = v31
                    }
                    var v87 int32 = 1
                    v3 -= v3
                    v88 := GenS3{F0: -337, F1: (("a" + ("a" + ("a" + "a"))) + "hello"), F2: -511}
                    var v89 int = (fi27 ^ (fi27 + (fi27 << ((65535) & 63))))
                    _ = v87
                    _ = v88
                    _ = v89
                }
                switch "" {
                case "\n\t":
                    v3 -= 41
                case "hello":
                    var v90 int32 = 2147483647
                    _ = v90
                case "":
                    v5 = v5
                }
                switch (fi27 & fi29) {
                case -2147483648:
                    if len(v2) > 0 {
                        v2[0] = "hello"
                    }
                default:
                    v5 = v5
                    v91 := GenS3{F0: 1000, F1: (((((("foo bar" + ((("foo bar" + "") + ((((("\n\t" + "\n\t") + "") + "foo bar") + "foo bar") + "hello")) + "a")) + "\n\t") + "foo bar") + "a") + "hello") + "a"), F2: (v3 % ((9223372036854775807) | 1))}
                    _ = v91
                }
            }
            fmt.Println(v3, fi27)
            _ = fi29
        }
        v92 := GenS9{F0: (v1 * 577), F1: v3}
        _ = fi27
        _ = v28
        _ = v92
    }
    for fi93 := 0; fi93 < 18; fi93++ {
        if true {
            if true {
                v5 -= (v5 >> ((42) & 63))
                v94 := GenS7{F0: 0, F1: 4294967295}
                _ = v94
            }
        }
        _ = fi93
    }
    if len(v4) > 0 {
        v4[0] = 929
    }
    switch ("" + "") {
    case "":
        switch "foo bar" {
        case "hello":
            v95 := GenS4{F0: 187}
            _ = v95
        case "foo bar":
            for fi96 := 0; fi96 < 16; fi96++ {
                for fi97 := 0; fi97 < 5; fi97++ {
                    _ = len(v6)
                    var v98 float32 = 1e10
                    _ = fi97
                    _ = v98
                }
                _ = fi96
            }
        case "\n\t":
            delete(v6, 2147483648)
        }
    case "hello":
        switch "foo bar" {
        case "":
            v1 = v1
        }
    case "foo bar":
        v3 -= -1
    default:
        v1 = (v1 * 42)
    }
    if (uint(v5) == uint(v1)) {
        var v99 uint8 = 1
        v100 := GenS10{F0: 1000}
        _ = v100
        _ = v99
    }
    v5 += v5
    var v101 string = "hello"
    fmt.Println("v1:", v1)
    fmt.Println("v101:", v101)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v101
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
