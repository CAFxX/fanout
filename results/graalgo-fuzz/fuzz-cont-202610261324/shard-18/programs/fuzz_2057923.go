package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint64 }
type GenS2 struct { F0 int32; F1 uint16 }
type GenS3 struct { F0 uint32; F1 int16 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 uint; F1 int32 }
type GenS6 struct { F0 int; F1 string; F2 uint }
type GenS7 struct { F0 int8; F1 uint }
func genFn0(p0 string, p1 int8, p2 uint64) uint {
    var v5 bool = true
    for fi6 := 0; fi6 < 14; fi6++ {
        p0 = p0
        v7 := GenS1{F0: uint(p2), F1: 42}
        var v8 float64 = -2.25
        _ = fi6
        _ = v7
        _ = v8
    }
    var v9 uint32 = 2147483647
    v10 := GenS1{F0: 18446744073709551615, F1: (p2 / (((p2 - 42)) | 1))}
    p1 = -128
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v5
    _ = v9
    return 4294967295
}
func genFn1(p0 uint, p1 uint) uint {
    p1 -= p1
    if (!false) {
        p0 += p1
        p1 = 1
        for fi11 := 0; fi11 < 7; fi11++ {
            if true {
                var v12 uint32 = 9
                p1 -= p1
                if true {
                    if false {
                        for fi13 := 0; fi13 < 9; fi13++ {
                            var v14 string = ("foo bar" + (("a" + "") + ""))
                            fmt.Println(v12, v14, p0)
                            v15 := GenS2{F0: int32(v12), F1: 255}
                            _ = fi13
                            _ = v14
                            _ = v15
                        }
                    } else {
                        var v16 int = (fi11 * -550)
                        v17 := []int{}
                        var v18 uint8 = 255
                        _ = v16
                        _ = v17
                        _ = v18
                    }
                    switch (fi11 / (((fi11 << ((665) & 63))) | 1)) {
                    case 128:
                        fmt.Println(p0, p0, fi11)
                    case 42:
                        v19 := []int{(fi11 >> ((genFn0("\n\t", -1, 768)) & 63)), fi11, fi11}
                        _ = v19
                    case 0:
                        v12 -= (v12 | (v12 / (((v12 % (((v12 << ((p1) & 63))) | 1))) | 1)))
                    default:
                        v20 := []int{fi11, -2147483648}
                        _ = v20
                    }
                    v21 := []int{}
                    _ = v21
                }
                _ = v12
            } else {
                switch fi11 {
                case -2147483648:
                    v22 := GenS1{F0: (p1 - genFn0(("hello" + "a"), -1, 90)), F1: 892}
                    _ = v22
                default:
                    var v23 int = (fi11 >> ((p0) & 63))
                    _ = v23
                }
                v24 := []int{}
                if true {
                    p0 = (^(p0))
                    var v25 uint8 = 1
                    _ = v25
                }
                _ = v24
            }
            v26 := []int{(fi11 & (fi11 * fi11)), (-(fi11)), (^(fi11))}
            if len(v26) > 0 {
                v26[0] = fi11
            }
            _ = fi11
            _ = v26
        }
    } else {
        v27 := GenS3{F0: 991, F1: 42}
        _ = v27
    }
    v28 := []int{-328, 1, 2147483647, 9223372036854775807}
    _ = p0
    _ = p1
    _ = v28
    return 2
}

func main() {
    var v1 int8 = 18
    var v2 int64 = 661
    v3 := []string{(("" + "") + ("\n\t" + "\n\t")), ("" + "foo bar")}
    v4 := make(map[string]uint32)
    v1 = 0
    for fi29 := 0; fi29 < 5; fi29++ {
        v1 = v1
        var v30 uint64 = 1
        v1 -= (v1 & (v1 * (-(v1))))
        _ = fi29
        _ = v30
    }
    v2 = (v2 + 9223372036854775807)
    var v31 int16 = 1
    fmt.Println(v31, v2, v1)
    v31 -= -32768
    v32 := GenS4{F0: (v1 & 127)}
    v31 = (-(v31))
    var v33 int32 = 189
    for fi34 := 0; fi34 < 15; fi34++ {
        fmt.Println(fi34, v2, v33)
        for fi35 := 0; fi35 < 16; fi35++ {
            v2 = int64(fi35)
            _ = fi35
        }
        _ = fi34
    }
    if (((false && false) && (v33 > v33)) || (2 >= 0)) {
        v33 += v33
        var v36 uint64 = 327
        v1 = (^(v1))
        _ = v36
    }
    if (true && (0 <= 1)) {
        v37 := GenS3{F0: 4294967295, F1: 177}
        _ = v37
    } else {
        if true {
            if (!((true && false) || (1 >= 127))) {
                v4["foo bar"] = 0
                if true {
                    for fi38 := 0; fi38 < 19; fi38++ {
                        if false {
                            v33 -= int32(v1)
                            if ((1e10 > ((1.5 + 0.1) * 3.14159)) || true) {
                                for fi39 := 0; fi39 < 5; fi39++ {
                                    switch "a" {
                                    case "foo bar":
                                        switch "\n\t" {
                                        case "":
                                            var v40 int32 = (v33 + (v33 & (v33 ^ v33)))
                                            _ = v40
                                        }
                                    case "\n\t":
                                        v31 += (v31 / (((v31 / (((v31 / (((v31 ^ (v31 + 1000))) | 1))) | 1))) | 1))
                                    case "":
                                        v41 := GenS4{F0: (^(v1))}
                                        _ = v41
                                    default:
                                        if true {
                                            switch "foo bar" {
                                            case "foo bar":
                                                v42 := GenS1{F0: 42, F1: 1}
                                                _ = v42
                                            case "a":
                                                v33 += (v33 ^ v33)
                                            default:
                                                fmt.Println(v2)
                                            }
                                            var v43 bool = (0 < -2147483648)
                                            v44 := v4[((("hello" + "a") + "") + "\n\t")]
                                            _ = v43
                                            _ = v44
                                        }
                                    }
                                    v1 = v1
                                    if false {
                                        fmt.Println(fi39, v31)
                                    }
                                    _ = fi39
                                }
                                for fi45 := 0; fi45 < 15; fi45++ {
                                    if (4294967295 < 0) {
                                        var v46 string = ""
                                        _ = len(v3)
                                        _ = v46
                                    } else {
                                        v33 += 0
                                        var v47 uint64 = 25
                                        v33 -= 659
                                        _ = v47
                                    }
                                    var v48 int64 = v2
                                    _ = fi45
                                    _ = v48
                                }
                                v33 = (v33 | v33)
                            } else {
                                v49 := GenS5{F0: genFn1(0, 2), F1: v33}
                                switch ("a" + "\n\t") {
                                case "hello":
                                    for fi50 := 0; fi50 < 10; fi50++ {
                                        if ((!false) || (true && (true && true))) {
                                            v1 = (v1 + (^(v1)))
                                            v51 := GenS5{F0: 161, F1: (v33 & -2147483648)}
                                            _ = v51
                                        } else {
                                            if true {
                                                var v52 uint8 = 128
                                                if (42 == 0) {
                                                    var v53 int = int(v31)
                                                    var v54 int = 1
                                                    switch (-(fi38)) {
                                                    case 127:
                                                        v54 = (v53 | (v54 ^ (fi38 >> ((4294967295) & 63))))
                                                    case 9223372036854775807:
                                                        switch v53 {
                                                        case -298:
                                                            v54 = fi38
                                                        case 0:
                                                            var v55 bool = (0.0 >= 3.14159)
                                                            if len(v3) >= 2 {
                                                                v3 = v3[1:len(v3)]
                                                            }
                                                            _ = v55
                                                        case 32768:
                                                            if true {
                                                                v56 := GenS2{F0: (v33 / (((v33 >> ((42) & 63))) | 1)), F1: 1}
                                                                if true {
                                                                    fmt.Println(v54, v2)
                                                                    fmt.Println(v1)
                                                                    v3 = append(v3, "")
                                                                }
                                                                v57 := GenS6{F0: fi38, F1: "a", F2: 1}
                                                                var v58 uint = 0
                                                                if (!(9223372036854775807 < (v2 * 876))) {
                                                                    v58 -= 18446744073709551615
                                                                    v59 := GenS5{F0: (v58 ^ v58), F1: 861}
                                                                    _ = v59
                                                                }
                                                                _ = v56
                                                                _ = v57
                                                                _ = v58
                                                            }
                                                        default:
                                                            if (int64(v1) == (v2 | v2)) {
                                                                v31 = 42
                                                            } else {
                                                                if true {
                                                                    v1 = -128
                                                                } else {
                                                                    v33 -= v33
                                                                    var v60 uint64 = 962
                                                                    _ = v60
                                                                }
                                                                var v61 uint32 = 4294967295
                                                                _ = v61
                                                            }
                                                        }
                                                    }
                                                    _ = v53
                                                    _ = v54
                                                }
                                                _ = v52
                                            }
                                            v31 -= v31
                                        }
                                        v2 -= v2
                                        var v62 uint32 = 42
                                        _ = fi50
                                        _ = v62
                                    }
                                case "foo bar":
                                    if ((fi38 * (fi38 - 32767)) != fi38) {
                                        v33 = 209
                                    } else {
                                        v33 = (v33 | (v33 << ((genFn1(939, 18446744073709551615)) & 63)))
                                    }
                                case "":
                                    v33 = v33
                                }
                                _ = v49
                            }
                        }
                        for fi63 := 0; fi63 < 5; fi63++ {
                            if true {
                                fmt.Println(v2)
                            } else {
                                var v64 uint8 = 42
                                fmt.Println(fi63, v2)
                                v33 -= (v33 | -2147483648)
                                _ = v64
                            }
                            v1 = 100
                            switch ("hello" + "a") {
                            case "foo bar":
                                var v65 string = ""
                                _ = v65
                            case "a":
                                if (false || (true || true)) {
                                    fmt.Println(v2, v31, v2)
                                    switch "" {
                                    case "\n\t":
                                        if len(v3) > 0 {
                                            v3[0] = "\n\t"
                                        }
                                    case "a":
                                        v66 := GenS5{F0: 65535, F1: (v33 + v33)}
                                        _ = v66
                                    case "hello":
                                        v3 = append(v3, "foo bar")
                                    }
                                }
                            case "":
                                fmt.Println(fi38)
                            default:
                                var v67 int32 = (v33 & (v33 * v33))
                                _ = v67
                            }
                            var v68 int64 = v2
                            _ = fi63
                            _ = v68
                        }
                        v31 = (v31 / ((-32768) | 1))
                        _ = fi38
                    }
                } else {
                    if (255 > 1) {
                        v33 = v33
                    }
                }
                v69 := GenS7{F0: 0, F1: 255}
                _ = v69
            }
        }
    }
    var v70 int = -333
    switch "\n\t" {
    case "":
        if ((false && ((v31 ^ v31) != (v31 - v31))) && ((!(true && true)) || (!(true && true)))) {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            var v71 uint = 2
            _ = v71
        } else {
            switch (("" + ("foo bar" + "a")) + "a") {
            case "\n\t":
                var v72 int8 = 117
                _ = v72
            case "foo bar":
                for fi73 := 0; fi73 < 6; fi73++ {
                    if (((("foo bar" + "foo bar") + "hello") + "hello") >= "") {
                        v31 -= (v31 << ((255) & 63))
                        v74 := GenS2{F0: (^(v33)), F1: 32767}
                        var v75 uint64 = 1
                        _ = v74
                        _ = v75
                    }
                    v70 = (fi73 % ((v70) | 1))
                    var v76 int = fi73
                    _ = fi73
                    _ = v76
                }
            case "a":
                for fi77 := 0; fi77 < 6; fi77++ {
                    v2 -= int64(v1)
                    _ = fi77
                }
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    fmt.Println("v33:", v33)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v70:", v70)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v4
    _ = v70
}
