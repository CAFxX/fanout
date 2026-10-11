package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16; F2 uint }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int; F1 int }
type GenS4 struct { F0 uint; F1 int; F2 int8 }
type GenS5 struct { F0 uint8; F1 int32 }
type GenS6 struct { F0 int16; F1 bool; F2 uint8 }
type GenS7 struct { F0 uint64 }
type GenS8 struct { F0 uint64; F1 bool }
type GenS9 struct { F0 int; F1 uint32 }
type GenS10 struct { F0 int8; F1 int8; F2 uint64 }
type GenS11 struct { F0 int; F1 string; F2 uint8 }
type GenS12 struct { F0 int; F1 uint64; F2 uint }
type GenS13 struct { F0 int16 }

func main() {
    var v1 float64 = (0.1 / (1.5 - 0.0))
    var v2 int64 = 9223372036854775807
    var v3 bool = true
    var v4 uint16 = 1
    var v5 float32 = 0.0
    var v6 int16 = -814
    v7 := GenS1{F0: "\n\t", F1: v6, F2: 4294967295}
    switch "a" {
    case "hello":
        if (true && ("foo bar" < ("foo bar" + "foo bar"))) {
            var v8 int8 = -21
            if false {
                for fi9 := 0; fi9 < 4; fi9++ {
                    fmt.Println(fi9, v2)
                    var v10 int8 = v8
                    v11 := []int{-2147483648}
                    _ = fi9
                    _ = v10
                    _ = v11
                }
                v12 := []int{42, -748, -9223372036854775808, 886}
                _ = v12
            }
            v13 := []int{42}
            _ = v13
            _ = v8
        } else {
            if (1 > 4294967295) {
                for fi14 := 0; fi14 < 20; fi14++ {
                    var v15 int16 = v6
                    v16 := GenS2{F0: 2147483647}
                    var v17 uint16 = (v4 | v4)
                    _ = fi14
                    _ = v15
                    _ = v16
                    _ = v17
                }
                v2 -= v2
                if (v5 == float32(v1)) {
                    v6 += 392
                    v18 := []int{}
                    v2 = v2
                    _ = v18
                } else {
                    fmt.Println(v6)
                }
            }
            v6 += v6
            var v19 uint8 = 255
            _ = v19
        }
    case "foo bar":
        for fi20 := 0; fi20 < 4; fi20++ {
            v2 = v2
            fmt.Println(v2)
            _ = fi20
        }
    default:
        v21 := []int{1, 9223372036854775807, -9223372036854775808}
        _ = v21
    }
    v22 := []int{-2147483648, 502, -428}
    if false {
        var v23 uint64 = 0
        if (float32((v1 - 1.5)) < -2.25) {
            for fi24 := 0; fi24 < 12; fi24++ {
                if (!v3) {
                    fmt.Println(v1, v6, v4)
                } else {
                    var v25 uint64 = v23
                    if v3 {
                        var v26 uint16 = v4
                        if v3 {
                            if (!(v3 || v3)) {
                                fmt.Println(v4, v1, v6)
                            } else {
                                v3 = (false || (2 <= 42))
                            }
                        } else {
                            v5 += -2.25
                        }
                        _ = v26
                    } else {
                        switch ("a" + "a") {
                        case "foo bar":
                            v27 := GenS2{F0: 1}
                            _ = v27
                        case "\n\t":
                            switch 128 {
                            case 64:
                                fmt.Println(v6, v3)
                            case -685:
                                v2 -= 42
                            default:
                                var v28 float32 = v5
                                _ = v28
                            }
                        case "a":
                            v29 := GenS3{F0: (fi24 - int(v23)), F1: (-(fi24))}
                            _ = v29
                        }
                        _ = len(v22)
                        var v30 uint64 = 733
                        _ = v30
                    }
                    v6 -= -32768
                    _ = v25
                }
                v31 := GenS3{F0: (^(fi24)), F1: fi24}
                _ = fi24
                _ = v31
            }
            v32 := GenS1{F0: ("foo bar" + "hello"), F1: (v6 << ((18446744073709551615) & 63)), F2: 65535}
            v3 = v3
            _ = v32
        } else {
            switch -9223372036854775808 {
            case 128:
                fmt.Println(v3, v6)
            case 32768:
                fmt.Println(v3, v2, v23)
            case 42:
                var v33 uint32 = 272
                _ = v33
            default:
                var v34 uint32 = 547
                _ = v34
            }
            if len(v22) > 0 {
                v22[0] = -672
            }
        }
        v2 = v2
        _ = v23
    }
    for fi35 := 0; fi35 < 6; fi35++ {
        _ = len(v22)
        _ = fi35
    }
    fmt.Println(v2)
    if (v3 || v3) {
        v2 -= (^(v2))
        for fi36 := 0; fi36 < 12; fi36++ {
            var v37 int64 = 1
            v6 = (v6 << ((871) & 63))
            v1 = 0.1
            _ = fi36
            _ = v37
        }
    } else {
        fmt.Println(v5, v1)
    }
    v38 := GenS2{F0: -447}
    var v39 int32 = -1
    var v40 bool = (920 != 1)
    for fi41 := 0; fi41 < 5; fi41++ {
        if (0 > int16(v2)) {
            if (false && v40) {
                var v42 bool = (fi41 != fi41)
                _ = v42
            } else {
                fmt.Println(v2, v6)
                switch "a" {
                case "a":
                    fmt.Println(v4, v39)
                case "foo bar":
                    var v43 uint16 = (^(v4))
                    _ = v43
                default:
                    v4 = v4
                }
                if (true || (v5 > v5)) {
                    switch "hello" {
                    case "foo bar":
                        v2 += v2
                    }
                    v40 = ((127 != 1) || true)
                } else {
                    if v40 {
                        v44 := GenS1{F0: "\n\t", F1: (v6 >> ((490) & 63)), F2: 255}
                        _ = v44
                    }
                }
            }
            var v45 uint16 = v4
            if len(v22) >= 2 {
                v22 = v22[1:len(v22)]
            }
            _ = v45
        } else {
            if (false && (((-(v2)) > 9223372036854775807) && true)) {
                fmt.Println(fi41, v1, v2)
            } else {
                for fi46 := 0; fi46 < 3; fi46++ {
                    if (false && v40) {
                        switch ("a" + "foo bar") {
                        case "foo bar":
                            var v47 uint = 255
                            _ = v47
                        case "a":
                            if len(v22) >= 2 {
                                v22 = v22[1:len(v22)]
                            }
                        }
                        if v40 {
                            var v48 uint64 = 0
                            _ = v48
                        }
                    } else {
                        v1 = 0.1
                        v49 := GenS4{F0: 83, F1: fi41, F2: 42}
                        fmt.Println(v3, v40)
                        if (!((((v5 + v5) + v5) > v5) && v3)) {
                            if v3 {
                                v50 := GenS4{F0: 1, F1: (fi46 << ((679) & 63)), F2: 66}
                                v2 = v2
                                v51 := GenS3{F0: (fi41 - 32767), F1: (fi41 << ((18446744073709551615) & 63))}
                                _ = v50
                                _ = v51
                            }
                        }
                        switch "" {
                        case "a":
                            _ = len(v22)
                        case "hello":
                            switch (^(fi46)) {
                            case 128:
                                switch "foo bar" {
                                case "hello":
                                    switch "a" {
                                    case "hello":
                                        if len(v22) > 0 {
                                            v22[0] = -839
                                        }
                                    case "a":
                                        v4 += (v4 + v4)
                                    }
                                    switch fi46 {
                                    case 0:
                                        v22 = append(v22, (fi46 / (((fi41 & -2147483648)) | 1)))
                                    case 4294967295:
                                        v22 = append(v22, fi46)
                                    }
                                    if false {
                                        var v52 uint = 65535
                                        var v53 bool = false
                                        _ = v52
                                        _ = v53
                                    }
                                case "":
                                    if len(v22) >= 2 {
                                        v22 = v22[1:len(v22)]
                                    }
                                    var v54 int32 = (v39 - (v39 + 0))
                                    _ = v54
                                default:
                                    v4 += 242
                                }
                            case 2147483647:
                                switch fi41 {
                                case 0:
                                    if ((v6 | 15) != (v6 ^ v6)) {
                                        v40 = v3
                                        fmt.Println(v40, v4)
                                    }
                                }
                            case 1:
                                fmt.Println(v3)
                            }
                        case "\n\t":
                            v22 = append(v22, (fi41 + (fi41 | fi41)))
                        default:
                            switch fi46 {
                            case 32768:
                                v6 -= -32768
                            case -376:
                                if (((v39 % ((int32(v6)) | 1)) > -2147483648) && v3) {
                                    var v55 uint64 = 1
                                    var v56 uint8 = 255
                                    v57 := GenS5{F0: (v56 >> ((18446744073709551615) & 63)), F1: -1}
                                    _ = v55
                                    _ = v56
                                    _ = v57
                                }
                            default:
                                var v58 uint = 18446744073709551615
                                _ = v58
                            }
                        }
                        switch "a" {
                        case "hello":
                            if true {
                                v1 = v1
                                if (!v40) {
                                    if v40 {
                                        fmt.Println(v1)
                                        var v59 string = "\n\t"
                                        var v60 uint16 = (^(v4))
                                        _ = v59
                                        _ = v60
                                    }
                                    switch (fi46 & -9223372036854775808) {
                                    case -596:
                                        var v61 uint32 = 1
                                        v22 = append(v22, fi41)
                                        _ = v61
                                    case 0:
                                        fmt.Println(v2)
                                    case 2147483647:
                                        switch 2147483647 {
                                        case 127:
                                            v5 += v5
                                        }
                                    default:
                                        var v62 float32 = float32(((v1 - (v1 + v1)) * (v1 - 1e10)))
                                        _ = v62
                                    }
                                }
                            }
                        case "":
                            var v63 uint16 = (v4 ^ v4)
                            _ = v63
                        }
                        _ = len(v22)
                        if v40 {
                            v40 = v3
                            v64 := GenS6{F0: (v6 + -326), F1: (42 >= 392), F2: 2}
                            var v65 int16 = -73
                            var v66 int8 = 127
                            _ = v64
                            _ = v65
                            _ = v66
                        } else {
                            if (!false) {
                                switch fi46 {
                                case 32767:
                                    v4 = v4
                                }
                                v67 := GenS6{F0: v6, F1: (753 < v39), F2: 0}
                                _ = v67
                            } else {
                                fmt.Println(v2)
                            }
                            v4 = (v4 - v4)
                            _ = len(v22)
                        }
                        v4 = 32767
                        if v40 {
                            v68 := GenS1{F0: "a", F1: v6, F2: 0}
                            var v69 int8 = 100
                            _ = len(v22)
                            _ = v68
                            _ = v69
                        } else {
                            fmt.Println(v40)
                            fmt.Println(v2, v5, v1)
                        }
                        _ = v49
                    }
                    v4 -= v4
                    _ = fi46
                }
                fmt.Println(v39)
            }
        }
        for fi70 := 0; fi70 < 13; fi70++ {
            _ = len(v22)
            switch "" {
            case "\n\t":
                var v71 int8 = 100
                _ = v71
            case "":
                if len(v22) > 0 {
                    v22[0] = -9223372036854775808
                }
            default:
                v22 = append(v22, fi41)
            }
            if v40 {
                v6 = (v6 / ((1) | 1))
                v4 = v4
                switch fi41 {
                case 127:
                    v72 := GenS7{F0: 4294967295}
                    _ = v72
                default:
                    fmt.Println(v6, v6, v39)
                }
            } else {
                if v40 {
                    if len(v22) >= 2 {
                        v22 = v22[1:len(v22)]
                    }
                    v73 := GenS8{F0: 770, F1: (480 != 2)}
                    if false {
                        var v74 bool = v3
                        if (-1 < v39) {
                            var v75 int32 = (v39 & 1)
                            v76 := GenS9{F0: 0, F1: 4294967295}
                            var v77 int8 = -76
                            _ = v75
                            _ = v76
                            _ = v77
                        } else {
                            v78 := GenS9{F0: (fi70 * -463), F1: 2147483648}
                            switch (fi41 << ((2) & 63)) {
                            case 127:
                                if true {
                                    var v79 string = "hello"
                                    var v80 int16 = 1000
                                    if len(v22) >= 2 {
                                        v22 = v22[1:len(v22)]
                                    }
                                    _ = v79
                                    _ = v80
                                }
                            case -9223372036854775808:
                                var v81 int64 = int64(v4)
                                _ = v81
                            case 32768:
                                v40 = (v74 && (false && v74))
                            }
                            _ = v78
                        }
                        _ = len(v22)
                        if false {
                            v82 := GenS10{F0: 1, F1: 127, F2: 0}
                            if (v40 && (!(!v40))) {
                                if false {
                                    v6 = (v6 + (v6 >> ((255) & 63)))
                                }
                            } else {
                                v1 = (v1 * 0.0)
                            }
                            v83 := GenS4{F0: 2, F1: (fi70 * (fi70 * 4294967295)), F2: int8(v6)}
                            _ = v82
                            _ = v83
                        }
                        _ = v74
                    }
                    if len(v22) > 0 {
                        v22[0] = 9223372036854775807
                    }
                    if len(v22) > 0 {
                        v22[0] = -137
                    }
                    switch "foo bar" {
                    case "foo bar":
                        switch "hello" {
                        case "a":
                            v1 -= float64((((0.0 - v5) * (((((v5 * 0.0) * v5) + 1e10) + float32(v1)) - ((-2.25 * (v5 - v5)) * v5))) - (v5 + -2.25)))
                        case "":
                            fmt.Println(v3, v3, v2)
                        }
                    default:
                        if len(v22) > 0 {
                            v22[0] = 32768
                        }
                    }
                    v84 := GenS1{F0: "hello", F1: v6, F2: uint(v4)}
                    _ = v73
                    _ = v84
                } else {
                    var v85 int32 = (v39 - v39)
                    _ = v85
                }
                if v40 {
                    v86 := GenS11{F0: fi41, F1: ((("\n\t" + ("\n\t" + ("foo bar" + "foo bar"))) + "hello") + "hello"), F2: 23}
                    var v87 uint16 = v4
                    v1 = v1
                    _ = v86
                    _ = v87
                } else {
                    v1 = 0.1
                    switch (("" + "") + ("foo bar" + "foo bar")) {
                    case "foo bar":
                        switch fi70 {
                        case 127:
                            if ((true || false) && (622 >= 0)) {
                                if len(v22) > 0 {
                                    v22[0] = fi70
                                }
                                if len(v22) >= 2 {
                                    v22 = v22[1:len(v22)]
                                }
                            }
                        case 32767:
                            v88 := GenS1{F0: "hello", F1: (v6 << ((2) & 63)), F2: 18446744073709551615}
                            _ = v88
                        case 2147483647:
                            var v89 int16 = v6
                            _ = v89
                        default:
                            fmt.Println(v4, v40)
                            v90 := GenS12{F0: (-(fi41)), F1: 1, F2: 4294967295}
                            v91 := GenS13{F0: v6}
                            _ = v90
                            _ = v91
                        }
                    case "\n\t":
                        v40 = true
                    case "":
                        switch "" {
                        case "hello":
                            v22 = append(v22, fi70)
                        case "":
                            var v92 uint16 = 0
                            _ = v92
                        }
                    }
                    v93 := GenS5{F0: 120, F1: v39}
                    _ = v93
                }
            }
            _ = fi70
        }
        var v94 string = ((("a" + "a") + ("foo bar" + "\n\t")) + "hello")
        _ = fi41
        _ = v94
    }
    var v95 uint8 = 1
    v96 := GenS3{F0: 177, F1: 4294967295}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v95:", v95)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v5
    _ = v6
    _ = v7
    _ = v95
    _ = v96
}
