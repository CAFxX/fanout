package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int16; F1 uint64; F2 uint }
type GenS3 struct { F0 int32; F1 int32 }
func genFn0(p0 int64) int16 {
    if false {
        switch 32767 {
        case 2147483647:
            fmt.Println(p0)
        case -9223372036854775808:
            v5 := []int{1, 128, -2147483648, 368}
            _ = v5
        case -345:
            var v6 float32 = 3.14159
            _ = v6
        default:
            var v7 int64 = (p0 & -989)
            _ = v7
        }
    }
    v8 := []int{int(p0), -1}
    if len(v8) >= 2 {
        v8 = v8[1:len(v8)]
    }
    fmt.Println(p0)
    var v9 int16 = -1
    _ = p0
    _ = v8
    _ = v9
    return v9
}
func genFn1() uint8 {
    var v10 int64 = 42
    v11 := []int{}
    if len(v11) >= 2 {
        v11 = v11[1:len(v11)]
    }
    var v12 uint64 = 256
    _ = v10
    _ = v11
    _ = v12
    return 39
}

func main() {
    var v1 int64 = -1
    var v2 uint = 4294967295
    var v3 float64 = 1e10
    var v4 int32 = -2147483648
    if (4294967295 <= 0) {
        fmt.Println(v1, v4, v2)
    }
    v4 = v4
    switch "foo bar" {
    case "foo bar":
        var v13 uint32 = 2
        _ = v13
    case "hello":
        v14 := []int{32768, 9223372036854775807, 35, 128}
        _ = v14
    case "\n\t":
        var v15 uint16 = 65535
        _ = v15
    }
    for fi16 := 0; fi16 < 15; fi16++ {
        v1 = v1
        fmt.Println(fi16)
        var v17 uint16 = 0
        _ = fi16
        _ = v17
    }
    v18 := []int{}
    v19 := GenS1{F0: 1}
    v4 += v4
    var v20 float32 = 1e10
    var v21 uint32 = 2147483647
    if (128 >= uint8(v2)) {
        var v22 float32 = v20
        _ = v22
    } else {
        switch "" {
        case "foo bar":
            var v23 string = ("" + (("a" + "a") + (("foo bar" + "hello") + ("foo bar" + ""))))
            _ = v23
        case "hello":
            if true {
                switch -2147483648 {
                case 925:
                    v24 := GenS2{F0: genFn0(867), F1: 0, F2: (v2 << (((^(v2))) & 63))}
                    _ = v24
                case 32767:
                    v25 := GenS1{F0: 1}
                    _ = v25
                default:
                    if false {
                        var v26 int64 = -27
                        v4 += v4
                        _ = v26
                    }
                }
                v21 = v21
            } else {
                if false {
                    var v27 int8 = -128
                    _ = v27
                }
                var v28 bool = (false && false)
                var v29 string = (("foo bar" + "foo bar") + "")
                _ = v28
                _ = v29
            }
        default:
            if true {
                switch "\n\t" {
                case "a":
                    for fi30 := 0; fi30 < 17; fi30++ {
                        var v31 uint64 = 42
                        v21 = uint32(v2)
                        _ = fi30
                        _ = v31
                    }
                case "":
                    v18 = append(v18, 32767)
                }
                fmt.Println(v1, v21)
                v20 = (v20 * float32(0.0))
            } else {
                switch 42 {
                case 32768:
                    v4 += (v4 - 1)
                case 127:
                    switch 1 {
                    case 676:
                        for fi32 := 0; fi32 < 11; fi32++ {
                            v33 := GenS1{F0: 0}
                            var v34 int8 = 42
                            if true {
                                v35 := GenS1{F0: (fi32 - -412)}
                                switch ("foo bar" + "hello") {
                                case "\n\t":
                                    if (!(false || false)) {
                                        var v36 uint16 = 2
                                        _ = v36
                                    }
                                case "foo bar":
                                    var v37 int16 = 32767
                                    _ = v37
                                case "":
                                    for fi38 := 0; fi38 < 5; fi38++ {
                                        var v39 uint32 = 2147483647
                                        if false {
                                            v39 -= (v21 % (((v39 >> (((^(v2))) & 63))) | 1))
                                            v40 := GenS3{F0: 0, F1: v4}
                                            _ = v40
                                        }
                                        fmt.Println(v34)
                                        _ = fi38
                                        _ = v39
                                    }
                                default:
                                    for fi41 := 0; fi41 < 4; fi41++ {
                                        fmt.Println(v1)
                                        _ = fi41
                                    }
                                }
                                _ = v35
                            } else {
                                var v42 float64 = (1e10 + v3)
                                var v43 string = "foo bar"
                                for fi44 := 0; fi44 < 1; fi44++ {
                                    if len(v18) > 0 {
                                        v18[0] = fi44
                                    }
                                    var v45 string = "hello"
                                    if len(v18) > 0 {
                                        v18[0] = (fi44 | fi44)
                                    }
                                    var v46 int8 = (v34 / ((25) | 1))
                                    _ = fi44
                                    _ = v45
                                    _ = v46
                                }
                                _ = v42
                                _ = v43
                            }
                            _ = fi32
                            _ = v33
                            _ = v34
                        }
                    }
                case -1:
                    v47 := GenS1{F0: 127}
                    _ = v47
                default:
                    for fi48 := 0; fi48 < 2; fi48++ {
                        v21 += (v21 | 900)
                        _ = fi48
                    }
                }
                fmt.Println(v1)
                v1 = (v1 + -949)
            }
        }
        v3 -= (v3 / v3)
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
}
