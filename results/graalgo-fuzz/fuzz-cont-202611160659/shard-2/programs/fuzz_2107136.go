package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int; F2 int }
type GenS2 struct { F0 uint32; F1 uint32 }
type GenS3 struct { F0 uint8; F1 int8; F2 uint64 }
func genFn0() int {
    var v7 int8 = 36
    if true {
        v8 := []int{-1, 9223372036854775807}
        fmt.Println(v7)
        _ = v8
    }
    v9 := []int{1, -1, 32767}
    switch "\n\t" {
    case "hello":
        fmt.Println(v7)
    case "a":
        v10 := GenS1{F0: true, F1: 964, F2: 42}
        _ = v10
    case "foo bar":
        v7 -= (v7 >> ((641) & 63))
    }
    _ = v7
    _ = v9
    return 2147483647
}

func main() {
    v1 := make(map[int]string)
    var v2 int16 = -699
    var v3 string = "foo bar"
    var v4 uint32 = 2
    var v5 bool = (true || true)
    var v6 float32 = 1.5
    var v11 int = 4294967295
    v12 := []int{v11}
    for fi13 := 0; fi13 < 16; fi13++ {
        fmt.Println(v11)
        for fi14 := 0; fi14 < 18; fi14++ {
            v15 := GenS2{F0: (^(v4)), F1: 2147483648}
            _ = fi14
            _ = v15
        }
        v2 = v2
        _ = fi13
    }
    switch (^(v11)) {
    case 9223372036854775807:
        for fi16 := 0; fi16 < 9; fi16++ {
            var v17 string = v3
            _ = fi16
            _ = v17
        }
    default:
        var v18 int32 = 969
        _ = v18
    }
    v5 = (952 >= 609)
    v3 = "a"
    if (498 > 282) {
        fmt.Println(v3, v2)
    } else {
        v12 = append(v12, -1)
        v19 := GenS3{F0: 42, F1: 121, F2: 285}
        _ = v19
    }
    if len(v12) >= 2 {
        v12 = v12[1:len(v12)]
    }
    for fi20 := 0; fi20 < 9; fi20++ {
        v2 = v2
        if len(v12) >= 2 {
            v12 = v12[1:len(v12)]
        }
        var v21 string = v3
        _ = fi20
        _ = v21
    }
    fmt.Println(v4)
    for fi22 := 0; fi22 < 17; fi22++ {
        var v23 uint64 = 2
        var v24 uint16 = 122
        _ = fi22
        _ = v23
        _ = v24
    }
    for fi25 := 0; fi25 < 13; fi25++ {
        var v26 int8 = -128
        fmt.Println(v4, v2, v26)
        v27 := GenS1{F0: v5, F1: (v11 / (((-(fi25))) | 1)), F2: -1}
        _ = fi25
        _ = v26
        _ = v27
    }
    fmt.Println(v3, v5)
    fmt.Println(v6, v2, v2)
    for fi28 := 0; fi28 < 16; fi28++ {
        var v29 bool = (3.14159 <= -2.25)
        _ = fi28
        _ = v29
    }
    if (v5 && (0.1 <= 3.14159)) {
        var v30 int8 = 0
        _ = v30
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
