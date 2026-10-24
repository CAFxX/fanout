package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int16; F1 uint32 }
type GenS3 struct { F0 uint64; F1 uint64 }
type GenS4 struct { F0 int64; F1 uint }
type GenS5 struct { F0 uint8; F1 int32; F2 uint }
type GenS6 struct { F0 int }

func main() {
    v1 := []string{"hello", "hello", "foo bar", "a"}
    v2 := []uint32{}
    var v3 uint16 = 255
    var v4 int64 = 42
    var v5 bool = (102 > 127)
    v6 := GenS1{F0: 127}
    switch "a" {
    case "foo bar":
        v7 := GenS2{F0: -492, F1: 2147483647}
        _ = v7
    case "":
        v8 := GenS1{F0: 242}
        _ = v8
    case "a":
        switch 32767 {
        case -2147483648:
            for fi9 := 0; fi9 < 19; fi9++ {
                if v5 {
                    v10 := GenS3{F0: 42, F1: 4294967295}
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    _ = v10
                }
                v5 = v5
                _ = fi9
            }
        case 9223372036854775807:
            v1 = append(v1, "foo bar")
        case 2147483647:
            v11 := GenS4{F0: v4, F1: 4294967295}
            _ = v11
        }
    default:
        if len(v2) > 0 {
            v2[0] = 123
        }
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v5 = ((-2.25 * 0.0) == float32(float64((1e10 / ((float32(1e10) / (0.1 / (3.14159 / 1.0))) + 1.5)))))
    for fi12 := 0; fi12 < 3; fi12++ {
        var v13 int = (fi12 & fi12)
        for fi14 := 0; fi14 < 12; fi14++ {
            v4 += v4
            var v15 uint64 = 42
            var v16 uint32 = 469
            _ = fi14
            _ = v15
            _ = v16
        }
        fmt.Println(v4)
        _ = fi12
        _ = v13
    }
    v4 = v4
    v17 := GenS5{F0: 1, F1: -2147483648, F2: 2}
    v3 = 733
    if true {
        v3 = v3
    } else {
        var v18 float32 = 3.14159
        v19 := GenS6{F0: 0}
        v4 -= v4
        _ = v18
        _ = v19
    }
    v3 += v3
    var v20 string = "\n\t"
    v21 := GenS1{F0: 81}
    v22 := GenS3{F0: 4294967295, F1: 430}
    for fi23 := 0; fi23 < 14; fi23++ {
        for fi24 := 0; fi24 < 2; fi24++ {
            if true {
                v4 -= v4
            } else {
                if ((610 > 0) || (18446744073709551615 <= 181)) {
                    switch (v20 + "") {
                    case "\n\t":
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    case "":
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                    default:
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                    }
                    var v25 int32 = 0
                    _ = v25
                }
            }
            _ = fi24
        }
        var v26 float64 = 3.14159
        _ = fi23
        _ = v26
    }
    for fi27 := 0; fi27 < 20; fi27++ {
        fmt.Println(v3)
        var v28 int16 = 1
        _ = fi27
        _ = v28
    }
    var v29 bool = ((3.14159 + 0.0) == 0.0)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v17
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
