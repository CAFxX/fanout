package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int64; F1 uint64; F2 uint }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 string; F1 int64; F2 int }

func main() {
    var v1 float64 = 1e10
    v2 := make(map[uint32]string)
    v3 := make(map[string]uint32)
    v4 := make(map[uint32]int)
    v5 := []int{933, 2147483647, 32768, 0}
    var v6 int = -471
    v6 += -156
    if (1 != -128) {
        v7 := GenS1{F0: -82}
        _ = v7
    } else {
        v6 = 9223372036854775807
        _ = len(v3)
        v6 += (v6 | 42)
    }
    v8 := GenS1{F0: 127}
    v6 = (v6 | (v6 ^ (v6 + v6)))
    for fi9 := 0; fi9 < 18; fi9++ {
        _ = len(v2)
        var v10 int32 = 501
        var v11 uint32 = 0
        _ = fi9
        _ = v10
        _ = v11
    }
    delete(v4, 2147483647)
    v12 := GenS2{F0: -766}
    v1 += 0.1
    v6 -= (v6 * v6)
    switch ("foo bar" + ("foo bar" + ("a" + "a"))) {
    case "\n\t":
        var v13 int = (v6 & (v6 ^ 240))
        _ = v13
    case "a":
        v6 = (-(v6))
    case "hello":
        v6 += (v6 << ((255) & 63))
    }
    switch v6 {
    case -82:
        v14 := GenS3{F0: -1, F1: 1, F2: 4294967295}
        _ = v14
    case 42:
        for fi15 := 0; fi15 < 20; fi15++ {
            v16 := GenS4{F0: 488}
            _ = fi15
            _ = v16
        }
    case 966:
        var v17 string = "\n\t"
        _ = v17
    default:
        v18 := GenS5{F0: "hello", F1: -9223372036854775808, F2: v6}
        _ = v18
    }
    for fi19 := 0; fi19 < 17; fi19++ {
        v2[2147483647] = ("\n\t" + ("a" + "\n\t"))
        _ = fi19
    }
    v1 += float64(0.1)
    switch "\n\t" {
    case "a":
        var v20 int64 = -416
        _ = v20
    case "hello":
        for fi21 := 0; fi21 < 10; fi21++ {
            var v22 int = 0
            if len(v5) > 0 {
                v5[0] = 0
            }
            _ = fi21
            _ = v22
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
