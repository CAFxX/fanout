package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int16; F1 uint; F2 string }
type GenS4 struct { F0 int8; F1 int8 }
type GenS5 struct { F0 uint; F1 uint16 }
type GenS6 struct { F0 int16; F1 uint32; F2 int64 }
type GenS7 struct { F0 int64 }
type GenS8 struct { F0 uint16; F1 int }
type GenS9 struct { F0 int64; F1 int32 }
type GenS10 struct { F0 string }
type GenS11 struct { F0 int; F1 int32; F2 bool }
type GenS12 struct { F0 int }
type GenS13 struct { F0 bool; F1 int64; F2 int8 }
func genFn0(p0 uint16) int {
    var v6 uint32 = 65535
    p0 = (p0 ^ 65535)
    var v7 float64 = (3.14159 * 1.5)
    if (!(!true)) {
        if (32767 >= -1) {
            if false {
                v7 = (v7 / v7)
                fmt.Println(v6, v6)
            } else {
                v7 = 3.14159
                var v8 float32 = 1.5
                _ = v8
            }
            for fi9 := 0; fi9 < 12; fi9++ {
                for fi10 := 0; fi10 < 13; fi10++ {
                    p0 = p0
                    fmt.Println(v6, fi9)
                    var v11 float64 = (v7 / 1.0)
                    if false {
                        v12 := GenS1{F0: -1}
                        v11 += v11
                        _ = v12
                    }
                    _ = fi10
                    _ = v11
                }
                _ = fi9
            }
        }
        v13 := []int{1, -2147483648, -9223372036854775808, 1}
        var v14 uint16 = (p0 << ((255) & 63))
        _ = v13
        _ = v14
    }
    v15 := GenS1{F0: 100}
    _ = p0
    _ = v15
    _ = v6
    _ = v7
    return -9223372036854775808
}

func main() {
    var v1 float64 = (float64((-2.25 - (3.14159 / 1e10))) + (((-2.25 - 3.14159) * (0.1 / 1.5)) - ((-2.25 - (3.14159 * 0.0)) - 1.5)))
    v2 := []uint32{2, 2, 162, 436}
    v3 := make(map[int]int)
    v4 := make(map[string]uint32)
    v5 := []uint32{1, 42, 806}
    v16 := GenS2{F0: 42}
    switch 32767 {
    case 128:
        v1 -= ((v1 + ((1e10 / (v1 - (v1 + (((v1 + ((((3.14159 + (v1 - (v1 * float64(0.1)))) + v1) / float64((3.14159 + (3.14159 + ((float32(v1) - -2.25) * -2.25))))) * float64(3.14159))) - 1.5) + v1)))) * (v1 - v1))) + v1)
    case 127:
        v1 += ((float64(float32((1.5 / v1))) / 1e10) * 1e10)
    case 42:
        v1 = v1
    default:
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    for fi17 := 0; fi17 < 12; fi17++ {
        for fi18 := 0; fi18 < 12; fi18++ {
            v1 = (1e10 - ((float64(1.5) * (v1 * (v1 + (float64(3.14159) / v1)))) - float64(0.1)))
            v19 := GenS1{F0: 42}
            v1 -= v1
            _ = fi18
            _ = v19
        }
        _ = fi17
    }
    switch -609 {
    case 127:
        if true {
            switch ("hello" + "a") {
            case "foo bar":
                switch "\n\t" {
                case "hello":
                    v2 = append(v2, 1)
                case "\n\t":
                    v20 := GenS3{F0: -817, F1: 843, F2: "foo bar"}
                    _ = v20
                case "":
                    var v21 int8 = 127
                    _ = v21
                default:
                    v1 = (v1 * (3.14159 - 1e10))
                }
            case "\n\t":
                var v22 float64 = v1
                _ = v22
            case "":
                v23 := GenS1{F0: 42}
                _ = v23
            }
        }
    default:
        v24 := GenS4{F0: 42, F1: 1}
        _ = v24
    }
    v25 := GenS4{F0: -128, F1: -1}
    for fi26 := 0; fi26 < 6; fi26++ {
        for fi27 := 0; fi27 < 19; fi27++ {
            if (32767 == -32768) {
                switch ("" + "foo bar") {
                case "a":
                    v1 = 0.1
                case "hello":
                    v28 := GenS4{F0: 58, F1: 127}
                    if false {
                        v5 = append(v5, 4294967295)
                    } else {
                        fmt.Println(v1, fi27, v1)
                    }
                    _ = v28
                case "\n\t":
                    if false {
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                    } else {
                        switch ("foo bar" + ("\n\t" + ("\n\t" + ""))) {
                        case "foo bar":
                            fmt.Println(fi26, fi27)
                        case "a":
                            v29 := v4["foo bar"]
                            _ = v29
                        case "\n\t":
                            var v30 int8 = -80
                            _ = v30
                        default:
                            var v31 int32 = -2147483648
                            _ = v31
                        }
                        var v32 uint32 = 661
                        _ = v32
                    }
                default:
                    v1 -= 1.5
                }
                v1 -= v1
                fmt.Println(fi27, v1, fi26)
                var v33 int32 = -1
                var v34 int32 = v33
                _ = v33
                _ = v34
            } else {
                v1 = ((v1 - 0.0) / v1)
            }
            _ = fi27
        }
        v1 = 3.14159
        _ = fi26
    }
    fmt.Println(v1)
    switch "" {
    case "hello":
        switch "hello" {
        case "\n\t":
            v35 := GenS5{F0: 4294967295, F1: 2}
            _ = v35
        case "":
            v36 := GenS1{F0: 42}
            _ = v36
        case "a":
            v1 += v1
        }
    case "":
        for fi37 := 0; fi37 < 3; fi37++ {
            for fi38 := 0; fi38 < 19; fi38++ {
                if false {
                    v39 := GenS6{F0: -32768, F1: 292, F2: -39}
                    var v40 int64 = -410
                    var v41 string = ("foo bar" + ("" + "\n\t"))
                    var v42 int64 = 0
                    if (1 != 158) {
                        v40 = 643
                        v2 = append(v2, 613)
                    } else {
                        v40 -= 1
                    }
                    _ = v39
                    _ = v40
                    _ = v41
                    _ = v42
                } else {
                    v43 := v3[fi37]
                    v44 := GenS3{F0: -1, F1: 65535, F2: "\n\t"}
                    v43 = -2147483648
                    var v45 int32 = -187
                    _ = v43
                    _ = v44
                    _ = v45
                }
                switch fi37 {
                case 9223372036854775807:
                    var v46 int32 = -1
                    _ = v46
                case -9223372036854775808:
                    v47 := GenS7{F0: -1}
                    _ = v47
                default:
                    if true {
                        v1 += v1
                    }
                }
                var v48 uint8 = 0
                _ = fi38
                _ = v48
            }
            _ = fi37
        }
    }
    for fi49 := 0; fi49 < 16; fi49++ {
        var v50 int = fi49
        var v51 uint64 = 795
        _ = fi49
        _ = v50
        _ = v51
    }
    if ((v1 > 0.1) || (!false)) {
        v1 -= -2.25
        if true {
            for fi52 := 0; fi52 < 5; fi52++ {
                for fi53 := 0; fi53 < 10; fi53++ {
                    if true {
                        v1 = v1
                        fmt.Println(fi53, fi53)
                    } else {
                        if (fi53 > (fi53 - fi53)) {
                            v54 := v4[(("" + "") + "")]
                            _ = v54
                        } else {
                            if len(v5) > 0 {
                                v5[0] = 2
                            }
                        }
                        if (-2147483648 < 0) {
                            v55 := v4["a"]
                            v56 := GenS8{F0: 32767, F1: 42}
                            v4["a"] = v55
                            if (128 > 255) {
                                var v57 uint = 872
                                switch "\n\t" {
                                case "foo bar":
                                    var v58 uint32 = (v55 % (((v55 - (v55 * (v55 ^ (v55 - v55))))) | 1))
                                    _ = v58
                                case "":
                                    v59 := GenS6{F0: -1, F1: v55, F2: 42}
                                    _ = v59
                                case "hello":
                                    v57 += (^(v57))
                                default:
                                    if false {
                                        switch "foo bar" {
                                        case "a":
                                            _ = len(v5)
                                        }
                                        v57 -= (v57 - (v57 % (((v57 << (((^(v57))) & 63))) | 1)))
                                        var v60 float64 = v1
                                        _ = v60
                                    }
                                }
                                var v61 float32 = -2.25
                                _ = v57
                                _ = v61
                            } else {
                                v62 := GenS1{F0: 14}
                                v3[genFn0(42)] = fi52
                                _ = v62
                            }
                            _ = v55
                            _ = v56
                        } else {
                            var v63 uint64 = 350
                            switch "" {
                            case "hello":
                                if true {
                                    fmt.Println(fi53)
                                    v64 := GenS9{F0: 42, F1: -2147483648}
                                    v63 += (v63 % (((v63 + (v63 - 354))) | 1))
                                    _ = v64
                                } else {
                                    var v65 string = ""
                                    var v66 int64 = 230
                                    var v67 uint8 = 0
                                    _ = v65
                                    _ = v66
                                    _ = v67
                                }
                            case "a":
                                v1 = v1
                            case "":
                                var v68 int16 = -994
                                _ = v68
                            }
                            _ = v63
                        }
                        if false {
                            var v69 uint8 = 128
                            v70 := GenS3{F0: 447, F1: 440, F2: "foo bar"}
                            switch fi52 {
                            case 128:
                                v71 := GenS5{F0: 42, F1: 487}
                                _ = v71
                            case 184:
                                v72 := v4["hello"]
                                _ = v72
                            default:
                                _ = len(v5)
                            }
                            _ = v69
                            _ = v70
                        } else {
                            v73 := v4[(("a" + "a") + "\n\t")]
                            v1 += 1.5
                            var v74 string = "a"
                            _ = v73
                            _ = v74
                        }
                        if (!(((42 != -1) && (128 != 157)) || (42 < 1000))) {
                            v1 -= float64(3.14159)
                            var v75 int8 = 100
                            if false {
                                var v76 float64 = v1
                                var v77 bool = false
                                _ = v76
                                _ = v77
                            } else {
                                if (2147483648 > 2147483648) {
                                    var v78 bool = true
                                    var v79 float64 = float64(1.5)
                                    var v80 int = -1
                                    var v81 uint = 255
                                    _ = v78
                                    _ = v79
                                    _ = v80
                                    _ = v81
                                } else {
                                    if true {
                                        v1 += 0.1
                                    } else {
                                        if true {
                                            switch fi53 {
                                            case -9223372036854775808:
                                                var v82 int16 = 1000
                                                _ = v82
                                            case 127:
                                                v83 := GenS10{F0: "foo bar"}
                                                _ = v83
                                            default:
                                                _ = len(v3)
                                            }
                                            v3[128] = (-(fi53))
                                        } else {
                                            v84 := GenS5{F0: 252, F1: 501}
                                            v85 := GenS11{F0: fi53, F1: -353, F2: false}
                                            var v86 uint8 = 2
                                            _ = v84
                                            _ = v85
                                            _ = v86
                                        }
                                    }
                                    v1 = v1
                                }
                            }
                            _ = v75
                        } else {
                            v1 = v1
                            v1 -= (-2.25 + v1)
                        }
                        if (true && (0.0 != 3.14159)) {
                            var v87 string = "\n\t"
                            var v88 int64 = 655
                            _ = v87
                            _ = v88
                        }
                    }
                    _ = fi53
                }
                fmt.Println(fi52)
                _ = fi52
            }
            if (2 < 42) {
                for fi89 := 0; fi89 < 4; fi89++ {
                    switch genFn0(32767) {
                    case -308:
                        for fi90 := 0; fi90 < 9; fi90++ {
                            var v91 uint8 = 254
                            v92 := GenS12{F0: (-(fi90))}
                            var v93 int = fi90
                            _ = fi90
                            _ = v91
                            _ = v92
                            _ = v93
                        }
                    case 32767:
                        if (float64(0.1) > v1) {
                            v1 = (v1 / ((float64((((-2.25 + -2.25) / 1e10) + 0.1)) - v1) * float64(((((0.1 * 3.14159) * 1e10) * 1e10) - (-2.25 / 0.1)))))
                            v94 := GenS3{F0: 0, F1: 42, F2: ("\n\t" + "hello")}
                            for fi95 := 0; fi95 < 1; fi95++ {
                                fmt.Println(v1)
                                v1 = v1
                                if (4294967295 > 967) {
                                    fmt.Println(fi89, v1, v1)
                                } else {
                                    var v96 int16 = -1
                                    _ = v96
                                }
                                v5 = append(v5, 4294967295)
                                _ = fi95
                            }
                            _ = v94
                        }
                    }
                    _ = fi89
                }
                v1 += v1
            } else {
                v1 += v1
            }
            var v97 uint64 = 42
            _ = v97
        }
        var v98 int = 9223372036854775807
        _ = v98
    }
    for fi99 := 0; fi99 < 14; fi99++ {
        var v100 uint64 = 302
        v101 := GenS13{F0: false, F1: 1, F2: 1}
        _ = fi99
        _ = v100
        _ = v101
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v16
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
