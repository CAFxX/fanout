package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int64; F1 bool }
type GenS3 struct { F0 int8; F1 int }
type GenS4 struct { F0 int16 }
func genFn0(p0 string, p1 int) uint64 {
    v5 := GenS1{F0: 5}
    var v6 uint16 = 65535
    var v7 uint8 = 0
    fmt.Println(p1)
    v8 := GenS1{F0: 1}
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    return 18446744073709551615
}

func main() {
    var v1 uint = 2
    v2 := []int{9223372036854775807, 9223372036854775807, 4294967295, 9223372036854775807}
    var v3 bool = true
    var v4 uint16 = 2
    fmt.Println(v3, v1)
    for fi9 := 0; fi9 < 9; fi9++ {
        var v10 uint32 = 255
        v11 := GenS2{F0: 1, F1: v3}
        _ = fi9
        _ = v10
        _ = v11
    }
    if v3 {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    var v12 uint32 = 1
    var v13 uint32 = 518
    v14 := GenS3{F0: -128, F1: -2147483648}
    if (!false) {
        _ = len(v2)
        v15 := GenS2{F0: 443, F1: ("foo bar" >= ((("hello" + "hello") + "foo bar") + "a"))}
        var v16 uint16 = (v4 | (v4 + (v4 << ((255) & 63))))
        _ = v15
        _ = v16
    }
    switch ("foo bar" + "\n\t") {
    case "\n\t":
        _ = len(v2)
    default:
        v17 := GenS4{F0: 868}
        _ = v17
    }
    v18 := GenS3{F0: -128, F1: 792}
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v19 uint32 = (v12 & 4294967295)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
}
