package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 string; F1 int }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int16; F1 int64; F2 string }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 int8; F1 int8; F2 bool }
type GenS7 struct { F0 uint64 }
type GenS8 struct { F0 int8; F1 int; F2 string }
type GenS9 struct { F0 uint32; F1 int64; F2 string }
type GenS10 struct { F0 bool; F1 uint; F2 int64 }
type GenS11 struct { F0 int; F1 uint }
type GenS12 struct { F0 uint32; F1 int; F2 uint32 }
type GenS13 struct { F0 uint32; F1 uint8 }
type GenS14 struct { F0 uint16 }
type GenS15 struct { F0 string; F1 uint16 }

func main() {
    var v1 uint16 = 0
    var v2 uint = 18446744073709551615
    var v3 uint32 = 0
    var v4 int = 42
    v5 := make(map[uint32]int)
    v6 := []int{(v4 % (((v4 ^ (v4 | (v4 * v4)))) | 1)), 32768, v4}
    if true {
        v2 = (^(v2))
    }
    v7 := GenS1{F0: -466}
    if (!false) {
        v4 -= -2147483648
        v4 -= (^(v4))
        if (!false) {
            v8 := GenS1{F0: -1}
            var v9 int8 = -1
            _ = v8
            _ = v9
        } else {
            if ((1 < -128) || false) {
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
            }
            for fi10 := 0; fi10 < 11; fi10++ {
                v3 += v3
                _ = fi10
            }
        }
    } else {
        var v11 int64 = 1
        v12 := GenS2{F0: "a", F1: v4}
        for fi13 := 0; fi13 < 12; fi13++ {
            for fi14 := 0; fi14 < 12; fi14++ {
                if (!true) {
                    var v15 float32 = 1.5
                    switch 32768 {
                    case 32768:
                        v3 += (^(v3))
                    case -244:
                        v16 := GenS3{F0: 9223372036854775807}
                        _ = v16
                    case 835:
                        if false {
                            fmt.Println(fi13)
                            var v17 int16 = -32768
                            _ = v17
                        }
                    default:
                        if true {
                            v11 += (v11 ^ 9223372036854775807)
                            v18 := GenS4{F0: 0, F1: (-(v11)), F2: "\n\t"}
                            var v19 uint16 = v1
                            _ = v18
                            _ = v19
                        }
                    }
                    if (-1 != (v11 & v11)) {
                        v2 = (v2 / (((v2 | v2)) | 1))
                        v20 := GenS5{F0: (v2 ^ v2)}
                        delete(v5, 23)
                        _ = v20
                    } else {
                        v4 += (^(fi14))
                    }
                    _ = v15
                } else {
                    v21 := GenS6{F0: 1, F1: 100, F2: true}
                    v22 := GenS7{F0: 18446744073709551615}
                    if ((0.1 * 0.1) <= -2.25) {
                        switch "" {
                        case "a":
                            if false {
                                var v23 bool = (true || ((((!(float32((1.5 * ((float64((-2.25 + 1.5)) / 3.14159) - 0.0))) < (3.14159 / 1.0))) || true) && true) || (!true)))
                                _ = v23
                            } else {
                                fmt.Println(v11, v1)
                            }
                        case "hello":
                            v1 = 255
                        case "\n\t":
                            var v24 bool = (!(42 == 1))
                            _ = v24
                        }
                        var v25 uint = 854
                        _ = v25
                    }
                    _ = v21
                    _ = v22
                }
                _ = fi14
            }
            if false {
                v26 := GenS8{F0: 100, F1: 1, F2: "hello"}
                if true {
                    var v27 int64 = v11
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                    _ = v27
                }
                _ = v26
            } else {
                switch fi13 {
                case 269:
                    switch ("foo bar" + (((("foo bar" + "a") + ("\n\t" + "")) + "") + ((("hello" + "") + "a") + ("\n\t" + ("hello" + "hello"))))) {
                    case "":
                        v3 -= (^(v3))
                    default:
                        v3 -= (v3 + v3)
                    }
                }
                for fi28 := 0; fi28 < 8; fi28++ {
                    fmt.Println(v4, v2, v2)
                    var v29 float32 = 3.14159
                    v29 -= 0.0
                    _ = fi28
                    _ = v29
                }
                v11 -= v11
            }
            _ = fi13
        }
        _ = v11
        _ = v12
    }
    if false {
        v3 += 2147483647
        delete(v5, v3)
    }
    v6 = append(v6, int(v3))
    fmt.Println(v2, v4, v3)
    for fi30 := 0; fi30 < 11; fi30++ {
        switch (^(fi30)) {
        case -2147483648:
            v5[65535] = (fi30 - (v4 * (fi30 * (v4 * fi30))))
        case 977:
            v2 += (v2 & v2)
        default:
            for fi31 := 0; fi31 < 7; fi31++ {
                var v32 string = ("\n\t" + "a")
                v33 := GenS2{F0: (v32 + (v32 + (("hello" + "a") + v32))), F1: (fi31 ^ -2147483648)}
                delete(v5, v3)
                _ = fi31
                _ = v32
                _ = v33
            }
        }
        for fi34 := 0; fi34 < 1; fi34++ {
            v2 = v2
            if (int(v3) > (fi30 >> ((v2) & 63))) {
                switch v4 {
                case 0:
                    v35 := GenS9{F0: uint32(v2), F1: -9223372036854775808, F2: "hello"}
                    _ = v35
                case 42:
                    fmt.Println(v2, v3)
                case -462:
                    if (!false) {
                        v3 = v3
                        v36 := GenS10{F0: ((1 <= 42) && false), F1: 1, F2: 1}
                        _ = v36
                    }
                default:
                    v2 += (^(v2))
                }
                if (2147483647 <= 726) {
                    fmt.Println(v4)
                    v4 += (v4 ^ v4)
                } else {
                    var v37 string = ("a" + "foo bar")
                    fmt.Println(v4)
                    _ = v37
                }
                switch "" {
                case "\n\t":
                    v38 := GenS11{F0: fi34, F1: v2}
                    _ = v38
                case "hello":
                    var v39 uint64 = 797
                    _ = v39
                }
                fmt.Println(v1, fi34)
                v40 := GenS4{F0: -611, F1: 1, F2: ((("hello" + "") + ((("" + ((("\n\t" + "foo bar") + "a") + "a")) + "a") + ("a" + "foo bar"))) + "")}
                _ = v40
            }
            _ = fi34
        }
        _ = fi30
    }
    fmt.Println(v4, v3, v4)
    var v41 float32 = 0.0
    v42 := v5[(v3 & 0)]
    var v43 int8 = -51
    for fi44 := 0; fi44 < 4; fi44++ {
        for fi45 := 0; fi45 < 8; fi45++ {
            if false {
                if true {
                    var v46 float64 = 1.5
                    if len(v6) > 0 {
                        v6[0] = (v4 * (v4 >> ((v2) & 63)))
                    }
                    var v47 uint8 = 128
                    var v48 uint64 = 42
                    fmt.Println(v43, v2)
                    _ = v46
                    _ = v47
                    _ = v48
                }
                if (0 == 0) {
                    v49 := GenS6{F0: 100, F1: 127, F2: true}
                    var v50 uint32 = (v3 & 4294967295)
                    fmt.Println(v42, v50)
                    _ = v49
                    _ = v50
                } else {
                    v2 = (v2 - v2)
                    fmt.Println(v2)
                }
            }
            _ = fi45
        }
        if true {
            if (fi44 != v4) {
                if (("a" + (("" + "") + (("a" + ("foo bar" + "hello")) + ((("hello" + "foo bar") + "foo bar") + "")))) <= ("a" + "")) {
                    v51 := GenS2{F0: "", F1: (^(fi44))}
                    v42 = int(v43)
                    for fi52 := 0; fi52 < 9; fi52++ {
                        if (!false) {
                            switch ("hello" + "") {
                            case "":
                                if true {
                                    if false {
                                        fmt.Println(v1, fi44, v2)
                                        v42 = (fi44 ^ (-(fi52)))
                                    }
                                    var v53 int32 = -1
                                    v1 = v1
                                    _ = v53
                                } else {
                                    v3 = 42
                                }
                            case "\n\t":
                                if true {
                                    var v54 uint8 = 1
                                    v41 = v41
                                    v5[2] = v4
                                    _ = v54
                                }
                            case "a":
                                var v55 int64 = -1
                                _ = v55
                            }
                        } else {
                            fmt.Println(v3)
                            v42 -= int(v2)
                            fmt.Println(v43)
                            v3 += (v3 >> ((v2) & 63))
                        }
                        _ = fi52
                    }
                    _ = v51
                } else {
                    v56 := GenS12{F0: v3, F1: v42, F2: v3}
                    _ = v56
                }
            } else {
                if true {
                    for fi57 := 0; fi57 < 3; fi57++ {
                        v41 = v41
                        var v58 uint16 = (v1 + v1)
                        v59 := GenS13{F0: v3, F1: 95}
                        if true {
                            fmt.Println(v41)
                            v43 -= v43
                        } else {
                            v60 := GenS12{F0: (v3 - (^(v3))), F1: 127, F2: (v3 + 65535)}
                            if true {
                                fmt.Println(v58, v4)
                            } else {
                                var v61 uint = (v2 >> ((65535) & 63))
                                _ = v61
                            }
                            v62 := GenS6{F0: (v43 & v43), F1: 100, F2: (v58 <= (v58 << ((v2) & 63)))}
                            _ = v60
                            _ = v62
                        }
                        _ = fi57
                        _ = v58
                        _ = v59
                    }
                    v5[(v3 - v3)] = v42
                } else {
                    var v63 float32 = (((1.5 / v41) / v41) - v41)
                    _ = v63
                }
                var v64 uint = 950
                if false {
                    switch 9223372036854775807 {
                    case 110:
                        var v65 uint64 = uint64(v3)
                        _ = v65
                    case -105:
                        fmt.Println(fi44)
                    case 493:
                        for fi66 := 0; fi66 < 17; fi66++ {
                            if (true || false) {
                                switch "" {
                                case "hello":
                                    fmt.Println(v4, v2, v4)
                                default:
                                    var v67 uint32 = v3
                                    _ = v67
                                }
                                var v68 uint = (v64 + 1)
                                _ = v68
                            } else {
                                var v69 bool = false
                                var v70 float32 = 1.5
                                _ = len(v5)
                                _ = v69
                                _ = v70
                            }
                            if false {
                                fmt.Println(v64, fi66, v64)
                            } else {
                                v2 -= (^(v64))
                            }
                            _ = fi66
                        }
                    default:
                        v71 := GenS14{F0: (v1 | v1)}
                        _ = v71
                    }
                }
                _ = v64
            }
        }
        v72 := GenS15{F0: "hello", F1: (v1 << (((v2 ^ (^(v2)))) & 63))}
        _ = fi44
        _ = v72
    }
    v73 := v5[(v3 & (v3 % ((4294967295) | 1)))]
    if true {
        fmt.Println(v2, v3)
        if true {
            var v74 int = (-(v42))
            _ = len(v5)
            _ = v74
        } else {
            fmt.Println(v2, v2)
            v43 += (v43 >> ((v2) & 63))
            if (128 >= 255) {
                for fi75 := 0; fi75 < 3; fi75++ {
                    var v76 float64 = 1e10
                    var v77 string = "hello"
                    var v78 int32 = 722
                    _ = fi75
                    _ = v76
                    _ = v77
                    _ = v78
                }
            } else {
                var v79 int16 = 0
                v42 = (v73 / (((-(v4))) | 1))
                if (!(42 >= 128)) {
                    var v80 int32 = 2147483647
                    v81 := GenS2{F0: ("a" + (("\n\t" + "a") + (("" + "hello") + ("a" + "foo bar")))), F1: v4}
                    fmt.Println(v41)
                    _ = v80
                    _ = v81
                }
                _ = v79
            }
        }
        v3 -= v3
    } else {
        switch v4 {
        case 42:
            var v82 uint32 = uint32(v2)
            _ = v82
        case 2147483647:
            for fi83 := 0; fi83 < 11; fi83++ {
                var v84 uint16 = v1
                _ = fi83
                _ = v84
            }
        case 571:
            fmt.Println(v4)
        }
        _ = len(v5)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v73:", v73)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v42
    _ = v43
    _ = v5
    _ = v6
    _ = v7
    _ = v73
}
