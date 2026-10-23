package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32 }
type GenS2 struct { F0 bool; F1 int; F2 uint16 }
type GenS3 struct { F0 bool; F1 int16 }
type GenS4 struct { F0 int; F1 uint64; F2 uint32 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 int8; F1 uint64 }
type GenS7 struct { F0 uint; F1 int32 }
func genFn0(p0 uint8, p1 string) int8 {
    p0 += p0
    fmt.Println(p1, p1)
    var v6 int32 = 966
    v7 := []int{0}
    switch "foo bar" {
    case "hello":
        var v8 uint8 = p0
        _ = v8
    case "\n\t":
        p1 = ""
    case "a":
        v7 = append(v7, int(p0))
    }
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 127
}

func main() {
    v1 := make(map[int]bool)
    v2 := []uint32{2147483647, 895, 898, 2147483648}
    var v3 string = "a"
    v4 := make(map[int]string)
    var v5 uint = 0
    if true {
        if true {
            v3 = "\n\t"
            var v9 uint = 492
            _ = v9
        } else {
            v5 = 984
            switch "hello" {
            case "a":
                var v10 bool = true
                _ = v10
            default:
                v3 = v3
            }
            v5 = v5
        }
    }
    var v11 int = 32768
    v11 -= (v11 << ((v5) & 63))
    for fi12 := 0; fi12 < 16; fi12++ {
        v5 = v5
        _ = fi12
    }
    if false {
        var v13 int32 = 42
        if ((!(-9223372036854775808 > -9223372036854775808)) || (true || ((0.0 + 0.1) > (1.5 + (float64(1.5) - (-2.25 * (3.14159 * 1.5))))))) {
            v5 = (v5 >> (((v5 / (((v5 % (((v5 | (v5 * (v5 + v5)))) | 1))) | 1))) & 63))
            v13 -= v13
        } else {
            for fi14 := 0; fi14 < 16; fi14++ {
                if (57 > -72) {
                    switch (v11 + (fi14 * -744)) {
                    case 2147483647:
                        v13 = (v13 - v13)
                    case 742:
                        fmt.Println(v11, v5)
                    case 32767:
                        var v15 int16 = 747
                        _ = v15
                    default:
                        if true {
                            v3 = v3
                            v16 := GenS1{F0: 2147483648, F1: 2147483648}
                            if (!true) {
                                for fi17 := 0; fi17 < 11; fi17++ {
                                    v18 := GenS2{F0: false, F1: v11, F2: 211}
                                    _ = fi17
                                    _ = v18
                                }
                                var v19 string = "hello"
                                if (42 > 42) {
                                    fmt.Println(v19, v13)
                                    for fi20 := 0; fi20 < 5; fi20++ {
                                        fmt.Println(v5, fi20, v11)
                                        switch int(v13) {
                                        case 0:
                                            if (!(-9223372036854775808 != 9223372036854775807)) {
                                                v13 -= (v13 + (v13 * v13))
                                                v3 = v19
                                            }
                                            v21 := GenS3{F0: (!false), F1: 1000}
                                            var v22 float32 = ((((0.0 / 1.0) - 3.14159) / (1.5 * 1.5)) * float32((((-2.25 - 3.14159) - 1.5) / 1.0)))
                                            _ = v21
                                            _ = v22
                                        case 749:
                                            switch fi14 {
                                            case 4294967295:
                                                var v23 string = ""
                                                _ = v23
                                            case 0:
                                                v3 = (v3 + "hello")
                                            case -506:
                                                v2 = append(v2, 6)
                                            default:
                                                var v24 uint = v5
                                                _ = v24
                                            }
                                        }
                                        _ = fi20
                                    }
                                } else {
                                    v5 += (^(v5))
                                }
                                _ = v19
                            } else {
                                var v25 int64 = -1
                                _ = v25
                            }
                            _ = v16
                        }
                    }
                    var v26 string = v3
                    _ = v26
                } else {
                    var v27 int = (v11 & 4294967295)
                    v28 := GenS4{F0: 896, F1: 18446744073709551615, F2: 0}
                    switch v11 {
                    case 32768:
                        v29 := v1[fi14]
                        _ = v29
                    case 1:
                        v2 = append(v2, 65535)
                    }
                    _ = v27
                    _ = v28
                }
                switch "" {
                case "hello":
                    if false {
                        if true {
                            v3 = ((v3 + ("a" + v3)) + v3)
                        }
                    } else {
                        switch "\n\t" {
                        case "hello":
                            v13 = 2147483647
                        case "a":
                            v2 = append(v2, 4294967295)
                        default:
                            if false {
                                switch (v11 | 2147483647) {
                                case -2147483648:
                                    _ = len(v4)
                                case 489:
                                    v11 = v11
                                }
                                v13 += (^(v13))
                                v3 = (v3 + v3)
                            }
                        }
                    }
                case "\n\t":
                    _ = len(v1)
                case "a":
                    if ((false || (2147483648 >= 835)) || false) {
                        var v30 float64 = 0.0
                        fmt.Println(v5, v11)
                        for fi31 := 0; fi31 < 19; fi31++ {
                            v32 := GenS5{F0: (("\n\t" + "\n\t") + v3)}
                            v11 = fi14
                            v3 = "hello"
                            _ = fi31
                            _ = v32
                        }
                        _ = v30
                    } else {
                        for fi33 := 0; fi33 < 17; fi33++ {
                            v13 = (v13 << ((v5) & 63))
                            switch "a" {
                            case "a":
                                fmt.Println(fi33, v13, v11)
                            }
                            _ = len(v2)
                            _ = fi33
                        }
                        v34 := GenS2{F0: (127 <= 2147483648), F1: (fi14 + fi14), F2: 255}
                        _ = v34
                    }
                }
                _ = fi14
            }
            if (127 > -128) {
                for fi35 := 0; fi35 < 14; fi35++ {
                    for fi36 := 0; fi36 < 20; fi36++ {
                        v3 = ("\n\t" + "")
                        _ = fi36
                    }
                    _ = fi35
                }
            }
        }
        _ = v13
    }
    for fi37 := 0; fi37 < 3; fi37++ {
        v38 := GenS3{F0: false, F1: 786}
        v39 := GenS5{F0: v3}
        if (!(false && true)) {
            v40 := v1[1]
            switch fi37 {
            case 32768:
                _ = len(v2)
            }
            switch (v3 + v3) {
            case "":
                v41 := GenS5{F0: ("foo bar" + ("a" + v3))}
                _ = v41
            case "hello":
                switch (fi37 ^ fi37) {
                case 1:
                    switch -1 {
                    case -1:
                        var v42 uint = (v5 | (v5 * v5))
                        _ = v42
                    }
                case 9223372036854775807:
                    if true {
                        v40 = (!(4294967295 != 65535))
                        fmt.Println(v5, v5)
                    }
                case 127:
                    if false {
                        v43 := GenS3{F0: (false && v40), F1: 1000}
                        switch "a" {
                        case "":
                            switch (fi37 >> ((v5) & 63)) {
                            case -512:
                                v40 = false
                            case 333:
                                v44 := GenS4{F0: (v11 % ((-1) | 1)), F1: 897, F2: 42}
                                _ = v44
                            default:
                                for fi45 := 0; fi45 < 20; fi45++ {
                                    v3 = (v3 + v3)
                                    v5 += v5
                                    _ = fi45
                                }
                            }
                        case "hello":
                            var v46 uint16 = 481
                            _ = v46
                        case "\n\t":
                            if (v3 <= v3) {
                                for fi47 := 0; fi47 < 14; fi47++ {
                                    if ((v3 + ("foo bar" + "hello")) <= "hello") {
                                        v11 -= (fi37 * (v11 << ((v5) & 63)))
                                    } else {
                                        switch (v11 ^ fi47) {
                                        case -2147483648:
                                            switch ("hello" + v3) {
                                            case "foo bar":
                                                v5 -= v5
                                            case "\n\t":
                                                var v48 uint32 = 1
                                                v11 += v11
                                                _ = v48
                                            case "a":
                                                delete(v1, 4294967295)
                                            }
                                        case 2147483647:
                                            var v49 uint64 = 711
                                            _ = v49
                                        }
                                        if (645 <= 2147483648) {
                                            if (0.1 != (-2.25 + 0.1)) {
                                                switch "\n\t" {
                                                case "foo bar":
                                                    var v50 int = fi47
                                                    _ = v50
                                                default:
                                                    if ((-1 > 42) || (791 == 721)) {
                                                        if (genFn0(15, v3) != 0) {
                                                            if true {
                                                                v3 = v3
                                                            } else {
                                                                v40 = v40
                                                            }
                                                            fmt.Println(fi47)
                                                            var v51 uint64 = 0
                                                            var v52 uint8 = 202
                                                            _ = v51
                                                            _ = v52
                                                        } else {
                                                            var v53 uint8 = 127
                                                            _ = v53
                                                        }
                                                        v54 := GenS6{F0: -6, F1: 42}
                                                        delete(v4, (^(fi37)))
                                                        _ = v54
                                                    } else {
                                                        if v40 {
                                                            v55 := GenS1{F0: 2, F1: 640}
                                                            _ = v55
                                                        } else {
                                                            fmt.Println(v3, v5, v3)
                                                        }
                                                    }
                                                }
                                            }
                                        } else {
                                            v11 = (fi37 | (-(fi37)))
                                        }
                                    }
                                    if true {
                                        var v56 int = fi37
                                        v56 = v11
                                        _ = v56
                                    } else {
                                        switch -507 {
                                        case 32767:
                                            var v57 int32 = -2147483648
                                            _ = v57
                                        default:
                                            if v40 {
                                                if true {
                                                    v4[32768] = v3
                                                    if true {
                                                        v40 = ((v3 + (v3 + ("" + (v3 + v3)))) > v3)
                                                        v3 = (v3 + "foo bar")
                                                    }
                                                    v5 = v5
                                                } else {
                                                    fmt.Println(fi37, v5, v40)
                                                    var v58 uint16 = 255
                                                    _ = v58
                                                }
                                            }
                                        }
                                        v59 := GenS4{F0: -520, F1: 4294967295, F2: 2147483648}
                                        var v60 uint = (v5 % ((v5) | 1))
                                        _ = v59
                                        _ = v60
                                    }
                                    _ = fi47
                                }
                                var v61 int8 = 127
                                for fi62 := 0; fi62 < 12; fi62++ {
                                    fmt.Println(v5, v5, fi37)
                                    fmt.Println(v3, v40)
                                    _ = fi62
                                }
                                _ = v61
                            }
                        default:
                            v5 = v5
                        }
                        _ = v43
                    }
                }
            default:
                var v63 string = "a"
                _ = v63
            }
            _ = v40
        }
        _ = fi37
        _ = v38
        _ = v39
    }
    if false {
        for fi64 := 0; fi64 < 14; fi64++ {
            var v65 string = ((("hello" + v3) + v3) + v3)
            _ = fi64
            _ = v65
        }
        fmt.Println(v5, v5)
        var v66 int8 = 1
        _ = v66
    }
    v11 += 9223372036854775807
    v67 := GenS5{F0: (("a" + ("\n\t" + (v3 + v3))) + ((v3 + (v3 + (v3 + v3))) + (v3 + v3)))}
    var v68 int16 = 32767
    v3 = "hello"
    switch (v3 + v3) {
    case "\n\t":
        var v69 bool = false
        _ = v69
    case "foo bar":
        v68 -= v68
    case "hello":
        v70 := GenS7{F0: v5, F1: 42}
        _ = v70
    }
    var v71 uint8 = 127
    v72 := GenS7{F0: (^(v5)), F1: int32(v5)}
    v5 += (^(v5))
    if false {
        var v73 int32 = -163
        var v74 uint8 = v71
        if true {
            v5 -= v5
            var v75 int32 = 531
            _ = v75
        }
        _ = v73
        _ = v74
    }
    var v76 int16 = v68
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v68:", v68)
    fmt.Println("v71:", v71)
    fmt.Println("v76:", v76)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v67
    _ = v68
    _ = v71
    _ = v72
    _ = v76
}
