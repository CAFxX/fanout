package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 bool; F1 string; F2 uint16 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 int8; F1 uint16 }
type GenS7 struct { F0 int8; F1 int64 }
type GenS8 struct { F0 int; F1 int64 }
type GenS9 struct { F0 string; F1 int; F2 bool }
func genFn0() uint8 {
    var v6 int8 = 127
    var v7 string = "hello"
    var v8 uint = 2
    _ = v6
    _ = v7
    _ = v8
    return uint8(v8)
}
func genFn1(p0 bool, p1 int16) uint8 {
    p1 = p1
    if false {
        var v9 float32 = 0.1
        _ = v9
    } else {
        var v10 uint8 = 5
        p0 = ((42 == 4294967295) && (!true))
        for fi11 := 0; fi11 < 20; fi11++ {
            if (0 <= 2147483647) {
                if (65535 < 65535) {
                    if true {
                        p0 = (252 < 9223372036854775807)
                    }
                    if p0 {
                        var v12 int16 = (-(p1))
                        _ = v12
                    } else {
                        v13 := GenS1{F0: 42}
                        v14 := GenS1{F0: v10}
                        p1 = p1
                        _ = v13
                        _ = v14
                    }
                    v10 = genFn0()
                }
                if (116 != 42) {
                    var v15 int64 = -883
                    var v16 int64 = (v15 | (v15 ^ (v15 - (^(v15)))))
                    fmt.Println(v10)
                    _ = v15
                    _ = v16
                } else {
                    var v17 bool = (!(int16(v10) > -71))
                    p0 = true
                    _ = v17
                }
                for fi18 := 0; fi18 < 3; fi18++ {
                    v19 := GenS2{F0: int16(fi11)}
                    _ = fi18
                    _ = v19
                }
            }
            _ = fi11
        }
        _ = v10
    }
    _ = p0
    _ = p1
    return 1
}

func main() {
    v1 := []uint32{590, 2147483648}
    v2 := make(map[uint32]string)
    var v3 float32 = 3.14159
    v4 := []uint32{271, 0}
    var v5 int64 = 9223372036854775807
    switch 468 {
    case -2147483648:
        fmt.Println(v3, v5)
    case -9223372036854775808:
        for fi20 := 0; fi20 < 3; fi20++ {
            for fi21 := 0; fi21 < 14; fi21++ {
                switch 1 {
                case 42:
                    v3 = 1.5
                }
                v5 = 1
                fmt.Println(v5, fi21)
                _ = fi21
            }
            _ = fi20
        }
    case -317:
        fmt.Println(v5, v3)
    }
    for fi22 := 0; fi22 < 8; fi22++ {
        v5 = (v5 + (-(v5)))
        _ = fi22
    }
    fmt.Println(v3, v5)
    v23 := GenS3{F0: 0}
    if len(v4) > 0 {
        v4[0] = 2
    }
    v24 := GenS3{F0: 886}
    if false {
        v25 := GenS4{F0: (!(-2147483648 >= 128)), F1: "", F2: 358}
        for fi26 := 0; fi26 < 10; fi26++ {
            switch "" {
            case "foo bar":
                var v27 string = ("a" + "hello")
                _ = v27
            case "\n\t":
                var v28 bool = ((("\n\t" + "a") + "hello") > "hello")
                _ = v28
            case "hello":
                delete(v2, 747)
            }
            v3 += 1.5
            _ = fi26
        }
        if false {
            var v29 int = -497
            v30 := GenS2{F0: -1}
            var v31 bool = true
            _ = v29
            _ = v30
            _ = v31
        } else {
            for fi32 := 0; fi32 < 11; fi32++ {
                var v33 int = fi32
                var v34 float32 = v3
                _ = fi32
                _ = v33
                _ = v34
            }
            v3 += v3
            for fi35 := 0; fi35 < 12; fi35++ {
                v5 = -104
                if (237 <= 2147483648) {
                    fmt.Println(v5, v5)
                    var v36 uint = 2
                    v36 += v36
                    _ = v36
                } else {
                    v1 = append(v1, 2)
                    v37 := GenS5{F0: 106}
                    _ = v37
                }
                v5 = 0
                _ = fi35
            }
        }
        _ = v25
    }
    if false {
        var v38 float64 = -2.25
        v38 = float64((v3 / v3))
        _ = v38
    } else {
        var v39 int = 642
        if (true && (-1 == 98)) {
            v3 -= ((float32(3.14159) * v3) * (float32((3.14159 / 1.5)) + v3))
        } else {
            if false {
                if true {
                    var v40 uint16 = 255
                    switch v39 {
                    case -2147483648:
                        var v41 int8 = 1
                        _ = v41
                    default:
                        v42 := GenS2{F0: -1}
                        _ = v42
                    }
                    if true {
                        v43 := GenS1{F0: 128}
                        v44 := GenS6{F0: 0, F1: (v40 * v40)}
                        _ = v43
                        _ = v44
                    }
                    _ = v40
                } else {
                    if false {
                        var v45 uint64 = 105
                        _ = v45
                    } else {
                        if true {
                            if false {
                                var v46 int = 165
                                if true {
                                    var v47 bool = ((("\n\t" + "") + "\n\t") < "\n\t")
                                    v48 := GenS5{F0: 1}
                                    if (-1 != 206) {
                                        v49 := GenS1{F0: 1}
                                        var v50 int32 = 42
                                        if true {
                                            var v51 int16 = -32768
                                            _ = v51
                                        }
                                        _ = v49
                                        _ = v50
                                    }
                                    _ = v47
                                    _ = v48
                                } else {
                                    var v52 bool = ((42 != -86) && (genFn0() >= 134))
                                    v53 := GenS4{F0: (816 <= 65535), F1: ("foo bar" + "hello"), F2: 160}
                                    _ = v52
                                    _ = v53
                                }
                                if ((false || (false || true)) && true) {
                                    if false {
                                        for fi54 := 0; fi54 < 4; fi54++ {
                                            for fi55 := 0; fi55 < 12; fi55++ {
                                                var v56 string = ("a" + (("a" + "hello") + ""))
                                                v57 := GenS7{F0: 0, F1: 691}
                                                _ = fi55
                                                _ = v56
                                                _ = v57
                                            }
                                            _ = fi54
                                        }
                                        if false {
                                            v46 = v39
                                        } else {
                                            var v58 float32 = v3
                                            _ = v58
                                        }
                                    } else {
                                        if false {
                                            v59 := GenS8{F0: (v39 | v39), F1: v5}
                                            _ = v59
                                        } else {
                                            if true {
                                                if ((v39 << ((42) & 63)) != (v39 | (-(v39)))) {
                                                    v60 := GenS2{F0: 0}
                                                    v61 := GenS5{F0: 255}
                                                    _ = v60
                                                    _ = v61
                                                } else {
                                                    fmt.Println(v46, v39, v46)
                                                    v3 = v3
                                                }
                                            }
                                            fmt.Println(v5, v46, v39)
                                        }
                                    }
                                    var v62 int64 = v5
                                    v63 := GenS6{F0: -128, F1: 32767}
                                    _ = v62
                                    _ = v63
                                }
                                _ = v46
                            } else {
                                var v64 uint32 = 706
                                _ = v64
                            }
                        } else {
                            v65 := GenS8{F0: (v39 | v39), F1: v5}
                            v5 += (v5 % ((v5) | 1))
                            var v66 int = v39
                            _ = v65
                            _ = v66
                        }
                        for fi67 := 0; fi67 < 12; fi67++ {
                            fmt.Println(fi67, v39)
                            var v68 int32 = 2147483647
                            _ = fi67
                            _ = v68
                        }
                    }
                }
                v39 += (v39 + -81)
                switch "" {
                case "hello":
                    for fi69 := 0; fi69 < 14; fi69++ {
                        if (false || true) {
                            fmt.Println(v39)
                            var v70 uint16 = 1
                            v39 -= 539
                            _ = v70
                        }
                        v5 = v5
                        _ = fi69
                    }
                default:
                    v5 -= 820
                }
            }
            for fi71 := 0; fi71 < 7; fi71++ {
                v5 = 1
                _ = fi71
            }
        }
        if false {
            var v72 int16 = 550
            v2[0] = "a"
            switch (v39 + int(v5)) {
            case -1:
                var v73 int8 = 1
                _ = v73
            case 546:
                for fi74 := 0; fi74 < 13; fi74++ {
                    switch (fi74 - 0) {
                    case -420:
                        var v75 string = ""
                        _ = v75
                    case -9223372036854775808:
                        switch (-(fi74)) {
                        case 128:
                            if true {
                                v2[2147483647] = "\n\t"
                            }
                        case -748:
                            if (!true) {
                                switch (("\n\t" + "foo bar") + ("a" + "hello")) {
                                case "\n\t":
                                    var v76 uint32 = 136
                                    _ = v76
                                case "a":
                                    v72 += v72
                                case "hello":
                                    for fi77 := 0; fi77 < 12; fi77++ {
                                        v3 = v3
                                        _ = fi77
                                    }
                                }
                            } else {
                                v5 = (v5 >> ((978) & 63))
                                v39 = (v39 % ((fi74) | 1))
                                v78 := GenS8{F0: 32767, F1: 1}
                                _ = v78
                            }
                        default:
                            var v79 int64 = (v5 - v5)
                            _ = v79
                        }
                    case 57:
                        delete(v2, 577)
                    }
                    switch (-(v39)) {
                    case -585:
                        v80 := v2[1]
                        _ = v80
                    }
                    _ = fi74
                }
            }
            _ = v72
        }
        _ = v39
    }
    for fi81 := 0; fi81 < 18; fi81++ {
        switch 128 {
        case 675:
            v3 = 0.1
        }
        v5 = (v5 % ((v5) | 1))
        _ = fi81
    }
    var v82 bool = (!((false && false) || (((((0.0 * (-2.25 / ((1.5 + 0.0) + 1e10))) / 0.1) * 0.0) * 1.5) > 1e10)))
    for fi83 := 0; fi83 < 9; fi83++ {
        var v84 uint8 = 223
        v85 := GenS9{F0: "foo bar", F1: fi83, F2: v82}
        _ = fi83
        _ = v84
        _ = v85
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v82:", v82)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v82
}
