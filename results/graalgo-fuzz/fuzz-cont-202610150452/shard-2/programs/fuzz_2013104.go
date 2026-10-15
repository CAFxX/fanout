package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int64; F1 uint8; F2 int }
type GenS3 struct { F0 uint64; F1 uint64; F2 string }
type GenS4 struct { F0 uint16; F1 int32; F2 int16 }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 int64; F1 int64; F2 uint32 }
type GenS7 struct { F0 int8 }

func main() {
    v1 := make(map[uint32]int)
    v2 := []uint32{4294967295}
    var v3 int32 = 2147483647
    var v4 int64 = -1
    v4 = v4
    var v5 float64 = float64((float32((1e10 * (3.14159 - 0.1))) + 1.5))
    v6 := GenS1{F0: 4294967295}
    if ((18446744073709551615 < 42) && (1e10 != 0.0)) {
        for fi7 := 0; fi7 < 5; fi7++ {
            var v8 int8 = 1
            switch "a" {
            case "a":
                v9 := GenS2{F0: (v4 ^ 1), F1: 128, F2: 0}
                _ = v9
            case "hello":
                v10 := GenS2{F0: v4, F1: 101, F2: (^(fi7))}
                _ = v10
            case "\n\t":
                v1[882] = fi7
            }
            var v11 int32 = 0
            _ = fi7
            _ = v11
            _ = v8
        }
    } else {
        for fi12 := 0; fi12 < 1; fi12++ {
            fmt.Println(v3, fi12, v3)
            v13 := GenS2{F0: (^(v4)), F1: 93, F2: (^(fi12))}
            _ = fi12
            _ = v13
        }
        for fi14 := 0; fi14 < 6; fi14++ {
            v4 -= v4
            for fi15 := 0; fi15 < 4; fi15++ {
                if false {
                    v5 -= float64((1.5 / (float32(v5) - 0.0)))
                    v4 += v4
                    switch fi15 {
                    case 1:
                        var v16 string = "a"
                        _ = v16
                    case 128:
                        switch "foo bar" {
                        case "foo bar":
                            var v17 string = ("\n\t" + "a")
                            _ = v17
                        case "\n\t":
                            var v18 int8 = -51
                            _ = v18
                        default:
                            var v19 float32 = 0.1
                            var v20 int16 = -32768
                            v21 := GenS3{F0: 18446744073709551615, F1: 0, F2: ""}
                            fmt.Println(fi15, v3, fi14)
                            _ = v19
                            _ = v20
                            _ = v21
                        }
                    case -9223372036854775808:
                        switch "\n\t" {
                        case "\n\t":
                            var v22 int16 = 1
                            if true {
                                if (2 > 1) {
                                    fmt.Println(v5, v3, fi14)
                                } else {
                                    v5 = v5
                                    v4 -= 371
                                    var v23 int16 = int16(fi14)
                                    v24 := GenS4{F0: 2, F1: v3, F2: 0}
                                    _ = v23
                                    _ = v24
                                }
                            }
                            _ = v22
                        default:
                            v25 := GenS3{F0: 760, F1: 988, F2: "hello"}
                            var v26 uint16 = 255
                            var v27 float64 = 0.1
                            _ = v25
                            _ = v26
                            _ = v27
                        }
                    default:
                        var v28 uint32 = 42
                        v29 := GenS1{F0: v28}
                        fmt.Println(fi15, v28, v5)
                        v5 = v5
                        v4 += v4
                        _ = v28
                        _ = v29
                    }
                } else {
                    var v30 int8 = 127
                    var v31 int64 = v4
                    _ = v30
                    _ = v31
                }
                _ = fi15
            }
            v32 := GenS3{F0: 4294967295, F1: 42, F2: "foo bar"}
            _ = fi14
            _ = v32
        }
        v33 := GenS4{F0: 32767, F1: (v3 & v3), F2: -1}
        _ = v33
    }
    for fi34 := 0; fi34 < 4; fi34++ {
        fmt.Println(v3, v3, fi34)
        _ = len(v1)
        _ = fi34
    }
    delete(v1, 679)
    if (9223372036854775807 >= -502) {
        for fi35 := 0; fi35 < 7; fi35++ {
            delete(v1, 302)
            _ = fi35
        }
    }
    if false {
        fmt.Println(v4, v4, v3)
        v3 = v3
    } else {
        var v36 bool = (42 != 2)
        if (!v36) {
            for fi37 := 0; fi37 < 6; fi37++ {
                if (!((v36 || v36) || v36)) {
                    if false {
                        v38 := GenS2{F0: (v4 % ((0) | 1)), F1: 0, F2: fi37}
                        var v39 int16 = -731
                        if v36 {
                            v40 := GenS5{F0: 2}
                            _ = v40
                        } else {
                            switch ("hello" + "") {
                            case "\n\t":
                                v3 = (-(v3))
                            case "a":
                                var v41 int8 = int8(v3)
                                _ = v41
                            }
                        }
                        _ = v38
                        _ = v39
                    }
                    var v42 int32 = v3
                    _ = v42
                } else {
                    for fi43 := 0; fi43 < 1; fi43++ {
                        v3 = (-(v3))
                        var v44 uint64 = 56
                        _ = fi43
                        _ = v44
                    }
                    if true {
                        if ((2 <= 2147483648) && v36) {
                            v5 += 0.0
                            var v45 float64 = 3.14159
                            _ = v45
                        }
                        if true {
                            fmt.Println(v4, v3, v5)
                            if (255 != 255) {
                                var v46 uint32 = 1
                                _ = v46
                            } else {
                                for fi47 := 0; fi47 < 9; fi47++ {
                                    switch (fi47 << ((0) & 63)) {
                                    case 990:
                                        v3 += v3
                                    }
                                    v36 = (1e10 >= 0.1)
                                    v48 := GenS6{F0: -1, F1: v4, F2: 447}
                                    _ = fi47
                                    _ = v48
                                }
                                v49 := v1[2147483648]
                                _ = v49
                            }
                            if (!(false || ("hello" <= ("foo bar" + "foo bar")))) {
                                for fi50 := 0; fi50 < 7; fi50++ {
                                    if len(v2) > 0 {
                                        v2[0] = 2
                                    }
                                    fmt.Println(v4, v36, v5)
                                    v51 := GenS1{F0: 2147483648}
                                    _ = fi50
                                    _ = v51
                                }
                                var v52 uint32 = 2147483647
                                _ = v52
                            } else {
                                var v53 uint8 = 52
                                v54 := GenS6{F0: v4, F1: v4, F2: uint32(v53)}
                                v55 := GenS1{F0: 42}
                                _ = v53
                                _ = v54
                                _ = v55
                            }
                        }
                        switch (fi37 - 778) {
                        case 42:
                            if v36 {
                                v56 := GenS4{F0: 1, F1: 923, F2: -853}
                                for fi57 := 0; fi57 < 13; fi57++ {
                                    v36 = v36
                                    v58 := v1[2]
                                    v3 += (-(v3))
                                    _ = fi57
                                    _ = v58
                                }
                                _ = v56
                            } else {
                                v4 -= (-(v4))
                                var v59 uint16 = 212
                                v36 = ((v4 ^ 42) > -440)
                                _ = v59
                            }
                        default:
                            if (v3 > (v3 << ((1) & 63))) {
                                if ("a" <= "hello") {
                                    v60 := GenS5{F0: 65535}
                                    for fi61 := 0; fi61 < 7; fi61++ {
                                        if v36 {
                                            var v62 int32 = 2147483647
                                            var v63 uint64 = 59
                                            _ = v62
                                            _ = v63
                                        } else {
                                            var v64 string = ((("\n\t" + "foo bar") + "\n\t") + "a")
                                            v65 := GenS2{F0: v4, F1: 204, F2: fi61}
                                            var v66 uint8 = 0
                                            _ = v64
                                            _ = v65
                                            _ = v66
                                        }
                                        var v67 uint32 = 2147483648
                                        fmt.Println(v4, v4, fi37)
                                        v5 = 1e10
                                        _ = fi61
                                        _ = v67
                                    }
                                    if v36 {
                                        fmt.Println(v3, fi37, v3)
                                        v2 = append(v2, 727)
                                        var v68 uint = 8
                                        _ = v68
                                    } else {
                                        if false {
                                            switch ((("foo bar" + "\n\t") + ("" + ("\n\t" + "\n\t"))) + ((("\n\t" + "foo bar") + "") + "a")) {
                                            case "hello":
                                                v3 -= v3
                                            case "":
                                                if false {
                                                    fmt.Println(v36, v4)
                                                    fmt.Println(v36, v5)
                                                }
                                            default:
                                                v2 = append(v2, 2147483647)
                                            }
                                            switch ("hello" + "") {
                                            case "foo bar":
                                                var v69 bool = (32767 == 1)
                                                _ = v69
                                            }
                                        }
                                    }
                                    _ = v60
                                }
                                fmt.Println(v36)
                                var v70 uint16 = 0
                                _ = v70
                            }
                        }
                    }
                }
                _ = fi37
            }
            var v71 uint64 = 672
            if v36 {
                var v72 int32 = (v3 | (v3 + 42))
                for fi73 := 0; fi73 < 13; fi73++ {
                    var v74 int64 = (v4 ^ (v4 % (((v4 * v4)) | 1)))
                    _ = fi73
                    _ = v74
                }
                _ = v72
            }
            _ = v71
        }
        _ = v36
    }
    for fi75 := 0; fi75 < 17; fi75++ {
        for fi76 := 0; fi76 < 8; fi76++ {
            v77 := GenS4{F0: 808, F1: -568, F2: 1000}
            v3 = (^(v3))
            _ = fi76
            _ = v77
        }
        v78 := GenS7{F0: 100}
        _ = fi75
        _ = v78
    }
    for fi79 := 0; fi79 < 9; fi79++ {
        v3 += (v3 ^ (v3 * (v3 / ((v3) | 1))))
        _ = fi79
    }
    switch "\n\t" {
    case "":
        switch 2147483647 {
        case 32768:
            var v80 uint = 804
            _ = v80
        case 4294967295:
            if (v3 > (v3 / ((v3) | 1))) {
                delete(v1, 2147483648)
            } else {
                var v81 int8 = 127
                var v82 uint64 = 1
                v81 += 0
                _ = v81
                _ = v82
            }
        case -2147483648:
            _ = len(v2)
        default:
            if ((v5 == v5) && true) {
                var v83 bool = ((v4 & v4) < -142)
                switch ("hello" + ("\n\t" + "\n\t")) {
                case "\n\t":
                    v84 := GenS1{F0: 61}
                    _ = v84
                case "a":
                    var v85 uint16 = 0
                    _ = v85
                case "foo bar":
                    var v86 float64 = (v5 - v5)
                    _ = v86
                default:
                    var v87 uint64 = 155
                    _ = v87
                }
                _ = v83
            } else {
                var v88 uint32 = 2147483648
                _ = v88
            }
        }
    default:
        v1[2] = -508
    }
    v4 += v4
    v5 = (v5 / ((1.5 / v5) * ((v5 * (((v5 + v5) - 1.5) / (float64((0.1 - 1e10)) - (v5 * v5)))) + ((v5 + v5) + v5))))
    v4 += int64(v3)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
