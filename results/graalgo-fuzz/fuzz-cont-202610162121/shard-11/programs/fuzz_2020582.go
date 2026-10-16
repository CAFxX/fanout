package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint }
type GenS2 struct { F0 int8; F1 int64; F2 uint8 }
type GenS3 struct { F0 uint; F1 int16 }
type GenS4 struct { F0 uint }
func genFn0(p0 int8, p1 uint8, p2 bool) uint32 {
    v5 := GenS1{F0: 1, F1: 526}
    var v6 uint32 = 65535
    switch 4294967295 {
    case -146:
        fmt.Println(p2, p2, p0)
    case -2147483648:
        if p2 {
            fmt.Println(p0, v6)
            fmt.Println(p2)
        } else {
            switch "foo bar" {
            case "\n\t":
                v7 := GenS2{F0: p0, F1: -1, F2: 42}
                _ = v7
            case "hello":
                v6 = 722
            case "foo bar":
                if true {
                    if false {
                        p2 = (0 <= 18446744073709551615)
                        fmt.Println(p1)
                        v8 := []int{}
                        _ = v8
                    } else {
                        v9 := GenS3{F0: 1, F1: int16(v6)}
                        fmt.Println(p1, v6)
                        _ = v9
                    }
                }
            default:
                var v10 uint = 42
                _ = v10
            }
        }
    case -1:
        if p2 {
            v11 := []int{}
            p1 -= p1
            _ = len(v11)
            _ = v11
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return 42
}

func main() {
    v1 := []int{}
    var v2 uint = 1
    var v3 uint64 = 4294967295
    var v4 float32 = 0.1
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    _ = len(v1)
    v12 := GenS4{F0: v2}
    var v13 int = -2147483648
    if len(v1) > 0 {
        v1[0] = v13
    }
    _ = len(v1)
    v14 := GenS2{F0: 42, F1: -9223372036854775808, F2: 150}
    _ = len(v1)
    v13 = v13
    for fi15 := 0; fi15 < 3; fi15++ {
        var v16 int32 = -1
        if false {
            v16 -= v16
            var v17 uint8 = 61
            switch ("hello" + "foo bar") {
            case "hello":
                var v18 float32 = 1.5
                _ = v18
            default:
                var v19 int8 = -1
                _ = v19
            }
            _ = v17
        } else {
            switch "a" {
            case "\n\t":
                var v20 uint16 = 2
                _ = v20
            case "hello":
                v2 -= (v2 << (((v2 >> (((v2 << ((v2) & 63))) & 63))) & 63))
            default:
                for fi21 := 0; fi21 < 6; fi21++ {
                    v4 -= (1e10 + v4)
                    _ = fi21
                }
            }
        }
        v16 = v16
        _ = fi15
        _ = v16
    }
    v2 = 440
    var v22 uint16 = 2
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v22
    _ = v3
    _ = v4
}
