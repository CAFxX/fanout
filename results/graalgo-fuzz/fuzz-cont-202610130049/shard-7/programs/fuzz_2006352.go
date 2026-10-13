package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 bool; F1 uint64; F2 uint32 }
func genFn0() int16 {
    var v7 uint16 = 32767
    v7 = 273
    switch "" {
    case "hello":
        v7 = (v7 << ((4294967295) & 63))
    case "":
        switch "a" {
        case "foo bar":
            v7 = (v7 - v7)
        case "hello":
            v8 := []int{-438, 511, 897, -4}
            _ = v8
        case "a":
            if (true || false) {
                v9 := []int{42, -482, 9223372036854775807, 1}
                _ = v9
            }
        default:
            v10 := GenS1{F0: ("a" + "")}
            _ = v10
        }
    }
    v11 := []int{32767, 128, 1, 1}
    var v12 int32 = 2147483647
    _ = v11
    _ = v12
    _ = v7
    return 1
}

func main() {
    var v1 bool = (1.5 == (0.0 + 1.5))
    v2 := []string{"a", "hello", "\n\t"}
    var v3 int16 = -32768
    var v4 string = "hello"
    v5 := make(map[uint32]int)
    v6 := []uint32{}
    v1 = (false && v1)
    for fi13 := 0; fi13 < 5; fi13++ {
        var v14 int8 = 0
        _ = fi13
        _ = v14
    }
    var v15 float32 = 0.1
    switch -217 {
    case 4294967295:
        if true {
            v16 := GenS2{F0: (v1 || v1), F1: 0, F2: 0}
            var v17 int64 = 26
            _ = v16
            _ = v17
        } else {
            var v18 int16 = (v3 ^ (v3 % ((v3) | 1)))
            var v19 int32 = 435
            _ = v18
            _ = v19
        }
    case 256:
        for fi20 := 0; fi20 < 18; fi20++ {
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
            v4 = v4
            _ = fi20
        }
    case 2147483647:
        var v21 uint32 = 2147483648
        _ = v21
    }
    fmt.Println(v15, v4)
    fmt.Println(v1)
    var v22 uint8 = 2
    _ = len(v2)
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v15
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
