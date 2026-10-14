package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint64; F2 uint16 }
type GenS2 struct { F0 uint32; F1 int32; F2 int }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint64; F1 int8; F2 int }
type GenS5 struct { F0 int16; F1 string; F2 int8 }
type GenS6 struct { F0 int32; F1 uint; F2 uint16 }
type GenS7 struct { F0 int8 }
type GenS8 struct { F0 uint8; F1 int8; F2 uint32 }
type GenS9 struct { F0 int64; F1 bool }
type GenS10 struct { F0 int16; F1 uint32 }
type GenS11 struct { F0 string }
type GenS12 struct { F0 int16 }
type GenS13 struct { F0 int }
type GenS14 struct { F0 int; F1 int; F2 uint64 }
func genFn0(p0 uint64, p1 uint64) uint64 {
    p0 = (p1 << ((255) & 63))
    v7 := GenS1{F0: 1, F1: p0, F2: 62}
    p1 += p0
    switch "hello" {
    case "":
        if (0 >= 873) {
            p1 -= (p1 % (((p1 / (((p0 % ((p0) | 1))) | 1))) | 1))
            fmt.Println(p0, p1)
            fmt.Println(p1)
        }
    case "\n\t":
        switch -201 {
        case -2147483648:
            if true {
                fmt.Println(p0, p1)
            } else {
                p0 += (p1 - p0)
            }
        case 42:
            p0 -= 18446744073709551615
        }
    default:
        if false {
            p0 -= (p0 / ((243) | 1))
        } else {
            v8 := []int{9223372036854775807}
            p1 = (p1 >> ((1) & 63))
            _ = v8
        }
    }
    _ = p0
    _ = p1
    _ = v7
    return p1
}
func genFn1(p0 int16, p1 string, p2 uint8) uint32 {
    var v9 float32 = -2.25
    for fi10 := 0; fi10 < 13; fi10++ {
        var v11 string = p1
        var v12 string = (("a" + (p1 + "")) + (v11 + p1))
        p2 -= p2
        _ = fi10
        _ = v11
        _ = v12
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return 42
}

func main() {
    v1 := []string{"foo bar", "\n\t", "hello"}
    var v2 uint64 = 42
    v3 := []int{0}
    var v4 uint16 = 255
    var v5 uint8 = 127
    var v6 uint8 = v5
    for fi13 := 0; fi13 < 12; fi13++ {
        switch (fi13 >> ((2) & 63)) {
        case 128:
            switch "a" {
            case "":
                _ = len(v1)
            case "hello":
                for fi14 := 0; fi14 < 18; fi14++ {
                    if (1.5 != 1.5) {
                        v15 := GenS2{F0: 1, F1: 42, F2: (^(fi13))}
                        v4 = 1
                        _ = v15
                    }
                    if (!false) {
                        v2 += 1
                        if true {
                            fmt.Println(v6)
                            if len(v3) > 0 {
                                v3[0] = -9223372036854775808
                            }
                            var v16 uint32 = 42
                            _ = v16
                        }
                    }
                    v17 := GenS3{F0: 4294967295}
                    _ = fi14
                    _ = v17
                }
            case "foo bar":
                if (false || true) {
                    if true {
                        var v18 int = -1
                        v6 = v5
                        var v19 int8 = -38
                        _ = v18
                        _ = v19
                    } else {
                        var v20 uint32 = 201
                        if len(v1) > 0 {
                            v1[0] = ("foo bar" + "foo bar")
                        }
                        switch "foo bar" {
                        case "hello":
                            if (v4 > (v4 * 32767)) {
                                for fi21 := 0; fi21 < 2; fi21++ {
                                    var v22 uint64 = uint64(v4)
                                    _ = fi21
                                    _ = v22
                                }
                            }
                        case "":
                            for fi23 := 0; fi23 < 6; fi23++ {
                                v1 = append(v1, "foo bar")
                                if ((v2 % (((v2 / ((genFn0(v2, (v2 & genFn0((v2 << ((65535) & 63)), (v2 - (^(v2))))))) | 1))) | 1)) <= (v2 ^ 903)) {
                                    v24 := GenS4{F0: genFn0(v2, genFn0(uint64(v5), uint64(v4))), F1: 127, F2: fi23}
                                    switch "\n\t" {
                                    case "\n\t":
                                        if len(v3) > 0 {
                                            v3[0] = fi23
                                        }
                                        if (!true) {
                                            v4 = 255
                                            if len(v3) >= 2 {
                                                v3 = v3[1:len(v3)]
                                            }
                                            v25 := GenS5{F0: 0, F1: "hello", F2: 1}
                                            _ = v25
                                        } else {
                                            if (!(42 > 42)) {
                                                switch ("" + "\n\t") {
                                                case "":
                                                    var v26 uint8 = 255
                                                    _ = v26
                                                case "hello":
                                                    v27 := GenS5{F0: -1, F1: ("a" + ""), F2: int8(v20)}
                                                    _ = v27
                                                case "foo bar":
                                                    fmt.Println(v20, v4)
                                                    fmt.Println(v6, v4)
                                                    var v28 uint64 = (v2 * 18446744073709551615)
                                                    _ = v28
                                                default:
                                                    v2 = uint64(v4)
                                                }
                                                v5 -= 127
                                                if len(v1) >= 2 {
                                                    v1 = v1[1:len(v1)]
                                                }
                                            } else {
                                                switch fi13 {
                                                case 127:
                                                    var v29 float32 = 1e10
                                                    _ = v29
                                                case -849:
                                                    v20 = 42
                                                case -1:
                                                    v20 = (v20 + (v20 / ((genFn1(154, "a", (v6 * v5))) | 1)))
                                                default:
                                                    var v30 string = ("foo bar" + "hello")
                                                    _ = v30
                                                }
                                                var v31 int64 = -1
                                                _ = v31
                                            }
                                            switch "hello" {
                                            case "":
                                                switch fi23 {
                                                case -524:
                                                    if true {
                                                        switch -2147483648 {
                                                        case 32768:
                                                            var v32 int8 = -58
                                                            _ = v32
                                                        case 128:
                                                            var v33 uint = 42
                                                            _ = v33
                                                        }
                                                    }
                                                case 127:
                                                    fmt.Println(fi23, v6, fi23)
                                                case 337:
                                                    fmt.Println(v6, v20, v5)
                                                default:
                                                    switch "foo bar" {
                                                    case "":
                                                        if (1 >= -632) {
                                                            switch 0 {
                                                            case 42:
                                                                v4 = v4
                                                            case 0:
                                                                v34 := GenS6{F0: 1, F1: 4294967295, F2: v4}
                                                                _ = v34
                                                            case 106:
                                                                v35 := GenS6{F0: -320, F1: 477, F2: v4}
                                                                _ = v35
                                                            default:
                                                                v20 = 65535
                                                            }
                                                            var v36 uint16 = (v4 << ((1) & 63))
                                                            var v37 uint64 = (v2 + v2)
                                                            _ = len(v1)
                                                            v38 := GenS7{F0: -4}
                                                            var v39 int32 = 994
                                                            fmt.Println(fi23, v37)
                                                            var v40 uint16 = v4
                                                            _ = v36
                                                            _ = v37
                                                            _ = v38
                                                            _ = v39
                                                            _ = v40
                                                        }
                                                    case "foo bar":
                                                        v6 = v6
                                                    default:
                                                        var v41 bool = ((v4 << ((uint(v5)) & 63)) <= (v4 + uint16(v20)))
                                                        var v42 uint8 = v6
                                                        var v43 uint64 = 2
                                                        _ = v41
                                                        _ = v42
                                                        _ = v43
                                                    }
                                                }
                                            case "a":
                                                v3 = append(v3, (fi23 ^ (fi13 - (fi23 / ((fi23) | 1)))))
                                            }
                                        }
                                    case "hello":
                                        v6 += 2
                                    case "a":
                                        if (!true) {
                                            v44 := GenS8{F0: (v5 % (((v6 | v5)) | 1)), F1: 100, F2: (v20 % ((v20) | 1))}
                                            _ = v44
                                        } else {
                                            fmt.Println(v6, v5, v2)
                                            v20 = genFn1(0, "foo bar", v5)
                                        }
                                    default:
                                        if (false || (961 < 376)) {
                                            var v45 bool = (v4 > (v4 + (^(v4))))
                                            v4 += (^(v4))
                                            var v46 uint8 = v5
                                            _ = v45
                                            _ = v46
                                        }
                                    }
                                    var v47 int64 = -1
                                    v48 := GenS8{F0: v6, F1: 100, F2: 2147483647}
                                    v47 += v47
                                    _ = v24
                                    _ = v47
                                    _ = v48
                                }
                                var v49 uint32 = (v20 * 2147483647)
                                _ = fi23
                                _ = v49
                            }
                        default:
                            v50 := GenS5{F0: -1, F1: "", F2: 0}
                            _ = v50
                        }
                        _ = v20
                    }
                    v6 = 0
                } else {
                    if false {
                        for fi51 := 0; fi51 < 17; fi51++ {
                            var v52 int64 = 517
                            _ = fi51
                            _ = v52
                        }
                    } else {
                        for fi53 := 0; fi53 < 12; fi53++ {
                            if (v4 > 0) {
                                if (127 == 102) {
                                    v54 := GenS9{F0: -9223372036854775808, F1: (int(v4) <= -9223372036854775808)}
                                    if (false && ("" == "foo bar")) {
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                        if true {
                                            var v55 int64 = -9223372036854775808
                                            var v56 int = (fi13 + fi13)
                                            _ = v55
                                            _ = v56
                                        } else {
                                            var v57 uint8 = (v5 << ((155) & 63))
                                            _ = v57
                                        }
                                        v58 := GenS10{F0: 378, F1: genFn1(42, "foo bar", v5)}
                                        _ = v58
                                    }
                                    _ = v54
                                }
                                var v59 uint8 = v6
                                v6 -= 1
                                v4 = v4
                                var v60 int32 = 42
                                _ = v59
                                _ = v60
                            }
                            _ = fi53
                        }
                        if false {
                            for fi61 := 0; fi61 < 16; fi61++ {
                                _ = len(v1)
                                switch "foo bar" {
                                case "\n\t":
                                    var v62 float64 = (3.14159 * float64(0.1))
                                    _ = v62
                                case "foo bar":
                                    var v63 int64 = 302
                                    _ = v63
                                }
                                var v64 uint8 = 34
                                _ = fi61
                                _ = v64
                            }
                            if false {
                                v65 := GenS11{F0: "hello"}
                                v66 := GenS10{F0: 895, F1: 695}
                                _ = v65
                                _ = v66
                            } else {
                                if (false && (!false)) {
                                    v2 = (v2 | (^(v2)))
                                    for fi67 := 0; fi67 < 15; fi67++ {
                                        v6 = (v6 & uint8(v4))
                                        var v68 int8 = 100
                                        v6 = v6
                                        var v69 uint32 = 31
                                        _ = fi67
                                        _ = v68
                                        _ = v69
                                    }
                                }
                                var v70 uint64 = genFn0(genFn0(uint64(v5), 0), v2)
                                _ = v70
                            }
                        }
                        v71 := GenS12{F0: int16(v6)}
                        _ = v71
                    }
                    v72 := GenS13{F0: fi13}
                    fmt.Println(v2)
                    _ = v72
                }
            }
        case 4294967295:
            v73 := GenS14{F0: fi13, F1: fi13, F2: genFn0(v2, v2)}
            _ = v73
        }
        _ = fi13
    }
    if false {
        v6 -= (v5 & v5)
    }
    switch 1 {
    case 0:
        v74 := GenS6{F0: 1, F1: 930, F2: v4}
        _ = v74
    default:
        var v75 uint32 = 42
        _ = v75
    }
    var v76 string = "hello"
    if false {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        if (32768 == -267) {
            if true {
                for fi77 := 0; fi77 < 7; fi77++ {
                    v6 = 255
                    v4 = 828
                    _ = fi77
                }
                var v78 float64 = 1.5
                _ = v78
            } else {
                v79 := GenS7{F0: int8(v4)}
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                var v80 string = v76
                _ = v79
                _ = v80
            }
        }
        var v81 float64 = float64(0.0)
        _ = v81
    }
    if true {
        _ = len(v1)
        v5 = v5
        if true {
            fmt.Println(v4, v2)
        } else {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            if len(v1) > 0 {
                v1[0] = v76
            }
            var v82 int = 127
            _ = v82
        }
    }
    var v83 int32 = -853
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    if (false || (v6 > 127)) {
        if len(v3) > 0 {
            v3[0] = 4294967295
        }
        var v84 uint8 = 42
        _ = v84
    }
    _ = len(v3)
    v83 = v83
    for fi85 := 0; fi85 < 12; fi85++ {
        v76 = v76
        for fi86 := 0; fi86 < 20; fi86++ {
            var v87 uint32 = genFn1(1, v76, v5)
            if (9223372036854775807 == 738) {
                fmt.Println(fi86, v6, v6)
                if (true && true) {
                    fmt.Println(fi86, v5)
                } else {
                    v87 -= uint32(v2)
                    v88 := GenS14{F0: -471, F1: fi85, F2: v2}
                    _ = v88
                }
                fmt.Println(v4)
            } else {
                v89 := GenS13{F0: fi86}
                _ = v89
            }
            fmt.Println(v76, v83, v83)
            _ = fi86
            _ = v87
        }
        _ = fi85
    }
    var v90 string = "hello"
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v76:", v76)
    fmt.Println("v83:", v83)
    fmt.Println("v90:", v90)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v76
    _ = v83
    _ = v90
}
