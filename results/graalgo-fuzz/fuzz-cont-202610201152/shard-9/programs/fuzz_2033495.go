package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int64; F2 int16 }
type GenS2 struct { F0 uint; F1 string; F2 int }
func genFn0(p0 uint16, p1 uint8, p2 string) int {
    for fi4 := 0; fi4 < 3; fi4++ {
        for fi5 := 0; fi5 < 17; fi5++ {
            p1 = (p1 % ((p1) | 1))
            switch "" {
            case "hello":
                v6 := GenS1{F0: p2, F1: 9223372036854775807, F2: -32768}
                _ = v6
            case "a":
                fmt.Println(fi4, fi4, p1)
            }
            _ = fi5
        }
        _ = fi4
    }
    fmt.Println(p0)
    for fi7 := 0; fi7 < 13; fi7++ {
        var v8 int32 = 2147483647
        v9 := []int{}
        v10 := GenS1{F0: p2, F1: 9223372036854775807, F2: 42}
        _ = fi7
        _ = v10
        _ = v8
        _ = v9
    }
    fmt.Println(p0)
    _ = p0
    _ = p1
    _ = p2
    return 127
}

func main() {
    var v1 int16 = -32768
    v2 := make(map[string]int)
    var v3 float64 = 1.5
    var v11 float64 = (v3 - v3)
    v1 -= (v1 << ((446) & 63))
    v12 := v2["foo bar"]
    v13 := v2[""]
    if (false || true) {
        switch "\n\t" {
        case "a":
            _ = len(v2)
        }
        v14 := v2["\n\t"]
        var v15 uint32 = 2
        _ = v14
        _ = v15
    } else {
        v12 = (v13 & (v12 - v13))
        v16 := GenS1{F0: "hello", F1: 9223372036854775807, F2: (v1 % ((v1) | 1))}
        _ = v16
    }
    var v17 uint64 = 1
    v18 := GenS2{F0: 0, F1: "", F2: (v12 / ((v12) | 1))}
    fmt.Println(v12)
    v19 := GenS1{F0: "", F1: -9223372036854775808, F2: -580}
    for fi20 := 0; fi20 < 12; fi20++ {
        v3 += (float64((1.5 * float32(((v3 * (v11 / float64((((((1e10 + float32(((0.1 + 0.0) - float64((0.0 * (-2.25 + 1.5)))))) / 1e10) / 1e10) - 3.14159) - (0.1 - (3.14159 + 3.14159)))))) / v3)))) * v3)
        for fi21 := 0; fi21 < 2; fi21++ {
            var v22 uint64 = v17
            _ = fi21
            _ = v22
        }
        v13 -= v13
        _ = fi20
    }
    fmt.Println(v13, v11)
    var v23 bool = (!(-2.25 != -2.25))
    switch "a" {
    case "hello":
        for fi24 := 0; fi24 < 8; fi24++ {
            v11 -= ((0.1 + float64(-2.25)) + 0.0)
            v25 := []int{-403}
            _ = fi24
            _ = v25
        }
    case "\n\t":
        v26 := []int{(v13 | v12), genFn0(65535, uint8(v17), "foo bar")}
        _ = v26
    case "foo bar":
        var v27 uint64 = (v17 << ((705) & 63))
        _ = v27
    }
    v1 = (v1 * (v1 >> ((1) & 63)))
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v23
    _ = v3
}
