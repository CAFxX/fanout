package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32 }
type GenS2 struct { F0 int8; F1 int; F2 uint32 }
type GenS3 struct { F0 string; F1 uint8 }
type GenS4 struct { F0 uint; F1 int8; F2 uint32 }
type GenS5 struct { F0 uint32; F1 uint16; F2 uint64 }
type GenS6 struct { F0 uint8; F1 string; F2 uint16 }
type GenS7 struct { F0 uint16 }
type GenS8 struct { F0 uint16; F1 int }
func genFn0(p0 string, p1 uint16, p2 uint64) int16 {
    v4 := GenS1{F0: (p2 & 42), F1: 1}
    fmt.Println(p0)
    fmt.Println(p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 673
}

func main() {
    var v1 uint = 0
    var v2 int8 = -128
    v3 := make(map[int]int)
    if false {
        fmt.Println(v1, v1)
    } else {
        _ = len(v3)
        delete(v3, 1)
    }
    for fi5 := 0; fi5 < 17; fi5++ {
        var v6 uint16 = 2
        var v7 uint32 = 4294967295
        var v8 int8 = -128
        _ = fi5
        _ = v6
        _ = v7
        _ = v8
    }
    switch 1 {
    case 32767:
        for fi9 := 0; fi9 < 14; fi9++ {
            for fi10 := 0; fi10 < 14; fi10++ {
                var v11 int16 = 32767
                _ = fi10
                _ = v11
            }
            if (1000 <= 959) {
                fmt.Println(v1, v1)
            }
            v2 -= v2
            _ = fi9
        }
    default:
        v1 -= 1
    }
    v12 := []int{2147483647, -529, 117, 71}
    v2 = (v2 / (((v2 << ((v1) & 63))) | 1))
    var v13 int = 4294967295
    if true {
        var v14 uint = v1
        v15 := GenS1{F0: 18446744073709551615, F1: 2147483647}
        for fi16 := 0; fi16 < 5; fi16++ {
            if true {
                v17 := GenS1{F0: 824, F1: 42}
                fmt.Println(v13, fi16, v13)
                v18 := GenS1{F0: 15, F1: -736}
                _ = v17
                _ = v18
            }
            switch (-(v13)) {
            case 127:
                var v19 int8 = (v2 + (v2 | (v2 + v2)))
                _ = v19
            case 4294967295:
                v20 := GenS2{F0: int8(v14), F1: (v13 % ((fi16) | 1)), F2: 4294967295}
                _ = v20
            case -1:
                for fi21 := 0; fi21 < 8; fi21++ {
                    v22 := GenS3{F0: "", F1: 2}
                    var v23 uint8 = 0
                    _ = len(v12)
                    _ = fi21
                    _ = v22
                    _ = v23
                }
            }
            _ = fi16
        }
        _ = v14
        _ = v15
    }
    if false {
        for fi24 := 0; fi24 < 19; fi24++ {
            v1 = 65535
            for fi25 := 0; fi25 < 10; fi25++ {
                var v26 int16 = 728
                _ = fi25
                _ = v26
            }
            _ = fi24
        }
        switch "foo bar" {
        case "a":
            var v27 uint16 = 0
            _ = v27
        case "":
            v3[(v13 | v13)] = v13
        default:
            var v28 float64 = (1e10 / 1.5)
            _ = v28
        }
        if (false || (!false)) {
            v29 := GenS2{F0: (v2 - (v2 * (^(v2)))), F1: 32768, F2: 2147483648}
            if (1 >= 2147483647) {
                var v30 int8 = 1
                if ((^(v30)) < v2) {
                    var v31 uint32 = 443
                    v32 := GenS2{F0: v2, F1: (^(v13)), F2: (v31 % (((^(v31))) | 1))}
                    v2 -= 85
                    _ = v31
                    _ = v32
                } else {
                    _ = len(v3)
                    var v33 int8 = v30
                    _ = v33
                }
                if false {
                    for fi34 := 0; fi34 < 8; fi34++ {
                        v1 -= v1
                        var v35 uint = v1
                        for fi36 := 0; fi36 < 9; fi36++ {
                            v37 := GenS4{F0: v1, F1: (v2 >> (((v1 * 1)) & 63)), F2: 921}
                            var v38 int32 = -1
                            v30 -= (v2 - (v30 + 42))
                            _ = fi36
                            _ = v37
                            _ = v38
                        }
                        _ = fi34
                        _ = v35
                    }
                }
                _ = v30
            }
            if (true && true) {
                fmt.Println(v13, v1, v13)
                if false {
                    if (106 == 0) {
                        fmt.Println(v1, v13, v1)
                        if (700 >= 65535) {
                            _ = len(v3)
                        }
                    }
                    for fi39 := 0; fi39 < 14; fi39++ {
                        fmt.Println(v1, v1)
                        _ = fi39
                    }
                } else {
                    var v40 uint16 = 722
                    var v41 int16 = -570
                    if false {
                        for fi42 := 0; fi42 < 8; fi42++ {
                            v3[v13] = (^(fi42))
                            var v43 uint16 = v40
                            _ = len(v3)
                            _ = fi42
                            _ = v43
                        }
                        var v44 uint64 = 943
                        _ = v44
                    }
                    _ = v40
                    _ = v41
                }
            } else {
                var v45 uint = v1
                switch "\n\t" {
                case "foo bar":
                    for fi46 := 0; fi46 < 9; fi46++ {
                        var v47 int32 = 0
                        var v48 uint8 = 255
                        _ = fi46
                        _ = v47
                        _ = v48
                    }
                case "":
                    if false {
                        if false {
                            v1 = v45
                            v2 = 0
                        } else {
                            for fi49 := 0; fi49 < 3; fi49++ {
                                v50 := v3[(fi49 & fi49)]
                                _ = fi49
                                _ = v50
                            }
                        }
                    } else {
                        var v51 uint32 = 42
                        fmt.Println(v2, v1, v13)
                        var v52 uint32 = v51
                        _ = v51
                        _ = v52
                    }
                default:
                    for fi53 := 0; fi53 < 13; fi53++ {
                        switch "foo bar" {
                        case "hello":
                            v54 := GenS2{F0: (v2 | (v2 - (v2 << ((v45) & 63)))), F1: v13, F2: 489}
                            _ = v54
                        case "":
                            switch (fi53 | -2147483648) {
                            case 0:
                                if false {
                                    for fi55 := 0; fi55 < 4; fi55++ {
                                        _ = len(v12)
                                        _ = fi55
                                    }
                                    v3[-2147483648] = (fi53 / ((32768) | 1))
                                } else {
                                    v13 -= fi53
                                }
                            }
                        }
                        v45 = (^(v1))
                        _ = fi53
                    }
                }
                _ = v45
            }
            _ = v29
        } else {
            var v56 int8 = v2
            _ = v56
        }
    } else {
        if len(v12) >= 2 {
            v12 = v12[1:len(v12)]
        }
    }
    v57 := GenS3{F0: "\n\t", F1: 1}
    fmt.Println(v1, v1)
    if (0 == 0) {
        v13 += (v13 + v13)
        fmt.Println(v13)
    } else {
        v12 = append(v12, v13)
        switch ("a" + ("hello" + ("\n\t" + "hello"))) {
        case "":
            v13 += (v13 / (((v13 * v13)) | 1))
        case "foo bar":
            v58 := GenS1{F0: 632, F1: -1}
            _ = v58
        case "\n\t":
            v59 := GenS1{F0: 4294967295, F1: 0}
            _ = v59
        default:
            for fi60 := 0; fi60 < 6; fi60++ {
                var v61 int32 = 0
                v62 := GenS2{F0: v2, F1: 128, F2: 2147483647}
                _ = fi60
                _ = v61
                _ = v62
            }
        }
    }
    v1 += 0
    v12 = append(v12, 32767)
    switch -827 {
    case -1:
        _ = len(v3)
    }
    if (true || (-389 != 9223372036854775807)) {
        for fi63 := 0; fi63 < 6; fi63++ {
            delete(v3, (-(v13)))
            if (1 <= uint64(v1)) {
                switch "foo bar" {
                case "hello":
                    for fi64 := 0; fi64 < 8; fi64++ {
                        v13 = v13
                        _ = fi64
                    }
                case "a":
                    var v65 uint32 = 1
                    _ = v65
                }
            } else {
                v66 := GenS5{F0: 2, F1: 834, F2: 2}
                v67 := v3[(fi63 << ((94) & 63))]
                _ = v66
                _ = v67
            }
            _ = fi63
        }
        var v68 int = 919
        _ = v68
    } else {
        var v69 uint = v1
        fmt.Println(v69, v2, v13)
        _ = v69
    }
    var v70 float64 = 0.0
    v71 := GenS6{F0: 2, F1: (("hello" + "a") + "foo bar"), F2: 65535}
    for fi72 := 0; fi72 < 19; fi72++ {
        for fi73 := 0; fi73 < 7; fi73++ {
            if len(v12) >= 2 {
                v12 = v12[1:len(v12)]
            }
            _ = fi73
        }
        if ((!(42 != 205)) && false) {
            v1 = v1
        } else {
            if (2 <= 0) {
                var v74 uint8 = 255
                v75 := GenS7{F0: 255}
                v3[-831] = v13
                _ = v74
                _ = v75
            }
            for fi76 := 0; fi76 < 20; fi76++ {
                switch ("a" + "") {
                case "foo bar":
                    fmt.Println(fi76, v70)
                }
                _ = fi76
            }
            v77 := GenS8{F0: 2, F1: (v13 - fi72)}
            _ = v77
        }
        v1 -= v1
        _ = fi72
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v70:", v70)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v57
    _ = v70
    _ = v71
}
