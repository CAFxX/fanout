package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint64; F2 uint8 }
type GenS2 struct { F0 int64; F1 int32 }
type GenS3 struct { F0 int; F1 int32 }
type GenS4 struct { F0 int16; F1 string; F2 uint64 }
type GenS5 struct { F0 int16; F1 int8 }
type GenS6 struct { F0 int; F1 string }
type GenS7 struct { F0 bool }

func main() {
    var v1 uint64 = 766
    var v2 uint16 = 1
    var v3 uint64 = (v1 % ((v1) | 1))
    var v4 uint16 = v2
    v5 := []int{205, 128, 1, 9223372036854775807}
    switch "" {
    case "foo bar":
        v6 := GenS1{F0: 33, F1: v1, F2: 255}
        _ = v6
    default:
        switch ("foo bar" + "\n\t") {
        case "":
            for fi7 := 0; fi7 < 1; fi7++ {
                var v8 int64 = -1
                if false {
                    if (100 < int8(v2)) {
                        if (int8(v8) < -18) {
                            var v9 uint64 = v1
                            _ = len(v5)
                            v3 -= 640
                            _ = v9
                        } else {
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                        }
                        v1 = 42
                    } else {
                        switch fi7 {
                        case -468:
                            for fi10 := 0; fi10 < 6; fi10++ {
                                switch ("foo bar" + ("hello" + "\n\t")) {
                                case "hello":
                                    if true {
                                        var v11 int8 = -128
                                        var v12 int8 = (-(v11))
                                        if false {
                                            if (1 == 4294967295) {
                                                _ = len(v5)
                                            } else {
                                                v2 -= 485
                                                var v13 string = ""
                                                var v14 float32 = (float32(0.1) / 1e10)
                                                _ = v13
                                                _ = v14
                                            }
                                            v4 -= (v4 ^ (v4 + 65535))
                                            v15 := GenS1{F0: 102, F1: (v3 << ((0) & 63)), F2: 106}
                                            _ = v15
                                        }
                                        _ = v11
                                        _ = v12
                                    } else {
                                        switch (fi7 ^ 1) {
                                        case 32767:
                                            v3 = uint64(v4)
                                            v16 := GenS1{F0: 127, F1: v3, F2: 238}
                                            _ = v16
                                        case 0:
                                            var v17 int = (-(fi10))
                                            _ = v17
                                        case -1:
                                            fmt.Println(v3, fi10)
                                        }
                                        if true {
                                            _ = len(v5)
                                            if len(v5) >= 2 {
                                                v5 = v5[1:len(v5)]
                                            }
                                            if ("\n\t" <= (("hello" + ("hello" + ("hello" + (("a" + "a") + ((((("hello" + "") + "") + "foo bar") + "a") + "\n\t"))))) + "hello")) {
                                                var v18 int32 = -881
                                                _ = v18
                                            } else {
                                                switch -28 {
                                                case -309:
                                                    if (42 <= 1) {
                                                        if (2147483647 == 2147483648) {
                                                            v19 := GenS2{F0: (-(v8)), F1: 0}
                                                            v8 = v8
                                                            var v20 string = ("foo bar" + "")
                                                            v5 = append(v5, 128)
                                                            _ = v19
                                                            _ = v20
                                                        } else {
                                                            if (127 != 128) {
                                                                if len(v5) >= 2 {
                                                                    v5 = v5[1:len(v5)]
                                                                }
                                                                fmt.Println(v8, v2)
                                                            } else {
                                                                if (-42 > 42) {
                                                                    var v21 int64 = (v8 / (((v8 % ((v8) | 1))) | 1))
                                                                    _ = v21
                                                                } else {
                                                                    if true {
                                                                        if true {
                                                                            var v22 float32 = 0.0
                                                                            if len(v5) >= 2 {
                                                                                v5 = v5[1:len(v5)]
                                                                            }
                                                                            v23 := GenS1{F0: 52, F1: (v3 + 1), F2: 42}
                                                                            v1 = v3
                                                                            v24 := GenS2{F0: 449, F1: 1}
                                                                            _ = v22
                                                                            _ = v23
                                                                            _ = v24
                                                                        }
                                                                        var v25 uint32 = 2147483647
                                                                        v25 -= 42
                                                                        _ = v25
                                                                    }
                                                                }
                                                            }
                                                        }
                                                        if (159 < (fi7 ^ fi7)) {
                                                            v26 := GenS1{F0: -1, F1: v1, F2: 255}
                                                            v8 -= v8
                                                            v1 -= (v1 % (((v3 & (v1 - v1))) | 1))
                                                            _ = v26
                                                        } else {
                                                            var v27 string = ("foo bar" + "hello")
                                                            v27 = v27
                                                            _ = v27
                                                        }
                                                        _ = len(v5)
                                                    }
                                                    if (!false) {
                                                        var v28 int8 = 1
                                                        v29 := GenS1{F0: v28, F1: 45, F2: uint8(v3)}
                                                        v4 += (v2 & v4)
                                                        _ = v28
                                                        _ = v29
                                                    } else {
                                                        if len(v5) > 0 {
                                                            v5[0] = fi7
                                                        }
                                                        var v30 int64 = v8
                                                        _ = v30
                                                    }
                                                    if len(v5) > 0 {
                                                        v5[0] = 343
                                                    }
                                                default:
                                                    if len(v5) > 0 {
                                                        v5[0] = 32767
                                                    }
                                                }
                                            }
                                        } else {
                                            var v31 int16 = -32768
                                            var v32 uint64 = (v1 & (^(v1)))
                                            _ = v31
                                            _ = v32
                                        }
                                        v4 = (v4 ^ (v4 & v4))
                                    }
                                case "foo bar":
                                    fmt.Println(v8)
                                }
                                v2 = v4
                                var v33 int8 = 42
                                _ = fi10
                                _ = v33
                            }
                        case -1:
                            var v34 int64 = int64(v2)
                            _ = v34
                        }
                        var v35 uint16 = (v4 >> ((255) & 63))
                        _ = v35
                    }
                } else {
                    var v36 uint16 = v2
                    _ = v36
                }
                v37 := GenS1{F0: -53, F1: v1, F2: 172}
                _ = fi7
                _ = v37
                _ = v8
            }
        case "a":
            if true {
                v2 = v4
                v2 += 0
                _ = len(v5)
            }
        case "foo bar":
            v38 := GenS3{F0: -625, F1: 2147483647}
            _ = v38
        }
    }
    v2 = (^(v2))
    for fi39 := 0; fi39 < 11; fi39++ {
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        v40 := GenS2{F0: -1, F1: -307}
        v3 = 31
        _ = fi39
        _ = v40
    }
    for fi41 := 0; fi41 < 15; fi41++ {
        var v42 uint32 = 42
        if (true || (false && true)) {
            v5 = append(v5, fi41)
            if false {
                fmt.Println(v4, v1, fi41)
            } else {
                for fi43 := 0; fi43 < 6; fi43++ {
                    if (-128 > 102) {
                        v42 = (^(v42))
                        v44 := GenS1{F0: 1, F1: (v1 | (v1 >> ((1) & 63))), F2: 2}
                        var v45 uint64 = v3
                        _ = v44
                        _ = v45
                    } else {
                        if len(v5) > 0 {
                            v5[0] = (fi43 ^ (fi43 + (fi41 / ((fi43) | 1))))
                        }
                        if false {
                            _ = len(v5)
                        }
                        var v46 uint = 2
                        v46 = 42
                        v4 = v2
                        v3 += v1
                        if false {
                            var v47 float64 = (-2.25 / 0.1)
                            _ = v47
                        } else {
                            v2 -= v2
                        }
                        fmt.Println(v42, v3)
                        if (!(v46 <= v46)) {
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                            fmt.Println(v4, v1, v1)
                            v4 -= 2
                        } else {
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                        }
                        var v48 string = "a"
                        _ = v46
                        _ = v48
                    }
                    _ = fi43
                }
            }
        }
        var v49 int16 = -32768
        _ = fi41
        _ = v42
        _ = v49
    }
    var v50 uint8 = 128
    v1 = (^(v3))
    var v51 float32 = 0.0
    if false {
        v5 = append(v5, 32768)
        var v52 uint32 = 2147483648
        _ = v52
    }
    v53 := GenS2{F0: 231, F1: 729}
    _ = len(v5)
    if (-1 >= 1) {
        var v54 int64 = -9223372036854775808
        if len(v5) > 0 {
            v5[0] = 0
        }
        for fi55 := 0; fi55 < 15; fi55++ {
            v56 := GenS4{F0: 1, F1: (("hello" + ("\n\t" + "hello")) + "\n\t"), F2: v1}
            v1 += 0
            _ = fi55
            _ = v56
        }
        _ = v54
    }
    var v57 int = 127
    fmt.Println(v57, v51, v57)
    v58 := GenS5{F0: 42, F1: 0}
    v1 += uint64(v50)
    switch 4294967295 {
    case -2147483648:
        for fi59 := 0; fi59 < 10; fi59++ {
            for fi60 := 0; fi60 < 17; fi60++ {
                var v61 uint8 = (v50 >> ((42) & 63))
                v62 := GenS6{F0: (-(fi60)), F1: ("hello" + "")}
                _ = fi60
                _ = v61
                _ = v62
            }
            if true {
                fmt.Println(v1)
            } else {
                v63 := GenS1{F0: 1, F1: (^(v1)), F2: (^(v50))}
                v1 = 2
                _ = v63
            }
            v64 := GenS3{F0: int(v2), F1: 2147483647}
            _ = fi59
            _ = v64
        }
    default:
        for fi65 := 0; fi65 < 16; fi65++ {
            fmt.Println(v2, v2)
            if (v51 >= v51) {
                v57 = v57
                switch "" {
                case "\n\t":
                    for fi66 := 0; fi66 < 8; fi66++ {
                        if true {
                            var v67 int16 = -354
                            _ = v67
                        } else {
                            fmt.Println(v50)
                        }
                        _ = fi66
                    }
                case "hello":
                    if false {
                        if true {
                            v1 += (v1 << ((334) & 63))
                            if (2147483647 == -565) {
                                for fi68 := 0; fi68 < 13; fi68++ {
                                    v69 := GenS6{F0: fi68, F1: ("foo bar" + "foo bar")}
                                    _ = fi68
                                    _ = v69
                                }
                                for fi70 := 0; fi70 < 1; fi70++ {
                                    v3 = (v1 * v3)
                                    _ = fi70
                                }
                                if true {
                                    v71 := GenS3{F0: v57, F1: 2147483647}
                                    _ = v71
                                } else {
                                    var v72 bool = (27 <= (v50 % (((v50 / ((234) | 1))) | 1)))
                                    for fi73 := 0; fi73 < 20; fi73++ {
                                        switch (fi65 + -5) {
                                        case 127:
                                            v2 -= (v4 >> ((42) & 63))
                                        case -2147483648:
                                            v74 := GenS7{F0: ((^(v4)) != (v2 - 1))}
                                            _ = v74
                                        case -1:
                                            var v75 string = ("" + ("\n\t" + "a"))
                                            _ = v75
                                        default:
                                            v51 = v51
                                        }
                                        fmt.Println(v2, fi65, v2)
                                        _ = fi73
                                    }
                                    v57 = 1
                                    _ = v72
                                }
                            }
                        } else {
                            var v76 string = ("\n\t" + "a")
                            v77 := GenS2{F0: -524, F1: 2147483647}
                            _ = v76
                            _ = v77
                        }
                    } else {
                        v78 := GenS7{F0: true}
                        _ = v78
                    }
                default:
                    var v79 float64 = 0.1
                    _ = v79
                }
            } else {
                for fi80 := 0; fi80 < 8; fi80++ {
                    v1 = (v1 / (((v1 * v1)) | 1))
                    _ = fi80
                }
            }
            _ = fi65
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v50:", v50)
    fmt.Println("v51:", v51)
    fmt.Println("v57:", v57)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v50
    _ = v51
    _ = v53
    _ = v57
    _ = v58
}
