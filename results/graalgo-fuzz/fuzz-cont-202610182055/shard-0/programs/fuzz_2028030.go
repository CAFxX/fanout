package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int64; F1 int16; F2 int }
type GenS4 struct { F0 string; F1 int16 }
type GenS5 struct { F0 int16; F1 bool; F2 string }
func genFn0(p0 int8, p1 int64, p2 uint64) int {
    p1 = (p1 & (p1 / ((0) | 1)))
    var v7 float64 = 0.0
    p2 = p2
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 26
}
func genFn1() int16 {
    v8 := GenS1{F0: 0, F1: 0}
    var v9 uint64 = 42
    v9 = 1
    fmt.Println(v9)
    var v10 int16 = 32767
    _ = v10
    _ = v8
    _ = v9
    return v10
}

func main() {
    v1 := make(map[string]int)
    var v2 uint32 = 819
    var v3 uint = 0
    var v4 uint64 = 2
    v5 := make(map[int]uint32)
    var v6 uint8 = 127
    for fi11 := 0; fi11 < 14; fi11++ {
        fmt.Println(v2, v3)
        _ = fi11
    }
    v2 += 2147483648
    fmt.Println(v2, v6, v3)
    if (((v4 % ((v4) | 1)) > (v4 - (v4 % ((18446744073709551615) | 1)))) || (-2147483648 == 221)) {
        fmt.Println(v6, v3, v3)
        fmt.Println(v2, v2, v6)
        var v12 uint16 = 42
        _ = v12
    }
    if (false || false) {
        v2 += v2
        v2 += v2
    } else {
        fmt.Println(v4)
        v13 := GenS2{F0: ((!(("hello" + "\n\t") == "a")) && true)}
        _ = v13
    }
    v6 -= 163
    var v14 int = 757
    if true {
        var v15 float64 = (float64(1.5) * ((1e10 / (1e10 / 0.1)) - float64(3.14159)))
        v14 += v14
        _ = v15
    } else {
        v16 := GenS1{F0: 0, F1: 255}
        var v17 int = v14
        for fi18 := 0; fi18 < 1; fi18++ {
            switch genFn0(100, -9223372036854775808, (v4 >> (((v3 - 898)) & 63))) {
            case 903:
                v3 = uint(v6)
            default:
                var v19 int16 = -32768
                _ = v19
            }
            v4 = (^(v4))
            _ = len(v1)
            _ = fi18
        }
        _ = v16
        _ = v17
    }
    for fi20 := 0; fi20 < 14; fi20++ {
        var v21 int16 = -379
        switch (fi20 / (((fi20 >> ((v3) & 63))) | 1)) {
        case -822:
            if false {
                v14 += 4294967295
            } else {
                var v22 int8 = -23
                for fi23 := 0; fi23 < 16; fi23++ {
                    if (!false) {
                        _ = len(v1)
                        v22 = (v22 ^ (v22 & v22))
                        v6 -= (v6 - v6)
                    }
                    v24 := []int{-9223372036854775808, (fi23 / ((v14) | 1))}
                    _ = fi23
                    _ = v24
                }
                v21 = (v21 / (((v21 ^ int16(v2))) | 1))
                _ = v22
            }
        case 42:
            var v25 int16 = genFn1()
            _ = v25
        case 128:
            if (0.0 <= ((0.1 / 3.14159) * (3.14159 / 1e10))) {
                v14 = 0
            } else {
                var v26 uint16 = 255
                _ = v26
            }
        default:
            v21 = (^(v21))
        }
        fmt.Println(v2)
        _ = fi20
        _ = v21
    }
    fmt.Println(v3, v2)
    if false {
        for fi27 := 0; fi27 < 11; fi27++ {
            fmt.Println(v4)
            var v28 float32 = 1.5
            _ = fi27
            _ = v28
        }
        v14 = v14
    }
    if true {
        v3 -= v3
    } else {
        v14 += genFn0(1, -9223372036854775808, (v4 >> (((v3 & v3)) & 63)))
        switch ("hello" + ("foo bar" + (("hello" + "a") + (("hello" + ("" + "foo bar")) + ("hello" + "\n\t"))))) {
        case "foo bar":
            var v29 uint8 = v6
            _ = v29
        case "hello":
            v30 := GenS3{F0: 476, F1: -230, F2: v14}
            _ = v30
        case "":
            v3 -= v3
        }
        var v31 uint8 = 255
        _ = v31
    }
    if false {
        v14 += 9223372036854775807
        for fi32 := 0; fi32 < 13; fi32++ {
            switch v14 {
            case 32767:
                fmt.Println(v14, v4)
            case 4294967295:
                var v33 float32 = (float32(((0.1 + ((1e10 - 3.14159) - 1e10)) + ((0.1 * 1.5) * float64(3.14159)))) - 1.5)
                _ = v33
            }
            switch "a" {
            case "hello":
                v14 = (v14 | v14)
            }
            _ = fi32
        }
        var v34 int = v14
        _ = v34
    } else {
        if false {
            switch "foo bar" {
            case "foo bar":
                v2 -= (^(v2))
            case "\n\t":
                v35 := GenS4{F0: "foo bar", F1: -1}
                _ = v35
            case "a":
                var v36 int8 = 1
                _ = v36
            }
            fmt.Println(v14, v3)
            v4 -= (v4 | v4)
        }
        v37 := GenS5{F0: -32768, F1: true, F2: ("hello" + "foo bar")}
        _ = v37
    }
    fmt.Println(v6, v3, v4)
    _ = len(v5)
    v3 = (v3 * v3)
    if true {
        v2 = v2
    } else {
        switch v14 {
        case 4294967295:
            v6 += (v6 * v6)
        case 42:
            if true {
                v3 += v3
                fmt.Println(v3, v6)
                for fi38 := 0; fi38 < 16; fi38++ {
                    v4 -= (v4 ^ v4)
                    if true {
                        var v39 float32 = (-2.25 * 1.5)
                        v4 = 115
                        if (!false) {
                            fmt.Println(v39, v3, v39)
                        } else {
                            v4 -= v4
                            switch ("" + (((("foo bar" + "\n\t") + "hello") + "a") + "hello")) {
                            case "foo bar":
                                v4 += 4294967295
                            default:
                                v40 := GenS3{F0: -9223372036854775808, F1: 1000, F2: v14}
                                _ = v40
                            }
                            fmt.Println(v2, v3, v2)
                        }
                        _ = v39
                    }
                    for fi41 := 0; fi41 < 19; fi41++ {
                        if ((genFn0(100, 0, (v4 * (v4 * v4))) >= fi41) && (!(!((!false) && true)))) {
                            var v42 uint32 = v2
                            fmt.Println(v3, v14)
                            switch (fi38 << ((4294967295) & 63)) {
                            case -1:
                                v43 := GenS5{F0: -1, F1: true, F2: "\n\t"}
                                _ = v43
                            }
                            _ = v42
                        } else {
                            var v44 int8 = 0
                            _ = len(v5)
                            delete(v5, genFn0(v44, -942, v4))
                            var v45 int = (-(v14))
                            _ = v44
                            _ = v45
                        }
                        _ = fi41
                    }
                    _ = fi38
                }
            } else {
                var v46 bool = (0 >= v6)
                _ = v46
            }
        }
    }
    var v47 uint32 = (v2 << (((^(v3))) & 63))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v47:", v47)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v47
    _ = v5
    _ = v6
}
