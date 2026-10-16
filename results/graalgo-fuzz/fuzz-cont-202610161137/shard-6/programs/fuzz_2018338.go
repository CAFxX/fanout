package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8; F2 uint32 }
func genFn0(p0 int8) uint {
    p0 = -55
    var v5 string = ((("foo bar" + "hello") + ("foo bar" + ("\n\t" + "a"))) + "foo bar")
    _ = p0
    _ = v5
    return 2
}

func main() {
    var v1 bool = false
    var v2 uint64 = 0
    var v3 uint32 = 2147483647
    v4 := []int{}
    var v6 uint8 = 0
    switch "hello" {
    case "":
        if v1 {
            v6 = (v6 & v6)
        } else {
            switch -542 {
            case -9223372036854775808:
                var v7 uint64 = 1
                _ = v7
            case 187:
                fmt.Println(v6, v1, v3)
            default:
                for fi8 := 0; fi8 < 2; fi8++ {
                    fmt.Println(fi8)
                    _ = fi8
                }
            }
            var v9 float64 = 0.0
            _ = v9
        }
    default:
        _ = len(v4)
    }
    switch "hello" {
    case "":
        v10 := GenS1{F0: -583, F1: (^(v6)), F2: v3}
        _ = v10
    }
    switch "foo bar" {
    case "foo bar":
        v4 = append(v4, -998)
    default:
        var v11 bool = false
        _ = v11
    }
    _ = len(v4)
    for fi12 := 0; fi12 < 15; fi12++ {
        v6 = (v6 >> ((genFn0(100)) & 63))
        _ = fi12
    }
    fmt.Println(v3, v6, v6)
    for fi13 := 0; fi13 < 20; fi13++ {
        fmt.Println(v3, v1, v1)
        _ = len(v4)
        _ = fi13
    }
    switch "\n\t" {
    case "hello":
        var v14 int64 = 1
        _ = v14
    case "":
        var v15 uint16 = 2
        _ = v15
    default:
        var v16 uint = 65535
        _ = v16
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    switch "foo bar" {
    case "\n\t":
        v2 = 2
    case "":
        switch -376 {
        case 0:
            var v17 int64 = -241
            _ = v17
        case 1:
            fmt.Println(v3, v3, v6)
        default:
            var v18 int = 286
            _ = v18
        }
    case "hello":
        switch "hello" {
        case "":
            if ((567 >= 430) || ((643 <= 128) || (4294967295 <= -9223372036854775808))) {
                v4 = append(v4, -790)
            } else {
                var v19 uint32 = v3
                var v20 int = 591
                _ = v19
                _ = v20
            }
        case "\n\t":
            switch "\n\t" {
            case "foo bar":
                var v21 uint8 = v6
                _ = v21
            case "hello":
                v4 = append(v4, 128)
            }
        }
    }
    v22 := GenS1{F0: 1, F1: 255, F2: 611}
    if ((v6 ^ v6) == (v6 / (((v6 + (v6 % ((v6) | 1)))) | 1))) {
        switch "a" {
        case "":
            if (42 == (^(v3))) {
                v4 = append(v4, int(v6))
                v2 = (^(v2))
            }
        case "\n\t":
            var v23 uint64 = v2
            _ = v23
        default:
            v2 = 4294967295
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v6
}
