package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int8 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 string; F1 uint16 }
type GenS4 struct { F0 bool; F1 int16 }
type GenS5 struct { F0 int32; F1 bool; F2 int }
type GenS6 struct { F0 int16; F1 uint16 }
type GenS7 struct { F0 uint; F1 uint; F2 int16 }
type GenS8 struct { F0 uint32; F1 uint16; F2 int }
type GenS9 struct { F0 string; F1 uint8 }
type GenS10 struct { F0 bool; F1 uint64 }
func genFn0(p0 int16, p1 int8) int64 {
    p1 -= 4
    var v7 int = 128
    if (-370 != int32(p0)) {
        v8 := GenS1{F0: 463, F1: (p1 & int8(v7))}
        _ = v8
    }
    p1 += p1
    var v9 uint8 = 1
    _ = p0
    _ = p1
    _ = v7
    _ = v9
    return -1
}
func genFn1(p0 bool) uint32 {
    fmt.Println(p0)
    p0 = p0
    var v10 int = 127
    _ = p0
    _ = v10
    return 0
}

func main() {
    v1 := []int{-1, -9223372036854775808, 128, -318}
    var v2 bool = (42 <= -69)
    var v3 bool = v2
    var v4 uint32 = 4294967295
    var v5 uint32 = v4
    var v6 int32 = 654
    for fi11 := 0; fi11 < 13; fi11++ {
        if ((!(v2 && (-1 >= -56))) || v3) {
            v1 = append(v1, fi11)
            v2 = v2
        }
        var v12 string = ("hello" + "hello")
        fmt.Println(fi11, v12)
        _ = fi11
        _ = v12
    }
    var v13 int = 127
    if false {
        if true {
            v3 = (!(65535 == 238))
        }
        v2 = (718 != 32767)
        var v14 uint32 = 824
        _ = v14
    } else {
        v5 = (v5 | 0)
        if len(v1) > 0 {
            v1[0] = v13
        }
        for fi15 := 0; fi15 < 9; fi15++ {
            if v2 {
                var v16 int = v13
                v2 = (65535 >= 416)
                for fi17 := 0; fi17 < 8; fi17++ {
                    v18 := GenS2{F0: "\n\t"}
                    switch int(v6) {
                    case 4294967295:
                        switch v13 {
                        case -9223372036854775808:
                            v19 := GenS1{F0: 445, F1: -128}
                            _ = v19
                        case 410:
                            if (v13 > 1) {
                                var v20 float32 = 3.14159
                                _ = v20
                            } else {
                                fmt.Println(v3, v13)
                                var v21 int16 = -32768
                                if true {
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                    v22 := GenS1{F0: 332, F1: 1}
                                    _ = v22
                                }
                                _ = v21
                            }
                            v16 -= v16
                        }
                    }
                    v23 := GenS2{F0: "a"}
                    _ = fi17
                    _ = v18
                    _ = v23
                }
                _ = v16
            } else {
                v13 -= (v13 & 18)
                if v3 {
                    var v24 float64 = 3.14159
                    switch "foo bar" {
                    case "":
                        v2 = ((("foo bar" + "hello") + "a") <= "")
                    case "a":
                        _ = len(v1)
                    case "\n\t":
                        if (v3 || true) {
                            if false {
                                _ = len(v1)
                                for fi25 := 0; fi25 < 8; fi25++ {
                                    fmt.Println(v6, v2, v4)
                                    var v26 uint32 = v4
                                    v27 := GenS1{F0: 315, F1: -128}
                                    if true {
                                        switch 47 {
                                        case -46:
                                            var v28 float32 = (((1.5 - ((-2.25 / float32(((3.14159 * 0.0) + ((0.0 / v24) * v24)))) - -2.25)) - (float32(float64((1.5 * 0.0))) * 1.5)) + 0.1)
                                            _ = v28
                                        case 127:
                                            v13 = 4294967295
                                            v29 := GenS2{F0: "a"}
                                            v24 -= v24
                                            _ = v29
                                        case 1:
                                            var v30 uint16 = 1
                                            _ = v30
                                        }
                                        v31 := GenS3{F0: ("" + ""), F1: 157}
                                        _ = v31
                                    } else {
                                        if (v2 && (32767 >= -730)) {
                                            if len(v1) >= 2 {
                                                v1 = v1[1:len(v1)]
                                            }
                                            v24 += (v24 - (v24 - v24))
                                            if len(v1) >= 2 {
                                                v1 = v1[1:len(v1)]
                                            }
                                            v2 = v2
                                            if (v2 && (42 >= 65535)) {
                                                v32 := GenS4{F0: true, F1: -1}
                                                var v33 uint32 = (v5 & 42)
                                                fmt.Println(v2)
                                                _ = v32
                                                _ = v33
                                            }
                                            v13 = fi15
                                            v4 -= (v4 / (((v4 % (((v5 << ((18446744073709551615) & 63))) | 1))) | 1))
                                        } else {
                                            v3 = (!v3)
                                        }
                                        switch fi25 {
                                        case 759:
                                            fmt.Println(fi15, v6, v3)
                                        case -9223372036854775808:
                                            var v34 int16 = 1000
                                            _ = v34
                                        }
                                        var v35 uint = 4294967295
                                        var v36 bool = v2
                                        _ = v35
                                        _ = v36
                                    }
                                    _ = fi25
                                    _ = v26
                                    _ = v27
                                }
                                var v37 uint = 0
                                _ = v37
                            } else {
                                v38 := GenS5{F0: v6, F1: (42 <= v4), F2: v13}
                                if (333 >= 0) {
                                    var v39 string = "a"
                                    var v40 int16 = 844
                                    fmt.Println(v39)
                                    _ = v39
                                    _ = v40
                                } else {
                                    v13 += v13
                                }
                                v13 = v13
                                _ = v38
                            }
                            _ = len(v1)
                            v2 = v2
                        } else {
                            fmt.Println(v5)
                        }
                    }
                    v41 := GenS4{F0: v3, F1: -32768}
                    _ = v24
                    _ = v41
                } else {
                    if (v2 && v2) {
                        for fi42 := 0; fi42 < 7; fi42++ {
                            v43 := GenS6{F0: -1, F1: 709}
                            switch "" {
                            case "hello":
                                v44 := GenS7{F0: 1, F1: 222, F2: 42}
                                _ = v44
                            default:
                                if false {
                                    if v2 {
                                        v4 -= v4
                                        var v45 uint32 = (v4 - v5)
                                        _ = v45
                                    } else {
                                        var v46 uint64 = 395
                                        v13 += fi42
                                        v13 -= fi15
                                        _ = v46
                                    }
                                    v6 = 2147483647
                                    v3 = (-2.25 != (1.5 + (-2.25 / 1.0)))
                                }
                            }
                            _ = fi42
                            _ = v43
                        }
                        if (4294967295 <= 958) {
                            var v47 uint = 1
                            _ = v47
                        } else {
                            for fi48 := 0; fi48 < 10; fi48++ {
                                var v49 int = fi15
                                v1 = append(v1, 2147483647)
                                v50 := GenS8{F0: 0, F1: 255, F2: fi48}
                                _ = fi48
                                _ = v49
                                _ = v50
                            }
                            v2 = ((true && v3) || true)
                        }
                    } else {
                        v13 = fi15
                        v13 = 128
                        v3 = v3
                    }
                    v4 = (v5 | v5)
                    v4 -= (^(v5))
                }
                v51 := GenS7{F0: 4294967295, F1: 4294967295, F2: 0}
                _ = v51
            }
            v4 -= (^(v4))
            _ = fi15
        }
    }
    if ((0.1 / -2.25) < ((0.1 * 0.1) - 0.1)) {
        v5 = 898
        v13 = 266
    } else {
        switch "\n\t" {
        case "a":
            switch (v13 >> ((7) & 63)) {
            case 32768:
                var v52 int64 = -1
                _ = v52
            case 127:
                fmt.Println(v5)
            default:
                if v3 {
                    var v53 uint8 = 242
                    _ = v53
                } else {
                    v5 -= (v5 << ((717) & 63))
                    var v54 string = ("" + "a")
                    var v55 float64 = 0.0
                    _ = v54
                    _ = v55
                }
            }
        }
        v6 -= (v6 | v6)
        if (1 >= -1) {
            v5 = (v5 / ((2) | 1))
            v4 = v4
        }
    }
    if ((v13 >> ((472) & 63)) >= v13) {
        fmt.Println(v5, v3)
        for fi56 := 0; fi56 < 3; fi56++ {
            if v3 {
                for fi57 := 0; fi57 < 10; fi57++ {
                    v58 := GenS9{F0: ("a" + "a"), F1: 231}
                    _ = fi57
                    _ = v58
                }
            } else {
                var v59 int64 = 42
                _ = v59
            }
            v60 := GenS9{F0: "a", F1: 1}
            if ((v5 ^ v4) <= (v5 + 65535)) {
                fmt.Println(v6, v3, fi56)
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            }
            _ = fi56
            _ = v60
        }
        for fi61 := 0; fi61 < 8; fi61++ {
            fmt.Println(v4, fi61, v5)
            _ = fi61
        }
    } else {
        var v62 int = v13
        _ = v62
    }
    var v63 uint8 = 120
    fmt.Println(v4, v2)
    for fi64 := 0; fi64 < 12; fi64++ {
        v6 = 438
        _ = fi64
    }
    switch 9223372036854775807 {
    case 1:
        v6 = v6
    }
    v65 := GenS10{F0: ("" >= (("foo bar" + "") + "\n\t")), F1: 42}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v63:", v63)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v63
    _ = v65
}
