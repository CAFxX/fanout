package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int8; F1 int32 }
type GenS3 struct { F0 uint32 }

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[string]string)
    var v3 int64 = -592
    var v4 uint = 2
    v5 := make(map[int]int)
    var v6 int8 = 76
    v7 := GenS1{F0: v4}
    var v8 int64 = v3
    var v9 uint = (v4 >> (((v4 ^ v4)) & 63))
    for fi10 := 0; fi10 < 15; fi10++ {
        switch ("" + "") {
        case "hello":
            for fi11 := 0; fi11 < 7; fi11++ {
                var v12 uint8 = 162
                _ = fi11
                _ = v12
            }
        case "a":
            var v13 bool = false
            _ = v13
        }
        v4 += (v4 << ((v4) & 63))
        _ = fi10
    }
    v4 -= v4
    var v14 uint64 = 4294967295
    for fi15 := 0; fi15 < 6; fi15++ {
        v16 := GenS2{F0: (v6 * (v6 * -1)), F1: -1}
        v17 := GenS3{F0: 296}
        _ = fi15
        _ = v16
        _ = v17
    }
    if false {
        switch "hello" {
        case "foo bar":
            v3 = (v8 ^ (v8 & (v8 / (((-(v8))) | 1))))
        case "\n\t":
            v18 := []int{42}
            _ = v18
        default:
            var v19 int16 = 1
            _ = v19
        }
        v3 -= (v3 << ((v4) & 63))
        var v20 bool = (128 == 32768)
        _ = v20
    }
    var v21 int16 = 616
    v8 += (v3 - v3)
    v22 := []int{0, 758}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v14
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
