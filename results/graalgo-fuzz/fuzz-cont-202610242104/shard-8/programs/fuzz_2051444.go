package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int32; F1 uint64 }

func main() {
    v1 := []uint32{2147483648}
    var v2 uint8 = 187
    v3 := []uint32{}
    switch "hello" {
    case "\n\t":
        if (!false) {
            v4 := GenS1{F0: 218}
            switch "\n\t" {
            case "hello":
                if (0.1 < (((1e10 / (-2.25 / 1.5)) * 1e10) - 0.0)) {
                    v2 -= v2
                    var v5 bool = (v2 > 255)
                    _ = v5
                } else {
                    fmt.Println(v2)
                }
            case "a":
                fmt.Println(v2)
            case "":
                if (int64(v2) <= 1) {
                    var v6 int32 = -2147483648
                    if false {
                        fmt.Println(v2, v2)
                    } else {
                        var v7 uint64 = 289
                        if ((^(v6)) > v6) {
                            if false {
                                v6 -= -2147483648
                                v6 = 701
                            } else {
                                var v8 string = ((("a" + "hello") + ("a" + (((("foo bar" + "") + "\n\t") + "hello") + ((("hello" + ("a" + "foo bar")) + "a") + "\n\t")))) + ("foo bar" + ("a" + ("foo bar" + (("hello" + "\n\t") + (("foo bar" + ((("\n\t" + "\n\t") + "") + "a")) + "hello"))))))
                                _ = v8
                            }
                            v7 = (v7 % (((v7 ^ 2)) | 1))
                            v2 -= 255
                        }
                        _ = v7
                    }
                    for fi9 := 0; fi9 < 5; fi9++ {
                        fmt.Println(v2, fi9, fi9)
                        if true {
                            var v10 uint16 = 42
                            if (2147483647 >= 207) {
                                var v11 uint = 0
                                fmt.Println(v6, v6)
                                _ = v11
                            }
                            _ = v10
                        }
                        _ = fi9
                    }
                    _ = v6
                }
            }
            _ = v4
        }
    case "hello":
        var v12 uint8 = (v2 << ((619) & 63))
        _ = v12
    default:
        v13 := GenS1{F0: 546}
        _ = v13
    }
    var v14 uint = 42
    switch ("foo bar" + "\n\t") {
    case "a":
        fmt.Println(v2, v14)
    case "foo bar":
        for fi15 := 0; fi15 < 13; fi15++ {
            var v16 int32 = -167
            _ = fi15
            _ = v16
        }
    default:
        if true {
            v14 = v14
        } else {
            if true {
                var v17 int = 580
                _ = v17
            }
            fmt.Println(v14, v14)
            switch "hello" {
            case "a":
                v14 -= 4294967295
            case "foo bar":
                v14 = v14
            case "":
                v18 := GenS1{F0: 18446744073709551615}
                _ = v18
            default:
                fmt.Println(v14, v14)
            }
        }
    }
    v19 := GenS1{F0: (^(v14))}
    v20 := GenS1{F0: v14}
    v21 := GenS1{F0: (^(v14))}
    if (692 <= 2147483647) {
        for fi22 := 0; fi22 < 2; fi22++ {
            fmt.Println(v2, v2)
            var v23 int = (fi22 * (fi22 % (((fi22 | 127)) | 1)))
            _ = fi22
            _ = v23
        }
        for fi24 := 0; fi24 < 17; fi24++ {
            switch fi24 {
            case -496:
                v25 := GenS1{F0: v14}
                _ = v25
            case -574:
                v26 := GenS1{F0: (v14 + (v14 + 18446744073709551615))}
                _ = v26
            case -9223372036854775808:
                for fi27 := 0; fi27 < 5; fi27++ {
                    var v28 int16 = 0
                    _ = fi27
                    _ = v28
                }
            default:
                v29 := GenS2{F0: 42, F1: 4294967295}
                _ = v29
            }
            _ = fi24
        }
        if true {
            _ = len(v3)
            v14 += (v14 >> ((v14) & 63))
        }
    }
    switch -996 {
    case 343:
        if len(v3) > 0 {
            v3[0] = 0
        }
    default:
        for fi30 := 0; fi30 < 9; fi30++ {
            v2 += 42
            var v31 int16 = 1000
            fmt.Println(fi30, v2)
            _ = fi30
            _ = v31
        }
    }
    var v32 float32 = (1.5 - -2.25)
    if (true && (-10 < 2147483647)) {
        var v33 int8 = 1
        _ = v33
    } else {
        if len(v1) > 0 {
            v1[0] = 1
        }
        fmt.Println(v14)
        v14 += 255
    }
    v32 += float32(0.1)
    v32 -= 0.0
    v1 = append(v1, 64)
    if true {
        for fi34 := 0; fi34 < 6; fi34++ {
            v35 := GenS1{F0: v14}
            _ = fi34
            _ = v35
        }
        fmt.Println(v2, v14)
    }
    for fi36 := 0; fi36 < 7; fi36++ {
        v32 -= v32
        _ = fi36
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v32:", v32)
    _ = v1
    _ = v14
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v32
}
