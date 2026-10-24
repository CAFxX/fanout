package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string; F2 int8 }
type GenS2 struct { F0 uint16; F1 int64; F2 uint }
type GenS3 struct { F0 uint16; F1 int32; F2 uint16 }
type GenS4 struct { F0 int16; F1 uint64 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 string; F1 int16; F2 uint64 }
func genFn0(p0 uint16) int {
    var v7 string = "foo bar"
    var v8 int = 0
    v9 := GenS1{F0: v8, F1: ("hello" + "\n\t"), F2: 42}
    _ = p0
    _ = v7
    _ = v8
    _ = v9
    return (^(v8))
}
func genFn1(p0 bool, p1 int32) bool {
    v10 := GenS1{F0: -994, F1: (("hello" + "a") + ""), F2: -84}
    v11 := GenS1{F0: genFn0(65535), F1: "foo bar", F2: -107}
    if (!true) {
        for fi12 := 0; fi12 < 9; fi12++ {
            for fi13 := 0; fi13 < 7; fi13++ {
                if (!(("" + (("\n\t" + "\n\t") + "")) <= "\n\t")) {
                    switch -844 {
                    case 434:
                        v14 := GenS2{F0: 32767, F1: 42, F2: 65535}
                        _ = v14
                    }
                } else {
                    var v15 int16 = 32767
                    if (float32(float64(1.5)) <= 3.14159) {
                        v16 := GenS2{F0: 42, F1: 1, F2: 4294967295}
                        _ = v16
                    }
                    _ = v15
                }
                _ = fi13
            }
            var v17 uint64 = 2
            for fi18 := 0; fi18 < 20; fi18++ {
                fmt.Println(v17, fi12)
                var v19 int64 = 42
                _ = fi18
                _ = v19
            }
            _ = fi12
            _ = v17
        }
    } else {
        switch 9223372036854775807 {
        case 771:
            v20 := GenS1{F0: -1, F1: "", F2: 100}
            _ = v20
        case 42:
            v21 := []int{9223372036854775807, 42, 538, 0}
            _ = v21
        case -31:
            fmt.Println(p0, p1)
        default:
            p0 = ((3.14159 != (((0.1 * ((-2.25 - (1e10 * 3.14159)) - -2.25)) + 1.5) + (1e10 / ((0.1 / 1.5) / 1.0)))) && p0)
        }
    }
    var v22 int16 = -32768
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    _ = v22
    return ((242 <= 282) && p0)
}

func main() {
    var v1 int32 = -2147483648
    v2 := []int{}
    var v3 int8 = -128
    var v4 uint32 = 4294967295
    var v5 uint = 42
    v6 := make(map[int]int)
    v23 := GenS2{F0: 516, F1: 917, F2: (^(v5))}
    for fi24 := 0; fi24 < 17; fi24++ {
        v25 := GenS2{F0: 1, F1: 257, F2: (v5 / ((v5) | 1))}
        v26 := GenS1{F0: int(v4), F1: "a", F2: v3}
        _ = fi24
        _ = v25
        _ = v26
    }
    var v27 int = -760
    var v28 uint16 = 61
    for fi29 := 0; fi29 < 8; fi29++ {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        for fi30 := 0; fi30 < 18; fi30++ {
            var v31 int16 = -202
            var v32 bool = false
            var v33 int16 = v31
            v32 = (v1 < v1)
            if false {
                switch (("" + "") + ("hello" + "\n\t")) {
                case "\n\t":
                    switch "hello" {
                    case "\n\t":
                        if (3.14159 >= (0.1 + 3.14159)) {
                            if false {
                                fmt.Println(fi29, fi30, fi29)
                                fmt.Println(fi29, v27, v3)
                                if ((0 <= 42) || v32) {
                                    v6[(fi29 << ((v5) & 63))] = int(v31)
                                    fmt.Println(v3, fi30)
                                    if len(v2) > 0 {
                                        v2[0] = genFn0((v28 % ((v28) | 1)))
                                    }
                                } else {
                                    switch "" {
                                    case "\n\t":
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                    case "hello":
                                        var v34 uint = (v5 >> ((uint(v28)) & 63))
                                        _ = v34
                                    case "foo bar":
                                        if false {
                                            var v35 uint8 = 42
                                            switch (fi29 + v27) {
                                            case 127:
                                                v4 += v4
                                            }
                                            v36 := v6[-199]
                                            switch (-(v27)) {
                                            case -9223372036854775808:
                                                switch ("a" + (("\n\t" + (("\n\t" + "") + "hello")) + "foo bar")) {
                                                case "":
                                                    switch (("hello" + "hello") + "\n\t") {
                                                    case "\n\t":
                                                        v2 = append(v2, v27)
                                                    case "foo bar":
                                                        v5 += v5
                                                    }
                                                case "hello":
                                                    var v37 bool = ((-(v33)) != (-(v31)))
                                                    _ = v37
                                                case "a":
                                                    v38 := GenS3{F0: v28, F1: v1, F2: (v28 ^ 889)}
                                                    if (680 > 18446744073709551615) {
                                                        switch (fi29 & (v36 << ((v5) & 63))) {
                                                        case -137:
                                                            v3 = v3
                                                        }
                                                    }
                                                    v39 := v6[(v27 & fi29)]
                                                    _ = v38
                                                    _ = v39
                                                }
                                            }
                                            _ = v35
                                            _ = v36
                                        } else {
                                            var v40 int32 = (v1 % ((-2147483648) | 1))
                                            _ = v40
                                        }
                                    }
                                    v5 = (v5 << (((v5 % (((v5 + v5)) | 1))) & 63))
                                    v41 := GenS1{F0: (fi29 - (-(v27))), F1: "", F2: 100}
                                    _ = v41
                                }
                                _ = len(v6)
                                delete(v6, (fi30 ^ (fi29 / (((fi29 & (v27 | v27))) | 1))))
                            }
                        }
                    case "a":
                        var v42 int64 = -9223372036854775808
                        _ = v42
                    default:
                        switch "" {
                        case "foo bar":
                            _ = len(v2)
                        }
                    }
                }
                if (v32 || false) {
                    var v43 int16 = v31
                    _ = v43
                }
            } else {
                fmt.Println(v4)
                if (42 >= int64(v4)) {
                    v44 := v6[genFn0((v28 % ((653) | 1)))]
                    _ = len(v2)
                    var v45 int16 = v31
                    _ = v44
                    _ = v45
                }
                v6[923] = (-(v27))
            }
            _ = fi30
            _ = v31
            _ = v32
            _ = v33
        }
        v46 := GenS4{F0: 32767, F1: 452}
        _ = fi29
        _ = v46
    }
    switch ("" + "foo bar") {
    case "a":
        v47 := GenS5{F0: true}
        _ = v47
    default:
        v2 = append(v2, (-(v27)))
    }
    _ = len(v6)
    for fi48 := 0; fi48 < 8; fi48++ {
        var v49 uint16 = (v28 - uint16(v5))
        if false {
            var v50 uint16 = (v28 - v49)
            var v51 int = 656
            if (628 > 32767) {
                for fi52 := 0; fi52 < 7; fi52++ {
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    _ = fi52
                }
            } else {
                var v53 int16 = -534
                fmt.Println(v49)
                if false {
                    var v54 string = "hello"
                    var v55 float64 = 0.0
                    v27 += genFn0(v49)
                    _ = v54
                    _ = v55
                } else {
                    v50 -= v49
                    if true {
                        fmt.Println(v51)
                        v56 := GenS4{F0: (v53 + (v53 * (v53 << ((v5) & 63)))), F1: 809}
                        var v57 int32 = -983
                        _ = v56
                        _ = v57
                    } else {
                        switch (("hello" + "foo bar") + "") {
                        case "hello":
                            var v58 int16 = v53
                            _ = v58
                        case "a":
                            if false {
                                for fi59 := 0; fi59 < 6; fi59++ {
                                    v53 += (v53 ^ (v53 - v53))
                                    _ = fi59
                                }
                                if (v5 <= 18446744073709551615) {
                                    var v60 bool = (!(0.1 < 0.1))
                                    _ = v60
                                } else {
                                    v61 := GenS3{F0: (v50 << ((v5) & 63)), F1: (v1 / ((-2147483648) | 1)), F2: (v49 % ((v50) | 1))}
                                    switch (fi48 | 4294967295) {
                                    case 128:
                                        for fi62 := 0; fi62 < 18; fi62++ {
                                            v63 := GenS3{F0: (v50 << ((v5) & 63)), F1: 42, F2: (^(v28))}
                                            fmt.Println(v28)
                                            _ = fi62
                                            _ = v63
                                        }
                                    }
                                    _ = v61
                                }
                            } else {
                                v5 += 1
                                v64 := GenS3{F0: v50, F1: (^(v1)), F2: (v49 % ((v50) | 1))}
                                _ = v64
                            }
                        case "foo bar":
                            v65 := v6[-1]
                            _ = v65
                        default:
                            fmt.Println(v27, v49)
                        }
                    }
                }
                _ = v53
            }
            _ = v50
            _ = v51
        }
        if (!(!(42 == -657))) {
            for fi66 := 0; fi66 < 7; fi66++ {
                var v67 string = ("foo bar" + "foo bar")
                _ = fi66
                _ = v67
            }
            v1 = v1
            if (-998 <= -32768) {
                v68 := GenS4{F0: -32768, F1: 42}
                _ = v68
            } else {
                if (false || false) {
                    var v69 int32 = (v1 << (((v5 / ((v5) | 1))) & 63))
                    v6[fi48] = (v27 | 32768)
                    for fi70 := 0; fi70 < 13; fi70++ {
                        v3 -= v3
                        _ = fi70
                    }
                    _ = v69
                }
                var v71 uint64 = uint64(v5)
                var v72 int8 = 113
                _ = v71
                _ = v72
            }
        } else {
            v5 += 0
            var v73 uint64 = 550
            _ = v73
        }
        _ = fi48
        _ = v49
    }
    if (false || true) {
        if false {
            if (127 == 2) {
                var v74 uint64 = 42
                var v75 int64 = 136
                _ = v74
                _ = v75
            }
        }
        for fi76 := 0; fi76 < 16; fi76++ {
            if true {
                for fi77 := 0; fi77 < 16; fi77++ {
                    v4 += v4
                    _ = fi77
                }
            } else {
                for fi78 := 0; fi78 < 14; fi78++ {
                    v28 = v28
                    switch 374 {
                    case -499:
                        v5 = (v5 << ((42) & 63))
                    default:
                        if true {
                            var v79 bool = ((1 < 32767) || (1e10 > float64(((3.14159 * 1e10) * 1.5))))
                            _ = v79
                        } else {
                            var v80 uint32 = (^(v4))
                            fmt.Println(v28, v1, v28)
                            _ = v80
                        }
                    }
                    if false {
                        delete(v6, (v27 % (((fi78 & 2147483647)) | 1)))
                    }
                    if false {
                        v81 := GenS6{F0: "", F1: -1, F2: 1}
                        v82 := GenS3{F0: 65535, F1: v1, F2: 65535}
                        fmt.Println(v1, v1)
                        _ = v81
                        _ = v82
                    }
                    _ = fi78
                }
                var v83 float64 = 3.14159
                _ = v83
            }
            _ = fi76
        }
    }
    var v84 int8 = -1
    v1 = 2147483647
    for fi85 := 0; fi85 < 1; fi85++ {
        fmt.Println(v27)
        _ = fi85
    }
    if false {
        var v86 int = int(v28)
        var v87 float64 = 1e10
        switch int(v1) {
        case 42:
            _ = len(v2)
        default:
            fmt.Println(v5, v86, v87)
        }
        _ = v86
        _ = v87
    } else {
        switch (^(v27)) {
        case -2147483648:
            fmt.Println(v1, v27, v28)
        case 2147483647:
            v5 -= v5
        case -835:
            if (v1 <= v1) {
                var v88 uint8 = 14
                _ = v88
            } else {
                var v89 string = "foo bar"
                v5 = 4294967295
                for fi90 := 0; fi90 < 10; fi90++ {
                    var v91 int16 = 1
                    _ = fi90
                    _ = v91
                }
                _ = v89
            }
        }
    }
    for fi92 := 0; fi92 < 13; fi92++ {
        v2 = append(v2, (v27 << (((^(v5))) & 63)))
        _ = fi92
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v84:", v84)
    _ = v1
    _ = v2
    _ = v23
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v84
}
