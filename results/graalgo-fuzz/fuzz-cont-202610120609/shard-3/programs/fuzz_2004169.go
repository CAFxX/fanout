package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int8; F1 uint64 }
type GenS3 struct { F0 uint; F1 uint32; F2 bool }
type GenS4 struct { F0 int32; F1 uint8; F2 int64 }
type GenS5 struct { F0 string; F1 int16 }
type GenS6 struct { F0 int32; F1 uint8 }
type GenS7 struct { F0 uint64 }
type GenS8 struct { F0 int32; F1 uint16; F2 uint64 }
type GenS9 struct { F0 uint8; F1 uint64 }
func genFn0(p0 uint16) int16 {
    var v7 uint = 255
    v8 := []int{-688, -529, 127, 32767}
    _ = p0
    _ = v7
    _ = v8
    return 1
}
func genFn1(p0 int64, p1 int) int32 {
    v9 := []int{-1, 42, p1, p1}
    p1 = (p1 * (-(p1)))
    _ = p0
    _ = p1
    _ = v9
    return 766
}

func main() {
    var v1 bool = (3.14159 > 1e10)
    var v2 int32 = 690
    var v3 uint8 = 0
    var v4 int32 = v2
    var v5 bool = (!(378 != 0))
    var v6 string = "a"
    if v5 {
        switch "a" {
        case "foo bar":
            fmt.Println(v5, v3, v2)
        default:
            var v10 int8 = 1
            _ = v10
        }
    } else {
        v11 := GenS1{F0: 1}
        for fi12 := 0; fi12 < 7; fi12++ {
            v13 := GenS1{F0: 255}
            _ = fi12
            _ = v13
        }
        var v14 bool = (!(!(1 < 0)))
        _ = v11
        _ = v14
    }
    if (!v1) {
        for fi15 := 0; fi15 < 18; fi15++ {
            for fi16 := 0; fi16 < 15; fi16++ {
                var v17 int64 = 1
                if v5 {
                    v18 := []int{}
                    v2 += (v2 | (v4 & (v2 / (((v4 + (-(v4)))) | 1))))
                    if (1 == 1) {
                        v19 := GenS2{F0: 127, F1: 1}
                        fmt.Println(fi16, v1, v1)
                        if v5 {
                            _ = len(v18)
                        } else {
                            if ((^(v17)) != v17) {
                                v2 = v4
                                v4 = int32(fi15)
                            } else {
                                v4 = v2
                            }
                        }
                        _ = v19
                    }
                    _ = v18
                }
                var v20 int16 = 693
                if (18446744073709551615 != 42) {
                    if (0 == 62) {
                        fmt.Println(v2)
                        v17 += 0
                    } else {
                        v2 += (v2 - genFn1(v17, (fi16 ^ -1)))
                        fmt.Println(v1)
                    }
                    v1 = (v1 && (fi16 == (fi15 % (((-(fi16))) | 1))))
                    v21 := GenS1{F0: 1}
                    var v22 int8 = 1
                    _ = v21
                    _ = v22
                } else {
                    if ((((3.14159 + (0.0 * 3.14159)) / 1.0) - (0.0 - 3.14159)) < (1.5 / float64(((float32(float64(0.1)) / ((1.5 + 3.14159) + ((1.5 - ((1e10 + 0.1) - ((0.1 + -2.25) * (3.14159 + 0.0)))) - 1e10))) * 1.5)))) {
                        v1 = ("foo bar" == "a")
                        if false {
                            fmt.Println(v1, v1)
                            if true {
                                v23 := []int{-758, fi15, (fi16 * 127), fi15}
                                if true {
                                    if ((255 >= v3) && true) {
                                        fmt.Println(v1, v2)
                                    }
                                    v4 = v4
                                } else {
                                    v1 = v1
                                    v3 = v3
                                    var v24 uint16 = 65535
                                    if false {
                                        v24 += (^(v24))
                                        if v1 {
                                            var v25 int32 = (v4 % (((v4 << ((1) & 63))) | 1))
                                            var v26 uint32 = 2147483647
                                            v27 := GenS1{F0: 65535}
                                            var v28 uint32 = (v26 % ((v26) | 1))
                                            _ = v25
                                            _ = v26
                                            _ = v27
                                            _ = v28
                                        } else {
                                            var v29 int = 4294967295
                                            v30 := GenS3{F0: 101, F1: 316, F2: false}
                                            _ = v29
                                            _ = v30
                                        }
                                        v6 = ("\n\t" + "\n\t")
                                        if len(v23) >= 2 {
                                            v23 = v23[1:len(v23)]
                                        }
                                        fmt.Println(v6, v6)
                                    }
                                    _ = v24
                                }
                                _ = v23
                            }
                            fmt.Println(v20)
                        } else {
                            fmt.Println(v20, v4, v4)
                            if v5 {
                                if (v1 && (v3 <= v3)) {
                                    if v1 {
                                        var v31 int = (fi15 * (fi16 & fi16))
                                        v32 := []int{}
                                        _ = v31
                                        _ = v32
                                    } else {
                                        v33 := GenS1{F0: 2}
                                        v17 = v17
                                        v17 = 1
                                        _ = v33
                                    }
                                } else {
                                    v34 := GenS4{F0: v4, F1: 112, F2: (v17 ^ v17)}
                                    switch v6 {
                                    case "a":
                                        var v35 bool = v1
                                        _ = v35
                                    case "":
                                        var v36 int16 = genFn0(1)
                                        _ = v36
                                    case "\n\t":
                                        fmt.Println(fi16, v17)
                                    }
                                    _ = v34
                                }
                                v3 -= v3
                            } else {
                                v20 += (v20 / ((1) | 1))
                            }
                            v1 = ((335 <= 65535) && (fi16 > (fi16 % ((int(v3)) | 1))))
                        }
                    }
                    var v37 uint = 255
                    if ((v20 | v20) <= (v20 - (^(v20)))) {
                        switch v6 {
                        case "":
                            if (255 == 162) {
                                v17 = (v17 + 0)
                            }
                        case "\n\t":
                            if v1 {
                                v38 := []int{(fi15 / (((fi15 * int(v37))) | 1)), (-(fi15)), fi15, fi15}
                                if len(v38) > 0 {
                                    v38[0] = 1
                                }
                                _ = v38
                            }
                        case "foo bar":
                            var v39 bool = false
                            switch -961 {
                            case -909:
                                fmt.Println(v17, v3, v4)
                            case -2147483648:
                                fmt.Println(v6)
                            case 211:
                                v40 := GenS5{F0: v6, F1: 0}
                                _ = v40
                            }
                            v41 := GenS3{F0: 217, F1: 65535, F2: (((v6 + "hello") == "foo bar") && (1e10 == 1.5))}
                            _ = v39
                            _ = v41
                        }
                    }
                    _ = v37
                }
                _ = fi16
                _ = v17
                _ = v20
            }
            _ = fi15
        }
        var v42 float32 = 0.1
        var v43 uint8 = (v3 ^ (v3 - v3))
        _ = v42
        _ = v43
    }
    switch (v6 + v6) {
    case "":
        for fi44 := 0; fi44 < 10; fi44++ {
            v45 := GenS1{F0: 67}
            for fi46 := 0; fi46 < 5; fi46++ {
                v47 := GenS5{F0: "hello", F1: -346}
                _ = fi46
                _ = v47
            }
            _ = fi44
            _ = v45
        }
    case "hello":
        v48 := GenS5{F0: "hello", F1: 1000}
        _ = v48
    default:
        if (!true) {
            for fi49 := 0; fi49 < 5; fi49++ {
                if true {
                    v50 := GenS6{F0: (v4 >> ((641) & 63)), F1: v3}
                    if v5 {
                        v4 -= v4
                    }
                    if v1 {
                        var v51 uint = 487
                        _ = v51
                    }
                    _ = v50
                } else {
                    v4 -= (v2 + (v2 & (^(v2))))
                    var v52 float32 = (3.14159 + (0.1 - 1.5))
                    _ = v52
                }
                if v5 {
                    if (!v1) {
                        for fi53 := 0; fi53 < 16; fi53++ {
                            var v54 int = (fi49 - fi53)
                            v55 := GenS5{F0: v6, F1: genFn0(572)}
                            _ = fi53
                            _ = v54
                            _ = v55
                        }
                        v56 := GenS4{F0: int32(v3), F1: 0, F2: -57}
                        v3 -= (v3 ^ 86)
                        _ = v56
                    } else {
                        var v57 string = ("\n\t" + v6)
                        switch "" {
                        case "\n\t":
                            for fi58 := 0; fi58 < 1; fi58++ {
                                v59 := []int{}
                                _ = fi58
                                _ = v59
                            }
                        case "":
                            fmt.Println(v4, v3, v4)
                        case "hello":
                            v57 = v6
                        default:
                            v60 := []int{}
                            _ = v60
                        }
                        _ = v57
                    }
                    switch v6 {
                    case "hello":
                        v1 = v5
                    }
                } else {
                    for fi61 := 0; fi61 < 3; fi61++ {
                        var v62 int = fi49
                        v4 += 0
                        v2 += v4
                        _ = fi61
                        _ = v62
                    }
                    fmt.Println(fi49, v5)
                    for fi63 := 0; fi63 < 16; fi63++ {
                        v64 := []int{(^(fi63))}
                        var v65 uint64 = 2
                        _ = fi63
                        _ = v64
                        _ = v65
                    }
                }
                _ = fi49
            }
            for fi66 := 0; fi66 < 11; fi66++ {
                v6 = v6
                var v67 int64 = -9223372036854775808
                var v68 int = fi66
                _ = fi66
                _ = v67
                _ = v68
            }
        }
    }
    var v69 int = 1
    if (2147483648 <= 4294967295) {
        if (-9223372036854775808 < v69) {
            v70 := GenS4{F0: v2, F1: (v3 % ((v3) | 1)), F2: -1}
            v1 = v5
            _ = v70
        } else {
            fmt.Println(v3, v4)
            var v71 uint16 = 348
            var v72 float32 = (0.1 * 1.5)
            _ = v71
            _ = v72
        }
    } else {
        for fi73 := 0; fi73 < 14; fi73++ {
            v74 := GenS3{F0: 255, F1: 65535, F2: ((32767 >= 255) && v1)}
            fmt.Println(v69, v69)
            _ = fi73
            _ = v74
        }
        if (v6 == "") {
            switch (v69 + -743) {
            case -9223372036854775808:
                fmt.Println(v69)
            case 9223372036854775807:
                v75 := []int{}
                _ = v75
            case 0:
                switch 128 {
                case 127:
                    v5 = (1.5 < ((float32(1.5) * ((float32(1.5) + 0.1) - 1e10)) + 3.14159))
                case 9223372036854775807:
                    var v76 uint64 = 947
                    _ = v76
                case 4294967295:
                    v77 := []int{-2147483648, -9223372036854775808, 42}
                    _ = v77
                default:
                    v78 := []int{v69, -285}
                    _ = v78
                }
            default:
                v79 := GenS7{F0: 4294967295}
                _ = v79
            }
        } else {
            v6 = "foo bar"
            for fi80 := 0; fi80 < 20; fi80++ {
                v2 = genFn1(42, -2147483648)
                for fi81 := 0; fi81 < 18; fi81++ {
                    if (228 > 32767) {
                        fmt.Println(fi81)
                        v3 = 127
                        v82 := []int{(v69 >> ((2) & 63)), (v69 >> ((18446744073709551615) & 63)), v69}
                        _ = v82
                    }
                    _ = fi81
                }
                _ = fi80
            }
            for fi83 := 0; fi83 < 18; fi83++ {
                v69 = -877
                for fi84 := 0; fi84 < 18; fi84++ {
                    var v85 float32 = 0.0
                    var v86 float64 = 3.14159
                    if (2147483648 == 4294967295) {
                        fmt.Println(v86, v3)
                        if false {
                            v87 := GenS8{F0: v2, F1: 32767, F2: 395}
                            v2 = (v4 >> ((980) & 63))
                            v3 = (v3 % ((2) | 1))
                            _ = v87
                        }
                    }
                    _ = fi84
                    _ = v85
                    _ = v86
                }
                var v88 int64 = 601
                _ = fi83
                _ = v88
            }
        }
    }
    switch v6 {
    case "":
        var v89 float32 = 0.0
        _ = v89
    case "hello":
        v69 -= v69
    case "\n\t":
        for fi90 := 0; fi90 < 7; fi90++ {
            var v91 uint8 = 2
            _ = fi90
            _ = v91
        }
    }
    var v92 int32 = -2147483648
    var v93 uint = 1
    fmt.Println(v69)
    var v94 uint8 = v3
    v95 := GenS9{F0: (v94 | v94), F1: 204}
    switch v6 {
    case "":
        v96 := GenS5{F0: (v6 + ""), F1: 0}
        _ = v96
    }
    v94 = 255
    if (v1 || v5) {
        v97 := []int{}
        _ = v97
    }
    v98 := []int{(v69 << ((127) & 63))}
    for fi99 := 0; fi99 < 13; fi99++ {
        v2 -= (v4 - (v4 | (v4 - 1)))
        _ = fi99
    }
    switch (^(v69)) {
    case 1:
        switch v69 {
        case 42:
            if len(v98) > 0 {
                v98[0] = (v69 - 1)
            }
        case -2147483648:
            var v100 string = "hello"
            _ = v100
        }
    default:
        v98 = append(v98, v69)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v69:", v69)
    fmt.Println("v92:", v92)
    fmt.Println("v93:", v93)
    fmt.Println("v94:", v94)
    fmt.Println("slicelen v98:", len(v98), v98)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v69
    _ = v92
    _ = v93
    _ = v94
    _ = v95
    _ = v98
}
