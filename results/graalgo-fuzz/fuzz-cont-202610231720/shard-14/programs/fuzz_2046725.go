package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int16 }
type GenS2 struct { F0 int16; F1 string; F2 int }
type GenS3 struct { F0 bool; F1 int8; F2 int64 }
type GenS4 struct { F0 uint; F1 uint16 }

func main() {
    var v1 float32 = (-2.25 * float32(-2.25))
    var v2 uint = 0
    v3 := make(map[uint32]uint32)
    var v4 string = "hello"
    for fi5 := 0; fi5 < 10; fi5++ {
        v6 := GenS1{F0: -2147483648, F1: 1000}
        v7 := []int{-2147483648, (fi5 % ((fi5) | 1)), (^(fi5)), fi5}
        switch "" {
        case "foo bar":
            if (2147483647 < 2147483647) {
                v7 = append(v7, 2147483647)
            }
        case "":
            v8 := GenS1{F0: 2147483647, F1: -291}
            _ = v8
        default:
            var v9 uint = 2
            _ = v9
        }
        _ = fi5
        _ = v6
        _ = v7
    }
    v2 = 4294967295
    switch v4 {
    case "a":
        v2 = v2
    }
    var v10 int64 = 14
    v4 = v4
    v11 := GenS1{F0: 381, F1: -114}
    for fi12 := 0; fi12 < 6; fi12++ {
        switch (fi12 ^ (fi12 * fi12)) {
        case 4294967295:
            v13 := []int{fi12, (fi12 ^ 1)}
            _ = v13
        }
        v14 := GenS2{F0: 1000, F1: v4, F2: (^(fi12))}
        _ = fi12
        _ = v14
    }
    v4 = "foo bar"
    for fi15 := 0; fi15 < 7; fi15++ {
        for fi16 := 0; fi16 < 17; fi16++ {
            v17 := GenS3{F0: (((0.0 == (3.14159 * v1)) || false) || (((((((v1 - v1) - ((1.5 / v1) * (v1 / 3.14159))) + (v1 + v1)) / 1.5) + v1) * v1) >= -2.25)), F1: 0, F2: (v10 + (v10 % ((v10) | 1)))}
            switch (fi15 << ((2) & 63)) {
            case 626:
                switch (fi16 << (((v2 & v2)) & 63)) {
                case 625:
                    v18 := []int{(fi15 % (((fi15 ^ (fi16 / ((fi15) | 1)))) | 1))}
                    v4 = (v4 + v4)
                    fmt.Println(v10, v4)
                    _ = v18
                }
            case 32767:
                v4 = (v4 + "")
            default:
                v19 := []int{fi15, -450, fi15}
                _ = v19
            }
            _ = fi16
            _ = v17
        }
        var v20 uint32 = 2
        for fi21 := 0; fi21 < 3; fi21++ {
            v4 = (v4 + "a")
            v3[(v20 & v20)] = 42
            _ = fi21
        }
        _ = fi15
        _ = v20
    }
    v22 := GenS3{F0: false, F1: 1, F2: v10}
    switch (((("foo bar" + "a") + v4) + ((("\n\t" + "a") + v4) + v4)) + (v4 + "hello")) {
    case "a":
        var v23 int32 = 1
        _ = v23
    case "hello":
        var v24 uint16 = 491
        _ = v24
    case "":
        switch -884 {
        case 4294967295:
            for fi25 := 0; fi25 < 13; fi25++ {
                for fi26 := 0; fi26 < 8; fi26++ {
                    v1 = v1
                    var v27 uint32 = 42
                    _ = len(v3)
                    fmt.Println(fi25, v10)
                    _ = fi26
                    _ = v27
                }
                _ = fi25
            }
        case -353:
            var v28 int8 = 100
            _ = v28
        default:
            v3[0] = 65535
        }
    default:
        for fi29 := 0; fi29 < 14; fi29++ {
            var v30 uint64 = 0
            fmt.Println(v30)
            _ = fi29
            _ = v30
        }
    }
    var v31 int = 42
    v32 := GenS4{F0: (^(v2)), F1: 2}
    if (!(-1 < -1)) {
        v4 = v4
    }
    var v33 float32 = (0.0 + ((v1 * (((v1 + float32(-2.25)) + v1) * v1)) / 1.0))
    if true {
        fmt.Println(v10)
        var v34 uint = 255
        v35 := []int{v31, v31, v31}
        _ = v34
        _ = v35
    }
    v1 = float32(1.5)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v22
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v4
}
