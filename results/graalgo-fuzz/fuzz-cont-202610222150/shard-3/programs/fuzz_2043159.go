package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 string; F2 uint16 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 bool }

func main() {
    v1 := make(map[int]string)
    var v2 int16 = -32768
    v3 := make(map[int]bool)
    var v4 int = 29
    v5 := make(map[uint32]int)
    for fi6 := 0; fi6 < 20; fi6++ {
        var v7 int64 = 9223372036854775807
        var v8 int32 = int32(v2)
        if true {
            v2 += (v2 >> ((338) & 63))
        }
        _ = fi6
        _ = v7
        _ = v8
    }
    switch 4294967295 {
    case 1:
        var v9 int16 = 42
        _ = v9
    }
    var v10 int8 = -1
    var v11 uint16 = 42
    for fi12 := 0; fi12 < 1; fi12++ {
        switch "foo bar" {
        case "foo bar":
            var v13 uint = 255
            _ = v13
        case "hello":
            fmt.Println(v10, fi12)
        case "a":
            fmt.Println(fi12, v4, v11)
        default:
            v14 := []int{int(v10), 1, 32768}
            _ = v14
        }
        switch v4 {
        case 0:
            v10 += v10
        default:
            v4 += v4
        }
        _ = fi12
    }
    for fi15 := 0; fi15 < 7; fi15++ {
        v16 := GenS1{F0: (v2 - (v2 ^ (v2 % ((v2) | 1)))), F1: "\n\t", F2: v11}
        _ = fi15
        _ = v16
    }
    fmt.Println(v11, v11, v11)
    v17 := GenS2{F0: v4}
    if true {
        v4 += v4
    }
    v18 := GenS3{F0: true}
    for fi19 := 0; fi19 < 16; fi19++ {
        switch fi19 {
        case -1:
            v11 = (^(v11))
        default:
            fmt.Println(v4)
        }
        if (!false) {
            for fi20 := 0; fi20 < 4; fi20++ {
                switch ("hello" + "foo bar") {
                case "hello":
                    var v21 uint8 = 223
                    _ = v21
                default:
                    v11 = 243
                }
                if true {
                    var v22 int64 = -9223372036854775808
                    fmt.Println(v2)
                    v23 := v1[int(v10)]
                    switch "\n\t" {
                    case "hello":
                        v24 := GenS3{F0: ((0.0 - 0.0) <= 3.14159)}
                        _ = v24
                    case "\n\t":
                        fmt.Println(fi20, v22)
                    }
                    var v25 int64 = (v22 + 469)
                    _ = v22
                    _ = v23
                    _ = v25
                }
                _ = fi20
            }
            var v26 int8 = v10
            v27 := GenS1{F0: (v2 + v2), F1: "hello", F2: (v11 * v11)}
            _ = v26
            _ = v27
        }
        _ = fi19
    }
    v10 = (v10 << ((0) & 63))
    fmt.Println(v4)
    v28 := GenS2{F0: v4}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v10
    _ = v11
    _ = v17
    _ = v18
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
