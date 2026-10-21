package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 bool }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int }
func genFn0(p0 uint, p1 int32) int32 {
    var v6 int = 68
    v6 -= 42
    p1 -= p1
    v7 := GenS1{F0: true, F1: false}
    v8 := []int{32767}
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    _ = v8
    return (^(p1))
}

func main() {
    var v1 int8 = 42
    var v2 uint = 2
    v3 := make(map[string]uint32)
    v4 := []int{0, 32768}
    var v5 float32 = 0.0
    v9 := GenS2{F0: 42}
    switch "\n\t" {
    case "foo bar":
        var v10 uint32 = 65535
        _ = v10
    default:
        v11 := GenS3{F0: -832}
        _ = v11
    }
    if false {
        fmt.Println(v1, v5, v5)
    }
    fmt.Println(v1)
    switch -9223372036854775808 {
    case -480:
        for fi12 := 0; fi12 < 13; fi12++ {
            switch fi12 {
            case -106:
                v13 := GenS4{F0: -452}
                _ = v13
            case 128:
                v1 -= (-(v1))
            case 2147483647:
                var v14 float32 = v5
                _ = v14
            }
            _ = fi12
        }
    default:
        delete(v3, "a")
    }
    for fi15 := 0; fi15 < 20; fi15++ {
        v16 := v3["a"]
        for fi17 := 0; fi17 < 19; fi17++ {
            v1 = (v1 & (v1 << (((v2 & 18446744073709551615)) & 63)))
            if false {
                var v18 float64 = 1.5
                _ = v18
            }
            fmt.Println(v2, fi15, v5)
            switch "a" {
            case "foo bar":
                var v19 uint16 = uint16(v16)
                _ = v19
            default:
                var v20 int = fi15
                _ = v20
            }
            _ = fi17
        }
        _ = fi15
        _ = v16
    }
    switch "hello" {
    case "hello":
        v1 = v1
    case "a":
        var v21 int8 = (v1 | (^(v1)))
        _ = v21
    }
    fmt.Println(v1)
    _ = len(v3)
    _ = len(v3)
    fmt.Println(v5)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
