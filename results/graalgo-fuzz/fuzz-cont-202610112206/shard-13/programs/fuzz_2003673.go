package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int; F1 string }
type GenS3 struct { F0 string; F1 int; F2 int16 }
type GenS4 struct { F0 bool; F1 string }
type GenS5 struct { F0 int32; F1 int; F2 uint16 }
type GenS6 struct { F0 uint; F1 int64; F2 uint64 }
type GenS7 struct { F0 uint64 }
type GenS8 struct { F0 int; F1 string; F2 int }
func genFn0() uint16 {
    var v7 int32 = 142
    v8 := GenS1{F0: 0}
    _ = v7
    _ = v8
    return 32767
}
func genFn1() uint8 {
    var v9 float64 = 1.5
    if true {
        var v10 uint64 = 254
        v11 := []int{-9223372036854775808}
        _ = v10
        _ = v11
    }
    var v12 uint8 = 255
    switch "\n\t" {
    case "":
        var v13 float64 = v9
        _ = v13
    case "foo bar":
        var v14 float64 = v9
        _ = v14
    default:
        for fi15 := 0; fi15 < 8; fi15++ {
            v12 = 2
            _ = fi15
        }
    }
    fmt.Println(v12, v12)
    _ = v12
    _ = v9
    return v12
}

func main() {
    v1 := []string{"\n\t", "a", ("a" + "")}
    v2 := []int{128, -2147483648, 0}
    var v3 uint16 = 492
    var v4 uint64 = 4294967295
    var v5 uint32 = uint32(v4)
    v6 := []int{281, 128, 2147483647}
    v4 = v4
    for fi16 := 0; fi16 < 17; fi16++ {
        if len(v1) > 0 {
            v1[0] = ""
        }
        v4 += (v4 << ((394) & 63))
        v17 := GenS1{F0: 903}
        _ = fi16
        _ = v17
    }
    var v18 int64 = -494
    if (100 == 42) {
        fmt.Println(v5)
    }
    if true {
        var v19 string = "foo bar"
        v3 = genFn0()
        v4 -= (v4 % (((v4 % ((4294967295) | 1))) | 1))
        _ = v19
    }
    v18 -= (v18 >> ((2) & 63))
    if len(v6) > 0 {
        v6[0] = int(v3)
    }
    var v20 uint = 42
    for fi21 := 0; fi21 < 16; fi21++ {
        for fi22 := 0; fi22 < 1; fi22++ {
            v23 := GenS2{F0: (fi21 / (((fi21 * (fi21 << (((^(v20))) & 63)))) | 1)), F1: ("a" + "hello")}
            var v24 int = 42
            _ = fi22
            _ = v23
            _ = v24
        }
        _ = fi21
    }
    v25 := GenS1{F0: (v18 ^ -1)}
    for fi26 := 0; fi26 < 5; fi26++ {
        fmt.Println(v3, v5)
        switch 128 {
        case 620:
            if true {
                v3 -= v3
            } else {
                if len(v2) > 0 {
                    v2[0] = (fi26 % (((fi26 * fi26)) | 1))
                }
                v27 := GenS3{F0: "hello", F1: 127, F2: 42}
                _ = v27
            }
        case 32768:
            v28 := GenS4{F0: (!false), F1: ("a" + "")}
            _ = v28
        default:
            for fi29 := 0; fi29 < 18; fi29++ {
                v1 = append(v1, "")
                v4 = uint64(v5)
                _ = fi29
            }
        }
        if ((v4 <= 79) || false) {
            if (!false) {
                switch ("\n\t" + "a") {
                case "foo bar":
                    var v30 int = fi26
                    _ = v30
                case "":
                    v31 := GenS3{F0: "hello", F1: int(v4), F2: -59}
                    _ = v31
                case "hello":
                    fmt.Println(v20, v18)
                default:
                    if len(v2) > 0 {
                        v2[0] = fi26
                    }
                }
                var v32 int = (fi26 ^ 32768)
                if true {
                    v3 = (v3 | genFn0())
                    switch v32 {
                    case -903:
                        for fi33 := 0; fi33 < 10; fi33++ {
                            var v34 float32 = -2.25
                            var v35 int32 = -1
                            v32 = 127
                            v32 -= fi26
                            _ = fi33
                            _ = v34
                            _ = v35
                        }
                    case 32767:
                        _ = len(v2)
                    default:
                        for fi36 := 0; fi36 < 7; fi36++ {
                            if true {
                                switch v32 {
                                case 849:
                                    v4 += 2
                                case -1:
                                    switch ("a" + "foo bar") {
                                    case "a":
                                        if len(v6) >= 2 {
                                            v6 = v6[1:len(v6)]
                                        }
                                    default:
                                        fmt.Println(v18, v32)
                                        v20 = v20
                                    }
                                default:
                                    var v37 bool = (875 > int32(v4))
                                    _ = v37
                                }
                            }
                            fmt.Println(v32, v5, fi36)
                            v38 := GenS5{F0: 0, F1: fi26, F2: v3}
                            _ = fi36
                            _ = v38
                        }
                    }
                } else {
                    var v39 int64 = -1
                    v40 := GenS6{F0: (v20 & (^(v20))), F1: v39, F2: uint64(v20)}
                    _ = v39
                    _ = v40
                }
                _ = v32
            } else {
                v41 := GenS2{F0: (fi26 * int(v18)), F1: "hello"}
                _ = v41
            }
        } else {
            if true {
                if ((-2147483648 >= int32(v4)) || true) {
                    v42 := GenS7{F0: 42}
                    for fi43 := 0; fi43 < 19; fi43++ {
                        if true {
                            var v44 float64 = ((-2.25 * (1e10 / ((1.5 - float64(float32(0.0))) + float64(-2.25)))) * ((((((1e10 - float64(((1.5 - 1.5) / (0.0 + 1.5)))) - (-2.25 + 3.14159)) - 0.1) / -2.25) - (1e10 + 1e10)) - (float64((-2.25 * -2.25)) + 0.0)))
                            v44 -= 0.1
                            switch fi43 {
                            case 0:
                                var v45 int32 = int32(v4)
                                _ = v45
                            }
                            _ = v44
                        } else {
                            var v46 int8 = -128
                            v4 -= 888
                            v46 += v46
                            _ = v46
                        }
                        _ = len(v6)
                        _ = fi43
                    }
                    _ = v42
                }
                var v47 bool = (((-2.25 / (float64(1.5) - -2.25)) >= 1.5) && false)
                _ = v47
            } else {
                for fi48 := 0; fi48 < 5; fi48++ {
                    var v49 int16 = -32768
                    v2 = append(v2, fi48)
                    if (1e10 == 3.14159) {
                        var v50 int = (-(fi26))
                        fmt.Println(v3, v50, v49)
                        _ = v50
                    }
                    _ = fi48
                    _ = v49
                }
            }
            var v51 float32 = 0.1
            if len(v2) > 0 {
                v2[0] = int(v4)
            }
            _ = v51
        }
        _ = fi26
    }
    var v52 uint32 = (v5 * 1)
    switch ("\n\t" + (("hello" + "foo bar") + "")) {
    case "":
        v53 := GenS8{F0: 1, F1: "", F2: 128}
        _ = v53
    }
    if false {
        switch 9223372036854775807 {
        case -1:
            v52 = (v52 / (((v5 | (v5 % ((v52) | 1)))) | 1))
        case -2147483648:
            if len(v6) > 0 {
                v6[0] = -544
            }
        default:
            v20 += v20
        }
    }
    for fi54 := 0; fi54 < 6; fi54++ {
        fmt.Println(v4, v52)
        v55 := GenS4{F0: (genFn0() == (^(v3))), F1: "a"}
        var v56 uint32 = 4294967295
        _ = fi54
        _ = v55
        _ = v56
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v52:", v52)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v18
    _ = v2
    _ = v20
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v52
    _ = v6
}
