package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint }
type GenS2 struct { F0 uint64; F1 bool; F2 uint32 }
type GenS3 struct { F0 uint32; F1 int }
type GenS4 struct { F0 uint64; F1 uint16; F2 uint8 }
type GenS5 struct { F0 int64; F1 string; F2 int64 }
type GenS6 struct { F0 int64; F1 uint8; F2 uint32 }
type GenS7 struct { F0 int64; F1 uint; F2 int32 }
type GenS8 struct { F0 int; F1 int32 }
type GenS9 struct { F0 int; F1 int32 }

func main() {
    var v1 int64 = -1
    v2 := []string{}
    var v3 uint32 = 4294967295
    var v4 string = ""
    var v5 int32 = 0
    v6 := []string{v4, v4, (v4 + ("" + v4))}
    v7 := GenS1{F0: 32767, F1: 18446744073709551615}
    v8 := GenS2{F0: 0, F1: (!(true && false)), F2: v3}
    v3 = 41
    if ((0 < 1) || true) {
        v1 -= v1
        switch 2147483647 {
        case -87:
            var v9 int8 = -128
            _ = v9
        case 0:
            if false {
                var v10 int8 = 1
                _ = v10
            } else {
                switch 0 {
                case 9223372036854775807:
                    fmt.Println(v1)
                default:
                    if false {
                        for fi11 := 0; fi11 < 13; fi11++ {
                            v3 += (v3 % (((v3 << ((uint(v3)) & 63))) | 1))
                            var v12 int8 = -1
                            _ = fi11
                            _ = v12
                        }
                        for fi13 := 0; fi13 < 8; fi13++ {
                            for fi14 := 0; fi14 < 14; fi14++ {
                                v6 = append(v6, v4)
                                _ = fi14
                            }
                            for fi15 := 0; fi15 < 15; fi15++ {
                                v1 -= (^(v1))
                                if len(v6) >= 2 {
                                    v6 = v6[1:len(v6)]
                                }
                                fmt.Println(fi13, fi13)
                                v5 -= (v5 << ((0) & 63))
                                switch "" {
                                case "foo bar":
                                    if true {
                                        v5 += 857
                                        v16 := GenS3{F0: (v3 | 699), F1: int(v5)}
                                        _ = v16
                                    } else {
                                        v17 := GenS1{F0: 128, F1: 65535}
                                        v18 := GenS4{F0: 42, F1: 0, F2: 98}
                                        var v19 uint = 65535
                                        _ = v17
                                        _ = v18
                                        _ = v19
                                    }
                                default:
                                    _ = len(v2)
                                }
                                _ = fi15
                            }
                            _ = fi13
                        }
                    }
                }
                if (67 != 122) {
                    v20 := GenS3{F0: 42, F1: -12}
                    if false {
                        if len(v6) >= 2 {
                            v6 = v6[1:len(v6)]
                        }
                        fmt.Println(v1, v4, v3)
                    } else {
                        v5 = v5
                    }
                    _ = v20
                }
                switch 4294967295 {
                case 175:
                    for fi21 := 0; fi21 < 8; fi21++ {
                        if (255 != 156) {
                            var v22 int16 = -1
                            var v23 uint = 255
                            _ = v22
                            _ = v23
                        } else {
                            v24 := GenS1{F0: 809, F1: 65535}
                            _ = v24
                        }
                        if true {
                            for fi25 := 0; fi25 < 15; fi25++ {
                                v3 = v3
                                var v26 bool = false
                                var v27 int64 = (v1 | (^(v1)))
                                _ = fi25
                                _ = v26
                                _ = v27
                            }
                        } else {
                            for fi28 := 0; fi28 < 18; fi28++ {
                                v4 = (v4 + "")
                                _ = fi28
                            }
                            fmt.Println(v3, v5, v1)
                        }
                        _ = fi21
                    }
                case 682:
                    v3 -= (v3 + (v3 & v3))
                }
            }
        }
        var v29 int16 = 0
        _ = v29
    }
    switch (v4 + (v4 + ("\n\t" + ((("foo bar" + "") + v4) + v4)))) {
    case "\n\t":
        for fi30 := 0; fi30 < 15; fi30++ {
            fmt.Println(v3, v1, v3)
            if (!(v1 > v1)) {
                if len(v2) > 0 {
                    v2[0] = "foo bar"
                }
                v31 := GenS2{F0: 2, F1: false, F2: (v3 + v3)}
                v32 := GenS3{F0: v3, F1: 2147483647}
                _ = v31
                _ = v32
            }
            _ = fi30
        }
    case "a":
        if ((false || false) && true) {
            v33 := GenS2{F0: 1, F1: (!false), F2: v3}
            switch v4 {
            case "\n\t":
                v34 := GenS2{F0: 896, F1: true, F2: (^(v3))}
                _ = v34
            case "":
                v1 = 9223372036854775807
            case "foo bar":
                v5 = v5
            default:
                v35 := GenS4{F0: 4294967295, F1: 255, F2: 83}
                _ = v35
            }
            switch "hello" {
            case "hello":
                var v36 uint = 948
                _ = v36
            }
            _ = v33
        } else {
            if false {
                if false {
                    for fi37 := 0; fi37 < 12; fi37++ {
                        var v38 int64 = v1
                        _ = fi37
                        _ = v38
                    }
                } else {
                    if (int16(v3) > 264) {
                        for fi39 := 0; fi39 < 5; fi39++ {
                            v4 = ("foo bar" + (("hello" + "a") + "foo bar"))
                            v40 := GenS1{F0: 987, F1: 362}
                            for fi41 := 0; fi41 < 11; fi41++ {
                                switch v4 {
                                case "\n\t":
                                    var v42 uint16 = 586
                                    _ = v42
                                case "foo bar":
                                    v3 = (v3 + (v3 / (((v3 & v3)) | 1)))
                                case "hello":
                                    if (((v4 + "") + "foo bar") > "hello") {
                                        v5 = v5
                                        var v43 int = fi39
                                        _ = v43
                                    } else {
                                        fmt.Println(fi39, v3, v5)
                                        var v44 int16 = 1000
                                        if true {
                                            v45 := GenS4{F0: 490, F1: 734, F2: 0}
                                            if (4294967295 < 892) {
                                                switch (fi39 - fi39) {
                                                case 42:
                                                    v4 = v4
                                                case 943:
                                                    fmt.Println(v1, v1, v5)
                                                case 128:
                                                    var v46 int64 = (-(v1))
                                                    _ = v46
                                                default:
                                                    var v47 int = (fi39 / (((fi39 >> ((327) & 63))) | 1))
                                                    _ = v47
                                                }
                                            } else {
                                                v48 := GenS4{F0: 2, F1: 0, F2: 16}
                                                _ = v48
                                            }
                                            var v49 int16 = v44
                                            _ = v45
                                            _ = v49
                                        }
                                        _ = v44
                                    }
                                default:
                                    v50 := GenS5{F0: 957, F1: v4, F2: v1}
                                    _ = v50
                                }
                                switch v4 {
                                case "foo bar":
                                    v51 := GenS6{F0: (^(v1)), F1: 226, F2: v3}
                                    _ = v51
                                case "\n\t":
                                    var v52 uint16 = 42
                                    _ = v52
                                case "a":
                                    if true {
                                        v53 := GenS7{F0: 42, F1: 42, F2: (^(v5))}
                                        _ = v53
                                    }
                                    fmt.Println(fi39, v4)
                                default:
                                    var v54 int64 = v1
                                    _ = v54
                                }
                                var v55 float32 = 3.14159
                                switch (^(fi39)) {
                                case 42:
                                    var v56 int16 = 518
                                    fmt.Println(v1, v1, v55)
                                    if true {
                                        var v57 int8 = -1
                                        _ = v57
                                    } else {
                                        v3 += 177
                                        if false {
                                            switch v4 {
                                            case "a":
                                                var v58 uint = 18446744073709551615
                                                _ = v58
                                            case "foo bar":
                                                var v59 uint64 = 4294967295
                                                _ = v59
                                            default:
                                                _ = len(v6)
                                            }
                                            v56 -= 42
                                            v55 += -2.25
                                        } else {
                                            switch (("" + ((v4 + "a") + v4)) + (("hello" + v4) + ("foo bar" + "\n\t"))) {
                                            case "a":
                                                var v60 float64 = (((0.0 - (1e10 * -2.25)) - 0.1) / (-2.25 * (float64(1e10) / (((0.1 + (0.0 / 1.0)) * (float64(v55) - 1e10)) + (-2.25 + 3.14159)))))
                                                _ = v60
                                            case "foo bar":
                                                v61 := GenS7{F0: (-(v1)), F1: 999, F2: 2147483647}
                                                _ = v61
                                            default:
                                                fmt.Println(v1, v5)
                                            }
                                        }
                                        v56 = v56
                                    }
                                    _ = v56
                                case 32767:
                                    v1 = (^(v1))
                                case 0:
                                    var v62 float64 = 1.5
                                    _ = v62
                                }
                                _ = fi41
                                _ = v55
                            }
                            _ = fi39
                            _ = v40
                        }
                    }
                    v4 = v4
                }
                v6 = append(v6, v4)
            }
        }
    case "hello":
        var v63 int = -2147483648
        _ = v63
    }
    var v64 float64 = 1e10
    for fi65 := 0; fi65 < 10; fi65++ {
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        for fi66 := 0; fi66 < 7; fi66++ {
            var v67 float32 = 0.1
            _ = fi66
            _ = v67
        }
        _ = fi65
    }
    switch ("hello" + (v4 + ("hello" + "\n\t"))) {
    case "a":
        if true {
            var v68 uint8 = 109
            _ = v68
        } else {
            v6 = append(v6, v4)
            v5 += -1
        }
    case "":
        v3 = v3
    case "foo bar":
        v69 := GenS8{F0: 32768, F1: v5}
        _ = v69
    default:
        var v70 uint32 = v3
        _ = v70
    }
    if false {
        var v71 int16 = 32767
        v72 := GenS9{F0: 228, F1: (^(v5))}
        fmt.Println(v64, v1)
        _ = v71
        _ = v72
    }
    for fi73 := 0; fi73 < 12; fi73++ {
        if true {
            for fi74 := 0; fi74 < 8; fi74++ {
                var v75 string = (v4 + (v4 + (v4 + "foo bar")))
                _ = fi74
                _ = v75
            }
            var v76 float64 = (0.0 + (v64 / v64))
            _ = v76
        } else {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            for fi77 := 0; fi77 < 19; fi77++ {
                var v78 float64 = v64
                v79 := GenS8{F0: (fi73 % ((-2147483648) | 1)), F1: v5}
                if len(v6) > 0 {
                    v6[0] = v4
                }
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                v1 += v1
                if len(v6) > 0 {
                    v6[0] = "foo bar"
                }
                fmt.Println(v78, fi73)
                _ = fi77
                _ = v78
                _ = v79
            }
        }
        v1 = v1
        _ = fi73
    }
    fmt.Println(v1)
    v80 := GenS4{F0: 18446744073709551615, F1: 394, F2: 157}
    for fi81 := 0; fi81 < 14; fi81++ {
        v82 := GenS7{F0: (-(v1)), F1: 0, F2: 1}
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        _ = fi81
        _ = v82
    }
    switch -2147483648 {
    case -1:
        v3 = 2147483647
    default:
        if len(v6) > 0 {
            v6[0] = ("hello" + v4)
        }
    }
    v3 -= (v3 % ((v3) | 1))
    fmt.Println(v3, v4, v4)
    for fi83 := 0; fi83 < 18; fi83++ {
        if false {
            var v84 int32 = (v5 - (^(v5)))
            var v85 uint16 = 925
            _ = v84
            _ = v85
        } else {
            v5 = v5
            var v86 bool = false
            _ = v86
        }
        for fi87 := 0; fi87 < 16; fi87++ {
            fmt.Println(v64, v1)
            var v88 uint16 = 65535
            v89 := GenS5{F0: (v1 | (v1 - v1)), F1: v4, F2: (^(v1))}
            _ = fi87
            _ = v88
            _ = v89
        }
        v5 = (v5 - (^(v5)))
        _ = fi83
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v64:", v64)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v64
    _ = v7
    _ = v8
    _ = v80
}
