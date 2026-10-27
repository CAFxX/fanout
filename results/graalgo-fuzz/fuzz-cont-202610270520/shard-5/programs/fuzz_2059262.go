package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int }
type GenS2 struct { F0 int; F1 int8; F2 uint32 }

func main() {
    var v1 uint = 0
    var v2 uint32 = 42
    var v3 float64 = -2.25
    var v4 int64 = -1
    var v5 float32 = 3.14159
    var v6 int16 = 959
    switch "a" {
    case "a":
        var v7 float64 = (0.1 / (0.0 + v3))
        _ = v7
    case "":
        fmt.Println(v1)
    }
    switch "hello" {
    case "":
        for fi8 := 0; fi8 < 8; fi8++ {
            v9 := []int{32767, 9223372036854775807}
            if len(v9) >= 2 {
                v9 = v9[1:len(v9)]
            }
            var v10 uint64 = 4294967295
            _ = fi8
            _ = v10
            _ = v9
        }
    case "foo bar":
        var v11 uint16 = 898
        _ = v11
    case "hello":
        switch "a" {
        case "foo bar":
            for fi12 := 0; fi12 < 15; fi12++ {
                var v13 string = "foo bar"
                _ = fi12
                _ = v13
            }
        case "hello":
            if true {
                fmt.Println(v1, v4)
                v14 := []int{222, 32768}
                fmt.Println(v2, v3, v6)
                _ = v14
            } else {
                v2 -= (v2 ^ (^(v2)))
            }
        default:
            v15 := []int{710, 128, 2147483647, -9223372036854775808}
            _ = v15
        }
    }
    v16 := []int{}
    v3 += float64(1.5)
    switch ("\n\t" + "\n\t") {
    case "hello":
        v1 += v1
    case "a":
        if len(v16) > 0 {
            v16[0] = 1
        }
    default:
        var v17 string = "foo bar"
        _ = v17
    }
    var v18 uint16 = 272
    v19 := GenS1{F0: v4, F1: 9223372036854775807}
    if true {
        fmt.Println(v18)
        fmt.Println(v2)
        if true {
            v18 -= v18
        } else {
            v20 := GenS1{F0: (v4 + (v4 << ((1) & 63))), F1: 615}
            v21 := GenS2{F0: -1, F1: 8, F2: (v2 + 2147483647)}
            _ = v20
            _ = v21
        }
    } else {
        fmt.Println(v6)
        fmt.Println(v6, v3)
        var v22 float32 = v5
        _ = v22
    }
    for fi23 := 0; fi23 < 7; fi23++ {
        var v24 uint64 = 42
        var v25 float32 = 0.1
        _ = fi23
        _ = v24
        _ = v25
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
