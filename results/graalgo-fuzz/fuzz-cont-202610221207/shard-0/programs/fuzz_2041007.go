package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8 }
type GenS2 struct { F0 uint; F1 int }
type GenS3 struct { F0 int32; F1 int16 }

func main() {
    v1 := make(map[uint32]string)
    var v2 uint64 = 462
    v3 := make(map[string]uint32)
    switch -620 {
    case 32768:
        v2 = v2
    case 128:
        var v4 int32 = -910
        _ = v4
    default:
        for fi5 := 0; fi5 < 12; fi5++ {
            switch "" {
            case "\n\t":
                var v6 float32 = ((1.5 / 1.5) + (0.1 * 0.1))
                _ = v6
            }
            if true {
                v7 := GenS1{F0: false, F1: -128}
                _ = v7
            } else {
                if (!(!(361 < 2147483647))) {
                    v8 := []int{(-(fi5)), fi5, (fi5 + int(v2)), (fi5 - 32768)}
                    _ = v8
                }
                var v9 float32 = 0.1
                _ = v9
            }
            _ = fi5
        }
    }
    fmt.Println(v2)
    switch 1 {
    case 4294967295:
        var v10 uint32 = 2147483647
        _ = v10
    case 32767:
        var v11 float64 = 0.0
        _ = v11
    case 32768:
        v12 := []int{127, 32767, 974, 1}
        _ = v12
    default:
        v13 := []int{}
        _ = v13
    }
    v14 := GenS2{F0: 1, F1: 2147483647}
    v15 := GenS3{F0: -2147483648, F1: -32768}
    v2 += v2
    delete(v3, "foo bar")
    for fi16 := 0; fi16 < 19; fi16++ {
        fmt.Println(v2)
        _ = fi16
    }
    fmt.Println(v2)
    var v17 uint = 0
    var v18 uint64 = 762
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v14
    _ = v15
    _ = v17
    _ = v18
    _ = v2
    _ = v3
}
