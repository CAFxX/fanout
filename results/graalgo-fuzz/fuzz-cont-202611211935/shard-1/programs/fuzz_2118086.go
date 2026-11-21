package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint }
type GenS2 struct { F0 int; F1 uint }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 int64 }

func main() {
    var v1 int = -240
    v2 := []int{(v1 | v1)}
    var v3 int8 = 1
    var v4 int32 = -2147483648
    var v5 uint32 = 2147483648
    var v6 int64 = -1
    v5 -= (v5 ^ v5)
    var v7 float32 = 0.1
    v7 = ((v7 + v7) - 1e10)
    switch "hello" {
    case "":
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    if false {
        var v8 int16 = 0
        if (!true) {
            if false {
                switch "hello" {
                case "\n\t":
                    fmt.Println(v6, v5)
                case "hello":
                    v5 -= (v5 << ((346) & 63))
                case "a":
                    v1 = v1
                }
                v9 := GenS1{F0: (^(v5)), F1: 49}
                _ = v9
            }
        }
        _ = v8
    } else {
        fmt.Println(v6)
        v2 = append(v2, (-(v1)))
        v1 -= v1
    }
    v5 = 4294967295
    v10 := GenS2{F0: -2147483648, F1: 42}
    v5 -= (v5 ^ v5)
    if (true || ((169 != 789) || ((true || (false && true)) || true))) {
        for fi11 := 0; fi11 < 13; fi11++ {
            v6 -= (v6 % (((v6 / ((42) | 1))) | 1))
            _ = fi11
        }
        v12 := GenS2{F0: v1, F1: 0}
        _ = v12
    } else {
        var v13 uint8 = 98
        v3 = v3
        _ = v13
    }
    if false {
        for fi14 := 0; fi14 < 5; fi14++ {
            var v15 int8 = -124
            v5 = v5
            v15 += 1
            _ = fi14
            _ = v15
        }
    } else {
        var v16 bool = false
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        _ = v16
    }
    for fi17 := 0; fi17 < 19; fi17++ {
        v4 = 42
        if (4294967295 >= 470) {
            var v18 string = "a"
            _ = v18
        }
        _ = fi17
    }
    for fi19 := 0; fi19 < 11; fi19++ {
        for fi20 := 0; fi20 < 20; fi20++ {
            var v21 uint16 = 42
            if true {
                fmt.Println(fi19)
                var v22 string = ("a" + "\n\t")
                _ = v22
            } else {
                if (fi19 < 821) {
                    if (false && (false && (42 < (fi20 & -365)))) {
                        var v23 int32 = (v4 & -792)
                        v24 := GenS3{F0: (v5 << ((2) & 63))}
                        _ = v23
                        _ = v24
                    }
                } else {
                    v5 += v5
                }
                v25 := GenS1{F0: (v5 - v5), F1: 0}
                _ = v25
            }
            _ = fi20
            _ = v21
        }
        fmt.Println(v4, fi19)
        if (3.14159 < (((0.1 * 3.14159) / 1.0) / 0.1)) {
            var v26 int = 201
            var v27 int64 = (v6 + (v6 ^ (v6 >> ((2) & 63))))
            v2 = append(v2, fi19)
            _ = v26
            _ = v27
        } else {
            var v28 bool = false
            for fi29 := 0; fi29 < 20; fi29++ {
                var v30 uint32 = 576
                if v28 {
                    fmt.Println(v7)
                    if (255 != 1) {
                        var v31 float32 = float32(1e10)
                        v5 += (v30 - (v30 >> ((65535) & 63)))
                        if v28 {
                            if (v3 < v3) {
                                v5 = 318
                                switch (fi29 | int(v30)) {
                                case -9223372036854775808:
                                    _ = len(v2)
                                }
                                switch ("\n\t" + "\n\t") {
                                case "\n\t":
                                    var v32 int8 = 0
                                    _ = v32
                                case "hello":
                                    v31 -= float32(0.1)
                                    var v33 uint32 = v30
                                    switch "" {
                                    case "\n\t":
                                        switch "hello" {
                                        case "\n\t":
                                            fmt.Println(fi29, v1, v1)
                                        case "a":
                                            var v34 uint8 = 128
                                            _ = v34
                                        }
                                    }
                                    _ = v33
                                case "foo bar":
                                    v30 = (v30 | 58)
                                default:
                                    if v28 {
                                        v2 = append(v2, (v1 % ((2147483647) | 1)))
                                    }
                                }
                            }
                            v4 = v4
                        } else {
                            switch (("foo bar" + ("a" + "hello")) + "a") {
                            case "foo bar":
                                v7 = v31
                            case "\n\t":
                                var v35 int32 = (v4 * v4)
                                _ = v35
                            default:
                                if len(v2) > 0 {
                                    v2[0] = (fi19 ^ fi29)
                                }
                            }
                        }
                        _ = v31
                    }
                    switch "" {
                    case "a":
                        switch v1 {
                        case -2147483648:
                            if v28 {
                                if v28 {
                                    fmt.Println(v1)
                                    v5 += v5
                                    fmt.Println(v6, v7, v5)
                                }
                                var v36 uint = 18446744073709551615
                                _ = v36
                            }
                        case 0:
                            if ((!false) || (true && (("" <= "a") && (v5 < (v5 % (((v5 >> ((573) & 63))) | 1)))))) {
                                var v37 string = "hello"
                                _ = v37
                            } else {
                                if false {
                                    if v28 {
                                        if true {
                                            if v28 {
                                                switch 9223372036854775807 {
                                                case 9223372036854775807:
                                                    v5 -= v30
                                                case 0:
                                                    if (v28 || v28) {
                                                        if true {
                                                            v30 = 0
                                                            v7 = (0.1 / (v7 - float32(3.14159)))
                                                        }
                                                    }
                                                default:
                                                    if (0.1 == 1.5) {
                                                        if (v3 > v3) {
                                                            if (false || (int16(v4) != 679)) {
                                                                v2 = append(v2, (fi19 & fi29))
                                                                fmt.Println(v30, v4, v30)
                                                                var v38 bool = false
                                                                v39 := GenS1{F0: v5, F1: 605}
                                                                fmt.Println(v6, v7)
                                                                _ = v38
                                                                _ = v39
                                                            }
                                                        } else {
                                                            switch ("\n\t" + (("" + "hello") + (((("a" + "\n\t") + "hello") + "\n\t") + "\n\t"))) {
                                                            case "foo bar":
                                                                if len(v2) > 0 {
                                                                    v2[0] = (v1 << ((0) & 63))
                                                                }
                                                            case "hello":
                                                                if (v6 != v6) {
                                                                    if len(v2) > 0 {
                                                                        v2[0] = v1
                                                                    }
                                                                    var v40 int = fi19
                                                                    _ = v40
                                                                } else {
                                                                    switch (v1 >> ((0) & 63)) {
                                                                    case -172:
                                                                        var v41 float64 = (0.1 / -2.25)
                                                                        _ = v41
                                                                    }
                                                                    v30 -= (^(v30))
                                                                }
                                                            case "a":
                                                                v6 = v6
                                                            default:
                                                                var v42 int = (v1 % ((fi19) | 1))
                                                                _ = v42
                                                            }
                                                            if (2 > 4294967295) {
                                                                fmt.Println(v30, v6, v4)
                                                            } else {
                                                                var v43 string = ""
                                                                _ = v43
                                                            }
                                                            fmt.Println(fi19, v6, v28)
                                                        }
                                                    } else {
                                                        if false {
                                                            var v44 float64 = 1.5
                                                            if (true && v28) {
                                                                var v45 float32 = (float32(v44) / 0.1)
                                                                v4 += v4
                                                                _ = v45
                                                            } else {
                                                                fmt.Println(v7, v28)
                                                                v5 += (v30 - 2)
                                                                v5 -= 0
                                                                v46 := GenS3{F0: 553}
                                                                switch "" {
                                                                case "a":
                                                                    var v47 int8 = (v3 + (v3 & (v3 << ((4294967295) & 63))))
                                                                    _ = v47
                                                                case "hello":
                                                                    if len(v2) >= 2 {
                                                                        v2 = v2[1:len(v2)]
                                                                    }
                                                                default:
                                                                    fmt.Println(v5, v4, fi29)
                                                                }
                                                                _ = v46
                                                            }
                                                            fmt.Println(v44)
                                                            var v48 int8 = (-(v3))
                                                            if len(v2) > 0 {
                                                                v2[0] = (v1 << ((42) & 63))
                                                            }
                                                            _ = v44
                                                            _ = v48
                                                        } else {
                                                            v49 := GenS1{F0: (v30 >> ((607) & 63)), F1: 517}
                                                            v50 := GenS1{F0: v5, F1: 18446744073709551615}
                                                            var v51 int32 = (v4 + 0)
                                                            _ = v49
                                                            _ = v50
                                                            _ = v51
                                                        }
                                                        var v52 uint32 = v5
                                                        var v53 float64 = ((1e10 - 3.14159) * -2.25)
                                                        if (128 > 127) {
                                                            v5 += 42
                                                        }
                                                        v28 = (4294967295 < 18446744073709551615)
                                                        _ = v52
                                                        _ = v53
                                                    }
                                                }
                                                v30 = 1
                                            }
                                            switch "hello" {
                                            case "hello":
                                                var v54 int32 = (v4 ^ (v4 / (((-(v4))) | 1)))
                                                _ = v54
                                            case "foo bar":
                                                fmt.Println(v7, v3, v7)
                                            }
                                            fmt.Println(v5, v1)
                                        }
                                        v55 := GenS4{F0: 4294967295}
                                        v56 := GenS5{F0: 97}
                                        var v57 uint32 = 766
                                        v58 := GenS3{F0: v5}
                                        var v59 uint16 = 34
                                        switch ("hello" + "\n\t") {
                                        case "\n\t":
                                            switch "a" {
                                            case "a":
                                                var v60 float32 = v7
                                                _ = v60
                                            }
                                        case "a":
                                            if (("\n\t" + "hello") != "") {
                                                if len(v2) > 0 {
                                                    v2[0] = (fi19 - fi19)
                                                }
                                                if v28 {
                                                    var v61 int64 = v6
                                                    _ = v61
                                                } else {
                                                    _ = len(v2)
                                                    var v62 string = ""
                                                    if v28 {
                                                        if len(v2) > 0 {
                                                            v2[0] = v1
                                                        }
                                                        v2 = append(v2, (fi29 ^ int(v3)))
                                                        v7 += 0.1
                                                    }
                                                    _ = v62
                                                }
                                                v63 := GenS5{F0: 255}
                                                v57 -= v5
                                                if len(v2) > 0 {
                                                    v2[0] = -1
                                                }
                                                _ = v63
                                            }
                                        case "":
                                            v64 := GenS6{F0: -1}
                                            _ = v64
                                        }
                                        v57 -= v5
                                        v65 := GenS5{F0: 0}
                                        v4 -= (v4 | v4)
                                        _ = v55
                                        _ = v56
                                        _ = v57
                                        _ = v58
                                        _ = v59
                                        _ = v65
                                    }
                                } else {
                                    var v66 float32 = v7
                                    _ = v66
                                }
                                v30 = v5
                            }
                        case 9223372036854775807:
                            v67 := GenS1{F0: (^(v5)), F1: 757}
                            _ = v67
                        }
                    default:
                        v68 := GenS6{F0: -475}
                        _ = v68
                    }
                } else {
                    v69 := GenS2{F0: 0, F1: 835}
                    switch "a" {
                    case "\n\t":
                        var v70 bool = (42 > 42)
                        _ = v70
                    default:
                        switch "" {
                        case "hello":
                            _ = len(v2)
                        }
                    }
                    var v71 float64 = 1.5
                    _ = v69
                    _ = v71
                }
                var v72 uint8 = 255
                v73 := GenS7{F0: v6}
                _ = fi29
                _ = v30
                _ = v72
                _ = v73
            }
            _ = v28
        }
        _ = fi19
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
