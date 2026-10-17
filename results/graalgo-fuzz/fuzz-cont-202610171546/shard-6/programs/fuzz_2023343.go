package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int32; F1 bool; F2 bool }
type GenS5 struct { F0 int32; F1 int32 }
type GenS6 struct { F0 uint16; F1 int16 }
type GenS7 struct { F0 string }
type GenS8 struct { F0 bool; F1 int64 }
type GenS9 struct { F0 bool; F1 string }

func main() {
    v1 := []int{}
    v2 := []string{"foo bar", "foo bar"}
    var v3 float32 = (float32(((-2.25 / 1e10) / 1.0)) + 3.14159)
    v4 := []uint32{103, 840, 42}
    var v5 int = 0
    switch "foo bar" {
    case "":
        v5 = (v5 + 4294967295)
    default:
        if (!true) {
            var v6 int16 = -1
            if true {
                switch 32767 {
                case -601:
                    var v7 int16 = (v6 ^ v6)
                    _ = v7
                default:
                    switch (v5 - v5) {
                    case -795:
                        v8 := GenS1{F0: 2147483648}
                        _ = v8
                    case 42:
                        v9 := GenS2{F0: 2147483647}
                        _ = v9
                    case 730:
                        _ = len(v2)
                    default:
                        for fi10 := 0; fi10 < 13; fi10++ {
                            var v11 int64 = 869
                            _ = fi10
                            _ = v11
                        }
                    }
                }
                var v12 uint = 18446744073709551615
                for fi13 := 0; fi13 < 10; fi13++ {
                    for fi14 := 0; fi14 < 13; fi14++ {
                        var v15 uint = 4294967295
                        v15 -= v12
                        fmt.Println(v15)
                        var v16 uint16 = 1
                        v12 -= (v15 >> ((v15) & 63))
                        _ = fi14
                        _ = v15
                        _ = v16
                    }
                    if (v12 <= v12) {
                        var v17 string = ""
                        if true {
                            v3 = v3
                            v6 += (v6 << (((v12 * v12)) & 63))
                        } else {
                            var v18 int16 = (-(v6))
                            fmt.Println(v3, v5, fi13)
                            _ = v18
                        }
                        var v19 float64 = 1.5
                        _ = v17
                        _ = v19
                    } else {
                        var v20 int32 = -1
                        v21 := GenS2{F0: v20}
                        _ = v20
                        _ = v21
                    }
                    _ = fi13
                }
                _ = v12
            }
            _ = v6
        } else {
            for fi22 := 0; fi22 < 18; fi22++ {
                var v23 int64 = 342
                _ = fi22
                _ = v23
            }
            if (!(127 == 127)) {
                if false {
                    var v24 int = v5
                    v24 += v24
                    _ = v24
                }
                if len(v1) > 0 {
                    v1[0] = (v5 >> ((65535) & 63))
                }
                for fi25 := 0; fi25 < 15; fi25++ {
                    if (0 > -562) {
                        var v26 bool = (1e10 != 0.0)
                        _ = v26
                    }
                    _ = fi25
                }
            } else {
                var v27 int = 579
                if (1 > 42) {
                    v5 -= (^(v27))
                    var v28 bool = true
                    _ = v28
                }
                _ = v27
            }
        }
    }
    for fi29 := 0; fi29 < 1; fi29++ {
        var v30 int = (fi29 ^ -800)
        var v31 int8 = 127
        fmt.Println(v3)
        _ = fi29
        _ = v30
        _ = v31
    }
    var v32 int16 = 1000
    fmt.Println(v32, v5)
    if false {
        var v33 uint = 255
        if true {
            v34 := GenS2{F0: -2147483648}
            if false {
                for fi35 := 0; fi35 < 5; fi35++ {
                    if false {
                        var v36 uint = 1
                        for fi37 := 0; fi37 < 7; fi37++ {
                            v38 := GenS3{F0: -1}
                            if (false || (v33 == v33)) {
                                switch "foo bar" {
                                case "a":
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                    switch 2147483647 {
                                    case 0:
                                        fmt.Println(fi35)
                                        if (213 == 12) {
                                            var v39 int16 = v32
                                            _ = v39
                                        } else {
                                            switch -616 {
                                            case -9223372036854775808:
                                                switch (fi37 >> ((v36) & 63)) {
                                                case -120:
                                                    var v40 string = ("foo bar" + "")
                                                    _ = v40
                                                case 140:
                                                    v41 := GenS1{F0: 2147483647}
                                                    if (!false) {
                                                        v42 := GenS1{F0: 927}
                                                        _ = v42
                                                    }
                                                    var v43 bool = true
                                                    _ = v41
                                                    _ = v43
                                                case 4294967295:
                                                    fmt.Println(fi37, fi37, v5)
                                                default:
                                                    switch ("" + "foo bar") {
                                                    case "hello":
                                                        v44 := GenS4{F0: 1, F1: ((("" + (("a" + "foo bar") + ("\n\t" + ""))) + "a") != "hello"), F2: (1e10 < -2.25)}
                                                        _ = v44
                                                    case "a":
                                                        if (594 > 255) {
                                                            v36 -= (^(v33))
                                                            v45 := GenS5{F0: 2147483647, F1: -1}
                                                            if (!(false || true)) {
                                                                var v46 float64 = float64(v3)
                                                                _ = v46
                                                            } else {
                                                                fmt.Println(v3)
                                                            }
                                                            _ = v45
                                                        }
                                                    case "foo bar":
                                                        var v47 int64 = -633
                                                        _ = v47
                                                    default:
                                                        var v48 string = "hello"
                                                        _ = v48
                                                    }
                                                    v5 = fi35
                                                    var v49 int = (fi37 | 42)
                                                    _ = v49
                                                }
                                            case 462:
                                                fmt.Println(fi35)
                                            case 128:
                                                fmt.Println(v33, fi37)
                                            default:
                                                v32 = (v32 * v32)
                                            }
                                            if false {
                                                _ = len(v2)
                                                var v50 float64 = float64((v3 / 1.0))
                                                v5 -= (fi37 / (((fi35 - v5)) | 1))
                                                _ = v50
                                            }
                                        }
                                    case -2147483648:
                                        v51 := GenS3{F0: 127}
                                        _ = v51
                                    default:
                                        v3 += 0.1
                                    }
                                    v52 := GenS5{F0: -1, F1: 0}
                                    _ = v52
                                case "\n\t":
                                    v53 := GenS4{F0: 0, F1: false, F2: false}
                                    _ = v53
                                case "":
                                    switch v5 {
                                    case -132:
                                        v32 = -32768
                                    case 42:
                                        var v54 string = "a"
                                        _ = v54
                                    }
                                }
                                fmt.Println(fi35, fi35, v5)
                                v33 += v36
                            } else {
                                var v55 uint16 = 65535
                                _ = v55
                            }
                            if (!(((float32(-2.25) * float32((3.14159 - float64((float32(0.1) * (v3 * v3)))))) - (v3 * v3)) >= 3.14159)) {
                                v36 = v36
                                if (-2147483648 <= 2147483647) {
                                    v33 = 65535
                                    v56 := GenS2{F0: -930}
                                    switch (fi35 * fi37) {
                                    case 127:
                                        v1 = append(v1, -9223372036854775808)
                                    case -2147483648:
                                        switch fi35 {
                                        case 9223372036854775807:
                                            v3 = (v3 + v3)
                                        }
                                    case 32767:
                                        var v57 float64 = 3.14159
                                        switch (fi35 << ((v36) & 63)) {
                                        case -157:
                                            if true {
                                                v32 = v32
                                            }
                                        case 32767:
                                            var v58 bool = (!false)
                                            _ = v58
                                        default:
                                            v2 = append(v2, "hello")
                                        }
                                        if len(v1) > 0 {
                                            v1[0] = fi37
                                        }
                                        if (1 >= 18446744073709551615) {
                                            v57 += (v57 / float64((v3 * v3)))
                                        }
                                        var v59 uint32 = 2147483647
                                        _ = v57
                                        _ = v59
                                    }
                                    v5 = fi37
                                    if len(v2) > 0 {
                                        v2[0] = "hello"
                                    }
                                    _ = v56
                                } else {
                                    var v60 int16 = (^(v32))
                                    var v61 int32 = 42
                                    _ = v60
                                    _ = v61
                                }
                                if (0 == int16(v33)) {
                                    if (!false) {
                                        var v62 bool = (0 == 4294967295)
                                        _ = v62
                                    } else {
                                        fmt.Println(fi37, v5)
                                    }
                                } else {
                                    switch ("a" + ("" + "")) {
                                    case "\n\t":
                                        v33 -= (^(v33))
                                    case "a":
                                        switch (fi37 * v5) {
                                        case 42:
                                            fmt.Println(v36, v36, v33)
                                        case 127:
                                            fmt.Println(v32, v33)
                                        }
                                    }
                                    var v63 uint = v36
                                    fmt.Println(v36)
                                    _ = v63
                                }
                            } else {
                                var v64 uint = (v36 + v33)
                                _ = v64
                            }
                            _ = fi37
                            _ = v38
                        }
                        _ = v36
                    }
                    v32 = (v32 * 42)
                    _ = fi35
                }
                var v65 int16 = 506
                _ = v65
            }
            _ = v34
        } else {
            if false {
                for fi66 := 0; fi66 < 2; fi66++ {
                    v32 -= v32
                    for fi67 := 0; fi67 < 1; fi67++ {
                        v68 := GenS6{F0: 2, F1: -1}
                        var v69 int8 = 1
                        v70 := GenS7{F0: "\n\t"}
                        _ = fi67
                        _ = v68
                        _ = v69
                        _ = v70
                    }
                    v3 -= 1.5
                    _ = fi66
                }
                _ = len(v2)
            }
            var v71 int64 = -523
            v32 = (v32 + -87)
            _ = v71
        }
        switch "hello" {
        case "foo bar":
            for fi72 := 0; fi72 < 8; fi72++ {
                switch "hello" {
                case "\n\t":
                    fmt.Println(v3)
                case "foo bar":
                    if false {
                        fmt.Println(v5, v33)
                        v33 = (v33 >> ((1) & 63))
                    } else {
                        _ = len(v4)
                        var v73 uint = 370
                        _ = v73
                    }
                }
                _ = fi72
            }
        }
        _ = v33
    } else {
        for fi74 := 0; fi74 < 17; fi74++ {
            if (!(v32 > (^(v32)))) {
                if true {
                    for fi75 := 0; fi75 < 13; fi75++ {
                        fmt.Println(v3, v5)
                        fmt.Println(fi75)
                        _ = fi75
                    }
                    switch (v5 - (v5 | (^(v5)))) {
                    case 42:
                        fmt.Println(fi74)
                    }
                } else {
                    v32 = (v32 + (v32 % ((238) | 1)))
                    fmt.Println(v3, v32, v32)
                    v32 = v32
                }
                v76 := GenS8{F0: false, F1: 9223372036854775807}
                _ = v76
            }
            v77 := GenS8{F0: (false && true), F1: -335}
            _ = fi74
            _ = v77
        }
        for fi78 := 0; fi78 < 11; fi78++ {
            v32 -= v32
            _ = len(v1)
            _ = fi78
        }
        var v79 uint64 = 18446744073709551615
        _ = v79
    }
    var v80 bool = (!true)
    v81 := GenS9{F0: (!v80), F1: ""}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v80:", v80)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v4
    _ = v5
    _ = v80
    _ = v81
}
