package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8; F2 int }
type GenS2 struct { F0 uint64; F1 int64; F2 uint }
type GenS3 struct { F0 string }
func genFn0(p0 int16, p1 int32, p2 uint32) int32 {
    p1 += (^(p1))
    switch 4294967295 {
    case 1:
        var v6 int32 = (p1 | -1)
        _ = v6
    case 57:
        v7 := GenS1{F0: 255, F1: uint8(p2), F2: 140}
        _ = v7
    }
    _ = p0
    _ = p1
    _ = p2
    return p1
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint16 = 32767
    v3 := []string{"hello", "\n\t", "hello"}
    v4 := make(map[string]bool)
    var v5 int = 32767
    if false {
        var v8 uint16 = 2
        _ = v8
    } else {
        var v9 int32 = 456
        _ = v9
    }
    for fi10 := 0; fi10 < 13; fi10++ {
        v11 := GenS2{F0: 4294967295, F1: -9223372036854775808, F2: 18446744073709551615}
        v5 += (fi10 << ((786) & 63))
        _ = fi10
        _ = v11
    }
    var v12 int = (-(v5))
    v13 := GenS1{F0: 40, F1: 42, F2: (v5 & (v5 & -832))}
    v2 += v2
    if false {
        v14 := GenS1{F0: 2, F1: 194, F2: (v12 ^ v5)}
        v2 = v2
        var v15 uint64 = 1
        _ = v14
        _ = v15
    }
    fmt.Println(v5, v2)
    v16 := GenS2{F0: 762, F1: -9223372036854775808, F2: 4294967295}
    var v17 float32 = float32(1.5)
    switch v5 {
    case 128:
        if true {
            switch (^(v5)) {
            case 9223372036854775807:
                var v18 uint8 = 2
                _ = v18
            }
            switch "\n\t" {
            case "\n\t":
                if true {
                    _ = len(v1)
                    for fi19 := 0; fi19 < 19; fi19++ {
                        var v20 uint64 = 534
                        _ = fi19
                        _ = v20
                    }
                    var v21 int64 = 42
                    _ = v21
                } else {
                    fmt.Println(v2, v12, v12)
                    v2 = 128
                    fmt.Println(v17)
                }
            case "foo bar":
                var v22 string = "a"
                _ = v22
            case "":
                var v23 int16 = 1000
                _ = v23
            default:
                v24 := GenS1{F0: 128, F1: 128, F2: v12}
                _ = v24
            }
        }
    case 1:
        var v25 uint64 = 1
        _ = v25
    }
    v26 := GenS1{F0: 127, F1: 255, F2: 32767}
    switch v12 {
    case 64:
        v2 = (v2 >> ((4294967295) & 63))
    default:
        var v27 int64 = 42
        _ = v27
    }
    switch -9223372036854775808 {
    case -648:
        var v28 float32 = v17
        _ = v28
    case 1:
        for fi29 := 0; fi29 < 18; fi29++ {
            v30 := GenS2{F0: 857, F1: -9223372036854775808, F2: uint(v2)}
            v3 = append(v3, "hello")
            _ = fi29
            _ = v30
        }
    case -1:
        _ = len(v3)
    }
    v31 := GenS2{F0: 18446744073709551615, F1: -930, F2: 0}
    v32 := GenS3{F0: "hello"}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v16
    _ = v17
    _ = v2
    _ = v26
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v5
}
