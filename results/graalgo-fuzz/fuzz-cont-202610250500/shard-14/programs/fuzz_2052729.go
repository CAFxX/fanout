package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint32; F1 uint32 }
type GenS4 struct { F0 bool; F1 uint16 }
type GenS5 struct { F0 uint32; F1 int32; F2 uint16 }
type GenS6 struct { F0 uint16; F1 uint16 }
type GenS7 struct { F0 uint32; F1 uint16 }
type GenS8 struct { F0 bool }
type GenS9 struct { F0 int; F1 string; F2 int64 }
type GenS10 struct { F0 int64; F1 int32 }
func genFn0(p0 uint32, p1 string, p2 bool) uint64 {
    v5 := GenS1{F0: 128}
    v6 := GenS1{F0: 479}
    if true {
        for fi7 := 0; fi7 < 16; fi7++ {
            for fi8 := 0; fi8 < 15; fi8++ {
                var v9 float32 = 0.1
                _ = fi8
                _ = v9
            }
            if true {
                p0 += p0
                p1 = ("a" + ((p1 + p1) + ""))
                p2 = (("a" <= p1) || ((^(p0)) > p0))
            }
            _ = fi7
        }
        v10 := []int{}
        v11 := GenS2{F0: p0}
        _ = v10
        _ = v11
    } else {
        v12 := []int{0, 610, -1, -90}
        _ = v12
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return uint64(p0)
}

func main() {
    var v1 uint64 = 4294967295
    var v2 float64 = 3.14159
    v3 := make(map[int]bool)
    var v4 int = 977
    delete(v3, (-(v4)))
    var v13 int = 32767
    var v14 int = (v4 | -9223372036854775808)
    v4 += v13
    for fi15 := 0; fi15 < 9; fi15++ {
        var v16 uint8 = 0
        var v17 uint64 = 2
        switch int(v17) {
        case 32768:
            if (v16 != v16) {
                for fi18 := 0; fi18 < 8; fi18++ {
                    v14 += int(v17)
                    _ = fi18
                }
                if (0 == 1) {
                    v19 := GenS1{F0: (-(v13))}
                    v20 := GenS2{F0: uint32(v16)}
                    _ = v19
                    _ = v20
                } else {
                    switch (("a" + ("a" + "foo bar")) + (("\n\t" + "") + "hello")) {
                    case "a":
                        var v21 bool = (255 != 42)
                        _ = v21
                    default:
                        v22 := GenS1{F0: (v13 / (((^(fi15))) | 1))}
                        _ = v22
                    }
                }
            } else {
                var v23 string = ((("a" + ("foo bar" + "\n\t")) + "\n\t") + "\n\t")
                _ = v23
            }
        case 446:
            fmt.Println(v16, v1, v2)
        default:
            var v24 int64 = 9223372036854775807
            _ = v24
        }
        _ = fi15
        _ = v16
        _ = v17
    }
    if ((-1 > (v4 / ((v13) | 1))) && true) {
        if false {
            var v25 int64 = -9223372036854775808
            if ("foo bar" >= (("" + ("foo bar" + "a")) + "hello")) {
                for fi26 := 0; fi26 < 2; fi26++ {
                    for fi27 := 0; fi27 < 17; fi27++ {
                        v2 -= (v2 - (0.1 - (v2 - 0.1)))
                        v13 = 42
                        _ = fi27
                    }
                    v4 = int(v1)
                    _ = fi26
                }
                var v28 int16 = -878
                v29 := []int{4294967295, v14, v14}
                _ = v28
                _ = v29
            } else {
                v2 -= ((v2 * v2) - (v2 * v2))
                var v30 uint64 = (v1 >> ((18446744073709551615) & 63))
                v31 := GenS1{F0: (v4 >> ((422) & 63))}
                _ = v30
                _ = v31
            }
            _ = v25
        } else {
            for fi32 := 0; fi32 < 17; fi32++ {
                fmt.Println(v2)
                v33 := v3[v14]
                var v34 float32 = -2.25
                _ = fi32
                _ = v33
                _ = v34
            }
            v2 = (0.1 / (1e10 / 1e10))
            delete(v3, v14)
        }
        _ = len(v3)
        v35 := v3[(v14 - v4)]
        _ = v35
    } else {
        fmt.Println(v2, v4, v4)
        fmt.Println(v2)
        var v36 string = ("hello" + "hello")
        _ = v36
    }
    v3[(v14 * int(v1))] = true
    var v37 string = "hello"
    if false {
        if (!(!(0.1 < 3.14159))) {
            v38 := GenS2{F0: 828}
            v39 := GenS3{F0: 2, F1: 2}
            _ = v38
            _ = v39
        } else {
            if ((v37 > ("\n\t" + ((v37 + "") + v37))) || false) {
                for fi40 := 0; fi40 < 13; fi40++ {
                    if true {
                        v4 = (fi40 | 32)
                        for fi41 := 0; fi41 < 6; fi41++ {
                            switch fi41 {
                            case 142:
                                var v42 int32 = 17
                                _ = v42
                            case 2147483647:
                                var v43 string = (("" + "foo bar") + v37)
                                var v44 int64 = 9223372036854775807
                                _ = v43
                                _ = v44
                            }
                            v45 := GenS4{F0: (!(-128 != -128)), F1: 407}
                            _ = fi41
                            _ = v45
                        }
                        if false {
                            if ((-128 <= 0) || false) {
                                if ((128 < 128) && false) {
                                    var v46 float32 = (float32(-2.25) * (-2.25 * 1.5))
                                    var v47 uint64 = v1
                                    for fi48 := 0; fi48 < 6; fi48++ {
                                        var v49 uint64 = (v1 * (v1 & (v1 ^ genFn0(4294967295, v37, true))))
                                        _ = fi48
                                        _ = v49
                                    }
                                    _ = v46
                                    _ = v47
                                }
                                for fi50 := 0; fi50 < 11; fi50++ {
                                    if (255 != 42) {
                                        _ = len(v3)
                                        fmt.Println(v14)
                                        delete(v3, v13)
                                    } else {
                                        v51 := GenS5{F0: 42, F1: -1, F2: 1}
                                        var v52 uint32 = 0
                                        var v53 int32 = -1
                                        v54 := []int{}
                                        _ = v51
                                        _ = v52
                                        _ = v53
                                        _ = v54
                                    }
                                    _ = fi50
                                }
                                for fi55 := 0; fi55 < 13; fi55++ {
                                    var v56 string = ""
                                    v37 = "hello"
                                    _ = fi55
                                    _ = v56
                                }
                            } else {
                                v2 = (3.14159 + -2.25)
                                if true {
                                    v4 -= fi40
                                }
                            }
                            if true {
                                for fi57 := 0; fi57 < 10; fi57++ {
                                    var v58 int64 = -9223372036854775808
                                    if false {
                                        fmt.Println(v13, fi57)
                                    } else {
                                        v59 := GenS5{F0: 960, F1: 557, F2: 714}
                                        v37 = v37
                                        _ = v59
                                    }
                                    _ = fi57
                                    _ = v58
                                }
                            } else {
                                fmt.Println(v14, v37, v14)
                                v60 := GenS2{F0: 42}
                                _ = v60
                            }
                        }
                    } else {
                        delete(v3, (v14 << ((0) & 63)))
                    }
                    _ = fi40
                }
                var v61 float32 = (0.1 + (((-2.25 / ((-2.25 / 3.14159) + 0.0)) / (-2.25 / 1.0)) - 1e10))
                _ = v61
            }
        }
        if true {
            var v62 uint = 1
            fmt.Println(v62, v1, v1)
            v63 := GenS6{F0: 65535, F1: 42}
            _ = v62
            _ = v63
        } else {
            if (!(v1 == (^(v1)))) {
                for fi64 := 0; fi64 < 19; fi64++ {
                    v14 -= 239
                    var v65 uint8 = 50
                    _ = fi64
                    _ = v65
                }
            } else {
                if false {
                    if true {
                        var v66 float64 = v2
                        if (true && false) {
                            switch v37 {
                            case "a":
                                v4 = (v13 << ((18446744073709551615) & 63))
                            default:
                                fmt.Println(v13, v1, v4)
                            }
                        } else {
                            if (!(true && ((false && (-2.25 > 3.14159)) && true))) {
                                for fi67 := 0; fi67 < 2; fi67++ {
                                    v4 = 32767
                                    _ = fi67
                                }
                            } else {
                                v68 := GenS7{F0: 42, F1: 255}
                                _ = v68
                            }
                            switch (("\n\t" + "\n\t") + "foo bar") {
                            case "":
                                v69 := GenS3{F0: uint32(v1), F1: 211}
                                _ = v69
                            case "foo bar":
                                var v70 uint = 957
                                _ = v70
                            case "a":
                                switch ("hello" + v37) {
                                case "":
                                    var v71 float64 = v2
                                    _ = v71
                                case "hello":
                                    for fi72 := 0; fi72 < 12; fi72++ {
                                        if (v4 == (v13 / ((42) | 1))) {
                                            v73 := GenS3{F0: 2147483647, F1: 0}
                                            v74 := []int{}
                                            _ = v73
                                            _ = v74
                                        }
                                        switch fi72 {
                                        case -9223372036854775808:
                                            if (uint32(v1) < 2) {
                                                delete(v3, (v4 + 1))
                                            }
                                        case 2147483647:
                                            if true {
                                                for fi75 := 0; fi75 < 7; fi75++ {
                                                    v76 := GenS6{F0: 0, F1: 255}
                                                    switch v37 {
                                                    case "a":
                                                        switch (fi75 / ((v4) | 1)) {
                                                        case 398:
                                                            fmt.Println(v2, v13)
                                                        case 641:
                                                            v77 := GenS4{F0: (false && (!false)), F1: 42}
                                                            var v78 uint64 = 116
                                                            v3[(fi72 / ((int(v1)) | 1))] = (127 == 127)
                                                            _ = v77
                                                            _ = v78
                                                        }
                                                    }
                                                    var v79 int8 = -128
                                                    _ = fi75
                                                    _ = v76
                                                    _ = v79
                                                }
                                                var v80 uint64 = genFn0(2, ((v37 + ("a" + v37)) + v37), (float64(3.14159) <= v2))
                                                v4 = (v14 + v13)
                                                _ = v80
                                            } else {
                                                if (((v37 + (v37 + v37)) >= "\n\t") || (-110 < 42)) {
                                                    delete(v3, v13)
                                                    var v81 uint8 = 231
                                                    _ = v81
                                                } else {
                                                    for fi82 := 0; fi82 < 18; fi82++ {
                                                        v83 := []int{-9223372036854775808, (v4 * (v4 + v13)), fi72, 635}
                                                        var v84 int8 = 1
                                                        switch v37 {
                                                        case "":
                                                            switch "hello" {
                                                            case "":
                                                                switch fi72 {
                                                                case 32768:
                                                                    fmt.Println(v13, v14, v1)
                                                                case -2147483648:
                                                                    v85 := GenS6{F0: 1, F1: 42}
                                                                    _ = v85
                                                                default:
                                                                    v66 -= v2
                                                                }
                                                            }
                                                        case "hello":
                                                            var v86 uint32 = 214
                                                            _ = v86
                                                        case "foo bar":
                                                            v87 := GenS8{F0: false}
                                                            _ = v87
                                                        }
                                                        _ = fi82
                                                        _ = v83
                                                        _ = v84
                                                    }
                                                    var v88 uint8 = 169
                                                    _ = v88
                                                }
                                                var v89 uint32 = 930
                                                _ = v89
                                            }
                                        case 127:
                                            for fi90 := 0; fi90 < 20; fi90++ {
                                                v91 := GenS8{F0: true}
                                                _ = fi90
                                                _ = v91
                                            }
                                        default:
                                            v92 := GenS9{F0: v13, F1: ((v37 + v37) + (("" + "hello") + "\n\t")), F2: -9223372036854775808}
                                            _ = v92
                                        }
                                        v93 := []int{v13, fi72}
                                        _ = fi72
                                        _ = v93
                                    }
                                }
                            }
                        }
                        var v94 uint32 = 516
                        _ = v66
                        _ = v94
                    } else {
                        var v95 int32 = 42
                        _ = v95
                    }
                    var v96 float64 = v2
                    switch v37 {
                    case "":
                        if (false && true) {
                            v14 = 0
                        } else {
                            var v97 uint = 65535
                            _ = v97
                        }
                    case "a":
                        if true {
                            v13 = 1
                            v13 += (v14 - v13)
                            v98 := []int{v4, v14, -2147483648}
                            _ = v98
                        }
                    case "hello":
                        v1 = genFn0(1, v37, false)
                    default:
                        var v99 uint8 = 3
                        _ = v99
                    }
                    _ = v96
                } else {
                    v14 = -1
                    for fi100 := 0; fi100 < 15; fi100++ {
                        v13 = (v13 - (fi100 - 9223372036854775807))
                        _ = fi100
                    }
                }
            }
            v101 := []int{(v13 << ((504) & 63)), 854, 0, v4}
            v102 := GenS5{F0: 0, F1: -529, F2: 0}
            _ = v101
            _ = v102
        }
    }
    if true {
        v103 := v3[v14]
        if (32767 != v13) {
            v104 := GenS7{F0: 42, F1: 65535}
            fmt.Println(v14, v14)
            v1 -= v1
            _ = v104
        }
        v105 := GenS10{F0: 9223372036854775807, F1: -633}
        _ = v103
        _ = v105
    } else {
        if false {
            switch v4 {
            case -1:
                if true {
                    v14 = (v14 / ((v14) | 1))
                }
            case 32768:
                if false {
                    v3[v13] = (453 < 255)
                    v106 := []int{}
                    _ = v106
                } else {
                    var v107 float32 = float32(v2)
                    _ = v107
                }
            case -90:
                if ((v14 & v13) < (v14 - (v14 >> ((255) & 63)))) {
                    var v108 string = ("" + v37)
                    _ = v108
                }
            }
            v2 = 0.1
        }
        var v109 bool = (377 < (v4 ^ -1))
        _ = v109
    }
    switch v14 {
    case 127:
        if (42 < 32767) {
            var v110 float32 = (1e10 - (0.0 * (float32((1e10 - (-2.25 + (v2 / v2)))) + (0.0 * ((0.0 / ((0.1 + 3.14159) - 0.1)) + (0.0 * (3.14159 - (1.5 * -2.25))))))))
            fmt.Println(v37, v2, v13)
            _ = v110
        }
    case 925:
        var v111 string = ("foo bar" + "hello")
        _ = v111
    default:
        switch (v37 + "a") {
        case "":
            v3[v14] = true
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v37
    _ = v4
}
