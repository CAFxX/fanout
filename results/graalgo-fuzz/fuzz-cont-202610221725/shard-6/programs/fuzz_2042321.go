package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64; F2 int16 }
type GenS2 struct { F0 uint16; F1 int8; F2 int }
type GenS3 struct { F0 bool; F1 uint64 }
type GenS4 struct { F0 uint8; F1 string; F2 uint64 }
type GenS5 struct { F0 bool; F1 int16; F2 uint }
func genFn0(p0 uint8, p1 int16) string {
    v7 := GenS1{F0: 156, F1: -191, F2: p1}
    p1 += (p1 << ((165) & 63))
    p0 = (p0 / (((p0 % ((1) | 1))) | 1))
    _ = p0
    _ = p1
    _ = v7
    return ""
}

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[string]int)
    v3 := []int{4294967295}
    v4 := make(map[uint32]int)
    v5 := []string{("foo bar" + "\n\t"), "foo bar", ("a" + "hello"), "foo bar"}
    v6 := make(map[uint32]uint32)
    var v8 bool = true
    if v8 {
        fmt.Println(v8)
        v8 = (!true)
    }
    if false {
        var v9 string = ("foo bar" + "a")
        switch 76 {
        case 4294967295:
            v10 := v4[1]
            _ = v10
        case -321:
            for fi11 := 0; fi11 < 12; fi11++ {
                if len(v3) > 0 {
                    v3[0] = 1
                }
                var v12 int = fi11
                v13 := GenS2{F0: 0, F1: -124, F2: 128}
                _ = fi11
                _ = v12
                _ = v13
            }
        }
        _ = v9
    }
    var v14 uint16 = 458
    var v15 int8 = int8(v14)
    fmt.Println(v14)
    switch "foo bar" {
    case "\n\t":
        var v16 uint64 = 271
        _ = v16
    case "hello":
        if (!v8) {
            v15 -= (^(v15))
        }
    default:
        fmt.Println(v8, v14, v15)
    }
    if (-1 > 9223372036854775807) {
        if true {
            for fi17 := 0; fi17 < 18; fi17++ {
                v14 = v14
                var v18 uint16 = 0
                _ = fi17
                _ = v18
            }
            var v19 uint64 = 0
            _ = v19
        } else {
            v14 -= v14
            _ = len(v4)
            v15 = v15
        }
        for fi20 := 0; fi20 < 4; fi20++ {
            v21 := GenS1{F0: 743, F1: -339, F2: -549}
            var v22 float64 = 0.0
            for fi23 := 0; fi23 < 12; fi23++ {
                fmt.Println(fi20)
                if true {
                    if (!false) {
                        if (!v8) {
                            v15 = (v15 ^ (v15 >> ((255) & 63)))
                        } else {
                            if v8 {
                                v14 += (v14 + (v14 | v14))
                                _ = len(v5)
                                if v8 {
                                    if ("hello" <= "") {
                                        v22 = (v22 / (v22 + v22))
                                        if v8 {
                                            if v8 {
                                                var v24 uint = 4294967295
                                                var v25 int = (fi23 >> ((v24) & 63))
                                                _ = v24
                                                _ = v25
                                            } else {
                                                v22 -= 0.1
                                            }
                                            var v26 uint = 42
                                            fmt.Println(v14, fi23, fi20)
                                            _ = v26
                                        } else {
                                            v27 := GenS3{F0: (v8 || (65535 > 415)), F1: 2}
                                            v8 = false
                                            _ = v27
                                        }
                                    } else {
                                        v5 = append(v5, "hello")
                                        switch (("" + "a") + "a") {
                                        case "foo bar":
                                            delete(v4, 42)
                                        case "\n\t":
                                            v28 := GenS3{F0: ((((v22 * (v22 / -2.25)) / v22) + ((1.5 * v22) / v22)) > float64(1.5)), F1: 4294967295}
                                            _ = v28
                                        default:
                                            v14 -= 0
                                        }
                                        if (!(v15 < (^(v15)))) {
                                            var v29 int16 = -600
                                            var v30 uint64 = 163
                                            _ = v29
                                            _ = v30
                                        } else {
                                            if ((fi20 == (fi20 - (fi20 / (((fi23 - (fi23 & (-(fi23))))) | 1)))) && (375 >= 2147483648)) {
                                                v14 = (v14 | 1)
                                                fmt.Println(fi23, fi20)
                                            } else {
                                                if len(v5) >= 2 {
                                                    v5 = v5[1:len(v5)]
                                                }
                                                var v31 uint16 = v14
                                                _ = v31
                                            }
                                        }
                                    }
                                }
                            }
                        }
                        var v32 uint8 = 93
                        _ = v32
                    } else {
                        v33 := GenS4{F0: 127, F1: "foo bar", F2: 772}
                        fmt.Println(v15, v8)
                        _ = v33
                    }
                    fmt.Println(v15, fi20)
                    v34 := GenS2{F0: (v14 << ((1) & 63)), F1: v15, F2: (fi23 | (fi20 / ((fi20) | 1)))}
                    v8 = (0 >= 320)
                    _ = v34
                } else {
                    v8 = v8
                }
                _ = fi23
            }
            _ = fi20
            _ = v21
            _ = v22
        }
    } else {
        for fi35 := 0; fi35 < 4; fi35++ {
            if v8 {
                var v36 int16 = -29
                v8 = (!(!((!(42 == 2147483647)) && (!(0 != 850)))))
                _ = v36
            } else {
                v37 := GenS2{F0: (v14 & (v14 / ((1) | 1))), F1: v15, F2: (fi35 ^ 42)}
                _ = v37
            }
            v15 -= v15
            for fi38 := 0; fi38 < 2; fi38++ {
                if (1 >= 255) {
                    var v39 uint64 = 4294967295
                    _ = v39
                } else {
                    var v40 int16 = -832
                    _ = len(v1)
                    _ = len(v4)
                    v14 = 315
                    switch 32767 {
                    case -2147483648:
                        if true {
                            switch fi35 {
                            case -952:
                                v41 := GenS2{F0: (v14 / (((v14 * v14)) | 1)), F1: (v15 >> ((18446744073709551615) & 63)), F2: 127}
                                _ = v41
                            case 2147483647:
                                v42 := GenS1{F0: 42, F1: 42, F2: (^(v40))}
                                _ = v42
                            }
                            v14 = 1
                            if ((3.14159 - ((0.0 - 3.14159) - 0.1)) != (1e10 / (0.0 + (1e10 / -2.25)))) {
                                v43 := GenS4{F0: 174, F1: ("hello" + ("hello" + ("foo bar" + "foo bar"))), F2: 18446744073709551615}
                                v15 += 1
                                var v44 uint64 = 42
                                _ = v43
                                _ = v44
                            } else {
                                v45 := GenS2{F0: v14, F1: v15, F2: (fi38 << ((2) & 63))}
                                if true {
                                    fmt.Println(v14, v14)
                                }
                                fmt.Println(v40, v14, fi38)
                                v46 := GenS5{F0: (0.1 >= -2.25), F1: (v40 & (^(v40))), F2: 307}
                                _ = v45
                                _ = v46
                            }
                        }
                    case -9223372036854775808:
                        v40 = (^(v40))
                    default:
                        if (!false) {
                            if true {
                                var v47 float32 = (-2.25 * 1.5)
                                var v48 int16 = (v40 * (v40 - (v40 - v40)))
                                _ = v47
                                _ = v48
                            }
                        } else {
                            fmt.Println(v40)
                            v49 := GenS1{F0: 2, F1: 1, F2: (-(v40))}
                            v50 := v4[4294967295]
                            _ = v49
                            _ = v50
                        }
                    }
                    _ = v40
                }
                _ = len(v6)
                v14 -= (^(v14))
                switch (("" + ("\n\t" + "foo bar")) + ("\n\t" + "hello")) {
                case "foo bar":
                    v51 := GenS5{F0: v8, F1: 1, F2: 255}
                    _ = v51
                case "\n\t":
                    v52 := v6[2147483648]
                    _ = v52
                case "a":
                    v8 = true
                }
                fmt.Println(v8, v8)
                _ = fi38
            }
            _ = fi35
        }
    }
    var v53 uint16 = 255
    switch 128 {
    case 32768:
        v15 = -95
    default:
        if (865 > 376) {
            v54 := GenS4{F0: 128, F1: "a", F2: 274}
            if v8 {
                for fi55 := 0; fi55 < 2; fi55++ {
                    v15 = (v15 & (-(v15)))
                    _ = fi55
                }
                v8 = (1 == 18446744073709551615)
            }
            _ = v54
        } else {
            v15 = -4
        }
    }
    var v56 int = 2147483647
    for fi57 := 0; fi57 < 18; fi57++ {
        var v58 uint32 = 701
        for fi59 := 0; fi59 < 7; fi59++ {
            var v60 uint8 = 150
            switch "hello" {
            case "":
                var v61 uint = 4294967295
                _ = v61
            case "foo bar":
                var v62 uint64 = 440
                _ = v62
            case "hello":
                var v63 uint16 = v14
                _ = v63
            default:
                v58 = (v58 + uint32(v60))
                var v64 uint = 712
                v65 := GenS5{F0: (v15 <= v15), F1: -705, F2: (v64 << (((v64 ^ 439)) & 63))}
                _ = v64
                _ = v65
            }
            var v66 bool = ((v53 * 0) == v14)
            _ = fi59
            _ = v60
            _ = v66
        }
        _ = fi57
        _ = v58
    }
    if (0 < 42) {
        var v67 uint = 42
        _ = v67
    }
    for fi68 := 0; fi68 < 15; fi68++ {
        v1[v56] = ((0 != 253) && false)
        if (255 >= 1) {
            v6[619] = 465
            v8 = (true && v8)
        }
        fmt.Println(v53, v15)
        _ = fi68
    }
    _ = len(v6)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v53:", v53)
    fmt.Println("v56:", v56)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v53
    _ = v56
    _ = v6
    _ = v8
}
