package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int32; F1 int32; F2 uint16 }
type GenS3 struct { F0 uint16; F1 uint; F2 uint32 }
type GenS4 struct { F0 int; F1 int8 }
type GenS5 struct { F0 string; F1 uint }
type GenS6 struct { F0 int32; F1 uint64 }
func genFn0(p0 uint64) int16 {
    v6 := []int{127, 176, 128, 4294967295}
    v7 := GenS1{F0: 640}
    if true {
        var v8 float64 = (-2.25 / 1.0)
        _ = v8
    } else {
        for fi9 := 0; fi9 < 6; fi9++ {
            var v10 bool = false
            if (38 < -32768) {
                for fi11 := 0; fi11 < 11; fi11++ {
                    var v12 int8 = -61
                    if (9223372036854775807 <= 1) {
                        if len(v6) >= 2 {
                            v6 = v6[1:len(v6)]
                        }
                        var v13 string = "\n\t"
                        v14 := GenS1{F0: fi11}
                        _ = v13
                        _ = v14
                    } else {
                        v12 += 45
                    }
                    switch "\n\t" {
                    case "a":
                        v15 := GenS1{F0: fi11}
                        _ = v15
                    }
                    _ = fi11
                    _ = v12
                }
                fmt.Println(fi9, p0)
            } else {
                var v16 string = "foo bar"
                _ = v16
            }
            _ = fi9
            _ = v10
        }
        var v17 uint8 = 170
        for fi18 := 0; fi18 < 13; fi18++ {
            fmt.Println(p0, v17)
            _ = fi18
        }
        _ = v17
    }
    var v19 uint8 = 127
    _ = p0
    _ = v19
    _ = v6
    _ = v7
    return -1
}
func genFn1(p0 uint64, p1 int8) int32 {
    for fi20 := 0; fi20 < 20; fi20++ {
        v21 := []int{fi20, fi20, fi20, fi20}
        fmt.Println(p0, p0, p1)
        _ = fi20
        _ = v21
    }
    v22 := []int{}
    switch 128 {
    case 32767:
        if true {
            if (false || false) {
                if (65535 == 2147483648) {
                    if false {
                        switch 42 {
                        case -2147483648:
                            if false {
                                p0 = p0
                                var v23 uint64 = p0
                                _ = v23
                            }
                        }
                        v24 := GenS1{F0: -2147483648}
                        for fi25 := 0; fi25 < 12; fi25++ {
                            switch ("hello" + ("\n\t" + "hello")) {
                            case "hello":
                                var v26 uint = 569
                                _ = v26
                            case "foo bar":
                                switch "a" {
                                case "\n\t":
                                    v27 := GenS1{F0: fi25}
                                    _ = v27
                                case "foo bar":
                                    var v28 bool = false
                                    _ = v28
                                case "":
                                    for fi29 := 0; fi29 < 9; fi29++ {
                                        p1 = (p1 * 127)
                                        var v30 uint = 65535
                                        _ = fi29
                                        _ = v30
                                    }
                                default:
                                    var v31 uint32 = 327
                                    _ = v31
                                }
                            case "":
                                var v32 string = ("a" + (((("a" + "\n\t") + "\n\t") + "") + ("\n\t" + "")))
                                _ = v32
                            }
                            _ = fi25
                        }
                        _ = v24
                    } else {
                        v33 := GenS2{F0: 2147483647, F1: -2147483648, F2: 2}
                        for fi34 := 0; fi34 < 19; fi34++ {
                            v35 := GenS2{F0: -2147483648, F1: 268, F2: 2}
                            var v36 int8 = p1
                            _ = fi34
                            _ = v35
                            _ = v36
                        }
                        fmt.Println(p1)
                        _ = v33
                    }
                } else {
                    v37 := GenS3{F0: 65535, F1: 661, F2: 42}
                    _ = v37
                }
                v38 := GenS4{F0: -148, F1: p1}
                if len(v22) > 0 {
                    v22[0] = 2147483647
                }
                _ = v38
            }
            var v39 int16 = 913
            _ = v39
        }
    case 9223372036854775807:
        for fi40 := 0; fi40 < 19; fi40++ {
            p1 = 100
            _ = fi40
        }
    default:
        v22 = append(v22, 172)
    }
    p1 = (p1 - (p1 + 127))
    p0 += p0
    _ = p0
    _ = p1
    _ = v22
    return -91
}

func main() {
    v1 := []uint32{42, 2147483648}
    var v2 uint8 = 2
    var v3 int = 2147483647
    var v4 string = "a"
    var v5 int = v3
    v41 := GenS4{F0: 128, F1: -43}
    switch ("hello" + v4) {
    case "":
        v42 := GenS4{F0: v3, F1: 1}
        _ = v42
    }
    for fi43 := 0; fi43 < 3; fi43++ {
        var v44 string = v4
        for fi45 := 0; fi45 < 14; fi45++ {
            switch (v3 << ((uint(v2)) & 63)) {
            case 0:
                v5 += (^(v5))
            case -2147483648:
                v1 = append(v1, 4294967295)
            }
            _ = fi45
        }
        _ = fi43
        _ = v44
    }
    var v46 int32 = 42
    if false {
        fmt.Println(v46, v3)
        v4 = v4
    } else {
        fmt.Println(v46, v5)
        for fi47 := 0; fi47 < 2; fi47++ {
            if len(v1) > 0 {
                v1[0] = 2147483647
            }
            switch v4 {
            case "foo bar":
                v48 := GenS5{F0: "", F1: 2}
                _ = v48
            case "hello":
                var v49 string = v4
                _ = v49
            default:
                v3 = 1
            }
            var v50 string = "hello"
            _ = fi47
            _ = v50
        }
    }
    var v51 int16 = -1
    if true {
        if (false || true) {
            v52 := GenS6{F0: genFn1(693, 70), F1: 0}
            if (!true) {
                fmt.Println(v51)
            } else {
                v46 = (v46 + 410)
                v1 = append(v1, 2147483647)
                for fi53 := 0; fi53 < 14; fi53++ {
                    var v54 bool = true
                    _ = fi53
                    _ = v54
                }
            }
            v55 := GenS5{F0: "a", F1: 255}
            _ = v52
            _ = v55
        } else {
            v2 += (v2 % (((^(v2))) | 1))
            for fi56 := 0; fi56 < 20; fi56++ {
                if true {
                    var v57 float32 = float32(1e10)
                    _ = v57
                }
                for fi58 := 0; fi58 < 20; fi58++ {
                    if len(v1) > 0 {
                        v1[0] = 2
                    }
                    var v59 uint16 = 283
                    var v60 string = ((((("a" + (v4 + (v4 + v4))) + ("\n\t" + v4)) + (v4 + v4)) + v4) + v4)
                    v61 := GenS2{F0: v46, F1: v46, F2: v59}
                    _ = fi58
                    _ = v59
                    _ = v60
                    _ = v61
                }
                _ = fi56
            }
            var v62 bool = (!((v51 << ((uint(v2)) & 63)) >= 858))
            _ = v62
        }
        switch v5 {
        case -951:
            var v63 int64 = int64(v46)
            _ = v63
        default:
            for fi64 := 0; fi64 < 16; fi64++ {
                switch v4 {
                case "foo bar":
                    v65 := GenS6{F0: 1, F1: 87}
                    _ = v65
                case "a":
                    var v66 float64 = ((float64((1.5 - -2.25)) - ((float64(1e10) + ((-2.25 + 3.14159) * float64(-2.25))) - 0.1)) - (1.5 - ((0.1 / 3.14159) + 0.1)))
                    _ = v66
                default:
                    var v67 int16 = -551
                    _ = v67
                }
                fmt.Println(v51)
                _ = fi64
            }
        }
    }
    for fi68 := 0; fi68 < 12; fi68++ {
        v1 = append(v1, 1)
        switch (v5 & fi68) {
        case 188:
            for fi69 := 0; fi69 < 13; fi69++ {
                v46 = v46
                var v70 int32 = int32(v51)
                v71 := GenS1{F0: 2147483647}
                _ = fi69
                _ = v70
                _ = v71
            }
        case 32768:
            var v72 int64 = 0
            _ = v72
        }
        for fi73 := 0; fi73 < 11; fi73++ {
            var v74 float32 = -2.25
            _ = fi73
            _ = v74
        }
        _ = fi68
    }
    fmt.Println(v3, v46)
    if len(v1) > 0 {
        v1[0] = 42
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v46:", v46)
    fmt.Println("v5:", v5)
    fmt.Println("v51:", v51)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v46
    _ = v5
    _ = v51
}
