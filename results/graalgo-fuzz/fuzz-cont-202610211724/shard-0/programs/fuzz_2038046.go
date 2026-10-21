package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int8; F1 int32 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 int16; F1 uint16; F2 int }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 uint64; F1 int; F2 int16 }
type GenS8 struct { F0 string; F1 uint8 }
type GenS9 struct { F0 bool; F1 bool; F2 int }
func genFn0(p0 int64) uint8 {
    for fi6 := 0; fi6 < 4; fi6++ {
        v7 := GenS1{F0: 42}
        var v8 uint = 4294967295
        for fi9 := 0; fi9 < 14; fi9++ {
            p0 += p0
            _ = fi9
        }
        _ = fi6
        _ = v7
        _ = v8
    }
    v10 := GenS2{F0: 1}
    var v11 int16 = -343
    if (313 < 18446744073709551615) {
        v11 += 42
        p0 += -9223372036854775808
        var v12 uint = 2
        _ = v12
    } else {
        v13 := GenS3{F0: 1, F1: -2147483648}
        p0 += p0
        for fi14 := 0; fi14 < 3; fi14++ {
            v15 := GenS4{F0: 42}
            for fi16 := 0; fi16 < 12; fi16++ {
                p0 -= (p0 - p0)
                switch "" {
                case "foo bar":
                    var v17 string = "a"
                    _ = v17
                case "hello":
                    v18 := []int{fi16, fi14}
                    _ = v18
                }
                _ = fi16
            }
            _ = fi14
            _ = v15
        }
        _ = v13
    }
    _ = p0
    _ = v10
    _ = v11
    return 0
}
func genFn1() bool {
    var v19 float64 = 1.5
    var v20 int8 = 42
    v19 += (0.0 / (0.0 + (v19 - float64((1e10 + 0.1)))))
    v21 := GenS5{F0: 1, F1: 1, F2: 739}
    _ = v19
    _ = v20
    _ = v21
    return (255 < 2)
}

func main() {
    var v1 uint64 = 2
    v2 := []int{-668}
    var v3 uint16 = 2
    var v4 int8 = -128
    var v5 uint = 254
    v22 := GenS3{F0: (v4 % (((-(v4))) | 1)), F1: -1}
    switch 127 {
    case 4294967295:
        v4 += -103
    default:
        if (42 == 458) {
            v3 -= v3
            switch 1 {
            case -909:
                v2 = append(v2, 42)
            case 9223372036854775807:
                v4 += (v4 & (v4 - (v4 / (((v4 / ((127) | 1))) | 1))))
            }
        } else {
            if (0 == -1) {
                if false {
                    _ = len(v2)
                } else {
                    var v23 int8 = (v4 | v4)
                    var v24 uint16 = 620
                    fmt.Println(v23)
                    _ = v23
                    _ = v24
                }
                var v25 int32 = 1
                _ = v25
            } else {
                var v26 bool = true
                _ = v26
            }
        }
    }
    var v27 string = "a"
    v3 -= 1
    switch v27 {
    case "a":
        for fi28 := 0; fi28 < 2; fi28++ {
            fmt.Println(v5, v5, v27)
            var v29 uint32 = 42
            var v30 float32 = (3.14159 - float32(3.14159))
            _ = fi28
            _ = v29
            _ = v30
        }
    }
    if false {
        v5 = v5
        if true {
            var v31 float64 = (0.1 / 1.5)
            _ = v31
        }
        var v32 uint32 = 2
        _ = v32
    }
    if (v4 != (-(v4))) {
        for fi33 := 0; fi33 < 3; fi33++ {
            if (true || (v3 > (v3 - (v3 + (v3 - (^(v3))))))) {
                v5 += (^(v5))
                if (!false) {
                    var v34 uint = (^(v5))
                    _ = v34
                }
                v35 := GenS6{F0: 1000}
                _ = v35
            } else {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            }
            _ = fi33
        }
    } else {
        var v36 int8 = 100
        var v37 float32 = (1.5 / 1e10)
        _ = v36
        _ = v37
    }
    for fi38 := 0; fi38 < 11; fi38++ {
        switch "hello" {
        case "foo bar":
            v2 = append(v2, fi38)
        case "\n\t":
            if true {
                _ = len(v2)
            } else {
                for fi39 := 0; fi39 < 1; fi39++ {
                    if true {
                        fmt.Println(fi38, v1, v5)
                    } else {
                        v40 := GenS6{F0: 1000}
                        v41 := GenS7{F0: v1, F1: (fi39 & 32768), F2: 651}
                        _ = v40
                        _ = v41
                    }
                    v3 += (v3 - 42)
                    _ = fi39
                }
                var v42 uint = v5
                _ = v42
            }
        default:
            v1 -= (v1 % ((v1) | 1))
        }
        for fi43 := 0; fi43 < 14; fi43++ {
            v44 := GenS4{F0: 1}
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            _ = fi43
            _ = v44
        }
        _ = fi38
    }
    var v45 float32 = 1.5
    v46 := GenS5{F0: -333, F1: (v3 - (v3 % (((^(v3))) | 1))), F2: 128}
    if (-2147483648 == -378) {
        switch v27 {
        case "hello":
            v47 := GenS8{F0: "foo bar", F1: 231}
            _ = v47
        case "":
            if ((("" + v27) + v27) < ((v27 + ((("hello" + "foo bar") + (v27 + "hello")) + ("" + ("a" + v27)))) + "")) {
                v2 = append(v2, 2147483647)
            } else {
                v5 = 4294967295
                v5 = v5
                fmt.Println(v3, v27, v1)
            }
        case "foo bar":
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        default:
            switch "" {
            case "\n\t":
                var v48 int64 = 9223372036854775807
                _ = v48
            case "hello":
                v4 -= v4
            case "a":
                v49 := GenS7{F0: (v1 << ((v5) & 63)), F1: 128, F2: -151}
                _ = v49
            default:
                if false {
                    v5 = v5
                    v50 := GenS1{F0: 1}
                    _ = v50
                } else {
                    for fi51 := 0; fi51 < 3; fi51++ {
                        v52 := GenS4{F0: 1}
                        fmt.Println(v27, v4, v5)
                        _ = fi51
                        _ = v52
                    }
                    for fi53 := 0; fi53 < 20; fi53++ {
                        fmt.Println(v3, v3)
                        if (true || ((v4 + 127) > v4)) {
                            v1 = (^(v1))
                        } else {
                            switch (fi53 - (fi53 ^ -1)) {
                            case 32768:
                                if len(v2) > 0 {
                                    v2[0] = (-(fi53))
                                }
                            default:
                                for fi54 := 0; fi54 < 12; fi54++ {
                                    var v55 uint8 = 2
                                    if (!true) {
                                        var v56 bool = false
                                        _ = v56
                                    } else {
                                        v27 = v27
                                        fmt.Println(v1, v1)
                                    }
                                    _ = fi54
                                    _ = v55
                                }
                            }
                        }
                        _ = fi53
                    }
                }
            }
        }
        v5 += 65535
    }
    switch -410 {
    case 775:
        v2 = append(v2, -2147483648)
    case 42:
        for fi57 := 0; fi57 < 7; fi57++ {
            var v58 int8 = v4
            _ = fi57
            _ = v58
        }
    case 606:
        for fi59 := 0; fi59 < 20; fi59++ {
            fmt.Println(fi59, v1)
            _ = fi59
        }
    default:
        v5 -= v5
    }
    if (2147483648 > 524) {
        v27 = ((v27 + "hello") + "\n\t")
        v60 := GenS6{F0: 32767}
        switch v27 {
        case "hello":
            var v61 int8 = -128
            _ = v61
        }
        _ = v60
    } else {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        fmt.Println(v45)
    }
    for fi62 := 0; fi62 < 9; fi62++ {
        if true {
            if (genFn0(-9223372036854775808) < 128) {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            } else {
                if (!true) {
                    v63 := GenS1{F0: 1}
                    _ = v63
                }
                var v64 uint16 = (v3 & (v3 - 65535))
                _ = v64
            }
        } else {
            var v65 string = v27
            var v66 int64 = 0
            if ((7 >= 255) && false) {
                var v67 bool = true
                for fi68 := 0; fi68 < 1; fi68++ {
                    fmt.Println(v65)
                    _ = fi68
                }
                v45 += ((v45 + (v45 - -2.25)) - v45)
                _ = v67
            }
            _ = v65
            _ = v66
        }
        _ = len(v2)
        v4 = v4
        _ = fi62
    }
    fmt.Println(v3)
    v69 := GenS9{F0: (142 != 2), F1: false, F2: 4294967295}
    v3 = (v3 << ((65535) & 63))
    var v70 uint64 = 0
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    fmt.Println("v5:", v5)
    fmt.Println("v70:", v70)
    _ = v1
    _ = v2
    _ = v22
    _ = v27
    _ = v3
    _ = v4
    _ = v45
    _ = v46
    _ = v5
    _ = v69
    _ = v70
}
