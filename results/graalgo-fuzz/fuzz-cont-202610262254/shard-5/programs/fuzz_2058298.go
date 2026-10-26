package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int64 }
type GenS2 struct { F0 bool; F1 uint32; F2 uint64 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 uint64; F1 int16; F2 int32 }
func genFn0() int16 {
    for fi7 := 0; fi7 < 13; fi7++ {
        v8 := GenS1{F0: 65535, F1: 0}
        var v9 uint16 = 0
        v9 -= (v9 % (((v9 * v9)) | 1))
        _ = fi7
        _ = v8
        _ = v9
    }
    for fi10 := 0; fi10 < 17; fi10++ {
        var v11 int32 = -2147483648
        var v12 uint16 = 0
        _ = fi10
        _ = v11
        _ = v12
    }
    var v13 uint = 2
    var v14 uint16 = 935
    v15 := GenS2{F0: (!true), F1: 2147483648, F2: 0}
    _ = v13
    _ = v14
    _ = v15
    return 67
}

func main() {
    var v1 int = -1
    v2 := make(map[string]uint32)
    v3 := make(map[uint32]int)
    var v4 string = ("foo bar" + "foo bar")
    v5 := []string{(v4 + "hello"), v4, v4}
    var v6 uint32 = 2147483648
    v6 = 2147483647
    v1 = v1
    if true {
        var v16 int = v1
        for fi17 := 0; fi17 < 7; fi17++ {
            var v18 int16 = 32767
            var v19 uint32 = v6
            _ = fi17
            _ = v18
            _ = v19
        }
        _ = v16
    } else {
        var v20 int16 = genFn0()
        v4 = v4
        var v21 int8 = 1
        _ = v20
        _ = v21
    }
    switch "hello" {
    case "foo bar":
        for fi22 := 0; fi22 < 4; fi22++ {
            for fi23 := 0; fi23 < 2; fi23++ {
                _ = len(v5)
                _ = fi23
            }
            v24 := GenS3{F0: v6}
            _ = fi22
            _ = v24
        }
    case "":
        v1 = -223
    }
    switch ("\n\t" + (v4 + v4)) {
    case "":
        fmt.Println(v6, v6)
    default:
        switch (v1 << ((0) & 63)) {
        case 1:
            v25 := GenS3{F0: v6}
            _ = v25
        case 111:
            v26 := GenS3{F0: (v6 ^ 949)}
            _ = v26
        default:
            fmt.Println(v1)
        }
    }
    v1 = 127
    var v27 int = (v1 ^ v1)
    _ = len(v5)
    var v28 uint8 = 128
    if (!false) {
        v28 -= v28
        for fi29 := 0; fi29 < 9; fi29++ {
            switch (v1 % ((4294967295) | 1)) {
            case 9223372036854775807:
                if true {
                    fmt.Println(v4, v28, v6)
                } else {
                    v28 = v28
                }
            case 1:
                var v30 uint32 = (v6 / (((v6 | v6)) | 1))
                _ = v30
            }
            _ = fi29
        }
        v4 = ""
    } else {
        v28 += 1
    }
    switch v27 {
    case -756:
        if (!true) {
            v27 = -262
            _ = len(v5)
        } else {
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            var v31 uint = 58
            var v32 int16 = -658
            _ = v31
            _ = v32
        }
    case 128:
        var v33 int16 = -1
        _ = v33
    default:
        v34 := GenS4{F0: v28}
        _ = v34
    }
    if (42 >= 2147483647) {
        switch v1 {
        case 827:
            if ((-1 < -1) || false) {
                if (-1 > 0) {
                    var v35 bool = (false && (!false))
                    for fi36 := 0; fi36 < 17; fi36++ {
                        for fi37 := 0; fi37 < 16; fi37++ {
                            fmt.Println(v27, v1)
                            _ = fi37
                        }
                        v28 -= v28
                        switch (v4 + v4) {
                        case "\n\t":
                            v38 := GenS4{F0: (v28 + v28)}
                            _ = v38
                        default:
                            v4 = (v4 + (("hello" + v4) + ("" + "hello")))
                        }
                        _ = fi36
                    }
                    _ = v35
                }
            } else {
                if true {
                    var v39 int = 9223372036854775807
                    _ = v39
                } else {
                    v4 = "hello"
                    for fi40 := 0; fi40 < 9; fi40++ {
                        var v41 int = fi40
                        for fi42 := 0; fi42 < 11; fi42++ {
                            v4 = v4
                            if (!true) {
                                fmt.Println(fi40, v27, v41)
                                v1 += (v1 << ((2) & 63))
                            }
                            v27 -= (fi42 + (fi42 ^ (-(v27))))
                            _ = fi42
                        }
                        var v43 int16 = 32767
                        _ = fi40
                        _ = v41
                        _ = v43
                    }
                    switch (-(v1)) {
                    case 2147483647:
                        v5 = append(v5, "a")
                    }
                }
                switch "a" {
                case "foo bar":
                    var v44 float32 = (-2.25 - (1e10 + (3.14159 + (1.5 * 1e10))))
                    _ = v44
                case "":
                    var v45 int8 = -11
                    _ = v45
                default:
                    v46 := GenS1{F0: 65535, F1: -29}
                    _ = v46
                }
            }
        case -1:
            var v47 int = v27
            _ = v47
        case -2147483648:
            var v48 uint16 = 319
            _ = v48
        default:
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
        }
        v4 = v4
        if true {
            if false {
                fmt.Println(v28, v27, v6)
                fmt.Println(v1)
            } else {
                for fi49 := 0; fi49 < 12; fi49++ {
                    for fi50 := 0; fi50 < 1; fi50++ {
                        if true {
                            if (!((-9223372036854775808 < 1) && false)) {
                                v51 := GenS1{F0: 42, F1: 42}
                                if (((((v4 + "\n\t") + "") + "\n\t") >= (v4 + "hello")) && ((false || false) || (true || (42 < 42)))) {
                                    v6 += v6
                                    var v52 int32 = 459
                                    if (777 != 18446744073709551615) {
                                        v27 = v27
                                    }
                                    _ = v52
                                } else {
                                    v6 = v6
                                    v27 += -2147483648
                                }
                                _ = v51
                            }
                            v4 = v4
                        }
                        var v53 int16 = 1000
                        if (0.1 <= -2.25) {
                            if (!(!true)) {
                                var v54 int64 = -1
                                var v55 int = v27
                                _ = v54
                                _ = v55
                            } else {
                                v56 := GenS5{F0: 431, F1: genFn0(), F2: 444}
                                var v57 uint8 = 0
                                _ = v56
                                _ = v57
                            }
                        }
                        _ = fi50
                        _ = v53
                    }
                    _ = fi49
                }
                if false {
                    v3[(v6 * 2)] = (v27 << ((65535) & 63))
                } else {
                    switch 1 {
                    case 128:
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                    case -1:
                        var v58 int8 = 127
                        _ = v58
                    default:
                        fmt.Println(v27, v1)
                    }
                    var v59 uint = 255
                    var v60 uint64 = uint64(v6)
                    _ = v59
                    _ = v60
                }
                v1 = v1
            }
        }
    }
    fmt.Println(v27)
    var v61 uint8 = 62
    var v62 int64 = 483
    if (0.0 != float32(float64(3.14159))) {
        if (1000 == 0) {
            if (false && true) {
                var v63 int = (v1 ^ 845)
                v6 = v6
                _ = v63
            }
            v3[v6] = v27
            var v64 float64 = float64(1e10)
            _ = v64
        }
        v1 = v27
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v61:", v61)
    fmt.Println("v62:", v62)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v61
    _ = v62
}
