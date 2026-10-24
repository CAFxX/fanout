package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int; F1 uint16; F2 uint32 }
type GenS3 struct { F0 bool; F1 int32 }

func main() {
    v1 := make(map[uint32]bool)
    var v2 int32 = 42
    v3 := make(map[int]int)
    if false {
        v2 = (v2 + v2)
        v1[0] = (100 <= -75)
        switch ("a" + (("hello" + "hello") + "hello")) {
        case "hello":
            var v4 float64 = -2.25
            _ = v4
        case "\n\t":
            v2 = (v2 >> ((4294967295) & 63))
        case "a":
            if true {
                var v5 int16 = 42
                fmt.Println(v2, v2)
                _ = v5
            }
        default:
            if (255 != 42) {
                v2 = v2
                v2 = (v2 * (v2 << ((1) & 63)))
                v6 := GenS1{F0: 0}
                _ = v6
            }
        }
    }
    switch 4294967295 {
    case -9223372036854775808:
        v7 := GenS2{F0: 0, F1: 311, F2: 1}
        _ = v7
    case 2147483647:
        var v8 uint8 = 45
        _ = v8
    case -322:
        fmt.Println(v2)
    }
    var v9 uint16 = 0
    if (false || true) {
        for fi10 := 0; fi10 < 8; fi10++ {
            v11 := GenS2{F0: 9223372036854775807, F1: v9, F2: 65535}
            _ = fi10
            _ = v11
        }
        var v12 float32 = (1e10 / -2.25)
        for fi13 := 0; fi13 < 3; fi13++ {
            fmt.Println(v9, fi13)
            _ = fi13
        }
        _ = v12
    } else {
        v14 := []int{42, 42}
        _ = v14
    }
    var v15 float64 = float64(0.0)
    var v16 bool = (9223372036854775807 < 42)
    fmt.Println(v15, v2)
    v17 := GenS3{F0: (v2 >= v2), F1: 0}
    v2 = -2147483648
    v16 = (v16 && (v16 || (1 > -1)))
    v16 = (101 > 100)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v9
}
