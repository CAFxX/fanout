package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint32; F1 int8 }
type GenS4 struct { F0 int32; F1 int8 }
type GenS5 struct { F0 bool; F1 int8; F2 int16 }
type GenS6 struct { F0 int64 }
func genFn0(p0 uint32) string {
    v7 := []int{128, 128, 647, -949}
    var v8 float64 = 1.5
    var v9 int16 = -1
    _ = p0
    _ = v7
    _ = v8
    _ = v9
    return "a"
}
func genFn1(p0 int64, p1 uint, p2 uint64) uint8 {
    v10 := GenS1{F0: 1, F1: 127}
    var v11 uint16 = 32767
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    return 127
}

func main() {
    v1 := make(map[int]int)
    v2 := make(map[string]string)
    v3 := []int{731}
    var v4 uint8 = 127
    var v5 string = (("hello" + ("foo bar" + "a")) + ("" + ""))
    v6 := []string{("" + ((("hello" + v5) + "") + v5)), (v5 + v5)}
    for fi12 := 0; fi12 < 16; fi12++ {
        v4 = v4
        for fi13 := 0; fi13 < 13; fi13++ {
            v14 := GenS1{F0: 0, F1: 13}
            v4 = v4
            _ = fi13
            _ = v14
        }
        var v15 float32 = 1e10
        _ = fi12
        _ = v15
    }
    switch ("hello" + "a") {
    case "":
        v4 -= genFn1(-9223372036854775808, 306, 18446744073709551615)
    case "\n\t":
        switch "\n\t" {
        case "a":
            v4 = (v4 | 255)
        default:
            if (false || false) {
                v16 := GenS1{F0: 100, F1: 1}
                _ = v16
            } else {
                switch v5 {
                case "hello":
                    v17 := GenS2{F0: (v4 ^ v4)}
                    _ = v17
                case "":
                    if false {
                        v18 := GenS3{F0: 917, F1: 42}
                        switch 410 {
                        case 32767:
                            v19 := v2[(v5 + v5)]
                            _ = v19
                        }
                        _ = v18
                    } else {
                        switch ("a" + v5) {
                        case "":
                            for fi20 := 0; fi20 < 9; fi20++ {
                                v21 := v1[fi20]
                                v6 = append(v6, (((v5 + (v5 + ("\n\t" + ""))) + v5) + (("a" + "") + ("hello" + v5))))
                                v22 := GenS4{F0: 2147483647, F1: 7}
                                _ = fi20
                                _ = v21
                                _ = v22
                            }
                        }
                        for fi23 := 0; fi23 < 14; fi23++ {
                            for fi24 := 0; fi24 < 1; fi24++ {
                                var v25 int8 = 100
                                fmt.Println(v4, fi24)
                                _ = fi24
                                _ = v25
                            }
                            _ = fi23
                        }
                        v5 = v5
                    }
                default:
                    fmt.Println(v5, v5)
                }
            }
        }
    case "hello":
        v5 = v5
    }
    var v26 uint8 = 170
    if true {
        fmt.Println(v5, v4)
        for fi27 := 0; fi27 < 9; fi27++ {
            for fi28 := 0; fi28 < 14; fi28++ {
                var v29 uint64 = 4294967295
                if false {
                    v29 = (v29 / (((v29 >> ((65535) & 63))) | 1))
                } else {
                    var v30 float64 = (3.14159 * 3.14159)
                    v31 := GenS4{F0: 0, F1: 1}
                    _ = v30
                    _ = v31
                }
                _ = fi28
                _ = v29
            }
            _ = fi27
        }
    } else {
        v32 := GenS5{F0: (42 < 65535), F1: -128, F2: -1}
        _ = v32
    }
    for fi33 := 0; fi33 < 7; fi33++ {
        if true {
            for fi34 := 0; fi34 < 12; fi34++ {
                switch fi33 {
                case -137:
                    var v35 uint64 = 4294967295
                    _ = v35
                default:
                    switch v5 {
                    case "foo bar":
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                    case "":
                        _ = len(v2)
                    case "\n\t":
                        v36 := GenS6{F0: 323}
                        if (true || ((^(fi33)) <= 32767)) {
                            switch v5 {
                            case "":
                                v2[(v5 + (v5 + v5))] = v5
                            case "\n\t":
                                v5 = ""
                            case "foo bar":
                                delete(v2, "")
                            }
                            var v37 string = v5
                            v5 = (((((v5 + "hello") + v37) + "foo bar") + "\n\t") + v37)
                            _ = v37
                        }
                        delete(v2, v5)
                        _ = v36
                    default:
                        fmt.Println(v26, v5, v5)
                    }
                }
                v38 := GenS6{F0: -193}
                _ = fi34
                _ = v38
            }
            var v39 uint16 = 255
            var v40 uint16 = (v39 << ((461) & 63))
            _ = v39
            _ = v40
        } else {
            var v41 float64 = -2.25
            if ((v4 * (v4 * (v4 - 2))) < genFn1(-9223372036854775808, 908, 983)) {
                for fi42 := 0; fi42 < 6; fi42++ {
                    var v43 int = (fi33 + (^(fi42)))
                    if false {
                        if true {
                            var v44 uint16 = 578
                            _ = v44
                        }
                    } else {
                        fmt.Println(v5, v26, v5)
                    }
                    _ = fi42
                    _ = v43
                }
                if (!(false && (!(2 > 128)))) {
                    v3 = append(v3, fi33)
                }
            }
            _ = v41
        }
        fmt.Println(v26, v5, v4)
        _ = fi33
    }
    v45 := GenS6{F0: 895}
    var v46 float32 = 3.14159
    fmt.Println(v26)
    fmt.Println(v46)
    if (42 > 2147483647) {
        v4 = 0
    }
    if ((1.5 + 1e10) <= 3.14159) {
        var v47 uint = 42
        var v48 uint = (v47 & (^(v47)))
        _ = v47
        _ = v48
    } else {
        switch ("foo bar" + "hello") {
        case "hello":
            if ((964 == 2147483648) || false) {
                v49 := GenS3{F0: 2147483647, F1: -128}
                _ = v49
            } else {
                _ = len(v3)
                var v50 uint8 = 0
                _ = v50
            }
        case "":
            v51 := GenS6{F0: 9223372036854775807}
            _ = v51
        default:
            v26 = 42
        }
    }
    v5 = ("hello" + "foo bar")
    v4 = (v26 + (v4 & v26))
    var v52 float64 = -2.25
    switch 2147483647 {
    case 865:
        v4 = genFn1(42, 42, 325)
    case 42:
        switch 363 {
        case 127:
            var v53 bool = false
            _ = v53
        default:
            var v54 int8 = 42
            _ = v54
        }
    case 2147483647:
        v46 -= -2.25
    default:
        switch 0 {
        case 127:
            var v55 uint64 = 764
            _ = v55
        case 42:
            v26 = v26
        }
    }
    var v56 int8 = 42
    var v57 float64 = float64(v46)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v46:", v46)
    fmt.Println("v5:", v5)
    fmt.Println("v52:", v52)
    fmt.Println("v56:", v56)
    fmt.Println("v57:", v57)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v45
    _ = v46
    _ = v5
    _ = v52
    _ = v56
    _ = v57
    _ = v6
}
