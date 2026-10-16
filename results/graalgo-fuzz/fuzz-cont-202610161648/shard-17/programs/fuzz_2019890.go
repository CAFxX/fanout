package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 string; F1 int32; F2 string }
type GenS4 struct { F0 uint16; F1 int32; F2 string }
type GenS5 struct { F0 uint8; F1 uint32; F2 int8 }
type GenS6 struct { F0 int }
type GenS7 struct { F0 uint64; F1 uint; F2 int16 }
type GenS8 struct { F0 int8; F1 uint16 }
type GenS9 struct { F0 bool; F1 int }

func main() {
    var v1 int32 = -1
    var v2 bool = (true || false)
    var v3 uint8 = 160
    fmt.Println(v2, v3, v3)
    fmt.Println(v1)
    if v2 {
        var v4 string = (("" + "foo bar") + "a")
        if (!((42 <= 123) || ((65535 < 4294967295) && ((-106 <= 0) || v2)))) {
            var v5 float32 = 0.1
            var v6 int64 = 1
            var v7 int = 128
            _ = v5
            _ = v6
            _ = v7
        } else {
            v8 := []int{745, -464, -1}
            if false {
                if (4294967295 == 1) {
                    if (-2.25 >= ((0.1 + ((float32((1.5 - 0.1)) * 1.5) + 1.5)) + 1.5)) {
                        if (!((v1 << ((255) & 63)) <= v1)) {
                            var v9 int8 = -44
                            _ = v9
                        } else {
                            for fi10 := 0; fi10 < 15; fi10++ {
                                if (v2 || v2) {
                                    var v11 uint16 = 879
                                    _ = v11
                                } else {
                                    for fi12 := 0; fi12 < 14; fi12++ {
                                        var v13 uint8 = 160
                                        var v14 int = 32768
                                        switch v4 {
                                        case "foo bar":
                                            v15 := GenS1{F0: 127}
                                            var v16 int8 = -1
                                            _ = v15
                                            _ = v16
                                        case "":
                                            if v2 {
                                                var v17 int32 = (-(v1))
                                                v2 = ((^(v3)) <= (v13 << ((1) & 63)))
                                                _ = v17
                                            }
                                        }
                                        _ = fi12
                                        _ = v13
                                        _ = v14
                                    }
                                    for fi18 := 0; fi18 < 18; fi18++ {
                                        v3 -= v3
                                        fmt.Println(fi18, v2, fi10)
                                        _ = fi18
                                    }
                                }
                                var v19 int64 = 0
                                v20 := GenS1{F0: v3}
                                _ = fi10
                                _ = v19
                                _ = v20
                            }
                            fmt.Println(v1, v3)
                        }
                        v21 := GenS2{F0: true}
                        _ = v21
                    }
                }
                v22 := GenS3{F0: "\n\t", F1: 42, F2: ((v4 + "a") + "foo bar")}
                _ = v22
            } else {
                v23 := GenS2{F0: v2}
                _ = v23
            }
            var v24 float64 = (((-2.25 + 1e10) + 0.1) - 0.1)
            _ = v24
            _ = v8
        }
        var v25 int16 = 290
        _ = v25
        _ = v4
    } else {
        for fi26 := 0; fi26 < 17; fi26++ {
            v27 := GenS4{F0: 65535, F1: v1, F2: ""}
            v1 = v1
            _ = fi26
            _ = v27
        }
    }
    v3 += 42
    v2 = true
    switch -54 {
    case -454:
        v28 := []int{2147483647, 127, 32767, 778}
        _ = v28
    case 760:
        switch 962 {
        case 1:
            var v29 float32 = 3.14159
            _ = v29
        default:
            v3 = v3
        }
    case 924:
        if (468 == 255) {
            if v2 {
                var v30 int = 127
                switch "" {
                case "\n\t":
                    var v31 uint64 = 18446744073709551615
                    _ = v31
                case "hello":
                    for fi32 := 0; fi32 < 2; fi32++ {
                        for fi33 := 0; fi33 < 10; fi33++ {
                            fmt.Println(fi33)
                            v34 := []int{32767, v30, v30, (fi32 % ((2147483647) | 1))}
                            _ = fi33
                            _ = v34
                        }
                        _ = fi32
                    }
                case "a":
                    v30 -= -9223372036854775808
                }
                v35 := []int{v30}
                _ = v30
                _ = v35
            } else {
                v2 = false
                v36 := GenS5{F0: v3, F1: 541, F2: 0}
                _ = v36
            }
            v37 := []int{4294967295, -131, -9223372036854775808, 127}
            if v2 {
                switch -39 {
                case -678:
                    if true {
                        for fi38 := 0; fi38 < 2; fi38++ {
                            v39 := GenS6{F0: fi38}
                            switch "a" {
                            case "\n\t":
                                fmt.Println(v3, fi38)
                            case "":
                                fmt.Println(v3, v1)
                            case "hello":
                                v2 = (-1 > 6)
                            default:
                                v1 -= (v1 * v1)
                            }
                            v40 := GenS6{F0: fi38}
                            _ = fi38
                            _ = v39
                            _ = v40
                        }
                        for fi41 := 0; fi41 < 7; fi41++ {
                            switch "hello" {
                            case "\n\t":
                                v2 = ("\n\t" <= "foo bar")
                            case "a":
                                var v42 float64 = -2.25
                                _ = v42
                            }
                            fmt.Println(v2)
                            if (986 != 680) {
                                if v2 {
                                    v2 = ((v2 || (255 <= 0)) || (v2 && ((^(v3)) != v3)))
                                }
                            } else {
                                var v43 uint = 18446744073709551615
                                var v44 int32 = -2147483648
                                v45 := GenS3{F0: "\n\t", F1: v44, F2: "hello"}
                                _ = v43
                                _ = v44
                                _ = v45
                            }
                            _ = fi41
                        }
                    } else {
                        switch -2147483648 {
                        case 1:
                            for fi46 := 0; fi46 < 19; fi46++ {
                                switch "\n\t" {
                                case "a":
                                    var v47 int32 = v1
                                    _ = v47
                                case "foo bar":
                                    for fi48 := 0; fi48 < 17; fi48++ {
                                        var v49 bool = v2
                                        v49 = v2
                                        _ = fi48
                                        _ = v49
                                    }
                                case "":
                                    var v50 int = fi46
                                    _ = v50
                                }
                                if ("a" >= (("" + (((("" + ("a" + ("hello" + ((("\n\t" + ("a" + ("foo bar" + ""))) + "a") + ("hello" + "a"))))) + "foo bar") + "\n\t") + ("\n\t" + ("foo bar" + "\n\t")))) + "hello")) {
                                    if len(v37) > 0 {
                                        v37[0] = fi46
                                    }
                                    for fi51 := 0; fi51 < 14; fi51++ {
                                        if v2 {
                                            var v52 float64 = 1.5
                                            _ = v52
                                        } else {
                                            v1 = v1
                                        }
                                        _ = fi51
                                    }
                                    for fi53 := 0; fi53 < 5; fi53++ {
                                        if v2 {
                                            var v54 uint64 = 521
                                            _ = v54
                                        } else {
                                            v37 = append(v37, fi46)
                                            v55 := GenS4{F0: 647, F1: (v1 >> ((uint(v3)) & 63)), F2: ""}
                                            var v56 int32 = 1
                                            v2 = (((v2 || ("" != "hello")) || (2147483648 != 148)) || (-1 <= 42))
                                            _ = v55
                                            _ = v56
                                        }
                                        var v57 float64 = 0.1
                                        v3 = (v3 - v3)
                                        _ = fi53
                                        _ = v57
                                    }
                                }
                                for fi58 := 0; fi58 < 2; fi58++ {
                                    var v59 int8 = 127
                                    var v60 int = -9223372036854775808
                                    _ = fi58
                                    _ = v59
                                    _ = v60
                                }
                                _ = fi46
                            }
                        case -15:
                            switch 32768 {
                            case -2147483648:
                                var v61 float32 = float32(-2.25)
                                _ = v61
                            case 668:
                                v2 = (true || ((-300 <= 0) || (1 == 42)))
                            }
                        case 984:
                            v2 = true
                        default:
                            for fi62 := 0; fi62 < 19; fi62++ {
                                v1 = v1
                                _ = fi62
                            }
                        }
                        var v63 int16 = 42
                        _ = v63
                    }
                case -9223372036854775808:
                    for fi64 := 0; fi64 < 5; fi64++ {
                        if (715 != 2) {
                            var v65 uint32 = 0
                            for fi66 := 0; fi66 < 2; fi66++ {
                                v2 = true
                                _ = len(v37)
                                if v2 {
                                    v2 = (3.14159 == 0.0)
                                    if true {
                                        fmt.Println(v65, fi66)
                                    } else {
                                        v2 = v2
                                    }
                                } else {
                                    v3 -= v3
                                    var v67 uint64 = 42
                                    _ = v67
                                }
                                _ = fi66
                            }
                            v37 = append(v37, fi64)
                            _ = v65
                        }
                        _ = fi64
                    }
                default:
                    _ = len(v37)
                }
                if (!((v3 ^ (v3 % ((202) | 1))) == 144)) {
                    if (103 <= 866) {
                        if (32767 == 136) {
                            v68 := GenS5{F0: v3, F1: 2147483647, F2: -1}
                            _ = len(v37)
                            v69 := GenS7{F0: 714, F1: 65535, F2: 32767}
                            _ = v68
                            _ = v69
                        }
                    } else {
                        var v70 int8 = 1
                        _ = v70
                    }
                    if len(v37) > 0 {
                        v37[0] = 9223372036854775807
                    }
                }
                fmt.Println(v2, v3)
            } else {
                v3 -= v3
                v2 = (((float64(0.0) - 3.14159) + float64(0.0)) == (1.5 * 0.0))
                var v71 uint32 = 1
                _ = v71
            }
            _ = v37
        } else {
            v3 = 2
            for fi72 := 0; fi72 < 19; fi72++ {
                var v73 uint16 = 2
                _ = fi72
                _ = v73
            }
        }
    default:
        var v74 float64 = (1.5 + 1e10)
        _ = v74
    }
    v75 := GenS6{F0: 0}
    v76 := GenS7{F0: 656, F1: 255, F2: -1}
    fmt.Println(v1, v3, v1)
    var v77 uint64 = 42
    switch "" {
    case "\n\t":
        var v78 int = 818
        _ = v78
    case "hello":
        switch ("hello" + "\n\t") {
        case "hello":
            v2 = (0 == 0)
        case "\n\t":
            var v79 int64 = 9223372036854775807
            _ = v79
        case "":
            if false {
                v80 := []int{42, -2147483648, 277, -750}
                v77 = v77
                if (65535 >= 65535) {
                    v81 := GenS8{F0: 1, F1: 1}
                    var v82 int8 = 0
                    _ = v81
                    _ = v82
                }
                _ = v80
            }
        default:
            v77 = 1
        }
    default:
        fmt.Println(v1, v3, v77)
    }
    var v83 int32 = (v1 >> ((2) & 63))
    v84 := GenS9{F0: (!(42 < -603)), F1: 1}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v77:", v77)
    fmt.Println("v83:", v83)
    _ = v1
    _ = v2
    _ = v3
    _ = v75
    _ = v76
    _ = v77
    _ = v83
    _ = v84
}
