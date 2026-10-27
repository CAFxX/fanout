package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int64 }
type GenS2 struct { F0 uint; F1 int8 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int8; F1 bool; F2 int8 }
type GenS5 struct { F0 int32; F1 uint; F2 int32 }
type GenS6 struct { F0 int; F1 uint64 }
type GenS7 struct { F0 int64 }
type GenS8 struct { F0 uint16; F1 uint32; F2 int8 }
type GenS9 struct { F0 int32; F1 uint8 }
type GenS10 struct { F0 int8; F1 uint32; F2 int32 }
type GenS11 struct { F0 int }
type GenS12 struct { F0 int8; F1 int }
func genFn0(p0 uint, p1 int64) bool {
    v7 := []int{32767}
    p1 -= 0
    p0 += (p0 - p0)
    if len(v7) >= 2 {
        v7 = v7[1:len(v7)]
    }
    p0 = p0
    _ = p0
    _ = p1
    _ = v7
    return false
}
func genFn1(p0 int, p1 string, p2 uint) uint {
    v8 := []int{(^(p0))}
    var v9 uint32 = 2147483648
    switch p0 {
    case -9223372036854775808:
        if true {
            switch ("foo bar" + p1) {
            case "hello":
                for fi10 := 0; fi10 < 4; fi10++ {
                    switch p1 {
                    case "hello":
                        if len(v8) > 0 {
                            v8[0] = 32767
                        }
                    case "":
                        v11 := GenS1{F0: (2147483647 != 1), F1: 0}
                        _ = v11
                    case "a":
                        if true {
                            var v12 int32 = 1
                            v13 := GenS1{F0: true, F1: -1}
                            for fi14 := 0; fi14 < 7; fi14++ {
                                v15 := GenS1{F0: (true && (-9223372036854775808 < 0)), F1: -532}
                                v16 := GenS2{F0: p2, F1: 100}
                                if len(v8) > 0 {
                                    v8[0] = fi10
                                }
                                _ = fi14
                                _ = v15
                                _ = v16
                            }
                            _ = v12
                            _ = v13
                        } else {
                            for fi17 := 0; fi17 < 13; fi17++ {
                                v18 := GenS3{F0: (true || false)}
                                v9 = v9
                                _ = fi17
                                _ = v18
                            }
                        }
                    default:
                        var v19 float64 = (1.5 - 0.0)
                        _ = v19
                    }
                    for fi20 := 0; fi20 < 20; fi20++ {
                        p1 = p1
                        var v21 int64 = 1
                        p0 += (fi20 >> ((p2) & 63))
                        p2 -= p2
                        _ = fi20
                        _ = v21
                    }
                    fmt.Println(fi10, v9, p0)
                    _ = fi10
                }
            case "a":
                v9 -= (^(v9))
            }
            for fi22 := 0; fi22 < 16; fi22++ {
                var v23 uint32 = v9
                switch ("\n\t" + ((p1 + ("\n\t" + p1)) + (p1 + (p1 + p1)))) {
                case "":
                    var v24 int64 = 1
                    _ = v24
                default:
                    v25 := GenS4{F0: -1, F1: false, F2: 127}
                    _ = v25
                }
                _ = fi22
                _ = v23
            }
            p2 = p2
        } else {
            for fi26 := 0; fi26 < 6; fi26++ {
                v27 := GenS5{F0: -1, F1: p2, F2: -47}
                _ = fi26
                _ = v27
            }
            v28 := GenS5{F0: -1, F1: (p2 - (p2 << (((p2 - (p2 ^ p2))) & 63))), F2: -1}
            switch (p1 + "hello") {
            case "hello":
                if (float32(float64(0.1)) > 3.14159) {
                    var v29 float32 = float32((0.1 / 0.1))
                    var v30 int16 = -227
                    v31 := GenS6{F0: 1, F1: 42}
                    _ = v29
                    _ = v30
                    _ = v31
                }
            default:
                var v32 int8 = -1
                _ = v32
            }
            _ = v28
        }
    }
    if len(v8) >= 2 {
        v8 = v8[1:len(v8)]
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    _ = v9
    return p2
}

func main() {
    var v1 float32 = 1.5
    v2 := []int{32767, 9223372036854775807}
    var v3 uint64 = 366
    v4 := make(map[string]string)
    v5 := []uint32{65535}
    var v6 int64 = 0
    if true {
        v33 := GenS7{F0: v6}
        _ = v33
    } else {
        if false {
            v34 := GenS2{F0: 806, F1: 42}
            _ = v34
        } else {
            for fi35 := 0; fi35 < 14; fi35++ {
                v5 = append(v5, 65535)
                var v36 uint16 = 0
                _ = fi35
                _ = v36
            }
            _ = len(v4)
        }
    }
    var v37 uint16 = 42
    if (2147483647 < 42) {
        var v38 int32 = 2147483647
        v39 := GenS4{F0: 100, F1: false, F2: 42}
        v4[("a" + "\n\t")] = (("a" + "a") + "hello")
        _ = v38
        _ = v39
    }
    var v40 uint = 2
    for fi41 := 0; fi41 < 13; fi41++ {
        for fi42 := 0; fi42 < 4; fi42++ {
            v1 += 0.0
            var v43 int = fi42
            _ = fi42
            _ = v43
        }
        if true {
            for fi44 := 0; fi44 < 11; fi44++ {
                v45 := GenS8{F0: 0, F1: 2147483647, F2: 125}
                _ = fi44
                _ = v45
            }
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
        }
        _ = len(v2)
        _ = fi41
    }
    _ = len(v2)
    if false {
        delete(v4, "")
        if true {
            if true {
                for fi46 := 0; fi46 < 19; fi46++ {
                    v5 = append(v5, 0)
                    var v47 int32 = 34
                    if true {
                        if (true && (!(true && ((1 >= (v3 ^ (v3 | (v3 + (v3 * v3))))) && false)))) {
                            for fi48 := 0; fi48 < 14; fi48++ {
                                var v49 int8 = 1
                                if (277 < fi48) {
                                    switch -726 {
                                    case -59:
                                        switch "foo bar" {
                                        case "hello":
                                            switch "a" {
                                            case "":
                                                v49 = v49
                                            case "\n\t":
                                                fmt.Println(fi48)
                                            default:
                                                if true {
                                                    v1 -= (1.5 - (v1 + v1))
                                                } else {
                                                    v47 = v47
                                                }
                                            }
                                        default:
                                            v50 := GenS9{F0: (v47 >> (((^(v40))) & 63)), F1: 127}
                                            var v51 uint64 = v3
                                            var v52 uint16 = v37
                                            _ = v50
                                            _ = v51
                                            _ = v52
                                        }
                                    case -1:
                                        v53 := v4[""]
                                        _ = v53
                                    }
                                    _ = len(v5)
                                    switch fi48 {
                                    case -932:
                                        v40 = 252
                                    }
                                }
                                if false {
                                    v54 := GenS6{F0: (fi46 >> (((v40 % ((uint(v3)) | 1))) & 63)), F1: v3}
                                    var v55 uint = genFn1(916, ("hello" + "foo bar"), (v40 * v40))
                                    switch fi46 {
                                    case 128:
                                        v56 := GenS1{F0: (true || (false && ((v55 << (((v40 ^ 1)) & 63)) < uint(v37)))), F1: (v6 * int64(v47))}
                                        _ = v56
                                    case 106:
                                        fmt.Println(v3)
                                    default:
                                        var v57 uint32 = 1
                                        _ = v57
                                    }
                                    _ = v54
                                    _ = v55
                                } else {
                                    v1 = v1
                                    v2 = append(v2, fi48)
                                }
                                if true {
                                    v58 := GenS9{F0: (v47 << ((genFn1((fi48 - (fi48 << (((v40 >> (((v40 >> ((42) & 63))) & 63))) & 63))), "a", (^(v40)))) & 63)), F1: 127}
                                    fmt.Println(fi48, fi46)
                                    _ = v58
                                }
                                _ = fi48
                                _ = v49
                            }
                            if ("hello" < "") {
                                var v59 bool = true
                                v60 := GenS4{F0: -1, F1: v59, F2: -101}
                                var v61 float32 = v1
                                _ = v59
                                _ = v60
                                _ = v61
                            } else {
                                if (v37 <= 0) {
                                    v6 = (v6 >> (((v40 << ((v40) & 63))) & 63))
                                } else {
                                    v1 -= (1e10 * v1)
                                    var v62 float64 = 3.14159
                                    v1 = ((0.0 * 1.5) - v1)
                                    _ = v62
                                }
                            }
                        } else {
                            if ((100 != -1) && ((!(false || (0 == 765))) || true)) {
                                var v63 float32 = v1
                                for fi64 := 0; fi64 < 10; fi64++ {
                                    var v65 uint16 = v37
                                    v3 += v3
                                    switch "hello" {
                                    case "":
                                        fmt.Println(v63, v40)
                                        fmt.Println(v65, v63, v37)
                                    case "a":
                                        var v66 int64 = -9223372036854775808
                                        _ = v66
                                    }
                                    _ = fi64
                                    _ = v65
                                }
                                _ = v63
                            }
                            for fi67 := 0; fi67 < 11; fi67++ {
                                var v68 float32 = 0.1
                                _ = fi67
                                _ = v68
                            }
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                        }
                    } else {
                        if true {
                            var v69 bool = false
                            _ = v69
                        }
                        fmt.Println(v40, v47)
                        v70 := GenS7{F0: v6}
                        _ = v70
                    }
                    _ = fi46
                    _ = v47
                }
                if true {
                    switch -672 {
                    case -1:
                        var v71 string = ""
                        _ = v71
                    case 533:
                        v1 -= v1
                    case 344:
                        if (128 == 32767) {
                            v72 := GenS10{F0: 0, F1: 0, F2: -167}
                            _ = v72
                        } else {
                            v73 := GenS9{F0: 396, F1: 62}
                            _ = v73
                        }
                    default:
                        for fi74 := 0; fi74 < 17; fi74++ {
                            var v75 int64 = v6
                            fmt.Println(v40)
                            switch (fi74 ^ (^(fi74))) {
                            case 32767:
                                v76 := GenS9{F0: 874, F1: 42}
                                _ = v76
                            case -1:
                                for fi77 := 0; fi77 < 1; fi77++ {
                                    var v78 bool = (v40 != genFn1(fi77, "", genFn1((fi74 + fi74), "hello", (v40 << (((v40 * genFn1(fi74, "", genFn1(fi74, "hello", v40)))) & 63)))))
                                    _ = fi77
                                    _ = v78
                                }
                            default:
                                v75 = 0
                            }
                            _ = fi74
                            _ = v75
                        }
                    }
                }
                var v79 uint32 = 1
                _ = v79
            } else {
                v1 = v1
                if true {
                    v4["a"] = "\n\t"
                    if ((!(-32768 < -1)) || (true && true)) {
                        v6 = v6
                        var v80 float64 = (0.0 / (-2.25 * 1.5))
                        _ = v80
                    }
                }
            }
            if (!(!(!false))) {
                var v81 uint32 = 2147483647
                switch ("hello" + "a") {
                case "foo bar":
                    fmt.Println(v40, v81)
                case "\n\t":
                    var v82 uint = v40
                    _ = v82
                default:
                    fmt.Println(v37, v81, v40)
                }
                _ = len(v5)
                _ = v81
            }
        } else {
            if len(v5) > 0 {
                v5[0] = 576
            }
            switch 42 {
            case 9223372036854775807:
                for fi83 := 0; fi83 < 4; fi83++ {
                    v84 := GenS5{F0: 0, F1: 761, F2: 2147483647}
                    fmt.Println(v3, v37, v1)
                    v85 := GenS11{F0: 4294967295}
                    _ = fi83
                    _ = v84
                    _ = v85
                }
            case 1:
                v86 := GenS12{F0: 100, F1: -2147483648}
                _ = v86
            case -446:
                fmt.Println(v6, v6, v37)
            }
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
        }
        var v87 float32 = 1.5
        _ = v87
    } else {
        var v88 int = 127
        _ = v88
    }
    var v89 int16 = -753
    v4["a"] = ""
    var v90 bool = true
    v40 += (v40 & v40)
    var v91 bool = (genFn1(443, "", v40) < 42)
    var v92 string = ("a" + ((("" + (("foo bar" + "foo bar") + "")) + "hello") + ("a" + "foo bar")))
    if v90 {
        v91 = (v1 < float32((((-2.25 * ((-2.25 / 1e10) + 1.5)) + ((1.5 + 3.14159) - (1.5 * 1e10))) * 1e10)))
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v89:", v89)
    fmt.Println("v90:", v90)
    fmt.Println("v91:", v91)
    fmt.Println("v92:", v92)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v4
    _ = v40
    _ = v5
    _ = v6
    _ = v89
    _ = v90
    _ = v91
    _ = v92
}
