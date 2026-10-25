package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint64; F2 int }
type GenS2 struct { F0 uint; F1 string; F2 uint64 }
type GenS3 struct { F0 int; F1 int64; F2 uint64 }

func main() {
    var v1 int16 = 1000
    v2 := []string{"\n\t"}
    v3 := make(map[uint32]bool)
    var v4 float32 = (0.0 * (1.5 * 0.1))
    for fi5 := 0; fi5 < 6; fi5++ {
        for fi6 := 0; fi6 < 20; fi6++ {
            if (v4 != v4) {
                var v7 bool = true
                if false {
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                } else {
                    v8 := GenS1{F0: -66, F1: 329, F2: (fi5 * fi5)}
                    v9 := v3[0]
                    var v10 float64 = -2.25
                    _ = v10
                    _ = v8
                    _ = v9
                }
                _ = v7
            }
            var v11 int32 = 42
            var v12 uint = 255
            _ = fi6
            _ = v11
            _ = v12
        }
        v4 = (v4 / 1.0)
        _ = fi5
    }
    var v13 float64 = (0.0 / 3.14159)
    var v14 int8 = 4
    v15 := GenS2{F0: 65535, F1: "", F2: 2}
    for fi16 := 0; fi16 < 16; fi16++ {
        if len(v2) > 0 {
            v2[0] = "foo bar"
        }
        switch 1 {
        case -2147483648:
            _ = len(v3)
        case 138:
            v17 := GenS3{F0: fi16, F1: 42, F2: 42}
            _ = v17
        }
        _ = fi16
    }
    if true {
        switch "foo bar" {
        case "hello":
            v18 := GenS2{F0: 1, F1: ("a" + ("\n\t" + "hello")), F2: 2}
            _ = v18
        case "":
            v4 += v4
        case "foo bar":
            if false {
                switch "\n\t" {
                case "":
                    var v19 int8 = (-(v14))
                    _ = v19
                case "\n\t":
                    for fi20 := 0; fi20 < 18; fi20++ {
                        var v21 int32 = -125
                        var v22 float32 = (3.14159 / (0.0 * v4))
                        _ = fi20
                        _ = v21
                        _ = v22
                    }
                case "a":
                    var v23 uint8 = 102
                    _ = v23
                }
                if (!true) {
                    fmt.Println(v4)
                    fmt.Println(v1)
                }
            } else {
                var v24 uint8 = 128
                v25 := GenS3{F0: 706, F1: 582, F2: 18446744073709551615}
                _ = v24
                _ = v25
            }
        }
        switch 9223372036854775807 {
        case 1:
            var v26 float32 = (0.0 * v4)
            _ = v26
        default:
            var v27 string = ("foo bar" + ("" + (("foo bar" + "hello") + "\n\t")))
            _ = v27
        }
    } else {
        var v28 int64 = 1
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        _ = v28
    }
    for fi29 := 0; fi29 < 11; fi29++ {
        var v30 float32 = v4
        _ = fi29
        _ = v30
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
}
