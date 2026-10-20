package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint16; F2 uint8 }
type GenS2 struct { F0 int32; F1 uint16; F2 uint16 }
type GenS3 struct { F0 int16; F1 int16 }
type GenS4 struct { F0 uint; F1 uint64; F2 uint8 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 int16; F1 int8 }
type GenS7 struct { F0 uint }
func genFn0() uint {
    var v6 float64 = (3.14159 * 0.1)
    if (false && (true && (true || (78 < -101)))) {
        for fi7 := 0; fi7 < 12; fi7++ {
            fmt.Println(fi7)
            if true {
                v6 += 3.14159
                var v8 float32 = (3.14159 * 0.0)
                fmt.Println(fi7, v6)
                _ = v8
            }
            _ = fi7
        }
        var v9 int16 = -32768
        var v10 string = ("hello" + ("\n\t" + "foo bar"))
        _ = v10
        _ = v9
    } else {
        var v11 int64 = 586
        v6 = v6
        _ = v11
    }
    v6 += v6
    v12 := GenS1{F0: ("foo bar" + "hello"), F1: 0, F2: 214}
    v6 -= (1e10 / ((v6 / (1.5 - v6)) - (3.14159 * v6)))
    _ = v12
    _ = v6
    return 65535
}
func genFn1(p0 uint32) uint {
    for fi13 := 0; fi13 < 8; fi13++ {
        p0 += 42
        fmt.Println(p0, p0)
        _ = fi13
    }
    if false {
        var v14 uint = 255
        _ = v14
    }
    _ = p0
    return 232
}

func main() {
    var v1 int8 = -1
    var v2 uint = 4294967295
    v3 := make(map[string]string)
    var v4 float64 = -2.25
    v5 := make(map[string]uint32)
    fmt.Println(v2, v2, v1)
    if ((v1 % ((v1) | 1)) > v1) {
        for fi15 := 0; fi15 < 11; fi15++ {
            if (1 != 2) {
                var v16 int32 = -60
                if ((!false) || true) {
                    if (18446744073709551615 != 42) {
                        if (!true) {
                            v2 += v2
                        } else {
                            v16 -= 42
                        }
                        v1 = v1
                    } else {
                        v17 := GenS2{F0: (v16 | 2147483647), F1: 2, F2: 65535}
                        for fi18 := 0; fi18 < 8; fi18++ {
                            v19 := []int{}
                            switch ("a" + "a") {
                            case "hello":
                                var v20 uint16 = 32767
                                _ = v20
                            case "foo bar":
                                var v21 uint8 = 223
                                _ = v21
                            case "a":
                                v16 = 0
                            default:
                                if ((true && ((fi18 & fi15) != fi15)) && ((0.1 + 0.1) > 0.1)) {
                                    var v22 string = "\n\t"
                                    if false {
                                        var v23 uint = 42
                                        if (4294967295 == 948) {
                                            v4 = (0.1 + ((v4 * 0.1) * 1.5))
                                            switch (fi18 ^ (fi15 % ((32767) | 1))) {
                                            case -9223372036854775808:
                                                if (("hello" >= ((v22 + v22) + (v22 + (((v22 + v22) + "hello") + "hello")))) && ((0.0 + -2.25) != 3.14159)) {
                                                    if true {
                                                        var v24 uint = genFn0()
                                                        _ = v24
                                                    } else {
                                                        delete(v5, v22)
                                                        v4 = (((((-2.25 + v4) + v4) + v4) - 1e10) / 0.1)
                                                        v22 = (v22 + ((v22 + ((((((v22 + ("foo bar" + ((((v22 + (v22 + "hello")) + "\n\t") + v22) + ((("" + ("foo bar" + v22)) + v22) + "a")))) + "foo bar") + v22) + ("foo bar" + ((v22 + "foo bar") + v22))) + "") + ("a" + ("" + ("\n\t" + v22))))) + ""))
                                                    }
                                                }
                                            default:
                                                v25 := GenS1{F0: "\n\t", F1: 65535, F2: 153}
                                                _ = v25
                                            }
                                        }
                                        if ((((2 != 771) && false) && false) || true) {
                                            fmt.Println(v4, fi18)
                                            var v26 uint64 = 820
                                            _ = v26
                                        } else {
                                            v27 := GenS3{F0: -32768, F1: int16(v16)}
                                            _ = v27
                                        }
                                        v16 = v16
                                        v22 = v22
                                        _ = v23
                                    }
                                    _ = v22
                                }
                            }
                            fmt.Println(v2, fi15, v16)
                            var v28 uint8 = 128
                            var v29 int8 = (v1 / (((v1 - -55)) | 1))
                            if len(v19) >= 2 {
                                v19 = v19[1:len(v19)]
                            }
                            _ = fi18
                            _ = v19
                            _ = v28
                            _ = v29
                        }
                        if (2 > 255) {
                            v30 := []int{fi15, fi15}
                            _ = v30
                        }
                        _ = v17
                    }
                    switch "" {
                    case "\n\t":
                        v2 = (^(v2))
                    default:
                        v16 -= v16
                    }
                }
                _ = v16
            }
            _ = fi15
        }
        v31 := GenS3{F0: -1, F1: 1}
        v1 -= v1
        _ = v31
    } else {
        if true {
            fmt.Println(v2)
            var v32 int16 = -664
            v33 := GenS2{F0: 1, F1: 42, F2: 65535}
            _ = v32
            _ = v33
        }
    }
    var v34 uint64 = 42
    var v35 int = 32767
    v2 += v2
    v4 = v4
    for fi36 := 0; fi36 < 18; fi36++ {
        v37 := GenS4{F0: v2, F1: v34, F2: 2}
        v34 = (v34 ^ v34)
        if (!true) {
            if true {
                if (true && (-1 < -186)) {
                    _ = len(v5)
                    fmt.Println(fi36, v1)
                    var v38 uint8 = 255
                    _ = v38
                } else {
                    var v39 uint32 = uint32(v34)
                    v40 := GenS5{F0: ("foo bar" == "foo bar")}
                    v41 := GenS6{F0: -32768, F1: v1}
                    _ = v39
                    _ = v40
                    _ = v41
                }
            }
            v42 := []int{(v35 * (v35 | v35)), fi36, v35, 32767}
            _ = v42
        }
        _ = fi36
        _ = v37
    }
    var v43 string = "foo bar"
    v44 := GenS3{F0: -682, F1: -32768}
    switch (v35 << (((v2 | genFn1(2147483648))) & 63)) {
    case 32768:
        v45 := []int{(v35 % ((v35) | 1)), v35, (-(v35))}
        _ = v45
    case 32767:
        v4 = 0.0
    case 128:
        for fi46 := 0; fi46 < 6; fi46++ {
            var v47 int = (fi46 + 76)
            _ = fi46
            _ = v47
        }
    }
    fmt.Println(v34)
    fmt.Println(v34)
    if ((!true) && true) {
        v48 := GenS7{F0: v2}
        v1 += (v1 / (((v1 | (v1 + v1))) | 1))
        _ = v48
    }
    if (!((v1 < v1) && false)) {
        fmt.Println(v2)
        delete(v3, ("" + v43))
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v4
    _ = v43
    _ = v44
    _ = v5
}
