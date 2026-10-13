package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint64; F2 uint8 }
type GenS2 struct { F0 uint; F1 int64 }
func genFn0(p0 uint64, p1 uint16) uint16 {
    switch "" {
    case "hello":
        fmt.Println(p1)
    default:
        switch -9223372036854775808 {
        case 1:
            p1 += (p1 >> ((18446744073709551615) & 63))
        case -819:
            p1 = 42
        default:
            p1 -= p1
        }
    }
    v5 := GenS1{F0: "a", F1: p0, F2: 22}
    _ = p0
    _ = p1
    _ = v5
    return 0
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[int]bool)
    v3 := make(map[int]int)
    v4 := make(map[uint32]uint32)
    var v6 uint32 = 1
    v7 := GenS1{F0: "foo bar", F1: 301, F2: 96}
    if false {
        v6 = (v6 ^ (^(v6)))
        fmt.Println(v6)
        for fi8 := 0; fi8 < 8; fi8++ {
            var v9 int = fi8
            switch "\n\t" {
            case "hello":
                delete(v2, (fi8 + v9))
            case "\n\t":
                var v10 uint32 = (v6 / ((v6) | 1))
                _ = v10
            case "foo bar":
                v6 -= v6
            default:
                v6 -= (v6 % ((v6) | 1))
            }
            _ = fi8
            _ = v9
        }
    }
    var v11 float32 = float32(1.5)
    for fi12 := 0; fi12 < 17; fi12++ {
        switch (fi12 | (^(fi12))) {
        case 32767:
            v13 := GenS1{F0: "", F1: 2, F2: 128}
            _ = v13
        case -25:
            v14 := []int{420, (fi12 + (fi12 << ((255) & 63)))}
            _ = v14
        }
        _ = fi12
    }
    for fi15 := 0; fi15 < 9; fi15++ {
        fmt.Println(fi15, fi15, v11)
        _ = fi15
    }
    for fi16 := 0; fi16 < 18; fi16++ {
        var v17 int8 = 1
        _ = fi16
        _ = v17
    }
    v18 := GenS2{F0: 1, F1: -47}
    var v19 int32 = 713
    for fi20 := 0; fi20 < 9; fi20++ {
        var v21 float32 = 3.14159
        for fi22 := 0; fi22 < 2; fi22++ {
            v19 = v19
            fmt.Println(v11, v21)
            _ = fi22
        }
        v23 := GenS2{F0: 557, F1: -287}
        _ = fi20
        _ = v21
        _ = v23
    }
    var v24 float32 = v11
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v18
    _ = v19
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v6
    _ = v7
}
