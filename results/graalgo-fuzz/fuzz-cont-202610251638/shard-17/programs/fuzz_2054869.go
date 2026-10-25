package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint32; F2 uint }
type GenS2 struct { F0 uint8; F1 int8; F2 int32 }
type GenS3 struct { F0 int; F1 string }

func main() {
    var v1 uint64 = 1
    v2 := make(map[uint32]uint32)
    v3 := make(map[int]bool)
    v1 -= 2
    v4 := []int{int(v1), -238, 325, 331}
    if (1 != -817) {
        v1 = v1
    }
    var v5 float64 = 0.1
    v1 = (^(v1))
    fmt.Println(v1, v1)
    v6 := GenS1{F0: 609, F1: 4294967295, F2: 42}
    if (true || ((!(-92 == 1)) && false)) {
        v5 -= v5
        for fi7 := 0; fi7 < 4; fi7++ {
            switch "hello" {
            case "":
                v8 := GenS1{F0: 551, F1: 4294967295, F2: 42}
                _ = v8
            default:
                v1 += (v1 % ((862) | 1))
            }
            var v9 uint16 = 1
            _ = fi7
            _ = v9
        }
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    } else {
        if len(v4) > 0 {
            v4[0] = 0
        }
        fmt.Println(v5, v1)
    }
    v10 := GenS1{F0: 558, F1: 4294967295, F2: 18446744073709551615}
    v11 := GenS1{F0: 42, F1: 65535, F2: 255}
    v12 := GenS1{F0: 42, F1: 1, F2: 4294967295}
    v4 = append(v4, 127)
    if false {
        for fi13 := 0; fi13 < 16; fi13++ {
            var v14 float32 = 1e10
            v15 := GenS2{F0: 164, F1: -1, F2: 2147483647}
            _ = fi13
            _ = v14
            _ = v15
        }
    } else {
        fmt.Println(v1)
        for fi16 := 0; fi16 < 6; fi16++ {
            var v17 uint = 0
            _ = fi16
            _ = v17
        }
    }
    v18 := GenS2{F0: 2, F1: 127, F2: -1}
    var v19 float32 = 3.14159
    for fi20 := 0; fi20 < 13; fi20++ {
        switch (fi20 - (fi20 + -824)) {
        case 350:
            var v21 uint64 = (v1 / (((v1 << ((607) & 63))) | 1))
            _ = v21
        default:
            if false {
                v1 = 236
                var v22 uint8 = 127
                v5 = 0.1
                _ = v22
            }
        }
        _ = fi20
    }
    v1 -= v1
    if (true || (0 <= 315)) {
        switch 109 {
        case 2147483647:
            v1 -= (v1 % ((18446744073709551615) | 1))
        case -2147483648:
            v1 += 18446744073709551615
        case -1:
            v23 := GenS3{F0: -1, F1: ""}
            _ = v23
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
