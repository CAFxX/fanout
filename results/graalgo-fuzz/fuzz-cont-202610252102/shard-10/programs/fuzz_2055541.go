package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int64 }

func main() {
    var v1 uint = 996
    v2 := make(map[uint32]int)
    var v3 int8 = 0
    v4 := make(map[int]int)
    var v5 uint8 = 1
    var v6 uint64 = 2
    var v7 int8 = (^(v3))
    switch 128 {
    case 128:
        for fi8 := 0; fi8 < 17; fi8++ {
            var v9 int16 = -983
            var v10 float32 = (3.14159 + ((0.0 / float32(0.1)) * 3.14159))
            v11 := GenS1{F0: v6, F1: 42}
            _ = fi8
            _ = v10
            _ = v11
            _ = v9
        }
    case 4294967295:
        var v12 int64 = -458
        _ = v12
    }
    v13 := GenS1{F0: v6, F1: 702}
    if true {
        if ((1 < 42) && true) {
            var v14 float32 = (3.14159 / 1e10)
            var v15 float32 = v14
            _ = v14
            _ = v15
        } else {
            v5 = (v5 | (v5 / (((v5 + v5)) | 1)))
        }
        fmt.Println(v3)
        switch ("a" + "hello") {
        case "hello":
            switch ("" + "\n\t") {
            case "foo bar":
                v7 -= v3
            case "":
                fmt.Println(v7)
            default:
                v3 = (v7 ^ (^(v7)))
            }
        default:
            var v16 uint16 = 425
            _ = v16
        }
    }
    var v17 float32 = (((float32(1e10) * (float32(1e10) - (-2.25 - 3.14159))) - (((0.0 / 3.14159) / 3.14159) * 3.14159)) / (1e10 / 3.14159))
    switch "foo bar" {
    case "\n\t":
        v7 = (v7 ^ v7)
    case "":
        var v18 int64 = -879
        _ = v18
    case "hello":
        v6 += (v6 << ((42) & 63))
    default:
        v19 := GenS1{F0: uint64(v1), F1: -9223372036854775808}
        _ = v19
    }
    for fi20 := 0; fi20 < 7; fi20++ {
        for fi21 := 0; fi21 < 4; fi21++ {
            _ = len(v2)
            _ = fi21
        }
        _ = fi20
    }
    switch "\n\t" {
    case "hello":
        var v22 bool = false
        _ = v22
    default:
        for fi23 := 0; fi23 < 14; fi23++ {
            switch (-(fi23)) {
            case -1:
                delete(v4, (^(fi23)))
            case 0:
                v6 -= uint64(v5)
            case 32767:
                fmt.Println(v5)
            }
            _ = fi23
        }
    }
    if false {
        for fi24 := 0; fi24 < 8; fi24++ {
            v6 -= (^(v6))
            _ = fi24
        }
        fmt.Println(v5, v5)
        for fi25 := 0; fi25 < 16; fi25++ {
            if true {
                fmt.Println(v6, v17)
                var v26 int64 = 42
                var v27 int16 = 0
                _ = v26
                _ = v27
            } else {
                if false {
                    var v28 bool = true
                    _ = v28
                }
            }
            for fi29 := 0; fi29 < 10; fi29++ {
                v6 -= uint64(v1)
                v3 += (v7 % ((127) | 1))
                var v30 int32 = -1
                _ = fi29
                _ = v30
            }
            _ = fi25
        }
    } else {
        switch 4294967295 {
        case 127:
            fmt.Println(v3, v5)
        default:
            var v31 int8 = v3
            _ = v31
        }
    }
    v3 = v3
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v13
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
