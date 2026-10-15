package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint8; F1 uint8; F2 bool }
type GenS3 struct { F0 bool; F1 int }

func main() {
    var v1 uint32 = 658
    var v2 int = 441
    v3 := make(map[uint32]string)
    var v4 float32 = float32((-2.25 / ((1.5 + 3.14159) + -2.25)))
    v5 := []int{(v2 | 591), v2, 9223372036854775807, (v2 | (^(v2)))}
    v4 = v4
    switch -923 {
    case -9223372036854775808:
        v6 := GenS1{F0: 42}
        _ = v6
    case 2147483647:
        if (false || false) {
            if ((v1 ^ (v1 >> ((42) & 63))) >= (v1 % ((v1) | 1))) {
                var v7 int32 = -812
                fmt.Println(v1)
                if false {
                    fmt.Println(v4, v2, v7)
                }
                _ = v7
            }
        }
    case 9223372036854775807:
        v8 := GenS1{F0: 42}
        _ = v8
    default:
        v4 = 3.14159
    }
    for fi9 := 0; fi9 < 14; fi9++ {
        switch (fi9 - -1) {
        case 128:
            switch (((((("" + "foo bar") + "") + "foo bar") + "\n\t") + "") + "a") {
            case "a":
                var v10 uint32 = (v1 % ((768) | 1))
                _ = v10
            }
        case 4294967295:
            v1 += v1
        case 9223372036854775807:
            v11 := GenS2{F0: 2, F1: 2, F2: ((0.0 * -2.25) >= ((-2.25 * (float64(0.0) + 1e10)) / 3.14159))}
            _ = v11
        }
        for fi12 := 0; fi12 < 1; fi12++ {
            fmt.Println(v2, v2)
            delete(v3, 42)
            _ = fi12
        }
        _ = fi9
    }
    v4 = (v4 + v4)
    v2 = 458
    var v13 float32 = v4
    switch "" {
    case "":
        v13 = 0.1
    case "a":
        if (("" + "") <= (((("a" + "foo bar") + "hello") + "hello") + "\n\t")) {
            v14 := GenS3{F0: false, F1: -2147483648}
            _ = v14
        }
    default:
        v15 := v3[v1]
        _ = v15
    }
    v4 = ((v4 / v13) * v13)
    var v16 uint = 353
    v17 := GenS2{F0: 255, F1: 0, F2: (573 == 0)}
    _ = len(v3)
    if true {
        v2 += (v2 & (v2 | v2))
        switch v2 {
        case -856:
            var v18 uint32 = 0
            _ = v18
        default:
            for fi19 := 0; fi19 < 16; fi19++ {
                v16 = (v16 + (v16 << ((v16) & 63)))
                _ = fi19
            }
        }
    }
    fmt.Println(v2, v4, v2)
    var v20 int = v2
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v13
    _ = v16
    _ = v17
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
