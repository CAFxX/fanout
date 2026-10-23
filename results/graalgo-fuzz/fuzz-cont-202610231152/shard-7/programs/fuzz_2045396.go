package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 string; F1 uint64; F2 bool }
type GenS3 struct { F0 uint; F1 int64 }
type GenS4 struct { F0 int8; F1 int32 }
type GenS5 struct { F0 uint32; F1 uint64; F2 uint16 }

func main() {
    v1 := make(map[int]bool)
    var v2 int32 = 0
    v3 := make(map[string]string)
    _ = len(v3)
    if (false || false) {
        var v4 uint8 = 1
        var v5 float32 = (0.1 - 1e10)
        for fi6 := 0; fi6 < 10; fi6++ {
            var v7 string = ("\n\t" + ("foo bar" + "\n\t"))
            _ = fi6
            _ = v7
        }
        _ = v4
        _ = v5
    }
    if true {
        v8 := []int{}
        _ = v8
    } else {
        fmt.Println(v2)
        if ((-2.25 - 3.14159) == float64(float32(0.1))) {
            v2 = (-(v2))
            for fi9 := 0; fi9 < 11; fi9++ {
                v10 := GenS1{F0: -1}
                _ = fi9
                _ = v10
            }
            v11 := GenS2{F0: "\n\t", F1: 4294967295, F2: (false && true)}
            _ = v11
        }
        for fi12 := 0; fi12 < 14; fi12++ {
            v13 := GenS2{F0: "hello", F1: 582, F2: ((1.5 * -2.25) <= (0.0 / -2.25))}
            var v14 uint64 = 2
            v14 -= 18446744073709551615
            _ = fi12
            _ = v13
            _ = v14
        }
    }
    v15 := GenS3{F0: 65535, F1: 633}
    fmt.Println(v2)
    var v16 int = 128
    if (!false) {
        v17 := GenS3{F0: 2, F1: -9223372036854775808}
        _ = v17
    }
    if (!true) {
        fmt.Println(v16)
        for fi18 := 0; fi18 < 7; fi18++ {
            if (128 > 128) {
                v16 -= (^(fi18))
                v19 := GenS4{F0: 42, F1: v2}
                v16 = 0
                _ = v19
            }
            var v20 int32 = v2
            fmt.Println(v2, fi18, v16)
            _ = fi18
            _ = v20
        }
    } else {
        switch (v16 % ((0) | 1)) {
        case 2147483647:
            fmt.Println(v16, v2)
        case 128:
            v21 := GenS5{F0: 625, F1: 150, F2: 42}
            _ = v21
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
}
