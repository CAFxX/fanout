package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint }
type GenS2 struct { F0 int16; F1 string; F2 string }
type GenS3 struct { F0 uint; F1 string }
type GenS4 struct { F0 uint32; F1 int64 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 uint8 }
func genFn0(p0 string, p1 int32) uint {
    var v4 int64 = 9223372036854775807
    switch p0 {
    case "a":
        if false {
            v5 := GenS1{F0: 1, F1: 255}
            _ = v5
        } else {
            for fi6 := 0; fi6 < 13; fi6++ {
                for fi7 := 0; fi7 < 19; fi7++ {
                    v8 := GenS1{F0: 432, F1: 255}
                    v9 := GenS1{F0: 2, F1: 255}
                    v10 := []int{720, fi7}
                    _ = fi7
                    _ = v10
                    _ = v8
                    _ = v9
                }
                var v11 int16 = 1
                _ = fi6
                _ = v11
            }
            fmt.Println(v4, v4, p0)
        }
    default:
        if (-2.25 < 0.1) {
            v12 := []int{962, 4294967295, 32768, 2147483647}
            for fi13 := 0; fi13 < 13; fi13++ {
                var v14 int16 = 1000
                v15 := GenS1{F0: 4294967295, F1: 162}
                _ = fi13
                _ = v14
                _ = v15
            }
            if len(v12) > 0 {
                v12[0] = 0
            }
            _ = v12
        }
    }
    var v16 uint64 = 18446744073709551615
    for fi17 := 0; fi17 < 17; fi17++ {
        v18 := GenS2{F0: -936, F1: p0, F2: (p0 + "a")}
        v19 := []int{fi17, 9223372036854775807, int(v16)}
        switch (fi17 & (fi17 % ((580) | 1))) {
        case 653:
            v20 := GenS3{F0: 2, F1: p0}
            _ = v20
        case 127:
            fmt.Println(fi17)
        default:
            var v21 uint16 = 32767
            _ = v21
        }
        _ = fi17
        _ = v18
        _ = v19
    }
    _ = p0
    _ = p1
    _ = v16
    _ = v4
    return 65535
}
func genFn1(p0 bool, p1 int16, p2 int32) uint8 {
    var v22 uint16 = 697
    v23 := []int{2147483647, 42, -9223372036854775808, 4294967295}
    if len(v23) > 0 {
        v23[0] = 127
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v22
    _ = v23
    return 2
}

func main() {
    v1 := make(map[int]bool)
    var v2 int16 = -120
    v3 := []uint32{}
    var v24 float32 = 3.14159
    v25 := GenS4{F0: 65535, F1: 382}
    v2 = (v2 / ((v2) | 1))
    var v26 uint32 = 4294967295
    for fi27 := 0; fi27 < 6; fi27++ {
        if (fi27 >= fi27) {
            v28 := GenS3{F0: 1, F1: ("hello" + "foo bar")}
            for fi29 := 0; fi29 < 9; fi29++ {
                var v30 float64 = (float64(-2.25) * -2.25)
                v26 += (^(v26))
                v31 := GenS3{F0: 2, F1: ""}
                _ = fi29
                _ = v30
                _ = v31
            }
            if (true || (3.14159 >= (((1e10 - 0.1) * 1e10) - (0.0 / 3.14159)))) {
                v24 -= 1e10
                var v32 uint32 = v26
                switch "\n\t" {
                case "foo bar":
                    v32 = v26
                case "hello":
                    for fi33 := 0; fi33 < 4; fi33++ {
                        var v34 float64 = 0.0
                        _ = fi33
                        _ = v34
                    }
                case "\n\t":
                    for fi35 := 0; fi35 < 12; fi35++ {
                        v24 += 1e10
                        if true {
                            v36 := GenS3{F0: 18446744073709551615, F1: ""}
                            v2 -= (v2 - 104)
                            v32 += v32
                            _ = v36
                        } else {
                            var v37 float64 = 0.0
                            v32 = v32
                            v2 = -1
                            _ = v37
                        }
                        switch ("\n\t" + "\n\t") {
                        case "\n\t":
                            fmt.Println(fi35)
                        case "foo bar":
                            if len(v3) > 0 {
                                v3[0] = 2
                            }
                        default:
                            var v38 float64 = 1e10
                            _ = v38
                        }
                        _ = fi35
                    }
                default:
                    fmt.Println(v26, v2)
                }
                _ = v32
            } else {
                var v39 uint64 = 2
                _ = v39
            }
            _ = v28
        }
        v24 = 0.1
        v2 += (v2 - v2)
        _ = fi27
    }
    if ((406 >= 2147483647) || (489 > 42)) {
        v2 -= v2
        switch 32767 {
        case -9223372036854775808:
            v40 := GenS2{F0: v2, F1: ("foo bar" + (("\n\t" + "\n\t") + ("" + "hello"))), F2: ""}
            _ = v40
        case 2147483647:
            switch "" {
            case "":
                v2 -= (v2 - (v2 >> ((471) & 63)))
            }
        case 128:
            if (true || ((false || (0 >= 65535)) && (v2 <= v2))) {
                fmt.Println(v26, v24, v2)
                var v41 uint64 = 150
                _ = v41
            } else {
                switch int(v26) {
                case 1:
                    v42 := v1[9223372036854775807]
                    _ = v42
                default:
                    v26 -= (v26 ^ (v26 >> ((2) & 63)))
                }
                v26 = (v26 + v26)
            }
        default:
            v43 := GenS5{F0: v2}
            _ = v43
        }
        if false {
            var v44 int16 = (-(v2))
            _ = v44
        }
    } else {
        var v45 int64 = -762
        for fi46 := 0; fi46 < 17; fi46++ {
            v47 := GenS2{F0: -32768, F1: "foo bar", F2: "\n\t"}
            var v48 bool = (((v26 * (^(v26))) > (^(v26))) || ((v2 - v2) == v2))
            var v49 int = (fi46 & fi46)
            _ = fi46
            _ = v47
            _ = v48
            _ = v49
        }
        _ = v45
    }
    var v50 float64 = 3.14159
    fmt.Println(v2, v2, v26)
    if false {
        v2 -= v2
        v2 = -32768
    }
    v26 = (v26 - v26)
    if (float64((v24 / ((((v24 + (float32(v50) - v24)) * float32(float64((0.1 + v24)))) / float32(1e10)) / ((1e10 + (v24 + v24)) * float32(float64(v24)))))) <= ((v50 * ((v50 * 0.1) * 3.14159)) + 0.0)) {
        v26 -= 2147483648
        switch -2147483648 {
        case -1:
            for fi51 := 0; fi51 < 3; fi51++ {
                v50 = ((1e10 - (3.14159 - (v50 - 1.5))) / v50)
                for fi52 := 0; fi52 < 3; fi52++ {
                    v53 := GenS6{F0: 130}
                    v2 -= (v2 * (v2 / (((v2 << ((0) & 63))) | 1)))
                    _ = fi52
                    _ = v53
                }
                _ = fi51
            }
        case 9223372036854775807:
            switch "\n\t" {
            case "foo bar":
                switch 0 {
                case 803:
                    v50 += 0.1
                case 4294967295:
                    v54 := GenS5{F0: (v2 ^ -745)}
                    _ = v54
                case 32768:
                    v26 += v26
                }
            case "":
                for fi55 := 0; fi55 < 16; fi55++ {
                    fmt.Println(v26, v26, v24)
                    _ = fi55
                }
            }
        case -398:
            for fi56 := 0; fi56 < 18; fi56++ {
                v26 = (v26 / ((v26) | 1))
                if false {
                    v24 = (v24 * (1e10 - float32(v50)))
                } else {
                    for fi57 := 0; fi57 < 3; fi57++ {
                        var v58 int8 = 127
                        _ = fi57
                        _ = v58
                    }
                    var v59 int16 = int16(v26)
                    _ = v59
                }
                _ = fi56
            }
        }
        v2 = -32768
    }
    for fi60 := 0; fi60 < 14; fi60++ {
        for fi61 := 0; fi61 < 6; fi61++ {
            var v62 uint16 = 1
            v63 := GenS2{F0: v2, F1: "", F2: (("" + ("hello" + "\n\t")) + "\n\t")}
            _ = fi61
            _ = v62
            _ = v63
        }
        delete(v1, fi60)
        var v64 int64 = 9223372036854775807
        _ = fi60
        _ = v64
    }
    for fi65 := 0; fi65 < 6; fi65++ {
        var v66 float32 = v24
        _ = fi65
        _ = v66
    }
    var v67 string = (("\n\t" + (("" + "hello") + "a")) + "a")
    var v68 string = v67
    switch 2147483647 {
    case 1:
        v69 := GenS3{F0: genFn0(v67, -278), F1: v68}
        _ = v69
    case 2147483647:
        v24 -= (v24 / 1.0)
    default:
        if (v26 != (^(v26))) {
            for fi70 := 0; fi70 < 1; fi70++ {
                v71 := GenS2{F0: v2, F1: v67, F2: ("a" + "foo bar")}
                _ = fi70
                _ = v71
            }
        } else {
            switch 2147483647 {
            case 4294967295:
                switch 128 {
                case 32767:
                    var v72 uint32 = 2147483648
                    _ = v72
                }
            case -9223372036854775808:
                if false {
                    for fi73 := 0; fi73 < 19; fi73++ {
                        v24 = v24
                        v74 := GenS5{F0: v2}
                        if true {
                            delete(v1, fi73)
                            switch "" {
                            case "foo bar":
                                switch (("foo bar" + v68) + ((v67 + v68) + "a")) {
                                case "hello":
                                    if (85 <= 42) {
                                        v75 := GenS6{F0: 172}
                                        if (!false) {
                                            for fi76 := 0; fi76 < 19; fi76++ {
                                                switch "" {
                                                case "foo bar":
                                                    var v77 uint = 553
                                                    _ = v77
                                                case "":
                                                    v24 += float32(v50)
                                                default:
                                                    _ = len(v1)
                                                }
                                                v24 = 1.5
                                                _ = fi76
                                            }
                                        } else {
                                            if len(v3) >= 2 {
                                                v3 = v3[1:len(v3)]
                                            }
                                            switch 32767 {
                                            case 32768:
                                                v78 := GenS6{F0: 165}
                                                _ = v78
                                            case 855:
                                                v79 := GenS3{F0: 427, F1: "foo bar"}
                                                _ = v79
                                            }
                                        }
                                        _ = v75
                                    }
                                case "a":
                                    var v80 uint = 266
                                    _ = v80
                                }
                            case "":
                                for fi81 := 0; fi81 < 18; fi81++ {
                                    if true {
                                        var v82 float32 = v24
                                        _ = v82
                                    }
                                    v24 = 0.0
                                    if (0 >= 42) {
                                        switch "foo bar" {
                                        case "hello":
                                            v24 = v24
                                        }
                                        v50 += (0.1 / v50)
                                    }
                                    _ = fi81
                                }
                            }
                            for fi83 := 0; fi83 < 13; fi83++ {
                                var v84 uint8 = 128
                                v84 -= v84
                                v85 := GenS1{F0: (v26 * (v26 & v26)), F1: 296}
                                if true {
                                    var v86 uint32 = v26
                                    v86 = v26
                                    var v87 int64 = 1
                                    _ = v86
                                    _ = v87
                                }
                                _ = fi83
                                _ = v84
                                _ = v85
                            }
                        } else {
                            v26 = 780
                            v24 += (v24 + v24)
                        }
                        _ = fi73
                        _ = v74
                    }
                }
            default:
                v50 += v50
            }
            for fi88 := 0; fi88 < 2; fi88++ {
                if len(v3) > 0 {
                    v3[0] = v26
                }
                _ = fi88
            }
            for fi89 := 0; fi89 < 6; fi89++ {
                switch "hello" {
                case "":
                    var v90 uint32 = v26
                    _ = v90
                case "\n\t":
                    v91 := GenS1{F0: v26, F1: 652}
                    _ = v91
                }
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                _ = fi89
            }
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v50:", v50)
    fmt.Println("v67:", v67)
    fmt.Println("v68:", v68)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v50
    _ = v67
    _ = v68
}
