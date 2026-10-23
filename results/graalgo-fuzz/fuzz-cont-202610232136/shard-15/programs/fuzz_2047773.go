package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint32; F2 int64 }
type GenS2 struct { F0 uint32; F1 int64; F2 uint }
type GenS3 struct { F0 int8; F1 uint32 }
type GenS4 struct { F0 uint8; F1 uint64; F2 int8 }

func main() {
    var v1 int = 4294967295
    var v2 uint64 = 563
    var v3 int16 = -32768
    v4 := make(map[uint32]uint32)
    var v5 uint = 4294967295
    v6 := []string{"hello"}
    if (!false) {
        var v7 uint8 = 1
        var v8 bool = (255 != 572)
        v3 = v3
        _ = v7
        _ = v8
    }
    switch "hello" {
    case "":
        var v9 float64 = 1.5
        _ = v9
    case "\n\t":
        v5 -= (v5 * uint(v2))
    case "foo bar":
        fmt.Println(v5, v3, v2)
    }
    fmt.Println(v5, v1, v3)
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    switch v1 {
    case 32767:
        var v10 uint64 = (v2 | v2)
        _ = v10
    case 4294967295:
        switch (("\n\t" + "") + ("" + "a")) {
        case "":
            var v11 int8 = 100
            _ = v11
        default:
            fmt.Println(v2)
        }
    default:
        var v12 int8 = -34
        _ = v12
    }
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    switch "\n\t" {
    case "a":
        if true {
            _ = len(v4)
        }
    case "":
        switch (v1 - (v1 ^ v1)) {
        case 4294967295:
            v13 := GenS1{F0: -9223372036854775808, F1: 0, F2: 9223372036854775807}
            _ = v13
        case 32767:
            v14 := GenS1{F0: 1, F1: 42, F2: 9223372036854775807}
            _ = v14
        case 1000:
            var v15 uint32 = 2147483648
            _ = v15
        default:
            var v16 uint16 = 65535
            _ = v16
        }
    case "hello":
        var v17 uint8 = 128
        _ = v17
    }
    switch "\n\t" {
    case "":
        v1 -= (v1 << (((^(v5))) & 63))
    default:
        switch ((("\n\t" + "foo bar") + "hello") + "foo bar") {
        case "":
            switch "\n\t" {
            case "a":
                v3 += v3
            case "foo bar":
                v18 := GenS2{F0: 2, F1: 9223372036854775807, F2: (^(v5))}
                _ = v18
            case "":
                v1 += v1
            }
        }
    }
    var v19 float32 = 0.0
    for fi20 := 0; fi20 < 13; fi20++ {
        v19 = v19
        v3 = v3
        v21 := GenS3{F0: 100, F1: 489}
        _ = fi20
        _ = v21
    }
    v3 = v3
    for fi22 := 0; fi22 < 10; fi22++ {
        fmt.Println(v2, v19)
        _ = fi22
    }
    v23 := GenS4{F0: 61, F1: 42, F2: -1}
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
