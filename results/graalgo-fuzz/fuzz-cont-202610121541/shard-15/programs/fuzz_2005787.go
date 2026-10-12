package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 bool; F2 uint32 }
type GenS2 struct { F0 uint; F1 uint }
type GenS3 struct { F0 int8; F1 uint16; F2 uint }
type GenS4 struct { F0 uint32; F1 uint64 }
type GenS5 struct { F0 int16; F1 int16 }
func genFn0(p0 uint8, p1 uint16) string {
    for fi6 := 0; fi6 < 5; fi6++ {
        var v7 uint = 18446744073709551615
        v8 := []int{fi6, (fi6 << (((v7 * (v7 & (v7 & v7)))) & 63))}
        var v9 uint = 255
        _ = fi6
        _ = v7
        _ = v8
        _ = v9
    }
    v10 := []int{int(p1)}
    fmt.Println(p1)
    _ = p0
    _ = p1
    _ = v10
    return ((("\n\t" + "") + "foo bar") + "")
}

func main() {
    v1 := make(map[int]bool)
    var v2 float64 = 0.1
    var v3 uint = 4294967295
    v4 := make(map[uint32]uint32)
    v5 := make(map[int]string)
    var v11 string = "hello"
    v3 = v3
    if false {
        v3 = 42
    } else {
        var v12 string = v11
        v13 := GenS1{F0: 4294967295, F1: (false && true), F2: 2}
        _ = v12
        _ = v13
    }
    v2 = -2.25
    var v14 uint32 = 4294967295
    v3 = (v3 + 137)
    v15 := GenS2{F0: v3, F1: (v3 | (v3 % ((v3) | 1)))}
    var v16 uint8 = 66
    var v17 bool = ((!(18446744073709551615 != 20)) && ((v3 & v3) == (v3 | v3)))
    v17 = (((false || false) || (v17 && (v16 == (v16 - (v16 ^ v16))))) || (!v17))
    v17 = (((v3 >> ((v3) & 63)) > v3) && (true && (60 > 137)))
    var v18 uint32 = 170
    fmt.Println(v17)
    if false {
        if false {
            var v19 uint32 = 0
            var v20 uint16 = 0
            _ = v19
            _ = v20
        } else {
            if v17 {
                var v21 float64 = 0.1
                _ = v21
            } else {
                _ = len(v5)
            }
            for fi22 := 0; fi22 < 20; fi22++ {
                var v23 uint = v3
                switch v11 {
                case "\n\t":
                    v24 := GenS2{F0: v23, F1: 322}
                    _ = v24
                case "hello":
                    v25 := GenS2{F0: (v3 | v3), F1: 2}
                    _ = v25
                case "foo bar":
                    fmt.Println(v14)
                default:
                    for fi26 := 0; fi26 < 4; fi26++ {
                        if v17 {
                            if v17 {
                                v11 = v11
                                fmt.Println(v3)
                            }
                            if (0 < 284) {
                                v27 := GenS3{F0: 42, F1: 280, F2: v3}
                                _ = v27
                            }
                            var v28 uint32 = 2147483648
                            _ = v28
                        } else {
                            v29 := v5[(-(fi22))]
                            _ = v29
                        }
                        v14 += v18
                        _ = fi26
                    }
                }
                if v17 {
                    v1[fi22] = v17
                    var v30 int64 = 1
                    _ = v30
                } else {
                    v31 := GenS3{F0: 16, F1: 1, F2: 764}
                    for fi32 := 0; fi32 < 6; fi32++ {
                        v2 += v2
                        _ = fi32
                    }
                    if false {
                        if false {
                            var v33 int16 = 42
                            var v34 int16 = v33
                            fmt.Println(fi22, v17, v23)
                            _ = v33
                            _ = v34
                        } else {
                            delete(v4, 1)
                            v18 = (v14 * v14)
                        }
                        if v17 {
                            v23 = (^(v23))
                        }
                        if (42 < 9223372036854775807) {
                            var v35 uint = (v3 + 48)
                            v4[(v18 & (v18 << ((v23) & 63)))] = (v14 % ((uint32(v16)) | 1))
                            _ = v35
                        } else {
                            if (((!((v16 << ((4294967295) & 63)) != v16)) && (!(v3 >= 310))) || v17) {
                                for fi36 := 0; fi36 < 19; fi36++ {
                                    if (951 < 42) {
                                        if v17 {
                                            var v37 int64 = 42
                                            if (!true) {
                                                v37 += v37
                                                var v38 uint = (^(v3))
                                                _ = v38
                                            } else {
                                                v17 = v17
                                                v3 = v3
                                            }
                                            v39 := GenS2{F0: 2, F1: v3}
                                            _ = v37
                                            _ = v39
                                        } else {
                                            v1[32768] = true
                                            if (((v18 - v14) <= (v18 | 4294967295)) && (("foo bar" + ((v11 + (v11 + "\n\t")) + ((v11 + (v11 + ("hello" + v11))) + v11))) < "hello")) {
                                                var v40 uint64 = 2
                                                delete(v1, fi22)
                                                _ = v40
                                            }
                                            v41 := v5[-2147483648]
                                            v42 := []int{(-(fi22)), 401}
                                            switch fi22 {
                                            case -9223372036854775808:
                                                v43 := GenS2{F0: (v3 & v23), F1: (v23 | v23)}
                                                _ = v43
                                            case 1:
                                                fmt.Println(v16)
                                            default:
                                                if v17 {
                                                    switch v41 {
                                                    case "":
                                                        var v44 int = (fi22 ^ fi22)
                                                        _ = v44
                                                    case "\n\t":
                                                        fmt.Println(fi22)
                                                    default:
                                                        _ = len(v5)
                                                    }
                                                    _ = len(v5)
                                                } else {
                                                    v45 := GenS4{F0: (v18 & uint32(v23)), F1: 0}
                                                    if (((v11 != "") || v17) || v17) {
                                                        v2 = float64(1.5)
                                                    }
                                                    _ = v45
                                                }
                                            }
                                            _ = v41
                                            _ = v42
                                        }
                                        v23 -= v23
                                        fmt.Println(v3, v23, v17)
                                    } else {
                                        if (-9223372036854775808 < 1) {
                                            v46 := v4[2147483647]
                                            fmt.Println(v14)
                                            _ = v46
                                        }
                                        v2 -= v2
                                    }
                                    _ = fi36
                                }
                                var v47 int64 = 42
                                _ = v47
                            } else {
                                if v17 {
                                    var v48 uint32 = v14
                                    _ = v48
                                } else {
                                    for fi49 := 0; fi49 < 12; fi49++ {
                                        var v50 uint32 = v14
                                        var v51 int64 = 42
                                        _ = fi49
                                        _ = v50
                                        _ = v51
                                    }
                                }
                            }
                            fmt.Println(v23, v16)
                        }
                    }
                    _ = v31
                }
                _ = fi22
                _ = v23
            }
        }
    } else {
        var v52 int64 = 9223372036854775807
        _ = len(v5)
        switch ("foo bar" + "a") {
        case "hello":
            v52 = v52
        case "":
            v16 = v16
        case "a":
            v2 -= v2
        default:
            if true {
                var v53 int32 = 1
                switch (v11 + v11) {
                case "foo bar":
                    for fi54 := 0; fi54 < 13; fi54++ {
                        if v17 {
                            v55 := GenS4{F0: (v18 + v14), F1: 550}
                            _ = v55
                        }
                        fmt.Println(v2, v17)
                        _ = fi54
                    }
                }
                if (v17 || (-9223372036854775808 >= 224)) {
                    switch 571 {
                    case -638:
                        _ = len(v5)
                    case 4294967295:
                        v3 = (v3 | (v3 | v3))
                    default:
                        _ = len(v5)
                    }
                    var v56 bool = false
                    _ = v56
                } else {
                    fmt.Println(v11)
                }
                _ = v53
            } else {
                var v57 uint16 = 2
                if (((0.1 - 1.5) <= (1.5 / float32(((v2 * 0.1) - (1.5 + v2))))) || (("" + "foo bar") == v11)) {
                    v16 += (v16 ^ (v16 % (((v16 | v16)) | 1)))
                    var v58 uint64 = 437
                    var v59 string = (("" + v11) + v11)
                    _ = v58
                    _ = v59
                }
                v60 := []int{-2147483648, 4294967295, -30, -944}
                _ = v57
                _ = v60
            }
        }
        _ = v52
    }
    if true {
        v61 := GenS2{F0: (^(v3)), F1: (v3 - (v3 / (((v3 << ((18446744073709551615) & 63))) | 1)))}
        _ = v61
    }
    v62 := GenS2{F0: uint(v18), F1: v3}
    for fi63 := 0; fi63 < 2; fi63++ {
        v64 := GenS5{F0: 0, F1: -830}
        var v65 string = ""
        var v66 uint = 1
        _ = fi63
        _ = v64
        _ = v65
        _ = v66
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v11
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v62
}
