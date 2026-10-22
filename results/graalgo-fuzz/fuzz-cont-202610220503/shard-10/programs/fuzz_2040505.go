package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint8; F1 uint8 }
type GenS3 struct { F0 uint64; F1 uint; F2 int8 }
type GenS4 struct { F0 uint32; F1 int32 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 uint8; F1 string; F2 uint64 }
type GenS7 struct { F0 uint8; F1 string; F2 uint }
func genFn0() uint16 {
    v6 := GenS1{F0: 1}
    if (0 != 18446744073709551615) {
        var v7 int32 = -492
        v8 := GenS2{F0: 0, F1: 128}
        _ = v7
        _ = v8
    }
    _ = v6
    return 2
}
func genFn1() int {
    if (false || true) {
        var v9 uint8 = 0
        var v10 float32 = (((0.1 + float32((0.0 * float64(((-2.25 - -2.25) - 0.0))))) - -2.25) - float32(((3.14159 + 0.1) / 1e10)))
        _ = v10
        _ = v9
    }
    v11 := []int{128}
    for fi12 := 0; fi12 < 11; fi12++ {
        var v13 int32 = 2147483647
        _ = fi12
        _ = v13
    }
    _ = v11
    return 9223372036854775807
}

func main() {
    v1 := make(map[int]bool)
    var v2 int64 = 1
    var v3 int32 = 42
    v4 := make(map[int]string)
    v5 := make(map[int]int)
    var v14 uint32 = 2
    _ = len(v5)
    var v15 int = -9223372036854775808
    delete(v4, v15)
    v2 -= -9223372036854775808
    v16 := []int{-362, (v15 ^ (v15 << ((65535) & 63))), -716}
    if ((v2 <= (v2 << ((203) & 63))) || true) {
        v14 -= 42
        var v17 int16 = -199
        _ = v17
    }
    var v18 int16 = 912
    var v19 int16 = 456
    if true {
        v3 = (^(v3))
        switch "\n\t" {
        case "":
            if false {
                if (v15 == v15) {
                    fmt.Println(v14, v18, v19)
                }
                v3 += (v3 >> ((0) & 63))
                fmt.Println(v19, v19)
            } else {
                v20 := GenS3{F0: 966, F1: 759, F2: 19}
                for fi21 := 0; fi21 < 6; fi21++ {
                    var v22 int32 = (v3 % (((v3 / ((v3) | 1))) | 1))
                    _ = fi21
                    _ = v22
                }
                if (!(-9223372036854775808 < 944)) {
                    if (!false) {
                        switch (v15 ^ (v15 & genFn1())) {
                        case 1:
                            var v23 uint16 = 65535
                            _ = v23
                        case 0:
                            for fi24 := 0; fi24 < 9; fi24++ {
                                for fi25 := 0; fi25 < 16; fi25++ {
                                    v14 = 42
                                    if len(v16) > 0 {
                                        v16[0] = (fi25 & fi24)
                                    }
                                    _ = fi25
                                }
                                for fi26 := 0; fi26 < 16; fi26++ {
                                    if (v3 == (v3 << ((804) & 63))) {
                                        var v27 int = 998
                                        switch fi26 {
                                        case -233:
                                            v2 += 1
                                        case -882:
                                            v28 := GenS4{F0: (v14 / (((v14 >> ((360) & 63))) | 1)), F1: v3}
                                            _ = v28
                                        case 0:
                                            if len(v16) >= 2 {
                                                v16 = v16[1:len(v16)]
                                            }
                                        default:
                                            v19 = v19
                                        }
                                        _ = v27
                                    } else {
                                        v29 := GenS3{F0: 42, F1: 2, F2: -15}
                                        _ = v29
                                    }
                                    _ = fi26
                                }
                                _ = len(v16)
                                _ = fi24
                            }
                        case -2147483648:
                            var v30 uint16 = uint16(v14)
                            _ = v30
                        }
                    } else {
                        v2 -= v2
                        v31 := GenS5{F0: v15}
                        var v32 int64 = v2
                        _ = v31
                        _ = v32
                    }
                } else {
                    v33 := GenS6{F0: 255, F1: (("a" + "") + ("foo bar" + "\n\t")), F2: 18446744073709551615}
                    v16 = append(v16, v15)
                    _ = v33
                }
                _ = v20
            }
        case "hello":
            fmt.Println(v14, v14)
        }
    } else {
        for fi34 := 0; fi34 < 20; fi34++ {
            if true {
                for fi35 := 0; fi35 < 3; fi35++ {
                    var v36 int64 = 9223372036854775807
                    var v37 uint = 255
                    v15 = int(v37)
                    var v38 int32 = (-(v3))
                    _ = fi35
                    _ = v36
                    _ = v37
                    _ = v38
                }
            } else {
                for fi39 := 0; fi39 < 1; fi39++ {
                    var v40 float32 = (0.1 + (float32(0.0) * (((float32((0.0 + 3.14159)) - ((3.14159 + 0.1) - ((float32(1e10) - 0.1) / 1.0))) / 1.5) - 0.0)))
                    v15 += (^(fi39))
                    v14 = v14
                    _ = fi39
                    _ = v40
                }
                var v41 uint8 = 42
                _ = v41
            }
            for fi42 := 0; fi42 < 5; fi42++ {
                if (!false) {
                    fmt.Println(v15)
                    var v43 uint64 = 950
                    _ = v43
                }
                _ = fi42
            }
            _ = fi34
        }
        for fi44 := 0; fi44 < 2; fi44++ {
            switch ("foo bar" + (("hello" + ("" + "foo bar")) + "foo bar")) {
            case "foo bar":
                v14 = 2147483647
            case "hello":
                var v45 int64 = (v2 << ((18446744073709551615) & 63))
                _ = v45
            }
            _ = fi44
        }
        v15 = (v15 - (^(v15)))
    }
    var v46 float32 = -2.25
    fmt.Println(v3)
    for fi47 := 0; fi47 < 5; fi47++ {
        if true {
            fmt.Println(v3)
            v48 := GenS5{F0: fi47}
            var v49 uint16 = 65535
            _ = v48
            _ = v49
        } else {
            var v50 int64 = v2
            _ = v50
        }
        v46 += 1.5
        switch 32768 {
        case 4294967295:
            if false {
                if true {
                    var v51 uint8 = 1
                    fmt.Println(v2, v18)
                    var v52 int64 = v2
                    _ = v51
                    _ = v52
                } else {
                    v53 := GenS2{F0: 127, F1: 1}
                    if ((int8(v3) >= 127) || (244 < 2)) {
                        v46 = 1.5
                    } else {
                        if (-2.25 > (1e10 - (((0.1 + ((0.1 - 1.5) - 0.0)) - 1.5) + (float64(v46) + (1e10 - 0.1))))) {
                            var v54 int64 = (v2 % ((v2) | 1))
                            var v55 uint32 = (v14 * v14)
                            _ = v54
                            _ = v55
                        } else {
                            var v56 uint64 = 4294967295
                            switch "" {
                            case "a":
                                if false {
                                    fmt.Println(v3)
                                } else {
                                    if (!false) {
                                        var v57 int16 = -32768
                                        var v58 uint32 = v14
                                        if true {
                                            if (0 > v14) {
                                                fmt.Println(v14, v15)
                                                v59 := v1[int(v19)]
                                                _ = v59
                                            } else {
                                                for fi60 := 0; fi60 < 9; fi60++ {
                                                    switch -2147483648 {
                                                    case 9223372036854775807:
                                                        var v61 uint = 42
                                                        _ = v61
                                                    case 0:
                                                        var v62 uint32 = v58
                                                        _ = v62
                                                    }
                                                    _ = fi60
                                                }
                                                switch genFn1() {
                                                case 127:
                                                    switch v15 {
                                                    case 127:
                                                        v4[(v15 ^ (fi47 + (v15 | genFn1())))] = ("a" + "a")
                                                    }
                                                default:
                                                    v63 := GenS1{F0: v2}
                                                    _ = v63
                                                }
                                                fmt.Println(fi47)
                                            }
                                        } else {
                                            v64 := GenS1{F0: int64(fi47)}
                                            var v65 uint32 = 229
                                            _ = v64
                                            _ = v65
                                        }
                                        _ = v57
                                        _ = v58
                                    }
                                    for fi66 := 0; fi66 < 1; fi66++ {
                                        var v67 uint8 = 255
                                        var v68 uint = 2
                                        var v69 uint16 = 42
                                        _ = fi66
                                        _ = v67
                                        _ = v68
                                        _ = v69
                                    }
                                }
                            }
                            v70 := GenS4{F0: v14, F1: v3}
                            _ = v56
                            _ = v70
                        }
                    }
                    _ = v53
                }
            }
        case 1:
            if (false && true) {
                var v71 uint64 = 2
                if len(v16) > 0 {
                    v16[0] = v15
                }
                _ = v71
            }
        case 128:
            switch ("hello" + ("" + ("\n\t" + "foo bar"))) {
            case "foo bar":
                v72 := GenS3{F0: 42, F1: 912, F2: 127}
                _ = v72
            case "":
                v19 -= (v18 * 1)
            case "hello":
                v73 := v4[1]
                _ = v73
            }
        }
        _ = fi47
    }
    var v74 int64 = (v2 ^ 0)
    for fi75 := 0; fi75 < 7; fi75++ {
        switch 979 {
        case 733:
            v46 -= 1.5
        case -2147483648:
            if (v2 < -1) {
                v19 = v18
                var v76 float32 = v46
                var v77 string = "a"
                _ = v76
                _ = v77
            } else {
                for fi78 := 0; fi78 < 19; fi78++ {
                    v3 = v3
                    _ = fi78
                }
                fmt.Println(v18, v15, v18)
                delete(v4, fi75)
            }
        default:
            for fi79 := 0; fi79 < 7; fi79++ {
                if (980 <= 42) {
                    if (genFn0() > 453) {
                        v14 = v14
                        var v80 int = -581
                        var v81 int8 = -128
                        _ = v80
                        _ = v81
                    }
                } else {
                    var v82 uint16 = 1
                    _ = v82
                }
                v15 -= (fi79 - genFn1())
                _ = fi79
            }
        }
        v2 = v74
        v83 := GenS7{F0: 128, F1: ("hello" + "a"), F2: 61}
        _ = fi75
        _ = v83
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v46:", v46)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v74:", v74)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v46
    _ = v5
    _ = v74
}
