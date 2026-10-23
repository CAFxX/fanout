package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 bool; F1 uint64; F2 int8 }
type GenS3 struct { F0 uint8; F1 uint }
type GenS4 struct { F0 int64; F1 uint; F2 int64 }
type GenS5 struct { F0 bool; F1 bool }
func genFn0(p0 int, p1 uint16) uint8 {
    v6 := []int{(p0 << ((202) & 63))}
    var v7 uint = 255
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 197
}

func main() {
    v1 := make(map[int]int)
    var v2 uint16 = 2
    var v3 bool = false
    var v4 int8 = 1
    v5 := []uint32{}
    v4 = 1
    v2 -= (v2 & 982)
    var v8 int64 = 706
    var v9 int64 = v8
    v3 = true
    fmt.Println(v8, v8)
    var v10 int32 = 42
    if len(v5) > 0 {
        v5[0] = 249
    }
    v4 -= v4
    if v3 {
        v11 := GenS1{F0: 0}
        var v12 int32 = v10
        _ = v11
        _ = v12
    } else {
        v3 = v3
    }
    if (true || v3) {
        if (2147483647 != 736) {
            v4 += 127
        }
        var v13 bool = v3
        _ = v13
    } else {
        switch "hello" {
        case "hello":
            var v14 int64 = 42
            _ = v14
        case "foo bar":
            fmt.Println(v3, v4, v9)
        }
    }
    var v15 uint16 = v2
    fmt.Println(v3, v2, v3)
    if v3 {
        for fi16 := 0; fi16 < 3; fi16++ {
            var v17 float32 = (0.0 / -2.25)
            var v18 uint32 = 381
            v2 = 65535
            _ = fi16
            _ = v17
            _ = v18
        }
    } else {
        var v19 int16 = 1
        var v20 bool = ((0.1 * (-2.25 * (((-2.25 * ((float64(float32(((0.1 - 1e10) + (-2.25 * (-2.25 + 3.14159))))) - 0.0) + 0.1)) / 1e10) * -2.25))) == float64(1e10))
        for fi21 := 0; fi21 < 8; fi21++ {
            switch "foo bar" {
            case "a":
                v20 = ((v4 << ((492) & 63)) >= 74)
            case "foo bar":
                v10 = -2147483648
            }
            switch "a" {
            case "foo bar":
                if v20 {
                    if v20 {
                        if v20 {
                            v15 += (v2 ^ 255)
                            for fi22 := 0; fi22 < 3; fi22++ {
                                var v23 bool = (v3 && v20)
                                v9 -= (^(v8))
                                _ = fi22
                                _ = v23
                            }
                            if v20 {
                                var v24 int16 = -1
                                var v25 uint = 1
                                _ = v24
                                _ = v25
                            }
                        }
                    }
                    for fi26 := 0; fi26 < 11; fi26++ {
                        if (!(v3 && false)) {
                            if v3 {
                                v19 -= v19
                                var v27 string = "foo bar"
                                _ = v27
                            }
                            delete(v1, fi21)
                            switch "a" {
                            case "":
                                if (v10 != v10) {
                                    switch fi26 {
                                    case -845:
                                        v3 = false
                                    case 0:
                                        if false {
                                            fmt.Println(fi21)
                                            if false {
                                                if (!v20) {
                                                    switch ("foo bar" + "") {
                                                    case "":
                                                        v8 = (^(v9))
                                                    }
                                                    var v28 int = fi21
                                                    var v29 int64 = v9
                                                    _ = v28
                                                    _ = v29
                                                }
                                                switch "foo bar" {
                                                case "a":
                                                    var v30 uint16 = 255
                                                    _ = v30
                                                case "foo bar":
                                                    var v31 float32 = ((((0.0 * ((1e10 / 3.14159) * 0.0)) + 0.0) + 3.14159) / ((1.5 / 0.1) - 0.1))
                                                    _ = v31
                                                }
                                                var v32 float32 = (1e10 + -2.25)
                                                _ = v32
                                            }
                                        }
                                    default:
                                        switch "\n\t" {
                                        case "\n\t":
                                            v33 := GenS1{F0: 0}
                                            _ = v33
                                        case "hello":
                                            v9 = v8
                                        case "foo bar":
                                            var v34 int16 = v19
                                            _ = v34
                                        }
                                    }
                                    v19 += (v19 * v19)
                                    v35 := GenS2{F0: (v19 > (v19 * (v19 & 32767))), F1: 0, F2: int8(v8)}
                                    _ = v35
                                } else {
                                    var v36 string = "hello"
                                    _ = len(v5)
                                    _ = v36
                                }
                            default:
                                switch (fi21 * (fi21 | (fi26 * -1))) {
                                case 9223372036854775807:
                                    v2 = (v15 / (((v15 + 255)) | 1))
                                case 797:
                                    if ("a" != "a") {
                                        if ((v15 & v15) < 65535) {
                                            if v3 {
                                                var v37 uint64 = 4294967295
                                                if v3 {
                                                    v38 := GenS1{F0: 797}
                                                    _ = v38
                                                } else {
                                                    var v39 uint32 = 910
                                                    v40 := GenS2{F0: v3, F1: (v37 % (((^(v37))) | 1)), F2: v4}
                                                    if (((!(2147483647 >= 955)) || v20) && false) {
                                                        if len(v5) > 0 {
                                                            v5[0] = (v39 | 4294967295)
                                                        }
                                                    } else {
                                                        if len(v5) > 0 {
                                                            v5[0] = (v39 | (^(v39)))
                                                        }
                                                        v41 := GenS2{F0: false, F1: v37, F2: -1}
                                                        _ = v41
                                                    }
                                                    _ = v39
                                                    _ = v40
                                                }
                                                fmt.Println(v8)
                                                _ = v37
                                            }
                                            var v42 int8 = int8(v15)
                                            v43 := GenS2{F0: v20, F1: 42, F2: v42}
                                            _ = v42
                                            _ = v43
                                        }
                                    } else {
                                        switch (fi21 & fi21) {
                                        case 163:
                                            switch (fi21 * fi21) {
                                            case -669:
                                                v20 = (0.1 < 1e10)
                                            case 32767:
                                                v15 = 255
                                            }
                                        case 284:
                                            v44 := GenS2{F0: (v20 || (!(v20 || v3))), F1: 272, F2: int8(v2)}
                                            _ = v44
                                        }
                                    }
                                default:
                                    var v45 bool = (((true || (float64((3.14159 - (float32(1.5) / 1.0))) > ((0.0 / 1.0) + 3.14159))) && false) || (v3 && (v3 || v20)))
                                    _ = v45
                                }
                            }
                            var v46 uint8 = 255
                            _ = v46
                        }
                        _ = fi26
                    }
                    v10 -= v10
                } else {
                    fmt.Println(v10, v10, fi21)
                    v47 := GenS1{F0: uint64(v15)}
                    var v48 uint = 18446744073709551615
                    _ = v47
                    _ = v48
                }
            case "hello":
                if false {
                    v4 -= int8(v10)
                    for fi49 := 0; fi49 < 20; fi49++ {
                        delete(v1, fi21)
                        v9 += -994
                        fmt.Println(v19)
                        _ = fi49
                    }
                } else {
                    v15 = v2
                    var v50 uint = 2
                    var v51 uint32 = 2147483647
                    _ = v50
                    _ = v51
                }
            }
            _ = fi21
        }
        _ = v19
        _ = v20
    }
    for fi52 := 0; fi52 < 11; fi52++ {
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        if len(v5) > 0 {
            v5[0] = 168
        }
        switch (("foo bar" + "hello") + "hello") {
        case "a":
            if v3 {
                switch (^(fi52)) {
                case 2147483647:
                    if v3 {
                        switch fi52 {
                        case 2147483647:
                            for fi53 := 0; fi53 < 10; fi53++ {
                                var v54 uint32 = 284
                                if true {
                                    v15 = 42
                                    fmt.Println(v15, v9)
                                } else {
                                    v5 = append(v5, v54)
                                    v55 := GenS3{F0: genFn0((fi52 ^ (fi53 + fi53)), uint16(v54)), F1: 1}
                                    fmt.Println(v3)
                                    _ = v55
                                }
                                _ = fi53
                                _ = v54
                            }
                        case -519:
                            v10 = (v10 & (v10 % (((v10 % ((v10) | 1))) | 1)))
                        }
                        if (v3 && v3) {
                            _ = len(v5)
                        } else {
                            var v56 string = "\n\t"
                            if (0 >= 53) {
                                v2 += (^(v2))
                            } else {
                                for fi57 := 0; fi57 < 16; fi57++ {
                                    var v58 uint = 447
                                    var v59 int16 = -32768
                                    var v60 uint16 = v2
                                    if (v56 >= "a") {
                                        var v61 float32 = (0.1 / ((1e10 - 1.5) * (0.0 + 3.14159)))
                                        if (42 < 18446744073709551615) {
                                            v61 = float32(-2.25)
                                            v3 = (false || v3)
                                            delete(v1, fi57)
                                        }
                                        delete(v1, (fi57 / (((fi52 & (^(fi52)))) | 1)))
                                        v62 := GenS4{F0: 270, F1: v58, F2: v8}
                                        _ = v61
                                        _ = v62
                                    } else {
                                        var v63 uint16 = v15
                                        var v64 uint8 = 2
                                        v65 := GenS5{F0: ((!(!v3)) && true), F1: false}
                                        if v3 {
                                            var v66 uint8 = v64
                                            _ = v66
                                        }
                                        _ = v63
                                        _ = v64
                                        _ = v65
                                    }
                                    v67 := GenS5{F0: v3, F1: v3}
                                    _ = fi57
                                    _ = v58
                                    _ = v59
                                    _ = v60
                                    _ = v67
                                }
                            }
                            _ = v56
                        }
                        switch ("" + "\n\t") {
                        case "foo bar":
                            var v68 uint16 = v2
                            _ = v68
                        case "\n\t":
                            v69 := v1[(fi52 / ((42) | 1))]
                            _ = v69
                        case "hello":
                            var v70 int32 = (v10 ^ -538)
                            _ = v70
                        default:
                            for fi71 := 0; fi71 < 4; fi71++ {
                                switch "a" {
                                case "":
                                    if len(v5) > 0 {
                                        v5[0] = 1
                                    }
                                default:
                                    var v72 string = ""
                                    _ = v72
                                }
                                fmt.Println(fi52, v10)
                                switch fi52 {
                                case -339:
                                    v73 := GenS2{F0: v3, F1: 1, F2: v4}
                                    _ = v73
                                case 32768:
                                    fmt.Println(v15, fi71)
                                }
                                _ = fi71
                            }
                        }
                    }
                }
            } else {
                var v74 uint = 815
                _ = v74
            }
        case "\n\t":
            v75 := GenS4{F0: -9223372036854775808, F1: 2, F2: v8}
            _ = v75
        case "hello":
            if v3 {
                v8 += (v9 >> ((835) & 63))
                v10 += v10
            }
        }
        _ = fi52
    }
    switch int(v4) {
    case 32767:
        fmt.Println(v3, v3, v4)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
