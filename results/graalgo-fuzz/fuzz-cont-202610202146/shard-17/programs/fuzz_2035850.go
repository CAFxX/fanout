package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool; F2 int32 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint8; F1 uint64 }
type GenS5 struct { F0 int8; F1 uint32 }
type GenS6 struct { F0 bool; F1 uint16 }
type GenS7 struct { F0 bool; F1 bool; F2 int8 }
type GenS8 struct { F0 int }
type GenS9 struct { F0 int32 }
func genFn0(p0 int, p1 int64) uint {
    p1 += p1
    v4 := GenS1{F0: 1000, F1: (-455 != p1), F2: -2147483648}
    _ = p0
    _ = p1
    _ = v4
    return 0
}
func genFn1() uint16 {
    v5 := GenS1{F0: 594, F1: false, F2: 173}
    var v6 int16 = -216
    _ = v5
    _ = v6
    return 363
}

func main() {
    v1 := make(map[string]uint32)
    v2 := []int{32767, -9223372036854775808}
    v3 := make(map[int]string)
    var v7 int8 = 1
    if (!(-2.25 > float32(1e10))) {
        for fi8 := 0; fi8 < 5; fi8++ {
            v7 = v7
            v9 := GenS1{F0: -738, F1: true, F2: -2147483648}
            _ = fi8
            _ = v9
        }
        v10 := GenS1{F0: -32768, F1: false, F2: 0}
        if true {
            fmt.Println(v7)
        }
        _ = v10
    }
    v3[128] = "\n\t"
    var v11 uint32 = 1
    if (!false) {
        for fi12 := 0; fi12 < 18; fi12++ {
            v7 -= v7
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            _ = fi12
        }
        v2 = append(v2, -83)
        switch "hello" {
        case "\n\t":
            for fi13 := 0; fi13 < 9; fi13++ {
                if ((-781 != int32(fi13)) && false) {
                    fmt.Println(v11)
                } else {
                    v11 = (v11 ^ 2147483647)
                    v14 := GenS1{F0: 32767, F1: ((true && (false && true)) || ((1000 > 1) && (false && false))), F2: 42}
                    v7 -= v7
                    _ = v14
                }
                var v15 float64 = 0.0
                for fi16 := 0; fi16 < 14; fi16++ {
                    fmt.Println(v7)
                    _ = fi16
                }
                _ = fi13
                _ = v15
            }
        default:
            for fi17 := 0; fi17 < 12; fi17++ {
                if (!((fi17 - 666) < fi17)) {
                    for fi18 := 0; fi18 < 3; fi18++ {
                        var v19 int32 = 42
                        var v20 float64 = 3.14159
                        _ = fi18
                        _ = v19
                        _ = v20
                    }
                    var v21 float32 = ((-2.25 * 0.0) + (float32(0.0) - 3.14159))
                    v11 = v11
                    _ = v21
                } else {
                    if ((0 != 18446744073709551615) || ((^(v7)) >= (v7 ^ (^(v7))))) {
                        _ = len(v3)
                        v22 := GenS1{F0: 32767, F1: true, F2: 2147483647}
                        v23 := GenS2{F0: 567}
                        _ = v22
                        _ = v23
                    } else {
                        for fi24 := 0; fi24 < 13; fi24++ {
                            var v25 int32 = 105
                            if true {
                                v7 = v7
                                if (false || true) {
                                    var v26 uint32 = 42
                                    _ = v26
                                } else {
                                    var v27 uint64 = 0
                                    v25 = (v25 & (v25 >> ((65535) & 63)))
                                    _ = len(v2)
                                    v25 += (v25 % (((v25 ^ (v25 % ((v25) | 1)))) | 1))
                                    _ = v27
                                }
                            } else {
                                v11 -= v11
                                v25 -= (v25 + (v25 >> ((1) & 63)))
                            }
                            v25 = (v25 % (((v25 % ((v25) | 1))) | 1))
                            _ = fi24
                            _ = v25
                        }
                        v7 = 1
                    }
                }
                switch -2147483648 {
                case -9223372036854775808:
                    if false {
                        switch "foo bar" {
                        case "foo bar":
                            var v28 int16 = -1
                            _ = v28
                        case "":
                            fmt.Println(v11, fi17, v7)
                        case "a":
                            if true {
                                v29 := GenS2{F0: -361}
                                _ = v29
                            }
                        }
                        switch ("\n\t" + (("\n\t" + "") + "")) {
                        case "":
                            var v30 float32 = -2.25
                            _ = v30
                        case "hello":
                            for fi31 := 0; fi31 < 1; fi31++ {
                                fmt.Println(v7)
                                v32 := GenS2{F0: -32768}
                                var v33 string = (("" + "\n\t") + "\n\t")
                                _ = fi31
                                _ = v32
                                _ = v33
                            }
                        case "\n\t":
                            switch fi17 {
                            case 32767:
                                if (fi17 >= fi17) {
                                    v11 = 1
                                    switch fi17 {
                                    case 32768:
                                        fmt.Println(v11)
                                    case 2147483647:
                                        fmt.Println(v7, v7)
                                    }
                                } else {
                                    var v34 int32 = 435
                                    var v35 int32 = v34
                                    if false {
                                        v36 := GenS2{F0: 32767}
                                        _ = v36
                                    } else {
                                        for fi37 := 0; fi37 < 15; fi37++ {
                                            fmt.Println(v35, v35, v35)
                                            v38 := GenS1{F0: 1, F1: false, F2: v35}
                                            v34 += (v35 & v34)
                                            _ = fi37
                                            _ = v38
                                        }
                                        var v39 int64 = 980
                                        _ = v39
                                    }
                                    _ = v34
                                    _ = v35
                                }
                            case 127:
                                var v40 string = (("\n\t" + "a") + (("a" + ("a" + "foo bar")) + ""))
                                _ = v40
                            }
                        default:
                            if ((-2.25 <= (-2.25 / 1e10)) && false) {
                                var v41 int64 = 56
                                v42 := GenS3{F0: v11}
                                v7 += (v7 >> ((255) & 63))
                                _ = v41
                                _ = v42
                            } else {
                                fmt.Println(v11, fi17, v11)
                                var v43 int8 = 1
                                v2 = append(v2, fi17)
                                _ = v43
                            }
                        }
                    }
                case 1:
                    for fi44 := 0; fi44 < 14; fi44++ {
                        if len(v2) > 0 {
                            v2[0] = fi44
                        }
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        _ = fi44
                    }
                default:
                    if true {
                        if ((1e10 / ((0.1 + float32((-2.25 * 1.5))) - -2.25)) >= (3.14159 + 1.5)) {
                            switch fi17 {
                            case 9223372036854775807:
                                for fi45 := 0; fi45 < 5; fi45++ {
                                    v7 -= v7
                                    if len(v2) > 0 {
                                        v2[0] = 127
                                    }
                                    v11 = (v11 - (v11 + 2147483648))
                                    _ = fi45
                                }
                            case -9223372036854775808:
                                for fi46 := 0; fi46 < 17; fi46++ {
                                    v11 = v11
                                    _ = fi46
                                }
                            case 538:
                                for fi47 := 0; fi47 < 7; fi47++ {
                                    _ = len(v1)
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                    _ = fi47
                                }
                            }
                            switch "" {
                            case "hello":
                                if (42 >= -2147483648) {
                                    if (-1 < 58) {
                                        _ = len(v2)
                                    }
                                    switch fi17 {
                                    case -278:
                                        var v48 uint32 = v11
                                        _ = v48
                                    case -691:
                                        v1[""] = v11
                                    case 127:
                                        v49 := GenS4{F0: 1, F1: 4294967295}
                                        _ = v49
                                    default:
                                        for fi50 := 0; fi50 < 3; fi50++ {
                                            fmt.Println(v11, fi50)
                                            if false {
                                                v3[fi50] = "hello"
                                                v7 -= v7
                                            } else {
                                                switch (fi50 - (fi17 & fi17)) {
                                                case -2147483648:
                                                    v7 = v7
                                                case 32767:
                                                    switch fi17 {
                                                    case 9223372036854775807:
                                                        v11 -= v11
                                                    }
                                                default:
                                                    switch fi17 {
                                                    case -351:
                                                        switch (fi17 >> ((2) & 63)) {
                                                        case 127:
                                                            var v51 uint = genFn0((fi50 + fi17), 967)
                                                            _ = v51
                                                        }
                                                    case 32767:
                                                        fmt.Println(fi17)
                                                        v11 = 2147483647
                                                    case 409:
                                                        if true {
                                                            fmt.Println(fi17, v7, fi17)
                                                        } else {
                                                            fmt.Println(fi50, fi17, fi50)
                                                        }
                                                    default:
                                                        var v52 int16 = 42
                                                        _ = v52
                                                    }
                                                }
                                                fmt.Println(v11)
                                            }
                                            _ = fi50
                                        }
                                    }
                                }
                            case "":
                                if (float32(0.1) > (1e10 * 1e10)) {
                                    if false {
                                        fmt.Println(v7, v11, fi17)
                                    }
                                    v53 := GenS5{F0: -1, F1: v11}
                                    if false {
                                        if true {
                                            v54 := GenS4{F0: 100, F1: 42}
                                            for fi55 := 0; fi55 < 13; fi55++ {
                                                var v56 uint64 = 1
                                                v57 := GenS2{F0: 42}
                                                var v58 int = (fi55 + fi17)
                                                _ = fi55
                                                _ = v56
                                                _ = v57
                                                _ = v58
                                            }
                                            if (42 == 2) {
                                                v11 = (v11 & v11)
                                                var v59 uint64 = 104
                                                _ = v59
                                            }
                                            _ = v54
                                        } else {
                                            v60 := GenS6{F0: false, F1: 65535}
                                            for fi61 := 0; fi61 < 4; fi61++ {
                                                v62 := GenS7{F0: true, F1: true, F2: v7}
                                                if (true && (((fi61 + fi61) > fi61) || (-130 < fi61))) {
                                                    var v63 int16 = -1
                                                    if (4294967295 == v11) {
                                                        switch fi17 {
                                                        case 856:
                                                            fmt.Println(fi61, v11, fi17)
                                                        default:
                                                            if true {
                                                                delete(v1, "hello")
                                                                var v64 string = "a"
                                                                _ = v64
                                                            }
                                                        }
                                                    } else {
                                                        var v65 uint32 = v11
                                                        if len(v2) >= 2 {
                                                            v2 = v2[1:len(v2)]
                                                        }
                                                        _ = v65
                                                    }
                                                    v66 := GenS6{F0: true, F1: genFn1()}
                                                    _ = v63
                                                    _ = v66
                                                } else {
                                                    if len(v2) >= 2 {
                                                        v2 = v2[1:len(v2)]
                                                    }
                                                }
                                                _ = fi61
                                                _ = v62
                                            }
                                            v7 += v7
                                            _ = v60
                                        }
                                    }
                                    _ = v53
                                }
                            case "a":
                                var v67 uint32 = v11
                                _ = v67
                            default:
                                for fi68 := 0; fi68 < 11; fi68++ {
                                    v69 := GenS5{F0: (v7 * v7), F1: v11}
                                    _ = fi68
                                    _ = v69
                                }
                            }
                        } else {
                            var v70 uint64 = 398
                            _ = v70
                        }
                        v71 := GenS3{F0: v11}
                        if (fi17 < -1) {
                            var v72 int64 = 42
                            _ = v72
                        } else {
                            v11 = (v11 >> ((genFn0(fi17, 1)) & 63))
                        }
                        _ = v71
                    }
                }
                var v73 int16 = -52
                _ = fi17
                _ = v73
            }
        }
    } else {
        if len(v2) > 0 {
            v2[0] = 9223372036854775807
        }
    }
    v74 := GenS6{F0: false, F1: 0}
    fmt.Println(v11)
    switch 598 {
    case -9223372036854775808:
        fmt.Println(v11, v11)
    case 127:
        if false {
            var v75 bool = true
            _ = v75
        } else {
            var v76 float64 = ((0.1 / 1e10) - 1.5)
            v11 += 4294967295
            _ = v76
        }
    }
    v77 := GenS4{F0: 255, F1: 686}
    var v78 uint64 = 42
    var v79 float32 = ((1.5 + 1e10) + 3.14159)
    v80 := GenS8{F0: 365}
    for fi81 := 0; fi81 < 4; fi81++ {
        var v82 string = "\n\t"
        v83 := GenS2{F0: 0}
        _ = fi81
        _ = v82
        _ = v83
    }
    v7 = 0
    v7 = -1
    var v84 bool = false
    v85 := GenS9{F0: -271}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v7:", v7)
    fmt.Println("v78:", v78)
    fmt.Println("v79:", v79)
    fmt.Println("v84:", v84)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v7
    _ = v74
    _ = v77
    _ = v78
    _ = v79
    _ = v80
    _ = v84
    _ = v85
}
