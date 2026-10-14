package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8 }
type GenS2 struct { F0 bool; F1 string; F2 uint }
type GenS3 struct { F0 string; F1 string }
type GenS4 struct { F0 string; F1 string }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 bool; F1 int32 }
type GenS8 struct { F0 string; F1 int16 }
type GenS9 struct { F0 int8; F1 string }
func genFn0(p0 int32, p1 uint64, p2 int32) int8 {
    for fi4 := 0; fi4 < 8; fi4++ {
        for fi5 := 0; fi5 < 19; fi5++ {
            var v6 uint = 1
            p0 = -2147483648
            var v7 uint64 = p1
            _ = fi5
            _ = v6
            _ = v7
        }
        switch fi4 {
        case -744:
            if ((!false) && (143 > 791)) {
                var v8 int8 = -42
                v9 := GenS1{F0: 18446744073709551615, F1: 228}
                _ = v8
                _ = v9
            }
        case 32767:
            for fi10 := 0; fi10 < 10; fi10++ {
                p1 -= p1
                _ = fi10
            }
        case 2147483647:
            v11 := []int{9223372036854775807}
            _ = v11
        }
        _ = fi4
    }
    p1 = 504
    if false {
        p1 -= p1
    } else {
        fmt.Println(p0)
        if true {
            v12 := []int{4294967295, int(p2)}
            _ = v12
        } else {
            for fi13 := 0; fi13 < 11; fi13++ {
                v14 := []int{(fi13 * fi13), fi13}
                v15 := GenS1{F0: 69, F1: 199}
                p0 += p2
                _ = fi13
                _ = v14
                _ = v15
            }
            p1 += 2
            for fi16 := 0; fi16 < 1; fi16++ {
                if false {
                    var v17 int32 = (p0 * (p0 / ((0) | 1)))
                    v18 := GenS1{F0: 255, F1: 0}
                    _ = v17
                    _ = v18
                } else {
                    var v19 string = "a"
                    if true {
                        if false {
                            for fi20 := 0; fi20 < 10; fi20++ {
                                var v21 float64 = (float64(-2.25) / 0.1)
                                if false {
                                    fmt.Println(v21)
                                    v22 := []int{-364}
                                    fmt.Println(p0)
                                    _ = v22
                                } else {
                                    v21 += v21
                                    if (1e10 < -2.25) {
                                        var v23 uint = 18446744073709551615
                                        _ = v23
                                    } else {
                                        p1 -= p1
                                    }
                                }
                                v24 := []int{-9223372036854775808, 411, 0}
                                _ = fi20
                                _ = v21
                                _ = v24
                            }
                        } else {
                            for fi25 := 0; fi25 < 15; fi25++ {
                                p1 = p1
                                var v26 int16 = -1
                                if false {
                                    var v27 int32 = (p0 - 870)
                                    v28 := GenS1{F0: 42, F1: 232}
                                    _ = v27
                                    _ = v28
                                } else {
                                    v29 := []int{fi25, (fi16 << ((uint(p1)) & 63))}
                                    v30 := GenS1{F0: 2, F1: 255}
                                    _ = v29
                                    _ = v30
                                }
                                _ = fi25
                                _ = v26
                            }
                        }
                    }
                    _ = v19
                }
                p2 = (-(p0))
                _ = fi16
            }
        }
        for fi31 := 0; fi31 < 11; fi31++ {
            v32 := GenS1{F0: 152, F1: 38}
            if false {
                var v33 int16 = 1
                fmt.Println(p1, p0)
                switch "foo bar" {
                case "foo bar":
                    v34 := GenS2{F0: false, F1: ("" + "\n\t"), F2: 1}
                    _ = v34
                case "hello":
                    switch (fi31 << ((18446744073709551615) & 63)) {
                    case 770:
                        for fi35 := 0; fi35 < 13; fi35++ {
                            fmt.Println(p0, fi35)
                            v36 := []int{}
                            _ = fi35
                            _ = v36
                        }
                    case 42:
                        for fi37 := 0; fi37 < 3; fi37++ {
                            var v38 int64 = 112
                            _ = fi37
                            _ = v38
                        }
                    case 32767:
                        v39 := GenS3{F0: ((("hello" + "hello") + "a") + "foo bar"), F1: ("\n\t" + "\n\t")}
                        _ = v39
                    }
                default:
                    var v40 string = "foo bar"
                    _ = v40
                }
                _ = v33
            }
            var v41 int8 = 66
            _ = fi31
            _ = v32
            _ = v41
        }
    }
    _ = p0
    _ = p1
    _ = p2
    return -39
}
func genFn1() string {
    var v42 float64 = 1e10
    v43 := GenS3{F0: "hello", F1: (("foo bar" + "") + ("hello" + "hello"))}
    var v44 bool = true
    var v45 uint32 = 95
    _ = v42
    _ = v43
    _ = v44
    _ = v45
    return "\n\t"
}

func main() {
    v1 := []string{}
    var v2 int64 = 0
    var v3 string = ("foo bar" + (("hello" + "a") + "foo bar"))
    if true {
        if true {
            fmt.Println(v3)
        } else {
            v3 = "\n\t"
            var v46 int16 = 1000
            _ = v46
        }
        v2 += (v2 - (v2 ^ -9223372036854775808))
        if (257 < 72) {
            for fi47 := 0; fi47 < 13; fi47++ {
                v48 := GenS2{F0: false, F1: v3, F2: 18446744073709551615}
                _ = fi47
                _ = v48
            }
        }
    }
    if true {
        v49 := GenS2{F0: (!false), F1: "hello", F2: 354}
        v50 := GenS4{F0: ("hello" + (("\n\t" + v3) + v3)), F1: "\n\t"}
        v2 = -558
        _ = v49
        _ = v50
    }
    if ((v3 + "hello") > "a") {
        switch "\n\t" {
        case "foo bar":
            var v51 string = "\n\t"
            _ = v51
        case "hello":
            switch v3 {
            case "a":
                if false {
                    v52 := GenS5{F0: 42}
                    _ = v52
                }
            case "":
                for fi53 := 0; fi53 < 19; fi53++ {
                    v1 = append(v1, v3)
                    fmt.Println(v2, fi53, v3)
                    _ = fi53
                }
            }
        default:
            switch "a" {
            case "hello":
                var v54 uint32 = 2
                _ = v54
            case "\n\t":
                v2 -= 1
            default:
                if len(v1) > 0 {
                    v1[0] = ((v3 + (("foo bar" + v3) + v3)) + ("a" + "foo bar"))
                }
            }
        }
    } else {
        var v55 int = 127
        var v56 int = v55
        v1 = append(v1, v3)
        _ = v55
        _ = v56
    }
    v2 -= v2
    if true {
        var v57 float32 = 0.0
        for fi58 := 0; fi58 < 8; fi58++ {
            v3 = ""
            v57 += v57
            _ = fi58
        }
        _ = v57
    }
    var v59 int8 = 100
    if false {
        for fi60 := 0; fi60 < 19; fi60++ {
            for fi61 := 0; fi61 < 13; fi61++ {
                var v62 uint16 = 1
                _ = fi61
                _ = v62
            }
            v3 = (v3 + (("" + ((v3 + v3) + v3)) + "\n\t"))
            v59 -= (v59 >> ((18446744073709551615) & 63))
            _ = fi60
        }
        var v63 uint8 = 240
        _ = v63
    } else {
        fmt.Println(v3, v59, v3)
        if true {
            v59 += (v59 + (v59 << ((4294967295) & 63)))
            if false {
                switch v3 {
                case "foo bar":
                    var v64 int8 = (v59 / (((v59 | 0)) | 1))
                    _ = v64
                case "":
                    for fi65 := 0; fi65 < 10; fi65++ {
                        for fi66 := 0; fi66 < 11; fi66++ {
                            if ("" != "\n\t") {
                                switch "\n\t" {
                                case "\n\t":
                                    v67 := GenS4{F0: "foo bar", F1: "hello"}
                                    _ = v67
                                case "":
                                    fmt.Println(v3, v3)
                                    switch v3 {
                                    case "":
                                        if true {
                                            _ = len(v1)
                                            switch (v3 + (v3 + (v3 + v3))) {
                                            case "\n\t":
                                                var v68 uint16 = 263
                                                _ = v68
                                            case "a":
                                                _ = len(v1)
                                            default:
                                                v3 = (v3 + "foo bar")
                                            }
                                        } else {
                                            if false {
                                                _ = len(v1)
                                                v3 = "hello"
                                                var v69 uint16 = 2
                                                v1 = append(v1, v3)
                                                var v70 float64 = (3.14159 + float64(0.0))
                                                _ = v69
                                                _ = v70
                                            } else {
                                                var v71 uint64 = 42
                                                switch v3 {
                                                case "\n\t":
                                                    switch v3 {
                                                    case "\n\t":
                                                        v59 -= 42
                                                    case "hello":
                                                        v72 := GenS6{F0: 591}
                                                        v59 -= (v59 | 1)
                                                        _ = v72
                                                    }
                                                case "hello":
                                                    var v73 uint16 = 0
                                                    _ = v73
                                                default:
                                                    fmt.Println(v59, v3, v2)
                                                }
                                                _ = v71
                                            }
                                            v3 = v3
                                        }
                                    case "hello":
                                        v2 = v2
                                    case "a":
                                        if (0.0 >= 1.5) {
                                            var v74 uint16 = 65535
                                            var v75 uint8 = 1
                                            _ = v74
                                            _ = v75
                                        } else {
                                            v3 = v3
                                            if len(v1) > 0 {
                                                v1[0] = "hello"
                                            }
                                            if true {
                                                var v76 int = -506
                                                switch ((v3 + (((v3 + v3) + v3) + (v3 + v3))) + (v3 + "a")) {
                                                case "foo bar":
                                                    if len(v1) >= 2 {
                                                        v1 = v1[1:len(v1)]
                                                    }
                                                    v59 += v59
                                                    v76 = fi65
                                                case "a":
                                                    v3 = ""
                                                case "hello":
                                                    switch v3 {
                                                    case "":
                                                        var v77 int16 = -32768
                                                        _ = v77
                                                    }
                                                }
                                                _ = v76
                                            }
                                            v3 = (("" + v3) + v3)
                                        }
                                        switch (fi65 ^ 127) {
                                        case 2147483647:
                                            if (0.1 >= 0.1) {
                                                var v78 float32 = 1.5
                                                v3 = (v3 + "")
                                                var v79 uint64 = 476
                                                _ = v78
                                                _ = v79
                                            } else {
                                                v1 = append(v1, ("" + (("a" + v3) + v3)))
                                                fmt.Println(fi65, fi65, v3)
                                            }
                                        case -9223372036854775808:
                                            var v80 uint8 = 40
                                            _ = v80
                                        case 128:
                                            if (((1e10 * (-2.25 * 3.14159)) >= 1.5) || (false || (false || (((1.5 - 0.1) <= 0.1) && true)))) {
                                                v59 = genFn0(-905, 664, 710)
                                                var v81 int16 = -32768
                                                v82 := GenS3{F0: v3, F1: ("" + v3)}
                                                _ = v81
                                                _ = v82
                                            }
                                        }
                                    default:
                                        var v83 float32 = (0.1 * (3.14159 - 3.14159))
                                        _ = v83
                                    }
                                    v59 -= (v59 % (((v59 + v59)) | 1))
                                case "foo bar":
                                    fmt.Println(v59, fi66)
                                default:
                                    v3 = v3
                                }
                            } else {
                                var v84 uint16 = 53
                                if len(v1) > 0 {
                                    v1[0] = v3
                                }
                                v85 := GenS5{F0: (v84 + (v84 & (v84 % (((^(v84))) | 1))))}
                                _ = v84
                                _ = v85
                            }
                            if len(v1) > 0 {
                                v1[0] = "foo bar"
                            }
                            if true {
                                v86 := GenS7{F0: ((v59 ^ v59) >= v59), F1: 42}
                                _ = v86
                            } else {
                                if false {
                                    v87 := GenS5{F0: 84}
                                    v88 := GenS3{F0: v3, F1: v3}
                                    var v89 uint = 20
                                    _ = v87
                                    _ = v88
                                    _ = v89
                                }
                                v2 = v2
                            }
                            _ = fi66
                        }
                        v3 = "hello"
                        _ = fi65
                    }
                }
                for fi90 := 0; fi90 < 13; fi90++ {
                    if len(v1) > 0 {
                        v1[0] = "foo bar"
                    }
                    v91 := GenS8{F0: v3, F1: -986}
                    _ = fi90
                    _ = v91
                }
                var v92 uint8 = 2
                _ = v92
            }
        } else {
            v3 = "hello"
            var v93 int16 = 0
            v3 = (v3 + (v3 + "a"))
            _ = v93
        }
    }
    if false {
        fmt.Println(v59, v59, v2)
    }
    for fi94 := 0; fi94 < 12; fi94++ {
        v3 = (v3 + v3)
        fmt.Println(fi94, v2)
        v95 := GenS4{F0: "\n\t", F1: (("hello" + "") + (v3 + ((v3 + "") + v3)))}
        _ = fi94
        _ = v95
    }
    v96 := GenS9{F0: (v59 >> ((65535) & 63)), F1: "\n\t"}
    if len(v1) > 0 {
        v1[0] = v3
    }
    fmt.Println(v2, v59)
    var v97 string = v3
    v3 = v3
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v59:", v59)
    fmt.Println("v97:", v97)
    _ = v1
    _ = v2
    _ = v3
    _ = v59
    _ = v96
    _ = v97
}
