package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint8; F1 uint32 }
type GenS3 struct { F0 int64; F1 uint8; F2 uint64 }
type GenS4 struct { F0 int; F1 int32; F2 uint32 }
type GenS5 struct { F0 int32; F1 int64 }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 uint8 }
type GenS8 struct { F0 int16; F1 uint16; F2 string }
func genFn0() int {
    var v7 uint16 = 42
    v7 += v7
    var v8 int64 = 138
    _ = v7
    _ = v8
    return -184
}
func genFn1() uint {
    var v9 bool = true
    if (127 <= 127) {
        fmt.Println(v9)
    } else {
        fmt.Println(v9)
        var v10 string = ("a" + ("\n\t" + "a"))
        _ = v10
    }
    _ = v9
    return 135
}

func main() {
    v1 := make(map[int]uint32)
    var v2 uint16 = 255
    var v3 int8 = 23
    var v4 float32 = 1.5
    v5 := []string{"hello"}
    var v6 uint64 = 2
    var v11 int = 128
    fmt.Println(v2, v2, v6)
    v12 := GenS1{F0: v2}
    var v13 string = ""
    v2 -= 395
    fmt.Println(v6)
    v6 -= (v6 / ((v6) | 1))
    var v14 uint8 = 176
    for fi15 := 0; fi15 < 14; fi15++ {
        if (1 >= 859) {
            v16 := GenS1{F0: v2}
            if true {
                v11 = (v11 ^ fi15)
                if (false && (460 >= 65535)) {
                    v17 := GenS2{F0: v14, F1: 65535}
                    switch (v13 + (v13 + (((v13 + v13) + v13) + "hello"))) {
                    case "hello":
                        for fi18 := 0; fi18 < 16; fi18++ {
                            switch "a" {
                            case "foo bar":
                                v19 := GenS2{F0: (v14 << ((255) & 63)), F1: 458}
                                _ = v19
                            case "":
                                var v20 int64 = 1
                                _ = v20
                            case "a":
                                var v21 int16 = -840
                                _ = v21
                            }
                            _ = fi18
                        }
                    default:
                        if true {
                            if (65535 > 18446744073709551615) {
                                if (4294967295 <= 4294967295) {
                                    var v22 string = ""
                                    fmt.Println(v3, v6)
                                    if (!((!true) && true)) {
                                        v13 = ""
                                        if (v14 < (v14 - 42)) {
                                            v11 = int(v3)
                                            var v23 uint = 4294967295
                                            for fi24 := 0; fi24 < 6; fi24++ {
                                                switch fi24 {
                                                case 4294967295:
                                                    switch genFn0() {
                                                    case -9223372036854775808:
                                                        var v25 uint16 = 32767
                                                        _ = v25
                                                    case 32767:
                                                        v26 := GenS2{F0: v14, F1: 4294967295}
                                                        _ = v26
                                                    case -255:
                                                        v23 += 2
                                                    default:
                                                        var v27 uint64 = v6
                                                        v1[1] = 42
                                                        _ = v27
                                                    }
                                                }
                                                _ = fi24
                                            }
                                            _ = v23
                                        } else {
                                            var v28 uint8 = v14
                                            if (!false) {
                                                var v29 float32 = v4
                                                _ = v29
                                            } else {
                                                var v30 uint32 = 4294967295
                                                _ = v30
                                            }
                                            var v31 uint8 = (v14 + 255)
                                            _ = v28
                                            _ = v31
                                        }
                                        v5 = append(v5, v13)
                                    } else {
                                        var v32 float32 = v4
                                        _ = v32
                                    }
                                    _ = v22
                                }
                                for fi33 := 0; fi33 < 10; fi33++ {
                                    v5 = append(v5, "hello")
                                    _ = fi33
                                }
                                switch "foo bar" {
                                case "\n\t":
                                    for fi34 := 0; fi34 < 15; fi34++ {
                                        var v35 int64 = 42
                                        _ = fi34
                                        _ = v35
                                    }
                                }
                            }
                        }
                    }
                    v3 = v3
                    _ = v17
                } else {
                    for fi36 := 0; fi36 < 8; fi36++ {
                        v37 := GenS3{F0: int64(v14), F1: v14, F2: v6}
                        _ = fi36
                        _ = v37
                    }
                }
                v6 += (v6 >> ((204) & 63))
            }
            var v38 uint8 = (v14 ^ 255)
            _ = v16
            _ = v38
        } else {
            switch genFn0() {
            case 167:
                v39 := GenS1{F0: v2}
                _ = v39
            }
        }
        for fi40 := 0; fi40 < 6; fi40++ {
            _ = len(v1)
            var v41 int64 = 403
            _ = fi40
            _ = v41
        }
        for fi42 := 0; fi42 < 5; fi42++ {
            switch v13 {
            case "hello":
                var v43 int = fi42
                _ = v43
            case "foo bar":
                if (255 == genFn1()) {
                    v2 = (^(v2))
                    fmt.Println(fi42, fi15)
                    v44 := GenS2{F0: (v14 >> ((364) & 63)), F1: 2147483648}
                    _ = v44
                } else {
                    v6 = v6
                    if true {
                        if (v6 == (^(v6))) {
                            _ = len(v1)
                            v3 -= v3
                        } else {
                            var v45 uint = 468
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                            v11 = 426
                            var v46 int8 = (v3 | v3)
                            v46 += (v3 ^ 100)
                            _ = v45
                            _ = v46
                        }
                        fmt.Println(v13)
                        v3 += (^(v3))
                    } else {
                        var v47 int16 = 1
                        _ = v47
                    }
                }
            case "a":
                var v48 bool = true
                var v49 int64 = -804
                delete(v1, 16)
                _ = v48
                _ = v49
            default:
                v4 -= v4
            }
            v50 := GenS4{F0: -557, F1: 2147483647, F2: 2147483647}
            v4 -= 3.14159
            _ = fi42
            _ = v50
        }
        _ = fi15
    }
    v51 := GenS1{F0: (v2 / ((v2) | 1))}
    var v52 uint8 = 1
    switch (v13 + (v13 + v13)) {
    case "":
        for fi53 := 0; fi53 < 19; fi53++ {
            var v54 float32 = (v4 / v4)
            _ = fi53
            _ = v54
        }
    }
    for fi55 := 0; fi55 < 19; fi55++ {
        if (!true) {
            switch fi55 {
            case -1:
                for fi56 := 0; fi56 < 13; fi56++ {
                    var v57 int16 = 1000
                    if ((fi56 & (v11 - (-(fi55)))) != -198) {
                        fmt.Println(v6, v11, v57)
                        var v58 uint8 = 1
                        _ = v58
                    }
                    _ = fi56
                    _ = v57
                }
            }
            switch fi55 {
            case -2147483648:
                for fi59 := 0; fi59 < 12; fi59++ {
                    var v60 float32 = v4
                    v61 := GenS5{F0: int32(v11), F1: 0}
                    _ = fi59
                    _ = v60
                    _ = v61
                }
            case 0:
                v4 += 0.1
            case 42:
                for fi62 := 0; fi62 < 13; fi62++ {
                    var v63 string = "hello"
                    v6 -= (^(v6))
                    _ = len(v1)
                    _ = fi62
                    _ = v63
                }
            default:
                var v64 int16 = int16(v6)
                _ = v64
            }
        } else {
            v65 := GenS6{F0: 1000}
            v13 = ((v13 + v13) + (("" + (v13 + ((((("a" + (v13 + (("\n\t" + v13) + ("foo bar" + v13)))) + (v13 + v13)) + v13) + (v13 + ((v13 + v13) + v13))) + v13))) + "hello"))
            _ = v65
        }
        _ = fi55
    }
    for fi66 := 0; fi66 < 9; fi66++ {
        if true {
            if (v3 >= 42) {
                var v67 bool = false
                _ = v67
            } else {
                switch genFn0() {
                case 668:
                    v68 := GenS7{F0: (^(v14))}
                    _ = v68
                default:
                    if true {
                        if (18446744073709551615 != v6) {
                            var v69 float64 = 3.14159
                            var v70 int16 = -224
                            var v71 int32 = 2147483647
                            _ = v69
                            _ = v70
                            _ = v71
                        } else {
                            switch (((((v13 + v13) + "") + ((v13 + "") + v13)) + (v13 + (v13 + (v13 + v13)))) + "") {
                            case "":
                                var v72 uint16 = 0
                                _ = v72
                            case "foo bar":
                                v3 += 94
                            case "hello":
                                var v73 int32 = 632
                                _ = v73
                            default:
                                delete(v1, (fi66 % (((fi66 + (v11 << ((18446744073709551615) & 63)))) | 1)))
                            }
                        }
                    }
                }
            }
        } else {
            switch ((v13 + v13) + v13) {
            case "foo bar":
                for fi74 := 0; fi74 < 3; fi74++ {
                    var v75 float32 = v4
                    v11 = 2147483647
                    switch v13 {
                    case "a":
                        var v76 int = 699
                        _ = v76
                    case "\n\t":
                        var v77 float64 = 3.14159
                        _ = v77
                    case "hello":
                        v11 += (-(v11))
                    default:
                        v11 = (fi74 - (v11 & fi66))
                    }
                    _ = fi74
                    _ = v75
                }
            case "\n\t":
                for fi78 := 0; fi78 < 19; fi78++ {
                    v11 -= -562
                    _ = fi78
                }
            case "a":
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
            }
            v14 = (v14 & (v52 << ((42) & 63)))
            switch v13 {
            case "":
                fmt.Println(fi66)
            case "\n\t":
                fmt.Println(v13, fi66)
            case "hello":
                v79 := GenS3{F0: 42, F1: v14, F2: v6}
                _ = v79
            }
        }
        _ = fi66
    }
    v80 := GenS8{F0: 1, F1: 108, F2: v13}
    fmt.Println(v13)
    var v81 int8 = (v3 | (v3 & v3))
    fmt.Println(v3, v11, v2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v52:", v52)
    fmt.Println("v6:", v6)
    fmt.Println("v81:", v81)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v51
    _ = v52
    _ = v6
    _ = v80
    _ = v81
}
