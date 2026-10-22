package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint8; F2 uint64 }
type GenS2 struct { F0 string; F1 int64; F2 bool }
type GenS3 struct { F0 uint16; F1 int8 }
type GenS4 struct { F0 int8; F1 string; F2 int64 }
type GenS5 struct { F0 int32; F1 uint8 }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 uint32; F1 int8; F2 bool }
type GenS8 struct { F0 uint32 }
func genFn0(p0 uint16) int8 {
    switch 9223372036854775807 {
    case -2147483648:
        switch -9223372036854775808 {
        case 128:
            var v5 uint = 255
            _ = v5
        case -2147483648:
            for fi6 := 0; fi6 < 19; fi6++ {
                var v7 int32 = 516
                v7 += (v7 * (v7 / ((v7) | 1)))
                _ = fi6
                _ = v7
            }
        case 466:
            switch ("\n\t" + "hello") {
            case "foo bar":
                for fi8 := 0; fi8 < 2; fi8++ {
                    switch "\n\t" {
                    case "":
                        p0 = (p0 * p0)
                    case "foo bar":
                        fmt.Println(p0, fi8)
                    }
                    p0 -= p0
                    p0 += p0
                    _ = fi8
                }
            }
        default:
            switch "" {
            case "\n\t":
                v9 := GenS1{F0: "\n\t", F1: 213, F2: 37}
                _ = v9
            default:
                p0 = (p0 % (((p0 * (p0 >> ((65535) & 63)))) | 1))
            }
        }
    case -605:
        v10 := []int{-756, 32768, 1, 32767}
        _ = v10
    default:
        var v11 int64 = -9223372036854775808
        _ = v11
    }
    v12 := GenS1{F0: (("\n\t" + "\n\t") + "hello"), F1: 0, F2: 1}
    v13 := []int{}
    fmt.Println(p0)
    if false {
        var v14 uint = 65535
        if ((v14 ^ 312) != v14) {
            p0 += (p0 ^ p0)
        }
        fmt.Println(p0, v14)
        _ = v14
    }
    _ = p0
    _ = v12
    _ = v13
    return 42
}

func main() {
    var v1 uint32 = 707
    var v2 uint64 = 925
    var v3 int64 = 9223372036854775807
    var v4 uint16 = 32767
    var v15 int32 = 0
    v16 := []int{}
    v17 := GenS2{F0: "a", F1: 0, F2: ((1 != uint(v4)) || (((376 >= 32767) && false) && true))}
    switch "" {
    case "foo bar":
        switch 4294967295 {
        case 4294967295:
            v18 := GenS2{F0: ("hello" + "a"), F1: v3, F2: ((-2.25 / -2.25) > (1.5 + 0.1))}
            _ = v18
        case 383:
            if len(v16) >= 2 {
                v16 = v16[1:len(v16)]
            }
        }
    case "\n\t":
        if (v1 < (v1 | v1)) {
            if false {
                if len(v16) > 0 {
                    v16[0] = 32767
                }
            } else {
                if ((v1 - (v1 >> ((722) & 63))) == 821) {
                    v1 += 529
                }
                v15 = v15
                if false {
                    if len(v16) >= 2 {
                        v16 = v16[1:len(v16)]
                    }
                    v16 = append(v16, -2147483648)
                } else {
                    v15 -= -30
                }
            }
        }
    case "":
        v2 = (v2 * (v2 % ((v2) | 1)))
    }
    switch -9223372036854775808 {
    case -70:
        for fi19 := 0; fi19 < 8; fi19++ {
            v20 := GenS3{F0: v4, F1: 127}
            var v21 uint8 = 31
            for fi22 := 0; fi22 < 6; fi22++ {
                v4 = 817
                switch fi19 {
                case 0:
                    var v23 uint = 42
                    _ = v23
                case 883:
                    if (!(-1 > -32768)) {
                        if (true && (((^(fi19)) >= int(v21)) && true)) {
                            fmt.Println(v21, v1, v4)
                            v3 -= int64(v2)
                            switch (-(fi22)) {
                            case -9223372036854775808:
                                v24 := GenS2{F0: "a", F1: (v3 - 468), F2: (true && true)}
                                _ = v24
                            case -17:
                                v1 += (v1 ^ (v1 * (v1 & (v1 ^ 2147483648))))
                            }
                        } else {
                            v3 = v3
                            v25 := GenS3{F0: (v4 / (((v4 << ((210) & 63))) | 1)), F1: 1}
                            if (int32(fi22) > (-(v15))) {
                                var v26 int8 = -1
                                v26 = int8(v2)
                                var v27 string = "foo bar"
                                _ = v26
                                _ = v27
                            }
                            v21 = v21
                            if false {
                                var v28 int16 = int16(v15)
                                _ = v28
                            }
                            v21 += v21
                            if true {
                                var v29 int32 = 42
                                v30 := GenS4{F0: 0, F1: "foo bar", F2: v3}
                                if ((genFn0(v4) == genFn0(v4)) || false) {
                                    v3 = 42
                                }
                                _ = v29
                                _ = v30
                            }
                            _ = v25
                        }
                        fmt.Println(v1)
                        var v31 uint32 = 4294967295
                        _ = v31
                    }
                }
                v32 := GenS5{F0: (v15 % ((42) | 1)), F1: (v21 + (^(v21)))}
                _ = fi22
                _ = v32
            }
            _ = fi19
            _ = v20
            _ = v21
        }
    default:
        switch "foo bar" {
        case "":
            v15 = 42
        case "\n\t":
            if true {
                _ = len(v16)
                for fi33 := 0; fi33 < 19; fi33++ {
                    var v34 bool = (84 != 1)
                    _ = fi33
                    _ = v34
                }
                switch "hello" {
                case "\n\t":
                    v4 += (v4 + uint16(v2))
                case "foo bar":
                    switch ("a" + "") {
                    case "foo bar":
                        fmt.Println(v3, v4, v15)
                    }
                case "hello":
                    var v35 int16 = 32767
                    _ = v35
                default:
                    for fi36 := 0; fi36 < 16; fi36++ {
                        v3 += v3
                        v3 = -1
                        if ((v1 / (((v1 << ((248) & 63))) | 1)) <= (v1 | v1)) {
                            switch fi36 {
                            case 42:
                                for fi37 := 0; fi37 < 3; fi37++ {
                                    var v38 uint64 = (v2 / ((42) | 1))
                                    _ = fi37
                                    _ = v38
                                }
                            case 32767:
                                v1 -= 2147483648
                            case -452:
                                v15 = v15
                            default:
                                for fi39 := 0; fi39 < 8; fi39++ {
                                    var v40 int16 = 42
                                    if false {
                                        var v41 float64 = 1e10
                                        v3 = (v3 | -237)
                                        switch (fi39 - (^(fi39))) {
                                        case -515:
                                            var v42 uint8 = 170
                                            _ = v42
                                        default:
                                            if true {
                                                v1 = v1
                                                var v43 int32 = v15
                                                fmt.Println(fi36, v15, fi36)
                                                _ = v43
                                            }
                                        }
                                        _ = v41
                                    }
                                    v40 -= v40
                                    _ = fi39
                                    _ = v40
                                }
                            }
                            if true {
                                for fi44 := 0; fi44 < 18; fi44++ {
                                    var v45 uint16 = v4
                                    if ("\n\t" == "a") {
                                        if true {
                                            v46 := GenS6{F0: v2}
                                            _ = v46
                                        }
                                        v47 := GenS6{F0: v2}
                                        v48 := GenS7{F0: v1, F1: 42, F2: false}
                                        _ = v47
                                        _ = v48
                                    }
                                    _ = fi44
                                    _ = v45
                                }
                                if (true && false) {
                                    fmt.Println(v2)
                                    v15 -= (v15 % ((v15) | 1))
                                }
                            }
                        }
                        _ = fi36
                    }
                }
            } else {
                fmt.Println(v15)
                v4 = 21
            }
        case "foo bar":
            v49 := GenS8{F0: v1}
            _ = v49
        default:
            var v50 int64 = (^(v3))
            _ = v50
        }
    }
    for fi51 := 0; fi51 < 13; fi51++ {
        fmt.Println(v1)
        v1 = v1
        _ = fi51
    }
    var v52 int64 = v3
    if len(v16) > 0 {
        v16[0] = 1
    }
    if len(v16) >= 2 {
        v16 = v16[1:len(v16)]
    }
    for fi53 := 0; fi53 < 17; fi53++ {
        v15 += -2147483648
        _ = fi53
    }
    var v54 uint = 19
    v55 := GenS3{F0: v4, F1: 100}
    for fi56 := 0; fi56 < 3; fi56++ {
        if (!false) {
            var v57 int64 = v3
            var v58 int = fi56
            _ = v57
            _ = v58
        } else {
            v54 += uint(v1)
            var v59 uint8 = 2
            _ = v59
        }
        switch (fi56 - 2147483647) {
        case 32767:
            var v60 float32 = -2.25
            _ = v60
        default:
            var v61 float64 = (float64(1e10) - 3.14159)
            _ = v61
        }
        _ = fi56
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v52:", v52)
    fmt.Println("v54:", v54)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v52
    _ = v54
    _ = v55
}
