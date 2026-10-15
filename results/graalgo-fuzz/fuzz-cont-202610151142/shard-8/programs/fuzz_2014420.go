package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int8; F2 uint16 }
type GenS2 struct { F0 int64; F1 uint32 }
type GenS3 struct { F0 uint64; F1 int32 }
func genFn0(p0 uint, p1 string, p2 bool) int {
    var v7 uint32 = 42
    switch (p1 + p1) {
    case "hello":
        if p2 {
            p1 = p1
            if p2 {
                switch -821 {
                case -9223372036854775808:
                    p2 = (42 != 42)
                }
                var v8 bool = false
                _ = v8
            }
        } else {
            fmt.Println(p0, p2, p1)
            v9 := []int{}
            _ = v9
        }
    default:
        p1 = p1
    }
    if true {
        p2 = (0 != -191)
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 42
}

func main() {
    v1 := make(map[string]uint32)
    var v2 uint16 = 721
    v3 := make(map[int]int)
    v4 := make(map[int]int)
    v5 := make(map[uint32]bool)
    v6 := make(map[uint32]uint32)
    v10 := GenS1{F0: -1, F1: -128, F2: (^(v2))}
    fmt.Println(v2)
    var v11 uint64 = 4294967295
    v11 += 622
    v12 := GenS1{F0: -761, F1: 42, F2: (v2 & (v2 ^ 0))}
    var v13 uint16 = 0
    if (false && ((0.1 <= (((0.0 * ((3.14159 / -2.25) / 1e10)) * 0.0) - (3.14159 / 1.0))) || false)) {
        fmt.Println(v11)
    } else {
        switch "a" {
        case "hello":
            var v14 float32 = 1.5
            _ = v14
        case "\n\t":
            fmt.Println(v2)
        case "foo bar":
            fmt.Println(v2)
        }
        if true {
            var v15 uint32 = 437
            _ = len(v5)
            v5[438] = (!true)
            _ = v15
        }
    }
    if ("foo bar" == "\n\t") {
        var v16 int64 = 7
        var v17 float32 = -2.25
        var v18 uint8 = 186
        _ = v16
        _ = v17
        _ = v18
    }
    v11 = v11
    v11 += (v11 & (^(v11)))
    switch "a" {
    case "hello":
        if true {
            var v19 uint32 = 4294967295
            v20 := GenS1{F0: -1, F1: 127, F2: (v2 % ((uint16(v11)) | 1))}
            _ = v19
            _ = v20
        } else {
            for fi21 := 0; fi21 < 4; fi21++ {
                v3[fi21] = (fi21 << ((0) & 63))
                _ = fi21
            }
        }
    case "a":
        v22 := GenS2{F0: -9223372036854775808, F1: uint32(v11)}
        _ = v22
    default:
        switch "a" {
        case "foo bar":
            v23 := GenS3{F0: (^(v11)), F1: 1}
            _ = v23
        case "a":
            _ = len(v5)
        default:
            if true {
                v11 -= (v11 & v11)
            }
        }
    }
    var v24 uint64 = uint64(v2)
    for fi25 := 0; fi25 < 6; fi25++ {
        for fi26 := 0; fi26 < 5; fi26++ {
            v24 = (^(v24))
            var v27 uint32 = 963
            fmt.Println(v2, v13)
            _ = fi26
            _ = v27
        }
        _ = fi25
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
