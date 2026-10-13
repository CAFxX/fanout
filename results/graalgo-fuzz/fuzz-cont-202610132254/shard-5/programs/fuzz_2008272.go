package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8; F2 int16 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint16; F1 uint16; F2 int16 }
type GenS4 struct { F0 string; F1 int64; F2 uint64 }
type GenS5 struct { F0 int16; F1 int64; F2 uint64 }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 int32 }
type GenS8 struct { F0 int32 }
type GenS9 struct { F0 string; F1 uint64 }
func genFn0() bool {
    var v4 string = ("a" + (("\n\t" + "foo bar") + "a"))
    if false {
        v5 := []int{-890}
        _ = v5
    }
    var v6 int8 = 127
    v7 := GenS1{F0: 221, F1: 128, F2: 1}
    _ = v4
    _ = v6
    _ = v7
    return true
}

func main() {
    v1 := make(map[string]string)
    var v2 uint8 = 111
    var v3 bool = (((^(v2)) < v2) || (-32768 >= 1))
    v2 = (v2 % (((v2 & (v2 >> ((199) & 63)))) | 1))
    fmt.Println(v2, v3)
    v8 := []int{-1, 4294967295, -136}
    v3 = true
    var v9 int32 = -1
    var v10 float64 = ((((1e10 * 0.1) - 3.14159) - 0.1) / 0.1)
    if v3 {
        if (true && ((v10 + ((((v10 - 1e10) - float64(((0.1 + ((0.0 * 1.5) * 1.5)) / ((0.1 + (3.14159 - 1.5)) * 0.1)))) - float64(-2.25)) - v10)) >= 3.14159)) {
            switch ("\n\t" + "hello") {
            case "hello":
                var v11 int64 = 1
                _ = v11
            case "a":
                fmt.Println(v3, v9, v2)
            default:
                switch "foo bar" {
                case "hello":
                    fmt.Println(v10, v2, v2)
                case "foo bar":
                    switch 32768 {
                    case 271:
                        v9 = 1
                    default:
                        var v12 int32 = 1
                        _ = v12
                    }
                default:
                    if (255 > 2) {
                        switch 128 {
                        case 501:
                            for fi13 := 0; fi13 < 19; fi13++ {
                                if false {
                                    v8 = append(v8, (^(fi13)))
                                    var v14 int8 = -128
                                    _ = v14
                                } else {
                                    var v15 string = ""
                                    if (32767 <= 109) {
                                        var v16 bool = (!(-109 > -94))
                                        _ = v16
                                    } else {
                                        var v17 int = int(v2)
                                        _ = v17
                                    }
                                    var v18 int32 = (v9 + 290)
                                    _ = v15
                                    _ = v18
                                }
                                if (2 > 255) {
                                    for fi19 := 0; fi19 < 11; fi19++ {
                                        switch ("" + "\n\t") {
                                        case "hello":
                                            v20 := GenS1{F0: 127, F1: v2, F2: int16(v2)}
                                            _ = v20
                                        case "foo bar":
                                            fmt.Println(fi13, v2)
                                        default:
                                            v8 = append(v8, fi19)
                                        }
                                        v21 := GenS2{F0: (583 >= 2)}
                                        _ = fi19
                                        _ = v21
                                    }
                                    if v3 {
                                        v22 := GenS2{F0: (12 >= 891)}
                                        switch "foo bar" {
                                        case "\n\t":
                                            for fi23 := 0; fi23 < 11; fi23++ {
                                                fmt.Println(v3)
                                                _ = fi23
                                            }
                                        default:
                                            v24 := GenS2{F0: (0 > 1000)}
                                            _ = v24
                                        }
                                        _ = v22
                                    }
                                } else {
                                    v25 := GenS3{F0: 42, F1: 492, F2: 42}
                                    _ = v25
                                }
                                _ = fi13
                            }
                        }
                    } else {
                        if (2 < 42) {
                            var v26 uint16 = 0
                            _ = v26
                        }
                        var v27 int8 = 117
                        _ = v27
                    }
                }
            }
        } else {
            var v28 string = ("\n\t" + ("\n\t" + ("a" + (("\n\t" + "hello") + "foo bar"))))
            if (!(!((!v3) || v3))) {
                for fi29 := 0; fi29 < 12; fi29++ {
                    v3 = (v28 != "a")
                    _ = fi29
                }
                v30 := GenS4{F0: ((v28 + ("\n\t" + (v28 + ("" + ("foo bar" + ("foo bar" + v28)))))) + v28), F1: -958, F2: 0}
                for fi31 := 0; fi31 < 15; fi31++ {
                    if v3 {
                        switch fi31 {
                        case 9223372036854775807:
                            for fi32 := 0; fi32 < 14; fi32++ {
                                if v3 {
                                    v28 = v28
                                    v10 += float64(3.14159)
                                } else {
                                    v28 = v28
                                    switch 763 {
                                    case 2147483647:
                                        switch fi31 {
                                        case 514:
                                            var v33 int16 = 525
                                            _ = v33
                                        }
                                    default:
                                        v9 = (v9 / ((v9) | 1))
                                    }
                                }
                                v9 -= v9
                                switch fi31 {
                                case 127:
                                    v3 = v3
                                }
                                _ = fi32
                            }
                        case -798:
                            switch v28 {
                            case "a":
                                var v34 int8 = 0
                                _ = v34
                            }
                        case 42:
                            var v35 float64 = (v10 - (3.14159 * v10))
                            _ = v35
                        default:
                            fmt.Println(fi31, v2, v10)
                        }
                        v36 := GenS2{F0: v3}
                        _ = v36
                    } else {
                        v1[v28] = ((v28 + ("a" + ("\n\t" + (("foo bar" + v28) + "hello")))) + (v28 + ""))
                        var v37 uint8 = (^(v2))
                        _ = v37
                    }
                    _ = len(v1)
                    v2 = (v2 / ((v2) | 1))
                    _ = fi31
                }
                _ = v30
            } else {
                fmt.Println(v9, v3, v9)
                for fi38 := 0; fi38 < 7; fi38++ {
                    var v39 bool = v3
                    fmt.Println(v39, v9)
                    _ = len(v1)
                    _ = fi38
                    _ = v39
                }
                var v40 uint64 = 4294967295
                _ = v40
            }
            if true {
                for fi41 := 0; fi41 < 3; fi41++ {
                    if (255 <= (v2 + v2)) {
                        v42 := GenS5{F0: 1, F1: 0, F2: 1}
                        for fi43 := 0; fi43 < 12; fi43++ {
                            if v3 {
                                v3 = (4294967295 != 32767)
                                var v44 float64 = float64((float32(float64(((0.0 / 1.5) * 0.0))) - float32((1.5 + v10))))
                                fmt.Println(v10)
                                _ = v44
                            } else {
                                var v45 int = (-(fi43))
                                var v46 int64 = -589
                                fmt.Println(v2, v3)
                                _ = v45
                                _ = v46
                            }
                            if ((true && false) && (!(!false))) {
                                if (fi43 >= fi43) {
                                    if len(v8) >= 2 {
                                        v8 = v8[1:len(v8)]
                                    }
                                }
                                v47 := GenS3{F0: 32767, F1: 805, F2: -1}
                                _ = v47
                            } else {
                                v48 := GenS4{F0: v28, F1: -9223372036854775808, F2: 18446744073709551615}
                                fmt.Println(v9, fi43, v28)
                                if v3 {
                                    var v49 float32 = (0.1 / 1e10)
                                    switch -199 {
                                    case -302:
                                        v28 = "a"
                                    case 32768:
                                        switch (v28 + "hello") {
                                        case "":
                                            var v50 uint64 = 656
                                            v28 = v28
                                            _ = v50
                                        case "foo bar":
                                            v10 += v10
                                        }
                                    case -94:
                                        switch v28 {
                                        case "a":
                                            if (0 == -42) {
                                                v51 := GenS5{F0: 133, F1: int64(v2), F2: 705}
                                                _ = v51
                                            } else {
                                                if false {
                                                    v2 -= v2
                                                    v9 = (v9 & 0)
                                                } else {
                                                    v28 = v28
                                                }
                                            }
                                        case "\n\t":
                                            var v52 float64 = v10
                                            _ = v52
                                        }
                                    default:
                                        if (!(!(!v3))) {
                                            v49 -= float32(3.14159)
                                            if true {
                                                fmt.Println(v49)
                                                v3 = true
                                                v53 := GenS1{F0: (^(v2)), F1: v2, F2: int16(v2)}
                                                _ = v53
                                            } else {
                                                v54 := GenS2{F0: (v3 && (v49 > float32((v10 / v10))))}
                                                _ = v54
                                            }
                                            v3 = true
                                            var v55 int16 = 201
                                            _ = v55
                                        } else {
                                            fmt.Println(v28, v3)
                                        }
                                    }
                                    switch "\n\t" {
                                    case "hello":
                                        fmt.Println(v2)
                                    default:
                                        var v56 string = v28
                                        _ = v56
                                    }
                                    _ = v49
                                }
                                v57 := GenS3{F0: 42, F1: 453, F2: 1}
                                var v58 uint8 = v2
                                _ = v48
                                _ = v57
                                _ = v58
                            }
                            if (2 != 4294967295) {
                                v59 := GenS6{F0: (!(v3 || v3))}
                                _ = v59
                            }
                            if v3 {
                                v60 := GenS2{F0: true}
                                var v61 uint32 = 42
                                _ = v60
                                _ = v61
                            }
                            v28 = v28
                            v62 := GenS7{F0: v9}
                            _ = fi43
                            _ = v62
                        }
                        _ = v42
                    } else {
                        v9 -= 0
                        if (127 == -1) {
                            if v3 {
                                if (!(0 > -778)) {
                                    for fi63 := 0; fi63 < 8; fi63++ {
                                        switch fi63 {
                                        case -2147483648:
                                            v64 := GenS8{F0: 218}
                                            _ = v64
                                        case -138:
                                            var v65 uint64 = 1
                                            _ = v65
                                        case 2147483647:
                                            v66 := GenS5{F0: 42, F1: 9223372036854775807, F2: 80}
                                            _ = v66
                                        }
                                        var v67 string = v28
                                        _ = fi63
                                        _ = v67
                                    }
                                } else {
                                    if len(v8) >= 2 {
                                        v8 = v8[1:len(v8)]
                                    }
                                    if v3 {
                                        var v68 bool = v3
                                        v9 -= (v9 / ((-2147483648) | 1))
                                        for fi69 := 0; fi69 < 11; fi69++ {
                                            v3 = v68
                                            _ = fi69
                                        }
                                        _ = v68
                                    } else {
                                        var v70 uint16 = 370
                                        fmt.Println(v70)
                                        var v71 int8 = 0
                                        _ = v70
                                        _ = v71
                                    }
                                    v3 = ((fi41 == fi41) && true)
                                }
                                v2 += 0
                            } else {
                                v72 := GenS9{F0: "\n\t", F1: 99}
                                _ = v72
                            }
                        }
                    }
                    _ = fi41
                }
                for fi73 := 0; fi73 < 1; fi73++ {
                    var v74 uint16 = 431
                    var v75 int = fi73
                    var v76 uint = 235
                    _ = fi73
                    _ = v74
                    _ = v75
                    _ = v76
                }
            } else {
                v77 := GenS2{F0: (1.5 > float32(float64(1e10)))}
                v9 = (v9 << ((1) & 63))
                _ = v77
            }
            _ = v28
        }
    } else {
        fmt.Println(v3)
        v1[("a" + "\n\t")] = "foo bar"
    }
    v8 = append(v8, -863)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v8:", len(v8), v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v8
    _ = v9
}
