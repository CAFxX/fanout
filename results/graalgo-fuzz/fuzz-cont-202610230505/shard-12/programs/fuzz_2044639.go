package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int64; F1 int64 }
type GenS4 struct { F0 uint8; F1 int16 }
type GenS5 struct { F0 int; F1 uint64; F2 int8 }

func main() {
    var v1 bool = ((false || (785 < 463)) || false)
    var v2 uint = 486
    v3 := []string{"foo bar"}
    for fi4 := 0; fi4 < 16; fi4++ {
        v1 = v1
        v2 += (v2 | (^(v2)))
        _ = fi4
    }
    for fi5 := 0; fi5 < 20; fi5++ {
        for fi6 := 0; fi6 < 6; fi6++ {
            var v7 int8 = -5
            if ((!(94 >= int8(v2))) && (false && (!v1))) {
                fmt.Println(v7)
                v8 := GenS1{F0: 1000, F1: 163}
                _ = v8
            }
            v1 = (4294967295 > 18446744073709551615)
            switch "foo bar" {
            case "a":
                v1 = true
            case "":
                if (2147483647 == 2) {
                    v9 := GenS1{F0: 32767, F1: 1}
                    _ = v9
                } else {
                    if (!v1) {
                        var v10 int16 = int16(fi6)
                        var v11 int64 = -9223372036854775808
                        _ = v10
                        _ = v11
                    }
                    if ((2147483647 >= -2147483648) || (!(!v1))) {
                        v12 := GenS2{F0: 2}
                        _ = v12
                    }
                }
            case "hello":
                v13 := GenS3{F0: 642, F1: 740}
                _ = v13
            default:
                v14 := GenS1{F0: 32767, F1: 510}
                _ = v14
            }
            v15 := GenS1{F0: -1, F1: 139}
            var v16 int64 = 803
            if v1 {
                if false {
                    v2 += v2
                    var v17 int32 = 2147483647
                    _ = v17
                }
                v2 = (v2 - (v2 >> (((v2 / ((543) | 1))) & 63)))
                if (!true) {
                    var v18 uint8 = 255
                    _ = v18
                } else {
                    if v1 {
                        fmt.Println(v1, fi6, v1)
                    } else {
                        v7 = v7
                        v2 = v2
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                    }
                    var v19 float32 = (-2.25 + (3.14159 - 0.1))
                    _ = v19
                }
                var v20 string = "foo bar"
                _ = v20
            }
            if v1 {
                fmt.Println(v16, fi6)
            }
            _ = fi6
            _ = v15
            _ = v16
            _ = v7
        }
        _ = fi5
    }
    switch -1 {
    case -2147483648:
        var v21 float32 = ((0.0 * -2.25) * 1.5)
        _ = v21
    case -630:
        switch "a" {
        case "\n\t":
            var v22 string = ("hello" + (("foo bar" + ("a" + "")) + ("" + "foo bar")))
            _ = v22
        case "hello":
            for fi23 := 0; fi23 < 18; fi23++ {
                if false {
                    switch 2147483647 {
                    case 277:
                        v24 := GenS4{F0: 255, F1: 1}
                        _ = v24
                    }
                }
                v25 := GenS5{F0: fi23, F1: 101, F2: 24}
                _ = fi23
                _ = v25
            }
        case "foo bar":
            v1 = true
        default:
            if false {
                v2 -= 4294967295
                for fi26 := 0; fi26 < 10; fi26++ {
                    for fi27 := 0; fi27 < 20; fi27++ {
                        v28 := GenS3{F0: 0, F1: 1}
                        _ = fi27
                        _ = v28
                    }
                    _ = fi26
                }
            } else {
                var v29 float32 = (((0.1 * 0.0) / -2.25) / 1.0)
                _ = v29
            }
        }
    default:
        for fi30 := 0; fi30 < 7; fi30++ {
            var v31 string = ("a" + "hello")
            _ = fi30
            _ = v31
        }
    }
    if len(v3) > 0 {
        v3[0] = (((("a" + "hello") + (((("" + ("foo bar" + "hello")) + "") + "foo bar") + ("foo bar" + "a"))) + "foo bar") + ("" + (("a" + (("a" + "") + ("\n\t" + ("foo bar" + "foo bar")))) + "a")))
    }
    v2 = v2
    for fi32 := 0; fi32 < 9; fi32++ {
        var v33 uint64 = 499
        var v34 uint16 = 42
        _ = fi32
        _ = v33
        _ = v34
    }
    fmt.Println(v2, v1)
    var v35 int16 = 1000
    if len(v3) > 0 {
        v3[0] = (("\n\t" + "\n\t") + "hello")
    }
    fmt.Println(v1, v1, v2)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v35:", v35)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
}
