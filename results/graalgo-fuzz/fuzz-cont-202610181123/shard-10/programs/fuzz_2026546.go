package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8; F2 string }
type GenS2 struct { F0 string; F1 int }
type GenS3 struct { F0 int; F1 int8 }
type GenS4 struct { F0 int }

func main() {
    var v1 int16 = 704
    var v2 uint8 = 1
    v3 := make(map[string]uint32)
    v4 := []uint32{}
    var v5 uint = uint(v2)
    var v6 int8 = 1
    v6 = v6
    var v7 float64 = 3.14159
    var v8 int64 = 9223372036854775807
    for fi9 := 0; fi9 < 7; fi9++ {
        var v10 int16 = int16(fi9)
        fmt.Println(v5, v8, v5)
        var v11 int8 = (-(v6))
        _ = fi9
        _ = v10
        _ = v11
    }
    for fi12 := 0; fi12 < 11; fi12++ {
        v13 := GenS1{F0: 255, F1: 0, F2: "hello"}
        v7 = 1e10
        v5 -= (v5 / (((v5 - 65535)) | 1))
        _ = fi12
        _ = v13
    }
    if (-1 >= -555) {
        if (float64((1e10 + float32(v7))) < v7) {
            v2 += v2
            for fi14 := 0; fi14 < 12; fi14++ {
                if false {
                    var v15 int32 = 824
                    _ = v15
                } else {
                    var v16 uint32 = 794
                    var v17 int64 = v8
                    _ = len(v4)
                    _ = v16
                    _ = v17
                }
                _ = fi14
            }
        } else {
            fmt.Println(v8, v7)
            var v18 uint16 = 32767
            _ = v18
        }
        v7 = float64(0.1)
    } else {
        for fi19 := 0; fi19 < 12; fi19++ {
            for fi20 := 0; fi20 < 4; fi20++ {
                fmt.Println(v5, v2)
                var v21 int32 = -1
                _ = fi20
                _ = v21
            }
            var v22 float64 = 0.1
            _ = fi19
            _ = v22
        }
        var v23 int32 = 993
        _ = v23
    }
    var v24 int8 = (v6 << ((v5) & 63))
    switch -2147483648 {
    case 32768:
        v4 = append(v4, 2147483648)
    default:
        v25 := GenS2{F0: ("" + "\n\t"), F1: 2147483647}
        _ = v25
    }
    v6 = (v6 >> ((v5) & 63))
    fmt.Println(v1, v7, v5)
    _ = len(v3)
    var v26 int = 42
    if (uint32(v5) <= 65535) {
        if (v6 > (v6 - -1)) {
            v27 := GenS2{F0: ((("foo bar" + ("" + ("a" + ("hello" + "\n\t")))) + "a") + ""), F1: v26}
            switch v26 {
            case 2147483647:
                v28 := GenS2{F0: ("\n\t" + "\n\t"), F1: v26}
                _ = v28
            case 32768:
                v29 := v3["\n\t"]
                _ = v29
            case -9223372036854775808:
                v7 += v7
            }
            _ = v27
        } else {
            if false {
                switch "\n\t" {
                case "\n\t":
                    if ((("a" != "\n\t") || (2147483648 == 4294967295)) || true) {
                        for fi30 := 0; fi30 < 1; fi30++ {
                            fmt.Println(fi30)
                            _ = fi30
                        }
                        v8 -= (v8 ^ v8)
                    } else {
                        v5 += v5
                        var v31 float64 = ((((((((3.14159 + ((1.5 / v7) / v7)) / float64(float32(v7))) * v7) / v7) / 3.14159) / -2.25) / -2.25) * 1e10)
                        _ = v31
                    }
                default:
                    switch (("hello" + "hello") + ("foo bar" + (("hello" + ("hello" + (("" + "hello") + ""))) + "hello"))) {
                    case "":
                        fmt.Println(v2)
                    default:
                        _ = len(v3)
                    }
                }
            }
        }
        for fi32 := 0; fi32 < 17; fi32++ {
            for fi33 := 0; fi33 < 18; fi33++ {
                var v34 int16 = (v1 * (v1 - v1))
                _ = fi33
                _ = v34
            }
            _ = fi32
        }
        var v35 bool = false
        _ = v35
    } else {
        var v36 int8 = 42
        fmt.Println(v2)
        _ = v36
    }
    for fi37 := 0; fi37 < 10; fi37++ {
        v38 := GenS3{F0: (fi37 + v26), F1: v6}
        v39 := GenS4{F0: (-(v26))}
        for fi40 := 0; fi40 < 8; fi40++ {
            v26 -= fi40
            _ = fi40
        }
        _ = fi37
        _ = v38
        _ = v39
    }
    fmt.Println(v5, v26)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v24
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
