package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 bool }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 bool; F1 uint16 }
type GenS5 struct { F0 uint8; F1 bool }
type GenS6 struct { F0 int8; F1 bool }
type GenS7 struct { F0 int64; F1 string; F2 int8 }
type GenS8 struct { F0 uint64; F1 uint8 }
type GenS9 struct { F0 uint64; F1 int16; F2 uint }
type GenS10 struct { F0 int8; F1 uint8; F2 uint8 }
func genFn0(p0 uint64) string {
    var v7 string = "a"
    fmt.Println(v7, v7)
    _ = p0
    _ = v7
    return v7
}
func genFn1() int8 {
    var v8 uint32 = 268
    v9 := []int{524, 128, 587, 691}
    fmt.Println(v8)
    var v10 uint8 = 55
    fmt.Println(v10, v8)
    _ = v10
    _ = v8
    _ = v9
    return 0
}

func main() {
    var v1 uint32 = 2
    var v2 uint = 720
    var v3 string = "foo bar"
    var v4 bool = ((v1 >> (((v2 + v2)) & 63)) >= (v1 >> ((1) & 63)))
    var v5 uint16 = 1
    var v6 uint32 = v1
    v11 := []int{9223372036854775807}
    switch -9223372036854775808 {
    case 9223372036854775807:
        v3 = ("\n\t" + ((("hello" + "") + v3) + v3))
    case 0:
        for fi12 := 0; fi12 < 13; fi12++ {
            if ((v2 ^ (v2 >> ((uint(v1)) & 63))) < (v2 * v2)) {
                if true {
                    var v13 uint32 = 1
                    _ = v13
                } else {
                    fmt.Println(v6, fi12)
                }
                fmt.Println(v3, v3, v5)
                var v14 uint32 = 0
                _ = v14
            } else {
                fmt.Println(v3)
                if len(v11) >= 2 {
                    v11 = v11[1:len(v11)]
                }
                v5 -= 987
            }
            if v4 {
                v2 += uint(v1)
                var v15 uint32 = v6
                v6 -= v1
                _ = v15
            }
            _ = fi12
        }
    default:
        v16 := GenS1{F0: 2147483647, F1: (!v4)}
        _ = v16
    }
    if false {
        v3 = v3
        v3 = ("foo bar" + ("a" + ""))
    }
    v1 -= (v1 * (v6 | 319))
    var v17 uint64 = 466
    var v18 int8 = 0
    for fi19 := 0; fi19 < 7; fi19++ {
        var v20 int8 = (v18 + (v18 & (v18 * genFn1())))
        _ = fi19
        _ = v20
    }
    v3 = "hello"
    if len(v11) >= 2 {
        v11 = v11[1:len(v11)]
    }
    var v21 uint8 = 128
    for fi22 := 0; fi22 < 10; fi22++ {
        if false {
            if false {
                v23 := GenS2{F0: -1}
                _ = v23
            } else {
                v4 = (!(647 > 2))
            }
            switch v3 {
            case "":
                switch v3 {
                case "a":
                    switch int(v17) {
                    case 32767:
                        fmt.Println(v1, v17)
                    case -2147483648:
                        if len(v11) >= 2 {
                            v11 = v11[1:len(v11)]
                        }
                    case 241:
                        switch (fi22 & (fi22 >> (((v2 - (v2 - v2))) & 63))) {
                        case -866:
                            if ((!v4) && (v3 > v3)) {
                                var v24 string = ""
                                var v25 uint = 65535
                                v11 = append(v11, 42)
                                _ = v24
                                _ = v25
                            }
                        default:
                            switch "\n\t" {
                            case "foo bar":
                                v26 := GenS3{F0: v17}
                                _ = v26
                            case "a":
                                var v27 int64 = -429
                                _ = v27
                            }
                        }
                    }
                }
            default:
                switch fi22 {
                case 551:
                    if len(v11) >= 2 {
                        v11 = v11[1:len(v11)]
                    }
                case -1:
                    v3 = ("a" + v3)
                case 1:
                    v3 = v3
                }
            }
            var v28 string = v3
            _ = v28
        }
        var v29 uint16 = (^(v5))
        if (v4 && true) {
            var v30 int8 = (v18 | v18)
            v11 = append(v11, (fi22 + fi22))
            fmt.Println(v3, v17)
            _ = v30
        }
        _ = fi22
        _ = v29
    }
    var v31 float32 = 0.0
    if (uint32(v2) != 2147483647) {
        var v32 int8 = 46
        v31 = v31
        if len(v11) >= 2 {
            v11 = v11[1:len(v11)]
        }
        _ = v32
    }
    for fi33 := 0; fi33 < 15; fi33++ {
        if len(v11) > 0 {
            v11[0] = (-(fi33))
        }
        for fi34 := 0; fi34 < 16; fi34++ {
            v35 := GenS3{F0: v17}
            fmt.Println(fi33)
            var v36 uint64 = (v17 * v17)
            _ = fi34
            _ = v35
            _ = v36
        }
        _ = fi33
    }
    switch (v3 + "\n\t") {
    case "":
        for fi37 := 0; fi37 < 12; fi37++ {
            var v38 bool = (v5 != (^(v5)))
            var v39 uint = v2
            v11 = append(v11, (fi37 ^ int(v21)))
            _ = fi37
            _ = v38
            _ = v39
        }
    }
    if len(v11) > 0 {
        v11[0] = int(v18)
    }
    if true {
        v11 = append(v11, -920)
        if v4 {
            v5 = 42
        } else {
            v40 := GenS1{F0: -370, F1: (v4 || false)}
            _ = v40
        }
    } else {
        switch "\n\t" {
        case "hello":
            v5 += (v5 % ((v5) | 1))
        case "a":
            if v4 {
                var v41 float32 = v31
                v42 := GenS4{F0: true, F1: (v5 & v5)}
                var v43 uint32 = v6
                _ = v41
                _ = v42
                _ = v43
            } else {
                var v44 uint32 = v6
                if false {
                    var v45 float64 = 3.14159
                    if v4 {
                        var v46 uint32 = uint32(v2)
                        if true {
                            v3 = "\n\t"
                            v45 = 0.0
                            switch 32767 {
                            case 127:
                                if v4 {
                                    for fi47 := 0; fi47 < 1; fi47++ {
                                        switch -1 {
                                        case 127:
                                            for fi48 := 0; fi48 < 11; fi48++ {
                                                v44 = (v6 ^ (v1 + (v44 >> ((4294967295) & 63))))
                                                var v49 uint64 = v17
                                                if false {
                                                    if ((fi48 ^ fi47) != fi47) {
                                                        var v50 float64 = v45
                                                        if false {
                                                            v18 = v18
                                                            v18 = (v18 - (v18 ^ genFn1()))
                                                            v44 = v1
                                                            fmt.Println(v2)
                                                        }
                                                        v51 := GenS1{F0: (fi47 ^ -2147483648), F1: (!true)}
                                                        _ = v50
                                                        _ = v51
                                                    }
                                                }
                                                v52 := GenS1{F0: (-(fi48)), F1: (v4 || ((v2 ^ (v2 | v2)) >= 80))}
                                                switch v3 {
                                                case "a":
                                                    var v53 float32 = ((v31 * v31) + v31)
                                                    _ = v53
                                                case "foo bar":
                                                    if (80 != -330) {
                                                        if v4 {
                                                            if (uint64(v5) >= (v17 | v17)) {
                                                                v4 = (v21 == (v21 * (^(v21))))
                                                                if (v4 && v4) {
                                                                    if false {
                                                                        v46 = v1
                                                                        v54 := GenS1{F0: 493, F1: v4}
                                                                        if len(v11) >= 2 {
                                                                            v11 = v11[1:len(v11)]
                                                                        }
                                                                        fmt.Println(v17, v18)
                                                                        _ = v54
                                                                    } else {
                                                                        v55 := GenS2{F0: 1}
                                                                        _ = v55
                                                                    }
                                                                    fmt.Println(fi48, v17)
                                                                }
                                                            } else {
                                                                _ = len(v11)
                                                                v1 = v44
                                                            }
                                                            v4 = true
                                                        }
                                                    } else {
                                                        v56 := GenS5{F0: (^(v21)), F1: v4}
                                                        v57 := GenS6{F0: (v18 >> (((^(v2))) & 63)), F1: (!(9223372036854775807 == -9223372036854775808))}
                                                        var v58 int32 = 2147483647
                                                        _ = v56
                                                        _ = v57
                                                        _ = v58
                                                    }
                                                case "":
                                                    v2 = 128
                                                }
                                                v11 = append(v11, int(v2))
                                                _ = fi48
                                                _ = v49
                                                _ = v52
                                            }
                                        case -9223372036854775808:
                                            var v59 int32 = 0
                                            _ = v59
                                        case -787:
                                            var v60 uint32 = (^(v44))
                                            _ = v60
                                        }
                                        _ = fi47
                                    }
                                } else {
                                    switch 32767 {
                                    case -2147483648:
                                        var v61 uint32 = (v6 * (v6 / ((2147483647) | 1)))
                                        _ = v61
                                    case -718:
                                        var v62 int8 = (v18 ^ v18)
                                        _ = v62
                                    case 580:
                                        if v4 {
                                            var v63 string = (v3 + (v3 + (((v3 + (v3 + v3)) + (v3 + (("foo bar" + v3) + "foo bar"))) + ("a" + v3))))
                                            _ = v63
                                        } else {
                                            v64 := GenS3{F0: (v17 >> ((v2) & 63))}
                                            _ = v64
                                        }
                                    default:
                                        var v65 float64 = (v45 / ((((-2.25 * v45) + 3.14159) * ((v45 - v45) / ((0.1 / v45) / v45))) + v45))
                                        _ = v65
                                    }
                                    var v66 int16 = int16(v21)
                                    _ = v66
                                }
                            case 128:
                                for fi67 := 0; fi67 < 7; fi67++ {
                                    if v4 {
                                        var v68 float32 = v31
                                        _ = v68
                                    } else {
                                        var v69 float32 = v31
                                        v70 := GenS1{F0: -380, F1: (!(!((v69 / ((v69 - v69) + float32(v45))) >= (v69 + 3.14159))))}
                                        v71 := GenS5{F0: v21, F1: (-613 == -461)}
                                        _ = v69
                                        _ = v70
                                        _ = v71
                                    }
                                    v72 := GenS2{F0: -497}
                                    v73 := GenS7{F0: -52, F1: (v3 + ((v3 + "\n\t") + v3)), F2: (v18 << ((1) & 63))}
                                    _ = fi67
                                    _ = v72
                                    _ = v73
                                }
                            }
                        } else {
                            if len(v11) > 0 {
                                v11[0] = 42
                            }
                            switch 9223372036854775807 {
                            case -9223372036854775808:
                                for fi74 := 0; fi74 < 10; fi74++ {
                                    switch -854 {
                                    case -9223372036854775808:
                                        var v75 int = fi74
                                        _ = v75
                                    default:
                                        if len(v11) >= 2 {
                                            v11 = v11[1:len(v11)]
                                        }
                                    }
                                    switch "\n\t" {
                                    case "foo bar":
                                        for fi76 := 0; fi76 < 10; fi76++ {
                                            var v77 int32 = -2147483648
                                            v4 = ("hello" != ("hello" + (("a" + "a") + (v3 + ((v3 + "a") + "\n\t")))))
                                            _ = fi76
                                            _ = v77
                                        }
                                    case "a":
                                        v78 := GenS6{F0: (v18 >> ((944) & 63)), F1: v4}
                                        _ = v78
                                    case "hello":
                                        v31 = float32(float64((v31 - 1e10)))
                                    default:
                                        var v79 uint64 = v17
                                        _ = v79
                                    }
                                    if v4 {
                                        v4 = false
                                        for fi80 := 0; fi80 < 17; fi80++ {
                                            if len(v11) >= 2 {
                                                v11 = v11[1:len(v11)]
                                            }
                                            v1 += v1
                                            var v81 uint32 = v1
                                            _ = fi80
                                            _ = v81
                                        }
                                        var v82 uint32 = v1
                                        _ = v82
                                    }
                                    _ = fi74
                                }
                            case -1:
                                var v83 int16 = 42
                                _ = v83
                            case 4294967295:
                                if (1 == 42) {
                                    var v84 float32 = float32(((float64((v31 + 1e10)) + (3.14159 / 0.1)) / (v45 * float64(0.0))))
                                    v85 := GenS8{F0: (^(v17)), F1: 127}
                                    _ = v84
                                    _ = v85
                                } else {
                                    if len(v11) >= 2 {
                                        v11 = v11[1:len(v11)]
                                    }
                                    var v86 bool = true
                                    _ = v86
                                }
                            default:
                                for fi87 := 0; fi87 < 13; fi87++ {
                                    if v4 {
                                        v88 := GenS9{F0: (v17 << (((v2 >> (((v2 % ((v2) | 1))) & 63))) & 63)), F1: 303, F2: uint(v21)}
                                        fmt.Println(v1)
                                        _ = v88
                                    } else {
                                        var v89 int32 = 1
                                        _ = v89
                                    }
                                    if (!v4) {
                                        fmt.Println(v17, v3)
                                        for fi90 := 0; fi90 < 13; fi90++ {
                                            var v91 int64 = 42
                                            var v92 float64 = v45
                                            _ = fi90
                                            _ = v91
                                            _ = v92
                                        }
                                    } else {
                                        if false {
                                            var v93 float64 = -2.25
                                            _ = v93
                                        } else {
                                            v11 = append(v11, (fi87 * fi87))
                                            var v94 float32 = (((0.1 / (v31 + (v31 / 1.0))) / ((0.0 + 0.1) + (3.14159 * v31))) + ((v31 * v31) + (v31 + float32(float64(0.0)))))
                                            _ = v94
                                        }
                                        var v95 string = v3
                                        _ = v95
                                    }
                                    _ = fi87
                                }
                            }
                            var v96 uint8 = v21
                            _ = v96
                        }
                        v97 := GenS10{F0: int8(v5), F1: uint8(v17), F2: v21}
                        _ = v46
                        _ = v97
                    } else {
                        for fi98 := 0; fi98 < 14; fi98++ {
                            if len(v11) >= 2 {
                                v11 = v11[1:len(v11)]
                            }
                            if len(v11) > 0 {
                                v11[0] = fi98
                            }
                            _ = fi98
                        }
                        v5 = 568
                    }
                    v4 = (143 >= 42)
                    _ = v45
                }
                _ = v44
            }
        case "foo bar":
            v31 = v31
        }
    }
    var v99 uint = v2
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v99:", v99)
    _ = v1
    _ = v11
    _ = v17
    _ = v18
    _ = v2
    _ = v21
    _ = v3
    _ = v31
    _ = v4
    _ = v5
    _ = v6
    _ = v99
}
