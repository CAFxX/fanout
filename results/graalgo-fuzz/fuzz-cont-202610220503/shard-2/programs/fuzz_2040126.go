package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint64; F2 int }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint; F1 uint }

func main() {
    var v1 int16 = -358
    v2 := []uint32{1, 4294967295}
    var v3 int8 = 1
    if len(v2) > 0 {
        v2[0] = 56
    }
    var v4 float64 = 0.0
    var v5 uint = 4294967295
    var v6 int = 9223372036854775807
    if (!true) {
        for fi7 := 0; fi7 < 3; fi7++ {
            v8 := GenS1{F0: (v1 << (((^(v5))) & 63)), F1: 18446744073709551615, F2: (fi7 ^ (v6 - (fi7 % (((v6 * (^(v6)))) | 1))))}
            _ = fi7
            _ = v8
        }
        v3 = v3
    }
    v9 := GenS1{F0: (v1 / (((v1 * 776)) | 1)), F1: 42, F2: v6}
    v5 += (v5 | v5)
    v1 = v1
    v10 := GenS2{F0: 511}
    for fi11 := 0; fi11 < 13; fi11++ {
        v5 += (v5 / ((v5) | 1))
        _ = fi11
    }
    v4 += 3.14159
    var v12 uint = 255
    var v13 int16 = 42
    for fi14 := 0; fi14 < 4; fi14++ {
        v3 -= (v3 % (((v3 + (v3 / ((-128) | 1)))) | 1))
        _ = fi14
    }
    fmt.Println(v1, v6)
    var v15 string = "\n\t"
    switch (-(v6)) {
    case -9223372036854775808:
        switch (v6 / ((v6) | 1)) {
        case 4294967295:
            for fi16 := 0; fi16 < 20; fi16++ {
                var v17 float32 = (3.14159 + (0.0 + (0.1 + 1e10)))
                _ = fi16
                _ = v17
            }
        }
    default:
        switch "a" {
        case "":
            fmt.Println(v4, v5, v1)
        case "hello":
            var v18 uint = (v5 - v5)
            _ = v18
        case "foo bar":
            v13 = int16(v3)
        default:
            v4 = (v4 * 3.14159)
        }
    }
    v19 := GenS3{F0: (v12 >> (((v5 * v12)) & 63)), F1: (v5 * v5)}
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v12
    _ = v13
    _ = v15
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
