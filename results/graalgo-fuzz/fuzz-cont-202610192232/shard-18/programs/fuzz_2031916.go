package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint64; F1 bool; F2 int8 }
type GenS3 struct { F0 uint8 }
func genFn0(p0 int32, p1 int, p2 uint16) int16 {
    switch int(p2) {
    case 32767:
        var v4 int8 = 0
        _ = v4
    case 128:
        fmt.Println(p0)
    default:
        for fi5 := 0; fi5 < 14; fi5++ {
            p0 = p0
            v6 := GenS1{F0: -1}
            _ = fi5
            _ = v6
        }
    }
    if (true && (1 < 255)) {
        if (2 <= 950) {
            if (0.0 <= 1.5) {
                fmt.Println(p2, p0)
                v7 := []int{(p1 / ((-610) | 1)), (p1 ^ (^(p1)))}
                v8 := GenS1{F0: 406}
                _ = v7
                _ = v8
            }
        }
    }
    _ = p0
    _ = p1
    _ = p2
    return -798
}
func genFn1(p0 uint64, p1 uint64) int32 {
    p1 += (p1 & p1)
    if true {
        v9 := GenS2{F0: 42, F1: true, F2: 1}
        switch "" {
        case "hello":
            v10 := GenS3{F0: 1}
            _ = v10
        case "foo bar":
            v11 := GenS3{F0: 195}
            _ = v11
        case "a":
            p1 -= p0
        default:
            switch "a" {
            case "":
                fmt.Println(p1, p0)
            case "foo bar":
                p0 = (p0 | (p1 ^ p0))
            case "\n\t":
                var v12 uint = 564
                _ = v12
            default:
                v13 := []int{4294967295, 524, 1}
                _ = v13
            }
        }
        p1 += 0
        _ = v9
    }
    var v14 uint16 = uint16(p0)
    _ = p0
    _ = p1
    _ = v14
    return -1
}

func main() {
    var v1 float64 = 1.5
    v2 := make(map[uint32]int)
    v3 := make(map[int]uint32)
    for fi15 := 0; fi15 < 5; fi15++ {
        fmt.Println(v1, v1)
        _ = fi15
    }
    fmt.Println(v1)
    fmt.Println(v1)
    var v16 uint8 = 42
    v1 = ((v1 * float64(-2.25)) - 0.0)
    for fi17 := 0; fi17 < 13; fi17++ {
        for fi18 := 0; fi18 < 4; fi18++ {
            fmt.Println(v1)
            _ = fi18
        }
        _ = fi17
    }
    var v19 int16 = -935
    v20 := GenS2{F0: 375, F1: true, F2: -128}
    var v21 int64 = -383
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v16
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
}
