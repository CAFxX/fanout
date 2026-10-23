package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64; F2 string }
type GenS2 struct { F0 string; F1 int }
type GenS3 struct { F0 uint; F1 uint32; F2 bool }
type GenS4 struct { F0 int16; F1 uint8 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 bool; F1 int8; F2 string }

func main() {
    v1 := make(map[int]uint32)
    var v2 uint64 = 42
    var v3 int16 = 32767
    v4 := []int{42, -54, 9223372036854775807, -472}
    var v5 string = "hello"
    if (1 < 0) {
        switch 42 {
        case -9223372036854775808:
            v2 = v2
        }
        if true {
            if (9223372036854775807 > 0) {
                v6 := GenS1{F0: 32768, F1: 854, F2: "hello"}
                var v7 uint = 0
                v1[2147483647] = 42
                _ = v6
                _ = v7
            } else {
                switch 917 {
                case 4294967295:
                    for fi8 := 0; fi8 < 10; fi8++ {
                        switch fi8 {
                        case 0:
                            if true {
                                fmt.Println(v3, fi8)
                                v9 := GenS2{F0: "foo bar", F1: fi8}
                                _ = v9
                            } else {
                                for fi10 := 0; fi10 < 5; fi10++ {
                                    if (0.0 == ((1e10 + (1e10 + ((3.14159 / float32(float64(float32(1e10)))) * (0.1 + (-2.25 * 0.1))))) * -2.25)) {
                                        var v11 float32 = 0.0
                                        _ = v11
                                    } else {
                                        switch (fi10 - (fi10 >> ((1) & 63))) {
                                        case 32768:
                                            if ((128 <= 199) || (42 >= 65535)) {
                                                fmt.Println(v2)
                                            }
                                            v2 -= v2
                                            var v12 float32 = (-2.25 / 1.5)
                                            var v13 string = (v5 + v5)
                                            _ = v12
                                            _ = v13
                                        case 373:
                                            var v14 string = ("foo bar" + ("foo bar" + (v5 + v5)))
                                            _ = v14
                                        case 9223372036854775807:
                                            var v15 int = (fi8 * fi8)
                                            _ = v15
                                        default:
                                            v2 += (v2 & v2)
                                        }
                                    }
                                    if false {
                                        if true {
                                            if (((!(true && ((v3 & 0) >= (-(v3))))) || (!((3.14159 <= 3.14159) && (v5 < "a")))) || (1 != 2147483647)) {
                                                v16 := GenS1{F0: -2147483648, F1: 9223372036854775807, F2: "a"}
                                                _ = v16
                                            } else {
                                                var v17 int32 = 1
                                                _ = v17
                                            }
                                            if len(v4) > 0 {
                                                v4[0] = fi8
                                            }
                                            v5 = ("foo bar" + (v5 + ((v5 + "hello") + v5)))
                                            v18 := v1[128]
                                            _ = v18
                                        } else {
                                            v2 = (v2 + (v2 % (((v2 >> ((255) & 63))) | 1)))
                                            v19 := GenS2{F0: "hello", F1: 2147483647}
                                            if (108 < 1) {
                                                fmt.Println(v2)
                                                v20 := GenS2{F0: (("" + "a") + "\n\t"), F1: fi8}
                                                var v21 bool = (!(true && ((v3 + (v3 ^ (v3 + v3))) != 32767)))
                                                var v22 uint = 297
                                                _ = v20
                                                _ = v21
                                                _ = v22
                                            } else {
                                                if true {
                                                    v23 := GenS3{F0: 321, F1: 42, F2: false}
                                                    _ = v23
                                                } else {
                                                    if (false && true) {
                                                        switch "hello" {
                                                        case "":
                                                            var v24 int64 = 9223372036854775807
                                                            _ = v24
                                                        case "hello":
                                                            var v25 int = fi10
                                                            _ = v25
                                                        case "\n\t":
                                                            fmt.Println(fi8, v3, fi10)
                                                        }
                                                        v3 -= 1
                                                    } else {
                                                        if false {
                                                            v4 = append(v4, -730)
                                                            v2 -= 1
                                                            v2 -= (v2 >> ((214) & 63))
                                                        } else {
                                                            var v26 uint16 = 65535
                                                            _ = v26
                                                        }
                                                        if len(v4) >= 2 {
                                                            v4 = v4[1:len(v4)]
                                                        }
                                                        _ = len(v1)
                                                    }
                                                    v3 -= -575
                                                    v2 += 42
                                                    var v27 string = ""
                                                    if (!false) {
                                                        var v28 bool = true
                                                        v4 = append(v4, (fi8 / ((fi10) | 1)))
                                                        _ = v28
                                                    }
                                                    _ = v27
                                                }
                                                var v29 int16 = 1
                                                _ = v29
                                            }
                                            _ = v19
                                        }
                                        var v30 int = -9223372036854775808
                                        v2 -= 18446744073709551615
                                        _ = v30
                                    }
                                    v2 += v2
                                    _ = fi10
                                }
                                v2 += v2
                                v2 += (v2 ^ 4294967295)
                            }
                        case -1:
                            if false {
                                v4 = append(v4, fi8)
                            } else {
                                v31 := GenS4{F0: v3, F1: 0}
                                _ = v31
                            }
                        default:
                            v4 = append(v4, fi8)
                        }
                        _ = fi8
                    }
                case -9223372036854775808:
                    delete(v1, -1)
                default:
                    v3 = v3
                }
                var v32 uint16 = 42
                if ("a" != ("foo bar" + "foo bar")) {
                    switch 32767 {
                    case 54:
                        for fi33 := 0; fi33 < 14; fi33++ {
                            v3 = -32768
                            var v34 int8 = -12
                            _ = fi33
                            _ = v34
                        }
                    case 32768:
                        var v35 int16 = 1000
                        _ = v35
                    case -9223372036854775808:
                        for fi36 := 0; fi36 < 2; fi36++ {
                            v32 += 2
                            var v37 uint64 = 960
                            var v38 int16 = v3
                            _ = fi36
                            _ = v37
                            _ = v38
                        }
                    default:
                        var v39 int64 = 1
                        _ = v39
                    }
                }
                _ = v32
            }
            var v40 int32 = 581
            v41 := GenS4{F0: 846, F1: 1}
            _ = v40
            _ = v41
        } else {
            v3 = (v3 >> ((65535) & 63))
        }
        for fi42 := 0; fi42 < 17; fi42++ {
            for fi43 := 0; fi43 < 7; fi43++ {
                var v44 uint8 = 196
                _ = fi43
                _ = v44
            }
            _ = fi42
        }
    }
    v2 += (^(v2))
    fmt.Println(v3)
    if false {
        var v45 int64 = 0
        _ = v45
    } else {
        var v46 uint8 = 2
        v1[466] = 42
        v47 := GenS3{F0: 399, F1: 65535, F2: false}
        _ = v46
        _ = v47
    }
    switch -9223372036854775808 {
    case 232:
        switch "" {
        case "a":
            if false {
                v48 := GenS5{F0: 2}
                for fi49 := 0; fi49 < 8; fi49++ {
                    v50 := GenS6{F0: 56}
                    v2 = v2
                    _ = fi49
                    _ = v50
                }
                var v51 float64 = -2.25
                _ = v48
                _ = v51
            }
        }
    case -1:
        for fi52 := 0; fi52 < 5; fi52++ {
            v3 += 1
            var v53 int32 = int32(v3)
            _ = fi52
            _ = v53
        }
    default:
        var v54 bool = false
        _ = v54
    }
    fmt.Println(v2)
    for fi55 := 0; fi55 < 9; fi55++ {
        var v56 uint64 = v2
        delete(v1, 2147483647)
        if false {
            v57 := GenS7{F0: (true && (32767 < 1)), F1: 1, F2: v5}
            if false {
                fmt.Println(v5, fi55, fi55)
            } else {
                _ = len(v4)
                var v58 uint8 = 203
                var v59 uint32 = 353
                _ = v58
                _ = v59
            }
            v56 = 2
            _ = v57
        } else {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            v2 -= 224
            fmt.Println(v3, v2)
        }
        _ = fi55
        _ = v56
    }
    if true {
        var v60 float32 = -2.25
        _ = v60
    } else {
        v2 = (v2 ^ (^(v2)))
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    }
    delete(v1, -2147483648)
    v5 = (v5 + v5)
    if false {
        fmt.Println(v5)
        if true {
            v3 = (v3 | (v3 * (v3 ^ 32767)))
            v3 -= -886
        } else {
            var v61 bool = (0 > 1)
            v2 = v2
            _ = v61
        }
    } else {
        v5 = v5
    }
    fmt.Println(v2)
    fmt.Println(v3)
    if (!(((!true) && (v2 <= v2)) && ((float64(3.14159) / 3.14159) != -2.25))) {
        for fi62 := 0; fi62 < 8; fi62++ {
            var v63 uint64 = v2
            var v64 float32 = (((float32(-2.25) / 1.0) * 0.0) / 1e10)
            _ = fi62
            _ = v63
            _ = v64
        }
        var v65 int32 = -1
        _ = v65
    } else {
        v3 -= 1000
    }
    switch 32767 {
    case 9223372036854775807:
        for fi66 := 0; fi66 < 4; fi66++ {
            var v67 uint = 18446744073709551615
            _ = fi66
            _ = v67
        }
    case -503:
        v3 = v3
    }
    for fi68 := 0; fi68 < 8; fi68++ {
        if (15 >= 0) {
            var v69 string = "foo bar"
            _ = v69
        }
        _ = fi68
    }
    v70 := GenS3{F0: uint(v2), F1: 42, F2: ((0.1 != float64(3.14159)) && (1 == 9223372036854775807))}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v70
}
