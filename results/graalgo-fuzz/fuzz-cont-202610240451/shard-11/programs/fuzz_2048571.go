package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint32; F2 uint }
type GenS2 struct { F0 int64; F1 int8 }
type GenS3 struct { F0 uint16; F1 uint16 }
type GenS4 struct { F0 uint; F1 uint64; F2 uint32 }
type GenS5 struct { F0 uint8; F1 int8 }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 uint8; F1 int8 }
type GenS8 struct { F0 uint64; F1 uint; F2 uint }
type GenS9 struct { F0 uint32; F1 uint; F2 bool }
type GenS10 struct { F0 uint8; F1 uint8 }
type GenS11 struct { F0 uint16; F1 int8 }
func genFn0(p0 bool, p1 int8, p2 bool) int64 {
    if true {
        fmt.Println(p0, p1, p2)
    } else {
        if p2 {
            var v7 int = 128
            _ = v7
        }
        for fi8 := 0; fi8 < 10; fi8++ {
            for fi9 := 0; fi9 < 5; fi9++ {
                p0 = p0
                v10 := []int{(fi9 & (^(fi8)))}
                if false {
                    p1 -= 100
                }
                _ = fi9
                _ = v10
            }
            _ = fi8
        }
    }
    p0 = p0
    p0 = p2
    _ = p0
    _ = p1
    _ = p2
    return 9223372036854775807
}
func genFn1(p0 uint64, p1 string) uint {
    v11 := GenS1{F0: 255, F1: 219, F2: 255}
    for fi12 := 0; fi12 < 10; fi12++ {
        var v13 bool = false
        p1 = p1
        p0 += (p0 - p0)
        _ = fi12
        _ = v13
    }
    v14 := []int{}
    _ = len(v14)
    _ = p0
    _ = p1
    _ = v11
    _ = v14
    return 18446744073709551615
}

func main() {
    var v1 int = 4294967295
    var v2 bool = ((32767 < 399) || (!((779 != 65535) || false)))
    var v3 uint32 = 65535
    var v4 uint32 = v3
    v5 := []string{"hello", "", ""}
    var v6 int8 = 100
    if v2 {
        var v15 int64 = -140
        _ = v15
    }
    v4 -= 2147483647
    v6 = -1
    var v16 string = ("" + "\n\t")
    v17 := GenS1{F0: 633, F1: v4, F2: genFn1(4294967295, v16)}
    for fi18 := 0; fi18 < 11; fi18++ {
        var v19 int32 = 210
        if v2 {
            fmt.Println(v1, fi18)
            if (v2 || v2) {
                var v20 int8 = 111
                v6 = 42
                switch v16 {
                case "":
                    v21 := GenS1{F0: 0, F1: 0, F2: 65535}
                    _ = v21
                case "a":
                    var v22 int32 = (v19 + 289)
                    _ = v22
                }
                _ = v20
            } else {
                switch v1 {
                case 9223372036854775807:
                    if v2 {
                        if v2 {
                            v5 = append(v5, v16)
                        } else {
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                        }
                        var v23 int16 = 1000
                        _ = v23
                    } else {
                        _ = len(v5)
                        if true {
                            var v24 uint32 = v3
                            if len(v5) > 0 {
                                v5[0] = ("\n\t" + v16)
                            }
                            _ = v24
                        } else {
                            var v25 int = (^(v1))
                            var v26 int = 127
                            v1 -= (fi18 + v26)
                            _ = v25
                            _ = v26
                        }
                        for fi27 := 0; fi27 < 6; fi27++ {
                            var v28 uint = 42
                            var v29 uint16 = 0
                            var v30 uint8 = 42
                            if ((9223372036854775807 <= 807) && v2) {
                                if v2 {
                                    if v2 {
                                        v5 = append(v5, v16)
                                        if len(v5) > 0 {
                                            v5[0] = v16
                                        }
                                        if len(v5) > 0 {
                                            v5[0] = "foo bar"
                                        }
                                    } else {
                                        v1 = (-(fi27))
                                    }
                                    v31 := GenS1{F0: (v28 + (v28 | v28)), F1: v3, F2: uint(v29)}
                                    _ = v31
                                } else {
                                    switch "\n\t" {
                                    case "hello":
                                        fmt.Println(v1, v29)
                                    case "foo bar":
                                        _ = len(v5)
                                        if (true || ((76 >= fi27) && false)) {
                                            v32 := GenS1{F0: (v28 + genFn1(2, v16)), F1: (v4 / (((v3 >> (((v28 - uint(v30))) & 63))) | 1)), F2: v28}
                                            _ = v32
                                        }
                                        var v33 int64 = -223
                                        _ = v33
                                    default:
                                        fmt.Println(v1, v1)
                                    }
                                    switch 42 {
                                    case 390:
                                        v30 = (v30 + 21)
                                    default:
                                        if v2 {
                                            v5 = append(v5, "hello")
                                        }
                                    }
                                    if false {
                                        if (true || v2) {
                                            v34 := GenS2{F0: int64(v28), F1: v6}
                                            v35 := GenS2{F0: 872, F1: v6}
                                            v36 := GenS3{F0: v29, F1: 42}
                                            var v37 int64 = -349
                                            _ = v34
                                            _ = v35
                                            _ = v36
                                            _ = v37
                                        }
                                        if (988 != 4294967295) {
                                            v16 = ""
                                            v1 -= v1
                                        }
                                        if len(v5) >= 2 {
                                            v5 = v5[1:len(v5)]
                                        }
                                    }
                                }
                                if v2 {
                                    var v38 int8 = v6
                                    v19 += (v19 - -762)
                                    _ = v38
                                } else {
                                    _ = len(v5)
                                    v39 := GenS2{F0: 9223372036854775807, F1: v6}
                                    var v40 uint = v28
                                    switch int(v19) {
                                    case 127:
                                        var v41 int64 = 882
                                        _ = v41
                                    case 9223372036854775807:
                                        v4 = 2
                                    }
                                    var v42 int64 = 0
                                    _ = v39
                                    _ = v40
                                    _ = v42
                                }
                                _ = len(v5)
                                if len(v5) > 0 {
                                    v5[0] = v16
                                }
                                v43 := GenS4{F0: genFn1(4294967295, (v16 + "hello")), F1: 2, F2: (v3 << ((0) & 63))}
                                _ = v43
                            }
                            _ = fi27
                            _ = v28
                            _ = v29
                            _ = v30
                        }
                    }
                }
                v19 = v19
                v44 := GenS5{F0: 139, F1: v6}
                _ = v44
            }
            v19 -= 894
        } else {
            var v45 int8 = (v6 + v6)
            fmt.Println(v6)
            if ((2 != 0) || false) {
                v5 = append(v5, v16)
                _ = len(v5)
                var v46 uint64 = 2
                _ = v46
            }
            _ = v45
        }
        _ = fi18
        _ = v19
    }
    var v47 uint16 = 254
    fmt.Println(v4, v1, v47)
    v48 := GenS1{F0: 42, F1: 2147483648, F2: 65535}
    if (((!v2) && v2) || false) {
        if (!(((0.0 + 0.0) != 0.0) || ((121 == 255) && (v47 > (v47 >> ((255) & 63)))))) {
            for fi49 := 0; fi49 < 19; fi49++ {
                _ = len(v5)
                _ = fi49
            }
        } else {
            if ((!(v1 == (v1 >> ((519) & 63)))) && (v16 != v16)) {
                v5 = append(v5, "hello")
                if v2 {
                    v50 := GenS5{F0: 1, F1: v6}
                    v4 += v4
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    _ = v50
                } else {
                    fmt.Println(v16)
                    switch ("foo bar" + v16) {
                    case "foo bar":
                        for fi51 := 0; fi51 < 18; fi51++ {
                            var v52 float32 = float32(3.14159)
                            switch 780 {
                            case -1:
                                fmt.Println(v47, fi51, v16)
                            case 269:
                                for fi53 := 0; fi53 < 12; fi53++ {
                                    v54 := GenS1{F0: 18446744073709551615, F1: 4294967295, F2: 18446744073709551615}
                                    v5 = append(v5, "")
                                    v1 -= 368
                                    _ = fi53
                                    _ = v54
                                }
                            case -2147483648:
                                if v2 {
                                    var v55 uint16 = (v47 >> ((genFn1(18446744073709551615, (v16 + v16))) & 63))
                                    v56 := GenS6{F0: 277}
                                    _ = v55
                                    _ = v56
                                }
                            }
                            v57 := GenS4{F0: 42, F1: uint64(v47), F2: v3}
                            _ = fi51
                            _ = v52
                            _ = v57
                        }
                    default:
                        var v58 float32 = ((0.1 * (1e10 * 0.1)) - 1.5)
                        _ = v58
                    }
                }
                _ = len(v5)
            } else {
                fmt.Println(v3, v4)
                fmt.Println(v47, v6)
            }
            var v59 uint = 65535
            if true {
                var v60 float32 = 0.0
                fmt.Println(v3, v4)
                _ = v60
            } else {
                for fi61 := 0; fi61 < 5; fi61++ {
                    v62 := GenS7{F0: 0, F1: 100}
                    var v63 bool = (!(3.14159 != (3.14159 * (1e10 - float64(((3.14159 * (float32(0.1) * 1e10)) * 1.5))))))
                    if false {
                        if true {
                            for fi64 := 0; fi64 < 2; fi64++ {
                                v65 := GenS8{F0: 694, F1: (v59 * genFn1(503, "")), F2: (^(v59))}
                                if (845 <= 18446744073709551615) {
                                    v66 := GenS9{F0: (v4 + v3), F1: (v59 | uint(v47)), F2: v63}
                                    switch (v1 | v1) {
                                    case 338:
                                        _ = len(v5)
                                    case 394:
                                        fmt.Println(v1)
                                    }
                                    _ = v66
                                }
                                v3 -= 126
                                _ = fi64
                                _ = v65
                            }
                            if v63 {
                                switch v16 {
                                case "\n\t":
                                    v59 -= v59
                                }
                            } else {
                                v6 = v6
                                var v67 int32 = 0
                                _ = v67
                            }
                            v3 += 2147483647
                        } else {
                            switch v16 {
                            case "":
                                v63 = ((v4 | (v3 & v4)) < 2)
                            case "hello":
                                v5 = append(v5, "foo bar")
                            }
                            v1 -= (v1 << (((v59 * v59)) & 63))
                        }
                    }
                    _ = fi61
                    _ = v62
                    _ = v63
                }
                v47 -= (^(v47))
            }
            _ = v59
        }
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    switch v16 {
    case "":
        if v2 {
            var v68 uint64 = 733
            _ = v68
        } else {
            var v69 string = v16
            var v70 uint32 = 65535
            v47 -= v47
            _ = v69
            _ = v70
        }
    case "\n\t":
        var v71 float32 = 0.0
        _ = v71
    case "a":
        var v72 uint16 = v47
        _ = v72
    default:
        if ((3.14159 + 1e10) > float64(((0.0 * (0.1 * (((((float32((0.0 / 1e10)) / (1.5 / 0.1)) - 0.1) - ((0.0 + (-2.25 - 3.14159)) + ((1.5 * ((-2.25 / (1e10 - 3.14159)) / 1.0)) / float32(1.5)))) * 3.14159) * float32(0.1)))) / 1.5))) {
            for fi73 := 0; fi73 < 10; fi73++ {
                for fi74 := 0; fi74 < 18; fi74++ {
                    v75 := GenS2{F0: 338, F1: (v6 >> ((735) & 63))}
                    _ = fi74
                    _ = v75
                }
                _ = fi73
            }
            switch "foo bar" {
            case "foo bar":
                var v76 uint = 255
                _ = v76
            case "\n\t":
                for fi77 := 0; fi77 < 13; fi77++ {
                    if false {
                        v4 = (^(v4))
                    }
                    for fi78 := 0; fi78 < 8; fi78++ {
                        var v79 int16 = 0
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                        _ = fi78
                        _ = v79
                    }
                    var v80 int32 = -1
                    _ = fi77
                    _ = v80
                }
            case "":
                v81 := GenS10{F0: 2, F1: 245}
                _ = v81
            default:
                if v2 {
                    var v82 uint64 = 453
                    if v2 {
                        if v2 {
                            v6 = -1
                            v5 = append(v5, "hello")
                            fmt.Println(v82, v1, v82)
                        } else {
                            for fi83 := 0; fi83 < 8; fi83++ {
                                var v84 uint16 = (^(v47))
                                _ = fi83
                                _ = v84
                            }
                            v85 := GenS5{F0: 1, F1: v6}
                            _ = v85
                        }
                        v6 += 95
                    } else {
                        if ("foo bar" == v16) {
                            if false {
                                var v86 uint32 = v3
                                v3 += v86
                                _ = v86
                            }
                            var v87 int = v1
                            if true {
                                var v88 uint16 = v47
                                v89 := GenS11{F0: v88, F1: 14}
                                if v2 {
                                    var v90 uint = 2
                                    _ = v90
                                }
                                _ = v88
                                _ = v89
                            } else {
                                var v91 uint64 = 18446744073709551615
                                for fi92 := 0; fi92 < 9; fi92++ {
                                    fmt.Println(fi92, fi92)
                                    _ = fi92
                                }
                                var v93 uint8 = 188
                                _ = v91
                                _ = v93
                            }
                            _ = v87
                        }
                    }
                    _ = v82
                }
            }
            var v94 string = v16
            _ = v94
        } else {
            v95 := GenS7{F0: 42, F1: v6}
            fmt.Println(v47, v2, v4)
            for fi96 := 0; fi96 < 15; fi96++ {
                switch ((((((v16 + v16) + v16) + v16) + "a") + v16) + v16) {
                case "hello":
                    fmt.Println(v6, v1, v1)
                case "":
                    for fi97 := 0; fi97 < 7; fi97++ {
                        v3 -= v3
                        v5 = append(v5, v16)
                        _ = fi97
                    }
                }
                _ = fi96
            }
            _ = v95
        }
    }
    v6 -= v6
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v47:", v47)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v47
    _ = v48
    _ = v5
    _ = v6
}
