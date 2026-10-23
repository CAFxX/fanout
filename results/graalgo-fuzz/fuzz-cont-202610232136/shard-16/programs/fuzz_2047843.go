package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint16; F1 int }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 string; F1 uint }
type GenS6 struct { F0 uint8; F1 uint; F2 int8 }
type GenS7 struct { F0 int64; F1 int8; F2 int8 }
type GenS8 struct { F0 bool; F1 int32; F2 string }
type GenS9 struct { F0 int16; F1 int32; F2 int }
type GenS10 struct { F0 uint8 }
func genFn0() uint32 {
    for fi7 := 0; fi7 < 16; fi7++ {
        for fi8 := 0; fi8 < 9; fi8++ {
            v9 := []int{(fi8 ^ fi8), 310, -2147483648}
            var v10 uint16 = 32767
            if len(v9) >= 2 {
                v9 = v9[1:len(v9)]
            }
            var v11 uint64 = 42
            if len(v9) > 0 {
                v9[0] = 128
            }
            v12 := GenS1{F0: 74}
            _ = fi8
            _ = v10
            _ = v11
            _ = v12
            _ = v9
        }
        for fi13 := 0; fi13 < 10; fi13++ {
            if (true && ("" != (("a" + "a") + "foo bar"))) {
                v14 := []int{}
                var v15 uint8 = 1
                var v16 uint16 = 581
                _ = v14
                _ = v15
                _ = v16
            } else {
                v17 := GenS1{F0: 127}
                v18 := GenS2{F0: (fi7 / ((fi7) | 1))}
                _ = v17
                _ = v18
            }
            switch ("" + "foo bar") {
            case "":
                var v19 uint32 = 1
                _ = v19
            default:
                switch (^(fi13)) {
                case -9223372036854775808:
                    if false {
                        var v20 int64 = 42
                        if (42 <= 65535) {
                            var v21 uint32 = 890
                            v21 = (v21 & (v21 / ((v21) | 1)))
                            v22 := GenS1{F0: 1}
                            v23 := GenS3{F0: 2, F1: fi13}
                            _ = v21
                            _ = v22
                            _ = v23
                        } else {
                            var v24 bool = (!true)
                            fmt.Println(v24, v20)
                            _ = v24
                        }
                        var v25 uint32 = 4294967295
                        var v26 uint8 = 1
                        _ = v20
                        _ = v25
                        _ = v26
                    } else {
                        fmt.Println(fi13, fi7)
                        fmt.Println(fi13, fi13)
                    }
                case -788:
                    switch "foo bar" {
                    case "foo bar":
                        if false {
                            v27 := GenS3{F0: 876, F1: 4294967295}
                            var v28 uint16 = 42
                            switch ("a" + "hello") {
                            case "a":
                                switch "\n\t" {
                                case "foo bar":
                                    if (true || false) {
                                        if (!true) {
                                            v29 := GenS4{F0: 65535}
                                            var v30 float32 = 1e10
                                            v30 = (((v30 * 0.0) / v30) / 3.14159)
                                            _ = v29
                                            _ = v30
                                        }
                                    } else {
                                        var v31 uint64 = 545
                                        _ = v31
                                    }
                                case "":
                                    v28 += 255
                                default:
                                    var v32 float32 = (1e10 + float32(float64(0.1)))
                                    _ = v32
                                }
                            case "foo bar":
                                var v33 string = ("hello" + "foo bar")
                                _ = v33
                            default:
                                if (true && (!(true && ((-81 > -122) || (!false))))) {
                                    var v34 bool = (false && false)
                                    if true {
                                        var v35 int8 = 100
                                        switch -2147483648 {
                                        case -957:
                                            v36 := GenS5{F0: "", F1: 18446744073709551615}
                                            _ = v36
                                        default:
                                            if (0 == 2) {
                                                var v37 uint8 = 42
                                                _ = v37
                                            }
                                        }
                                        _ = v35
                                    }
                                    v34 = (361 > 42)
                                    v34 = v34
                                    if false {
                                        var v38 bool = v34
                                        var v39 int = fi13
                                        var v40 uint8 = 42
                                        if ((3.14159 + 3.14159) != float32((float64(-2.25) * float64(3.14159)))) {
                                            fmt.Println(v40, v39)
                                        } else {
                                            switch "a" {
                                            case "foo bar":
                                                if v38 {
                                                    var v41 int64 = 0
                                                    _ = v41
                                                }
                                            case "hello":
                                                if (("hello" + "hello") != "hello") {
                                                    fmt.Println(v39, v34, fi7)
                                                    v28 = v28
                                                    fmt.Println(fi13, v40, v28)
                                                }
                                            case "":
                                                var v42 uint = 4294967295
                                                _ = v42
                                            default:
                                                v34 = v38
                                            }
                                            v34 = true
                                            v38 = (1 < 1)
                                        }
                                        var v43 uint = 18446744073709551615
                                        v44 := []int{128, (fi13 ^ (^(fi7))), (fi13 >> (((v43 % ((v43) | 1))) & 63)), v39}
                                        v34 = v38
                                        _ = v38
                                        _ = v39
                                        _ = v40
                                        _ = v43
                                        _ = v44
                                    } else {
                                        v45 := []int{}
                                        var v46 uint8 = 58
                                        v47 := GenS6{F0: v46, F1: 255, F2: 0}
                                        var v48 int64 = 1
                                        _ = v45
                                        _ = v46
                                        _ = v47
                                        _ = v48
                                    }
                                    var v49 float32 = (1e10 * 1.5)
                                    _ = v34
                                    _ = v49
                                } else {
                                    v28 = v28
                                    var v50 uint = 0
                                    _ = v50
                                }
                            }
                            v51 := GenS4{F0: 255}
                            _ = v27
                            _ = v28
                            _ = v51
                        }
                    case "\n\t":
                        fmt.Println(fi7, fi13)
                    }
                case -692:
                    var v52 int32 = 42
                    _ = v52
                default:
                    fmt.Println(fi7, fi13)
                }
                if true {
                    if false {
                        v53 := []int{(fi13 / ((fi13) | 1)), fi13, 575}
                        var v54 float64 = 0.1
                        var v55 int16 = 0
                        _ = v53
                        _ = v54
                        _ = v55
                    } else {
                        var v56 uint = 2
                        _ = v56
                    }
                }
                var v57 uint32 = 0
                if (21 >= 127) {
                    fmt.Println(fi13)
                    var v58 uint64 = 10
                    _ = v58
                }
                var v59 int8 = 0
                var v60 float64 = 0.0
                var v61 uint64 = 832
                var v62 int64 = -658
                var v63 uint64 = (v61 >> ((42) & 63))
                _ = v57
                _ = v59
                _ = v60
                _ = v61
                _ = v62
                _ = v63
            }
            var v64 float32 = (0.0 - 0.0)
            _ = fi13
            _ = v64
        }
        _ = fi7
    }
    var v65 string = "hello"
    fmt.Println(v65)
    var v66 float64 = 1.5
    _ = v65
    _ = v66
    return 1
}

func main() {
    var v1 uint16 = 290
    var v2 int64 = -9223372036854775808
    v3 := []string{"foo bar", "", "foo bar"}
    var v4 string = "a"
    v5 := make(map[uint32]uint32)
    var v6 int32 = 2147483647
    v5[1] = 590
    for fi67 := 0; fi67 < 13; fi67++ {
        var v68 int32 = v6
        _ = fi67
        _ = v68
    }
    switch 42 {
    case -303:
        v3 = append(v3, v4)
    case -70:
        fmt.Println(v6, v1, v6)
    }
    for fi69 := 0; fi69 < 3; fi69++ {
        if (!true) {
            if false {
                v6 = v6
                switch (("" + "") + "\n\t") {
                case "":
                    for fi70 := 0; fi70 < 5; fi70++ {
                        _ = len(v3)
                        switch v4 {
                        case "foo bar":
                            if (!((((-2.25 - float64(0.1)) / 1.5) - (((0.0 + -2.25) / 1.5) + (-2.25 + 3.14159))) < float64(((1e10 / 1.5) * 1.5)))) {
                                var v71 uint32 = 339
                                v1 = 42
                                if ((((1.5 + (1e10 / -2.25)) + -2.25) * -2.25) != 1.5) {
                                    v72 := GenS1{F0: 0}
                                    if false {
                                        v73 := GenS5{F0: v4, F1: 1}
                                        _ = v73
                                    }
                                    _ = v72
                                }
                                _ = v71
                            }
                        case "\n\t":
                            v74 := GenS7{F0: (^(v2)), F1: 1, F2: int8(fi69)}
                            _ = v74
                        }
                        _ = fi70
                    }
                case "\n\t":
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                case "foo bar":
                    if (fi69 != (fi69 << ((255) & 63))) {
                        for fi75 := 0; fi75 < 10; fi75++ {
                            switch "hello" {
                            case "hello":
                                v76 := GenS8{F0: false, F1: v6, F2: (v4 + v4)}
                                _ = v76
                            default:
                                var v77 int64 = (v2 | int64(fi75))
                                _ = v77
                            }
                            switch 32768 {
                            case -2147483648:
                                v6 += (-(v6))
                            case 0:
                                _ = len(v3)
                                var v78 uint8 = 1
                                _ = v78
                            case 9223372036854775807:
                                var v79 uint16 = (v1 * v1)
                                _ = v79
                            default:
                                v80 := GenS2{F0: (fi69 << ((4294967295) & 63))}
                                _ = v80
                            }
                            fmt.Println(v4, fi69)
                            _ = fi75
                        }
                        var v81 uint16 = v1
                        _ = v81
                    } else {
                        v6 += v6
                        v82 := GenS3{F0: 759, F1: -208}
                        for fi83 := 0; fi83 < 8; fi83++ {
                            v2 = (-(v2))
                            var v84 uint16 = v1
                            if (v2 < v2) {
                                v3 = append(v3, v4)
                                var v85 int8 = 100
                                v86 := GenS5{F0: "a", F1: 65535}
                                switch "" {
                                case "hello":
                                    v2 = (-(v2))
                                case "":
                                    v85 = 127
                                case "a":
                                    fmt.Println(v4)
                                default:
                                    var v87 int8 = v85
                                    _ = v87
                                }
                                v6 = (-(v6))
                                v85 = int8(v2)
                                var v88 uint64 = 1
                                _ = v85
                                _ = v86
                                _ = v88
                            } else {
                                fmt.Println(fi83, v6)
                                fmt.Println(v4, v2, fi69)
                            }
                            v89 := GenS9{F0: -480, F1: -1, F2: fi83}
                            v3 = append(v3, (v4 + v4))
                            v90 := GenS7{F0: 42, F1: 42, F2: -128}
                            _ = fi83
                            _ = v84
                            _ = v89
                            _ = v90
                        }
                        _ = v82
                    }
                }
            } else {
                v2 = (-(v2))
                var v91 int8 = -22
                var v92 bool = (!((v6 & v6) == (v6 + v6)))
                _ = v91
                _ = v92
            }
            if len(v3) > 0 {
                v3[0] = v4
            }
        }
        var v93 int = -1
        _ = fi69
        _ = v93
    }
    for fi94 := 0; fi94 < 3; fi94++ {
        v4 = "\n\t"
        for fi95 := 0; fi95 < 17; fi95++ {
            v96 := GenS3{F0: v1, F1: (fi95 << ((676) & 63))}
            _ = fi95
            _ = v96
        }
        _ = fi94
    }
    fmt.Println(v6)
    if (false || ((v1 & v1) <= (v1 | 99))) {
        for fi97 := 0; fi97 < 12; fi97++ {
            var v98 float64 = ((-2.25 * (1e10 / -2.25)) - 0.1)
            v99 := GenS10{F0: 1}
            if len(v3) > 0 {
                v3[0] = v4
            }
            _ = fi97
            _ = v98
            _ = v99
        }
        v1 = (v1 * 32767)
        for fi100 := 0; fi100 < 17; fi100++ {
            for fi101 := 0; fi101 < 16; fi101++ {
                fmt.Println(v6, fi100)
                _ = fi101
            }
            _ = fi100
        }
    } else {
        v2 += v2
    }
    v102 := GenS8{F0: false, F1: v6, F2: (v4 + "")}
    v103 := GenS8{F0: (true && (42 != 42)), F1: -135, F2: ("" + v4)}
    switch v4 {
    case "hello":
        var v104 uint64 = 154
        _ = v104
    case "":
        v1 = 2
    case "a":
        if (791 == (v6 >> ((42) & 63))) {
            switch "\n\t" {
            case "hello":
                fmt.Println(v1)
            default:
                v4 = v4
            }
            v1 -= (^(v1))
            switch 1 {
            case 32768:
                switch 128 {
                case 697:
                    var v105 int32 = v6
                    _ = v105
                case 42:
                    fmt.Println(v1, v6)
                default:
                    fmt.Println(v6, v1)
                }
            case -9223372036854775808:
                var v106 float32 = 3.14159
                _ = v106
            default:
                var v107 float64 = 1.5
                _ = v107
            }
        } else {
            var v108 int8 = 127
            _ = v108
        }
    default:
        fmt.Println(v6)
    }
    v3 = append(v3, "")
    v2 -= (v2 | (v2 | (v2 << ((1) & 63))))
    var v109 uint8 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v109:", v109)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v102
    _ = v103
    _ = v109
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
