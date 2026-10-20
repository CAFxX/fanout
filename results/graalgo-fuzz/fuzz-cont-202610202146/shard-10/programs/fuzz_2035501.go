package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64; F2 int16 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int16; F1 bool }
type GenS4 struct { F0 int; F1 string }
type GenS5 struct { F0 uint; F1 int }
type GenS6 struct { F0 int64; F1 uint32; F2 string }
type GenS7 struct { F0 int16; F1 uint32; F2 int32 }
type GenS8 struct { F0 uint32 }

func main() {
    var v1 int8 = -1
    v2 := make(map[string]int)
    v3 := []int{32768, 42}
    var v4 int16 = 1
    var v5 float64 = float64(3.14159)
    if true {
        if true {
            switch -892 {
            case 127:
                if (0.0 > 0.0) {
                    v6 := GenS1{F0: 2147483648, F1: 42, F2: v4}
                    v4 -= (v4 & v4)
                    delete(v2, "foo bar")
                    _ = v6
                } else {
                    for fi7 := 0; fi7 < 15; fi7++ {
                        delete(v2, "")
                        v3 = append(v3, fi7)
                        if false {
                            v5 = v5
                            v8 := GenS1{F0: 2147483648, F1: 4294967295, F2: v4}
                            for fi9 := 0; fi9 < 5; fi9++ {
                                if true {
                                    var v10 uint16 = 65535
                                    var v11 uint8 = 1
                                    _ = v10
                                    _ = v11
                                }
                                _ = fi9
                            }
                            _ = v8
                        }
                        _ = fi7
                    }
                    for fi12 := 0; fi12 < 7; fi12++ {
                        var v13 string = ("foo bar" + (("foo bar" + "\n\t") + "hello"))
                        if false {
                            for fi14 := 0; fi14 < 9; fi14++ {
                                delete(v2, "\n\t")
                                var v15 int64 = -9223372036854775808
                                _ = fi14
                                _ = v15
                            }
                            v5 = 1e10
                        }
                        v16 := GenS2{F0: v1}
                        _ = fi12
                        _ = v13
                        _ = v16
                    }
                    var v17 int8 = 0
                    _ = v17
                }
            }
            _ = len(v3)
            var v18 uint8 = 214
            _ = v18
        }
        if true {
            var v19 uint32 = 2147483647
            v5 = float64(3.14159)
            _ = v19
        } else {
            if false {
                var v20 uint32 = 642
                v20 = (v20 + (v20 & v20))
                for fi21 := 0; fi21 < 4; fi21++ {
                    if (false || false) {
                        for fi22 := 0; fi22 < 1; fi22++ {
                            v3 = append(v3, (fi21 ^ (^(fi22))))
                            v5 -= ((v5 / ((v5 - v5) + v5)) * 0.0)
                            _ = fi22
                        }
                    } else {
                        for fi23 := 0; fi23 < 20; fi23++ {
                            if len(v3) > 0 {
                                v3[0] = (fi21 << ((588) & 63))
                            }
                            v2["hello"] = 54
                            _ = fi23
                        }
                    }
                    fmt.Println(v4, v4, v1)
                    fmt.Println(fi21)
                    _ = fi21
                }
                _ = v20
            }
        }
    } else {
        var v24 int8 = 1
        _ = v24
    }
    v5 = v5
    fmt.Println(v1, v4, v5)
    v4 = v4
    v25 := GenS2{F0: (v1 / ((v1) | 1))}
    var v26 uint = 255
    v5 -= ((v5 * 1e10) / 1.0)
    for fi27 := 0; fi27 < 10; fi27++ {
        v28 := GenS3{F0: int16(fi27), F1: (true || (!false))}
        if false {
            switch fi27 {
            case 128:
                switch fi27 {
                case -523:
                    var v29 float32 = 3.14159
                    _ = v29
                }
            case 0:
                switch (-(fi27)) {
                case 793:
                    var v30 string = ("hello" + "\n\t")
                    _ = v30
                case -9223372036854775808:
                    if len(v3) > 0 {
                        v3[0] = (-(fi27))
                    }
                case 127:
                    v5 -= ((v5 / (0.0 + -2.25)) / v5)
                default:
                    if true {
                        var v31 uint64 = 232
                        var v32 int64 = 42
                        v3 = append(v3, (fi27 << (((v26 ^ v26)) & 63)))
                        _ = v31
                        _ = v32
                    } else {
                        if false {
                            var v33 int = 4294967295
                            v34 := GenS4{F0: (^(v33)), F1: "foo bar"}
                            _ = v33
                            _ = v34
                        } else {
                            if false {
                                var v35 int8 = (v1 | v1)
                                _ = v35
                            } else {
                                for fi36 := 0; fi36 < 17; fi36++ {
                                    if true {
                                        v5 -= (v5 * 1e10)
                                        v4 = (-(v4))
                                        v4 -= (v4 >> (((v26 + v26)) & 63))
                                    } else {
                                        v1 -= (v1 % ((v1) | 1))
                                    }
                                    v3 = append(v3, 9223372036854775807)
                                    _ = len(v3)
                                    v37 := GenS3{F0: v4, F1: (float64(3.14159) >= 1e10)}
                                    var v38 bool = ((255 > 128) || (fi36 < fi36))
                                    var v39 uint16 = 255
                                    _ = fi36
                                    _ = v37
                                    _ = v38
                                    _ = v39
                                }
                                v3 = append(v3, (fi27 ^ fi27))
                            }
                            if true {
                                var v40 uint32 = 317
                                if ((!(1 != 128)) || (("" + "") == "hello")) {
                                    switch "\n\t" {
                                    case "":
                                        for fi41 := 0; fi41 < 17; fi41++ {
                                            fmt.Println(fi41, v40)
                                            switch fi27 {
                                            case 4294967295:
                                                fmt.Println(v1, v4)
                                            case -541:
                                                fmt.Println(v5, v26)
                                            case 359:
                                                v42 := v2["hello"]
                                                _ = v42
                                            }
                                            _ = fi41
                                        }
                                    case "hello":
                                        v40 = 4294967295
                                    default:
                                        for fi43 := 0; fi43 < 1; fi43++ {
                                            var v44 string = "a"
                                            fmt.Println(v40)
                                            _ = fi43
                                            _ = v44
                                        }
                                    }
                                    delete(v2, ((("hello" + "") + "") + "foo bar"))
                                    for fi45 := 0; fi45 < 2; fi45++ {
                                        if (uint16(v40) == 437) {
                                            fmt.Println(v26)
                                        }
                                        if false {
                                            var v46 uint = uint(v40)
                                            _ = v46
                                        } else {
                                            v1 = 42
                                            if len(v3) >= 2 {
                                                v3 = v3[1:len(v3)]
                                            }
                                        }
                                        if (4294967295 < 0) {
                                            v47 := GenS5{F0: (v26 % ((v26) | 1)), F1: (fi45 << ((v26) & 63))}
                                            delete(v2, ("foo bar" + ("hello" + ("hello" + "\n\t"))))
                                            _ = len(v2)
                                            v2[""] = fi27
                                            _ = v47
                                        } else {
                                            if (false || (!(-9223372036854775808 >= 0))) {
                                                v40 += v40
                                            }
                                            var v48 bool = (true && true)
                                            _ = v48
                                        }
                                        _ = fi45
                                    }
                                }
                                var v49 int32 = int32(v1)
                                _ = v40
                                _ = v49
                            }
                            for fi50 := 0; fi50 < 11; fi50++ {
                                switch ("\n\t" + "hello") {
                                case "\n\t":
                                    if true {
                                        if ((false && true) && ((v4 & (v4 % ((32767) | 1))) == v4)) {
                                            var v51 uint32 = 721
                                            v1 = -108
                                            _ = v51
                                        }
                                        switch fi50 {
                                        case -1:
                                            if true {
                                                v52 := GenS6{F0: 0, F1: 11, F2: "hello"}
                                                _ = v52
                                            }
                                            var v53 int16 = 708
                                            _ = v53
                                        case 4294967295:
                                            if true {
                                                v1 += (v1 >> (((v26 | (v26 + v26))) & 63))
                                                v54 := GenS4{F0: fi50, F1: ("" + "\n\t")}
                                                fmt.Println(v26, fi27)
                                                _ = v54
                                            }
                                            v2[("hello" + "a")] = -503
                                        case 0:
                                            var v55 string = "\n\t"
                                            _ = v55
                                        }
                                    }
                                case "foo bar":
                                    v56 := GenS7{F0: -696, F1: 4294967295, F2: -735}
                                    _ = v56
                                case "":
                                    fmt.Println(v5, v26, v1)
                                }
                                _ = fi50
                            }
                        }
                    }
                }
            case 127:
                switch (fi27 ^ (fi27 + (fi27 & (fi27 << ((800) & 63))))) {
                case 527:
                    fmt.Println(v1)
                case 2147483647:
                    for fi57 := 0; fi57 < 3; fi57++ {
                        v1 -= int8(v4)
                        var v58 int64 = -9223372036854775808
                        _ = fi57
                        _ = v58
                    }
                default:
                    for fi59 := 0; fi59 < 17; fi59++ {
                        var v60 int = (fi27 - (fi59 % ((32768) | 1)))
                        if len(v3) > 0 {
                            v3[0] = -9223372036854775808
                        }
                        v1 += (v1 | (v1 + (v1 - (v1 << (((v26 + (v26 * (v26 | (v26 << (((v26 / ((126) | 1))) & 63)))))) & 63)))))
                        _ = fi59
                        _ = v60
                    }
                }
            }
        } else {
            fmt.Println(fi27, fi27)
        }
        _ = fi27
        _ = v28
    }
    var v61 int = 802
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v62 := GenS8{F0: 255}
    var v63 uint32 = 432
    v64 := GenS6{F0: 909, F1: v63, F2: ""}
    for fi65 := 0; fi65 < 18; fi65++ {
        v3 = append(v3, 0)
        _ = fi65
    }
    if true {
        switch "hello" {
        case "":
            _ = len(v2)
        case "hello":
            var v66 int = 840
            _ = v66
        default:
            fmt.Println(v1, v1)
        }
        _ = len(v3)
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v61:", v61)
    fmt.Println("v63:", v63)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v61
    _ = v62
    _ = v63
    _ = v64
}
