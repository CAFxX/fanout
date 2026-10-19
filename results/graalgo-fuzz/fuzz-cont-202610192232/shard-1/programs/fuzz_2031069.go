package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8 }
type GenS2 struct { F0 uint8; F1 int32; F2 int32 }
type GenS3 struct { F0 int16; F1 bool; F2 uint32 }
type GenS4 struct { F0 int32; F1 int }

func main() {
    var v1 uint8 = 128
    var v2 int = 42
    var v3 int64 = 9223372036854775807
    v4 := make(map[int]uint32)
    v5 := make(map[int]int)
    for fi6 := 0; fi6 < 2; fi6++ {
        delete(v4, fi6)
        var v7 int = (^(v2))
        _ = fi6
        _ = v7
    }
    switch (("\n\t" + "\n\t") + "foo bar") {
    case "foo bar":
        if ((float64(3.14159) * float64(3.14159)) != (0.1 + (1e10 / 1.0))) {
            var v8 uint64 = 4294967295
            _ = len(v5)
            _ = v8
        } else {
            v2 = (v2 | -928)
            var v9 int8 = 42
            v10 := []int{}
            _ = v10
            _ = v9
        }
    case "\n\t":
        v11 := GenS1{F0: 18446744073709551615, F1: 255}
        _ = v11
    }
    for fi12 := 0; fi12 < 15; fi12++ {
        for fi13 := 0; fi13 < 20; fi13++ {
            v14 := GenS2{F0: v1, F1: 1, F2: -537}
            _ = len(v4)
            if false {
                var v15 int32 = 0
                _ = v15
            } else {
                var v16 bool = false
                _ = v16
            }
            _ = fi13
            _ = v14
        }
        if true {
            var v17 float32 = 3.14159
            if (!(42 > 127)) {
                var v18 float32 = v17
                for fi19 := 0; fi19 < 17; fi19++ {
                    v20 := GenS3{F0: 10, F1: false, F2: 65535}
                    _ = fi19
                    _ = v20
                }
                _ = v18
            } else {
                var v21 int64 = (^(v3))
                _ = v21
            }
            _ = v17
        }
        _ = fi12
    }
    v2 += v2
    switch "" {
    case "foo bar":
        for fi22 := 0; fi22 < 13; fi22++ {
            switch "\n\t" {
            case "hello":
                v23 := GenS1{F0: 550, F1: v1}
                _ = v23
            default:
                var v24 int16 = 1000
                _ = v24
            }
            _ = fi22
        }
    case "\n\t":
        var v25 uint = 726
        _ = v25
    default:
        v26 := GenS4{F0: -2147483648, F1: v2}
        _ = v26
    }
    v1 = 255
    switch (v2 >> ((65535) & 63)) {
    case 32768:
        var v27 string = (("" + "") + "")
        _ = v27
    }
    v28 := []int{(v2 ^ (v2 % ((0) | 1))), v2}
    if len(v28) >= 2 {
        v28 = v28[1:len(v28)]
    }
    v2 -= (v2 % ((v2) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v28:", len(v28), v28)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
