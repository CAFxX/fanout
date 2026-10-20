package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32 }
type GenS2 struct { F0 uint32; F1 int }
type GenS3 struct { F0 uint64; F1 uint16 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 string; F1 uint8; F2 int }
type GenS6 struct { F0 uint64; F1 uint8 }
type GenS7 struct { F0 uint }
type GenS8 struct { F0 uint8; F1 string }

func main() {
    v1 := make(map[int]uint32)
    var v2 int16 = -943
    var v3 uint32 = 147
    v3 += 2
    _ = len(v1)
    if true {
        var v4 uint64 = 0
        _ = v4
    }
    var v5 int64 = int64(v2)
    v2 = v2
    var v6 uint = 908
    v7 := []int{-893, 9223372036854775807, 2147483647}
    v7 = append(v7, 32768)
    v5 = (v5 ^ 146)
    if ((v2 - v2) < (^(v2))) {
        var v8 float64 = (0.1 + (0.1 / (1e10 / 1e10)))
        fmt.Println(v3)
        _ = v8
    } else {
        v9 := GenS1{F0: 218, F1: (v3 ^ 4294967295)}
        var v10 uint8 = 138
        _ = v10
        _ = v9
    }
    _ = len(v1)
    v2 = v2
    if false {
        switch 699 {
        case 915:
            if true {
                if false {
                    if false {
                        v11 := GenS2{F0: (v3 | 2), F1: 1}
                        if true {
                            fmt.Println(v6, v3, v6)
                            v12 := GenS1{F0: 2, F1: (^(v3))}
                            _ = v12
                        }
                        var v13 uint64 = 42
                        _ = v11
                        _ = v13
                    } else {
                        fmt.Println(v2, v3, v3)
                        v6 = (v6 / ((v6) | 1))
                        if len(v7) >= 2 {
                            v7 = v7[1:len(v7)]
                        }
                    }
                    v7 = append(v7, 295)
                } else {
                    switch ("foo bar" + "a") {
                    case "foo bar":
                        v14 := GenS2{F0: v3, F1: 9223372036854775807}
                        _ = v14
                    case "\n\t":
                        var v15 uint16 = 32767
                        _ = v15
                    case "":
                        v5 = (^(v5))
                    }
                }
                v3 = v3
                if (0.1 == 1e10) {
                    fmt.Println(v3)
                    for fi16 := 0; fi16 < 13; fi16++ {
                        v2 = (v2 % ((v2) | 1))
                        var v17 uint = (^(v6))
                        _ = fi16
                        _ = v17
                    }
                    v2 = 1
                } else {
                    switch "" {
                    case "\n\t":
                        v18 := GenS3{F0: 700, F1: 230}
                        _ = v18
                    case "foo bar":
                        switch 555 {
                        case 1:
                            for fi19 := 0; fi19 < 5; fi19++ {
                                var v20 int8 = -57
                                var v21 int64 = 42
                                var v22 int16 = (v2 << ((v6) & 63))
                                _ = fi19
                                _ = v20
                                _ = v21
                                _ = v22
                            }
                        case -392:
                            for fi23 := 0; fi23 < 8; fi23++ {
                                var v24 uint8 = 0
                                _ = fi23
                                _ = v24
                            }
                        case 379:
                            if false {
                                var v25 int16 = v2
                                _ = v25
                            } else {
                                for fi26 := 0; fi26 < 4; fi26++ {
                                    switch -750 {
                                    case 127:
                                        switch "" {
                                        case "a":
                                            fmt.Println(v6, v3, v2)
                                        default:
                                            switch (fi26 ^ fi26) {
                                            case -680:
                                                v27 := GenS4{F0: 18446744073709551615}
                                                _ = v27
                                            case 9223372036854775807:
                                                switch 32767 {
                                                case -9223372036854775808:
                                                    fmt.Println(v2, v5)
                                                default:
                                                    v28 := GenS5{F0: ((((("\n\t" + (("hello" + ("" + "a")) + "\n\t")) + "hello") + "a") + "\n\t") + "hello"), F1: 42, F2: int(v2)}
                                                    _ = v28
                                                }
                                            default:
                                                if ((v6 - (^(v6))) > (v6 * v6)) {
                                                    v29 := GenS6{F0: 1, F1: 128}
                                                    var v30 uint = 42
                                                    var v31 uint8 = 128
                                                    _ = v29
                                                    _ = v30
                                                    _ = v31
                                                }
                                            }
                                        }
                                    case 756:
                                        var v32 uint32 = (v3 >> (((v6 - (v6 & (v6 ^ (v6 << ((42) & 63)))))) & 63))
                                        _ = v32
                                    case 9223372036854775807:
                                        var v33 int = (^(fi26))
                                        _ = v33
                                    default:
                                        var v34 bool = (0.0 >= (0.0 + 0.1))
                                        _ = v34
                                    }
                                    if (int32(v5) > 42) {
                                        v5 = v5
                                        switch "foo bar" {
                                        case "\n\t":
                                            v3 += 956
                                        default:
                                            var v35 uint32 = v3
                                            _ = v35
                                        }
                                    } else {
                                        if true {
                                            if ((v5 * (v5 + v5)) == (v5 + 42)) {
                                                v6 += v6
                                                fmt.Println(v2, v2, v6)
                                            } else {
                                                var v36 uint32 = v3
                                                _ = v36
                                            }
                                        } else {
                                            var v37 float64 = 3.14159
                                            for fi38 := 0; fi38 < 17; fi38++ {
                                                v5 = (v5 % ((v5) | 1))
                                                v3 -= (v3 >> ((440) & 63))
                                                _ = fi38
                                            }
                                            v39 := v1[(fi26 & -467)]
                                            _ = v37
                                            _ = v39
                                        }
                                        v3 += v3
                                        v3 += v3
                                    }
                                    _ = fi26
                                }
                            }
                        }
                    }
                }
            } else {
                var v40 int = 1
                switch 2147483647 {
                case 128:
                    var v41 bool = true
                    _ = v41
                case 2147483647:
                    fmt.Println(v3)
                case 1:
                    v42 := GenS7{F0: v6}
                    _ = v42
                default:
                    for fi43 := 0; fi43 < 14; fi43++ {
                        if (!true) {
                            fmt.Println(v6, v2)
                            v44 := v1[(fi43 << ((v6) & 63))]
                            _ = v44
                        }
                        _ = fi43
                    }
                }
                v45 := GenS7{F0: 945}
                _ = v40
                _ = v45
            }
        case 927:
            var v46 uint8 = 2
            _ = v46
        default:
            v47 := GenS8{F0: 19, F1: (("\n\t" + "") + "\n\t")}
            _ = v47
        }
        v1[128] = v3
    } else {
        var v48 uint16 = 1
        fmt.Println(v48, v2, v5)
        _ = v48
    }
    if false {
        var v49 string = (("a" + "") + "\n\t")
        v3 += v3
        _ = v49
    } else {
        for fi50 := 0; fi50 < 18; fi50++ {
            var v51 uint32 = (v3 % (((v3 % ((v3) | 1))) | 1))
            _ = fi50
            _ = v51
        }
        fmt.Println(v3, v2)
    }
    v5 += v5
    var v52 uint = (v6 ^ (v6 - v6))
    var v53 string = "hello"
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    fmt.Println("v52:", v52)
    fmt.Println("v53:", v53)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v5
    _ = v52
    _ = v53
    _ = v6
    _ = v7
}
