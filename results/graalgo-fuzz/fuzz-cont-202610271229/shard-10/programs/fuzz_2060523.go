package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint16; F1 uint64; F2 int }
type GenS3 struct { F0 uint64; F1 uint }
type GenS4 struct { F0 int16; F1 int16 }
type GenS5 struct { F0 uint; F1 uint }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 bool; F1 int; F2 string }
type GenS8 struct { F0 int }
func genFn0() int {
    for fi6 := 0; fi6 < 3; fi6++ {
        v7 := GenS1{F0: ("foo bar" + "\n\t")}
        _ = fi6
        _ = v7
    }
    switch 535 {
    case 459:
        var v8 uint32 = 2147483647
        _ = v8
    case -671:
        v9 := GenS1{F0: "foo bar"}
        _ = v9
    case 32768:
        v10 := GenS2{F0: 537, F1: 986, F2: -479}
        _ = v10
    }
    var v11 int8 = -4
    v12 := []int{-47, 42, 32768, -2147483648}
    _ = v11
    _ = v12
    return 2147483647
}
func genFn1(p0 bool, p1 uint64) uint8 {
    var v13 int8 = 42
    for fi14 := 0; fi14 < 4; fi14++ {
        var v15 float32 = ((1e10 - float32(1e10)) / 1.5)
        var v16 int64 = 42
        _ = fi14
        _ = v15
        _ = v16
    }
    if true {
        for fi17 := 0; fi17 < 8; fi17++ {
            var v18 int16 = 135
            v19 := []int{genFn0(), genFn0(), 32768, (fi17 * (fi17 ^ fi17))}
            v19 = append(v19, 42)
            _ = fi17
            _ = v18
            _ = v19
        }
        if p0 {
            switch -9223372036854775808 {
            case 127:
                var v20 uint = 42
                _ = v20
            default:
                v21 := GenS1{F0: ""}
                _ = v21
            }
            p0 = (!true)
            p1 += p1
        }
        p0 = true
    } else {
        var v22 int = 32767
        for fi23 := 0; fi23 < 16; fi23++ {
            fmt.Println(p1, v13, v13)
            v24 := []int{}
            p0 = true
            _ = fi23
            _ = v24
        }
        v13 = 1
        _ = v22
    }
    _ = p0
    _ = p1
    _ = v13
    return 0
}

func main() {
    var v1 uint = 124
    var v2 int64 = 536
    var v3 int = 32767
    var v4 uint8 = 0
    v5 := []uint32{414}
    v1 -= 2
    if true {
        for fi25 := 0; fi25 < 3; fi25++ {
            v1 -= (v1 << (((v1 - 25)) & 63))
            _ = fi25
        }
        v3 -= v3
        var v26 float32 = 3.14159
        _ = v26
    }
    var v27 uint32 = 51
    v4 = v4
    if len(v5) > 0 {
        v5[0] = v27
    }
    if false {
        var v28 uint64 = 963
        for fi29 := 0; fi29 < 9; fi29++ {
            switch fi29 {
            case 32768:
                for fi30 := 0; fi30 < 9; fi30++ {
                    _ = len(v5)
                    if false {
                        v31 := GenS2{F0: 726, F1: (v28 >> ((1) & 63)), F2: 2147483647}
                        v4 = (^(v4))
                        switch "foo bar" {
                        case "\n\t":
                            fmt.Println(fi29, v28)
                        }
                        v2 = -9223372036854775808
                        _ = v31
                    }
                    var v32 int32 = 1
                    switch ("hello" + "a") {
                    case "a":
                        v33 := GenS2{F0: 0, F1: v28, F2: fi30}
                        _ = v33
                    case "\n\t":
                        var v34 int8 = 36
                        _ = v34
                    case "hello":
                        _ = len(v5)
                    default:
                        fmt.Println(v32, v28, fi30)
                    }
                    _ = fi30
                    _ = v32
                }
            default:
                var v35 float32 = 1e10
                _ = v35
            }
            _ = len(v5)
            if (((1e10 + 3.14159) / (-2.25 + 3.14159)) >= 1.5) {
                v36 := GenS3{F0: 2, F1: (v1 - (v1 + (v1 - (^(v1)))))}
                _ = v36
            }
            _ = fi29
        }
        if (v2 >= (v2 - (v2 - -202))) {
            var v37 string = "\n\t"
            for fi38 := 0; fi38 < 7; fi38++ {
                var v39 int64 = 42
                var v40 int8 = -30
                _ = len(v5)
                _ = fi38
                _ = v39
                _ = v40
            }
            for fi41 := 0; fi41 < 10; fi41++ {
                v42 := GenS3{F0: 0, F1: v1}
                if true {
                    var v43 bool = (37 != 100)
                    _ = v43
                } else {
                    var v44 int16 = 1
                    _ = v44
                }
                fmt.Println(v27, v3, v37)
                _ = fi41
                _ = v42
            }
            _ = v37
        } else {
            fmt.Println(v28, v4, v1)
            v45 := GenS4{F0: 1, F1: 42}
            v4 = uint8(v1)
            _ = v45
        }
        _ = v28
    } else {
        var v46 int8 = 0
        _ = v46
    }
    if false {
        v2 -= v2
        v27 = v27
        v2 = (^(v2))
    } else {
        for fi47 := 0; fi47 < 11; fi47++ {
            if len(v5) > 0 {
                v5[0] = (v27 ^ (v27 << ((v1) & 63)))
            }
            var v48 int64 = -1
            _ = fi47
            _ = v48
        }
        v49 := GenS2{F0: 2, F1: 42, F2: v3}
        _ = v49
    }
    for fi50 := 0; fi50 < 14; fi50++ {
        if true {
            fmt.Println(v2, v27)
        } else {
            v1 -= v1
            if (((false || (genFn1(false, 0) < v4)) && false) || (v1 == v1)) {
                fmt.Println(v4, fi50)
                fmt.Println(fi50, v27, v3)
                switch "foo bar" {
                case "a":
                    var v51 float64 = (((3.14159 / (0.0 - (3.14159 / 1e10))) - 3.14159) / (3.14159 * 0.1))
                    _ = v51
                case "":
                    fmt.Println(v3, v27)
                default:
                    if (!(1e10 > (0.0 / (1.5 / ((1.5 + 0.1) + 0.1))))) {
                        for fi52 := 0; fi52 < 11; fi52++ {
                            _ = len(v5)
                            if ((v1 == v1) && (0 == 529)) {
                                fmt.Println(v4)
                                var v53 bool = true
                                _ = v53
                            } else {
                                if true {
                                    v2 -= v2
                                } else {
                                    var v54 float32 = -2.25
                                    if (!(1e10 == (0.0 * (float64(v54) / -2.25)))) {
                                        v3 = genFn0()
                                        v2 = 9223372036854775807
                                        _ = len(v5)
                                    } else {
                                        var v55 bool = true
                                        fmt.Println(v55, v4, fi50)
                                        v54 = v54
                                        var v56 string = "a"
                                        _ = v55
                                        _ = v56
                                    }
                                    var v57 bool = (!(1 <= (v1 - 0)))
                                    _ = v54
                                    _ = v57
                                }
                            }
                            var v58 uint = 848
                            if (2 > v58) {
                                var v59 uint = (v58 & (v1 >> ((18446744073709551615) & 63)))
                                v3 = int(v27)
                                _ = v59
                            }
                            _ = fi52
                            _ = v58
                        }
                    }
                }
            } else {
                _ = len(v5)
                v60 := GenS2{F0: 65535, F1: 1, F2: fi50}
                _ = v60
            }
        }
        _ = fi50
    }
    v61 := GenS2{F0: 255, F1: 2, F2: v3}
    var v62 uint16 = 440
    var v63 int8 = 0
    v2 += v2
    var v64 string = ("foo bar" + ("hello" + "a"))
    if false {
        fmt.Println(v64, v3)
        var v65 int16 = -32768
        _ = v65
    }
    switch 32767 {
    case -1:
        v66 := GenS5{F0: v1, F1: uint(v27)}
        _ = v66
    case -9223372036854775808:
        v2 = v2
    default:
        var v67 uint64 = 1
        _ = v67
    }
    v63 = v63
    if (42 < 18446744073709551615) {
        if len(v5) > 0 {
            v5[0] = (v27 / ((v27) | 1))
        }
        fmt.Println(v1)
    }
    if ((820 >= int32(v1)) || true) {
        if ((!(954 == uint(v4))) || (-679 >= 42)) {
            v68 := GenS5{F0: (v1 << ((uint(v4)) & 63)), F1: (v1 - v1)}
            fmt.Println(v3, v62, v63)
            if true {
                v69 := GenS6{F0: int32(v62)}
                _ = len(v5)
                _ = v69
            } else {
                for fi70 := 0; fi70 < 19; fi70++ {
                    switch fi70 {
                    case -1:
                        for fi71 := 0; fi71 < 11; fi71++ {
                            v72 := GenS4{F0: 1000, F1: 42}
                            if len(v5) > 0 {
                                v5[0] = 1
                            }
                            _ = fi71
                            _ = v72
                        }
                    }
                    var v73 int = genFn0()
                    _ = fi70
                    _ = v73
                }
                var v74 uint32 = 988
                var v75 uint16 = 32767
                _ = v74
                _ = v75
            }
            _ = v68
        }
    } else {
        if (42 != -11) {
            v62 = 249
            for fi76 := 0; fi76 < 6; fi76++ {
                var v77 uint64 = 2
                fmt.Println(v64, v64)
                v78 := GenS7{F0: (1 < 1), F1: (v3 >> (((v1 % ((v1) | 1))) & 63)), F2: v64}
                _ = fi76
                _ = v77
                _ = v78
            }
            var v79 string = ("" + (((v64 + "\n\t") + "\n\t") + v64))
            _ = v79
        } else {
            for fi80 := 0; fi80 < 17; fi80++ {
                fmt.Println(v64, v27, v4)
                v64 = "hello"
                v27 -= v27
                _ = fi80
            }
            var v81 int8 = (v63 ^ -89)
            var v82 string = ""
            _ = v81
            _ = v82
        }
        v83 := GenS8{F0: genFn0()}
        _ = v83
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v62:", v62)
    fmt.Println("v63:", v63)
    fmt.Println("v64:", v64)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v61
    _ = v62
    _ = v63
    _ = v64
}
