package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint }
type GenS2 struct { F0 uint32; F1 uint64; F2 uint16 }
type GenS3 struct { F0 int8; F1 int16; F2 int8 }

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint = 2
    v3 := make(map[string]int)
    var v4 float64 = 0.1
    var v5 uint = (v2 - v2)
    if false {
        var v6 int = -338
        v2 = (v5 << ((v5) & 63))
        v7 := GenS1{F0: "foo bar", F1: 42}
        _ = v6
        _ = v7
    } else {
        var v8 uint = v5
        v9 := []int{763, 32768, 128, 2147483647}
        _ = v8
        _ = v9
    }
    switch 42 {
    case 4:
        var v10 int16 = -125
        _ = v10
    case 9223372036854775807:
        var v11 int32 = -2147483648
        _ = v11
    default:
        var v12 bool = false
        _ = v12
    }
    v5 -= v5
    delete(v1, 2147483648)
    if (true || true) {
        if false {
            v13 := GenS1{F0: "\n\t", F1: (^(v5))}
            var v14 uint16 = 749
            var v15 float64 = 1.5
            _ = v13
            _ = v14
            _ = v15
        }
        if true {
            fmt.Println(v5, v4, v5)
        }
        v16 := []int{431}
        _ = v16
    }
    for fi17 := 0; fi17 < 9; fi17++ {
        v4 = v4
        fmt.Println(v4, v2)
        for fi18 := 0; fi18 < 20; fi18++ {
            switch (fi18 | (fi17 * fi18)) {
            case 9223372036854775807:
                if (1 == 1) {
                    var v19 uint32 = 2
                    _ = v19
                }
            case 32768:
                v20 := GenS1{F0: "", F1: v5}
                _ = v20
            case 32767:
                v21 := GenS1{F0: "", F1: v2}
                _ = v21
            }
            v22 := GenS2{F0: 2147483647, F1: 4294967295, F2: 255}
            switch "a" {
            case "a":
                if ((("\n\t" + "\n\t") + (("a" + "foo bar") + "foo bar")) > "a") {
                    v5 -= v5
                    v2 -= (v2 & (^(v2)))
                }
            }
            var v23 int = 285
            _ = fi18
            _ = v22
            _ = v23
        }
        _ = fi17
    }
    switch 42 {
    case 9223372036854775807:
        v24 := GenS1{F0: "foo bar", F1: (^(v2))}
        _ = v24
    default:
        v2 -= v2
    }
    v25 := GenS3{F0: -128, F1: 1000, F2: -1}
    var v26 int8 = -38
    var v27 int32 = 0
    v2 = (^(v2))
    switch 1 {
    case 2147483647:
        var v28 float32 = 1.5
        _ = v28
    case -667:
        v29 := GenS1{F0: "\n\t", F1: 53}
        _ = v29
    case 127:
        v2 -= v2
    default:
        for fi30 := 0; fi30 < 11; fi30++ {
            v31 := []int{}
            v2 = 908
            _ = fi30
            _ = v31
        }
    }
    var v32 uint64 = 454
    v32 = (v32 << ((v5) & 63))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v32
    _ = v4
    _ = v5
}
