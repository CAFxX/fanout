package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32; F2 int64 }
type GenS2 struct { F0 uint32; F1 uint16 }
type GenS3 struct { F0 int32; F1 uint32; F2 int64 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 int8; F1 string; F2 uint64 }
type GenS7 struct { F0 uint64; F1 int32 }
func genFn0(p0 int8, p1 uint8, p2 uint) int {
    v4 := GenS1{F0: 32767, F1: 2, F2: -1}
    p1 = p1
    var v5 uint64 = 1
    if (int(p1) > 127) {
        var v6 string = "a"
        var v7 uint32 = 2147483647
        switch "\n\t" {
        case "\n\t":
            var v8 int = 0
            _ = v8
        case "hello":
            if ((2 <= 595) || (543 == 42)) {
                p1 = 0
                fmt.Println(p1, v7)
            }
        case "foo bar":
            var v9 uint8 = (p1 * (p1 * (p1 & p1)))
            _ = v9
        }
        _ = v6
        _ = v7
    } else {
        if (v5 >= v5) {
            fmt.Println(v5, p2, p1)
            for fi10 := 0; fi10 < 13; fi10++ {
                v5 -= v5
                v11 := []int{fi10, fi10, (fi10 * fi10), 128}
                p1 += 0
                _ = fi10
                _ = v11
            }
            v12 := []int{}
            _ = v12
        } else {
            var v13 uint8 = p1
            _ = v13
        }
    }
    for fi14 := 0; fi14 < 15; fi14++ {
        v15 := GenS1{F0: (fi14 + 4294967295), F1: 4294967295, F2: 42}
        _ = fi14
        _ = v15
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return 127
}

func main() {
    v1 := []int{9223372036854775807}
    var v2 float32 = 1.5
    var v3 uint8 = 255
    var v16 uint8 = (v3 ^ v3)
    switch ("foo bar" + "\n\t") {
    case "\n\t":
        var v17 float64 = 0.0
        _ = v17
    case "":
        v18 := GenS2{F0: 42, F1: 32767}
        _ = v18
    default:
        var v19 uint16 = 907
        _ = v19
    }
    if (v3 >= 42) {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    }
    switch 32767 {
    case 684:
        v20 := GenS3{F0: 811, F1: 109, F2: 0}
        _ = v20
    case -576:
        if true {
            if (42 >= 0) {
                v21 := GenS4{F0: 690}
                _ = v21
            }
            v1 = append(v1, 32768)
            if (-1 != -1) {
                v16 += v3
                v22 := GenS5{F0: 255}
                switch "" {
                case "foo bar":
                    var v23 uint16 = 21
                    _ = v23
                default:
                    var v24 uint = 0
                    _ = v24
                }
                _ = v22
            } else {
                if true {
                    switch ("" + "hello") {
                    case "hello":
                        v25 := GenS4{F0: 1}
                        _ = v25
                    case "\n\t":
                        for fi26 := 0; fi26 < 1; fi26++ {
                            if true {
                                for fi27 := 0; fi27 < 9; fi27++ {
                                    var v28 int64 = 1
                                    _ = fi27
                                    _ = v28
                                }
                                v29 := GenS5{F0: 1}
                                var v30 int64 = -9223372036854775808
                                _ = v29
                                _ = v30
                            } else {
                                switch (fi26 + (fi26 << ((65535) & 63))) {
                                case 753:
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                default:
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                }
                                v3 -= v16
                            }
                            if ("" > "foo bar") {
                                switch ("a" + "") {
                                case "":
                                    v2 = v2
                                }
                            } else {
                                v31 := GenS6{F0: -1, F1: "", F2: 0}
                                switch (fi26 - fi26) {
                                case 6:
                                    v2 += 0.1
                                case 246:
                                    var v32 int16 = 1
                                    _ = v32
                                case -1000:
                                    fmt.Println(v3)
                                default:
                                    for fi33 := 0; fi33 < 18; fi33++ {
                                        v2 += v2
                                        var v34 uint32 = 2147483647
                                        _ = fi33
                                        _ = v34
                                    }
                                }
                                _ = v31
                            }
                            _ = fi26
                        }
                    }
                    fmt.Println(v16)
                    var v35 uint16 = 0
                    _ = v35
                } else {
                    fmt.Println(v3, v2)
                }
                if (!false) {
                    fmt.Println(v3, v16)
                    v36 := GenS1{F0: 2147483647, F1: 4294967295, F2: -9223372036854775808}
                    _ = len(v1)
                    _ = v36
                } else {
                    switch -2147483648 {
                    case 9223372036854775807:
                        var v37 int8 = -56
                        _ = v37
                    case -595:
                        fmt.Println(v16, v3, v2)
                    }
                    for fi38 := 0; fi38 < 13; fi38++ {
                        for fi39 := 0; fi39 < 13; fi39++ {
                            v2 -= (3.14159 * v2)
                            v2 = -2.25
                            v40 := GenS7{F0: 4294967295, F1: 2147483647}
                            _ = fi39
                            _ = v40
                        }
                        for fi41 := 0; fi41 < 17; fi41++ {
                            var v42 int = -9223372036854775808
                            var v43 int = (-(fi38))
                            _ = fi41
                            _ = v42
                            _ = v43
                        }
                        _ = fi38
                    }
                    fmt.Println(v3, v2)
                }
                if (false || (-22 < -61)) {
                    var v44 float64 = 0.0
                    _ = v44
                } else {
                    var v45 int8 = 96
                    if false {
                        v16 = 0
                        v1 = append(v1, 127)
                        v3 -= v3
                    }
                    _ = v45
                }
            }
        } else {
            for fi46 := 0; fi46 < 6; fi46++ {
                if false {
                    if (!false) {
                        for fi47 := 0; fi47 < 14; fi47++ {
                            v3 = 1
                            _ = fi47
                        }
                        if (!(1 <= 0)) {
                            for fi48 := 0; fi48 < 4; fi48++ {
                                v3 += 128
                                _ = fi48
                            }
                            v2 = v2
                        } else {
                            if (0 == -690) {
                                if len(v1) > 0 {
                                    v1[0] = genFn0(100, (^(v16)), 218)
                                }
                                if true {
                                    v2 = v2
                                }
                                var v49 int64 = -647
                                _ = v49
                            }
                            if (2147483648 >= 2147483647) {
                                fmt.Println(v16, v2)
                                var v50 uint = 951
                                _ = v50
                            }
                        }
                        fmt.Println(v16, fi46)
                    }
                    v51 := GenS1{F0: fi46, F1: 2147483648, F2: 1}
                    v2 += (float32((1.5 - float64(float32(0.0)))) * v2)
                    _ = v51
                }
                _ = fi46
            }
        }
    default:
        for fi52 := 0; fi52 < 12; fi52++ {
            var v53 uint64 = 18446744073709551615
            _ = fi52
            _ = v53
        }
    }
    var v54 int32 = -712
    v55 := GenS1{F0: -343, F1: 2147483647, F2: -9223372036854775808}
    var v56 uint32 = 0
    fmt.Println(v16)
    v16 = (v3 % (((v3 / ((v16) | 1))) | 1))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v54:", v54)
    fmt.Println("v56:", v56)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v54
    _ = v55
    _ = v56
}
