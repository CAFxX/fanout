package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int32; F2 uint64 }
type GenS2 struct { F0 string; F1 string }
type GenS3 struct { F0 bool }

func main() {
    var v1 int32 = 42
    var v2 uint = 18446744073709551615
    var v3 string = "foo bar"
    var v4 int = 128
    v5 := GenS1{F0: "\n\t", F1: v1, F2: 18446744073709551615}
    var v6 uint8 = 0
    v7 := GenS1{F0: "hello", F1: -1, F2: 4294967295}
    for fi8 := 0; fi8 < 17; fi8++ {
        v9 := GenS2{F0: v3, F1: "foo bar"}
        fmt.Println(v6, v6, fi8)
        _ = fi8
        _ = v9
    }
    v10 := []int{v4, v4}
    switch v3 {
    case "foo bar":
        for fi11 := 0; fi11 < 15; fi11++ {
            fmt.Println(fi11, fi11, v1)
            for fi12 := 0; fi12 < 15; fi12++ {
                _ = len(v10)
                _ = fi12
            }
            if (-18 == -562) {
                var v13 int32 = (v1 | v1)
                var v14 uint64 = 619
                v4 += 127
                _ = v13
                _ = v14
            } else {
                fmt.Println(v1, fi11)
                fmt.Println(v2, fi11)
                v4 = v4
            }
            _ = fi11
        }
    case "":
        for fi15 := 0; fi15 < 19; fi15++ {
            v16 := GenS1{F0: v3, F1: v1, F2: uint64(v6)}
            _ = fi15
            _ = v16
        }
    case "\n\t":
        var v17 uint = v2
        _ = v17
    }
    switch v4 {
    case 32767:
        v2 = 461
    case 32768:
        if len(v10) > 0 {
            v10[0] = v4
        }
    default:
        v18 := GenS3{F0: (!false)}
        _ = v18
    }
    v6 -= v6
    v1 += v1
    var v19 int64 = 0
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
