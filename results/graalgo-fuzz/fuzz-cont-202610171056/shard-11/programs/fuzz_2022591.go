package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int16; F1 int32 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 int8; F1 string; F2 int8 }
type GenS5 struct { F0 uint16; F1 uint16; F2 uint16 }
type GenS6 struct { F0 int64; F1 int64 }
type GenS7 struct { F0 uint16 }
type GenS8 struct { F0 uint16; F1 int; F2 int }
type GenS9 struct { F0 int8 }
type GenS10 struct { F0 uint16 }
type GenS11 struct { F0 int8 }
type GenS12 struct { F0 uint64 }
type GenS13 struct { F0 int64; F1 int32 }
func genFn0(p0 uint, p1 uint8) string {
    fmt.Println(p1, p1)
    var v7 int32 = 2147483647
    fmt.Println(v7, p0)
    var v8 uint32 = uint32(p0)
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return "foo bar"
}
func genFn1(p0 uint16, p1 int16, p2 uint16) bool {
    if false {
        for fi9 := 0; fi9 < 11; fi9++ {
            var v10 uint16 = p2
            _ = fi9
            _ = v10
        }
    } else {
        for fi11 := 0; fi11 < 1; fi11++ {
            v12 := []int{-917, fi11}
            _ = fi11
            _ = v12
        }
    }
    v13 := []int{-1}
    if true {
        for fi14 := 0; fi14 < 9; fi14++ {
            var v15 uint8 = 2
            p0 = (p2 % ((2) | 1))
            fmt.Println(p0, v15, fi14)
            _ = fi14
            _ = v15
        }
    }
    for fi16 := 0; fi16 < 18; fi16++ {
        var v17 float32 = 1.5
        _ = fi16
        _ = v17
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    return false
}

func main() {
    var v1 int = 739
    v2 := []uint32{469}
    v3 := []string{}
    var v4 uint = 18446744073709551615
    v5 := []uint32{255, 4294967295, 2147483648, 455}
    var v6 int64 = 656
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v18 uint8 = 1
    v19 := GenS1{F0: 32767}
    v1 = v1
    var v20 int8 = 87
    for fi21 := 0; fi21 < 16; fi21++ {
        _ = len(v2)
        if true {
            for fi22 := 0; fi22 < 9; fi22++ {
                v20 += v20
                fmt.Println(v6)
                _ = fi22
            }
        }
        for fi23 := 0; fi23 < 16; fi23++ {
            v24 := GenS1{F0: -480}
            var v25 bool = true
            if true {
                fmt.Println(v25)
                v18 -= (v18 >> ((411) & 63))
                if v25 {
                    switch "hello" {
                    case "":
                        _ = len(v5)
                    }
                } else {
                    if v25 {
                        if v25 {
                            switch ("" + ("\n\t" + "")) {
                            case "hello":
                                _ = len(v5)
                            }
                            var v26 float64 = -2.25
                            fmt.Println(v1, v1, fi23)
                            _ = v26
                        }
                    } else {
                        v18 = v18
                        var v27 int32 = 2147483647
                        fmt.Println(v18, fi23, v6)
                        if ((((3.14159 * (((0.1 - float64(1e10)) + (-2.25 / 1.0)) - (float64(3.14159) - float64(0.1)))) * float64(((0.0 + 3.14159) - 0.1))) <= float64(-2.25)) && false) {
                            v28 := GenS2{F0: -1, F1: 0}
                            v25 = true
                            var v29 uint32 = 65535
                            _ = v28
                            _ = v29
                        } else {
                            fmt.Println(fi23, v27)
                            if v25 {
                                if v25 {
                                    v25 = v25
                                    if v25 {
                                        v3 = append(v3, "\n\t")
                                        v30 := GenS2{F0: -1, F1: v27}
                                        _ = v30
                                    }
                                    if v25 {
                                        if (!false) {
                                            v31 := GenS3{F0: 554}
                                            if (v27 <= -931) {
                                                v6 = v6
                                                v1 += 127
                                            } else {
                                                var v32 uint32 = 65535
                                                if len(v2) >= 2 {
                                                    v2 = v2[1:len(v2)]
                                                }
                                                var v33 string = "a"
                                                var v34 uint16 = 42
                                                var v35 uint16 = v34
                                                _ = v32
                                                _ = v33
                                                _ = v34
                                                _ = v35
                                            }
                                            _ = v31
                                        } else {
                                            switch (v1 & 279) {
                                            case 688:
                                                v36 := GenS2{F0: 239, F1: (v27 * v27)}
                                                _ = v36
                                            }
                                            if (!((!v25) || v25)) {
                                                if len(v3) > 0 {
                                                    v3[0] = "hello"
                                                }
                                                v6 -= (v6 / ((v6) | 1))
                                                var v37 int = int(v20)
                                                _ = v37
                                            } else {
                                                if len(v3) > 0 {
                                                    v3[0] = ("hello" + ("\n\t" + ("foo bar" + "hello")))
                                                }
                                                v6 = (v6 - (^(v6)))
                                            }
                                        }
                                    }
                                }
                            }
                            if ((v20 | (v20 / (((v20 ^ 127)) | 1))) > (v20 & v20)) {
                                v18 -= v18
                            } else {
                                _ = len(v2)
                                _ = len(v2)
                                v27 = v27
                            }
                        }
                        if false {
                            var v38 uint16 = 1
                            var v39 string = ""
                            v40 := GenS4{F0: v20, F1: "a", F2: (v20 >> (((^(v4))) & 63))}
                            if false {
                                v41 := GenS5{F0: 65535, F1: (^(v38)), F2: uint16(v4)}
                                v42 := GenS6{F0: 591, F1: (v6 << ((v4) & 63))}
                                if false {
                                    switch (^(fi21)) {
                                    case 9223372036854775807:
                                        _ = len(v5)
                                    }
                                    fmt.Println(fi21, v20, fi21)
                                    v43 := GenS7{F0: v38}
                                    fmt.Println(v27, v25)
                                    if (true && false) {
                                        v44 := GenS7{F0: (^(v38))}
                                        if len(v2) > 0 {
                                            v2[0] = 932
                                        }
                                        _ = v44
                                    } else {
                                        if v25 {
                                            switch v1 {
                                            case 32768:
                                                switch v39 {
                                                case "foo bar":
                                                    var v45 float32 = (1.5 - 0.0)
                                                    _ = v45
                                                default:
                                                    fmt.Println(v4, v6)
                                                }
                                            case -708:
                                                if false {
                                                    if v25 {
                                                        _ = len(v2)
                                                        var v46 float32 = 3.14159
                                                        v27 += (v27 % ((42) | 1))
                                                        _ = v46
                                                    } else {
                                                        if ((((0.0 / (1e10 * (3.14159 / 1.0))) - float32((((-2.25 / (float64(((3.14159 - float32(0.1)) * 3.14159)) / 3.14159)) + ((1e10 - 3.14159) / 1.0)) * 1e10))) / (3.14159 - float32(0.0))) <= 3.14159) {
                                                            var v47 int32 = 0
                                                            v48 := GenS1{F0: -697}
                                                            v20 = (v20 + (v20 - v20))
                                                            _ = v47
                                                            _ = v48
                                                        }
                                                        v49 := GenS8{F0: 0, F1: (fi21 ^ (v1 ^ v1)), F2: v1}
                                                        _ = v49
                                                    }
                                                    fmt.Println(fi21, v4)
                                                    v39 = v39
                                                    v50 := GenS7{F0: uint16(v18)}
                                                    _ = len(v2)
                                                    _ = v50
                                                }
                                            case -940:
                                                switch fi23 {
                                                case 585:
                                                    v5 = append(v5, 65535)
                                                case -9223372036854775808:
                                                    var v51 int = int(v27)
                                                    _ = v51
                                                case -1:
                                                    var v52 int32 = -2147483648
                                                    _ = v52
                                                default:
                                                    var v53 float32 = float32(-2.25)
                                                    _ = v53
                                                }
                                                v2 = append(v2, 42)
                                                v5 = append(v5, 610)
                                            }
                                        } else {
                                            v54 := GenS8{F0: v38, F1: v1, F2: int(v38)}
                                            v1 -= 127
                                            var v55 bool = (-846 == -591)
                                            v38 += (v38 / (((v38 & v38)) | 1))
                                            if len(v2) >= 2 {
                                                v2 = v2[1:len(v2)]
                                            }
                                            switch (v39 + v39) {
                                            case "":
                                                var v56 int32 = (v27 + (v27 << ((v4) & 63)))
                                                _ = v56
                                            default:
                                                var v57 bool = (v55 && true)
                                                _ = v57
                                            }
                                            _ = v54
                                            _ = v55
                                        }
                                    }
                                    _ = v43
                                }
                                var v58 uint8 = v18
                                v27 = (-(v27))
                                if v25 {
                                    var v59 uint32 = 2147483647
                                    if (v25 || (uint8(v38) != (v18 >> (((v4 << (((v4 >> (((v4 ^ (v4 ^ 255))) & 63))) & 63))) & 63)))) {
                                        _ = len(v5)
                                        var v60 int32 = (v27 % ((v27) | 1))
                                        _ = v60
                                    }
                                    _ = v59
                                } else {
                                    fmt.Println(v58)
                                    if (v6 != (v6 ^ v6)) {
                                        v6 = (^(v6))
                                    }
                                    v61 := GenS9{F0: (v20 << ((0) & 63))}
                                    _ = v61
                                }
                                _ = v41
                                _ = v42
                                _ = v58
                            }
                            v62 := GenS5{F0: 32767, F1: (^(v38)), F2: (^(v38))}
                            _ = v38
                            _ = v39
                            _ = v40
                            _ = v62
                        } else {
                            v63 := GenS10{F0: 692}
                            if (!(v25 && v25)) {
                                var v64 int64 = (v6 * (v6 << ((v4) & 63)))
                                v65 := GenS7{F0: 255}
                                var v66 int16 = 32767
                                _ = v64
                                _ = v65
                                _ = v66
                            } else {
                                if v25 {
                                    v20 = v20
                                }
                                if v25 {
                                    var v67 uint8 = v18
                                    fmt.Println(v27)
                                    _ = v67
                                } else {
                                    if (v25 && v25) {
                                        v2 = append(v2, 2147483648)
                                        v4 = 191
                                        var v68 uint = (v4 & v4)
                                        switch (-(fi23)) {
                                        case 128:
                                            var v69 int16 = 1000
                                            _ = v69
                                        case 32767:
                                            var v70 int8 = v20
                                            _ = v70
                                        case -1:
                                            v27 += v27
                                        default:
                                            v25 = true
                                        }
                                        if true {
                                            if v25 {
                                                var v71 string = "hello"
                                                var v72 int32 = v27
                                                var v73 uint64 = 0
                                                _ = v71
                                                _ = v72
                                                _ = v73
                                            } else {
                                                var v74 uint32 = 4294967295
                                                if len(v2) > 0 {
                                                    v2[0] = (v74 | (v74 * (v74 / (((v74 & (v74 + (v74 + (v74 | v74))))) | 1))))
                                                }
                                                v5 = append(v5, v74)
                                                _ = v74
                                            }
                                        }
                                        _ = v68
                                    } else {
                                        var v75 float64 = 0.1
                                        if v25 {
                                            v76 := GenS11{F0: 127}
                                            if len(v2) > 0 {
                                                v2[0] = 299
                                            }
                                            var v77 bool = v25
                                            _ = v76
                                            _ = v77
                                        } else {
                                            if ((!v25) && (true || (-9223372036854775808 == 32767))) {
                                                v1 += (v1 << (((^(v4))) & 63))
                                            } else {
                                                v78 := GenS2{F0: -32768, F1: (-(v27))}
                                                v4 = 18446744073709551615
                                                _ = v78
                                            }
                                            v20 = v20
                                            if v25 {
                                                var v79 uint16 = 65535
                                                switch v1 {
                                                case 127:
                                                    if len(v5) > 0 {
                                                        v5[0] = 2
                                                    }
                                                    v80 := GenS3{F0: (v27 + -1)}
                                                    v75 = 1e10
                                                    if ((fi21 * 1) > 32767) {
                                                        var v81 string = "hello"
                                                        _ = v81
                                                    }
                                                    _ = v80
                                                case -591:
                                                    fmt.Println(v20, v1, v79)
                                                }
                                                switch "" {
                                                case "\n\t":
                                                    _ = len(v3)
                                                }
                                                v27 = (v27 % (((v27 / ((v27) | 1))) | 1))
                                                _ = v79
                                            }
                                        }
                                        _ = v75
                                    }
                                    v4 -= v4
                                }
                                if (v25 && v25) {
                                    switch v1 {
                                    case -128:
                                        v82 := GenS8{F0: 42, F1: int(v4), F2: (fi21 + (fi23 / ((fi21) | 1)))}
                                        _ = v82
                                    case -24:
                                        switch fi21 {
                                        case 572:
                                            v6 += (v6 & -1)
                                            if v25 {
                                                v83 := GenS9{F0: -1}
                                                var v84 int32 = 2147483647
                                                _ = v83
                                                _ = v84
                                            }
                                            var v85 uint16 = 255
                                            _ = v85
                                        case 32768:
                                            v86 := GenS12{F0: 116}
                                            _ = v86
                                        case 783:
                                            _ = len(v5)
                                        }
                                    case 151:
                                        switch "a" {
                                        case "hello":
                                            v1 = -1
                                        case "\n\t":
                                            v1 -= (^(fi21))
                                        }
                                    default:
                                        switch int(v6) {
                                        case 444:
                                            var v87 int16 = 1
                                            _ = v87
                                        case 127:
                                            if false {
                                                v25 = (v25 || ((v25 || (!("a" >= "\n\t"))) || ((v20 + v20) == (v20 * v20))))
                                            } else {
                                                v88 := GenS10{F0: 2}
                                                v25 = (v18 != (^(v18)))
                                                var v89 int8 = v20
                                                _ = v88
                                                _ = v89
                                            }
                                        }
                                    }
                                } else {
                                    v1 += (fi21 | fi23)
                                }
                                if (!v25) {
                                    switch ("a" + "foo bar") {
                                    case "":
                                        v4 += (v4 << (((^(v4))) & 63))
                                    case "foo bar":
                                        var v90 float32 = -2.25
                                        var v91 uint8 = 68
                                        _ = v90
                                        _ = v91
                                    }
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                }
                                switch (fi23 | (fi23 & (fi21 - fi23))) {
                                case 203:
                                    v92 := GenS3{F0: v27}
                                    _ = v92
                                }
                            }
                            fmt.Println(v27, v1)
                            _ = v63
                        }
                        v93 := GenS13{F0: v6, F1: v27}
                        fmt.Println(v1, v25, v20)
                        _ = v27
                        _ = v93
                    }
                    fmt.Println(v20, fi23, v1)
                    switch -1 {
                    case 0:
                        v25 = (-176 > 42)
                    case -457:
                        v6 = int64(v18)
                    case 4294967295:
                        v18 = (^(v18))
                    default:
                        var v94 int16 = 32767
                        if true {
                            var v95 int = (fi23 & fi23)
                            v1 = (v1 * (fi23 * -875))
                            _ = v95
                        }
                        _ = v94
                    }
                }
            } else {
                v96 := GenS13{F0: (v6 - (^(v6))), F1: -691}
                v18 = (v18 & 128)
                if v25 {
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    var v97 uint = 939
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    _ = v97
                }
                _ = v96
            }
            _ = fi23
            _ = v24
            _ = v25
        }
        _ = fi21
    }
    fmt.Println(v1, v4)
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    var v98 int = v1
    var v99 uint16 = 255
    for fi100 := 0; fi100 < 20; fi100++ {
        v98 -= -1
        var v101 uint = (v4 | 42)
        _ = fi100
        _ = v101
    }
    v102 := GenS9{F0: 25}
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v98:", v98)
    fmt.Println("v99:", v99)
    _ = v1
    _ = v102
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v98
    _ = v99
}
