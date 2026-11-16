package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int16; F1 uint64 }
type GenS4 struct { F0 int16; F1 uint32; F2 int32 }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 bool; F1 int32 }
func genFn0() int16 {
    v6 := GenS1{F0: ""}
    var v7 int8 = 1
    _ = v6
    _ = v7
    return 825
}
func genFn1(p0 string, p1 uint8, p2 uint8) uint32 {
    switch -9223372036854775808 {
    case 144:
        v8 := GenS1{F0: "foo bar"}
        _ = v8
    case 32768:
        switch p0 {
        case "foo bar":
            switch 32768 {
            case 9223372036854775807:
                var v9 uint = 670
                _ = v9
            case 0:
                v10 := []int{-1, 32767, 0}
                _ = v10
            }
        default:
            var v11 float32 = (1e10 / 1e10)
            _ = v11
        }
    default:
        var v12 uint8 = 253
        _ = v12
    }
    if (false && ((0.1 / float32((((1e10 + 1.5) - 1.5) - (float64(-2.25) * 1e10)))) != 1.5)) {
        v13 := []int{}
        fmt.Println(p1)
        p1 += p2
        _ = v13
    }
    _ = p0
    _ = p1
    _ = p2
    return 493
}

func main() {
    var v1 int = -1
    v2 := make(map[string]bool)
    v3 := []string{"hello", (("hello" + "\n\t") + ("\n\t" + ("a" + ("hello" + "\n\t")))), ("hello" + "hello")}
    v4 := make(map[int]bool)
    v5 := make(map[int]uint32)
    fmt.Println(v1)
    for fi14 := 0; fi14 < 16; fi14++ {
        var v15 float64 = 0.0
        _ = fi14
        _ = v15
    }
    for fi16 := 0; fi16 < 4; fi16++ {
        if (!(0 > 338)) {
            var v17 int8 = -1
            fmt.Println(v1)
            delete(v5, v1)
            _ = v17
        } else {
            var v18 uint16 = 0
            _ = v18
        }
        var v19 string = ((("" + "a") + "foo bar") + "\n\t")
        _ = fi16
        _ = v19
    }
    v1 = v1
    var v20 string = "hello"
    switch (v1 << ((255) & 63)) {
    case 71:
        var v21 float64 = (3.14159 + 0.0)
        _ = v21
    default:
        v20 = (v20 + v20)
    }
    if (9223372036854775807 < 9223372036854775807) {
        v1 = v1
    } else {
        var v22 bool = true
        var v23 uint16 = 255
        fmt.Println(v22, v23, v22)
        _ = v22
        _ = v23
    }
    v2["a"] = (!true)
    v24 := v5[v1]
    var v25 int8 = 100
    var v26 float64 = 0.1
    v26 = ((-2.25 - (v26 - v26)) / v26)
    var v27 bool = ((v20 + ("a" + (v20 + ((v20 + "\n\t") + "")))) >= "hello")
    var v28 int8 = v25
    _ = len(v5)
    fmt.Println(v25)
    for fi29 := 0; fi29 < 4; fi29++ {
        switch fi29 {
        case 32768:
            delete(v4, (fi29 & -9223372036854775808))
        case -704:
            v30 := GenS1{F0: "foo bar"}
            _ = v30
        case 1:
            fmt.Println(v25, v26)
        default:
            v31 := GenS2{F0: 1}
            _ = v31
        }
        switch fi29 {
        case 0:
            if v27 {
                v32 := GenS3{F0: 32767, F1: 4294967295}
                switch (("foo bar" + v20) + (("\n\t" + (v20 + v20)) + "a")) {
                case "":
                    var v33 string = v20
                    _ = v33
                case "a":
                    v24 -= v24
                case "\n\t":
                    for fi34 := 0; fi34 < 19; fi34++ {
                        if v27 {
                            v28 = v28
                            var v35 uint32 = (v24 | v24)
                            var v36 int16 = 0
                            _ = v35
                            _ = v36
                        } else {
                            var v37 uint64 = 4294967295
                            var v38 float32 = 0.1
                            if (-306 > -383) {
                                var v39 uint32 = (^(v24))
                                v40 := GenS4{F0: -126, F1: v24, F2: 1}
                                _ = v39
                                _ = v40
                            } else {
                                switch (fi34 + 127) {
                                case -182:
                                    if false {
                                        var v41 uint32 = v24
                                        fmt.Println(v37, v25, fi29)
                                        _ = v41
                                    } else {
                                        var v42 string = "\n\t"
                                        switch (v20 + "a") {
                                        case "\n\t":
                                            switch -625 {
                                            case 4294967295:
                                                if v27 {
                                                    _ = len(v3)
                                                    _ = len(v4)
                                                } else {
                                                    v25 += 17
                                                }
                                                v37 = (v37 * 951)
                                            case 32767:
                                                var v43 uint16 = 502
                                                _ = v43
                                            default:
                                                v44 := GenS3{F0: genFn0(), F1: (v37 << ((814) & 63))}
                                                _ = v44
                                            }
                                        case "":
                                            var v45 bool = v27
                                            _ = v45
                                        case "hello":
                                            v38 += float32(3.14159)
                                            if (433 < 785) {
                                                v25 += (v25 * v28)
                                                var v46 float32 = 3.14159
                                                fmt.Println(v24, v38, v38)
                                                v47 := GenS5{F0: uint16(v24)}
                                                var v48 int = fi29
                                                _ = v46
                                                _ = v47
                                                _ = v48
                                            }
                                            fmt.Println(fi34, v20, v25)
                                            var v49 int16 = 42
                                            v42 = (((v42 + ((("\n\t" + (v42 + "a")) + "hello") + "\n\t")) + "a") + v20)
                                            _ = v49
                                        default:
                                            fmt.Println(v20, v24)
                                        }
                                        _ = v42
                                    }
                                case 32768:
                                    v27 = ((v1 | 42) > fi34)
                                default:
                                    var v50 int8 = (v25 >> ((uint(v37)) & 63))
                                    _ = v50
                                }
                                v38 = v38
                            }
                            _ = v37
                            _ = v38
                        }
                        var v51 uint = 255
                        _ = fi34
                        _ = v51
                    }
                default:
                    v26 -= v26
                }
                var v52 float64 = v26
                _ = v32
                _ = v52
            } else {
                if (!v27) {
                    if v27 {
                        v24 = (v24 - genFn1(v20, 2, 47))
                        v1 += fi29
                        v1 = fi29
                    } else {
                        switch (fi29 / (((fi29 & (fi29 * int(v28)))) | 1)) {
                        case -170:
                            switch int(v24) {
                            case -746:
                                for fi53 := 0; fi53 < 8; fi53++ {
                                    v27 = (0 != 523)
                                    if (v20 == (v20 + v20)) {
                                        var v54 uint64 = 1
                                        _ = v54
                                    } else {
                                        if (true || false) {
                                            fmt.Println(v26)
                                            var v55 int8 = v25
                                            v56 := GenS2{F0: 61}
                                            _ = v55
                                            _ = v56
                                        } else {
                                            if v27 {
                                                v57 := GenS3{F0: 42, F1: 18446744073709551615}
                                                if len(v3) >= 2 {
                                                    v3 = v3[1:len(v3)]
                                                }
                                                _ = v57
                                            } else {
                                                var v58 float64 = 1e10
                                                if true {
                                                    fmt.Println(v28)
                                                } else {
                                                    if v27 {
                                                        var v59 uint16 = 1
                                                        _ = v59
                                                    }
                                                }
                                                _ = v58
                                            }
                                            v25 += v25
                                            v26 -= float64((0.1 - 0.1))
                                        }
                                        switch "hello" {
                                        case "hello":
                                            var v60 uint = 4294967295
                                            v61 := GenS2{F0: 242}
                                            var v62 uint = (v60 << ((1) & 63))
                                            _ = v60
                                            _ = v61
                                            _ = v62
                                        }
                                        if v27 {
                                            var v63 int16 = 0
                                            delete(v5, (v1 + v1))
                                            _ = v63
                                        } else {
                                            v4[128] = (2147483647 >= 396)
                                            switch "foo bar" {
                                            case "":
                                                var v64 int64 = -252
                                                _ = v64
                                            case "\n\t":
                                                var v65 uint8 = 160
                                                _ = v65
                                            }
                                        }
                                    }
                                    _ = fi53
                                }
                            case 568:
                                fmt.Println(v26, fi29, v27)
                            }
                        default:
                            var v66 bool = true
                            _ = v66
                        }
                    }
                }
                for fi67 := 0; fi67 < 1; fi67++ {
                    v3 = append(v3, (v20 + ""))
                    fmt.Println(v1, v1, v26)
                    v68 := GenS5{F0: 65535}
                    if (-1 < 32767) {
                        if (v27 && (v27 && v27)) {
                            v69 := GenS6{F0: v24}
                            fmt.Println(v24, fi67, v26)
                            _ = v69
                        } else {
                            switch v20 {
                            case "foo bar":
                                v70 := GenS4{F0: -1, F1: 14, F2: -595}
                                _ = v70
                            default:
                                v71 := v4[fi67]
                                _ = v71
                            }
                        }
                    } else {
                        var v72 int8 = (v28 & (v28 % ((v28) | 1)))
                        var v73 int32 = 42
                        delete(v2, "foo bar")
                        _ = v72
                        _ = v73
                    }
                    var v74 uint8 = 255
                    _ = fi67
                    _ = v68
                    _ = v74
                }
            }
        case 32767:
            switch (-(v1)) {
            case 42:
                _ = len(v3)
            case 4294967295:
                switch "foo bar" {
                case "a":
                    for fi75 := 0; fi75 < 4; fi75++ {
                        v76 := GenS7{F0: ((((((1e10 / (3.14159 * ((((((0.1 - (float32(v26) * (1e10 + 0.0))) / 1.0) - (1.5 + 0.1)) + 0.0) / 3.14159) + 0.1))) - 3.14159) + ((1.5 / ((-2.25 + (1e10 / float32(v26))) - -2.25)) + 1e10)) * 1.5) * float32(0.1)) == float32((((v26 - 1.5) - float64(0.1)) * float64(0.1)))), F1: int32(v24)}
                        _ = fi75
                        _ = v76
                    }
                case "foo bar":
                    var v77 int16 = genFn0()
                    _ = v77
                }
            default:
                var v78 bool = v27
                _ = v78
            }
        case -2147483648:
            var v79 uint = 980
            _ = v79
        default:
            for fi80 := 0; fi80 < 9; fi80++ {
                var v81 int8 = 100
                _ = len(v4)
                _ = fi80
                _ = v81
            }
        }
        _ = fi29
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v20
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
