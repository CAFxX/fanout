package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32 }
type GenS2 struct { F0 uint; F1 bool }
type GenS3 struct { F0 string; F1 int64; F2 uint64 }
type GenS4 struct { F0 int16; F1 uint16 }
type GenS5 struct { F0 string; F1 uint; F2 uint32 }
type GenS6 struct { F0 string }
type GenS7 struct { F0 string; F1 int32 }
type GenS8 struct { F0 string; F1 int16; F2 uint64 }
type GenS9 struct { F0 uint64; F1 int8 }
func genFn0(p0 uint8, p1 string, p2 bool) uint32 {
    p1 = (p1 + p1)
    for fi6 := 0; fi6 < 3; fi6++ {
        v7 := []int{(fi6 * (fi6 % (((fi6 - fi6)) | 1)))}
        var v8 uint32 = 1
        _ = fi6
        _ = v7
        _ = v8
    }
    p0 -= p0
    _ = p0
    _ = p1
    _ = p2
    return 2147483648
}
func genFn1(p0 uint64, p1 int8, p2 uint) uint64 {
    if (-9223372036854775808 != -423) {
        if true {
            p0 = 18446744073709551615
        } else {
            for fi9 := 0; fi9 < 8; fi9++ {
                fmt.Println(p0, fi9, p2)
                var v10 uint16 = 32767
                _ = fi9
                _ = v10
            }
            fmt.Println(p1)
        }
    }
    v11 := []int{42, 1, 32768, 97}
    var v12 uint64 = p0
    for fi13 := 0; fi13 < 11; fi13++ {
        fmt.Println(p1)
        fmt.Println(p1, fi13)
        fmt.Println(v12, v12, fi13)
        _ = fi13
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v12
    return (^(v12))
}

func main() {
    v1 := []uint32{65535, 859}
    v2 := []uint32{4294967295}
    var v3 int16 = -1
    v4 := make(map[string]uint32)
    var v5 uint8 = 183
    switch 127 {
    case 720:
        v14 := GenS1{F0: v5, F1: 42}
        _ = v14
    case 162:
        if (!true) {
            switch "foo bar" {
            case "":
                v5 += v5
            case "hello":
                var v15 uint16 = 2
                _ = v15
            case "foo bar":
                v16 := GenS2{F0: uint(v5), F1: (!(384 <= v3))}
                _ = v16
            default:
                v17 := GenS3{F0: "foo bar", F1: -9223372036854775808, F2: 0}
                _ = v17
            }
        } else {
            fmt.Println(v3, v3)
        }
    case 127:
        v18 := GenS2{F0: 147, F1: false}
        _ = v18
    }
    v19 := GenS4{F0: 84, F1: 942}
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    for fi20 := 0; fi20 < 5; fi20++ {
        for fi21 := 0; fi21 < 15; fi21++ {
            switch (fi20 + 32767) {
            case 127:
                switch ("hello" + ("a" + "\n\t")) {
                case "\n\t":
                    var v22 uint32 = 2
                    _ = v22
                case "foo bar":
                    var v23 uint8 = v5
                    _ = v23
                default:
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                }
            case 128:
                var v24 float64 = ((float64(0.0) / ((((-2.25 / 3.14159) * ((3.14159 / 0.1) * (0.0 * 1e10))) + 1.5) + float64(1.5))) * float64((float32((0.0 / float64(1.5))) - 0.0)))
                _ = v24
            }
            if true {
                var v25 int8 = 1
                _ = v25
            }
            _ = fi21
        }
        _ = fi20
    }
    for fi26 := 0; fi26 < 9; fi26++ {
        var v27 uint32 = 14
        var v28 string = "a"
        if (!false) {
            switch (v28 + ((("" + v28) + (("a" + ("" + "foo bar")) + "a")) + ((v28 + "") + "foo bar"))) {
            case "hello":
                if len(v1) > 0 {
                    v1[0] = (^(v27))
                }
            }
        } else {
            fmt.Println(v27, v5, fi26)
            if (0 > -60) {
                var v29 float32 = 3.14159
                v28 = ((v28 + v28) + (v28 + (v28 + "foo bar")))
                for fi30 := 0; fi30 < 4; fi30++ {
                    var v31 bool = (42 < 0)
                    var v32 uint8 = 2
                    fmt.Println(v31, v5, v32)
                    v32 += v32
                    v33 := GenS3{F0: (v28 + v28), F1: 0, F2: 42}
                    v29 += (v29 - (v29 * v29))
                    _ = fi30
                    _ = v31
                    _ = v32
                    _ = v33
                }
                _ = v29
            } else {
                v4["\n\t"] = (v27 + v27)
                for fi34 := 0; fi34 < 10; fi34++ {
                    if false {
                        v5 -= (^(v5))
                        v35 := v4[v28]
                        var v36 bool = false
                        _ = v35
                        _ = v36
                    } else {
                        v28 = "\n\t"
                    }
                    if len(v1) > 0 {
                        v1[0] = (v27 >> ((42) & 63))
                    }
                    var v37 float32 = 1.5
                    v1 = append(v1, (v27 << ((18446744073709551615) & 63)))
                    _ = fi34
                    _ = v37
                }
            }
            for fi38 := 0; fi38 < 4; fi38++ {
                v3 = 1000
                _ = fi38
            }
        }
        _ = fi26
        _ = v27
        _ = v28
    }
    v5 += (v5 % ((v5) | 1))
    if false {
        for fi39 := 0; fi39 < 12; fi39++ {
            v3 = v3
            if (false && false) {
                var v40 string = ("foo bar" + "")
                v4[(v40 + "\n\t")] = 4294967295
                if true {
                    for fi41 := 0; fi41 < 14; fi41++ {
                        v1 = append(v1, 4294967295)
                        _ = fi41
                    }
                    if (42 > -9223372036854775808) {
                        v42 := GenS1{F0: v5, F1: 0}
                        v43 := GenS2{F0: 119, F1: (v40 <= "")}
                        _ = v42
                        _ = v43
                    }
                } else {
                    if false {
                        v40 = "\n\t"
                        if true {
                            v44 := GenS3{F0: ((v40 + v40) + (v40 + v40)), F1: 520, F2: genFn1(400, -82, 126)}
                            _ = v44
                        }
                    } else {
                        _ = len(v4)
                        for fi45 := 0; fi45 < 10; fi45++ {
                            var v46 int32 = 2147483647
                            var v47 int8 = -128
                            _ = fi45
                            _ = v46
                            _ = v47
                        }
                        v40 = ("hello" + ((v40 + v40) + ((v40 + "a") + v40)))
                    }
                }
                _ = v40
            }
            _ = fi39
        }
        v48 := GenS4{F0: v3, F1: 32767}
        _ = v48
    } else {
        fmt.Println(v5, v5)
        for fi49 := 0; fi49 < 12; fi49++ {
            var v50 uint16 = 2
            v50 = v50
            _ = fi49
            _ = v50
        }
    }
    if len(v2) > 0 {
        v2[0] = 2147483648
    }
    var v51 int32 = 0
    v1 = append(v1, 843)
    v51 -= (v51 % ((1) | 1))
    if false {
        switch 140 {
        case 127:
            var v52 bool = (!(!true))
            _ = v52
        case -2147483648:
            var v53 int32 = 0
            _ = v53
        }
        switch "foo bar" {
        case "\n\t":
            var v54 int32 = 42
            _ = v54
        case "a":
            switch 0 {
            case 4294967295:
                if (false && true) {
                    if (42 <= int64(v5)) {
                        v5 -= (v5 / ((128) | 1))
                        var v55 uint = 920
                        _ = v55
                    } else {
                        var v56 uint8 = 255
                        switch "foo bar" {
                        case "":
                            if false {
                                v3 -= (-(v3))
                                var v57 int = 9223372036854775807
                                var v58 int = v57
                                _ = v57
                                _ = v58
                            }
                        case "hello":
                            delete(v4, "hello")
                        case "a":
                            if false {
                                switch ("" + "") {
                                case "":
                                    var v59 string = ("a" + "hello")
                                    _ = v59
                                case "\n\t":
                                    if true {
                                        switch ("" + "\n\t") {
                                        case "hello":
                                            var v60 float64 = ((0.0 + float64(0.1)) + float64(float32(float64(float32(1.5)))))
                                            _ = v60
                                        case "a":
                                            var v61 bool = (100 > 1)
                                            _ = v61
                                        case "":
                                            v51 -= (v51 - (v51 - v51))
                                        default:
                                            for fi62 := 0; fi62 < 8; fi62++ {
                                                for fi63 := 0; fi63 < 3; fi63++ {
                                                    v4[("hello" + "\n\t")] = 2147483648
                                                    delete(v4, (("\n\t" + ("\n\t" + ("\n\t" + "foo bar"))) + ("" + "foo bar")))
                                                    _ = fi63
                                                }
                                                v3 -= v3
                                                _ = fi62
                                            }
                                        }
                                        if ((v56 >> ((840) & 63)) >= v5) {
                                            var v64 int32 = (v51 % (((v51 / (((v51 * v51)) | 1))) | 1))
                                            var v65 float64 = float64(3.14159)
                                            _ = v64
                                            _ = v65
                                        } else {
                                            for fi66 := 0; fi66 < 18; fi66++ {
                                                var v67 uint8 = v5
                                                var v68 uint8 = (v5 | (v56 << ((0) & 63)))
                                                _ = fi66
                                                _ = v67
                                                _ = v68
                                            }
                                        }
                                        fmt.Println(v3, v5, v51)
                                    }
                                default:
                                    var v69 uint16 = 1
                                    _ = v69
                                }
                                for fi70 := 0; fi70 < 9; fi70++ {
                                    delete(v4, "a")
                                    _ = fi70
                                }
                            } else {
                                for fi71 := 0; fi71 < 18; fi71++ {
                                    for fi72 := 0; fi72 < 19; fi72++ {
                                        fmt.Println(v51)
                                        v56 -= (v56 * v5)
                                        fmt.Println(v51)
                                        _ = fi72
                                    }
                                    if (1 != genFn0(v5, ("a" + "hello"), false)) {
                                        v73 := v4["\n\t"]
                                        if false {
                                            var v74 uint64 = 0
                                            var v75 uint32 = (^(v73))
                                            v56 -= (v5 >> ((1) & 63))
                                            _ = v74
                                            _ = v75
                                        } else {
                                            if false {
                                                if false {
                                                    for fi76 := 0; fi76 < 17; fi76++ {
                                                        v51 = (v51 ^ v51)
                                                        _ = fi76
                                                    }
                                                    if len(v2) >= 2 {
                                                        v2 = v2[1:len(v2)]
                                                    }
                                                } else {
                                                    if ((v56 | 140) == v5) {
                                                        v56 = v56
                                                        for fi77 := 0; fi77 < 10; fi77++ {
                                                            var v78 int = (fi71 << ((2) & 63))
                                                            _ = fi77
                                                            _ = v78
                                                        }
                                                        v51 = v51
                                                    } else {
                                                        if (v51 <= 576) {
                                                            var v79 uint32 = v73
                                                            var v80 int16 = 0
                                                            for fi81 := 0; fi81 < 17; fi81++ {
                                                                var v82 int32 = 397
                                                                fmt.Println(v79)
                                                                _ = fi81
                                                                _ = v82
                                                            }
                                                            _ = v79
                                                            _ = v80
                                                        }
                                                        fmt.Println(v73, v73)
                                                        v73 += genFn0((v5 / (((^(v56))) | 1)), (("foo bar" + "") + "hello"), false)
                                                    }
                                                    switch "\n\t" {
                                                    case "hello":
                                                        var v83 string = "foo bar"
                                                        _ = v83
                                                    case "a":
                                                        var v84 int = (fi71 >> ((0) & 63))
                                                        _ = v84
                                                    default:
                                                        v85 := GenS3{F0: ((("" + ("hello" + ("" + ""))) + ((((("hello" + "") + ("\n\t" + "foo bar")) + "hello") + ("" + "foo bar")) + "a")) + (("hello" + (("\n\t" + "\n\t") + ("\n\t" + ("a" + ("foo bar" + ""))))) + "\n\t")), F1: int64(v3), F2: 2}
                                                        _ = v85
                                                    }
                                                    var v86 int8 = 127
                                                    _ = v86
                                                }
                                                _ = len(v4)
                                            }
                                        }
                                        _ = v73
                                    }
                                    var v87 string = ""
                                    _ = fi71
                                    _ = v87
                                }
                            }
                        }
                        var v88 bool = (-29 >= 100)
                        _ = v56
                        _ = v88
                    }
                } else {
                    _ = len(v4)
                    var v89 uint16 = 0
                    _ = v89
                }
            default:
                switch ("" + "foo bar") {
                case "hello":
                    for fi90 := 0; fi90 < 2; fi90++ {
                        var v91 uint64 = 1
                        _ = fi90
                        _ = v91
                    }
                }
            }
        }
    } else {
        var v92 int = 2147483647
        for fi93 := 0; fi93 < 16; fi93++ {
            v51 = v51
            _ = fi93
        }
        v92 = v92
        _ = v92
    }
    for fi94 := 0; fi94 < 4; fi94++ {
        v95 := GenS5{F0: "hello", F1: 4294967295, F2: 0}
        if len(v1) > 0 {
            v1[0] = 951
        }
        _ = fi94
        _ = v95
    }
    for fi96 := 0; fi96 < 5; fi96++ {
        var v97 uint = 735
        fmt.Println(v3)
        v98 := v4[("hello" + "foo bar")]
        _ = fi96
        _ = v97
        _ = v98
    }
    if true {
        v99 := GenS6{F0: "hello"}
        v100 := GenS7{F0: (((("" + ("" + (("" + (((("a" + "") + ("a" + ("a" + "hello"))) + "a") + "\n\t")) + (("a" + "hello") + "hello")))) + "") + ("foo bar" + "foo bar")) + ("a" + "foo bar")), F1: 777}
        _ = v100
        _ = v99
    } else {
        switch 32768 {
        case 32768:
            var v101 uint8 = (v5 ^ v5)
            _ = v101
        default:
            var v102 int8 = -1
            _ = v102
        }
        for fi103 := 0; fi103 < 13; fi103++ {
            v104 := GenS8{F0: "", F1: (v3 << ((156) & 63)), F2: 783}
            var v105 bool = ((!(!((1 <= 1) && true))) || (65535 <= 2))
            _ = fi103
            _ = v104
            _ = v105
        }
        var v106 uint64 = 0
        _ = v106
    }
    if ((v5 > v5) && (false || false)) {
        if (true && (true || false)) {
            for fi107 := 0; fi107 < 10; fi107++ {
                fmt.Println(v5)
                var v108 uint8 = v5
                _ = fi107
                _ = v108
            }
            delete(v4, ("" + (("foo bar" + "foo bar") + "a")))
            for fi109 := 0; fi109 < 7; fi109++ {
                v110 := GenS9{F0: 1, F1: -128}
                _ = fi109
                _ = v110
            }
        }
    } else {
        _ = len(v1)
        v111 := GenS6{F0: "hello"}
        _ = v111
    }
    var v112 uint64 = 836
    for fi113 := 0; fi113 < 13; fi113++ {
        var v114 uint = 42
        _ = fi113
        _ = v114
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v112:", v112)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v51:", v51)
    _ = v1
    _ = v112
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v51
}
