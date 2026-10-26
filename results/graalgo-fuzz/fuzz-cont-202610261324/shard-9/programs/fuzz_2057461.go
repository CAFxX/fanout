package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int8 }
type GenS2 struct { F0 string; F1 string; F2 int32 }
type GenS3 struct { F0 int16; F1 string }
type GenS4 struct { F0 bool; F1 int32; F2 int32 }
type GenS5 struct { F0 string; F1 int64; F2 uint64 }
type GenS6 struct { F0 int32; F1 bool }
func genFn0(p0 uint16, p1 uint, p2 bool) uint {
    v4 := []int{-1}
    if true {
        v5 := GenS1{F0: 1, F1: -13}
        _ = v5
    } else {
        for fi6 := 0; fi6 < 10; fi6++ {
            switch 1 {
            case 127:
                fmt.Println(p0, p0)
            case 1:
                var v7 uint8 = 42
                _ = v7
            case 100:
                var v8 bool = true
                _ = v8
            default:
                if p2 {
                    p2 = p2
                    if p2 {
                        switch (("a" + "") + "hello") {
                        case "a":
                            if len(v4) > 0 {
                                v4[0] = fi6
                            }
                        case "\n\t":
                            var v9 int8 = -128
                            _ = v9
                        default:
                            if (!p2) {
                                if (128 != 2) {
                                    v4 = append(v4, -790)
                                    p2 = (!false)
                                } else {
                                    for fi10 := 0; fi10 < 1; fi10++ {
                                        switch "\n\t" {
                                        case "\n\t":
                                            if p2 {
                                                switch -400 {
                                                case 32768:
                                                    if len(v4) > 0 {
                                                        v4[0] = (fi10 + fi6)
                                                    }
                                                case 42:
                                                    if len(v4) >= 2 {
                                                        v4 = v4[1:len(v4)]
                                                    }
                                                default:
                                                    if len(v4) >= 2 {
                                                        v4 = v4[1:len(v4)]
                                                    }
                                                }
                                                p1 -= p1
                                                v11 := GenS2{F0: "\n\t", F1: "hello", F2: -1}
                                                _ = v11
                                            } else {
                                                p1 = 475
                                            }
                                        case "a":
                                            if len(v4) > 0 {
                                                v4[0] = (fi6 ^ -874)
                                            }
                                        default:
                                            if p2 {
                                                var v12 int64 = -1
                                                v12 -= (v12 << (((p1 % ((p1) | 1))) & 63))
                                                _ = v12
                                            }
                                        }
                                        if len(v4) >= 2 {
                                            v4 = v4[1:len(v4)]
                                        }
                                        v13 := GenS1{F0: 1, F1: -96}
                                        _ = fi10
                                        _ = v13
                                    }
                                }
                            } else {
                                for fi14 := 0; fi14 < 13; fi14++ {
                                    var v15 bool = false
                                    var v16 uint = p1
                                    v16 -= v16
                                    _ = fi14
                                    _ = v15
                                    _ = v16
                                }
                                var v17 string = ("" + "\n\t")
                                fmt.Println(fi6, p0, fi6)
                                _ = v17
                            }
                        }
                        _ = len(v4)
                        if len(v4) > 0 {
                            v4[0] = fi6
                        }
                    }
                }
            }
            _ = fi6
        }
        _ = len(v4)
        for fi18 := 0; fi18 < 8; fi18++ {
            fmt.Println(p1)
            _ = fi18
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return (p1 / (((p1 >> (((p1 % ((p1) | 1))) & 63))) | 1))
}
func genFn1() int8 {
    var v19 uint32 = 688
    v19 = v19
    for fi20 := 0; fi20 < 15; fi20++ {
        v19 = (v19 * (^(v19)))
        for fi21 := 0; fi21 < 17; fi21++ {
            v19 = 0
            v19 = v19
            switch (fi20 * fi20) {
            case 776:
                var v22 float32 = float32(1e10)
                _ = v22
            case -9223372036854775808:
                fmt.Println(fi21, fi20, fi20)
            case -2147483648:
                switch fi20 {
                case 9223372036854775807:
                    var v23 int32 = 2147483647
                    _ = v23
                default:
                    if ((0.0 <= (3.14159 / 3.14159)) && (1000 >= 1000)) {
                        v19 = v19
                        fmt.Println(fi21, fi20)
                    } else {
                        if true {
                            v24 := GenS2{F0: ("hello" + "\n\t"), F1: ("\n\t" + ("foo bar" + ("\n\t" + ("\n\t" + ("\n\t" + "a"))))), F2: -2147483648}
                            _ = v24
                        } else {
                            if (!(255 >= 2)) {
                                var v25 float32 = 1e10
                                v26 := []int{}
                                _ = v25
                                _ = v26
                            }
                            v27 := GenS3{F0: -1, F1: "a"}
                            _ = v27
                        }
                    }
                }
            }
            v19 = 1
            _ = fi21
        }
        _ = fi20
    }
    _ = v19
    return -128
}

func main() {
    var v1 int8 = 127
    var v2 bool = false
    v3 := []string{"\n\t", "foo bar"}
    fmt.Println(v1)
    var v28 float64 = 0.0
    for fi29 := 0; fi29 < 7; fi29++ {
        for fi30 := 0; fi30 < 10; fi30++ {
            fmt.Println(v28)
            _ = fi30
        }
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        fmt.Println(v2, fi29, v28)
        _ = fi29
    }
    if ("" == "a") {
        for fi31 := 0; fi31 < 20; fi31++ {
            if (v2 && (1.5 != (v28 * 0.0))) {
                for fi32 := 0; fi32 < 12; fi32++ {
                    switch "a" {
                    case "foo bar":
                        var v33 string = "foo bar"
                        _ = v33
                    case "\n\t":
                        if v2 {
                            v34 := GenS3{F0: -32768, F1: ("" + ("a" + (("a" + "") + "\n\t")))}
                            _ = v34
                        }
                    default:
                        v28 += v28
                    }
                    _ = fi32
                }
                v28 += v28
            }
            var v35 uint8 = 1
            fmt.Println(v28, v2)
            _ = fi31
            _ = v35
        }
        var v36 float64 = v28
        _ = v36
    }
    var v37 uint32 = 2
    var v38 uint32 = v37
    if false {
        v39 := GenS1{F0: 2, F1: v1}
        _ = v39
    }
    if true {
        var v40 uint16 = 579
        for fi41 := 0; fi41 < 12; fi41++ {
            v42 := GenS1{F0: 1, F1: 0}
            var v43 int8 = v1
            _ = fi41
            _ = v42
            _ = v43
        }
        for fi44 := 0; fi44 < 13; fi44++ {
            v3 = append(v3, (("" + "a") + "a"))
            _ = fi44
        }
        _ = v40
    } else {
        v45 := GenS4{F0: (float32(float64(1e10)) != ((((-2.25 / ((float32((v28 + 1.5)) - (((((((1e10 / float32(3.14159)) * (((1e10 * ((((0.1 / 1.5) * 1e10) + 1e10) - 1.5)) + (0.0 - (((float32((0.1 + v28)) - 1e10) / 1e10) / 1e10))) + (-2.25 * 1.5))) * float32((v28 + 1.5))) * ((float32(1.5) - 0.1) * float32(v28))) + 0.1) * float32(float64((1e10 * float32(v28))))) - 0.0)) - (1e10 / -2.25))) + -2.25) * 0.1) - 1e10)), F1: 1, F2: 0}
        _ = v45
    }
    v1 += genFn1()
    for fi46 := 0; fi46 < 2; fi46++ {
        var v47 string = ""
        _ = fi46
        _ = v47
    }
    if (65535 == 255) {
        var v48 uint = 42
        v49 := GenS5{F0: ("hello" + "foo bar"), F1: 652, F2: 18446744073709551615}
        var v50 int = 2147483647
        _ = v48
        _ = v49
        _ = v50
    } else {
        v1 = (^(v1))
    }
    v51 := GenS6{F0: 42, F1: v2}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v37
    _ = v38
    _ = v51
}
