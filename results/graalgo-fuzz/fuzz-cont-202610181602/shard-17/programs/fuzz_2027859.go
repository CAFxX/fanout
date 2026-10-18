package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 bool; F1 int }
type GenS4 struct { F0 int32; F1 uint64; F2 int32 }
type GenS5 struct { F0 uint64; F1 int }
type GenS6 struct { F0 int; F1 uint32 }
type GenS7 struct { F0 bool }
type GenS8 struct { F0 int }
type GenS9 struct { F0 uint16; F1 int32 }
type GenS10 struct { F0 bool; F1 uint32; F2 int }
type GenS11 struct { F0 uint32; F1 string; F2 int16 }
type GenS12 struct { F0 int }
type GenS13 struct { F0 int16; F1 int16 }
type GenS14 struct { F0 bool; F1 bool; F2 uint16 }
type GenS15 struct { F0 int16; F1 uint32 }
func genFn0(p0 int32, p1 int16, p2 uint16) int64 {
    for fi5 := 0; fi5 < 14; fi5++ {
        var v6 int = (fi5 % ((42) | 1))
        _ = fi5
        _ = v6
    }
    var v7 string = ("a" + "")
    p0 = (p0 & (p0 & (p0 + (p0 / ((p0) | 1)))))
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 0
}
func genFn1(p0 bool) uint64 {
    p0 = p0
    for fi8 := 0; fi8 < 16; fi8++ {
        p0 = ("hello" != "\n\t")
        _ = fi8
    }
    for fi9 := 0; fi9 < 18; fi9++ {
        p0 = (true || p0)
        p0 = p0
        _ = fi9
    }
    for fi10 := 0; fi10 < 14; fi10++ {
        for fi11 := 0; fi11 < 10; fi11++ {
            var v12 int = (-(fi11))
            _ = fi11
            _ = v12
        }
        switch "a" {
        case "hello":
            v13 := []int{fi10, (fi10 - fi10)}
            _ = v13
        case "\n\t":
            fmt.Println(fi10)
        default:
            v14 := GenS1{F0: 1}
            _ = v14
        }
        _ = fi10
    }
    if (557 > 9223372036854775807) {
        fmt.Println(p0)
        p0 = false
        v15 := GenS1{F0: 524}
        _ = v15
    }
    _ = p0
    return 566
}

func main() {
    v1 := []uint32{171, 4294967295, 2, 1}
    v2 := []int{-308}
    var v3 uint32 = 2147483648
    v4 := make(map[string]string)
    if (-2.25 > 3.14159) {
        v16 := GenS2{F0: 247}
        if true {
            var v17 float32 = ((3.14159 * -2.25) + 1.5)
            _ = v17
        }
        _ = v16
    } else {
        v1 = append(v1, (v3 ^ (v3 - (v3 >> ((65535) & 63)))))
        v18 := GenS3{F0: (!(false || false)), F1: 344}
        fmt.Println(v3)
        _ = v18
    }
    if false {
        switch -2147483648 {
        case 128:
            v19 := GenS1{F0: (^(v3))}
            _ = v19
        case 32768:
            v4["\n\t"] = ""
        case 32767:
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        default:
            if len(v1) > 0 {
                v1[0] = (v3 << ((255) & 63))
            }
        }
        if false {
            v3 = (v3 / (((v3 << ((4294967295) & 63))) | 1))
            v3 -= 2147483647
        } else {
            switch "a" {
            case "foo bar":
                var v20 float64 = 0.0
                _ = v20
            case "":
                if false {
                    delete(v4, "foo bar")
                    var v21 uint32 = 2147483648
                    _ = v21
                } else {
                    switch "hello" {
                    case "hello":
                        if false {
                            v3 += 611
                            switch -9223372036854775808 {
                            case -1:
                                var v22 int16 = 32767
                                _ = v22
                            case -461:
                                if (!(128 == 128)) {
                                    var v23 int16 = 384
                                    _ = v23
                                }
                            case 42:
                                v3 = 0
                            }
                        }
                    case "\n\t":
                        v3 = (v3 | 65535)
                    case "":
                        for fi24 := 0; fi24 < 1; fi24++ {
                            if true {
                                switch 32768 {
                                case 0:
                                    var v25 int8 = -1
                                    _ = v25
                                default:
                                    v3 += (v3 + (v3 * v3))
                                }
                                var v26 int64 = 683
                                _ = v26
                            }
                            v3 -= (v3 - (v3 + (v3 ^ v3)))
                            _ = fi24
                        }
                    }
                    for fi27 := 0; fi27 < 4; fi27++ {
                        if false {
                            for fi28 := 0; fi28 < 8; fi28++ {
                                fmt.Println(v3, fi27)
                                _ = fi28
                            }
                            var v29 uint = 4294967295
                            _ = v29
                        } else {
                            v3 += 2147483648
                            v3 += 1
                            var v30 uint32 = v3
                            _ = v30
                        }
                        var v31 float64 = 3.14159
                        _ = fi27
                        _ = v31
                    }
                }
            default:
                var v32 int = 32768
                _ = v32
            }
            switch (("hello" + "hello") + "a") {
            case "hello":
                fmt.Println(v3)
            case "":
                v3 += v3
            default:
                v3 = (v3 ^ (v3 ^ v3))
            }
            v3 = 4294967295
        }
        v33 := GenS4{F0: 42, F1: 1, F2: 209}
        _ = v33
    } else {
        var v34 int = 127
        v35 := GenS5{F0: genFn1(true), F1: (v34 ^ 128)}
        _ = v34
        _ = v35
    }
    var v36 uint = 595
    v37 := GenS1{F0: (v3 ^ (v3 / ((v3) | 1)))}
    switch 641 {
    case 103:
        for fi38 := 0; fi38 < 20; fi38++ {
            switch "" {
            case "":
                var v39 int64 = 714
                _ = v39
            case "hello":
                var v40 int = -9223372036854775808
                _ = v40
            case "foo bar":
                v36 -= 40
            }
            _ = fi38
        }
    }
    v41 := GenS6{F0: 4294967295, F1: (^(v3))}
    var v42 uint8 = 127
    if false {
        v43 := GenS7{F0: true}
        if (4294967295 != 2) {
            for fi44 := 0; fi44 < 6; fi44++ {
                v45 := GenS8{F0: (fi44 % ((4294967295) | 1))}
                _ = len(v4)
                _ = fi44
                _ = v45
            }
        }
        if ((0.0 + -2.25) == ((1.5 - 0.1) / float64(((1e10 * 1.5) + 0.1)))) {
            v42 += v42
            v4["\n\t"] = "a"
            v46 := GenS3{F0: (!true), F1: -554}
            _ = v46
        } else {
            for fi47 := 0; fi47 < 16; fi47++ {
                v3 = (v3 % ((v3) | 1))
                for fi48 := 0; fi48 < 12; fi48++ {
                    v3 -= v3
                    if false {
                        _ = len(v2)
                    } else {
                        v49 := GenS7{F0: true}
                        v36 = v36
                        if (!true) {
                            switch ("hello" + "a") {
                            case "\n\t":
                                v50 := GenS7{F0: true}
                                _ = v50
                            case "":
                                var v51 int16 = -747
                                _ = v51
                            case "foo bar":
                                var v52 uint8 = v42
                                _ = v52
                            }
                            var v53 int = 0
                            _ = v53
                        } else {
                            if (!false) {
                                if true {
                                    v54 := GenS9{F0: 255, F1: 2147483647}
                                    v4["hello"] = ""
                                    _ = v54
                                } else {
                                    var v55 uint64 = 0
                                    switch "" {
                                    case "a":
                                        fmt.Println(fi47)
                                    case "\n\t":
                                        switch "\n\t" {
                                        case "a":
                                            v3 = 2147483647
                                        default:
                                            if (9223372036854775807 <= -9223372036854775808) {
                                                var v56 uint8 = (^(v42))
                                                _ = v56
                                            } else {
                                                if true {
                                                    var v57 uint32 = v3
                                                    v36 = v36
                                                    v42 = (v42 % (((^(v42))) | 1))
                                                    _ = v57
                                                }
                                            }
                                        }
                                    default:
                                        fmt.Println(v3)
                                    }
                                    v58 := GenS3{F0: (fi48 == (fi47 & fi47)), F1: (fi48 & (-(fi47)))}
                                    _ = v55
                                    _ = v58
                                }
                                v3 = uint32(v42)
                            } else {
                                var v59 int64 = -462
                                _ = v59
                            }
                            v60 := GenS3{F0: (true || false), F1: fi48}
                            if (874 == 841) {
                                _ = len(v4)
                            } else {
                                v36 += v36
                                v36 += v36
                                if (float32(1e10) == 0.1) {
                                    v61 := GenS4{F0: 1, F1: 2, F2: -717}
                                    v62 := GenS10{F0: (!true), F1: 140, F2: fi47}
                                    _ = v61
                                    _ = v62
                                }
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                                v36 += (v36 & (v36 * v36))
                            }
                            _ = v60
                        }
                        _ = v49
                    }
                    _ = fi48
                }
                _ = fi47
            }
            if (!((^(v3)) < 884)) {
                for fi63 := 0; fi63 < 2; fi63++ {
                    switch (fi63 * (-(fi63))) {
                    case 9223372036854775807:
                        switch (("a" + "\n\t") + "\n\t") {
                        case "a":
                            switch "foo bar" {
                            case "foo bar":
                                switch "a" {
                                case "hello":
                                    fmt.Println(v42, v36)
                                case "a":
                                    v36 = v36
                                case "foo bar":
                                    v42 = 42
                                }
                            case "hello":
                                if ("\n\t" >= ("hello" + ("a" + "\n\t"))) {
                                    if true {
                                        if false {
                                            v3 = (v3 & (v3 % (((v3 - (^(v3)))) | 1)))
                                            v64 := GenS11{F0: (v3 >> ((uint(v3)) & 63)), F1: "hello", F2: 208}
                                            _ = v64
                                        } else {
                                            fmt.Println(v36, fi63, v3)
                                            for fi65 := 0; fi65 < 3; fi65++ {
                                                v66 := GenS12{F0: (^(fi65))}
                                                if (false && true) {
                                                    v3 += 2147483647
                                                    var v67 uint64 = genFn1(((564 != 18446744073709551615) || ((!(!false)) && false)))
                                                    _ = v67
                                                }
                                                _ = fi65
                                                _ = v66
                                            }
                                            v36 = v36
                                        }
                                        var v68 float32 = (3.14159 / (1e10 / 1e10))
                                        _ = v68
                                    }
                                    var v69 int8 = 0
                                    _ = v69
                                }
                            case "\n\t":
                                switch fi63 {
                                case 9223372036854775807:
                                    v36 = (v36 ^ (^(v36)))
                                case 795:
                                    if (true && true) {
                                        fmt.Println(v36, v36, fi63)
                                        v70 := GenS13{F0: -1, F1: 1000}
                                        switch (^(fi63)) {
                                        case -10:
                                            for fi71 := 0; fi71 < 9; fi71++ {
                                                v72 := GenS14{F0: false, F1: false, F2: 255}
                                                _ = fi71
                                                _ = v72
                                            }
                                        case 1:
                                            var v73 float32 = -2.25
                                            _ = v73
                                        }
                                        _ = v70
                                    } else {
                                        v42 -= v42
                                        var v74 string = ((("hello" + "\n\t") + "\n\t") + "foo bar")
                                        _ = v74
                                    }
                                case 127:
                                    if false {
                                        fmt.Println(v3, v3)
                                        var v75 int16 = -262
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                        _ = v75
                                    } else {
                                        for fi76 := 0; fi76 < 15; fi76++ {
                                            var v77 uint32 = 566
                                            _ = fi76
                                            _ = v77
                                        }
                                        var v78 float64 = float64(3.14159)
                                        _ = v78
                                    }
                                }
                            default:
                                var v79 float64 = (float64((-2.25 + float32(0.0))) - (0.0 + 1.5))
                                _ = v79
                            }
                        case "\n\t":
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                        }
                    default:
                        v80 := GenS3{F0: (((v3 >> ((65535) & 63)) >= 2147483648) || (-974 >= -329)), F1: fi63}
                        _ = v80
                    }
                    _ = fi63
                }
                v36 -= v36
                var v81 int = 9223372036854775807
                _ = v81
            } else {
                if false {
                    switch 1 {
                    case -565:
                        switch -9223372036854775808 {
                        case 0:
                            var v82 int32 = 42
                            _ = v82
                        case 32768:
                            var v83 int64 = 666
                            _ = v83
                        }
                    }
                    var v84 bool = (0 <= 0)
                    switch "" {
                    case "a":
                        fmt.Println(v3, v42)
                    case "foo bar":
                        var v85 uint8 = uint8(v36)
                        _ = v85
                    case "hello":
                        v84 = (((!v84) && v84) || (0 == 42))
                    }
                    _ = v84
                } else {
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                }
            }
        }
        _ = v43
    }
    fmt.Println(v42, v3)
    v3 = (v3 << ((243) & 63))
    v86 := GenS11{F0: (v3 | (^(v3))), F1: "", F2: 0}
    var v87 int8 = 81
    v42 = v42
    v88 := GenS15{F0: 0, F1: v3}
    for fi89 := 0; fi89 < 20; fi89++ {
        var v90 int32 = -1
        _ = fi89
        _ = v90
    }
    v42 += 2
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v42:", v42)
    fmt.Println("v87:", v87)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v4
    _ = v41
    _ = v42
    _ = v86
    _ = v87
    _ = v88
}
