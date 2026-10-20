package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64; F2 int64 }
type GenS2 struct { F0 uint8; F1 bool }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint16 }

func main() {
    var v1 int16 = -95
    var v2 uint16 = 32767
    var v3 bool = false
    var v4 uint64 = 371
    for fi5 := 0; fi5 < 6; fi5++ {
        var v6 uint32 = 183
        _ = fi5
        _ = v6
    }
    if true {
        fmt.Println(v2, v3, v2)
        v3 = (!v3)
        switch 4294967295 {
        case 1:
            for fi7 := 0; fi7 < 13; fi7++ {
                v8 := GenS1{F0: 4294967295, F1: 42, F2: 1}
                v1 += v1
                for fi9 := 0; fi9 < 14; fi9++ {
                    var v10 int32 = 16
                    _ = fi9
                    _ = v10
                }
                _ = fi7
                _ = v8
            }
        }
    } else {
        var v11 float64 = (1.5 / 1.0)
        fmt.Println(v1, v1, v11)
        var v12 int8 = 127
        _ = v11
        _ = v12
    }
    v13 := []int{-2147483648, -2147483648, 127, 553}
    switch 4294967295 {
    case 1:
        for fi14 := 0; fi14 < 20; fi14++ {
            v15 := GenS2{F0: 128, F1: ((v2 >> ((544) & 63)) == (v2 | 32767))}
            _ = fi14
            _ = v15
        }
    case -9223372036854775808:
        for fi16 := 0; fi16 < 2; fi16++ {
            if (0.1 != (float64(0.1) + (1e10 / 1.0))) {
                v17 := GenS3{F0: ("a" + "\n\t")}
                if (("a" + "a") <= "") {
                    if len(v13) > 0 {
                        v13[0] = 4294967295
                    }
                } else {
                    v13 = append(v13, (-(fi16)))
                    if v3 {
                        if len(v13) >= 2 {
                            v13 = v13[1:len(v13)]
                        }
                    }
                    fmt.Println(v1)
                }
                if ((0.0 == 0.0) && true) {
                    if len(v13) > 0 {
                        v13[0] = fi16
                    }
                } else {
                    if ((!false) && ("hello" != ("hello" + "a"))) {
                        if (!(!(!(true || v3)))) {
                            v2 += 32767
                            v2 -= 65535
                        } else {
                            v13 = append(v13, 42)
                        }
                        fmt.Println(v1, v3)
                    }
                    v18 := GenS2{F0: 127, F1: ((v4 | v4) < 2)}
                    _ = v18
                }
                _ = v17
            }
            v4 = (v4 / ((v4) | 1))
            var v19 uint = 48
            _ = fi16
            _ = v19
        }
    case 9223372036854775807:
        if len(v13) >= 2 {
            v13 = v13[1:len(v13)]
        }
    default:
        for fi20 := 0; fi20 < 2; fi20++ {
            fmt.Println(fi20)
            _ = fi20
        }
    }
    if v3 {
        fmt.Println(v4, v3)
        for fi21 := 0; fi21 < 20; fi21++ {
            var v22 uint8 = 1
            var v23 float64 = 1e10
            fmt.Println(v3, v22, v1)
            _ = fi21
            _ = v22
            _ = v23
        }
        if (!v3) {
            v1 += 0
            v3 = false
            v24 := GenS4{F0: (v2 | (v2 << ((uint(v2)) & 63)))}
            _ = v24
        }
    }
    v13 = append(v13, 252)
    switch "foo bar" {
    case "\n\t":
        var v25 int64 = 9223372036854775807
        _ = v25
    }
    v4 = 0
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
}
