package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint16 }
type GenS2 struct { F0 int8; F1 uint16; F2 int16 }
type GenS3 struct { F0 int64; F1 int16; F2 int }
type GenS4 struct { F0 int8 }

func main() {
    var v1 uint32 = 118
    var v2 float32 = 0.0
    var v3 bool = (0 < 297)
    var v4 uint = 42
    var v5 int16 = 221
    v6 := make(map[string]int)
    for fi7 := 0; fi7 < 6; fi7++ {
        v8 := GenS1{F0: (v3 || (!(!(v3 || (v5 != v5))))), F1: 255}
        v1 += 4294967295
        for fi9 := 0; fi9 < 2; fi9++ {
            v10 := []int{(fi7 >> (((v4 % ((v4) | 1))) & 63))}
            if (((^(fi9)) != fi7) || v3) {
                v5 = (v5 / ((v5) | 1))
                switch "" {
                case "foo bar":
                    var v11 uint8 = 127
                    _ = v11
                case "":
                    v4 = (v4 ^ v4)
                case "\n\t":
                    v2 = float32(float64(1.5))
                }
                switch "" {
                case "hello":
                    if ((!v3) || (-9223372036854775808 > 1)) {
                        if len(v10) >= 2 {
                            v10 = v10[1:len(v10)]
                        }
                    } else {
                        if len(v10) > 0 {
                            v10[0] = fi7
                        }
                        fmt.Println(fi7, v2)
                        var v12 float32 = float32(1.5)
                        var v13 float32 = v2
                        _ = v12
                        _ = v13
                    }
                    var v14 int8 = -63
                    v4 += (v4 << ((18446744073709551615) & 63))
                    _ = v14
                case "":
                    fmt.Println(v2)
                case "a":
                    var v15 int16 = v5
                    _ = v15
                default:
                    var v16 bool = (-108 == 42)
                    _ = v16
                }
            }
            _ = fi9
            _ = v10
        }
        _ = fi7
        _ = v8
    }
    v17 := []int{-9223372036854775808, 732}
    v1 = 65535
    switch ("\n\t" + "") {
    case "\n\t":
        var v18 uint32 = (v1 - (v1 << (((^(v4))) & 63)))
        _ = v18
    case "hello":
        if (42 >= (v4 / (((^(v4))) | 1))) {
            var v19 uint64 = 495
            if v3 {
                if len(v17) >= 2 {
                    v17 = v17[1:len(v17)]
                }
                v3 = v3
                v1 = v1
            }
            _ = v19
        }
    case "":
        v3 = false
    }
    v20 := GenS2{F0: 1, F1: 255, F2: -478}
    if len(v17) > 0 {
        v17[0] = -9223372036854775808
    }
    _ = len(v6)
    var v21 bool = (128 > 42)
    v22 := GenS2{F0: 0, F1: 620, F2: v5}
    v23 := GenS3{F0: int64(v4), F1: v5, F2: -618}
    v24 := GenS4{F0: 127}
    if (2147483647 != (^(v1))) {
        fmt.Println(v2)
        if ((!true) && ("a" <= (((((("hello" + (("\n\t" + "hello") + "a")) + "hello") + "foo bar") + (("" + ("a" + "\n\t")) + ("foo bar" + ""))) + ("\n\t" + "a")) + ""))) {
            switch "hello" {
            case "hello":
                for fi25 := 0; fi25 < 16; fi25++ {
                    var v26 string = "foo bar"
                    if v21 {
                        if len(v17) >= 2 {
                            v17 = v17[1:len(v17)]
                        }
                        v4 -= (v4 << ((4294967295) & 63))
                    } else {
                        fmt.Println(v4, v1, v2)
                    }
                    v27 := GenS3{F0: 0, F1: v5, F2: fi25}
                    _ = fi25
                    _ = v26
                    _ = v27
                }
            }
        } else {
            v1 = 2
        }
        var v28 float32 = (v2 + v2)
        _ = v28
    } else {
        var v29 string = "\n\t"
        var v30 uint64 = uint64(v4)
        _ = v29
        _ = v30
    }
    _ = len(v17)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v17
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
