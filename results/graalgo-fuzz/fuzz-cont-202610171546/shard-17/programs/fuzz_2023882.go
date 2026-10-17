package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint8 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint; F1 int16 }
type GenS4 struct { F0 string; F1 uint }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 int64; F1 int32 }
type GenS7 struct { F0 bool; F1 uint16 }
type GenS8 struct { F0 int8; F1 bool }
type GenS9 struct { F0 uint16; F1 uint32; F2 int16 }
type GenS10 struct { F0 int64 }
func genFn0(p0 uint32, p1 int8, p2 bool) uint {
    var v5 float32 = 1.5
    for fi6 := 0; fi6 < 20; fi6++ {
        var v7 uint16 = 0
        v8 := []int{9223372036854775807}
        _ = fi6
        _ = v7
        _ = v8
    }
    var v9 int64 = 9223372036854775807
    var v10 uint64 = 4294967295
    for fi11 := 0; fi11 < 17; fi11++ {
        for fi12 := 0; fi12 < 15; fi12++ {
            var v13 float32 = v5
            v5 = 0.1
            _ = fi12
            _ = v13
        }
        _ = fi11
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v5
    _ = v9
    return 384
}
func genFn1(p0 uint64, p1 uint, p2 uint8) bool {
    var v14 uint32 = 368
    for fi15 := 0; fi15 < 3; fi15++ {
        fmt.Println(p0, p1)
        _ = fi15
    }
    v14 += v14
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    return false
}

func main() {
    var v1 bool = (!true)
    v2 := make(map[int]bool)
    v3 := []string{"foo bar", "a", "\n\t", "\n\t"}
    var v4 int8 = 127
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v16 := GenS1{F0: "", F1: 1}
    v1 = v1
    var v17 uint32 = 65535
    var v18 uint8 = 181
    fmt.Println(v18, v4)
    v19 := GenS2{F0: false}
    for fi20 := 0; fi20 < 1; fi20++ {
        v21 := GenS3{F0: 2, F1: -718}
        v22 := GenS4{F0: "foo bar", F1: 4294967295}
        _ = fi20
        _ = v21
        _ = v22
    }
    v23 := GenS3{F0: 2, F1: -1}
    var v24 uint64 = 820
    var v25 uint32 = (v17 * v17)
    v26 := GenS5{F0: 2147483647}
    if v1 {
        var v27 uint16 = 1
        if v1 {
            for fi28 := 0; fi28 < 18; fi28++ {
                if v1 {
                    v18 -= 251
                    if ((622 < 42) || false) {
                        _ = len(v2)
                        v29 := GenS6{F0: 864, F1: -2147483648}
                        v30 := GenS5{F0: -169}
                        _ = v29
                        _ = v30
                    }
                }
                var v31 int8 = v4
                _ = fi28
                _ = v31
            }
        } else {
            if (false || true) {
                if (1000 > -300) {
                    if ((true || v1) && false) {
                        for fi32 := 0; fi32 < 20; fi32++ {
                            for fi33 := 0; fi33 < 16; fi33++ {
                                switch fi32 {
                                case 4294967295:
                                    fmt.Println(fi32)
                                    v34 := GenS4{F0: "\n\t", F1: 781}
                                    var v35 bool = ((v18 << ((genFn0(0, (-(v4)), (!true))) & 63)) == v18)
                                    _ = v34
                                    _ = v35
                                case 345:
                                    v17 = (v17 % ((v17) | 1))
                                }
                                _ = fi33
                            }
                            fmt.Println(v4, fi32)
                            if (!(v27 == v27)) {
                                v17 = (v17 - (^(v17)))
                            } else {
                                if v1 {
                                    var v36 uint64 = v24
                                    var v37 bool = (-950 < -1)
                                    v17 += v25
                                    _ = v36
                                    _ = v37
                                }
                                if (v1 || v1) {
                                    switch int(v25) {
                                    case 32768:
                                        var v38 string = "\n\t"
                                        _ = v38
                                    case -519:
                                        v39 := GenS7{F0: ((v1 && (v24 != (^(v24)))) && false), F1: v27}
                                        _ = v39
                                    case 2147483647:
                                        var v40 int32 = 1
                                        _ = v40
                                    }
                                } else {
                                    v1 = (v27 != v27)
                                    if (32767 >= 1000) {
                                        v4 = v4
                                    } else {
                                        v41 := GenS8{F0: 1, F1: (-2.25 >= 0.1)}
                                        _ = v41
                                    }
                                    var v42 int32 = 641
                                    _ = v42
                                }
                                if v1 {
                                    switch "a" {
                                    case "foo bar":
                                        v18 -= 255
                                    case "\n\t":
                                        v3 = append(v3, "\n\t")
                                    default:
                                        var v43 uint8 = (v18 & (^(v18)))
                                        _ = v43
                                    }
                                    for fi44 := 0; fi44 < 8; fi44++ {
                                        var v45 float64 = 1e10
                                        v46 := GenS9{F0: (v27 - (v27 >> ((4294967295) & 63))), F1: (v17 / ((v25) | 1)), F2: 1}
                                        if (42 <= -132) {
                                            v17 = (v17 * 4294967295)
                                            _ = len(v3)
                                            if (-32768 < 1) {
                                                v17 += uint32(v27)
                                                v25 -= v17
                                                var v47 int = fi44
                                                _ = v47
                                            }
                                        } else {
                                            v48 := v2[32768]
                                            var v49 string = "foo bar"
                                            var v50 uint = 43
                                            _ = v48
                                            _ = v49
                                            _ = v50
                                        }
                                        _ = fi44
                                        _ = v45
                                        _ = v46
                                    }
                                    if ((-503 != 1000) && (!((0.0 < 0.0) || false))) {
                                        var v51 uint16 = (v27 << ((1) & 63))
                                        var v52 uint = 1
                                        v53 := GenS7{F0: (467 == (fi32 + -1)), F1: (v51 % (((v27 << (((v52 * 0)) & 63))) | 1))}
                                        _ = v51
                                        _ = v52
                                        _ = v53
                                    }
                                } else {
                                    fmt.Println(fi32, v4, v24)
                                    if true {
                                        v17 += (v17 ^ uint32(v24))
                                        _ = len(v3)
                                        if (9223372036854775807 != -286) {
                                            if true {
                                                for fi54 := 0; fi54 < 16; fi54++ {
                                                    switch ("foo bar" + "a") {
                                                    case "":
                                                        v55 := GenS2{F0: (v1 || v1)}
                                                        _ = v55
                                                    case "\n\t":
                                                        if ("a" < "\n\t") {
                                                            v25 = (v17 & 65535)
                                                            switch (-(fi54)) {
                                                            case 415:
                                                                v56 := v2[32768]
                                                                _ = v56
                                                            case -250:
                                                                v57 := GenS3{F0: 4294967295, F1: -1}
                                                                _ = v57
                                                            case 32767:
                                                                if v1 {
                                                                    var v58 uint64 = 4294967295
                                                                    _ = v58
                                                                }
                                                            default:
                                                                if (-32768 < 658) {
                                                                    v27 = v27
                                                                    v59 := GenS9{F0: v27, F1: (v17 / (((v25 - v25)) | 1)), F2: -974}
                                                                    _ = v59
                                                                }
                                                            }
                                                        } else {
                                                            v1 = v1
                                                        }
                                                    case "foo bar":
                                                        switch (fi54 & fi54) {
                                                        case -726:
                                                            if len(v3) > 0 {
                                                                v3[0] = "hello"
                                                            }
                                                            v17 += (v17 << ((65535) & 63))
                                                            v4 -= -13
                                                        }
                                                    default:
                                                        v3 = append(v3, ("a" + "\n\t"))
                                                    }
                                                    v60 := GenS10{F0: 9223372036854775807}
                                                    v24 = v24
                                                    var v61 float64 = (0.0 + 3.14159)
                                                    _ = fi54
                                                    _ = v60
                                                    _ = v61
                                                }
                                                v4 -= v4
                                                v4 -= 42
                                            } else {
                                                var v62 uint32 = v17
                                                if len(v3) >= 2 {
                                                    v3 = v3[1:len(v3)]
                                                }
                                                _ = v62
                                            }
                                            switch fi32 {
                                            case 128:
                                                if (!((288 >= (v27 << ((867) & 63))) || v1)) {
                                                    for fi63 := 0; fi63 < 5; fi63++ {
                                                        if len(v3) > 0 {
                                                            v3[0] = "\n\t"
                                                        }
                                                        v25 = v25
                                                        _ = fi63
                                                    }
                                                }
                                            case 1:
                                                for fi64 := 0; fi64 < 10; fi64++ {
                                                    if ((!(((-2.25 == (float64(1.5) - (1.5 + (0.1 + 3.14159)))) || v1) || (((v24 | uint64(v18)) == 42) || v1))) || v1) {
                                                        v24 -= v24
                                                    }
                                                    _ = fi64
                                                }
                                            case 2147483647:
                                                for fi65 := 0; fi65 < 1; fi65++ {
                                                    v24 = (^(v24))
                                                    _ = fi65
                                                }
                                            default:
                                                switch (("hello" + "") + ("foo bar" + "foo bar")) {
                                                case "\n\t":
                                                    var v66 string = ((("\n\t" + ("" + "\n\t")) + (("" + "a") + "a")) + ("\n\t" + ("\n\t" + ("\n\t" + "\n\t"))))
                                                    _ = v66
                                                case "a":
                                                    if (!v1) {
                                                        v1 = (0.0 >= -2.25)
                                                        var v67 uint16 = 333
                                                        if (v4 <= v4) {
                                                            v1 = v1
                                                        } else {
                                                            for fi68 := 0; fi68 < 15; fi68++ {
                                                                fmt.Println(v25)
                                                                _ = fi68
                                                            }
                                                        }
                                                        _ = v67
                                                    } else {
                                                        var v69 float64 = 0.1
                                                        _ = v69
                                                    }
                                                default:
                                                    fmt.Println(v4, v1)
                                                }
                                            }
                                            v18 = (v18 / ((v18) | 1))
                                        } else {
                                            v25 += 542
                                            delete(v2, (fi32 >> ((42) & 63)))
                                            v27 = v27
                                        }
                                    }
                                    v1 = (!(!v1))
                                }
                            }
                            _ = fi32
                        }
                    } else {
                        for fi70 := 0; fi70 < 6; fi70++ {
                            v1 = (2 == 18446744073709551615)
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            _ = fi70
                        }
                    }
                    if v1 {
                        switch 0 {
                        case -9223372036854775808:
                            v1 = (0.1 < (0.0 - 1.5))
                        case -422:
                            v71 := v2[9223372036854775807]
                            _ = v71
                        case 9223372036854775807:
                            if len(v3) > 0 {
                                v3[0] = "hello"
                            }
                        }
                    }
                }
                _ = len(v2)
            }
        }
        v24 = v24
        _ = v27
    }
    var v72 uint8 = v18
    v18 = (v18 & 1)
    switch 42 {
    case 42:
        if v1 {
            var v73 int64 = 9223372036854775807
            _ = len(v2)
            v24 = (v24 + (v24 << ((2) & 63)))
            _ = v73
        } else {
            v4 -= 100
            fmt.Println(v4, v4, v25)
        }
    case 127:
        v3 = append(v3, ("" + ""))
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v72:", v72)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v72
}
