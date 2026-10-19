package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int64; F1 string }
func genFn0(p0 uint64) int16 {
    var v5 uint8 = 0
    fmt.Println(v5)
    v5 += 0
    for fi6 := 0; fi6 < 8; fi6++ {
        v7 := []int{}
        v5 -= v5
        _ = fi6
        _ = v7
    }
    v5 = v5
    _ = p0
    _ = v5
    return int16(p0)
}
func genFn1(p0 int8, p1 uint, p2 int) uint {
    var v8 uint8 = 237
    p0 = 1
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return (p1 % (((p1 % ((312) | 1))) | 1))
}

func main() {
    var v1 string = "a"
    v2 := make(map[uint32]string)
    var v3 int8 = -76
    v4 := make(map[int]uint32)
    v9 := []int{}
    switch "\n\t" {
    case "":
        if true {
            v10 := GenS1{F0: v3}
            v1 = ((v1 + (v1 + "")) + ("foo bar" + "hello"))
            v11 := v2[1]
            _ = v10
            _ = v11
        } else {
            if len(v9) > 0 {
                v9[0] = 9223372036854775807
            }
        }
    }
    v12 := v2[0]
    fmt.Println(v1, v12, v3)
    v3 += v3
    _ = len(v9)
    v1 = v12
    switch v1 {
    case "hello":
        var v13 float32 = 0.1
        _ = v13
    }
    v14 := GenS2{F0: 485, F1: "foo bar"}
    switch (v1 + "") {
    case "\n\t":
        fmt.Println(v1, v1, v1)
    case "foo bar":
        v3 += (v3 % (((v3 << ((2) & 63))) | 1))
    case "a":
        if len(v9) > 0 {
            v9[0] = 9223372036854775807
        }
    }
    for fi15 := 0; fi15 < 7; fi15++ {
        v1 = "a"
        v16 := GenS1{F0: (^(v3))}
        switch "a" {
        case "":
            var v17 float64 = (-2.25 / (float64(3.14159) / float64((0.1 + ((3.14159 - -2.25) / float32(-2.25))))))
            _ = v17
        case "hello":
            for fi18 := 0; fi18 < 15; fi18++ {
                v3 += v3
                v12 = v1
                _ = fi18
            }
        default:
            var v19 int8 = (v3 & 100)
            _ = v19
        }
        _ = fi15
        _ = v16
    }
    fmt.Println(v1)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v12
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
