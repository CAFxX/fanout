package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int8; F2 int64 }
type GenS2 struct { F0 int; F1 uint }
type GenS3 struct { F0 uint16; F1 uint; F2 int64 }
type GenS4 struct { F0 int32; F1 uint32; F2 uint8 }
type GenS5 struct { F0 string; F1 uint64; F2 int8 }
type GenS6 struct { F0 string }
type GenS7 struct { F0 int64 }
type GenS8 struct { F0 int8; F1 uint32; F2 int8 }

func main() {
    var v1 bool = (1.5 != (1e10 + 1.5))
    var v2 bool = (false || v1)
    v3 := []int{313}
    var v4 int32 = 2147483647
    for fi5 := 0; fi5 < 6; fi5++ {
        switch (fi5 << ((65535) & 63)) {
        case 4294967295:
            if true {
                var v6 int64 = 42
                if (0 > 842) {
                    if len(v3) > 0 {
                        v3[0] = (fi5 ^ fi5)
                    }
                }
                if (1 < fi5) {
                    if (v6 <= v6) {
                        var v7 int64 = (-(v6))
                        for fi8 := 0; fi8 < 17; fi8++ {
                            v7 = v6
                            var v9 int8 = 0
                            _ = fi8
                            _ = v9
                        }
                        _ = v7
                    }
                    fmt.Println(fi5)
                } else {
                    switch "foo bar" {
                    case "foo bar":
                        for fi10 := 0; fi10 < 6; fi10++ {
                            fmt.Println(v2, v2, v1)
                            v1 = false
                            if (v2 && (930 < (-(fi10)))) {
                                v11 := GenS1{F0: 2, F1: -88, F2: (v6 >> ((65535) & 63))}
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                _ = len(v3)
                                _ = v11
                            }
                            _ = fi10
                        }
                    case "\n\t":
                        var v12 int32 = (v4 / ((v4) | 1))
                        _ = v12
                    case "":
                        for fi13 := 0; fi13 < 8; fi13++ {
                            v4 += 998
                            if (512 >= 18446744073709551615) {
                                v1 = v1
                                var v14 float64 = 3.14159
                                _ = v14
                            } else {
                                fmt.Println(fi5)
                                var v15 int = -2147483648
                                _ = v15
                            }
                            v16 := GenS2{F0: fi5, F1: 2}
                            _ = fi13
                            _ = v16
                        }
                    }
                    v4 = v4
                    v1 = v1
                }
                _ = v6
            } else {
                v1 = (-1 <= 9223372036854775807)
                _ = len(v3)
            }
        }
        _ = fi5
    }
    var v17 int16 = 32767
    if v2 {
        if (v2 && v1) {
            var v18 float64 = 1e10
            var v19 string = "foo bar"
            v2 = v1
            _ = v18
            _ = v19
        } else {
            v2 = v2
        }
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        if (42 > -126) {
            fmt.Println(v4, v4)
        }
    }
    fmt.Println(v2, v2, v17)
    for fi20 := 0; fi20 < 15; fi20++ {
        if (-2.25 >= ((0.1 / 0.1) / 3.14159)) {
            v4 = (v4 - (v4 >> ((18446744073709551615) & 63)))
        } else {
            for fi21 := 0; fi21 < 12; fi21++ {
                if v2 {
                    var v22 int32 = -1
                    v2 = v2
                    fmt.Println(v2)
                    _ = v22
                } else {
                    var v23 string = ("\n\t" + "hello")
                    v24 := GenS1{F0: 42, F1: int8(v17), F2: 388}
                    _ = v23
                    _ = v24
                }
                _ = fi21
            }
            switch fi20 {
            case -2147483648:
                v25 := GenS3{F0: 65535, F1: 18446744073709551615, F2: 0}
                _ = v25
            default:
                for fi26 := 0; fi26 < 8; fi26++ {
                    v27 := GenS1{F0: 2, F1: -83, F2: 926}
                    v2 = (v1 && (2 > 65535))
                    _ = fi26
                    _ = v27
                }
            }
            for fi28 := 0; fi28 < 16; fi28++ {
                v4 = 400
                v4 -= v4
                if ((255 >= 122) && v2) {
                    var v29 bool = v1
                    if (32767 > 201) {
                        v30 := GenS2{F0: fi28, F1: 255}
                        fmt.Println(fi20, fi20)
                        _ = v30
                    } else {
                        var v31 int8 = -128
                        _ = v31
                    }
                    if len(v3) > 0 {
                        v3[0] = 288
                    }
                    _ = v29
                } else {
                    switch fi28 {
                    case -893:
                        v1 = v2
                    case 948:
                        switch 128 {
                        case 32767:
                            switch fi28 {
                            case 9223372036854775807:
                                v3 = append(v3, (fi20 | 382))
                            default:
                                switch fi28 {
                                case 1:
                                    v32 := GenS3{F0: 1, F1: 0, F2: -960}
                                    _ = v32
                                case 2147483647:
                                    switch ("hello" + "\n\t") {
                                    case "":
                                        if v1 {
                                            var v33 int64 = 0
                                            v2 = v2
                                            fmt.Println(v17, v17, v33)
                                            _ = v33
                                        }
                                    case "foo bar":
                                        v34 := GenS1{F0: 747, F1: -128, F2: 9223372036854775807}
                                        _ = v34
                                    case "hello":
                                        v35 := GenS4{F0: v4, F1: 0, F2: 128}
                                        _ = v35
                                    }
                                case 128:
                                    v2 = (97 < 18)
                                    switch "foo bar" {
                                    case "a":
                                        v1 = true
                                    }
                                default:
                                    if v2 {
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                        v36 := GenS4{F0: 595, F1: 174, F2: 214}
                                        _ = v36
                                    }
                                }
                            }
                        case -380:
                            fmt.Println(v2, v4)
                        case -2147483648:
                            if ((1.5 * 3.14159) == 0.0) {
                                if ("hello" <= ("" + "hello")) {
                                    if len(v3) > 0 {
                                        v3[0] = (fi28 >> ((881) & 63))
                                    }
                                    v17 = (v17 * 1000)
                                }
                                var v37 float32 = -2.25
                                _ = v37
                            } else {
                                v4 = v4
                                v1 = v2
                            }
                        }
                    case -2147483648:
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                    }
                    v1 = (-2.25 == 0.1)
                }
                switch fi28 {
                case 42:
                    if v1 {
                        _ = len(v3)
                        var v38 int64 = -9223372036854775808
                        _ = v38
                    } else {
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        if ((!v2) || true) {
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            var v39 int8 = -1
                            var v40 int32 = v4
                            v3 = append(v3, fi20)
                            if (false && v1) {
                                v4 = (v40 | 1)
                                var v41 int = fi28
                                _ = v41
                            } else {
                                v42 := GenS3{F0: 255, F1: 4294967295, F2: 9223372036854775807}
                                fmt.Println(fi28, fi28, v40)
                                if v1 {
                                    _ = len(v3)
                                    if ((1.5 + (3.14159 + -2.25)) > 1e10) {
                                        v39 += (v39 + int8(fi28))
                                        var v43 float64 = -2.25
                                        _ = v43
                                    } else {
                                        if v2 {
                                            v44 := GenS5{F0: "a", F1: 2, F2: (v39 - v39)}
                                            var v45 uint64 = 2
                                            v39 -= v39
                                            var v46 float32 = float32(((1.5 * 1.5) - float64((0.1 - (1e10 - (1e10 + 0.0))))))
                                            _ = v44
                                            _ = v45
                                            _ = v46
                                        } else {
                                            var v47 int16 = (v17 + v17)
                                            v1 = v1
                                            _ = v47
                                        }
                                        var v48 uint16 = 93
                                        if false {
                                            switch (((((("a" + "foo bar") + "hello") + "foo bar") + "a") + "foo bar") + "foo bar") {
                                            case "foo bar":
                                                var v49 float64 = 1e10
                                                _ = v49
                                            case "":
                                                switch fi20 {
                                                case 32767:
                                                    fmt.Println(fi20, v2)
                                                case 208:
                                                    var v50 float64 = (1e10 / -2.25)
                                                    _ = v50
                                                default:
                                                    var v51 uint = 836
                                                    _ = v51
                                                }
                                            default:
                                                switch (((("hello" + "a") + "\n\t") + "") + "a") {
                                                case "foo bar":
                                                    var v52 bool = v1
                                                    _ = v52
                                                case "":
                                                    if true {
                                                        var v53 float32 = 1e10
                                                        var v54 string = "\n\t"
                                                        _ = v53
                                                        _ = v54
                                                    }
                                                case "\n\t":
                                                    switch (("" + "foo bar") + "foo bar") {
                                                    case "a":
                                                        fmt.Println(v2, v1)
                                                    case "hello":
                                                        switch fi28 {
                                                        case 1:
                                                            v1 = (1e10 >= 0.1)
                                                        default:
                                                            var v55 bool = true
                                                            _ = v55
                                                        }
                                                        v39 += 127
                                                        v2 = v1
                                                    default:
                                                        v2 = v2
                                                    }
                                                }
                                            }
                                            var v56 bool = (42 == 30)
                                            _ = v56
                                        }
                                        v39 += (-(v39))
                                        _ = v48
                                    }
                                } else {
                                    v40 = 0
                                    if ((v39 ^ v39) >= (v39 / ((-128) | 1))) {
                                        v57 := GenS6{F0: "\n\t"}
                                        _ = v57
                                    }
                                }
                                _ = v42
                            }
                            _ = v39
                            _ = v40
                        }
                        v4 = (v4 ^ v4)
                    }
                case 0:
                    v58 := GenS3{F0: 6, F1: 4294967295, F2: -1}
                    _ = v58
                default:
                    v3 = append(v3, (fi20 % (((fi20 ^ 1)) | 1)))
                }
                _ = fi28
            }
        }
        v3 = append(v3, (fi20 | (fi20 >> ((2) & 63))))
        v59 := GenS3{F0: 1, F1: 255, F2: -285}
        _ = fi20
        _ = v59
    }
    v2 = v1
    if (1 != 1) {
        var v60 uint32 = 661
        _ = v60
    } else {
        fmt.Println(v2, v2)
        v61 := GenS5{F0: "a", F1: 42, F2: 22}
        var v62 int = -2147483648
        _ = v61
        _ = v62
    }
    switch -2147483648 {
    case 1:
        var v63 uint64 = 4294967295
        _ = v63
    }
    if (42 < 1) {
        if ((false && (255 < 42)) || v2) {
            for fi64 := 0; fi64 < 1; fi64++ {
                for fi65 := 0; fi65 < 8; fi65++ {
                    v2 = (42 > 4294967295)
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    _ = fi65
                }
                var v66 int64 = 816
                _ = fi64
                _ = v66
            }
            v67 := GenS7{F0: 1}
            if len(v3) > 0 {
                v3[0] = 2147483647
            }
            _ = v67
        }
        if ((v4 >> ((276) & 63)) >= v4) {
            _ = len(v3)
            v1 = (0.0 > 0.0)
            switch 0 {
            case 32768:
                for fi68 := 0; fi68 < 20; fi68++ {
                    if (false || v1) {
                        fmt.Println(v1)
                        var v69 float64 = 0.1
                        _ = v69
                    }
                    for fi70 := 0; fi70 < 5; fi70++ {
                        _ = len(v3)
                        var v71 float64 = 1.5
                        var v72 uint32 = 42
                        _ = fi70
                        _ = v71
                        _ = v72
                    }
                    _ = fi68
                }
            case -391:
                fmt.Println(v2)
            case 128:
                v73 := GenS8{F0: -128, F1: 356, F2: 20}
                _ = v73
            default:
                for fi74 := 0; fi74 < 6; fi74++ {
                    var v75 int = (fi74 ^ (fi74 ^ 32768))
                    v2 = (((((1 < 100) && v1) && v1) && v1) && v2)
                    _ = fi74
                    _ = v75
                }
            }
        }
    } else {
        var v76 int8 = 42
        _ = v76
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v4
}
