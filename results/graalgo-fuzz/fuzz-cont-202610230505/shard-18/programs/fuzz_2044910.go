package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool }
type GenS2 struct { F0 int; F1 uint32; F2 string }
type GenS3 struct { F0 uint64; F1 int }
type GenS4 struct { F0 uint64; F1 bool; F2 int8 }
type GenS5 struct { F0 int8; F1 int16 }
type GenS6 struct { F0 uint16 }
func genFn0(p0 uint32) int {
    v6 := []int{}
    if true {
        var v7 bool = true
        _ = v7
    }
    var v8 uint32 = p0
    for fi9 := 0; fi9 < 17; fi9++ {
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        for fi10 := 0; fi10 < 12; fi10++ {
            var v11 uint16 = 42
            _ = fi10
            _ = v11
        }
        var v12 int16 = -1
        _ = fi9
        _ = v12
    }
    _ = p0
    _ = v6
    _ = v8
    return 32768
}
func genFn1() int {
    v13 := []int{1, -768}
    if len(v13) > 0 {
        v13[0] = -417
    }
    v14 := GenS1{F0: 127, F1: ((((32767 <= -1) && (1 >= 4294967295)) && (100 < 1)) && (-763 > 42))}
    _ = v13
    _ = v14
    return 793
}

func main() {
    var v1 uint64 = 101
    var v2 uint8 = 255
    var v3 int16 = -84
    v4 := []string{("" + "a"), ("foo bar" + "hello"), ""}
    v5 := make(map[int]int)
    v15 := GenS1{F0: v2, F1: true}
    for fi16 := 0; fi16 < 2; fi16++ {
        v2 += v2
        for fi17 := 0; fi17 < 17; fi17++ {
            switch (("\n\t" + ("\n\t" + "hello")) + "a") {
            case "hello":
                v2 += 1
            case "":
                if true {
                    v5[-1] = genFn0(1)
                    v18 := GenS1{F0: v2, F1: (42 < 63)}
                    fmt.Println(v1)
                    _ = v18
                } else {
                    var v19 uint32 = 2
                    _ = v19
                }
            }
            _ = fi17
        }
        _ = fi16
    }
    for fi20 := 0; fi20 < 9; fi20++ {
        v3 = v3
        _ = fi20
    }
    v21 := GenS1{F0: v2, F1: (v2 < (v2 % (((v2 ^ 1)) | 1)))}
    v2 = v2
    if (float32((float64((-2.25 / ((0.0 - ((0.1 / (-2.25 * (float32((-2.25 / 0.1)) * 1e10))) - 1e10)) - float32(float64((((float32(float64(1e10)) + 0.0) - ((1e10 * -2.25) + (3.14159 + -2.25))) + (-2.25 / (1.5 + (-2.25 + ((-2.25 * 1.5) * 3.14159)))))))))) / 1.0)) > 1.5) {
        for fi22 := 0; fi22 < 1; fi22++ {
            var v23 float32 = (1.5 * 0.1)
            _ = fi22
            _ = v23
        }
        if true {
            v24 := GenS2{F0: 2147483647, F1: 741, F2: "foo bar"}
            var v25 int64 = 1
            _ = v24
            _ = v25
        } else {
            var v26 string = (("\n\t" + "\n\t") + ("a" + ""))
            v1 = 1
            _ = v26
        }
        var v27 int64 = -511
        _ = v27
    } else {
        if false {
            for fi28 := 0; fi28 < 2; fi28++ {
                fmt.Println(v3, fi28)
                switch fi28 {
                case 255:
                    var v29 uint8 = v2
                    _ = v29
                case -2147483648:
                    if false {
                        v2 = 2
                        v2 += v2
                    } else {
                        v30 := GenS2{F0: fi28, F1: 268, F2: "\n\t"}
                        v2 -= (^(v2))
                        switch "hello" {
                        case "\n\t":
                            delete(v5, fi28)
                        case "":
                            for fi31 := 0; fi31 < 19; fi31++ {
                                var v32 string = (("a" + "foo bar") + "")
                                fmt.Println(fi28, v32)
                                var v33 uint64 = (v1 - (v1 | (v1 >> ((0) & 63))))
                                v34 := GenS2{F0: fi31, F1: 0, F2: "a"}
                                _ = fi31
                                _ = v32
                                _ = v33
                                _ = v34
                            }
                        case "a":
                            if false {
                                var v35 uint = 214
                                _ = v35
                            } else {
                                v1 += uint64(v2)
                                switch fi28 {
                                case 9223372036854775807:
                                    v3 += 0
                                case -642:
                                    if (!false) {
                                        if len(v4) > 0 {
                                            v4[0] = "\n\t"
                                        }
                                    }
                                default:
                                    if (false && false) {
                                        for fi36 := 0; fi36 < 16; fi36++ {
                                            v3 += (v3 % ((int16(v2)) | 1))
                                            _ = fi36
                                        }
                                        if len(v4) > 0 {
                                            v4[0] = "hello"
                                        }
                                    } else {
                                        fmt.Println(v2)
                                        if len(v4) >= 2 {
                                            v4 = v4[1:len(v4)]
                                        }
                                        if false {
                                            v3 += -32768
                                        }
                                    }
                                }
                                v37 := GenS3{F0: (v1 * v1), F1: (fi28 * (fi28 * 0))}
                                _ = v37
                            }
                        default:
                            var v38 int = (^(fi28))
                            _ = v38
                        }
                        _ = v30
                    }
                }
                switch ("" + "") {
                case "foo bar":
                    if len(v4) > 0 {
                        v4[0] = "\n\t"
                    }
                default:
                    var v39 uint = 18446744073709551615
                    _ = v39
                }
                _ = fi28
            }
            var v40 uint32 = 42
            _ = v40
        }
    }
    var v41 uint8 = v2
    _ = len(v5)
    var v42 uint32 = 1
    var v43 int32 = -453
    var v44 int64 = 42
    var v45 int32 = v43
    switch "foo bar" {
    case "\n\t":
        v1 = (v1 / (((^(v1))) | 1))
    case "":
        v46 := GenS1{F0: uint8(v1), F1: (true || false)}
        _ = v46
    case "hello":
        v41 += (v41 / ((0) | 1))
    }
    v41 += (v41 - (v2 & uint8(v1)))
    for fi47 := 0; fi47 < 17; fi47++ {
        for fi48 := 0; fi48 < 12; fi48++ {
            fmt.Println(v41, fi47, fi48)
            fmt.Println(v44, v1)
            _ = fi48
        }
        _ = fi47
    }
    if true {
        for fi49 := 0; fi49 < 6; fi49++ {
            for fi50 := 0; fi50 < 20; fi50++ {
                var v51 uint16 = 65535
                switch fi50 {
                case -383:
                    v52 := GenS4{F0: v1, F1: (42 >= 0), F2: int8(v44)}
                    _ = v52
                default:
                    if false {
                        var v53 int16 = -32768
                        v2 = (^(v2))
                        fmt.Println(v42, v1, v42)
                        var v54 int = (fi49 ^ -1)
                        v45 = v45
                        if true {
                            v55 := GenS3{F0: (^(v1)), F1: 584}
                            v56 := GenS3{F0: 2, F1: (v54 % ((genFn1()) | 1))}
                            var v57 int16 = (v53 << ((4294967295) & 63))
                            if ((v53 ^ v3) == v57) {
                                switch -2147483648 {
                                case 9223372036854775807:
                                    var v58 int = (fi49 % (((fi49 ^ genFn0(4294967295))) | 1))
                                    _ = v58
                                }
                                if ((65535 > 0) && false) {
                                    v57 -= int16(v1)
                                    v59 := GenS2{F0: fi50, F1: (v42 << ((743) & 63)), F2: "a"}
                                    _ = v59
                                }
                                v42 = (v42 / ((v42) | 1))
                            }
                            if true {
                                v42 = 4294967295
                            }
                            _ = v55
                            _ = v56
                            _ = v57
                        } else {
                            var v60 int16 = v53
                            _ = v60
                        }
                        _ = v53
                        _ = v54
                    } else {
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        v3 = -511
                        var v61 uint = 0
                        if false {
                            v44 = v44
                            var v62 uint32 = (^(v42))
                            v3 = (v3 + v3)
                            var v63 uint = 775
                            _ = v62
                            _ = v63
                        }
                        fmt.Println(v44, v44)
                        _ = v61
                    }
                }
                v64 := v5[int(v51)]
                _ = fi50
                _ = v51
                _ = v64
            }
            v2 += v41
            switch (fi49 << ((630) & 63)) {
            case 32767:
                v65 := GenS5{F0: 0, F1: v3}
                _ = v65
            case 127:
                switch "hello" {
                case "a":
                    if (v2 < (v41 * (v41 % ((2) | 1)))) {
                        v43 -= (v43 + 140)
                        switch ("hello" + (("foo bar" + ("\n\t" + ("a" + ("" + "\n\t")))) + ("\n\t" + ("\n\t" + "foo bar")))) {
                        case "":
                            delete(v5, 32768)
                        }
                    }
                default:
                    v66 := GenS6{F0: 84}
                    _ = v66
                }
            }
            _ = fi49
        }
    }
    v41 += 69
    fmt.Println(v44, v1)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v15
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v41
    _ = v42
    _ = v43
    _ = v44
    _ = v45
    _ = v5
}
