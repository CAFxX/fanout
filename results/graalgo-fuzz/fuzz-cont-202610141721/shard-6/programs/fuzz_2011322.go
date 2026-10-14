package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int64; F2 uint8 }
type GenS2 struct { F0 bool; F1 uint8; F2 int16 }
type GenS3 struct { F0 uint32; F1 uint8 }
type GenS4 struct { F0 uint64; F1 uint }
type GenS5 struct { F0 bool; F1 string; F2 uint64 }
type GenS6 struct { F0 string }
type GenS7 struct { F0 int; F1 int64; F2 bool }
type GenS8 struct { F0 uint32; F1 uint64; F2 int64 }
type GenS9 struct { F0 string; F1 uint64; F2 int8 }
type GenS10 struct { F0 uint64; F1 uint8 }
type GenS11 struct { F0 int; F1 uint8 }
func genFn0(p0 int16) uint8 {
    switch 1 {
    case -961:
        var v5 uint16 = 65535
        _ = v5
    case 907:
        if ((43 <= -1) || (!(-2.25 == float64(-2.25)))) {
            fmt.Println(p0)
            p0 -= 1
        } else {
            var v6 uint = 4294967295
            var v7 uint64 = 0
            var v8 float32 = (0.1 - 0.0)
            _ = v6
            _ = v7
            _ = v8
        }
    }
    fmt.Println(p0)
    fmt.Println(p0)
    p0 = p0
    var v9 bool = false
    _ = p0
    _ = v9
    return 0
}
func genFn1() int64 {
    var v10 bool = true
    v11 := GenS1{F0: 2147483647, F1: 0, F2: 42}
    var v12 float32 = 1e10
    if v10 {
        v13 := []int{}
        for fi14 := 0; fi14 < 14; fi14++ {
            for fi15 := 0; fi15 < 14; fi15++ {
                var v16 uint64 = 42
                v12 -= v12
                _ = fi15
                _ = v16
            }
            for fi17 := 0; fi17 < 7; fi17++ {
                var v18 string = "foo bar"
                _ = fi17
                _ = v18
            }
            _ = fi14
        }
        if (11 < 0) {
            for fi19 := 0; fi19 < 2; fi19++ {
                v12 += (3.14159 + (v12 * 1.5))
                switch "\n\t" {
                case "":
                    for fi20 := 0; fi20 < 19; fi20++ {
                        switch fi19 {
                        case 128:
                            v10 = v10
                            var v21 string = ""
                            var v22 bool = (!v10)
                            _ = v21
                            _ = v22
                        case 767:
                            var v23 int64 = 9223372036854775807
                            _ = v23
                        case 9223372036854775807:
                            v12 += float32(-2.25)
                        default:
                            var v24 int8 = int8(fi19)
                            var v25 int32 = 1
                            _ = v24
                            _ = v25
                        }
                        v10 = (v10 && (!(!v10)))
                        switch "foo bar" {
                        case "\n\t":
                            switch 0 {
                            case 774:
                                v12 -= float32(1.5)
                            case 957:
                                v26 := GenS2{F0: (594 < 4294967295), F1: 59, F2: 42}
                                _ = v26
                            default:
                                if v10 {
                                    v27 := GenS3{F0: 42, F1: genFn0(42)}
                                    _ = v27
                                } else {
                                    switch "hello" {
                                    case "foo bar":
                                        _ = len(v13)
                                    case "":
                                        fmt.Println(v12)
                                    case "\n\t":
                                        var v28 uint8 = 128
                                        _ = v28
                                    }
                                    v13 = append(v13, fi20)
                                    if len(v13) > 0 {
                                        v13[0] = (fi20 + (fi20 | (^(fi20))))
                                    }
                                    var v29 uint16 = 381
                                    _ = v29
                                }
                            }
                        case "hello":
                            var v30 uint8 = 42
                            _ = v30
                        default:
                            switch fi19 {
                            case 127:
                                var v31 int64 = 9223372036854775807
                                _ = v31
                            case 567:
                                v32 := GenS3{F0: 1, F1: 54}
                                _ = v32
                            case 32768:
                                var v33 uint64 = 0
                                _ = v33
                            default:
                                var v34 float32 = 0.1
                                _ = v34
                            }
                        }
                        _ = fi20
                    }
                case "foo bar":
                    if len(v13) > 0 {
                        v13[0] = fi19
                    }
                default:
                    v10 = (1.5 != 3.14159)
                }
                _ = fi19
            }
            switch 1 {
            case -612:
                var v35 uint32 = 65535
                _ = v35
            case 9223372036854775807:
                var v36 int16 = 0
                _ = v36
            }
        }
        _ = v13
    }
    _ = v10
    _ = v11
    _ = v12
    return 9223372036854775807
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint = 42
    var v3 uint = (^(v2))
    var v4 float64 = 0.1
    v37 := []int{}
    v38 := GenS4{F0: 1, F1: (v2 - v3)}
    v39 := GenS3{F0: 2, F1: 128}
    fmt.Println(v2, v3)
    v2 = 2
    var v40 uint64 = 4294967295
    delete(v1, 477)
    switch "\n\t" {
    case "foo bar":
        for fi41 := 0; fi41 < 6; fi41++ {
            for fi42 := 0; fi42 < 9; fi42++ {
                if true {
                    var v43 string = "hello"
                    switch 42 {
                    case 2147483647:
                        var v44 uint64 = 18446744073709551615
                        _ = v44
                    }
                    _ = v43
                } else {
                    var v45 float32 = ((1e10 - 1e10) / (0.0 + 0.1))
                    if true {
                        v46 := GenS1{F0: -1, F1: genFn1(), F2: 255}
                        _ = v46
                    }
                    var v47 uint32 = 2
                    if ((fi42 - (fi42 | -872)) > (fi41 + (fi41 + (fi41 + fi42)))) {
                        delete(v1, v47)
                        _ = len(v1)
                        if false {
                            fmt.Println(fi41, v45)
                        }
                    } else {
                        _ = len(v1)
                        if true {
                            v48 := GenS5{F0: (false && ((v40 ^ v40) >= (v40 + (^(v40))))), F1: "foo bar", F2: (v40 ^ v40)}
                            _ = v48
                        }
                        var v49 int = fi41
                        _ = v49
                    }
                    v40 = v40
                    if ((v47 & 1) <= v47) {
                        if true {
                            v40 = v40
                            if len(v37) >= 2 {
                                v37 = v37[1:len(v37)]
                            }
                            if false {
                                if (35 > uint8(v40)) {
                                    var v50 int32 = -199
                                    _ = len(v37)
                                    v50 += 0
                                    _ = v50
                                } else {
                                    var v51 int16 = 42
                                    _ = v51
                                }
                                if false {
                                    var v52 int = fi42
                                    v53 := GenS6{F0: ("foo bar" + "")}
                                    switch "a" {
                                    case "hello":
                                        v2 += (v2 % ((v3) | 1))
                                    case "\n\t":
                                        var v54 int = (fi42 ^ fi42)
                                        _ = v54
                                    case "foo bar":
                                        if false {
                                            v37 = append(v37, (fi41 >> (((v3 / ((v3) | 1))) & 63)))
                                            v2 = v2
                                        } else {
                                            if false {
                                                v55 := GenS3{F0: v47, F1: 42}
                                                v52 = 214
                                                _ = v55
                                            }
                                            switch "a" {
                                            case "":
                                                if true {
                                                    v3 += v3
                                                    v56 := GenS7{F0: fi42, F1: -1, F2: (1 > -32768)}
                                                    v57 := v1[151]
                                                    var v58 int8 = -1
                                                    v3 -= v3
                                                    _ = v56
                                                    _ = v57
                                                    _ = v58
                                                }
                                            case "a":
                                                v59 := GenS2{F0: ((!false) || true), F1: 32, F2: 397}
                                                _ = v59
                                            case "hello":
                                                v37 = append(v37, 42)
                                            default:
                                                v47 -= v47
                                            }
                                            switch 106 {
                                            case 0:
                                                switch "foo bar" {
                                                case "":
                                                    fmt.Println(fi41, v40)
                                                case "hello":
                                                    var v60 string = ("hello" + "\n\t")
                                                    _ = v60
                                                default:
                                                    v61 := GenS8{F0: 42, F1: v40, F2: -9223372036854775808}
                                                    _ = v61
                                                }
                                                if true {
                                                    switch "\n\t" {
                                                    case "hello":
                                                        switch "hello" {
                                                        case "foo bar":
                                                            v62 := GenS9{F0: "", F1: (v40 >> (((v3 + v2)) & 63)), F2: -1}
                                                            _ = v62
                                                        case "":
                                                            fmt.Println(v3, fi42)
                                                        default:
                                                            v40 -= v40
                                                        }
                                                    case "foo bar":
                                                        v63 := GenS10{F0: (v40 - (v40 >> (((v2 & (^(v3)))) & 63))), F1: 2}
                                                        _ = v63
                                                    case "":
                                                        var v64 float64 = ((((v4 - v4) - float64((((v45 + v45) / v45) + (3.14159 - (float32(1e10) * ((((v45 - v45) / v45) - ((v45 / v45) / 3.14159)) * (float32(3.14159) * 3.14159))))))) / (v4 / v4)) + v4)
                                                        _ = v64
                                                    default:
                                                        fmt.Println(v3, v2)
                                                    }
                                                    v52 = v52
                                                } else {
                                                    fmt.Println(fi42, v4, fi41)
                                                    var v65 int8 = -35
                                                    _ = v65
                                                }
                                                var v66 uint = v2
                                                v4 -= v4
                                                var v67 int64 = 264
                                                _ = v66
                                                _ = v67
                                            case -552:
                                                v37 = append(v37, (v52 / ((-580) | 1)))
                                            case 4294967295:
                                                fmt.Println(v40, v3, v45)
                                            default:
                                                if false {
                                                    v47 = v47
                                                } else {
                                                    fmt.Println(v52, v45)
                                                    var v68 float32 = 0.0
                                                    fmt.Println(v2)
                                                    fmt.Println(v47, fi41)
                                                    _ = v68
                                                }
                                            }
                                        }
                                    }
                                    _ = v52
                                    _ = v53
                                } else {
                                    var v69 float32 = (float32((v4 + 0.1)) + ((1.5 / ((v45 - (-2.25 / v45)) + 0.0)) + v45))
                                    fmt.Println(fi41, v2, fi42)
                                    _ = v69
                                }
                            }
                        }
                    } else {
                        var v70 float64 = v4
                        if true {
                            if true {
                                v47 = (v47 >> ((v2) & 63))
                                v40 = v40
                                v2 = (^(v2))
                                v40 += 1
                                if (false && true) {
                                    v71 := GenS11{F0: (fi42 << (((v3 << ((uint(v40)) & 63))) & 63)), F1: 127}
                                    v4 = 3.14159
                                    _ = v71
                                }
                                var v72 float32 = v45
                                _ = v72
                            }
                            fmt.Println(fi41, fi41, v45)
                        }
                        _ = v70
                    }
                    _ = v45
                    _ = v47
                }
                switch "\n\t" {
                case "\n\t":
                    if len(v37) >= 2 {
                        v37 = v37[1:len(v37)]
                    }
                case "":
                    var v73 float32 = ((3.14159 * 1e10) - 3.14159)
                    _ = v73
                default:
                    var v74 uint16 = 32767
                    _ = v74
                }
                v2 += 1
                v75 := GenS2{F0: true, F1: 42, F2: 0}
                v76 := v1[65535]
                _ = fi42
                _ = v75
                _ = v76
            }
            _ = fi41
        }
    case "hello":
        var v77 int8 = -122
        _ = v77
    case "a":
        var v78 int16 = 32767
        _ = v78
    }
    var v79 int32 = 94
    var v80 float64 = v4
    for fi81 := 0; fi81 < 11; fi81++ {
        var v82 uint16 = 1
        var v83 int = fi81
        _ = len(v1)
        _ = fi81
        _ = v82
        _ = v83
    }
    if len(v37) > 0 {
        v37[0] = -605
    }
    var v84 string = "a"
    fmt.Println(v2)
    v37 = append(v37, 131)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v37:", len(v37), v37)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v79:", v79)
    fmt.Println("v80:", v80)
    fmt.Println("v84:", v84)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v79
    _ = v80
    _ = v84
}
