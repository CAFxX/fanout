package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint; F2 uint8 }
type GenS2 struct { F0 int16; F1 uint16 }
type GenS3 struct { F0 string; F1 int32; F2 int }
type GenS4 struct { F0 int; F1 string }
type GenS5 struct { F0 int16; F1 uint }
type GenS6 struct { F0 bool; F1 bool; F2 uint64 }
type GenS7 struct { F0 int8 }
func genFn0(p0 int) int16 {
    if true {
        if true {
            p0 -= p0
            fmt.Println(p0)
            fmt.Println(p0)
        }
        v6 := []int{}
        v7 := GenS1{F0: 735, F1: 255, F2: 0}
        _ = v6
        _ = v7
    }
    var v8 float64 = 1.5
    var v9 uint = 431
    _ = p0
    _ = v8
    _ = v9
    return 0
}

func main() {
    var v1 uint32 = 1
    var v2 int64 = 960
    var v3 int = 0
    var v4 int32 = int32(v1)
    var v5 float64 = 0.0
    v10 := []int{(v3 * (v3 * (v3 << ((255) & 63))))}
    if false {
        for fi11 := 0; fi11 < 12; fi11++ {
            var v12 uint32 = (v1 - 1)
            v4 += (v4 >> ((18446744073709551615) & 63))
            _ = fi11
            _ = v12
        }
    } else {
        for fi13 := 0; fi13 < 13; fi13++ {
            if ((1e10 + 3.14159) >= ((float64(0.0) + (((float64(((1.5 / 1e10) - (3.14159 - 0.0))) / ((float64((0.0 * 0.0)) - (0.0 + 0.1)) - v5)) / v5) * (1e10 / (0.1 + ((v5 + (0.1 * (1.5 - (((1e10 * v5) * v5) / float64(-2.25))))) * v5))))) * 3.14159)) {
                v1 -= 0
                if len(v10) >= 2 {
                    v10 = v10[1:len(v10)]
                }
                if len(v10) > 0 {
                    v10[0] = v3
                }
            }
            _ = fi13
        }
    }
    v4 += (v4 * int32(v2))
    var v14 string = ""
    v4 = 1
    v15 := GenS1{F0: 42, F1: 506, F2: 42}
    v16 := GenS1{F0: 65535, F1: 65535, F2: 1}
    v5 -= 0.0
    v17 := GenS2{F0: 1, F1: 255}
    if true {
        if true {
            v5 = v5
            var v18 uint16 = 65535
            _ = len(v10)
            _ = v18
        } else {
            if false {
                v19 := GenS2{F0: 1, F1: 65535}
                for fi20 := 0; fi20 < 9; fi20++ {
                    if len(v10) >= 2 {
                        v10 = v10[1:len(v10)]
                    }
                    _ = fi20
                }
                for fi21 := 0; fi21 < 5; fi21++ {
                    var v22 uint64 = 2
                    if false {
                        v22 += v22
                        v23 := GenS2{F0: 1, F1: 32767}
                        _ = v23
                    } else {
                        for fi24 := 0; fi24 < 19; fi24++ {
                            if (v14 <= "hello") {
                                if len(v10) > 0 {
                                    v10[0] = fi21
                                }
                                var v25 float64 = v5
                                _ = len(v10)
                                _ = v25
                            } else {
                                v3 += fi24
                                if true {
                                    var v26 uint64 = v22
                                    if len(v10) >= 2 {
                                        v10 = v10[1:len(v10)]
                                    }
                                    v27 := GenS2{F0: genFn0((fi24 & 0)), F1: 255}
                                    _ = v26
                                    _ = v27
                                }
                                v28 := GenS2{F0: 4, F1: uint16(v22)}
                                _ = v28
                            }
                            v14 = v14
                            _ = fi24
                        }
                    }
                    _ = fi21
                    _ = v22
                }
                _ = v19
            }
            _ = len(v10)
        }
        v10 = append(v10, (v3 | v3))
        v29 := GenS3{F0: v14, F1: 1, F2: (v3 % ((0) | 1))}
        _ = v29
    }
    var v30 bool = true
    v31 := GenS1{F0: 1, F1: 1, F2: 28}
    var v32 uint8 = 127
    v33 := GenS4{F0: (^(v3)), F1: (v14 + v14)}
    for fi34 := 0; fi34 < 9; fi34++ {
        for fi35 := 0; fi35 < 8; fi35++ {
            v36 := GenS3{F0: (v14 + v14), F1: (v4 - v4), F2: fi35}
            fmt.Println(v2, v5, fi35)
            if (!v30) {
                v32 += 128
                _ = len(v10)
                switch fi34 {
                case 437:
                    if ((v32 % ((v32) | 1)) == (v32 >> ((18446744073709551615) & 63))) {
                        v32 -= 42
                    }
                case -2147483648:
                    v37 := GenS5{F0: -1, F1: 815}
                    if len(v10) > 0 {
                        v10[0] = (fi35 % (((fi34 << ((18446744073709551615) & 63))) | 1))
                    }
                    v38 := GenS1{F0: 18446744073709551615, F1: uint(v1), F2: (^(v32))}
                    var v39 int32 = v4
                    _ = v37
                    _ = v38
                    _ = v39
                default:
                    fmt.Println(fi35)
                }
            } else {
                var v40 uint8 = 2
                if (false && v30) {
                    if len(v10) > 0 {
                        v10[0] = (^(fi35))
                    }
                    if v30 {
                        v40 += (^(v32))
                        var v41 int = fi34
                        if false {
                            if len(v10) >= 2 {
                                v10 = v10[1:len(v10)]
                            }
                            if true {
                                fmt.Println(v3)
                            }
                        } else {
                            switch 444 {
                            case 128:
                                if false {
                                    var v42 string = (v14 + "a")
                                    _ = v42
                                } else {
                                    if (fi35 >= 2147483647) {
                                        v43 := GenS4{F0: (v41 << ((1) & 63)), F1: "foo bar"}
                                        var v44 uint64 = 0
                                        v3 += (-(fi34))
                                        _ = v43
                                        _ = v44
                                    }
                                    if (10 == 0) {
                                        var v45 uint8 = (v40 % ((v40) | 1))
                                        if v30 {
                                            var v46 int64 = (v2 & 52)
                                            if false {
                                                v47 := GenS5{F0: genFn0(-812), F1: 247}
                                                v30 = (v4 < v4)
                                                _ = v47
                                            } else {
                                                if v30 {
                                                    v1 -= v1
                                                    var v48 uint = 4294967295
                                                    _ = v48
                                                } else {
                                                    v49 := GenS4{F0: fi35, F1: ((("a" + (v14 + ("a" + (("\n\t" + v14) + v14)))) + ("foo bar" + (v14 + v14))) + v14)}
                                                    _ = v49
                                                }
                                                switch "a" {
                                                case "a":
                                                    if (v1 <= (v1 / ((v1) | 1))) {
                                                        if len(v10) > 0 {
                                                            v10[0] = (v41 | -1)
                                                        }
                                                        var v50 string = "\n\t"
                                                        if v30 {
                                                            fmt.Println(v41, v32, v45)
                                                        }
                                                        _ = v50
                                                    } else {
                                                        v41 = fi35
                                                        var v51 uint = 65535
                                                        v10 = append(v10, (^(v3)))
                                                        _ = v51
                                                    }
                                                case "":
                                                    var v52 float64 = float64(1e10)
                                                    if false {
                                                        if len(v10) >= 2 {
                                                            v10 = v10[1:len(v10)]
                                                        }
                                                    }
                                                    _ = v52
                                                case "foo bar":
                                                    if len(v10) > 0 {
                                                        v10[0] = v3
                                                    }
                                                }
                                            }
                                            v32 += 0
                                            _ = v46
                                        } else {
                                            var v53 string = (v14 + ((v14 + v14) + v14))
                                            var v54 bool = ((!(0 > (v1 + (v1 >> ((952) & 63))))) || (((v1 - v1) < v1) || v30))
                                            var v55 float64 = float64(3.14159)
                                            switch 0 {
                                            case -1:
                                                switch (v14 + (v14 + (v14 + ((v53 + v53) + v14)))) {
                                                case "\n\t":
                                                    v30 = true
                                                case "hello":
                                                    var v56 uint8 = v40
                                                    _ = v56
                                                case "a":
                                                    var v57 uint32 = v1
                                                    _ = v57
                                                default:
                                                    v30 = (false || true)
                                                }
                                            case 9223372036854775807:
                                                switch (fi34 << ((18446744073709551615) & 63)) {
                                                case -1:
                                                    var v58 float32 = ((0.0 + 0.1) + ((1.5 * 0.1) / 1.5))
                                                    _ = v58
                                                case -73:
                                                    if len(v10) >= 2 {
                                                        v10 = v10[1:len(v10)]
                                                    }
                                                    var v59 uint64 = 497
                                                    _ = v59
                                                default:
                                                    var v60 uint32 = v1
                                                    _ = v60
                                                }
                                            default:
                                                v30 = (v54 || v54)
                                                v61 := GenS6{F0: v54, F1: false, F2: 0}
                                                _ = v61
                                            }
                                            v55 += v5
                                            var v62 uint64 = 0
                                            _ = v53
                                            _ = v54
                                            _ = v55
                                            _ = v62
                                        }
                                        v5 += v5
                                        v63 := GenS6{F0: v30, F1: (!(v30 && (2 <= 42))), F2: 2}
                                        _ = v45
                                        _ = v63
                                    } else {
                                        _ = len(v10)
                                        var v64 uint32 = (^(v1))
                                        if v30 {
                                            v65 := GenS3{F0: (v14 + "\n\t"), F1: -2147483648, F2: (v41 + int(v4))}
                                            var v66 int = 0
                                            v10 = append(v10, v66)
                                            _ = v65
                                            _ = v66
                                        } else {
                                            var v67 int16 = 0
                                            switch -335 {
                                            case 987:
                                                if len(v10) >= 2 {
                                                    v10 = v10[1:len(v10)]
                                                }
                                                if len(v10) >= 2 {
                                                    v10 = v10[1:len(v10)]
                                                }
                                                switch (^(fi34)) {
                                                case 127:
                                                    if len(v10) > 0 {
                                                        v10[0] = v3
                                                    }
                                                case 91:
                                                    var v68 string = v14
                                                    _ = v68
                                                case 128:
                                                    var v69 int = v3
                                                    _ = v69
                                                }
                                                if v30 {
                                                    var v70 int = fi34
                                                    _ = v70
                                                }
                                            }
                                            if ((v4 % ((int32(v67)) | 1)) > v4) {
                                                switch ((v14 + v14) + v14) {
                                                case "foo bar":
                                                    switch ((("hello" + "foo bar") + v14) + v14) {
                                                    case "\n\t":
                                                        v71 := GenS7{F0: 1}
                                                        _ = v71
                                                    case "foo bar":
                                                        v10 = append(v10, int(v40))
                                                    case "hello":
                                                        var v72 uint = 65535
                                                        _ = v72
                                                    }
                                                }
                                                v5 -= 1e10
                                            }
                                            _ = v67
                                        }
                                        _ = v64
                                    }
                                    v40 = 223
                                    if v30 {
                                        v40 += (^(v40))
                                        v30 = v30
                                    }
                                }
                            case 42:
                                var v73 string = (v14 + v14)
                                _ = v73
                            case -1:
                                if ((!(-2147483648 != v41)) && ((0.1 / 0.1) < (1e10 * (-2.25 * 1.5)))) {
                                    if (!true) {
                                        var v74 int8 = int8(v1)
                                        switch -9223372036854775808 {
                                        case 128:
                                            if v30 {
                                                v40 = v32
                                                var v75 string = "foo bar"
                                                _ = v75
                                            }
                                        case 1:
                                            v3 += fi34
                                        }
                                        switch v14 {
                                        case "foo bar":
                                            var v76 bool = (32767 == int16(v74))
                                            _ = v76
                                        default:
                                            var v77 int16 = -1
                                            _ = v77
                                        }
                                        _ = v74
                                    } else {
                                        var v78 uint64 = 42
                                        v40 = 49
                                        _ = v78
                                    }
                                    if len(v10) > 0 {
                                        v10[0] = (v3 * -295)
                                    }
                                }
                            }
                            fmt.Println(v3)
                        }
                        v79 := GenS3{F0: (v14 + "a"), F1: (v4 / ((v4) | 1)), F2: (v41 ^ (v41 ^ (v41 ^ fi34)))}
                        _ = v41
                        _ = v79
                    }
                    var v80 uint = 0
                    _ = v80
                }
                _ = v40
            }
            v14 = v14
            var v81 uint64 = 18446744073709551615
            _ = fi35
            _ = v36
            _ = v81
        }
        if len(v10) > 0 {
            v10[0] = (fi34 << ((79) & 63))
        }
        var v82 uint8 = (v32 << ((0) & 63))
        _ = fi34
        _ = v82
    }
    if len(v10) > 0 {
        v10[0] = -996
    }
    if len(v10) >= 2 {
        v10 = v10[1:len(v10)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v4
    _ = v5
}
