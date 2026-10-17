package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 bool; F1 uint16 }
type GenS4 struct { F0 bool; F1 int8 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 string; F1 bool }
type GenS7 struct { F0 bool; F1 int16 }
type GenS8 struct { F0 int }

func main() {
    var v1 uint32 = 2147483647
    v2 := make(map[string]bool)
    var v3 int64 = 0
    v4 := []int{-321}
    if (18446744073709551615 >= 2) {
        var v5 int = 128
        _ = v5
    } else {
        var v6 uint = 652
        v2["\n\t"] = false
        var v7 int = 4294967295
        _ = v6
        _ = v7
    }
    v1 = 42
    v3 -= v3
    var v8 int64 = (v3 / (((v3 >> ((0) & 63))) | 1))
    switch -2147483648 {
    case -794:
        if true {
            if (((v1 != (v1 << ((2) & 63))) || false) && (-2147483648 > 0)) {
                var v9 int = 32768
                _ = v9
            } else {
                v1 = 46
                v10 := GenS1{F0: 128, F1: (v8 & v3)}
                v3 += 42
                _ = v10
            }
            if false {
                var v11 uint16 = 2
                v1 -= v1
                v12 := GenS1{F0: 32767, F1: (-(v8))}
                _ = v11
                _ = v12
            } else {
                var v13 uint8 = 97
                _ = v13
            }
            var v14 string = ("hello" + "foo bar")
            _ = v14
        } else {
            var v15 uint16 = 267
            v15 += 406
            _ = v15
        }
    default:
        v3 -= (^(v3))
    }
    v1 = v1
    if (-1 != -966) {
        var v16 int8 = -36
        switch -1 {
        case -1:
            var v17 float32 = 0.1
            _ = v17
        case 415:
            v8 = (v8 << ((0) & 63))
        }
        v1 = v1
        _ = v16
    }
    var v18 int = -1
    if (1 > 44) {
        fmt.Println(v3)
        var v19 string = ("hello" + "")
        var v20 uint64 = 2
        _ = v19
        _ = v20
    }
    var v21 uint16 = 445
    if (127 != 127) {
        v18 = v18
        v21 -= 0
        delete(v2, ("\n\t" + "foo bar"))
    } else {
        v22 := v2["a"]
        if v22 {
            if ((v18 + v18) == 127) {
                for fi23 := 0; fi23 < 17; fi23++ {
                    if len(v4) > 0 {
                        v4[0] = -494
                    }
                    fmt.Println(v1, fi23, v18)
                    v2["\n\t"] = v22
                    _ = fi23
                }
                v24 := GenS2{F0: 255}
                _ = v24
            } else {
                v25 := GenS1{F0: 2147483647, F1: v8}
                var v26 int = (v18 >> ((431) & 63))
                v3 = (v3 ^ (v8 | -610))
                _ = v25
                _ = v26
            }
            v18 = v18
            if false {
                v18 = (v18 + v18)
                switch 32767 {
                case 0:
                    switch "" {
                    case "\n\t":
                        switch v18 {
                        case 2147483647:
                            if true {
                                var v27 int32 = 1
                                _ = v27
                            }
                        case 4294967295:
                            switch "\n\t" {
                            case "":
                                v3 = int64(v1)
                            default:
                                v21 += v21
                            }
                        case 0:
                            v28 := GenS1{F0: (v18 << ((5) & 63)), F1: int64(v1)}
                            _ = v28
                        }
                    default:
                        _ = len(v4)
                    }
                case 32767:
                    fmt.Println(v22, v21)
                }
            } else {
                v29 := GenS3{F0: (!(141 > v1)), F1: 627}
                v8 = (v3 & v8)
                v1 += v1
                _ = v29
            }
        } else {
            v18 = 32767
        }
        _ = v22
    }
    var v30 int = v18
    for fi31 := 0; fi31 < 5; fi31++ {
        v32 := GenS4{F0: false, F1: -128}
        if (false || false) {
            fmt.Println(v21)
            var v33 int8 = int8(v3)
            v21 += 42
            _ = v33
        } else {
            v34 := GenS5{F0: (!(true && false))}
            for fi35 := 0; fi35 < 6; fi35++ {
                _ = len(v4)
                switch 42 {
                case 127:
                    var v36 bool = (-2.25 < ((3.14159 * ((1.5 * 0.0) * float32(float64((((3.14159 - -2.25) + -2.25) * (3.14159 + 1e10)))))) * 0.0))
                    _ = v36
                case 9223372036854775807:
                    if (128 > 0) {
                        fmt.Println(v8, v3, fi31)
                    }
                default:
                    fmt.Println(v3)
                }
                v37 := GenS6{F0: (("foo bar" + (("a" + "") + "hello")) + "a"), F1: ((v21 & 634) == v21)}
                _ = fi35
                _ = v37
            }
            _ = v34
        }
        _ = fi31
        _ = v32
    }
    var v38 int16 = int16(v30)
    if true {
        v8 -= 9223372036854775807
        v39 := GenS7{F0: (26 != 255), F1: v38}
        _ = v39
    }
    for fi40 := 0; fi40 < 20; fi40++ {
        if false {
            switch (v18 << ((402) & 63)) {
            case -2147483648:
                var v41 uint64 = 60
                _ = v41
            case 135:
                var v42 float32 = (0.1 + ((1e10 / (((float32(((((0.0 * 1.5) + 1.5) - 1.5) - (0.1 / 1.5))) + (0.0 / 1.0)) + (-2.25 * -2.25)) - -2.25)) / 3.14159))
                _ = v42
            default:
                for fi43 := 0; fi43 < 5; fi43++ {
                    v44 := GenS8{F0: (^(fi43))}
                    var v45 string = "\n\t"
                    _ = fi43
                    _ = v44
                    _ = v45
                }
            }
        }
        v4 = append(v4, v18)
        _ = fi40
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v38:", v38)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v18
    _ = v2
    _ = v21
    _ = v3
    _ = v30
    _ = v38
    _ = v4
    _ = v8
}
