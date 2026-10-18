package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 bool; F1 string; F2 string }
type GenS3 struct { F0 uint64; F1 uint32 }
type GenS4 struct { F0 int }

func main() {
    var v1 int32 = 52
    var v2 bool = true
    var v3 uint64 = 4294967295
    var v4 int16 = 144
    var v5 string = "hello"
    var v6 uint64 = (v3 * 2)
    if false {
        var v7 uint = 65535
        switch 42 {
        case 146:
            v5 = (v5 + (v5 + ("\n\t" + "a")))
        }
        v8 := GenS1{F0: v7}
        _ = v7
        _ = v8
    }
    for fi9 := 0; fi9 < 8; fi9++ {
        if false {
            v4 += 448
            var v10 uint8 = 1
            _ = v10
        } else {
            switch "" {
            case "":
                switch "a" {
                case "\n\t":
                    var v11 int = fi9
                    _ = v11
                case "hello":
                    if true {
                        var v12 float64 = -2.25
                        v13 := GenS2{F0: true, F1: (v5 + ("\n\t" + v5)), F2: v5}
                        v14 := []int{}
                        _ = v12
                        _ = v13
                        _ = v14
                    }
                default:
                    var v15 uint = 4294967295
                    _ = v15
                }
            case "a":
                v6 = v3
            case "hello":
                if (189 != 572) {
                    v5 = v5
                    v16 := GenS2{F0: (2 == 42), F1: ("a" + v5), F2: v5}
                    if (v1 <= 158) {
                        var v17 string = v5
                        v18 := []int{}
                        v1 = (^(v1))
                        _ = v17
                        _ = v18
                    }
                    _ = v16
                } else {
                    if v2 {
                        v1 = v1
                    } else {
                        v6 -= v6
                        fmt.Println(v3, v4)
                        var v19 int = fi9
                        _ = v19
                    }
                    v20 := GenS3{F0: 1, F1: 2}
                    _ = v20
                }
            }
            fmt.Println(v6, fi9, v3)
        }
        _ = fi9
    }
    fmt.Println(v6, v1)
    for fi21 := 0; fi21 < 13; fi21++ {
        v2 = (731 < 255)
        switch v5 {
        case "foo bar":
            for fi22 := 0; fi22 < 10; fi22++ {
                switch (fi22 % ((127) | 1)) {
                case 9223372036854775807:
                    v23 := GenS1{F0: 1}
                    _ = v23
                default:
                    v24 := GenS2{F0: v2, F1: (v5 + (((v5 + (v5 + v5)) + ((v5 + v5) + ("a" + v5))) + (v5 + (v5 + v5)))), F2: v5}
                    _ = v24
                }
                switch ((v5 + v5) + v5) {
                case "\n\t":
                    v5 = (v5 + ("a" + v5))
                case "a":
                    v5 = v5
                }
                if v2 {
                    if (("hello" + v5) > v5) {
                        v4 = (v4 / (((v4 + 316)) | 1))
                        if (!v2) {
                            var v25 bool = (-9223372036854775808 == -771)
                            if v2 {
                                switch v5 {
                                case "hello":
                                    v4 = (v4 - 227)
                                }
                            } else {
                                var v26 float64 = ((1e10 - 1e10) - float64(0.1))
                                _ = v26
                            }
                            if true {
                                v6 = 1
                            } else {
                                v27 := GenS4{F0: fi21}
                                _ = v27
                            }
                            _ = v25
                        }
                    }
                }
                _ = fi22
            }
        }
        _ = fi21
    }
    var v28 uint8 = 75
    v29 := []int{1}
    var v30 float64 = ((1e10 + float64((float32((float64(((float32(3.14159) + float32(-2.25)) * ((3.14159 + (0.1 / 0.1)) / ((1e10 - (0.1 * 0.1)) + 3.14159)))) / (float64(1.5) - -2.25))) + 0.1))) - -2.25)
    if v2 {
        v6 = v3
        switch 489 {
        case -881:
            v29 = append(v29, 32767)
        }
        var v31 int32 = v1
        _ = v31
    }
    if len(v29) > 0 {
        v29[0] = 522
    }
    v6 = 960
    var v32 int16 = (^(v4))
    v32 = (^(v4))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v29:", len(v29), v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
