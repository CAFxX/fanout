package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64 }
type GenS2 struct { F0 uint16; F1 uint64; F2 uint8 }
type GenS3 struct { F0 uint; F1 int64 }
type GenS4 struct { F0 uint32; F1 int32; F2 int }
type GenS5 struct { F0 int32; F1 string }
type GenS6 struct { F0 int32; F1 string; F2 uint8 }
type GenS7 struct { F0 int8; F1 uint8 }
func genFn0() string {
    if true {
        v7 := GenS1{F0: 1, F1: 4294967295}
        var v8 uint32 = 0
        _ = v7
        _ = v8
    }
    for fi9 := 0; fi9 < 20; fi9++ {
        for fi10 := 0; fi10 < 13; fi10++ {
            var v11 float32 = 1e10
            v11 = 0.0
            var v12 int32 = 1
            _ = fi10
            _ = v11
            _ = v12
        }
        _ = fi9
    }
    for fi13 := 0; fi13 < 15; fi13++ {
        var v14 bool = false
        fmt.Println(fi13, fi13)
        _ = fi13
        _ = v14
    }
    if (!false) {
        var v15 uint8 = 127
        var v16 uint16 = 2
        _ = v15
        _ = v16
    } else {
        if false {
            if true {
                var v17 int64 = 1
                v18 := GenS2{F0: 495, F1: 1, F2: 2}
                _ = v17
                _ = v18
            } else {
                switch "a" {
                case "a":
                    switch "hello" {
                    case "a":
                        var v19 uint = 1
                        _ = v19
                    }
                }
                var v20 float32 = (1.5 * 0.1)
                switch 128 {
                case 935:
                    var v21 float32 = (0.1 + v20)
                    _ = v21
                case -6:
                    var v22 int8 = 127
                    _ = v22
                case 399:
                    var v23 uint32 = 692
                    _ = v23
                }
                _ = v20
            }
            v24 := []int{32767, 4294967295, 4}
            _ = v24
        } else {
            var v25 uint16 = 185
            v25 = 0
            if (1e10 != 0.0) {
                for fi26 := 0; fi26 < 19; fi26++ {
                    fmt.Println(v25, v25)
                    v25 = v25
                    _ = fi26
                }
                v27 := GenS1{F0: 42, F1: 1}
                v28 := []int{127, -2147483648}
                _ = v27
                _ = v28
            }
            _ = v25
        }
    }
    return "a"
}

func main() {
    var v1 int16 = -32768
    var v2 int64 = -492
    v3 := []string{"foo bar", ("\n\t" + "foo bar")}
    v4 := make(map[string]uint32)
    v5 := []uint32{2}
    v6 := make(map[string]int)
    for fi29 := 0; fi29 < 8; fi29++ {
        if true {
            v6[("" + ((("\n\t" + ((("hello" + "") + ("hello" + "foo bar")) + ("foo bar" + "\n\t"))) + "foo bar") + ""))] = (fi29 << ((255) & 63))
            fmt.Println(v1, fi29, fi29)
        } else {
            if (true || (((!(false && false)) || (false || (((v2 & v2) == v2) || true))) && (-2147483648 == fi29))) {
                fmt.Println(v2, v2, v1)
                var v30 uint = 720
                var v31 string = "a"
                _ = v30
                _ = v31
            }
            v32 := v4["hello"]
            for fi33 := 0; fi33 < 11; fi33++ {
                v32 = v32
                _ = fi33
            }
            _ = v32
        }
        _ = fi29
    }
    if (474 == 4294967295) {
        v34 := GenS2{F0: 2, F1: 583, F2: 127}
        for fi35 := 0; fi35 < 2; fi35++ {
            delete(v4, "\n\t")
            for fi36 := 0; fi36 < 2; fi36++ {
                var v37 int32 = 2147483647
                var v38 bool = (10 < 18)
                _ = fi36
                _ = v37
                _ = v38
            }
            _ = fi35
        }
        _ = v34
    } else {
        switch 2147483647 {
        case 128:
            if (2147483647 <= -1) {
                if (!(!(919 > 1))) {
                    if (!false) {
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                    } else {
                        if true {
                            v39 := GenS3{F0: 65535, F1: (-(v2))}
                            if (true || true) {
                                var v40 int16 = -32768
                                _ = v40
                            }
                            _ = v39
                        } else {
                            v2 = v2
                            v1 = v1
                        }
                        if true {
                            if false {
                                fmt.Println(v2)
                                fmt.Println(v2)
                            }
                            if true {
                                v2 -= -314
                                var v41 int8 = 127
                                if len(v3) > 0 {
                                    v3[0] = (("\n\t" + ("" + (((("\n\t" + "foo bar") + ("foo bar" + ("" + "\n\t"))) + "\n\t") + ("" + ("a" + "hello"))))) + "hello")
                                }
                                _ = v41
                            } else {
                                for fi42 := 0; fi42 < 5; fi42++ {
                                    var v43 float32 = (0.0 / 1e10)
                                    var v44 bool = false
                                    var v45 string = ""
                                    _ = fi42
                                    _ = v43
                                    _ = v44
                                    _ = v45
                                }
                                var v46 string = ("\n\t" + "a")
                                _ = v46
                            }
                        }
                    }
                    for fi47 := 0; fi47 < 8; fi47++ {
                        for fi48 := 0; fi48 < 1; fi48++ {
                            if true {
                                var v49 int32 = 42
                                _ = v49
                            } else {
                                if false {
                                    var v50 int64 = v2
                                    _ = v50
                                } else {
                                    fmt.Println(fi47, fi48, v1)
                                    var v51 uint64 = 1
                                    _ = v51
                                }
                                fmt.Println(v1)
                            }
                            _ = fi48
                        }
                        if len(v5) > 0 {
                            v5[0] = 124
                        }
                        for fi52 := 0; fi52 < 11; fi52++ {
                            var v53 int64 = v2
                            if true {
                                if false {
                                    v53 -= v2
                                } else {
                                    v54 := GenS3{F0: 2, F1: (v2 % ((-65) | 1))}
                                    v55 := GenS4{F0: 2, F1: -1, F2: 9223372036854775807}
                                    _ = v54
                                    _ = v55
                                }
                                v53 -= v2
                                fmt.Println(fi52, v1)
                            }
                            _ = fi52
                            _ = v53
                        }
                        _ = fi47
                    }
                    switch "hello" {
                    case "foo bar":
                        for fi56 := 0; fi56 < 16; fi56++ {
                            switch "hello" {
                            case "\n\t":
                                var v57 int64 = (v2 & v2)
                                _ = v57
                            case "hello":
                                if (255 > 284) {
                                    v4["hello"] = 616
                                    var v58 uint64 = 0
                                    _ = v58
                                }
                            case "a":
                                for fi59 := 0; fi59 < 20; fi59++ {
                                    if false {
                                        if (int8(v2) >= 42) {
                                            fmt.Println(fi56, v2)
                                            v2 += -300
                                        }
                                        var v60 uint64 = 2
                                        if (50 >= 39) {
                                            switch ("foo bar" + (("foo bar" + "a") + "foo bar")) {
                                            case "foo bar":
                                                v1 = (v1 + (v1 / (((v1 * (v1 | (v1 | 0)))) | 1)))
                                            case "":
                                                fmt.Println(fi56)
                                            }
                                            if (-17 <= 42) {
                                                var v61 uint32 = 65535
                                                _ = len(v5)
                                                fmt.Println(fi56)
                                                _ = v61
                                            } else {
                                                v62 := GenS1{F0: v60, F1: (^(v60))}
                                                _ = v62
                                            }
                                        } else {
                                            var v63 float32 = 1e10
                                            _ = v63
                                        }
                                        _ = v60
                                    }
                                    if len(v5) > 0 {
                                        v5[0] = 2147483648
                                    }
                                    v1 = (-(v1))
                                    _ = fi59
                                }
                            }
                            v64 := GenS1{F0: 1, F1: 0}
                            v1 = (v1 % ((0) | 1))
                            _ = fi56
                            _ = v64
                        }
                    case "a":
                        var v65 int8 = -6
                        _ = v65
                    case "":
                        for fi66 := 0; fi66 < 20; fi66++ {
                            var v67 uint = 0
                            v1 += (v1 & (v1 & -850))
                            v1 += (v1 * (v1 | -1))
                            _ = fi66
                            _ = v67
                        }
                    default:
                        var v68 uint = 255
                        _ = v68
                    }
                } else {
                    v69 := v6[("" + "hello")]
                    _ = v69
                }
                var v70 int = -1
                _ = v70
            }
        case 0:
            var v71 uint16 = 439
            _ = v71
        case 42:
            fmt.Println(v2)
        default:
            for fi72 := 0; fi72 < 7; fi72++ {
                v1 = v1
                _ = fi72
            }
        }
    }
    v73 := GenS3{F0: 2, F1: (v2 << ((1) & 63))}
    v1 -= v1
    fmt.Println(v1)
    v1 -= (v1 * v1)
    v1 = v1
    switch "" {
    case "a":
        for fi74 := 0; fi74 < 16; fi74++ {
            var v75 string = ("\n\t" + "\n\t")
            v76 := GenS1{F0: 510, F1: 18446744073709551615}
            _ = fi74
            _ = v75
            _ = v76
        }
    case "foo bar":
        var v77 string = "hello"
        _ = v77
    case "\n\t":
        for fi78 := 0; fi78 < 9; fi78++ {
            if true {
                for fi79 := 0; fi79 < 9; fi79++ {
                    var v80 float64 = (-2.25 + -2.25)
                    _ = fi79
                    _ = v80
                }
                var v81 int64 = (v2 & (v2 >> ((2) & 63)))
                _ = v81
            }
            _ = fi78
        }
    default:
        for fi82 := 0; fi82 < 1; fi82++ {
            if (2 != 541) {
                var v83 bool = true
                _ = v83
            } else {
                v84 := GenS5{F0: 42, F1: "\n\t"}
                if true {
                    var v85 int64 = (v2 ^ v2)
                    _ = v85
                }
                fmt.Println(v1, v2, v2)
                _ = v84
            }
            fmt.Println(fi82, v1, fi82)
            var v86 uint32 = 65535
            _ = fi82
            _ = v86
        }
    }
    switch -9223372036854775808 {
    case 258:
        switch "a" {
        case "hello":
            v3 = append(v3, "a")
        case "":
            fmt.Println(v1)
        default:
            var v87 float64 = 1e10
            _ = v87
        }
    case -9223372036854775808:
        var v88 uint8 = 0
        _ = v88
    }
    v89 := GenS6{F0: 394, F1: "a", F2: 0}
    delete(v6, "foo bar")
    if true {
        v90 := GenS2{F0: 255, F1: 2, F2: 127}
        var v91 uint64 = 42
        switch 2147483647 {
        case -9223372036854775808:
            var v92 int = -950
            _ = v92
        default:
            v93 := v4[("hello" + "hello")]
            _ = v93
        }
        _ = v90
        _ = v91
    } else {
        for fi94 := 0; fi94 < 3; fi94++ {
            if false {
                v95 := v6[("a" + "hello")]
                _ = v95
            }
            v2 += 9223372036854775807
            _ = fi94
        }
        v96 := GenS7{F0: 0, F1: 127}
        v97 := GenS5{F0: 2147483647, F1: "hello"}
        _ = v96
        _ = v97
    }
    v1 = (v1 << ((933) & 63))
    if len(v5) > 0 {
        v5[0] = 4294967295
    }
    if (true && ("" <= "")) {
        v1 = -64
    } else {
        if (32768 <= 32768) {
            var v98 int64 = (v2 | 0)
            _ = v98
        } else {
            var v99 int16 = (^(v1))
            switch -1 {
            case -1:
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
            default:
                v2 = v2
            }
            _ = v99
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v73
    _ = v89
}
