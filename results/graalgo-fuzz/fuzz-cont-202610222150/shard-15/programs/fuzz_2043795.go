package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32; F2 bool }
type GenS2 struct { F0 string; F1 uint8 }
type GenS3 struct { F0 int8; F1 uint8 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 string; F1 int32; F2 uint }
type GenS6 struct { F0 int8; F1 uint8; F2 int8 }
type GenS7 struct { F0 uint32; F1 int16; F2 bool }
func genFn0(p0 int64) int32 {
    var v4 int16 = -275
    switch -9223372036854775808 {
    case -1:
        v5 := []int{32768, 0, 127}
        _ = v5
    case 42:
        switch ("\n\t" + ("foo bar" + (("a" + "") + ("" + "\n\t")))) {
        case "foo bar":
            v6 := GenS1{F0: 1, F1: 4294967295, F2: true}
            _ = v6
        case "\n\t":
            var v7 uint16 = 65535
            _ = v7
        case "hello":
            p0 -= p0
        default:
            fmt.Println(v4, v4)
        }
    default:
        var v8 bool = (572 <= 516)
        _ = v8
    }
    switch 715 {
    case 127:
        p0 = (-(p0))
    case -1:
        v9 := []int{}
        _ = v9
    case 75:
        v10 := []int{-9223372036854775808, 32768, 815, 128}
        _ = v10
    }
    _ = p0
    _ = v4
    return 2147483647
}

func main() {
    var v1 string = "\n\t"
    var v2 float64 = 1e10
    var v3 int16 = -172
    v11 := []int{9223372036854775807, 92, -172, 9223372036854775807}
    for fi12 := 0; fi12 < 18; fi12++ {
        var v13 int8 = 127
        _ = fi12
        _ = v13
    }
    v2 += -2.25
    switch 128 {
    case 2147483647:
        switch "\n\t" {
        case "\n\t":
            v3 = (-(v3))
        }
    default:
        if false {
            v14 := GenS2{F0: "hello", F1: 176}
            if false {
                fmt.Println(v1, v1, v1)
            }
            _ = v14
        } else {
            var v15 int16 = v3
            var v16 float64 = v2
            for fi17 := 0; fi17 < 1; fi17++ {
                v15 += -885
                _ = fi17
            }
            _ = v15
            _ = v16
        }
    }
    var v18 int = 32767
    if len(v11) > 0 {
        v11[0] = (v18 & (v18 % (((v18 >> ((42) & 63))) | 1)))
    }
    var v19 int32 = 998
    var v20 uint64 = 1
    v19 = (-(v19))
    switch v1 {
    case "\n\t":
        for fi21 := 0; fi21 < 7; fi21++ {
            v3 += v3
            if len(v11) >= 2 {
                v11 = v11[1:len(v11)]
            }
            switch fi21 {
            case -5:
                if (716 == 65535) {
                    switch v1 {
                    case "":
                        switch (((v1 + v1) + v1) + "") {
                        case "\n\t":
                            switch v1 {
                            case "\n\t":
                                var v22 int8 = 100
                                _ = v22
                            default:
                                for fi23 := 0; fi23 < 11; fi23++ {
                                    var v24 string = v1
                                    _ = fi23
                                    _ = v24
                                }
                            }
                        case "":
                            v3 = v3
                        case "hello":
                            v19 += 0
                        default:
                            var v25 int16 = (v3 | v3)
                            _ = v25
                        }
                    }
                    if (1e10 < float64(3.14159)) {
                        v1 = (v1 + ("a" + "a"))
                    } else {
                        v20 = (v20 << ((339) & 63))
                        for fi26 := 0; fi26 < 11; fi26++ {
                            var v27 uint16 = 255
                            switch v1 {
                            case "\n\t":
                                v18 -= 42
                            case "foo bar":
                                v28 := GenS3{F0: 100, F1: 1}
                                _ = v28
                            }
                            v29 := GenS4{F0: (fi21 + 32768)}
                            v30 := GenS5{F0: ("a" + "hello"), F1: -2147483648, F2: 279}
                            _ = fi26
                            _ = v27
                            _ = v29
                            _ = v30
                        }
                        v1 = (("foo bar" + (("foo bar" + v1) + v1)) + (v1 + ""))
                    }
                }
            case 2147483647:
                if ((v2 > v2) && false) {
                    v20 += 4294967295
                    v1 = "a"
                    v11 = append(v11, 2147483647)
                } else {
                    v19 = v19
                    var v31 uint16 = 32767
                    var v32 uint32 = 2147483647
                    _ = v31
                    _ = v32
                }
            default:
                var v33 float64 = (v2 / v2)
                _ = v33
            }
            _ = fi21
        }
    case "a":
        var v34 uint32 = 65535
        _ = v34
    case "hello":
        if (-2.25 != 1.5) {
            v1 = v1
            v35 := GenS3{F0: 121, F1: 128}
            v36 := GenS6{F0: 1, F1: 0, F2: 0}
            _ = v35
            _ = v36
        }
    }
    switch (v1 + v1) {
    case "hello":
        fmt.Println(v19)
    case "\n\t":
        var v37 int8 = 100
        _ = v37
    default:
        v20 -= v20
    }
    v38 := GenS7{F0: 42, F1: -32768, F2: true}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v11
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v38
}
