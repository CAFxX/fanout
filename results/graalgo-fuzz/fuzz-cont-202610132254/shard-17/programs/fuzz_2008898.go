package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint; F2 uint32 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint64; F1 string; F2 uint8 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 int64; F1 bool; F2 string }
type GenS7 struct { F0 uint8; F1 string }
type GenS8 struct { F0 uint16 }

func main() {
    var v1 int = -9223372036854775808
    var v2 uint8 = 188
    var v3 int8 = 42
    var v4 int8 = 0
    var v5 int = -529
    var v6 float32 = 1e10
    v1 = (v1 + int(v3))
    for fi7 := 0; fi7 < 18; fi7++ {
        v5 -= v5
        fmt.Println(v3, v5)
        v6 = v6
        _ = fi7
    }
    fmt.Println(v5, v6, v5)
    var v8 int64 = 1
    switch v1 {
    case 42:
        v8 = (-(v8))
    case 0:
        for fi9 := 0; fi9 < 2; fi9++ {
            for fi10 := 0; fi10 < 2; fi10++ {
                v4 = v4
                _ = fi10
            }
            _ = fi9
        }
    case 4294967295:
        fmt.Println(v6)
    default:
        var v11 float32 = v6
        _ = v11
    }
    if true {
        v12 := GenS1{F0: 635, F1: 255, F2: 635}
        v4 -= 1
        v2 = (v2 ^ 2)
        _ = v12
    } else {
        v3 = v3
        switch 32768 {
        case 9223372036854775807:
            v2 += v2
        case 127:
            var v13 int16 = 32767
            _ = v13
        default:
            switch "a" {
            case "a":
                if ("foo bar" == "a") {
                    var v14 int = v1
                    var v15 int16 = 973
                    if true {
                        fmt.Println(v14, v8, v6)
                    } else {
                        v6 -= v6
                        v16 := GenS2{F0: 32767}
                        _ = v16
                    }
                    _ = v14
                    _ = v15
                }
            case "hello":
                v1 = (^(v5))
            case "":
                v17 := GenS2{F0: 0}
                _ = v17
            default:
                v4 -= 0
            }
        }
    }
    if false {
        v18 := []int{(v1 + (^(v1)))}
        if (382 != 668) {
            var v19 int32 = 42
            _ = v19
        } else {
            var v20 int64 = 9223372036854775807
            if true {
                for fi21 := 0; fi21 < 6; fi21++ {
                    v18 = append(v18, v5)
                    _ = fi21
                }
                v6 -= (v6 / (float32(0.0) + 1.5))
            } else {
                var v22 uint64 = 0
                _ = v22
            }
            _ = v20
        }
        _ = v18
    } else {
        v23 := GenS1{F0: 42, F1: 42, F2: 2147483648}
        if true {
            for fi24 := 0; fi24 < 2; fi24++ {
                v2 += (^(v2))
                _ = fi24
            }
        } else {
            v4 -= (v3 & (^(v4)))
            switch 384 {
            case 4294967295:
                v25 := GenS1{F0: 262, F1: 792, F2: 565}
                _ = v25
            case -87:
                fmt.Println(v6)
            }
            switch (v5 & (v1 - v1)) {
            case 565:
                var v26 bool = (601 >= 1)
                _ = v26
            case 679:
                if (42 <= -1) {
                    var v27 uint16 = 42
                    for fi28 := 0; fi28 < 1; fi28++ {
                        if false {
                            if (1 == v4) {
                                v1 = -1
                            }
                            if (true && (!(true && false))) {
                                if false {
                                    for fi29 := 0; fi29 < 6; fi29++ {
                                        if (!(false || false)) {
                                            if (false && true) {
                                                v2 = v2
                                            }
                                        }
                                        var v30 int = fi29
                                        fmt.Println(v4)
                                        _ = fi29
                                        _ = v30
                                    }
                                    v2 += (^(v2))
                                }
                            }
                        } else {
                            for fi31 := 0; fi31 < 14; fi31++ {
                                var v32 int = (v5 | fi28)
                                v5 = v1
                                _ = fi31
                                _ = v32
                            }
                            v5 = v1
                            for fi33 := 0; fi33 < 12; fi33++ {
                                v34 := []int{(-(v5)), (v1 - fi33), v1, (-(fi28))}
                                _ = fi33
                                _ = v34
                            }
                        }
                        _ = fi28
                    }
                    _ = v27
                } else {
                    var v35 uint8 = 42
                    v1 = (-(v5))
                    v36 := []int{}
                    _ = v35
                    _ = v36
                }
            case 32768:
                for fi37 := 0; fi37 < 12; fi37++ {
                    if (true || (!false)) {
                        if (!(597 < 2147483648)) {
                            if (!(255 == 0)) {
                                v1 = 116
                                for fi38 := 0; fi38 < 15; fi38++ {
                                    var v39 string = "hello"
                                    fmt.Println(v2)
                                    _ = fi38
                                    _ = v39
                                }
                            }
                            var v40 string = "hello"
                            if false {
                                for fi41 := 0; fi41 < 13; fi41++ {
                                    v4 = (v4 | (v4 / (((v3 ^ v4)) | 1)))
                                    _ = fi41
                                }
                                v42 := GenS3{F0: 167, F1: "foo bar", F2: 128}
                                v43 := GenS1{F0: 236, F1: 0, F2: 2147483648}
                                _ = v42
                                _ = v43
                            } else {
                                v4 = (v4 ^ 42)
                                fmt.Println(v40)
                            }
                            _ = v40
                        } else {
                            for fi44 := 0; fi44 < 13; fi44++ {
                                fmt.Println(v6, v5)
                                _ = fi44
                            }
                            var v45 int = v1
                            _ = v45
                        }
                        v46 := []int{-172, (-(v5)), fi37}
                        v47 := GenS4{F0: 2147483647}
                        _ = v46
                        _ = v47
                    } else {
                        if (42 <= 0) {
                            v48 := []int{(-(v5))}
                            v49 := GenS2{F0: 829}
                            _ = v48
                            _ = v49
                        }
                        v50 := GenS3{F0: 42, F1: "\n\t", F2: (^(v2))}
                        var v51 int32 = 1
                        _ = v50
                        _ = v51
                    }
                    v6 = v6
                    v52 := []int{}
                    _ = fi37
                    _ = v52
                }
            default:
                v53 := GenS5{F0: 32767}
                _ = v53
            }
        }
        _ = v23
    }
    v6 = v6
    if true {
        v6 += float32(0.0)
        if ((0.0 * ((0.1 * float64(0.1)) * (1.5 - (1e10 / 1e10)))) >= -2.25) {
            v54 := []int{128, v1, v1, 128}
            _ = v54
        }
    } else {
        v55 := GenS6{F0: -823, F1: false, F2: "foo bar"}
        var v56 int32 = 0
        _ = v55
        _ = v56
    }
    v8 = 9223372036854775807
    for fi57 := 0; fi57 < 1; fi57++ {
        var v58 int16 = 32767
        if false {
            v4 = (^(v3))
            if false {
                switch (^(v5)) {
                case 127:
                    v2 = v2
                case 4294967295:
                    var v59 bool = true
                    _ = v59
                case -798:
                    fmt.Println(v3, v58)
                default:
                    fmt.Println(v4, v1)
                }
            } else {
                for fi60 := 0; fi60 < 16; fi60++ {
                    var v61 int = v5
                    v62 := GenS2{F0: 710}
                    v58 += -1
                    _ = fi60
                    _ = v61
                    _ = v62
                }
                fmt.Println(v6, fi57)
                fmt.Println(v5)
            }
            for fi63 := 0; fi63 < 14; fi63++ {
                if false {
                    v5 = fi57
                } else {
                    if (!false) {
                        v4 += v4
                        if (2 != 2) {
                            v1 = (fi57 * (v5 + (fi63 << ((1) & 63))))
                            switch -1 {
                            case 32768:
                                if (false || (v2 == v2)) {
                                    fmt.Println(v6, v5, v5)
                                    v64 := GenS7{F0: (^(v2)), F1: ("hello" + ((("a" + "") + ("a" + "hello")) + ""))}
                                    v65 := []int{v1, v5, fi57}
                                    _ = v64
                                    _ = v65
                                }
                            case -1:
                                if ((v8 ^ v8) != 0) {
                                    if false {
                                        if (v6 >= (float32((0.1 * (float64(float32(0.0)) + (3.14159 + (0.0 - 0.1))))) * 0.1)) {
                                            var v66 uint16 = 78
                                            _ = v66
                                        }
                                        v2 -= v2
                                    } else {
                                        v58 += v58
                                        if false {
                                            v67 := GenS8{F0: 65535}
                                            if (18446744073709551615 < 2) {
                                                fmt.Println(fi63)
                                                fmt.Println(v2, v5, v58)
                                                v68 := []int{v1, (v5 + (fi57 % (((^(fi63))) | 1))), v1}
                                                _ = v68
                                            } else {
                                                v4 -= (-(v3))
                                            }
                                            _ = v67
                                        }
                                    }
                                } else {
                                    v6 = v6
                                    v69 := []int{42}
                                    _ = v69
                                }
                            case 42:
                                v58 += v58
                            }
                            fmt.Println(v6, v3, v58)
                        } else {
                            v70 := GenS8{F0: 622}
                            if (v58 != (v58 / ((v58) | 1))) {
                                var v71 int = fi63
                                _ = v71
                            }
                            _ = v70
                        }
                    }
                }
                switch "foo bar" {
                case "a":
                    v6 = v6
                }
                if false {
                    v58 -= (v58 & (v58 << ((408) & 63)))
                }
                _ = fi63
            }
        }
        _ = fi57
        _ = v58
    }
    if false {
        v5 = v1
        v72 := []int{(v1 & (-(v1))), (v5 & int(v2)), (v5 | (v5 >> ((663) & 63))), v5}
        switch "\n\t" {
        case "":
            var v73 uint16 = 42
            _ = v73
        case "a":
            if len(v72) >= 2 {
                v72 = v72[1:len(v72)]
            }
        case "\n\t":
            for fi74 := 0; fi74 < 3; fi74++ {
                v4 += 62
                if (2 < 4294967295) {
                    v6 -= v6
                    var v75 int32 = -1
                    _ = v75
                } else {
                    var v76 uint16 = 255
                    _ = v76
                }
                v8 += (v8 - (v8 | (v8 & v8)))
                _ = fi74
            }
        }
        _ = v72
    }
    var v77 int16 = 1
    v78 := []int{(v1 + 32768), (v1 ^ (v1 << ((0) & 63))), v1}
    v79 := GenS6{F0: v8, F1: false, F2: "foo bar"}
    for fi80 := 0; fi80 < 8; fi80++ {
        switch fi80 {
        case 0:
            v77 += v77
        case -2147483648:
            if (!false) {
                v77 = v77
            } else {
                var v81 int64 = v8
                _ = v81
            }
        case 128:
            fmt.Println(fi80, v8)
        default:
            if len(v78) > 0 {
                v78[0] = (v1 | v5)
            }
        }
        var v82 int8 = v4
        var v83 int16 = (^(v77))
        _ = fi80
        _ = v82
        _ = v83
    }
    var v84 int16 = v77
    v5 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v77:", v77)
    fmt.Println("slicelen v78:", len(v78), v78)
    fmt.Println("v8:", v8)
    fmt.Println("v84:", v84)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v77
    _ = v78
    _ = v79
    _ = v8
    _ = v84
}
