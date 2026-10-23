package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int8; F1 uint64 }
type GenS3 struct { F0 int8; F1 uint32; F2 string }
type GenS4 struct { F0 int8; F1 uint8; F2 int8 }
type GenS5 struct { F0 uint8; F1 bool; F2 uint32 }

func main() {
    v1 := make(map[uint32]int)
    var v2 uint = 4294967295
    v3 := make(map[int]bool)
    v2 = (^(v2))
    switch ("foo bar" + "") {
    case "\n\t":
        v2 -= v2
    case "hello":
        v4 := GenS1{F0: 18446744073709551615}
        _ = v4
    case "a":
        v5 := []int{127, 9223372036854775807, 158}
        _ = v5
    default:
        v2 += (v2 * (v2 >> ((v2) & 63)))
    }
    v2 = (v2 * v2)
    if false {
        var v6 float64 = 1.5
        var v7 uint16 = 32767
        _ = v6
        _ = v7
    } else {
        fmt.Println(v2)
    }
    v8 := GenS1{F0: 186}
    for fi9 := 0; fi9 < 2; fi9++ {
        v10 := []int{564, -644, fi9}
        v2 -= v2
        if (-882 <= 307) {
            for fi11 := 0; fi11 < 20; fi11++ {
                fmt.Println(fi9, fi11)
                var v12 uint8 = 250
                v2 = v2
                _ = fi11
                _ = v12
            }
            _ = len(v10)
            switch 4294967295 {
            case 32768:
                v2 -= (v2 << (((v2 ^ v2)) & 63))
            default:
                var v13 uint8 = 1
                _ = v13
            }
        }
        _ = fi9
        _ = v10
    }
    var v14 bool = true
    for fi15 := 0; fi15 < 9; fi15++ {
        v16 := GenS1{F0: 790}
        if v14 {
            v2 -= (v2 % (((v2 | 304)) | 1))
            if false {
                v2 = 1
                for fi17 := 0; fi17 < 6; fi17++ {
                    v18 := GenS1{F0: 2}
                    switch "foo bar" {
                    case "hello":
                        v19 := GenS1{F0: 22}
                        _ = v19
                    case "":
                        if false {
                            fmt.Println(fi15)
                            fmt.Println(fi17, v2)
                            var v20 int64 = -275
                            _ = v20
                        } else {
                            v21 := GenS1{F0: 42}
                            if v14 {
                                switch "a" {
                                case "":
                                    if ((v2 < (^(v2))) || (!(0 > 100))) {
                                        v2 -= v2
                                        v22 := GenS2{F0: -7, F1: 0}
                                        _ = v22
                                    }
                                case "\n\t":
                                    v2 += v2
                                }
                                var v23 int = fi15
                                _ = v23
                            }
                            v24 := GenS2{F0: -1, F1: 894}
                            _ = v21
                            _ = v24
                        }
                    case "a":
                        v14 = (v14 || ((false || v14) && false))
                        if ((!false) || v14) {
                            v14 = v14
                        }
                    }
                    var v25 int64 = -272
                    _ = fi17
                    _ = v18
                    _ = v25
                }
                v26 := GenS3{F0: 42, F1: 2147483647, F2: "hello"}
                _ = v26
            }
        } else {
            var v27 int8 = -63
            v2 -= 255
            if (-1 > -1) {
                switch -1 {
                case -1:
                    v28 := GenS2{F0: 1, F1: 2}
                    _ = v28
                case 1:
                    switch "hello" {
                    case "":
                        var v29 int8 = 86
                        _ = v29
                    case "a":
                        for fi30 := 0; fi30 < 7; fi30++ {
                            v31 := GenS3{F0: 1, F1: 2147483648, F2: "foo bar"}
                            v32 := GenS1{F0: 999}
                            switch "foo bar" {
                            case "":
                                v33 := []int{}
                                _ = v33
                            }
                            _ = fi30
                            _ = v31
                            _ = v32
                        }
                    default:
                        fmt.Println(v2, fi15)
                    }
                default:
                    var v34 int16 = -19
                    _ = v34
                }
                var v35 uint32 = 1
                _ = v35
            }
            _ = v27
        }
        _ = fi15
        _ = v16
    }
    v14 = (1 == 1)
    v36 := []int{}
    for fi37 := 0; fi37 < 8; fi37++ {
        switch fi37 {
        case 128:
            switch ("foo bar" + "") {
            case "foo bar":
                var v38 int64 = 9223372036854775807
                _ = v38
            }
        case 4294967295:
            switch (("foo bar" + "hello") + "") {
            case "foo bar":
                v39 := GenS2{F0: 127, F1: 634}
                _ = v39
            case "":
                for fi40 := 0; fi40 < 19; fi40++ {
                    switch (fi40 * fi40) {
                    case 128:
                        var v41 float32 = 0.0
                        _ = v41
                    case 0:
                        v14 = v14
                    default:
                        if len(v36) > 0 {
                            v36[0] = fi37
                        }
                    }
                    _ = fi40
                }
            default:
                v42 := GenS4{F0: 1, F1: 128, F2: -2}
                _ = v42
            }
        default:
            v14 = v14
        }
        v14 = v14
        for fi43 := 0; fi43 < 11; fi43++ {
            if len(v36) >= 2 {
                v36 = v36[1:len(v36)]
            }
            v2 = (v2 % ((v2) | 1))
            var v44 uint8 = 42
            if (false && true) {
                v36 = append(v36, 4294967295)
            } else {
                delete(v3, (fi37 + (fi37 ^ (fi37 ^ fi43))))
            }
            switch "hello" {
            case "\n\t":
                if false {
                    v44 += (v44 << ((v2) & 63))
                    if (0 != uint32(v2)) {
                        if false {
                            var v45 uint64 = 0
                            if len(v36) >= 2 {
                                v36 = v36[1:len(v36)]
                            }
                            _ = v45
                        } else {
                            if v14 {
                                v46 := GenS3{F0: -7, F1: 216, F2: ("hello" + ("foo bar" + "a"))}
                                _ = v46
                            } else {
                                var v47 int = (fi37 | (fi43 >> (((v2 / ((v2) | 1))) & 63)))
                                v48 := v3[v47]
                                _ = v47
                                _ = v48
                            }
                        }
                        if v14 {
                            var v49 uint = v2
                            _ = v49
                        } else {
                            fmt.Println(v14, v14, v14)
                            v36 = append(v36, (fi43 % ((fi37) | 1)))
                            v50 := GenS1{F0: 18446744073709551615}
                            _ = v50
                        }
                        if (!(2147483647 < 514)) {
                            if ((!(v2 <= (v2 & v2))) || ((v2 * (v2 / ((0) | 1))) >= (v2 / ((v2) | 1)))) {
                                if ((v44 & v44) <= (v44 | v44)) {
                                    if false {
                                        v51 := GenS4{F0: 0, F1: v44, F2: -1}
                                        var v52 uint32 = 18
                                        _ = v51
                                        _ = v52
                                    }
                                    var v53 float32 = 1e10
                                    v36 = append(v36, 977)
                                    _ = v53
                                } else {
                                    fmt.Println(v14, v14, fi37)
                                }
                            } else {
                                v14 = (float64(1.5) != -2.25)
                                var v54 float32 = 3.14159
                                var v55 int16 = 42
                                _ = v54
                                _ = v55
                            }
                            v2 -= (^(v2))
                            if (!(0.1 < (1.5 + 1.5))) {
                                fmt.Println(fi43, v2)
                            }
                        } else {
                            v3[(-(fi37))] = true
                            v14 = v14
                        }
                    } else {
                        var v56 int8 = 32
                        var v57 int64 = 1
                        v58 := GenS4{F0: (v56 + v56), F1: (v44 % ((2) | 1)), F2: (v56 ^ -44)}
                        _ = v56
                        _ = v57
                        _ = v58
                    }
                } else {
                    var v59 uint8 = 166
                    _ = v59
                }
            case "foo bar":
                if v14 {
                    var v60 int = (^(fi37))
                    _ = v60
                }
            default:
                v14 = true
            }
            _ = fi43
            _ = v44
        }
        _ = fi37
    }
    v2 += v2
    switch "hello" {
    case "foo bar":
        for fi61 := 0; fi61 < 1; fi61++ {
            if (!(-128 > 127)) {
                fmt.Println(v14, fi61, fi61)
                for fi62 := 0; fi62 < 12; fi62++ {
                    if v14 {
                        if true {
                            var v63 float32 = 3.14159
                            if len(v36) > 0 {
                                v36[0] = fi61
                            }
                            v14 = v14
                            var v64 int32 = 1
                            _ = v63
                            _ = v64
                        }
                        switch 32768 {
                        case 1:
                            v65 := GenS5{F0: 253, F1: v14, F2: 4294967295}
                            _ = v65
                        case 32768:
                            var v66 int64 = 9223372036854775807
                            _ = v66
                        case 42:
                            v2 += (^(v2))
                        }
                        delete(v1, 42)
                    }
                    v36 = append(v36, (fi62 & (fi61 + (-(fi61)))))
                    fmt.Println(v14, v14, fi61)
                    _ = fi62
                }
                v2 = v2
            } else {
                v67 := GenS4{F0: 0, F1: 100, F2: 99}
                switch "" {
                case "\n\t":
                    var v68 int32 = 42
                    _ = v68
                }
                _ = v67
            }
            _ = len(v1)
            _ = fi61
        }
    case "a":
        v14 = (!true)
    case "\n\t":
        switch 9223372036854775807 {
        case 0:
            switch "a" {
            case "hello":
                var v69 int8 = -72
                _ = v69
            case "a":
                fmt.Println(v14, v2)
            }
        }
    }
    switch (("a" + "hello") + "a") {
    case "\n\t":
        for fi70 := 0; fi70 < 18; fi70++ {
            _ = len(v3)
            _ = fi70
        }
    default:
        v71 := GenS3{F0: -88, F1: 42, F2: "a"}
        _ = v71
    }
    fmt.Println(v14, v2)
    v14 = (18446744073709551615 <= 42)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v36:", len(v36), v36)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v36
    _ = v8
}
