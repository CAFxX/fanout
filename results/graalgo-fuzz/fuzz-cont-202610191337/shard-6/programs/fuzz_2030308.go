package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int64; F1 bool; F2 bool }
type GenS3 struct { F0 string; F1 int }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 uint16; F1 int64; F2 int8 }

func main() {
    v1 := []int{}
    var v2 float64 = 3.14159
    var v3 uint64 = 4294967295
    var v4 uint16 = 32767
    var v5 uint16 = (v4 / (((v4 + (v4 | (v4 * v4)))) | 1))
    var v6 int64 = -9223372036854775808
    v7 := GenS1{F0: true}
    if len(v1) > 0 {
        v1[0] = 2147483647
    }
    var v8 uint16 = (v5 | v4)
    fmt.Println(v6, v8, v8)
    v9 := GenS1{F0: (!true)}
    v10 := GenS1{F0: (0 >= 431)}
    _ = len(v1)
    fmt.Println(v6, v2, v2)
    var v11 int8 = -36
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    if len(v1) > 0 {
        v1[0] = 2147483647
    }
    v12 := GenS1{F0: (v6 != -441)}
    if (653 == -2147483648) {
        v13 := GenS2{F0: -9223372036854775808, F1: (false || false), F2: (true || (false && ((-818 < 4294967295) && (false && true))))}
        for fi14 := 0; fi14 < 4; fi14++ {
            switch fi14 {
            case -563:
                if (v2 <= 3.14159) {
                    if len(v1) > 0 {
                        v1[0] = (fi14 >> ((255) & 63))
                    }
                    var v15 uint16 = (v8 / ((uint16(v3)) | 1))
                    for fi16 := 0; fi16 < 4; fi16++ {
                        v1 = append(v1, fi14)
                        _ = len(v1)
                        if ((("" + "") + "foo bar") > (("hello" + "\n\t") + ("" + "a"))) {
                            var v17 uint8 = 1
                            v18 := GenS2{F0: (-(v6)), F1: false, F2: (!(42 == 127))}
                            var v19 int32 = 0
                            v2 = v2
                            _ = v17
                            _ = v18
                            _ = v19
                        } else {
                            if (-854 == 0) {
                                fmt.Println(v2)
                                v20 := GenS2{F0: 1, F1: true, F2: false}
                                switch "" {
                                case "\n\t":
                                    switch "" {
                                    case "hello":
                                        v15 = (v15 >> ((435) & 63))
                                        if true {
                                            var v21 uint16 = (^(v8))
                                            v15 = 255
                                            _ = v21
                                        }
                                        v6 -= (v6 ^ (v6 << ((1) & 63)))
                                    default:
                                        var v22 int16 = -1
                                        _ = v22
                                    }
                                case "a":
                                    var v23 uint = 2
                                    var v24 int = int(v6)
                                    if true {
                                        v23 = (^(v23))
                                    }
                                    _ = v23
                                    _ = v24
                                case "":
                                    v1 = append(v1, (fi16 ^ (fi14 + 2147483647)))
                                    if true {
                                        if len(v1) > 0 {
                                            v1[0] = (fi14 * fi14)
                                        }
                                    }
                                    _ = len(v1)
                                    v25 := GenS2{F0: (v6 - 9223372036854775807), F1: true, F2: (!false)}
                                    _ = v25
                                default:
                                    v1 = append(v1, 2147483647)
                                    v8 -= 42
                                    fmt.Println(v3, v3, v6)
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                }
                                _ = v20
                            } else {
                                v8 -= 0
                                if true {
                                    v1 = append(v1, fi14)
                                    var v26 int64 = (^(v6))
                                    _ = len(v1)
                                    _ = v26
                                } else {
                                    v3 += 2
                                    var v27 uint8 = 127
                                    v28 := GenS2{F0: (v6 % ((42) | 1)), F1: (v15 == (^(v4))), F2: true}
                                    v29 := GenS3{F0: (("a" + ("" + ("hello" + "hello"))) + (("hello" + "foo bar") + "a")), F1: 9223372036854775807}
                                    var v30 int16 = 42
                                    _ = v27
                                    _ = v28
                                    _ = v29
                                    _ = v30
                                }
                            }
                            var v31 uint64 = (v3 - v3)
                            _ = v31
                        }
                        var v32 uint16 = v8
                        fmt.Println(v3, fi14, v3)
                        _ = fi16
                        _ = v32
                    }
                    _ = v15
                } else {
                    v6 = int64(fi14)
                    v3 = v3
                }
            case -2147483648:
                var v33 float64 = float64(0.1)
                _ = v33
            case 1:
                var v34 string = "foo bar"
                _ = v34
            default:
                v6 = 42
            }
            for fi35 := 0; fi35 < 20; fi35++ {
                switch "a" {
                case "":
                    v6 = (-(v6))
                case "hello":
                    v8 = v8
                case "a":
                    v36 := GenS4{F0: (v11 & -1)}
                    _ = v36
                default:
                    v37 := GenS2{F0: v6, F1: true, F2: true}
                    _ = v37
                }
                if true {
                    fmt.Println(v4)
                    var v38 int8 = v11
                    switch ("foo bar" + ("" + "a")) {
                    case "a":
                        v39 := GenS5{F0: v8, F1: (v6 * (v6 + (v6 ^ v6))), F2: (^(v11))}
                        _ = v39
                    }
                    _ = v38
                }
                fmt.Println(fi35)
                _ = fi35
            }
            if (false && ((fi14 << ((uint(v4)) & 63)) <= 32767)) {
                var v40 float32 = 3.14159
                v6 -= 0
                var v41 int = fi14
                _ = v40
                _ = v41
            } else {
                var v42 uint64 = (v3 - 626)
                var v43 int = (fi14 | 1)
                _ = v42
                _ = v43
            }
            _ = fi14
        }
        var v44 bool = (((false && false) || (v4 <= v4)) && true)
        _ = v13
        _ = v44
    } else {
        fmt.Println(v3)
    }
    for fi45 := 0; fi45 < 2; fi45++ {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        var v46 uint16 = v8
        for fi47 := 0; fi47 < 8; fi47++ {
            var v48 uint = 2
            if true {
                v48 += (v48 | v48)
                var v49 string = "a"
                _ = v49
            } else {
                v48 -= (v48 % ((515) | 1))
                var v50 int = (-(fi47))
                _ = v50
            }
            v48 = v48
            _ = fi47
            _ = v48
        }
        _ = fi45
        _ = v46
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
