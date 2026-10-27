package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint16; F1 string }
type GenS5 struct { F0 int; F1 bool; F2 int }
type GenS6 struct { F0 int8; F1 int8 }
type GenS7 struct { F0 uint; F1 int32; F2 uint }
func genFn0() uint {
    for fi5 := 0; fi5 < 10; fi5++ {
        v6 := GenS1{F0: 480, F1: "a"}
        var v7 int64 = 9223372036854775807
        for fi8 := 0; fi8 < 19; fi8++ {
            v9 := GenS1{F0: 42, F1: (("\n\t" + ("\n\t" + "foo bar")) + "foo bar")}
            v10 := GenS1{F0: 32767, F1: "a"}
            var v11 int16 = -992
            if (4294967295 == 0) {
                v11 = (-(v11))
                v11 -= v11
                if false {
                    v12 := []int{(fi5 & (fi8 >> ((1) & 63))), (fi5 << ((1) & 63))}
                    _ = v12
                }
                fmt.Println(v11)
            }
            v13 := []int{(fi5 & 4294967295)}
            _ = fi8
            _ = v10
            _ = v11
            _ = v13
            _ = v9
        }
        _ = fi5
        _ = v6
        _ = v7
    }
    var v14 uint8 = 128
    if false {
        if (0 != -1) {
            var v15 int8 = 127
            for fi16 := 0; fi16 < 19; fi16++ {
                if (true && false) {
                    v17 := []int{fi16}
                    fmt.Println(v14, fi16)
                    _ = v17
                } else {
                    v18 := []int{fi16}
                    for fi19 := 0; fi19 < 5; fi19++ {
                        v15 = (v15 % ((v15) | 1))
                        v15 = (v15 % ((-60) | 1))
                        _ = fi19
                    }
                    switch ("" + "foo bar") {
                    case "":
                        for fi20 := 0; fi20 < 11; fi20++ {
                            var v21 int8 = v15
                            fmt.Println(v14, v15, fi16)
                            _ = fi20
                            _ = v21
                        }
                    }
                    _ = v18
                }
                if true {
                    switch (fi16 + (fi16 % (((-(fi16))) | 1))) {
                    case 570:
                        if (float64((-2.25 * -2.25)) == -2.25) {
                            v22 := []int{}
                            _ = v22
                        } else {
                            for fi23 := 0; fi23 < 2; fi23++ {
                                if true {
                                    if false {
                                        v24 := []int{(fi23 << ((357) & 63)), (^(fi16)), (fi23 - (fi23 / (((fi16 % ((fi16) | 1))) | 1))), 128}
                                        _ = v24
                                    }
                                }
                                var v25 string = "\n\t"
                                _ = fi23
                                _ = v25
                            }
                            if true {
                                v26 := []int{(fi16 / ((0) | 1)), -1}
                                _ = v26
                            }
                        }
                    case 32768:
                        fmt.Println(v15, v15, v14)
                    case 1:
                        var v27 uint32 = 0
                        _ = v27
                    default:
                        v28 := []int{fi16}
                        _ = v28
                    }
                } else {
                    v29 := []int{127, fi16}
                    v30 := GenS1{F0: 0, F1: "foo bar"}
                    v14 = v14
                    _ = v29
                    _ = v30
                }
                _ = fi16
            }
            fmt.Println(v15)
            _ = v15
        } else {
            var v31 uint8 = 245
            v32 := []int{}
            if true {
                v33 := GenS1{F0: 328, F1: "\n\t"}
                fmt.Println(v14)
                var v34 int64 = -272
                _ = v33
                _ = v34
            } else {
                for fi35 := 0; fi35 < 14; fi35++ {
                    var v36 float64 = 1e10
                    var v37 uint64 = 0
                    _ = fi35
                    _ = v36
                    _ = v37
                }
                fmt.Println(v14)
            }
            _ = v31
            _ = v32
        }
        if ((true || (!(false && false))) || false) {
            var v38 int16 = -341
            var v39 uint16 = 840
            var v40 int = 393
            _ = v38
            _ = v39
            _ = v40
        } else {
            v14 = (v14 ^ (v14 >> ((170) & 63)))
            if ("\n\t" <= "hello") {
                if (false || (false || (2 < 585))) {
                    if false {
                        v41 := GenS1{F0: 700, F1: ("foo bar" + ("\n\t" + "foo bar"))}
                        fmt.Println(v14)
                        fmt.Println(v14)
                        _ = v41
                    }
                } else {
                    v14 = v14
                    v14 -= v14
                }
                v14 += (v14 % ((v14) | 1))
            } else {
                for fi42 := 0; fi42 < 11; fi42++ {
                    var v43 uint32 = 65535
                    _ = fi42
                    _ = v43
                }
                for fi44 := 0; fi44 < 17; fi44++ {
                    var v45 uint64 = 778
                    var v46 int64 = -1
                    _ = fi44
                    _ = v45
                    _ = v46
                }
                var v47 float64 = 1e10
                _ = v47
            }
            switch "hello" {
            case "hello":
                if ("" > "a") {
                    var v48 uint16 = 32767
                    v48 = 1
                    _ = v48
                }
            }
        }
    }
    _ = v14
    return 974
}

func main() {
    v1 := []int{-9223372036854775808}
    var v2 uint16 = 42
    v3 := []uint32{2147483648, 2147483648, 65535, 0}
    var v4 string = "foo bar"
    v49 := GenS2{F0: 1}
    v4 = (v4 + v4)
    v2 = (v2 ^ (v2 * 194))
    switch 32767 {
    case -648:
        for fi50 := 0; fi50 < 12; fi50++ {
            if true {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                v4 = v4
                switch (v4 + (((v4 + v4) + ((v4 + (("a" + v4) + ("" + "\n\t"))) + v4)) + v4)) {
                case "a":
                    var v51 int64 = -9223372036854775808
                    _ = v51
                default:
                    v52 := GenS2{F0: 676}
                    _ = v52
                }
            }
            v2 += v2
            _ = fi50
        }
    case -207:
        switch 106 {
        case 32768:
            var v53 int16 = 1
            _ = v53
        case 691:
            var v54 uint16 = 2
            _ = v54
        case -2147483648:
            if (176 <= 127) {
                var v55 bool = false
                var v56 uint = genFn0()
                _ = v55
                _ = v56
            } else {
                if (9223372036854775807 <= -1) {
                    if len(v3) > 0 {
                        v3[0] = 929
                    }
                    switch ((("foo bar" + ("" + "a")) + v4) + "a") {
                    case "":
                        v57 := GenS3{F0: 876}
                        _ = v57
                    default:
                        var v58 float64 = (0.1 - -2.25)
                        _ = v58
                    }
                    v59 := GenS1{F0: v2, F1: v4}
                    _ = v59
                }
            }
        default:
            for fi60 := 0; fi60 < 5; fi60++ {
                for fi61 := 0; fi61 < 9; fi61++ {
                    if (42 > 65535) {
                        fmt.Println(fi61)
                        var v62 uint16 = v2
                        if false {
                            v63 := GenS4{F0: v2, F1: "a"}
                            _ = v63
                        } else {
                            var v64 string = (v4 + "hello")
                            var v65 uint32 = 1
                            _ = v64
                            _ = v65
                        }
                        _ = v62
                    } else {
                        v66 := GenS5{F0: (fi61 << ((2) & 63)), F1: true, F2: fi61}
                        v2 -= 245
                        switch fi61 {
                        case -1:
                            var v67 uint8 = 255
                            _ = v67
                        case 9223372036854775807:
                            switch (fi61 & (fi61 - (-(fi60)))) {
                            case 32768:
                                v2 = v2
                            case 4294967295:
                                v2 = (v2 + v2)
                            default:
                                v2 = (v2 & v2)
                            }
                        case 113:
                            v2 = 1
                        default:
                            var v68 uint16 = 32767
                            _ = v68
                        }
                        _ = v66
                    }
                    switch (-(fi61)) {
                    case 939:
                        v4 = v4
                    case 104:
                        v4 = "hello"
                    default:
                        fmt.Println(v2, v2, fi61)
                    }
                    _ = fi61
                }
                _ = fi60
            }
        }
    case 2147483647:
        v69 := GenS6{F0: 100, F1: 42}
        _ = v69
    }
    switch 4294967295 {
    case 1:
        var v70 uint8 = 127
        _ = v70
    default:
        fmt.Println(v2, v2)
    }
    if (-422 != -969) {
        v4 = v4
        if (675 <= 255) {
            fmt.Println(v2)
            v1 = append(v1, 0)
            for fi71 := 0; fi71 < 19; fi71++ {
                v1 = append(v1, fi71)
                for fi72 := 0; fi72 < 1; fi72++ {
                    if true {
                        var v73 int16 = -410
                        var v74 uint16 = (v2 / ((v2) | 1))
                        if true {
                            v75 := GenS4{F0: v74, F1: v4}
                            var v76 uint64 = 42
                            _ = v75
                            _ = v76
                        } else {
                            v77 := GenS7{F0: 18446744073709551615, F1: -1, F2: 1}
                            _ = v77
                        }
                        _ = v73
                        _ = v74
                    } else {
                        v78 := GenS6{F0: -128, F1: 100}
                        _ = v78
                    }
                    v4 = v4
                    _ = fi72
                }
                _ = fi71
            }
        } else {
            v79 := GenS1{F0: v2, F1: "\n\t"}
            _ = v79
        }
    }
    var v80 float32 = 0.1
    for fi81 := 0; fi81 < 6; fi81++ {
        var v82 uint32 = 0
        if false {
            v83 := GenS7{F0: 4294967295, F1: -2147483648, F2: 18446744073709551615}
            _ = v83
        }
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = fi81
        _ = v82
    }
    v4 = v4
    for fi84 := 0; fi84 < 8; fi84++ {
        var v85 uint8 = 1
        var v86 uint64 = 0
        _ = fi84
        _ = v85
        _ = v86
    }
    var v87 int8 = 0
    v4 = v4
    var v88 uint8 = 128
    v2 = (v2 + 32767)
    v87 -= (v87 | (v87 + (v87 & (^(v87)))))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v80:", v80)
    fmt.Println("v87:", v87)
    fmt.Println("v88:", v88)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v49
    _ = v80
    _ = v87
    _ = v88
}
