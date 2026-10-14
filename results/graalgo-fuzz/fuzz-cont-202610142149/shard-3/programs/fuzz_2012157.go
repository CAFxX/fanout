package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint16; F1 uint32; F2 int64 }
type GenS3 struct { F0 uint32; F1 uint16 }
type GenS4 struct { F0 int16; F1 uint16 }
type GenS5 struct { F0 int32; F1 int8; F2 int64 }
type GenS6 struct { F0 bool; F1 uint64 }
type GenS7 struct { F0 int; F1 int8; F2 int32 }
type GenS8 struct { F0 uint32 }

func main() {
    v1 := make(map[uint32]uint32)
    var v2 string = "hello"
    var v3 uint16 = 32767
    v4 := make(map[uint32]uint32)
    if ((v2 + ("\n\t" + v2)) > v2) {
        v5 := GenS1{F0: -2147483648}
        v6 := GenS2{F0: v3, F1: 2, F2: -486}
        v2 = v2
        _ = v5
        _ = v6
    }
    v7 := GenS2{F0: v3, F1: 2147483647, F2: 514}
    for fi8 := 0; fi8 < 3; fi8++ {
        v3 += v3
        var v9 uint8 = 127
        _ = fi8
        _ = v9
    }
    v10 := GenS2{F0: v3, F1: 4294967295, F2: 9223372036854775807}
    fmt.Println(v3, v3)
    fmt.Println(v2, v2)
    v3 += v3
    if false {
        fmt.Println(v3, v3)
        var v11 int16 = -32768
        if false {
            v12 := GenS1{F0: -1}
            v13 := GenS2{F0: 2, F1: 592, F2: 42}
            if true {
                if false {
                    v2 = ("hello" + v2)
                    v2 = ((v2 + "foo bar") + v2)
                }
                switch v2 {
                case "\n\t":
                    var v14 int16 = (v11 & (v11 * v11))
                    _ = v14
                case "":
                    v15 := GenS2{F0: (^(v3)), F1: 114, F2: 1}
                    _ = v15
                default:
                    for fi16 := 0; fi16 < 12; fi16++ {
                        v17 := GenS1{F0: 4294967295}
                        v18 := []int{4294967295, -731, 1, (fi16 * fi16)}
                        if true {
                            v11 += (^(v11))
                        }
                        _ = fi16
                        _ = v17
                        _ = v18
                    }
                }
                var v19 uint64 = 0
                _ = v19
            }
            _ = v12
            _ = v13
        }
        _ = v11
    } else {
        for fi20 := 0; fi20 < 11; fi20++ {
            var v21 string = (v2 + ("foo bar" + v2))
            var v22 uint32 = 0
            v23 := v4[(^(v22))]
            _ = fi20
            _ = v21
            _ = v22
            _ = v23
        }
        if (30 < 255) {
            v24 := v4[2147483648]
            v24 += (v24 << ((4294967295) & 63))
            _ = v24
        }
        v3 -= 255
    }
    v2 = (("\n\t" + v2) + "hello")
    var v25 bool = true
    fmt.Println(v25)
    v25 = (v3 >= (^(v3)))
    if v25 {
        for fi26 := 0; fi26 < 10; fi26++ {
            if v25 {
                var v27 bool = (!true)
                for fi28 := 0; fi28 < 7; fi28++ {
                    v3 = v3
                    _ = fi28
                }
                for fi29 := 0; fi29 < 1; fi29++ {
                    var v30 bool = (2147483647 == 2)
                    var v31 string = ""
                    _ = fi29
                    _ = v30
                    _ = v31
                }
                _ = v27
            } else {
                var v32 string = ("" + v2)
                delete(v4, 2147483648)
                _ = v32
            }
            _ = fi26
        }
    }
    v33 := GenS1{F0: 9223372036854775807}
    var v34 int = -1
    for fi35 := 0; fi35 < 9; fi35++ {
        v2 = "a"
        v34 = 9223372036854775807
        if ((42 >= -9223372036854775808) && ((v25 && ((v25 || true) || (-1 >= 646))) || (1 < 2))) {
            switch ((v2 + ((v2 + v2) + "hello")) + "hello") {
            case "foo bar":
                for fi36 := 0; fi36 < 13; fi36++ {
                    if (0 == -1) {
                        switch v2 {
                        case "a":
                            var v37 int = (v34 >> ((745) & 63))
                            _ = v37
                        case "\n\t":
                            if v25 {
                                v38 := GenS1{F0: (fi36 >> ((478) & 63))}
                                v34 -= (fi36 >> ((18446744073709551615) & 63))
                                _ = v38
                            } else {
                                var v39 uint64 = 812
                                if v25 {
                                    switch (fi35 >> ((65535) & 63)) {
                                    case 409:
                                        if v25 {
                                            var v40 bool = false
                                            fmt.Println(v25, v25, v39)
                                            _ = v40
                                        } else {
                                            var v41 uint64 = 690
                                            v39 -= (v41 - v41)
                                            v25 = true
                                            _ = v41
                                        }
                                    case -2147483648:
                                        if (v25 || v25) {
                                            fmt.Println(fi36)
                                        } else {
                                            fmt.Println(v39, v3)
                                        }
                                    case 910:
                                        switch (("a" + v2) + v2) {
                                        case "a":
                                            var v42 int8 = -119
                                            _ = v42
                                        }
                                    default:
                                        if false {
                                            v43 := GenS3{F0: 2, F1: 211}
                                            v3 = (v3 ^ (v3 | (v3 / (((v3 % ((v3) | 1))) | 1))))
                                            v44 := GenS4{F0: 42, F1: (v3 + v3)}
                                            _ = v43
                                            _ = v44
                                        }
                                    }
                                    delete(v1, 42)
                                    v34 = fi35
                                }
                                _ = v39
                            }
                        case "hello":
                            switch (fi36 << ((42) & 63)) {
                            case 4294967295:
                                fmt.Println(fi35)
                            default:
                                _ = len(v1)
                                v45 := []int{fi36, 42}
                                _ = v45
                            }
                        }
                        v46 := []int{(^(fi36)), -2147483648, 690, fi36}
                        if (213 >= 1) {
                            v47 := GenS1{F0: v34}
                            fmt.Println(v25, v2)
                            _ = v47
                        }
                        switch (v34 / (((^(fi35))) | 1)) {
                        case 2147483647:
                            switch "\n\t" {
                            case "hello":
                                if v25 {
                                    if true {
                                        v48 := GenS4{F0: -1, F1: (v3 >> ((1) & 63))}
                                        if true {
                                            fmt.Println(v3, fi36)
                                            if true {
                                                var v49 uint = 2
                                                _ = v49
                                            }
                                            v50 := GenS5{F0: 1, F1: -1, F2: int64(v3)}
                                            _ = v50
                                        }
                                        var v51 uint16 = 711
                                        _ = v48
                                        _ = v51
                                    }
                                } else {
                                    fmt.Println(fi35)
                                    v1[2] = 6
                                }
                            case "":
                                v25 = ((0 <= 1) || (true || true))
                            case "foo bar":
                                v52 := GenS5{F0: 279, F1: 73, F2: 293}
                                _ = v52
                            }
                        case 9223372036854775807:
                            v4[0] = 4294967295
                        case -9223372036854775808:
                            var v53 int64 = 0
                            _ = v53
                        default:
                            var v54 uint64 = 1
                            _ = v54
                        }
                        _ = v46
                    }
                    v25 = false
                    v34 = fi36
                    fmt.Println(v2)
                    _ = fi36
                }
            }
            v34 = (v34 - 2147483647)
        } else {
            switch fi35 {
            case 366:
                var v55 int = 649
                _ = v55
            default:
                v56 := GenS1{F0: (-(v34))}
                _ = v56
            }
            fmt.Println(v25, v2)
            for fi57 := 0; fi57 < 11; fi57++ {
                v58 := GenS6{F0: (2 > 915), F1: 42}
                if v25 {
                    v25 = v25
                    var v59 uint16 = (v3 | (v3 | (^(v3))))
                    if (0 >= 1) {
                        var v60 int = fi57
                        var v61 int64 = -1
                        _ = v60
                        _ = v61
                    } else {
                        if ((82 != 127) || true) {
                            if v25 {
                                fmt.Println(v2, fi57)
                            } else {
                                fmt.Println(v59)
                                v62 := GenS7{F0: fi57, F1: 1, F2: 42}
                                _ = v62
                            }
                        }
                        v34 -= v34
                        var v63 float64 = (float64(3.14159) * float64(-2.25))
                        _ = v63
                    }
                    _ = v59
                } else {
                    fmt.Println(v25, v2, v34)
                    v3 -= 32767
                    if (!true) {
                        v2 = ""
                        fmt.Println(fi35)
                        v64 := GenS4{F0: -403, F1: v3}
                        _ = v64
                    } else {
                        switch fi57 {
                        case -2147483648:
                            switch "hello" {
                            case "a":
                                var v65 uint = 1
                                _ = v65
                            case "\n\t":
                                v66 := GenS8{F0: 4294967295}
                                _ = v66
                            default:
                                if true {
                                    v67 := []int{}
                                    _ = len(v4)
                                    if v25 {
                                        if (!true) {
                                            _ = len(v67)
                                            _ = len(v67)
                                        }
                                    }
                                    _ = v67
                                }
                            }
                        default:
                            var v68 string = "\n\t"
                            _ = v68
                        }
                        var v69 int16 = 42
                        _ = v69
                    }
                    var v70 int8 = 1
                    _ = v70
                }
                switch -999 {
                case 4294967295:
                    v3 -= (v3 * (v3 >> ((886) & 63)))
                case 9223372036854775807:
                    v71 := GenS6{F0: false, F1: 4294967295}
                    _ = v71
                default:
                    _ = len(v4)
                }
                _ = fi57
                _ = v58
            }
        }
        _ = fi35
    }
    var v72 int8 = 127
    if ((!(834 == 0)) && v25) {
        if (2 >= 42) {
            for fi73 := 0; fi73 < 7; fi73++ {
                v72 += v72
                _ = len(v4)
                _ = fi73
            }
            var v74 float64 = 1.5
            _ = v74
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v72:", v72)
    _ = v1
    _ = v10
    _ = v2
    _ = v25
    _ = v3
    _ = v33
    _ = v34
    _ = v4
    _ = v7
    _ = v72
}
