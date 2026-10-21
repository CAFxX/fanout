package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint16 }
type GenS2 struct { F0 int64; F1 uint8; F2 int16 }
type GenS3 struct { F0 int16 }
func genFn0(p0 uint16, p1 uint64) uint32 {
    var v5 float64 = -2.25
    fmt.Println(p0, v5, p0)
    if false {
        v6 := GenS1{F0: -9223372036854775808, F1: p0}
        _ = v6
    } else {
        p1 += (^(p1))
    }
    v7 := []int{32768, 9223372036854775807, 724}
    for fi8 := 0; fi8 < 11; fi8++ {
        var v9 float32 = ((0.1 + 0.1) * (1.5 / 1e10))
        p1 -= p1
        _ = fi8
        _ = v9
    }
    _ = p0
    _ = p1
    _ = v5
    _ = v7
    return uint32(p0)
}

func main() {
    var v1 int64 = -759
    v2 := make(map[uint32]uint32)
    var v3 int = 127
    v4 := make(map[uint32]string)
    v10 := GenS2{F0: v1, F1: 1, F2: 1}
    var v11 string = ("foo bar" + ("" + "foo bar"))
    if false {
        v3 -= (v3 + 9223372036854775807)
    }
    v1 += v1
    switch v11 {
    case "a":
        fmt.Println(v1, v1)
    case "":
        if true {
            switch (v11 + v11) {
            case "a":
                v12 := GenS1{F0: (v1 & v1), F1: 42}
                _ = v12
            case "\n\t":
                v13 := []int{4294967295, (v3 * v3)}
                _ = v13
            default:
                v14 := []int{9223372036854775807}
                _ = v14
            }
            v11 = ("foo bar" + "\n\t")
            fmt.Println(v11)
        } else {
            var v15 int = (v3 / (((v3 >> ((42) & 63))) | 1))
            var v16 string = v11
            var v17 float32 = (-2.25 + (0.1 * 1e10))
            _ = v15
            _ = v16
            _ = v17
        }
    case "hello":
        if (((-(v1)) <= (v1 | (v1 - v1))) && (!(!false))) {
            for fi18 := 0; fi18 < 19; fi18++ {
                fmt.Println(v11, v1, v11)
                switch fi18 {
                case -1:
                    fmt.Println(v3, v11, fi18)
                default:
                    v19 := GenS2{F0: -9223372036854775808, F1: 127, F2: -32768}
                    _ = v19
                }
                _ = fi18
            }
        }
    }
    switch v11 {
    case "\n\t":
        for fi20 := 0; fi20 < 3; fi20++ {
            if (3.14159 < 1e10) {
                if (false && false) {
                    v3 += (fi20 + (^(fi20)))
                } else {
                    for fi21 := 0; fi21 < 9; fi21++ {
                        v4[0] = (v11 + "\n\t")
                        var v22 float64 = (-2.25 - float64((1.5 + 0.0)))
                        fmt.Println(v3)
                        _ = fi21
                        _ = v22
                    }
                    fmt.Println(v11)
                }
                var v23 float32 = 1e10
                v24 := GenS3{F0: 960}
                _ = v23
                _ = v24
            } else {
                v11 = (("foo bar" + "hello") + v11)
            }
            _ = fi20
        }
    }
    var v25 uint = 185
    if ((v3 * 9223372036854775807) != (v3 ^ (v3 / ((356) | 1)))) {
        fmt.Println(v11, v25, v25)
        v1 -= v1
    }
    switch v3 {
    case -9223372036854775808:
        switch (("foo bar" + v11) + "a") {
        case "hello":
            v26 := []int{(v3 + v3), (v3 * v3)}
            _ = v26
        case "\n\t":
            v25 -= (v25 ^ v25)
        case "foo bar":
            var v27 uint16 = 1
            _ = v27
        }
    case -2147483648:
        fmt.Println(v3, v11, v3)
    case 42:
        v3 -= v3
    default:
        v28 := GenS1{F0: v1, F1: 32767}
        _ = v28
    }
    fmt.Println(v11, v11, v1)
    if true {
        if true {
            v29 := []int{(v3 >> (((v25 % ((42) | 1))) & 63)), v3}
            var v30 uint64 = 94
            _ = v29
            _ = v30
        }
        v31 := []int{}
        var v32 float64 = float64(3.14159)
        _ = v31
        _ = v32
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v25
    _ = v3
    _ = v4
}
