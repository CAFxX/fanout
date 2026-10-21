package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int; F2 uint16 }
type GenS2 struct { F0 int8; F1 int }
type GenS3 struct { F0 string; F1 bool; F2 uint }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 int16; F1 int32; F2 uint8 }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 uint32 }
type GenS8 struct { F0 int16; F1 uint }
type GenS9 struct { F0 uint64 }
func genFn0(p0 int8) string {
    var v7 string = "\n\t"
    for fi8 := 0; fi8 < 8; fi8++ {
        p0 -= p0
        switch "foo bar" {
        case "":
            var v9 bool = true
            _ = v9
        case "hello":
            for fi10 := 0; fi10 < 15; fi10++ {
                var v11 string = ("a" + "")
                var v12 bool = false
                if true {
                    fmt.Println(v11)
                } else {
                    v12 = v12
                    fmt.Println(v7, p0)
                    v13 := []int{fi8}
                    switch v7 {
                    case "foo bar":
                        var v14 uint8 = 128
                        _ = v14
                    case "":
                        v15 := GenS1{F0: 9223372036854775807, F1: 814, F2: 1}
                        _ = v15
                    default:
                        if (p0 <= p0) {
                            switch fi10 {
                            case -975:
                                v11 = ""
                            case -552:
                                if (140 > 462) {
                                    v16 := GenS2{F0: (p0 << ((481) & 63)), F1: -2147483648}
                                    _ = v16
                                }
                            default:
                                v7 = ("hello" + v7)
                            }
                            v11 = "foo bar"
                            var v17 int = fi8
                            _ = v17
                        } else {
                            var v18 string = "a"
                            _ = v18
                        }
                    }
                    _ = v13
                }
                _ = fi10
                _ = v11
                _ = v12
            }
        case "\n\t":
            if false {
                if (!((false || false) || (118 > (p0 << ((7) & 63))))) {
                    v19 := GenS2{F0: (p0 + p0), F1: fi8}
                    fmt.Println(fi8, p0)
                    _ = v19
                }
                var v20 uint64 = 268
                v7 = ((v7 + "a") + ("a" + v7))
                _ = v20
            } else {
                v21 := []int{}
                v7 = "hello"
                v7 = (v7 + (v7 + v7))
                _ = v21
            }
        default:
            fmt.Println(fi8, p0, p0)
        }
        _ = fi8
    }
    v7 = (v7 + v7)
    v7 = ("a" + "")
    _ = p0
    _ = v7
    return (v7 + ("a" + v7))
}
func genFn1(p0 int16, p1 uint64, p2 int) uint32 {
    var v22 string = "a"
    v23 := []int{(p2 >> ((888) & 63)), (-(p2)), 9223372036854775807}
    switch (p2 * p2) {
    case -2147483648:
        if ((p2 | 807) <= (p2 * p2)) {
            var v24 uint = 390
            _ = v24
        } else {
            p1 += 42
            fmt.Println(p2, p2)
        }
    case -2:
        for fi25 := 0; fi25 < 7; fi25++ {
            v23 = append(v23, p2)
            var v26 float32 = 0.0
            if (false || true) {
                var v27 int = fi25
                for fi28 := 0; fi28 < 8; fi28++ {
                    if len(v23) > 0 {
                        v23[0] = fi28
                    }
                    p1 = 2
                    _ = fi28
                }
                _ = v27
            }
            _ = fi25
            _ = v26
        }
    default:
        fmt.Println(p1, p1, v22)
    }
    if false {
        v22 = v22
        _ = len(v23)
    }
    switch p2 {
    case 32767:
        v29 := GenS3{F0: ((v22 + v22) + "hello"), F1: (true && (65535 < 32767)), F2: 39}
        _ = v29
    case 32768:
        for fi30 := 0; fi30 < 4; fi30++ {
            for fi31 := 0; fi31 < 16; fi31++ {
                if (true || true) {
                    var v32 int16 = 1000
                    _ = len(v23)
                    if (!(false || false)) {
                        v22 = ""
                        switch 32767 {
                        case -855:
                            fmt.Println(v32)
                        }
                        v22 = (v22 + v22)
                    } else {
                        v33 := GenS4{F0: 324}
                        if true {
                            var v34 uint8 = uint8(p1)
                            v34 -= (v34 + v34)
                            v22 = v22
                            _ = v34
                        }
                        switch "foo bar" {
                        case "":
                            v32 = int16(fi30)
                        case "a":
                            fmt.Println(v32)
                        default:
                            if (true || false) {
                                v35 := GenS3{F0: ("a" + (v22 + (("a" + v22) + v22))), F1: (!false), F2: 42}
                                v36 := GenS4{F0: 0}
                                if (true && false) {
                                    var v37 uint64 = p1
                                    _ = v37
                                }
                                fmt.Println(p1, p2, fi31)
                                _ = v35
                                _ = v36
                            }
                        }
                        var v38 uint64 = p1
                        _ = v33
                        _ = v38
                    }
                    if len(v23) > 0 {
                        v23[0] = p2
                    }
                    v39 := GenS4{F0: 0}
                    _ = v32
                    _ = v39
                } else {
                    v22 = (v22 + (v22 + v22))
                    var v40 int8 = 0
                    p1 = p1
                    _ = v40
                }
                _ = fi31
            }
            if len(v23) > 0 {
                v23[0] = p2
            }
            _ = fi30
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v22
    _ = v23
    return 2
}

func main() {
    var v1 uint32 = 2
    var v2 int8 = -128
    var v3 uint16 = 42
    var v4 int64 = 285
    v5 := make(map[string]bool)
    var v6 int = 32768
    fmt.Println(v3)
    if true {
        for fi41 := 0; fi41 < 20; fi41++ {
            v42 := GenS5{F0: -1, F1: 0, F2: 1}
            fmt.Println(v1)
            var v43 uint = 1
            _ = fi41
            _ = v42
            _ = v43
        }
        var v44 uint = 4294967295
        _ = v44
    } else {
        var v45 float32 = 1e10
        var v46 int8 = -42
        _ = v45
        _ = v46
    }
    fmt.Println(v3, v4, v6)
    for fi47 := 0; fi47 < 19; fi47++ {
        if true {
            var v48 int8 = v2
            switch "foo bar" {
            case "foo bar":
                for fi49 := 0; fi49 < 7; fi49++ {
                    var v50 uint8 = 116
                    fmt.Println(fi47)
                    _ = fi49
                    _ = v50
                }
            default:
                v4 -= v4
            }
            _ = v48
        }
        fmt.Println(fi47, v6)
        v1 += 42
        _ = fi47
    }
    fmt.Println(v4, v6, v2)
    for fi51 := 0; fi51 < 13; fi51++ {
        var v52 uint = 4294967295
        v5["\n\t"] = (v3 != 1)
        if false {
            var v53 uint = v52
            delete(v5, "a")
            _ = v53
        }
        _ = fi51
        _ = v52
    }
    var v54 int = 2147483647
    v4 -= (v4 / ((int64(v6)) | 1))
    if false {
        v55 := GenS6{F0: (v4 & v4)}
        v54 = v6
        if (float64(float32(float64(1e10))) < float64(((3.14159 + (0.0 - 0.0)) * (((1.5 * 1.5) + 1.5) + (float32(3.14159) - -2.25))))) {
            v1 -= (v1 - (v1 ^ (v1 & (v1 * (^(v1))))))
            var v56 int32 = 1
            var v57 int = (v6 >> ((1) & 63))
            _ = v56
            _ = v57
        }
        _ = v55
    } else {
        if ((0.0 - (-2.25 - 0.1)) > 0.1) {
            if (131 > 255) {
                v58 := []int{(v6 << ((1) & 63)), (v54 - v6), v54}
                v6 -= v54
                v5[""] = true
                _ = v58
            } else {
                switch "hello" {
                case "a":
                    var v59 uint8 = 115
                    _ = v59
                case "\n\t":
                    var v60 uint16 = (v3 >> ((51) & 63))
                    _ = v60
                case "foo bar":
                    if true {
                        v5["hello"] = true
                    }
                default:
                    v4 += v4
                }
                for fi61 := 0; fi61 < 16; fi61++ {
                    v62 := GenS7{F0: v1}
                    _ = fi61
                    _ = v62
                }
            }
        }
    }
    for fi63 := 0; fi63 < 5; fi63++ {
        var v64 uint = 1
        _ = fi63
        _ = v64
    }
    if false {
        v2 = v2
    } else {
        v2 = v2
        if (true || false) {
            v3 += (v3 ^ v3)
            v65 := GenS6{F0: int64(v3)}
            if ((v1 - genFn1(1000, 1, 32767)) > 138) {
                switch "\n\t" {
                case "foo bar":
                    for fi66 := 0; fi66 < 8; fi66++ {
                        v67 := GenS8{F0: -365, F1: 761}
                        for fi68 := 0; fi68 < 9; fi68++ {
                            if false {
                                v6 += v6
                                var v69 int16 = -1
                                _ = v69
                            }
                            v70 := []int{}
                            v54 += 9223372036854775807
                            v1 -= (v1 << ((0) & 63))
                            _ = fi68
                            _ = v70
                        }
                        switch ("" + "a") {
                        case "\n\t":
                            var v71 int32 = 2147483647
                            _ = v71
                        default:
                            var v72 bool = true
                            _ = v72
                        }
                        _ = fi66
                        _ = v67
                    }
                case "\n\t":
                    switch v54 {
                    case 127:
                        var v73 float64 = 3.14159
                        _ = v73
                    case 1:
                        fmt.Println(v54, v6)
                    }
                case "":
                    var v74 uint32 = v1
                    _ = v74
                default:
                    var v75 int16 = 42
                    _ = v75
                }
                fmt.Println(v4)
                switch 32768 {
                case 0:
                    var v76 uint = 42
                    _ = v76
                case 891:
                    var v77 bool = true
                    _ = v77
                case 127:
                    v2 = (v2 ^ 100)
                }
            }
            _ = v65
        } else {
            v78 := GenS9{F0: 0}
            _ = v78
        }
    }
    for fi79 := 0; fi79 < 19; fi79++ {
        v80 := []int{(-(fi79)), 569, (fi79 ^ 1)}
        _ = fi79
        _ = v80
    }
    v3 = v3
    if ((v2 - 100) <= v2) {
        v81 := []int{(v6 - v54), (v54 - (v54 >> ((42) & 63)))}
        _ = v81
    }
    for fi82 := 0; fi82 < 3; fi82++ {
        v6 = v54
        _ = fi82
    }
    var v83 uint = 18446744073709551615
    v6 += v54
    v4 = (v4 + (v4 >> (((^(v83))) & 63)))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v54:", v54)
    fmt.Println("v6:", v6)
    fmt.Println("v83:", v83)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v54
    _ = v6
    _ = v83
}
