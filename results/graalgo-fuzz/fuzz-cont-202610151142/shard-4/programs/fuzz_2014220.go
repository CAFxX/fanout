package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 bool; F1 uint32; F2 uint32 }
type GenS5 struct { F0 int8; F1 uint8 }
type GenS6 struct { F0 string; F1 string; F2 string }
type GenS7 struct { F0 uint16 }
type GenS8 struct { F0 string; F1 int8; F2 int16 }
type GenS9 struct { F0 int; F1 int; F2 bool }
type GenS10 struct { F0 uint32; F1 uint32 }
type GenS11 struct { F0 int16; F1 int32; F2 uint8 }
type GenS12 struct { F0 uint32; F1 int32; F2 int16 }
type GenS13 struct { F0 int64; F1 int64 }
func genFn0(p0 bool, p1 int32) uint32 {
    var v6 string = ""
    switch 2147483647 {
    case -9223372036854775808:
        v6 = "a"
    case 2147483647:
        if p0 {
            p1 += (-(p1))
        } else {
            v7 := GenS1{F0: (((("a" + ((v6 + (v6 + (v6 + v6))) + (v6 + "hello"))) + v6) + ("" + "\n\t")) + v6), F1: 1}
            _ = v7
        }
    }
    fmt.Println(v6)
    v8 := []int{4294967295, -259, 788}
    if p0 {
        for fi9 := 0; fi9 < 4; fi9++ {
            var v10 uint = 2
            v8 = append(v8, fi9)
            _ = fi9
            _ = v10
        }
    }
    _ = p0
    _ = p1
    _ = v6
    _ = v8
    return 42
}
func genFn1(p0 uint, p1 uint16, p2 uint8) uint32 {
    p0 = (p0 << ((p0) & 63))
    fmt.Println(p2)
    if (!(true && true)) {
        v11 := []int{128, -2147483648, 128}
        if true {
            v11 = append(v11, 32768)
        }
        _ = v11
    }
    _ = p0
    _ = p1
    _ = p2
    return uint32(p1)
}

func main() {
    var v1 float32 = 1e10
    var v2 uint32 = 4294967295
    var v3 uint16 = 1
    var v4 uint32 = v2
    var v5 uint8 = 42
    for fi12 := 0; fi12 < 17; fi12++ {
        v3 -= 65535
        v13 := []int{(fi12 >> ((1) & 63)), fi12, fi12, (fi12 >> ((42) & 63))}
        var v14 int64 = 9223372036854775807
        _ = fi12
        _ = v13
        _ = v14
    }
    var v15 float64 = 1e10
    v16 := []int{127, 1, 1, 859}
    switch -2147483648 {
    case -246:
        var v17 int = 2147483647
        _ = v17
    case -785:
        _ = len(v16)
    }
    if true {
        for fi18 := 0; fi18 < 16; fi18++ {
            var v19 int16 = -678
            for fi20 := 0; fi20 < 7; fi20++ {
                v15 = (float64((v1 + (0.0 / (v1 + ((1e10 * 1e10) / (((0.1 * (0.1 / v1)) - float32(v15)) - v1)))))) - (1e10 - v15))
                if (float64(v1) != (((0.0 * float64(v1)) + v15) * ((v15 - float64(((v1 - 0.0) - 3.14159))) / (v15 / ((v15 / float64(v1)) * ((float64(v1) * v15) * v15)))))) {
                    if true {
                        v21 := GenS2{F0: -1}
                        _ = v21
                    } else {
                        v19 = v19
                        var v22 uint = 255
                        var v23 uint64 = 4294967295
                        if (true || true) {
                            if true {
                                var v24 uint16 = 2
                                _ = v24
                            } else {
                                switch fi20 {
                                case 978:
                                    var v25 string = ""
                                    _ = v25
                                case 32767:
                                    var v26 uint32 = (v2 & v4)
                                    _ = v26
                                }
                                fmt.Println(v5, v22)
                            }
                            var v27 bool = true
                            fmt.Println(fi20)
                            _ = v27
                        }
                        _ = v22
                        _ = v23
                    }
                }
                v1 += v1
                _ = fi20
            }
            _ = fi18
            _ = v19
        }
        if (uint(v4) > 18446744073709551615) {
            for fi28 := 0; fi28 < 8; fi28++ {
                for fi29 := 0; fi29 < 12; fi29++ {
                    v30 := GenS1{F0: "a", F1: (v2 ^ 2)}
                    v31 := GenS1{F0: "\n\t", F1: v4}
                    v32 := GenS3{F0: -32768}
                    _ = fi29
                    _ = v30
                    _ = v31
                    _ = v32
                }
                _ = fi28
            }
            v33 := GenS2{F0: 104}
            if true {
                var v34 uint8 = 2
                switch 32768 {
                case 128:
                    fmt.Println(v2)
                case -104:
                    var v35 int64 = -506
                    _ = v35
                case 179:
                    v5 -= 129
                default:
                    if true {
                        for fi36 := 0; fi36 < 2; fi36++ {
                            var v37 float32 = (v1 * float32(float64(v1)))
                            fmt.Println(v37)
                            _ = fi36
                            _ = v37
                        }
                        for fi38 := 0; fi38 < 6; fi38++ {
                            for fi39 := 0; fi39 < 5; fi39++ {
                                if false {
                                    if (v4 == genFn0(false, 1)) {
                                        var v40 uint16 = (v3 | (v3 ^ v3))
                                        v41 := GenS4{F0: ((v4 ^ (v4 + (v4 - 65535))) >= v4), F1: (v4 << ((255) & 63)), F2: (v4 - v4)}
                                        _ = v40
                                        _ = v41
                                    }
                                    var v42 int = fi38
                                    v1 = 1e10
                                    _ = v42
                                } else {
                                    v43 := GenS5{F0: 81, F1: v34}
                                    v44 := GenS6{F0: "a", F1: "foo bar", F2: "\n\t"}
                                    fmt.Println(v2, v34)
                                    _ = v43
                                    _ = v44
                                }
                                _ = fi39
                            }
                            if (false && true) {
                                if len(v16) >= 2 {
                                    v16 = v16[1:len(v16)]
                                }
                                for fi45 := 0; fi45 < 11; fi45++ {
                                    fmt.Println(v1)
                                    if ((!(false || ((!(true || false)) || (!false)))) || (-632 > 42)) {
                                        v46 := GenS2{F0: 6}
                                        v47 := GenS7{F0: (v3 | (v3 % (((v3 % (((v3 | 461)) | 1))) | 1)))}
                                        var v48 int64 = 1
                                        var v49 string = "hello"
                                        var v50 uint8 = (v34 + 42)
                                        _ = v46
                                        _ = v47
                                        _ = v48
                                        _ = v49
                                        _ = v50
                                    } else {
                                        switch "" {
                                        case "a":
                                            var v51 uint16 = v3
                                            _ = v51
                                        case "hello":
                                            v52 := GenS8{F0: "foo bar", F1: 1, F2: 565}
                                            _ = v52
                                        case "":
                                            if true {
                                                var v53 int = (fi38 / ((fi45) | 1))
                                                fmt.Println(v2)
                                                _ = v53
                                            }
                                        default:
                                            if true {
                                                _ = len(v16)
                                                switch ("a" + ("" + "a")) {
                                                case "":
                                                    if ((^(v34)) < v34) {
                                                        var v54 int8 = 29
                                                        if true {
                                                            v2 = (v2 + (v2 | 1))
                                                        }
                                                        _ = v54
                                                    } else {
                                                        var v55 float64 = v15
                                                        _ = v55
                                                    }
                                                case "\n\t":
                                                    if true {
                                                        if true {
                                                            fmt.Println(fi38, v4)
                                                            var v56 int64 = -1
                                                            var v57 string = ("" + ("" + ("\n\t" + "")))
                                                            _ = v56
                                                            _ = v57
                                                        }
                                                        v58 := GenS9{F0: fi38, F1: fi38, F2: (-32768 != 1)}
                                                        v5 += v5
                                                        _ = v58
                                                    } else {
                                                        if ((v1 - (1.5 / v1)) < v1) {
                                                            if (!true) {
                                                                if true {
                                                                    if len(v16) > 0 {
                                                                        v16[0] = 127
                                                                    }
                                                                } else {
                                                                    v59 := GenS10{F0: (v2 & 2), F1: 1}
                                                                    _ = v59
                                                                }
                                                                fmt.Println(v1)
                                                                if false {
                                                                    var v60 int = int(v3)
                                                                    v16 = append(v16, v60)
                                                                    if len(v16) > 0 {
                                                                        v16[0] = (^(fi38))
                                                                    }
                                                                    _ = v60
                                                                } else {
                                                                    var v61 int64 = 1
                                                                    if true {
                                                                        fmt.Println(v15, v34)
                                                                        var v62 float32 = (0.0 + (v1 + v1))
                                                                        v63 := GenS1{F0: ("\n\t" + "\n\t"), F1: v2}
                                                                        _ = v62
                                                                        _ = v63
                                                                    } else {
                                                                        v15 -= 1e10
                                                                        if true {
                                                                            if true {
                                                                                v16 = append(v16, fi38)
                                                                                if ((!false) || true) {
                                                                                    if false {
                                                                                        var v64 float64 = float64(1e10)
                                                                                        if true {
                                                                                            v16 = append(v16, 32767)
                                                                                            if len(v16) > 0 {
                                                                                                v16[0] = (fi38 + (fi45 & 32767))
                                                                                            }
                                                                                        }
                                                                                        switch (fi45 & (-(fi45))) {
                                                                                        case 32768:
                                                                                            v64 -= v64
                                                                                        case 128:
                                                                                            var v65 int8 = 127
                                                                                            _ = v65
                                                                                        case 9223372036854775807:
                                                                                            if (v61 != -9223372036854775808) {
                                                                                                var v66 uint8 = v5
                                                                                                _ = v66
                                                                                            }
                                                                                        default:
                                                                                            var v67 float32 = v1
                                                                                            _ = v67
                                                                                        }
                                                                                        _ = v64
                                                                                    } else {
                                                                                        v16 = append(v16, fi45)
                                                                                        if len(v16) >= 2 {
                                                                                            v16 = v16[1:len(v16)]
                                                                                        }
                                                                                        v3 += (v3 * (^(v3)))
                                                                                    }
                                                                                }
                                                                                v68 := GenS8{F0: "a", F1: 1, F2: -1}
                                                                                _ = v68
                                                                            } else {
                                                                                var v69 float32 = (v1 + v1)
                                                                                _ = v69
                                                                            }
                                                                            var v70 int16 = -544
                                                                            var v71 int = 9223372036854775807
                                                                            switch fi45 {
                                                                            case -2147483648:
                                                                                var v72 uint64 = 4294967295
                                                                                _ = v72
                                                                            default:
                                                                                var v73 uint16 = v3
                                                                                v74 := GenS11{F0: int16(v2), F1: 42, F2: (^(v5))}
                                                                                _ = v73
                                                                                _ = v74
                                                                            }
                                                                            v75 := GenS12{F0: v4, F1: 42, F2: (-(v70))}
                                                                            if ((1.5 / v15) <= ((v15 + 1e10) + v15)) {
                                                                                var v76 uint64 = 1
                                                                                _ = v76
                                                                            }
                                                                            _ = v70
                                                                            _ = v71
                                                                            _ = v75
                                                                        }
                                                                    }
                                                                    v77 := GenS1{F0: "\n\t", F1: genFn1(uint(v3), (v3 | v3), (v5 + v34))}
                                                                    _ = v61
                                                                    _ = v77
                                                                }
                                                                v78 := GenS6{F0: (("a" + "hello") + "hello"), F1: "", F2: ("foo bar" + "\n\t")}
                                                                switch 42 {
                                                                case 128:
                                                                    var v79 uint16 = v3
                                                                    _ = v79
                                                                case 2147483647:
                                                                    fmt.Println(v15, fi45)
                                                                }
                                                                _ = v78
                                                            }
                                                            v1 = 1e10
                                                        }
                                                    }
                                                case "hello":
                                                    v15 -= 1.5
                                                default:
                                                    var v80 int16 = -82
                                                    _ = v80
                                                }
                                                fmt.Println(v2, v34, v15)
                                            } else {
                                                if len(v16) >= 2 {
                                                    v16 = v16[1:len(v16)]
                                                }
                                            }
                                        }
                                    }
                                    _ = fi45
                                }
                                v3 = 1
                            } else {
                                fmt.Println(fi38, v34, v1)
                                v4 = (v4 << ((42) & 63))
                                for fi81 := 0; fi81 < 11; fi81++ {
                                    if len(v16) >= 2 {
                                        v16 = v16[1:len(v16)]
                                    }
                                    _ = fi81
                                }
                            }
                            _ = fi38
                        }
                        v4 -= (v2 * v2)
                    } else {
                        v15 -= ((1e10 - v15) / v15)
                    }
                }
                v34 = 255
                _ = v34
            }
            _ = v33
        } else {
            v16 = append(v16, -1)
            if true {
                v3 -= 32767
                if (!(v4 >= (^(v4)))) {
                    switch 29 {
                    case 42:
                        var v82 uint16 = (v3 % ((v3) | 1))
                        _ = v82
                    case 32767:
                        var v83 int8 = 0
                        _ = v83
                    }
                } else {
                    v84 := GenS13{F0: int64(v5), F1: -200}
                    fmt.Println(v4, v5, v5)
                    v1 -= float32(0.0)
                    _ = v84
                }
                if (0.0 <= (((float32(v15) / v1) * 1e10) - 1e10)) {
                    var v85 int16 = 29
                    fmt.Println(v3, v4)
                    var v86 int16 = v85
                    _ = v85
                    _ = v86
                }
            } else {
                var v87 uint64 = 18446744073709551615
                _ = v87
            }
            if len(v16) >= 2 {
                v16 = v16[1:len(v16)]
            }
        }
        v88 := GenS1{F0: "\n\t", F1: uint32(v3)}
        _ = v88
    }
    fmt.Println(v2, v5, v4)
    if false {
        var v89 string = "foo bar"
        v90 := GenS7{F0: 0}
        _ = v89
        _ = v90
    }
    var v91 uint8 = v5
    v16 = append(v16, 32767)
    var v92 int16 = -32768
    var v93 uint8 = (v5 - v91)
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v91:", v91)
    fmt.Println("v92:", v92)
    fmt.Println("v93:", v93)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v91
    _ = v92
    _ = v93
}
