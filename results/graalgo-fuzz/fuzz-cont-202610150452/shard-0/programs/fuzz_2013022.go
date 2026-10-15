package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 string; F2 bool }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint8; F1 int; F2 string }
type GenS4 struct { F0 uint32; F1 int16 }
type GenS5 struct { F0 uint16; F1 int32; F2 int32 }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 int16; F1 bool }
type GenS8 struct { F0 bool; F1 uint16; F2 uint8 }
type GenS9 struct { F0 uint32; F1 int64; F2 uint32 }
type GenS10 struct { F0 uint16 }
func genFn0() bool {
    var v6 uint = 42
    v6 -= (v6 - 65535)
    fmt.Println(v6)
    for fi7 := 0; fi7 < 14; fi7++ {
        v6 += v6
        if (927 < 18446744073709551615) {
            v8 := []int{(fi7 << ((v6) & 63)), fi7}
            var v9 uint8 = 3
            _ = v8
            _ = v9
        }
        switch "a" {
        case "":
            v10 := GenS1{F0: 1, F1: (("" + "hello") + ("hello" + (("\n\t" + "a") + "\n\t"))), F2: true}
            _ = v10
        case "a":
            if true {
                if true {
                    var v11 int8 = 125
                    _ = v11
                }
            }
        case "\n\t":
            var v12 int16 = 0
            _ = v12
        default:
            if false {
                v6 += 4294967295
                for fi13 := 0; fi13 < 10; fi13++ {
                    if false {
                        v14 := []int{}
                        fmt.Println(fi13, v6, fi7)
                        v6 += v6
                        var v15 uint32 = 65535
                        switch fi13 {
                        case 32767:
                            v6 = 827
                        case -9223372036854775808:
                            if (144 != 164) {
                                v16 := GenS1{F0: v6, F1: "", F2: false}
                                if false {
                                    v17 := GenS1{F0: v6, F1: ("" + ((("hello" + ("foo bar" + ("\n\t" + "hello"))) + (("a" + "\n\t") + ("hello" + ""))) + "hello")), F2: true}
                                    _ = v17
                                }
                                _ = v16
                            }
                        default:
                            v6 = (v6 & v6)
                        }
                        _ = v14
                        _ = v15
                    }
                    var v18 bool = false
                    v6 = 2
                    v19 := []int{}
                    _ = fi13
                    _ = v18
                    _ = v19
                }
            }
        }
        _ = fi7
    }
    switch 997 {
    case 127:
        v6 -= v6
    case 4294967295:
        if false {
            v20 := []int{127, 527, -915, -915}
            if false {
                if true {
                    if (-2.25 > 0.1) {
                        var v21 string = "foo bar"
                        for fi22 := 0; fi22 < 3; fi22++ {
                            switch v21 {
                            case "\n\t":
                                fmt.Println(v6, v6)
                            case "a":
                                var v23 uint16 = 719
                                _ = v23
                            case "":
                                for fi24 := 0; fi24 < 17; fi24++ {
                                    if (false || false) {
                                        var v25 int32 = 0
                                        _ = len(v20)
                                        _ = v25
                                    } else {
                                        var v26 float32 = 0.1
                                        var v27 bool = true
                                        _ = v26
                                        _ = v27
                                    }
                                    var v28 uint8 = 128
                                    v6 -= 0
                                    _ = fi24
                                    _ = v28
                                }
                            }
                            v21 = v21
                            _ = fi22
                        }
                        switch -1 {
                        case 0:
                            switch v21 {
                            case "":
                                var v29 int16 = 0
                                _ = v29
                            case "foo bar":
                                v30 := GenS2{F0: 100}
                                _ = v30
                            case "a":
                                var v31 uint8 = 2
                                _ = v31
                            }
                        case -821:
                            v32 := GenS2{F0: 0}
                            _ = v32
                        default:
                            fmt.Println(v21, v6)
                        }
                        _ = v21
                    }
                } else {
                    var v33 int32 = -2147483648
                    v6 = (v6 + (v6 ^ (v6 >> ((v6) & 63))))
                    if len(v20) > 0 {
                        v20[0] = -192
                    }
                    _ = v33
                }
            } else {
                if false {
                    v34 := GenS3{F0: 255, F1: 127, F2: ("" + "hello")}
                    if len(v20) > 0 {
                        v20[0] = 42
                    }
                    _ = v34
                }
                if false {
                    var v35 uint64 = 860
                    v6 -= v6
                    _ = v35
                } else {
                    for fi36 := 0; fi36 < 13; fi36++ {
                        if true {
                            v6 = v6
                        }
                        _ = fi36
                    }
                }
                for fi37 := 0; fi37 < 5; fi37++ {
                    v38 := GenS4{F0: 703, F1: 220}
                    if ((false && true) || ((0.0 * -2.25) != -2.25)) {
                        for fi39 := 0; fi39 < 10; fi39++ {
                            var v40 float32 = 1e10
                            _ = fi39
                            _ = v40
                        }
                    }
                    _ = fi37
                    _ = v38
                }
            }
            _ = v20
        } else {
            var v41 int = 494
            var v42 float64 = (0.0 * -2.25)
            var v43 uint8 = 1
            _ = v41
            _ = v42
            _ = v43
        }
    default:
        for fi44 := 0; fi44 < 18; fi44++ {
            if true {
                v45 := []int{9223372036854775807}
                _ = v45
            }
            if false {
                var v46 float32 = (0.1 - (0.0 + -2.25))
                var v47 float32 = float32((-2.25 / 1.5))
                _ = v46
                _ = v47
            }
            v48 := []int{fi44, (fi44 | (fi44 * fi44))}
            _ = fi44
            _ = v48
        }
    }
    _ = v6
    return true
}

func main() {
    v1 := make(map[int]bool)
    var v2 int = 127
    v3 := []uint32{}
    var v4 uint32 = 428
    var v5 uint32 = v4
    switch -2147483648 {
    case 2147483647:
        switch "hello" {
        case "a":
            if ((!true) || false) {
                v5 += 475
                v4 = 65535
            }
        case "":
            v4 += (v4 / ((42) | 1))
        case "foo bar":
            for fi49 := 0; fi49 < 14; fi49++ {
                for fi50 := 0; fi50 < 5; fi50++ {
                    fmt.Println(v5, fi50, v4)
                    _ = fi50
                }
                switch "" {
                case "\n\t":
                    _ = len(v1)
                case "foo bar":
                    if (128 >= 106) {
                        var v51 float64 = -2.25
                        var v52 int64 = -9223372036854775808
                        var v53 float32 = 1.5
                        _ = v51
                        _ = v52
                        _ = v53
                    }
                case "":
                    for fi54 := 0; fi54 < 17; fi54++ {
                        switch "hello" {
                        case "a":
                            var v55 uint64 = 18446744073709551615
                            _ = v55
                        case "foo bar":
                            fmt.Println(v2, v2, v2)
                        case "\n\t":
                            _ = len(v1)
                            var v56 bool = false
                            v57 := GenS2{F0: 23}
                            _ = v56
                            _ = v57
                        default:
                            v58 := GenS5{F0: 2, F1: -2147483648, F2: -1}
                            _ = v58
                        }
                        v2 = -611
                        _ = fi54
                    }
                default:
                    fmt.Println(v2, fi49, v2)
                }
                _ = fi49
            }
        }
    case 32768:
        var v59 int32 = 775
        _ = v59
    }
    fmt.Println(v2)
    var v60 bool = (!true)
    v60 = ((2147483647 <= int32(v4)) && (65535 != 817))
    v2 = (v2 | (^(v2)))
    v60 = (!v60)
    fmt.Println(v5, v5, v5)
    v4 += 887
    switch ("" + ("" + "hello")) {
    case "hello":
        for fi61 := 0; fi61 < 20; fi61++ {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            fmt.Println(v2, fi61, v60)
            for fi62 := 0; fi62 < 6; fi62++ {
                v63 := GenS2{F0: -1}
                fmt.Println(v5, v4)
                v5 += (v5 % (((v5 & v5)) | 1))
                _ = fi62
                _ = v63
            }
            _ = fi61
        }
    }
    for fi64 := 0; fi64 < 7; fi64++ {
        v4 = 446
        var v65 uint = 1
        if v60 {
            v66 := GenS6{F0: 0}
            v2 = fi64
            _ = v66
        } else {
            if v60 {
                for fi67 := 0; fi67 < 17; fi67++ {
                    v68 := GenS2{F0: -128}
                    v65 = (v65 * 0)
                    switch (fi67 / ((fi67) | 1)) {
                    case -505:
                        v69 := GenS7{F0: -1, F1: v60}
                        _ = v69
                    case -875:
                        if ((^(v4)) == v4) {
                            var v70 bool = v60
                            if v60 {
                                if v60 {
                                    if (2 < 65535) {
                                        if (898 == 2147483648) {
                                            v65 += 71
                                            switch 312 {
                                            case 9223372036854775807:
                                                var v71 int16 = -524
                                                _ = v71
                                            }
                                        }
                                    } else {
                                        if v60 {
                                            if len(v3) >= 2 {
                                                v3 = v3[1:len(v3)]
                                            }
                                            v4 += v5
                                            v72 := GenS6{F0: 839}
                                            _ = v72
                                        } else {
                                            if true {
                                                v73 := v1[(fi64 << ((v65) & 63))]
                                                if ((!(!v73)) || v73) {
                                                    v74 := v1[9223372036854775807]
                                                    if v70 {
                                                        var v75 int8 = 1
                                                        _ = v75
                                                    }
                                                    var v76 bool = (-2.25 == float64(0.1))
                                                    _ = v74
                                                    _ = v76
                                                } else {
                                                    _ = len(v3)
                                                    if (0 == 1) {
                                                        _ = len(v3)
                                                        _ = len(v3)
                                                    }
                                                    var v77 uint64 = 440
                                                    _ = v77
                                                }
                                                _ = v73
                                            } else {
                                                var v78 uint = v65
                                                _ = v78
                                            }
                                            v4 = uint32(v65)
                                        }
                                        switch -9223372036854775808 {
                                        case 32767:
                                            var v79 int = (v2 | -9223372036854775808)
                                            _ = v79
                                        default:
                                            v2 = -9223372036854775808
                                        }
                                        v4 = 2147483647
                                    }
                                }
                                _ = len(v1)
                                v80 := GenS5{F0: 2, F1: -479, F2: 42}
                                v65 += v65
                                switch 531 {
                                case 4294967295:
                                    v70 = v70
                                case 32767:
                                    var v81 int8 = 104
                                    _ = v81
                                }
                                _ = v80
                            }
                            _ = v70
                        }
                    }
                    var v82 uint64 = 1
                    _ = fi67
                    _ = v68
                    _ = v82
                }
            } else {
                if false {
                    var v83 float64 = 1.5
                    if true {
                        if ((1.5 <= (0.1 + -2.25)) && (0.1 <= (-2.25 + -2.25))) {
                            v5 = (v5 >> (((v65 / (((v65 - (v65 - (v65 >> ((18446744073709551615) & 63))))) | 1))) & 63))
                        }
                        if (!false) {
                            var v84 bool = (!(1 < 558))
                            v84 = (!v84)
                            switch "a" {
                            case "hello":
                                var v85 int8 = 0
                                _ = v85
                            case "":
                                var v86 int16 = -32768
                                _ = v86
                            case "foo bar":
                                var v87 uint64 = 1
                                _ = v87
                            }
                            _ = v84
                        } else {
                            fmt.Println(v83)
                        }
                    }
                    for fi88 := 0; fi88 < 3; fi88++ {
                        v89 := GenS8{F0: v60, F1: 2, F2: 127}
                        if (("foo bar" + ("hello" + ("a" + ("hello" + ("\n\t" + ""))))) > "foo bar") {
                            fmt.Println(v65, v65, v60)
                            var v90 float64 = -2.25
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            var v91 float64 = 1e10
                            var v92 uint16 = 183
                            _ = v90
                            _ = v91
                            _ = v92
                        } else {
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            v93 := GenS9{F0: 4294967295, F1: 9223372036854775807, F2: v5}
                            fmt.Println(v60, v5, v83)
                            _ = v93
                        }
                        v2 -= (v2 & (v2 / (((v2 << ((v65) & 63))) | 1)))
                        _ = fi88
                        _ = v89
                    }
                    _ = v83
                }
                for fi94 := 0; fi94 < 10; fi94++ {
                    v5 = (v5 - 690)
                    if v60 {
                        if false {
                            if false {
                                delete(v1, fi94)
                            } else {
                                _ = len(v3)
                            }
                        }
                        v4 = v5
                        v95 := GenS10{F0: 505}
                        _ = v95
                    } else {
                        v65 += 426
                    }
                    _ = fi94
                }
                v65 -= v65
            }
            v2 -= v2
            v96 := GenS3{F0: uint8(v65), F1: (fi64 - -385), F2: ""}
            _ = v96
        }
        _ = fi64
        _ = v65
    }
    switch 32767 {
    case -9223372036854775808:
        if v60 {
            v97 := GenS9{F0: 2, F1: -1, F2: 325}
            _ = v97
        }
    case 32768:
        v60 = v60
    case 127:
        var v98 string = (("" + "hello") + ("a" + ("" + ("foo bar" + ((((("\n\t" + "foo bar") + "foo bar") + "hello") + "hello") + "\n\t")))))
        _ = v98
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v60:", v60)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v60
}
