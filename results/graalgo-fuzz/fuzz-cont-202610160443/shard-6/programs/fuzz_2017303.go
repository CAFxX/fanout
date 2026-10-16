package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16; F2 bool }

func main() {
    var v1 int64 = -9223372036854775808
    var v2 int16 = -32768
    v3 := make(map[uint32]int)
    var v4 uint32 = 65535
    switch 127 {
    case 1:
        _ = len(v3)
    case 50:
        if ((!(!false)) && ((v2 >> ((255) & 63)) >= v2)) {
            v5 := []int{-933, -1}
            _ = v5
        }
    default:
        fmt.Println(v1, v4, v4)
    }
    fmt.Println(v2)
    var v6 int64 = -791
    v1 -= -1
    v2 += 32767
    var v7 int32 = 0
    fmt.Println(v6, v1)
    fmt.Println(v6)
    v8 := GenS1{F0: -1, F1: 885, F2: (2147483648 <= v4)}
    fmt.Println(v2, v4)
    for fi9 := 0; fi9 < 18; fi9++ {
        fmt.Println(v6, v4)
        var v10 uint8 = 0
        _ = fi9
        _ = v10
    }
    fmt.Println(v7, v7, v7)
    v3[2147483647] = -795
    if false {
        var v11 bool = (2 == 190)
        var v12 uint32 = v4
        if true {
            fmt.Println(v6, v2, v7)
            v4 = 1
        }
        _ = v11
        _ = v12
    }
    switch -444 {
    case 128:
        v13 := GenS1{F0: -1, F1: 2, F2: false}
        _ = v13
    case 32768:
        for fi14 := 0; fi14 < 3; fi14++ {
            v15 := []int{(fi14 % ((128) | 1)), 32767}
            _ = len(v15)
            _ = fi14
            _ = v15
        }
    case 1:
        switch "\n\t" {
        case "a":
            if false {
                v2 -= 1000
                v4 = (v4 % ((v4) | 1))
                var v16 uint8 = 35
                _ = v16
            }
        case "foo bar":
            var v17 uint16 = 1
            _ = v17
        case "hello":
            v18 := []int{}
            _ = v18
        default:
            var v19 int32 = v7
            _ = v19
        }
    }
    var v20 int64 = -430
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v6
    _ = v7
    _ = v8
}
