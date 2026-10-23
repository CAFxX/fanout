package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint16; F1 uint }
type GenS3 struct { F0 int; F1 int8 }
type GenS4 struct { F0 int64; F1 bool; F2 string }
type GenS5 struct { F0 int8; F1 int32 }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 int32; F1 int32 }
type GenS8 struct { F0 uint; F1 int8; F2 uint }

func main() {
    v1 := []string{"hello", "foo bar", "", "hello"}
    var v2 string = "a"
    var v3 int = 32767
    v4 := []int{-527, v3, (v3 + (v3 * (-(v3))))}
    var v5 float32 = -2.25
    v6 := make(map[int]uint32)
    if true {
        v7 := GenS1{F0: -1}
        _ = v7
    }
    if ((float64(v5) * (-2.25 * 3.14159)) > 0.0) {
        v8 := GenS1{F0: 0}
        v5 += ((-2.25 / ((3.14159 - v5) / (v5 - 1.5))) + 1e10)
        _ = v8
    }
    if true {
        var v9 float64 = 1e10
        switch -696 {
        case 0:
            var v10 int16 = -538
            _ = v10
        case -9223372036854775808:
            v11 := GenS2{F0: 1, F1: 18446744073709551615}
            _ = v11
        }
        v9 -= v9
        _ = v9
    } else {
        var v12 uint16 = 0
        v5 = v5
        _ = v12
    }
    for fi13 := 0; fi13 < 12; fi13++ {
        if (484 == 1) {
            for fi14 := 0; fi14 < 16; fi14++ {
                switch "a" {
                case "\n\t":
                    var v15 uint8 = 128
                    _ = v15
                case "hello":
                    v2 = "a"
                }
                switch "" {
                case "hello":
                    var v16 float64 = -2.25
                    _ = v16
                }
                if len(v4) > 0 {
                    v4[0] = (v3 - (fi13 % (((^(v3))) | 1)))
                }
                _ = fi14
            }
            v2 = v2
        }
        fmt.Println(v5, fi13)
        _ = fi13
    }
    v17 := GenS3{F0: v3, F1: -1}
    var v18 int32 = 0
    switch (-(v3)) {
    case 1:
        fmt.Println(v18, v2)
    case 32767:
        v19 := GenS2{F0: 857, F1: 67}
        _ = v19
    case 517:
        v6[-9223372036854775808] = 4294967295
    default:
        for fi20 := 0; fi20 < 20; fi20++ {
            fmt.Println(v18, v3, v5)
            _ = len(v1)
            _ = fi20
        }
    }
    var v21 float32 = (v5 - v5)
    v2 = v2
    if (v18 < -1) {
        var v22 uint32 = 711
        var v23 int8 = 42
        switch v3 {
        case 4294967295:
            _ = len(v4)
        case 127:
            switch -1 {
            case 313:
                switch 32767 {
                case 2147483647:
                    v24 := GenS2{F0: 326, F1: 629}
                    _ = v24
                case 1:
                    fmt.Println(v22, v3, v5)
                case 32767:
                    v23 -= (v23 << ((65535) & 63))
                default:
                    v18 -= (v18 / ((v18) | 1))
                }
            default:
                var v25 uint32 = v22
                _ = v25
            }
        }
        _ = v22
        _ = v23
    } else {
        _ = len(v6)
    }
    if (78 <= -1) {
        if (!(v5 <= v21)) {
            if (-9223372036854775808 != 1) {
                delete(v6, -335)
                var v26 int = v3
                v18 -= int32(v26)
                _ = v26
            } else {
                v6[(-(v3))] = 4294967295
                switch v2 {
                case "a":
                    if false {
                        v27 := GenS1{F0: 82}
                        _ = v27
                    }
                case "foo bar":
                    v2 = (v2 + ("hello" + (v2 + v2)))
                case "\n\t":
                    v2 = ""
                }
                v5 = (v21 / 1.5)
            }
            if (1.5 != 1e10) {
                if false {
                    switch ("hello" + "hello") {
                    case "\n\t":
                        switch v2 {
                        case "\n\t":
                            v6[(^(v3))] = 42
                        case "":
                            var v28 bool = (18446744073709551615 >= 65535)
                            _ = v28
                        case "a":
                            if false {
                                var v29 int8 = 0
                                _ = v29
                            }
                        }
                    default:
                        v30 := GenS4{F0: 0, F1: false, F2: "a"}
                        _ = v30
                    }
                    if (v18 > (v18 - (-(v18)))) {
                        if len(v4) > 0 {
                            v4[0] = (v3 & (v3 % (((v3 | v3)) | 1)))
                        }
                        for fi31 := 0; fi31 < 1; fi31++ {
                            var v32 int = v3
                            v33 := GenS5{F0: -128, F1: 2147483647}
                            if true {
                                for fi34 := 0; fi34 < 18; fi34++ {
                                    v32 += fi31
                                    if false {
                                        v35 := GenS1{F0: 0}
                                        _ = v35
                                    } else {
                                        if ((false || (65535 <= 0)) || true) {
                                            fmt.Println(fi31, v2, v2)
                                            v21 = v5
                                        } else {
                                            var v36 bool = true
                                            v36 = true
                                            _ = len(v4)
                                            v37 := GenS6{F0: v36}
                                            _ = v36
                                            _ = v37
                                        }
                                        var v38 int64 = 804
                                        v38 += 42
                                        var v39 int64 = 42
                                        _ = v38
                                        _ = v39
                                    }
                                    if (42 != 1) {
                                        v21 -= (-2.25 - (0.1 - -2.25))
                                    } else {
                                        v40 := GenS5{F0: 127, F1: v18}
                                        if (1 <= 8) {
                                            var v41 uint16 = 2
                                            _ = v41
                                        }
                                        if true {
                                            if len(v1) > 0 {
                                                v1[0] = (v2 + (v2 + v2))
                                            }
                                        } else {
                                            _ = len(v6)
                                            if (32767 > 0) {
                                                var v42 string = "foo bar"
                                                var v43 uint8 = 124
                                                fmt.Println(v32, v5)
                                                if true {
                                                    switch v2 {
                                                    case "foo bar":
                                                        var v44 uint64 = 42
                                                        _ = v44
                                                    case "hello":
                                                        switch v42 {
                                                        case "hello":
                                                            var v45 int16 = 0
                                                            _ = v45
                                                        case "foo bar":
                                                            delete(v6, (^(v3)))
                                                        default:
                                                            if false {
                                                                var v46 int64 = 42
                                                                var v47 int64 = v46
                                                                v48 := GenS3{F0: v32, F1: 0}
                                                                _ = v46
                                                                _ = v47
                                                                _ = v48
                                                            } else {
                                                                if (((v21 * 0.1) / v21) <= v5) {
                                                                    v49 := GenS3{F0: v32, F1: 15}
                                                                    var v50 string = v2
                                                                    _ = v49
                                                                    _ = v50
                                                                }
                                                                var v51 int32 = 434
                                                                var v52 int = v32
                                                                _ = v51
                                                                _ = v52
                                                            }
                                                        }
                                                        var v53 bool = ((-128 >= 127) && (0 >= 127))
                                                        _ = v53
                                                    case "":
                                                        v54 := GenS2{F0: 379, F1: 0}
                                                        _ = v54
                                                    }
                                                    v55 := GenS7{F0: (v18 << ((65535) & 63)), F1: (v18 ^ (v18 >> ((2) & 63)))}
                                                    v56 := GenS1{F0: 127}
                                                    _ = v55
                                                    _ = v56
                                                }
                                                v57 := GenS5{F0: 42, F1: v18}
                                                _ = v42
                                                _ = v43
                                                _ = v57
                                            }
                                            var v58 float64 = float64((v5 / v21))
                                            v59 := GenS3{F0: (v3 | 1), F1: 92}
                                            _ = v58
                                            _ = v59
                                        }
                                        _ = v40
                                    }
                                    _ = fi34
                                }
                                var v60 int16 = -541
                                v3 = (v3 ^ (^(fi31)))
                                _ = v60
                            }
                            _ = fi31
                            _ = v32
                            _ = v33
                        }
                        for fi61 := 0; fi61 < 14; fi61++ {
                            var v62 int = (v3 - (^(fi61)))
                            _ = fi61
                            _ = v62
                        }
                    } else {
                        var v63 uint = 5
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        v6[v3] = 65535
                        _ = v63
                    }
                    for fi64 := 0; fi64 < 9; fi64++ {
                        v65 := v6[-172]
                        v5 -= v5
                        _ = fi64
                        _ = v65
                    }
                }
                var v66 float64 = (-2.25 + 0.1)
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                _ = v66
            } else {
                var v67 int16 = -369
                _ = v67
            }
            var v68 float64 = 0.1
            _ = v68
        }
        _ = len(v6)
    } else {
        var v69 uint8 = 128
        v69 = (v69 | (v69 * (^(v69))))
        for fi70 := 0; fi70 < 2; fi70++ {
            if false {
                var v71 int32 = v18
                _ = v71
            } else {
                switch v2 {
                case "\n\t":
                    v5 += float32(0.1)
                case "foo bar":
                    var v72 bool = false
                    _ = v72
                case "hello":
                    v2 = (("a" + v2) + ("hello" + (v2 + "")))
                default:
                    v73 := GenS8{F0: 1, F1: -1, F2: 338}
                    _ = v73
                }
                v3 += v3
                fmt.Println(v5, v18, v21)
            }
            for fi74 := 0; fi74 < 3; fi74++ {
                v3 = v3
                _ = fi74
            }
            _ = fi70
        }
        _ = v69
    }
    delete(v6, -9223372036854775808)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
