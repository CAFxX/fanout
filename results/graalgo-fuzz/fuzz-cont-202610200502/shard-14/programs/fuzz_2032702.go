package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int16; F2 bool }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int32; F1 uint16; F2 uint16 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 int8 }
func genFn0(p0 int16, p1 int16) int {
    var v7 uint8 = 42
    for fi8 := 0; fi8 < 10; fi8++ {
        if (false || ("a" >= ("a" + "a"))) {
            fmt.Println(p0, v7)
        }
        var v9 uint = 42
        _ = fi8
        _ = v9
    }
    v10 := []int{-2147483648, 32767}
    switch "a" {
    case "\n\t":
        p1 = 0
    case "foo bar":
        var v11 uint = 309
        _ = v11
    default:
        if len(v10) > 0 {
            v10[0] = -747
        }
    }
    _ = p0
    _ = p1
    _ = v10
    _ = v7
    return 0
}
func genFn1() uint32 {
    v12 := GenS1{F0: -9223372036854775808, F1: -65, F2: (((0.1 / 1.5) * float32((0.1 + 1e10))) != float32((float64(3.14159) + 1e10)))}
    var v13 int64 = -705
    v14 := []int{}
    if true {
        var v15 float32 = 3.14159
        var v16 string = ("a" + "foo bar")
        _ = v15
        _ = v16
    }
    if true {
        if false {
            v13 -= 314
            if true {
                v13 -= v13
                v13 -= (^(v13))
                fmt.Println(v13)
            } else {
                fmt.Println(v13)
                fmt.Println(v13)
            }
        }
        var v17 float32 = 0.1
        _ = v17
    } else {
        switch -920 {
        case 0:
            if len(v14) > 0 {
                v14[0] = 32768
            }
        default:
            if false {
                var v18 int64 = (v13 >> ((42) & 63))
                if len(v14) >= 2 {
                    v14 = v14[1:len(v14)]
                }
                fmt.Println(v18, v13)
                _ = v18
            } else {
                for fi19 := 0; fi19 < 9; fi19++ {
                    _ = len(v14)
                    v13 += 9223372036854775807
                    _ = fi19
                }
                v13 -= v13
            }
        }
        v13 += (^(v13))
        var v20 uint64 = 4294967295
        _ = v20
    }
    _ = v12
    _ = v13
    _ = v14
    return 2
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := []string{"\n\t", "hello", "foo bar"}
    var v3 int = 32767
    v4 := make(map[string]string)
    var v5 uint16 = 1
    v6 := make(map[uint32]int)
    switch "a" {
    case "hello":
        fmt.Println(v3, v3)
    default:
        v5 += (^(v5))
    }
    if false {
        if true {
            for fi21 := 0; fi21 < 10; fi21++ {
                for fi22 := 0; fi22 < 8; fi22++ {
                    v23 := GenS2{F0: 42}
                    _ = fi22
                    _ = v23
                }
                _ = fi21
            }
            switch "foo bar" {
            case "a":
                v3 += (v3 ^ genFn0(1, 1))
            default:
                if (0 <= 686) {
                    switch ("" + "") {
                    case "hello":
                        v24 := GenS3{F0: -750, F1: v5, F2: 255}
                        _ = v24
                    case "\n\t":
                        v5 += v5
                    case "a":
                        fmt.Println(v5, v3)
                    }
                    fmt.Println(v3, v3)
                    var v25 int16 = -945
                    _ = v25
                } else {
                    _ = len(v1)
                    switch "hello" {
                    case "foo bar":
                        _ = len(v2)
                    case "hello":
                        switch "\n\t" {
                        case "":
                            v3 -= (v3 - v3)
                        case "foo bar":
                            v26 := GenS1{F0: 9223372036854775807, F1: 449, F2: false}
                            _ = v26
                        default:
                            for fi27 := 0; fi27 < 1; fi27++ {
                                v3 = (fi27 * (v3 & (fi27 << ((42) & 63))))
                                var v28 uint16 = (v5 % ((v5) | 1))
                                _ = fi27
                                _ = v28
                            }
                        }
                    default:
                        v29 := GenS2{F0: v5}
                        _ = v29
                    }
                    switch v3 {
                    case 32767:
                        var v30 int32 = 841
                        _ = v30
                    }
                }
            }
            switch v3 {
            case 0:
                switch "\n\t" {
                case "":
                    v31 := GenS4{F0: 255}
                    _ = v31
                default:
                    v2 = append(v2, "foo bar")
                }
            case 128:
                if ("a" != ("" + "foo bar")) {
                    var v32 int8 = 127
                    if true {
                        for fi33 := 0; fi33 < 2; fi33++ {
                            v3 = 4294967295
                            switch v3 {
                            case 127:
                                fmt.Println(v3, v5, v5)
                            }
                            _ = fi33
                        }
                    } else {
                        var v34 uint32 = 2147483648
                        if false {
                            v35 := GenS2{F0: (v5 / ((v5) | 1))}
                            if ((v34 | uint32(v5)) != v34) {
                                var v36 int = 2147483647
                                _ = v36
                            }
                            _ = v35
                        }
                        _ = v34
                    }
                    _ = v32
                }
            case 9223372036854775807:
                for fi37 := 0; fi37 < 5; fi37++ {
                    v6[160] = (fi37 - (v3 << ((42) & 63)))
                    _ = fi37
                }
            default:
                var v38 uint16 = (^(v5))
                _ = v38
            }
        }
    }
    v5 = v5
    var v39 uint64 = 658
    for fi40 := 0; fi40 < 14; fi40++ {
        switch "foo bar" {
        case "hello":
            var v41 uint32 = 868
            _ = v41
        case "foo bar":
            if false {
                v42 := GenS3{F0: 0, F1: (v5 * (v5 ^ (v5 * 692))), F2: 152}
                _ = v42
            }
        default:
            v43 := GenS5{F0: 32}
            _ = v43
        }
        v39 -= 209
        _ = fi40
    }
    v3 = v3
    if len(v2) > 0 {
        v2[0] = "\n\t"
    }
    for fi44 := 0; fi44 < 9; fi44++ {
        v45 := GenS3{F0: -1, F1: v5, F2: v5}
        var v46 int8 = -128
        v3 = (v3 >> ((255) & 63))
        _ = fi44
        _ = v45
        _ = v46
    }
    if false {
        v5 = (^(v5))
        for fi47 := 0; fi47 < 19; fi47++ {
            v3 = (fi47 & (v3 % (((^(fi47))) | 1)))
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            _ = fi47
        }
        for fi48 := 0; fi48 < 19; fi48++ {
            for fi49 := 0; fi49 < 3; fi49++ {
                var v50 float32 = 0.1
                var v51 bool = true
                _ = fi49
                _ = v50
                _ = v51
            }
            for fi52 := 0; fi52 < 2; fi52++ {
                var v53 uint = uint(v5)
                var v54 uint16 = v5
                _ = fi52
                _ = v53
                _ = v54
            }
            var v55 int16 = -97
            _ = fi48
            _ = v55
        }
    }
    v3 += (v3 >> ((786) & 63))
    var v56 int32 = -1
    var v57 int32 = (v56 + (v56 / ((-2147483648) | 1)))
    if true {
        var v58 uint = 236
        v59 := GenS4{F0: 213}
        _ = v58
        _ = v59
    } else {
        v60 := GenS4{F0: 42}
        _ = v60
    }
    for fi61 := 0; fi61 < 3; fi61++ {
        v3 = fi61
        v3 += (fi61 / (((v3 % ((v3) | 1))) | 1))
        var v62 int64 = -9223372036854775808
        _ = fi61
        _ = v62
    }
    for fi63 := 0; fi63 < 11; fi63++ {
        for fi64 := 0; fi64 < 16; fi64++ {
            v39 = v39
            switch "\n\t" {
            case "hello":
                if true {
                    var v65 uint16 = (v5 * (v5 << ((18446744073709551615) & 63)))
                    v66 := GenS6{F0: 1}
                    _ = v65
                    _ = v66
                }
            case "":
                v67 := GenS6{F0: 1}
                _ = v67
            }
            fmt.Println(v39, v39)
            _ = fi64
        }
        _ = fi63
    }
    for fi68 := 0; fi68 < 15; fi68++ {
        switch int(v57) {
        case 9223372036854775807:
            var v69 uint = 65535
            _ = v69
        default:
            for fi70 := 0; fi70 < 16; fi70++ {
                v56 += (v57 | v56)
                var v71 uint32 = 2
                _ = fi70
                _ = v71
            }
        }
        v5 -= v5
        _ = fi68
    }
    fmt.Println(v5, v56, v5)
    switch (v3 ^ (v3 * v3)) {
    case 42:
        fmt.Println(v56, v5, v39)
    case -2147483648:
        v3 = v3
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v56:", v56)
    fmt.Println("v57:", v57)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v5
    _ = v56
    _ = v57
    _ = v6
}
