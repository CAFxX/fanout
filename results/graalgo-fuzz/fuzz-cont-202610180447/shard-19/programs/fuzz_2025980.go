package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16 }
type GenS2 struct { F0 int8; F1 uint32; F2 int }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint16; F1 int64 }
type GenS5 struct { F0 uint32 }
func genFn0(p0 uint16) int {
    v7 := GenS1{F0: 1000, F1: 42}
    for fi8 := 0; fi8 < 13; fi8++ {
        var v9 string = ""
        v10 := GenS1{F0: 42, F1: 1000}
        v11 := []int{(^(fi8))}
        _ = fi8
        _ = v10
        _ = v11
        _ = v9
    }
    _ = p0
    _ = v7
    return 42
}

func main() {
    var v1 uint32 = 42
    var v2 int32 = -610
    var v3 float32 = (0.1 + (float32(-2.25) + ((0.1 * 1e10) * -2.25)))
    var v4 int32 = (-(v2))
    var v5 uint8 = 127
    v6 := make(map[uint32]string)
    v2 = 0
    var v12 int16 = 32767
    v13 := []int{4294967295}
    for fi14 := 0; fi14 < 13; fi14++ {
        v15 := GenS1{F0: v12, F1: v12}
        for fi16 := 0; fi16 < 9; fi16++ {
            v17 := GenS2{F0: 83, F1: v1, F2: (^(fi16))}
            switch "foo bar" {
            case "":
                if true {
                    v18 := GenS3{F0: int16(v5)}
                    _ = v18
                } else {
                    v19 := GenS3{F0: (^(v12))}
                    _ = v19
                }
            }
            v20 := GenS4{F0: 608, F1: -1}
            _ = fi16
            _ = v17
            _ = v20
        }
        switch fi14 {
        case 1:
            fmt.Println(v5)
        }
        _ = fi14
        _ = v15
    }
    switch "hello" {
    case "hello":
        fmt.Println(v12, v5, v1)
    default:
        v3 = float32(1.5)
    }
    v2 += int32(v12)
    fmt.Println(v3, v3)
    for fi21 := 0; fi21 < 2; fi21++ {
        var v22 uint = 42
        var v23 uint16 = 32767
        v24 := GenS3{F0: (v12 + (^(v12)))}
        _ = fi21
        _ = v22
        _ = v23
        _ = v24
    }
    v25 := GenS5{F0: 348}
    var v26 int16 = 891
    v1 += (v1 << ((0) & 63))
    for fi27 := 0; fi27 < 7; fi27++ {
        v28 := GenS2{F0: -1, F1: (v1 - (v1 | v1)), F2: fi27}
        var v29 float32 = v3
        _ = fi27
        _ = v28
        _ = v29
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
