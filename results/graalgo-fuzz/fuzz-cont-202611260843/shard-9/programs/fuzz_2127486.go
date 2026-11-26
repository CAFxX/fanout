package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int64 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int8; F1 int32; F2 uint }
type GenS5 struct { F0 string }
type GenS6 struct { F0 string; F1 int32; F2 bool }
type GenS7 struct { F0 int; F1 uint8 }
type GenS8 struct { F0 int16; F1 uint8 }
type GenS9 struct { F0 uint; F1 string; F2 string }
type GenS10 struct { F0 int8; F1 uint32; F2 int32 }
type GenS11 struct { F0 uint32; F1 string; F2 uint16 }
type GenS12 struct { F0 int32; F1 int8; F2 uint }
type GenS13 struct { F0 uint8 }
type GenS14 struct { F0 int8; F1 uint64 }
type GenS15 struct { F0 uint32 }
func genFn0() int64 {
    var v6 uint8 = 128
    v6 -= v6
    v6 -= (v6 | (v6 | (v6 % ((119) | 1))))
    switch -1 {
    case 9223372036854775807:
        var v7 bool = (60 > -1)
        _ = v7
    case -744:
        fmt.Println(v6)
    }
    _ = v6
    return 42
}

func main() {
    var v1 int64 = 1
    var v2 uint16 = 0
    v3 := []int{127, 717}
    var v4 int16 = 42
    var v5 int32 = 0
    v1 += v1
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v8 string = "foo bar"
    _ = len(v3)
    if false {
        v4 = 42
    }
    if true {
        switch -249 {
        case 372:
            v3 = append(v3, -1)
        default:
            v9 := GenS1{F0: (v2 | v2), F1: v1}
            _ = v9
        }
        var v10 int8 = 74
        _ = v10
    }
    switch 69 {
    case 939:
        v11 := GenS2{F0: (v4 - -150)}
        _ = v11
    case 183:
        v8 = ""
    case 505:
        if (0 == 42) {
            if ((-9223372036854775808 >= 9223372036854775807) || false) {
                v12 := GenS3{F0: 104}
                var v13 uint32 = 2
                switch ("hello" + (v8 + "a")) {
                case "hello":
                    if true {
                        fmt.Println(v8, v13)
                        v3 = append(v3, -1)
                        var v14 uint16 = (v2 | 255)
                        _ = v14
                    } else {
                        var v15 uint = 2
                        _ = v15
                    }
                case "a":
                    for fi16 := 0; fi16 < 8; fi16++ {
                        v3 = append(v3, fi16)
                        _ = fi16
                    }
                case "":
                    if (v2 < v2) {
                        switch v8 {
                        case "":
                            switch (v8 + "\n\t") {
                            case "foo bar":
                                v17 := GenS4{F0: 1, F1: v5, F2: 65535}
                                _ = v17
                            case "a":
                                fmt.Println(v13)
                            }
                        case "foo bar":
                            switch -493 {
                            case 2147483647:
                                fmt.Println(v8, v8, v5)
                            case -9223372036854775808:
                                for fi18 := 0; fi18 < 18; fi18++ {
                                    v5 += 2147483647
                                    _ = fi18
                                }
                            }
                        default:
                            if (true || (4294967295 < 42)) {
                                for fi19 := 0; fi19 < 11; fi19++ {
                                    var v20 string = v8
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                    _ = fi19
                                    _ = v20
                                }
                            }
                        }
                        var v21 int32 = (v5 * v5)
                        v22 := GenS3{F0: 255}
                        _ = v21
                        _ = v22
                    }
                default:
                    var v23 uint64 = 4294967295
                    _ = v23
                }
                _ = v12
                _ = v13
            } else {
                v24 := GenS1{F0: (v2 | v2), F1: (v1 % (((v1 & (v1 ^ (v1 & v1)))) | 1))}
                _ = v24
            }
            for fi25 := 0; fi25 < 9; fi25++ {
                if true {
                    switch ((v8 + ("\n\t" + "foo bar")) + v8) {
                    case "":
                        v26 := GenS5{F0: (v8 + v8)}
                        _ = v26
                    case "\n\t":
                        v4 = 42
                    default:
                        if ((-(fi25)) < 42) {
                            fmt.Println(fi25)
                            if false {
                                v27 := GenS6{F0: (("hello" + (((v8 + "hello") + ((((v8 + "\n\t") + v8) + v8) + (v8 + v8))) + "foo bar")) + v8), F1: v5, F2: false}
                                var v28 uint64 = 0
                                v29 := GenS7{F0: fi25, F1: 83}
                                _ = v27
                                _ = v28
                                _ = v29
                            } else {
                                v4 -= v4
                                v30 := GenS8{F0: (-(v4)), F1: 0}
                                _ = v30
                            }
                            switch fi25 {
                            case 32767:
                                fmt.Println(v2)
                            case -798:
                                for fi31 := 0; fi31 < 1; fi31++ {
                                    v32 := GenS4{F0: 0, F1: 0, F2: 42}
                                    if (false || false) {
                                        switch fi25 {
                                        case 32767:
                                            fmt.Println(fi25, fi25)
                                        case 239:
                                            fmt.Println(v2, v8, fi31)
                                            if false {
                                                v33 := GenS2{F0: 1000}
                                                _ = v33
                                            }
                                            fmt.Println(v8, v8, v8)
                                        default:
                                            v5 = -402
                                        }
                                        v34 := GenS6{F0: "", F1: 1, F2: true}
                                        _ = v34
                                    }
                                    _ = fi31
                                    _ = v32
                                }
                            case 36:
                                if len(v3) > 0 {
                                    v3[0] = fi25
                                }
                            }
                        }
                    }
                    var v35 float64 = 0.0
                    _ = v35
                } else {
                    if ((!false) || false) {
                        var v36 int16 = (v4 * (v4 - int16(v2)))
                        var v37 uint = 255
                        switch "foo bar" {
                        case "\n\t":
                            switch v8 {
                            case "":
                                v38 := GenS9{F0: v37, F1: v8, F2: (v8 + v8)}
                                _ = v38
                            case "\n\t":
                                var v39 int64 = genFn0()
                                _ = v39
                            default:
                                v4 += v36
                            }
                        default:
                            var v40 bool = (false || false)
                            _ = v40
                        }
                        _ = v36
                        _ = v37
                    }
                    if (!false) {
                        var v41 int64 = 705
                        _ = v41
                    } else {
                        switch ("foo bar" + "") {
                        case "\n\t":
                            if false {
                                for fi42 := 0; fi42 < 18; fi42++ {
                                    if ((v2 ^ v2) != v2) {
                                        if len(v3) > 0 {
                                            v3[0] = fi42
                                        }
                                    } else {
                                        if false {
                                            v43 := GenS10{F0: 1, F1: 0, F2: v5}
                                            _ = v43
                                        }
                                    }
                                    switch (fi42 - 2147483647) {
                                    case 32767:
                                        v4 += v4
                                    case 4294967295:
                                        v8 = v8
                                    default:
                                        v8 = v8
                                    }
                                    switch v8 {
                                    case "\n\t":
                                        var v44 float64 = 1e10
                                        _ = v44
                                    case "foo bar":
                                        v1 -= v1
                                    case "a":
                                        _ = len(v3)
                                    default:
                                        _ = len(v3)
                                    }
                                    _ = fi42
                                }
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                v3 = append(v3, 4294967295)
                            }
                        case "a":
                            if (v5 == (v5 >> ((18446744073709551615) & 63))) {
                                var v45 uint16 = (v2 << ((990) & 63))
                                v4 -= (v4 | (v4 - (-(v4))))
                                _ = v45
                            }
                        }
                        fmt.Println(v2, v1, fi25)
                        v46 := GenS5{F0: v8}
                        _ = v46
                    }
                    if len(v3) > 0 {
                        v3[0] = 862
                    }
                }
                if len(v3) > 0 {
                    v3[0] = (fi25 * 42)
                }
                _ = fi25
            }
            for fi47 := 0; fi47 < 9; fi47++ {
                switch fi47 {
                case -575:
                    v8 = ("a" + v8)
                default:
                    if ((true || (501 > (v2 * (^(v2))))) || (65535 >= 42)) {
                        v8 = (v8 + "\n\t")
                        for fi48 := 0; fi48 < 9; fi48++ {
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            switch fi47 {
                            case 9223372036854775807:
                                if false {
                                    v49 := GenS7{F0: (fi48 >> ((65535) & 63)), F1: 128}
                                    _ = v49
                                }
                            case 2147483647:
                                var v50 float64 = (0.0 + (0.0 - 1e10))
                                _ = v50
                            }
                            _ = fi48
                        }
                        v1 -= (v1 + (v1 >> ((1) & 63)))
                    }
                }
                for fi51 := 0; fi51 < 5; fi51++ {
                    if false {
                        if (0 == -107) {
                            if (!(((3.14159 / (0.1 - (-2.25 - (-2.25 / 3.14159)))) / 1.5) < -2.25)) {
                                var v52 int = -1
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                if true {
                                    var v53 bool = false
                                    _ = v53
                                } else {
                                    fmt.Println(fi51, fi47)
                                    fmt.Println(v52, v1)
                                }
                                v5 += (v5 % ((2147483647) | 1))
                                v52 = 1
                                _ = v52
                            } else {
                                v4 = (v4 & 132)
                                var v54 uint16 = (v2 + 65535)
                                _ = v54
                            }
                            if false {
                                fmt.Println(v5)
                            }
                            fmt.Println(v1, v2)
                        }
                        fmt.Println(v4)
                        fmt.Println(v1)
                        var v55 uint64 = 1
                        _ = v55
                    } else {
                        if true {
                            v56 := GenS11{F0: 0, F1: v8, F2: 873}
                            _ = v56
                        } else {
                            if true {
                                var v57 float64 = (((0.0 / (0.0 - float64(3.14159))) * 1.5) - 1.5)
                                v58 := GenS12{F0: (v5 & v5), F1: 68, F2: 2}
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                _ = v57
                                _ = v58
                            }
                            fmt.Println(fi51)
                            v3 = append(v3, (fi51 + fi47))
                        }
                        if ((!((v5 + (v5 & v5)) < 187)) || (false || false)) {
                            var v59 string = v8
                            _ = v59
                        } else {
                            v1 -= (v1 ^ (v1 | v1))
                            v5 -= v5
                            var v60 bool = false
                            _ = v60
                        }
                        fmt.Println(v1)
                    }
                    _ = fi51
                }
                if true {
                    v61 := GenS4{F0: 127, F1: v5, F2: 0}
                    _ = v61
                } else {
                    var v62 string = (v8 + v8)
                    _ = v62
                }
                _ = fi47
            }
        }
    default:
        if len(v3) > 0 {
            v3[0] = 4294967295
        }
    }
    if (!((v2 % (((^(v2))) | 1)) > (v2 & (^(v2))))) {
        v63 := GenS11{F0: 42, F1: (v8 + v8), F2: v2}
        v1 += v1
        if len(v3) > 0 {
            v3[0] = -626
        }
        _ = v63
    }
    fmt.Println(v5)
    for fi64 := 0; fi64 < 8; fi64++ {
        v8 = v8
        for fi65 := 0; fi65 < 7; fi65++ {
            v66 := GenS13{F0: 12}
            if false {
                switch "" {
                case "hello":
                    v67 := GenS14{F0: 100, F1: 18446744073709551615}
                    _ = v67
                case "":
                    _ = len(v3)
                default:
                    v68 := GenS14{F0: -99, F1: 2}
                    _ = v68
                }
                v4 -= v4
            } else {
                v8 = v8
                v69 := GenS15{F0: 42}
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                _ = v69
            }
            var v70 uint8 = 127
            _ = fi65
            _ = v66
            _ = v70
        }
        v1 += genFn0()
        _ = fi64
    }
    fmt.Println(v4, v2, v4)
    for fi71 := 0; fi71 < 14; fi71++ {
        v8 = v8
        _ = fi71
    }
    var v72 int32 = v5
    _ = len(v3)
    switch "a" {
    case "\n\t":
        if false {
            fmt.Println(v4)
        } else {
            var v73 int = -9223372036854775808
            _ = v73
        }
    case "":
        v4 = (v4 - -32768)
    case "hello":
        fmt.Println(v2, v72, v2)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v72:", v72)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v72
    _ = v8
}
