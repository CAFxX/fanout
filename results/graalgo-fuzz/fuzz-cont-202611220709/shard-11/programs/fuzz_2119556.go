package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int64 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint32; F1 bool; F2 int64 }
type GenS4 struct { F0 uint64; F1 bool }
type GenS5 struct { F0 bool; F1 int16 }
type GenS6 struct { F0 int32; F1 int }
func genFn0(p0 int8, p1 int64) uint16 {
    var v7 int8 = p0
    var v8 bool = (-2147483648 > 127)
    v9 := GenS1{F0: 42, F1: -9223372036854775808}
    for fi10 := 0; fi10 < 9; fi10++ {
        var v11 uint64 = 18446744073709551615
        _ = fi10
        _ = v11
    }
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    _ = v9
    return 1
}
func genFn1(p0 uint) uint64 {
    p0 = p0
    v12 := []int{-71, -9223372036854775808, 1}
    var v13 int32 = 780
    switch 2147483647 {
    case 666:
        for fi14 := 0; fi14 < 2; fi14++ {
            p0 -= 565
            _ = fi14
        }
    }
    switch (("a" + ((("foo bar" + "foo bar") + "a") + "a")) + ("" + "")) {
    case "hello":
        v13 -= (v13 + 1)
    default:
        if len(v12) > 0 {
            v12[0] = 128
        }
    }
    _ = p0
    _ = v12
    _ = v13
    return 42
}

func main() {
    var v1 int8 = -89
    var v2 uint32 = 2147483647
    var v3 int64 = -1
    v4 := []uint32{v2, (^(v2)), v2}
    var v5 float64 = (1.5 / 1.0)
    v6 := make(map[string]string)
    var v15 int8 = (v1 + (v1 * (v1 + v1)))
    if false {
        fmt.Println(v15, v1)
        if false {
            v1 -= 58
            _ = len(v6)
            if (false && (!(((!false) && ((v2 >> ((106) & 63)) < v2)) || false))) {
                for fi16 := 0; fi16 < 11; fi16++ {
                    v4 = append(v4, 409)
                    if false {
                        v4 = append(v4, v2)
                        fmt.Println(fi16, v5, v2)
                        v17 := GenS1{F0: 127, F1: (v3 + (-(v3)))}
                        _ = v17
                    } else {
                        switch fi16 {
                        case 2147483647:
                            switch "\n\t" {
                            case "foo bar":
                                v18 := GenS1{F0: 2, F1: (v3 * (-(v3)))}
                                _ = v18
                            case "\n\t":
                                if true {
                                    if (true && (18446744073709551615 <= 18446744073709551615)) {
                                        if true {
                                            v3 = (v3 % ((-873) | 1))
                                            var v19 float64 = v5
                                            if false {
                                                v3 = (v3 / (((v3 ^ 42)) | 1))
                                                switch "a" {
                                                case "a":
                                                    var v20 float32 = 1e10
                                                    _ = v20
                                                case "\n\t":
                                                    var v21 uint64 = 495
                                                    _ = v21
                                                case "hello":
                                                    if (!false) {
                                                        v15 = v1
                                                        if true {
                                                            fmt.Println(v3, v19, v5)
                                                            v5 += v19
                                                            for fi22 := 0; fi22 < 4; fi22++ {
                                                                var v23 uint32 = v2
                                                                v24 := GenS1{F0: 99, F1: (v3 << ((754) & 63))}
                                                                var v25 float32 = 1e10
                                                                v6[("foo bar" + ("a" + "\n\t"))] = ""
                                                                var v26 uint16 = 42
                                                                _ = fi22
                                                                _ = v23
                                                                _ = v24
                                                                _ = v25
                                                                _ = v26
                                                            }
                                                        }
                                                    } else {
                                                        var v27 int64 = (v3 % (((v3 << ((1) & 63))) | 1))
                                                        _ = v27
                                                    }
                                                }
                                                fmt.Println(v5)
                                            } else {
                                                var v28 int32 = -2147483648
                                                switch (fi16 * -9223372036854775808) {
                                                case -939:
                                                    for fi29 := 0; fi29 < 13; fi29++ {
                                                        v6["a"] = ""
                                                        _ = fi29
                                                    }
                                                }
                                                var v30 string = "a"
                                                _ = v28
                                                _ = v30
                                            }
                                            _ = v19
                                        }
                                        var v31 uint8 = 2
                                        var v32 bool = false
                                        _ = v31
                                        _ = v32
                                    }
                                    v3 = -1
                                } else {
                                    var v33 int = (fi16 + (fi16 / ((fi16) | 1)))
                                    var v34 uint = 0
                                    v35 := GenS1{F0: 128, F1: v3}
                                    _ = v33
                                    _ = v34
                                    _ = v35
                                }
                            default:
                                v15 = (v15 + (v15 + v15))
                            }
                        case -9223372036854775808:
                            for fi36 := 0; fi36 < 8; fi36++ {
                                var v37 int = (fi36 >> ((1) & 63))
                                v38 := GenS1{F0: 1, F1: v3}
                                _ = fi36
                                _ = v37
                                _ = v38
                            }
                        default:
                            var v39 uint8 = 134
                            _ = v39
                        }
                        var v40 int32 = 1
                        if (575 >= 4) {
                            v41 := GenS2{F0: (2 < (v2 >> ((0) & 63)))}
                            _ = v41
                        }
                        _ = v40
                    }
                    _ = fi16
                }
                _ = len(v4)
                var v42 uint32 = (^(v2))
                _ = v42
            } else {
                fmt.Println(v3, v1, v5)
                if false {
                    v5 -= v5
                    for fi43 := 0; fi43 < 2; fi43++ {
                        fmt.Println(v3, v3)
                        v2 -= v2
                        if ((-(v1)) >= 117) {
                            _ = len(v6)
                            v44 := GenS2{F0: ((783 > genFn0((^(v1)), v3)) && false)}
                            var v45 int16 = -546
                            _ = v44
                            _ = v45
                        } else {
                            v3 += v3
                            if true {
                                if (v2 < (v2 % ((v2) | 1))) {
                                    v46 := GenS2{F0: true}
                                    _ = v46
                                } else {
                                    switch (fi43 << ((65535) & 63)) {
                                    case 42:
                                        switch "foo bar" {
                                        case "hello":
                                            if (111 == 255) {
                                                _ = len(v4)
                                                if (42 < 1) {
                                                    delete(v6, ("" + (("\n\t" + ("\n\t" + "")) + "a")))
                                                    v47 := GenS3{F0: (^(v2)), F1: false, F2: (v3 / (((-(v3))) | 1))}
                                                    _ = v47
                                                } else {
                                                    if (false && ("a" <= "hello")) {
                                                        _ = len(v4)
                                                        v15 -= -10
                                                        v48 := GenS2{F0: (!((v2 % ((v2) | 1)) <= v2))}
                                                        _ = v48
                                                    }
                                                }
                                                var v49 string = "a"
                                                _ = v49
                                            } else {
                                                if false {
                                                    switch fi43 {
                                                    case 9223372036854775807:
                                                        switch fi43 {
                                                        case 896:
                                                            if (271 >= 18446744073709551615) {
                                                                v50 := GenS4{F0: 930, F1: true}
                                                                delete(v6, (("foo bar" + "\n\t") + "hello"))
                                                                _ = v50
                                                            } else {
                                                                fmt.Println(v5, fi43, v15)
                                                                v51 := GenS5{F0: true, F1: -306}
                                                                var v52 int8 = 1
                                                                _ = v51
                                                                _ = v52
                                                            }
                                                        case 128:
                                                            var v53 float64 = v5
                                                            _ = v53
                                                        default:
                                                            switch 0 {
                                                            case 32768:
                                                                v54 := GenS5{F0: (!(((v3 & v3) != v3) && (true && true))), F1: 42}
                                                                _ = v54
                                                            case 1:
                                                                delete(v6, (("hello" + "foo bar") + "hello"))
                                                            }
                                                        }
                                                    case -680:
                                                        fmt.Println(v15, v5, v1)
                                                    default:
                                                        _ = len(v4)
                                                    }
                                                    v15 -= v1
                                                } else {
                                                    var v55 uint32 = 245
                                                    var v56 int64 = (v3 * (v3 - (v3 ^ v3)))
                                                    for fi57 := 0; fi57 < 9; fi57++ {
                                                        v58 := GenS4{F0: uint64(v55), F1: false}
                                                        v3 = (v3 >> ((117) & 63))
                                                        var v59 int64 = (-(v3))
                                                        _ = fi57
                                                        _ = v58
                                                        _ = v59
                                                    }
                                                    _ = v55
                                                    _ = v56
                                                }
                                                for fi60 := 0; fi60 < 19; fi60++ {
                                                    switch (fi43 | fi60) {
                                                    case -817:
                                                        v2 += v2
                                                    case -9223372036854775808:
                                                        v61 := GenS1{F0: 107, F1: v3}
                                                        _ = v61
                                                    }
                                                    _ = fi60
                                                }
                                                v15 = v1
                                            }
                                        case "\n\t":
                                            switch fi43 {
                                            case -2147483648:
                                                v62 := v6[((("a" + "foo bar") + ("" + "hello")) + "hello")]
                                                _ = v62
                                            }
                                        default:
                                            if (!true) {
                                                var v63 int16 = 1
                                                var v64 string = "foo bar"
                                                _ = v63
                                                _ = v64
                                            }
                                        }
                                    case -2147483648:
                                        v2 = (v2 * (^(v2)))
                                    }
                                }
                            } else {
                                var v65 float64 = float64((0.1 / float32(v5)))
                                switch "a" {
                                case "foo bar":
                                    if false {
                                        v3 -= 1
                                    } else {
                                        v66 := v6["a"]
                                        v67 := GenS4{F0: 148, F1: false}
                                        _ = v66
                                        _ = v67
                                    }
                                case "hello":
                                    var v68 int = (fi43 | 0)
                                    _ = v68
                                case "a":
                                    for fi69 := 0; fi69 < 12; fi69++ {
                                        if (-53 > -2147483648) {
                                            switch (fi43 * (fi43 - int(v3))) {
                                            case 0:
                                                fmt.Println(fi43, v5, v15)
                                            case 2147483647:
                                                if (1 < genFn1(4294967295)) {
                                                    var v70 float32 = -2.25
                                                    _ = v70
                                                }
                                            case 1:
                                                var v71 uint32 = v2
                                                _ = v71
                                            }
                                            v5 += v5
                                        } else {
                                            if true {
                                                if false {
                                                    switch 414 {
                                                    case -9223372036854775808:
                                                        if (v65 <= v65) {
                                                            var v72 int64 = 910
                                                            var v73 uint16 = 0
                                                            _ = v72
                                                            _ = v73
                                                        }
                                                    case -765:
                                                        v15 = v1
                                                    default:
                                                        var v74 bool = true
                                                        _ = v74
                                                    }
                                                } else {
                                                    v2 -= v2
                                                    v75 := GenS2{F0: false}
                                                    var v76 int8 = (v15 - v1)
                                                    _ = v75
                                                    _ = v76
                                                }
                                            }
                                            switch "foo bar" {
                                            case "a":
                                                var v77 float32 = float32(((((v5 * 0.1) + (v5 / v65)) - v5) / float64(1.5)))
                                                _ = v77
                                            case "hello":
                                                var v78 bool = (1000 != 42)
                                                _ = v78
                                            }
                                            fmt.Println(v1)
                                            var v79 uint16 = 65535
                                            var v80 int8 = v1
                                            v4 = append(v4, v2)
                                            fmt.Println(v2)
                                            _ = v79
                                            _ = v80
                                        }
                                        v1 += (v15 & (v15 << ((255) & 63)))
                                        _ = len(v6)
                                        _ = fi69
                                    }
                                default:
                                    var v81 int32 = 1
                                    _ = v81
                                }
                                _ = v65
                            }
                            switch "" {
                            case "a":
                                v82 := GenS1{F0: 255, F1: -569}
                                _ = v82
                            case "":
                                v5 -= 0.1
                            }
                        }
                        _ = fi43
                    }
                    v83 := GenS4{F0: 0, F1: (2147483647 < 1)}
                    _ = v83
                }
                var v84 uint8 = 2
                _ = v84
            }
        } else {
            v3 -= v3
        }
        v85 := GenS5{F0: true, F1: -921}
        _ = v85
    } else {
        v86 := GenS3{F0: (v2 ^ v2), F1: true, F2: (v3 - v3)}
        _ = v86
    }
    if false {
        var v87 uint = 255
        _ = v87
    }
    switch ("\n\t" + "") {
    case "a":
        var v88 float64 = (v5 * (v5 + float64((3.14159 * float32(v5)))))
        _ = v88
    case "\n\t":
        var v89 int64 = v3
        _ = v89
    }
    v90 := v6["a"]
    v91 := GenS6{F0: 883, F1: 42}
    var v92 int = -9223372036854775808
    var v93 int16 = 561
    v4 = append(v4, (v2 % (((^(v2))) | 1)))
    v92 = v92
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v90:", v90)
    fmt.Println("v92:", v92)
    fmt.Println("v93:", v93)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v90
    _ = v91
    _ = v92
    _ = v93
}
