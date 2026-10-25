package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int8 }
type GenS2 struct { F0 uint; F1 bool; F2 int16 }

func main() {
    v1 := make(map[string]int)
    v2 := []uint32{65535, 1}
    v3 := []uint32{65535, 627, 429, 65535}
    var v4 int = 32767
    fmt.Println(v4)
    if ((42 > -809) && false) {
        v4 -= -603
    } else {
        fmt.Println(v4)
        v4 += 4294967295
    }
    if (!false) {
        v4 += 32767
    } else {
        if false {
            v1[((("" + "") + "hello") + "a")] = v4
            v4 = (v4 ^ 460)
            for fi5 := 0; fi5 < 6; fi5++ {
                var v6 int8 = -128
                switch "foo bar" {
                case "":
                    _ = len(v1)
                case "hello":
                    for fi7 := 0; fi7 < 14; fi7++ {
                        fmt.Println(fi5, v4)
                        v8 := GenS1{F0: 1, F1: (v6 / ((v6) | 1))}
                        fmt.Println(v6, fi7)
                        _ = fi7
                        _ = v8
                    }
                case "a":
                    v6 += 1
                default:
                    switch "hello" {
                    case "":
                        fmt.Println(fi5)
                    default:
                        var v9 string = "\n\t"
                        _ = v9
                    }
                }
                for fi10 := 0; fi10 < 11; fi10++ {
                    switch "hello" {
                    case "foo bar":
                        v1["a"] = 42
                    case "":
                        fmt.Println(fi10)
                    case "\n\t":
                        switch (("" + "foo bar") + ((("" + "foo bar") + ("\n\t" + "\n\t")) + "foo bar")) {
                        case "a":
                            var v11 uint64 = 0
                            _ = v11
                        case "foo bar":
                            v4 = 4294967295
                            fmt.Println(fi10)
                        }
                    }
                    _ = fi10
                }
                _ = fi5
                _ = v6
            }
        } else {
            if true {
                v12 := GenS2{F0: 438, F1: (!false), F2: -271}
                _ = v12
            }
            v4 += (v4 >> ((42) & 63))
        }
        for fi13 := 0; fi13 < 11; fi13++ {
            v4 += fi13
            if len(v2) > 0 {
                v2[0] = 2147483647
            }
            _ = fi13
        }
        var v14 uint32 = 2147483647
        _ = v14
    }
    if ((0 <= 0) && false) {
        var v15 int64 = -984
        for fi16 := 0; fi16 < 3; fi16++ {
            v17 := GenS2{F0: 1, F1: false, F2: -32768}
            var v18 uint32 = 42
            fmt.Println(v18, v18)
            _ = fi16
            _ = v17
            _ = v18
        }
        v19 := GenS1{F0: 2147483647, F1: 0}
        _ = v15
        _ = v19
    }
    for fi20 := 0; fi20 < 12; fi20++ {
        var v21 uint64 = 854
        _ = fi20
        _ = v21
    }
    for fi22 := 0; fi22 < 4; fi22++ {
        var v23 float64 = float64((((-2.25 / (0.0 + float32(0.1))) - ((-2.25 - float32(((1e10 * 0.0) - -2.25))) + 0.0)) / 0.1))
        _ = len(v1)
        _ = fi22
        _ = v23
    }
    switch ("\n\t" + ("" + "hello")) {
    case "foo bar":
        for fi24 := 0; fi24 < 5; fi24++ {
            for fi25 := 0; fi25 < 11; fi25++ {
                switch (-(fi24)) {
                case -993:
                    v26 := GenS2{F0: 0, F1: true, F2: 363}
                    _ = v26
                case 942:
                    fmt.Println(fi24, fi25, fi25)
                case 572:
                    var v27 uint = 2
                    _ = v27
                }
                v28 := GenS2{F0: 190, F1: (-1 == -9223372036854775808), F2: -787}
                fmt.Println(fi24, fi25)
                _ = fi25
                _ = v28
            }
            if (((3.14159 * 3.14159) * float64(3.14159)) < float64(float32(1.5))) {
                fmt.Println(fi24, v4)
                switch "hello" {
                case "hello":
                    for fi29 := 0; fi29 < 9; fi29++ {
                        v4 += 9223372036854775807
                        var v30 uint64 = 893
                        _ = fi29
                        _ = v30
                    }
                case "\n\t":
                    fmt.Println(fi24)
                default:
                    v4 += v4
                }
            } else {
                var v31 int16 = -497
                v4 += 128
                _ = v31
            }
            for fi32 := 0; fi32 < 6; fi32++ {
                v4 -= (^(v4))
                if len(v3) > 0 {
                    v3[0] = 2147483648
                }
                v33 := GenS1{F0: -2147483648, F1: -128}
                _ = fi32
                _ = v33
            }
            _ = fi24
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
}
