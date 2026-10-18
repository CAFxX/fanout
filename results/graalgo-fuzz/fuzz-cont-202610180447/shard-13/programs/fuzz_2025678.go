package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32 }
type GenS2 struct { F0 int; F1 uint8; F2 bool }
type GenS3 struct { F0 int; F1 string; F2 uint8 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 uint16; F1 int8 }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 bool }
type GenS8 struct { F0 uint8; F1 int16; F2 uint8 }
func genFn0() string {
    v5 := GenS1{F0: -1, F1: -663}
    if ("" != "foo bar") {
        var v6 uint64 = 4294967295
        _ = v6
    } else {
        var v7 bool = false
        _ = v7
    }
    _ = v5
    return ("hello" + "hello")
}
func genFn1() uint32 {
    switch 32767 {
    case 0:
        var v8 uint16 = 255
        _ = v8
    }
    for fi9 := 0; fi9 < 18; fi9++ {
        var v10 string = ""
        fmt.Println(v10)
        _ = fi9
        _ = v10
    }
    return 4294967295
}

func main() {
    var v1 uint16 = 0
    var v2 float32 = 0.1
    var v3 int64 = -997
    v4 := make(map[uint32]uint32)
    fmt.Println(v1, v1, v2)
    v2 = (v2 * float32((-2.25 + 0.1)))
    delete(v4, 243)
    if false {
        v3 -= v3
        if true {
            _ = len(v4)
            fmt.Println(v3, v2)
            fmt.Println(v2, v1)
        } else {
            var v11 uint8 = 24
            if false {
                var v12 uint32 = 65535
                _ = v12
            } else {
                for fi13 := 0; fi13 < 14; fi13++ {
                    fmt.Println(v11, v2, v3)
                    fmt.Println(fi13)
                    if (2 == 779) {
                        var v14 uint16 = (v1 + (^(v1)))
                        switch "a" {
                        case "\n\t":
                            v15 := GenS1{F0: -9223372036854775808, F1: -332}
                            _ = v15
                        case "foo bar":
                            v14 = 559
                        }
                        _ = v14
                    }
                    _ = fi13
                }
                if false {
                    for fi16 := 0; fi16 < 13; fi16++ {
                        if true {
                            v3 = v3
                            v11 = 0
                        }
                        v17 := []int{(fi16 % ((fi16) | 1)), (fi16 >> ((65535) & 63))}
                        _ = fi16
                        _ = v17
                    }
                    for fi18 := 0; fi18 < 4; fi18++ {
                        switch "hello" {
                        case "a":
                            if false {
                                if (100 != 127) {
                                    var v19 int = 0
                                    var v20 float64 = 1.5
                                    _ = v19
                                    _ = v20
                                } else {
                                    v21 := GenS2{F0: (fi18 + (fi18 / (((fi18 - (-(fi18)))) | 1))), F1: 0, F2: (true || (42 <= 1000))}
                                    if (true || true) {
                                        var v22 uint16 = (^(v1))
                                        _ = v22
                                    }
                                    _ = v21
                                }
                            } else {
                                var v23 bool = true
                                v23 = false
                                _ = v23
                            }
                        case "":
                            var v24 float32 = v2
                            _ = v24
                        }
                        if false {
                            v25 := GenS2{F0: (fi18 | (fi18 * int(v3))), F1: 1, F2: false}
                            var v26 uint = 440
                            _ = v25
                            _ = v26
                        } else {
                            fmt.Println(v11, v2)
                            v2 += 1.5
                            if (!(false && ((fi18 | fi18) < fi18))) {
                                v2 += (1.5 * v2)
                                fmt.Println(v1)
                            }
                        }
                        _ = fi18
                    }
                    var v27 float32 = 0.1
                    _ = v27
                }
                var v28 float64 = 0.1
                _ = v28
            }
            if false {
                v1 += (^(v1))
            }
            _ = v11
        }
        for fi29 := 0; fi29 < 15; fi29++ {
            var v30 uint32 = 2147483647
            for fi31 := 0; fi31 < 7; fi31++ {
                fmt.Println(v2)
                switch "hello" {
                case "\n\t":
                    v1 -= v1
                case "":
                    var v32 uint32 = (v30 + (v30 ^ (v30 >> ((uint(v1)) & 63))))
                    _ = v32
                default:
                    if (!(1 > 390)) {
                        v33 := GenS3{F0: fi31, F1: ("a" + ""), F2: 2}
                        v34 := GenS4{F0: 255}
                        var v35 int8 = 42
                        _ = v33
                        _ = v34
                        _ = v35
                    }
                }
                var v36 string = "foo bar"
                var v37 string = v36
                v2 += v2
                _ = fi31
                _ = v36
                _ = v37
            }
            _ = fi29
            _ = v30
        }
    } else {
        v3 = (v3 << ((255) & 63))
        for fi38 := 0; fi38 < 16; fi38++ {
            v2 -= ((v2 + (v2 / float32(3.14159))) + float32(1e10))
            v39 := GenS5{F0: v1, F1: -128}
            v3 = (^(v3))
            _ = fi38
            _ = v39
        }
    }
    switch "hello" {
    case "foo bar":
        if false {
            switch "hello" {
            case "hello":
                for fi40 := 0; fi40 < 3; fi40++ {
                    for fi41 := 0; fi41 < 14; fi41++ {
                        v2 = (v2 / v2)
                        var v42 uint32 = 107
                        _ = fi41
                        _ = v42
                    }
                    _ = fi40
                }
            default:
                var v43 int8 = 1
                _ = v43
            }
            if true {
                v1 -= (^(v1))
                var v44 uint32 = genFn1()
                _ = v44
            }
        }
    }
    if (false && (!(4294967295 <= 2147483648))) {
        for fi45 := 0; fi45 < 7; fi45++ {
            switch (^(fi45)) {
            case 128:
                switch 2147483647 {
                case -663:
                    switch fi45 {
                    case 9223372036854775807:
                        var v46 int64 = 1
                        _ = v46
                    case 42:
                        for fi47 := 0; fi47 < 2; fi47++ {
                            v3 = (v3 + (v3 - (v3 + v3)))
                            v2 += 0.1
                            v3 = v3
                            _ = fi47
                        }
                    case 128:
                        switch 9223372036854775807 {
                        case 302:
                            v1 += v1
                        case -1:
                            for fi48 := 0; fi48 < 11; fi48++ {
                                var v49 uint32 = genFn1()
                                switch 127 {
                                case 4294967295:
                                    v1 = (v1 * 255)
                                case 307:
                                    v50 := []int{fi48}
                                    _ = v50
                                case 991:
                                    var v51 int16 = 42
                                    _ = v51
                                default:
                                    switch ((((("foo bar" + "foo bar") + ("hello" + ((("hello" + ((("hello" + "a") + (("foo bar" + ("\n\t" + "a")) + "hello")) + ("\n\t" + ("\n\t" + "")))) + "hello") + ("foo bar" + "\n\t")))) + ("\n\t" + "")) + ("\n\t" + "foo bar")) + ("foo bar" + "hello")) {
                                    case "foo bar":
                                        v3 += v3
                                    default:
                                        v2 += v2
                                    }
                                }
                                v49 -= v49
                                v52 := GenS5{F0: v1, F1: 1}
                                switch "\n\t" {
                                case "\n\t":
                                    var v53 string = "\n\t"
                                    _ = v53
                                case "":
                                    _ = len(v4)
                                case "a":
                                    if ((2 != 183) && (1e10 < float64(float32(0.0)))) {
                                        v54 := GenS4{F0: 2}
                                        _ = v54
                                    } else {
                                        if true {
                                            if ((18446744073709551615 < 2) && false) {
                                                delete(v4, genFn1())
                                                switch "" {
                                                case "":
                                                    v55 := GenS6{F0: (v3 > v3)}
                                                    _ = v55
                                                case "\n\t":
                                                    v56 := v4[v49]
                                                    _ = v56
                                                case "foo bar":
                                                    var v57 uint32 = genFn1()
                                                    _ = v57
                                                }
                                            }
                                        } else {
                                            if true {
                                                fmt.Println(v3)
                                                v58 := []int{-651}
                                                var v59 uint64 = 1000
                                                if (-1 != -2147483648) {
                                                    v49 += (v49 % ((genFn1()) | 1))
                                                } else {
                                                    switch "\n\t" {
                                                    case "\n\t":
                                                        var v60 int16 = 362
                                                        _ = v60
                                                    }
                                                    v1 -= (v1 / ((v1) | 1))
                                                    switch (fi45 + (fi45 & (fi48 >> ((65535) & 63)))) {
                                                    case 127:
                                                        v61 := GenS7{F0: false}
                                                        _ = v61
                                                    case -931:
                                                        var v62 bool = true
                                                        _ = v62
                                                    default:
                                                        fmt.Println(v49)
                                                        v63 := GenS4{F0: 255}
                                                        v64 := GenS4{F0: 310}
                                                        v65 := GenS3{F0: 1, F1: ("a" + "foo bar"), F2: 252}
                                                        fmt.Println(v59)
                                                        v1 = 0
                                                        _ = v63
                                                        _ = v64
                                                        _ = v65
                                                    }
                                                }
                                                v49 = v49
                                                v49 -= v49
                                                delete(v4, genFn1())
                                                var v66 int64 = v3
                                                fmt.Println(fi45, fi48, fi48)
                                                _ = v58
                                                _ = v59
                                                _ = v66
                                            }
                                            switch "" {
                                            case "foo bar":
                                                if (2 <= 192) {
                                                    v2 -= ((0.1 / ((v2 * 3.14159) + v2)) * float32(1.5))
                                                    var v67 int32 = 1
                                                    _ = v67
                                                } else {
                                                    v1 += (v1 | (v1 * (v1 + v1)))
                                                    var v68 string = "foo bar"
                                                    v69 := []int{fi45, fi48}
                                                    _ = v68
                                                    _ = v69
                                                }
                                            case "hello":
                                                if false {
                                                    v2 += v2
                                                    var v70 uint16 = (v1 | (v1 << ((1) & 63)))
                                                    v71 := GenS3{F0: 128, F1: "foo bar", F2: 128}
                                                    _ = v70
                                                    _ = v71
                                                } else {
                                                    v1 = (v1 * v1)
                                                    var v72 int8 = 117
                                                    var v73 uint16 = (v1 / ((v1) | 1))
                                                    v74 := []int{42}
                                                    _ = v72
                                                    _ = v73
                                                    _ = v74
                                                }
                                            case "":
                                                delete(v4, (v49 + (v49 >> ((984) & 63))))
                                            }
                                            var v75 int8 = 127
                                            fmt.Println(fi48, v49)
                                            _ = v75
                                        }
                                    }
                                }
                                _ = fi48
                                _ = v49
                                _ = v52
                            }
                        default:
                            v1 -= 212
                        }
                    default:
                        if ((!true) || (((v3 ^ 0) <= 9223372036854775807) && true)) {
                            if (int8(fi45) > -1) {
                                if true {
                                    if (42 <= 834) {
                                        v1 += (^(v1))
                                        v76 := []int{(-(fi45)), (fi45 + 569), fi45, fi45}
                                        _ = v76
                                    } else {
                                        fmt.Println(v3, v1)
                                        for fi77 := 0; fi77 < 12; fi77++ {
                                            v2 += v2
                                            _ = fi77
                                        }
                                        var v78 int32 = 2147483647
                                        _ = v78
                                    }
                                    delete(v4, 253)
                                }
                            }
                            v79 := []int{}
                            _ = v79
                        }
                    }
                case 42:
                    if true {
                        fmt.Println(v1)
                    } else {
                        for fi80 := 0; fi80 < 5; fi80++ {
                            var v81 uint16 = (v1 % ((345) | 1))
                            switch "a" {
                            case "\n\t":
                                v2 = float32(1.5)
                            case "hello":
                                v82 := GenS1{F0: (fi45 * fi45), F1: -1}
                                _ = v82
                            default:
                                v1 += (v81 * v1)
                            }
                            _ = fi80
                            _ = v81
                        }
                        switch ("" + "a") {
                        case "\n\t":
                            if ((42 >= 128) && (1 <= -128)) {
                                fmt.Println(v2, v1)
                                switch "a" {
                                case "":
                                    fmt.Println(v3, fi45)
                                case "\n\t":
                                    v2 = v2
                                case "foo bar":
                                    v3 += v3
                                }
                                if false {
                                    v1 -= 1
                                    var v83 int64 = (v3 % ((-1) | 1))
                                    _ = v83
                                }
                            }
                        case "hello":
                            if (-128 > 48) {
                                v1 += v1
                                var v84 int32 = -2147483648
                                fmt.Println(v84, v1)
                                _ = v84
                            }
                        default:
                            if (!false) {
                                for fi85 := 0; fi85 < 15; fi85++ {
                                    v86 := GenS6{F0: (v1 < v1)}
                                    _ = fi85
                                    _ = v86
                                }
                                v3 = 1
                                fmt.Println(v1, v2, v2)
                            }
                        }
                    }
                case -2147483648:
                    var v87 uint8 = 253
                    _ = v87
                }
            default:
                v1 -= v1
            }
            if false {
                v2 += 1e10
                for fi88 := 0; fi88 < 8; fi88++ {
                    switch fi88 {
                    case -813:
                        if (0 != 1) {
                            fmt.Println(fi88, v1, fi88)
                        }
                    case 1:
                        v89 := GenS8{F0: 154, F1: 32767, F2: 1}
                        _ = v89
                    default:
                        v90 := GenS1{F0: 127, F1: -1}
                        _ = v90
                    }
                    v91 := []int{(-(fi45)), (fi45 << ((328) & 63)), 9223372036854775807, (fi88 - fi88)}
                    _ = fi88
                    _ = v91
                }
                for fi92 := 0; fi92 < 11; fi92++ {
                    _ = len(v4)
                    fmt.Println(v3)
                    _ = fi92
                }
            }
            var v93 uint8 = 0
            _ = fi45
            _ = v93
        }
        v94 := []int{2147483647, 32767}
        var v95 uint64 = 919
        _ = v94
        _ = v95
    } else {
        v3 -= (v3 ^ v3)
    }
    for fi96 := 0; fi96 < 3; fi96++ {
        var v97 float64 = 0.1
        v98 := GenS4{F0: 728}
        _ = fi96
        _ = v97
        _ = v98
    }
    if (44 <= 127) {
        if ((!true) || false) {
            v99 := v4[2147483648]
            var v100 uint16 = v1
            if (-2.25 > 1.5) {
                var v101 float64 = 0.0
                var v102 float32 = ((float32((((1e10 - (v101 - 0.0)) - 3.14159) / float64(float32((v101 + (v101 / v101)))))) - (v2 * (v2 - ((v2 / v2) - (v2 * v2))))) - ((0.0 * 3.14159) * float32((float64((float32(0.1) - (v2 / 1.0))) - (float64(0.1) / v101)))))
                _ = v101
                _ = v102
            } else {
                for fi103 := 0; fi103 < 4; fi103++ {
                    var v104 int32 = 2147483647
                    _ = fi103
                    _ = v104
                }
            }
            _ = v100
            _ = v99
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
}
