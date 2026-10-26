package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int8 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint8; F1 int16; F2 uint }
type GenS4 struct { F0 uint64; F1 int16; F2 int16 }
type GenS5 struct { F0 int32 }

func main() {
    v1 := make(map[int]string)
    var v2 uint64 = 4294967295
    v3 := make(map[string]string)
    var v4 float64 = (((3.14159 / 3.14159) + (float64(3.14159) * (0.0 / float64((float32(1.5) / 0.1))))) - 0.0)
    v5 := []int{128, -1}
    v4 = v4
    v6 := GenS1{F0: ("a" + "\n\t"), F1: 0}
    var v7 int32 = 2147483647
    v8 := GenS1{F0: ("hello" + "\n\t"), F1: -1}
    fmt.Println(v4)
    switch ("a" + ("" + ("\n\t" + ("" + ("foo bar" + "foo bar"))))) {
    case "":
        var v9 uint8 = 42
        _ = v9
    case "foo bar":
        var v10 uint = 0
        _ = v10
    case "\n\t":
        if len(v5) > 0 {
            v5[0] = 1
        }
    }
    var v11 uint16 = 255
    for fi12 := 0; fi12 < 19; fi12++ {
        fmt.Println(v2)
        if true {
            switch "\n\t" {
            case "\n\t":
                if (1e10 <= 0.0) {
                    switch ("foo bar" + ("foo bar" + "hello")) {
                    case "":
                        fmt.Println(v11, v2, fi12)
                    case "foo bar":
                        v2 += (v2 ^ (v2 * 560))
                    case "\n\t":
                        for fi13 := 0; fi13 < 13; fi13++ {
                            switch ("\n\t" + (("foo bar" + "") + "foo bar")) {
                            case "\n\t":
                                v11 = v11
                            }
                            _ = fi13
                        }
                    }
                } else {
                    v14 := GenS2{F0: 0}
                    v15 := GenS1{F0: "hello", F1: 99}
                    _ = v14
                    _ = v15
                }
            case "a":
                for fi16 := 0; fi16 < 16; fi16++ {
                    if ((^(v7)) != 1) {
                        _ = len(v3)
                    } else {
                        v17 := GenS3{F0: 255, F1: 1000, F2: 0}
                        _ = v17
                    }
                    v18 := GenS4{F0: v2, F1: -514, F2: 32767}
                    v19 := GenS1{F0: ("foo bar" + ("" + "\n\t")), F1: -128}
                    _ = fi16
                    _ = v18
                    _ = v19
                }
            }
        } else {
            var v20 float32 = float32(((float64((float32((float64(1e10) + ((-2.25 + (((v4 / 1e10) * ((float64((-2.25 / -2.25)) - 1e10) / ((-2.25 / (v4 * 0.0)) / v4))) + 0.0)) / float64(-2.25)))) * ((1e10 + 0.0) - ((3.14159 * (float32(-2.25) - (1.5 / 0.1))) + 0.1)))) * v4) + v4))
            if true {
                v7 += (v7 | v7)
            }
            _ = v20
        }
        _ = fi12
    }
    var v21 int32 = (v7 | (-(v7)))
    fmt.Println(v7)
    delete(v3, "\n\t")
    v5 = append(v5, 1)
    fmt.Println(v11)
    switch "\n\t" {
    case "hello":
        switch int(v2) {
        case 42:
            for fi22 := 0; fi22 < 2; fi22++ {
                switch (fi22 & (fi22 >> ((1) & 63))) {
                case 2147483647:
                    var v23 uint = 255
                    _ = v23
                case 264:
                    for fi24 := 0; fi24 < 16; fi24++ {
                        v11 += v11
                        if false {
                            _ = len(v1)
                            var v25 uint64 = v2
                            _ = v25
                        }
                        if false {
                            v26 := GenS2{F0: -32768}
                            _ = v26
                        }
                        _ = fi24
                    }
                }
                _ = fi22
            }
        case 9223372036854775807:
            var v27 float64 = 1.5
            _ = v27
        case 32767:
            for fi28 := 0; fi28 < 8; fi28++ {
                var v29 int32 = (v21 + -44)
                switch fi28 {
                case 340:
                    var v30 uint8 = 233
                    _ = v30
                case 4294967295:
                    _ = len(v1)
                case 339:
                    v31 := GenS5{F0: v29}
                    _ = v31
                default:
                    v2 -= v2
                }
                _ = fi28
                _ = v29
            }
        default:
            var v32 int16 = 1000
            _ = v32
        }
    case "":
        for fi33 := 0; fi33 < 12; fi33++ {
            for fi34 := 0; fi34 < 9; fi34++ {
                var v35 string = ""
                _ = fi34
                _ = v35
            }
            switch 9223372036854775807 {
            case 127:
                v4 = (v4 * -2.25)
            case 2147483647:
                for fi36 := 0; fi36 < 7; fi36++ {
                    var v37 int64 = -1
                    v37 += -749
                    var v38 int64 = v37
                    _ = fi36
                    _ = v37
                    _ = v38
                }
            }
            v11 += (v11 - v11)
            _ = fi33
        }
    default:
        if (3.14159 != 0.0) {
            v2 += (v2 + v2)
            fmt.Println(v21)
            var v39 float32 = (-2.25 * float32(v4))
            _ = v39
        }
    }
    v21 = v7
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
