package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint; F1 string; F2 int32 }
type GenS3 struct { F0 string; F1 uint32 }
type GenS4 struct { F0 int; F1 uint32 }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 int16; F1 uint64 }
func genFn0(p0 uint, p1 uint16, p2 int32) uint16 {
    v7 := GenS1{F0: int8(p1)}
    var v8 uint64 = 42
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 87
}

func main() {
    var v1 float64 = 3.14159
    var v2 uint64 = 1
    var v3 bool = (true || true)
    var v4 bool = ((!true) || (!(2 == 38)))
    var v5 int16 = 32767
    var v6 uint32 = 333
    v5 = (v5 % ((v5) | 1))
    v5 = (-(v5))
    v9 := GenS1{F0: -61}
    v10 := GenS2{F0: 2, F1: "", F2: 2147483647}
    v11 := []int{-2147483648, -862}
    var v12 bool = (uint16(v6) < 32767)
    v3 = v12
    switch ("" + "") {
    case "":
        var v13 int64 = 42
        _ = v13
    default:
        v4 = (v4 || v3)
    }
    var v14 int8 = 57
    switch ("a" + ("hello" + "a")) {
    case "a":
        var v15 uint64 = (v2 >> ((813) & 63))
        _ = v15
    case "\n\t":
        if len(v11) >= 2 {
            v11 = v11[1:len(v11)]
        }
    case "":
        if len(v11) >= 2 {
            v11 = v11[1:len(v11)]
        }
    default:
        var v16 uint64 = (^(v2))
        _ = v16
    }
    for fi17 := 0; fi17 < 7; fi17++ {
        var v18 uint32 = 42
        for fi19 := 0; fi19 < 16; fi19++ {
            v3 = ((!(2 != 801)) && ((-32768 != v5) || v3))
            v3 = (v2 > (^(v2)))
            _ = fi19
        }
        _ = fi17
        _ = v18
    }
    switch 42 {
    case 4294967295:
        for fi20 := 0; fi20 < 18; fi20++ {
            v1 = float64(0.1)
            if v12 {
                if (("" > "") || v3) {
                    switch ("foo bar" + ("foo bar" + (("a" + ("\n\t" + "")) + ("foo bar" + "")))) {
                    case "foo bar":
                        for fi21 := 0; fi21 < 1; fi21++ {
                            var v22 string = ""
                            if len(v11) >= 2 {
                                v11 = v11[1:len(v11)]
                            }
                            if v3 {
                                if (!((2 < 187) || v12)) {
                                    v12 = (9223372036854775807 > 0)
                                    var v23 int32 = 0
                                    if (3.14159 < 0.0) {
                                        var v24 int16 = v5
                                        _ = v24
                                    }
                                    v4 = (255 > 18446744073709551615)
                                    _ = v23
                                } else {
                                    var v25 int32 = -48
                                    _ = v25
                                }
                                v26 := GenS2{F0: 42, F1: (v22 + "foo bar"), F2: 1}
                                v3 = v3
                                _ = v26
                            }
                            v1 += float64((((3.14159 + -2.25) + 0.1) * 0.1))
                            _ = fi21
                            _ = v22
                        }
                    case "a":
                        switch ("\n\t" + "") {
                        case "":
                            fmt.Println(v6, fi20, v3)
                        case "a":
                            fmt.Println(v2, v12)
                        case "hello":
                            _ = len(v11)
                        default:
                            v3 = (false && v12)
                        }
                    default:
                        v2 = v2
                    }
                    for fi27 := 0; fi27 < 1; fi27++ {
                        v28 := GenS2{F0: 2, F1: ("" + "\n\t"), F2: -2147483648}
                        if len(v11) > 0 {
                            v11[0] = 127
                        }
                        _ = fi27
                        _ = v28
                    }
                    v29 := GenS2{F0: 876, F1: "a", F2: 2147483647}
                    _ = v29
                } else {
                    var v30 uint32 = (v6 << ((65535) & 63))
                    _ = v30
                }
            }
            _ = fi20
        }
    case 342:
        for fi31 := 0; fi31 < 15; fi31++ {
            for fi32 := 0; fi32 < 19; fi32++ {
                if v4 {
                    v33 := GenS2{F0: 255, F1: ("foo bar" + "foo bar"), F2: -59}
                    v12 = (true || (!((190 < 128) && v3)))
                    v3 = v4
                    _ = len(v11)
                    v12 = (0.1 == (1e10 / (0.0 + ((1e10 - 0.1) / -2.25))))
                    v14 = v14
                    _ = v33
                } else {
                    v2 += (^(v2))
                    v34 := GenS2{F0: 4294967295, F1: "", F2: -1}
                    v3 = false
                    var v35 uint32 = 4294967295
                    _ = v34
                    _ = v35
                }
                if len(v11) > 0 {
                    v11[0] = fi31
                }
                fmt.Println(v4, v4, v6)
                v36 := GenS2{F0: 53, F1: ("" + "\n\t"), F2: 670}
                v12 = (v6 == 42)
                _ = fi32
                _ = v36
            }
            for fi37 := 0; fi37 < 3; fi37++ {
                _ = len(v11)
                var v38 float64 = v1
                if ((!(v12 || true)) || (v14 < 100)) {
                    switch ("hello" + ("\n\t" + "foo bar")) {
                    case "a":
                        v39 := GenS3{F0: "foo bar", F1: (v6 / ((2147483648) | 1))}
                        _ = v39
                    case "hello":
                        if ((-(fi31)) <= fi37) {
                            if (false || ("" > "hello")) {
                                v14 = v14
                            }
                        }
                    case "":
                        if (-2.25 <= 0.1) {
                            var v40 float32 = ((3.14159 - 0.0) - float32(v38))
                            switch -761 {
                            case 32767:
                                if v3 {
                                    if len(v11) >= 2 {
                                        v11 = v11[1:len(v11)]
                                    }
                                } else {
                                    fmt.Println(v38, v38)
                                }
                            case 128:
                                switch "a" {
                                case "a":
                                    var v41 bool = v3
                                    var v42 string = ("hello" + ("foo bar" + ""))
                                    _ = v41
                                    _ = v42
                                case "\n\t":
                                    if ("a" == "foo bar") {
                                        if (v3 || false) {
                                            if ((^(v6)) == v6) {
                                                _ = len(v11)
                                                v3 = ((v5 / ((0) | 1)) <= v5)
                                                v5 -= v5
                                            }
                                            var v43 float64 = 3.14159
                                            switch (fi31 ^ fi31) {
                                            case 32768:
                                                if (!v12) {
                                                    v11 = append(v11, (fi31 * (fi31 + fi37)))
                                                }
                                            case 128:
                                                var v44 int = (^(fi37))
                                                _ = v44
                                            }
                                            _ = v43
                                        } else {
                                            var v45 int32 = 769
                                            v6 -= v6
                                            var v46 int64 = -877
                                            _ = v45
                                            _ = v46
                                        }
                                        fmt.Println(v1, v3)
                                        switch ((("" + "foo bar") + ("\n\t" + "\n\t")) + ("foo bar" + "")) {
                                        case "a":
                                            var v47 uint16 = uint16(v6)
                                            _ = v47
                                        }
                                    }
                                }
                            case 2147483647:
                                v48 := GenS4{F0: (fi37 | (^(fi31))), F1: v6}
                                _ = v48
                            default:
                                var v49 int8 = (v14 + (^(v14)))
                                _ = v49
                            }
                            _ = v40
                        }
                        if true {
                            var v50 uint64 = (v2 & (v2 << ((2) & 63)))
                            var v51 int = int(v14)
                            var v52 float32 = -2.25
                            _ = v50
                            _ = v51
                            _ = v52
                        } else {
                            switch "hello" {
                            case "foo bar":
                                switch (fi37 - (fi37 * fi31)) {
                                case 127:
                                    if (32767 != 827) {
                                        if (v1 >= (-2.25 - v1)) {
                                            _ = len(v11)
                                            var v53 uint8 = 182
                                            fmt.Println(v38)
                                            _ = v53
                                        }
                                        var v54 float32 = 3.14159
                                        v55 := GenS1{F0: 0}
                                        _ = v54
                                        _ = v55
                                    } else {
                                        var v56 uint8 = 127
                                        v14 += (v14 ^ v14)
                                        if v12 {
                                            v57 := GenS1{F0: v14}
                                            _ = v57
                                        }
                                        switch fi37 {
                                        case 32767:
                                            fmt.Println(v14, fi37)
                                        case 128:
                                            v2 = 0
                                        default:
                                            if (971 == 42) {
                                                if len(v11) >= 2 {
                                                    v11 = v11[1:len(v11)]
                                                }
                                                var v58 int16 = v5
                                                _ = v58
                                            }
                                        }
                                        if v4 {
                                            if (-2147483648 != -5) {
                                                v5 = (v5 & (v5 % ((-32768) | 1)))
                                                if (v2 != (^(v2))) {
                                                    v56 += (v56 % ((2) | 1))
                                                    if len(v11) >= 2 {
                                                        v11 = v11[1:len(v11)]
                                                    }
                                                    var v59 int8 = (v14 >> ((0) & 63))
                                                    _ = v59
                                                } else {
                                                    v60 := GenS3{F0: ((("foo bar" + ("a" + "\n\t")) + "\n\t") + ""), F1: 42}
                                                    var v61 int8 = int8(v5)
                                                    _ = v60
                                                    _ = v61
                                                }
                                            } else {
                                                v3 = v3
                                            }
                                            var v62 uint32 = v6
                                            var v63 uint16 = 255
                                            if (-2.25 < 1e10) {
                                                v3 = v4
                                                if false {
                                                    if v3 {
                                                        v12 = (((v12 && (629 != 426)) || (!((fi37 > (fi37 | (^(fi31)))) || (!(!v12))))) && v4)
                                                    } else {
                                                        switch 2147483647 {
                                                        case 128:
                                                            v3 = v4
                                                        case 42:
                                                            if (0.0 > 1e10) {
                                                                var v64 uint16 = 2
                                                                _ = v64
                                                            }
                                                        default:
                                                            v12 = false
                                                        }
                                                        switch "" {
                                                        case "hello":
                                                            v2 -= v2
                                                        case "\n\t":
                                                            if v3 {
                                                                v65 := GenS5{F0: 270}
                                                                _ = v65
                                                            }
                                                        default:
                                                            fmt.Println(v63)
                                                        }
                                                    }
                                                    var v66 int = -2147483648
                                                    _ = v66
                                                } else {
                                                    var v67 uint16 = (v63 / (((v63 ^ uint16(v56))) | 1))
                                                    v68 := GenS5{F0: 2}
                                                    v69 := GenS1{F0: 127}
                                                    _ = v67
                                                    _ = v68
                                                    _ = v69
                                                }
                                                var v70 int64 = 42
                                                var v71 uint16 = (v63 >> ((255) & 63))
                                                if (21 <= int16(v70)) {
                                                    v70 -= (v70 + (^(v70)))
                                                } else {
                                                    if len(v11) > 0 {
                                                        v11[0] = (fi31 >> ((924) & 63))
                                                    }
                                                    v5 += 1000
                                                    v72 := GenS1{F0: (-(v14))}
                                                    var v73 int32 = 419
                                                    v5 += (v5 | -513)
                                                    _ = v72
                                                    _ = v73
                                                }
                                                v1 += v38
                                                v74 := GenS2{F0: 2, F1: (("\n\t" + "hello") + "hello"), F2: 0}
                                                if len(v11) >= 2 {
                                                    v11 = v11[1:len(v11)]
                                                }
                                                _ = v70
                                                _ = v71
                                                _ = v74
                                            }
                                            _ = v62
                                            _ = v63
                                        }
                                        _ = v56
                                    }
                                case 9223372036854775807:
                                    v75 := GenS1{F0: int8(fi37)}
                                    _ = v75
                                }
                            case "\n\t":
                                if (((!v4) && (0.1 < 3.14159)) || v12) {
                                    if len(v11) > 0 {
                                        v11[0] = (fi31 + 452)
                                    }
                                } else {
                                    v2 = 0
                                }
                            case "a":
                                var v76 uint = 255
                                _ = v76
                            }
                            var v77 uint16 = 65535
                            _ = v77
                        }
                    default:
                        _ = len(v11)
                    }
                    v2 = 18446744073709551615
                    fmt.Println(v38)
                } else {
                    var v78 uint64 = v2
                    v2 = 422
                    v6 += 65535
                    _ = v78
                }
                if true {
                    fmt.Println(v38)
                    var v79 uint16 = 2
                    _ = v79
                }
                var v80 uint = 18446744073709551615
                _ = fi37
                _ = v38
                _ = v80
            }
            v11 = append(v11, (fi31 - (fi31 * fi31)))
            _ = fi31
        }
    case 32768:
        var v81 bool = v12
        _ = v81
    default:
        v11 = append(v11, 32767)
    }
    switch ((("a" + "hello") + "a") + "foo bar") {
    case "foo bar":
        v82 := GenS6{F0: (v5 * 761), F1: 18446744073709551615}
        _ = v82
    }
    var v83 uint16 = 255
    if len(v11) > 0 {
        v11[0] = 42
    }
    v84 := GenS4{F0: 32767, F1: 1}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v83:", v83)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v83
    _ = v84
    _ = v9
}
