package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int32; F2 int32 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int32; F1 int; F2 bool }
type GenS4 struct { F0 uint8; F1 uint32 }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 string; F1 int16; F2 bool }
type GenS7 struct { F0 uint16; F1 string; F2 int }
type GenS8 struct { F0 uint32; F1 int16; F2 bool }
type GenS9 struct { F0 int; F1 uint32 }
type GenS10 struct { F0 int64 }
func genFn0(p0 bool, p1 uint16) int8 {
    v4 := GenS1{F0: -44, F1: 42, F2: 253}
    var v5 uint64 = 4294967295
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return -1
}

func main() {
    var v1 uint = 255
    var v2 int8 = 100
    var v3 float64 = 0.1
    for fi6 := 0; fi6 < 20; fi6++ {
        v7 := GenS2{F0: 2}
        fmt.Println(v3)
        v8 := GenS3{F0: -299, F1: (fi6 & (fi6 ^ (fi6 | 127))), F2: true}
        _ = fi6
        _ = v7
        _ = v8
    }
    v2 = (-(v2))
    fmt.Println(v2, v2)
    var v9 int8 = genFn0((0.0 > ((((3.14159 * 1e10) / float32(1.5)) / float32(v3)) + 1e10)), 65535)
    var v10 uint = 1
    v11 := GenS4{F0: 2, F1: 600}
    v12 := []int{42, 707}
    switch 42 {
    case 1:
        if len(v12) >= 2 {
            v12 = v12[1:len(v12)]
        }
    case 0:
        v13 := GenS1{F0: (v2 / (((v9 - (^(v2)))) | 1)), F1: -2147483648, F2: 2147483647}
        _ = v13
    case -1:
        if true {
            v2 -= (^(v2))
            if true {
                fmt.Println(v3, v1)
            } else {
                fmt.Println(v1, v1, v9)
                if false {
                    switch 2147483647 {
                    case 9223372036854775807:
                        var v14 float32 = 0.0
                        _ = v14
                    case 127:
                        for fi15 := 0; fi15 < 11; fi15++ {
                            if (!(!(true || false))) {
                                v1 = v10
                                for fi16 := 0; fi16 < 13; fi16++ {
                                    v17 := GenS5{F0: 42}
                                    v12 = append(v12, (fi15 << (((v1 << ((65535) & 63))) & 63)))
                                    _ = fi16
                                    _ = v17
                                }
                            }
                            v2 += v9
                            v18 := GenS3{F0: 686, F1: (fi15 / (((fi15 | (^(fi15)))) | 1)), F2: true}
                            _ = fi15
                            _ = v18
                        }
                    case 32767:
                        if false {
                            for fi19 := 0; fi19 < 14; fi19++ {
                                if ((819 != 2147483647) && (608 <= 2)) {
                                    if len(v12) >= 2 {
                                        v12 = v12[1:len(v12)]
                                    }
                                    switch "hello" {
                                    case "hello":
                                        if (1000 <= 1) {
                                            fmt.Println(v2, v3, v10)
                                            v9 += (v9 | (-(v9)))
                                            v3 -= v3
                                        }
                                    case "a":
                                        _ = len(v12)
                                    case "":
                                        switch (fi19 << (((^(v1))) & 63)) {
                                        case -9223372036854775808:
                                            v10 += v1
                                        case 1:
                                            v20 := GenS5{F0: 499}
                                            _ = v20
                                        }
                                    default:
                                        if false {
                                            if true {
                                                var v21 int = -526
                                                for fi22 := 0; fi22 < 13; fi22++ {
                                                    if true {
                                                        if false {
                                                            fmt.Println(v10, v10, v3)
                                                            if false {
                                                                if (!(false || ((650 == 394) || (369 <= 219)))) {
                                                                    switch "" {
                                                                    case "a":
                                                                        v12 = append(v12, fi22)
                                                                    default:
                                                                        var v23 uint8 = 2
                                                                        _ = v23
                                                                    }
                                                                } else {
                                                                    v24 := GenS6{F0: "", F1: 1000, F2: (-2.25 < (1.5 - v3))}
                                                                    _ = v24
                                                                }
                                                            } else {
                                                                var v25 float64 = v3
                                                                switch (("\n\t" + (("" + (("hello" + "foo bar") + "\n\t")) + "foo bar")) + "hello") {
                                                                case "\n\t":
                                                                    _ = len(v12)
                                                                case "hello":
                                                                    switch (fi19 >> (((v10 % (((v10 ^ v1)) | 1))) & 63)) {
                                                                    case 128:
                                                                        if (1 != 0) {
                                                                            switch "hello" {
                                                                            case "":
                                                                                var v26 uint8 = 146
                                                                                _ = v26
                                                                            case "a":
                                                                                if len(v12) >= 2 {
                                                                                    v12 = v12[1:len(v12)]
                                                                                }
                                                                            case "\n\t":
                                                                                if (4294967295 != 1) {
                                                                                    if len(v12) >= 2 {
                                                                                        v12 = v12[1:len(v12)]
                                                                                    }
                                                                                    var v27 uint32 = 2147483647
                                                                                    _ = v27
                                                                                } else {
                                                                                    v10 = (v1 + v10)
                                                                                }
                                                                            }
                                                                            if ("foo bar" <= ("foo bar" + "a")) {
                                                                                v9 = v2
                                                                                v21 -= v21
                                                                            } else {
                                                                                fmt.Println(fi22)
                                                                            }
                                                                            var v28 int16 = -32768
                                                                            _ = v28
                                                                        }
                                                                    case 127:
                                                                        var v29 float64 = ((-2.25 - v25) + float64(0.1))
                                                                        _ = v29
                                                                    case 32767:
                                                                        v1 -= (v1 ^ (v1 >> (((^(v10))) & 63)))
                                                                    default:
                                                                        v30 := GenS4{F0: 128, F1: 4294967295}
                                                                        _ = v30
                                                                    }
                                                                default:
                                                                    v1 = v1
                                                                }
                                                                switch ("" + "\n\t") {
                                                                case "hello":
                                                                    switch "\n\t" {
                                                                    case "foo bar":
                                                                        v31 := GenS4{F0: 0, F1: 42}
                                                                        _ = v31
                                                                    case "a":
                                                                        v32 := GenS5{F0: 0}
                                                                        var v33 string = "hello"
                                                                        _ = v32
                                                                        _ = v33
                                                                    default:
                                                                        if (v10 >= v10) {
                                                                            var v34 uint8 = 42
                                                                            var v35 int64 = 0
                                                                            _ = v34
                                                                            _ = v35
                                                                        }
                                                                    }
                                                                case "\n\t":
                                                                    fmt.Println(fi19, v21, v2)
                                                                    fmt.Println(v1)
                                                                    if true {
                                                                        _ = len(v12)
                                                                        if (4294967295 > 18446744073709551615) {
                                                                            fmt.Println(fi22, v10, v10)
                                                                            _ = len(v12)
                                                                        } else {
                                                                            if len(v12) > 0 {
                                                                                v12[0] = (fi22 >> (((v1 << ((v10) & 63))) & 63))
                                                                            }
                                                                            var v36 int = fi19
                                                                            fmt.Println(v1, v10)
                                                                            _ = v36
                                                                        }
                                                                        v9 -= v2
                                                                    }
                                                                default:
                                                                    fmt.Println(fi19, v2)
                                                                }
                                                                _ = v25
                                                            }
                                                        }
                                                    } else {
                                                        var v37 int16 = 1
                                                        var v38 uint = 1
                                                        fmt.Println(v2)
                                                        _ = v37
                                                        _ = v38
                                                    }
                                                    _ = fi22
                                                }
                                                var v39 string = "hello"
                                                _ = v21
                                                _ = v39
                                            }
                                            for fi40 := 0; fi40 < 8; fi40++ {
                                                fmt.Println(fi19, v10, v1)
                                                fmt.Println(fi19, v3)
                                                var v41 string = "\n\t"
                                                fmt.Println(v1, fi19)
                                                if false {
                                                    v9 += v9
                                                    v42 := GenS5{F0: 2147483647}
                                                    v1 = 1
                                                    _ = v42
                                                }
                                                _ = fi40
                                                _ = v41
                                            }
                                        }
                                    }
                                } else {
                                    switch "a" {
                                    case "":
                                        fmt.Println(v10, v3, v3)
                                    case "hello":
                                        for fi43 := 0; fi43 < 7; fi43++ {
                                            fmt.Println(v10, v3)
                                            switch -46 {
                                            case -2147483648:
                                                if len(v12) >= 2 {
                                                    v12 = v12[1:len(v12)]
                                                }
                                            default:
                                                var v44 int16 = int16(v1)
                                                _ = v44
                                            }
                                            _ = fi43
                                        }
                                    case "a":
                                        if false {
                                            for fi45 := 0; fi45 < 7; fi45++ {
                                                fmt.Println(fi19, fi45)
                                                v46 := GenS7{F0: 2, F1: ((("a" + "a") + "a") + "a"), F2: 42}
                                                _ = fi45
                                                _ = v46
                                            }
                                            for fi47 := 0; fi47 < 9; fi47++ {
                                                v12 = append(v12, fi19)
                                                v48 := GenS6{F0: "hello", F1: 70, F2: false}
                                                _ = fi47
                                                _ = v48
                                            }
                                        } else {
                                            if ((false && false) && false) {
                                                var v49 string = ((("\n\t" + "a") + "hello") + (("a" + ("foo bar" + "foo bar")) + "\n\t"))
                                                if len(v12) >= 2 {
                                                    v12 = v12[1:len(v12)]
                                                }
                                                _ = v49
                                            }
                                            var v50 int8 = (v9 * 17)
                                            if len(v12) > 0 {
                                                v12[0] = 9223372036854775807
                                            }
                                            _ = v50
                                        }
                                    default:
                                        var v51 bool = false
                                        _ = v51
                                    }
                                }
                                fmt.Println(v3, v10, v10)
                                _ = fi19
                            }
                            v52 := GenS8{F0: 1, F1: -141, F2: true}
                            if len(v12) > 0 {
                                v12[0] = -2147483648
                            }
                            _ = v52
                        } else {
                            v53 := GenS2{F0: 42}
                            _ = v53
                        }
                    default:
                        if true {
                            for fi54 := 0; fi54 < 6; fi54++ {
                                var v55 int = (fi54 ^ fi54)
                                v56 := GenS9{F0: fi54, F1: 1}
                                _ = fi54
                                _ = v55
                                _ = v56
                            }
                            v57 := GenS9{F0: 128, F1: 1}
                            _ = v57
                        } else {
                            if (v2 <= 1) {
                                var v58 uint16 = 42
                                v12 = append(v12, 2147483647)
                                _ = len(v12)
                                _ = v58
                            }
                            var v59 int16 = 789
                            v60 := GenS7{F0: 32767, F1: "", F2: 115}
                            _ = v59
                            _ = v60
                        }
                    }
                    var v61 int16 = 42
                    fmt.Println(v10, v1)
                    _ = v61
                } else {
                    v10 = (^(v1))
                    var v62 uint64 = 541
                    var v63 int16 = 42
                    _ = v62
                    _ = v63
                }
            }
            _ = len(v12)
        } else {
            var v64 uint64 = 0
            for fi65 := 0; fi65 < 6; fi65++ {
                v1 = 42
                _ = fi65
            }
            _ = v64
        }
    default:
        switch "a" {
        case "a":
            fmt.Println(v9)
        }
    }
    switch 9223372036854775807 {
    case 42:
        var v66 int64 = -9223372036854775808
        _ = v66
    default:
        v67 := GenS7{F0: 32767, F1: "\n\t", F2: -411}
        _ = v67
    }
    _ = len(v12)
    v10 -= (v10 << ((v1) & 63))
    for fi68 := 0; fi68 < 16; fi68++ {
        switch "foo bar" {
        case "\n\t":
            v1 += v1
        case "foo bar":
            v3 = v3
        case "":
            if len(v12) >= 2 {
                v12 = v12[1:len(v12)]
            }
        }
        if true {
            var v69 float32 = float32((0.1 * float64(0.0)))
            var v70 int16 = 42
            fmt.Println(v2, v69, v9)
            _ = v69
            _ = v70
        }
        _ = len(v12)
        _ = fi68
    }
    var v71 float32 = ((1.5 * 3.14159) - 3.14159)
    v72 := GenS10{F0: 599}
    var v73 int16 = 32767
    v1 -= v1
    v10 += v10
    switch -874 {
    case 0:
        var v74 int = 9223372036854775807
        _ = v74
    case 42:
        var v75 string = (("a" + "a") + (("a" + ((("foo bar" + "foo bar") + "\n\t") + "a")) + "hello"))
        _ = v75
    default:
        v73 += 186
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v71:", v71)
    fmt.Println("v73:", v73)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v71
    _ = v72
    _ = v73
    _ = v9
}
