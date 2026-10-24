package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint16; F1 uint32 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int32; F1 bool; F2 int32 }
type GenS5 struct { F0 uint8; F1 int16; F2 int }
type GenS6 struct { F0 int }

func main() {
    v1 := []uint32{42, 678, 580}
    v2 := make(map[int]bool)
    v3 := make(map[int]int)
    var v4 int = 9223372036854775807
    var v5 float32 = float32(3.14159)
    v6 := make(map[int]int)
    _ = len(v2)
    _ = len(v1)
    var v7 int64 = 1
    switch ("hello" + "hello") {
    case "hello":
        var v8 uint = 207
        _ = v8
    default:
        var v9 uint32 = 422
        _ = v9
    }
    for fi10 := 0; fi10 < 1; fi10++ {
        var v11 uint8 = 151
        v7 = 1
        v7 += 622
        _ = fi10
        _ = v11
    }
    v4 = -1
    v12 := GenS1{F0: 2147483648}
    switch v4 {
    case -1:
        v13 := GenS1{F0: 1}
        _ = v13
    case 32767:
        v14 := GenS2{F0: 42, F1: 16}
        _ = v14
    case 127:
        _ = len(v1)
    default:
        var v15 uint = 255
        _ = v15
    }
    _ = len(v1)
    for fi16 := 0; fi16 < 16; fi16++ {
        for fi17 := 0; fi17 < 2; fi17++ {
            if len(v1) > 0 {
                v1[0] = 0
            }
            v18 := GenS3{F0: 42}
            v19 := GenS4{F0: -497, F1: (18446744073709551615 <= 2), F2: -37}
            _ = fi17
            _ = v18
            _ = v19
        }
        for fi20 := 0; fi20 < 7; fi20++ {
            v7 -= v7
            v21 := GenS5{F0: 128, F1: 32767, F2: 0}
            switch fi16 {
            case 32768:
                _ = len(v1)
                v7 = v7
            default:
                v7 -= 42
            }
            v22 := GenS6{F0: fi20}
            var v23 uint32 = 42
            _ = fi20
            _ = v21
            _ = v22
            _ = v23
        }
        v6[(fi16 * (v4 & (fi16 / (((fi16 % (((-(fi16))) | 1))) | 1))))] = (fi16 / ((fi16) | 1))
        _ = fi16
    }
    var v24 int64 = (v7 | v7)
    v7 += 42
    var v25 int = v4
    v26 := GenS5{F0: 1, F1: 1, F2: int(v24)}
    var v27 uint32 = 2147483648
    var v28 uint8 = 2
    var v29 uint32 = 769
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v12
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
