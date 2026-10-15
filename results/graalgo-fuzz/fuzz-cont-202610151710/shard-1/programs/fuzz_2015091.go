package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string }
type GenS2 struct { F0 int32; F1 int32 }
type GenS3 struct { F0 int64 }

func main() {
    v1 := make(map[string]uint32)
    var v2 uint16 = 32767
    var v3 int32 = 2147483647
    v4 := GenS1{F0: -581, F1: ""}
    switch "foo bar" {
    case "\n\t":
        fmt.Println(v2)
    }
    v5 := GenS2{F0: v3, F1: v3}
    for fi6 := 0; fi6 < 13; fi6++ {
        switch fi6 {
        case 9223372036854775807:
            v7 := GenS3{F0: 42}
            _ = v7
        }
        v8 := []int{}
        for fi9 := 0; fi9 < 5; fi9++ {
            fmt.Println(fi6)
            if false {
                fmt.Println(v3, v2, v2)
                v2 = 1
            } else {
                v2 = v2
            }
            _ = fi9
        }
        _ = fi6
        _ = v8
    }
    v10 := []int{-520}
    fmt.Println(v3, v2)
    v2 += (v2 + v2)
    switch -587 {
    case 42:
        var v11 int16 = 32767
        _ = v11
    case 0:
        v2 -= v2
    case 887:
        for fi12 := 0; fi12 < 12; fi12++ {
            if len(v10) >= 2 {
                v10 = v10[1:len(v10)]
            }
            _ = fi12
        }
    }
    for fi13 := 0; fi13 < 20; fi13++ {
        v2 = v2
        v2 = 524
        _ = fi13
    }
    fmt.Println(v3, v2)
    v14 := GenS2{F0: v3, F1: v3}
    v15 := GenS3{F0: 1}
    var v16 bool = false
    var v17 string = "hello"
    v2 -= v2
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
