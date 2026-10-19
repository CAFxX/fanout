package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 bool; F1 uint16 }
type GenS3 struct { F0 int8; F1 int16; F2 string }
type GenS4 struct { F0 int16; F1 uint64; F2 uint32 }
func genFn0() uint {
    var v6 int = 128
    var v7 int32 = 0
    for fi8 := 0; fi8 < 11; fi8++ {
        if ((42 >= 0) || true) {
            v9 := []int{106, (-(v6)), 32767, fi8}
            v7 = 0
            _ = v9
        } else {
            v10 := []int{}
            _ = v10
        }
        fmt.Println(fi8, v6, v7)
        _ = fi8
    }
    _ = v6
    _ = v7
    return 199
}
func genFn1(p0 int8, p1 int8, p2 uint32) string {
    p0 -= p0
    v11 := []int{-902, 44}
    v11 = append(v11, -2147483648)
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    return "a"
}

func main() {
    var v1 string = ("a" + ("hello" + ("" + ("" + ("\n\t" + "a")))))
    v2 := []int{665, -752}
    v3 := make(map[uint32]bool)
    v4 := make(map[uint32]int)
    var v5 uint8 = 42
    var v12 bool = true
    var v13 int8 = -5
    if ((float64((float32((-2.25 / 0.1)) - (0.1 - 0.0))) == -2.25) || (1e10 <= 0.0)) {
        v5 -= (v5 + 2)
        for fi14 := 0; fi14 < 6; fi14++ {
            var v15 int8 = v13
            v12 = v12
            _ = fi14
            _ = v15
        }
        switch v1 {
        case "\n\t":
            if v12 {
                switch "hello" {
                case "\n\t":
                    v12 = v12
                case "hello":
                    if v12 {
                        var v16 uint8 = (v5 + (v5 % (((v5 >> ((42) & 63))) | 1)))
                        if (!(v13 <= (v13 >> ((4294967295) & 63)))) {
                            v5 = v16
                            if v12 {
                                var v17 int8 = 0
                                if v12 {
                                    v1 = v1
                                }
                                if v12 {
                                    var v18 int64 = 1
                                    _ = v18
                                }
                                _ = v17
                            }
                        }
                        v19 := GenS1{F0: 65535}
                        _ = v16
                        _ = v19
                    } else {
                        v20 := GenS1{F0: 147}
                        v5 -= (v5 >> ((genFn0()) & 63))
                        v12 = v12
                        _ = v20
                    }
                }
                var v21 int = 128
                _ = len(v3)
                _ = v21
            } else {
                var v22 int16 = -32768
                fmt.Println(v13, v12, v1)
                for fi23 := 0; fi23 < 8; fi23++ {
                    v22 = (^(v22))
                    for fi24 := 0; fi24 < 7; fi24++ {
                        var v25 int64 = 1
                        var v26 float64 = float64(1e10)
                        if (2 != 2147483647) {
                            var v27 uint64 = 2
                            v13 -= -1
                            _ = v27
                        } else {
                            v5 = v5
                            if ((v13 << ((18446744073709551615) & 63)) == v13) {
                                v28 := GenS2{F0: v12, F1: 1}
                                _ = v28
                            }
                        }
                        _ = fi24
                        _ = v25
                        _ = v26
                    }
                    _ = fi23
                }
                _ = v22
            }
        }
    } else {
        var v29 uint8 = (v5 / ((v5) | 1))
        v30 := v3[4294967295]
        for fi31 := 0; fi31 < 10; fi31++ {
            fmt.Println(v5, fi31)
            _ = fi31
        }
        _ = v29
        _ = v30
    }
    delete(v4, 959)
    v5 -= v5
    if true {
        if (4294967295 == 2147483648) {
            v2 = append(v2, 32767)
            switch "foo bar" {
            case "a":
                switch "hello" {
                case "hello":
                    for fi32 := 0; fi32 < 14; fi32++ {
                        switch fi32 {
                        case 4294967295:
                            if (fi32 > (fi32 ^ int(v5))) {
                                var v33 uint = 65535
                                _ = v33
                            } else {
                                for fi34 := 0; fi34 < 6; fi34++ {
                                    var v35 uint8 = v5
                                    switch v1 {
                                    case "foo bar":
                                        _ = len(v4)
                                    }
                                    _ = fi34
                                    _ = v35
                                }
                                for fi36 := 0; fi36 < 20; fi36++ {
                                    fmt.Println(fi32, fi32)
                                    v12 = v12
                                    _ = fi36
                                }
                            }
                        case -1:
                            var v37 float64 = ((((3.14159 + ((1e10 - 0.1) - (-2.25 / 1.0))) - 0.0) + 1.5) / float64(0.1))
                            _ = v37
                        case 0:
                            fmt.Println(v13, v12, v5)
                        default:
                            v38 := GenS1{F0: 1}
                            _ = v38
                        }
                        var v39 float32 = 0.1
                        _ = fi32
                        _ = v39
                    }
                case "foo bar":
                    var v40 int16 = 32767
                    _ = v40
                case "":
                    if (((610 >= 42) || (42 != 313)) && v12) {
                        v1 = (v1 + v1)
                        var v41 int8 = (-(v13))
                        _ = v41
                    }
                default:
                    v1 = v1
                }
            }
            for fi42 := 0; fi42 < 7; fi42++ {
                var v43 string = "foo bar"
                _ = fi42
                _ = v43
            }
        } else {
            v44 := GenS1{F0: 1}
            var v45 int32 = 2147483647
            _ = v44
            _ = v45
        }
        v46 := GenS2{F0: v12, F1: 255}
        for fi47 := 0; fi47 < 6; fi47++ {
            if len(v2) > 0 {
                v2[0] = (fi47 * fi47)
            }
            switch ("\n\t" + "\n\t") {
            case "hello":
                var v48 uint8 = (v5 >> ((genFn0()) & 63))
                _ = v48
            }
            _ = fi47
        }
        _ = v46
    } else {
        if (0.0 > (-2.25 / 0.1)) {
            for fi49 := 0; fi49 < 12; fi49++ {
                fmt.Println(v5, v1, v5)
                v50 := GenS3{F0: (^(v13)), F1: 958, F2: v1}
                _ = fi49
                _ = v50
            }
            v51 := GenS3{F0: 127, F1: -267, F2: v1}
            _ = v51
        } else {
            var v52 float32 = 0.0
            var v53 bool = ((v1 + "") != ("" + v1))
            fmt.Println(v53)
            _ = v52
            _ = v53
        }
    }
    v54 := GenS3{F0: -46, F1: -911, F2: v1}
    v55 := GenS3{F0: (v13 ^ v13), F1: 42, F2: v1}
    var v56 int = -1
    var v57 int16 = 1000
    switch v1 {
    case "":
        var v58 int = (^(v56))
        _ = v58
    default:
        var v59 string = v1
        _ = v59
    }
    var v60 bool = (v12 || (false && true))
    var v61 uint = 4294967295
    v2 = append(v2, (v56 + -353))
    if v60 {
        for fi62 := 0; fi62 < 19; fi62++ {
            var v63 int32 = -380
            v64 := GenS3{F0: v13, F1: (v57 * 42), F2: (v1 + (((((v1 + "") + ("hello" + "a")) + "foo bar") + "foo bar") + ((v1 + "") + "a")))}
            var v65 string = v1
            _ = fi62
            _ = v63
            _ = v64
            _ = v65
        }
        if (-9223372036854775808 < 649) {
            if (true && (820 >= 550)) {
                v66 := GenS4{F0: -32768, F1: 2, F2: 42}
                for fi67 := 0; fi67 < 19; fi67++ {
                    v61 = v61
                    switch (-(v56)) {
                    case -2147483648:
                        var v68 int64 = 9223372036854775807
                        _ = v68
                    }
                    _ = fi67
                }
                v5 += v5
                _ = v66
            }
            switch (v56 & (^(v56))) {
            case 128:
                v61 -= v61
            case 2147483647:
                v1 = (("foo bar" + "") + (((("a" + "foo bar") + v1) + (v1 + "a")) + v1))
            default:
                fmt.Println(v5, v60, v12)
            }
        } else {
            v4[2147483647] = (v56 >> ((947) & 63))
            for fi69 := 0; fi69 < 15; fi69++ {
                v5 = (v5 & (v5 << ((v61) & 63)))
                var v70 bool = v60
                _ = fi69
                _ = v70
            }
            var v71 uint32 = 2147483647
            _ = v71
        }
    }
    v61 -= (^(v61))
    v4[0] = (v56 % ((v56) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v56:", v56)
    fmt.Println("v57:", v57)
    fmt.Println("v60:", v60)
    fmt.Println("v61:", v61)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v54
    _ = v55
    _ = v56
    _ = v57
    _ = v60
    _ = v61
}
