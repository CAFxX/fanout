package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8 }
type GenS2 struct { F0 int64; F1 uint; F2 int }
type GenS3 struct { F0 bool; F1 int8 }

func main() {
    var v1 int8 = 38
    var v2 uint64 = 485
    var v3 int = 42
    var v4 bool = true
    v5 := make(map[string]int)
    v6 := make(map[int]uint32)
    fmt.Println(v1, v3)
    var v7 float64 = 0.0
    for fi8 := 0; fi8 < 7; fi8++ {
        fmt.Println(v2)
        var v9 uint64 = (v2 % ((v2) | 1))
        _ = fi8
        _ = v9
    }
    var v10 int16 = 633
    v11 := GenS1{F0: (v10 ^ -1), F1: 207}
    fmt.Println(v7, v10, v7)
    switch "hello" {
    case "foo bar":
        v10 += (v10 & 46)
    case "\n\t":
        switch "foo bar" {
        case "":
            var v12 int = 127
            _ = v12
        case "foo bar":
            fmt.Println(v1, v10)
        case "\n\t":
            v13 := GenS1{F0: (v10 - (v10 | (v10 / ((v10) | 1)))), F1: 2}
            _ = v13
        }
    case "hello":
        v14 := GenS1{F0: 42, F1: 25}
        _ = v14
    default:
        v15 := GenS1{F0: (v10 | int16(v1)), F1: 1}
        _ = v15
    }
    v7 += v7
    if false {
        v3 = v3
        if (0 <= 719) {
            var v16 uint64 = v2
            if ((395 < v10) || (false && (42 <= (v10 >> ((42) & 63))))) {
                v2 -= v2
            }
            _ = v16
        } else {
            fmt.Println(v10, v2, v7)
            _ = len(v5)
            v4 = v4
        }
        v17 := []int{v3, v3, v3}
        _ = v17
    } else {
        v18 := []int{(v3 - v3)}
        if v4 {
            for fi19 := 0; fi19 < 11; fi19++ {
                if (v4 && (v4 || (v10 < v10))) {
                    var v20 uint8 = 2
                    _ = v20
                } else {
                    v21 := GenS2{F0: 9223372036854775807, F1: 346, F2: (fi19 - fi19)}
                    _ = v21
                }
                v22 := GenS1{F0: 42, F1: 42}
                _ = fi19
                _ = v22
            }
        } else {
            var v23 bool = true
            v24 := GenS3{F0: false, F1: v1}
            _ = v23
            _ = v24
        }
        _ = v18
    }
    fmt.Println(v2, v4, v10)
    for fi25 := 0; fi25 < 1; fi25++ {
        delete(v6, (^(v3)))
        _ = fi25
    }
    if v4 {
        if (2147483648 < 65535) {
            fmt.Println(v1, v2)
        }
    } else {
        if (!v4) {
            fmt.Println(v3, v2, v10)
            if v4 {
                v3 = (v3 >> ((255) & 63))
                _ = len(v6)
            }
            v26 := []int{9223372036854775807, (v3 ^ (v3 + (-(v3)))), v3}
            _ = v26
        } else {
            v10 = (v10 ^ (v10 & v10))
        }
    }
    delete(v6, v3)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
