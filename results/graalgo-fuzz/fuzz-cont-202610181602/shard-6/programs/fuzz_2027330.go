package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool; F2 uint8 }
type GenS2 struct { F0 int16; F1 int16; F2 uint16 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 string; F1 uint8; F2 uint }
type GenS5 struct { F0 uint; F1 int16; F2 int }
type GenS6 struct { F0 int16; F1 uint8; F2 uint16 }
type GenS7 struct { F0 int16 }
type GenS8 struct { F0 int32 }
func genFn0(p0 bool, p1 int, p2 int8) int8 {
    v6 := GenS1{F0: 51, F1: (495 > -495), F2: 156}
    v7 := []int{p1, (p1 << ((0) & 63)), -904, p1}
    v8 := GenS2{F0: 1, F1: int16(p1), F2: 438}
    fmt.Println(p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    _ = v8
    return p2
}
func genFn1(p0 int16) int8 {
    if true {
        for fi9 := 0; fi9 < 11; fi9++ {
            var v10 int = fi9
            _ = fi9
            _ = v10
        }
        v11 := []int{32767, -2147483648}
        for fi12 := 0; fi12 < 18; fi12++ {
            var v13 bool = (true && (!(255 <= 42)))
            p0 = (p0 & 470)
            _ = fi12
            _ = v13
        }
        _ = v11
    } else {
        if false {
            p0 = (p0 - p0)
            if false {
                var v14 uint = 65535
                switch "" {
                case "":
                    if false {
                        p0 = p0
                    }
                default:
                    switch 4294967295 {
                    case 1:
                        p0 += p0
                    }
                }
                v15 := GenS3{F0: (255 != 255)}
                _ = v14
                _ = v15
            } else {
                for fi16 := 0; fi16 < 5; fi16++ {
                    var v17 int16 = p0
                    _ = fi16
                    _ = v17
                }
                fmt.Println(p0)
            }
            var v18 uint32 = 1
            _ = v18
        } else {
            v19 := []int{93, -1}
            for fi20 := 0; fi20 < 15; fi20++ {
                v19 = append(v19, fi20)
                var v21 int32 = -19
                switch (^(fi20)) {
                case -331:
                    var v22 bool = false
                    _ = v22
                case 528:
                    _ = len(v19)
                case -969:
                    p0 += p0
                }
                _ = fi20
                _ = v21
            }
            _ = v19
        }
    }
    for fi23 := 0; fi23 < 13; fi23++ {
        v24 := []int{}
        p0 = -455
        _ = fi23
        _ = v24
    }
    _ = p0
    return -128
}

func main() {
    v1 := []string{"foo bar", ("" + "a")}
    var v2 uint8 = 2
    var v3 int16 = -32768
    var v4 uint8 = v2
    v5 := []string{}
    if (!true) {
        fmt.Println(v2, v2, v2)
        var v25 bool = (100 <= genFn0(true, 1, 0))
        v4 += (v4 / ((v2) | 1))
        _ = v25
    }
    v3 = (v3 >> ((0) & 63))
    for fi26 := 0; fi26 < 16; fi26++ {
        switch ((("" + "") + "") + ("\n\t" + ("\n\t" + "foo bar"))) {
        case "\n\t":
            switch "" {
            case "a":
                _ = len(v5)
            case "foo bar":
                var v27 uint8 = v2
                _ = v27
            case "hello":
                v28 := GenS3{F0: false}
                _ = v28
            }
        case "foo bar":
            var v29 uint16 = 909
            _ = v29
        case "a":
            v30 := GenS4{F0: "foo bar", F1: (v2 * v2), F2: 580}
            _ = v30
        }
        v4 += v2
        _ = fi26
    }
    v3 += 1000
    if true {
        switch ("\n\t" + "") {
        case "hello":
            switch -9223372036854775808 {
            case 295:
                if (-9223372036854775808 < 9223372036854775807) {
                    v4 = 127
                } else {
                    if len(v5) > 0 {
                        v5[0] = ""
                    }
                }
            case 9223372036854775807:
                var v31 uint64 = 321
                _ = v31
            case -597:
                v3 = (^(v3))
            }
        }
        for fi32 := 0; fi32 < 14; fi32++ {
            v2 += (v2 ^ 0)
            v33 := GenS2{F0: -1, F1: v3, F2: 42}
            switch ("foo bar" + "a") {
            case "hello":
                v3 += (v3 >> ((4294967295) & 63))
            case "":
                var v34 int64 = -1
                _ = v34
            default:
                v35 := GenS1{F0: 2147483647, F1: (!false), F2: 1}
                _ = v35
            }
            _ = fi32
            _ = v33
        }
        var v36 int = 127
        _ = v36
    }
    switch ("\n\t" + "a") {
    case "":
        var v37 int = 4294967295
        _ = v37
    case "\n\t":
        var v38 uint32 = 2147483647
        _ = v38
    case "hello":
        var v39 uint = 0
        _ = v39
    }
    var v40 int16 = v3
    v3 = v40
    fmt.Println(v3, v3)
    for fi41 := 0; fi41 < 15; fi41++ {
        if ((^(v4)) >= 54) {
            if len(v1) > 0 {
                v1[0] = ("a" + ("a" + ("hello" + ("" + "\n\t"))))
            }
            v4 -= (v4 % (((v2 ^ (v2 | (v2 - (v2 ^ 255))))) | 1))
            _ = len(v5)
        }
        var v42 uint8 = (v2 & (v2 % (((v4 & v2)) | 1)))
        if (!true) {
            _ = len(v1)
        }
        _ = fi41
        _ = v42
    }
    for fi43 := 0; fi43 < 15; fi43++ {
        v3 -= v40
        if (false || true) {
            switch "foo bar" {
            case "hello":
                var v44 int64 = -9223372036854775808
                _ = v44
            case "\n\t":
                var v45 uint64 = 878
                _ = v45
            case "a":
                if (-128 < -120) {
                    var v46 int = (-(fi43))
                    if (v4 == v2) {
                        v47 := GenS2{F0: (v3 * (v3 >> ((42) & 63))), F1: v40, F2: 0}
                        _ = v47
                    } else {
                        fmt.Println(fi43, v40, fi43)
                        for fi48 := 0; fi48 < 2; fi48++ {
                            v40 += (v40 ^ v40)
                            if (-1 < 82) {
                                switch fi43 {
                                case 128:
                                    v46 -= fi43
                                case 127:
                                    v49 := GenS1{F0: 577, F1: false, F2: 129}
                                    _ = v49
                                case -489:
                                    v50 := GenS1{F0: 26, F1: (479 == 255), F2: (v2 / ((v2) | 1))}
                                    _ = v50
                                default:
                                    switch v46 {
                                    case 947:
                                        v51 := GenS5{F0: 1, F1: (v3 | v3), F2: -144}
                                        _ = v51
                                    case 32767:
                                        if true {
                                            v52 := GenS6{F0: 0, F1: (^(v2)), F2: 407}
                                            fmt.Println(v46, fi48, v2)
                                            _ = v52
                                        } else {
                                            switch v46 {
                                            case -369:
                                                v40 = 0
                                            case 729:
                                                var v53 uint = 1
                                                _ = v53
                                            case 600:
                                                v2 += v4
                                            }
                                        }
                                    case -160:
                                        v54 := GenS2{F0: 1000, F1: -32768, F2: uint16(v2)}
                                        _ = v54
                                    }
                                }
                                v2 -= v4
                                v4 = (v2 | 44)
                                v55 := GenS7{F0: -32768}
                                var v56 uint64 = 23
                                v56 = v56
                                _ = v55
                                _ = v56
                            } else {
                                v57 := GenS8{F0: -2147483648}
                                fmt.Println(v46)
                                _ = v57
                            }
                            _ = fi48
                        }
                    }
                    var v58 uint64 = 830
                    _ = v46
                    _ = v58
                }
            default:
                v3 = v3
            }
            fmt.Println(v3, v3)
        }
        _ = fi43
    }
    for fi59 := 0; fi59 < 6; fi59++ {
        var v60 int16 = (^(v40))
        _ = fi59
        _ = v60
    }
    for fi61 := 0; fi61 < 6; fi61++ {
        for fi62 := 0; fi62 < 13; fi62++ {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            _ = fi62
        }
        fmt.Println(v3)
        var v63 int32 = 2147483647
        _ = fi61
        _ = v63
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v5
}
