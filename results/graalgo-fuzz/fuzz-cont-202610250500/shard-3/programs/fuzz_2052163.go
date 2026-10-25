package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int16 }
type GenS2 struct { F0 uint8; F1 uint32; F2 int32 }
type GenS3 struct { F0 uint64; F1 int16; F2 uint32 }
type GenS4 struct { F0 bool; F1 uint; F2 int64 }
type GenS5 struct { F0 uint16; F1 int16 }
type GenS6 struct { F0 string }
func genFn0(p0 int32, p1 uint, p2 int8) int64 {
    p1 = (^(p1))
    var v7 int16 = 0
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 0
}

func main() {
    var v1 uint64 = 681
    v2 := make(map[string]int)
    var v3 float32 = 0.0
    var v4 uint64 = v1
    v5 := []int{42}
    var v6 bool = ((0.0 == 0.0) && true)
    for fi8 := 0; fi8 < 12; fi8++ {
        var v9 uint32 = 65535
        var v10 int = fi8
        _ = fi8
        _ = v10
        _ = v9
    }
    if ((4294967295 < 32767) && false) {
        switch "foo bar" {
        case "\n\t":
            switch int(v4) {
            case 32768:
                v6 = true
            default:
                var v11 int32 = 0
                _ = v11
            }
        case "foo bar":
            var v12 float64 = 0.1
            _ = v12
        default:
            if (genFn0(-881, 4294967295, 42) > -9223372036854775808) {
                for fi13 := 0; fi13 < 16; fi13++ {
                    v6 = (v1 != v4)
                    switch "" {
                    case "hello":
                        switch 32768 {
                        case 128:
                            for fi14 := 0; fi14 < 8; fi14++ {
                                v15 := GenS1{F0: (v4 << ((372) & 63)), F1: -1}
                                fmt.Println(v6, v6)
                                _ = fi14
                                _ = v15
                            }
                        case -2147483648:
                            var v16 uint16 = 234
                            _ = v16
                        case 32767:
                            _ = len(v5)
                        }
                    case "a":
                        switch 127 {
                        case 2147483647:
                            var v17 bool = (2147483647 >= 65535)
                            _ = v17
                        }
                    case "\n\t":
                        if (0 < 127) {
                            var v18 int16 = 1000
                            fmt.Println(v3, v1, v4)
                            v19 := GenS2{F0: 255, F1: 2147483647, F2: 2147483647}
                            _ = v18
                            _ = v19
                        }
                    }
                    _ = fi13
                }
                v20 := GenS2{F0: 128, F1: 2147483648, F2: -11}
                if (v6 || (!("" == "foo bar"))) {
                    fmt.Println(v3)
                    _ = len(v2)
                }
                _ = v20
            } else {
                fmt.Println(v4, v6)
                v21 := GenS3{F0: (v4 & 969), F1: 0, F2: 4294967295}
                _ = v21
            }
        }
        v3 += 0.0
        var v22 uint64 = (^(v4))
        _ = v22
    } else {
        for fi23 := 0; fi23 < 3; fi23++ {
            v24 := GenS2{F0: 127, F1: 718, F2: 2147483647}
            switch (^(fi23)) {
            case 208:
                if ((!v6) || v6) {
                    if (100 >= 127) {
                        var v25 uint32 = 2147483648
                        if false {
                            v3 = v3
                            if v6 {
                                for fi26 := 0; fi26 < 13; fi26++ {
                                    var v27 uint = 2
                                    v28 := GenS4{F0: true, F1: v27, F2: 9223372036854775807}
                                    v4 += (v1 * (v1 & 0))
                                    _ = fi26
                                    _ = v27
                                    _ = v28
                                }
                                if len(v5) >= 2 {
                                    v5 = v5[1:len(v5)]
                                }
                                for fi29 := 0; fi29 < 2; fi29++ {
                                    switch ("a" + "hello") {
                                    case "\n\t":
                                        v30 := v2["foo bar"]
                                        _ = v30
                                    default:
                                        var v31 float32 = 0.0
                                        _ = v31
                                    }
                                    fmt.Println(v4)
                                    _ = fi29
                                }
                            } else {
                                fmt.Println(fi23, v3)
                                fmt.Println(v4, v6, v25)
                            }
                        } else {
                            v1 = (v1 | (^(v4)))
                            var v32 int8 = 42
                            _ = v32
                        }
                        var v33 int8 = 0
                        _ = v25
                        _ = v33
                    } else {
                        v34 := GenS1{F0: 2, F1: 32767}
                        _ = v34
                    }
                    v2["foo bar"] = (fi23 / ((fi23) | 1))
                }
            }
            _ = fi23
            _ = v24
        }
    }
    v3 += v3
    fmt.Println(v3, v3, v1)
    if (0.1 > 0.0) {
        var v35 uint16 = 0
        _ = v35
    } else {
        if true {
            v3 = float32(1e10)
            if (false || v6) {
                if true {
                    var v36 uint8 = 1
                    fmt.Println(v4, v3)
                    if v6 {
                        _ = len(v5)
                    } else {
                        for fi37 := 0; fi37 < 10; fi37++ {
                            var v38 int16 = 32767
                            _ = fi37
                            _ = v38
                        }
                        switch 539 {
                        case 9223372036854775807:
                            var v39 int8 = 42
                            _ = v39
                        case -9223372036854775808:
                            if ((true || v6) || false) {
                                if ((v3 * (v3 + 1e10)) == v3) {
                                    if false {
                                        v2[("foo bar" + "\n\t")] = 42
                                    } else {
                                        if len(v5) > 0 {
                                            v5[0] = 1
                                        }
                                        var v40 int8 = 123
                                        _ = v40
                                    }
                                } else {
                                    v3 = (-2.25 / ((0.0 * v3) * v3))
                                }
                                _ = len(v5)
                            } else {
                                if (!(1.5 == -2.25)) {
                                    var v41 uint32 = 2147483648
                                    _ = v41
                                } else {
                                    v5 = append(v5, -1)
                                    v42 := GenS3{F0: 42, F1: -32768, F2: 668}
                                    if v6 {
                                        switch -2147483648 {
                                        case 127:
                                            var v43 string = "a"
                                            _ = v43
                                        }
                                        fmt.Println(v3, v6)
                                        for fi44 := 0; fi44 < 15; fi44++ {
                                            v1 = (v4 - uint64(v36))
                                            v1 = 1
                                            _ = fi44
                                        }
                                    }
                                    _ = v42
                                }
                                if ((-570 >= -195) || true) {
                                    var v45 float64 = ((-2.25 - 0.0) - 0.1)
                                    for fi46 := 0; fi46 < 11; fi46++ {
                                        var v47 int16 = -1
                                        _ = fi46
                                        _ = v47
                                    }
                                    var v48 uint16 = 2
                                    _ = v45
                                    _ = v48
                                }
                            }
                        case 32768:
                            if (1e10 != (1e10 + 3.14159)) {
                                var v49 int8 = 100
                                _ = v49
                            }
                        }
                        v50 := v2["foo bar"]
                        _ = v50
                    }
                    _ = v36
                }
                if (127 <= 127) {
                    var v51 int = 32767
                    _ = v51
                }
                var v52 int32 = 2147483647
                _ = v52
            }
        }
        for fi53 := 0; fi53 < 18; fi53++ {
            v6 = (!true)
            _ = fi53
        }
        var v54 uint16 = 0
        _ = v54
    }
    if v6 {
        if (v3 >= (v3 / v3)) {
            fmt.Println(v3)
            if v6 {
                v55 := GenS2{F0: 242, F1: 2, F2: 646}
                _ = v55
            }
            v1 = v4
        } else {
            v4 -= v4
        }
    } else {
        for fi56 := 0; fi56 < 4; fi56++ {
            var v57 bool = (2147483647 > -62)
            v4 = (v1 << ((436) & 63))
            _ = fi56
            _ = v57
        }
        v3 += (v3 - v3)
        var v58 int16 = -1
        _ = v58
    }
    if ((!false) || ((9223372036854775807 >= 0) && v6)) {
        v5 = append(v5, 1)
        if (790 < 403) {
            if (!false) {
                v6 = (!(0.0 > 3.14159))
                v6 = (genFn0(412, 2, 32) <= genFn0(-2147483648, 65535, 1))
                for fi59 := 0; fi59 < 1; fi59++ {
                    v60 := GenS2{F0: 124, F1: 252, F2: 42}
                    v61 := GenS1{F0: v4, F1: -378}
                    _ = fi59
                    _ = v60
                    _ = v61
                }
            } else {
                v62 := GenS3{F0: 1, F1: 604, F2: 2}
                var v63 uint = 1
                _ = v62
                _ = v63
            }
            v5 = append(v5, 4294967295)
        }
        fmt.Println(v4, v1, v3)
    }
    v6 = ((v1 | (v1 & 2)) < (v4 - (^(v4))))
    for fi64 := 0; fi64 < 9; fi64++ {
        for fi65 := 0; fi65 < 6; fi65++ {
            switch fi64 {
            case -753:
                v66 := GenS1{F0: v4, F1: 477}
                _ = v66
            case -724:
                delete(v2, "hello")
            case 287:
                v6 = (1 != 42)
            default:
                var v67 int = 127
                v6 = false
                _ = v67
            }
            _ = fi65
        }
        _ = fi64
    }
    v68 := GenS4{F0: false, F1: 437, F2: -9223372036854775808}
    for fi69 := 0; fi69 < 11; fi69++ {
        v3 -= ((((v3 / 0.1) + 0.1) - v3) * (v3 / (1e10 - ((0.1 - float32(float64(((((((v3 * (v3 + float32(1.5))) / v3) - v3) + (float32(-2.25) / 1.5)) - 3.14159) * v3)))) + 3.14159))))
        v4 += v1
        if (1 >= 1) {
            fmt.Println(v3, v3)
            for fi70 := 0; fi70 < 11; fi70++ {
                var v71 uint = 789
                switch "a" {
                case "":
                    v3 = (((v3 * 3.14159) / (v3 + (1.5 - 0.0))) / 1.5)
                case "hello":
                    switch "a" {
                    case "hello":
                        var v72 int = (fi69 << ((v71) & 63))
                        _ = v72
                    case "a":
                        v73 := GenS5{F0: 1, F1: 32767}
                        _ = v73
                    case "foo bar":
                        v1 += v1
                    default:
                        v74 := GenS6{F0: "hello"}
                        _ = v74
                    }
                default:
                    var v75 int32 = 892
                    _ = v75
                }
                if ((fi70 > 185) && v6) {
                    fmt.Println(fi70, v3)
                }
                v6 = (225 > 180)
                v76 := GenS2{F0: 184, F1: 975, F2: -2147483648}
                v6 = (!(!(fi70 < (fi69 >> (((v71 + v71)) & 63)))))
                _ = fi70
                _ = v71
                _ = v76
            }
            if (!(-9223372036854775808 == 938)) {
                fmt.Println(v6, v1, v4)
            } else {
                if v6 {
                    fmt.Println(v6)
                }
                v3 = v3
            }
        } else {
            for fi77 := 0; fi77 < 19; fi77++ {
                v6 = (!(0 <= 1))
                _ = fi77
            }
        }
        _ = fi69
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    v3 -= v3
    for fi78 := 0; fi78 < 11; fi78++ {
        v3 = ((v3 * float32(0.0)) - v3)
        var v79 uint8 = 0
        _ = fi78
        _ = v79
    }
    v6 = (!("\n\t" <= "hello"))
    v6 = (2 == 4294967295)
    for fi80 := 0; fi80 < 6; fi80++ {
        v81 := v2[(("foo bar" + (("a" + "hello") + "hello")) + (("foo bar" + "hello") + ("\n\t" + "\n\t")))]
        fmt.Println(v1)
        _ = fi80
        _ = v81
    }
    var v82 float32 = v3
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v82:", v82)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v68
    _ = v82
}
