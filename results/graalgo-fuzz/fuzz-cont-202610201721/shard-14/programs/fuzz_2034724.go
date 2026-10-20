package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16 }
type GenS2 struct { F0 uint32; F1 uint16 }

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[uint32]bool)
    var v3 uint64 = 0
    var v4 int8 = 127
    v3 -= v3
    switch "" {
    case "a":
        var v5 float64 = 0.0
        _ = v5
    case "":
        for fi6 := 0; fi6 < 8; fi6++ {
            v4 -= v4
            var v7 int8 = -128
            _ = fi6
            _ = v7
        }
    case "\n\t":
        var v8 uint64 = (v3 + 18446744073709551615)
        _ = v8
    }
    for fi9 := 0; fi9 < 4; fi9++ {
        switch fi9 {
        case -386:
            v10 := []int{fi9, -333, fi9, 0}
            _ = v10
        case 32768:
            for fi11 := 0; fi11 < 14; fi11++ {
                v12 := []int{}
                delete(v2, 802)
                fmt.Println(fi9, fi9, v3)
                _ = fi11
                _ = v12
            }
        default:
            var v13 bool = true
            _ = v13
        }
        switch fi9 {
        case 42:
            v4 = v4
        }
        _ = fi9
    }
    v14 := GenS1{F0: true, F1: -481}
    v4 = (v4 / ((v4) | 1))
    var v15 float64 = 0.0
    switch "\n\t" {
    case "hello":
        var v16 uint32 = 1
        _ = v16
    default:
        var v17 uint = 0
        _ = v17
    }
    switch -498 {
    case 594:
        switch "a" {
        case "\n\t":
            v18 := []int{-1}
            _ = v18
        case "a":
            v19 := v1[65535]
            _ = v19
        }
    case 32767:
        v20 := GenS1{F0: (!(!false)), F1: -32768}
        _ = v20
    case 286:
        if (2 == 0) {
            v15 += float64(0.0)
            v4 = (-(v4))
            fmt.Println(v15)
        } else {
            switch 9223372036854775807 {
            case 128:
                v15 -= -2.25
            case 0:
                for fi21 := 0; fi21 < 8; fi21++ {
                    for fi22 := 0; fi22 < 3; fi22++ {
                        v3 -= 1
                        _ = fi22
                    }
                    _ = fi21
                }
            case 32767:
                v23 := GenS1{F0: (false && false), F1: 1}
                _ = v23
            default:
                var v24 float32 = 0.1
                _ = v24
            }
        }
    }
    if (0 <= 151) {
        v3 = 741
        v25 := GenS2{F0: 2147483647, F1: 2}
        var v26 uint64 = v3
        _ = v25
        _ = v26
    }
    v15 = 1e10
    var v27 int64 = -1
    for fi28 := 0; fi28 < 10; fi28++ {
        var v29 uint32 = 42
        _ = fi28
        _ = v29
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v27
    _ = v3
    _ = v4
}
