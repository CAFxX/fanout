package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64 }
type GenS2 struct { F0 uint; F1 uint8; F2 string }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint32; F1 int32; F2 int16 }
type GenS5 struct { F0 bool; F1 uint32 }
type GenS6 struct { F0 bool; F1 int; F2 uint }
type GenS7 struct { F0 string; F1 uint16; F2 uint8 }
type GenS8 struct { F0 uint16; F1 int8 }
type GenS9 struct { F0 string; F1 bool; F2 uint16 }
type GenS10 struct { F0 int8; F1 uint64; F2 int }

func main() {
    var v1 int64 = 9223372036854775807
    var v2 uint64 = 21
    var v3 uint64 = (v2 >> ((1) & 63))
    var v4 int32 = 0
    var v5 int = 86
    switch int(v3) {
    case 127:
        v6 := GenS1{F0: 0, F1: 218}
        _ = v6
    case -9223372036854775808:
        switch 0 {
        case 0:
            fmt.Println(v1, v3)
        case 32767:
            v2 = v3
        default:
            v7 := []int{795, 32768}
            _ = v7
        }
    default:
        v8 := GenS1{F0: -764, F1: (v3 | 2)}
        _ = v8
    }
    v9 := GenS1{F0: v5, F1: (^(v3))}
    if (v3 >= (v2 ^ (v3 % ((568) | 1)))) {
        switch v5 {
        case -350:
            if (128 >= (v5 + (v5 << ((18446744073709551615) & 63)))) {
                fmt.Println(v2)
                v5 -= (v5 >> ((0) & 63))
                if true {
                    v2 -= (v3 ^ (v3 | (v3 << ((65535) & 63))))
                }
            } else {
                if true {
                    var v10 uint16 = 461
                    for fi11 := 0; fi11 < 1; fi11++ {
                        var v12 uint = 1
                        switch v5 {
                        case 128:
                            var v13 uint = 2
                            _ = v13
                        case 32768:
                            v12 = (v12 - v12)
                        case -199:
                            v5 -= fi11
                        default:
                            var v14 int8 = -128
                            _ = v14
                        }
                        for fi15 := 0; fi15 < 15; fi15++ {
                            v5 += fi15
                            _ = fi15
                        }
                        _ = fi11
                        _ = v12
                    }
                    if false {
                        v3 = v2
                    }
                    _ = v10
                } else {
                    for fi16 := 0; fi16 < 14; fi16++ {
                        v17 := GenS1{F0: (v5 >> ((0) & 63)), F1: v2}
                        v18 := GenS1{F0: (fi16 - (fi16 << ((1) & 63))), F1: v2}
                        _ = fi16
                        _ = v17
                        _ = v18
                    }
                    if false {
                        v3 -= v2
                        var v19 string = "hello"
                        v4 = (-(v4))
                        _ = v19
                    } else {
                        var v20 uint = 0
                        v21 := GenS2{F0: v20, F1: 22, F2: ("\n\t" + "hello")}
                        _ = v20
                        _ = v21
                    }
                }
                if (true && (127 != 100)) {
                    v3 -= (^(v2))
                    fmt.Println(v4, v3, v3)
                }
            }
        case 32768:
            v22 := []int{(v5 & v5), (v5 * v5), 2147483647, int(v1)}
            _ = v22
        default:
            var v23 int64 = (v1 + v1)
            _ = v23
        }
    }
    if true {
        v3 = v3
        v24 := []int{v5}
        switch (v5 & (v5 * (v5 - v5))) {
        case -992:
            for fi25 := 0; fi25 < 5; fi25++ {
                v1 -= (v1 << ((42) & 63))
                if false {
                    if ((v4 / (((v4 / (((v4 / ((v4) | 1))) | 1))) | 1)) >= v4) {
                        var v26 uint32 = 42
                        var v27 uint16 = 65535
                        v26 = v26
                        _ = v26
                        _ = v27
                    } else {
                        switch -1 {
                        case 32767:
                            fmt.Println(v1, v5)
                        default:
                            for fi28 := 0; fi28 < 3; fi28++ {
                                if false {
                                    v5 += (fi28 << ((171) & 63))
                                    fmt.Println(v3, fi25)
                                } else {
                                    if false {
                                        v29 := GenS3{F0: 4294967295}
                                        v2 = (^(v3))
                                        var v30 float32 = 1e10
                                        _ = v29
                                        _ = v30
                                    }
                                    if true {
                                        switch (fi25 - fi28) {
                                        case 4294967295:
                                            if len(v24) >= 2 {
                                                v24 = v24[1:len(v24)]
                                            }
                                        default:
                                            switch (fi28 / ((9223372036854775807) | 1)) {
                                            case 32768:
                                                v24 = append(v24, (^(v5)))
                                            case 42:
                                                var v31 int = (fi28 >> ((413) & 63))
                                                v4 -= v4
                                                var v32 float64 = (0.0 + (3.14159 + 0.1))
                                                _ = v31
                                                _ = v32
                                            case 2147483647:
                                                v24 = append(v24, v5)
                                            }
                                            var v33 int16 = -1
                                            _ = v33
                                        }
                                        v2 += 42
                                        var v34 int64 = v1
                                        _ = v34
                                    }
                                    v35 := GenS2{F0: uint(v3), F1: 1, F2: ("" + (("foo bar" + ("a" + "")) + "hello"))}
                                    v5 -= (v5 ^ int(v4))
                                    var v36 int64 = v1
                                    _ = v35
                                    _ = v36
                                }
                                v37 := GenS4{F0: 605, F1: 0, F2: 1}
                                _ = fi28
                                _ = v37
                            }
                        }
                        v5 = (fi25 ^ (fi25 ^ (^(fi25))))
                    }
                    v1 += (v1 * (v1 - (v1 * v1)))
                    v4 += v4
                } else {
                    fmt.Println(v1, fi25, fi25)
                }
                var v38 uint8 = 128
                _ = fi25
                _ = v38
            }
        case -63:
            v4 = -1
        default:
            for fi39 := 0; fi39 < 14; fi39++ {
                v5 -= v5
                var v40 uint64 = 657
                _ = fi39
                _ = v40
            }
        }
        _ = v24
    } else {
        var v41 float64 = (((1e10 / (((0.1 + (float64((float32(1e10) / 1e10)) / 1e10)) - float64(1.5)) * 3.14159)) + -2.25) * 1e10)
        v42 := []int{}
        if len(v42) >= 2 {
            v42 = v42[1:len(v42)]
        }
        _ = v41
        _ = v42
    }
    switch "foo bar" {
    case "\n\t":
        fmt.Println(v2, v4)
    }
    switch "a" {
    case "":
        v3 = (^(v2))
    case "foo bar":
        for fi43 := 0; fi43 < 11; fi43++ {
            v44 := []int{v5}
            v3 += (v2 ^ 65)
            _ = fi43
            _ = v44
        }
    case "\n\t":
        v2 = v3
    default:
        v3 += 569
    }
    for fi45 := 0; fi45 < 3; fi45++ {
        switch v5 {
        case -130:
            v3 -= (v2 * v3)
        case 128:
            switch "\n\t" {
            case "\n\t":
                var v46 string = "foo bar"
                _ = v46
            case "foo bar":
                for fi47 := 0; fi47 < 9; fi47++ {
                    var v48 uint16 = 32767
                    _ = fi47
                    _ = v48
                }
            case "a":
                fmt.Println(v2)
            }
        }
        if true {
            switch "hello" {
            case "\n\t":
                v5 += fi45
            case "":
                switch "hello" {
                case "hello":
                    v1 -= v1
                case "\n\t":
                    v49 := GenS5{F0: (true || (1 >= 439)), F1: 4294967295}
                    _ = v49
                default:
                    v50 := []int{(v5 << ((2) & 63)), (v5 << ((2) & 63)), v5, (fi45 * 592)}
                    _ = v50
                }
            case "foo bar":
                if (((((v3 > (v2 - (v2 - (v2 ^ (^(v3)))))) || true) || (!true)) && true) || false) {
                    for fi51 := 0; fi51 < 9; fi51++ {
                        v3 -= v2
                        v52 := GenS5{F0: ((-(v5)) < 4294967295), F1: 2}
                        v53 := GenS6{F0: false, F1: -830, F2: 4294967295}
                        _ = fi51
                        _ = v52
                        _ = v53
                    }
                    v54 := []int{v5, fi45, (v5 & (v5 & fi45)), (-(fi45))}
                    var v55 int = 2147483647
                    _ = v54
                    _ = v55
                } else {
                    fmt.Println(v1, v4, v2)
                    var v56 uint32 = 2
                    switch v5 {
                    case -814:
                        v57 := []int{fi45, (-(v5)), v5, fi45}
                        _ = v57
                    case 442:
                        fmt.Println(fi45, v1, v2)
                    case 877:
                        v2 = (^(v2))
                    default:
                        if (true && (2 < 0)) {
                            switch fi45 {
                            case 0:
                                if false {
                                    v56 = v56
                                    var v58 int16 = 0
                                    var v59 int32 = 1
                                    _ = v58
                                    _ = v59
                                }
                            case 998:
                                var v60 int16 = -1
                                _ = v60
                            default:
                                v61 := GenS6{F0: false, F1: (v5 | 32768), F2: 534}
                                _ = v61
                            }
                        }
                    }
                    _ = v56
                }
            }
            v62 := GenS7{F0: ("" + "\n\t"), F1: 255, F2: 2}
            fmt.Println(v3)
            _ = v62
        } else {
            v63 := []int{fi45, v5, (fi45 & fi45), (fi45 >> ((4294967295) & 63))}
            _ = v63
        }
        fmt.Println(v4)
        _ = fi45
    }
    v64 := []int{2147483647, (v5 + v5), (v5 % (((v5 | v5)) | 1))}
    v2 = v2
    var v65 int = v5
    if true {
        for fi66 := 0; fi66 < 5; fi66++ {
            for fi67 := 0; fi67 < 16; fi67++ {
                var v68 uint8 = 1
                fmt.Println(v3, v3)
                _ = fi67
                _ = v68
            }
            _ = fi66
        }
    }
    var v69 bool = (false && ((v4 << ((4294967295) & 63)) < (v4 ^ v4)))
    switch "" {
    case "hello":
        if v69 {
            if (42 > 255) {
                if len(v64) >= 2 {
                    v64 = v64[1:len(v64)]
                }
                if len(v64) > 0 {
                    v64[0] = (v65 & (v5 % ((406) | 1)))
                }
                switch -823 {
                case 4294967295:
                    v69 = v69
                case 128:
                    v65 = (-(v5))
                }
            } else {
                var v70 uint64 = 42
                v4 = (v4 ^ v4)
                for fi71 := 0; fi71 < 18; fi71++ {
                    fmt.Println(v2, v1)
                    switch "" {
                    case "a":
                        if v69 {
                            for fi72 := 0; fi72 < 15; fi72++ {
                                v3 = (v70 | 18446744073709551615)
                                _ = fi72
                            }
                        } else {
                            v64 = append(v64, v5)
                            var v73 uint8 = 31
                            if ((!v69) || v69) {
                                fmt.Println(v69)
                                v3 = (v2 * (v70 + 2))
                                v74 := GenS8{F0: 255, F1: -42}
                                _ = v74
                            }
                            _ = v73
                        }
                    }
                    v75 := GenS6{F0: v69, F1: (-(v65)), F2: 65535}
                    _ = fi71
                    _ = v75
                }
                _ = v70
            }
            v76 := GenS9{F0: "foo bar", F1: (9223372036854775807 < (^(v65))), F2: 1}
            var v77 int8 = 42
            _ = v76
            _ = v77
        }
    case "\n\t":
        v78 := GenS10{F0: 42, F1: (v2 - v2), F2: v5}
        _ = v78
    case "a":
        v1 += (v1 + v1)
    }
    var v79 uint8 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v64:", len(v64), v64)
    fmt.Println("v65:", v65)
    fmt.Println("v69:", v69)
    fmt.Println("v79:", v79)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v64
    _ = v65
    _ = v69
    _ = v79
    _ = v9
}
