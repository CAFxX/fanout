package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16 }
type GenS2 struct { F0 int64; F1 int32 }
type GenS3 struct { F0 int64; F1 string; F2 int }
type GenS4 struct { F0 int64; F1 uint64 }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 int }
type GenS7 struct { F0 int16; F1 uint8 }
type GenS8 struct { F0 bool }
type GenS9 struct { F0 int32; F1 uint64; F2 int64 }
type GenS10 struct { F0 int16; F1 uint }

func main() {
    v1 := []uint32{2, 42, 4294967295}
    var v2 float64 = 1e10
    var v3 bool = false
    var v4 uint64 = 1
    if false {
        for fi5 := 0; fi5 < 13; fi5++ {
            fmt.Println(v4)
            if len(v1) > 0 {
                v1[0] = 2147483647
            }
            var v6 uint16 = 2
            _ = fi5
            _ = v6
        }
        for fi7 := 0; fi7 < 15; fi7++ {
            var v8 int64 = 428
            _ = fi7
            _ = v8
        }
    }
    for fi9 := 0; fi9 < 15; fi9++ {
        v2 -= v2
        v4 += 0
        _ = fi9
    }
    v10 := GenS1{F0: 50, F1: 0}
    if false {
        v4 += v4
    } else {
        fmt.Println(v4)
        v1 = append(v1, 526)
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    switch 32768 {
    case 255:
        var v11 string = "\n\t"
        _ = v11
    default:
        v12 := GenS1{F0: 127, F1: -32768}
        _ = v12
    }
    for fi13 := 0; fi13 < 12; fi13++ {
        if v3 {
            v14 := GenS1{F0: 127, F1: 32767}
            if (!v3) {
                if false {
                    for fi15 := 0; fi15 < 8; fi15++ {
                        var v16 string = "a"
                        switch v16 {
                        case "":
                            v4 = (^(v4))
                        case "\n\t":
                            var v17 uint64 = (v4 | (^(v4)))
                            var v18 int64 = 1
                            _ = v17
                            _ = v18
                        case "hello":
                            if ((-588 > 610) && false) {
                                switch "\n\t" {
                                case "a":
                                    if (!v3) {
                                        v19 := GenS1{F0: 1, F1: 1}
                                        var v20 int = int(v4)
                                        _ = v19
                                        _ = v20
                                    }
                                case "\n\t":
                                    fmt.Println(fi15)
                                default:
                                    if len(v1) > 0 {
                                        v1[0] = 42
                                    }
                                }
                                switch 261 {
                                case 42:
                                    v2 = 0.1
                                    v4 = (v4 | (v4 << ((717) & 63)))
                                }
                                v3 = true
                                v21 := GenS2{F0: 0, F1: 0}
                                _ = v21
                            } else {
                                fmt.Println(v4, v16)
                                v2 += (v2 * float64(3.14159))
                                v22 := GenS3{F0: -813, F1: "\n\t", F2: 32768}
                                var v23 uint = 255
                                var v24 float32 = 3.14159
                                _ = v22
                                _ = v23
                                _ = v24
                            }
                        }
                        _ = fi15
                        _ = v16
                    }
                } else {
                    for fi25 := 0; fi25 < 6; fi25++ {
                        if (2 < 255) {
                            fmt.Println(v4, fi25, fi13)
                            if (!true) {
                                v26 := GenS1{F0: 1, F1: -1}
                                var v27 uint32 = 42
                                fmt.Println(fi13, v3)
                                v3 = v3
                                var v28 string = "foo bar"
                                v4 += (v4 + (v4 << ((58) & 63)))
                                _ = v26
                                _ = v27
                                _ = v28
                            }
                            v29 := GenS4{F0: -1000, F1: v4}
                            _ = v29
                        } else {
                            var v30 uint8 = 1
                            var v31 uint8 = (v30 | (v30 >> ((1) & 63)))
                            if (1 >= 921) {
                                v32 := GenS5{F0: -1}
                                v1 = append(v1, 13)
                                v1 = append(v1, 466)
                                _ = v32
                            }
                            _ = v30
                            _ = v31
                        }
                        _ = fi25
                    }
                }
                v33 := GenS6{F0: fi13}
                for fi34 := 0; fi34 < 6; fi34++ {
                    v1 = append(v1, 1)
                    _ = fi34
                }
                _ = v33
            }
            _ = v14
        } else {
            v2 = (1e10 - v2)
        }
        fmt.Println(v2)
        fmt.Println(v3, v4)
        _ = fi13
    }
    var v35 float64 = float64(0.0)
    switch 2147483647 {
    case 4294967295:
        if (-32768 != -877) {
            for fi36 := 0; fi36 < 19; fi36++ {
                if v3 {
                    v4 += v4
                }
                _ = fi36
            }
        }
    }
    if false {
        var v37 float64 = v35
        if (127 != 1) {
            v2 = ((1e10 / (0.1 * -2.25)) / (-2.25 + (0.0 * v35)))
            for fi38 := 0; fi38 < 16; fi38++ {
                v3 = (v4 == 18446744073709551615)
                v39 := GenS2{F0: -1, F1: 1}
                _ = fi38
                _ = v39
            }
            if (!true) {
                if v3 {
                    _ = len(v1)
                    if v3 {
                        var v40 int32 = 0
                        switch "foo bar" {
                        case "a":
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                        }
                        _ = v40
                    } else {
                        var v41 int8 = 62
                        _ = v41
                    }
                    if v3 {
                        fmt.Println(v2, v2)
                        v2 = 0.1
                    } else {
                        var v42 int = -2147483648
                        if v3 {
                            for fi43 := 0; fi43 < 10; fi43++ {
                                v44 := GenS7{F0: -32768, F1: 42}
                                _ = fi43
                                _ = v44
                            }
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                        }
                        for fi45 := 0; fi45 < 10; fi45++ {
                            fmt.Println(fi45, fi45, v37)
                            v46 := GenS6{F0: (fi45 | 721)}
                            var v47 int8 = 1
                            _ = fi45
                            _ = v46
                            _ = v47
                        }
                        _ = v42
                    }
                }
                switch -503 {
                case 9223372036854775807:
                    fmt.Println(v37, v4, v2)
                default:
                    v48 := GenS3{F0: -452, F1: "\n\t", F2: 254}
                    _ = v48
                }
                var v49 int64 = 9223372036854775807
                _ = v49
            } else {
                if len(v1) > 0 {
                    v1[0] = 1
                }
                switch -604 {
                case 1:
                    v3 = (v3 && ((-2.25 - (((float32(1e10) / 1e10) / 3.14159) / 1.0)) >= (float32(v35) / (1.5 - float32(v2)))))
                default:
                    switch ("a" + "a") {
                    case "hello":
                        switch "a" {
                        case "foo bar":
                            var v50 bool = (176 >= 0)
                            _ = v50
                        case "\n\t":
                            fmt.Println(v37)
                        case "":
                            v35 = v35
                        default:
                            if (883 == -221) {
                                v51 := GenS4{F0: 0, F1: v4}
                                _ = v51
                            } else {
                                v52 := GenS5{F0: -1}
                                fmt.Println(v37)
                                v53 := GenS5{F0: 1}
                                _ = v52
                                _ = v53
                            }
                        }
                    default:
                        if (187 == 999) {
                            var v54 int = 154
                            v4 = (v4 * (v4 + v4))
                            _ = v54
                        } else {
                            var v55 string = "a"
                            if (false && (v3 && true)) {
                                v56 := GenS5{F0: 42}
                                _ = v56
                            } else {
                                fmt.Println(v55, v37)
                            }
                            if (v3 || ((186 == 2) && v3)) {
                                v57 := GenS8{F0: (v3 || true)}
                                _ = v57
                            } else {
                                switch 32768 {
                                case 32768:
                                    v58 := GenS9{F0: -2147483648, F1: v4, F2: 9223372036854775807}
                                    _ = v58
                                case 997:
                                    if false {
                                        if v3 {
                                            fmt.Println(v4, v3)
                                            var v59 float64 = v35
                                            _ = v59
                                        }
                                        for fi60 := 0; fi60 < 2; fi60++ {
                                            v55 = "\n\t"
                                            _ = fi60
                                        }
                                        v3 = (2147483647 >= 36)
                                    }
                                default:
                                    v61 := GenS4{F0: 42, F1: v4}
                                    _ = v61
                                }
                            }
                            _ = v55
                        }
                    }
                }
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            }
        } else {
            var v62 int = 0
            v4 = 740
            _ = v62
        }
        _ = v37
    }
    if (-64 != -1) {
        var v63 int8 = 100
        v35 -= 1.5
        var v64 uint64 = v4
        _ = v63
        _ = v64
    } else {
        var v65 uint = 627
        v4 -= 1
        switch "\n\t" {
        case "foo bar":
            v2 -= v2
        case "hello":
            for fi66 := 0; fi66 < 5; fi66++ {
                var v67 float32 = 3.14159
                if len(v1) > 0 {
                    v1[0] = 706
                }
                _ = fi66
                _ = v67
            }
        case "\n\t":
            if (721 != -1) {
                if true {
                    v4 -= (v4 * 2)
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                } else {
                    v2 -= float64((0.1 * 1.5))
                    fmt.Println(v65)
                    var v68 int32 = 89
                    _ = v68
                }
                switch 37 {
                case 32767:
                    if (!true) {
                        if (255 == uint16(v65)) {
                            v2 = v2
                        } else {
                            var v69 int32 = 491
                            v1 = append(v1, 1)
                            _ = v69
                        }
                        _ = len(v1)
                        v70 := GenS10{F0: 1000, F1: v65}
                        _ = v70
                    }
                case -851:
                    if len(v1) > 0 {
                        v1[0] = 245
                    }
                }
                var v71 string = ("a" + "a")
                _ = v71
            } else {
                v4 = (v4 << ((v65) & 63))
            }
        default:
            fmt.Println(v65)
        }
        _ = v65
    }
    v72 := GenS8{F0: v3}
    fmt.Println(v2)
    _ = len(v1)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v35
    _ = v4
    _ = v72
}
