package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int16; F1 uint16; F2 int16 }
type GenS4 struct { F0 uint64; F1 uint8; F2 int16 }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 bool; F1 int64; F2 int32 }
type GenS7 struct { F0 bool; F1 int8; F2 int8 }
type GenS8 struct { F0 uint64; F1 int64 }
type GenS9 struct { F0 uint64; F1 int8; F2 int8 }
type GenS10 struct { F0 string }
type GenS11 struct { F0 int64; F1 int; F2 bool }
type GenS12 struct { F0 uint64; F1 uint16; F2 int16 }
type GenS13 struct { F0 string; F1 int8; F2 uint8 }
type GenS14 struct { F0 int64; F1 uint32; F2 int16 }
func genFn0(p0 string, p1 int) uint32 {
    v6 := GenS1{F0: (p0 + (p0 + (((p0 + "") + "") + "a")))}
    p1 = p1
    switch (p1 >> ((255) & 63)) {
    case -9223372036854775808:
        p1 += -156
    case 247:
        p1 = p1
    }
    fmt.Println(p0, p0)
    _ = p0
    _ = p1
    _ = v6
    return 887
}

func main() {
    v1 := []string{}
    var v2 int8 = 82
    v3 := []string{}
    var v4 int8 = v2
    v5 := make(map[uint32]bool)
    var v7 int16 = 32767
    var v8 int = 586
    v8 = v8
    var v9 bool = (!(127 > 255))
    v9 = (2147483648 <= 65535)
    v9 = (2 == 2147483647)
    var v10 uint32 = 1
    v8 = (v8 & 32768)
    for fi11 := 0; fi11 < 12; fi11++ {
        var v12 bool = v9
        var v13 int16 = (v7 * (v7 + 32767))
        _ = fi11
        _ = v12
        _ = v13
    }
    v14 := GenS2{F0: v9}
    for fi15 := 0; fi15 < 2; fi15++ {
        if true {
            switch "hello" {
            case "\n\t":
                switch ("foo bar" + ("\n\t" + "foo bar")) {
                case "\n\t":
                    v9 = v9
                case "":
                    v16 := GenS3{F0: (-(v7)), F1: 42, F2: int16(v10)}
                    _ = v16
                default:
                    delete(v5, (v10 | genFn0(("\n\t" + "foo bar"), fi15)))
                }
            default:
                for fi17 := 0; fi17 < 4; fi17++ {
                    var v18 uint64 = 2
                    _ = fi17
                    _ = v18
                }
            }
            v19 := GenS4{F0: 2, F1: 127, F2: 32767}
            delete(v5, (v10 >> ((255) & 63)))
            _ = v19
        } else {
            if v9 {
                for fi20 := 0; fi20 < 20; fi20++ {
                    switch fi15 {
                    case 470:
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                    case 128:
                        var v21 int = fi20
                        _ = v21
                    case 42:
                        v9 = v9
                    }
                    _ = fi20
                }
                v22 := GenS5{F0: 935}
                _ = len(v5)
                _ = v22
            } else {
                if true {
                    if ((v7 / (((v7 ^ v7)) | 1)) != (v7 + v7)) {
                        delete(v5, (^(v10)))
                        v23 := GenS1{F0: ("\n\t" + "a")}
                        v10 += (v10 * 408)
                        _ = v23
                    } else {
                        v24 := GenS6{F0: v9, F1: 9223372036854775807, F2: 42}
                        for fi25 := 0; fi25 < 13; fi25++ {
                            var v26 string = ""
                            if v9 {
                                var v27 float64 = (3.14159 / (0.0 - (-2.25 / ((-2.25 - 1e10) / 1.0))))
                                var v28 float64 = v27
                                _ = v27
                                _ = v28
                            } else {
                                var v29 float64 = 0.1
                                v10 -= v10
                                fmt.Println(v29, v26)
                                _ = v29
                            }
                            v7 += v7
                            switch -9223372036854775808 {
                            case -327:
                                v7 += int16(v10)
                            case 127:
                                if true {
                                    if (fi25 >= (fi15 * 32768)) {
                                        if (0.1 != 1e10) {
                                            switch 9223372036854775807 {
                                            case 32767:
                                                switch ((v26 + ("hello" + v26)) + (("" + v26) + v26)) {
                                                case "a":
                                                    v30 := GenS7{F0: (3.14159 == -2.25), F1: (^(v4)), F2: 1}
                                                    _ = v30
                                                case "":
                                                    var v31 bool = v9
                                                    _ = v31
                                                case "foo bar":
                                                    v32 := GenS7{F0: v9, F1: v2, F2: (v4 / (((v4 | -77)) | 1))}
                                                    _ = v32
                                                }
                                                v2 = (v4 - (v4 | int8(v8)))
                                                fmt.Println(v4)
                                                v8 += (-(v8))
                                            }
                                            v8 -= (fi25 % ((fi15) | 1))
                                            var v33 int32 = -2147483648
                                            _ = v33
                                        }
                                        v34 := GenS8{F0: 2, F1: 940}
                                        fmt.Println(v26)
                                        _ = v34
                                    } else {
                                        var v35 uint8 = 127
                                        if false {
                                            var v36 uint32 = v10
                                            v37 := GenS5{F0: 65535}
                                            v38 := GenS2{F0: (549 != 54)}
                                            _ = v36
                                            _ = v37
                                            _ = v38
                                        } else {
                                            v10 += (v10 * uint32(v35))
                                            if len(v3) > 0 {
                                                v3[0] = v26
                                            }
                                        }
                                        _ = len(v5)
                                        _ = v35
                                    }
                                    switch v26 {
                                    case "foo bar":
                                        if v9 {
                                            v39 := GenS5{F0: 631}
                                            v10 -= (v10 >> ((42) & 63))
                                            _ = v39
                                        }
                                    case "a":
                                        v10 = (^(v10))
                                    case "\n\t":
                                        var v40 uint64 = 42
                                        _ = v40
                                    }
                                    var v41 int64 = 9223372036854775807
                                    _ = v41
                                } else {
                                    v42 := GenS2{F0: v9}
                                    _ = v42
                                }
                            default:
                                fmt.Println(fi25, v2, v2)
                            }
                            v5[2147483648] = (!v9)
                            _ = fi25
                            _ = v26
                        }
                        _ = v24
                    }
                }
                switch 0 {
                case 0:
                    for fi43 := 0; fi43 < 12; fi43++ {
                        v3 = append(v3, ("" + "foo bar"))
                        _ = fi43
                    }
                case 127:
                    v7 -= v7
                }
            }
            var v44 uint = 0
            _ = v44
        }
        var v45 uint = 42
        v5[(v10 * v10)] = v9
        _ = fi15
        _ = v45
    }
    if (!(827 >= (v10 % ((v10) | 1)))) {
        v9 = true
        if (1 == 484) {
            if false {
                fmt.Println(v8)
                var v46 uint16 = 42
                _ = v46
            } else {
                v47 := GenS8{F0: 18446744073709551615, F1: 9223372036854775807}
                var v48 int16 = v7
                fmt.Println(v7)
                _ = v47
                _ = v48
            }
        }
    }
    v2 = 41
    if (!true) {
        switch -2147483648 {
        case 127:
            v10 = v10
        default:
            v49 := GenS2{F0: true}
            _ = v49
        }
        var v50 uint32 = (v10 % ((genFn0(((("a" + (("hello" + ("foo bar" + "a")) + ("a" + "a"))) + ("a" + "\n\t")) + "a"), v8)) | 1))
        _ = v50
    }
    var v51 int64 = -255
    if v9 {
        fmt.Println(v4)
        v52 := GenS3{F0: v7, F1: 643, F2: (v7 & 1)}
        if (!(0 <= (^(v10)))) {
            for fi53 := 0; fi53 < 2; fi53++ {
                var v54 int64 = (v51 >> ((1) & 63))
                _ = len(v3)
                _ = fi53
                _ = v54
            }
            if (157 < (v10 % ((v10) | 1))) {
                fmt.Println(v2, v8)
                for fi55 := 0; fi55 < 6; fi55++ {
                    if v9 {
                        for fi56 := 0; fi56 < 8; fi56++ {
                            fmt.Println(v2, v8, v51)
                            _ = fi56
                        }
                        if v9 {
                            for fi57 := 0; fi57 < 20; fi57++ {
                                var v58 float32 = 1.5
                                v59 := GenS9{F0: 2, F1: -109, F2: v2}
                                _ = fi57
                                _ = v58
                                _ = v59
                            }
                            v60 := v5[genFn0("\n\t", (v8 * v8))]
                            _ = v60
                        } else {
                            v7 = (v7 / (((v7 / ((1000) | 1))) | 1))
                        }
                        fmt.Println(v7, fi55)
                    } else {
                        var v61 string = "hello"
                        v4 -= v4
                        _ = v61
                    }
                    switch fi55 {
                    case 127:
                        for fi62 := 0; fi62 < 17; fi62++ {
                            fmt.Println(v10, v51, v4)
                            v63 := GenS10{F0: ""}
                            v9 = false
                            _ = fi62
                            _ = v63
                        }
                    case 32768:
                        v2 = v2
                    }
                    _ = fi55
                }
                var v64 int32 = 42
                _ = v64
            }
        }
        _ = v52
    } else {
        fmt.Println(v9, v2)
        v65 := GenS11{F0: v51, F1: (v8 * v8), F2: v9}
        fmt.Println(v4, v2)
        _ = v65
    }
    switch v8 {
    case 4294967295:
        v4 -= (v2 - (v2 / ((v4) | 1)))
    case 32768:
        for fi66 := 0; fi66 < 20; fi66++ {
            v10 -= (v10 * (^(v10)))
            for fi67 := 0; fi67 < 15; fi67++ {
                if ((uint64(v10) >= 18446744073709551615) || v9) {
                    switch 1 {
                    case -515:
                        if (v9 && ("" >= "a")) {
                            v3 = append(v3, ((("a" + ("foo bar" + "")) + ("" + ("hello" + ""))) + ""))
                            var v68 int64 = v51
                            switch "a" {
                            case "\n\t":
                                var v69 float64 = 0.0
                                _ = v69
                            case "foo bar":
                                var v70 int32 = 42
                                _ = v70
                            default:
                                var v71 int = (^(fi66))
                                _ = v71
                            }
                            v2 = v4
                            v7 += 252
                            _ = v68
                        }
                    }
                    v72 := GenS4{F0: 198, F1: 163, F2: (v7 | v7)}
                    fmt.Println(v2)
                    _ = v72
                } else {
                    if false {
                        var v73 int32 = 1
                        var v74 uint8 = 128
                        v9 = (!v9)
                        _ = v73
                        _ = v74
                    }
                    v75 := v5[(^(v10))]
                    if ((v51 >> ((767) & 63)) >= (v51 >> ((627) & 63))) {
                        var v76 uint32 = (v10 - 42)
                        _ = v76
                    } else {
                        var v77 float32 = 0.1
                        _ = v77
                    }
                    if v9 {
                        var v78 float32 = 0.0
                        switch ("\n\t" + ("hello" + ("\n\t" + "a"))) {
                        case "hello":
                            var v79 bool = (v10 < genFn0("", fi66))
                            _ = v79
                        }
                        v7 = 1
                        _ = v78
                    } else {
                        fmt.Println(fi67, v4)
                        if (!(!(42 < 750))) {
                            var v80 int = -252
                            _ = v80
                        }
                        if ((v7 ^ 42) == 1000) {
                            switch "hello" {
                            case "\n\t":
                                v81 := GenS12{F0: 42, F1: 385, F2: (v7 % ((1000) | 1))}
                                _ = v81
                            case "":
                                var v82 bool = ((362 >= 2) || (9223372036854775807 <= v51))
                                var v83 uint32 = v10
                                v10 += (v10 >> ((333) & 63))
                                var v84 uint = 1
                                v85 := GenS11{F0: (v51 ^ v51), F1: -9223372036854775808, F2: (0 >= 4294967295)}
                                _ = v82
                                _ = v83
                                _ = v84
                                _ = v85
                            default:
                                if (true || false) {
                                    if v75 {
                                        var v86 uint = 887
                                        _ = v86
                                    } else {
                                        v3 = append(v3, (("a" + "foo bar") + ""))
                                        var v87 float32 = 1e10
                                        delete(v5, v10)
                                        _ = v87
                                    }
                                    v10 = genFn0(("\n\t" + ("\n\t" + "")), (fi67 >> ((311) & 63)))
                                    var v88 int32 = 0
                                    _ = v88
                                }
                            }
                            var v89 uint32 = v10
                            v90 := GenS4{F0: 4294967295, F1: 128, F2: int16(v4)}
                            _ = v89
                            _ = v90
                        }
                    }
                    if (!v75) {
                        v75 = (true && (uint16(v10) < 42))
                        v91 := GenS8{F0: 18446744073709551615, F1: (v51 >> ((18446744073709551615) & 63))}
                        _ = v91
                    }
                    var v92 int16 = int16(v51)
                    v2 += v4
                    _ = v75
                    _ = v92
                }
                fmt.Println(v7, v4)
                var v93 int16 = -32768
                _ = fi67
                _ = v93
            }
            for fi94 := 0; fi94 < 19; fi94++ {
                switch 32768 {
                case 302:
                    v3 = append(v3, "hello")
                case -760:
                    var v95 uint16 = 576
                    _ = v95
                case 846:
                    var v96 bool = false
                    _ = v96
                default:
                    switch (v8 + 0) {
                    case 128:
                        v9 = (!true)
                    case 32768:
                        v97 := GenS13{F0: (("" + ("\n\t" + "\n\t")) + ("\n\t" + "foo bar")), F1: v2, F2: 0}
                        _ = v97
                    case -471:
                        switch fi94 {
                        case -1:
                            v98 := GenS14{F0: v51, F1: v10, F2: (v7 % ((v7) | 1))}
                            _ = v98
                        }
                    }
                }
                _ = fi94
            }
            _ = fi66
        }
    }
    var v99 float32 = ((-2.25 / (-2.25 + (1.5 - 1e10))) - (float32(((0.1 * (0.0 - -2.25)) / float64(float32(-2.25)))) / 0.1))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v51:", v51)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    fmt.Println("v99:", v99)
    _ = v1
    _ = v10
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v51
    _ = v7
    _ = v8
    _ = v9
    _ = v99
}
