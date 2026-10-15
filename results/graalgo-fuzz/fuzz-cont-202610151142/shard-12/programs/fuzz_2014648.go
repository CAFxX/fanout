package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int16; F1 uint64; F2 uint64 }
type GenS3 struct { F0 int64 }
func genFn0() uint16 {
    var v4 int16 = -1
    fmt.Println(v4)
    v4 += (v4 + -32768)
    _ = v4
    return 387
}
func genFn1(p0 int64, p1 uint16, p2 uint8) int8 {
    v5 := []int{32768, -9223372036854775808, -791, 0}
    for fi6 := 0; fi6 < 12; fi6++ {
        p2 -= (p2 >> ((550) & 63))
        p0 = p0
        p2 = (p2 + p2)
        _ = fi6
    }
    for fi7 := 0; fi7 < 2; fi7++ {
        fmt.Println(p1, p2, p2)
        fmt.Println(p2, p2, p1)
        _ = fi7
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 0
}

func main() {
    var v1 string = ""
    var v2 uint16 = 933
    var v3 uint16 = (v2 + (v2 / (((v2 >> ((65535) & 63))) | 1)))
    var v8 int64 = 956
    fmt.Println(v1, v3, v8)
    for fi9 := 0; fi9 < 19; fi9++ {
        v3 = v2
        v10 := []int{fi9, (fi9 & fi9)}
        _ = fi9
        _ = v10
    }
    for fi11 := 0; fi11 < 11; fi11++ {
        fmt.Println(v8)
        _ = fi11
    }
    switch v1 {
    case "a":
        for fi12 := 0; fi12 < 15; fi12++ {
            v13 := GenS1{F0: 221}
            switch "foo bar" {
            case "a":
                var v14 uint8 = 1
                _ = v14
            case "":
                v15 := []int{(-(fi12)), fi12, (fi12 ^ 813)}
                _ = v15
            case "\n\t":
                v16 := []int{fi12, fi12, (fi12 / ((-9223372036854775808) | 1))}
                _ = v16
            }
            var v17 uint8 = 127
            _ = fi12
            _ = v13
            _ = v17
        }
    case "":
        fmt.Println(v8)
    case "hello":
        var v18 float64 = 1.5
        _ = v18
    }
    switch 128 {
    case 0:
        var v19 float32 = (-2.25 / -2.25)
        _ = v19
    case -9223372036854775808:
        v20 := GenS2{F0: -1, F1: 762, F2: 0}
        _ = v20
    }
    var v21 int64 = (v8 ^ (-(v8)))
    v22 := []int{-9223372036854775808, -2147483648, 9223372036854775807, -627}
    v22 = append(v22, 2147483647)
    if len(v22) >= 2 {
        v22 = v22[1:len(v22)]
    }
    fmt.Println(v3, v2)
    if false {
        for fi23 := 0; fi23 < 15; fi23++ {
            v24 := GenS3{F0: 1}
            _ = fi23
            _ = v24
        }
    }
    v25 := GenS3{F0: (v8 >> ((255) & 63))}
    fmt.Println(v3, v2)
    fmt.Println(v3, v3)
    var v26 string = v1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v25
    _ = v26
    _ = v3
    _ = v8
}
