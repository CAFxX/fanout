package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int32; F1 uint64 }
type GenS3 struct { F0 uint64; F1 uint16; F2 uint64 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 uint16; F1 uint }
type GenS7 struct { F0 bool; F1 string }
func genFn0() uint8 {
    v7 := []int{250}
    var v8 uint64 = 0
    _ = v7
    _ = v8
    return 41
}
func genFn1(p0 int16, p1 uint) int64 {
    p0 += (p0 - p0)
    var v9 int8 = -1
    var v10 float32 = 3.14159
    _ = p0
    _ = p1
    _ = v10
    _ = v9
    return 1
}

func main() {
    var v1 string = "\n\t"
    var v2 uint = 1
    var v3 int16 = -87
    var v4 float64 = 0.1
    v5 := []string{v1, (v1 + v1), (("\n\t" + v1) + v1)}
    v6 := make(map[string]uint32)
    v11 := GenS1{F0: (^(v2))}
    var v12 int8 = -128
    v3 = 1
    var v13 float32 = float32((-2.25 - (v4 + ((v4 - v4) - (((0.1 / v4) + 3.14159) / (v4 + v4))))))
    if (v1 != v1) {
        v3 += v3
        if true {
            var v14 uint32 = 4294967295
            var v15 int = -2147483648
            _ = v14
            _ = v15
        } else {
            if len(v5) > 0 {
                v5[0] = "hello"
            }
            v13 += v13
            var v16 int64 = 789
            _ = v16
        }
        for fi17 := 0; fi17 < 16; fi17++ {
            if true {
                for fi18 := 0; fi18 < 8; fi18++ {
                    if false {
                        v19 := GenS1{F0: v2}
                        v4 = v4
                        _ = v19
                    } else {
                        v20 := GenS2{F0: 1, F1: 779}
                        if ((4294967295 == 399) || false) {
                            fmt.Println(v2)
                        }
                        _ = v20
                    }
                    _ = fi18
                }
                for fi21 := 0; fi21 < 2; fi21++ {
                    v3 -= (v3 ^ v3)
                    _ = fi21
                }
                var v22 uint64 = 1
                _ = v22
            } else {
                if ((v2 > (v2 >> (((v2 - 2)) & 63))) && (-1 >= -1)) {
                    var v23 string = "hello"
                    if true {
                        switch (-(fi17)) {
                        case -1:
                            var v24 uint64 = 38
                            _ = v24
                        }
                        v13 = -2.25
                        for fi25 := 0; fi25 < 20; fi25++ {
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                            v1 = v1
                            _ = fi25
                        }
                    }
                    _ = v23
                }
                v3 = 0
                var v26 bool = false
                _ = v26
            }
            for fi27 := 0; fi27 < 16; fi27++ {
                switch fi17 {
                case 4294967295:
                    var v28 int32 = 42
                    _ = v28
                case 1:
                    fmt.Println(v4, fi27)
                    switch "a" {
                    case "a":
                        var v29 int16 = 32767
                        _ = v29
                    case "foo bar":
                        fmt.Println(v12)
                    case "":
                        var v30 uint = 65535
                        _ = v30
                    }
                    var v31 uint16 = 255
                    v32 := GenS1{F0: v2}
                    _ = v31
                    _ = v32
                case 127:
                    v3 += v3
                    var v33 bool = false
                    _ = v33
                }
                var v34 int8 = (v12 % ((v12) | 1))
                if ((764 <= -9223372036854775808) && (!true)) {
                    switch v1 {
                    case "hello":
                        if true {
                            if true {
                                var v35 uint8 = 255
                                _ = v35
                            }
                        } else {
                            _ = len(v5)
                            var v36 uint8 = 224
                            _ = v36
                        }
                    default:
                        if true {
                            fmt.Println(fi17, v34)
                        } else {
                            v37 := GenS1{F0: v2}
                            _ = v37
                        }
                    }
                    v38 := GenS3{F0: 578, F1: 2, F2: 635}
                    var v39 int8 = v34
                    v40 := GenS2{F0: 0, F1: 0}
                    _ = v38
                    _ = v39
                    _ = v40
                } else {
                    v41 := GenS3{F0: 705, F1: 1, F2: 18446744073709551615}
                    if false {
                        if true {
                            switch v1 {
                            case "\n\t":
                                v13 -= ((v13 + float32(v4)) - (v13 * (0.0 + v13)))
                            case "hello":
                                v12 = (v12 >> ((2) & 63))
                            case "a":
                                switch v1 {
                                case "foo bar":
                                    var v42 int32 = -447
                                    _ = v42
                                case "\n\t":
                                    var v43 uint16 = 255
                                    _ = v43
                                case "":
                                    if (25 > (fi27 | (fi17 >> ((v2) & 63)))) {
                                        v44 := GenS4{F0: 0}
                                        _ = v44
                                    }
                                default:
                                    var v45 uint = v2
                                    _ = v45
                                }
                                switch "foo bar" {
                                case "hello":
                                    v2 = v2
                                }
                            default:
                                v2 = 0
                            }
                            v46 := GenS5{F0: 1}
                            v34 = int8(v3)
                            var v47 string = "foo bar"
                            _ = v46
                            _ = v47
                        }
                    }
                    switch v1 {
                    case "\n\t":
                        v34 = (v34 % ((102) | 1))
                    case "foo bar":
                        var v48 uint8 = 255
                        _ = v48
                    default:
                        var v49 uint64 = 619
                        _ = v49
                    }
                    _ = v41
                }
                _ = fi27
                _ = v34
            }
            _ = fi17
        }
    } else {
        v4 = v4
        var v50 uint32 = 2147483647
        _ = v50
    }
    fmt.Println(v3, v4)
    v4 = 0.0
    if (1 != v3) {
        var v51 uint32 = 603
        v52 := GenS5{F0: 979}
        switch ((v1 + v1) + v1) {
        case "foo bar":
            fmt.Println(v13, v4, v4)
        case "\n\t":
            var v53 float64 = v4
            _ = v53
        default:
            v2 = (^(v2))
        }
        _ = v51
        _ = v52
    } else {
        v2 -= v2
        switch -9223372036854775808 {
        case 32767:
            var v54 int16 = (v3 - (v3 & (v3 << (((^(v2))) & 63))))
            _ = v54
        case -721:
            if true {
                var v55 int32 = 1
                var v56 int8 = (v12 / ((-81) | 1))
                _ = v55
                _ = v56
            } else {
                var v57 string = (("\n\t" + "hello") + ("hello" + (("hello" + v1) + ((("foo bar" + ((v1 + "a") + "\n\t")) + v1) + v1))))
                v57 = v57
                var v58 bool = true
                _ = v57
                _ = v58
            }
        case 851:
            switch (v1 + v1) {
            case "a":
                switch ("a" + ("" + (v1 + "hello"))) {
                case "foo bar":
                    fmt.Println(v13)
                case "":
                    v59 := GenS5{F0: 18446744073709551615}
                    _ = v59
                case "\n\t":
                    v1 = (v1 + v1)
                default:
                    var v60 uint8 = 0
                    _ = v60
                }
            case "":
                v61 := GenS6{F0: 42, F1: (v2 % ((v2) | 1))}
                _ = v61
            default:
                v1 = (v1 + v1)
            }
        default:
            _ = len(v5)
        }
        v12 = v12
    }
    v3 = (v3 ^ (^(v3)))
    if true {
        v62 := GenS7{F0: (42 < 404), F1: ("foo bar" + ((v1 + ("" + "\n\t")) + ("hello" + ("foo bar" + "hello"))))}
        var v63 uint16 = 65535
        fmt.Println(v1)
        _ = v62
        _ = v63
    }
    var v64 int16 = v3
    var v65 int32 = 175
    v1 = ("a" + "\n\t")
    v4 = (float64(float32(1e10)) + v4)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v64:", v64)
    fmt.Println("v65:", v65)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v64
    _ = v65
}
