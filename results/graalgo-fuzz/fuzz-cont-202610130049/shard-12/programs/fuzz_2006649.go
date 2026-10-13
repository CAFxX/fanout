package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint64; F1 int32 }
type GenS4 struct { F0 int64; F1 uint8; F2 int8 }
type GenS5 struct { F0 int64; F1 int32 }

func main() {
    v1 := make(map[int]string)
    var v2 uint64 = 4294967295
    var v3 int16 = -32768
    var v4 uint32 = 4294967295
    for fi5 := 0; fi5 < 16; fi5++ {
        v6 := GenS1{F0: 42, F1: 4294967295}
        fmt.Println(v4)
        for fi7 := 0; fi7 < 7; fi7++ {
            var v8 uint8 = 255
            switch (fi5 | (fi5 << ((65535) & 63))) {
            case -1:
                var v9 float32 = (float32(float64(((-2.25 + 0.0) / 3.14159))) * 1.5)
                _ = v9
            case 4294967295:
                if false {
                    v10 := []int{4294967295, 9223372036854775807, -178, fi5}
                    v2 -= (v2 / ((v2) | 1))
                    v3 = v3
                    _ = v10
                }
            case 42:
                if false {
                    var v11 int32 = int32(v8)
                    _ = v11
                } else {
                    v12 := GenS2{F0: (2 >= (v8 | (^(v8))))}
                    fmt.Println(v3, fi7)
                    v2 -= 1
                    _ = v12
                }
            }
            fmt.Println(v2, v8)
            v4 = v4
            _ = fi7
            _ = v8
        }
        _ = fi5
        _ = v6
    }
    v13 := GenS3{F0: v2, F1: 42}
    var v14 float64 = (0.1 * 3.14159)
    if false {
        if (!false) {
            var v15 bool = ((^(v4)) < v4)
            for fi16 := 0; fi16 < 15; fi16++ {
                v17 := GenS4{F0: 1, F1: 111, F2: 42}
                if (128 <= 1) {
                    v15 = (!v15)
                } else {
                    v18 := v1[fi16]
                    for fi19 := 0; fi19 < 11; fi19++ {
                        var v20 int = fi16
                        v18 = ((v18 + "foo bar") + "")
                        _ = fi19
                        _ = v20
                    }
                    _ = v18
                }
                _ = fi16
                _ = v17
            }
            var v21 int8 = 0
            _ = v15
            _ = v21
        }
        v22 := []int{-1}
        var v23 float64 = float64(float32(((v14 / (v14 + (v14 * v14))) - v14)))
        _ = v22
        _ = v23
    }
    var v24 uint64 = v2
    var v25 int32 = -2147483648
    v3 -= (-(v3))
    var v26 int32 = (v25 % ((-1) | 1))
    v27 := GenS2{F0: (100 != 93)}
    fmt.Println(v4, v2, v26)
    if true {
        switch "a" {
        case "\n\t":
            fmt.Println(v3)
        case "a":
            v28 := []int{}
            _ = v28
        case "":
            v29 := GenS4{F0: -9223372036854775808, F1: 115, F2: -1}
            _ = v29
        }
    }
    v30 := GenS3{F0: 146, F1: 0}
    v31 := GenS5{F0: -1, F1: v26}
    v32 := GenS4{F0: -30, F1: 147, F2: 42}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
}
