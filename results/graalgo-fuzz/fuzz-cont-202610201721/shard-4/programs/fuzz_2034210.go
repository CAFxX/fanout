package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8; F2 int8 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int32; F1 uint }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 int32; F1 uint16; F2 int16 }
type GenS6 struct { F0 uint64; F1 string; F2 uint16 }
type GenS7 struct { F0 int32; F1 bool; F2 uint16 }
type GenS8 struct { F0 uint8 }
type GenS9 struct { F0 bool }
type GenS10 struct { F0 int64; F1 int8; F2 uint8 }
type GenS11 struct { F0 int; F1 uint64; F2 uint8 }
type GenS12 struct { F0 int64; F1 uint }
type GenS13 struct { F0 uint8 }
type GenS14 struct { F0 int64; F1 int32; F2 uint }
type GenS15 struct { F0 int32; F1 uint8 }
type GenS16 struct { F0 string; F1 uint16 }
type GenS17 struct { F0 uint16; F1 string }
func genFn0(p0 uint) int {
    v7 := GenS1{F0: 182, F1: 42, F2: 8}
    fmt.Println(p0)
    if true {
        var v8 int = 32767
        _ = v8
    }
    v9 := []int{127}
    _ = p0
    _ = v7
    _ = v9
    return 128
}
func genFn1() int64 {
    var v10 uint64 = 0
    v11 := GenS1{F0: 638, F1: -31, F2: -1}
    v12 := GenS2{F0: -128}
    var v13 uint = 65535
    var v14 int = -2147483648
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    return 0
}

func main() {
    var v1 int16 = -781
    v2 := make(map[uint32]bool)
    var v3 uint8 = 127
    var v4 uint32 = 4294967295
    var v5 bool = false
    var v6 int32 = -1
    v15 := GenS1{F0: (-(v1)), F1: -38, F2: 100}
    v1 = v1
    if v5 {
        for fi16 := 0; fi16 < 6; fi16++ {
            for fi17 := 0; fi17 < 5; fi17++ {
                var v18 uint64 = 4294967295
                delete(v2, 2147483647)
                _ = fi17
                _ = v18
            }
            v19 := []int{(fi16 / ((genFn0(0)) | 1)), (fi16 >> ((255) & 63)), (fi16 - (fi16 & (fi16 | int(v6)))), 32768}
            v6 = v6
            _ = fi16
            _ = v19
        }
    } else {
        v20 := GenS1{F0: (v1 >> ((727) & 63)), F1: 127, F2: -122}
        var v21 uint16 = 843
        v6 -= v6
        _ = v20
        _ = v21
    }
    for fi22 := 0; fi22 < 12; fi22++ {
        for fi23 := 0; fi23 < 15; fi23++ {
            delete(v2, v4)
            if v5 {
                switch 635 {
                case 32768:
                    fmt.Println(v3)
                case -341:
                    if v5 {
                        var v24 uint16 = 255
                        v25 := GenS3{F0: v6, F1: 472}
                        v26 := GenS3{F0: (v6 << ((65535) & 63)), F1: 42}
                        _ = v24
                        _ = v25
                        _ = v26
                    }
                default:
                    if (((-(fi22)) == 128) || (42 > 1)) {
                        v27 := GenS1{F0: v1, F1: 100, F2: -1}
                        v4 += 2
                        if v5 {
                            v4 = v4
                            v28 := GenS4{F0: 296}
                            v3 += v3
                            _ = v28
                        } else {
                            var v29 uint32 = (v4 << ((18446744073709551615) & 63))
                            v30 := []int{(-(fi23)), (-(fi23))}
                            fmt.Println(v4, fi22, v1)
                            _ = v29
                            _ = v30
                        }
                        _ = v27
                    }
                }
                v4 = 2147483647
                v6 = -517
            } else {
                if false {
                    v31 := GenS5{F0: -600, F1: 65535, F2: (v1 ^ v1)}
                    if true {
                        var v32 int64 = genFn1()
                        var v33 uint64 = 4294967295
                        _ = v32
                        _ = v33
                    } else {
                        if false {
                            fmt.Println(fi22, fi23)
                            v1 -= v1
                            var v34 bool = ((fi23 >> ((1) & 63)) > fi22)
                            _ = v34
                        } else {
                            if (v4 == 1) {
                                switch fi22 {
                                case 917:
                                    var v35 string = ("a" + "hello")
                                    _ = v35
                                case 0:
                                    v36 := GenS6{F0: 18446744073709551615, F1: "", F2: 872}
                                    _ = v36
                                }
                            } else {
                                v2[42] = (v5 && v5)
                                v2[(v4 % ((v4) | 1))] = (v4 > 42)
                            }
                            v3 = v3
                            switch "\n\t" {
                            case "a":
                                if false {
                                    var v37 int8 = -1
                                    _ = v37
                                }
                            case "\n\t":
                                v38 := v2[(^(v4))]
                                _ = v38
                            case "":
                                var v39 uint16 = 255
                                _ = v39
                            }
                        }
                    }
                    switch ("\n\t" + "hello") {
                    case "a":
                        v6 += (v6 * (v6 % ((v6) | 1)))
                    }
                    _ = v31
                }
                switch ("\n\t" + (("hello" + "\n\t") + ("foo bar" + "\n\t"))) {
                case "":
                    if v5 {
                        var v40 int = fi22
                        _ = v40
                    } else {
                        var v41 float32 = float32(3.14159)
                        _ = v41
                    }
                }
                if v5 {
                    delete(v2, (v4 % (((v4 & v4)) | 1)))
                } else {
                    var v42 int16 = v1
                    var v43 float32 = 1e10
                    _ = len(v2)
                    if v5 {
                        fmt.Println(v1, v42)
                        _ = len(v2)
                        if v5 {
                            if (v5 || (((("hello" + "") + "hello") < "hello") || v5)) {
                                var v44 uint8 = v3
                                v45 := GenS7{F0: v6, F1: v5, F2: 255}
                                _ = v44
                                _ = v45
                            }
                            var v46 uint = 266
                            _ = v46
                        }
                    }
                    v1 = (v1 << ((0) & 63))
                    v43 = (v43 / 1.5)
                    _ = v42
                    _ = v43
                }
            }
            switch fi23 {
            case 128:
                fmt.Println(v5, fi22)
            }
            _ = fi23
        }
        _ = fi22
    }
    if (4294967295 < 0) {
        for fi47 := 0; fi47 < 3; fi47++ {
            if (!v5) {
                v48 := []int{2147483647, (fi47 ^ (fi47 << ((2) & 63)))}
                if (!((0.0 == float32((0.1 / ((-2.25 / 1e10) / 1.0)))) && v5)) {
                    for fi49 := 0; fi49 < 16; fi49++ {
                        switch "hello" {
                        case "":
                            switch "a" {
                            case "\n\t":
                                v1 = 42
                            case "":
                                if ((v1 | (-(v1))) < v1) {
                                    switch fi49 {
                                    case 0:
                                        if false {
                                            var v50 int8 = 42
                                            var v51 int = (-(fi49))
                                            if len(v48) >= 2 {
                                                v48 = v48[1:len(v48)]
                                            }
                                            _ = v50
                                            _ = v51
                                        } else {
                                            if v5 {
                                                delete(v2, 4294967295)
                                                var v52 float64 = (((0.1 / 1.0) / 0.1) - -2.25)
                                                _ = v52
                                            } else {
                                                v53 := GenS6{F0: 259, F1: ("" + ("a" + "foo bar")), F2: 42}
                                                fmt.Println(v6, fi47)
                                                v5 = v5
                                                _ = v53
                                            }
                                            v6 -= v6
                                            if (false && v5) {
                                                v3 -= (v3 ^ (v3 * v3))
                                                v54 := GenS1{F0: 1000, F1: 0, F2: int8(fi47)}
                                                _ = v54
                                            } else {
                                                var v55 int32 = v6
                                                _ = v55
                                            }
                                        }
                                    }
                                }
                            default:
                                if (42 == 127) {
                                    v2[(v4 & 539)] = v5
                                    v4 = (v4 / ((0) | 1))
                                    var v56 int32 = -570
                                    fmt.Println(v56)
                                    var v57 uint16 = 2
                                    var v58 int = 32768
                                    var v59 string = ((("\n\t" + ("foo bar" + "a")) + "a") + "a")
                                    _ = v56
                                    _ = v57
                                    _ = v58
                                    _ = v59
                                } else {
                                    if len(v48) >= 2 {
                                        v48 = v48[1:len(v48)]
                                    }
                                    v6 -= (^(v6))
                                    if (1.5 <= float64(1e10)) {
                                        v3 = (v3 - (v3 >> ((18446744073709551615) & 63)))
                                        if (0 != 18446744073709551615) {
                                            v60 := GenS5{F0: 0, F1: 65535, F2: v1}
                                            v61 := GenS8{F0: (v3 >> ((65535) & 63))}
                                            var v62 uint8 = v3
                                            switch "foo bar" {
                                            case "":
                                                switch "\n\t" {
                                                case "\n\t":
                                                    var v63 uint64 = 2
                                                    fmt.Println(v4, v62)
                                                    if (1.5 > 0.1) {
                                                        if (((v62 << ((42) & 63)) <= v3) || (((float64(0.0) - 0.1) - 1e10) >= ((1e10 - (0.0 / float64(float32((0.1 - 0.0))))) - 0.1))) {
                                                            if v5 {
                                                                v64 := GenS9{F0: v5}
                                                                var v65 int16 = (-(v1))
                                                                _ = v64
                                                                _ = v65
                                                            }
                                                            delete(v2, v4)
                                                        } else {
                                                            var v66 int = (fi49 * 32768)
                                                            delete(v2, uint32(v63))
                                                            _ = v66
                                                        }
                                                    } else {
                                                        fmt.Println(v62, v63, v62)
                                                    }
                                                    v63 -= (v63 ^ 18446744073709551615)
                                                    _ = v63
                                                default:
                                                    if true {
                                                        v67 := GenS10{F0: -152, F1: 100, F2: v3}
                                                        delete(v2, v4)
                                                        v68 := GenS11{F0: -1, F1: 0, F2: (v62 >> ((255) & 63))}
                                                        _ = v67
                                                        _ = v68
                                                    }
                                                }
                                            }
                                            if v5 {
                                                var v69 int = int(v4)
                                                _ = v69
                                            }
                                            v6 = (v6 << ((421) & 63))
                                            v5 = (2 < (v62 | (^(v3))))
                                            var v70 int16 = (v1 % ((v1) | 1))
                                            _ = v60
                                            _ = v61
                                            _ = v62
                                            _ = v70
                                        }
                                    } else {
                                        v4 = (v4 % ((2) | 1))
                                        _ = len(v48)
                                        v71 := GenS4{F0: 1}
                                        v4 = v4
                                        _ = v71
                                    }
                                    v72 := v2[0]
                                    _ = v72
                                }
                            }
                        }
                        _ = fi49
                    }
                    if (v1 == (v1 & (-(v1)))) {
                        fmt.Println(v3)
                        if false {
                            v3 -= 128
                            var v73 int = (fi47 + genFn0(554))
                            v4 += (v4 | (v4 + 1))
                            _ = v73
                        }
                    } else {
                        v74 := GenS11{F0: fi47, F1: 640, F2: 42}
                        v75 := GenS12{F0: int64(v4), F1: 610}
                        _ = v74
                        _ = v75
                    }
                } else {
                    v76 := GenS13{F0: v3}
                    fmt.Println(v1, v3)
                    v77 := GenS10{F0: 0, F1: 1, F2: 128}
                    _ = v76
                    _ = v77
                }
                if (0 != 1) {
                    switch (fi47 + fi47) {
                    case -9223372036854775808:
                        if len(v48) >= 2 {
                            v48 = v48[1:len(v48)]
                        }
                    case 9223372036854775807:
                        v78 := GenS14{F0: 860, F1: (v6 >> ((219) & 63)), F2: 768}
                        _ = v78
                    }
                    var v79 int = 0
                    v80 := v2[v4]
                    _ = v79
                    _ = v80
                }
                _ = v48
            }
            v4 = 1
            v81 := GenS15{F0: v6, F1: v3}
            _ = fi47
            _ = v81
        }
    }
    v1 += 0
    v82 := v2[(v4 | v4)]
    if v82 {
        if (1e10 == (1e10 - (float64(1e10) / float64(((0.0 + (float32(float64(((0.1 + float32(0.0)) / -2.25))) / 3.14159)) * 0.1))))) {
            v2[(v4 << ((65535) & 63))] = v82
        }
        v83 := GenS16{F0: "a", F1: 1}
        _ = v83
    }
    switch 4294967295 {
    case 1:
        var v84 float32 = 1e10
        _ = v84
    default:
        v85 := GenS14{F0: -1, F1: (v6 ^ v6), F2: 968}
        _ = v85
    }
    v86 := GenS11{F0: -706, F1: 118, F2: 109}
    v1 = v1
    v87 := GenS17{F0: 42, F1: "\n\t"}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v82:", v82)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v82
    _ = v86
    _ = v87
}
