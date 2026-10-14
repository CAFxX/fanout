package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int16; F1 bool; F2 uint32 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint8; F1 uint8 }
type GenS5 struct { F0 bool; F1 uint64 }
type GenS6 struct { F0 uint32; F1 uint8 }
type GenS7 struct { F0 uint }
func genFn0(p0 uint8) uint8 {
    v6 := GenS1{F0: "foo bar"}
    var v7 int16 = -32768
    v8 := []int{0, -944}
    for fi9 := 0; fi9 < 1; fi9++ {
        if len(v8) > 0 {
            v8[0] = fi9
        }
        _ = fi9
    }
    if len(v8) >= 2 {
        v8 = v8[1:len(v8)]
    }
    _ = p0
    _ = v6
    _ = v7
    _ = v8
    return 202
}
func genFn1() int32 {
    v10 := []int{-902, -904, 132}
    var v11 int16 = -32768
    if false {
        var v12 int8 = 61
        v11 = (v11 % ((v11) | 1))
        _ = v12
    } else {
        for fi13 := 0; fi13 < 1; fi13++ {
            v10 = append(v10, -184)
            _ = fi13
        }
        fmt.Println(v11)
        v11 = 881
    }
    _ = v10
    _ = v11
    return 599
}

func main() {
    v1 := make(map[int]string)
    var v2 uint32 = 358
    var v3 bool = false
    var v4 int8 = -1
    var v5 uint16 = uint16(v2)
    if ((("a" <= "\n\t") || v3) || false) {
        var v14 string = ("a" + "")
        _ = v14
    }
    v15 := GenS1{F0: ("hello" + ("" + "hello"))}
    for fi16 := 0; fi16 < 7; fi16++ {
        fmt.Println(v3)
        fmt.Println(v3, v4)
        _ = fi16
    }
    if (("" >= "a") || v3) {
        var v17 float32 = 1.5
        v4 += -122
        _ = v17
    }
    v18 := []int{-9223372036854775808, 1}
    for fi19 := 0; fi19 < 12; fi19++ {
        v5 -= v5
        for fi20 := 0; fi20 < 7; fi20++ {
            var v21 string = "a"
            _ = fi20
            _ = v21
        }
        _ = fi19
    }
    if (251 < v5) {
        if len(v18) > 0 {
            v18[0] = 4294967295
        }
        delete(v1, 1)
        for fi22 := 0; fi22 < 9; fi22++ {
            v23 := v1[(fi22 - fi22)]
            var v24 float32 = (0.1 * 0.1)
            switch int(v2) {
            case 9223372036854775807:
                if len(v18) >= 2 {
                    v18 = v18[1:len(v18)]
                }
            case -9223372036854775808:
                if v3 {
                    v18 = append(v18, -9223372036854775808)
                    v2 = v2
                } else {
                    if ((v3 || v3) || (((v2 & (v2 / ((v2) | 1))) < (v2 | (^(v2)))) && (!v3))) {
                        if (0.0 <= v24) {
                            if (2147483647 < 0) {
                                var v25 bool = (2 != 128)
                                _ = v25
                            }
                        } else {
                            var v26 int64 = -9223372036854775808
                            if len(v18) > 0 {
                                v18[0] = (fi22 << ((643) & 63))
                            }
                            _ = v26
                        }
                        for fi27 := 0; fi27 < 5; fi27++ {
                            _ = len(v1)
                            var v28 int32 = int32(v5)
                            v29 := GenS1{F0: v23}
                            _ = fi27
                            _ = v28
                            _ = v29
                        }
                    }
                    v24 -= (0.0 / float32(float64((v24 * 0.1))))
                }
            default:
                v30 := GenS1{F0: ("foo bar" + ("hello" + ("\n\t" + v23)))}
                _ = v30
            }
            _ = fi22
            _ = v23
            _ = v24
        }
    }
    var v31 uint8 = 103
    v5 += (v5 - 188)
    v32 := GenS1{F0: ""}
    switch 4294967295 {
    case 32767:
        fmt.Println(v2)
    default:
        v4 = 94
    }
    fmt.Println(v5, v2, v2)
    v31 = v31
    if v3 {
        if (1000 == 681) {
            var v33 uint32 = v2
            if ("\n\t" < "foo bar") {
                v4 -= 127
            }
            _ = v33
        } else {
            if (2 == 669) {
                if v3 {
                    if true {
                        var v34 uint64 = 42
                        if ((1e10 * (0.0 / 3.14159)) < (float64((3.14159 * (1.5 / (1e10 * 1e10)))) + 3.14159)) {
                            if (v3 && true) {
                                fmt.Println(v2, v4)
                                v4 = -20
                            } else {
                                if v3 {
                                    v4 = (v4 ^ (v4 >> ((4294967295) & 63)))
                                }
                                for fi35 := 0; fi35 < 13; fi35++ {
                                    var v36 int64 = -656
                                    v5 -= 255
                                    v5 = v5
                                    _ = fi35
                                    _ = v36
                                }
                            }
                        } else {
                            var v37 uint8 = (^(v31))
                            var v38 int16 = 735
                            if (genFn0(v31) >= 195) {
                                switch -972 {
                                case 4294967295:
                                    v4 -= 100
                                }
                                var v39 bool = (1.5 < (0.0 + (((-2.25 * float32(3.14159)) * 1e10) + 1.5)))
                                _ = v39
                            } else {
                                if (735 > 1) {
                                    for fi40 := 0; fi40 < 5; fi40++ {
                                        v41 := GenS1{F0: "foo bar"}
                                        _ = fi40
                                        _ = v41
                                    }
                                    var v42 float64 = (1.5 + 1e10)
                                    fmt.Println(v42, v34)
                                    _ = v42
                                } else {
                                    var v43 uint16 = v5
                                    _ = v43
                                }
                                if v3 {
                                    fmt.Println(v4, v5)
                                }
                                _ = len(v1)
                            }
                            _ = v37
                            _ = v38
                        }
                        var v44 bool = true
                        _ = v34
                        _ = v44
                    } else {
                        if false {
                            var v45 int16 = 1000
                            _ = v45
                        } else {
                            if v3 {
                                for fi46 := 0; fi46 < 3; fi46++ {
                                    _ = len(v1)
                                    _ = fi46
                                }
                                var v47 float64 = (0.1 - (0.1 / 3.14159))
                                if v3 {
                                    v18 = append(v18, 0)
                                    var v48 int = 2147483647
                                    _ = v48
                                } else {
                                    if v3 {
                                        if v3 {
                                            if len(v18) >= 2 {
                                                v18 = v18[1:len(v18)]
                                            }
                                            var v49 int32 = -883
                                            _ = v49
                                        } else {
                                            _ = len(v1)
                                            switch "hello" {
                                            case "hello":
                                                v50 := GenS2{F0: 214, F1: (!v3), F2: 636}
                                                _ = v50
                                            }
                                            if v3 {
                                                var v51 int16 = -294
                                                var v52 uint32 = (v2 / ((v2) | 1))
                                                v53 := v1[127]
                                                _ = v51
                                                _ = v52
                                                _ = v53
                                            }
                                        }
                                        var v54 float64 = 0.0
                                        var v55 int16 = 32767
                                        _ = v54
                                        _ = v55
                                    } else {
                                        var v56 uint = 42
                                        var v57 string = ("a" + "\n\t")
                                        v1[-2147483648] = "\n\t"
                                        _ = v56
                                        _ = v57
                                    }
                                    if ((v2 << ((255) & 63)) != v2) {
                                        if v3 {
                                            v58 := GenS3{F0: 911}
                                            var v59 int32 = 622
                                            _ = v58
                                            _ = v59
                                        }
                                    }
                                }
                                _ = v47
                            } else {
                                v4 += v4
                                _ = len(v1)
                            }
                        }
                        if v3 {
                            if v3 {
                                var v60 string = "foo bar"
                                _ = v60
                            } else {
                                if true {
                                    v61 := GenS4{F0: genFn0((v31 % (((v31 & v31)) | 1))), F1: 1}
                                    v31 += (v31 - uint8(v5))
                                    _ = v61
                                }
                                fmt.Println(v31, v31, v5)
                            }
                        } else {
                            if v3 {
                                if v3 {
                                    fmt.Println(v5, v2)
                                    for fi62 := 0; fi62 < 6; fi62++ {
                                        var v63 uint32 = uint32(v31)
                                        v3 = (0 <= genFn1())
                                        switch "foo bar" {
                                        case "foo bar":
                                            fmt.Println(v3, v3, v4)
                                        case "":
                                            v64 := GenS5{F0: v3, F1: 779}
                                            _ = v64
                                        default:
                                            for fi65 := 0; fi65 < 10; fi65++ {
                                                v66 := GenS1{F0: ("\n\t" + "a")}
                                                var v67 bool = true
                                                v67 = (true || ((genFn1() == genFn1()) && (v3 || (genFn1() != 0))))
                                                if v3 {
                                                    if v3 {
                                                        var v68 bool = (!(217 == fi62))
                                                        var v69 int = int(v5)
                                                        if v67 {
                                                            if ((fi62 >> ((18446744073709551615) & 63)) <= (v69 >> ((0) & 63))) {
                                                                if (2 != 522) {
                                                                    switch (("" + ((("a" + "foo bar") + (("\n\t" + "foo bar") + "a")) + ("\n\t" + ("\n\t" + "hello")))) + "\n\t") {
                                                                    case "":
                                                                        v4 = (^(v4))
                                                                    default:
                                                                        v70 := GenS6{F0: (v2 ^ v63), F1: v31}
                                                                        _ = v70
                                                                    }
                                                                }
                                                                v71 := GenS7{F0: 2}
                                                                _ = v71
                                                            } else {
                                                                fmt.Println(v69, v67)
                                                                v2 -= 1
                                                            }
                                                        }
                                                        _ = v68
                                                        _ = v69
                                                    } else {
                                                        if v3 {
                                                            v2 -= v63
                                                            var v72 float64 = 0.0
                                                            v31 = (v31 % ((v31) | 1))
                                                            _ = v72
                                                        }
                                                        v2 += (v2 * v2)
                                                    }
                                                    if v3 {
                                                        v73 := GenS2{F0: 1000, F1: v3, F2: v2}
                                                        _ = v73
                                                    }
                                                    var v74 bool = (9 != 18446744073709551615)
                                                    _ = v74
                                                }
                                                _ = fi65
                                                _ = v66
                                                _ = v67
                                            }
                                        }
                                        _ = fi62
                                        _ = v63
                                    }
                                    var v75 float64 = 1e10
                                    _ = v75
                                }
                                if v3 {
                                    v76 := GenS3{F0: 2}
                                    _ = v76
                                }
                            }
                        }
                    }
                }
                fmt.Println(v31, v5)
            }
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v18
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v5
}
