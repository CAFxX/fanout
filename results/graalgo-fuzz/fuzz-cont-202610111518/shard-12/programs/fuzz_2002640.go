package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 string }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int64; F1 uint8 }

func main() {
    var v1 int32 = 0
    var v2 uint8 = 127
    v3 := make(map[string]string)
    var v4 uint64 = 4294967295
    v5 := make(map[int]int)
    var v6 float32 = 3.14159
    if true {
        switch 311 {
        case 32767:
            v2 = (v2 >> ((65535) & 63))
        case 2147483647:
            if false {
                var v7 int8 = 17
                var v8 uint16 = 0
                _ = v7
                _ = v8
            }
        case 0:
            v1 += v1
        }
    }
    var v9 int8 = -60
    fmt.Println(v9, v2)
    v1 += v1
    v4 += (v4 | 2)
    var v10 string = "hello"
    for fi11 := 0; fi11 < 6; fi11++ {
        fmt.Println(v6)
        _ = fi11
    }
    for fi12 := 0; fi12 < 19; fi12++ {
        v13 := GenS1{F0: v1, F1: ""}
        var v14 int32 = (v1 >> ((uint(v2)) & 63))
        _ = fi12
        _ = v13
        _ = v14
    }
    var v15 uint16 = 65535
    v3[(((v10 + ("foo bar" + v10)) + v10) + v10)] = (v10 + "hello")
    switch 611 {
    case 42:
        for fi16 := 0; fi16 < 18; fi16++ {
            v17 := []int{}
            v18 := GenS2{F0: v9}
            v9 += v9
            _ = fi16
            _ = v17
            _ = v18
        }
    case -956:
        switch "" {
        case "hello":
            for fi19 := 0; fi19 < 15; fi19++ {
                var v20 uint64 = v4
                v21 := GenS3{F0: 9223372036854775807, F1: v2}
                var v22 bool = (false || false)
                _ = fi19
                _ = v20
                _ = v21
                _ = v22
            }
        case "a":
            v9 += (^(v9))
        case "\n\t":
            if true {
                var v23 int32 = v1
                _ = v23
            }
        }
    case 32767:
        v1 -= (v1 ^ 2147483647)
    default:
        _ = len(v5)
    }
    var v24 int64 = 42
    v25 := []int{-2147483648, -2147483648}
    var v26 int = 32767
    var v27 int32 = 42
    var v28 int16 = 304
    var v29 int16 = (v28 | v28)
    v1 = (-(v1))
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v25:", len(v25), v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v15
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
