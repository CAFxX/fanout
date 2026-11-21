package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 string }
type GenS2 struct { F0 int64; F1 int8; F2 int8 }
type GenS3 struct { F0 int16; F1 uint64; F2 uint16 }
type GenS4 struct { F0 int64; F1 uint16 }
type GenS5 struct { F0 uint8; F1 uint64 }
type GenS6 struct { F0 int; F1 int8 }
type GenS7 struct { F0 uint32 }
type GenS8 struct { F0 uint; F1 int64; F2 int64 }
func genFn0(p0 uint64, p1 uint) uint8 {
    var v7 uint8 = 1
    v8 := GenS1{F0: p1, F1: "hello"}
    var v9 string = "\n\t"
    fmt.Println(v7, v7, p0)
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    _ = v9
    return v7
}
func genFn1(p0 uint32, p1 uint8) uint8 {
    p1 = genFn0(316, 84)
    v10 := GenS1{F0: 852, F1: "a"}
    v11 := []int{32768, 9223372036854775807, -753}
    p0 -= p0
    p1 = p1
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    return genFn0(4294967295, 745)
}

func main() {
    var v1 uint = 44
    var v2 string = "hello"
    v3 := []string{((((v2 + ("foo bar" + v2)) + "foo bar") + "foo bar") + "\n\t")}
    var v4 int32 = -76
    v5 := make(map[uint32]uint32)
    var v6 uint16 = 266
    fmt.Println(v2)
    for fi12 := 0; fi12 < 3; fi12++ {
        var v13 bool = false
        _ = fi12
        _ = v13
    }
    var v14 int32 = (v4 & 1)
    v14 -= (-(v14))
    var v15 uint = 65535
    var v16 int32 = 1
    v1 -= 65535
    for fi17 := 0; fi17 < 6; fi17++ {
        if ((1e10 * 3.14159) != (3.14159 / 3.14159)) {
            v18 := GenS1{F0: v15, F1: (v2 + (v2 + v2))}
            v5[2] = 740
            _ = v18
        }
        v15 -= v1
        fmt.Println(v1, v2, v15)
        _ = fi17
    }
    var v19 uint32 = 233
    var v20 bool = (128 <= genFn1((v19 ^ (v19 % ((v19) | 1))), 158))
    var v21 uint = (^(v15))
    v22 := GenS2{F0: -1, F1: 115, F2: 18}
    if false {
        _ = len(v3)
        var v23 uint32 = 0
        if false {
            v24 := GenS3{F0: -982, F1: 0, F2: (^(v6))}
            _ = v24
        } else {
            if (!(0 >= int8(v1))) {
                v25 := v5[(v19 & v19)]
                v26 := GenS3{F0: 1, F1: 42, F2: v6}
                v19 += uint32(v15)
                _ = v25
                _ = v26
            }
            v19 -= (v19 >> (((v21 - (v21 | 18446744073709551615))) & 63))
            if false {
                delete(v5, v19)
                var v27 float64 = 1e10
                v28 := GenS3{F0: 0, F1: 42, F2: v6}
                _ = v27
                _ = v28
            } else {
                fmt.Println(v2, v6, v15)
            }
        }
        _ = v23
    } else {
        var v29 float32 = float32((0.0 - 1.5))
        _ = v29
    }
    v6 -= (v6 % ((uint16(v19)) | 1))
    if v20 {
        for fi30 := 0; fi30 < 9; fi30++ {
            for fi31 := 0; fi31 < 18; fi31++ {
                if ((v16 << ((342) & 63)) <= (v16 % (((v16 + (v14 - 641))) | 1))) {
                    switch (fi30 % (((fi30 & 891)) | 1)) {
                    case 0:
                        v14 += -1
                    case 1:
                        var v32 float32 = ((float32((0.0 / (-2.25 + float64((1e10 * ((3.14159 - 3.14159) / (0.1 + -2.25))))))) * (-2.25 / ((((-2.25 - 0.0) - 0.0) - float32(0.1)) - float32(-2.25)))) * ((float32(0.0) / 1.5) + 0.0))
                        _ = v32
                    }
                } else {
                    v15 += 42
                    v1 -= v15
                    v33 := v5[(v19 << ((2) & 63))]
                    if true {
                        if true {
                            v34 := v5[v33]
                            var v35 int16 = 0
                            _ = len(v5)
                            _ = v34
                            _ = v35
                        } else {
                            var v36 uint16 = (v6 >> ((v15) & 63))
                            var v37 int8 = 0
                            _ = v36
                            _ = v37
                        }
                        if v20 {
                            v38 := GenS4{F0: 9223372036854775807, F1: 32767}
                            switch -741 {
                            case 502:
                                v15 = (v1 + 2)
                            }
                            if v20 {
                                switch (v2 + v2) {
                                case "hello":
                                    switch v2 {
                                    case "foo bar":
                                        v39 := GenS2{F0: 736, F1: -128, F2: -7}
                                        _ = v39
                                    case "hello":
                                        fmt.Println(fi30, fi30, v19)
                                    case "\n\t":
                                        v6 = v6
                                    }
                                    v20 = ((v2 > "foo bar") || v20)
                                }
                            }
                            var v40 int32 = 1
                            var v41 int16 = 42
                            _ = v38
                            _ = v40
                            _ = v41
                        }
                        v42 := GenS5{F0: genFn0(uint64(v19), v21), F1: 42}
                        _ = v42
                    } else {
                        if len(v3) > 0 {
                            v3[0] = ((v2 + v2) + v2)
                        }
                        if v20 {
                            v19 -= (v19 - (^(v33)))
                        } else {
                            _ = len(v3)
                        }
                    }
                    _ = v33
                }
                var v43 int32 = v14
                _ = fi31
                _ = v43
            }
            var v44 int32 = (v14 & 42)
            var v45 uint8 = uint8(v19)
            _ = fi30
            _ = v44
            _ = v45
        }
        for fi46 := 0; fi46 < 7; fi46++ {
            v47 := GenS3{F0: 1, F1: 4294967295, F2: 255}
            switch (fi46 + 32768) {
            case 1:
                var v48 string = v2
                _ = v48
            case 2147483647:
                v3 = append(v3, "\n\t")
            case 4294967295:
                v15 += v15
            default:
                v1 -= v15
            }
            switch v2 {
            case "\n\t":
                var v49 uint64 = 769
                _ = v49
            default:
                switch v2 {
                case "":
                    var v50 uint16 = v6
                    _ = v50
                case "\n\t":
                    var v51 int = fi46
                    _ = v51
                default:
                    delete(v5, 569)
                }
            }
            _ = fi46
            _ = v47
        }
    } else {
        v52 := GenS1{F0: (^(v1)), F1: v2}
        for fi53 := 0; fi53 < 2; fi53++ {
            switch ("hello" + (v2 + ("\n\t" + v2))) {
            case "hello":
                switch 32767 {
                case -333:
                    fmt.Println(v19, v20)
                case 127:
                    if v20 {
                        for fi54 := 0; fi54 < 15; fi54++ {
                            if len(v3) > 0 {
                                v3[0] = (v2 + "\n\t")
                            }
                            v55 := GenS6{F0: (fi54 * fi53), F1: 100}
                            _ = fi54
                            _ = v55
                        }
                    }
                case -1:
                    v56 := GenS5{F0: 127, F1: 2}
                    _ = v56
                default:
                    if (((v16 * (v16 % ((v4) | 1))) != v16) && true) {
                        for fi57 := 0; fi57 < 12; fi57++ {
                            var v58 bool = v20
                            var v59 int8 = 0
                            switch ("" + ("a" + "hello")) {
                            case "a":
                                switch (v2 + "foo bar") {
                                case "\n\t":
                                    var v60 int16 = 32767
                                    _ = v60
                                case "":
                                    v20 = ((v6 ^ v6) > 42)
                                case "hello":
                                    var v61 int64 = 155
                                    _ = v61
                                default:
                                    var v62 int16 = -217
                                    _ = v62
                                }
                                v63 := GenS7{F0: 1}
                                _ = len(v5)
                                _ = v63
                            case "":
                                v64 := GenS4{F0: 148, F1: v6}
                                if len(v3) > 0 {
                                    v3[0] = ((v2 + "") + v2)
                                }
                                if ((v6 - 255) > v6) {
                                    v3 = append(v3, "a")
                                    switch (fi57 >> ((v15) & 63)) {
                                    case 4294967295:
                                        if v20 {
                                            v3 = append(v3, "")
                                            v58 = ((((1e10 + 3.14159) - (-2.25 + 1.5)) + 1.5) == 1.5)
                                        }
                                    case 327:
                                        if false {
                                            if false {
                                                switch -1 {
                                                case 32767:
                                                    v65 := GenS5{F0: 1, F1: 4294967295}
                                                    _ = v65
                                                case 4294967295:
                                                    fmt.Println(v1, v1)
                                                case -2147483648:
                                                    if (0 != 324) {
                                                        var v66 uint32 = v19
                                                        var v67 uint64 = 4294967295
                                                        _ = v66
                                                        _ = v67
                                                    } else {
                                                        var v68 uint32 = v19
                                                        v19 -= v68
                                                        _ = v68
                                                    }
                                                }
                                                fmt.Println(v20, v4)
                                            }
                                            delete(v5, 65535)
                                            v4 = (v16 | v16)
                                            if (v58 && v58) {
                                                v15 = v15
                                                var v69 float32 = (1.5 * 0.0)
                                                _ = v69
                                            }
                                            v21 -= 494
                                        }
                                    default:
                                        var v70 uint8 = 128
                                        fmt.Println(v14, v16, v59)
                                        v6 = v6
                                        _ = v70
                                    }
                                }
                                _ = v64
                            case "foo bar":
                                var v71 int32 = v16
                                _ = v71
                            default:
                                switch (fi57 & (fi57 % ((int(v59)) | 1))) {
                                case 779:
                                    v14 -= 42
                                    fmt.Println(v2, v4, v1)
                                    fmt.Println(v14, fi53)
                                }
                            }
                            _ = fi57
                            _ = v58
                            _ = v59
                        }
                    }
                }
            case "a":
                v72 := GenS8{F0: 0, F1: 0, F2: 0}
                _ = v72
            case "":
                switch (fi53 ^ fi53) {
                case 9223372036854775807:
                    v73 := GenS7{F0: v19}
                    _ = v73
                case -198:
                    var v74 int8 = 127
                    _ = v74
                case -665:
                    var v75 int16 = 898
                    _ = v75
                }
            }
            _ = fi53
        }
        _ = v52
    }
    fmt.Println(v4)
    for fi76 := 0; fi76 < 15; fi76++ {
        var v77 uint64 = 18446744073709551615
        _ = fi76
        _ = v77
    }
    if v20 {
        var v78 uint = v21
        for fi79 := 0; fi79 < 20; fi79++ {
            for fi80 := 0; fi80 < 19; fi80++ {
                var v81 float64 = 3.14159
                fmt.Println(v21, fi80, v19)
                _ = fi80
                _ = v81
            }
            _ = fi79
        }
        _ = v78
    } else {
        v14 = v16
        if false {
            for fi82 := 0; fi82 < 6; fi82++ {
                for fi83 := 0; fi83 < 13; fi83++ {
                    v6 += v6
                    _ = fi83
                }
                _ = fi82
            }
            if v20 {
                v15 -= (^(v21))
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
