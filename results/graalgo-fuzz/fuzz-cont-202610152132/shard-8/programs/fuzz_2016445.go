package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 int; F1 uint }
type GenS6 struct { F0 uint64; F1 uint64; F2 uint64 }
type GenS7 struct { F0 uint64; F1 string }
type GenS8 struct { F0 uint; F1 uint64 }
type GenS9 struct { F0 bool; F1 uint8; F2 int }
type GenS10 struct { F0 uint16; F1 uint8; F2 uint8 }
type GenS11 struct { F0 string; F1 int16; F2 int64 }
type GenS12 struct { F0 int32; F1 int8 }

func main() {
    var v1 int8 = -107
    var v2 int8 = (v1 - -128)
    var v3 bool = (1 <= -2147483648)
    var v4 int16 = 42
    var v5 uint8 = 0
    var v6 uint32 = 978
    v3 = v3
    for fi7 := 0; fi7 < 19; fi7++ {
        if false {
            for fi8 := 0; fi8 < 17; fi8++ {
                var v9 float32 = float32(1.5)
                fmt.Println(v6)
                switch "" {
                case "\n\t":
                    v10 := GenS1{F0: v6}
                    _ = v10
                case "foo bar":
                    fmt.Println(v6)
                case "a":
                    var v11 uint8 = v5
                    _ = v11
                default:
                    fmt.Println(v9, v2, v9)
                }
                _ = fi8
                _ = v9
            }
            v4 = (v4 & v4)
        } else {
            var v12 float64 = 0.1
            _ = v12
        }
        _ = fi7
    }
    if (!(2 != 56)) {
        v13 := GenS1{F0: v6}
        if v3 {
            v14 := []int{9223372036854775807, 2147483647}
            var v15 int64 = 0
            v4 -= (v4 + int16(v1))
            _ = v14
            _ = v15
        }
        _ = v13
    }
    switch (("a" + "a") + (("foo bar" + "a") + "")) {
    case "\n\t":
        for fi16 := 0; fi16 < 14; fi16++ {
            switch fi16 {
            case 407:
                var v17 int = fi16
                _ = v17
            case 128:
                for fi18 := 0; fi18 < 7; fi18++ {
                    if (-9223372036854775808 == 738) {
                        v19 := GenS1{F0: v6}
                        var v20 int32 = -908
                        _ = v19
                        _ = v20
                    } else {
                        var v21 int64 = 9223372036854775807
                        _ = v21
                    }
                    v6 -= 1
                    _ = fi18
                }
            case 402:
                v6 = v6
            default:
                v22 := GenS1{F0: (v6 & v6)}
                _ = v22
            }
            _ = fi16
        }
    }
    if v3 {
        v23 := GenS2{F0: v4}
        v4 += (v4 - 120)
        var v24 int8 = (v1 & v1)
        _ = v23
        _ = v24
    }
    switch "hello" {
    case "a":
        for fi25 := 0; fi25 < 9; fi25++ {
            fmt.Println(v5, fi25)
            _ = fi25
        }
    case "\n\t":
        var v26 int = -183
        _ = v26
    }
    for fi27 := 0; fi27 < 8; fi27++ {
        for fi28 := 0; fi28 < 6; fi28++ {
            var v29 int16 = v4
            _ = fi28
            _ = v29
        }
        _ = fi27
    }
    if v3 {
        for fi30 := 0; fi30 < 4; fi30++ {
            v1 += v1
            _ = fi30
        }
        if (!v3) {
            var v31 uint64 = 0
            fmt.Println(v1, v5, v3)
            _ = v31
        } else {
            switch "foo bar" {
            case "foo bar":
                if (!false) {
                    var v32 uint64 = 844
                    var v33 bool = ((1e10 - 3.14159) <= (float64(3.14159) / (-2.25 * (0.1 * -2.25))))
                    _ = v32
                    _ = v33
                }
            default:
                fmt.Println(v6)
            }
            if v3 {
                v34 := GenS3{F0: 465}
                _ = v34
            } else {
                if v3 {
                    for fi35 := 0; fi35 < 16; fi35++ {
                        v36 := []int{}
                        v37 := GenS3{F0: uint(v6)}
                        _ = fi35
                        _ = v36
                        _ = v37
                    }
                    for fi38 := 0; fi38 < 9; fi38++ {
                        if ((0.0 / 1.0) <= -2.25) {
                            v39 := []int{(fi38 / (((-(fi38))) | 1))}
                            _ = v39
                        } else {
                            switch "foo bar" {
                            case "":
                                for fi40 := 0; fi40 < 13; fi40++ {
                                    v3 = true
                                    _ = fi40
                                }
                            case "foo bar":
                                if v3 {
                                    v4 = v4
                                    var v41 int8 = (v2 << ((529) & 63))
                                    _ = v41
                                }
                            case "\n\t":
                                v42 := []int{128}
                                _ = v42
                            default:
                                var v43 int64 = -9223372036854775808
                                _ = v43
                            }
                            v6 = (v6 | v6)
                        }
                        _ = fi38
                    }
                } else {
                    fmt.Println(v4, v3, v6)
                    v4 -= (v4 * (v4 >> ((95) & 63)))
                }
                v44 := GenS4{F0: 561}
                _ = v44
            }
        }
    } else {
        for fi45 := 0; fi45 < 3; fi45++ {
            var v46 uint32 = (v6 & v6)
            v47 := []int{}
            if (!(1 == -1)) {
                fmt.Println(v4)
                for fi48 := 0; fi48 < 9; fi48++ {
                    v4 = v4
                    if (uint(v46) <= 18446744073709551615) {
                        if len(v47) >= 2 {
                            v47 = v47[1:len(v47)]
                        }
                    }
                    if v3 {
                        v5 = v5
                        v46 += (v6 & (v6 * (v6 - uint32(v5))))
                    } else {
                        if len(v47) >= 2 {
                            v47 = v47[1:len(v47)]
                        }
                        if (v3 || (v4 < (v4 + (v4 | 795)))) {
                            v49 := GenS5{F0: 32768, F1: 18446744073709551615}
                            _ = v49
                        }
                        v50 := GenS3{F0: 1}
                        _ = v50
                    }
                    _ = fi48
                }
                if (false || (!v3)) {
                    v3 = false
                    if len(v47) > 0 {
                        v47[0] = -1
                    }
                    v6 = 1
                }
            } else {
                if (v6 == (v6 | 2)) {
                    v5 += 29
                    for fi51 := 0; fi51 < 2; fi51++ {
                        v5 -= 128
                        fmt.Println(v5, v1)
                        if v3 {
                            if v3 {
                                if v3 {
                                    var v52 int32 = -1
                                    var v53 float32 = 1e10
                                    _ = v52
                                    _ = v53
                                } else {
                                    if ((3.14159 >= 3.14159) && v3) {
                                        switch ("foo bar" + ("hello" + "\n\t")) {
                                        case "":
                                            _ = len(v47)
                                        case "foo bar":
                                            v4 = v4
                                        default:
                                            if (36 >= 4294967295) {
                                                fmt.Println(v46)
                                                if len(v47) >= 2 {
                                                    v47 = v47[1:len(v47)]
                                                }
                                            } else {
                                                switch ("\n\t" + "") {
                                                case "a":
                                                    if v3 {
                                                        if (fi51 >= 9223372036854775807) {
                                                            v4 += (v4 << ((255) & 63))
                                                        }
                                                        var v54 int32 = -657
                                                        _ = v54
                                                    }
                                                case "foo bar":
                                                    var v55 string = "foo bar"
                                                    _ = v55
                                                default:
                                                    v6 += v6
                                                }
                                            }
                                        }
                                        v1 += v2
                                    }
                                    v6 += (v6 >> ((42) & 63))
                                    var v56 int = fi45
                                    _ = v56
                                }
                                var v57 int64 = -9223372036854775808
                                _ = v57
                            } else {
                                v1 += 73
                            }
                            v1 = v1
                        }
                        fmt.Println(v4)
                        v46 -= 2147483648
                        _ = fi51
                    }
                    for fi58 := 0; fi58 < 2; fi58++ {
                        var v59 uint = 4294967295
                        v2 -= v1
                        var v60 int8 = v1
                        _ = fi58
                        _ = v59
                        _ = v60
                    }
                }
                if true {
                    var v61 int8 = 0
                    if (((!(32767 <= 935)) || (((fi45 ^ (fi45 | (fi45 >> ((1) & 63)))) > fi45) || (v2 != v61))) || v3) {
                        v62 := GenS6{F0: 1, F1: 2, F2: 42}
                        fmt.Println(fi45, v1)
                        var v63 uint16 = 16
                        _ = v62
                        _ = v63
                    } else {
                        var v64 int64 = 551
                        _ = v64
                    }
                    _ = v61
                } else {
                    if v3 {
                        switch "hello" {
                        case "a":
                            for fi65 := 0; fi65 < 8; fi65++ {
                                if true {
                                    var v66 string = "\n\t"
                                    v67 := GenS4{F0: 9223372036854775807}
                                    var v68 string = "foo bar"
                                    _ = v66
                                    _ = v67
                                    _ = v68
                                }
                                var v69 uint64 = 18446744073709551615
                                _ = fi65
                                _ = v69
                            }
                        case "":
                            for fi70 := 0; fi70 < 16; fi70++ {
                                v71 := GenS7{F0: 348, F1: "hello"}
                                if (float64((3.14159 + 1e10)) == 1e10) {
                                    if (fi70 > fi45) {
                                        v2 -= 127
                                        v72 := GenS8{F0: 42, F1: 42}
                                        v73 := GenS9{F0: v3, F1: v5, F2: fi70}
                                        _ = v72
                                        _ = v73
                                    } else {
                                        v47 = append(v47, (fi70 - -1))
                                        v1 = (v1 | v2)
                                        v1 -= v2
                                    }
                                    var v74 float64 = 1.5
                                    v75 := GenS7{F0: 425, F1: "\n\t"}
                                    v76 := GenS10{F0: 581, F1: uint8(v6), F2: 2}
                                    var v77 int32 = 1
                                    _ = v74
                                    _ = v75
                                    _ = v76
                                    _ = v77
                                }
                                _ = fi70
                                _ = v71
                            }
                        case "foo bar":
                            if v3 {
                                fmt.Println(v5, v2)
                                v78 := GenS11{F0: "", F1: v4, F2: 9223372036854775807}
                                _ = v78
                            }
                        }
                    } else {
                        var v79 uint8 = v5
                        fmt.Println(v6, v4)
                        _ = v79
                    }
                }
            }
            _ = fi45
            _ = v46
            _ = v47
        }
        var v80 int64 = 0
        _ = v80
    }
    switch "\n\t" {
    case "":
        var v81 float64 = (float64(float32((1e10 - (1e10 + ((float64(0.0) - (-2.25 * float64((3.14159 / (((1.5 - 0.1) / 1.5) + 1e10))))) / -2.25))))) * (1e10 * 3.14159))
        _ = v81
    case "hello":
        v82 := GenS12{F0: 1, F1: (v2 / ((v1) | 1))}
        _ = v82
    case "foo bar":
        switch int(v1) {
        case 42:
            fmt.Println(v6, v4, v5)
        case 2147483647:
            var v83 uint = 1
            _ = v83
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
