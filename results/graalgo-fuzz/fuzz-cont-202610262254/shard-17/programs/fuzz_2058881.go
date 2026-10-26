package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint8; F1 uint64; F2 int }
type GenS3 struct { F0 int; F1 uint; F2 string }
type GenS4 struct { F0 int16; F1 int32; F2 int8 }
type GenS5 struct { F0 bool; F1 uint16 }

func main() {
    var v1 int16 = 1000
    var v2 uint16 = 0
    var v3 uint8 = 42
    var v4 int64 = -29
    if false {
        switch ("a" + "foo bar") {
        case "":
            v4 -= v4
        }
        fmt.Println(v3, v1)
        switch -582 {
        case 42:
            var v5 float32 = 0.0
            _ = v5
        case 128:
            fmt.Println(v2, v1)
        }
    } else {
        v6 := GenS1{F0: 0}
        v3 += v3
        _ = v6
    }
    for fi7 := 0; fi7 < 16; fi7++ {
        v2 = (v2 >> ((42) & 63))
        var v8 uint64 = 212
        _ = fi7
        _ = v8
    }
    v1 -= v1
    v9 := GenS2{F0: (v3 + (v3 << ((42) & 63))), F1: 151, F2: 2147483647}
    v1 = 1
    var v10 int32 = 42
    v11 := GenS2{F0: 0, F1: 42, F2: 760}
    var v12 bool = true
    v10 = (^(v10))
    if v12 {
        v10 -= (v10 / (((v10 % (((^(v10))) | 1))) | 1))
        if v12 {
            switch 4294967295 {
            case -1:
                v1 += 1000
            case -9223372036854775808:
                if (!((!v12) || (0 >= 4294967295))) {
                    var v13 uint8 = (v3 / ((128) | 1))
                    v14 := GenS3{F0: 128, F1: 522, F2: ""}
                    v15 := GenS2{F0: v3, F1: 57, F2: 4294967295}
                    _ = v13
                    _ = v14
                    _ = v15
                } else {
                    v2 += v2
                    var v16 float32 = 0.1
                    var v17 uint16 = v2
                    _ = v16
                    _ = v17
                }
            case 0:
                switch -2147483648 {
                case -1:
                    fmt.Println(v4)
                case 4294967295:
                    v18 := []int{}
                    _ = v18
                }
            default:
                v19 := []int{617, -716, 868, 32768}
                _ = v19
            }
            v3 -= (v3 + v3)
        }
        var v20 uint8 = v3
        _ = v20
    } else {
        var v21 uint32 = 2147483647
        var v22 uint = 1
        _ = v21
        _ = v22
    }
    if v12 {
        for fi23 := 0; fi23 < 3; fi23++ {
            v24 := GenS2{F0: v3, F1: 2, F2: (-(fi23))}
            for fi25 := 0; fi25 < 19; fi25++ {
                v4 = (v4 | (v4 & v4))
                switch "a" {
                case "a":
                    v3 = v3
                case "hello":
                    switch 509 {
                    case -591:
                        v26 := []int{}
                        _ = v26
                    case 49:
                        switch fi25 {
                        case 1:
                            switch "\n\t" {
                            case "":
                                if v12 {
                                    fmt.Println(v12)
                                } else {
                                    var v27 uint = 341
                                    switch "\n\t" {
                                    case "hello":
                                        v1 += int16(v2)
                                    case "foo bar":
                                        v28 := GenS1{F0: v27}
                                        _ = v28
                                    case "\n\t":
                                        v29 := []int{(fi25 / ((0) | 1)), (fi25 / ((fi25) | 1)), -2147483648}
                                        _ = v29
                                    default:
                                        fmt.Println(v4, v3)
                                    }
                                    _ = v27
                                }
                            case "foo bar":
                                fmt.Println(v2, v4, v3)
                            case "hello":
                                v30 := []int{fi23, fi23, 4294967295, (fi25 >> ((65535) & 63))}
                                _ = v30
                            default:
                                v12 = (1 != 4294967295)
                                v31 := GenS2{F0: 214, F1: 1, F2: (fi25 | fi23)}
                                v32 := GenS3{F0: (-(fi25)), F1: 42, F2: ""}
                                _ = v31
                                _ = v32
                            }
                        case 32767:
                            v33 := GenS1{F0: 1}
                            _ = v33
                        case 133:
                            v34 := []int{}
                            _ = v34
                        }
                    case -302:
                        v4 -= (v4 - v4)
                    default:
                        v2 = (v2 >> ((682) & 63))
                    }
                default:
                    v3 -= v3
                }
                var v35 int16 = 42
                v36 := GenS1{F0: uint(v2)}
                var v37 uint8 = (^(v3))
                v35 += v1
                if v12 {
                    var v38 int32 = v10
                    fmt.Println(v35, v10)
                    var v39 uint16 = v2
                    v3 -= v3
                    v38 -= v10
                    if ("a" < "") {
                        v2 = 32767
                    } else {
                        v40 := GenS3{F0: (^(fi23)), F1: 9, F2: "hello"}
                        _ = v40
                    }
                    _ = v38
                    _ = v39
                } else {
                    v41 := GenS2{F0: (v3 | v37), F1: 2, F2: 128}
                    v10 -= v10
                    _ = v41
                }
                _ = fi25
                _ = v35
                _ = v36
                _ = v37
            }
            _ = fi23
            _ = v24
        }
        if (false && ((v2 - v2) > (v2 & (^(v2))))) {
            for fi42 := 0; fi42 < 8; fi42++ {
                v4 -= v4
                v3 += 127
                _ = fi42
            }
            v43 := []int{2147483647, 0, 32768}
            var v44 string = ((("foo bar" + "") + "foo bar") + "")
            _ = v43
            _ = v44
        }
        v4 = (^(v4))
    } else {
        if (32767 >= 275) {
            fmt.Println(v1, v10)
            v45 := GenS4{F0: (v1 - (v1 & (v1 | (v1 ^ 0)))), F1: v10, F2: 100}
            _ = v45
        } else {
            if v12 {
                var v46 string = "hello"
                for fi47 := 0; fi47 < 9; fi47++ {
                    switch v46 {
                    case "\n\t":
                        switch v46 {
                        case "":
                            v10 = v10
                        case "hello":
                            if false {
                                v48 := []int{fi47, fi47, fi47, (^(fi47))}
                                v3 += 2
                                _ = v48
                            } else {
                                var v49 uint64 = 384
                                if (v10 <= int32(v2)) {
                                    for fi50 := 0; fi50 < 12; fi50++ {
                                        if (fi47 < (^(fi47))) {
                                            var v51 int = 32768
                                            switch (fi50 / (((v51 ^ fi47)) | 1)) {
                                            case 127:
                                                v51 -= (v51 / ((127) | 1))
                                            case 128:
                                                fmt.Println(fi50, v2)
                                            case 42:
                                                v52 := GenS3{F0: (^(fi50)), F1: 0, F2: "foo bar"}
                                                _ = v52
                                            }
                                            var v53 bool = (((-28 > v4) || true) && ((!v12) && v12))
                                            _ = v51
                                            _ = v53
                                        } else {
                                            var v54 uint64 = v49
                                            _ = v54
                                        }
                                        var v55 uint8 = uint8(v49)
                                        _ = fi50
                                        _ = v55
                                    }
                                    v10 = (^(v10))
                                    if (v12 && (fi47 <= int(v49))) {
                                        var v56 uint8 = uint8(v2)
                                        v56 += (v3 & v56)
                                        _ = v56
                                    }
                                } else {
                                    if true {
                                        v57 := GenS3{F0: int(v10), F1: 255, F2: "\n\t"}
                                        _ = v57
                                    } else {
                                        var v58 int = -2147483648
                                        v59 := GenS4{F0: v1, F1: (^(v10)), F2: -128}
                                        _ = v58
                                        _ = v59
                                    }
                                }
                                _ = v49
                            }
                        default:
                            v60 := []int{42, (fi47 / ((32767) | 1)), (fi47 | 884)}
                            _ = v60
                        }
                    case "":
                        if v12 {
                            v61 := []int{(fi47 + 32767), fi47, fi47}
                            _ = v61
                        }
                    case "hello":
                        v4 = 0
                    default:
                        v62 := []int{(fi47 + 32768)}
                        _ = v62
                    }
                    if true {
                        for fi63 := 0; fi63 < 6; fi63++ {
                            var v64 uint16 = v2
                            v12 = true
                            var v65 string = (v46 + ("" + "a"))
                            _ = fi63
                            _ = v64
                            _ = v65
                        }
                        for fi66 := 0; fi66 < 4; fi66++ {
                            v12 = v12
                            var v67 float64 = 1e10
                            _ = fi66
                            _ = v67
                        }
                    }
                    v3 = (v3 * v3)
                    _ = fi47
                }
                _ = v46
            }
        }
    }
    fmt.Println(v1, v12, v12)
    for fi68 := 0; fi68 < 17; fi68++ {
        fmt.Println(v2, v4)
        switch ("foo bar" + "a") {
        case "\n\t":
            var v69 uint64 = 0
            _ = v69
        default:
            fmt.Println(v2, v12)
        }
        v70 := []int{fi68, 32768}
        _ = fi68
        _ = v70
    }
    var v71 int16 = (v1 & 1)
    for fi72 := 0; fi72 < 15; fi72++ {
        for fi73 := 0; fi73 < 19; fi73++ {
            if (719 != 49) {
                v74 := GenS1{F0: 18446744073709551615}
                v12 = ((v12 || v12) || v12)
                _ = v74
            }
            v10 += v10
            switch "\n\t" {
            case "\n\t":
                v75 := GenS1{F0: 965}
                _ = v75
            case "a":
                v71 = 1
            case "":
                if false {
                    if (56 > -128) {
                        v71 += 0
                        v2 += v2
                        v12 = false
                    }
                    if true {
                        v12 = v12
                        var v76 uint16 = uint16(v3)
                        _ = v76
                    }
                    v77 := GenS5{F0: v12, F1: v2}
                    _ = v77
                } else {
                    var v78 uint64 = 204
                    if v12 {
                        if v12 {
                            fmt.Println(fi72, v12)
                        }
                    }
                    v79 := []int{fi73}
                    _ = v78
                    _ = v79
                }
            }
            _ = fi73
        }
        v12 = v12
        _ = fi72
    }
    fmt.Println(v3, v10)
    v80 := []int{}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v71:", v71)
    fmt.Println("slicelen v80:", len(v80), v80)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v71
    _ = v80
    _ = v9
}
