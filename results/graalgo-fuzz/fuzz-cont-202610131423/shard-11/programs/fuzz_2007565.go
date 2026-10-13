package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int }
type GenS2 struct { F0 int32 }
func genFn0(p0 int16, p1 int16, p2 int16) int8 {
    var v4 int16 = 1000
    var v5 uint8 = 127
    var v6 int = 0
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    _ = v6
    return int8(v4)
}

func main() {
    v1 := []int{4294967295}
    var v2 uint16 = 32767
    var v3 string = ((("foo bar" + "\n\t") + "\n\t") + "hello")
    v2 = v2
    var v7 uint8 = 148
    for fi8 := 0; fi8 < 16; fi8++ {
        fmt.Println(v7)
        _ = fi8
    }
    v9 := GenS1{F0: -375, F1: 0}
    if false {
        v10 := GenS1{F0: 9223372036854775807, F1: 128}
        var v11 float32 = 1.5
        _ = v10
        _ = v11
    }
    v2 += 144
    v3 = v3
    for fi12 := 0; fi12 < 13; fi12++ {
        switch v3 {
        case "hello":
            var v13 int64 = -1
            _ = v13
        case "foo bar":
            v14 := GenS2{F0: -791}
            _ = v14
        case "\n\t":
            var v15 bool = true
            _ = v15
        default:
            v3 = v3
        }
        _ = fi12
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v7
    _ = v9
}
