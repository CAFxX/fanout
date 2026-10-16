package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int32; F2 string }
type GenS2 struct { F0 int32; F1 uint64 }

func main() {
    v1 := make(map[int]bool)
    v2 := []int{29, -2147483648}
    var v3 uint = 65535
    v4 := make(map[uint32]int)
    if false {
        v5 := GenS1{F0: 1000, F1: -172, F2: "a"}
        _ = len(v2)
        var v6 int32 = 42
        _ = v5
        _ = v6
    } else {
        var v7 bool = false
        _ = v7
    }
    switch ("" + "\n\t") {
    case "foo bar":
        var v8 uint8 = 42
        _ = v8
    case "hello":
        fmt.Println(v3)
    default:
        switch 32767 {
        case 9223372036854775807:
            v3 += (v3 * (v3 & 1))
        }
    }
    v9 := GenS2{F0: 0, F1: 767}
    fmt.Println(v3)
    v10 := GenS1{F0: 1, F1: -1, F2: ""}
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    for fi11 := 0; fi11 < 5; fi11++ {
        v3 += v3
        var v12 float32 = (1.5 / -2.25)
        v3 = (v3 << (((^(v3))) & 63))
        _ = fi11
        _ = v12
    }
    fmt.Println(v3)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
