package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint8; F1 uint8 }
type GenS3 struct { F0 int; F1 string; F2 uint64 }
type GenS4 struct { F0 uint32; F1 int8 }
type GenS5 struct { F0 bool; F1 uint32; F2 uint64 }
type GenS6 struct { F0 bool; F1 int32; F2 int16 }
type GenS7 struct { F0 uint32; F1 int64; F2 int }
type GenS8 struct { F0 int8; F1 int; F2 uint16 }
type GenS9 struct { F0 int8; F1 uint32; F2 int16 }
func genFn0(p0 uint32, p1 int) int64 {
    for fi5 := 0; fi5 < 5; fi5++ {
        for fi6 := 0; fi6 < 16; fi6++ {
            var v7 uint16 = 184
            v8 := GenS1{F0: 42}
            _ = fi6
            _ = v7
            _ = v8
        }
        _ = fi5
    }
    p1 = (p1 << ((18446744073709551615) & 63))
    for fi9 := 0; fi9 < 17; fi9++ {
        for fi10 := 0; fi10 < 8; fi10++ {
            var v11 string = ("hello" + "hello")
            _ = fi10
            _ = v11
        }
        if true {
            p1 += p1
            v12 := GenS1{F0: 1}
            _ = v12
        } else {
            switch ("\n\t" + "hello") {
            case "hello":
                var v13 int64 = 1
                _ = v13
            case "\n\t":
                p1 = (^(p1))
            case "a":
                for fi14 := 0; fi14 < 16; fi14++ {
                    var v15 uint = 401
                    if true {
                        var v16 int32 = -618
                        fmt.Println(fi14, fi14)
                        _ = v16
                    } else {
                        v15 = (v15 / ((v15) | 1))
                    }
                    _ = fi14
                    _ = v15
                }
            }
        }
        for fi17 := 0; fi17 < 19; fi17++ {
            var v18 int = fi9
            var v19 bool = (!true)
            _ = fi17
            _ = v18
            _ = v19
        }
        _ = fi9
    }
    _ = p0
    _ = p1
    return 9223372036854775807
}
func genFn1(p0 uint8, p1 uint, p2 int) int {
    p0 = (p0 & p0)
    var v20 int16 = 1
    switch "\n\t" {
    case "\n\t":
        fmt.Println(v20, p2, v20)
    case "":
        var v21 string = ("a" + "")
        _ = v21
    case "foo bar":
        v20 = 32767
    default:
        v22 := []int{(-(p2)), p2}
        _ = v22
    }
    if (1 >= -128) {
        p2 -= 4294967295
    } else {
        v23 := GenS2{F0: 2, F1: p0}
        fmt.Println(p2, v20, p1)
        if true {
            v24 := GenS3{F0: 32768, F1: "a", F2: 18446744073709551615}
            _ = v24
        }
        _ = v23
    }
    var v25 uint16 = 0
    _ = p0
    _ = p1
    _ = p2
    _ = v20
    _ = v25
    return 42
}

func main() {
    v1 := []uint32{71}
    var v2 int32 = 1
    v3 := []int{}
    var v4 uint32 = 1
    v26 := GenS4{F0: (v4 - v4), F1: -89}
    if (v2 <= 1) {
        v2 = 159
        if true {
            var v27 string = ""
            v27 = v27
            v2 = 0
            _ = v27
        }
    }
    switch genFn1(225, 449, -330) {
    case 32767:
        if true {
            switch -293 {
            case 9223372036854775807:
                v4 = 42
            case 32768:
                var v28 int32 = (^(v2))
                _ = v28
            case -254:
                _ = len(v1)
            default:
                if len(v1) > 0 {
                    v1[0] = 0
                }
            }
            switch 128 {
            case -762:
                v2 += v2
            case 128:
                var v29 uint64 = 816
                _ = v29
            default:
                var v30 int64 = int64(v4)
                _ = v30
            }
            if len(v1) > 0 {
                v1[0] = (v4 & v4)
            }
        }
    case 330:
        var v31 string = "foo bar"
        _ = v31
    case -416:
        var v32 bool = true
        _ = v32
    }
    var v33 int = genFn1(42, 255, 128)
    for fi34 := 0; fi34 < 10; fi34++ {
        for fi35 := 0; fi35 < 6; fi35++ {
            switch 2147483647 {
            case 127:
                if ((14 < (^(v4))) || ("a" <= "foo bar")) {
                    switch (v33 & genFn1(16, 7, genFn1(uint8(v4), 2, 4294967295))) {
                    case 1:
                        v3 = append(v3, fi35)
                    case -943:
                        var v36 int = (fi35 / ((fi34) | 1))
                        _ = v36
                    case 42:
                        var v37 uint32 = 4294967295
                        _ = v37
                    default:
                        v4 += (v4 * v4)
                    }
                    if false {
                        v4 += (v4 & v4)
                        if true {
                            switch (("foo bar" + "foo bar") + "a") {
                            case "":
                                fmt.Println(v33, v33, v33)
                            case "foo bar":
                                var v38 int16 = 42
                                _ = v38
                            default:
                                v2 += int32(v33)
                            }
                            var v39 uint = 320
                            if (!false) {
                                if false {
                                    v40 := GenS5{F0: true, F1: (v4 % (((v4 << ((237) & 63))) | 1)), F2: 0}
                                    v1 = append(v1, 2147483647)
                                    _ = v40
                                }
                            }
                            _ = v39
                        } else {
                            fmt.Println(fi34, v33)
                        }
                        if ((v33 ^ (^(fi35))) != (v33 >> ((42) & 63))) {
                            var v41 uint = 2
                            _ = v41
                        }
                        v42 := GenS4{F0: (v4 % (((v4 % (((^(v4))) | 1))) | 1)), F1: 60}
                        _ = len(v3)
                        _ = v42
                    } else {
                        fmt.Println(fi34, fi35, fi35)
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                    }
                    v4 = (v4 - (v4 - v4))
                }
            default:
                v43 := GenS6{F0: true, F1: v2, F2: 0}
                _ = v43
            }
            _ = fi35
        }
        for fi44 := 0; fi44 < 14; fi44++ {
            v4 += 42
            if false {
                if len(v3) > 0 {
                    v3[0] = 1
                }
                if true {
                    switch "hello" {
                    case "hello":
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        if (!(6 >= 42)) {
                            v4 = v4
                        } else {
                            v45 := GenS3{F0: (fi44 & v33), F1: "a", F2: 6}
                            switch ("\n\t" + "\n\t") {
                            case "hello":
                                fmt.Println(fi34, fi34)
                            }
                            _ = v45
                        }
                    case "\n\t":
                        v33 += fi44
                    case "a":
                        if true {
                            v46 := GenS7{F0: (^(v4)), F1: 42, F2: v33}
                            if ((((^(v4)) == (v4 - (^(v4)))) || true) && true) {
                                var v47 int = (v33 << ((0) & 63))
                                _ = v47
                            } else {
                                if false {
                                    fmt.Println(v2)
                                }
                                _ = len(v3)
                                fmt.Println(fi34, v4)
                            }
                            var v48 bool = (0 > 1)
                            _ = v46
                            _ = v48
                        } else {
                            if (false && false) {
                                v49 := GenS3{F0: genFn1(2, 1, genFn1(0, 65535, (^(fi44)))), F1: ("" + "hello"), F2: 4294967295}
                                fmt.Println(fi34, v4)
                                v4 -= v4
                                _ = v49
                            }
                            v33 -= v33
                        }
                    default:
                        v3 = append(v3, v33)
                        var v50 int8 = 127
                        _ = v50
                    }
                    var v51 int = fi44
                    var v52 int16 = -522
                    _ = v51
                    _ = v52
                } else {
                    v53 := GenS8{F0: 127, F1: (fi44 + (^(fi34))), F2: 504}
                    v33 = 4294967295
                    v54 := GenS2{F0: 0, F1: 247}
                    _ = v53
                    _ = v54
                }
            } else {
                if true {
                    v2 = -1
                    if len(v3) > 0 {
                        v3[0] = fi44
                    }
                } else {
                    var v55 int8 = -106
                    _ = v55
                }
                if len(v1) > 0 {
                    v1[0] = v4
                }
                v3 = append(v3, fi44)
            }
            var v56 uint32 = 594
            _ = fi44
            _ = v56
        }
        _ = fi34
    }
    var v57 float32 = (0.0 / -2.25)
    v3 = append(v3, (v33 | 127))
    v2 += (v2 ^ (v2 - v2))
    v58 := GenS3{F0: (^(v33)), F1: "a", F2: 42}
    var v59 string = "a"
    v57 = (v57 - (float32(float64(v57)) / v57))
    v60 := GenS9{F0: 7, F1: v4, F2: 0}
    v57 = float32(0.0)
    var v61 uint16 = 0
    v2 = v2
    for fi62 := 0; fi62 < 16; fi62++ {
        v59 = (v59 + ("\n\t" + (v59 + "")))
        _ = fi62
    }
    if false {
        switch (v33 & v33) {
        case -9223372036854775808:
            for fi63 := 0; fi63 < 19; fi63++ {
                for fi64 := 0; fi64 < 12; fi64++ {
                    var v65 uint16 = v61
                    _ = fi64
                    _ = v65
                }
                v59 = (((v59 + "foo bar") + "foo bar") + (("a" + ("hello" + v59)) + ("foo bar" + (("a" + ("a" + ("hello" + v59))) + v59))))
                for fi66 := 0; fi66 < 17; fi66++ {
                    var v67 int64 = -898
                    v68 := GenS6{F0: true, F1: (v2 % ((42) | 1)), F2: 919}
                    _ = fi66
                    _ = v67
                    _ = v68
                }
                _ = fi63
            }
        case 9223372036854775807:
            v61 = (v61 << ((2) & 63))
        case 32768:
            var v69 bool = false
            _ = v69
        }
        if len(v1) > 0 {
            v1[0] = v4
        }
    } else {
        if (-652 <= 107) {
            fmt.Println(v59, v2)
            switch (^(v33)) {
            case 32768:
                v70 := GenS3{F0: genFn1(128, 42, (-(v33))), F1: "hello", F2: 4294967295}
                _ = v70
            case 127:
                v33 -= 1
            }
        } else {
            var v71 float32 = (float32(3.14159) / v57)
            var v72 uint = 441
            var v73 float64 = ((0.0 / 1.5) / -2.25)
            _ = v71
            _ = v72
            _ = v73
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v57:", v57)
    fmt.Println("v59:", v59)
    fmt.Println("v61:", v61)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v33
    _ = v4
    _ = v57
    _ = v58
    _ = v59
    _ = v60
    _ = v61
}
