package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint64; F2 uint8 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int8; F1 int }
type GenS4 struct { F0 bool; F1 uint32; F2 uint }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 uint8; F1 string; F2 uint8 }
func genFn0() int64 {
    v6 := GenS1{F0: 32767, F1: 4294967295, F2: 221}
    v7 := GenS1{F0: -894, F1: 1, F2: 0}
    var v8 float64 = ((((0.0 - (0.0 / 1e10)) - float64(float32((-2.25 * 1.5)))) * 1e10) - 0.1)
    fmt.Println(v8)
    v8 -= (v8 / 1.0)
    _ = v6
    _ = v7
    _ = v8
    return -89
}
func genFn1() int64 {
    for fi9 := 0; fi9 < 20; fi9++ {
        if (18446744073709551615 > 2) {
            var v10 bool = true
            v10 = true
            var v11 float32 = 0.1
            _ = v10
            _ = v11
        }
        if true {
            var v12 int64 = -9223372036854775808
            _ = v12
        }
        _ = fi9
    }
    v13 := []int{9223372036854775807, -679}
    for fi14 := 0; fi14 < 3; fi14++ {
        var v15 int32 = -707
        for fi16 := 0; fi16 < 7; fi16++ {
            v17 := GenS2{F0: 32767}
            if true {
                switch fi14 {
                case 128:
                    var v18 uint16 = 42
                    _ = v18
                case 42:
                    fmt.Println(v15, fi14)
                }
                fmt.Println(fi16)
                var v19 uint8 = 83
                var v20 uint8 = v19
                _ = v19
                _ = v20
            } else {
                var v21 uint64 = 505
                if ((-76 < 97) && (true && ((-2.25 * -2.25) == 1.5))) {
                    switch fi14 {
                    case 271:
                        if (1 == 128) {
                            v21 -= 1
                            if (1 != 61) {
                                var v22 int16 = 32767
                                _ = v22
                            } else {
                                var v23 bool = (255 <= 0)
                                v15 = 1
                                _ = v23
                            }
                        }
                    case -9223372036854775808:
                        v21 = v21
                    case 896:
                        var v24 uint32 = 65535
                        _ = v24
                    default:
                        fmt.Println(fi14, fi16)
                        switch fi14 {
                        case -782:
                            v21 = (^(v21))
                        case -810:
                            v21 -= v21
                        case 128:
                            if len(v13) > 0 {
                                v13[0] = fi16
                            }
                        default:
                            var v25 int32 = 765
                            _ = v25
                        }
                        v13 = append(v13, (fi16 * fi16))
                    }
                }
                _ = v21
            }
            v13 = append(v13, fi14)
            v15 -= (-(v15))
            var v26 uint32 = 2147483648
            _ = fi16
            _ = v17
            _ = v26
        }
        switch "" {
        case "hello":
            if (42 != 42) {
                var v27 bool = (false && true)
                if len(v13) > 0 {
                    v13[0] = fi14
                }
                var v28 uint8 = 1
                _ = v27
                _ = v28
            }
        }
        _ = fi14
        _ = v15
    }
    if len(v13) > 0 {
        v13[0] = -9223372036854775808
    }
    var v29 int = 2147483647
    _ = v13
    _ = v29
    return 42
}

func main() {
    v1 := []int{-1, 884, 9223372036854775807, 9223372036854775807}
    v2 := make(map[uint32]uint32)
    var v3 int64 = 9223372036854775807
    v4 := make(map[int]uint32)
    v5 := make(map[string]string)
    var v30 uint8 = 0
    var v31 float64 = 1e10
    var v32 int = -996
    for fi33 := 0; fi33 < 4; fi33++ {
        v30 = 127
        v34 := GenS2{F0: 703}
        if (530 < 0) {
            switch "foo bar" {
            case "hello":
                v1 = append(v1, (fi33 * (v32 + fi33)))
            case "a":
                switch (^(v32)) {
                case 1:
                    if false {
                        v35 := GenS3{F0: -98, F1: (v32 - v32)}
                        v36 := GenS3{F0: -128, F1: 9223372036854775807}
                        for fi37 := 0; fi37 < 8; fi37++ {
                            v30 = v30
                            v38 := v2[1]
                            _ = fi37
                            _ = v38
                        }
                        _ = v35
                        _ = v36
                    } else {
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        for fi39 := 0; fi39 < 17; fi39++ {
                            _ = len(v1)
                            switch (fi39 >> ((18446744073709551615) & 63)) {
                            case -935:
                                switch 2147483647 {
                                case 42:
                                    var v40 string = ""
                                    _ = v40
                                case 32767:
                                    v41 := GenS2{F0: 0}
                                    _ = v41
                                case 32768:
                                    fmt.Println(v32, v30, v3)
                                default:
                                    v31 -= v31
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                    var v42 float64 = (3.14159 + 1.5)
                                    _ = v42
                                }
                            }
                            if true {
                                _ = len(v1)
                                var v43 int64 = -9223372036854775808
                                _ = v43
                            } else {
                                if len(v1) > 0 {
                                    v1[0] = (v32 - (v32 + (v32 + fi33)))
                                }
                            }
                            _ = fi39
                        }
                        for fi44 := 0; fi44 < 15; fi44++ {
                            var v45 int16 = 32767
                            _ = fi44
                            _ = v45
                        }
                    }
                case -9223372036854775808:
                    v32 = fi33
                }
            default:
                fmt.Println(v31)
            }
        }
        _ = fi33
        _ = v34
    }
    v30 -= 0
    for fi46 := 0; fi46 < 9; fi46++ {
        for fi47 := 0; fi47 < 7; fi47++ {
            if ((42 == 42) && true) {
                var v48 uint32 = 283
                _ = v48
            } else {
                v49 := GenS1{F0: 953, F1: 2, F2: (^(v30))}
                switch (^(fi47)) {
                case -911:
                    switch 9223372036854775807 {
                    case 0:
                        v50 := GenS3{F0: -1, F1: -2147483648}
                        _ = v50
                    case 62:
                        if len(v1) > 0 {
                            v1[0] = fi47
                        }
                        v31 = (float64(((0.1 * (3.14159 * 1e10)) + 1.5)) / ((v31 + ((3.14159 * 0.0) / (((v31 * v31) * v31) + (3.14159 * 0.1)))) - (((1.5 / v31) + float64(3.14159)) + float64((float32(v31) - (((0.1 / float32(((((3.14159 + v31) + (1e10 - v31)) - v31) / 0.1))) - 0.1) * float32(0.1)))))))
                        v30 -= (v30 | (v30 - (v30 / ((42) | 1))))
                    case 128:
                        var v51 uint16 = 65535
                        _ = v51
                    default:
                        fmt.Println(v32, v3)
                    }
                case 128:
                    var v52 string = "\n\t"
                    _ = v52
                }
                v32 += 0
                _ = v49
            }
            _ = fi47
        }
        _ = fi46
    }
    fmt.Println(v31, v3, v32)
    v31 -= 1e10
    delete(v2, 0)
    v53 := v4[2147483647]
    v53 = 65535
    if (((v31 - 3.14159) == v31) && (true && (true || (!true)))) {
        var v54 uint = 390
        v55 := GenS4{F0: (false || (!false)), F1: 2147483647, F2: v54}
        _ = v54
        _ = v55
    }
    if ((!((false && (1 != 409)) && (!false))) || ("a" == "\n\t")) {
        v3 = (v3 - genFn1())
        v56 := v4[0]
        if false {
            v31 += 0.1
        } else {
            var v57 uint32 = (v53 * (v53 / ((823) | 1)))
            v31 = (3.14159 + (0.1 / (v31 / -2.25)))
            _ = v57
        }
        _ = v56
    } else {
        var v58 int64 = (v3 ^ v3)
        v59 := GenS3{F0: 100, F1: (^(v32))}
        if true {
            v60 := GenS5{F0: v53}
            _ = v60
        } else {
            for fi61 := 0; fi61 < 9; fi61++ {
                v62 := GenS6{F0: (v30 << ((406) & 63)), F1: "foo bar", F2: v30}
                _ = fi61
                _ = v62
            }
        }
        _ = v58
        _ = v59
    }
    for fi63 := 0; fi63 < 6; fi63++ {
        if (!false) {
            for fi64 := 0; fi64 < 7; fi64++ {
                v31 -= v31
                if false {
                    if false {
                        v32 -= v32
                    } else {
                        if false {
                            if len(v1) > 0 {
                                v1[0] = 32767
                            }
                        } else {
                            var v65 uint8 = uint8(v53)
                            _ = len(v1)
                            _ = len(v1)
                            if false {
                                var v66 int16 = int16(v53)
                                v53 += v53
                                var v67 uint32 = v53
                                _ = v66
                                _ = v67
                            } else {
                                var v68 float64 = 0.1
                                var v69 bool = true
                                _ = v68
                                _ = v69
                            }
                            if true {
                                if (73 == -128) {
                                    v65 = 255
                                    fmt.Println(v53, fi63, v65)
                                } else {
                                    if (v3 == -436) {
                                        if false {
                                            v70 := GenS2{F0: 32767}
                                            _ = v70
                                        }
                                    } else {
                                        var v71 float32 = 0.0
                                        fmt.Println(v31, v30)
                                        _ = v71
                                    }
                                    var v72 float32 = 0.1
                                    _ = v72
                                }
                                v32 += (fi64 + int(v3))
                                var v73 uint8 = 127
                                switch "hello" {
                                case "\n\t":
                                    fmt.Println(v31, v53, v73)
                                default:
                                    delete(v4, (fi64 ^ fi63))
                                }
                                _ = v73
                            } else {
                                v31 = 3.14159
                                if (42 != 255) {
                                    var v74 int64 = v3
                                    _ = v74
                                }
                                fmt.Println(v53, fi64, v32)
                                v75 := GenS6{F0: (v30 / ((v65) | 1)), F1: "", F2: v65}
                                var v76 int16 = -994
                                _ = v75
                                _ = v76
                            }
                            _ = v65
                        }
                    }
                    switch (("" + ("hello" + ("foo bar" + (("a" + ("" + "")) + "\n\t")))) + ((((("\n\t" + "hello") + "\n\t") + ("foo bar" + "foo bar")) + "a") + ("" + "a"))) {
                    case "hello":
                        v32 += int(v30)
                    case "\n\t":
                        _ = len(v1)
                    default:
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                    }
                    switch (fi64 | fi63) {
                    case 42:
                        var v77 uint16 = uint16(v53)
                        _ = v77
                    }
                    v78 := v2[v53]
                    v31 -= v31
                    _ = v78
                }
                v4[fi63] = (v53 ^ 42)
                _ = fi64
            }
            switch "\n\t" {
            case "foo bar":
                var v79 uint64 = 0
                _ = v79
            case "a":
                var v80 uint16 = uint16(v30)
                _ = v80
            }
            var v81 uint8 = 42
            _ = v81
        } else {
            if false {
                switch ("foo bar" + "") {
                case "hello":
                    v31 += v31
                case "foo bar":
                    if true {
                        for fi82 := 0; fi82 < 15; fi82++ {
                            v2[4294967295] = (v53 << ((2) & 63))
                            delete(v5, "a")
                            _ = fi82
                        }
                    } else {
                        v83 := GenS2{F0: 0}
                        var v84 string = ("foo bar" + ("hello" + "a"))
                        _ = v83
                        _ = v84
                    }
                default:
                    var v85 uint = 1
                    _ = v85
                }
                v86 := GenS1{F0: 32767, F1: 1, F2: 2}
                for fi87 := 0; fi87 < 16; fi87++ {
                    var v88 uint8 = (v30 << ((65535) & 63))
                    var v89 float64 = 0.1
                    v90 := GenS5{F0: v53}
                    _ = fi87
                    _ = v88
                    _ = v89
                    _ = v90
                }
                _ = v86
            } else {
                v3 = int64(v32)
            }
            if true {
                var v91 int64 = (^(v3))
                var v92 uint64 = 2
                var v93 int = -9223372036854775808
                _ = v91
                _ = v92
                _ = v93
            }
        }
        _ = fi63
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v53:", v53)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v53
}
