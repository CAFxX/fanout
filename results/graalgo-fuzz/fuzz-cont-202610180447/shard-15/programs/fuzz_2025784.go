package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int64 }
type GenS2 struct { F0 uint32; F1 uint32; F2 uint }
type GenS3 struct { F0 uint16; F1 int64; F2 uint16 }
type GenS4 struct { F0 bool; F1 int8 }
type GenS5 struct { F0 bool; F1 bool; F2 bool }
type GenS6 struct { F0 uint64; F1 uint8 }
type GenS7 struct { F0 int32; F1 string; F2 int }
type GenS8 struct { F0 int; F1 uint16 }
type GenS9 struct { F0 string; F1 bool; F2 uint8 }
type GenS10 struct { F0 string; F1 int; F2 uint8 }
type GenS11 struct { F0 int64; F1 int8; F2 int8 }
func genFn0(p0 int64) int64 {
    v7 := []int{4294967295}
    var v8 float32 = 0.1
    v9 := GenS1{F0: 799, F1: p0}
    _ = p0
    _ = v7
    _ = v8
    _ = v9
    return 42
}
func genFn1(p0 int32, p1 bool) uint32 {
    p0 -= p0
    var v10 uint32 = 2147483647
    if (!p1) {
        if true {
            for fi11 := 0; fi11 < 17; fi11++ {
                p1 = p1
                _ = fi11
            }
        }
        for fi12 := 0; fi12 < 4; fi12++ {
            var v13 bool = false
            v14 := GenS2{F0: (v10 + 748), F1: (^(v10)), F2: 255}
            _ = fi12
            _ = v13
            _ = v14
        }
        v15 := GenS2{F0: v10, F1: (^(v10)), F2: 42}
        _ = v15
    }
    v16 := GenS3{F0: 0, F1: 0, F2: 210}
    _ = p0
    _ = p1
    _ = v10
    _ = v16
    return 0
}

func main() {
    v1 := make(map[int]int)
    var v2 uint8 = 2
    v3 := []int{}
    var v4 int8 = 100
    v5 := []int{-1}
    var v6 uint64 = 749
    v17 := GenS4{F0: true, F1: v4}
    _ = len(v5)
    v2 += v2
    v18 := GenS5{F0: (!(v6 == (v6 % (((^(v6))) | 1)))), F1: ((v6 - v6) != 0), F2: true}
    var v19 float64 = -2.25
    var v20 uint64 = 165
    v4 = v4
    v4 = (v4 | v4)
    if (-903 < 2147483647) {
        v2 = (v2 << ((18446744073709551615) & 63))
        if true {
            for fi21 := 0; fi21 < 10; fi21++ {
                if (fi21 == 555) {
                    _ = len(v5)
                    v22 := GenS6{F0: v6, F1: (^(v2))}
                    v23 := GenS7{F0: 42, F1: "hello", F2: fi21}
                    _ = v22
                    _ = v23
                }
                var v24 uint8 = v2
                v25 := GenS5{F0: true, F1: ((v24 ^ v24) >= v24), F2: (1 < 0)}
                _ = fi21
                _ = v24
                _ = v25
            }
            for fi26 := 0; fi26 < 16; fi26++ {
                if false {
                    if ((^(v4)) != (v4 - 100)) {
                        var v27 int32 = 42
                        _ = v27
                    } else {
                        var v28 uint8 = 127
                        if false {
                            var v29 uint8 = v2
                            for fi30 := 0; fi30 < 10; fi30++ {
                                v6 += 354
                                _ = fi30
                            }
                            _ = v29
                        } else {
                            fmt.Println(v20)
                            v4 -= (^(v4))
                            switch -478 {
                            case 32768:
                                for fi31 := 0; fi31 < 14; fi31++ {
                                    var v32 float32 = 0.0
                                    fmt.Println(v32, fi31)
                                    _ = fi31
                                    _ = v32
                                }
                            }
                        }
                        switch fi26 {
                        case 32768:
                            if true {
                                var v33 string = "a"
                                _ = v33
                            }
                        }
                        _ = v28
                    }
                    fmt.Println(fi26)
                } else {
                    if (("\n\t" + ("a" + ("\n\t" + ""))) > ((("hello" + "\n\t") + "a") + "")) {
                        v19 = 3.14159
                        v19 = ((0.0 - 1e10) - 3.14159)
                        v2 += 127
                    }
                    var v34 int = (fi26 - fi26)
                    _ = v34
                }
                _ = fi26
            }
        } else {
            if (-745 < int64(v20)) {
                var v35 int16 = 1
                _ = v35
            }
            if false {
                switch 2147483647 {
                case 127:
                    v2 = 42
                default:
                    v36 := v1[690]
                    _ = v36
                }
                switch "\n\t" {
                case "hello":
                    if false {
                        if (true && true) {
                            var v37 uint16 = 65535
                            v4 = (v4 >> ((uint(v37)) & 63))
                            _ = v37
                        }
                        for fi38 := 0; fi38 < 3; fi38++ {
                            fmt.Println(v2, v6)
                            switch (^(fi38)) {
                            case 127:
                                v4 = (v4 & 127)
                            default:
                                var v39 int16 = 42
                                _ = v39
                            }
                            v40 := v1[fi38]
                            _ = fi38
                            _ = v40
                        }
                    } else {
                        var v41 uint32 = uint32(v2)
                        v42 := GenS5{F0: (0.0 >= (1.5 * 1.5)), F1: false, F2: (!((v41 + (^(v41))) > v41))}
                        v2 -= v2
                        _ = v41
                        _ = v42
                    }
                case "\n\t":
                    v43 := GenS8{F0: 30, F1: 988}
                    _ = v43
                case "":
                    v6 -= (^(v6))
                }
                if false {
                    if true {
                        if len(v5) > 0 {
                            v5[0] = 127
                        }
                    }
                    for fi44 := 0; fi44 < 18; fi44++ {
                        v45 := GenS4{F0: false, F1: v4}
                        _ = fi44
                        _ = v45
                    }
                    v46 := GenS5{F0: true, F1: true, F2: false}
                    _ = v46
                }
            } else {
                fmt.Println(v19)
            }
        }
    }
    fmt.Println(v2)
    for fi47 := 0; fi47 < 7; fi47++ {
        switch "a" {
        case "a":
            switch "hello" {
            case "":
                switch (fi47 - fi47) {
                case -2147483648:
                    v19 += v19
                case -343:
                    if false {
                        v2 = 255
                        v48 := GenS1{F0: v6, F1: genFn0(genFn0(410))}
                        switch "\n\t" {
                        case "foo bar":
                            var v49 int = (^(fi47))
                            _ = v49
                        case "":
                            fmt.Println(v4)
                        }
                        _ = v48
                    } else {
                        if true {
                            if (("hello" + (("foo bar" + "foo bar") + "a")) >= "") {
                                for fi50 := 0; fi50 < 12; fi50++ {
                                    if false {
                                        var v51 string = "\n\t"
                                        _ = v51
                                    } else {
                                        delete(v1, fi47)
                                    }
                                    fmt.Println(fi50, fi47, fi47)
                                    v52 := GenS9{F0: "", F1: ("" < ""), F2: v2}
                                    _ = fi50
                                    _ = v52
                                }
                                v53 := GenS8{F0: (fi47 % ((fi47) | 1)), F1: 255}
                                _ = v53
                            }
                            for fi54 := 0; fi54 < 14; fi54++ {
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                var v55 string = ((("a" + "") + (("\n\t" + "a") + "a")) + "foo bar")
                                fmt.Println(fi47)
                                _ = fi54
                                _ = v55
                            }
                        }
                        v19 += (v19 - -2.25)
                    }
                default:
                    var v56 int16 = 32767
                    _ = v56
                }
            case "hello":
                delete(v1, (fi47 * fi47))
            default:
                if (((true || true) || (((v19 * float64(0.1)) * ((1e10 - 1.5) - v19)) != v19)) || (((v19 - v19) >= 0.1) && ((fi47 * 996) >= fi47))) {
                    _ = len(v1)
                } else {
                    for fi57 := 0; fi57 < 6; fi57++ {
                        v58 := GenS1{F0: v20, F1: 42}
                        v2 = (v2 / ((255) | 1))
                        _ = fi57
                        _ = v58
                    }
                }
            }
        case "foo bar":
            v3 = append(v3, fi47)
        case "\n\t":
            if true {
                var v59 uint64 = 18446744073709551615
                v6 = v59
                _ = v59
            }
        }
        _ = fi47
    }
    for fi60 := 0; fi60 < 2; fi60++ {
        for fi61 := 0; fi61 < 3; fi61++ {
            v4 += v4
            _ = fi61
        }
        v62 := GenS6{F0: (v20 | (v6 / ((v20) | 1))), F1: v2}
        if true {
            for fi63 := 0; fi63 < 4; fi63++ {
                var v64 uint16 = 404
                if false {
                    if true {
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                        v20 = (^(v6))
                    } else {
                        var v65 uint32 = 2
                        if true {
                            var v66 int16 = 1000
                            var v67 uint8 = 218
                            if false {
                                var v68 int = 362
                                switch (fi63 / (((v68 | fi63)) | 1)) {
                                case -2147483648:
                                    if len(v5) >= 2 {
                                        v5 = v5[1:len(v5)]
                                    }
                                case 1:
                                    fmt.Println(v2)
                                default:
                                    fmt.Println(v2, v19, v2)
                                }
                                v64 = (v64 / (((v64 ^ (v64 % ((42) | 1)))) | 1))
                                _ = v68
                            } else {
                                fmt.Println(v66)
                                if true {
                                    fmt.Println(v66, v4, v65)
                                    var v69 int = fi60
                                    fmt.Println(v2, v4, v20)
                                    _ = v69
                                } else {
                                    v20 += uint64(v64)
                                    switch (fi63 << ((18446744073709551615) & 63)) {
                                    case 127:
                                        fmt.Println(v19, v2, v20)
                                    case -671:
                                        v64 = v64
                                    default:
                                        var v70 int8 = v4
                                        _ = v70
                                    }
                                }
                            }
                            _ = v66
                            _ = v67
                        }
                        v19 += v19
                        _ = v65
                    }
                    var v71 int8 = int8(v6)
                    if true {
                        v71 = (v71 - v4)
                        v72 := GenS7{F0: -553, F1: "\n\t", F2: 366}
                        _ = v72
                    } else {
                        var v73 int64 = 1
                        var v74 float32 = 1e10
                        v6 = v6
                        _ = v73
                        _ = v74
                    }
                    _ = v71
                } else {
                    v2 = v2
                    if (!(443 != -265)) {
                        v19 += (v19 - 3.14159)
                    } else {
                        if (("a" + "\n\t") == ("hello" + ("a" + ("hello" + (((("hello" + "\n\t") + ("\n\t" + "foo bar")) + ("a" + ("" + "\n\t"))) + "foo bar"))))) {
                            v75 := GenS4{F0: false, F1: (v4 << ((42) & 63))}
                            switch -132 {
                            case 4294967295:
                                v1[-712] = fi60
                            case -2147483648:
                                var v76 int64 = -1
                                _ = v76
                            case 444:
                                switch ("a" + "a") {
                                case "foo bar":
                                    var v77 uint16 = uint16(v6)
                                    _ = v77
                                }
                                var v78 int64 = 1
                                v64 -= 2
                                _ = v78
                            default:
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                            }
                            _ = v75
                        } else {
                            var v79 int16 = -667
                            var v80 float32 = (-2.25 * 1e10)
                            fmt.Println(v4, v80)
                            _ = v79
                            _ = v80
                        }
                        if len(v5) > 0 {
                            v5[0] = 42
                        }
                        _ = len(v1)
                    }
                    v4 -= 127
                }
                if false {
                    var v81 int64 = genFn0(164)
                    var v82 int32 = -2147483648
                    v83 := GenS10{F0: "\n\t", F1: fi60, F2: (v2 - (v2 + (v2 ^ (v2 - 87))))}
                    fmt.Println(v19, v19, fi63)
                    var v84 uint = 2
                    _ = v81
                    _ = v82
                    _ = v83
                    _ = v84
                }
                v64 += (v64 & v64)
                _ = fi63
                _ = v64
            }
        }
        _ = fi60
        _ = v62
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    if true {
        v20 = 4294967295
    }
    if (9223372036854775807 != 1) {
        v19 -= 0.1
        switch "foo bar" {
        case "foo bar":
            v4 -= 21
        default:
            v85 := GenS11{F0: 409, F1: v4, F2: (v4 & v4)}
            _ = v85
        }
    }
    v20 = 18446744073709551615
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
