package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 string; F2 uint32 }
type GenS2 struct { F0 int16; F1 int64; F2 uint }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 int8; F1 bool }
type GenS6 struct { F0 int32; F1 int32; F2 int8 }
type GenS7 struct { F0 uint64; F1 uint8 }
type GenS8 struct { F0 int64; F1 uint }
type GenS9 struct { F0 uint }

func main() {
    var v1 bool = (-387 == 9223372036854775807)
    var v2 uint = 1
    v3 := make(map[int]string)
    var v4 float32 = (1e10 / 1.0)
    v5 := make(map[uint32]uint32)
    v6 := []int{128, 2147483647, 130}
    _ = len(v6)
    v7 := GenS1{F0: "", F1: "a", F2: 2147483648}
    fmt.Println(v2, v1)
    v8 := GenS1{F0: "\n\t", F1: "hello", F2: 1}
    v6 = append(v6, -621)
    v2 = (^(v2))
    v5[0] = 2
    var v9 string = "foo bar"
    for fi10 := 0; fi10 < 17; fi10++ {
        var v11 int8 = 100
        if (float32(1e10) == v4) {
            switch v9 {
            case "hello":
                var v12 uint16 = 2
                _ = v12
            case "foo bar":
                var v13 int16 = -402
                _ = v13
            case "":
                v14 := GenS2{F0: -1, F1: -80, F2: v2}
                _ = v14
            default:
                if v1 {
                    v15 := GenS3{F0: 2147483647}
                    fmt.Println(v1, v4, v9)
                    if true {
                        v16 := GenS2{F0: 1000, F1: -1, F2: v2}
                        if v1 {
                            for fi17 := 0; fi17 < 10; fi17++ {
                                switch fi10 {
                                case -1:
                                    var v18 uint16 = 1
                                    _ = v18
                                case 9223372036854775807:
                                    switch v9 {
                                    case "foo bar":
                                        v2 += (v2 >> (((v2 % ((v2) | 1))) & 63))
                                    case "":
                                        v1 = v1
                                    case "a":
                                        var v19 float64 = (3.14159 / 3.14159)
                                        _ = v19
                                    default:
                                        var v20 float64 = 0.1
                                        _ = v20
                                    }
                                }
                                _ = len(v6)
                                v3[(fi10 ^ (fi17 >> (((v2 % ((4294967295) | 1))) & 63)))] = v9
                                _ = fi17
                            }
                        } else {
                            v21 := GenS1{F0: ("foo bar" + v9), F1: ((v9 + (v9 + v9)) + (("a" + ("foo bar" + v9)) + "\n\t")), F2: 2}
                            _ = v21
                        }
                        if (v1 && false) {
                            v22 := GenS2{F0: -1, F1: 1, F2: v2}
                            _ = v22
                        }
                        _ = v16
                    } else {
                        if false {
                            v6 = append(v6, fi10)
                            for fi23 := 0; fi23 < 10; fi23++ {
                                v9 = v9
                                _ = fi23
                            }
                        } else {
                            v24 := GenS4{F0: 196}
                            for fi25 := 0; fi25 < 20; fi25++ {
                                switch (fi25 >> ((708) & 63)) {
                                case -438:
                                    if len(v6) > 0 {
                                        v6[0] = fi25
                                    }
                                case 73:
                                    var v26 int = (fi25 ^ -2147483648)
                                    _ = v26
                                case 9223372036854775807:
                                    if false {
                                        v2 -= (^(v2))
                                    }
                                default:
                                    var v27 uint16 = 0
                                    _ = v27
                                }
                                v28 := GenS5{F0: (v11 >> ((v2) & 63)), F1: v1}
                                _ = fi25
                                _ = v28
                            }
                            v29 := GenS1{F0: (v9 + "a"), F1: v9, F2: 2147483648}
                            _ = v24
                            _ = v29
                        }
                        var v30 uint8 = 127
                        v5[626] = 2147483648
                        _ = v30
                    }
                    _ = v15
                } else {
                    if ((0.0 * 1e10) == -2.25) {
                        var v31 uint32 = 324
                        fmt.Println(v2)
                        if v1 {
                            v6 = append(v6, 2147483647)
                            v32 := v5[2]
                            v33 := GenS4{F0: -436}
                            _ = v32
                            _ = v33
                        }
                        _ = v31
                    } else {
                        v34 := GenS4{F0: -9223372036854775808}
                        v4 = 1.5
                        _ = v34
                    }
                    v1 = true
                }
            }
            switch fi10 {
            case 752:
                if (((((float64(0.0) + (1.5 / (float64(((v4 * 3.14159) + (v4 * (((v4 + float32(0.1)) / (float32((0.0 + 0.0)) / (v4 - float32(float64(((((0.1 + v4) - v4) / (-2.25 * v4)) - float32((0.1 / 1.5)))))))) - 0.1)))) * 1e10))) * 1.5) * -2.25) - float64(v4)) >= ((0.1 - 0.1) / float64(v4))) {
                    _ = len(v6)
                    fmt.Println(v1, fi10)
                    var v35 int32 = 0
                    _ = v35
                } else {
                    v36 := GenS2{F0: 32767, F1: -927, F2: (v2 - v2)}
                    for fi37 := 0; fi37 < 16; fi37++ {
                        v11 -= (v11 / (((^(v11))) | 1))
                        if (true && (!(v2 < (v2 ^ (v2 & v2))))) {
                            v38 := GenS5{F0: v11, F1: v1}
                            _ = v38
                        } else {
                            if (!(420 != 65535)) {
                                if v1 {
                                    v9 = (v9 + "hello")
                                    var v39 uint8 = 1
                                    _ = v39
                                } else {
                                    fmt.Println(v1, v9)
                                }
                                if (!(2 >= 18446744073709551615)) {
                                    _ = len(v3)
                                    v3[fi37] = (v9 + v9)
                                }
                                switch v9 {
                                case "foo bar":
                                    v40 := GenS1{F0: v9, F1: ((v9 + v9) + (v9 + "a")), F2: 2147483648}
                                    if len(v6) >= 2 {
                                        v6 = v6[1:len(v6)]
                                    }
                                    switch v9 {
                                    case "foo bar":
                                        delete(v5, 65535)
                                    case "a":
                                        v41 := GenS1{F0: "a", F1: v9, F2: 242}
                                        _ = v41
                                    case "":
                                        var v42 float32 = v4
                                        var v43 float64 = 3.14159
                                        var v44 uint16 = 2
                                        switch fi10 {
                                        case -9223372036854775808:
                                            fmt.Println(v4, v43, fi10)
                                        }
                                        v45 := GenS6{F0: -2147483648, F1: -216, F2: 1}
                                        _ = v42
                                        _ = v43
                                        _ = v44
                                        _ = v45
                                    default:
                                        if v1 {
                                            if (4294967295 <= 1) {
                                                _ = len(v6)
                                                v11 = 127
                                            }
                                        } else {
                                            var v46 int32 = 906
                                            fmt.Println(v1, fi37)
                                            _ = v46
                                        }
                                    }
                                    v1 = v1
                                    if (v1 || false) {
                                        v47 := GenS1{F0: (v9 + v9), F1: ((((v9 + "\n\t") + v9) + v9) + "a"), F2: 733}
                                        switch v9 {
                                        case "hello":
                                            if v1 {
                                                v48 := GenS1{F0: (v9 + v9), F1: "hello", F2: 4294967295}
                                                if v1 {
                                                    v6 = append(v6, (fi10 - (fi10 / ((fi37) | 1))))
                                                    v9 = ("a" + v9)
                                                    switch fi10 {
                                                    case 2147483647:
                                                        v5[156] = 2147483647
                                                    default:
                                                        if ("foo bar" < "") {
                                                            var v49 int16 = -32768
                                                            _ = v49
                                                        }
                                                    }
                                                }
                                                _ = v48
                                            } else {
                                                if true {
                                                    if (1 <= 1) {
                                                        var v50 uint16 = 1
                                                        var v51 int8 = -103
                                                        var v52 string = ""
                                                        var v53 uint64 = 4294967295
                                                        _ = v50
                                                        _ = v51
                                                        _ = v52
                                                        _ = v53
                                                    } else {
                                                        var v54 uint64 = 1
                                                        if ((true || v1) || (!v1)) {
                                                            fmt.Println(v9, v54)
                                                            if v1 {
                                                                v55 := GenS1{F0: v9, F1: v9, F2: 716}
                                                                var v56 uint64 = (^(v54))
                                                                _ = v55
                                                                _ = v56
                                                            }
                                                        }
                                                        _ = v54
                                                    }
                                                } else {
                                                    var v57 uint64 = 1
                                                    if v1 {
                                                        fmt.Println(v57, v9)
                                                        if v1 {
                                                            fmt.Println(fi10, v9, fi10)
                                                            fmt.Println(fi10, v4, v57)
                                                            fmt.Println(v11)
                                                        }
                                                        _ = len(v3)
                                                    }
                                                    var v58 uint8 = 127
                                                    _ = v57
                                                    _ = v58
                                                }
                                            }
                                        }
                                        _ = v47
                                    } else {
                                        if true {
                                            switch (v9 + "") {
                                            case "":
                                                var v59 string = v9
                                                _ = v59
                                            case "a":
                                                v1 = true
                                            case "foo bar":
                                                var v60 string = ""
                                                var v61 int64 = 9223372036854775807
                                                _ = v60
                                                _ = v61
                                            }
                                            var v62 float64 = (float64(v4) / (0.1 * 0.1))
                                            var v63 int16 = -814
                                            _ = v62
                                            _ = v63
                                        }
                                        v64 := GenS5{F0: (^(v11)), F1: v1}
                                        var v65 int32 = -2147483648
                                        _ = v64
                                        _ = v65
                                    }
                                    _ = v40
                                default:
                                    if v1 {
                                        if (2 != 0) {
                                            fmt.Println(v11)
                                            fmt.Println(v4, fi37)
                                        }
                                        v1 = ((fi10 | (fi37 | fi10)) <= (^(fi10)))
                                        v1 = ((v1 && v1) && (!v1))
                                    } else {
                                        v66 := GenS3{F0: -2147483648}
                                        _ = v66
                                    }
                                }
                            }
                            var v67 uint64 = 2
                            _ = v67
                        }
                        _ = fi37
                    }
                    _ = v36
                }
            case 0:
                fmt.Println(v11, v11, v11)
            default:
                fmt.Println(v9)
            }
            var v68 int = fi10
            _ = v68
        } else {
            for fi69 := 0; fi69 < 3; fi69++ {
                if ((^(v11)) >= -126) {
                    if v1 {
                        if v1 {
                            v70 := GenS7{F0: 0, F1: 115}
                            var v71 float64 = 0.1
                            v72 := GenS3{F0: 1}
                            _ = v70
                            _ = v71
                            _ = v72
                        } else {
                            if v1 {
                                var v73 int32 = 1
                                var v74 uint16 = 255
                                _ = v73
                                _ = v74
                            } else {
                                if v1 {
                                    v75 := v3[600]
                                    _ = v75
                                }
                            }
                            v6 = append(v6, (fi10 ^ fi69))
                            v76 := GenS8{F0: int64(v2), F1: (v2 + (v2 % (((v2 + (^(v2)))) | 1)))}
                            fmt.Println(v2, v11)
                            _ = v76
                        }
                    }
                    var v77 uint = v2
                    var v78 uint = (v2 & (^(v77)))
                    _ = v77
                    _ = v78
                }
                _ = fi69
            }
            v1 = ("foo bar" < (v9 + v9))
            switch fi10 {
            case 2147483647:
                var v79 float64 = (-2.25 / float64(1e10))
                _ = v79
            case 32768:
                v3[(^(fi10))] = (v9 + (((v9 + v9) + v9) + (v9 + v9)))
            default:
                var v80 uint8 = 255
                _ = v80
            }
        }
        for fi81 := 0; fi81 < 12; fi81++ {
            v82 := GenS9{F0: v2}
            if v1 {
                v9 = "foo bar"
                if v1 {
                    if v1 {
                        if (2147483647 == 42) {
                            var v83 bool = v1
                            _ = v83
                        } else {
                            v84 := GenS6{F0: 757, F1: -1, F2: (v11 + v11)}
                            var v85 uint = 863
                            _ = v84
                            _ = v85
                        }
                    }
                }
            }
            _ = fi81
            _ = v82
        }
        _ = fi10
        _ = v11
    }
    var v86 int32 = -900
    v6 = append(v6, -1)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v86:", v86)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v86
    _ = v9
}
