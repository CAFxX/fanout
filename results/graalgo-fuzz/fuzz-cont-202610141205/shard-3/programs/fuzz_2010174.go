package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 bool }
type GenS2 struct { F0 uint16; F1 string }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 bool; F1 int32; F2 uint32 }
type GenS6 struct { F0 uint64; F1 string; F2 int }

func main() {
    var v1 uint16 = 490
    var v2 float32 = ((float32(0.0) * 0.1) - 1e10)
    var v3 uint32 = 1
    v4 := []int{1, 32768, 176, 2147483647}
    fmt.Println(v3)
    fmt.Println(v2, v1, v3)
    var v5 int8 = 127
    for fi6 := 0; fi6 < 1; fi6++ {
        fmt.Println(v5, fi6)
        if (!((-736 != 0) || false)) {
            var v7 uint32 = (^(v3))
            v1 -= (v1 - (v1 / ((v1) | 1)))
            _ = v7
        }
        _ = fi6
    }
    for fi8 := 0; fi8 < 16; fi8++ {
        if false {
            var v9 uint16 = v1
            _ = v9
        }
        v4 = append(v4, (fi8 ^ -625))
        _ = fi8
    }
    var v10 int = 128
    if (false || false) {
        var v11 int = (v10 ^ (-(v10)))
        v12 := GenS1{F0: (42 <= 0), F1: (343 > 849)}
        for fi13 := 0; fi13 < 13; fi13++ {
            for fi14 := 0; fi14 < 18; fi14++ {
                if ((v10 << ((596) & 63)) <= 42) {
                    if (true || false) {
                        var v15 bool = true
                        var v16 uint8 = 255
                        var v17 uint = 1
                        _ = v15
                        _ = v16
                        _ = v17
                    }
                    if false {
                        v1 = v1
                        v5 = (v5 & v5)
                        v3 = (v3 << ((65535) & 63))
                        v4 = append(v4, v11)
                    } else {
                        if false {
                            switch (("" + (("\n\t" + "foo bar") + ("hello" + ""))) + (("foo bar" + "a") + "foo bar")) {
                            case "a":
                                if false {
                                    v1 -= v1
                                    _ = len(v4)
                                    fmt.Println(v5, fi13)
                                    if false {
                                        var v18 string = ""
                                        if len(v4) > 0 {
                                            v4[0] = fi14
                                        }
                                        _ = v18
                                    }
                                } else {
                                    v5 = 127
                                }
                            case "\n\t":
                                v5 = v5
                            case "foo bar":
                                var v19 uint32 = (v3 * (v3 >> ((4294967295) & 63)))
                                _ = v19
                            }
                            v10 += -952
                        } else {
                            var v20 uint = 42
                            switch (v10 * v10) {
                            case 4294967295:
                                if (!(!false)) {
                                    if (!true) {
                                        v21 := GenS1{F0: (!false), F1: false}
                                        v1 -= (v1 * v1)
                                        _ = v21
                                    } else {
                                        v22 := GenS2{F0: 559, F1: ("\n\t" + (("\n\t" + "hello") + (("" + "") + "foo bar")))}
                                        _ = v22
                                    }
                                    v5 += v5
                                    v23 := GenS1{F0: false, F1: (1 == 1)}
                                    _ = v23
                                } else {
                                    if true {
                                        if false {
                                            if len(v4) > 0 {
                                                v4[0] = (fi14 & (fi14 * v11))
                                            }
                                            var v24 uint16 = (v1 % ((v1) | 1))
                                            _ = v24
                                        }
                                        if len(v4) > 0 {
                                            v4[0] = -107
                                        }
                                        v25 := GenS1{F0: (v2 >= (0.0 + v2)), F1: true}
                                        if true {
                                            v26 := GenS3{F0: (v1 + v1)}
                                            fmt.Println(v3, v5)
                                            switch ("foo bar" + ("foo bar" + "")) {
                                            case "":
                                                v11 = 443
                                                var v27 float32 = v2
                                                if len(v4) > 0 {
                                                    v4[0] = v10
                                                }
                                                _ = v27
                                            case "hello":
                                                var v28 uint32 = uint32(v20)
                                                _ = v28
                                            case "a":
                                                v4 = append(v4, (v10 / ((-372) | 1)))
                                            }
                                            _ = v26
                                        } else {
                                            if false {
                                                v29 := GenS2{F0: (v1 << (((v20 & (v20 & (v20 & v20)))) & 63)), F1: "foo bar"}
                                                v30 := GenS2{F0: (v1 ^ 56), F1: ("" + "")}
                                                if false {
                                                    v31 := GenS4{F0: (v3 / ((v3) | 1))}
                                                    switch "hello" {
                                                    case "":
                                                        var v32 int8 = v5
                                                        _ = v32
                                                    case "\n\t":
                                                        v33 := GenS5{F0: true, F1: -2147483648, F2: v3}
                                                        _ = v33
                                                    case "foo bar":
                                                        var v34 uint = 2
                                                        _ = v34
                                                    default:
                                                        if (!false) {
                                                            v35 := GenS1{F0: true, F1: (!(true && true))}
                                                            v36 := GenS4{F0: (^(v3))}
                                                            var v37 int32 = 2147483647
                                                            v38 := GenS1{F0: (v3 == uint32(v20)), F1: (((v37 > 320) && true) && false)}
                                                            var v39 uint8 = 81
                                                            _ = v35
                                                            _ = v36
                                                            _ = v37
                                                            _ = v38
                                                            _ = v39
                                                        } else {
                                                            if len(v4) > 0 {
                                                                v4[0] = fi13
                                                            }
                                                            v2 = 1e10
                                                            v10 = (-(fi14))
                                                            var v40 uint16 = uint16(v20)
                                                            _ = v40
                                                        }
                                                        fmt.Println(v5, v3, v20)
                                                        var v41 int64 = 1
                                                        fmt.Println(fi14, v41)
                                                        var v42 uint16 = v1
                                                        _ = v41
                                                        _ = v42
                                                    }
                                                    fmt.Println(v20)
                                                    _ = v31
                                                }
                                                _ = v29
                                                _ = v30
                                            } else {
                                                if len(v4) > 0 {
                                                    v4[0] = fi13
                                                }
                                                fmt.Println(v3, fi14, fi14)
                                            }
                                            v5 -= v5
                                            v11 = int(v1)
                                        }
                                        _ = v25
                                    } else {
                                        if (!true) {
                                            if ((v2 + (v2 + 1.5)) != (v2 / -2.25)) {
                                                if (v2 <= v2) {
                                                    var v43 bool = (!(v5 != (v5 * 127)))
                                                    v44 := GenS5{F0: false, F1: 0, F2: 530}
                                                    _ = v43
                                                    _ = v44
                                                } else {
                                                    switch "foo bar" {
                                                    case "a":
                                                        if false {
                                                            var v45 int16 = 0
                                                            v1 = 0
                                                            _ = v45
                                                        }
                                                    case "foo bar":
                                                        if (!(3.14159 != float64(v2))) {
                                                            if (66 < 4294967295) {
                                                                v20 = v20
                                                                v46 := GenS1{F0: false, F1: true}
                                                                var v47 float32 = (float32(1e10) + 1.5)
                                                                _ = v46
                                                                _ = v47
                                                            } else {
                                                                if true {
                                                                    v10 = v11
                                                                    v10 -= (v11 * 69)
                                                                    fmt.Println(v3, v2, fi13)
                                                                    v2 -= (0.0 + v2)
                                                                } else {
                                                                    var v48 bool = true
                                                                    _ = v48
                                                                }
                                                                var v49 float32 = 3.14159
                                                                fmt.Println(v20, fi14)
                                                                var v50 string = "hello"
                                                                _ = v49
                                                                _ = v50
                                                            }
                                                            v10 -= (fi14 | fi14)
                                                            var v51 string = "a"
                                                            _ = v51
                                                        } else {
                                                            fmt.Println(v1)
                                                            v10 = 2147483647
                                                            if len(v4) >= 2 {
                                                                v4 = v4[1:len(v4)]
                                                            }
                                                        }
                                                        v52 := GenS2{F0: 0, F1: ""}
                                                        var v53 uint64 = 18446744073709551615
                                                        if len(v4) >= 2 {
                                                            v4 = v4[1:len(v4)]
                                                        }
                                                        if len(v4) >= 2 {
                                                            v4 = v4[1:len(v4)]
                                                        }
                                                        var v54 float32 = (v2 / v2)
                                                        fmt.Println(fi13, v1, v20)
                                                        _ = v52
                                                        _ = v53
                                                        _ = v54
                                                    default:
                                                        var v55 uint = (v20 - uint(v3))
                                                        _ = v55
                                                    }
                                                }
                                                v11 += fi13
                                                if len(v4) > 0 {
                                                    v4[0] = 32767
                                                }
                                            }
                                        }
                                    }
                                    v10 = 9223372036854775807
                                    var v56 uint32 = (v3 | uint32(v1))
                                    _ = v56
                                }
                            default:
                                if ((174 >= 562) || (18446744073709551615 >= 1)) {
                                    if (fi13 <= 577) {
                                        if true {
                                            if len(v4) > 0 {
                                                v4[0] = v10
                                            }
                                            var v57 uint8 = 235
                                            v2 = (0.0 * (3.14159 / (float32(-2.25) + (v2 * v2))))
                                            v58 := GenS5{F0: (((v2 * v2) / (1e10 - 1.5)) >= 0.1), F1: -372, F2: 550}
                                            if len(v4) > 0 {
                                                v4[0] = (v11 ^ 42)
                                            }
                                            if true {
                                                if ((0.0 * 1.5) <= ((1.5 - (1e10 - 3.14159)) + float64(0.0))) {
                                                    v59 := GenS4{F0: v3}
                                                    v4 = append(v4, -9223372036854775808)
                                                    _ = v59
                                                }
                                                switch ("\n\t" + "hello") {
                                                case "foo bar":
                                                    fmt.Println(v11, fi14, v5)
                                                default:
                                                    v20 = v20
                                                }
                                                var v60 string = ""
                                                _ = v60
                                            } else {
                                                v2 += float32((0.1 + (1e10 - 0.1)))
                                                v61 := GenS6{F0: 1, F1: "a", F2: (v11 | int(v57))}
                                                if false {
                                                    if true {
                                                        v3 -= v3
                                                        var v62 int64 = -64
                                                        fmt.Println(v1, v11)
                                                        v4 = append(v4, (fi14 ^ v10))
                                                        if false {
                                                            fmt.Println(v10)
                                                        }
                                                        _ = v62
                                                    } else {
                                                        var v63 int32 = 1
                                                        _ = v63
                                                    }
                                                    fmt.Println(v2)
                                                    if true {
                                                        v20 -= 669
                                                        fmt.Println(v11)
                                                    }
                                                }
                                                _ = v61
                                            }
                                            _ = v57
                                            _ = v58
                                        }
                                        var v64 uint8 = 42
                                        switch "a" {
                                        case "a":
                                            var v65 int16 = 1000
                                            _ = v65
                                        case "foo bar":
                                            fmt.Println(v5)
                                        }
                                        _ = v64
                                    } else {
                                        v10 = fi13
                                        var v66 uint32 = 187
                                        v67 := GenS4{F0: v66}
                                        _ = v66
                                        _ = v67
                                    }
                                } else {
                                    fmt.Println(v11, v2)
                                    switch fi13 {
                                    case 127:
                                        var v68 bool = true
                                        _ = v68
                                    case -1:
                                        v10 = int(v1)
                                    }
                                    switch "foo bar" {
                                    case "foo bar":
                                        var v69 uint32 = (v3 & (v3 | (v3 + (^(v3)))))
                                        _ = v69
                                    }
                                }
                            }
                            v11 += fi14
                            _ = v20
                        }
                    }
                }
                _ = fi14
            }
            _ = fi13
        }
        _ = v11
        _ = v12
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    var v70 int16 = 0
    var v71 int32 = -266
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v70:", v70)
    fmt.Println("v71:", v71)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v70
    _ = v71
}
