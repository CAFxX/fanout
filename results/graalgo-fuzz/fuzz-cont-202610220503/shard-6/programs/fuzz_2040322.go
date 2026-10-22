package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 string; F2 uint }
type GenS2 struct { F0 int32; F1 int16 }

func main() {
    var v1 int32 = -513
    v2 := make(map[uint32]uint32)
    var v3 int32 = v1
    var v4 uint8 = 1
    v5 := GenS1{F0: (v4 << ((65535) & 63)), F1: "\n\t", F2: 605}
    var v6 string = "a"
    for fi7 := 0; fi7 < 10; fi7++ {
        if false {
            fmt.Println(v1, v4)
            for fi8 := 0; fi8 < 3; fi8++ {
                v9 := GenS2{F0: v1, F1: 42}
                v10 := GenS1{F0: (v4 + (v4 & 127)), F1: (("a" + v6) + v6), F2: 255}
                var v11 int16 = int16(v4)
                fmt.Println(v4, v4, v3)
                _ = fi8
                _ = v10
                _ = v11
                _ = v9
            }
            v12 := v2[uint32(v4)]
            _ = v12
        }
        var v13 uint = 65535
        _ = fi7
        _ = v13
    }
    v14 := []int{-9223372036854775808, 1}
    if false {
        var v15 uint32 = 721
        _ = v15
    }
    v16 := GenS2{F0: (^(v3)), F1: 216}
    if false {
        if (940 != 42) {
            switch v6 {
            case "":
                v6 = "\n\t"
            case "foo bar":
                v17 := GenS2{F0: v1, F1: 1}
                _ = v17
            default:
                v1 -= -2147483648
            }
            var v18 int8 = -128
            switch v6 {
            case "":
                var v19 int8 = (v18 - (v18 / ((v18) | 1)))
                _ = v19
            case "foo bar":
                v4 += 42
            }
            _ = v18
        }
    }
    v20 := GenS2{F0: (v1 % ((v1) | 1)), F1: 191}
    for fi21 := 0; fi21 < 16; fi21++ {
        v3 -= (v1 & 39)
        _ = fi21
    }
    switch -9223372036854775808 {
    case 4294967295:
        if len(v14) > 0 {
            v14[0] = 32767
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v16
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
