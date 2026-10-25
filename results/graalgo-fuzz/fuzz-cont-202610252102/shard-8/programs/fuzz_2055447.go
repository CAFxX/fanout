package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint64; F1 int8 }
type GenS4 struct { F0 int64; F1 int8; F2 string }
type GenS5 struct { F0 int16; F1 int8; F2 uint16 }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 int64; F1 int8; F2 int }
type GenS8 struct { F0 int16; F1 uint }
type GenS9 struct { F0 bool; F1 int64 }
type GenS10 struct { F0 uint64 }
type GenS11 struct { F0 int32; F1 uint32; F2 int64 }
type GenS12 struct { F0 bool; F1 uint32; F2 int }
type GenS13 struct { F0 int16; F1 uint32 }
type GenS14 struct { F0 int; F1 int; F2 int8 }

func main() {
    var v1 uint16 = 0
    var v2 int = -2147483648
    v3 := []int{}
    v4 := make(map[uint32]string)
    v5 := []uint32{2147483647, 42, 2147483647}
    v1 = v1
    v1 = (v1 >> ((0) & 63))
    if true {
        for fi6 := 0; fi6 < 5; fi6++ {
            v2 = fi6
            var v7 int16 = 1
            v8 := GenS1{F0: v7, F1: 1}
            _ = fi6
            _ = v7
            _ = v8
        }
        var v9 int32 = -1
        _ = v9
    }
    if ((v2 % ((v2) | 1)) >= 160) {
        var v10 uint16 = (v1 ^ 928)
        if ("foo bar" <= "") {
            switch -466 {
            case 32768:
                v2 = (v2 - v2)
            case -969:
                fmt.Println(v2)
            }
        } else {
            var v11 uint16 = (v1 | v10)
            v11 = v10
            if (-128 >= 100) {
                for fi12 := 0; fi12 < 12; fi12++ {
                    if (1e10 == 0.0) {
                        _ = len(v5)
                        for fi13 := 0; fi13 < 8; fi13++ {
                            if true {
                                v2 -= (v2 - -610)
                                if (!true) {
                                    v2 = 2147483647
                                    var v14 uint16 = v11
                                    _ = v14
                                }
                            } else {
                                var v15 float32 = 3.14159
                                if len(v5) > 0 {
                                    v5[0] = 0
                                }
                                v11 = v10
                                _ = v15
                            }
                            v16 := GenS2{F0: (42 == 42)}
                            var v17 uint = 0
                            _ = fi13
                            _ = v16
                            _ = v17
                        }
                        if true {
                            v11 = (v11 / ((v11) | 1))
                            switch (v2 * (-(fi12))) {
                            case -700:
                                if (-2147483648 != -2147483648) {
                                    fmt.Println(v1, v2, fi12)
                                    fmt.Println(v10)
                                    fmt.Println(v11, v11, fi12)
                                } else {
                                    fmt.Println(v2, v1)
                                }
                            default:
                                v18 := GenS2{F0: (3.14159 <= 1.5)}
                                _ = v18
                            }
                            for fi19 := 0; fi19 < 8; fi19++ {
                                if false {
                                    v20 := GenS3{F0: 42, F1: 1}
                                    _ = v20
                                } else {
                                    v21 := GenS4{F0: 1, F1: -128, F2: ""}
                                    var v22 uint8 = 2
                                    if false {
                                        if true {
                                            v4[0] = "foo bar"
                                            v23 := GenS3{F0: 2, F1: 127}
                                            if true {
                                                v24 := GenS1{F0: 0, F1: (v10 + v10)}
                                                if (!(true || false)) {
                                                    switch (-(fi19)) {
                                                    case 551:
                                                        if (fi12 != fi19) {
                                                            v25 := GenS5{F0: 42, F1: int8(v10), F2: (v11 % (((v10 ^ (v11 >> ((929) & 63)))) | 1))}
                                                            var v26 uint32 = 1
                                                            _ = v25
                                                            _ = v26
                                                        } else {
                                                            v27 := GenS6{F0: 1}
                                                            _ = v27
                                                        }
                                                    case 4294967295:
                                                        var v28 float32 = (0.1 + (-2.25 * 1e10))
                                                        _ = v28
                                                    }
                                                    v29 := GenS7{F0: -9223372036854775808, F1: 122, F2: v2}
                                                    _ = v29
                                                }
                                                _ = v24
                                            }
                                            _ = v23
                                        } else {
                                            switch "\n\t" {
                                            case "hello":
                                                switch "foo bar" {
                                                case "a":
                                                    var v30 uint8 = 0
                                                    _ = v30
                                                case "":
                                                    var v31 uint32 = 0
                                                    _ = v31
                                                case "hello":
                                                    v5 = append(v5, 915)
                                                }
                                            default:
                                                switch "foo bar" {
                                                case "foo bar":
                                                    v32 := GenS8{F0: -668, F1: 608}
                                                    _ = v32
                                                }
                                            }
                                            v33 := GenS2{F0: (!true)}
                                            v34 := GenS9{F0: (("a" + "foo bar") != ("\n\t" + ("hello" + ("\n\t" + ("a" + (("foo bar" + "foo bar") + "foo bar")))))), F1: 1}
                                            fmt.Println(fi19, v10, v22)
                                            _ = v33
                                            _ = v34
                                        }
                                        v1 = (v1 & v1)
                                        fmt.Println(v10, v11)
                                    }
                                    _ = v21
                                    _ = v22
                                }
                                v1 -= v11
                                if true {
                                    v35 := GenS10{F0: 246}
                                    _ = v35
                                }
                                _ = fi19
                            }
                        }
                    } else {
                        v2 = (v2 & v2)
                        v36 := GenS10{F0: 4294967295}
                        _ = v36
                    }
                    var v37 uint16 = (^(v1))
                    _ = fi12
                    _ = v37
                }
                if (255 >= 18446744073709551615) {
                    switch (v2 * (v2 % ((v2) | 1))) {
                    case 32767:
                        if true {
                            switch 127 {
                            case -792:
                                var v38 string = "hello"
                                _ = v38
                            case 42:
                                _ = len(v4)
                            }
                        } else {
                            if false {
                                v10 = (v10 - v1)
                            } else {
                                v39 := GenS11{F0: 42, F1: 1, F2: 0}
                                var v40 int8 = -128
                                _ = v39
                                _ = v40
                            }
                        }
                    case -2147483648:
                        fmt.Println(v10, v1, v2)
                    case 1:
                        for fi41 := 0; fi41 < 5; fi41++ {
                            var v42 uint32 = 2147483647
                            var v43 uint16 = 805
                            _ = len(v5)
                            _ = fi41
                            _ = v42
                            _ = v43
                        }
                    default:
                        v44 := GenS2{F0: (0.0 > ((0.0 - ((float64((1.5 / -2.25)) - 0.0) * 1.5)) + float64(1e10)))}
                        _ = v44
                    }
                    var v45 int16 = 42
                    _ = v45
                } else {
                    v46 := v4[2147483648]
                    v46 = ("" + "hello")
                    _ = v46
                }
            } else {
                if true {
                    if (true || false) {
                        v47 := GenS5{F0: -340, F1: -123, F2: v1}
                        _ = v47
                    } else {
                        v48 := GenS4{F0: -9223372036854775808, F1: int8(v10), F2: "\n\t"}
                        fmt.Println(v1)
                        for fi49 := 0; fi49 < 20; fi49++ {
                            v1 = v11
                            _ = fi49
                        }
                        _ = v48
                    }
                    var v50 int16 = 1000
                    v51 := GenS6{F0: 672}
                    _ = v50
                    _ = v51
                }
                switch (v2 + v2) {
                case 0:
                    var v52 uint32 = 2147483647
                    _ = v52
                case 521:
                    var v53 bool = false
                    _ = v53
                case -575:
                    var v54 uint8 = 78
                    _ = v54
                }
            }
            _ = v11
        }
        for fi55 := 0; fi55 < 4; fi55++ {
            v2 -= 32768
            v2 += 32768
            if (true && (144 <= 42)) {
                fmt.Println(v10, v2, fi55)
            } else {
                v1 = 255
            }
            _ = fi55
        }
        _ = v10
    } else {
        v1 -= (^(v1))
    }
    v2 -= v2
    v1 += v1
    v2 += v2
    v56 := GenS7{F0: 1, F1: 69, F2: (^(v2))}
    v5 = append(v5, 42)
    var v57 uint32 = 470
    fmt.Println(v57)
    if false {
        v58 := GenS5{F0: 1, F1: 2, F2: v1}
        for fi59 := 0; fi59 < 13; fi59++ {
            v60 := GenS7{F0: 0, F1: 1, F2: v2}
            v61 := GenS12{F0: (!false), F1: v57, F2: 32767}
            var v62 int16 = 0
            _ = fi59
            _ = v60
            _ = v61
            _ = v62
        }
        for fi63 := 0; fi63 < 17; fi63++ {
            var v64 float32 = -2.25
            v1 -= v1
            for fi65 := 0; fi65 < 7; fi65++ {
                switch "\n\t" {
                case "hello":
                    if false {
                        fmt.Println(fi63, v1)
                        _ = len(v4)
                    }
                case "foo bar":
                    v4[(v57 + (v57 ^ v57))] = ("a" + "a")
                default:
                    fmt.Println(v64, v57)
                }
                _ = fi65
            }
            _ = fi63
            _ = v64
        }
        _ = v58
    } else {
        var v66 float64 = 1e10
        if (952 >= int16(v1)) {
            for fi67 := 0; fi67 < 12; fi67++ {
                if (v1 > v1) {
                    var v68 float32 = ((((3.14159 - (0.1 - (0.1 - 0.1))) / (-2.25 / float32((v66 * (0.1 * (v66 / (-2.25 + ((v66 * 0.1) / v66)))))))) / 1.5) - 1e10)
                    v69 := GenS13{F0: -32768, F1: v57}
                    _ = v68
                    _ = v69
                }
                v66 = v66
                _ = fi67
            }
        } else {
            v70 := GenS14{F0: v2, F1: int(v1), F2: 1}
            v5 = append(v5, v57)
            for fi71 := 0; fi71 < 20; fi71++ {
                v72 := v4[(v57 | 627)]
                var v73 float64 = 1e10
                _ = fi71
                _ = v72
                _ = v73
            }
            _ = v70
        }
        if (("\n\t" + "foo bar") >= "hello") {
            v1 += 32767
            fmt.Println(v57, v2)
        } else {
            var v74 int32 = 42
            v5 = append(v5, (v57 << ((1) & 63)))
            fmt.Println(v2, v57, v57)
            _ = v74
        }
        _ = v66
    }
    switch -897 {
    case -9223372036854775808:
        v57 = 441
    default:
        fmt.Println(v57, v2, v1)
    }
    var v75 string = "\n\t"
    fmt.Println(v75)
    fmt.Println(v1, v2, v57)
    if (v1 == (v1 % ((1) | 1))) {
        v76 := GenS6{F0: 9223372036854775807}
        _ = len(v3)
        var v77 int16 = int16(v57)
        _ = v76
        _ = v77
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v57:", v57)
    fmt.Println("v75:", v75)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v56
    _ = v57
    _ = v75
}
