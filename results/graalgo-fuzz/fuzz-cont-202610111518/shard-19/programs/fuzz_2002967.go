package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int32; F1 int16; F2 int }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 uint32; F1 int64; F2 uint32 }
type GenS7 struct { F0 bool }

func main() {
    var v1 uint32 = 42
    var v2 float32 = -2.25
    var v3 int32 = -2147483648
    var v4 uint8 = 149
    var v5 bool = (127 < 127)
    var v6 int16 = -758
    var v7 int32 = (v3 & v3)
    for fi8 := 0; fi8 < 15; fi8++ {
        var v9 bool = false
        _ = fi8
        _ = v9
    }
    switch -955 {
    case 9223372036854775807:
        switch "foo bar" {
        case "hello":
            v10 := GenS1{F0: ("" + "hello")}
            _ = v10
        case "foo bar":
            v11 := GenS1{F0: ("" + "")}
            _ = v11
        case "":
            v12 := GenS1{F0: ""}
            _ = v12
        default:
            v3 = int32(v4)
        }
    case -2147483648:
        v6 += 1
    case -629:
        for fi13 := 0; fi13 < 7; fi13++ {
            for fi14 := 0; fi14 < 3; fi14++ {
                if ("hello" == ("foo bar" + "\n\t")) {
                    v15 := GenS1{F0: "foo bar"}
                    v16 := []int{(-(fi14)), fi14}
                    _ = v15
                    _ = v16
                } else {
                    fmt.Println(fi14)
                    fmt.Println(v7)
                    var v17 uint32 = (^(v1))
                    _ = v17
                }
                _ = fi14
            }
            fmt.Println(v7, v5)
            var v18 int8 = 100
            _ = fi13
            _ = v18
        }
    }
    v19 := GenS1{F0: ("" + "foo bar")}
    v20 := []int{}
    switch 2147483647 {
    case -2147483648:
        var v21 uint8 = (v4 - 255)
        _ = v21
    case 0:
        v22 := GenS1{F0: "\n\t"}
        _ = v22
    case -55:
        if (false || (true || true)) {
            v23 := GenS1{F0: "a"}
            if ((v4 & (v4 / ((v4) | 1))) != (v4 | v4)) {
                var v24 float64 = (((-2.25 - ((3.14159 + float64(((v2 / v2) - v2))) / 1.0)) * 3.14159) + 1e10)
                switch "" {
                case "hello":
                    fmt.Println(v2, v2, v6)
                case "foo bar":
                    _ = len(v20)
                default:
                    for fi25 := 0; fi25 < 1; fi25++ {
                        var v26 uint32 = (v1 & (v1 / ((2147483648) | 1)))
                        if (-1 > v6) {
                            for fi27 := 0; fi27 < 19; fi27++ {
                                v28 := GenS1{F0: ""}
                                _ = fi27
                                _ = v28
                            }
                            var v29 bool = v5
                            v29 = (v29 && v5)
                            _ = v29
                        }
                        var v30 float32 = (v2 - 1.5)
                        _ = fi25
                        _ = v26
                        _ = v30
                    }
                }
                _ = v24
            } else {
                fmt.Println(v3, v5, v3)
                if (float64((v2 / (0.1 / (v2 * v2)))) == ((-2.25 / float64(1.5)) * ((((3.14159 / 0.1) + 0.1) - (0.1 / (1.5 * 0.1))) + (((-2.25 * (-2.25 + 0.1)) / 1.0) * 0.1)))) {
                    var v31 uint = 311
                    v2 -= (3.14159 * 3.14159)
                    _ = v31
                }
                var v32 int64 = 0
                _ = v32
            }
            v33 := GenS1{F0: ("foo bar" + "")}
            _ = v23
            _ = v33
        }
    default:
        _ = len(v20)
    }
    fmt.Println(v4, v2, v3)
    v34 := GenS2{F0: 1}
    for fi35 := 0; fi35 < 20; fi35++ {
        switch "hello" {
        case "":
            v20 = append(v20, int(v1))
        case "foo bar":
            if len(v20) >= 2 {
                v20 = v20[1:len(v20)]
            }
        case "hello":
            if (!(1.5 < (float32(1.5) + v2))) {
                switch "" {
                case "a":
                    if true {
                        _ = len(v20)
                        v36 := GenS2{F0: 42}
                        _ = v36
                    }
                default:
                    var v37 string = ("\n\t" + ((("hello" + ("a" + "foo bar")) + "hello") + "a"))
                    _ = v37
                }
            } else {
                _ = len(v20)
            }
        default:
            switch "" {
            case "foo bar":
                v38 := GenS3{F0: v3, F1: 912, F2: fi35}
                _ = v38
            case "":
                if len(v20) > 0 {
                    v20[0] = (fi35 ^ fi35)
                }
            default:
                v7 = 1
            }
        }
        if v5 {
            v39 := GenS2{F0: 4294967295}
            v40 := GenS4{F0: v4}
            _ = v39
            _ = v40
        } else {
            for fi41 := 0; fi41 < 2; fi41++ {
                if (52 == uint8(v1)) {
                    if v5 {
                        if (42 <= 18446744073709551615) {
                            if (4294967295 > 18446744073709551615) {
                                var v42 int32 = (v7 | v7)
                                _ = v42
                            } else {
                                if true {
                                    v4 = (v4 - v4)
                                    v3 = (v7 << ((4294967295) & 63))
                                    switch (fi41 & fi41) {
                                    case -239:
                                        var v43 int = fi35
                                        if (v3 != v3) {
                                            if (false && false) {
                                                v44 := GenS1{F0: "foo bar"}
                                                v45 := GenS5{F0: 0}
                                                _ = v44
                                                _ = v45
                                            }
                                            v46 := GenS5{F0: 18446744073709551615}
                                            var v47 uint8 = (v4 ^ 1)
                                            _ = v46
                                            _ = v47
                                        } else {
                                            if v5 {
                                                var v48 uint16 = 32767
                                                if false {
                                                    v4 += (^(v4))
                                                    if (-1 == -478) {
                                                        fmt.Println(v2, v2, v3)
                                                        var v49 int64 = 9223372036854775807
                                                        var v50 int = fi41
                                                        v20 = append(v20, v43)
                                                        _ = v49
                                                        _ = v50
                                                    }
                                                    if true {
                                                        v51 := GenS2{F0: 2}
                                                        _ = v51
                                                    }
                                                }
                                                _ = v48
                                            }
                                        }
                                        _ = v43
                                    case 1:
                                        _ = len(v20)
                                    }
                                    switch "foo bar" {
                                    case "\n\t":
                                        var v52 uint32 = 2147483647
                                        _ = v52
                                    default:
                                        var v53 int32 = v7
                                        _ = v53
                                    }
                                }
                            }
                            fmt.Println(v6, fi41, v2)
                            var v54 bool = (v6 > (v6 + v6))
                            _ = v54
                        } else {
                            v20 = append(v20, 32767)
                            v6 += 42
                        }
                        var v55 bool = true
                        _ = v55
                    } else {
                        var v56 bool = ((0 >= 0) || (!(!v5)))
                        fmt.Println(v56, v3, v6)
                        var v57 int16 = (-(v6))
                        v5 = (v1 > uint32(v4))
                        v1 -= (^(v1))
                        _ = v56
                        _ = v57
                    }
                    v58 := GenS4{F0: 42}
                    _ = v58
                } else {
                    var v59 int16 = v6
                    var v60 uint32 = v1
                    _ = len(v20)
                    _ = v59
                    _ = v60
                }
                if (!(!v5)) {
                    var v61 int64 = 9223372036854775807
                    _ = v61
                } else {
                    if false {
                        var v62 float64 = 3.14159
                        v20 = append(v20, fi35)
                        v6 = (-(v6))
                        _ = v62
                    } else {
                        if true {
                            v4 = v4
                            v4 = (v4 - (^(v4)))
                            var v63 int32 = v3
                            _ = v63
                        }
                        if ((("foo bar" + "foo bar") + "\n\t") != "\n\t") {
                            v64 := GenS4{F0: (v4 * 255)}
                            if len(v20) >= 2 {
                                v20 = v20[1:len(v20)]
                            }
                            switch "a" {
                            case "":
                                if len(v20) > 0 {
                                    v20[0] = 9223372036854775807
                                }
                            case "hello":
                                var v65 uint32 = 597
                                _ = v65
                            case "foo bar":
                                v66 := GenS6{F0: (v1 + v1), F1: 386, F2: (v1 & (v1 % ((2) | 1)))}
                                _ = v66
                            }
                            _ = v64
                        } else {
                            fmt.Println(v1, v5, v6)
                            v2 -= v2
                            v20 = append(v20, int(v3))
                            v6 = 32767
                            var v67 float64 = 0.0
                            fmt.Println(v2, v5)
                            if len(v20) > 0 {
                                v20[0] = (fi41 % ((fi41) | 1))
                            }
                            var v68 uint64 = 299
                            switch 282 {
                            case 0:
                                v67 += 1e10
                                switch "foo bar" {
                                case "":
                                    v69 := GenS7{F0: v5}
                                    _ = v69
                                case "foo bar":
                                    var v70 uint8 = v4
                                    _ = v70
                                }
                                v20 = append(v20, (fi35 / ((2147483647) | 1)))
                            case 32768:
                                var v71 float32 = v2
                                _ = v71
                            case 817:
                                if true {
                                    if (4294967295 == 4294967295) {
                                        var v72 int8 = -26
                                        v72 += v72
                                        _ = v72
                                    }
                                    v73 := GenS2{F0: v68}
                                    v3 = int32(v4)
                                    _ = v73
                                }
                            }
                            _ = v67
                            _ = v68
                        }
                    }
                    fmt.Println(v1)
                    v1 += (v1 << ((843) & 63))
                }
                _ = fi41
            }
            v20 = append(v20, (-(fi35)))
        }
        _ = fi35
    }
    var v74 int16 = (v6 % (((v6 >> ((1) & 63))) | 1))
    if v5 {
        if len(v20) > 0 {
            v20[0] = -9223372036854775808
        }
        v2 = float32((-2.25 - -2.25))
        fmt.Println(v7, v6, v3)
    } else {
        v75 := GenS7{F0: ("hello" != ("foo bar" + "hello"))}
        var v76 bool = (v7 >= v7)
        v77 := GenS4{F0: v4}
        _ = v75
        _ = v76
        _ = v77
    }
    v7 = (v3 * (v3 * v3))
    _ = len(v20)
    if false {
        fmt.Println(v74)
    } else {
        if len(v20) > 0 {
            v20[0] = 32767
        }
    }
    var v78 uint64 = 654
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v74:", v74)
    fmt.Println("v78:", v78)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v34
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v74
    _ = v78
}
