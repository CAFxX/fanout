package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16; F2 uint32 }
type GenS2 struct { F0 uint8; F1 int64; F2 int16 }
type GenS3 struct { F0 int }

func main() {
    v1 := []string{"a", ((("\n\t" + "hello") + (("" + "foo bar") + ("a" + "foo bar"))) + "")}
    v2 := make(map[int]string)
    var v3 bool = (1 < 0)
    var v4 int16 = -446
    var v5 uint8 = 72
    var v6 int16 = (v4 & -121)
    for fi7 := 0; fi7 < 4; fi7++ {
        for fi8 := 0; fi8 < 8; fi8++ {
            v6 = v4
            var v9 int = fi7
            v10 := v2[fi7]
            _ = fi8
            _ = v10
            _ = v9
        }
        _ = fi7
    }
    v6 = v4
    var v11 int64 = -1
    for fi12 := 0; fi12 < 1; fi12++ {
        var v13 int16 = v6
        _ = fi12
        _ = v13
    }
    v6 -= (v4 & (v4 >> ((18446744073709551615) & 63)))
    var v14 uint64 = 189
    v15 := GenS1{F0: 0, F1: 65535, F2: 623}
    v16 := GenS2{F0: uint8(v14), F1: v11, F2: (v6 + v4)}
    v6 = (^(v4))
    v17 := GenS3{F0: -1}
    fmt.Println(v14)
    var v18 int64 = (v11 + 75)
    v19 := GenS3{F0: 471}
    for fi20 := 0; fi20 < 15; fi20++ {
        if v3 {
            switch 9223372036854775807 {
            case -529:
                var v21 uint = 588
                _ = v21
            }
        } else {
            for fi22 := 0; fi22 < 16; fi22++ {
                v23 := v2[(fi20 << ((1) & 63))]
                var v24 int = 127
                _ = fi22
                _ = v23
                _ = v24
            }
            var v25 int32 = 0
            for fi26 := 0; fi26 < 16; fi26++ {
                v27 := GenS1{F0: 100, F1: 0, F2: 65535}
                v5 += v5
                v4 = v6
                _ = fi26
                _ = v27
            }
            _ = v25
        }
        var v28 float32 = 0.1
        fmt.Println(v14, v28, v5)
        _ = fi20
        _ = v28
    }
    fmt.Println(v11, v3, v3)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
