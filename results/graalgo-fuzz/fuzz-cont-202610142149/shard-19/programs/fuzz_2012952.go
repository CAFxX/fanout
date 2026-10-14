package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32; F2 uint32 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int64; F1 uint8; F2 int32 }
type GenS5 struct { F0 uint16 }
func genFn0(p0 uint16, p1 uint8, p2 int) uint {
    v7 := GenS1{F0: 65535, F1: 2147483647, F2: 2147483648}
    v8 := GenS1{F0: 0, F1: 2147483647, F2: 479}
    var v9 int = p2
    var v10 int8 = 0
    if ((v9 >> ((65535) & 63)) != 0) {
        if false {
            var v11 uint8 = (^(p1))
            fmt.Println(p0, p1)
            v12 := GenS1{F0: 4294967295, F1: 42, F2: 65535}
            _ = v11
            _ = v12
        }
        fmt.Println(p1, v9)
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return 292
}
func genFn1(p0 int64) bool {
    v13 := []int{}
    var v14 int8 = -69
    if len(v13) > 0 {
        v13[0] = -765
    }
    _ = p0
    _ = v13
    _ = v14
    return false
}

func main() {
    v1 := make(map[string]string)
    var v2 uint32 = 889
    var v3 uint8 = 219
    var v4 uint64 = 194
    v5 := make(map[string]bool)
    v6 := []uint32{(^(v2)), (v2 % (((^(v2))) | 1)), (^(v2))}
    switch "\n\t" {
    case "\n\t":
        _ = len(v5)
    case "a":
        for fi15 := 0; fi15 < 19; fi15++ {
            var v16 float32 = (0.1 * float32(1e10))
            _ = fi15
            _ = v16
        }
    case "hello":
        var v17 uint16 = 1
        _ = v17
    default:
        var v18 int32 = 655
        _ = v18
    }
    var v19 uint16 = 1
    v19 -= v19
    for fi20 := 0; fi20 < 15; fi20++ {
        v3 += 42
        v4 = v4
        _ = fi20
    }
    v2 -= (v2 | v2)
    var v21 uint = 42
    for fi22 := 0; fi22 < 11; fi22++ {
        v4 = 821
        v23 := GenS1{F0: (v21 % (((v21 << ((339) & 63))) | 1)), F1: -1, F2: (v2 << ((genFn0(v19, v3, (fi22 / (((fi22 << ((genFn0(65535, (^(v3)), (fi22 & fi22))) & 63))) | 1)))) & 63))}
        if (!false) {
            v6 = append(v6, (v2 & v2))
            if false {
                switch fi22 {
                case 128:
                    var v24 bool = ((v19 % ((v19) | 1)) < v19)
                    _ = v24
                case 292:
                    for fi25 := 0; fi25 < 15; fi25++ {
                        var v26 bool = (!true)
                        v27 := GenS2{F0: 1000}
                        v26 = (genFn0(uint16(v21), uint8(v19), (fi25 + fi22)) == v21)
                        if true {
                            v2 = v2
                            switch (fi22 ^ (fi22 << (((v21 * uint(v3))) & 63))) {
                            case -72:
                                var v28 float32 = 0.1
                                _ = v28
                            default:
                                fmt.Println(v19, fi22)
                            }
                            var v29 int64 = -398
                            _ = v29
                        }
                        if len(v6) > 0 {
                            v6[0] = (v2 % ((v2) | 1))
                        }
                        switch "foo bar" {
                        case "hello":
                            var v30 int8 = -63
                            _ = v30
                        case "a":
                            var v31 float64 = 0.0
                            _ = v31
                        default:
                            var v32 float64 = 1.5
                            v26 = true
                            v33 := GenS1{F0: 255, F1: -2147483648, F2: (v2 ^ v2)}
                            _ = v32
                            _ = v33
                        }
                        _ = fi25
                        _ = v26
                        _ = v27
                    }
                case 42:
                    var v34 uint32 = 2147483647
                    _ = v34
                }
            }
            if (((42 != 128) || true) && false) {
                v3 += v3
                v21 = 443
            }
        } else {
            v5["a"] = (v3 >= v3)
            for fi35 := 0; fi35 < 4; fi35++ {
                v36 := GenS1{F0: (^(v21)), F1: 2147483647, F2: 0}
                if false {
                    if ((31 < 100) || false) {
                        if false {
                            if (1e10 >= (1.5 * 0.0)) {
                                v37 := GenS3{F0: 675}
                                var v38 uint = (v21 >> ((v21) & 63))
                                fmt.Println(v38)
                                _ = v37
                                _ = v38
                            }
                        }
                        delete(v5, ((("hello" + "a") + ("" + ("hello" + (("a" + "") + ("foo bar" + "a"))))) + ("\n\t" + (("foo bar" + ("a" + ("foo bar" + "a"))) + ("foo bar" + (("a" + "") + "a"))))))
                    }
                } else {
                    var v39 string = ("a" + (("hello" + "") + "hello"))
                    v2 -= v2
                    _ = v39
                }
                if (v21 == (^(v21))) {
                    v4 -= uint64(v21)
                }
                _ = fi35
                _ = v36
            }
            if ((uint16(v21) == (^(v19))) || (!("hello" <= ("" + "\n\t")))) {
                for fi40 := 0; fi40 < 10; fi40++ {
                    v41 := v1[("a" + ("a" + ""))]
                    var v42 float64 = (0.0 / (1e10 / 1.5))
                    v41 = ("\n\t" + v41)
                    v5[(("hello" + (v41 + (v41 + "foo bar"))) + v41)] = (!true)
                    _ = fi40
                    _ = v41
                    _ = v42
                }
            } else {
                v3 -= v3
                switch fi22 {
                case 32767:
                    if true {
                        v43 := GenS1{F0: (v21 + (v21 + v21)), F1: 42, F2: (^(v2))}
                        v1["foo bar"] = ("foo bar" + "foo bar")
                        _ = v43
                    } else {
                        var v44 string = ""
                        if true {
                            v45 := GenS1{F0: genFn0(1, (v3 * v3), (fi22 | -336)), F1: -1, F2: (^(v2))}
                            _ = v45
                        }
                        if true {
                            for fi46 := 0; fi46 < 20; fi46++ {
                                v4 = 2
                                var v47 int = (-(fi22))
                                v2 = (v2 & v2)
                                var v48 int = (^(fi46))
                                _ = fi46
                                _ = v47
                                _ = v48
                            }
                            fmt.Println(v4, v21, fi22)
                        } else {
                            switch (fi22 / (((fi22 - fi22)) | 1)) {
                            case -1:
                                switch fi22 {
                                case -9223372036854775808:
                                    fmt.Println(v21, fi22, v44)
                                }
                            default:
                                fmt.Println(v44)
                            }
                            if true {
                                var v49 float32 = 0.0
                                _ = v49
                            } else {
                                fmt.Println(v3, v2)
                                if false {
                                    if false {
                                        v50 := GenS4{F0: -393, F1: v3, F2: -624}
                                        if false {
                                            var v51 uint = uint(v3)
                                            v52 := GenS2{F0: -584}
                                            _ = v51
                                            _ = v52
                                        }
                                        if (-2147483648 > (-(fi22))) {
                                            fmt.Println(v19, v21, v2)
                                        }
                                        _ = v50
                                    } else {
                                        v44 = "hello"
                                        delete(v5, "hello")
                                    }
                                    switch v44 {
                                    case "":
                                        v4 -= v4
                                    }
                                }
                            }
                        }
                        _ = v44
                    }
                case 127:
                    if (!false) {
                        var v53 int64 = 750
                        if len(v6) > 0 {
                            v6[0] = 1
                        }
                        _ = v53
                    } else {
                        if false {
                            for fi54 := 0; fi54 < 17; fi54++ {
                                switch "" {
                                case "a":
                                    if false {
                                        v21 += genFn0(232, v3, (fi22 * (fi54 << ((18446744073709551615) & 63))))
                                        v4 += (v4 << (((v21 / ((uint(v2)) | 1))) & 63))
                                        v4 += uint64(v3)
                                    } else {
                                        if false {
                                            v55 := GenS3{F0: -377}
                                            v19 -= v19
                                            v2 = (v2 - 65535)
                                            v56 := GenS4{F0: -219, F1: uint8(v2), F2: -2147483648}
                                            v57 := GenS4{F0: -924, F1: v3, F2: -702}
                                            _ = v55
                                            _ = v56
                                            _ = v57
                                        } else {
                                            v4 -= uint64(v21)
                                            fmt.Println(v4)
                                            v21 -= 65535
                                        }
                                        var v58 int8 = 0
                                        v4 = v4
                                        _ = v58
                                    }
                                case "foo bar":
                                    v4 -= v4
                                case "":
                                    v3 = uint8(v4)
                                }
                                v2 = 207
                                fmt.Println(v19)
                                delete(v1, (((("a" + "") + ("" + "")) + ("" + ("hello" + "a"))) + ""))
                                if len(v6) >= 2 {
                                    v6 = v6[1:len(v6)]
                                }
                                _ = fi54
                            }
                        }
                        var v59 int8 = 90
                        v2 -= (v2 ^ v2)
                        _ = v59
                    }
                case -540:
                    for fi60 := 0; fi60 < 14; fi60++ {
                        if true {
                            _ = len(v5)
                            v61 := GenS5{F0: (^(v19))}
                            v62 := GenS2{F0: 190}
                            _ = v61
                            _ = v62
                        } else {
                            switch (fi22 % ((-540) | 1)) {
                            case -1:
                                _ = len(v1)
                            case -428:
                                if ((((0.0 + 1e10) - 3.14159) == -2.25) || (("a" + "") < ("a" + "foo bar"))) {
                                    _ = len(v6)
                                }
                            }
                            v4 = uint64(v21)
                            if ((v21 % ((42) | 1)) < v21) {
                                if true {
                                    var v63 uint64 = (v4 & (v4 ^ v4))
                                    v21 += genFn0(0, (v3 >> (((v21 | (v21 - genFn0((v19 / ((uint16(v3)) | 1)), (v3 << ((458) & 63)), fi60)))) & 63)), (fi22 % (((fi22 % ((951) | 1))) | 1)))
                                    v3 += (v3 | (v3 | v3))
                                    _ = v63
                                }
                                v2 = 2147483647
                                var v64 uint32 = (v2 / ((v2) | 1))
                                _ = v64
                            }
                        }
                        delete(v5, ((("a" + "hello") + "hello") + ("\n\t" + ("a" + "a"))))
                        _ = fi60
                    }
                default:
                    if len(v6) > 0 {
                        v6[0] = (^(v2))
                    }
                }
                fmt.Println(v3)
            }
        }
        _ = fi22
        _ = v23
    }
    v2 -= v2
    _ = len(v6)
    v21 -= genFn0((v19 - 32767), v3, -1)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
