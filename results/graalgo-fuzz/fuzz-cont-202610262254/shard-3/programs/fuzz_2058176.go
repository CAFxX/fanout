package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint; F2 int64 }
type GenS2 struct { F0 int64; F1 uint32 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 uint32; F1 uint8; F2 int8 }
type GenS6 struct { F0 string }
type GenS7 struct { F0 string; F1 uint64 }
type GenS8 struct { F0 int32; F1 bool; F2 string }
type GenS9 struct { F0 uint32; F1 uint32; F2 uint16 }
type GenS10 struct { F0 string; F1 uint8 }
type GenS11 struct { F0 int; F1 uint8; F2 int32 }
func genFn0(p0 int, p1 int64, p2 uint32) uint {
    if (int16(p0) > 42) {
        fmt.Println(p1)
        p1 += p1
    }
    v6 := []int{(p0 << ((582) & 63))}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 65535
}

func main() {
    v1 := []int{32768}
    v2 := []int{175}
    var v3 uint = 42
    var v4 int8 = 1
    var v5 bool = true
    for fi7 := 0; fi7 < 17; fi7++ {
        v8 := GenS1{F0: 2147483647, F1: v3, F2: 0}
        _ = fi7
        _ = v8
    }
    var v9 int64 = -1
    v10 := GenS1{F0: 1, F1: v3, F2: (v9 - v9)}
    if (v3 < 4294967295) {
        var v11 uint = (v3 + 373)
        var v12 uint8 = 127
        switch (("hello" + "") + ("a" + "")) {
        case "":
            v13 := GenS1{F0: uint32(v11), F1: v3, F2: (v9 * v9)}
            _ = v13
        case "\n\t":
            for fi14 := 0; fi14 < 18; fi14++ {
                v1 = append(v1, (fi14 % (((fi14 * (fi14 | 799))) | 1)))
                _ = fi14
            }
        case "hello":
            if (v5 || v5) {
                v9 = (v9 % (((v9 / ((1) | 1))) | 1))
            }
        }
        _ = v11
        _ = v12
    }
    fmt.Println(v3, v3)
    _ = len(v1)
    if (42 > 210) {
        fmt.Println(v9)
    } else {
        switch -9223372036854775808 {
        case 0:
            switch 128 {
            case 442:
                v9 = (^(v9))
            }
        default:
            for fi15 := 0; fi15 < 14; fi15++ {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                _ = fi15
            }
        }
        v9 += v9
        if (-1 < 32768) {
            if (2 >= 2) {
                fmt.Println(v5)
                if (-953 < -2147483648) {
                    v16 := GenS2{F0: (v9 | v9), F1: 4294967295}
                    _ = v16
                }
                v4 = v4
            }
        } else {
            if (47 <= 233) {
                if false {
                    v4 -= (v4 + -87)
                    _ = len(v2)
                }
                var v17 float64 = (((1e10 + (0.0 * (3.14159 / ((0.0 - 0.1) + -2.25)))) / -2.25) / 0.1)
                var v18 float64 = 3.14159
                _ = v17
                _ = v18
            }
            switch 9223372036854775807 {
            case 1:
                v19 := GenS2{F0: (v9 % ((v9) | 1)), F1: 1}
                _ = v19
            }
            switch 128 {
            case 37:
                v9 = -835
            case -633:
                if v5 {
                    for fi20 := 0; fi20 < 17; fi20++ {
                        v21 := GenS1{F0: 398, F1: v3, F2: v9}
                        switch (^(fi20)) {
                        case 32768:
                            fmt.Println(v9, v4)
                        case 128:
                            var v22 int16 = 1
                            _ = v22
                        }
                        v2 = append(v2, fi20)
                        _ = fi20
                        _ = v21
                    }
                    var v23 int32 = 2147483647
                    for fi24 := 0; fi24 < 1; fi24++ {
                        if (853 == 65535) {
                            var v25 uint64 = 994
                            var v26 float64 = 1.5
                            _ = v25
                            _ = v26
                        } else {
                            var v27 float32 = 0.0
                            if v5 {
                                switch "hello" {
                                case "\n\t":
                                    if (1 >= 714) {
                                        v28 := GenS1{F0: 70, F1: (^(v3)), F2: (v9 | -581)}
                                        if false {
                                            var v29 int = 994
                                            if (v9 > v9) {
                                                v2 = append(v2, (v29 - (^(v29))))
                                                v29 = fi24
                                                v30 := GenS2{F0: v9, F1: 2}
                                                _ = v30
                                            } else {
                                                var v31 int = 9223372036854775807
                                                v2 = append(v2, v31)
                                                v3 += (v3 >> (((v3 % ((629) | 1))) & 63))
                                                _ = v31
                                            }
                                            _ = v29
                                        } else {
                                            switch "" {
                                            case "foo bar":
                                                switch ("a" + (("a" + "\n\t") + "a")) {
                                                case "\n\t":
                                                    for fi32 := 0; fi32 < 15; fi32++ {
                                                        var v33 int32 = v23
                                                        var v34 int64 = -9223372036854775808
                                                        v35 := GenS3{F0: 2}
                                                        _ = fi32
                                                        _ = v33
                                                        _ = v34
                                                        _ = v35
                                                    }
                                                default:
                                                    v36 := GenS4{F0: 472}
                                                    _ = v36
                                                }
                                            case "\n\t":
                                                var v37 uint64 = 2
                                                _ = v37
                                            }
                                        }
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                        _ = v28
                                    }
                                default:
                                    switch fi24 {
                                    case 42:
                                        v5 = ((fi24 % ((fi24) | 1)) >= fi24)
                                    case 32767:
                                        for fi38 := 0; fi38 < 15; fi38++ {
                                            var v39 string = ("foo bar" + "\n\t")
                                            if (!true) {
                                                var v40 int16 = 27
                                                _ = v40
                                            } else {
                                                v5 = (!(v5 || ((!(v27 == 0.0)) || v5)))
                                                if (!((v3 + 4294967295) != genFn0(fi38, int64(v23), 531))) {
                                                    var v41 uint32 = 2147483647
                                                    v42 := GenS5{F0: v41, F1: uint8(v41), F2: 100}
                                                    _ = v41
                                                    _ = v42
                                                }
                                            }
                                            var v43 int32 = (v23 - (v23 << ((genFn0(fi24, int64(v23), 65535)) & 63)))
                                            _ = fi38
                                            _ = v39
                                            _ = v43
                                        }
                                    }
                                }
                            } else {
                                if v5 {
                                    for fi44 := 0; fi44 < 5; fi44++ {
                                        v3 = 255
                                        v23 += (-(v23))
                                        fmt.Println(v4)
                                        _ = fi44
                                    }
                                    var v45 float32 = 1.5
                                    _ = v45
                                }
                                for fi46 := 0; fi46 < 10; fi46++ {
                                    v47 := GenS4{F0: 718}
                                    v48 := GenS6{F0: "a"}
                                    _ = fi46
                                    _ = v47
                                    _ = v48
                                }
                                if len(v1) > 0 {
                                    v1[0] = (fi24 >> (((v3 / ((1) | 1))) & 63))
                                }
                            }
                            _ = v27
                        }
                        v23 += -948
                        _ = fi24
                    }
                    _ = v23
                }
            case 2147483647:
                switch 4294967295 {
                case 2147483647:
                    for fi49 := 0; fi49 < 3; fi49++ {
                        var v50 bool = v5
                        v5 = (!v50)
                        _ = fi49
                        _ = v50
                    }
                case -919:
                    var v51 int64 = 9223372036854775807
                    _ = v51
                case -554:
                    var v52 string = ((((("foo bar" + ("a" + "foo bar")) + "\n\t") + (("" + "hello") + "a")) + ("foo bar" + "a")) + "\n\t")
                    _ = v52
                default:
                    v9 = (v9 / ((v9) | 1))
                }
            }
        }
    }
    var v53 int64 = 1
    if v5 {
        var v54 bool = false
        v53 = (v9 - (-(v53)))
        _ = v54
    }
    v5 = ((v9 <= (-(v9))) && v5)
    v53 = (^(v53))
    if v5 {
        for fi55 := 0; fi55 < 7; fi55++ {
            fmt.Println(v9)
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            _ = fi55
        }
        _ = len(v1)
        fmt.Println(v4, v53, v9)
    } else {
        var v56 int64 = v53
        var v57 int64 = v56
        v58 := GenS6{F0: ("foo bar" + "a")}
        _ = v56
        _ = v57
        _ = v58
    }
    var v59 bool = ((v3 ^ 164) < (v3 / ((v3) | 1)))
    for fi60 := 0; fi60 < 20; fi60++ {
        v61 := GenS5{F0: 2, F1: 42, F2: (v4 | v4)}
        _ = fi60
        _ = v61
    }
    var v62 uint16 = 32767
    v63 := GenS4{F0: 18446744073709551615}
    v64 := GenS7{F0: "", F1: 461}
    for fi65 := 0; fi65 < 1; fi65++ {
        switch ("a" + "a") {
        case "a":
            switch (fi65 - (fi65 - (fi65 / ((fi65) | 1)))) {
            case -159:
                v5 = (true || ((4294967295 < 2147483647) || true))
            case 127:
                var v66 uint16 = (v62 & (v62 | v62))
                _ = v66
            case 1:
                v62 += v62
            default:
                var v67 uint16 = (v62 >> ((173) & 63))
                _ = v67
            }
        case "":
            switch (("a" + "") + "foo bar") {
            case "foo bar":
                if (-9223372036854775808 == v53) {
                    for fi68 := 0; fi68 < 12; fi68++ {
                        v69 := GenS1{F0: 162, F1: 160, F2: v53}
                        _ = fi68
                        _ = v69
                    }
                    if len(v2) > 0 {
                        v2[0] = fi65
                    }
                    for fi70 := 0; fi70 < 15; fi70++ {
                        var v71 uint64 = uint64(v62)
                        fmt.Println(v5)
                        if len(v2) > 0 {
                            v2[0] = fi65
                        }
                        _ = fi70
                        _ = v71
                    }
                } else {
                    v72 := GenS8{F0: 0, F1: false, F2: "hello"}
                    if len(v2) > 0 {
                        v2[0] = (fi65 % ((fi65) | 1))
                    }
                    _ = v72
                }
            default:
                for fi73 := 0; fi73 < 10; fi73++ {
                    fmt.Println(v53, v53)
                    v4 = (^(v4))
                    if (-2147483648 < 0) {
                        switch ((("\n\t" + ((("hello" + "") + ("foo bar" + "foo bar")) + (("a" + "a") + "hello"))) + "") + "hello") {
                        case "":
                            var v74 uint8 = 230
                            _ = v74
                        case "hello":
                            var v75 uint16 = v62
                            _ = v75
                        case "foo bar":
                            if v59 {
                                fmt.Println(fi73)
                                if v5 {
                                    v59 = (v5 || false)
                                    if false {
                                        v76 := GenS9{F0: 2147483648, F1: 65535, F2: v62}
                                        v77 := GenS10{F0: "hello", F1: 137}
                                        _ = v76
                                        _ = v77
                                    }
                                    v78 := GenS3{F0: 4294967295}
                                    var v79 int8 = (v4 % (((^(v4))) | 1))
                                    _ = v78
                                    _ = v79
                                }
                                fmt.Println(v53, v62, fi73)
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                            }
                        default:
                            v59 = v59
                        }
                    }
                    _ = fi73
                }
            }
        case "\n\t":
            v80 := GenS11{F0: int(v9), F1: 255, F2: 0}
            _ = v80
        }
        _ = fi65
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v53:", v53)
    fmt.Println("v59:", v59)
    fmt.Println("v62:", v62)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v53
    _ = v59
    _ = v62
    _ = v63
    _ = v64
    _ = v9
}
