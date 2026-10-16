package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int; F1 int32 }
type GenS3 struct { F0 uint16; F1 uint16; F2 int64 }
type GenS4 struct { F0 int64; F1 int8 }
type GenS5 struct { F0 int32; F1 uint32; F2 int32 }
type GenS6 struct { F0 uint64; F1 uint64; F2 int16 }
type GenS7 struct { F0 bool }

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[int]bool)
    v3 := []string{((("" + "foo bar") + "hello") + "a")}
    v4 := make(map[uint32]string)
    var v5 int8 = 0
    var v6 int32 = 1
    if (2147483647 >= 2) {
        switch "foo bar" {
        case "a":
            if true {
                delete(v1, ("" + "\n\t"))
            } else {
                if true {
                    v7 := GenS1{F0: 26}
                    _ = v7
                } else {
                    v6 -= (v6 * v6)
                }
                var v8 float32 = 0.0
                _ = v8
            }
        case "foo bar":
            if (!(236 <= 0)) {
                for fi9 := 0; fi9 < 2; fi9++ {
                    v5 += (v5 << ((2) & 63))
                    if (65535 <= 521) {
                        if true {
                            v4[42] = "\n\t"
                            v5 -= v5
                        }
                        if true {
                            if false {
                                v6 += (v6 * v6)
                                var v10 uint16 = 255
                                v6 = v6
                                _ = v10
                            }
                            for fi11 := 0; fi11 < 9; fi11++ {
                                var v12 uint8 = 127
                                _ = fi11
                                _ = v12
                            }
                        } else {
                            var v13 float32 = 1.5
                            if len(v3) > 0 {
                                v3[0] = "hello"
                            }
                            v6 = (v6 >> ((42) & 63))
                            _ = v13
                        }
                        var v14 int = (fi9 | (fi9 * int(v6)))
                        _ = v14
                    } else {
                        for fi15 := 0; fi15 < 13; fi15++ {
                            fmt.Println(fi9, fi15, fi9)
                            fmt.Println(fi9, fi9, fi15)
                            var v16 int = 32767
                            var v17 uint = 255
                            var v18 int8 = v5
                            if false {
                                var v19 uint32 = 2147483648
                                var v20 int32 = (-(v6))
                                _ = v19
                                _ = v20
                            } else {
                                if (!("a" == ("" + ("" + "\n\t")))) {
                                    v17 -= v17
                                    fmt.Println(v17, v16, v17)
                                }
                            }
                            _ = fi15
                            _ = v16
                            _ = v17
                            _ = v18
                        }
                        for fi21 := 0; fi21 < 20; fi21++ {
                            if false {
                                if len(v3) > 0 {
                                    v3[0] = "a"
                                }
                                var v22 float32 = ((0.0 * 1.5) / ((-2.25 / 0.1) + 1e10))
                                var v23 float64 = (float64((v22 / 1e10)) * 0.0)
                                _ = v22
                                _ = v23
                            } else {
                                var v24 uint64 = 1
                                var v25 string = "hello"
                                v2[fi9] = (104 <= uint8(v24))
                                _ = v24
                                _ = v25
                            }
                            _ = fi21
                        }
                    }
                    if (42 != 255) {
                        v5 = (-(v5))
                        for fi26 := 0; fi26 < 16; fi26++ {
                            if true {
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                if (!false) {
                                    v27 := GenS1{F0: 2}
                                    _ = v27
                                } else {
                                    var v28 int16 = -26
                                    v29 := v1["foo bar"]
                                    var v30 uint64 = 990
                                    _ = v28
                                    _ = v29
                                    _ = v30
                                }
                            } else {
                                v31 := GenS1{F0: 18446744073709551615}
                                _ = v31
                            }
                            _ = fi26
                        }
                        v6 += (^(v6))
                    } else {
                        if false {
                            var v32 uint16 = 32767
                            v6 = 1
                            _ = v32
                        }
                        v3 = append(v3, "hello")
                        fmt.Println(v6)
                    }
                    _ = fi9
                }
                switch ("a" + (((("hello" + "foo bar") + ("hello" + "\n\t")) + ("a" + ("a" + "foo bar"))) + ("hello" + "foo bar"))) {
                case "a":
                    if true {
                        for fi33 := 0; fi33 < 7; fi33++ {
                            if ((false || ((0 > 9223372036854775807) && true)) && false) {
                                var v34 uint16 = 0
                                for fi35 := 0; fi35 < 15; fi35++ {
                                    v3 = append(v3, "a")
                                    _ = fi35
                                }
                                _ = v34
                            } else {
                                for fi36 := 0; fi36 < 3; fi36++ {
                                    fmt.Println(fi36)
                                    v5 += v5
                                    _ = fi36
                                }
                                _ = len(v4)
                            }
                            if false {
                                fmt.Println(fi33, fi33)
                                v37 := GenS2{F0: fi33, F1: v6}
                                _ = v37
                            }
                            _ = fi33
                        }
                        fmt.Println(v5, v6)
                        if true {
                            for fi38 := 0; fi38 < 8; fi38++ {
                                switch fi38 {
                                case 4294967295:
                                    v6 = v6
                                case 2147483647:
                                    var v39 float32 = 0.0
                                    _ = v39
                                }
                                _ = fi38
                            }
                            if ((2 >= 2147483648) || ((4294967295 > 18446744073709551615) || false)) {
                                for fi40 := 0; fi40 < 20; fi40++ {
                                    v5 = 0
                                    v5 -= 78
                                    v5 -= v5
                                    _ = fi40
                                }
                            } else {
                                v5 -= (v5 / ((127) | 1))
                            }
                            v6 = -2147483648
                        }
                    } else {
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        var v41 bool = true
                        _ = v41
                    }
                case "\n\t":
                    v6 += 0
                case "foo bar":
                    if true {
                        v42 := GenS3{F0: 2, F1: 782, F2: 9223372036854775807}
                        if true {
                            switch "" {
                            case "":
                                v6 += -636
                            default:
                                if true {
                                    if false {
                                        v5 += (v5 - v5)
                                        var v43 float64 = -2.25
                                        _ = v43
                                    } else {
                                        v5 -= 1
                                        v5 -= v5
                                        for fi44 := 0; fi44 < 12; fi44++ {
                                            if true {
                                                fmt.Println(v5)
                                                var v45 int16 = -666
                                                v46 := GenS1{F0: 4294967295}
                                                _ = v45
                                                _ = v46
                                            }
                                            fmt.Println(v5)
                                            _ = fi44
                                        }
                                    }
                                    var v47 int16 = 0
                                    switch "" {
                                    case "foo bar":
                                        _ = len(v3)
                                    case "hello":
                                        _ = len(v1)
                                    default:
                                        fmt.Println(v47, v5, v5)
                                    }
                                    _ = v47
                                } else {
                                    v48 := v4[2]
                                    _ = v48
                                }
                            }
                        }
                        for fi49 := 0; fi49 < 10; fi49++ {
                            for fi50 := 0; fi50 < 8; fi50++ {
                                v5 = (^(v5))
                                switch fi49 {
                                case -1:
                                    v51 := GenS4{F0: 0, F1: (-(v5))}
                                    _ = v51
                                case 2147483647:
                                    var v52 bool = true
                                    _ = v52
                                case 32768:
                                    if ((-1 < -1) || true) {
                                        fmt.Println(fi50, v6)
                                        switch (fi50 & (fi50 ^ (fi50 & (fi50 / (((-(fi49))) | 1))))) {
                                        case 1:
                                            var v53 int8 = (v5 % (((v5 | 100)) | 1))
                                            _ = v53
                                        case 4:
                                            var v54 int16 = -32768
                                            _ = v54
                                        case -1:
                                            if (v6 <= 42) {
                                                var v55 uint64 = 381
                                                fmt.Println(v5)
                                                var v56 uint = 65535
                                                _ = v55
                                                _ = v56
                                            } else {
                                                v5 = v5
                                            }
                                        }
                                        var v57 uint = 313
                                        _ = v57
                                    }
                                default:
                                    fmt.Println(fi50)
                                }
                                _ = fi50
                            }
                            v58 := GenS5{F0: (v6 % ((-1) | 1)), F1: 65535, F2: v6}
                            _ = fi49
                            _ = v58
                        }
                        _ = v42
                    }
                }
                v59 := GenS2{F0: 42, F1: 1}
                _ = v59
            } else {
                v60 := GenS2{F0: -538, F1: v6}
                _ = v60
            }
        case "":
            var v61 uint8 = 189
            _ = v61
        }
    } else {
        v6 -= v6
        var v62 uint16 = 42
        _ = v62
    }
    v5 = -36
    if false {
        v6 -= v6
        for fi63 := 0; fi63 < 4; fi63++ {
            for fi64 := 0; fi64 < 20; fi64++ {
                v65 := v2[-9223372036854775808]
                switch ("a" + (("" + "foo bar") + "\n\t")) {
                case "\n\t":
                    v6 += (v6 | v6)
                case "hello":
                    var v66 int64 = 0
                    _ = v66
                case "a":
                    v5 = v5
                }
                var v67 uint32 = 77
                _ = fi64
                _ = v65
                _ = v67
            }
            switch ("" + "") {
            case "":
                if true {
                    fmt.Println(fi63, v5)
                    v68 := GenS6{F0: 374, F1: 113, F2: 795}
                    v69 := GenS1{F0: 66}
                    _ = v68
                    _ = v69
                } else {
                    var v70 uint = 378
                    _ = v70
                }
            case "a":
                if (2 == 2147483647) {
                    v5 = 1
                }
            }
            v6 -= v6
            _ = fi63
        }
        var v71 int16 = -231
        _ = v71
    }
    for fi72 := 0; fi72 < 10; fi72++ {
        v73 := GenS1{F0: 1}
        _ = fi72
        _ = v73
    }
    var v74 int32 = v6
    if true {
        v5 = v5
    } else {
        if (0 == (v6 + v74)) {
            switch "" {
            case "hello":
                v6 = (-(v6))
            }
            v5 -= (v5 >> ((42) & 63))
            for fi75 := 0; fi75 < 2; fi75++ {
                for fi76 := 0; fi76 < 10; fi76++ {
                    if ((v6 < (v6 / ((-1) | 1))) && true) {
                        var v77 int32 = v74
                        v3 = append(v3, "")
                        _ = v77
                    }
                    fmt.Println(fi75, fi76, v6)
                    if ((1 < 199) || true) {
                        var v78 uint64 = 0
                        v79 := GenS4{F0: 1, F1: v5}
                        var v80 uint32 = 0
                        _ = len(v3)
                        _ = v78
                        _ = v79
                        _ = v80
                    }
                    v81 := GenS7{F0: ((true && (v6 == (-(v74)))) && true)}
                    _ = fi76
                    _ = v81
                }
                fmt.Println(v74, fi75, v74)
                v5 += (v5 << ((42) & 63))
                _ = fi75
            }
        }
    }
    if true {
        var v82 int32 = v6
        _ = v82
    } else {
        var v83 int8 = (v5 << ((0) & 63))
        _ = v83
    }
    v6 = -324
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v74:", v74)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v74
}
