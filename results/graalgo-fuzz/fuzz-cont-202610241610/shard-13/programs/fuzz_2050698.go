package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32; F2 uint64 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint64; F1 int32; F2 bool }
type GenS4 struct { F0 int32; F1 int64; F2 uint32 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 uint64 }
func genFn0() uint32 {
    for fi4 := 0; fi4 < 20; fi4++ {
        if false {
            var v5 uint = 2
            fmt.Println(v5, v5)
            _ = v5
        }
        _ = fi4
    }
    var v6 string = ("hello" + (("\n\t" + "") + "foo bar"))
    if false {
        var v7 bool = (721 < 1)
        _ = v7
    } else {
        v8 := []int{}
        _ = v8
    }
    if (-2.25 != 1e10) {
        for fi9 := 0; fi9 < 3; fi9++ {
            v6 = "foo bar"
            v10 := GenS1{F0: 2, F1: 65535, F2: 292}
            _ = fi9
            _ = v10
        }
        switch "" {
        case "\n\t":
            v6 = (v6 + v6)
        case "a":
            v6 = v6
        case "hello":
            fmt.Println(v6)
        default:
            var v11 uint64 = 398
            _ = v11
        }
        fmt.Println(v6)
    } else {
        v6 = "hello"
        for fi12 := 0; fi12 < 18; fi12++ {
            var v13 int = fi12
            var v14 bool = ("hello" < (v6 + v6))
            _ = fi12
            _ = v13
            _ = v14
        }
        for fi15 := 0; fi15 < 9; fi15++ {
            v6 = ""
            _ = fi15
        }
    }
    _ = v6
    return 0
}
func genFn1(p0 uint32, p1 int8) uint32 {
    p0 += genFn0()
    v16 := []int{4294967295, 381, -2147483648, 42}
    _ = p0
    _ = p1
    _ = v16
    return p0
}

func main() {
    var v1 uint64 = 0
    var v2 bool = (14 < 1)
    var v3 uint = 18446744073709551615
    for fi17 := 0; fi17 < 10; fi17++ {
        v18 := []int{6, (-(fi17)), fi17}
        var v19 uint16 = 255
        v19 -= 0
        _ = fi17
        _ = v18
        _ = v19
    }
    for fi20 := 0; fi20 < 8; fi20++ {
        for fi21 := 0; fi21 < 9; fi21++ {
            v2 = (125 < -840)
            v22 := []int{0}
            _ = fi21
            _ = v22
        }
        fmt.Println(v1)
        for fi23 := 0; fi23 < 2; fi23++ {
            if v2 {
                var v24 int32 = -219
                v25 := []int{fi20, (fi20 % ((32767) | 1))}
                var v26 bool = (9223372036854775807 < 9223372036854775807)
                _ = v24
                _ = v25
                _ = v26
            } else {
                var v27 int32 = -2147483648
                switch fi23 {
                case 42:
                    var v28 float32 = 3.14159
                    _ = v28
                }
                _ = v27
            }
            _ = fi23
        }
        _ = fi20
    }
    v29 := GenS1{F0: 0, F1: 807, F2: (v1 << ((18446744073709551615) & 63))}
    for fi30 := 0; fi30 < 18; fi30++ {
        var v31 int16 = 32767
        _ = fi30
        _ = v31
    }
    v32 := []int{}
    fmt.Println(v3)
    var v33 uint64 = v1
    for fi34 := 0; fi34 < 6; fi34++ {
        v33 = (^(v1))
        switch fi34 {
        case -2147483648:
            fmt.Println(v3, v3, fi34)
        case 32767:
            var v35 string = ("" + "a")
            _ = v35
        case 4294967295:
            if v2 {
                switch (((("a" + "hello") + "\n\t") + "foo bar") + "") {
                case "a":
                    for fi36 := 0; fi36 < 1; fi36++ {
                        if len(v32) > 0 {
                            v32[0] = (fi34 / (((^(fi36))) | 1))
                        }
                        _ = fi36
                    }
                case "hello":
                    _ = len(v32)
                case "":
                    var v37 uint = v3
                    _ = v37
                default:
                    var v38 uint16 = 585
                    _ = v38
                }
                var v39 int32 = 0
                var v40 int = (fi34 + 42)
                _ = v39
                _ = v40
            } else {
                if true {
                    switch "foo bar" {
                    case "a":
                        switch ("hello" + "hello") {
                        case "foo bar":
                            for fi41 := 0; fi41 < 2; fi41++ {
                                if (128 != 179) {
                                    var v42 int = 32768
                                    v1 += (v33 ^ v33)
                                    _ = v42
                                } else {
                                    v32 = append(v32, -1)
                                    var v43 float64 = float64(3.14159)
                                    _ = len(v32)
                                    _ = v43
                                }
                                var v44 uint32 = 221
                                _ = fi41
                                _ = v44
                            }
                        case "\n\t":
                            switch "\n\t" {
                            case "":
                                if len(v32) >= 2 {
                                    v32 = v32[1:len(v32)]
                                }
                            }
                        case "hello":
                            for fi45 := 0; fi45 < 9; fi45++ {
                                if len(v32) >= 2 {
                                    v32 = v32[1:len(v32)]
                                }
                                var v46 uint32 = 1
                                v46 = v46
                                var v47 int8 = int8(fi45)
                                if true {
                                    fmt.Println(v1)
                                    v48 := GenS1{F0: 2, F1: v46, F2: (^(v1))}
                                    v2 = ("foo bar" > "a")
                                    if false {
                                        v46 -= v46
                                    } else {
                                        var v49 int32 = 2147483647
                                        v50 := GenS1{F0: 32767, F1: 2, F2: v33}
                                        var v51 string = (("foo bar" + "hello") + ("" + ("a" + "foo bar")))
                                        _ = v49
                                        _ = v50
                                        _ = v51
                                    }
                                    _ = v48
                                } else {
                                    v52 := GenS2{F0: 243}
                                    _ = v52
                                }
                                _ = fi45
                                _ = v46
                                _ = v47
                            }
                        default:
                            fmt.Println(v1)
                        }
                    case "\n\t":
                        for fi53 := 0; fi53 < 13; fi53++ {
                            var v54 int32 = -769
                            if len(v32) >= 2 {
                                v32 = v32[1:len(v32)]
                            }
                            v55 := GenS3{F0: v33, F1: 42, F2: (!v2)}
                            _ = fi53
                            _ = v54
                            _ = v55
                        }
                    }
                    for fi56 := 0; fi56 < 16; fi56++ {
                        v3 += v3
                        _ = fi56
                    }
                    var v57 int8 = 18
                    _ = v57
                } else {
                    if v2 {
                        v58 := GenS2{F0: -758}
                        _ = v58
                    }
                    v3 -= (v3 & (v3 + v3))
                }
                _ = len(v32)
            }
        }
        v1 += 42
        _ = fi34
    }
    if (true || (!true)) {
        if ((0.0 == (3.14159 * (1.5 + 0.0))) || (-678 >= -9223372036854775808)) {
            var v59 int8 = -51
            _ = v59
        } else {
            if (-2147483648 < -14) {
                var v60 uint8 = 255
                if (!v2) {
                    var v61 int16 = -32768
                    _ = v61
                }
                var v62 int8 = 0
                _ = v60
                _ = v62
            } else {
                var v63 uint8 = 2
                _ = v63
            }
            var v64 uint = v3
            v65 := GenS4{F0: -1, F1: 42, F2: 420}
            _ = v64
            _ = v65
        }
    }
    v1 = 18446744073709551615
    var v66 uint16 = 65535
    if v2 {
        var v67 uint = (v3 * 4294967295)
        fmt.Println(v1, v1)
        var v68 float64 = 1.5
        _ = v67
        _ = v68
    } else {
        var v69 string = (((("" + ("" + "")) + "\n\t") + "\n\t") + "a")
        v3 = v3
        v70 := GenS3{F0: (v33 & (^(v1))), F1: 0, F2: v2}
        _ = v69
        _ = v70
    }
    for fi71 := 0; fi71 < 18; fi71++ {
        v66 += v66
        if true {
            _ = len(v32)
            if v2 {
                var v72 int64 = 1
                v73 := GenS2{F0: -9223372036854775808}
                _ = v72
                _ = v73
            }
            for fi74 := 0; fi74 < 1; fi74++ {
                if len(v32) >= 2 {
                    v32 = v32[1:len(v32)]
                }
                if (v2 && v2) {
                    v75 := GenS5{F0: (0 != 42)}
                    v33 -= v1
                    _ = v75
                }
                v66 = (v66 % ((v66) | 1))
                _ = fi74
            }
        } else {
            if v2 {
                fmt.Println(v66, v33)
            } else {
                v33 += v33
            }
        }
        var v76 uint8 = 54
        _ = fi71
        _ = v76
    }
    v32 = append(v32, -807)
    var v77 float32 = 0.0
    if (-2147483648 >= 262) {
        if (v2 && false) {
            v78 := GenS3{F0: (v33 % ((v33) | 1)), F1: 1, F2: (((v3 + (^(v3))) != (^(v3))) || v2)}
            for fi79 := 0; fi79 < 15; fi79++ {
                v80 := GenS5{F0: v2}
                v81 := GenS5{F0: (1 == (v66 / ((uint16(v3)) | 1)))}
                _ = fi79
                _ = v80
                _ = v81
            }
            _ = v78
        } else {
            switch "" {
            case "\n\t":
                for fi82 := 0; fi82 < 1; fi82++ {
                    v32 = append(v32, fi82)
                    var v83 uint32 = 956
                    _ = fi82
                    _ = v83
                }
            case "a":
                for fi84 := 0; fi84 < 3; fi84++ {
                    v85 := GenS6{F0: v1}
                    v33 = v1
                    _ = fi84
                    _ = v85
                }
            default:
                fmt.Println(v77, v77)
            }
            fmt.Println(v2, v1)
        }
        var v86 uint32 = 0
        var v87 uint16 = (v66 | v66)
        _ = v86
        _ = v87
    } else {
        if len(v32) > 0 {
            v32[0] = 2147483647
        }
        switch 993 {
        case 1:
            if (1 >= 0) {
                v66 -= (v66 * (v66 & 0))
                if len(v32) > 0 {
                    v32[0] = 4294967295
                }
                if (v2 && ("a" <= (((("foo bar" + ("foo bar" + "a")) + "\n\t") + "\n\t") + ""))) {
                    if false {
                        var v88 int32 = -570
                        var v89 float32 = 1e10
                        var v90 int = 0
                        _ = v88
                        _ = v89
                        _ = v90
                    } else {
                        if v2 {
                            switch "" {
                            case "a":
                                v2 = (337 < (^(v1)))
                            }
                            var v91 uint32 = 1
                            _ = v91
                        }
                    }
                } else {
                    if len(v32) >= 2 {
                        v32 = v32[1:len(v32)]
                    }
                    var v92 uint64 = (v1 % ((v33) | 1))
                    for fi93 := 0; fi93 < 9; fi93++ {
                        v33 -= (v33 - uint64(v66))
                        _ = fi93
                    }
                    _ = v92
                }
            }
        }
        var v94 int16 = -1
        _ = v94
    }
    v66 += (v66 - (v66 + 819))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v32:", len(v32), v32)
    fmt.Println("v33:", v33)
    fmt.Println("v66:", v66)
    fmt.Println("v77:", v77)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v32
    _ = v33
    _ = v66
    _ = v77
}
