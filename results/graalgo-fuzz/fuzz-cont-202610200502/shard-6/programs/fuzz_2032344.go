package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint16; F1 uint16; F2 uint8 }
type GenS3 struct { F0 uint8; F1 uint; F2 int16 }
type GenS4 struct { F0 int64; F1 int8; F2 uint64 }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 bool; F1 int8 }
func genFn0() int64 {
    v5 := []int{2147483647, -9223372036854775808, -974}
    var v6 uint64 = 1
    if false {
        for fi7 := 0; fi7 < 8; fi7++ {
            if (1.5 <= (((((float64(3.14159) - 1.5) / -2.25) / 1e10) + 1e10) - float64(1e10))) {
                v6 -= v6
                if (13 > -126) {
                    switch fi7 {
                    case 564:
                        v6 = (v6 | v6)
                    default:
                        v6 += (v6 ^ (v6 - v6))
                    }
                    if true {
                        var v8 bool = false
                        _ = v8
                    } else {
                        var v9 float64 = 1.5
                        _ = v9
                    }
                } else {
                    v6 = (^(v6))
                    if true {
                        var v10 string = "foo bar"
                        var v11 uint16 = 2
                        v10 = v10
                        _ = v10
                        _ = v11
                    } else {
                        if len(v5) > 0 {
                            v5[0] = (fi7 ^ (fi7 - fi7))
                        }
                        var v12 int32 = 922
                        _ = v12
                    }
                }
            }
            _ = fi7
        }
    } else {
        if ((true || false) && false) {
            for fi13 := 0; fi13 < 13; fi13++ {
                if false {
                    var v14 bool = true
                    fmt.Println(v6, v6, fi13)
                    _ = v14
                }
                v6 -= v6
                fmt.Println(fi13, fi13)
                _ = fi13
            }
            fmt.Println(v6)
        } else {
            var v15 int32 = -586
            for fi16 := 0; fi16 < 14; fi16++ {
                switch fi16 {
                case -349:
                    switch (-(fi16)) {
                    case 32767:
                        if (v6 == v6) {
                            var v17 int64 = 0
                            _ = v17
                        }
                    case 127:
                        var v18 uint8 = 0
                        _ = v18
                    }
                }
                switch (fi16 + fi16) {
                case -156:
                    v15 -= 0
                default:
                    v15 += v15
                }
                _ = fi16
            }
            switch 42 {
            case -1:
                v19 := GenS1{F0: true}
                _ = v19
            case 42:
                var v20 int16 = 42
                _ = v20
            case 1:
                for fi21 := 0; fi21 < 6; fi21++ {
                    _ = len(v5)
                    _ = fi21
                }
            }
            _ = v15
        }
        for fi22 := 0; fi22 < 12; fi22++ {
            var v23 int16 = -1
            _ = fi22
            _ = v23
        }
    }
    for fi24 := 0; fi24 < 3; fi24++ {
        var v25 uint = 295
        if ((1.5 + 0.1) >= (1.5 + -2.25)) {
            var v26 float64 = -2.25
            v26 -= 1e10
            if true {
                switch "" {
                case "foo bar":
                    _ = len(v5)
                default:
                    if (true || (!(true || (!true)))) {
                        switch 0 {
                        case -766:
                            v6 -= v6
                        default:
                            if (100 == -13) {
                                for fi27 := 0; fi27 < 15; fi27++ {
                                    var v28 string = "a"
                                    fmt.Println(v25, v28)
                                    v6 -= 0
                                    _ = fi27
                                    _ = v28
                                }
                            } else {
                                if (128 <= uint8(v6)) {
                                    if false {
                                        for fi29 := 0; fi29 < 8; fi29++ {
                                            _ = len(v5)
                                            var v30 int32 = 2147483647
                                            _ = fi29
                                            _ = v30
                                        }
                                        v31 := GenS1{F0: ((("a" + "foo bar") <= "hello") && true)}
                                        v5 = append(v5, (fi24 - fi24))
                                        _ = v31
                                    }
                                    if (true && true) {
                                        for fi32 := 0; fi32 < 17; fi32++ {
                                            if false {
                                                if (2 >= 255) {
                                                    var v33 int = -9223372036854775808
                                                    v34 := GenS2{F0: 32767, F1: 234, F2: 128}
                                                    _ = v33
                                                    _ = v34
                                                }
                                                v35 := GenS3{F0: 127, F1: v25, F2: -540}
                                                _ = v35
                                            }
                                            _ = fi32
                                        }
                                        if (1e10 < (1.5 / 1.5)) {
                                            var v36 int16 = 32767
                                            _ = v36
                                        } else {
                                            var v37 int = (fi24 / ((fi24) | 1))
                                            _ = v37
                                        }
                                        var v38 int64 = -9223372036854775808
                                        _ = v38
                                    } else {
                                        if ((0.0 + -2.25) != -2.25) {
                                            var v39 int8 = 1
                                            _ = v39
                                        }
                                    }
                                }
                            }
                        }
                    } else {
                        var v40 float64 = (((v26 + v26) - v26) * (float64(1e10) - 3.14159))
                        v25 = (v25 + (^(v25)))
                        _ = v40
                    }
                }
            } else {
                for fi41 := 0; fi41 < 5; fi41++ {
                    v25 += 194
                    _ = fi41
                }
            }
            _ = v26
        } else {
            v6 -= v6
            var v42 string = (("\n\t" + "") + (("hello" + "hello") + "\n\t"))
            _ = v42
        }
        _ = fi24
        _ = v25
    }
    _ = v5
    _ = v6
    return -662
}

func main() {
    var v1 int32 = -65
    v2 := []int{}
    v3 := []uint32{444, 311}
    var v4 uint16 = 42
    _ = len(v3)
    switch "" {
    case "foo bar":
        if true {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            var v43 uint32 = 763
            var v44 int32 = 42
            _ = v43
            _ = v44
        }
    }
    if true {
        var v45 int32 = v1
        _ = v45
    }
    switch 4294967295 {
    case -172:
        v46 := GenS2{F0: (^(v4)), F1: (v4 % ((32767) | 1)), F2: 0}
        _ = v46
    default:
        for fi47 := 0; fi47 < 4; fi47++ {
            v1 = v1
            _ = fi47
        }
    }
    v48 := GenS1{F0: true}
    switch 32768 {
    case 32768:
        v49 := GenS4{F0: 599, F1: -1, F2: 956}
        _ = v49
    case 4294967295:
        v50 := GenS5{F0: 65535}
        _ = v50
    case 508:
        v3 = append(v3, 0)
    }
    switch (("a" + "hello") + "hello") {
    case "foo bar":
        v2 = append(v2, 57)
    case "hello":
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    case "\n\t":
        fmt.Println(v4)
    }
    if false {
        if (18446744073709551615 > 42) {
            var v51 int16 = 32767
            v52 := GenS6{F0: (-1 == -9223372036854775808), F1: 58}
            _ = v51
            _ = v52
        }
    }
    switch "a" {
    case "foo bar":
        v1 -= 42
    case "":
        switch "hello" {
        case "foo bar":
            v53 := GenS4{F0: -9223372036854775808, F1: 42, F2: 0}
            _ = v53
        }
    }
    _ = len(v2)
    for fi54 := 0; fi54 < 9; fi54++ {
        _ = len(v2)
        v55 := GenS6{F0: true, F1: 0}
        _ = fi54
        _ = v55
    }
    v4 -= v4
    if (genFn0() >= genFn0()) {
        var v56 uint64 = 42
        if false {
            var v57 int16 = 42
            fmt.Println(v57, v4, v1)
            _ = v57
        } else {
            var v58 bool = (127 == 42)
            _ = v58
        }
        v1 += v1
        _ = v56
    } else {
        var v59 uint8 = 42
        _ = v59
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v48
}
