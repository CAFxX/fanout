package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int }
type GenS2 struct { F0 uint16; F1 uint8 }
type GenS3 struct { F0 int64; F1 int; F2 int }

func main() {
    v1 := make(map[uint32]int)
    v2 := []uint32{935, 41}
    v3 := []int{22, -9223372036854775808, 32768, 787}
    var v4 int16 = 1000
    var v5 uint = 2
    v6 := make(map[int]int)
    switch (("hello" + "a") + "foo bar") {
    case "foo bar":
        var v7 int16 = v4
        _ = v7
    case "\n\t":
        v5 += v5
    case "":
        var v8 float64 = 1e10
        _ = v8
    default:
        fmt.Println(v5)
    }
    v5 -= (v5 >> (((v5 >> ((320) & 63))) & 63))
    if (!(!false)) {
        for fi9 := 0; fi9 < 11; fi9++ {
            var v10 uint32 = 0
            for fi11 := 0; fi11 < 4; fi11++ {
                v12 := GenS1{F0: v5, F1: fi11}
                _ = fi11
                _ = v12
            }
            _ = fi9
            _ = v10
        }
        if true {
            v5 = 0
            if true {
                var v13 int8 = 42
                var v14 int8 = 100
                _ = v13
                _ = v14
            }
            for fi15 := 0; fi15 < 13; fi15++ {
                for fi16 := 0; fi16 < 4; fi16++ {
                    if (!(1 == 890)) {
                        v17 := GenS1{F0: (v5 - v5), F1: 9223372036854775807}
                        _ = v17
                    }
                    _ = fi16
                }
                switch (^(fi15)) {
                case 4294967295:
                    var v18 uint64 = 689
                    _ = v18
                case 32768:
                    v1[19] = 0
                default:
                    v4 -= v4
                }
                v4 = (v4 * 85)
                _ = fi15
            }
        }
    }
    switch (("\n\t" + "\n\t") + "a") {
    case "hello":
        _ = len(v1)
    }
    fmt.Println(v5, v4)
    if true {
        for fi19 := 0; fi19 < 8; fi19++ {
            var v20 int16 = (v4 | 0)
            for fi21 := 0; fi21 < 14; fi21++ {
                fmt.Println(fi19)
                var v22 int16 = 0
                _ = fi21
                _ = v22
            }
            var v23 float64 = 3.14159
            _ = fi19
            _ = v20
            _ = v23
        }
        v4 += -50
    }
    v4 -= v4
    switch "hello" {
    case "hello":
        v24 := GenS2{F0: 65535, F1: 2}
        _ = v24
    case "foo bar":
        var v25 float32 = float32((3.14159 * 1.5))
        _ = v25
    case "":
        for fi26 := 0; fi26 < 12; fi26++ {
            v27 := GenS3{F0: 42, F1: fi26, F2: fi26}
            v4 += v4
            _ = fi26
            _ = v27
        }
    }
    v4 = (v4 ^ v4)
    v5 += (^(v5))
    v4 += (v4 / ((-76) | 1))
    var v28 bool = (false || (-70 >= 0))
    v28 = ((v28 && true) || false)
    var v29 float32 = -2.25
    if v28 {
        switch "a" {
        case "hello":
            if v28 {
                fmt.Println(v4)
            }
        default:
            v28 = false
        }
        v29 = (float32((0.1 - 1.5)) - v29)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
