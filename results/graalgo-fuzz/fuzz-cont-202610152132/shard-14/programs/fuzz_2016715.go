package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64 }
type GenS2 struct { F0 uint; F1 uint32; F2 int8 }
type GenS3 struct { F0 uint32; F1 int8 }
type GenS4 struct { F0 int32; F1 int64 }
type GenS5 struct { F0 uint8; F1 uint16 }
type GenS6 struct { F0 int8; F1 bool }
type GenS7 struct { F0 int; F1 string }
func genFn0() int8 {
    var v7 int8 = 100
    for fi8 := 0; fi8 < 11; fi8++ {
        v9 := GenS1{F0: 96, F1: 2}
        switch fi8 {
        case 127:
            var v10 int8 = 99
            _ = v10
        case -2147483648:
            switch fi8 {
            case 36:
                if true {
                    var v11 string = ("hello" + "")
                    fmt.Println(v7, fi8, v7)
                    _ = v11
                } else {
                    fmt.Println(fi8, fi8)
                }
            default:
                v12 := GenS2{F0: 2, F1: 0, F2: v7}
                _ = v12
            }
        }
        var v13 int8 = -128
        _ = fi8
        _ = v13
        _ = v9
    }
    fmt.Println(v7)
    fmt.Println(v7)
    v14 := GenS3{F0: 2147483648, F1: v7}
    _ = v14
    _ = v7
    return v7
}
func genFn1(p0 string, p1 uint8) uint8 {
    for fi15 := 0; fi15 < 4; fi15++ {
        switch "hello" {
        case "":
            if true {
                v16 := GenS1{F0: (p1 + (p1 * p1)), F1: 0}
                p0 = p0
                p0 = p0
                _ = v16
            } else {
                v17 := GenS1{F0: p1, F1: 829}
                fmt.Println(p1)
                switch (-(fi15)) {
                case 299:
                    if ((false && ((p0 + p0) < p0)) && (50 >= genFn0())) {
                        var v18 uint = 2
                        if (false && ((-836 <= 907) && true)) {
                            if false {
                                if true {
                                    var v19 int16 = 42
                                    v20 := GenS4{F0: 42, F1: -9223372036854775808}
                                    _ = v19
                                    _ = v20
                                } else {
                                    var v21 uint64 = 4294967295
                                    _ = v21
                                }
                            }
                            for fi22 := 0; fi22 < 2; fi22++ {
                                var v23 uint = uint(p1)
                                _ = fi22
                                _ = v23
                            }
                            if (575 < 62) {
                                var v24 uint16 = 65535
                                p0 = p0
                                _ = v24
                            } else {
                                fmt.Println(p0, v18)
                                fmt.Println(v18)
                                v25 := GenS3{F0: 1, F1: 42}
                                _ = v25
                            }
                        } else {
                            var v26 uint16 = uint16(v18)
                            v27 := GenS2{F0: (^(v18)), F1: 2147483648, F2: 127}
                            var v28 uint64 = 18446744073709551615
                            _ = v26
                            _ = v27
                            _ = v28
                        }
                        _ = v18
                    }
                case 32768:
                    var v29 uint16 = 562
                    _ = v29
                case -9223372036854775808:
                    fmt.Println(p1, p0, p1)
                default:
                    if (1 > 609) {
                        v30 := GenS5{F0: p1, F1: 255}
                        p0 = p0
                        switch (p0 + ((("a" + ("hello" + p0)) + "") + (p0 + (p0 + p0)))) {
                        case "\n\t":
                            for fi31 := 0; fi31 < 18; fi31++ {
                                p1 -= 127
                                p1 += p1
                                _ = fi31
                            }
                        default:
                            var v32 uint64 = 42
                            _ = v32
                        }
                        _ = v30
                    } else {
                        if (((54 != 787) || (!true)) || (((p0 + ("hello" + "foo bar")) >= "a") && ((255 < 760) || ((!false) || false)))) {
                            p1 = p1
                        } else {
                            p0 = ((("foo bar" + "foo bar") + p0) + ("" + "hello"))
                        }
                        fmt.Println(fi15)
                    }
                }
                _ = v17
            }
        case "hello":
            var v33 string = p0
            _ = v33
        case "\n\t":
            for fi34 := 0; fi34 < 18; fi34++ {
                if true {
                    if (((p0 + "foo bar") == "hello") || (-1 == 42)) {
                        v35 := []int{(fi34 & fi34), (fi15 | (fi34 + (fi15 << ((970) & 63)))), (fi15 * fi34), fi15}
                        if (281 != (fi15 | (fi15 ^ fi15))) {
                            var v36 uint = 768
                            if len(v35) > 0 {
                                v35[0] = fi15
                            }
                            _ = v36
                        } else {
                            var v37 uint16 = 0
                            v38 := GenS6{F0: 1, F1: ((v37 != 552) && false)}
                            p1 += (p1 ^ (^(p1)))
                            _ = v37
                            _ = v38
                        }
                        fmt.Println(fi15)
                        _ = v35
                    } else {
                        var v39 float64 = (3.14159 + ((1.5 / 1.0) * ((float64(3.14159) - (0.0 + 1e10)) / 1.0)))
                        p1 = (^(p1))
                        var v40 uint16 = 255
                        _ = v39
                        _ = v40
                    }
                    var v41 int16 = -935
                    if ((42 < 42) || false) {
                        var v42 int8 = -1
                        var v43 int32 = -2147483648
                        _ = v42
                        _ = v43
                    } else {
                        p1 += p1
                    }
                    _ = v41
                }
                _ = fi34
            }
        default:
            var v44 string = (p0 + "foo bar")
            _ = v44
        }
        var v45 int8 = 1
        if true {
            v46 := []int{(-(fi15))}
            _ = v46
        } else {
            if false {
                p1 = 255
            } else {
                v47 := GenS6{F0: (v45 % (((v45 / ((0) | 1))) | 1)), F1: false}
                v48 := []int{(fi15 + fi15)}
                _ = v47
                _ = v48
            }
            fmt.Println(p1, p0)
        }
        _ = fi15
        _ = v45
    }
    var v49 int32 = 0
    var v50 int = -884
    _ = p0
    _ = p1
    _ = v49
    _ = v50
    return (p1 & 42)
}

func main() {
    var v1 int64 = 42
    v2 := make(map[string]int)
    var v3 uint16 = 65535
    var v4 float64 = float64((1e10 + ((0.1 * (1.5 - 1.5)) / 1.0)))
    var v5 uint = 125
    v6 := []uint32{4294967295, 804, 413}
    if (true && (3.14159 == (1.5 / ((v4 - float64(1.5)) / v4)))) {
        _ = len(v2)
        if false {
            if (0 >= 1) {
                if (!(1 >= genFn0())) {
                    v3 -= v3
                }
                fmt.Println(v3, v3)
            } else {
                if (662 == 1) {
                    fmt.Println(v3, v3, v5)
                    fmt.Println(v1, v4)
                    if (!true) {
                        v51 := GenS1{F0: 42, F1: 18446744073709551615}
                        var v52 uint8 = 68
                        for fi53 := 0; fi53 < 2; fi53++ {
                            for fi54 := 0; fi54 < 4; fi54++ {
                                v52 = 128
                                v6 = append(v6, 1)
                                _ = fi54
                            }
                            var v55 uint8 = (v52 ^ 1)
                            _ = fi53
                            _ = v55
                        }
                        _ = v51
                        _ = v52
                    } else {
                        for fi56 := 0; fi56 < 15; fi56++ {
                            for fi57 := 0; fi57 < 9; fi57++ {
                                if false {
                                    switch 9223372036854775807 {
                                    case -681:
                                        _ = len(v6)
                                    case 127:
                                        v58 := GenS7{F0: (-(fi57)), F1: "foo bar"}
                                        var v59 uint64 = 4294967295
                                        _ = v58
                                        _ = v59
                                    }
                                    v1 = -1
                                    v6 = append(v6, 712)
                                }
                                _ = fi57
                            }
                            _ = fi56
                        }
                        v5 = uint(v3)
                    }
                } else {
                    v4 = 0.1
                    if len(v6) > 0 {
                        v6[0] = 37
                    }
                }
                for fi60 := 0; fi60 < 2; fi60++ {
                    if false {
                        fmt.Println(v5, v1)
                        var v61 string = "\n\t"
                        var v62 int = (-(fi60))
                        _ = v61
                        _ = v62
                    }
                    _ = fi60
                }
                v63 := GenS5{F0: 93, F1: v3}
                _ = v63
            }
        }
        v64 := GenS3{F0: uint32(v5), F1: -39}
        _ = v64
    } else {
        var v65 int64 = (v1 ^ -681)
        delete(v2, (("foo bar" + "foo bar") + (("a" + ("a" + ("foo bar" + ("hello" + "\n\t")))) + "\n\t")))
        v65 -= v65
        _ = v65
    }
    for fi66 := 0; fi66 < 19; fi66++ {
        var v67 int16 = -32768
        _ = fi66
        _ = v67
    }
    v5 -= 4294967295
    v68 := GenS5{F0: genFn1("foo bar", 2), F1: (v3 | v3)}
    if false {
        if false {
            var v69 int16 = 1
            switch "foo bar" {
            case "a":
                var v70 float32 = (((0.1 - 1.5) - (0.1 / 0.1)) + -2.25)
                _ = v70
            case "":
                _ = len(v6)
            case "\n\t":
                for fi71 := 0; fi71 < 1; fi71++ {
                    var v72 uint64 = 0
                    _ = fi71
                    _ = v72
                }
            }
            _ = len(v6)
            _ = v69
        }
        var v73 int8 = 127
        var v74 int8 = int8(v5)
        _ = v73
        _ = v74
    }
    if true {
        for fi75 := 0; fi75 < 17; fi75++ {
            v76 := GenS6{F0: -1, F1: (!(int32(v5) != -2147483648))}
            fmt.Println(v4)
            _ = fi75
            _ = v76
        }
        if len(v6) > 0 {
            v6[0] = 2
        }
    } else {
        fmt.Println(v5)
        if false {
            switch (("\n\t" + ("\n\t" + "foo bar")) + "\n\t") {
            case "hello":
                var v77 string = ""
                _ = v77
            case "a":
                for fi78 := 0; fi78 < 11; fi78++ {
                    var v79 uint32 = 986
                    switch (fi78 * fi78) {
                    case -9223372036854775808:
                        for fi80 := 0; fi80 < 3; fi80++ {
                            if true {
                                var v81 string = ("" + "hello")
                                switch 42 {
                                case -687:
                                    if false {
                                        v1 += (v1 * (^(v1)))
                                    } else {
                                        switch (-(fi80)) {
                                        case -3:
                                            v2[(v81 + v81)] = 32767
                                        }
                                    }
                                default:
                                    var v82 int32 = 1
                                    _ = v82
                                }
                                fmt.Println(fi80)
                                _ = v81
                            }
                            _ = fi80
                        }
                    default:
                        for fi83 := 0; fi83 < 18; fi83++ {
                            if len(v6) > 0 {
                                v6[0] = (v79 + v79)
                            }
                            _ = fi83
                        }
                    }
                    _ = fi78
                    _ = v79
                }
            }
            for fi84 := 0; fi84 < 10; fi84++ {
                v85 := GenS4{F0: 2147483647, F1: (v1 % ((v1) | 1))}
                v4 = v4
                v3 -= v3
                _ = fi84
                _ = v85
            }
        } else {
            var v86 float32 = 1e10
            switch ((((("a" + ("a" + "hello")) + "a") + "\n\t") + "\n\t") + "\n\t") {
            case "hello":
                v4 = float64(0.0)
            case "\n\t":
                switch 803 {
                case -1:
                    if true {
                        v86 = (-2.25 - v86)
                        var v87 bool = false
                        var v88 int8 = genFn0()
                        _ = v87
                        _ = v88
                    }
                case -118:
                    v89 := GenS7{F0: 128, F1: "foo bar"}
                    _ = v89
                }
            }
            _ = v86
        }
    }
    if true {
        fmt.Println(v3, v3)
        v3 = (v3 - 65535)
    }
    for fi90 := 0; fi90 < 15; fi90++ {
        _ = len(v6)
        for fi91 := 0; fi91 < 10; fi91++ {
            fmt.Println(v3)
            v92 := GenS3{F0: 1, F1: 90}
            v4 = 0.1
            _ = fi91
            _ = v92
        }
        _ = fi90
    }
    for fi93 := 0; fi93 < 3; fi93++ {
        var v94 float32 = 1e10
        _ = fi93
        _ = v94
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v68
}
