package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 bool; F1 uint16; F2 int32 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int8; F1 int }

func main() {
    v1 := make(map[int]bool)
    var v2 uint = 1
    v3 := make(map[string]bool)
    v4 := make(map[string]bool)
    v5 := []string{"a", "hello", ("hello" + "foo bar")}
    var v6 uint32 = 42
    for fi7 := 0; fi7 < 7; fi7++ {
        var v8 int64 = 311
        _ = len(v4)
        v6 = (v6 >> (((v2 / (((v2 - v2)) | 1))) & 63))
        _ = fi7
        _ = v8
    }
    var v9 string = ("a" + "a")
    v10 := GenS1{F0: true}
    v2 = (v2 & 0)
    for fi11 := 0; fi11 < 10; fi11++ {
        v12 := GenS2{F0: (779 <= 255), F1: 42, F2: -1}
        _ = len(v5)
        switch v9 {
        case "hello":
            fmt.Println(v6, fi11)
        case "foo bar":
            var v13 uint16 = 32767
            _ = v13
        }
        _ = fi11
        _ = v12
    }
    v9 = v9
    if true {
        if (127 <= 255) {
            delete(v1, -9223372036854775808)
            v14 := GenS3{F0: uint8(v6)}
            fmt.Println(v9)
            _ = v14
        } else {
            v2 = v2
        }
        if true {
            var v15 int8 = 114
            var v16 uint = 2
            _ = v15
            _ = v16
        } else {
            v2 += v2
            var v17 bool = true
            _ = v17
        }
    } else {
        for fi18 := 0; fi18 < 15; fi18++ {
            v2 += (v2 * 627)
            _ = fi18
        }
        var v19 uint16 = 32767
        _ = v19
    }
    var v20 int64 = 1
    var v21 float64 = -2.25
    v22 := GenS4{F0: 0, F1: -311}
    fmt.Println(v9, v21)
    v20 -= v20
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
