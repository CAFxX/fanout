package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 string }
type GenS2 struct { F0 uint8; F1 uint32; F2 uint }
type GenS3 struct { F0 uint64; F1 int; F2 uint64 }
type GenS4 struct { F0 uint32; F1 uint; F2 int }
type GenS5 struct { F0 bool; F1 int; F2 int16 }
type GenS6 struct { F0 uint64; F1 int8 }
type GenS7 struct { F0 int8 }
type GenS8 struct { F0 int; F1 uint }
type GenS9 struct { F0 int32; F1 int8 }

func main() {
    v1 := []int{32768, -233}
    v2 := make(map[int]uint32)
    var v3 int64 = 0
    var v4 string = "\n\t"
    v5 := make(map[int]uint32)
    if true {
        var v6 uint32 = 707
        delete(v2, 32767)
        _ = v6
    } else {
        var v7 float32 = 0.0
        _ = v7
    }
    var v8 uint = 482
    switch 32768 {
    case -9223372036854775808:
        v8 = v8
    case 256:
        for fi9 := 0; fi9 < 8; fi9++ {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            if false {
                v10 := GenS1{F0: 18446744073709551615, F1: (("\n\t" + v4) + "a")}
                var v11 int32 = 0
                var v12 string = v4
                _ = v10
                _ = v11
                _ = v12
            } else {
                delete(v2, 32767)
                for fi13 := 0; fi13 < 13; fi13++ {
                    if true {
                        if false {
                            v3 += (v3 % (((-(v3))) | 1))
                            var v14 uint32 = 973
                            var v15 int64 = v3
                            _ = v14
                            _ = v15
                        }
                        v3 = (v3 % ((-1) | 1))
                    }
                    var v16 int64 = (v3 / (((v3 ^ 0)) | 1))
                    _ = fi13
                    _ = v16
                }
            }
            v17 := GenS2{F0: 127, F1: 505, F2: v8}
            _ = fi9
            _ = v17
        }
    case -1:
        v18 := GenS3{F0: 4294967295, F1: 32767, F2: 465}
        _ = v18
    default:
        if false {
            var v19 int64 = (v3 ^ (v3 | (v3 * (v3 ^ 1))))
            v8 = v8
            _ = v19
        } else {
            fmt.Println(v4)
        }
    }
    if true {
        v20 := GenS4{F0: 1, F1: (v8 >> (((v8 & v8)) & 63)), F2: 32767}
        v21 := GenS4{F0: 77, F1: 4294967295, F2: 2147483647}
        _ = v20
        _ = v21
    }
    v1 = append(v1, 32767)
    delete(v5, 1)
    if len(v1) > 0 {
        v1[0] = 127
    }
    switch 32768 {
    case 96:
        v22 := GenS5{F0: (-2.25 > (1e10 * 1.5)), F1: -1, F2: 1}
        _ = v22
    default:
        var v23 int = 4294967295
        _ = v23
    }
    switch "\n\t" {
    case "foo bar":
        if false {
            _ = len(v1)
        }
    default:
        var v24 float64 = (0.1 * 0.1)
        _ = v24
    }
    v4 = ((("hello" + "\n\t") + "") + "\n\t")
    var v25 float32 = ((0.1 + -2.25) + 0.0)
    if (-9223372036854775808 != 9223372036854775807) {
        if true {
            var v26 int16 = -1
            _ = v26
        } else {
            var v27 uint = v8
            _ = v27
        }
    } else {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        var v28 int8 = -1
        _ = v28
    }
    v2[1] = 158
    var v29 int8 = -1
    for fi30 := 0; fi30 < 19; fi30++ {
        v29 = (v29 & -128)
        switch (fi30 + fi30) {
        case -61:
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        case 9223372036854775807:
            v31 := GenS5{F0: false, F1: (-(fi30)), F2: -1}
            _ = v31
        }
        if ((v25 * 3.14159) < (((1e10 / float32(-2.25)) * 0.0) + float32((0.1 / 3.14159)))) {
            for fi32 := 0; fi32 < 12; fi32++ {
                v25 += (0.1 + float32(float64(v25)))
                _ = fi32
            }
        } else {
            if false {
                v33 := GenS2{F0: 224, F1: 2, F2: 42}
                for fi34 := 0; fi34 < 12; fi34++ {
                    switch (fi30 - (fi34 << ((18446744073709551615) & 63))) {
                    case -1:
                        switch "a" {
                        case "":
                            if true {
                                v25 = float32(-2.25)
                                if false {
                                    v4 = v4
                                    if true {
                                        if (v29 > 40) {
                                            v29 = -1
                                            delete(v2, (fi34 & fi30))
                                        }
                                        if (fi30 <= (fi34 >> ((242) & 63))) {
                                            v2[(fi34 << (((v8 | (^(v8)))) & 63))] = 993
                                        } else {
                                            var v35 int16 = 247
                                            v36 := GenS6{F0: uint64(v8), F1: (v29 % ((v29) | 1))}
                                            v5[-9223372036854775808] = 65535
                                            _ = v35
                                            _ = v36
                                        }
                                    }
                                    if true {
                                        var v37 int32 = -2147483648
                                        v37 = (v37 + v37)
                                        var v38 int16 = -779
                                        fmt.Println(fi34, v29)
                                        _ = v37
                                        _ = v38
                                    } else {
                                        switch "\n\t" {
                                        case "\n\t":
                                            if false {
                                                if (v29 > v29) {
                                                    v8 = v8
                                                } else {
                                                    fmt.Println(v3, v8)
                                                    v4 = (v4 + (v4 + v4))
                                                    var v39 uint64 = 399
                                                    _ = v39
                                                }
                                                if len(v1) >= 2 {
                                                    v1 = v1[1:len(v1)]
                                                }
                                                if (("foo bar" + ("foo bar" + v4)) > "a") {
                                                    switch "\n\t" {
                                                    case "hello":
                                                        var v40 uint8 = 127
                                                        _ = v40
                                                    default:
                                                        v3 -= (v3 % ((-1) | 1))
                                                    }
                                                }
                                            } else {
                                                switch v4 {
                                                case "":
                                                    v25 += 1e10
                                                case "\n\t":
                                                    v2[(fi30 & 108)] = 1
                                                case "hello":
                                                    if (true && (float64(0.1) != float64((3.14159 * 0.0)))) {
                                                        if (2147483647 == 1) {
                                                            v29 -= v29
                                                            fmt.Println(v29, v25)
                                                        } else {
                                                            if len(v1) >= 2 {
                                                                v1 = v1[1:len(v1)]
                                                            }
                                                            var v41 bool = (545 != (fi30 & (fi30 - (fi34 >> ((v8) & 63)))))
                                                            _ = v41
                                                        }
                                                        v42 := GenS7{F0: v29}
                                                        _ = v42
                                                    }
                                                }
                                                if false {
                                                    if true {
                                                        v8 -= (v8 + v8)
                                                    } else {
                                                        v43 := GenS8{F0: -1, F1: (v8 & 18446744073709551615)}
                                                        _ = v43
                                                    }
                                                    v44 := GenS8{F0: (fi30 & fi30), F1: (v8 / (((v8 + (v8 | v8))) | 1))}
                                                    fmt.Println(v4)
                                                    _ = v44
                                                } else {
                                                    delete(v2, fi34)
                                                }
                                                _ = len(v1)
                                            }
                                        }
                                    }
                                } else {
                                    switch "" {
                                    case "a":
                                        v45 := GenS9{F0: -1, F1: v29}
                                        _ = v45
                                    case "foo bar":
                                        var v46 int64 = v3
                                        _ = v46
                                    case "\n\t":
                                        if len(v1) > 0 {
                                            v1[0] = fi30
                                        }
                                    default:
                                        fmt.Println(v29)
                                    }
                                    switch (fi30 + 4294967295) {
                                    case 797:
                                        _ = len(v1)
                                    case -656:
                                        var v47 uint = (v8 & v8)
                                        _ = v47
                                    case 32768:
                                        switch fi34 {
                                        case 816:
                                            var v48 bool = true
                                            _ = v48
                                        case 2147483647:
                                            var v49 uint16 = 451
                                            _ = v49
                                        case -2147483648:
                                            v50 := GenS9{F0: 42, F1: (v29 >> ((4294967295) & 63))}
                                            _ = v50
                                        default:
                                            var v51 int16 = 1000
                                            _ = v51
                                        }
                                    default:
                                        switch fi30 {
                                        case 1:
                                            v52 := v2[(fi34 ^ -31)]
                                            _ = v52
                                        case 891:
                                            v4 = ((v4 + "a") + "\n\t")
                                        default:
                                            v2[379] = 2147483647
                                        }
                                    }
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                }
                                delete(v2, (fi34 & fi30))
                            }
                        default:
                            var v53 int32 = -2147483648
                            _ = v53
                        }
                    default:
                        if false {
                            v54 := GenS8{F0: fi34, F1: (v8 % ((682) | 1))}
                            v25 -= (v25 / 1.5)
                            var v55 uint32 = 2
                            _ = v54
                            _ = v55
                        }
                    }
                    v56 := GenS2{F0: 0, F1: 809, F2: v8}
                    switch v4 {
                    case "":
                        v29 = (-(v29))
                    case "\n\t":
                        fmt.Println(v29, v8, v4)
                    }
                    var v57 uint16 = 32767
                    v57 = (v57 * 0)
                    if true {
                        var v58 float32 = 0.1
                        if (!((v4 >= v4) && false)) {
                            if (4294967295 >= 4294967295) {
                                v3 += v3
                                v59 := GenS4{F0: 2147483647, F1: (v8 + (v8 - (v8 % (((v8 % ((v8) | 1))) | 1)))), F2: 127}
                                var v60 uint16 = 1
                                _ = v59
                                _ = v60
                            }
                            switch "foo bar" {
                            case "foo bar":
                                if false {
                                    var v61 int = fi34
                                    v62 := GenS6{F0: 0, F1: v29}
                                    var v63 int16 = 399
                                    _ = v61
                                    _ = v62
                                    _ = v63
                                } else {
                                    v3 += int64(v29)
                                    fmt.Println(v29, fi30)
                                    if true {
                                        v64 := v2[fi34]
                                        _ = v64
                                    } else {
                                        v65 := GenS8{F0: 417, F1: 18446744073709551615}
                                        fmt.Println(v25)
                                        _ = v65
                                    }
                                    var v66 uint8 = 1
                                    _ = v66
                                }
                            case "\n\t":
                                fmt.Println(fi30)
                            case "":
                                switch ("" + v4) {
                                case "\n\t":
                                    v8 = 42
                                default:
                                    _ = len(v2)
                                }
                            }
                            switch v4 {
                            case "":
                                switch (fi30 | fi30) {
                                case 4294967295:
                                    fmt.Println(v29)
                                case -2147483648:
                                    v8 -= 889
                                }
                            case "\n\t":
                                v4 = v4
                            }
                        }
                        _ = v58
                    }
                    _ = fi34
                    _ = v56
                    _ = v57
                }
                _ = v33
            }
        }
        _ = fi30
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v25
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
