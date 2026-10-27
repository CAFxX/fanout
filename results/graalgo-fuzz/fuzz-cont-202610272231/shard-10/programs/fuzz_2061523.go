package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint }
type GenS2 struct { F0 string; F1 string; F2 int16 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint8; F1 bool; F2 int }
func genFn0(p0 uint) int16 {
    switch "a" {
    case "":
        if false {
            v4 := []int{128, 13, 128}
            fmt.Println(p0)
            _ = v4
        } else {
            var v5 uint16 = 974
            p0 = (p0 * (^(p0)))
            _ = v5
        }
    case "hello":
        fmt.Println(p0)
    }
    for fi6 := 0; fi6 < 4; fi6++ {
        if true {
            var v7 int8 = 0
            p0 = p0
            v8 := GenS1{F0: 1, F1: 42}
            _ = v7
            _ = v8
        } else {
            for fi9 := 0; fi9 < 7; fi9++ {
                fmt.Println(fi9, p0, fi6)
                _ = fi9
            }
        }
        p0 += 2
        switch fi6 {
        case 782:
            fmt.Println(p0)
        default:
            var v10 int16 = 233
            _ = v10
        }
        _ = fi6
    }
    p0 = 993
    var v11 int8 = -1
    _ = p0
    _ = v11
    return 1
}

func main() {
    v1 := []string{"", "", "a", "hello"}
    v2 := make(map[int]string)
    v3 := []uint32{}
    var v12 uint = 42
    for fi13 := 0; fi13 < 3; fi13++ {
        if false {
            if false {
                v14 := GenS2{F0: "a", F1: "", F2: 250}
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                switch "hello" {
                case "foo bar":
                    switch "\n\t" {
                    case "foo bar":
                        v12 = 255
                    case "\n\t":
                        for fi15 := 0; fi15 < 20; fi15++ {
                            v12 = v12
                            var v16 int16 = genFn0(v12)
                            var v17 int64 = -1
                            _ = fi15
                            _ = v16
                            _ = v17
                        }
                    }
                case "a":
                    v12 += 314
                default:
                    v18 := GenS1{F0: 0, F1: (^(v12))}
                    _ = v18
                }
                _ = v14
            }
            v19 := v2[-517]
            var v20 uint32 = uint32(v12)
            _ = v19
            _ = v20
        } else {
            var v21 uint16 = 133
            v2[(fi13 ^ (fi13 + fi13))] = "a"
            _ = v21
        }
        _ = fi13
    }
    v12 += 255
    var v22 uint32 = 759
    v22 -= v22
    v2[-1] = "foo bar"
    if (((("foo bar" + "a") + "\n\t") + "foo bar") < ("foo bar" + (("hello" + ("hello" + ("\n\t" + ("" + "hello")))) + "hello"))) {
        for fi23 := 0; fi23 < 7; fi23++ {
            v24 := GenS1{F0: 1, F1: 18446744073709551615}
            _ = fi23
            _ = v24
        }
        _ = len(v3)
        v22 = v22
    } else {
        if false {
            if false {
                v25 := GenS3{F0: 2}
                for fi26 := 0; fi26 < 11; fi26++ {
                    v22 -= 42
                    _ = fi26
                }
                _ = v25
            } else {
                var v27 int8 = 1
                _ = v27
            }
            var v28 bool = (1 <= 404)
            for fi29 := 0; fi29 < 11; fi29++ {
                _ = len(v1)
                _ = fi29
            }
            _ = v28
        } else {
            if (255 <= 32767) {
                switch ("" + ("hello" + (("foo bar" + (("hello" + ("" + "hello")) + (("hello" + "") + ("a" + "a")))) + "\n\t"))) {
                case "\n\t":
                    for fi30 := 0; fi30 < 15; fi30++ {
                        v31 := GenS4{F0: 98, F1: false, F2: (-(fi30))}
                        _ = fi30
                        _ = v31
                    }
                case "":
                    var v32 uint32 = v22
                    _ = v32
                default:
                    fmt.Println(v22)
                }
                fmt.Println(v22, v12)
            }
        }
        v12 += 42
    }
    fmt.Println(v22, v22)
    v12 = v12
    var v33 int = 958
    for fi34 := 0; fi34 < 13; fi34++ {
        var v35 uint32 = (v22 << (((v12 << ((255) & 63))) & 63))
        var v36 int8 = -128
        _ = fi34
        _ = v35
        _ = v36
    }
    switch "" {
    case "hello":
        var v37 int64 = -1
        _ = v37
    case "\n\t":
        if true {
            var v38 bool = false
            v38 = (((0.1 / 3.14159) >= 0.0) || v38)
            fmt.Println(v38, v33, v38)
            _ = v38
        } else {
            if (0.1 == 1.5) {
                v33 += (v33 * (v33 + (v33 >> (((v12 << (((v12 + (v12 << ((255) & 63)))) & 63))) & 63))))
            }
        }
    case "foo bar":
        for fi39 := 0; fi39 < 19; fi39++ {
            switch fi39 {
            case -9223372036854775808:
                for fi40 := 0; fi40 < 2; fi40++ {
                    v12 = (^(v12))
                    fmt.Println(v12)
                    _ = fi40
                }
            case 127:
                if (!(211 != 147)) {
                    var v41 float64 = 0.1
                    if len(v3) > 0 {
                        v3[0] = v22
                    }
                    _ = v41
                }
            default:
                v33 += (fi39 | (v33 / ((1) | 1)))
            }
            _ = fi39
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v33:", v33)
    _ = v1
    _ = v12
    _ = v2
    _ = v22
    _ = v3
    _ = v33
}
