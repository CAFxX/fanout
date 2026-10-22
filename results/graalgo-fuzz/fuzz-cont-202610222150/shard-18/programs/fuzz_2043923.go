package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool }
type GenS2 struct { F0 int }
type GenS3 struct { F0 string; F1 uint32; F2 int8 }
type GenS4 struct { F0 uint64; F1 uint; F2 int8 }
type GenS5 struct { F0 uint16; F1 uint16 }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 int8; F1 uint16; F2 int }
type GenS8 struct { F0 int; F1 int8; F2 int32 }
type GenS9 struct { F0 uint8; F1 uint; F2 int8 }
type GenS10 struct { F0 int }
func genFn0(p0 int, p1 int64) uint32 {
    v5 := GenS1{F0: 255, F1: false}
    var v6 uint64 = 2
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return 2147483647
}
func genFn1(p0 uint16, p1 uint8) uint64 {
    fmt.Println(p1, p0)
    p0 -= 219
    for fi7 := 0; fi7 < 16; fi7++ {
        if true {
            if ("hello" < "\n\t") {
                fmt.Println(p1, fi7, p1)
                v8 := GenS1{F0: 2, F1: true}
                v9 := GenS2{F0: 671}
                _ = v8
                _ = v9
            }
            v10 := GenS3{F0: "a", F1: 364, F2: 84}
            _ = v10
        } else {
            if true {
                switch fi7 {
                case -311:
                    v11 := GenS1{F0: 42, F1: (100 < -31)}
                    _ = v11
                default:
                    var v12 float64 = 1e10
                    _ = v12
                }
                switch "hello" {
                case "\n\t":
                    v13 := GenS4{F0: 0, F1: 990, F2: -128}
                    _ = v13
                case "":
                    v14 := []int{fi7, (-(fi7))}
                    _ = v14
                }
                v15 := GenS3{F0: ("" + "\n\t"), F1: 42, F2: 0}
                _ = v15
            } else {
                switch (fi7 - fi7) {
                case 0:
                    var v16 int16 = 0
                    _ = v16
                case 694:
                    p0 = p0
                case 42:
                    var v17 float32 = 0.1
                    _ = v17
                default:
                    for fi18 := 0; fi18 < 19; fi18++ {
                        v19 := GenS5{F0: 2, F1: uint16(p1)}
                        var v20 uint8 = (p1 + p1)
                        _ = fi18
                        _ = v19
                        _ = v20
                    }
                }
                for fi21 := 0; fi21 < 5; fi21++ {
                    v22 := GenS6{F0: -1}
                    p0 = 1
                    var v23 uint = 2
                    _ = fi21
                    _ = v22
                    _ = v23
                }
                v24 := GenS3{F0: ("\n\t" + "\n\t"), F1: genFn0(fi7, -1), F2: -1}
                _ = v24
            }
            p1 = p1
            fmt.Println(fi7, p1)
        }
        _ = fi7
    }
    _ = p0
    _ = p1
    return 237
}

func main() {
    v1 := []uint32{}
    var v2 uint = 0
    var v3 uint8 = 133
    var v4 string = ""
    if (1 >= -32768) {
        for fi25 := 0; fi25 < 17; fi25++ {
            var v26 float64 = float64(1e10)
            _ = fi25
            _ = v26
        }
        for fi27 := 0; fi27 < 20; fi27++ {
            for fi28 := 0; fi28 < 6; fi28++ {
                v1 = append(v1, 65535)
                v2 += v2
                var v29 uint = v2
                if false {
                    v3 -= (v3 - (v3 / ((v3) | 1)))
                    var v30 int32 = 2147483647
                    if true {
                        switch (v4 + "") {
                        case "\n\t":
                            if len(v1) > 0 {
                                v1[0] = 42
                            }
                        case "":
                            if (("" + "") >= v4) {
                                var v31 uint8 = (v3 + uint8(v2))
                                _ = v31
                            } else {
                                _ = len(v1)
                                var v32 float64 = ((0.0 + (1e10 + -2.25)) * 1e10)
                                _ = v32
                            }
                        default:
                            if (!((true || true) && ((v3 | (v3 + v3)) != (^(v3))))) {
                                if len(v1) > 0 {
                                    v1[0] = 0
                                }
                            } else {
                                var v33 uint16 = 2
                                _ = v33
                            }
                        }
                    } else {
                        var v34 bool = false
                        v35 := GenS7{F0: int8(fi27), F1: 856, F2: -9223372036854775808}
                        v30 = v30
                        _ = v34
                        _ = v35
                    }
                    _ = v30
                } else {
                    var v36 int = fi27
                    _ = v36
                }
                v29 -= 42
                _ = fi28
                _ = v29
            }
            _ = fi27
        }
        if ((2 >= genFn1(0, 177)) || (169 >= 18446744073709551615)) {
            v1 = append(v1, 349)
            v1 = append(v1, 2)
        }
    } else {
        fmt.Println(v2)
        v4 = (v4 + (("hello" + ("a" + (v4 + "a"))) + v4))
    }
    var v37 uint16 = 0
    v38 := GenS7{F0: 1, F1: v37, F2: 42}
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v39 := GenS1{F0: v2, F1: true}
    for fi40 := 0; fi40 < 7; fi40++ {
        fmt.Println(v2, v4)
        for fi41 := 0; fi41 < 16; fi41++ {
            _ = len(v1)
            v42 := GenS2{F0: fi40}
            _ = fi41
            _ = v42
        }
        _ = fi40
    }
    v37 -= (v37 | v37)
    var v43 int8 = -41
    switch "\n\t" {
    case "":
        switch -1 {
        case -9223372036854775808:
            for fi44 := 0; fi44 < 3; fi44++ {
                fmt.Println(v4)
                _ = fi44
            }
        case 670:
            if len(v1) > 0 {
                v1[0] = 1
            }
        case 482:
            v2 += 4294967295
        }
    case "\n\t":
        if ((878 == -9223372036854775808) && (true || (!(!(int32(v3) != -165))))) {
            if false {
                if true {
                    switch v4 {
                    case "":
                        var v45 int = 127
                        _ = v45
                    case "hello":
                        switch 42 {
                        case 42:
                            v46 := GenS8{F0: -2147483648, F1: -1, F2: 47}
                            _ = v46
                        default:
                            for fi47 := 0; fi47 < 8; fi47++ {
                                v4 = ("hello" + (("foo bar" + (((v4 + (v4 + (v4 + ("foo bar" + ("" + "\n\t"))))) + v4) + "foo bar")) + "a"))
                                v43 = 0
                                _ = fi47
                            }
                        }
                    case "foo bar":
                        v3 = (v3 >> ((65535) & 63))
                    }
                    if true {
                        v48 := GenS7{F0: v43, F1: v37, F2: 32767}
                        var v49 int = -1
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        _ = v48
                        _ = v49
                    }
                } else {
                    fmt.Println(v3, v3, v2)
                    v50 := GenS5{F0: 2, F1: uint16(v2)}
                    _ = len(v1)
                    _ = v50
                }
            } else {
                fmt.Println(v37, v43, v2)
                var v51 uint = (v2 % ((4294967295) | 1))
                v52 := GenS2{F0: 976}
                _ = v51
                _ = v52
            }
        }
    }
    for fi53 := 0; fi53 < 8; fi53++ {
        var v54 float64 = (0.0 - -2.25)
        for fi55 := 0; fi55 < 1; fi55++ {
            v3 -= v3
            _ = fi55
        }
        _ = fi53
        _ = v54
    }
    var v56 uint16 = 42
    v4 = (v4 + v4)
    fmt.Println(v56)
    if true {
        v57 := GenS3{F0: "", F1: 653, F2: (^(v43))}
        v58 := GenS4{F0: 42, F1: 255, F2: v43}
        for fi59 := 0; fi59 < 2; fi59++ {
            var v60 bool = true
            _ = fi59
            _ = v60
        }
        _ = v57
        _ = v58
    } else {
        v43 -= (v43 >> (((v2 / (((v2 << ((v2) & 63))) | 1))) & 63))
        if (!false) {
            if (!true) {
                if ((2147483647 > 2147483647) || false) {
                    var v61 uint8 = (v3 * v3)
                    fmt.Println(v2, v3)
                    _ = v61
                } else {
                    fmt.Println(v2, v4, v37)
                    for fi62 := 0; fi62 < 18; fi62++ {
                        switch v4 {
                        case "\n\t":
                            var v63 uint = v2
                            _ = v63
                        case "a":
                            switch v4 {
                            case "":
                                switch "foo bar" {
                                case "":
                                    v3 -= uint8(v56)
                                case "hello":
                                    var v64 string = ("hello" + "a")
                                    _ = v64
                                default:
                                    if true {
                                        switch 128 {
                                        case -9223372036854775808:
                                            v65 := GenS9{F0: v3, F1: (^(v2)), F2: (v43 << ((1) & 63))}
                                            _ = v65
                                        case 9223372036854775807:
                                            v56 -= (v37 | v37)
                                        }
                                        for fi66 := 0; fi66 < 6; fi66++ {
                                            var v67 int32 = 1
                                            _ = fi66
                                            _ = v67
                                        }
                                        if false {
                                            var v68 bool = true
                                            var v69 uint32 = 2147483648
                                            _ = v68
                                            _ = v69
                                        }
                                    }
                                }
                            default:
                                for fi70 := 0; fi70 < 3; fi70++ {
                                    fmt.Println(fi70, v3, v37)
                                    _ = fi70
                                }
                            }
                        case "foo bar":
                            fmt.Println(fi62)
                        }
                        v3 -= (v3 / ((v3) | 1))
                        for fi71 := 0; fi71 < 3; fi71++ {
                            v72 := GenS5{F0: v37, F1: v37}
                            v2 += (^(v2))
                            switch fi71 {
                            case 4294967295:
                                fmt.Println(fi71)
                            }
                            _ = fi71
                            _ = v72
                        }
                        _ = fi62
                    }
                    var v73 uint16 = 369
                    _ = v73
                }
                v4 = ((v4 + ((v4 + "hello") + v4)) + v4)
            }
            for fi74 := 0; fi74 < 1; fi74++ {
                var v75 float32 = float32(0.0)
                v76 := GenS4{F0: 4294967295, F1: v2, F2: v43}
                var v77 int8 = v43
                _ = fi74
                _ = v75
                _ = v76
                _ = v77
            }
            for fi78 := 0; fi78 < 8; fi78++ {
                for fi79 := 0; fi79 < 6; fi79++ {
                    var v80 string = v4
                    var v81 string = v80
                    _ = fi79
                    _ = v80
                    _ = v81
                }
                for fi82 := 0; fi82 < 2; fi82++ {
                    var v83 int32 = 1
                    v84 := GenS10{F0: (fi82 / ((fi82) | 1))}
                    v1 = append(v1, 166)
                    _ = fi82
                    _ = v83
                    _ = v84
                }
                if len(v1) > 0 {
                    v1[0] = 65535
                }
                _ = fi78
            }
        }
        var v85 float32 = (((0.0 * 1e10) + -2.25) - 1e10)
        _ = v85
    }
    v86 := GenS1{F0: v2, F1: (float32(((((float64(float32(0.1)) - -2.25) * 0.0) * 3.14159) - 1.5)) > 3.14159)}
    v87 := GenS9{F0: (v3 & (v3 >> ((v2) & 63))), F1: 177, F2: (-(v43))}
    var v88 bool = ((false || (!(32767 != 1))) || false)
    for fi89 := 0; fi89 < 17; fi89++ {
        v1 = append(v1, genFn0(127, 690))
        _ = fi89
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v56:", v56)
    fmt.Println("v88:", v88)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v43
    _ = v56
    _ = v86
    _ = v87
    _ = v88
}
