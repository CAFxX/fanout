package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint64; F1 string; F2 string }
type GenS3 struct { F0 int16; F1 uint8; F2 string }
type GenS4 struct { F0 int32; F1 int32 }
type GenS5 struct { F0 int32; F1 string }
type GenS6 struct { F0 int32; F1 uint; F2 uint }
type GenS7 struct { F0 int16; F1 int16 }
type GenS8 struct { F0 int64; F1 uint32; F2 string }
type GenS9 struct { F0 string; F1 uint64 }
func genFn0() uint {
    var v5 int8 = -1
    for fi6 := 0; fi6 < 13; fi6++ {
        v7 := []int{(fi6 * (fi6 & 128)), (fi6 >> ((1) & 63)), fi6, (fi6 - (-(fi6)))}
        _ = fi6
        _ = v7
    }
    for fi8 := 0; fi8 < 10; fi8++ {
        var v9 int8 = v5
        switch ("foo bar" + "") {
        case "a":
            v10 := []int{(fi8 ^ fi8), (fi8 << ((240) & 63))}
            _ = v10
        case "\n\t":
            v5 = v5
        }
        _ = fi8
        _ = v9
    }
    _ = v5
    return 255
}
func genFn1() uint64 {
    if false {
        var v11 int = 4294967295
        for fi12 := 0; fi12 < 10; fi12++ {
            for fi13 := 0; fi13 < 3; fi13++ {
                if (true && false) {
                    var v14 uint8 = 1
                    switch "a" {
                    case "hello":
                        var v15 uint16 = 32767
                        if (42 <= 100) {
                            var v16 float64 = -2.25
                            v15 -= v15
                            if (4294967295 < 42) {
                                fmt.Println(v14)
                            } else {
                                var v17 int64 = 0
                                v15 = 62
                                if (!(1 <= genFn0())) {
                                    v18 := GenS1{F0: 1}
                                    v16 = float64(-2.25)
                                    _ = v18
                                } else {
                                    var v19 int8 = 42
                                    var v20 float64 = (v16 - v16)
                                    var v21 int = 127
                                    switch (fi13 * (-(v21))) {
                                    case 99:
                                        var v22 int8 = (v19 << ((255) & 63))
                                        _ = v22
                                    case -1:
                                        switch v11 {
                                        case -9223372036854775808:
                                            switch (v21 >> ((18446744073709551615) & 63)) {
                                            case 128:
                                                v23 := []int{2147483647, v21, (v21 + v11)}
                                                var v24 int64 = 294
                                                if ((v19 >> ((255) & 63)) <= (v19 / ((int8(v24)) | 1))) {
                                                    v21 += (fi13 & (v21 % ((906) | 1)))
                                                    fmt.Println(fi12, fi12, v19)
                                                }
                                                _ = v23
                                                _ = v24
                                            default:
                                                v25 := GenS1{F0: 2}
                                                _ = v25
                                            }
                                        case 32768:
                                            var v26 float32 = 0.0
                                            _ = v26
                                        }
                                    case 32767:
                                        v27 := GenS2{F0: uint64(v15), F1: ("foo bar" + "a"), F2: "\n\t"}
                                        _ = v27
                                    }
                                    var v28 uint8 = (^(v14))
                                    _ = v19
                                    _ = v20
                                    _ = v21
                                    _ = v28
                                }
                                v29 := GenS1{F0: 664}
                                switch (fi12 + fi13) {
                                case 9223372036854775807:
                                    if true {
                                        v17 = (-(v17))
                                        if true {
                                            v30 := []int{(-(v11))}
                                            v17 += v17
                                            _ = v30
                                        } else {
                                            fmt.Println(v15, v15)
                                            var v31 float32 = 1e10
                                            v11 = 983
                                            _ = v31
                                        }
                                        var v32 uint16 = (v15 / (((v15 / (((v15 * (^(v15)))) | 1))) | 1))
                                        _ = v32
                                    }
                                case 32767:
                                    v33 := []int{(fi13 / (((v11 % (((v11 + 54)) | 1))) | 1)), (fi13 | fi12), (v11 & 32767)}
                                    _ = v33
                                case 1:
                                    v34 := GenS1{F0: 0}
                                    fmt.Println(fi12, fi13)
                                    _ = v34
                                default:
                                    fmt.Println(v14)
                                    fmt.Println(fi13, v15, v14)
                                    var v35 int16 = 1000
                                    _ = v35
                                }
                                _ = v17
                                _ = v29
                            }
                            v36 := []int{0, (v11 + fi13), v11}
                            fmt.Println(v15, v14, v14)
                            var v37 int = fi13
                            _ = v16
                            _ = v36
                            _ = v37
                        }
                        var v38 int32 = int32(v15)
                        var v39 uint16 = 339
                        v40 := []int{v11, 42, 32768}
                        var v41 int32 = -2147483648
                        _ = v15
                        _ = v38
                        _ = v39
                        _ = v40
                        _ = v41
                    }
                    _ = v14
                } else {
                    fmt.Println(fi12, fi13)
                    if ((723 != 42) || (false || ((0 <= 109) || (true || (false || true))))) {
                        var v42 int32 = 42
                        v43 := GenS3{F0: 42, F1: 128, F2: ("a" + "")}
                        _ = v42
                        _ = v43
                    }
                    if (491 < 2) {
                        var v44 float32 = 0.0
                        var v45 uint = genFn0()
                        switch 128 {
                        case 42:
                            if (false || true) {
                                v46 := []int{-2147483648}
                                if true {
                                    v47 := GenS1{F0: 2}
                                    _ = v47
                                }
                                _ = v46
                            }
                        case 657:
                            var v48 float64 = ((((1e10 / float64(v44)) + -2.25) - 0.1) + 0.1)
                            _ = v48
                        case 128:
                            v49 := GenS4{F0: 2147483647, F1: 69}
                            _ = v49
                        default:
                            var v50 bool = (42 < -256)
                            _ = v50
                        }
                        _ = v44
                        _ = v45
                    }
                }
                v51 := GenS5{F0: 2147483647, F1: (("" + ("hello" + "\n\t")) + ("\n\t" + "\n\t"))}
                _ = fi13
                _ = v51
            }
            for fi52 := 0; fi52 < 6; fi52++ {
                v11 -= (v11 << ((18446744073709551615) & 63))
                switch "a" {
                case "hello":
                    v11 = (fi52 >> ((497) & 63))
                    v53 := GenS6{F0: -1, F1: 787, F2: 18446744073709551615}
                    v54 := GenS2{F0: 1, F1: (("a" + "") + ("hello" + "foo bar")), F2: "\n\t"}
                    _ = v53
                    _ = v54
                case "foo bar":
                    v55 := []int{32768}
                    _ = v55
                case "a":
                    v11 += (^(fi12))
                default:
                    v56 := []int{128}
                    _ = len(v56)
                    _ = v56
                }
                _ = fi52
            }
            _ = fi12
        }
        _ = v11
    } else {
        if false {
            var v57 uint32 = 915
            v58 := []int{-2147483648, 2147483647, -1}
            switch ("\n\t" + ("\n\t" + "a")) {
            case "foo bar":
                if (42 > 0) {
                    var v59 float64 = ((float64(0.1) * ((1.5 - -2.25) / (1e10 + 0.1))) + (1.5 - 3.14159))
                    _ = len(v58)
                    _ = v59
                } else {
                    var v60 string = "a"
                    _ = v60
                }
            }
            _ = v57
            _ = v58
        }
        var v61 int16 = 32767
        _ = v61
    }
    var v62 uint16 = 65535
    for fi63 := 0; fi63 < 8; fi63++ {
        switch 0 {
        case -1:
            for fi64 := 0; fi64 < 6; fi64++ {
                switch "hello" {
                case "\n\t":
                    v65 := []int{982, fi63}
                    _ = v65
                }
                var v66 uint32 = 1
                _ = fi64
                _ = v66
            }
        case 648:
            v62 -= (^(v62))
        }
        _ = fi63
    }
    v67 := []int{42}
    v68 := GenS7{F0: 32767, F1: -1}
    _ = v62
    _ = v67
    _ = v68
    return 186
}

func main() {
    var v1 uint8 = 128
    var v2 int8 = 0
    v3 := make(map[uint32]bool)
    v4 := []string{"a", "hello"}
    v69 := GenS6{F0: 42, F1: 1, F2: genFn0()}
    fmt.Println(v1, v2)
    fmt.Println(v1, v2)
    var v70 uint = 965
    v2 -= 27
    for fi71 := 0; fi71 < 7; fi71++ {
        if false {
            var v72 uint64 = 1
            _ = v72
        } else {
            fmt.Println(fi71, fi71, v2)
            for fi73 := 0; fi73 < 19; fi73++ {
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                fmt.Println(v1)
                v3[42] = false
                _ = fi73
            }
            v74 := GenS5{F0: -1, F1: "foo bar"}
            _ = v74
        }
        for fi75 := 0; fi75 < 5; fi75++ {
            if (1 != 9223372036854775807) {
                v70 -= (v70 - v70)
                switch "a" {
                case "\n\t":
                    if false {
                        var v76 uint16 = 255
                        _ = v76
                    }
                case "hello":
                    if ((0 > 65535) || (0 < 494)) {
                        var v77 bool = (-40 <= 1000)
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        _ = v77
                    }
                case "":
                    switch "\n\t" {
                    case "a":
                        if true {
                            if (!(true || (!(0 >= 1000)))) {
                                switch "a" {
                                case "":
                                    v78 := GenS8{F0: -1, F1: 2, F2: ("\n\t" + "\n\t")}
                                    _ = v78
                                }
                                v1 = (v1 ^ v1)
                            } else {
                                var v79 uint8 = v1
                                fmt.Println(v2)
                                _ = v79
                            }
                            if false {
                                var v80 uint = v70
                                _ = v80
                            } else {
                                v70 -= (v70 + (v70 | (v70 * (v70 % (((v70 & (v70 / ((v70) | 1)))) | 1)))))
                                var v81 float32 = 0.1
                                _ = v81
                            }
                            switch 9223372036854775807 {
                            case -2147483648:
                                delete(v3, 42)
                            case 32767:
                                if (!(!false)) {
                                    v82 := GenS9{F0: "a", F1: 18446744073709551615}
                                    v83 := GenS6{F0: -1, F1: v70, F2: 65535}
                                    _ = v82
                                    _ = v83
                                } else {
                                    var v84 uint = v70
                                    v1 = v1
                                    fmt.Println(v1, fi75)
                                    _ = v84
                                }
                            case -556:
                                fmt.Println(fi71, v1)
                            default:
                                v3[2147483648] = false
                            }
                            var v85 uint32 = 1
                            switch 2147483647 {
                            case 32768:
                                switch (^(fi71)) {
                                case 127:
                                    var v86 int8 = (v2 & (v2 + (-(v2))))
                                    v4 = append(v4, "hello")
                                    _ = v86
                                case 182:
                                    v2 -= (v2 << ((v70) & 63))
                                case -1:
                                    v2 += v2
                                }
                            }
                            _ = v85
                        }
                    case "\n\t":
                        v1 -= v1
                    }
                default:
                    switch -938 {
                    case 42:
                        var v87 int32 = 0
                        _ = v87
                    case 4294967295:
                        fmt.Println(fi75, v2, v70)
                        fmt.Println(v70, v1)
                    default:
                        if len(v4) > 0 {
                            v4[0] = ("a" + ("foo bar" + ""))
                        }
                    }
                    _ = len(v4)
                }
            }
            fmt.Println(v1, fi75, v70)
            v70 -= v70
            _ = fi75
        }
        _ = fi71
    }
    var v88 uint32 = 238
    switch "hello" {
    case "\n\t":
        v70 -= (v70 | 621)
    case "foo bar":
        var v89 uint16 = 584
        _ = v89
    }
    v90 := GenS5{F0: -2147483648, F1: "hello"}
    var v91 float64 = 3.14159
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v70:", v70)
    fmt.Println("v88:", v88)
    fmt.Println("v91:", v91)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v69
    _ = v70
    _ = v88
    _ = v90
    _ = v91
}
