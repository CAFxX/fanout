package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint8; F1 int }
func genFn0(p0 int64) int16 {
    if false {
        for fi5 := 0; fi5 < 4; fi5++ {
            fmt.Println(p0)
            _ = fi5
        }
        v6 := GenS1{F0: 815}
        _ = v6
    } else {
        switch "\n\t" {
        case "a":
            for fi7 := 0; fi7 < 13; fi7++ {
                fmt.Println(p0)
                _ = fi7
            }
        }
    }
    v8 := GenS1{F0: 42}
    _ = p0
    _ = v8
    return 1
}
func genFn1(p0 int8) bool {
    switch ("" + "hello") {
    case "hello":
        v9 := GenS2{F0: ""}
        _ = v9
    }
    for fi10 := 0; fi10 < 18; fi10++ {
        fmt.Println(fi10, p0)
        _ = fi10
    }
    var v11 int = 42
    _ = p0
    _ = v11
    return (255 >= 1)
}

func main() {
    var v1 int32 = -226
    var v2 int64 = 9223372036854775807
    v3 := []int{}
    v4 := []int{1, 127, 2147483647}
    v1 -= v1
    v1 -= v1
    var v12 uint32 = 746
    v4 = append(v4, 2147483647)
    _ = len(v3)
    fmt.Println(v1, v12, v12)
    v13 := GenS2{F0: "\n\t"}
    var v14 uint = 0
    v2 -= (v2 >> (((v14 * v14)) & 63))
    fmt.Println(v12)
    for fi15 := 0; fi15 < 17; fi15++ {
        v2 += (v2 * 21)
        _ = fi15
    }
    v3 = append(v3, 1)
    v16 := GenS3{F0: 127, F1: 32767}
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v16
    _ = v2
    _ = v3
    _ = v4
}
